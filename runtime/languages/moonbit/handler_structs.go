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
	"log"
	"reflect"
	"strings"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/logger"
	"github.com/gmlewis/modus/runtime/utils"
)

const maxDepth = 5 // TODO: make this based on the depth requested in the query

func (p *planner) NewStructHandler(ctx context.Context, ti langsupport.TypeInfo) (langsupport.TypeHandler, error) {
	handler := &structHandler{
		typeHandler: *NewTypeHandler(ti),
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
}

func (h *structHandler) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	gmlPrintf("GML: handler_structs.go: structHandler.Read(offset: %v)", debugShowOffset(offset))

	// Check for recursion
	visitedPtrs := wa.(*wasmAdapter).visitedPtrs
	if visitedPtrs[offset] >= maxDepth {
		logger.Warn(ctx).Bool("user_visible", true).Msgf("Excessive recursion detected in %s. Stopping at depth %d.", h.typeInfo.Name(), maxDepth)
		return nil, nil
	}
	visitedPtrs[offset]++
	defer func() {
		n := visitedPtrs[offset]
		if n == 1 {
			delete(visitedPtrs, offset)
		} else {
			visitedPtrs[offset] = n - 1
		}
	}()

	// For debugging purposes only:
	_, _, _, err := memoryBlockAtOffset(wa, offset, 0, true)
	if err != nil {
		return nil, fmt.Errorf("structHandler failed to read memory block at offset %v: %w", debugShowOffset(offset), err)
	}

	fieldOffsets := h.typeInfo.ObjectFieldOffsets()

	m := make(map[string]any, len(h.fieldHandlers))
	for i, field := range h.typeDef.Fields {
		handler := h.fieldHandlers[i]
		fieldOffset := offset + 8 + fieldOffsets[i]
		val, err := handler.Read(ctx, wa, fieldOffset)
		if err != nil {
			return nil, err
		}
		m[field.Name] = val
	}

	return h.getStructOutput(m)
}

func (h *structHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	var mapObj map[string]any
	var rvObj reflect.Value
	if m, ok := obj.(map[string]any); ok {
		mapObj = m
	} else {
		rvObj = reflect.ValueOf(obj)
		if rvObj.Kind() != reflect.Struct {
			return nil, fmt.Errorf("expected a struct, got %s", rvObj.Kind())
		}
	}

	numFields := len(h.typeDef.Fields)
	fieldOffsets := h.typeInfo.ObjectFieldOffsets()
	cleaner := utils.NewCleanerN(numFields)

	for i, field := range h.typeDef.Fields {
		var fieldObj any
		if mapObj != nil {
			// case sensitive when reading from map
			fieldObj = mapObj[field.Name]
		} else {
			// case insensitive when reading from struct
			fieldObj = rvObj.FieldByNameFunc(func(s string) bool { return strings.EqualFold(s, field.Name) }).Interface()
		}

		fieldOffset := offset + fieldOffsets[i]
		handler := h.fieldHandlers[i]
		cln, err := handler.Write(ctx, wa, fieldOffset, fieldObj)
		cleaner.AddCleaner(cln)
		if err != nil {
			return cleaner, err
		}
	}

	return cleaner, nil
}

func (h *structHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	gmlPrintf("GML: handler_structs.go: structHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value when decoding a primitive slice but got %v: %+v", len(vals), vals)
	}

	memBlockPtr := uint32(vals[0])
	memBlock, classID, _, err := memoryBlockAtOffset(wa, memBlockPtr, 0, true)
	if err != nil {
		return nil, err
	}

	if classID != TupleBlockType {
		return nil, fmt.Errorf("expected a tuple block but got classID %v", classID)
	}

	numFields := len(h.typeDef.Fields)
	m := make(map[string]any, numFields)

	memBlockOffset := uint32(8)
	for i, field := range h.typeDef.Fields {
		handler := h.fieldHandlers[i]
		// Is this a hack? It seems that MoonBit encodes primitives _directly_ into the struct without
		// pointing to them. So we need to read the data directly from the memBlock.
		if handler.TypeInfo().IsPrimitive() {
			val, err := handler.Read(ctx, wa, memBlockPtr+memBlockOffset)
			if err != nil {
				return nil, err
			}
			m[field.Name] = val
		} else {
			// fieldOffset := binary.LittleEndian.Uint32(memBlock[8+i*4:])
			fieldOffset := binary.LittleEndian.Uint32(memBlock[memBlockOffset:])
			fieldObj, err := handler.Decode(ctx, wa, []uint64{uint64(fieldOffset)})
			if err != nil {
				return nil, err
			}
			m[field.Name] = fieldObj
		}
		memBlockOffset += handler.TypeInfo().Alignment() // TODO: Is this correct, or should it be Size() or DataSize() or EncodingLength()?
		gmlPrintf("GML: handler_structs.go: structHandler.Decode: field.Name: '%v', Alignment: %v, DataSize: %v, EncodingLength: %v, Size: %v", field.Name, handler.TypeInfo().Alignment(), handler.TypeInfo().DataSize(), handler.TypeInfo().EncodingLength(), handler.TypeInfo().Size())
	}

	return h.getStructOutput(m)

	// switch len(h.fieldHandlers) {
	// case 0:
	// 	return nil, nil
	// case 1:
	// 	data, err := h.fieldHandlers[0].Decode(ctx, wa, vals)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	m := map[string]any{h.typeDef.Fields[0].Name: data}
	// 	return h.getStructOutput(m)
	// }

	// // this doesn't need to be supported until TinyGo implements multi-value returns
	// return nil, fmt.Errorf("decoding struct of type %s is not supported", h.typeInfo.Name())
}

func (h *structHandler) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	var mapObj map[string]any
	var rvObj reflect.Value
	if m, ok := obj.(map[string]any); ok {
		mapObj = m
	} else {
		rvObj = reflect.ValueOf(obj)
		if rvObj.Kind() != reflect.Struct {
			return nil, nil, fmt.Errorf("expected a struct, got %s", rvObj.Kind())
		}
	}

	numFields := len(h.typeDef.Fields)
	results := make([]uint64, 0, numFields*2)
	cleaner := utils.NewCleanerN(numFields)

	for i, field := range h.typeDef.Fields {
		var fieldObj any
		if mapObj != nil {
			// case sensitive when reading from map
			fieldObj = mapObj[field.Name]
		} else {
			// case insensitive when reading from struct
			fieldObj = rvObj.FieldByNameFunc(func(s string) bool { return strings.EqualFold(s, field.Name) }).Interface()
		}

		handler := h.fieldHandlers[i]
		vals, cln, err := handler.Encode(ctx, wa, fieldObj)
		cleaner.AddCleaner(cln)
		if err != nil {
			return nil, cleaner, err
		}
		results = append(results, vals...)
	}

	return results, cleaner, nil
}

func (h *structHandler) getStructOutput(data map[string]any) (any, error) {
	if strings.HasPrefix(h.typeInfo.Name(), "(") {
		// Handle tuple output
		result := make([]any, 0, len(data))
		for i := 0; i < len(data); i++ {
			result = append(result, data[fmt.Sprintf("%v", i)])
		}
		return result, nil
	}

	rt := h.typeInfo.ReflectedType()
	log.Printf("GML: handler_structs.go: getStructOutput: rt.Kind()=%v", rt.Kind())
	if rt.Kind() == reflect.Map {
		return data, nil
	}

	rv := reflect.New(rt)
	if err := utils.MapToStruct(data, rv.Interface()); err != nil {
		return nil, err
	}
	return rv.Elem().Interface(), nil
}
