// -*- compile-command: "go test -run ^TestTestablePostProcess_Runtime ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package codegen

import (
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
)

func TestTestablePostProcess_Runtime(t *testing.T) {
	t.Parallel()

	config := &config.Config{
		SourceDir: "../testdata/runtime-testdata",
	}

	meta := postProcessTestSetup(t, config)

	body, header := testablePostProcess(meta)

	wg := &postProcessDiffs{
		wantPostProcessBody:   wantRuntimePostProcessBody,
		gotPostProcessBody:    body.String(),
		wantPostProcessHeader: wantRuntimePostProcessHeader,
		gotPostProcessHeader:  header.String(),
	}
	reportPostProcessDiffs(t, "runtime", wg)
}

var wantRuntimePostProcessBody = `
///|
pub fn read_map(
  key_type_name_ptr : Int,
  value_type_name_ptr : Int,
  map_ptr : Int
) -> Int64 {
  let key_type_name = ptr2str(key_type_name_ptr + 8)
  let value_type_name = ptr2str(value_type_name_ptr + 8)
  match (key_type_name, value_type_name) {
    ("Int", "Double") => read_map_helper_0(map_ptr)
    ("Int", "Float") => read_map_helper_1(map_ptr)
    ("String", "HttpHeader?") => read_map_helper_2(map_ptr)
    ("String", "String") => read_map_helper_3(map_ptr)
    _ => 0
  }
}
///|
fn read_map_helper_0(map_ptr : Int) -> Int64 {
  let m : Map[Int, Double] = cast(map_ptr)
  let pairs = m.to_array()
  let keys = pairs.map(fn(t) { t.0 })
  let values = pairs.map(fn(t) { t.1 })
  let keys_ptr : Int = cast(keys)
  let values_ptr : Int = cast(values)
  (keys_ptr.to_int64() << 32) | values_ptr.to_int64()
}

///|
fn read_map_helper_1(map_ptr : Int) -> Int64 {
  let m : Map[Int, Float] = cast(map_ptr)
  let pairs = m.to_array()
  let keys = pairs.map(fn(t) { t.0 })
  let values = pairs.map(fn(t) { t.1 })
  let keys_ptr : Int = cast(keys)
  let values_ptr : Int = cast(values)
  (keys_ptr.to_int64() << 32) | values_ptr.to_int64()
}

///|
fn read_map_helper_2(map_ptr : Int) -> Int64 {
  let m : Map[String, HttpHeader?] = cast(map_ptr)
  let pairs = m.to_array()
  let keys = pairs.map(fn(t) { t.0 })
  let values = pairs.map(fn(t) { t.1 })
  let keys_ptr : Int = cast(keys)
  let values_ptr : Int = cast(values)
  (keys_ptr.to_int64() << 32) | values_ptr.to_int64()
}

///|
fn read_map_helper_3(map_ptr : Int) -> Int64 {
  let m : Map[String, String] = cast(map_ptr)
  let pairs = m.to_array()
  let keys = pairs.map(fn(t) { t.0 })
  let values = pairs.map(fn(t) { t.1 })
  let keys_ptr : Int = cast(keys)
  let values_ptr : Int = cast(values)
  (keys_ptr.to_int64() << 32) | values_ptr.to_int64()
}

///|
pub fn write_map(key_type_name_ptr : Int, value_type_name_ptr : Int, keys_ptr : Int, values_ptr : Int) -> Int {
  let key_type_name = ptr2str(key_type_name_ptr + 8)
  let value_type_name = ptr2str(value_type_name_ptr + 8)
  match (key_type_name, value_type_name) {
    ("Int", "Double") => write_map_helper_0(keys_ptr, values_ptr)
    ("Int", "Float") => write_map_helper_1(keys_ptr, values_ptr)
    ("String", "HttpHeader?") => write_map_helper_2(keys_ptr, values_ptr)
    ("String", "String") => write_map_helper_3(keys_ptr, values_ptr)
    _ => 0
  }
}
///|
fn write_map_helper_0(keys_ptr: Int, values_ptr: Int) -> Int {
  let keys : Array[Int] = cast(keys_ptr)
  let values : Array[Double] = cast(values_ptr)
  let m : Map[Int, Double] = Map::new(capacity=keys.length())
  for i in 0..<keys.length() {
    m[keys[i]] = values[i]
  }
  cast(m)
}

///|
fn write_map_helper_1(keys_ptr: Int, values_ptr: Int) -> Int {
  let keys : Array[Int] = cast(keys_ptr)
  let values : Array[Float] = cast(values_ptr)
  let m : Map[Int, Float] = Map::new(capacity=keys.length())
  for i in 0..<keys.length() {
    m[keys[i]] = values[i]
  }
  cast(m)
}

///|
fn write_map_helper_2(keys_ptr: Int, values_ptr: Int) -> Int {
  let keys : Array[String] = cast(keys_ptr)
  let values : Array[HttpHeader?] = cast(values_ptr)
  let m : Map[String, HttpHeader?] = Map::new(capacity=keys.length())
  for i in 0..<keys.length() {
    m[keys[i]] = values[i]
  }
  cast(m)
}

///|
fn write_map_helper_3(keys_ptr: Int, values_ptr: Int) -> Int {
  let keys : Array[String] = cast(keys_ptr)
  let values : Array[String] = cast(values_ptr)
  let m : Map[String, String] = Map::new(capacity=keys.length())
  for i in 0..<keys.length() {
    m[keys[i]] = values[i]
  }
  cast(m)
}
`

var wantRuntimePostProcessHeader = `// Code generated by modus-moonbit-build. DO NOT EDIT.

///|
pub fn cabi_realloc(
  src_offset : Int,
  src_size : Int,
  _dst_alignment : Int,
  dst_size : Int
) -> Int {
  // malloc
  if src_offset == 0 && src_size == 0 {
    return malloc(dst_size)
  }
  // free
  if dst_size <= 0 {
    free(src_offset)
    return 0
  }
  // realloc
  let dst = malloc(dst_size)
  copy(dst, src_offset)
  free(src_offset)
  dst
}

// Generated by wit-bindgen 0.36.0. DO NOT EDIT!

// ///|
// pub extern "wasm" fn extend16(value : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.extend16_s)

// ///|
// pub extern "wasm" fn extend8(value : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.extend8_s)

///|
pub extern "wasm" fn store8(offset : Int, value : Int) =
  #|(func (param i32) (param i32) local.get 0 local.get 1 i32.store8)

// ///|
// pub extern "wasm" fn load8_u(offset : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.load8_u)

// ///|
// pub extern "wasm" fn load8(offset : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.load8_s)

// ///|
// pub extern "wasm" fn store16(offset : Int, value : Int) =
//   #|(func (param i32) (param i32) local.get 0 local.get 1 i32.store16)

// ///|
// pub extern "wasm" fn load16(offset : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.load16_s)

// ///|
// pub extern "wasm" fn load16_u(offset : Int) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.load16_u)

///|
pub extern "wasm" fn store32(offset : Int, value : Int) =
  #|(func (param i32) (param i32) local.get 0 local.get 1 i32.store)

///|
pub extern "wasm" fn load32(offset : Int) -> Int =
  #|(func (param i32) (result i32) local.get 0 i32.load)

// ///|
// pub extern "wasm" fn store64(offset : Int, value : Int64) =
//   #|(func (param i32) (param i64) local.get 0 local.get 1 i64.store)

// ///|
// pub extern "wasm" fn load64(offset : Int) -> Int64 =
//   #|(func (param i32) (result i64) local.get 0 i64.load)

// ///|
// pub extern "wasm" fn storef32(offset : Int, value : Float) =
//   #|(func (param i32) (param f32) local.get 0 local.get 1 f32.store)

// ///|
// pub extern "wasm" fn loadf32(offset : Int) -> Float =
//   #|(func (param i32) (result f32) local.get 0 f32.load)

// ///|
// pub extern "wasm" fn storef64(offset : Int, value : Double) =
//   #|(func (param i32) (param f64) local.get 0 local.get 1 f64.store)

// ///|
// pub extern "wasm" fn loadf64(offset : Int) -> Double =
//   #|(func (param i32) (result f64) local.get 0 f64.load)

// ///|
// // pub extern "wasm" fn f32_to_i32(value : Float) -> Int =
// //   #|(func (param f32) (result i32) local.get 0 f32.convert_i32_s)

// ///|
// // pub extern "wasm" fn f32_to_i64(value : Float) -> Int64 =
// //   #|(func (param f32) (result i64) local.get 0 f32.convert_i64_s)

///|
extern "wasm" fn malloc_inline(size : Int) -> Int =
  #|(func (param i32) (result i32) local.get 0 call $moonbit.malloc)

///|
pub fn malloc(size : Int) -> Int {
  let words = size / 4 + 1
  let address = malloc_inline(8 + words * 4)
  store32(address, 1)
  store32(address + 4, (words << 8) | 246)
  store8(address + words * 4 + 7, 3 - size % 4)
  address + 8
}

///|
pub extern "wasm" fn free(position : Int) =
  #|(func (param i32) local.get 0 i32.const 8 i32.sub call $moonbit.decref)

///|
pub fn copy(dest : Int, src : Int) -> Unit {
  let src = src - 8
  let dest = dest - 8
  let src_len = load32(src + 4) & 0xFFFFFF
  let dest_len = load32(dest + 4) & 0xFFFFFF
  let min = if src_len < dest_len { src_len } else { dest_len }
  copy_inline(dest, src, min)
}

///|
extern "wasm" fn copy_inline(dest : Int, src : Int, len : Int) =
  #|(func (param i32) (param i32) (param i32) local.get 0 local.get 1 local.get 2 memory.copy)

// ///|
// pub extern "wasm" fn str2ptr(str : String) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

///|
pub extern "wasm" fn ptr2str(ptr : Int) -> String =
  #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 243 i32.store8 local.get 0 i32.const 8 i32.sub)

// ///|
// pub extern "wasm" fn bytes2ptr(bytes : Bytes) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn ptr2bytes(ptr : Int) -> Bytes =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

// ///|
// pub extern "wasm" fn uint_array2ptr(array : FixedArray[UInt]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn uint64_array2ptr(array : FixedArray[UInt64]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn int_array2ptr(array : FixedArray[Int]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn int64_array2ptr(array : FixedArray[Int64]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn float_array2ptr(array : FixedArray[Float]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn double_array2ptr(array : FixedArray[Double]) -> Int =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.add)

// ///|
// pub extern "wasm" fn ptr2uint_array(ptr : Int) -> FixedArray[UInt] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)

// ///|
// pub extern "wasm" fn ptr2int_array(ptr : Int) -> FixedArray[Int] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)

// ///|
// pub extern "wasm" fn ptr2float_array(ptr : Int) -> FixedArray[Float] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 4 i32.sub i32.const 241 i32.store8 local.get 0 i32.const 8 i32.sub)

// ///|
// extern "wasm" fn ptr2uint64_array_ffi(ptr : Int) -> FixedArray[UInt64] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

// ///|
// pub fn ptr2uint64_array(ptr : Int) -> FixedArray[UInt64] {
//   set_64_header_ffi(ptr - 4)
//   ptr2uint64_array_ffi(ptr)
// }

// ///|
// extern "wasm" fn ptr2int64_array_ffi(ptr : Int) -> FixedArray[Int64] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

// ///|
// pub fn ptr2int64_array(ptr : Int) -> FixedArray[Int64] {
//   set_64_header_ffi(ptr - 4)
//   ptr2int64_array_ffi(ptr)
// }

// ///|
// extern "wasm" fn ptr2double_array_ffi(ptr : Int) -> FixedArray[Double] =
//   #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

// ///|
// pub fn ptr2double_array(ptr : Int) -> FixedArray[Double] {
//   set_64_header_ffi(ptr - 4)
//   ptr2double_array_ffi(ptr)
// }

// ///|
// fn set_64_header_ffi(offset : Int) -> Unit {
//   let len = load32(offset)
//   store32(offset, len >> 1)
//   store8(offset, 241)
// }

// ///|
// pub(open) trait Any {}

// ///|
// pub(all) struct Cleanup {
//   address : Int
//   size : Int
//   align : Int
// }

pub fn ptr_to_none() -> Int {
  let val : Double? = None
  cast(val)
}

///|
fn cast[A, B](a : A) -> B = "%identity"

///|
pub fn zoned_date_time_from_unix_seconds_and_nanos(second : Int64, nanos : Int64) -> @time.ZonedDateTime!Error {
  let nanosecond = (nanos % 1_000_000_000).to_int()
  @time.unix!(second, nanosecond~)
}

///|
pub fn duration_from_nanos(nanoseconds : Int64) -> @time.Duration!Error {
  @time.Duration::of!(nanoseconds~)
}
`
