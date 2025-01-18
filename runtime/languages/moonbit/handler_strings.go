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
	"unicode/utf16"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"

	"github.com/spf13/cast"
)

const stringHeaderSize = 8

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

	data, size, err := h.readStringHeader(ctx, wa, offset)
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

	data, size, err := h.readStringHeader(ctx, wa, ptr)
	if err != nil {
		return cln, err
	}

	return cln, h.writeStringHeader(wa, data, size, offset)
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

	offset, size, err := h.readStringHeader(ctx, wa, uint32(vals[0]))
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
	// res[0] += stringHeaderSize

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

func (h *stringHandler) readStringHeader(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (data, size uint32, err error) {
	log.Printf("GML: handler_strings.go: stringHandler.readStringHeader(offset: %v)", offset)

	if offset == 0 {
		return 0, 0, nil
	}

	// First, call str2ptr to get the pointer to the string data.
	res, err := wa.(*wasmAdapter).fnStr2ptr.Call(ctx, uint64(offset))
	if err != nil {
		return 0, 0, err
	}
	log.Printf("GML: handler_strings.go: stringHandler.readStringHeader: fnStr2ptr.Call(%v) returned %+v", offset, res)

	// val, ok := wa.Memory().ReadUint64Le(offset)
	// if !ok {
	// 	return 0, 0, errors.New("failed to read string header from WASM memory")
	// }

	// data = uint32(val)
	// size = uint32(val >> 32)

	// return data, size, nil

	bytes, ok := wa.Memory().Read(offset, uint32(8))
	if !ok {
		return 0, 0, fmt.Errorf("failed to read string data from WASM memory A: (offset: %v, size: %v)", offset, size)
	}

	if bytes[0] != 0xff || bytes[1] != 0xff || bytes[2] != 0xff || bytes[3] != 0xff {
		log.Printf("WARNING: stringHandler.Decode: expected 'ffffffff' but got '%02x%02x%02x%02x'; continuing",
			bytes[0], bytes[1], bytes[2], bytes[3])
	}
	size = binary.LittleEndian.Uint32(bytes[4:8])

	return offset, size, nil
}

func (h *stringHandler) writeStringHeader(wa langsupport.WasmAdapter, data, size, offset uint32) error {
	log.Printf("GML: handler_strings.go: stringHandler.writeStringHeader(data: %v, size: %v, offset: %v)", data, size, offset)

	val := uint64(size)<<32 | uint64(data)
	if ok := wa.Memory().WriteUint64Le(offset, val); !ok {
		return errors.New("failed to write string header to WASM memory")
	}

	return nil
}

// convertMoonBitUTF16ToUTF8 converts a wasm-encoded MoonBit UTF-16 String to a Go UTF-8 string.
func convertMoonBitUTF16ToUTF8(data []byte) (string, error) {
	if len(data) < 8 {
		return "", fmt.Errorf("data too short to contain metadata")
	}
	data = data[8:]

	var codeUnits []uint16
	for i := 0; i+1 < len(data); i += 2 {
		codeUnit := binary.LittleEndian.Uint16(data[i:])
		if codeUnit == 0 || codeUnit == 256 {
			break
		}
		codeUnits = append(codeUnits, codeUnit)
	}

	runes := utf16.Decode(codeUnits)
	return string(runes), nil
}

// convertGoUTF8ToUTF16 converts a Go UTF-8 string to a wasm-encoded MoonBit UTF-16 String.
func convertGoUTF8ToUTF16(str string) []byte {
	runes := []rune(str)
	codeUnits := utf16.Encode(runes)

	size := (len(codeUnits) + 1) * 2 // null-terminated UTF-16 string
	data := make([]byte, size)
	for i, codeUnit := range codeUnits {
		binary.LittleEndian.PutUint16(data[i*2:], codeUnit)
	}
	binary.LittleEndian.PutUint16(data[size-2:], 0) // null-terminate the string

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
