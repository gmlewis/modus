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
	"github.com/google/go-cmp/cmp"
)

func TestGenBuffers(t *testing.T) {
	t.Parallel()

	config := &config.Config{
		SourceDir: "testdata",
	}
	pkg, err := getMainPackage(config.SourceDir)
	if err != nil {
		t.Fatalf("getMainPackage failed: %v", err)
	}

	functions := getFunctionsNeedingWrappers(pkg)
	imports := getRequiredImports(functions)

	body, header, moonPkgJSON, err := genBuffers(pkg, imports, functions)
	if err != nil {
		t.Fatalf("genBuffers failed: %v", err)
	}

	if diff := cmp.Diff(wantBody, body.String()); diff != "" {
		t.Errorf("body mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantHeader, header.String()); diff != "" {
		t.Errorf("header mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantMoonPkgJSON, moonPkgJSON.String()); diff != "" {
		t.Errorf("moonPkgJSON mismatch (-want +got):\n%v", diff)
	}
}

var wantBody = `pub fn __modus_log_message(message : String) -> Unit {
  log_message(message)
}

pub fn __modus_add(x : Int, y : Int) -> Int {
  add(x, y)
}

pub fn __modus_add3(a : Int, b : Int, c : Int) -> Int {
  add3(a, b, c~)
}

pub fn __modus_add3_WithDefaults(a : Int, b : Int) -> Int {
  add3(a, b)
}

pub fn __modus_add_n(args : Array[Int]) -> Int {
  add_n(args)
}

pub fn __modus_get_current_time(now : @wallClock.Datetime) -> @time.ZonedDateTime!Error {
  try get_current_time!(now~) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_current_time_WithDefaults() -> @time.ZonedDateTime!Error {
  try get_current_time!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_current_time_formatted(now : @wallClock.Datetime) -> String!Error {
  try get_current_time_formatted!(now~) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_current_time_formatted_WithDefaults() -> String!Error {
  try get_current_time_formatted!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_full_name(first_name : String, last_name : String) -> String {
  get_full_name(first_name, last_name)
}

pub fn __modus_say_hello(name : String?) -> String {
  say_hello(name~)
}

pub fn __modus_say_hello_WithDefaults() -> String {
  say_hello()
}

pub fn __modus_get_person() -> Person {
  get_person()
}

pub fn __modus_get_random_person() -> Person {
  get_random_person()
}

pub fn __modus_get_people() -> Array[Person] {
  get_people()
}

pub fn __modus_get_name_and_age() -> (String, Int) {
  get_name_and_age()
}

pub fn __modus_test_normal_error(input : String) -> String!Error {
  try test_normal_error!(input) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_alternative_error(input : String) -> String {
  test_alternative_error(input)
}

pub fn __modus_test_abort() -> Unit {
  test_abort()
}

pub fn __modus_test_exit() -> Unit {
  test_exit()
}

pub fn __modus_test_logging() -> Unit {
  test_logging()
}

`

var wantHeader = `// Code generated by modus-moonbit-build. DO NOT EDIT.

`

var wantMoonPkgJSON = `{
  "import": [
    "gmlewis/modus/pkg/console",
    "gmlewis/modus/wit/interface/imports/wasi/clocks/wallClock",
    "moonbitlang/x/sys",
    "moonbitlang/x/time"
  ],
  "test-import": [
    "gmlewis/modus/pkg/testutils"
  ],
  "targets": {
    "modus_post_generated.mbt": [
      "wasm"
    ],
    "modus_pre_generated.mbt": [
      "wasm"
    ]
  },
  "link": {
    "wasm": {
      "exports": [
        "__modus_add3:add3",
        "__modus_add3_WithDefaults:add3_WithDefaults",
        "__modus_add:add",
        "__modus_add_n:add_n",
        "__modus_get_current_time:get_current_time",
        "__modus_get_current_time_WithDefaults:get_current_time_WithDefaults",
        "__modus_get_current_time_formatted:get_current_time_formatted",
        "__modus_get_current_time_formatted_WithDefaults:get_current_time_formatted_WithDefaults",
        "__modus_get_full_name:get_full_name",
        "__modus_get_name_and_age:get_name_and_age",
        "__modus_get_people:get_people",
        "__modus_get_person:get_person",
        "__modus_get_random_person:get_random_person",
        "__modus_log_message:log_message",
        "__modus_say_hello:say_hello",
        "__modus_say_hello_WithDefaults:say_hello_WithDefaults",
        "__modus_test_abort:test_abort",
        "__modus_test_alternative_error:test_alternative_error",
        "__modus_test_exit:test_exit",
        "__modus_test_logging:test_logging",
        "__modus_test_normal_error:test_normal_error",
        "bytes2ptr",
        "cabi_realloc",
        "copy",
        "double_array2ptr",
        "duration_from_nanos",
        "extend16",
        "extend8",
        "f32_to_i32",
        "f32_to_i64",
        "float_array2ptr",
        "free",
        "int64_array2ptr",
        "int_array2ptr",
        "load16",
        "load16_u",
        "load32",
        "load64",
        "load8",
        "load8_u",
        "loadf32",
        "loadf64",
        "malloc",
        "ptr2bytes",
        "ptr2double_array",
        "ptr2float_array",
        "ptr2int64_array",
        "ptr2int_array",
        "ptr2str",
        "ptr2uint64_array",
        "ptr2uint_array",
        "store16",
        "store32",
        "store64",
        "store8",
        "storef32",
        "storef64",
        "str2ptr",
        "uint64_array2ptr",
        "uint_array2ptr",
        "zoned_date_time_from_unix_seconds_and_nanos"
      ],
      "export-memory-name": "memory"
    }
  }
}`
