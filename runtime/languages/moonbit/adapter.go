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
		mod:         mod,
		visitedPtrs: make(map[uint32]int),
		fnRealloc:   mod.ExportedFunction("cabi_realloc"), // pub fn cabi_realloc(src_offset : Int, src_size : Int, _dst_alignment : Int, dst_size : Int) -> Int
		fnExtend16:  mod.ExportedFunction("extend16"),     // pub fn extend16(value : Int) -> Int
		fnExtend8:   mod.ExportedFunction("extend8"),      // pub fn extend8(value : Int) -> Int
		fnStore8:    mod.ExportedFunction("store8"),       // pub fn store8(offset : Int, value : Int)
		fnLoad8_u:   mod.ExportedFunction("load8_u"),      // pub fn load8_u(offset : Int) -> Int
		fnLoad8:     mod.ExportedFunction("load8"),        // pub fn load8(offset : Int) -> Int
		fnStore16:   mod.ExportedFunction("store16"),      // pub fn store16(offset : Int, value : Int)
		fnLoad16:    mod.ExportedFunction("load16"),       // pub fn load16(offset : Int) -> Int
		fnLoad16_u:  mod.ExportedFunction("load16_u"),     // pub fn load16_u(offset : Int) -> Int
		fnStore32:   mod.ExportedFunction("store32"),      // pub fn store32(offset : Int, value : Int)
		fnLoad32:    mod.ExportedFunction("load32"),       // pub fn load32(offset : Int) -> Int
		fnStore64:   mod.ExportedFunction("store64"),      // pub fn store64(offset : Int, value : Int64)
		fnLoad64:    mod.ExportedFunction("load64"),       // pub fn load64(offset : Int) -> Int64
		fnStoref32:  mod.ExportedFunction("storef32"),     // pub fn storef32(offset : Int, value : Float)
		fnLoadf32:   mod.ExportedFunction("loadf32"),      // pub fn loadf32(offset : Int) -> Float
		fnStoref64:  mod.ExportedFunction("storef64"),     // pub fn storef64(offset : Int, value : Double)
		fnLoadf64:   mod.ExportedFunction("loadf64"),      // pub fn loadf64(offset : Int) -> Double
		// fnF32_to_i32:       mod.ExportedFunction("f32_to_i32"),       // pub fn f32_to_i32(value : Float) -> Int
		// fnF32_to_i64:       mod.ExportedFunction("f32_to_i64"),       // pub fn f32_to_i64(value : Float) -> Int64
		fnMalloc:           mod.ExportedFunction("malloc"),           // pub fn malloc(size : Int) -> Int
		fnFree:             mod.ExportedFunction("free"),             // pub fn free(position : Int)
		fnCopy:             mod.ExportedFunction("copy"),             // pub fn copy(dest : Int, src : Int) -> Unit
		fnStr2ptr:          mod.ExportedFunction("str2ptr"),          // pub fn str2ptr(str : String) -> Int
		fnPtr2str:          mod.ExportedFunction("ptr2str"),          // pub fn ptr2str(ptr : Int) -> String
		fnBytes2ptr:        mod.ExportedFunction("bytes2ptr"),        // pub fn bytes2ptr(bytes : Bytes) -> Int
		fnPtr2bytes:        mod.ExportedFunction("ptr2bytes"),        // pub fn ptr2bytes(ptr : Int) -> Bytes
		fnUint_array2ptr:   mod.ExportedFunction("uint_array2ptr"),   // pub fn uint_array2ptr(array : FixedArray[UInt]) -> Int
		fnUint64_array2ptr: mod.ExportedFunction("uint64_array2ptr"), // pub fn uint64_array2ptr(array : FixedArray[UInt64]) -> Int
		fnInt_array2ptr:    mod.ExportedFunction("int_array2ptr"),    // pub fn int_array2ptr(array : FixedArray[Int]) -> Int
		fnInt64_array2ptr:  mod.ExportedFunction("int64_array2ptr"),  // pub fn int64_array2ptr(array : FixedArray[Int64]) -> Int
		fnFloat_array2ptr:  mod.ExportedFunction("float_array2ptr"),  // pub fn float_array2ptr(array : FixedArray[Float]) -> Int
		fnDouble_array2ptr: mod.ExportedFunction("double_array2ptr"), // pub fn double_array2ptr(array : FixedArray[Double]) -> Int
		fnPtr2uint_array:   mod.ExportedFunction("ptr2uint_array"),   // pub fn ptr2uint_array(ptr : Int) -> FixedArray[UInt]
		fnPtr2int_array:    mod.ExportedFunction("ptr2int_array"),    // pub fn ptr2int_array(ptr : Int) -> FixedArray[Int]
		fnPtr2float_array:  mod.ExportedFunction("ptr2float_array"),  // pub fn ptr2float_array(ptr : Int) -> FixedArray[Float]
		fnPtr2uint64_array: mod.ExportedFunction("ptr2uint64_array"), // pub fn ptr2uint64_array(ptr : Int) -> FixedArray[UInt64]
		fnPtr2int64_array:  mod.ExportedFunction("ptr2int64_array"),  // pub fn ptr2int64_array(ptr : Int) -> FixedArray[Int64]
		fnPtr2double_array: mod.ExportedFunction("ptr2double_array"), // pub fn ptr2double_array(ptr : Int) -> FixedArray[Double]
		// pub fn zoned_date_time_from_unix_seconds_and_nanos(second : Int64, nanos : Int64) -> @time.ZonedDateTime!Error
		fnZonedDateTimeFromUnixSecondsAndNanos: mod.ExportedFunction("zoned_date_time_from_unix_seconds_and_nanos"),
		// pub fn duration_from_nanos(nanoseconds : Int64) -> @time.Duration!Error
		fnDurationFromNanos: mod.ExportedFunction("duration_from_nanos"),
		// pub fn write_map(key_type_name_ptr : Int, value_type_name_ptr : Int, key_ptr : Int, value_ptr : Int) -> Int
		fnWriteMap: mod.ExportedFunction("write_map"),
	}
}

type wasmAdapter struct {
	mod         wasm.Module
	visitedPtrs map[uint32]int
	fnRealloc   wasm.Function
	fnExtend16  wasm.Function
	fnExtend8   wasm.Function
	fnStore8    wasm.Function
	fnLoad8_u   wasm.Function
	fnLoad8     wasm.Function
	fnStore16   wasm.Function
	fnLoad16    wasm.Function
	fnLoad16_u  wasm.Function
	fnStore32   wasm.Function
	fnLoad32    wasm.Function
	fnStore64   wasm.Function
	fnLoad64    wasm.Function
	fnStoref32  wasm.Function
	fnLoadf32   wasm.Function
	fnStoref64  wasm.Function
	fnLoadf64   wasm.Function
	// fnF32_to_i32       wasm.Function
	// fnF32_to_i64       wasm.Function
	fnMalloc           wasm.Function
	fnFree             wasm.Function
	fnCopy             wasm.Function
	fnStr2ptr          wasm.Function
	fnPtr2str          wasm.Function
	fnBytes2ptr        wasm.Function
	fnPtr2bytes        wasm.Function
	fnUint_array2ptr   wasm.Function
	fnUint64_array2ptr wasm.Function
	fnInt_array2ptr    wasm.Function
	fnInt64_array2ptr  wasm.Function
	fnFloat_array2ptr  wasm.Function
	fnDouble_array2ptr wasm.Function
	fnPtr2uint_array   wasm.Function
	fnPtr2int_array    wasm.Function
	fnPtr2float_array  wasm.Function
	fnPtr2uint64_array wasm.Function
	fnPtr2int64_array  wasm.Function
	fnPtr2double_array wasm.Function
	// used to convert Go time.Time to MoonBit @time.ZonedDateTime
	fnZonedDateTimeFromUnixSecondsAndNanos wasm.Function
	// used to convert Go time.Duration to MoonBit @time.Duration
	fnDurationFromNanos wasm.Function
	fnWriteMap          wasm.Function
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
	gmlPrintf("GML: adapter.go: wasmAdapter.PreInvoke")
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

	gmlPrintf("GML: wasmAdapter.allocateWasmMemory(size: %v, classID: %v): offset: %v", size, classID, offset)
	return offset, nil
}

// Free memory within the MoonBit module.
// TODO: Figure out _WHEN_ to clean up the memory so that results can be passed to other wasm functions.
// func (wa *wasmAdapter) freeWasmMemory(ctx context.Context, offset uint32) error {
// 	gmlPrintf("GML: wasmAdapter.freeWasmMemory(offset: %v)", offset)
// 	res, err := wa.fnRealloc.Call(ctx, uint64(offset), 0, 0, 0)
// 	if err != nil {
// 		return fmt.Errorf("failed to free WASM memory (offset: %v): %w", offset, err)
// 	}

// 	ptr := uint32(res[0])
// 	if ptr != 0 {
// 		return fmt.Errorf("failed to free WASM memory: non-zero result: %v", ptr)
// 	}

// 	return nil
// }
