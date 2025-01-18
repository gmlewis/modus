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

pub fn __modus_get_current_time(now : @wallClock.Datetime) -> @time.PlainDateTime {
  try get_current_time!(now~) {
    e => {
      @console.error(e.to_string())
      abort(e.to_string())
    }
  }
}

pub fn __modus_get_current_time_WithDefaults() -> @time.PlainDateTime {
  try get_current_time!() {
    e => {
      @console.error(e.to_string())
      abort(e.to_string())
    }
  }
}

pub fn __modus_get_current_time_formatted(now : @wallClock.Datetime) -> String {
  try get_current_time_formatted!(now~) {
    e => {
      @console.error(e.to_string())
      abort(e.to_string())
    }
  }
}

pub fn __modus_get_current_time_formatted_WithDefaults() -> String {
  try get_current_time_formatted!() {
    e => {
      @console.error(e.to_string())
      abort(e.to_string())
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

pub fn __modus_test_normal_error(input : String) -> String {
  try test_normal_error!(input) {
    e => {
      @console.error(e.to_string())
      abort(e.to_string())
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
        "__modus_test_normal_error:test_normal_error"
      ],
      "export-memory-name": "memory"
    }
  }
}`
