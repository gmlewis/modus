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
		return nil, err
	}
	handler.typeDef = typeDef

	keyType := ti.MapKeyType()
	valueType := ti.MapValueType()
	keysHandler, err := p.GetHandler(ctx, "Array["+keyType.Name()+"]")
	if err != nil {
		return nil, err
	}
	handler.keysHandler = keysHandler

	valuesHandler, err := p.GetHandler(ctx, "Array["+valueType.Name()+"]")
	if err != nil {
		return nil, err
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
	typeDef          *metadata.TypeDefinition
	keysHandler      langsupport.TypeHandler
	valuesHandler    langsupport.TypeHandler
	usePseudoMap     bool
	rtPseudoMap      reflect.Type
	rtPseudoMapSlice reflect.Type
}

func (h *mapHandler) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	log.Printf("GML: handler_maps.go: mapHandler.Read(offset: %v)", offset)

	log.Printf("GML: handler_maps.go: mapHandler.Read is not yet implemented for MoonBit")
	return nil, nil
	/*
	   	if offset == 0 {
	   		return nil, nil
	   	}

	   res, err := wa.(*wasmAdapter).fnReadMap.Call(ctx, uint64(h.typeDef.Id), uint64(offset))

	   	if err != nil {
	   		return nil, fmt.Errorf("failed to read %s from WASM memory: %w", h.typeInfo.Name(), err)
	   	}

	   r := res[0]

	   pKeys := uint32(r >> 32)
	   pVals := uint32(r)

	   keys, err := h.keysHandler.Read(ctx, wa, pKeys)

	   	if err != nil {
	   		return nil, err
	   	}

	   vals, err := h.valuesHandler.Read(ctx, wa, pVals)

	   	if err != nil {
	   		return nil, err
	   	}

	   rvKeys := reflect.ValueOf(keys)
	   rvVals := reflect.ValueOf(vals)
	   size := rvKeys.Len()

	   rtKey := h.keysHandler.TypeInfo().ReflectedType().Elem()

	   	if rtKey.Comparable() {
	   		// return a map
	   		m := reflect.MakeMapWithSize(h.typeInfo.ReflectedType(), size)
	   		for i := 0; i < size; i++ {
	   			m.SetMapIndex(rvKeys.Index(i), rvVals.Index(i))
	   		}
	   		return m.Interface(), nil
	   	} else {

	   		s := reflect.MakeSlice(h.rtPseudoMapSlice, size, size)
	   		for i := 0; i < size; i++ {
	   			s.Index(i).Field(0).Set(rvKeys.Index(i))
	   			s.Index(i).Field(1).Set(rvVals.Index(i))
	   		}

	   		m := reflect.New(h.rtPseudoMap).Elem()
	   		m.Field(0).Set(s)
	   		return m.Interface(), nil
	   	}
	*/
}

func (h *mapHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	ptr, cln, err := h.doWriteMap(ctx, wa, obj)
	if err != nil {
		return cln, err
	}

	mapPtr, ok := wa.Memory().ReadUint32Le(ptr)
	if !ok {
		return cln, errors.New("failed to read map internal pointer from memory")
	}

	if ok := wa.Memory().WriteUint32Le(offset, mapPtr); !ok {
		return cln, errors.New("failed to write map pointer to memory")
	}

	return cln, nil
}

func (h *mapHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	log.Printf("GML: handler_maps.go: DEBUG mapHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value when decoding a map but got %v: %+v", len(vals), vals)
	}

	if vals[0] == 0 {
		return nil, nil
	}

	memBlock, _, err := memoryBlockAtOffset(wa, uint32(vals[0]), true)
	if err != nil {
		return nil, err
	}

	if len(memBlock) != 40 {
		log.Printf("GML: handler_maps.go: mapHandler.Decode: expected memBlock length of 40, got %v", len(memBlock))
	}

	sliceOffset := binary.LittleEndian.Uint32(memBlock[8:])
	log.Printf("GML: handler_maps.go: mapHandler.Decode: sliceOffset=%+v=%v", memBlock[8:12], sliceOffset)
	ptr1 := binary.LittleEndian.Uint32(memBlock[12:])
	log.Printf("GML: handler_maps.go: mapHandler.Decode: ptr1=%+v=%v", memBlock[12:16], ptr1)
	numElements := binary.LittleEndian.Uint32(memBlock[16:])
	log.Printf("GML: handler_maps.go: mapHandler.Decode: numElements=%+v=%v", memBlock[16:20], numElements)

	if numElements == 0 {
		// return an empty map
		m := reflect.MakeMapWithSize(h.typeInfo.ReflectedType(), 0)
		return m.Interface(), nil
	}

	const2 := binary.LittleEndian.Uint32(memBlock[20:])
	log.Printf("GML: handler_maps.go: mapHandler.Decode: const2=%+v=%v", memBlock[20:24], const2)
	const3 := binary.LittleEndian.Uint32(memBlock[24:])
	log.Printf("GML: handler_maps.go: mapHandler.Decode: const3=%+v=%v", memBlock[24:28], const3)
	const4 := binary.LittleEndian.Uint32(memBlock[28:])
	log.Printf("GML: handler_maps.go: mapHandler.Decode: const4=%+v=%v", memBlock[28:32], const4)
	firstKVPair := binary.LittleEndian.Uint32(memBlock[32:])
	log.Printf("GML: handler_maps.go: mapHandler.Decode: firstKVPair=%+v=%v", memBlock[32:36], firstKVPair)
	lastKVPair := binary.LittleEndian.Uint32(memBlock[36:])
	log.Printf("GML: handler_maps.go: mapHandler.Decode: lastKVPair=%+v=%v", memBlock[36:40], lastKVPair)
	// size := numElements * uint32(h.converter.TypeSize())

	sliceMemBlock, _, err := memoryBlockAtOffset(wa, sliceOffset)
	if err != nil {
		return nil, err
	}
	ptr1MemBlock, _, err := memoryBlockAtOffset(wa, ptr1)
	if err != nil {
		return nil, err
	}
	firstKVPairMemBlock, _, err := memoryBlockAtOffset(wa, firstKVPair)
	if err != nil {
		return nil, err
	}
	lastKVPairMemBlock, _, err := memoryBlockAtOffset(wa, lastKVPair)
	if err != nil {
		return nil, err
	}
	log.Printf("GML: handler_maps.go: mapHandler.Decode: (sliceOffset: %v, numElements: %v), sliceMemBlock(%+v=@%v)=(%v bytes)=%+v", sliceOffset, numElements, memBlock[8:12], sliceOffset, len(sliceMemBlock), sliceMemBlock)
	dumpMemBlock(wa, "sliceMemBlock", sliceMemBlock)
	log.Printf("GML: handler_maps.go: mapHandler.Decode: (sliceOffset: %v, numElements: %v), ptr1(%+v=@%v)=(%v bytes)=%+v", sliceOffset, numElements, memBlock[12:16], ptr1, len(ptr1MemBlock), ptr1MemBlock)
	dumpMemBlock(wa, "ptr1", ptr1MemBlock)
	log.Printf("GML: handler_maps.go: mapHandler.Decode: (sliceOffset: %v, numElements: %v), firstKVPair(%+v=@%v)=(%v bytes)=%+v", sliceOffset, numElements, memBlock[32:36], firstKVPair, len(firstKVPairMemBlock), firstKVPairMemBlock)
	dumpMemBlock(wa, "firstKVPair", firstKVPairMemBlock)
	log.Printf("GML: handler_maps.go: mapHandler.Decode: (sliceOffset: %v, numElements: %v), lastKVPair(%+v=@%v)=(%v bytes)=%+v", sliceOffset, numElements, memBlock[36:40], lastKVPair, len(lastKVPairMemBlock), lastKVPairMemBlock)
	dumpMemBlock(wa, "lastKVPair", lastKVPairMemBlock)

	log.Printf("GML: handler_maps.go: ORIGINAL mapHandler.Decode(vals: %+v)", vals)

	mapPtr := uint32(vals[0])
	if mapPtr == 0 {
		return nil, nil
	}

	// we need to make a pointer to the map data
	ptr, cln, err := wa.(*wasmAdapter).AllocateMemory(ctx, 4)
	defer func() {
		if cln != nil {
			if e := cln.Clean(); e != nil && err == nil {
				err = e
			}
		}
	}()
	if err != nil {
		return nil, fmt.Errorf("failed to allocate memory for map pointer: %w", err)
	}

	if ok := wa.Memory().WriteUint32Le(ptr, mapPtr); !ok {
		return nil, errors.New("failed to write map pointer data to memory")
	}

	// now we can use that pointer to read the map
	return h.Read(ctx, wa, ptr)
}

func dumpMemBlock(wa langsupport.WasmAdapter, namePrefix string, memBlock []byte) {
	for i := 8; i < len(memBlock); i += 4 {
		ptr := binary.LittleEndian.Uint32(memBlock[i : i+4])
		if ptr == 0 {
			continue
		}
		newMemBlock, words, err := memoryBlockAtOffset(wa, ptr)
		if err != nil {
			continue
		}
		if len(newMemBlock) > 8 && newMemBlock[0] == 255 && newMemBlock[1] == 255 && newMemBlock[2] == 255 && newMemBlock[3] == 255 && newMemBlock[4] == 243 {
			stringData, err := stringDataFromMemBlock(newMemBlock, words)
			if err == nil {
				s, err := doReadString(stringData)
				if err == nil {
					log.Printf("GML: handler_maps.go: dumpMemBlock: %v(%v bytes): memBlock[%v:%v]=%+v=@%v=%+v = '%v'", namePrefix, len(memBlock), i, i+4, memBlock[i:i+4], ptr, newMemBlock, s)
					continue
				}
			}
		}
		log.Printf("GML: handler_maps.go: dumpMemBlock: %v(%v bytes): memBlock[%v:%v]=%+v=@%v=%+v", namePrefix, len(memBlock), i, i+4, memBlock[i:i+4], ptr, newMemBlock)
	}
}

func (h *mapHandler) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	ptr, cln, err := h.doWriteMap(ctx, wa, obj)
	if err != nil {
		return nil, cln, err
	}

	mapPtr, ok := wa.Memory().ReadUint32Le(ptr)
	if !ok {
		return nil, cln, errors.New("failed to read map internal pointer from memory")
	}

	return []uint64{uint64(mapPtr)}, cln, nil
}

func (h *mapHandler) doWriteMap(ctx context.Context, wa langsupport.WasmAdapter, obj any) (pMap uint32, cln utils.Cleaner, err error) {
	log.Printf("GML: handler_maps.go: mapHandler.doWriteMap is not yet implemented for MoonBit")
	return 0, nil, nil
	/*
	   	if utils.HasNil(obj) {
	   		return 0, nil, nil
	   	}

	   m, err := utils.ConvertToMap(obj)

	   	if err != nil {
	   		return 0, nil, err
	   	}

	   keys, vals := utils.MapKeysAndValues(m)

	   pMap, cln, err = wa.(*wasmAdapter).makeWasmObject(ctx, h.typeDef.Id, uint32(len(m)))

	   	if err != nil {
	   		return 0, nil, err
	   	}

	   innerCln := utils.NewCleanerN(2)

	   	defer func() {
	   		// clean up the slices after the map is written to memory
	   		if e := innerCln.Clean(); e != nil && err == nil {
	   			err = e
	   		}
	   	}()

	   pKeys, c, err := h.keysHandler.(sliceWriter).doWriteSlice(ctx, wa, keys)
	   innerCln.AddCleaner(c)

	   	if err != nil {
	   		return 0, cln, err
	   	}

	   pVals, c, err := h.valuesHandler.(sliceWriter).doWriteSlice(ctx, wa, vals)
	   innerCln.AddCleaner(c)

	   	if err != nil {
	   		return 0, cln, err
	   	}

	   	if _, err := wa.(*wasmAdapter).fnWriteMap.Call(ctx, uint64(h.typeDef.Id), uint64(pMap), uint64(pKeys), uint64(pVals)); err != nil {
	   		return 0, cln, fmt.Errorf("failed to write %s to WASM memory: %w", h.typeInfo.Name(), err)
	   	}

	   return pMap, cln, nil
	*/
}
