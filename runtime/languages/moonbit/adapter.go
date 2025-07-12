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
	"errors"
	"fmt"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"

	wasm "github.com/tetratelabs/wazero/api"
)

func NewWasmAdapter(mod wasm.Module) langsupport.WasmAdapter {
	return &wasmAdapter{
		mod:                                    mod,
		visitedPtrs:                            make(map[uint32]int),
		fnRealloc:                              mod.ExportedFunction("cabi_realloc"),
		fnPtr2str:                              mod.ExportedFunction("ptr2str"),
		fnZonedDateTimeFromUnixSecondsAndNanos: mod.ExportedFunction("zoned_date_time_from_unix_seconds_and_nanos"),
		fnDurationFromNanos:                    mod.ExportedFunction("duration_from_nanos"),
		fnReadMap:                              mod.ExportedFunction("read_map"),
		fnWriteMap:                             mod.ExportedFunction("write_map"),
		fnPtrToNone:                            mod.ExportedFunction("ptr_to_none"),
	}
}

type wasmAdapter struct {
	mod         wasm.Module
	visitedPtrs map[uint32]int
	fnRealloc   wasm.Function
	fnPtr2str   wasm.Function
	// used to convert Go time.Time to MoonBit @time.ZonedDateTime
	fnZonedDateTimeFromUnixSecondsAndNanos wasm.Function
	// used to convert Go time.Duration to MoonBit @time.Duration
	fnDurationFromNanos wasm.Function
	fnReadMap           wasm.Function
	fnWriteMap          wasm.Function
	fnPtrToNone         wasm.Function
}

func (*wasmAdapter) TypeInfo() langsupport.LanguageTypeInfo {
	return _langTypeInfo
}

func (wa *wasmAdapter) Memory() wasm.Memory {
	return wa.mod.Memory()
}

func (wa *wasmAdapter) GetFunction(name string) wasm.Function {
	return wa.mod.ExportedFunction(name)
}

func (wa *wasmAdapter) PreInvoke(ctx context.Context, plan langsupport.ExecutionPlan) error {
	return nil
}

func (wa *wasmAdapter) AllocateMemory(ctx context.Context, size uint32) (uint32, utils.Cleaner, error) {
	return wa.allocateAndPinMemory(ctx, size, 1)
}

// Allocate and pin memory within the MoonBit module.
// The cleaner returned will unpin the memory when invoked.
func (wa *wasmAdapter) allocateAndPinMemory(ctx context.Context, size, classID uint32) (uint32, utils.Cleaner, error) {
	ptr, err := wa.allocateWasmMemory(ctx, size, classID)
	if err != nil {
		return 0, nil, err
	}

	cln := utils.NewCleanerN(1)
	// TODO: Figure out _WHEN_ to clean up the memory so that results can be passed to other wasm functions.
	// cln.AddCleanup(func() error {
	// 	return wa.freeWasmMemory(ctx, ptr)
	// })

	return ptr, cln, nil
}

// Allocate memory within the MoonBit module.
func (wa *wasmAdapter) allocateWasmMemory(ctx context.Context, size, classID uint32) (uint32, error) {
	res, err := wa.fnRealloc.Call(ctx, 0, 0, 0, uint64(size))
	if err != nil {
		return 0, fmt.Errorf("failed to allocate WASM memory (size: %v, id: %v): %w", size, classID, err)
	}

	offset := uint32(res[0])
	if offset == 0 {
		return 0, errors.New("failed to allocate WASM memory")
	}

	refCount := uint32(1)
	var memType uint32

	// For strings, use the correct format: classID in upper 4 bits, length in lower 28 bits
	if classID == 80 { // StringBlockType
		// Size parameter represents the string length in characters for strings
		// Put classID in upper 4 bits (bits 28-31) and length in lower 28 bits (bits 0-27)
		memType = (classID << 24) | (size & 0x0FFFFFFF)
	} else {
		// For other types, use the original format
		memType = ((size / 4) << 8) | classID
	}

	wa.Memory().WriteUint32Le(offset-8, refCount)
	wa.Memory().WriteUint32Le(offset-4, memType)

	return offset, nil
}

// Free memory within the MoonBit module.
// TODO: Figure out _WHEN_ to clean up the memory so that results can be passed to other wasm functions.
// func (wa *wasmAdapter) freeWasmMemory(ctx context.Context, offset uint32) error {
// 	res, err := wa.fnRealloc.Call(ctx, uint64(offset), 0, 0, 0)
// 	if err != nil {
// 		return fmt.Errorf("failed to free WASM memory (offset: %v): %w", offset, err)
// 	}
//
// 	ptr := uint32(res[0])
// 	if ptr != 0 {
// 		return fmt.Errorf("failed to free WASM memory: non-zero result: %v", ptr)
// 	}
//
// 	return nil
// }
