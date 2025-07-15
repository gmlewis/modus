// -*- compile-command: "NO_COLOR=1 go test -timeout 5s ./..."; -*-

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
	// MoonBit memory block type identifiers (classID values)
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

	// Additional classID values discovered through testing
	BoolByteCharClassID = 96  // Used for Bool?, Byte?, Char?, Int16?, UInt arrays
	Int64DoubleClassID  = 112 // Used for Int64, UInt64, Double arrays
	RefArrayClassID     = 160 // Used for Double?, Float?, Int64?, UInt64? arrays

	// None singleton values for optional types
	NoneSentinelUInt32      = 0xFFFFFFFF // None value for Bool?, Byte?, Char?
	NoneSingletonPointer    = 10248      // None singleton pointer for 64-bit reference types (confirmed by WAT analysis)
	NoneValueInt16          = 32768      // None value for Int16?
	NoneValueInt            = 4          // None value for Int?/UInt? (determined from runtime testing)// Memory layout constants
	MemoryBlockHeaderSize   = 8          // Standard memory block header size
	MemoryBlockHeaderSizeLg = 16         // Extended header size for some arrays
	MinValidMemoryOffset    = 1000       // Minimum valid memory address threshold
	EmptyArrayMarker1       = 4294967295 // Special marker for empty arrays
	EmptyArrayMarker2       = 1610612736 // Special marker for empty arrays

	// Bit manipulation constants
	ClassIDShift           = 24         // Bit shift for classID extraction
	StringLengthMask       = 0x0FFFFFFF // 28-bit mask for string length
	WordsCountMask         = 0x00ffffff // 24-bit mask for word count
	ClassIDMask            = 0xff       // 8-bit mask for classID
	FourByteAlignmentMask  = 3          // Mask for 4-byte alignment
	ByteAlignmentIncrement = 3          // Increment for alignment calculations

	// Array element sizes
	MoonBitBoolSize = 4 // MoonBit Bool is 4 bytes (vs Go's 1 byte)
	MoonBitCharSize = 4 // MoonBit Char is 4 bytes (vs Go's 2 bytes)
	StandardPtrSize = 4 // Standard pointer size in MoonBit
	Int64Size       = 8 // Size of Int64/UInt64 types
	StringCharSize  = 2 // UTF-16 character size in strings
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

	// Try to read the memory block header directly
	memBlockHeader, ok := wa.Memory().Read(offset, MemoryBlockHeaderSize)
	if !ok {
		return nil, 0, 0, fmt.Errorf("failed to read memBlockHeader from WASM memory: (offset: %v, size: %d)", debugShowOffset(offset), MemoryBlockHeaderSize)
	}
	part2 := binary.LittleEndian.Uint32(memBlockHeader[4:8])
	classID = byte(part2 >> ClassIDShift)
	var size uint32

	// For strings, extract the length from the lower 28 bits
	if classID == StringBlockType {
		// String length is in the lower 28 bits (in characters)
		words = part2 & StringLengthMask
		// Size is 8 bytes header + (length * 2) bytes for UTF-16 data, padded to 4-byte boundary
		dataSize := words * StringCharSize
		paddedSize := (dataSize + FourByteAlignmentMask) & ^uint32(FourByteAlignmentMask) // Round up to next 4-byte boundary
		size = MemoryBlockHeaderSize + paddedSize
	} else if classID == TupleBlockType {
		classID = byte(part2 & ClassIDMask)
		words = part2 >> MemoryBlockHeaderSize
		size = uint32(MemoryBlockHeaderSize + words*StandardPtrSize)
	} else {
		words = part2 & WordsCountMask
		size = uint32(MemoryBlockHeaderSize * (2 + (words >> 2)))
	}
	if sizeOverride > 0 {
		size = MemoryBlockHeaderSize + sizeOverride
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
