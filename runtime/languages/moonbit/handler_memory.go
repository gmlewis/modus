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

func memoryBlockAtOffset(wa wasmMemoryReader, offset, sizeOverride uint32) (data []byte, classID byte, words uint32, err error) {
	if offset == 0 {
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
