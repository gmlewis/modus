// -*- compile-command: "NO_COLOR=1 go test -run ^TestTestablePostProcess_Runtime ."; -*-

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
		wantPostProcessHeader: postProcessHeader + wantRuntimePostProcessHeader,
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

var wantRuntimePostProcessHeader = `
///|
pub fn zoned_date_time_from_unix_seconds_and_nanos(second : Int64, nanos : Int64) -> @time.ZonedDateTime raise Error {
  let nanosecond = (nanos % 1_000_000_000).to_int()
  @time.unix!(second, nanosecond~)
}

///|
pub fn duration_from_nanos(nanoseconds : Int64) -> @time.Duration raise Error {
  @time.Duration::of!(nanoseconds~)
}
`
