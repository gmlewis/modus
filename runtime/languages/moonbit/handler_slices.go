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
	gmlPrintf("GML: handler_slices.go: sliceHandler.Read(offset: %v), type=%T", debugShowOffset(offset), h.emptyValue)
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

func (h *sliceHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	gmlPrintf("GML: handler_slices.go: sliceHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value when decoding a slice but got %v: %+v", len(vals), vals)
	}

	memBlock, _, err := memoryBlockAtOffset(wa, uint32(vals[0]), 0, true)
	if err != nil {
		return nil, err
	}

	part2 := binary.LittleEndian.Uint32(memBlock[4:8])
	classID := part2 & 0xff
	words := part2 >> 8
	if words == 0 {
		return h.emptyValue, nil // empty slice
	}

	numElements := uint32(words)
	elemType := h.typeInfo.ListElementType()
	elemTypeSize := uint32(4) // elemType.Size()
	isNullable := elemType.IsNullable()
	if isNullable {
		elemTypeSize = 8
	}
	if classID == ArrayBlockType && elemType.IsPrimitive() && isNullable {
		memBlock, _, err = memoryBlockAtOffset(wa, uint32(vals[0]), words*elemTypeSize, true)
		if err != nil {
			return nil, err
		}
	} else {
		sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
		if sliceOffset == 0 {
			return nil, nil // nil slice
		}

		// For debugging:
		_, _, _ = memoryBlockAtOffset(wa, sliceOffset, 0, true)

		numElements = binary.LittleEndian.Uint32(memBlock[12:16])
		if numElements == 0 {
			return h.emptyValue, nil // empty slice
		}

		size := numElements * uint32(elemTypeSize)
		gmlPrintf("GML: handler_slices.go: sliceHandler.Decode: sliceOffset=%v, numElements=%v, size=%v", debugShowOffset(sliceOffset), numElements, size)

		memBlock, _, err = memoryBlockAtOffset(wa, sliceOffset, size, true)
		if err != nil {
			return nil, err
		}
	}

	// return h.doReadSlice(ctx, wa, sliceOffset, size)

	// elementSize := h.elementHandler.TypeInfo().Size()
	items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
	for i := uint32(0); i < numElements; i++ {
		if elemType.IsPrimitive() && isNullable {
			var value uint64
			if words > 1 {
				isNone := binary.LittleEndian.Uint32(memBlock[12+i*elemTypeSize:]) != 0
				if isNone {
					// items.Index(int(i)).Set(reflect.Zero(h.elementHandler.TypeInfo().ReflectedType()))
					continue
				}
				value = binary.LittleEndian.Uint64(memBlock[8+i*elemTypeSize:])
			} else {
				value32 := binary.LittleEndian.Uint32(memBlock[8+i*elemTypeSize:])
				value = uint64(value32)
			}
			item, err := h.elementHandler.Decode(ctx, wa, []uint64{value})
			if err != nil {
				return nil, err
			}
			if !utils.HasNil(item) {
				items.Index(int(i)).Set(reflect.ValueOf(item))
			}
			continue
		}
		itemOffset := binary.LittleEndian.Uint32(memBlock[8+i*elemTypeSize:])
		item, err := h.elementHandler.Read(ctx, wa, itemOffset)
		if err != nil {
			return nil, err
		}
		if !utils.HasNil(item) {
			items.Index(int(i)).Set(reflect.ValueOf(item))
		}
	}

	return items.Interface(), nil
}

func (h *sliceHandler) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	ptr, cln, err := h.doWriteSlice(ctx, wa, obj)
	if err != nil {
		return nil, cln, err
	}

	return []uint64{uint64(ptr)}, cln, nil
}

/*
func (h *sliceHandler) doReadSlice(ctx context.Context, wa langsupport.WasmAdapter, data, size uint32) (any, error) {
	if data == 0 {
		// nil slice
		return nil, nil
	}

	if size == 0 {
		// empty slice
		return h.emptyValue, nil
	}

	elementSize := h.elementHandler.TypeInfo().Size()
	items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(size), int(size))
	for i := uint32(0); i < size; i++ {
		itemOffset := data + i*elementSize
		item, err := h.elementHandler.Read(ctx, wa, itemOffset)
		if err != nil {
			return nil, err
		}
		if !utils.HasNil(item) {
			items.Index(int(i)).Set(reflect.ValueOf(item))
		}
	}

	return items.Interface(), nil
}
*/

// type sliceWriter interface {
// 	doWriteSlice(ctx context.Context, wa langsupport.WasmAdapter, obj any) (ptr uint32, cln utils.Cleaner, err error)
// }

func (h *sliceHandler) doWriteSlice(ctx context.Context, wa langsupport.WasmAdapter, obj any) (ptr uint32, cln utils.Cleaner, err error) {
	gmlPrintf("GML: handler_slices.go: sliceHandler.doWriteSlice(obj: %+v)", obj)
	if utils.HasNil(obj) {
		return 0, nil, nil
	}

	slice, err := utils.ConvertToSlice(obj)
	if err != nil {
		return 0, nil, err
	}

	numElements := uint32(len(slice))
	elemType := h.typeInfo.ListElementType()
	elemTypeSize := uint32(4) // elemType.Size()
	isNullable := elemType.IsNullable()
	if elemType.IsPrimitive() && isNullable {
		elemTypeSize = 8
	}
	size := numElements * uint32(elemTypeSize)
	// headerValue = (numElements << 8) | 241 // 241 is the int array header type
	memBlockClassID := uint32(ArrayBlockType)

	//  ptr, cln, err = wa.(*wasmAdapter).makeWasmObject(ctx, h.typeDef.Id, uint32(len(slice)))
	ptr, cln, err = wa.(*wasmAdapter).allocateAndPinMemory(ctx, size, memBlockClassID)
	if err != nil {
		return 0, cln, err
	}

	// For debugging purposes:
	_, _, _ = memoryBlockAtOffset(wa, ptr-8, 0, true)

	innerCln := utils.NewCleanerN(len(slice))

	defer func() {
		// unpin slice elements after the slice is written to memory
		if e := innerCln.Clean(); e != nil && err == nil {
			err = e
		}
	}()

	elementSize := h.elementHandler.TypeInfo().Size()

	for i, val := range slice {
		if elemType.IsPrimitive() && isNullable {
			if !utils.HasNil(val) {
				if _, err := h.elementHandler.Write(ctx, wa, ptr+uint32(i)*elemTypeSize, val); err != nil {
					return 0, cln, err
				}
				wa.(*wasmAdapter).Memory().Write(ptr+4+uint32(i)*elemTypeSize, []byte{0, 0, 0, 0})
			} else {
				wa.(*wasmAdapter).Memory().Write(ptr+uint32(i)*elemTypeSize, []byte{0, 0, 0, 0, 1, 0, 0, 0}) // None
			}
			continue
		}
		if !utils.HasNil(val) {
			c, err := h.elementHandler.Write(ctx, wa, ptr+uint32(i)*elementSize, val)
			innerCln.AddCleaner(c)
			if err != nil {
				return 0, cln, err
			}
		}
	}

	// For debugging purposes:
	_, _, _ = memoryBlockAtOffset(wa, ptr-8, 0, true)

	// Finally, write the slice memory bock.
	slicePtr, sliceCln, err := wa.(*wasmAdapter).allocateAndPinMemory(ctx, 8, 0)
	innerCln.AddCleaner(sliceCln)
	if err != nil {
		return 0, cln, err
	}
	wa.(*wasmAdapter).Memory().WriteUint32Le(slicePtr, ptr-8)
	wa.(*wasmAdapter).Memory().WriteUint32Le(slicePtr+4, numElements)

	// For debugging purposes:
	_, _, _ = memoryBlockAtOffset(wa, slicePtr-8, 0, true)

	return slicePtr - 8, cln, nil
}
