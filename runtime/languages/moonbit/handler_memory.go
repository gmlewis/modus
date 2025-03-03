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

	"github.com/gmlewis/modus/runtime/utils"
	wasm "github.com/tetratelabs/wazero/api"
)

// Ptr is a helper routine that allocates a new T value
// to store v and returns a pointer to it.
func Ptr[T any](v T) *T {
	return &v
}

var moonBitBlockType = map[byte]string{
	0:   "Tuple",
	241: "FixedArray[Primitive]",
	242: "FixedArray[String]", // Also FixedArray[Int64?], FixedArray[UInt64?]
	243: "String",
	246: "FixedArray[Byte]",
}

const (
	FixedArrayPrimitiveBlockType = 241
	PtrArrayBlockType            = 242
	StringBlockType              = 243
	FixedArrayByteBlockType      = 246
	TupleBlockType               = 0
	ZonedDateTimeBlockType       = 3
	ZoneBlockType                = 3
	ZoneOffsetBlockType          = 4
	PlainDateTimeBlockType       = 2
	PlainDateBlockType           = 3
	PlainTimeBlockType           = 4
	OptionBlockType              = 1 // TODO
)

// For testing purposes:
type wasmMemoryReader interface {
	Memory() wasm.Memory
}

type wasmMemoryWriter interface {
	allocateAndPinMemory(ctx context.Context, size, blockType uint32) (uint32, utils.Cleaner, error)
	Memory() wasm.Memory
}

func memoryBlockAtOffset(wa wasmMemoryReader, offset, sizeOverride uint32, dbgHackToRemove ...bool) (data []byte, classID byte, words uint32, err error) {
	if offset == 0 {
		gmlPrintf("  // memoryBlockAtOffset(offset: 0) = (data=0, size=0)")
		return nil, 0, 0, nil
	}

	memBlockHeader, ok := wa.Memory().Read(offset, uint32(8))
	if !ok {
		return nil, 0, 0, fmt.Errorf("failed to read memBlockHeader from WASM memory: (offset: %v, size: 8)", debugShowOffset(offset))
	}
	part2 := binary.LittleEndian.Uint32(memBlockHeader[4:8])
	classID = byte(part2 & 0xff)
	words = part2 >> 8
	size := uint32(8 + words*4)
	if sizeOverride > 0 {
		// if sizeOverride+8 == size {
		// 	sizeOverride = 0 // for debugging - no need for override
		// } else {
		size = 8 + sizeOverride
		// }
	}
	memBlock, ok := wa.Memory().Read(offset, size)
	if !ok {
		return nil, 0, 0, fmt.Errorf("failed to read memBlock from WASM memory: (offset: %v, size: %v)", debugShowOffset(offset), size)
	}
	if len(dbgHackToRemove) > 0 { // && sizeOverride == 0 {
		classIDName := moonBitBlockType[classID]
		if classIDName == "String" {
			data, err := stringDataFromMemBlock(memBlock, words)
			if err != nil {
				log.Printf("DEBUGGING ERROR: handler_memory.go: memoryBlockAtOffset(offset: %v, size: %v=8+words*4), classID=%v(%v), words=%v, memBlock=%+v, err=%v", debugShowOffset(offset), size, classID, classIDName, words, memBlock, err)
			}
			s, err := doReadString(data)
			if err != nil {
				log.Printf("DEBUGGING ERROR: handler_memory.go: memoryBlockAtOffset(offset: %v, size: %v=8+words*4), classID=%v(%v), words=%v, memBlock=%+v, err=%v", debugShowOffset(offset), size, classID, classIDName, words, memBlock, err)
			}
			gmlPrintf("  // memoryBlockAtOffset(offset: %v, size: %v=8+words*4), classID=%v(%v), words=%v, memBlock=%+v = '%v'",
				debugShowOffset(offset), size, classID, classIDName, words, memBlock, s)
		} else {
			gmlPrintf("  // memoryBlockAtOffset(offset: %v, size: %v=8+words*4), classID=%v(%v), words=%v, memBlock=%+v",
				debugShowOffset(offset), size, classID, classIDName, words, memBlock)
		}
	}
	return memBlock, classID, words, nil
}

// func writeMemoryBlockHeader(wa langsupport.WasmAdapter, data, size, offset uint32) error {
// 	gmlPrintf("GML: handler_memory.go: writeMemoryBlockHeader(data: %v, size: %v, offset: %v)", data, size, debugShowOffset(offset))

// 	val := uint64(size)<<32 | uint64(data)
// 	if ok := wa.Memory().WriteUint64Le(offset, val); !ok {
// 		return errors.New("failed to write string header to WASM memory")
// 	}

// 	return nil
// }

func debugShowOffset(offset uint32) string {
	return fmt.Sprintf("%v=0x%08X=[%v %v %v %v]", offset, offset, byte(offset), byte(offset>>8), byte(offset>>16), byte(offset>>24))
}
