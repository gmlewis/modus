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
	"log"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
)

func (p *planner) NewPointerHandler(ctx context.Context, ti langsupport.TypeInfo) (langsupport.TypeHandler, error) {
	handler := &pointerHandler{
		typeHandler: *NewTypeHandler(ti),
	}
	p.AddHandler(handler)

	typeDef, err := p.metadata.GetTypeDefinition(ti.Name())
	if err != nil {
		return nil, err
	}
	handler.typeDef = typeDef

	elementHandler, err := p.GetHandler(ctx, ti.UnderlyingType().Name())
	if err != nil {
		return nil, err
	}
	handler.elementHandler = elementHandler

	return handler, nil
}

type pointerHandler struct {
	typeHandler
	typeDef        *metadata.TypeDefinition
	elementHandler langsupport.TypeHandler
}

func (h *pointerHandler) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	log.Printf("GML: handler_pointers.go: pointerHandler.Read(offset: %v)", offset)

	ptr, ok := wa.Memory().ReadUint32Le(offset)
	if !ok {
		return nil, errors.New("failed to read pointer from memory")
	}
	log.Printf("GML: handler_pointers.go: pointerHandler.Read(offset: %v): ptr=%v", offset, ptr)

	return h.readData(ctx, wa, ptr)
}

func (h *pointerHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	ptr, cln, err := h.writeData(ctx, wa, obj)
	if err != nil {
		return cln, err
	}

	if ok := wa.Memory().WriteUint32Le(offset, ptr); !ok {
		return cln,
			errors.New("failed to write object pointer to memory")
	}
	return cln, nil
}

func (h *pointerHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	log.Printf("GML: handler_pointers.go: pointerHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value, got %v: %+v", len(vals), vals)
	}

	// memBlock, _, err := memoryBlockAtOffset(wa, uint32(vals[0]))
	// if err != nil {
	// 	return nil, err
	// }
	// log.Printf("GML: handler_pointers.go: pointerHandler.Decode: memBlock=%+v", memBlock)
	// if len(memBlock) == 0 {
	// 	return nil, nil
	// }

	// sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
	// // numElements := binary.LittleEndian.Uint32(memBlock[12:16])
	// // log.Printf("GML: handler_pointers.go: pointerHandler.Decode: sliceOffset=%v, numElements=%v", sliceOffset, numElements)
	// log.Printf("GML: handler_pointers.go: pointerHandler.Decode: sliceOffset=%v", sliceOffset)

	// sliceMemBlock, _, err := memoryBlockAtOffset(wa, sliceOffset)
	// if err != nil {
	// 	return nil, err
	// }

	// // log.Printf("GML: handler_pointers.go: pointerHandler.Decode: (sliceOffset: %v, numElements: %v), sliceMemBlock=%+v", sliceOffset, numElements, sliceMemBlock)
	// log.Printf("GML: handler_pointers.go: pointerHandler.Decode: (sliceOffset: %v), sliceMemBlock(%v bytes)=%+v", sliceOffset, len(sliceMemBlock), sliceMemBlock)

	return h.readData(ctx, wa, uint32(vals[0]))
}

func (h *pointerHandler) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	ptr, cln, err := h.writeData(ctx, wa, obj)
	if err != nil {
		return nil, cln, err
	}

	return []uint64{uint64(ptr)}, cln, nil
}

func (h *pointerHandler) readData(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	if offset == 0 {
		// nil pointer
		return nil, nil
	}

	data, err := h.elementHandler.Read(ctx, wa, offset)
	if err != nil {
		return nil, err
	}
	log.Printf("GML: handler_pointers.go: pointerHandler.readData: data=%#v", data)

	ptr := utils.MakePointer(data)
	return ptr, nil
}

func (h *pointerHandler) writeData(ctx context.Context, wa langsupport.WasmAdapter, obj any) (uint32, utils.Cleaner, error) {
	log.Printf("GML: handler_pointers.go: pointerHandler.writeData is not yet implemented for MoonBit")
	return 0, nil, nil
	/*
	   	if utils.HasNil(obj) {
	   		// nil pointer
	   		return 0, nil, nil
	   	}

	   data := utils.DereferencePointer(obj)

	   ptr, cln, err := wa.(*wasmAdapter).newWasmObject(ctx, h.typeDef.Id)

	   	if err != nil {
	   		return 0, cln, nil
	   	}

	   c, err := h.elementHandler.Write(ctx, wa, ptr, data)
	   cln.AddCleaner(c)

	   	if err != nil {
	   		return 0, cln, err
	   	}

	   return ptr, cln, nil
	*/
}
