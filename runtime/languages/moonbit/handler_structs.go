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
	"log"
	"reflect"
	"strings"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
)

// const maxDepth = 5 // TODO: make this based on the depth requested in the query

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
	gmlPrintf("GML: handler_structs.go: structHandler.Read(offset: %v)", debugShowOffset(offset))

	if v, ok := h.seenOffsets[offset]; ok {
		gmlPrintf("GML: structHandler.Read: RECURSION DETECTED for offset=%v", offset)
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
	log.Printf("GML: structHandler.Read: h.seenOffsets[%v]=%v=0x%[2]x", offset, reflect.ValueOf(h.seenOffsets[offset]).Pointer())
	defer func() {
		delete(h.seenOffsets, offset)
	}()

	// // Check for recursion
	// visitedPtrs := wa.(*wasmAdapter).visitedPtrs
	// if visitedPtrs[offset] >= maxDepth {
	// 	logger.Warn(ctx).Bool("user_visible", true).Msgf("Excessive recursion detected in %s. Stopping at depth %d.", h.typeInfo.Name(), maxDepth)
	// 	return nil, nil
	// }
	// visitedPtrs[offset]++
	// defer func() {
	// 	n := visitedPtrs[offset]
	// 	if n == 1 {
	// 		delete(visitedPtrs, offset)
	// 	} else {
	// 		visitedPtrs[offset] = n - 1
	// 	}
	// }()

	memBlock, _, _, err := memoryBlockAtOffset(wa, offset, 0, true)
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
			gmlPrintf("GML: handler_structs.go: structHandler.Read: fieldName: '%v', type: '%v', fieldOffset: %v, uint64 ptr: %v", fieldName, handler.TypeInfo().Name(), fieldOffset, ptr)
		default:
			// gmlPrintf("GML: handler_structs.go: structHandler.Read: fieldName: '%v', type: '%v', fieldOffset: %v", fieldName, handler.TypeInfo().Name(), fieldOffset)
			// return nil, fmt.Errorf("unsupported size for type '%v': %v", handler.TypeInfo().Name(), handler.TypeInfo().Size())
			// case 4:
			// Go primitive types (e.g. "uint16", "int16") return a size of 2, but all MoonBit sizes in structs
			// are either 4 or 8, so make the default be case 4.
			ptr = uint64(binary.LittleEndian.Uint32(memBlock[8+fieldOffset:]))
			gmlPrintf("GML: handler_structs.go: structHandler.Read: fieldName: '%v', type: '%v', fieldOffset: %v, uint32 ptr: %v", fieldName, handler.TypeInfo().Name(), fieldOffset, debugShowOffset(uint32(ptr)))
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
		log.Printf("GML: structHandler.Read: rt.Kind() == reflect.Map: m=%v=0x%[1]x", reflect.ValueOf(m).Pointer())
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
		gmlPrintf("GML: handler_structs.go: structHandler.Encode: obj=%v=0x%[1]x, m=%[2]v=0x%[2]x, mapObj=%[3]v=0x%[3]x", reflect.ValueOf(obj).Pointer(), reflect.ValueOf(m).Pointer(), reflect.ValueOf(mapObj).Pointer())
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
		gmlPrintf("GML: structHandler: RECURSION DETECTED for uintptr=%v: returning %v", thisObjPtr, v)
		return []uint64{uint64(v)}, nil, nil
	}

	numFields := len(h.typeDef.Fields)
	cleaner := utils.NewCleanerN(numFields)

	totalSize := h.typeInfo.Size()
	gmlPrintf("GML: handler_structs.go: structHandler.Encode: totalSize=%v", totalSize)

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

	var fieldOffset uint32
	for i, field := range h.typeDef.Fields {
		var fieldObj any
		fieldName := strings.TrimPrefix(field.Name, "mut ")
		// var isRecursiveRef bool
		if mapObj != nil {
			// case sensitive when reading from map
			fieldObj = mapObj[fieldName]
			// if reflect.TypeOf(fieldObj).Kind() == reflect.Map {
			// 	fieldObjAddr := int64(reflect.ValueOf(fieldObj).Pointer())
			// 	mapObjAddr := int64(reflect.ValueOf(mapObj).Pointer())
			// 	if fieldObjAddr == mapObjAddr {
			// 		isRecursiveRef = true
			// 	}
			// }
		} else {
			fieldObj = rvObj.FieldByIndex([]int{i}).Interface()
			// case insensitive when reading from struct
			// fieldObj = rvObj.FieldByNameFunc(func(s string) bool { return strings.EqualFold(s, fieldName) }).Interface()
			// if reflect.TypeOf(fieldObj).Kind() == reflect.Ptr && reflect.TypeOf(obj).Kind() == reflect.Ptr {
			// 	fieldObjAddr := int64(reflect.ValueOf(fieldObj).Pointer())
			// 	objAddr := int64(reflect.ValueOf(obj).Pointer())
			// 	if fieldObjAddr == objAddr {
			// 		isRecursiveRef = true
			// 	}
			// }
		}

		if utils.HasNil(fieldObj) {
			continue
		}

		handler := h.fieldHandlers[i]
		// if isRecursiveRef {
		// 	// This is a self-referencing field. We need to write the pointer to the struct itself.
		// 	wa.Memory().WriteUint32Le(offset+fieldOffset, offset-8)
		// } else {
		cln, err := handler.Write(ctx, wasmAdapter, offset+fieldOffset, fieldObj)
		cleaner.AddCleaner(cln)
		if err != nil {
			return nil, cleaner, err
		}
		// }

		fieldOffset += handler.TypeInfo().Size()
	}

	return []uint64{uint64(offset - 8)}, cleaner, nil
}

func (h *structHandler) getStructOutput(rv reflect.Value, data map[string]any, recursionOnFields map[int]any) (any, error) {
	// rt := h.typeInfo.ReflectedType()
	// if rt.Kind() == reflect.Map {
	// 	// for name := range recursionOnFields {
	// 	// 	data[name] = data
	// 	// }
	// 	return data, nil
	// }

	// rv := reflect.New(rt)
	if err := utils.MapToStruct(data, rv.Interface()); err != nil {
		return nil, err
	}

	for index, v := range recursionOnFields {
		fieldObj := rv.Elem().FieldByIndex([]int{index})
		// ptrValue := rv // .Elem()
		// s := ptrValue.Interface()
		// log.Printf("GML1: TypeOf(s)=%v, s=%p=%[2]T", reflect.TypeOf(s), s)
		// log.Printf("GML1: TypeOf(&s)=%v, &s=%p=%[2]T", reflect.TypeOf(&s), &s)
		// fieldObj.Set(reflect.ValueOf(s))
		fieldObj.Set(reflect.ValueOf(v))
	}

	result := rv.Elem().Interface()
	// if len(recursionOnFields) > 0 {
	// 	// We need to return a pointer to the struct itself or the return-struct-by-value
	// 	// will have an invalid self-reference.
	// 	// ptrValue := rv.Elem()
	// 	// return ptrValue.Pointer(), nil
	// 	// return rv.Pointer(), nil
	// 	// s := rv.Elem()
	// 	// return &s, nil
	// 	ptrValue := rv // .Elem()
	// 	s := ptrValue.Interface()
	// 	log.Printf("GML2: s=%p=%[1]T", s)
	// 	return s, nil
	// }
	return result, nil
}
