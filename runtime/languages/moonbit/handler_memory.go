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
	"encoding/binary"
	"errors"
	"fmt"
	"log"

	"github.com/gmlewis/modus/runtime/langsupport"
)

var moonBitBlockType = map[byte]string{
	0:   "Tuple",
	241: "FixedArray[Int]", // Also FixedArray[UInt]
	243: "String",
	246: "FixedArray[Byte]",
}

func memoryBlockAtOffset(wa langsupport.WasmAdapter, offset uint32, dbgHackToRemove ...bool) (data []byte, words uint32, err error) {
	if offset == 0 {
		log.Printf("GML: handler_memory.go: memoryBlockAtOffset(offset: %v) = (data=0, size=0)", offset)
		return nil, 0, nil
	}

	memBlockHeader, ok := wa.Memory().Read(offset, uint32(8))
	if !ok {
		return nil, 0, fmt.Errorf("failed to read memBlockHeader from WASM memory: (offset: %v, size: 8)", offset)
	}
	words = binary.LittleEndian.Uint32(memBlockHeader[4:8]) >> 8
	size := uint32(8 + words*4)
	memBlock, ok := wa.Memory().Read(offset, size)
	if !ok {
		return nil, 0, fmt.Errorf("failed to read memBlock from WASM memory: (offset: %v, size: %v)", offset, size)
	}
	if len(dbgHackToRemove) > 0 {
		moonBitType := memBlockHeader[4]
		moonBitTypeName := moonBitBlockType[moonBitType]
		if moonBitTypeName == "String" {
			data, _ := stringDataFromMemBlock(memBlock, words) // ignore errors during debugging
			s, _ := doReadString(data)
			log.Printf("GML: handler_memory.go: memoryBlockAtOffset(offset: %v, size: %v=8+words*4), moonBitType=%v(%v), words=%v, memBlock=%+v = '%v'",
				offset, size, moonBitType, moonBitTypeName, words, memBlock, s)
		} else {
			log.Printf("GML: handler_memory.go: memoryBlockAtOffset(offset: %v, size: %v=8+words*4), moonBitType=%v(%v), words=%v, memBlock=%+v",
				offset, size, moonBitType, moonBitTypeName, words, memBlock)
		}
	}
	return memBlock, words, nil
}

func writeMemoryBlockHeader(wa langsupport.WasmAdapter, data, size, offset uint32) error {
	log.Printf("GML: handler_memory.go: writeMemoryBlockHeader(data: %v, size: %v, offset: %v)", data, size, offset)

	val := uint64(size)<<32 | uint64(data)
	if ok := wa.Memory().WriteUint64Le(offset, val); !ok {
		return errors.New("failed to write string header to WASM memory")
	}

	return nil
}
