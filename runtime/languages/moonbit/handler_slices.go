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
	"errors"
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

	if ok := utils.CopyMemory(wa.Memory(), ptr, offset, 12); !ok {
		return cln, errors.New("failed to copy slice header")
	}

	return cln, nil
}

func (h *sliceHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	gmlPrintf("GML: handler_slices.go: sliceHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value when decoding a slice but got %v: %+v", len(vals), vals)
	}

	// note: capacity is not used here
	// data, size := uint32(vals[0]), uint32(vals[1])

	memBlock, _, err := memoryBlockAtOffset(wa, uint32(vals[0]), true)
	if err != nil {
		return nil, err
	}

	sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
	if sliceOffset == 0 {
		return nil, nil // nil slice
	}
	numElements := binary.LittleEndian.Uint32(memBlock[12:16])
	if numElements == 0 {
		return h.emptyValue, nil // empty slice
	}
	gmlPrintf("GML: handler_slices.go: sliceHandler.Decode: sliceOffset=%v, numElements=%v", debugShowOffset(sliceOffset), numElements)

	memBlock, _, err = memoryBlockAtOffset(wa, sliceOffset, true)
	if err != nil {
		return nil, err
	}

	// return h.doReadSlice(ctx, wa, sliceOffset, size)

	// elementSize := h.elementHandler.TypeInfo().Size()
	items := reflect.MakeSlice(h.typeInfo.ReflectedType(), int(numElements), int(numElements))
	for i := uint32(0); i < numElements; i++ {
		// itemOffset := data + i*elementSize
		itemOffset := binary.LittleEndian.Uint32(memBlock[8+i*4:])
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
	size := numElements * 4
	// headerValue = (numElements << 8) | 241 // 241 is the int array header type
	memBlockClassID := uint32(ArrayBlockType)

	//  ptr, cln, err = wa.(*wasmAdapter).makeWasmObject(ctx, h.typeDef.Id, uint32(len(slice)))
	ptr, cln, err = wa.(*wasmAdapter).allocateAndPinMemory(ctx, size, memBlockClassID)
	if err != nil {
		return 0, cln, err
	}

	innerCln := utils.NewCleanerN(len(slice))

	defer func() {
		// unpin slice elements after the slice is written to memory
		if e := innerCln.Clean(); e != nil && err == nil {
			err = e
		}
	}()

	elementSize := h.elementHandler.TypeInfo().Size()

	for i, val := range slice {
		if !utils.HasNil(val) {
			c, err := h.elementHandler.Write(ctx, wa, ptr+uint32(i)*elementSize, val)
			innerCln.AddCleaner(c)
			if err != nil {
				return 0, cln, err
			}
		}
	}

	return ptr, cln, nil
}
