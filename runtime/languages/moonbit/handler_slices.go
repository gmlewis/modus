/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package moonbit

import (
	"context"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
)

func (p *planner) NewSliceHandler(ctx context.Context, ti langsupport.TypeInfo) (langsupport.TypeHandler, error) {
	handler := &sliceHandler{
		typeHandler: *NewTypeHandler(ti),
	}
	p.AddHandler(handler)

	typeDef, err := p.metadata.GetTypeDefinition(ti.Name())
	if err != nil {
		return nil, fmt.Errorf("planner.NewSliceHandler: p.metadata.GetTypeDefinition('%v'): %w", ti.Name(), err)
	}
	handler.typeDef = typeDef

	elementHandler, err := p.GetHandler(ctx, ti.ListElementType().Name())
	if err != nil {
		return nil, fmt.Errorf("planner.NewSliceHandler: p.GetHandler('%v'): %w", ti.ListElementType().Name(), err)
	}
	handler.elementHandler = elementHandler

	// an empty slice (not nil)
	handler.emptyValue = reflect.MakeSlice(ti.ReflectedType(), 0, 0).Interface()

	return handler, nil
}

type sliceHandler struct {
	typeHandler
	typeDef        *metadata.TypeDefinition
	elementHandler langsupport.TypeHandler
	emptyValue     any
}

func (h *sliceHandler) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	return h.Decode(ctx, wa, []uint64{uint64(offset)})
}

func (h *sliceHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	ptr, cln, err := h.doWriteSlice(ctx, wa, obj)
	if err != nil {
		return cln, err
	}

	wa.Memory().WriteUint32Le(offset, ptr)

	return cln, nil
}

// Decode is always passed an address to a slice in memory.
func (h *sliceHandler) Decode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, vals []uint64) (any, error) {
	wa, ok := wasmAdapter.(wasmMemoryReader)
	if !ok {
		return nil, fmt.Errorf("expected a wasmMemoryReader, got %T", wasmAdapter)
	}

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value when decoding a slice but got %v: %+v", len(vals), vals)
	}

	if vals[0] == 0 {
		return nil, nil
	}

	memBlock, classID, words, err := memoryBlockAtOffset(wa, uint32(vals[0]), 0)
	if err != nil {
		return nil, err
	}

	if words == 0 {
		return h.emptyValue, nil // empty slice
	}

	numElements := uint32(words)
	elemType := h.typeInfo.ListElementType()
	elemTypeSize := uint32(4) // elemType.Size()
	isNullable := elemType.IsNullable()
	if isNullable && elemType.IsPrimitive() &&
		(elemType.Name() == "Int?" || elemType.Name() == "UInt?" || elemType.Name() == "String?") { // TODO: "String?" is not a "primitive" type, probably can be removed.
		elemTypeSize = 8
	}
	if classID == FixedArrayPrimitiveBlockType || classID == PtrArrayBlockType {
		memBlock, _, _, err = memoryBlockAtOffset(wa, uint32(vals[0]), words*elemTypeSize)
		if err != nil {
			return nil, err
		}
	} else {
		sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
		if sliceOffset == 0 {
			return nil, nil // nil slice
		}

		if words == 1 {
			// sliceOffset is the pointer to the single-slice element.
			items := reflect.MakeSlice(h.typeInfo.ReflectedType(), 1, 1)
			item, err := h.elementHandler.Read(ctx, wasmAdapter, sliceOffset)
			if err != nil {
				return nil, err
			}
			if !utils.HasNil(item) {
				items.Index(0).Set(reflect.ValueOf(item))
			}
			return items.Interface(), nil
		}

		numElements = binary.LittleEndian.Uint32(memBlock[12:16])
		if numElements == 0 {
			return h.emptyValue, nil // empty slice
		}

		size := numElements * uint32(elemTypeSize)

		memBlock, _, _, err = memoryBlockAtOffset(wa, sliceOffset, size)
		if err != nil {
			return nil, err
		}
	}

	items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
	for i := uint32(0); i < numElements; i++ {
		// TODO: This is all quite a hack - figure out how to make this an elegant solution.
		if elemType.IsPrimitive() && isNullable {
			var value uint64
			if elemType.Name() == "Int?" || elemType.Name() == "UInt?" || elemType.Name() == "String?" {
				value = binary.LittleEndian.Uint64(memBlock[8+i*elemTypeSize:])
			} else {
				value32 := binary.LittleEndian.Uint32(memBlock[8+i*elemTypeSize:])
				value = uint64(value32)
			}
			item, err := h.elementHandler.Decode(ctx, wasmAdapter, []uint64{value})
			if err != nil {
				return nil, err
			}
			if !utils.HasNil(item) {
				items.Index(int(i)).Set(reflect.ValueOf(item))
			}
			continue
		}
		ptr := binary.LittleEndian.Uint32(memBlock[8+i*elemTypeSize:])
		item, err := h.elementHandler.Decode(ctx, wasmAdapter, []uint64{uint64(ptr)})
		if err != nil {
			return nil, err
		}
		if !utils.HasNil(item) {
			items.Index(int(i)).Set(reflect.ValueOf(item))
		}
	}

	return items.Interface(), nil
}

func (h *sliceHandler) Encode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	ptr, cln, err := h.doWriteSlice(ctx, wasmAdapter, obj)
	if err != nil {
		return nil, cln, err
	}

	return []uint64{uint64(ptr)}, cln, nil
}

func (h *sliceHandler) doWriteSlice(ctx context.Context, wasmAdapter langsupport.WasmAdapter, obj any) (ptr uint32, cln utils.Cleaner, err error) {
	wa, ok := wasmAdapter.(wasmMemoryWriter)
	if !ok {
		return 0, nil, fmt.Errorf("expected a wasmMemoryWriter, got %T", wasmAdapter)
	}

	if utils.HasNil(obj) {
		return 0, nil, nil
	}

	slice, err := utils.ConvertToSlice(obj)
	if err != nil {
		return 0, nil, err
	}

	numElements := uint32(len(slice))
	elemType := h.typeInfo.ListElementType()
	elemTypeSize := uint32(4)
	isNullable := elemType.IsNullable()
	if elemType.IsPrimitive() && isNullable &&
		(elemType.Name() == "Int?" || elemType.Name() == "UInt?") {
		elemTypeSize = 8
	}
	size := numElements * uint32(elemTypeSize)
	memBlockClassID := uint32(PtrArrayBlockType)
	if elemType.Name() == "Byte?" || elemType.Name() == "Bool?" || elemType.Name() == "Char?" ||
		elemType.Name() == "Int?" || elemType.Name() == "UInt?" ||
		elemType.Name() == "Int16?" || elemType.Name() == "UInt16?" {
		memBlockClassID = uint32(FixedArrayPrimitiveBlockType)
	}

	// Allocate memory
	if size == 0 {
		ptr, cln, err = wa.allocateAndPinMemory(ctx, 1, memBlockClassID) // cannot allocate 0 bytes
		if err != nil {
			return 0, cln, err
		}
		wa.Memory().WriteByte(ptr-3, 0) // overwrite size=1 to size=0
	} else {
		ptr, cln, err = wa.allocateAndPinMemory(ctx, size, memBlockClassID)
		if err != nil {
			return 0, cln, err
		}

		// For `Int?`, `UInt?`, the `words` portion of the memory block
		// indicates the number of elements in the slice, not the number of 16-bit words.
		if elemType.Name() == "Int?" || elemType.Name() == "UInt?" {
			memType := ((size / 8) << 8) | memBlockClassID
			wa.Memory().WriteUint32Le(ptr-4, memType)
		}
	}

	innerCln := utils.NewCleanerN(len(slice))

	defer func() {
		// unpin slice elements after the slice is written to memory
		if e := innerCln.Clean(); e != nil && err == nil {
			err = e
		}
	}()

	for i, val := range slice {
		c, err := h.elementHandler.Write(ctx, wasmAdapter, ptr+uint32(i)*elemTypeSize, val)
		innerCln.AddCleaner(c)
		if err != nil {
			return 0, cln, err
		}
	}

	if strings.HasPrefix(h.typeDef.Name, "FixedArray[") {
		return ptr - 8, cln, nil
	}

	// Finally, write the slice memory block.
	slicePtr, sliceCln, err := wa.allocateAndPinMemory(ctx, 8, TupleBlockType)
	innerCln.AddCleaner(sliceCln)
	if err != nil {
		return 0, cln, err
	}
	wa.Memory().WriteUint32Le(slicePtr, ptr-8)
	wa.Memory().WriteUint32Le(slicePtr+4, numElements)

	return slicePtr - 8, cln, nil
}
