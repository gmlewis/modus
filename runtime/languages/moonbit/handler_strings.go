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
	"unicode/utf16"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"

	"github.com/spf13/cast"
)

func (p *planner) NewStringHandler(ti langsupport.TypeInfo) (langsupport.TypeHandler, error) {
	handler := &stringHandler{*NewTypeHandler(ti)}
	p.AddHandler(handler)
	return handler, nil
}

type stringHandler struct {
	typeHandler
}

func (h *stringHandler) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	log.Printf("GML: handler_strings.go: stringHandler.Read(offset: %v)", offset)

	if offset == 0 {
		return "", nil
	}

	data, size, err := memoryBlockAtOffset(wa, offset)
	if err != nil {
		return "", err
	}

	return h.doReadString(wa, data, size)
}

func (h *stringHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	log.Printf("GML: handler_strings.go: stringHandler.Write(offset: %v, obj: '%v')", offset, obj)

	str, err := cast.ToStringE(obj)
	if err != nil {
		return nil, err
	}

	log.Printf("stringHandler.Write: Calling doWriteString: str='%v'", str)

	ptr, cln, err := h.doWriteString(ctx, wa, str)
	if err != nil {
		return cln, err
	}

	data, size, err := memoryBlockAtOffset(wa, ptr)
	if err != nil {
		return cln, err
	}

	return cln, writeMemoryBlockHeader(wa, data, size, offset)
}

func (h *stringHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	log.Printf("GML: handler_strings.go: stringHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("MoonBit: expected 1 value when decoding a string but got %v: %+v", len(vals), vals)
	}

	// offset := uint32(vals[0])
	// size := uint32(8)
	// bytes, ok := wa.Memory().Read(offset, size)
	// if !ok {
	// 	return "", fmt.Errorf("failed to read string data from WASM memory A: (offset: %v, size: %v)", offset, size)
	// }

	// if bytes[0] != 0xff || bytes[1] != 0xff || bytes[2] != 0xff || bytes[3] != 0xff {
	// 	log.Printf("WARNING: stringHandler.Decode: expected 'ffffffff' but got '%02x%02x%02x%02x'; continuing",
	// 		bytes[0], bytes[1], bytes[2], bytes[3])
	// }
	// size = binary.LittleEndian.Uint32(bytes[4:8])

	offset, size, err := memoryBlockAtOffset(wa, uint32(vals[0]))
	if err != nil {
		return "", err
	}

	// bytes, ok = wa.Memory().Read(offset, size)
	// if !ok {
	// 	return "", fmt.Errorf("failed to read string data from WASM memory B: (offset: %v, size: %v)", offset, size)
	// }

	// return convertMoonBitUTF16ToUTF8(bytes)
	return h.doReadString(wa, offset, size)
}

func (h *stringHandler) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	log.Printf("GML: handler_strings.go: stringHandler.Encode(obj: '%v')", obj)

	str, err := cast.ToStringE(obj)
	if err != nil {
		return nil, nil, err
	}

	bytes := convertGoUTF8ToUTF16(str)

	offset, cln, err := h.doWriteBytes(ctx, wa, bytes)
	if err != nil {
		return nil, cln, err
	}

	// Call ptr2str on the raw memory to convert the data to a MoonBit String (i32) that can be passed to a function.
	log.Printf("GML: handler_strings.go: stringHandler.Encode: Calling fnPtr2str.Call(%v)", offset)
	res, err := wa.(*wasmAdapter).fnPtr2str.Call(ctx, uint64(offset))
	if err != nil {
		return nil, cln, err
	}

	// For debugging purposes:
	memBytes, ok := wa.Memory().Read(offset-8, uint32(len(bytes)+16))
	if !ok {
		log.Printf("GML: handler_strings.go: stringHandler.Encode: failed to read memory bytes")
	} else {
		log.Printf("GML: handler_strings.go: stringHandler.Encode: memBytes: %+v", memBytes)
	}

	return res, cln, nil
}

func (h *stringHandler) doReadString(wa langsupport.WasmAdapter, offset, size uint32) (string, error) {
	log.Printf("GML: handler_strings.go: stringHandler.doReadString(offset: %v, size: %v)", offset, size)

	if offset == 0 || size == 0 {
		return "", nil
	}

	// bytes, ok := wa.Memory().Read(offset, size)
	// if !ok {
	// 	return "", fmt.Errorf("failed to read string data from WASM memory C: (offset: %v, size: %v)", offset, size)
	// }

	// return unsafe.String(&bytes[0], size), nil

	bytes, ok := wa.Memory().Read(offset, size)
	if !ok {
		return "", fmt.Errorf("failed to read string data from WASM memory B: (offset: %v, size: %v)", offset, size)
	}
	log.Printf("GML: handler_strings.go: stringHandler.doReadString: bytes: %+v", bytes)

	return convertMoonBitUTF16ToUTF8(bytes)
}

func (h *stringHandler) doWriteString(ctx context.Context, wa langsupport.WasmAdapter, str string) (uint32, utils.Cleaner, error) {
	log.Printf("GML: handler_strings.go: stringHandler.doWriteString(str: '%v') is not yet implemented for MoonBit", str)
	return 0, nil, nil
	/*
	   const id = 2 // ID for string is always 2
	   ptr, cln, err := wa.(*wasmAdapter).makeWasmObject(ctx, id, uint32(len(str)))

	   	if err != nil {
	   		return 0, cln, err
	   	}

	   offset, ok := wa.Memory().ReadUint32Le(ptr)

	   	if !ok {
	   		return 0, cln, errors.New("failed to read string data pointer from WASM memory")
	   	}

	   	if ok := wa.Memory().WriteString(offset, str); !ok {
	   		return 0, cln, fmt.Errorf("failed to write string data to WASM memory (size: %d)", len(str))
	   	}

	   return ptr, cln, nil
	*/
}

// convertMoonBitUTF16ToUTF8 converts a wasm-encoded MoonBit UTF-16 String to a Go UTF-8 string.
func convertMoonBitUTF16ToUTF8(data []byte) (string, error) {
	codeUnits := make([]uint16, 0, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		codeUnit := binary.LittleEndian.Uint16(data[i:])
		codeUnits = append(codeUnits, codeUnit)
	}

	runes := utf16.Decode(codeUnits)
	return string(runes), nil
}

// convertGoUTF8ToUTF16 converts a Go UTF-8 string to a wasm-encoded MoonBit UTF-16 String.
func convertGoUTF8ToUTF16(str string) []byte {
	runes := []rune(str)
	codeUnits := utf16.Encode(runes)

	size := len(codeUnits) * 2
	data := make([]byte, size)
	for i, codeUnit := range codeUnits {
		binary.LittleEndian.PutUint16(data[i*2:], codeUnit)
	}

	return data
}

func (h *stringHandler) doWriteBytes(ctx context.Context, wa langsupport.WasmAdapter, bytes []byte) (uint32, utils.Cleaner, error) {
	const id = 2 // ID for string is always 2
	size := uint32(len(bytes))
	offset, cln, err := wa.(*wasmAdapter).allocateAndPinMemory(ctx, size, id)
	if err != nil {
		return 0, cln, err
	}

	log.Printf("GML: handler_strings.go: stringHandler.doWriteBytes(bytes(%v): %+v), got offset: %v", size, bytes, offset)
	if ok := wa.Memory().Write(offset, bytes); !ok {
		return 0, cln, fmt.Errorf("failed to write string data to WASM memory (offset: %v, size: %v)", offset, size)
	}

	return offset, cln, nil
}
