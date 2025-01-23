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

func memoryBlockAtOffset(wa langsupport.WasmAdapter, offset uint32) (data, size uint32, err error) {
	if offset == 0 {
		log.Printf("GML: handler_memory.go: memoryBlockAtOffset(offset: %v) = (data=0, size=0)", offset)
		return 0, 0, nil
	}

	memBlockHeader, ok := wa.Memory().Read(offset, uint32(8))
	if !ok {
		return 0, 0, fmt.Errorf("failed to read memBlockHeader from WASM memory: (offset: %v, size: 8)", offset)
	}
	words := binary.LittleEndian.Uint32(memBlockHeader[4:8]) >> 8
	size = uint32(8 + words*4)
	memBlock, ok := wa.Memory().Read(offset, size)
	if !ok {
		return 0, 0, fmt.Errorf("failed to read memBlock from WASM memory: (offset: %v, size: %v)", offset, size)
	}
	remainderOffset := words*4 + 7
	remainder := uint32(3 - memBlock[remainderOffset]%4)
	size = (words-1)*4 + remainder
	log.Printf("GML: handler_memory.go: memoryBlockAtOffset: memBlock: %+v, words: %v, remainderOffset: %v, remainder: %v, size: %v",
		memBlock, words, remainderOffset, remainder, size)

	log.Printf("GML: handler_memory.go: memoryBlockAtOffset(offset: %v) = (data=%v, size=%v)", offset, offset+8, size)
	return offset + 8, size, nil
}

func writeMemoryBlockHeader(wa langsupport.WasmAdapter, data, size, offset uint32) error {
	log.Printf("GML: handler_memory.go: writeMemoryBlockHeader(data: %v, size: %v, offset: %v)", data, size, offset)

	val := uint64(size)<<32 | uint64(data)
	if ok := wa.Memory().WriteUint64Le(offset, val); !ok {
		return errors.New("failed to write string header to WASM memory")
	}

	return nil
}
