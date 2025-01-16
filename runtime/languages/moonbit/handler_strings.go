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
	"unsafe"

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
	if offset == 0 {
		return "", nil
	}

	data, size, err := h.readStringHeader(wa, offset)
	if err != nil {
		return "", err
	}

	return h.doReadString(wa, data, size)
}

func (h *stringHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	str, err := cast.ToStringE(obj)
	if err != nil {
		return nil, err
	}

	log.Printf("stringHandler.Write: Calling doWriteString: str='%v'", str)

	ptr, cln, err := h.doWriteString(ctx, wa, str)
	if err != nil {
		return cln, err
	}

	data, size, err := h.readStringHeader(wa, ptr)
	if err != nil {
		return cln, err
	}

	return cln, h.writeStringHeader(wa, data, size, offset)
}

func (h *stringHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	if len(vals) != 1 {
		return nil, fmt.Errorf("MoonBit: expected 1 value when decoding a string but got %v: %+v", len(vals), vals)
	}

	offset := uint32(vals[0])
	size := uint32(8)
	bytes, ok := wa.Memory().Read(offset, size)
	if !ok {
		return "", fmt.Errorf("failed to read string data from WASM memory (offset: %v, size: %v)", offset, size)
	}

	if bytes[0] != 0xff || bytes[1] != 0xff || bytes[2] != 0xff || bytes[3] != 0xff {
		log.Printf("WARNING: stringHandler.Decode: expected 'ffffffff' but got '%02x%02x%02x%02x'; continuing",
			bytes[0], bytes[1], bytes[2], bytes[3])
	}
	size = binary.LittleEndian.Uint32(bytes[4:8])
	bytes, ok = wa.Memory().Read(offset, size)
	if !ok {
		return "", fmt.Errorf("failed to read string data from WASM memory (offset: %v, size: %v)", offset, size)
	}

	return convertMoonBitUTF16ToUTF8(bytes)
}

func (h *stringHandler) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	str, err := cast.ToStringE(obj)
	if err != nil {
		return nil, nil, err
	}

	bytes := convertGoUTF8ToUTF16(str)

	ptr, cln, err := h.doWriteBytes(ctx, wa, bytes)
	if err != nil {
		return nil, cln, err
	}

	return []uint64{uint64(ptr)}, cln, nil
}

func (h *stringHandler) doReadString(wa langsupport.WasmAdapter, offset, size uint32) (string, error) {
	if offset == 0 || size == 0 {
		return "", nil
	}

	bytes, ok := wa.Memory().Read(offset, size)
	if !ok {
		return "", fmt.Errorf("failed to read string data from WASM memory (size: %d)", size)
	}

	return unsafe.String(&bytes[0], size), nil
}

func (h *stringHandler) doWriteString(ctx context.Context, wa langsupport.WasmAdapter, str string) (uint32, utils.Cleaner, error) {
	log.Printf("GML: handler_strings.go: stringHandler.doWriteString is not yet implemented for MoonBit")
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

func (h *stringHandler) readStringHeader(wa langsupport.WasmAdapter, offset uint32) (data, size uint32, err error) {
	if offset == 0 {
		return 0, 0, nil
	}

	val, ok := wa.Memory().ReadUint64Le(offset)
	if !ok {
		return 0, 0, errors.New("failed to read string header from WASM memory")
	}

	data = uint32(val)
	size = uint32(val >> 32)

	return data, size, nil
}

func (h *stringHandler) writeStringHeader(wa langsupport.WasmAdapter, data, size, offset uint32) error {
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

	size := 8 + len(codeUnits)*2 + 6
	data := make([]byte, size)
	binary.LittleEndian.PutUint32(data[0:], 0xffffffff)
	binary.LittleEndian.PutUint32(data[4:], uint32(size)) // TODO: Is this the correct length to report?
	for i, codeUnit := range codeUnits {
		binary.LittleEndian.PutUint16(data[8+i*2:], codeUnit)
	}

	return data
}

func (h *stringHandler) doWriteBytes(ctx context.Context, wa langsupport.WasmAdapter, bytes []byte) (uint32, utils.Cleaner, error) {
	const id = 2 // ID for string is always 2
	size := uint32(len(bytes))
	ptr, cln, err := wa.(*wasmAdapter).allocateAndPinMemory(ctx, size, id)
	if err != nil {
		return 0, cln, err
	}

	if ok := wa.Memory().Write(ptr, bytes); !ok {
		return 0, cln, fmt.Errorf("failed to write string data to WASM memory (size: %d)", size)
	}

	return ptr, cln, nil
}
