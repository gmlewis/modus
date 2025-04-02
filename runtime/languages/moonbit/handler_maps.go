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
	"errors"
	"fmt"
	"reflect"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
)

func (p *planner) NewMapHandler(ctx context.Context, ti langsupport.TypeInfo) (langsupport.TypeHandler, error) {
	handler := &mapHandler{
		typeHandler: *NewTypeHandler(ti),
	}
	p.AddHandler(handler)

	typeDef, err := p.metadata.GetTypeDefinition(ti.Name())
	if err != nil {
		return nil, fmt.Errorf("planner.NewMapHandler: p.metadata.GetTypeDefinition('%v'): %w", ti.Name(), err)
	}
	handler.typeDef = typeDef

	keyType := ti.MapKeyType()
	valueType := ti.MapValueType()

	// A mapHandler always needs a stringsHandler so that it can call the "write_map" function to generate
	// a map based on the key and value type names (as Strings).
	stringsHandler, err := p.GetHandler(ctx, "String")
	if err != nil {
		return nil, fmt.Errorf("planner.NewMapHandler: p.GetHandler('String'): %w", err)
	}
	handler.stringsHandler = stringsHandler

	sliceOfKeysType := "Array[" + keyType.Name() + "]"
	sliceOfKeysHandler, err := p.GetHandler(ctx, sliceOfKeysType)
	if err != nil {
		return nil, fmt.Errorf("planner.NewMapHandler: p.GetHandler('%v'): %w", sliceOfKeysType, err)
	}
	handler.sliceOfKeysHandler = sliceOfKeysHandler
	p.AddHandler(sliceOfKeysHandler)

	keysHandler, err := p.GetHandler(ctx, keyType.Name())
	if err != nil {
		return nil, fmt.Errorf("planner.NewMapHandler: p.GetHandler('%v'): %w", keyType.Name(), err)
	}
	handler.keysHandler = keysHandler

	sliceOfValuesType := "Array[" + valueType.Name() + "]"
	sliceOfValuesHandler, err := p.GetHandler(ctx, sliceOfValuesType)
	if err != nil {
		return nil, fmt.Errorf("planner.NewMapHandler: p.GetHandler('%v'): %w", sliceOfValuesType, err)
	}
	handler.sliceOfValuesHandler = sliceOfValuesHandler
	p.AddHandler(sliceOfValuesHandler)

	valuesHandler, err := p.GetHandler(ctx, valueType.Name())
	if err != nil {
		return nil, fmt.Errorf("planner.NewMapHandler: p.GetHandler('%v'): %w", valueType.Name(), err)
	}
	handler.valuesHandler = valuesHandler

	rtKey := keyType.ReflectedType()
	rtValue := valueType.ReflectedType()
	if !rtKey.Comparable() {
		handler.usePseudoMap = true
		handler.rtPseudoMapSlice = reflect.SliceOf(reflect.StructOf([]reflect.StructField{
			{
				Name: "Key",
				Type: rtKey,
				Tag:  `json:"key"`,
			},
			{
				Name: "Value",
				Type: rtValue,
				Tag:  `json:"value"`,
			},
		}))

		handler.rtPseudoMap = reflect.StructOf([]reflect.StructField{
			{
				Name: "Data",
				Type: handler.rtPseudoMapSlice,
				Tag:  `json:"$mapdata"`,
			},
		})
	}

	return handler, nil
}

type mapHandler struct {
	typeHandler
	typeDef              *metadata.TypeDefinition
	sliceOfKeysHandler   langsupport.TypeHandler
	sliceOfValuesHandler langsupport.TypeHandler
	keysHandler          langsupport.TypeHandler
	valuesHandler        langsupport.TypeHandler
	stringsHandler       langsupport.TypeHandler
	usePseudoMap         bool
	rtPseudoMap          reflect.Type
	rtPseudoMapSlice     reflect.Type
}

func (h *mapHandler) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	if offset == 0 {
		return nil, nil
	}

	keyTypeName := h.keysHandler.TypeInfo().Name()
	valueTypeName := h.valuesHandler.TypeInfo().Name()

	keyTypeNamePtr, _, err := h.stringsHandler.Encode(ctx, wa, keyTypeName)
	// TODO: innerCln.AddCleaner(c)
	if err != nil {
		return nil, err
	}
	valueTypeNamePtr, _, err := h.stringsHandler.Encode(ctx, wa, valueTypeName)
	// TODO: innerCln.AddCleaner(c)
	if err != nil {
		return nil, err
	}

	params := []uint64{keyTypeNamePtr[0], valueTypeNamePtr[0], uint64(offset)}
	res, err := wa.(*wasmAdapter).fnReadMap.Call(ctx, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s from WASM memory: %w", h.typeInfo.Name(), err)
	}

	r := res[0]
	pKeys := uint32(r >> 32)
	pVals := uint32(r)

	keys, err := h.sliceOfKeysHandler.Read(ctx, wa, pKeys)
	if err != nil {
		return nil, err
	}

	vals, err := h.sliceOfValuesHandler.Read(ctx, wa, pVals)
	if err != nil {
		return nil, err
	}

	rvKeys := reflect.ValueOf(keys)
	rvVals := reflect.ValueOf(vals)
	size := rvKeys.Len()

	// A MoonBit Map must always use comparable keys (as defined by Go), right?
	m := reflect.MakeMapWithSize(h.typeInfo.ReflectedType(), size)
	for i := 0; i < size; i++ {
		m.SetMapIndex(rvKeys.Index(i), rvVals.Index(i))
	}
	return m.Interface(), nil
}

func (h *mapHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	ptr, cln, err := h.doWriteMap(ctx, wa, obj)
	if err != nil {
		return cln, err
	}

	if ok := wa.Memory().WriteUint32Le(offset, ptr); !ok {
		return cln, errors.New("failed to write map pointer to memory")
	}

	return cln, nil
}

// Decode should always be passed an address to a Map[] in memory.
func (h *mapHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value when decoding a map but got %v: %+v", len(vals), vals)
	}

	return h.Read(ctx, wa, uint32(vals[0]))
}

func (h *mapHandler) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	ptr, cln, err := h.doWriteMap(ctx, wa, obj)
	if err != nil {
		return nil, cln, err
	}

	return []uint64{uint64(ptr)}, cln, nil
}

func (h *mapHandler) doWriteMap(ctx context.Context, wa langsupport.WasmAdapter, obj any) (pMap uint32, cln utils.Cleaner, err error) {
	if utils.HasNil(obj) {
		return 0, nil, nil
	}

	m, err := utils.ConvertToMap(obj)

	if err != nil {
		return 0, nil, err
	}

	keys, vals := utils.MapKeysAndValues(m)

	innerCln := utils.NewCleanerN(4)

	defer func() {
		// clean up the slices after the map is written to memory
		if e := innerCln.Clean(); e != nil && err == nil {
			err = e
		}
	}()

	keyTypeName := h.keysHandler.TypeInfo().Name()
	valueTypeName := h.valuesHandler.TypeInfo().Name()

	keyTypeNamePtr, c, err := h.stringsHandler.Encode(ctx, wa, keyTypeName)
	innerCln.AddCleaner(c)
	if err != nil {
		return 0, cln, err
	}
	valueTypeNamePtr, c, err := h.stringsHandler.Encode(ctx, wa, valueTypeName)
	innerCln.AddCleaner(c)
	if err != nil {
		return 0, cln, err
	}

	pKeys, c, err := h.sliceOfKeysHandler.Encode(ctx, wa, keys)
	innerCln.AddCleaner(c)
	if err != nil {
		return 0, cln, err
	}

	pVals, c, err := h.sliceOfValuesHandler.Encode(ctx, wa, vals)
	innerCln.AddCleaner(c)
	if err != nil {
		return 0, cln, err
	}

	params := []uint64{keyTypeNamePtr[0], valueTypeNamePtr[0], pKeys[0], pVals[0]}
	res, err := wa.(*wasmAdapter).fnWriteMap.Call(ctx, params...)
	if err != nil {
		return 0, cln, fmt.Errorf("failed to write %s to WASM memory: %w", h.typeInfo.Name(), err)
	}

	return uint32(res[0]), cln, nil
}
