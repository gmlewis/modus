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
	"strings"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
)

func (p *planner) NewStructHandler(ctx context.Context, ti langsupport.TypeInfo) (langsupport.TypeHandler, error) {
	handler := &structHandler{
		typeHandler: *NewTypeHandler(ti),
		seenPtrs:    map[uintptr]uint32{},
		seenOffsets: map[uint32]any{},
	}
	p.AddHandler(handler)

	typeDef, err := p.metadata.GetTypeDefinition(ti.Name())
	if err != nil {
		return nil, err
	}
	handler.typeDef = typeDef

	fieldTypes := ti.ObjectFieldTypes()
	fieldHandlers := make([]langsupport.TypeHandler, len(fieldTypes))
	for i, fieldType := range fieldTypes {
		fieldHandler, err := p.GetHandler(ctx, fieldType.Name())
		if err != nil {
			return nil, err
		}
		fieldHandlers[i] = fieldHandler
	}

	handler.fieldHandlers = fieldHandlers
	return handler, nil
}

type structHandler struct {
	typeHandler
	typeDef       *metadata.TypeDefinition
	fieldHandlers []langsupport.TypeHandler

	// seen pointers prevents infinite recursion when encoding structs or maps (representing structs).
	seenPtrs map[uintptr]uint32
	// seen offsets prevents infinite recursion when decoding structs or maps (representing structs).
	seenOffsets map[uint32]any
}

func (h *structHandler) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {

	if v, ok := h.seenOffsets[offset]; ok {
		return v, nil
	}

	rt := h.typeInfo.ReflectedType()
	rv := reflect.New(rt) // only used to build a struct
	// `m` is used both for maps and to build a struct
	m := make(map[string]any, len(h.fieldHandlers))
	objPtr := rv.Interface()
	switch objPtr.(type) {
	case *map[string]any:
		h.seenOffsets[offset] = m
	default:
		h.seenOffsets[offset] = objPtr
	}
	defer func() {
		delete(h.seenOffsets, offset)
	}()

	memBlock, _, _, err := memoryBlockAtOffset(wa, offset, 0)
	if err != nil {
		return nil, fmt.Errorf("structHandler failed to read memory block at offset %v: %w", debugShowOffset(offset), err)
	}

	// TODO: fieldOffsets are unreliable because in the MoonBit metadata, the option
	// types do not have the full underlying struct info. When a Go map is traversed
	// to get the struct info, the option type may be processed before the underlying
	// type which causes zeros to be returned for size and alignment.
	// To fix this, the type cache would need to be always created by processing
	// the underlying types first.
	// fieldOffsets := h.typeInfo.ObjectFieldOffsets()

	recursionOnFields := map[int]any{}

	var fieldOffset uint32
	for i, field := range h.typeDef.Fields {
		handler := h.fieldHandlers[i]
		fieldName := strings.TrimPrefix(field.Name, "mut ")

		elementSize := handler.TypeInfo().Size()
		elementTypeName := handler.TypeInfo().Name()
		if elementTypeName == "Int?" || elementTypeName == "UInt?" {
			elementSize = 8
		} else if elementSize < 4 {
			elementSize = 4
		}

		var ptr uint64
		switch elementSize {
		case 8:
			ptr = binary.LittleEndian.Uint64(memBlock[8+fieldOffset:])
		default:
			// case 4:
			// Go primitive types (e.g. "uint16", "int16") return a size of 2, but all MoonBit sizes in structs
			// are either 4 or 8, so make the default be case 4.
			ptr = uint64(binary.LittleEndian.Uint32(memBlock[8+fieldOffset:]))
		}

		v, ok := h.seenOffsets[uint32(ptr)]
		if ok && handler.TypeInfo().IsPointer() {
			recursionOnFields[i] = v
			if rt.Kind() == reflect.Map {
				m[fieldName] = v
			} else {
				m[fieldName] = nil
			}
		} else {
			val, err := handler.Decode(ctx, wa, []uint64{ptr})
			if err != nil {
				return nil, err
			}
			m[fieldName] = val
		}

		// switch handler.TypeInfo().Size() {
		// case 8:
		// 	fieldOffset += 8
		// default:
		// 	// case 4:
		// 	// Go primitive types (e.g. "uint16", "int16") return a size of 2, but all MoonBit sizes in structs
		// 	// are either 4 or 8, so make the default be case 4.
		// 	fieldOffset += 4
		// }
		fieldOffset += elementSize
	}

	// Handle tuple output
	if strings.HasPrefix(h.typeInfo.Name(), "(") {
		result := make([]any, 0, len(m))
		for i := 0; i < len(m); i++ {
			result = append(result, m[fmt.Sprintf("%v", i)])
		}
		return result, nil
	}

	// Handle map output
	if rt.Kind() == reflect.Map {
		return m, nil
	}

	// Handle struct output
	result, err := h.getStructOutput(rv, m, recursionOnFields)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (h *structHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	return h.Read(ctx, wa, uint32(vals[0]))
}

func (h *structHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	ptr, cln, err := h.Encode(ctx, wa, obj)
	if err != nil {
		return cln, err
	}

	if ok := wa.Memory().WriteUint32Le(offset, uint32(ptr[0])); !ok {
		return cln, errors.New("failed to write struct pointer to memory")
	}

	return cln, nil
}

func (h *structHandler) Encode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	wa, ok := wasmAdapter.(wasmMemoryWriter)
	if !ok {
		return nil, nil, fmt.Errorf("expected a wasmMemoryWriter, got %T", wasmAdapter)
	}

	var mapObj map[string]any
	var rvObj reflect.Value
	var thisObjPtr uintptr // populated for maps or pointers to structs
	if m, ok := obj.(*map[string]any); ok {
		obj = utils.DereferencePointer(m)
	}

	if m, ok := obj.(map[string]any); ok {
		mapObj = m
		thisObjPtr = reflect.ValueOf(obj).Pointer()
	} else {
		rvObj = reflect.ValueOf(obj)
		if rvObj.Kind() == reflect.Ptr {
			thisObjPtr = reflect.ValueOf(obj).Pointer()
			// Dereference the pointer but leave `obj` pointing to the original object to
			// detect and handle self-referencing fields.
			structObj := utils.DereferencePointer(obj)
			rvObj = reflect.ValueOf(structObj)
		}
		if rvObj.Kind() != reflect.Struct {
			return nil, nil, fmt.Errorf("expected a struct, got %s", rvObj.Kind())
		}
	}

	if v, ok := h.seenPtrs[thisObjPtr]; ok {
		return []uint64{uint64(v)}, nil, nil
	}

	numFields := len(h.typeDef.Fields)
	cleaner := utils.NewCleanerN(numFields)

	totalSize := uint32(0) // h.typeInfo.Size()
	// NOTE: Due to the fact that the Go AST is being used to represent MoonBit types
	// and due to MoonBit primitives not taking the same size as Go primitives, the
	// totalSize is not correct.  Therefore, we need to calculate the totalSize based
	// on the individual field sizes.
	fieldOffsets := make([]uint32, 0, numFields)
	for i := range h.typeDef.Fields {
		handler := h.fieldHandlers[i]
		elementSize := handler.TypeInfo().Size()
		elementTypeName := handler.TypeInfo().Name()
		if elementTypeName == "Int?" || elementTypeName == "UInt?" {
			elementSize = 8
		} else if elementSize < 4 {
			elementSize = 4
		}
		fieldOffsets = append(fieldOffsets, totalSize)
		totalSize += elementSize
	}

	offset, cln, err := wa.allocateAndPinMemory(ctx, totalSize, TupleBlockType)
	if err != nil {
		return nil, cln, err
	}
	cleaner.AddCleaner(cln)

	if thisObjPtr != 0 {
		h.seenPtrs[thisObjPtr] = offset - 8
		defer func() {
			delete(h.seenPtrs, thisObjPtr)
		}()
	}

	for i, field := range h.typeDef.Fields {
		var fieldObj any
		fieldName := strings.TrimPrefix(field.Name, "mut ")
		if mapObj != nil {
			// case sensitive when reading from map
			fieldObj = mapObj[fieldName]
		} else {
			fieldObj = rvObj.FieldByIndex([]int{i}).Interface()
		}

		// In MoonBit, 'None's/nils need to be written, so don't skip them.

		handler := h.fieldHandlers[i]
		cln, err := handler.Write(ctx, wasmAdapter, offset+fieldOffsets[i], fieldObj)
		cleaner.AddCleaner(cln)
		if err != nil {
			return nil, cleaner, err
		}
	}

	return []uint64{uint64(offset - 8)}, cleaner, nil
}

func (h *structHandler) getStructOutput(rv reflect.Value, data map[string]any, recursionOnFields map[int]any) (any, error) {
	if err := utils.MapToStruct(data, rv.Interface()); err != nil {
		return nil, err
	}

	for index, v := range recursionOnFields {
		fieldObj := rv.Elem().FieldByIndex([]int{index})
		fieldObj.Set(reflect.ValueOf(v))
	}

	return rv.Elem().Interface(), nil
}
