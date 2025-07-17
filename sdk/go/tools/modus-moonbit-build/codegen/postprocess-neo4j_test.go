// -*- compile-command: "NO_COLOR=1 go test -run ^TestTestablePostProcess_Neo4j ."; -*-

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

func TestTestablePostProcess_Neo4j(t *testing.T) {
	t.Parallel()

	config := &config.Config{
		SourceDir: "../testdata/neo4j-example",
	}

	meta := postProcessTestSetup(t, config)

	body, header := testablePostProcess(meta)

	wg := &postProcessDiffs{
		wantPostProcessBody:   wantNeo4jPostProcessBody,
		gotPostProcessBody:    body.String(),
		wantPostProcessHeader: postProcessHeader,
		gotPostProcessHeader:  header.String(),
	}
	reportPostProcessDiffs(t, "neo4j", wg)
}

var wantNeo4jPostProcessBody = `
///|
pub fn read_map(
  key_type_name_ptr : Int,
  value_type_name_ptr : Int,
  map_ptr : Int
) -> Int64 {
  let key_type_name = ptr2str(key_type_name_ptr + 8)
  let value_type_name = ptr2str(value_type_name_ptr + 8)
  match (key_type_name, value_type_name) {
    ("String", "Json") => read_map_helper_0(map_ptr)
    ("String", "String") => read_map_helper_1(map_ptr)
    _ => 0
  }
}
///|
fn read_map_helper_0(map_ptr : Int) -> Int64 {
  let m : Map[String, Json] = cast(map_ptr)
  let pairs = m.to_array()
  let keys = pairs.map(fn(t) { t.0 })
  let values = pairs.map(fn(t) { t.1 })
  let keys_ptr : Int = cast(keys)
  let values_ptr : Int = cast(values)
  (keys_ptr.to_int64() << 32) | values_ptr.to_int64()
}

///|
fn read_map_helper_1(map_ptr : Int) -> Int64 {
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
    ("String", "Json") => write_map_helper_0(keys_ptr, values_ptr)
    ("String", "String") => write_map_helper_1(keys_ptr, values_ptr)
    _ => 0
  }
}
///|
fn write_map_helper_0(keys_ptr: Int, values_ptr: Int) -> Int {
  let keys : Array[String] = cast(keys_ptr)
  let values : Array[Json] = cast(values_ptr)
  let m : Map[String, Json] = Map::new(capacity=keys.length())
  for i in 0..<keys.length() {
    m[keys[i]] = values[i]
  }
  cast(m)
}

///|
fn write_map_helper_1(keys_ptr: Int, values_ptr: Int) -> Int {
  let keys : Array[String] = cast(keys_ptr)
  let values : Array[String] = cast(values_ptr)
  let m : Map[String, String] = Map::new(capacity=keys.length())
  for i in 0..<keys.length() {
    m[keys[i]] = values[i]
  }
  cast(m)
}
`
