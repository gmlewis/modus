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

	"github.com/gmlewis/modus/runtime/utils"
	wasm "github.com/tetratelabs/wazero/api"
)

// Ptr is a helper routine that allocates a new T value
// to store v and returns a pointer to it.
func Ptr[T any](v T) *T {
	return &v
}

const (
	// TODO: Fix these:
	FixedArrayPrimitiveBlockType = 241
	PtrArrayBlockType            = 242
	StringBlockType              = 80
	FixedArrayByteBlockType      = 246
	TupleBlockType               = 0
	ZonedDateTimeBlockType       = 3
	ZoneBlockType                = 3
	ZoneOffsetBlockType          = 4
	PlainDateTimeBlockType       = 2
	PlainDateBlockType           = 3
	PlainTimeBlockType           = 4
	OptionBlockType              = 1
)

// For testing purposes:
type wasmMemoryReader interface {
	Memory() wasm.Memory
}

type wasmMemoryWriter interface {
	allocateAndPinMemory(ctx context.Context, size, blockType uint32) (uint32, utils.Cleaner, error)
	Memory() wasm.Memory
}

func memoryBlockAtOffset(wa wasmMemoryReader, offset, sizeOverride uint32) (data []byte, classID byte, words uint32, err error) {
	if offset == 0 {
		return nil, 0, 0, nil
	}

	// Handle None singleton pointer for optional types
	if offset == 10248 {
		// This is the None singleton - return a special marker that can be detected
		// The None singleton has structure [255 255 255 255] [0 0 0 0]
		// Return mock data that will be interpreted as None
		return []byte{255, 255, 255, 255, 0, 0, 0, 0}, 0, 0, nil
	}

	// Check for invalid small offsets that suggest incorrect function return handling
	if offset < 1000 {
		return nil, 0, 0, fmt.Errorf("invalid memory offset %d: function may be returning direct values instead of pointers (check function metadata/compilation)", offset)
	}

	memBlockHeader, ok := wa.Memory().Read(offset, uint32(8))
	if !ok {
		return nil, 0, 0, fmt.Errorf("failed to read memBlockHeader from WASM memory: (offset: %v, size: 8)", debugShowOffset(offset))
	}
	part2 := binary.LittleEndian.Uint32(memBlockHeader[4:8])
	classID = byte(part2 >> 24)
	var size uint32

	// For strings, extract the length from the lower 28 bits
	if classID == StringBlockType {
		// String length is in the lower 28 bits (in characters)
		words = part2 & 0x0FFFFFFF
		// Size is 8 bytes header + (length * 2) bytes for UTF-16 data, padded to 4-byte boundary
		dataSize := words * 2
		paddedSize := (dataSize + 3) & ^uint32(3) // Round up to next 4-byte boundary
		size = 8 + paddedSize
	} else if classID == 0 {
		classID = byte(part2 & 0xff)
		words = part2 >> 8
		size = uint32(8 + words*4)
	} else {
		words = part2 & 0x00ffffff
		size = uint32(8 * (2 + (words >> 2)))
	}
	if sizeOverride > 0 {
		size = 8 + sizeOverride
	}

	memBlock, ok := wa.Memory().Read(offset, size)
	if !ok {
		return nil, 0, 0, fmt.Errorf("failed to read memBlock from WASM memory: (offset: %v, size: %v)", debugShowOffset(offset), size)
	}
	return memBlock, classID, words, nil
}

func debugShowOffset(offset uint32) string {
	return fmt.Sprintf("%v=0x%08X=[%v %v %v %v]", offset, offset, byte(offset), byte(offset>>8), byte(offset>>16), byte(offset>>24))
}
