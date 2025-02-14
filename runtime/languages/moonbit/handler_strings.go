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
	"unicode/utf16"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
	wasm "github.com/tetratelabs/wazero/api"

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
	gmlPrintf("GML: handler_strings.go: stringHandler.Read(offset: %v)", debugShowOffset(offset))
	if offset == 0 {
		if h.typeInfo.IsPointer() {
			return nil, nil
		}
		return "", nil
	}

	// TODO: How do we know if Read should directly intpret the offset as pointing to the String itself
	// (as in a sliceHandler) or if it should interpret the offset the address of the String (as in a structHandler)?
	// For now, first attempt to decode the offset as a string, and if that doesn't work, interpret it as a pointer to a string.
	if s, err := h.Decode(ctx, wa, []uint64{uint64(offset)}); err == nil {
		return s, nil
	}

	// Read pointer to MoonBit String
	ptr, ok := wa.Memory().ReadUint32Le(offset)
	if !ok {
		gmlPrintf("GML: handler_strings.go: stringHandler.Read: failed to read pointer to String")
	} else {
		gmlPrintf("GML: handler_strings.go: stringHandler.Read: ptr: %v", debugShowOffset(ptr))
	}

	vals := []uint64{uint64(ptr)}
	return h.Decode(ctx, wa, vals)
}

func (h *stringHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	gmlPrintf("GML: handler_strings.go: stringHandler.Write(offset: %v, obj: '%v')", offset, obj)

	str, err := cast.ToStringE(obj)
	if err != nil {
		return nil, err
	}

	gmlPrintf("stringHandler.Write: Calling doWriteString: str='%v'", str)

	ptr, cln, err := h.doWriteString(ctx, wa, str)
	if err != nil {
		return cln, err
	}

	// memBlock, err := stringDataAtOffset(wa, ptr)
	// if err != nil {
	// 	return cln, err
	// }
	// data, size := offset+8, uint32(len(memBlock))

	// return cln, writeMemoryBlockHeader(wa, data, size, offset)

	if ok := wa.Memory().WriteUint32Le(offset, ptr); !ok {
		return cln, fmt.Errorf("failed to write pointer to string data to WASM memory (offset: %v, ptr: %v)", offset, ptr)
	}

	return cln, nil
}

func (h *stringHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	gmlPrintf("GML: handler_strings.go: stringHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("MoonBit: expected 1 value when decoding a string but got %v: %+v", len(vals), vals)
	}

	if vals[0] == 0 {
		return nil, nil
	}

	data, err := stringDataAtOffset(wa, uint32(vals[0]))
	if err != nil {
		return "", err
	}

	s, err := doReadString(data)
	if err != nil {
		return nil, err
	}

	if h.typeInfo.IsPointer() {
		return &s, nil
	}

	return s, nil
}

func (h *stringHandler) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	gmlPrintf("GML: handler_strings.go: stringHandler.Encode(obj: '%v')", obj)

	str, err := cast.ToStringE(obj)
	if err != nil {
		return nil, nil, err
	}

	bytes := convertGoUTF8ToUTF16(str)
	wasmWriter, ok := wa.(*wasmAdapter)
	if !ok {
		return nil, nil, fmt.Errorf("expected *wasmAdapter, got %T", wa)
	}
	offset, cln, err := h.doWriteStringBytes(ctx, wasmWriter, bytes)
	if err != nil {
		return nil, cln, err
	}

	// // Call ptr2str on the raw memory to convert the data to a MoonBit String (i32) that can be passed to a function.
	// gmlPrintf("GML: handler_strings.go: stringHandler.Encode: Calling fnPtr2str.Call(%v)", offset)
	// res, err := wa.(*wasmAdapter).fnPtr2str.Call(ctx, uint64(offset))
	// if err != nil {
	// 	return nil, cln, err
	// }

	// For debugging purposes:
	// _, _, _ = memoryBlockAtOffset(wa, uint32(res[0]), 0, true)
	_, _, _ = memoryBlockAtOffset(wa, offset, 0, true)

	// return res, cln, nil
	return []uint64{uint64(offset)}, cln, nil
}

func doReadString(bytes []byte) (string, error) {
	if len(bytes) == 0 {
		gmlPrintf("GML: handler_strings.go: stringHandler.doReadString: bytes: [] = ''")
		return "", nil
	}

	s, err := convertMoonBitUTF16ToUTF8(bytes)
	// gmlPrintf("GML: handler_strings.go: stringHandler.doReadString: bytes: %+v = '%v'", bytes, s)
	return s, err
}

func (h *stringHandler) doWriteString(ctx context.Context, wa langsupport.WasmAdapter, str string) (uint32, utils.Cleaner, error) {
	res, cln, err := h.Encode(ctx, wa, str)
	if err != nil {
		return 0, cln, err
	}

	return uint32(res[0]), cln, nil
}

// convertMoonBitUTF16ToUTF8 converts a wasm-encoded MoonBit UTF-16 String to a Go UTF-8 string.
func convertMoonBitUTF16ToUTF8(data []byte) (string, error) {
	codeUnits := make([]uint16, 0, len(data)/2)
	for i := 0; i+1 < len(data); i += 2 {
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

// For testing purposes:
type wasmMemoryWriter interface {
	allocateAndPinMemory(ctx context.Context, size, blockType uint32) (uint32, utils.Cleaner, error)
	Memory() wasm.Memory
}

func (h *stringHandler) doWriteStringBytes(ctx context.Context, wa wasmMemoryWriter, bytes []byte) (uint32, utils.Cleaner, error) {
	size := uint32(len(bytes))
	words := uint32((size + 5) / 4)
	totalSize := words * 4
	offset, cln, err := wa.allocateAndPinMemory(ctx, totalSize, StringBlockType)
	if err != nil {
		return 0, cln, err
	}

	remainderOffset := words*4 + 7
	remainder := uint8((size + 3) % 4)
	wa.Memory().WriteByte(offset-8+remainderOffset, remainder)

	gmlPrintf("GML: handler_strings.go: stringHandler.doWriteStringBytes(bytes(%v): %+v), got offset: %v, remainderOffset: %v, remainder: %v", size, bytes, offset, remainderOffset, remainder)
	if ok := wa.Memory().Write(offset, bytes); !ok {
		return 0, cln, fmt.Errorf("failed to write string data to WASM memory (offset: %v, size: %v)", offset, size)
	}

	// For debugging:
	// _, _, _ = memoryBlockAtOffset(wa, offset-8, 0, true)

	return offset - 8, cln, nil
}

func stringDataAtOffset(wa wasmMemoryReader, offset uint32) (data []byte, err error) {
	memBlock, words, err := memoryBlockAtOffset(wa, offset, 0, true)
	if err != nil {
		return nil, err
	}

	result, err := stringDataFromMemBlock(memBlock, words)
	// gmlPrintf("GML: handler_memory.go: stringDataAtOffset(offset: %v) = (data=%v, size=%v)", offset, offset+8, len(result))
	return result, err
}

func stringDataFromMemBlock(memBlock []byte, words uint32) (data []byte, err error) {
	if memBlock[4] != StringBlockType {
		return nil, fmt.Errorf("expected MoonBit String block type %v, got %v", StringBlockType, memBlock[4])
	}
	remainderOffset := words*4 + 7
	remainder := uint32(3 - memBlock[remainderOffset]%4)
	size := (words-1)*4 + remainder
	// gmlPrintf("GML: handler_memory.go: stringDataFromMemBlock: memBlockHeader: %+v, memBlock: %+v, words: %v, remainderOffset: %v, remainder: %v, size: %v",
	// 	memBlock[0:8], memBlock, words, remainderOffset, remainder, size)
	if size <= 0 {
		return nil, nil
	}
	if int(size)+8 > len(memBlock) {
		return nil, fmt.Errorf("expected string data size %v, got %v", size, len(memBlock))
	}

	return memBlock[8 : size+8], nil
}
