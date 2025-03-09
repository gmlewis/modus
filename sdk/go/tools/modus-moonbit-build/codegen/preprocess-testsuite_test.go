// -*- compile-command: "go test -run ^TestTestablePreProcess_Testsuite ."; -*-

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

func TestTestablePreProcess_Testsuite(t *testing.T) {
	t.Parallel()

	config := &config.Config{
		SourceDir: "testdata/test-suite",
	}

	mod := preProcessTestSetup(t, config)

	body, header, moonPkgJSON, err := testablePreProcess(config, mod)
	if err != nil {
		t.Fatal(err)
	}

	wg := &preProcessDiffs{
		wantPreProcessBody:        wantTestsuitePreProcessBody,
		gotPreProcessBody:         body.String(),
		wantPreProcessHeader:      wantTestsuitePreProcessHeader,
		gotPreProcessHeader:       header.String(),
		wantPreProcessMoonPkgJSON: wantTestsuitePreProcessMoonPkgJSON,
		gotPreProcessMoonPkgJSON:  moonPkgJSON.String(),
	}
	reportPreProcessDiffs(t, "testsuite", wg)
}

var wantTestsuitePreProcessBody = `pub fn __modus_hello_array_of_ints(n : Int) -> Array[Int] {
  hello_array_of_ints(n~)
}

pub fn __modus_hello_array_of_ints_WithDefaults() -> Array[Int] {
  hello_array_of_ints()
}

pub fn __modus_hello_array_of_ints_option(n : Int) -> Array[Int]? {
  hello_array_of_ints_option(n~)
}

pub fn __modus_hello_array_of_ints_option_WithDefaults() -> Array[Int]? {
  hello_array_of_ints_option()
}

pub fn __modus_hello_array_of_strings(n : Int) -> Array[String] {
  hello_array_of_strings(n~)
}

pub fn __modus_hello_array_of_strings_WithDefaults() -> Array[String] {
  hello_array_of_strings()
}

pub fn __modus_hello_array_of_strings_option(n : Int) -> Array[String]? {
  hello_array_of_strings_option(n~)
}

pub fn __modus_hello_array_of_strings_option_WithDefaults() -> Array[String]? {
  hello_array_of_strings_option()
}

pub fn __modus_hello_maps_n_items(n : Int) -> Map[String, String] {
  hello_maps_n_items(n~)
}

pub fn __modus_hello_maps_n_items_WithDefaults() -> Map[String, String] {
  hello_maps_n_items()
}

pub fn __modus_hello_maps_n_items_option(n : Int) -> Map[String, String]? {
  hello_maps_n_items_option(n~)
}

pub fn __modus_hello_maps_n_items_option_WithDefaults() -> Map[String, String]? {
  hello_maps_n_items_option()
}

pub fn __modus_cabi_realloc(src_offset : Int, src_size : Int, _dst_alignment : Int, dst_size : Int) -> Int {
  cabi_realloc(src_offset, src_size, _dst_alignment, dst_size)
}

pub fn __modus_malloc(size : Int) -> Int {
  malloc(size)
}

pub fn __modus_copy(dest : Int, src : Int) -> Unit {
  copy(dest, src)
}

pub fn __modus_ptr2uint64_array(ptr : Int) -> FixedArray[UInt64] {
  ptr2uint64_array(ptr)
}

pub fn __modus_ptr2int64_array(ptr : Int) -> FixedArray[Int64] {
  ptr2int64_array(ptr)
}

pub fn __modus_ptr2double_array(ptr : Int) -> FixedArray[Double] {
  ptr2double_array(ptr)
}

pub fn __modus___modus_hello_array_of_ints(n : Int) -> Array[Int] {
  __modus_hello_array_of_ints(n)
}

pub fn __modus___modus_hello_array_of_ints_WithDefaults() -> Array[Int] {
  __modus_hello_array_of_ints()
}

pub fn __modus___modus_hello_array_of_ints_option(n : Int) -> Array[Int]? {
  __modus_hello_array_of_ints_option(n)
}

pub fn __modus___modus_hello_array_of_ints_option_WithDefaults() -> Array[Int]? {
  __modus_hello_array_of_ints_option()
}

pub fn __modus___modus_hello_array_of_strings(n : Int) -> Array[String] {
  __modus_hello_array_of_strings(n)
}

pub fn __modus___modus_hello_array_of_strings_WithDefaults() -> Array[String] {
  __modus_hello_array_of_strings()
}

pub fn __modus___modus_hello_array_of_strings_option(n : Int) -> Array[String]? {
  __modus_hello_array_of_strings_option(n)
}

pub fn __modus___modus_hello_array_of_strings_option_WithDefaults() -> Array[String]? {
  __modus_hello_array_of_strings_option()
}

pub fn __modus___modus_hello_maps_n_items(n : Int) -> Map[String, String] {
  __modus_hello_maps_n_items(n)
}

pub fn __modus___modus_hello_maps_n_items_WithDefaults() -> Map[String, String] {
  __modus_hello_maps_n_items()
}

pub fn __modus___modus_hello_maps_n_items_option(n : Int) -> Map[String, String]? {
  __modus_hello_maps_n_items_option(n)
}

pub fn __modus___modus_hello_maps_n_items_option_WithDefaults() -> Map[String, String]? {
  __modus_hello_maps_n_items_option()
}

pub fn __modus___modus_hello_primitive_bool_min() -> Bool {
  __modus_hello_primitive_bool_min()
}

pub fn __modus___modus_hello_primitive_bool_max() -> Bool {
  __modus_hello_primitive_bool_max()
}

pub fn __modus___modus_hello_primitive_byte_min() -> Byte {
  __modus_hello_primitive_byte_min()
}

pub fn __modus___modus_hello_primitive_byte_max() -> Byte {
  __modus_hello_primitive_byte_max()
}

pub fn __modus___modus_hello_primitive_char_min() -> Char {
  __modus_hello_primitive_char_min()
}

pub fn __modus___modus_hello_primitive_char_max() -> Char {
  __modus_hello_primitive_char_max()
}

pub fn __modus___modus_hello_primitive_double_min() -> Double {
  __modus_hello_primitive_double_min()
}

pub fn __modus___modus_hello_primitive_double_max() -> Double {
  __modus_hello_primitive_double_max()
}

pub fn __modus___modus_hello_primitive_float_min() -> Float {
  __modus_hello_primitive_float_min()
}

pub fn __modus___modus_hello_primitive_float_max() -> Float {
  __modus_hello_primitive_float_max()
}

pub fn __modus___modus_hello_primitive_int_min() -> Int {
  __modus_hello_primitive_int_min()
}

pub fn __modus___modus_hello_primitive_int_max() -> Int {
  __modus_hello_primitive_int_max()
}

pub fn __modus___modus_hello_primitive_int16_min() -> Int16 {
  __modus_hello_primitive_int16_min()
}

pub fn __modus___modus_hello_primitive_int16_max() -> Int16 {
  __modus_hello_primitive_int16_max()
}

pub fn __modus___modus_hello_primitive_int64_min() -> Int64 {
  __modus_hello_primitive_int64_min()
}

pub fn __modus___modus_hello_primitive_int64_max() -> Int64 {
  __modus_hello_primitive_int64_max()
}

pub fn __modus___modus_hello_primitive_uint_min() -> UInt {
  __modus_hello_primitive_uint_min()
}

pub fn __modus___modus_hello_primitive_uint_max() -> UInt {
  __modus_hello_primitive_uint_max()
}

pub fn __modus___modus_hello_primitive_uint16_min() -> UInt16 {
  __modus_hello_primitive_uint16_min()
}

pub fn __modus___modus_hello_primitive_uint16_max() -> UInt16 {
  __modus_hello_primitive_uint16_max()
}

pub fn __modus___modus_hello_primitive_uint64_min() -> UInt64 {
  __modus_hello_primitive_uint64_min()
}

pub fn __modus___modus_hello_primitive_uint64_max() -> UInt64 {
  __modus_hello_primitive_uint64_max()
}

pub fn __modus___modus_log_message(message : String) -> Unit {
  __modus_log_message(message)
}

pub fn __modus___modus_add(x : Int, y : Int) -> Int {
  __modus_add(x, y)
}

pub fn __modus___modus_add3(a : Int, b : Int, c : Int) -> Int {
  __modus_add3(a, b, c)
}

pub fn __modus___modus_add3_WithDefaults(a : Int, b : Int) -> Int {
  __modus_add3(a, b)
}

pub fn __modus___modus_get_current_time() -> @time.ZonedDateTime!Error {
  try __modus_get_current_time!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus___modus_get_current_time_formatted() -> String!Error {
  try __modus_get_current_time_formatted!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus___modus_get_full_name(first_name : String, last_name : String) -> String {
  __modus_get_full_name(first_name, last_name)
}

pub fn __modus___modus_get_person() -> Person {
  __modus_get_person()
}

pub fn __modus___modus_get_random_person() -> Person {
  __modus_get_random_person()
}

pub fn __modus___modus_get_people() -> Array[Person] {
  __modus_get_people()
}

pub fn __modus___modus_test_normal_error(input : String) -> String!Error {
  try __modus_test_normal_error!(input) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus___modus_test_alternative_error(input : String) -> String {
  __modus_test_alternative_error(input)
}

pub fn __modus___modus_test_abort() -> Unit {
  __modus_test_abort()
}

pub fn __modus___modus_test_exit() -> Unit {
  __modus_test_exit()
}

pub fn __modus___modus_test_logging() -> Unit {
  __modus_test_logging()
}

pub fn __modus___modus_hello_option_empty_string(some : Bool) -> String? {
  __modus_hello_option_empty_string(some)
}

pub fn __modus___modus_hello_option_empty_string_WithDefaults() -> String? {
  __modus_hello_option_empty_string()
}

pub fn __modus___modus_hello_option_none(some : Bool) -> String? {
  __modus_hello_option_none(some)
}

pub fn __modus___modus_hello_option_none_WithDefaults() -> String? {
  __modus_hello_option_none()
}

pub fn __modus___modus_hello_option_some_string(some : Bool) -> String? {
  __modus_hello_option_some_string(some)
}

pub fn __modus___modus_hello_option_some_string_WithDefaults() -> String? {
  __modus_hello_option_some_string()
}

pub fn __modus___modus_hello_world_with_arg(name : String) -> String {
  __modus_hello_world_with_arg(name)
}

pub fn __modus___modus_hello_world_with_optional_arg(name : String) -> String {
  __modus_hello_world_with_optional_arg(name)
}

pub fn __modus___modus_hello_world_with_optional_arg_WithDefaults() -> String {
  __modus_hello_world_with_optional_arg()
}

pub fn __modus___modus_hello_world() -> String {
  __modus_hello_world()
}

pub fn __modus___modus_get_utc_time() -> @time.ZonedDateTime!Error {
  try __modus_get_utc_time!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus___modus_get_local_time() -> String!Error {
  try __modus_get_local_time!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus___modus_get_time_in_zone(tz : String) -> String!Error {
  try __modus_get_time_in_zone!(tz) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus___modus_get_local_time_zone_id() -> String {
  __modus_get_local_time_zone_id()
}

pub fn __modus___modus_get_time_zone_info(tz : String) -> TimeZoneInfo!Error {
  try __modus_get_time_zone_info!(tz) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_hello_primitive_bool_min() -> Bool {
  hello_primitive_bool_min()
}

pub fn __modus_hello_primitive_bool_max() -> Bool {
  hello_primitive_bool_max()
}

pub fn __modus_hello_primitive_byte_min() -> Byte {
  hello_primitive_byte_min()
}

pub fn __modus_hello_primitive_byte_max() -> Byte {
  hello_primitive_byte_max()
}

pub fn __modus_hello_primitive_char_min() -> Char {
  hello_primitive_char_min()
}

pub fn __modus_hello_primitive_char_max() -> Char {
  hello_primitive_char_max()
}

pub fn __modus_hello_primitive_double_min() -> Double {
  hello_primitive_double_min()
}

pub fn __modus_hello_primitive_double_max() -> Double {
  hello_primitive_double_max()
}

pub fn __modus_hello_primitive_float_min() -> Float {
  hello_primitive_float_min()
}

pub fn __modus_hello_primitive_float_max() -> Float {
  hello_primitive_float_max()
}

pub fn __modus_hello_primitive_int_min() -> Int {
  hello_primitive_int_min()
}

pub fn __modus_hello_primitive_int_max() -> Int {
  hello_primitive_int_max()
}

pub fn __modus_hello_primitive_int16_min() -> Int16 {
  hello_primitive_int16_min()
}

pub fn __modus_hello_primitive_int16_max() -> Int16 {
  hello_primitive_int16_max()
}

pub fn __modus_hello_primitive_int64_min() -> Int64 {
  hello_primitive_int64_min()
}

pub fn __modus_hello_primitive_int64_max() -> Int64 {
  hello_primitive_int64_max()
}

pub fn __modus_hello_primitive_uint_min() -> UInt {
  hello_primitive_uint_min()
}

pub fn __modus_hello_primitive_uint_max() -> UInt {
  hello_primitive_uint_max()
}

pub fn __modus_hello_primitive_uint16_min() -> UInt16 {
  hello_primitive_uint16_min()
}

pub fn __modus_hello_primitive_uint16_max() -> UInt16 {
  hello_primitive_uint16_max()
}

pub fn __modus_hello_primitive_uint64_min() -> UInt64 {
  hello_primitive_uint64_min()
}

pub fn __modus_hello_primitive_uint64_max() -> UInt64 {
  hello_primitive_uint64_max()
}

pub fn __modus_log_message(message : String) -> Unit {
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

pub fn __modus_get_current_time() -> @time.ZonedDateTime!Error {
  try get_current_time!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_current_time_formatted() -> String!Error {
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

pub fn __modus_get_person() -> Person {
  get_person()
}

pub fn __modus_get_random_person() -> Person {
  get_random_person()
}

pub fn __modus_get_people() -> Array[Person] {
  get_people()
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

pub fn __modus_hello_option_empty_string(some : Bool) -> String? {
  hello_option_empty_string(some~)
}

pub fn __modus_hello_option_empty_string_WithDefaults() -> String? {
  hello_option_empty_string()
}

pub fn __modus_hello_option_none(some : Bool) -> String? {
  hello_option_none(some~)
}

pub fn __modus_hello_option_none_WithDefaults() -> String? {
  hello_option_none()
}

pub fn __modus_hello_option_some_string(some : Bool) -> String? {
  hello_option_some_string(some~)
}

pub fn __modus_hello_option_some_string_WithDefaults() -> String? {
  hello_option_some_string()
}

pub fn __modus_hello_world_with_arg(name : String) -> String {
  hello_world_with_arg(name)
}

pub fn __modus_hello_world_with_optional_arg(name : String) -> String {
  hello_world_with_optional_arg(name~)
}

pub fn __modus_hello_world_with_optional_arg_WithDefaults() -> String {
  hello_world_with_optional_arg()
}

pub fn __modus_hello_world() -> String {
  hello_world()
}

pub fn __modus_get_utc_time() -> @time.ZonedDateTime!Error {
  try get_utc_time!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_local_time() -> String!Error {
  try get_local_time!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_time_in_zone(tz : String) -> String!Error {
  try get_time_in_zone!(tz) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_local_time_zone_id() -> String {
  get_local_time_zone_id()
}

pub fn __modus_get_time_zone_info(tz : String) -> TimeZoneInfo!Error {
  try get_time_zone_info!(tz) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

`

var wantTestsuitePreProcessHeader = `// Code generated by modus-moonbit-build. DO NOT EDIT.

`

var wantTestsuitePreProcessMoonPkgJSON = `{
  "import": [
    "gmlewis/modus/pkg/console",
    "gmlewis/modus/pkg/localtime",
    "gmlewis/modus/wit/interface/wasi",
    "gmlewis/modus/pkg/time",
    "moonbitlang/x/sys"
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
        "__modus___modus_add3:__modus_add3",
        "__modus___modus_add3_WithDefaults:__modus_add3_WithDefaults",
        "__modus___modus_add:__modus_add",
        "__modus___modus_get_current_time:__modus_get_current_time",
        "__modus___modus_get_current_time_formatted:__modus_get_current_time_formatted",
        "__modus___modus_get_full_name:__modus_get_full_name",
        "__modus___modus_get_local_time:__modus_get_local_time",
        "__modus___modus_get_local_time_zone_id:__modus_get_local_time_zone_id",
        "__modus___modus_get_people:__modus_get_people",
        "__modus___modus_get_person:__modus_get_person",
        "__modus___modus_get_random_person:__modus_get_random_person",
        "__modus___modus_get_time_in_zone:__modus_get_time_in_zone",
        "__modus___modus_get_time_zone_info:__modus_get_time_zone_info",
        "__modus___modus_get_utc_time:__modus_get_utc_time",
        "__modus___modus_hello_array_of_ints:__modus_hello_array_of_ints",
        "__modus___modus_hello_array_of_ints_WithDefaults:__modus_hello_array_of_ints_WithDefaults",
        "__modus___modus_hello_array_of_ints_option:__modus_hello_array_of_ints_option",
        "__modus___modus_hello_array_of_ints_option_WithDefaults:__modus_hello_array_of_ints_option_WithDefaults",
        "__modus___modus_hello_array_of_strings:__modus_hello_array_of_strings",
        "__modus___modus_hello_array_of_strings_WithDefaults:__modus_hello_array_of_strings_WithDefaults",
        "__modus___modus_hello_array_of_strings_option:__modus_hello_array_of_strings_option",
        "__modus___modus_hello_array_of_strings_option_WithDefaults:__modus_hello_array_of_strings_option_WithDefaults",
        "__modus___modus_hello_maps_n_items:__modus_hello_maps_n_items",
        "__modus___modus_hello_maps_n_items_WithDefaults:__modus_hello_maps_n_items_WithDefaults",
        "__modus___modus_hello_maps_n_items_option:__modus_hello_maps_n_items_option",
        "__modus___modus_hello_maps_n_items_option_WithDefaults:__modus_hello_maps_n_items_option_WithDefaults",
        "__modus___modus_hello_option_empty_string:__modus_hello_option_empty_string",
        "__modus___modus_hello_option_empty_string_WithDefaults:__modus_hello_option_empty_string_WithDefaults",
        "__modus___modus_hello_option_none:__modus_hello_option_none",
        "__modus___modus_hello_option_none_WithDefaults:__modus_hello_option_none_WithDefaults",
        "__modus___modus_hello_option_some_string:__modus_hello_option_some_string",
        "__modus___modus_hello_option_some_string_WithDefaults:__modus_hello_option_some_string_WithDefaults",
        "__modus___modus_hello_primitive_bool_max:__modus_hello_primitive_bool_max",
        "__modus___modus_hello_primitive_bool_min:__modus_hello_primitive_bool_min",
        "__modus___modus_hello_primitive_byte_max:__modus_hello_primitive_byte_max",
        "__modus___modus_hello_primitive_byte_min:__modus_hello_primitive_byte_min",
        "__modus___modus_hello_primitive_char_max:__modus_hello_primitive_char_max",
        "__modus___modus_hello_primitive_char_min:__modus_hello_primitive_char_min",
        "__modus___modus_hello_primitive_double_max:__modus_hello_primitive_double_max",
        "__modus___modus_hello_primitive_double_min:__modus_hello_primitive_double_min",
        "__modus___modus_hello_primitive_float_max:__modus_hello_primitive_float_max",
        "__modus___modus_hello_primitive_float_min:__modus_hello_primitive_float_min",
        "__modus___modus_hello_primitive_int16_max:__modus_hello_primitive_int16_max",
        "__modus___modus_hello_primitive_int16_min:__modus_hello_primitive_int16_min",
        "__modus___modus_hello_primitive_int64_max:__modus_hello_primitive_int64_max",
        "__modus___modus_hello_primitive_int64_min:__modus_hello_primitive_int64_min",
        "__modus___modus_hello_primitive_int_max:__modus_hello_primitive_int_max",
        "__modus___modus_hello_primitive_int_min:__modus_hello_primitive_int_min",
        "__modus___modus_hello_primitive_uint16_max:__modus_hello_primitive_uint16_max",
        "__modus___modus_hello_primitive_uint16_min:__modus_hello_primitive_uint16_min",
        "__modus___modus_hello_primitive_uint64_max:__modus_hello_primitive_uint64_max",
        "__modus___modus_hello_primitive_uint64_min:__modus_hello_primitive_uint64_min",
        "__modus___modus_hello_primitive_uint_max:__modus_hello_primitive_uint_max",
        "__modus___modus_hello_primitive_uint_min:__modus_hello_primitive_uint_min",
        "__modus___modus_hello_world:__modus_hello_world",
        "__modus___modus_hello_world_with_arg:__modus_hello_world_with_arg",
        "__modus___modus_hello_world_with_optional_arg:__modus_hello_world_with_optional_arg",
        "__modus___modus_hello_world_with_optional_arg_WithDefaults:__modus_hello_world_with_optional_arg_WithDefaults",
        "__modus___modus_log_message:__modus_log_message",
        "__modus___modus_test_abort:__modus_test_abort",
        "__modus___modus_test_alternative_error:__modus_test_alternative_error",
        "__modus___modus_test_exit:__modus_test_exit",
        "__modus___modus_test_logging:__modus_test_logging",
        "__modus___modus_test_normal_error:__modus_test_normal_error",
        "__modus_add3:add3",
        "__modus_add3_WithDefaults:add3_WithDefaults",
        "__modus_add:add",
        "__modus_cabi_realloc:cabi_realloc",
        "__modus_copy:copy",
        "__modus_get_current_time:get_current_time",
        "__modus_get_current_time_formatted:get_current_time_formatted",
        "__modus_get_full_name:get_full_name",
        "__modus_get_local_time:get_local_time",
        "__modus_get_local_time_zone_id:get_local_time_zone_id",
        "__modus_get_people:get_people",
        "__modus_get_person:get_person",
        "__modus_get_random_person:get_random_person",
        "__modus_get_time_in_zone:get_time_in_zone",
        "__modus_get_time_zone_info:get_time_zone_info",
        "__modus_get_utc_time:get_utc_time",
        "__modus_hello_array_of_ints:hello_array_of_ints",
        "__modus_hello_array_of_ints_WithDefaults:hello_array_of_ints_WithDefaults",
        "__modus_hello_array_of_ints_option:hello_array_of_ints_option",
        "__modus_hello_array_of_ints_option_WithDefaults:hello_array_of_ints_option_WithDefaults",
        "__modus_hello_array_of_strings:hello_array_of_strings",
        "__modus_hello_array_of_strings_WithDefaults:hello_array_of_strings_WithDefaults",
        "__modus_hello_array_of_strings_option:hello_array_of_strings_option",
        "__modus_hello_array_of_strings_option_WithDefaults:hello_array_of_strings_option_WithDefaults",
        "__modus_hello_maps_n_items:hello_maps_n_items",
        "__modus_hello_maps_n_items_WithDefaults:hello_maps_n_items_WithDefaults",
        "__modus_hello_maps_n_items_option:hello_maps_n_items_option",
        "__modus_hello_maps_n_items_option_WithDefaults:hello_maps_n_items_option_WithDefaults",
        "__modus_hello_option_empty_string:hello_option_empty_string",
        "__modus_hello_option_empty_string_WithDefaults:hello_option_empty_string_WithDefaults",
        "__modus_hello_option_none:hello_option_none",
        "__modus_hello_option_none_WithDefaults:hello_option_none_WithDefaults",
        "__modus_hello_option_some_string:hello_option_some_string",
        "__modus_hello_option_some_string_WithDefaults:hello_option_some_string_WithDefaults",
        "__modus_hello_primitive_bool_max:hello_primitive_bool_max",
        "__modus_hello_primitive_bool_min:hello_primitive_bool_min",
        "__modus_hello_primitive_byte_max:hello_primitive_byte_max",
        "__modus_hello_primitive_byte_min:hello_primitive_byte_min",
        "__modus_hello_primitive_char_max:hello_primitive_char_max",
        "__modus_hello_primitive_char_min:hello_primitive_char_min",
        "__modus_hello_primitive_double_max:hello_primitive_double_max",
        "__modus_hello_primitive_double_min:hello_primitive_double_min",
        "__modus_hello_primitive_float_max:hello_primitive_float_max",
        "__modus_hello_primitive_float_min:hello_primitive_float_min",
        "__modus_hello_primitive_int16_max:hello_primitive_int16_max",
        "__modus_hello_primitive_int16_min:hello_primitive_int16_min",
        "__modus_hello_primitive_int64_max:hello_primitive_int64_max",
        "__modus_hello_primitive_int64_min:hello_primitive_int64_min",
        "__modus_hello_primitive_int_max:hello_primitive_int_max",
        "__modus_hello_primitive_int_min:hello_primitive_int_min",
        "__modus_hello_primitive_uint16_max:hello_primitive_uint16_max",
        "__modus_hello_primitive_uint16_min:hello_primitive_uint16_min",
        "__modus_hello_primitive_uint64_max:hello_primitive_uint64_max",
        "__modus_hello_primitive_uint64_min:hello_primitive_uint64_min",
        "__modus_hello_primitive_uint_max:hello_primitive_uint_max",
        "__modus_hello_primitive_uint_min:hello_primitive_uint_min",
        "__modus_hello_world:hello_world",
        "__modus_hello_world_with_arg:hello_world_with_arg",
        "__modus_hello_world_with_optional_arg:hello_world_with_optional_arg",
        "__modus_hello_world_with_optional_arg_WithDefaults:hello_world_with_optional_arg_WithDefaults",
        "__modus_log_message:log_message",
        "__modus_malloc:malloc",
        "__modus_ptr2double_array:ptr2double_array",
        "__modus_ptr2int64_array:ptr2int64_array",
        "__modus_ptr2uint64_array:ptr2uint64_array",
        "__modus_test_abort:test_abort",
        "__modus_test_alternative_error:test_alternative_error",
        "__modus_test_exit:test_exit",
        "__modus_test_logging:test_logging",
        "__modus_test_normal_error:test_normal_error",
        "cabi_realloc",
        "copy",
        "duration_from_nanos",
        "free",
        "load32",
        "malloc",
        "ptr2str",
        "ptr_to_none",
        "read_map",
        "store32",
        "store8",
        "write_map",
        "zoned_date_time_from_unix_seconds_and_nanos"
      ],
      "export-memory-name": "memory"
    }
  }
}`
