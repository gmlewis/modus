// -*- compile-command: "go test -run ^TestTestablePreProcess_Runtime ."; -*-

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

func TestTestablePreProcess_Runtime(t *testing.T) {
	t.Parallel()

	config := &config.Config{
		SourceDir: "../testdata/runtime-testdata",
	}

	mod := preProcessTestSetup(t, config)

	body, header, moonPkgJSON, err := testablePreProcess(config, mod)
	if err != nil {
		t.Fatal(err)
	}

	wg := &preProcessDiffs{
		wantPreProcessBody:        wantRuntimePreProcessBody,
		gotPreProcessBody:         body.String(),
		wantPreProcessHeader:      wantRuntimePreProcessHeader,
		gotPreProcessHeader:       header.String(),
		wantPreProcessMoonPkgJSON: wantRuntimePreProcessMoonPkgJSON,
		gotPreProcessMoonPkgJSON:  moonPkgJSON.String(),
	}
	reportPreProcessDiffs(t, "runtime", wg)
}

var wantRuntimePreProcessBody = `pub fn __modus_test_array_output_bool_0() -> Array[Bool] {
  test_array_output_bool_0()
}

pub fn __modus_test_array_output_bool_1() -> Array[Bool] {
  test_array_output_bool_1()
}

pub fn __modus_test_array_output_bool_2() -> Array[Bool] {
  test_array_output_bool_2()
}

pub fn __modus_test_array_output_bool_3() -> Array[Bool] {
  test_array_output_bool_3()
}

pub fn __modus_test_array_output_bool_4() -> Array[Bool] {
  test_array_output_bool_4()
}

pub fn __modus_test_array_output_bool_option_0() -> Array[Bool?] {
  test_array_output_bool_option_0()
}

pub fn __modus_test_array_output_bool_option_1_none() -> Array[Bool?] {
  test_array_output_bool_option_1_none()
}

pub fn __modus_test_array_output_bool_option_1_false() -> Array[Bool?] {
  test_array_output_bool_option_1_false()
}

pub fn __modus_test_array_output_bool_option_1_true() -> Array[Bool?] {
  test_array_output_bool_option_1_true()
}

pub fn __modus_test_array_output_bool_option_2() -> Array[Bool?] {
  test_array_output_bool_option_2()
}

pub fn __modus_test_array_output_bool_option_3() -> Array[Bool?] {
  test_array_output_bool_option_3()
}

pub fn __modus_test_array_output_bool_option_4() -> Array[Bool?] {
  test_array_output_bool_option_4()
}

pub fn __modus_test_array_input_bool_0(val : Array[Bool]) -> Unit!Error {
  try test_array_input_bool_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_1(val : Array[Bool]) -> Unit!Error {
  try test_array_input_bool_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_2(val : Array[Bool]) -> Unit!Error {
  try test_array_input_bool_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_3(val : Array[Bool]) -> Unit!Error {
  try test_array_input_bool_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_4(val : Array[Bool]) -> Unit!Error {
  try test_array_input_bool_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_option_0(val : Array[Bool?]) -> Unit!Error {
  try test_array_input_bool_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_option_1_none(val : Array[Bool?]) -> Unit!Error {
  try test_array_input_bool_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_option_1_false(val : Array[Bool?]) -> Unit!Error {
  try test_array_input_bool_option_1_false!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_option_1_true(val : Array[Bool?]) -> Unit!Error {
  try test_array_input_bool_option_1_true!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_option_2(val : Array[Bool?]) -> Unit!Error {
  try test_array_input_bool_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_option_3(val : Array[Bool?]) -> Unit!Error {
  try test_array_input_bool_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_bool_option_4(val : Array[Bool?]) -> Unit!Error {
  try test_array_input_bool_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_byte_0() -> Array[Byte] {
  test_array_output_byte_0()
}

pub fn __modus_test_array_output_byte_1() -> Array[Byte] {
  test_array_output_byte_1()
}

pub fn __modus_test_array_output_byte_2() -> Array[Byte] {
  test_array_output_byte_2()
}

pub fn __modus_test_array_output_byte_3() -> Array[Byte] {
  test_array_output_byte_3()
}

pub fn __modus_test_array_output_byte_4() -> Array[Byte] {
  test_array_output_byte_4()
}

pub fn __modus_test_array_output_byte_option_0() -> Array[Byte?] {
  test_array_output_byte_option_0()
}

pub fn __modus_test_array_output_byte_option_1() -> Array[Byte?] {
  test_array_output_byte_option_1()
}

pub fn __modus_test_array_output_byte_option_2() -> Array[Byte?] {
  test_array_output_byte_option_2()
}

pub fn __modus_test_array_output_byte_option_3() -> Array[Byte?] {
  test_array_output_byte_option_3()
}

pub fn __modus_test_array_output_byte_option_4() -> Array[Byte?] {
  test_array_output_byte_option_4()
}

pub fn __modus_test_array_input_byte_0(val : Array[Byte]) -> Unit!Error {
  try test_array_input_byte_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_byte_1(val : Array[Byte]) -> Unit!Error {
  try test_array_input_byte_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_byte_2(val : Array[Byte]) -> Unit!Error {
  try test_array_input_byte_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_byte_3(val : Array[Byte]) -> Unit!Error {
  try test_array_input_byte_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_byte_4(val : Array[Byte]) -> Unit!Error {
  try test_array_input_byte_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_byte_option_0(val : Array[Byte?]) -> Unit!Error {
  try test_array_input_byte_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_byte_option_1(val : Array[Byte?]) -> Unit!Error {
  try test_array_input_byte_option_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_byte_option_2(val : Array[Byte?]) -> Unit!Error {
  try test_array_input_byte_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_byte_option_3(val : Array[Byte?]) -> Unit!Error {
  try test_array_input_byte_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_byte_option_4(val : Array[Byte?]) -> Unit!Error {
  try test_array_input_byte_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_input_char_empty(val : Array[Char]) -> Unit!Error {
  try test_array_input_char_empty!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_char_0() -> Array[Char] {
  test_array_output_char_0()
}

pub fn __modus_test_array_output_char_1() -> Array[Char] {
  test_array_output_char_1()
}

pub fn __modus_test_array_output_char_2() -> Array[Char] {
  test_array_output_char_2()
}

pub fn __modus_test_array_output_char_3() -> Array[Char] {
  test_array_output_char_3()
}

pub fn __modus_test_array_output_char_4() -> Array[Char] {
  test_array_output_char_4()
}

pub fn __modus_test_array_output_char_option_0() -> Array[Char?] {
  test_array_output_char_option_0()
}

pub fn __modus_test_array_output_char_option_1_none() -> Array[Char?] {
  test_array_output_char_option_1_none()
}

pub fn __modus_test_array_output_char_option_1_some() -> Array[Char?] {
  test_array_output_char_option_1_some()
}

pub fn __modus_test_array_output_char_option_2() -> Array[Char?] {
  test_array_output_char_option_2()
}

pub fn __modus_test_array_output_char_option_3() -> Array[Char?] {
  test_array_output_char_option_3()
}

pub fn __modus_test_array_output_char_option_4() -> Array[Char?] {
  test_array_output_char_option_4()
}

pub fn __modus_test_array_input_char_option(val : Array[Char?]) -> Unit!Error {
  try test_array_input_char_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_char_option() -> Array[Char?] {
  test_array_output_char_option()
}

pub fn __modus_test_array_input_double_empty(val : Array[Double]) -> Unit!Error {
  try test_array_input_double_empty!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_double_0() -> Array[Double] {
  test_array_output_double_0()
}

pub fn __modus_test_array_output_double_1() -> Array[Double] {
  test_array_output_double_1()
}

pub fn __modus_test_array_output_double_2() -> Array[Double] {
  test_array_output_double_2()
}

pub fn __modus_test_array_output_double_3() -> Array[Double] {
  test_array_output_double_3()
}

pub fn __modus_test_array_output_double_4() -> Array[Double] {
  test_array_output_double_4()
}

pub fn __modus_test_array_output_double_option_0() -> Array[Double?] {
  test_array_output_double_option_0()
}

pub fn __modus_test_array_output_double_option_1_none() -> Array[Double?] {
  test_array_output_double_option_1_none()
}

pub fn __modus_test_array_output_double_option_1_some() -> Array[Double?] {
  test_array_output_double_option_1_some()
}

pub fn __modus_test_array_output_double_option_2() -> Array[Double?] {
  test_array_output_double_option_2()
}

pub fn __modus_test_array_output_double_option_3() -> Array[Double?] {
  test_array_output_double_option_3()
}

pub fn __modus_test_array_output_double_option_4() -> Array[Double?] {
  test_array_output_double_option_4()
}

pub fn __modus_test_array_input_double_option(val : Array[Double?]) -> Unit!Error {
  try test_array_input_double_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_double_option() -> Array[Double?] {
  test_array_output_double_option()
}

pub fn __modus_test_array_input_float_empty(val : Array[Float]) -> Unit!Error {
  try test_array_input_float_empty!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_float_0() -> Array[Float] {
  test_array_output_float_0()
}

pub fn __modus_test_array_output_float_1() -> Array[Float] {
  test_array_output_float_1()
}

pub fn __modus_test_array_output_float_2() -> Array[Float] {
  test_array_output_float_2()
}

pub fn __modus_test_array_output_float_3() -> Array[Float] {
  test_array_output_float_3()
}

pub fn __modus_test_array_output_float_4() -> Array[Float] {
  test_array_output_float_4()
}

pub fn __modus_test_array_output_float_option_0() -> Array[Float?] {
  test_array_output_float_option_0()
}

pub fn __modus_test_array_output_float_option_1_none() -> Array[Float?] {
  test_array_output_float_option_1_none()
}

pub fn __modus_test_array_output_float_option_1_some() -> Array[Float?] {
  test_array_output_float_option_1_some()
}

pub fn __modus_test_array_output_float_option_2() -> Array[Float?] {
  test_array_output_float_option_2()
}

pub fn __modus_test_array_output_float_option_3() -> Array[Float?] {
  test_array_output_float_option_3()
}

pub fn __modus_test_array_output_float_option_4() -> Array[Float?] {
  test_array_output_float_option_4()
}

pub fn __modus_test_array_input_float_option(val : Array[Float?]) -> Unit!Error {
  try test_array_input_float_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_float_option() -> Array[Float?] {
  test_array_output_float_option()
}

pub fn __modus_test_array_input_int_empty(val : Array[Int]) -> Unit!Error {
  try test_array_input_int_empty!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_int_0() -> Array[Int] {
  test_array_output_int_0()
}

pub fn __modus_test_array_output_int_1() -> Array[Int] {
  test_array_output_int_1()
}

pub fn __modus_test_array_output_int_1_min() -> Array[Int] {
  test_array_output_int_1_min()
}

pub fn __modus_test_array_output_int_1_max() -> Array[Int] {
  test_array_output_int_1_max()
}

pub fn __modus_test_array_output_int_2() -> Array[Int] {
  test_array_output_int_2()
}

pub fn __modus_test_array_output_int_3() -> Array[Int] {
  test_array_output_int_3()
}

pub fn __modus_test_array_output_int_4() -> Array[Int] {
  test_array_output_int_4()
}

pub fn __modus_test_array_output_int_option_0() -> Array[Int?] {
  test_array_output_int_option_0()
}

pub fn __modus_test_array_output_int_option_1_none() -> Array[Int?] {
  test_array_output_int_option_1_none()
}

pub fn __modus_test_array_output_int_option_1_min() -> Array[Int?] {
  test_array_output_int_option_1_min()
}

pub fn __modus_test_array_output_int_option_1_max() -> Array[Int?] {
  test_array_output_int_option_1_max()
}

pub fn __modus_test_array_output_int_option_2() -> Array[Int?] {
  test_array_output_int_option_2()
}

pub fn __modus_test_array_output_int_option_3() -> Array[Int?] {
  test_array_output_int_option_3()
}

pub fn __modus_test_array_output_int_option_4() -> Array[Int?] {
  test_array_output_int_option_4()
}

pub fn __modus_test_array_input_int_option(val : Array[Int?]) -> Unit!Error {
  try test_array_input_int_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_int_option() -> Array[Int?] {
  test_array_output_int_option()
}

pub fn __modus_test_array_input_int16_empty(val : Array[Int16]) -> Unit!Error {
  try test_array_input_int16_empty!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_int16_0() -> Array[Int16] {
  test_array_output_int16_0()
}

pub fn __modus_test_array_output_int16_1() -> Array[Int16] {
  test_array_output_int16_1()
}

pub fn __modus_test_array_output_int16_1_min() -> Array[Int16] {
  test_array_output_int16_1_min()
}

pub fn __modus_test_array_output_int16_1_max() -> Array[Int16] {
  test_array_output_int16_1_max()
}

pub fn __modus_test_array_output_int16_2() -> Array[Int16] {
  test_array_output_int16_2()
}

pub fn __modus_test_array_output_int16_3() -> Array[Int16] {
  test_array_output_int16_3()
}

pub fn __modus_test_array_output_int16_4() -> Array[Int16] {
  test_array_output_int16_4()
}

pub fn __modus_test_array_output_int16_option_0() -> Array[Int16?] {
  test_array_output_int16_option_0()
}

pub fn __modus_test_array_output_int16_option_1_none() -> Array[Int16?] {
  test_array_output_int16_option_1_none()
}

pub fn __modus_test_array_output_int16_option_1_min() -> Array[Int16?] {
  test_array_output_int16_option_1_min()
}

pub fn __modus_test_array_output_int16_option_1_max() -> Array[Int16?] {
  test_array_output_int16_option_1_max()
}

pub fn __modus_test_array_output_int16_option_2() -> Array[Int16?] {
  test_array_output_int16_option_2()
}

pub fn __modus_test_array_output_int16_option_3() -> Array[Int16?] {
  test_array_output_int16_option_3()
}

pub fn __modus_test_array_output_int16_option_4() -> Array[Int16?] {
  test_array_output_int16_option_4()
}

pub fn __modus_test_array_input_int16_option(val : Array[Int16?]) -> Unit!Error {
  try test_array_input_int16_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_int16_option() -> Array[Int16?] {
  test_array_output_int16_option()
}

pub fn __modus_test_array_output_int64_0() -> Array[Int64] {
  test_array_output_int64_0()
}

pub fn __modus_test_array_output_int64_1() -> Array[Int64] {
  test_array_output_int64_1()
}

pub fn __modus_test_array_output_int64_1_min() -> Array[Int64] {
  test_array_output_int64_1_min()
}

pub fn __modus_test_array_output_int64_1_max() -> Array[Int64] {
  test_array_output_int64_1_max()
}

pub fn __modus_test_array_output_int64_2() -> Array[Int64] {
  test_array_output_int64_2()
}

pub fn __modus_test_array_output_int64_3() -> Array[Int64] {
  test_array_output_int64_3()
}

pub fn __modus_test_array_output_int64_4() -> Array[Int64] {
  test_array_output_int64_4()
}

pub fn __modus_test_array_output_int64_option_0() -> Array[Int64?] {
  test_array_output_int64_option_0()
}

pub fn __modus_test_array_output_int64_option_1_none() -> Array[Int64?] {
  test_array_output_int64_option_1_none()
}

pub fn __modus_test_array_output_int64_option_1_min() -> Array[Int64?] {
  test_array_output_int64_option_1_min()
}

pub fn __modus_test_array_output_int64_option_1_max() -> Array[Int64?] {
  test_array_output_int64_option_1_max()
}

pub fn __modus_test_array_output_int64_option_2() -> Array[Int64?] {
  test_array_output_int64_option_2()
}

pub fn __modus_test_array_output_int64_option_3() -> Array[Int64?] {
  test_array_output_int64_option_3()
}

pub fn __modus_test_array_output_int64_option_4() -> Array[Int64?] {
  test_array_output_int64_option_4()
}

pub fn __modus_test_array_input_string(val : Array[String]) -> Unit!Error {
  try test_array_input_string!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_string() -> Array[String] {
  test_array_output_string()
}

pub fn __modus_test_array_input_string_none(val : Array[String]?) -> Unit!Error {
  try test_array_input_string_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_string_none() -> Array[String]? {
  test_array_output_string_none()
}

pub fn __modus_test_array_input_string_empty(val : Array[String]) -> Unit!Error {
  try test_array_input_string_empty!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_string_empty() -> Array[String] {
  test_array_output_string_empty()
}

pub fn __modus_test_array_input_string_option(val : Array[String?]) -> Unit!Error {
  try test_array_input_string_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_string_option() -> Array[String?] {
  test_array_output_string_option()
}

pub fn __modus_test2d_array_input_string(val : Array[Array[String]]) -> Unit!Error {
  try test2d_array_input_string!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test2d_array_output_string() -> Array[Array[String]] {
  test2d_array_output_string()
}

pub fn __modus_test2d_array_input_string_none(val : Array[Array[String]]?) -> Unit!Error {
  try test2d_array_input_string_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test2d_array_output_string_none() -> Array[Array[String]]? {
  test2d_array_output_string_none()
}

pub fn __modus_test2d_array_input_string_empty(val : Array[Array[String]]) -> Unit!Error {
  try test2d_array_input_string_empty!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test2d_array_output_string_empty() -> Array[Array[String]] {
  test2d_array_output_string_empty()
}

pub fn __modus_test2d_array_input_string_inner_empty(val : Array[Array[String]]) -> Unit!Error {
  try test2d_array_input_string_inner_empty!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test2d_array_output_string_inner_empty() -> Array[Array[String]] {
  test2d_array_output_string_inner_empty()
}

pub fn __modus_test2d_array_input_string_inner_none(val : Array[Array[String]?]) -> Unit!Error {
  try test2d_array_input_string_inner_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test2d_array_output_string_inner_none() -> Array[Array[String]?] {
  test2d_array_output_string_inner_none()
}

pub fn __modus_test_array_input_uint_empty(val : Array[UInt]) -> Unit!Error {
  try test_array_input_uint_empty!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_uint_0() -> Array[UInt] {
  test_array_output_uint_0()
}

pub fn __modus_test_array_output_uint_1() -> Array[UInt] {
  test_array_output_uint_1()
}

pub fn __modus_test_array_output_uint_1_min() -> Array[UInt] {
  test_array_output_uint_1_min()
}

pub fn __modus_test_array_output_uint_1_max() -> Array[UInt] {
  test_array_output_uint_1_max()
}

pub fn __modus_test_array_output_uint_2() -> Array[UInt] {
  test_array_output_uint_2()
}

pub fn __modus_test_array_output_uint_3() -> Array[UInt] {
  test_array_output_uint_3()
}

pub fn __modus_test_array_output_uint_4() -> Array[UInt] {
  test_array_output_uint_4()
}

pub fn __modus_test_array_output_uint_option_0() -> Array[UInt?] {
  test_array_output_uint_option_0()
}

pub fn __modus_test_array_output_uint_option_1_none() -> Array[UInt?] {
  test_array_output_uint_option_1_none()
}

pub fn __modus_test_array_output_uint_option_1_min() -> Array[UInt?] {
  test_array_output_uint_option_1_min()
}

pub fn __modus_test_array_output_uint_option_1_max() -> Array[UInt?] {
  test_array_output_uint_option_1_max()
}

pub fn __modus_test_array_output_uint_option_2() -> Array[UInt?] {
  test_array_output_uint_option_2()
}

pub fn __modus_test_array_output_uint_option_3() -> Array[UInt?] {
  test_array_output_uint_option_3()
}

pub fn __modus_test_array_output_uint_option_4() -> Array[UInt?] {
  test_array_output_uint_option_4()
}

pub fn __modus_test_array_input_uint_option(val : Array[UInt?]) -> Unit!Error {
  try test_array_input_uint_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_uint_option() -> Array[UInt?] {
  test_array_output_uint_option()
}

pub fn __modus_test_array_input_uint16_empty(val : Array[UInt16]) -> Unit!Error {
  try test_array_input_uint16_empty!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_uint16_0() -> Array[UInt16] {
  test_array_output_uint16_0()
}

pub fn __modus_test_array_output_uint16_1() -> Array[UInt16] {
  test_array_output_uint16_1()
}

pub fn __modus_test_array_output_uint16_1_min() -> Array[UInt16] {
  test_array_output_uint16_1_min()
}

pub fn __modus_test_array_output_uint16_1_max() -> Array[UInt16] {
  test_array_output_uint16_1_max()
}

pub fn __modus_test_array_output_uint16_2() -> Array[UInt16] {
  test_array_output_uint16_2()
}

pub fn __modus_test_array_output_uint16_3() -> Array[UInt16] {
  test_array_output_uint16_3()
}

pub fn __modus_test_array_output_uint16_4() -> Array[UInt16] {
  test_array_output_uint16_4()
}

pub fn __modus_test_array_output_uint16_option_0() -> Array[UInt16?] {
  test_array_output_uint16_option_0()
}

pub fn __modus_test_array_output_uint16_option_1_none() -> Array[UInt16?] {
  test_array_output_uint16_option_1_none()
}

pub fn __modus_test_array_output_uint16_option_1_min() -> Array[UInt16?] {
  test_array_output_uint16_option_1_min()
}

pub fn __modus_test_array_output_uint16_option_1_max() -> Array[UInt16?] {
  test_array_output_uint16_option_1_max()
}

pub fn __modus_test_array_output_uint16_option_2() -> Array[UInt16?] {
  test_array_output_uint16_option_2()
}

pub fn __modus_test_array_output_uint16_option_3() -> Array[UInt16?] {
  test_array_output_uint16_option_3()
}

pub fn __modus_test_array_output_uint16_option_4() -> Array[UInt16?] {
  test_array_output_uint16_option_4()
}

pub fn __modus_test_array_input_uint16_option(val : Array[UInt16?]) -> Unit!Error {
  try test_array_input_uint16_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_array_output_uint16_option() -> Array[UInt16?] {
  test_array_output_uint16_option()
}

pub fn __modus_test_array_output_uint64_0() -> Array[UInt64] {
  test_array_output_uint64_0()
}

pub fn __modus_test_array_output_uint64_1() -> Array[UInt64] {
  test_array_output_uint64_1()
}

pub fn __modus_test_array_output_uint64_1_min() -> Array[UInt64] {
  test_array_output_uint64_1_min()
}

pub fn __modus_test_array_output_uint64_1_max() -> Array[UInt64] {
  test_array_output_uint64_1_max()
}

pub fn __modus_test_array_output_uint64_2() -> Array[UInt64] {
  test_array_output_uint64_2()
}

pub fn __modus_test_array_output_uint64_3() -> Array[UInt64] {
  test_array_output_uint64_3()
}

pub fn __modus_test_array_output_uint64_4() -> Array[UInt64] {
  test_array_output_uint64_4()
}

pub fn __modus_test_array_output_uint64_option_0() -> Array[UInt64?] {
  test_array_output_uint64_option_0()
}

pub fn __modus_test_array_output_uint64_option_1_none() -> Array[UInt64?] {
  test_array_output_uint64_option_1_none()
}

pub fn __modus_test_array_output_uint64_option_1_min() -> Array[UInt64?] {
  test_array_output_uint64_option_1_min()
}

pub fn __modus_test_array_output_uint64_option_1_max() -> Array[UInt64?] {
  test_array_output_uint64_option_1_max()
}

pub fn __modus_test_array_output_uint64_option_2() -> Array[UInt64?] {
  test_array_output_uint64_option_2()
}

pub fn __modus_test_array_output_uint64_option_3() -> Array[UInt64?] {
  test_array_output_uint64_option_3()
}

pub fn __modus_test_array_output_uint64_option_4() -> Array[UInt64?] {
  test_array_output_uint64_option_4()
}

pub fn __modus_test_fixedarray_input0_byte(val : FixedArray[Byte]) -> Unit!Error {
  try test_fixedarray_input0_byte!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output0_byte() -> FixedArray[Byte] {
  test_fixedarray_output0_byte()
}

pub fn __modus_test_fixedarray_input0_string(val : FixedArray[String]) -> Unit!Error {
  try test_fixedarray_input0_string!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output0_string() -> FixedArray[String] {
  test_fixedarray_output0_string()
}

pub fn __modus_test_fixedarray_input0_string_option(val : FixedArray[String?]) -> Unit!Error {
  try test_fixedarray_input0_string_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output0_string_option() -> FixedArray[String?] {
  test_fixedarray_output0_string_option()
}

pub fn __modus_test_fixedarray_input0_int_option(val : FixedArray[Int?]) -> Unit!Error {
  try test_fixedarray_input0_int_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output0_int_option() -> FixedArray[Int?] {
  test_fixedarray_output0_int_option()
}

pub fn __modus_test_fixedarray_input1_byte(val : FixedArray[Byte]) -> Unit!Error {
  try test_fixedarray_input1_byte!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output1_byte() -> FixedArray[Byte] {
  test_fixedarray_output1_byte()
}

pub fn __modus_test_fixedarray_input1_string(val : FixedArray[String]) -> Unit!Error {
  try test_fixedarray_input1_string!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output1_string() -> FixedArray[String] {
  test_fixedarray_output1_string()
}

pub fn __modus_test_fixedarray_input1_string_option(val : FixedArray[String?]) -> Unit!Error {
  try test_fixedarray_input1_string_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output1_string_option() -> FixedArray[String?] {
  test_fixedarray_output1_string_option()
}

pub fn __modus_get_string_option_fixedarray1() -> FixedArray[String?] {
  get_string_option_fixedarray1()
}

pub fn __modus_test_fixedarray_input1_int_option(val : FixedArray[Int?]) -> Unit!Error {
  try test_fixedarray_input1_int_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output1_int_option() -> FixedArray[Int?] {
  test_fixedarray_output1_int_option()
}

pub fn __modus_get_int_option_fixedarray1() -> FixedArray[Int?] {
  get_int_option_fixedarray1()
}

pub fn __modus_test_fixedarray_input2_byte(val : FixedArray[Byte]) -> Unit!Error {
  try test_fixedarray_input2_byte!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output2_byte() -> FixedArray[Byte] {
  test_fixedarray_output2_byte()
}

pub fn __modus_test_fixedarray_input2_string(val : FixedArray[String]) -> Unit!Error {
  try test_fixedarray_input2_string!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output2_string() -> FixedArray[String] {
  test_fixedarray_output2_string()
}

pub fn __modus_test_fixedarray_input2_string_option(val : FixedArray[String?]) -> Unit!Error {
  try test_fixedarray_input2_string_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output2_string_option() -> FixedArray[String?] {
  test_fixedarray_output2_string_option()
}

pub fn __modus_get_string_option_array2() -> FixedArray[String?] {
  get_string_option_array2()
}

pub fn __modus_test_fixedarray_input2_int_option(val : FixedArray[Int?]) -> Unit!Error {
  try test_fixedarray_input2_int_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output2_int_option() -> FixedArray[Int?] {
  test_fixedarray_output2_int_option()
}

pub fn __modus_get_int_ptr_array2() -> FixedArray[Int?] {
  get_int_ptr_array2()
}

pub fn __modus_test_fixedarray_input2_struct(val : FixedArray[TestStruct2]) -> Unit!Error {
  try test_fixedarray_input2_struct!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output2_struct() -> FixedArray[TestStruct2] {
  test_fixedarray_output2_struct()
}

pub fn __modus_test_fixedarray_input2_struct_option(val : FixedArray[TestStruct2?]) -> Unit!Error {
  try test_fixedarray_input2_struct_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output2_struct_option() -> FixedArray[TestStruct2?] {
  test_fixedarray_output2_struct_option()
}

pub fn __modus_test_fixedarray_input2_map(val : FixedArray[Map[String, String]]) -> Unit!Error {
  try test_fixedarray_input2_map!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output2_map() -> FixedArray[Map[String, String]] {
  test_fixedarray_output2_map()
}

pub fn __modus_test_fixedarray_input2_map_option(val : FixedArray[Map[String, String]?]) -> Unit!Error {
  try test_fixedarray_input2_map_option!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output2_map_option() -> FixedArray[Map[String, String]?] {
  test_fixedarray_output2_map_option()
}

pub fn __modus_get_map_array2() -> FixedArray[Map[String, String]] {
  get_map_array2()
}

pub fn __modus_get_map_ptr_array2() -> FixedArray[Map[String, String]?] {
  get_map_ptr_array2()
}

pub fn __modus_test_option_fixedarray_input1_int(val : FixedArray[Int]?) -> Unit!Error {
  try test_option_fixedarray_input1_int!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_option_fixedarray_output1_int() -> FixedArray[Int]? {
  test_option_fixedarray_output1_int()
}

pub fn __modus_get_option_int_fixedarray1() -> FixedArray[Int]? {
  get_option_int_fixedarray1()
}

pub fn __modus_test_option_fixedarray_input2_int(val : FixedArray[Int]?) -> Unit!Error {
  try test_option_fixedarray_input2_int!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_option_fixedarray_output2_int() -> FixedArray[Int]? {
  test_option_fixedarray_output2_int()
}

pub fn __modus_get_option_int_fixedarray2() -> FixedArray[Int]? {
  get_option_int_fixedarray2()
}

pub fn __modus_test_option_fixedarray_input1_string(val : FixedArray[String]?) -> Unit!Error {
  try test_option_fixedarray_input1_string!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_option_fixedarray_output1_string() -> FixedArray[String]? {
  test_option_fixedarray_output1_string()
}

pub fn __modus_get_option_string_fixedarray1() -> FixedArray[String]? {
  get_option_string_fixedarray1()
}

pub fn __modus_get_option_string_fixedarray2() -> FixedArray[String]? {
  get_option_string_fixedarray2()
}

pub fn __modus_test_option_fixedarray_input2_string(val : FixedArray[String]?) -> Unit!Error {
  try test_option_fixedarray_input2_string!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_option_fixedarray_output2_string() -> FixedArray[String]? {
  test_option_fixedarray_output2_string()
}

pub fn __modus_test_fixedarray_output_bool_0() -> FixedArray[Bool] {
  test_fixedarray_output_bool_0()
}

pub fn __modus_test_fixedarray_output_bool_1() -> FixedArray[Bool] {
  test_fixedarray_output_bool_1()
}

pub fn __modus_test_fixedarray_output_bool_2() -> FixedArray[Bool] {
  test_fixedarray_output_bool_2()
}

pub fn __modus_test_fixedarray_output_bool_3() -> FixedArray[Bool] {
  test_fixedarray_output_bool_3()
}

pub fn __modus_test_fixedarray_output_bool_4() -> FixedArray[Bool] {
  test_fixedarray_output_bool_4()
}

pub fn __modus_test_fixedarray_output_bool_option_0() -> FixedArray[Bool?] {
  test_fixedarray_output_bool_option_0()
}

pub fn __modus_test_fixedarray_output_bool_option_1_none() -> FixedArray[Bool?] {
  test_fixedarray_output_bool_option_1_none()
}

pub fn __modus_test_fixedarray_output_bool_option_1_false() -> FixedArray[Bool?] {
  test_fixedarray_output_bool_option_1_false()
}

pub fn __modus_test_fixedarray_output_bool_option_1_true() -> FixedArray[Bool?] {
  test_fixedarray_output_bool_option_1_true()
}

pub fn __modus_test_fixedarray_output_bool_option_2() -> FixedArray[Bool?] {
  test_fixedarray_output_bool_option_2()
}

pub fn __modus_test_fixedarray_output_bool_option_3() -> FixedArray[Bool?] {
  test_fixedarray_output_bool_option_3()
}

pub fn __modus_test_fixedarray_output_bool_option_4() -> FixedArray[Bool?] {
  test_fixedarray_output_bool_option_4()
}

pub fn __modus_test_fixedarray_input_bool_0(val : FixedArray[Bool]) -> Unit!Error {
  try test_fixedarray_input_bool_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_1(val : FixedArray[Bool]) -> Unit!Error {
  try test_fixedarray_input_bool_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_2(val : FixedArray[Bool]) -> Unit!Error {
  try test_fixedarray_input_bool_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_3(val : FixedArray[Bool]) -> Unit!Error {
  try test_fixedarray_input_bool_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_4(val : FixedArray[Bool]) -> Unit!Error {
  try test_fixedarray_input_bool_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_option_0(val : FixedArray[Bool?]) -> Unit!Error {
  try test_fixedarray_input_bool_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_option_1_none(val : FixedArray[Bool?]) -> Unit!Error {
  try test_fixedarray_input_bool_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_option_1_false(val : FixedArray[Bool?]) -> Unit!Error {
  try test_fixedarray_input_bool_option_1_false!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_option_1_true(val : FixedArray[Bool?]) -> Unit!Error {
  try test_fixedarray_input_bool_option_1_true!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_option_2(val : FixedArray[Bool?]) -> Unit!Error {
  try test_fixedarray_input_bool_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_option_3(val : FixedArray[Bool?]) -> Unit!Error {
  try test_fixedarray_input_bool_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_bool_option_4(val : FixedArray[Bool?]) -> Unit!Error {
  try test_fixedarray_input_bool_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_byte_0() -> FixedArray[Byte] {
  test_fixedarray_output_byte_0()
}

pub fn __modus_test_fixedarray_output_byte_1() -> FixedArray[Byte] {
  test_fixedarray_output_byte_1()
}

pub fn __modus_test_fixedarray_output_byte_2() -> FixedArray[Byte] {
  test_fixedarray_output_byte_2()
}

pub fn __modus_test_fixedarray_output_byte_3() -> FixedArray[Byte] {
  test_fixedarray_output_byte_3()
}

pub fn __modus_test_fixedarray_output_byte_4() -> FixedArray[Byte] {
  test_fixedarray_output_byte_4()
}

pub fn __modus_test_fixedarray_output_byte_option_0() -> FixedArray[Byte?] {
  test_fixedarray_output_byte_option_0()
}

pub fn __modus_test_fixedarray_output_byte_option_1() -> FixedArray[Byte?] {
  test_fixedarray_output_byte_option_1()
}

pub fn __modus_test_fixedarray_output_byte_option_2() -> FixedArray[Byte?] {
  test_fixedarray_output_byte_option_2()
}

pub fn __modus_test_fixedarray_output_byte_option_3() -> FixedArray[Byte?] {
  test_fixedarray_output_byte_option_3()
}

pub fn __modus_test_fixedarray_output_byte_option_4() -> FixedArray[Byte?] {
  test_fixedarray_output_byte_option_4()
}

pub fn __modus_test_fixedarray_input_byte_0(val : FixedArray[Byte]) -> Unit!Error {
  try test_fixedarray_input_byte_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_byte_1(val : FixedArray[Byte]) -> Unit!Error {
  try test_fixedarray_input_byte_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_byte_2(val : FixedArray[Byte]) -> Unit!Error {
  try test_fixedarray_input_byte_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_byte_3(val : FixedArray[Byte]) -> Unit!Error {
  try test_fixedarray_input_byte_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_byte_4(val : FixedArray[Byte]) -> Unit!Error {
  try test_fixedarray_input_byte_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_byte_option_0(val : FixedArray[Byte?]) -> Unit!Error {
  try test_fixedarray_input_byte_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_byte_option_1(val : FixedArray[Byte?]) -> Unit!Error {
  try test_fixedarray_input_byte_option_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_byte_option_2(val : FixedArray[Byte?]) -> Unit!Error {
  try test_fixedarray_input_byte_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_byte_option_3(val : FixedArray[Byte?]) -> Unit!Error {
  try test_fixedarray_input_byte_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_byte_option_4(val : FixedArray[Byte?]) -> Unit!Error {
  try test_fixedarray_input_byte_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_char_0() -> FixedArray[Char] {
  test_fixedarray_output_char_0()
}

pub fn __modus_test_fixedarray_output_char_1() -> FixedArray[Char] {
  test_fixedarray_output_char_1()
}

pub fn __modus_test_fixedarray_output_char_2() -> FixedArray[Char] {
  test_fixedarray_output_char_2()
}

pub fn __modus_test_fixedarray_output_char_3() -> FixedArray[Char] {
  test_fixedarray_output_char_3()
}

pub fn __modus_test_fixedarray_output_char_4() -> FixedArray[Char] {
  test_fixedarray_output_char_4()
}

pub fn __modus_test_fixedarray_output_char_option_0() -> FixedArray[Char?] {
  test_fixedarray_output_char_option_0()
}

pub fn __modus_test_fixedarray_output_char_option_1_none() -> FixedArray[Char?] {
  test_fixedarray_output_char_option_1_none()
}

pub fn __modus_test_fixedarray_output_char_option_1_some() -> FixedArray[Char?] {
  test_fixedarray_output_char_option_1_some()
}

pub fn __modus_test_fixedarray_output_char_option_2() -> FixedArray[Char?] {
  test_fixedarray_output_char_option_2()
}

pub fn __modus_test_fixedarray_output_char_option_3() -> FixedArray[Char?] {
  test_fixedarray_output_char_option_3()
}

pub fn __modus_test_fixedarray_output_char_option_4() -> FixedArray[Char?] {
  test_fixedarray_output_char_option_4()
}

pub fn __modus_test_fixedarray_input_char_0(val : FixedArray[Char]) -> Unit!Error {
  try test_fixedarray_input_char_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_char_1(val : FixedArray[Char]) -> Unit!Error {
  try test_fixedarray_input_char_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_char_2(val : FixedArray[Char]) -> Unit!Error {
  try test_fixedarray_input_char_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_char_3(val : FixedArray[Char]) -> Unit!Error {
  try test_fixedarray_input_char_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_char_4(val : FixedArray[Char]) -> Unit!Error {
  try test_fixedarray_input_char_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_char_option_0(val : FixedArray[Char?]) -> Unit!Error {
  try test_fixedarray_input_char_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_char_option_1_none(val : FixedArray[Char?]) -> Unit!Error {
  try test_fixedarray_input_char_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_char_option_1_some(val : FixedArray[Char?]) -> Unit!Error {
  try test_fixedarray_input_char_option_1_some!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_char_option_2(val : FixedArray[Char?]) -> Unit!Error {
  try test_fixedarray_input_char_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_char_option_3(val : FixedArray[Char?]) -> Unit!Error {
  try test_fixedarray_input_char_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_char_option_4(val : FixedArray[Char?]) -> Unit!Error {
  try test_fixedarray_input_char_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_double_0() -> FixedArray[Double] {
  test_fixedarray_output_double_0()
}

pub fn __modus_test_fixedarray_output_double_1() -> FixedArray[Double] {
  test_fixedarray_output_double_1()
}

pub fn __modus_test_fixedarray_output_double_2() -> FixedArray[Double] {
  test_fixedarray_output_double_2()
}

pub fn __modus_test_fixedarray_output_double_3() -> FixedArray[Double] {
  test_fixedarray_output_double_3()
}

pub fn __modus_test_fixedarray_output_double_4() -> FixedArray[Double] {
  test_fixedarray_output_double_4()
}

pub fn __modus_test_fixedarray_output_double_option_0() -> FixedArray[Double?] {
  test_fixedarray_output_double_option_0()
}

pub fn __modus_test_fixedarray_output_double_option_1_none() -> FixedArray[Double?] {
  test_fixedarray_output_double_option_1_none()
}

pub fn __modus_test_fixedarray_output_double_option_1_some() -> FixedArray[Double?] {
  test_fixedarray_output_double_option_1_some()
}

pub fn __modus_test_fixedarray_output_double_option_2() -> FixedArray[Double?] {
  test_fixedarray_output_double_option_2()
}

pub fn __modus_test_fixedarray_output_double_option_3() -> FixedArray[Double?] {
  test_fixedarray_output_double_option_3()
}

pub fn __modus_test_fixedarray_output_double_option_4() -> FixedArray[Double?] {
  test_fixedarray_output_double_option_4()
}

pub fn __modus_test_fixedarray_input_double_0(val : FixedArray[Double]) -> Unit!Error {
  try test_fixedarray_input_double_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_double_1(val : FixedArray[Double]) -> Unit!Error {
  try test_fixedarray_input_double_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_double_2(val : FixedArray[Double]) -> Unit!Error {
  try test_fixedarray_input_double_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_double_3(val : FixedArray[Double]) -> Unit!Error {
  try test_fixedarray_input_double_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_double_4(val : FixedArray[Double]) -> Unit!Error {
  try test_fixedarray_input_double_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_double_option_0(val : FixedArray[Double?]) -> Unit!Error {
  try test_fixedarray_input_double_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_double_option_1_none(val : FixedArray[Double?]) -> Unit!Error {
  try test_fixedarray_input_double_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_double_option_1_some(val : FixedArray[Double?]) -> Unit!Error {
  try test_fixedarray_input_double_option_1_some!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_double_option_2(val : FixedArray[Double?]) -> Unit!Error {
  try test_fixedarray_input_double_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_double_option_3(val : FixedArray[Double?]) -> Unit!Error {
  try test_fixedarray_input_double_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_double_option_4(val : FixedArray[Double?]) -> Unit!Error {
  try test_fixedarray_input_double_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_float_0() -> FixedArray[Float] {
  test_fixedarray_output_float_0()
}

pub fn __modus_test_fixedarray_output_float_1() -> FixedArray[Float] {
  test_fixedarray_output_float_1()
}

pub fn __modus_test_fixedarray_output_float_2() -> FixedArray[Float] {
  test_fixedarray_output_float_2()
}

pub fn __modus_test_fixedarray_output_float_3() -> FixedArray[Float] {
  test_fixedarray_output_float_3()
}

pub fn __modus_test_fixedarray_output_float_4() -> FixedArray[Float] {
  test_fixedarray_output_float_4()
}

pub fn __modus_test_fixedarray_output_float_option_0() -> FixedArray[Float?] {
  test_fixedarray_output_float_option_0()
}

pub fn __modus_test_fixedarray_output_float_option_1_none() -> FixedArray[Float?] {
  test_fixedarray_output_float_option_1_none()
}

pub fn __modus_test_fixedarray_output_float_option_1_some() -> FixedArray[Float?] {
  test_fixedarray_output_float_option_1_some()
}

pub fn __modus_test_fixedarray_output_float_option_2() -> FixedArray[Float?] {
  test_fixedarray_output_float_option_2()
}

pub fn __modus_test_fixedarray_output_float_option_3() -> FixedArray[Float?] {
  test_fixedarray_output_float_option_3()
}

pub fn __modus_test_fixedarray_output_float_option_4() -> FixedArray[Float?] {
  test_fixedarray_output_float_option_4()
}

pub fn __modus_test_fixedarray_input_float_0(val : FixedArray[Float]) -> Unit!Error {
  try test_fixedarray_input_float_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_float_1(val : FixedArray[Float]) -> Unit!Error {
  try test_fixedarray_input_float_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_float_2(val : FixedArray[Float]) -> Unit!Error {
  try test_fixedarray_input_float_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_float_3(val : FixedArray[Float]) -> Unit!Error {
  try test_fixedarray_input_float_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_float_4(val : FixedArray[Float]) -> Unit!Error {
  try test_fixedarray_input_float_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_float_option_0(val : FixedArray[Float?]) -> Unit!Error {
  try test_fixedarray_input_float_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_float_option_1_none(val : FixedArray[Float?]) -> Unit!Error {
  try test_fixedarray_input_float_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_float_option_1_some(val : FixedArray[Float?]) -> Unit!Error {
  try test_fixedarray_input_float_option_1_some!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_float_option_2(val : FixedArray[Float?]) -> Unit!Error {
  try test_fixedarray_input_float_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_float_option_3(val : FixedArray[Float?]) -> Unit!Error {
  try test_fixedarray_input_float_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_float_option_4(val : FixedArray[Float?]) -> Unit!Error {
  try test_fixedarray_input_float_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_int_0() -> FixedArray[Int] {
  test_fixedarray_output_int_0()
}

pub fn __modus_test_fixedarray_output_int_1() -> FixedArray[Int] {
  test_fixedarray_output_int_1()
}

pub fn __modus_test_fixedarray_output_int_1_min() -> FixedArray[Int] {
  test_fixedarray_output_int_1_min()
}

pub fn __modus_test_fixedarray_output_int_1_max() -> FixedArray[Int] {
  test_fixedarray_output_int_1_max()
}

pub fn __modus_test_fixedarray_output_int_2() -> FixedArray[Int] {
  test_fixedarray_output_int_2()
}

pub fn __modus_test_fixedarray_output_int_3() -> FixedArray[Int] {
  test_fixedarray_output_int_3()
}

pub fn __modus_test_fixedarray_output_int_4() -> FixedArray[Int] {
  test_fixedarray_output_int_4()
}

pub fn __modus_test_fixedarray_output_int_option_0() -> FixedArray[Int?] {
  test_fixedarray_output_int_option_0()
}

pub fn __modus_test_fixedarray_output_int_option_1_none() -> FixedArray[Int?] {
  test_fixedarray_output_int_option_1_none()
}

pub fn __modus_test_fixedarray_output_int_option_1_min() -> FixedArray[Int?] {
  test_fixedarray_output_int_option_1_min()
}

pub fn __modus_test_fixedarray_output_int_option_1_max() -> FixedArray[Int?] {
  test_fixedarray_output_int_option_1_max()
}

pub fn __modus_test_fixedarray_output_int_option_2() -> FixedArray[Int?] {
  test_fixedarray_output_int_option_2()
}

pub fn __modus_test_fixedarray_output_int_option_3() -> FixedArray[Int?] {
  test_fixedarray_output_int_option_3()
}

pub fn __modus_test_fixedarray_output_int_option_4() -> FixedArray[Int?] {
  test_fixedarray_output_int_option_4()
}

pub fn __modus_test_fixedarray_input_int_0(val : FixedArray[Int]) -> Unit!Error {
  try test_fixedarray_input_int_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_1(val : FixedArray[Int]) -> Unit!Error {
  try test_fixedarray_input_int_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_1_min(val : FixedArray[Int]) -> Unit!Error {
  try test_fixedarray_input_int_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_1_max(val : FixedArray[Int]) -> Unit!Error {
  try test_fixedarray_input_int_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_2(val : FixedArray[Int]) -> Unit!Error {
  try test_fixedarray_input_int_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_3(val : FixedArray[Int]) -> Unit!Error {
  try test_fixedarray_input_int_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_4(val : FixedArray[Int]) -> Unit!Error {
  try test_fixedarray_input_int_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_option_0(val : FixedArray[Int?]) -> Unit!Error {
  try test_fixedarray_input_int_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_option_1_none(val : FixedArray[Int?]) -> Unit!Error {
  try test_fixedarray_input_int_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_option_1_min(val : FixedArray[Int?]) -> Unit!Error {
  try test_fixedarray_input_int_option_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_option_1_max(val : FixedArray[Int?]) -> Unit!Error {
  try test_fixedarray_input_int_option_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_option_2(val : FixedArray[Int?]) -> Unit!Error {
  try test_fixedarray_input_int_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_option_3(val : FixedArray[Int?]) -> Unit!Error {
  try test_fixedarray_input_int_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int_option_4(val : FixedArray[Int?]) -> Unit!Error {
  try test_fixedarray_input_int_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_int16_0() -> FixedArray[Int16] {
  test_fixedarray_output_int16_0()
}

pub fn __modus_test_fixedarray_output_int16_1() -> FixedArray[Int16] {
  test_fixedarray_output_int16_1()
}

pub fn __modus_test_fixedarray_output_int16_1_min() -> FixedArray[Int16] {
  test_fixedarray_output_int16_1_min()
}

pub fn __modus_test_fixedarray_output_int16_1_max() -> FixedArray[Int16] {
  test_fixedarray_output_int16_1_max()
}

pub fn __modus_test_fixedarray_output_int16_2() -> FixedArray[Int16] {
  test_fixedarray_output_int16_2()
}

pub fn __modus_test_fixedarray_output_int16_3() -> FixedArray[Int16] {
  test_fixedarray_output_int16_3()
}

pub fn __modus_test_fixedarray_output_int16_4() -> FixedArray[Int16] {
  test_fixedarray_output_int16_4()
}

pub fn __modus_test_fixedarray_output_int16_option_0() -> FixedArray[Int16?] {
  test_fixedarray_output_int16_option_0()
}

pub fn __modus_test_fixedarray_output_int16_option_1_none() -> FixedArray[Int16?] {
  test_fixedarray_output_int16_option_1_none()
}

pub fn __modus_test_fixedarray_output_int16_option_1_min() -> FixedArray[Int16?] {
  test_fixedarray_output_int16_option_1_min()
}

pub fn __modus_test_fixedarray_output_int16_option_1_max() -> FixedArray[Int16?] {
  test_fixedarray_output_int16_option_1_max()
}

pub fn __modus_test_fixedarray_output_int16_option_2() -> FixedArray[Int16?] {
  test_fixedarray_output_int16_option_2()
}

pub fn __modus_test_fixedarray_output_int16_option_3() -> FixedArray[Int16?] {
  test_fixedarray_output_int16_option_3()
}

pub fn __modus_test_fixedarray_output_int16_option_4() -> FixedArray[Int16?] {
  test_fixedarray_output_int16_option_4()
}

pub fn __modus_test_fixedarray_input_int16_0(val : FixedArray[Int16]) -> Unit!Error {
  try test_fixedarray_input_int16_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_1(val : FixedArray[Int16]) -> Unit!Error {
  try test_fixedarray_input_int16_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_1_min(val : FixedArray[Int16]) -> Unit!Error {
  try test_fixedarray_input_int16_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_1_max(val : FixedArray[Int16]) -> Unit!Error {
  try test_fixedarray_input_int16_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_2(val : FixedArray[Int16]) -> Unit!Error {
  try test_fixedarray_input_int16_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_3(val : FixedArray[Int16]) -> Unit!Error {
  try test_fixedarray_input_int16_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_4(val : FixedArray[Int16]) -> Unit!Error {
  try test_fixedarray_input_int16_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_option_0(val : FixedArray[Int16?]) -> Unit!Error {
  try test_fixedarray_input_int16_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_option_1_none(val : FixedArray[Int16?]) -> Unit!Error {
  try test_fixedarray_input_int16_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_option_1_min(val : FixedArray[Int16?]) -> Unit!Error {
  try test_fixedarray_input_int16_option_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_option_1_max(val : FixedArray[Int16?]) -> Unit!Error {
  try test_fixedarray_input_int16_option_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_option_2(val : FixedArray[Int16?]) -> Unit!Error {
  try test_fixedarray_input_int16_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_option_3(val : FixedArray[Int16?]) -> Unit!Error {
  try test_fixedarray_input_int16_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int16_option_4(val : FixedArray[Int16?]) -> Unit!Error {
  try test_fixedarray_input_int16_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_int64_0() -> FixedArray[Int64] {
  test_fixedarray_output_int64_0()
}

pub fn __modus_test_fixedarray_output_int64_1() -> FixedArray[Int64] {
  test_fixedarray_output_int64_1()
}

pub fn __modus_test_fixedarray_output_int64_1_min() -> FixedArray[Int64] {
  test_fixedarray_output_int64_1_min()
}

pub fn __modus_test_fixedarray_output_int64_1_max() -> FixedArray[Int64] {
  test_fixedarray_output_int64_1_max()
}

pub fn __modus_test_fixedarray_output_int64_2() -> FixedArray[Int64] {
  test_fixedarray_output_int64_2()
}

pub fn __modus_test_fixedarray_output_int64_3() -> FixedArray[Int64] {
  test_fixedarray_output_int64_3()
}

pub fn __modus_test_fixedarray_output_int64_4() -> FixedArray[Int64] {
  test_fixedarray_output_int64_4()
}

pub fn __modus_test_fixedarray_output_int64_option_0() -> FixedArray[Int64?] {
  test_fixedarray_output_int64_option_0()
}

pub fn __modus_test_fixedarray_output_int64_option_1_none() -> FixedArray[Int64?] {
  test_fixedarray_output_int64_option_1_none()
}

pub fn __modus_test_fixedarray_output_int64_option_1_min() -> FixedArray[Int64?] {
  test_fixedarray_output_int64_option_1_min()
}

pub fn __modus_test_fixedarray_output_int64_option_1_max() -> FixedArray[Int64?] {
  test_fixedarray_output_int64_option_1_max()
}

pub fn __modus_test_fixedarray_output_int64_option_2() -> FixedArray[Int64?] {
  test_fixedarray_output_int64_option_2()
}

pub fn __modus_test_fixedarray_output_int64_option_3() -> FixedArray[Int64?] {
  test_fixedarray_output_int64_option_3()
}

pub fn __modus_test_fixedarray_output_int64_option_4() -> FixedArray[Int64?] {
  test_fixedarray_output_int64_option_4()
}

pub fn __modus_test_fixedarray_input_int64_0(val : FixedArray[Int64]) -> Unit!Error {
  try test_fixedarray_input_int64_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_1(val : FixedArray[Int64]) -> Unit!Error {
  try test_fixedarray_input_int64_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_1_min(val : FixedArray[Int64]) -> Unit!Error {
  try test_fixedarray_input_int64_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_1_max(val : FixedArray[Int64]) -> Unit!Error {
  try test_fixedarray_input_int64_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_2(val : FixedArray[Int64]) -> Unit!Error {
  try test_fixedarray_input_int64_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_3(val : FixedArray[Int64]) -> Unit!Error {
  try test_fixedarray_input_int64_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_4(val : FixedArray[Int64]) -> Unit!Error {
  try test_fixedarray_input_int64_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_option_0(val : FixedArray[Int64?]) -> Unit!Error {
  try test_fixedarray_input_int64_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_option_1_none(val : FixedArray[Int64?]) -> Unit!Error {
  try test_fixedarray_input_int64_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_option_1_min(val : FixedArray[Int64?]) -> Unit!Error {
  try test_fixedarray_input_int64_option_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_option_1_max(val : FixedArray[Int64?]) -> Unit!Error {
  try test_fixedarray_input_int64_option_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_option_2(val : FixedArray[Int64?]) -> Unit!Error {
  try test_fixedarray_input_int64_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_option_3(val : FixedArray[Int64?]) -> Unit!Error {
  try test_fixedarray_input_int64_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_int64_option_4(val : FixedArray[Int64?]) -> Unit!Error {
  try test_fixedarray_input_int64_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_string_0() -> FixedArray[String] {
  test_fixedarray_output_string_0()
}

pub fn __modus_test_fixedarray_output_string_1() -> FixedArray[String] {
  test_fixedarray_output_string_1()
}

pub fn __modus_test_fixedarray_output_string_2() -> FixedArray[String] {
  test_fixedarray_output_string_2()
}

pub fn __modus_test_fixedarray_output_string_3() -> FixedArray[String] {
  test_fixedarray_output_string_3()
}

pub fn __modus_test_fixedarray_output_string_4() -> FixedArray[String] {
  test_fixedarray_output_string_4()
}

pub fn __modus_test_fixedarray_output_string_option_0() -> FixedArray[String?] {
  test_fixedarray_output_string_option_0()
}

pub fn __modus_test_fixedarray_output_string_option_1_none() -> FixedArray[String?] {
  test_fixedarray_output_string_option_1_none()
}

pub fn __modus_test_fixedarray_output_string_option_1_some() -> FixedArray[String?] {
  test_fixedarray_output_string_option_1_some()
}

pub fn __modus_test_fixedarray_output_string_option_2() -> FixedArray[String?] {
  test_fixedarray_output_string_option_2()
}

pub fn __modus_test_fixedarray_output_string_option_3() -> FixedArray[String?] {
  test_fixedarray_output_string_option_3()
}

pub fn __modus_test_fixedarray_output_string_option_4() -> FixedArray[String?] {
  test_fixedarray_output_string_option_4()
}

pub fn __modus_test_fixedarray_input_string_0(val : FixedArray[String]) -> Unit!Error {
  try test_fixedarray_input_string_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_string_1(val : FixedArray[String]) -> Unit!Error {
  try test_fixedarray_input_string_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_string_2(val : FixedArray[String]) -> Unit!Error {
  try test_fixedarray_input_string_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_string_3(val : FixedArray[String]) -> Unit!Error {
  try test_fixedarray_input_string_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_string_4(val : FixedArray[String]) -> Unit!Error {
  try test_fixedarray_input_string_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_string_option_0(val : FixedArray[String?]) -> Unit!Error {
  try test_fixedarray_input_string_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_string_option_1_none(val : FixedArray[String?]) -> Unit!Error {
  try test_fixedarray_input_string_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_string_option_1_some(val : FixedArray[String?]) -> Unit!Error {
  try test_fixedarray_input_string_option_1_some!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_string_option_2(val : FixedArray[String?]) -> Unit!Error {
  try test_fixedarray_input_string_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_string_option_3(val : FixedArray[String?]) -> Unit!Error {
  try test_fixedarray_input_string_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_string_option_4(val : FixedArray[String?]) -> Unit!Error {
  try test_fixedarray_input_string_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_uint_0() -> FixedArray[UInt] {
  test_fixedarray_output_uint_0()
}

pub fn __modus_test_fixedarray_output_uint_1() -> FixedArray[UInt] {
  test_fixedarray_output_uint_1()
}

pub fn __modus_test_fixedarray_output_uint_1_min() -> FixedArray[UInt] {
  test_fixedarray_output_uint_1_min()
}

pub fn __modus_test_fixedarray_output_uint_1_max() -> FixedArray[UInt] {
  test_fixedarray_output_uint_1_max()
}

pub fn __modus_test_fixedarray_output_uint_2() -> FixedArray[UInt] {
  test_fixedarray_output_uint_2()
}

pub fn __modus_test_fixedarray_output_uint_3() -> FixedArray[UInt] {
  test_fixedarray_output_uint_3()
}

pub fn __modus_test_fixedarray_output_uint_4() -> FixedArray[UInt] {
  test_fixedarray_output_uint_4()
}

pub fn __modus_test_fixedarray_output_uint_option_0() -> FixedArray[UInt?] {
  test_fixedarray_output_uint_option_0()
}

pub fn __modus_test_fixedarray_output_uint_option_1_none() -> FixedArray[UInt?] {
  test_fixedarray_output_uint_option_1_none()
}

pub fn __modus_test_fixedarray_output_uint_option_1_min() -> FixedArray[UInt?] {
  test_fixedarray_output_uint_option_1_min()
}

pub fn __modus_test_fixedarray_output_uint_option_1_max() -> FixedArray[UInt?] {
  test_fixedarray_output_uint_option_1_max()
}

pub fn __modus_test_fixedarray_output_uint_option_2() -> FixedArray[UInt?] {
  test_fixedarray_output_uint_option_2()
}

pub fn __modus_test_fixedarray_output_uint_option_3() -> FixedArray[UInt?] {
  test_fixedarray_output_uint_option_3()
}

pub fn __modus_test_fixedarray_output_uint_option_4() -> FixedArray[UInt?] {
  test_fixedarray_output_uint_option_4()
}

pub fn __modus_test_fixedarray_input_uint_0(val : FixedArray[UInt]) -> Unit!Error {
  try test_fixedarray_input_uint_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_1(val : FixedArray[UInt]) -> Unit!Error {
  try test_fixedarray_input_uint_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_1_min(val : FixedArray[UInt]) -> Unit!Error {
  try test_fixedarray_input_uint_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_1_max(val : FixedArray[UInt]) -> Unit!Error {
  try test_fixedarray_input_uint_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_2(val : FixedArray[UInt]) -> Unit!Error {
  try test_fixedarray_input_uint_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_3(val : FixedArray[UInt]) -> Unit!Error {
  try test_fixedarray_input_uint_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_4(val : FixedArray[UInt]) -> Unit!Error {
  try test_fixedarray_input_uint_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_option_0(val : FixedArray[UInt?]) -> Unit!Error {
  try test_fixedarray_input_uint_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_option_1_none(val : FixedArray[UInt?]) -> Unit!Error {
  try test_fixedarray_input_uint_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_option_1_min(val : FixedArray[UInt?]) -> Unit!Error {
  try test_fixedarray_input_uint_option_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_option_1_max(val : FixedArray[UInt?]) -> Unit!Error {
  try test_fixedarray_input_uint_option_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_option_2(val : FixedArray[UInt?]) -> Unit!Error {
  try test_fixedarray_input_uint_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_option_3(val : FixedArray[UInt?]) -> Unit!Error {
  try test_fixedarray_input_uint_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint_option_4(val : FixedArray[UInt?]) -> Unit!Error {
  try test_fixedarray_input_uint_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_uint16_0() -> FixedArray[UInt16] {
  test_fixedarray_output_uint16_0()
}

pub fn __modus_test_fixedarray_output_uint16_1() -> FixedArray[UInt16] {
  test_fixedarray_output_uint16_1()
}

pub fn __modus_test_fixedarray_output_uint16_1_min() -> FixedArray[UInt16] {
  test_fixedarray_output_uint16_1_min()
}

pub fn __modus_test_fixedarray_output_uint16_1_max() -> FixedArray[UInt16] {
  test_fixedarray_output_uint16_1_max()
}

pub fn __modus_test_fixedarray_output_uint16_2() -> FixedArray[UInt16] {
  test_fixedarray_output_uint16_2()
}

pub fn __modus_test_fixedarray_output_uint16_3() -> FixedArray[UInt16] {
  test_fixedarray_output_uint16_3()
}

pub fn __modus_test_fixedarray_output_uint16_4() -> FixedArray[UInt16] {
  test_fixedarray_output_uint16_4()
}

pub fn __modus_test_fixedarray_output_uint16_option_0() -> FixedArray[UInt16?] {
  test_fixedarray_output_uint16_option_0()
}

pub fn __modus_test_fixedarray_output_uint16_option_1_none() -> FixedArray[UInt16?] {
  test_fixedarray_output_uint16_option_1_none()
}

pub fn __modus_test_fixedarray_output_uint16_option_1_min() -> FixedArray[UInt16?] {
  test_fixedarray_output_uint16_option_1_min()
}

pub fn __modus_test_fixedarray_output_uint16_option_1_max() -> FixedArray[UInt16?] {
  test_fixedarray_output_uint16_option_1_max()
}

pub fn __modus_test_fixedarray_output_uint16_option_2() -> FixedArray[UInt16?] {
  test_fixedarray_output_uint16_option_2()
}

pub fn __modus_test_fixedarray_output_uint16_option_3() -> FixedArray[UInt16?] {
  test_fixedarray_output_uint16_option_3()
}

pub fn __modus_test_fixedarray_output_uint16_option_4() -> FixedArray[UInt16?] {
  test_fixedarray_output_uint16_option_4()
}

pub fn __modus_test_fixedarray_input_uint16_0(val : FixedArray[UInt16]) -> Unit!Error {
  try test_fixedarray_input_uint16_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_1(val : FixedArray[UInt16]) -> Unit!Error {
  try test_fixedarray_input_uint16_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_1_min(val : FixedArray[UInt16]) -> Unit!Error {
  try test_fixedarray_input_uint16_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_1_max(val : FixedArray[UInt16]) -> Unit!Error {
  try test_fixedarray_input_uint16_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_2(val : FixedArray[UInt16]) -> Unit!Error {
  try test_fixedarray_input_uint16_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_3(val : FixedArray[UInt16]) -> Unit!Error {
  try test_fixedarray_input_uint16_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_4(val : FixedArray[UInt16]) -> Unit!Error {
  try test_fixedarray_input_uint16_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_option_0(val : FixedArray[UInt16?]) -> Unit!Error {
  try test_fixedarray_input_uint16_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_option_1_none(val : FixedArray[UInt16?]) -> Unit!Error {
  try test_fixedarray_input_uint16_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_option_1_min(val : FixedArray[UInt16?]) -> Unit!Error {
  try test_fixedarray_input_uint16_option_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_option_1_max(val : FixedArray[UInt16?]) -> Unit!Error {
  try test_fixedarray_input_uint16_option_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_option_2(val : FixedArray[UInt16?]) -> Unit!Error {
  try test_fixedarray_input_uint16_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_option_3(val : FixedArray[UInt16?]) -> Unit!Error {
  try test_fixedarray_input_uint16_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint16_option_4(val : FixedArray[UInt16?]) -> Unit!Error {
  try test_fixedarray_input_uint16_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_output_uint64_0() -> FixedArray[UInt64] {
  test_fixedarray_output_uint64_0()
}

pub fn __modus_test_fixedarray_output_uint64_1() -> FixedArray[UInt64] {
  test_fixedarray_output_uint64_1()
}

pub fn __modus_test_fixedarray_output_uint64_1_min() -> FixedArray[UInt64] {
  test_fixedarray_output_uint64_1_min()
}

pub fn __modus_test_fixedarray_output_uint64_1_max() -> FixedArray[UInt64] {
  test_fixedarray_output_uint64_1_max()
}

pub fn __modus_test_fixedarray_output_uint64_2() -> FixedArray[UInt64] {
  test_fixedarray_output_uint64_2()
}

pub fn __modus_test_fixedarray_output_uint64_3() -> FixedArray[UInt64] {
  test_fixedarray_output_uint64_3()
}

pub fn __modus_test_fixedarray_output_uint64_4() -> FixedArray[UInt64] {
  test_fixedarray_output_uint64_4()
}

pub fn __modus_test_fixedarray_output_uint64_option_0() -> FixedArray[UInt64?] {
  test_fixedarray_output_uint64_option_0()
}

pub fn __modus_test_fixedarray_output_uint64_option_1_none() -> FixedArray[UInt64?] {
  test_fixedarray_output_uint64_option_1_none()
}

pub fn __modus_test_fixedarray_output_uint64_option_1_min() -> FixedArray[UInt64?] {
  test_fixedarray_output_uint64_option_1_min()
}

pub fn __modus_test_fixedarray_output_uint64_option_1_max() -> FixedArray[UInt64?] {
  test_fixedarray_output_uint64_option_1_max()
}

pub fn __modus_test_fixedarray_output_uint64_option_2() -> FixedArray[UInt64?] {
  test_fixedarray_output_uint64_option_2()
}

pub fn __modus_test_fixedarray_output_uint64_option_3() -> FixedArray[UInt64?] {
  test_fixedarray_output_uint64_option_3()
}

pub fn __modus_test_fixedarray_output_uint64_option_4() -> FixedArray[UInt64?] {
  test_fixedarray_output_uint64_option_4()
}

pub fn __modus_test_fixedarray_input_uint64_0(val : FixedArray[UInt64]) -> Unit!Error {
  try test_fixedarray_input_uint64_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_1(val : FixedArray[UInt64]) -> Unit!Error {
  try test_fixedarray_input_uint64_1!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_1_min(val : FixedArray[UInt64]) -> Unit!Error {
  try test_fixedarray_input_uint64_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_1_max(val : FixedArray[UInt64]) -> Unit!Error {
  try test_fixedarray_input_uint64_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_2(val : FixedArray[UInt64]) -> Unit!Error {
  try test_fixedarray_input_uint64_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_3(val : FixedArray[UInt64]) -> Unit!Error {
  try test_fixedarray_input_uint64_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_4(val : FixedArray[UInt64]) -> Unit!Error {
  try test_fixedarray_input_uint64_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_option_0(val : FixedArray[UInt64?]) -> Unit!Error {
  try test_fixedarray_input_uint64_option_0!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_option_1_none(val : FixedArray[UInt64?]) -> Unit!Error {
  try test_fixedarray_input_uint64_option_1_none!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_option_1_min(val : FixedArray[UInt64?]) -> Unit!Error {
  try test_fixedarray_input_uint64_option_1_min!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_option_1_max(val : FixedArray[UInt64?]) -> Unit!Error {
  try test_fixedarray_input_uint64_option_1_max!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_option_2(val : FixedArray[UInt64?]) -> Unit!Error {
  try test_fixedarray_input_uint64_option_2!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_option_3(val : FixedArray[UInt64?]) -> Unit!Error {
  try test_fixedarray_input_uint64_option_3!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_fixedarray_input_uint64_option_4(val : FixedArray[UInt64?]) -> Unit!Error {
  try test_fixedarray_input_uint64_option_4!(val) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_add(a : Int, b : Int) -> Int {
  add(a, b)
}

pub fn __modus_echo1(message : String) -> String {
  echo1(message)
}

pub fn __modus_echo2(message : String?) -> String {
  echo2(message)
}

pub fn __modus_echo3(message : String) -> String? {
  echo3(message)
}

pub fn __modus_echo4(message : String?) -> String? {
  echo4(message)
}

pub fn __modus_encode_strings1(items : Array[String]?) -> String? {
  encode_strings1(items)
}

pub fn __modus_encode_strings2(items : Array[String?]?) -> String? {
  encode_strings2(items)
}

pub fn __modus_test_http_response_headers(r : HttpResponse?) -> Unit {
  test_http_response_headers(r)
}

pub fn __modus_test_http_response_headers_output() -> HttpResponse? {
  test_http_response_headers_output()
}

pub fn __modus_test_http_headers(h : HttpHeaders) -> Unit {
  test_http_headers(h)
}

pub fn __modus_test_http_header_map(m : Map[String, HttpHeader?]) -> Unit {
  test_http_header_map(m)
}

pub fn __modus_test_http_header(h : HttpHeader?) -> Unit {
  test_http_header(h)
}

pub fn __modus_host_echo1(message : String) -> String {
  host_echo1(message)
}

pub fn __modus_host_echo2(message : String?) -> String {
  host_echo2(message)
}

pub fn __modus_host_echo3(message : String) -> String? {
  host_echo3(message)
}

pub fn __modus_host_echo4(message : String?) -> String? {
  host_echo4(message)
}

pub fn __modus_test_map_input_string_string(m : Map[String, String]) -> Unit!Error {
  try test_map_input_string_string!(m) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_map_option_input_string_string(m : Map[String, String]?) -> Unit!Error {
  try test_map_option_input_string_string!(m) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_map_output_string_string() -> Map[String, String] {
  test_map_output_string_string()
}

pub fn __modus_test_map_option_output_string_string() -> Map[String, String]? {
  test_map_option_output_string_string()
}

pub fn __modus_test_iterate_map_string_string(m : Map[String, String]) -> Unit {
  test_iterate_map_string_string(m)
}

pub fn __modus_test_map_lookup_string_string(m : Map[String, String], key : String) -> String {
  test_map_lookup_string_string(m, key)
}

pub fn __modus_test_struct_containing_map_input_string_string(s : TestStructWithMap) -> Unit!Error {
  try test_struct_containing_map_input_string_string!(s) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_containing_map_output_string_string() -> TestStructWithMap {
  test_struct_containing_map_output_string_string()
}

pub fn __modus_test_map_input_int_float(m : Map[Int, Float]) -> Unit!Error {
  try test_map_input_int_float!(m) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_map_output_int_float() -> Map[Int, Float] {
  test_map_output_int_float()
}

pub fn __modus_test_map_input_int_double(m : Map[Int, Double]) -> Unit!Error {
  try test_map_input_int_double!(m) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_map_output_int_double() -> Map[Int, Double] {
  test_map_output_int_double()
}

pub fn __modus_test_generate_map_string_string_output() -> Map[String, String] {
  test_generate_map_string_string_output()
}

pub fn __modus_test_bool_input_false(b : Bool) -> Unit!Error {
  try test_bool_input_false!(b) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_bool_input_true(b : Bool) -> Unit!Error {
  try test_bool_input_true!(b) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_bool_output_false() -> Bool {
  test_bool_output_false()
}

pub fn __modus_test_bool_output_true() -> Bool {
  test_bool_output_true()
}

pub fn __modus_test_bool_option_input_false(b : Bool?) -> Unit!Error {
  try test_bool_option_input_false!(b) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_bool_option_input_true(b : Bool?) -> Unit!Error {
  try test_bool_option_input_true!(b) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_bool_option_input_none(b : Bool?) -> Unit!Error {
  try test_bool_option_input_none!(b) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_bool_option_output_false() -> Bool? {
  test_bool_option_output_false()
}

pub fn __modus_test_bool_option_output_true() -> Bool? {
  test_bool_option_output_true()
}

pub fn __modus_test_bool_option_output_none() -> Bool? {
  test_bool_option_output_none()
}

pub fn __modus_test_byte_input_min(b : Byte) -> Unit!Error {
  try test_byte_input_min!(b) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_byte_input_max(b : Byte) -> Unit!Error {
  try test_byte_input_max!(b) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_byte_output_min() -> Byte {
  test_byte_output_min()
}

pub fn __modus_test_byte_output_max() -> Byte {
  test_byte_output_max()
}

pub fn __modus_test_byte_option_input_min(b : Byte?) -> Unit!Error {
  try test_byte_option_input_min!(b) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_byte_option_input_max(b : Byte?) -> Unit!Error {
  try test_byte_option_input_max!(b) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_byte_option_input_none(b : Byte?) -> Unit!Error {
  try test_byte_option_input_none!(b) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_byte_option_output_min() -> Byte? {
  test_byte_option_output_min()
}

pub fn __modus_test_byte_option_output_max() -> Byte? {
  test_byte_option_output_max()
}

pub fn __modus_test_byte_option_output_none() -> Byte? {
  test_byte_option_output_none()
}

pub fn __modus_test_char_input_min(c : Char) -> Unit!Error {
  try test_char_input_min!(c) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_char_input_max(c : Char) -> Unit!Error {
  try test_char_input_max!(c) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_char_output_min() -> Char {
  test_char_output_min()
}

pub fn __modus_test_char_output_max() -> Char {
  test_char_output_max()
}

pub fn __modus_test_char_option_input_min(c : Char?) -> Unit!Error {
  try test_char_option_input_min!(c) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_char_option_input_max(c : Char?) -> Unit!Error {
  try test_char_option_input_max!(c) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_char_option_input_none(c : Char?) -> Unit!Error {
  try test_char_option_input_none!(c) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_char_option_output_min() -> Char? {
  test_char_option_output_min()
}

pub fn __modus_test_char_option_output_max() -> Char? {
  test_char_option_output_max()
}

pub fn __modus_test_char_option_output_none() -> Char? {
  test_char_option_output_none()
}

pub fn __modus_test_double_input_min(n : Double) -> Unit!Error {
  try test_double_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_double_input_max(n : Double) -> Unit!Error {
  try test_double_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_double_output_min() -> Double {
  test_double_output_min()
}

pub fn __modus_test_double_output_max() -> Double {
  test_double_output_max()
}

pub fn __modus_test_double_option_input_min(n : Double?) -> Unit!Error {
  try test_double_option_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_double_option_input_max(n : Double?) -> Unit!Error {
  try test_double_option_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_double_option_input_none(n : Double?) -> Unit!Error {
  try test_double_option_input_none!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_double_option_output_min() -> Double? {
  test_double_option_output_min()
}

pub fn __modus_test_double_option_output_max() -> Double? {
  test_double_option_output_max()
}

pub fn __modus_test_double_option_output_none() -> Double? {
  test_double_option_output_none()
}

pub fn __modus_test_float_input_min(n : Float) -> Unit!Error {
  try test_float_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_float_input_max(n : Float) -> Unit!Error {
  try test_float_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_float_output_min() -> Float {
  test_float_output_min()
}

pub fn __modus_test_float_output_max() -> Float {
  test_float_output_max()
}

pub fn __modus_test_float_option_input_min(n : Float?) -> Unit!Error {
  try test_float_option_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_float_option_input_max(n : Float?) -> Unit!Error {
  try test_float_option_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_float_option_input_none(n : Float?) -> Unit!Error {
  try test_float_option_input_none!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_float_option_output_min() -> Float? {
  test_float_option_output_min()
}

pub fn __modus_test_float_option_output_max() -> Float? {
  test_float_option_output_max()
}

pub fn __modus_test_float_option_output_none() -> Float? {
  test_float_option_output_none()
}

pub fn __modus_test_int_input_min(n : Int) -> Unit!Error {
  try test_int_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int_input_max(n : Int) -> Unit!Error {
  try test_int_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int_output_min() -> Int {
  test_int_output_min()
}

pub fn __modus_test_int_output_max() -> Int {
  test_int_output_max()
}

pub fn __modus_test_int_option_input_min(n : Int?) -> Unit!Error {
  try test_int_option_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int_option_input_max(n : Int?) -> Unit!Error {
  try test_int_option_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int_option_input_none(n : Int?) -> Unit!Error {
  try test_int_option_input_none!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int_option_output_min() -> Int? {
  test_int_option_output_min()
}

pub fn __modus_test_int_option_output_max() -> Int? {
  test_int_option_output_max()
}

pub fn __modus_test_int_option_output_none() -> Int? {
  test_int_option_output_none()
}

pub fn __modus_test_int16_input_min(n : Int16) -> Unit!Error {
  try test_int16_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int16_input_max(n : Int16) -> Unit!Error {
  try test_int16_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int16_output_min() -> Int16 {
  test_int16_output_min()
}

pub fn __modus_test_int16_output_max() -> Int16 {
  test_int16_output_max()
}

pub fn __modus_test_int16_option_input_min(n : Int16?) -> Unit!Error {
  try test_int16_option_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int16_option_input_max(n : Int16?) -> Unit!Error {
  try test_int16_option_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int16_option_input_none(n : Int16?) -> Unit!Error {
  try test_int16_option_input_none!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int16_option_output_min() -> Int16? {
  test_int16_option_output_min()
}

pub fn __modus_test_int16_option_output_max() -> Int16? {
  test_int16_option_output_max()
}

pub fn __modus_test_int16_option_output_none() -> Int16? {
  test_int16_option_output_none()
}

pub fn __modus_test_int64_input_min(n : Int64) -> Unit!Error {
  try test_int64_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int64_input_max(n : Int64) -> Unit!Error {
  try test_int64_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int64_output_min() -> Int64 {
  test_int64_output_min()
}

pub fn __modus_test_int64_output_max() -> Int64 {
  test_int64_output_max()
}

pub fn __modus_test_int64_option_input_min(n : Int64?) -> Unit!Error {
  try test_int64_option_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int64_option_input_max(n : Int64?) -> Unit!Error {
  try test_int64_option_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int64_option_input_none(n : Int64?) -> Unit!Error {
  try test_int64_option_input_none!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_int64_option_output_min() -> Int64? {
  test_int64_option_output_min()
}

pub fn __modus_test_int64_option_output_max() -> Int64? {
  test_int64_option_output_max()
}

pub fn __modus_test_int64_option_output_none() -> Int64? {
  test_int64_option_output_none()
}

pub fn __modus_test_uint_input_min(n : UInt) -> Unit!Error {
  try test_uint_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint_input_max(n : UInt) -> Unit!Error {
  try test_uint_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint_output_min() -> UInt {
  test_uint_output_min()
}

pub fn __modus_test_uint_output_max() -> UInt {
  test_uint_output_max()
}

pub fn __modus_test_uint_option_input_min(n : UInt?) -> Unit!Error {
  try test_uint_option_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint_option_input_max(n : UInt?) -> Unit!Error {
  try test_uint_option_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint_option_input_none(n : UInt?) -> Unit!Error {
  try test_uint_option_input_none!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint_option_output_min() -> UInt? {
  test_uint_option_output_min()
}

pub fn __modus_test_uint_option_output_max() -> UInt? {
  test_uint_option_output_max()
}

pub fn __modus_test_uint_option_output_none() -> UInt? {
  test_uint_option_output_none()
}

pub fn __modus_test_uint16_input_min(n : UInt16) -> Unit!Error {
  try test_uint16_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint16_input_max(n : UInt16) -> Unit!Error {
  try test_uint16_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint16_output_min() -> UInt16 {
  test_uint16_output_min()
}

pub fn __modus_test_uint16_output_max() -> UInt16 {
  test_uint16_output_max()
}

pub fn __modus_test_uint16_option_input_min(n : UInt16?) -> Unit!Error {
  try test_uint16_option_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint16_option_input_max(n : UInt16?) -> Unit!Error {
  try test_uint16_option_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint16_option_input_none(n : UInt16?) -> Unit!Error {
  try test_uint16_option_input_none!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint16_option_output_min() -> UInt16? {
  test_uint16_option_output_min()
}

pub fn __modus_test_uint16_option_output_max() -> UInt16? {
  test_uint16_option_output_max()
}

pub fn __modus_test_uint16_option_output_none() -> UInt16? {
  test_uint16_option_output_none()
}

pub fn __modus_test_uint64_input_min(n : UInt64) -> Unit!Error {
  try test_uint64_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint64_input_max(n : UInt64) -> Unit!Error {
  try test_uint64_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint64_output_min() -> UInt64 {
  test_uint64_output_min()
}

pub fn __modus_test_uint64_output_max() -> UInt64 {
  test_uint64_output_max()
}

pub fn __modus_test_uint64_option_input_min(n : UInt64?) -> Unit!Error {
  try test_uint64_option_input_min!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint64_option_input_max(n : UInt64?) -> Unit!Error {
  try test_uint64_option_input_max!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint64_option_input_none(n : UInt64?) -> Unit!Error {
  try test_uint64_option_input_none!(n) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_uint64_option_output_min() -> UInt64? {
  test_uint64_option_output_min()
}

pub fn __modus_test_uint64_option_output_max() -> UInt64? {
  test_uint64_option_output_max()
}

pub fn __modus_test_uint64_option_output_none() -> UInt64? {
  test_uint64_option_output_none()
}

pub fn __modus_test_string_input(s : String) -> Unit!Error {
  try test_string_input!(s) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_string_option_input(s : String?) -> Unit!Error {
  try test_string_option_input!(s) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_string_option_input_none(s : String?) -> Unit!Error {
  try test_string_option_input_none!(s) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_string_output() -> String {
  test_string_output()
}

pub fn __modus_test_string_option_output() -> String? {
  test_string_option_output()
}

pub fn __modus_test_string_option_output_none() -> String? {
  test_string_option_output_none()
}

pub fn __modus_test_struct_input1(o : TestStruct1) -> Unit!Error {
  try test_struct_input1!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_input2(o : TestStruct2) -> Unit!Error {
  try test_struct_input2!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_input3(o : TestStruct3) -> Unit!Error {
  try test_struct_input3!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_input4(o : TestStruct4) -> Unit!Error {
  try test_struct_input4!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_input5(o : TestStruct5) -> Unit!Error {
  try test_struct_input5!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_input4_with_none(o : TestStruct4) -> Unit!Error {
  try test_struct_input4_with_none!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_smorgasbord_struct_input(o : TestSmorgasbordStruct) -> Unit!Error {
  try test_smorgasbord_struct_input!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_recursive_struct_input(r1 : TestRecursiveStruct) -> Unit!Error {
  try test_recursive_struct_input!(r1) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input1(o : TestStruct1?) -> Unit!Error {
  try test_struct_option_input1!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input2(o : TestStruct2?) -> Unit!Error {
  try test_struct_option_input2!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input3(o : TestStruct3?) -> Unit!Error {
  try test_struct_option_input3!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input4(o : TestStruct4?) -> Unit!Error {
  try test_struct_option_input4!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input5(o : TestStruct5?) -> Unit!Error {
  try test_struct_option_input5!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input4_with_none(o : TestStruct4?) -> Unit!Error {
  try test_struct_option_input4_with_none!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_recursive_struct_option_input(o : TestRecursiveStruct?) -> Unit!Error {
  try test_recursive_struct_option_input!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_smorgasbord_struct_option_input(o : TestSmorgasbordStruct?) -> Unit!Error {
  try test_smorgasbord_struct_option_input!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input1_none(o : TestStruct1?) -> Unit!Error {
  try test_struct_option_input1_none!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input2_none(o : TestStruct2?) -> Unit!Error {
  try test_struct_option_input2_none!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input3_none(o : TestStruct3?) -> Unit!Error {
  try test_struct_option_input3_none!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input4_none(o : TestStruct4?) -> Unit!Error {
  try test_struct_option_input4_none!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_option_input5_none(o : TestStruct5?) -> Unit!Error {
  try test_struct_option_input5_none!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_recursive_struct_option_input_none(o : TestRecursiveStruct?) -> Unit!Error {
  try test_recursive_struct_option_input_none!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_smorgasbord_struct_option_input_none(o : TestSmorgasbordStruct?) -> Unit!Error {
  try test_smorgasbord_struct_option_input_none!(o) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_struct_output1() -> TestStruct1 {
  test_struct_output1()
}

pub fn __modus_test_struct_output1_map() -> TestStruct1_map {
  test_struct_output1_map()
}

pub fn __modus_test_struct_output2() -> TestStruct2 {
  test_struct_output2()
}

pub fn __modus_test_struct_output2_map() -> TestStruct2_map {
  test_struct_output2_map()
}

pub fn __modus_test_struct_output3() -> TestStruct3 {
  test_struct_output3()
}

pub fn __modus_test_struct_output3_map() -> TestStruct3_map {
  test_struct_output3_map()
}

pub fn __modus_test_struct_output4() -> TestStruct4 {
  test_struct_output4()
}

pub fn __modus_test_struct_output4_map() -> TestStruct4_map {
  test_struct_output4_map()
}

pub fn __modus_test_struct_output5() -> TestStruct5 {
  test_struct_output5()
}

pub fn __modus_test_struct_output5_map() -> TestStruct5_map {
  test_struct_output5_map()
}

pub fn __modus_test_struct_output4_with_none() -> TestStruct4 {
  test_struct_output4_with_none()
}

pub fn __modus_test_struct_output4_map_with_none() -> TestStruct4_map {
  test_struct_output4_map_with_none()
}

pub fn __modus_test_recursive_struct_output() -> TestRecursiveStruct {
  test_recursive_struct_output()
}

pub fn __modus_test_recursive_struct_output_map() -> TestRecursiveStruct_map {
  test_recursive_struct_output_map()
}

pub fn __modus_test_smorgasbord_struct_output() -> TestSmorgasbordStruct {
  test_smorgasbord_struct_output()
}

pub fn __modus_test_smorgasbord_struct_output_map() -> TestSmorgasbordStruct_map {
  test_smorgasbord_struct_output_map()
}

pub fn __modus_test_struct_option_output1() -> TestStruct1? {
  test_struct_option_output1()
}

pub fn __modus_test_struct_option_output1_map() -> TestStruct1_map? {
  test_struct_option_output1_map()
}

pub fn __modus_test_struct_option_output2() -> TestStruct2? {
  test_struct_option_output2()
}

pub fn __modus_test_struct_option_output2_map() -> TestStruct2_map? {
  test_struct_option_output2_map()
}

pub fn __modus_test_struct_option_output3() -> TestStruct3? {
  test_struct_option_output3()
}

pub fn __modus_test_struct_option_output3_map() -> TestStruct3_map? {
  test_struct_option_output3_map()
}

pub fn __modus_test_struct_option_output4() -> TestStruct4? {
  test_struct_option_output4()
}

pub fn __modus_test_struct_option_output4_map() -> TestStruct4_map? {
  test_struct_option_output4_map()
}

pub fn __modus_test_struct_option_output5() -> TestStruct5? {
  test_struct_option_output5()
}

pub fn __modus_test_struct_option_output5_map() -> TestStruct5_map? {
  test_struct_option_output5_map()
}

pub fn __modus_test_struct_option_output4_with_none() -> TestStruct4? {
  test_struct_option_output4_with_none()
}

pub fn __modus_test_struct_option_output4_map_with_none() -> TestStruct4_map? {
  test_struct_option_output4_map_with_none()
}

pub fn __modus_test_recursive_struct_option_output() -> TestRecursiveStruct? {
  test_recursive_struct_option_output()
}

pub fn __modus_test_recursive_struct_option_output_map() -> TestRecursiveStruct_map? {
  test_recursive_struct_option_output_map()
}

pub fn __modus_test_smorgasbord_struct_option_output() -> TestSmorgasbordStruct? {
  test_smorgasbord_struct_option_output()
}

pub fn __modus_test_smorgasbord_struct_option_output_map() -> TestSmorgasbordStruct_map? {
  test_smorgasbord_struct_option_output_map()
}

pub fn __modus_test_struct_option_output1_none() -> TestStruct1? {
  test_struct_option_output1_none()
}

pub fn __modus_test_struct_option_output2_none() -> TestStruct2? {
  test_struct_option_output2_none()
}

pub fn __modus_test_struct_option_output3_none() -> TestStruct3? {
  test_struct_option_output3_none()
}

pub fn __modus_test_struct_option_output4_none() -> TestStruct4? {
  test_struct_option_output4_none()
}

pub fn __modus_test_struct_option_output5_none() -> TestStruct5? {
  test_struct_option_output5_none()
}

pub fn __modus_test_recursive_struct_option_output_none() -> TestRecursiveStruct? {
  test_recursive_struct_option_output_none()
}

pub fn __modus_test_smorgasbord_struct_option_output_none() -> TestSmorgasbordStruct? {
  test_smorgasbord_struct_option_output_none()
}

pub fn __modus_test_smorgasbord_struct_option_output_map_none() -> TestSmorgasbordStruct_map? {
  test_smorgasbord_struct_option_output_map_none()
}

pub fn __modus_test_time_input(t : @time.ZonedDateTime) -> Unit!Error {
  try test_time_input!(t) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_time_option_input(t : @time.ZonedDateTime?) -> Unit!Error {
  try test_time_option_input!(t) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_call_test_time_option_input_some() -> @time.ZonedDateTime? {
  call_test_time_option_input_some()
}

pub fn __modus_call_test_time_option_input_none() -> @time.ZonedDateTime? {
  call_test_time_option_input_none()
}

pub fn __modus_test_time_option_input_style2(t? : @time.ZonedDateTime) -> Unit!Error {
  try test_time_option_input_style2!(t?) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_time_option_input_none(t : @time.ZonedDateTime?) -> Unit!Error {
  try test_time_option_input_none!(t) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_time_option_input_none_style2(t? : @time.ZonedDateTime) -> Unit!Error {
  try test_time_option_input_none_style2!(t?) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_time_output() -> @time.ZonedDateTime {
  test_time_output()
}

pub fn __modus_test_time_option_output() -> @time.ZonedDateTime? {
  test_time_option_output()
}

pub fn __modus_test_time_option_output_none() -> @time.ZonedDateTime? {
  test_time_option_output_none()
}

pub fn __modus_test_duration_input(d : @time.Duration) -> Unit!Error {
  try test_duration_input!(d) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_duration_option_input(d : @time.Duration?) -> Unit!Error {
  try test_duration_option_input!(d) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_duration_option_input_style2(d? : @time.Duration) -> Unit!Error {
  try test_duration_option_input_style2!(d?) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_duration_option_input_none(d : @time.Duration?) -> Unit!Error {
  try test_duration_option_input_none!(d) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_duration_option_input_none_style2(d? : @time.Duration) -> Unit!Error {
  try test_duration_option_input_none_style2!(d?) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_test_duration_output() -> @time.Duration {
  test_duration_output()
}

pub fn __modus_test_duration_option_output() -> @time.Duration? {
  test_duration_option_output()
}

pub fn __modus_test_duration_option_output_none() -> @time.Duration? {
  test_duration_option_output_none()
}

pub fn __modus_test_tuple_output() -> (Int, Bool, String) {
  test_tuple_output()
}

`

var wantRuntimePreProcessHeader = `// Code generated by modus-moonbit-build. DO NOT EDIT.

`

var wantRuntimePreProcessMoonPkgJSON = `{
  "import": [
    "gmlewis/modus/pkg/console",
    "gmlewis/modus/wit/ffi",
    "moonbitlang/x/sys",
    "moonbitlang/x/time"
  ],
  "targets": {
    "debug-memory_notwasm.mbt": [
      "not",
      "wasm"
    ],
    "debug-memory_wasm.mbt": [
      "wasm"
    ],
    "imports_notwasm.mbt": [
      "not",
      "wasm"
    ],
    "imports_wasm.mbt": [
      "wasm"
    ],
    "modus_post_generated.mbt": [
      "wasm"
    ]
  },
  "link": {
    "wasm": {
      "exports": [
        "__modus_add:add",
        "__modus_call_test_time_option_input_none:call_test_time_option_input_none",
        "__modus_call_test_time_option_input_some:call_test_time_option_input_some",
        "__modus_echo1:echo1",
        "__modus_echo2:echo2",
        "__modus_echo3:echo3",
        "__modus_echo4:echo4",
        "__modus_encode_strings1:encode_strings1",
        "__modus_encode_strings2:encode_strings2",
        "__modus_get_int_option_fixedarray1:get_int_option_fixedarray1",
        "__modus_get_int_ptr_array2:get_int_ptr_array2",
        "__modus_get_map_array2:get_map_array2",
        "__modus_get_map_ptr_array2:get_map_ptr_array2",
        "__modus_get_option_int_fixedarray1:get_option_int_fixedarray1",
        "__modus_get_option_int_fixedarray2:get_option_int_fixedarray2",
        "__modus_get_option_string_fixedarray1:get_option_string_fixedarray1",
        "__modus_get_option_string_fixedarray2:get_option_string_fixedarray2",
        "__modus_get_string_option_array2:get_string_option_array2",
        "__modus_get_string_option_fixedarray1:get_string_option_fixedarray1",
        "__modus_host_echo1:host_echo1",
        "__modus_host_echo2:host_echo2",
        "__modus_host_echo3:host_echo3",
        "__modus_host_echo4:host_echo4",
        "__modus_modus_test.add:modus_test.add",
        "__modus_modus_test.echo1:modus_test.echo1",
        "__modus_modus_test.echo2:modus_test.echo2",
        "__modus_modus_test.echo3:modus_test.echo3",
        "__modus_modus_test.echo4:modus_test.echo4",
        "__modus_modus_test.encodeStrings1:modus_test.encodeStrings1",
        "__modus_modus_test.encodeStrings2:modus_test.encodeStrings2",
        "__modus_test2d_array_input_string:test2d_array_input_string",
        "__modus_test2d_array_input_string_empty:test2d_array_input_string_empty",
        "__modus_test2d_array_input_string_inner_empty:test2d_array_input_string_inner_empty",
        "__modus_test2d_array_input_string_inner_none:test2d_array_input_string_inner_none",
        "__modus_test2d_array_input_string_none:test2d_array_input_string_none",
        "__modus_test2d_array_output_string:test2d_array_output_string",
        "__modus_test2d_array_output_string_empty:test2d_array_output_string_empty",
        "__modus_test2d_array_output_string_inner_empty:test2d_array_output_string_inner_empty",
        "__modus_test2d_array_output_string_inner_none:test2d_array_output_string_inner_none",
        "__modus_test2d_array_output_string_none:test2d_array_output_string_none",
        "__modus_test_array_input_bool_0:test_array_input_bool_0",
        "__modus_test_array_input_bool_1:test_array_input_bool_1",
        "__modus_test_array_input_bool_2:test_array_input_bool_2",
        "__modus_test_array_input_bool_3:test_array_input_bool_3",
        "__modus_test_array_input_bool_4:test_array_input_bool_4",
        "__modus_test_array_input_bool_option_0:test_array_input_bool_option_0",
        "__modus_test_array_input_bool_option_1_false:test_array_input_bool_option_1_false",
        "__modus_test_array_input_bool_option_1_none:test_array_input_bool_option_1_none",
        "__modus_test_array_input_bool_option_1_true:test_array_input_bool_option_1_true",
        "__modus_test_array_input_bool_option_2:test_array_input_bool_option_2",
        "__modus_test_array_input_bool_option_3:test_array_input_bool_option_3",
        "__modus_test_array_input_bool_option_4:test_array_input_bool_option_4",
        "__modus_test_array_input_byte_0:test_array_input_byte_0",
        "__modus_test_array_input_byte_1:test_array_input_byte_1",
        "__modus_test_array_input_byte_2:test_array_input_byte_2",
        "__modus_test_array_input_byte_3:test_array_input_byte_3",
        "__modus_test_array_input_byte_4:test_array_input_byte_4",
        "__modus_test_array_input_byte_option_0:test_array_input_byte_option_0",
        "__modus_test_array_input_byte_option_1:test_array_input_byte_option_1",
        "__modus_test_array_input_byte_option_2:test_array_input_byte_option_2",
        "__modus_test_array_input_byte_option_3:test_array_input_byte_option_3",
        "__modus_test_array_input_byte_option_4:test_array_input_byte_option_4",
        "__modus_test_array_input_char_empty:test_array_input_char_empty",
        "__modus_test_array_input_char_option:test_array_input_char_option",
        "__modus_test_array_input_double_empty:test_array_input_double_empty",
        "__modus_test_array_input_double_option:test_array_input_double_option",
        "__modus_test_array_input_float_empty:test_array_input_float_empty",
        "__modus_test_array_input_float_option:test_array_input_float_option",
        "__modus_test_array_input_int16_empty:test_array_input_int16_empty",
        "__modus_test_array_input_int16_option:test_array_input_int16_option",
        "__modus_test_array_input_int_empty:test_array_input_int_empty",
        "__modus_test_array_input_int_option:test_array_input_int_option",
        "__modus_test_array_input_string:test_array_input_string",
        "__modus_test_array_input_string_empty:test_array_input_string_empty",
        "__modus_test_array_input_string_none:test_array_input_string_none",
        "__modus_test_array_input_string_option:test_array_input_string_option",
        "__modus_test_array_input_uint16_empty:test_array_input_uint16_empty",
        "__modus_test_array_input_uint16_option:test_array_input_uint16_option",
        "__modus_test_array_input_uint_empty:test_array_input_uint_empty",
        "__modus_test_array_input_uint_option:test_array_input_uint_option",
        "__modus_test_array_output_bool_0:test_array_output_bool_0",
        "__modus_test_array_output_bool_1:test_array_output_bool_1",
        "__modus_test_array_output_bool_2:test_array_output_bool_2",
        "__modus_test_array_output_bool_3:test_array_output_bool_3",
        "__modus_test_array_output_bool_4:test_array_output_bool_4",
        "__modus_test_array_output_bool_option_0:test_array_output_bool_option_0",
        "__modus_test_array_output_bool_option_1_false:test_array_output_bool_option_1_false",
        "__modus_test_array_output_bool_option_1_none:test_array_output_bool_option_1_none",
        "__modus_test_array_output_bool_option_1_true:test_array_output_bool_option_1_true",
        "__modus_test_array_output_bool_option_2:test_array_output_bool_option_2",
        "__modus_test_array_output_bool_option_3:test_array_output_bool_option_3",
        "__modus_test_array_output_bool_option_4:test_array_output_bool_option_4",
        "__modus_test_array_output_byte_0:test_array_output_byte_0",
        "__modus_test_array_output_byte_1:test_array_output_byte_1",
        "__modus_test_array_output_byte_2:test_array_output_byte_2",
        "__modus_test_array_output_byte_3:test_array_output_byte_3",
        "__modus_test_array_output_byte_4:test_array_output_byte_4",
        "__modus_test_array_output_byte_option_0:test_array_output_byte_option_0",
        "__modus_test_array_output_byte_option_1:test_array_output_byte_option_1",
        "__modus_test_array_output_byte_option_2:test_array_output_byte_option_2",
        "__modus_test_array_output_byte_option_3:test_array_output_byte_option_3",
        "__modus_test_array_output_byte_option_4:test_array_output_byte_option_4",
        "__modus_test_array_output_char_0:test_array_output_char_0",
        "__modus_test_array_output_char_1:test_array_output_char_1",
        "__modus_test_array_output_char_2:test_array_output_char_2",
        "__modus_test_array_output_char_3:test_array_output_char_3",
        "__modus_test_array_output_char_4:test_array_output_char_4",
        "__modus_test_array_output_char_option:test_array_output_char_option",
        "__modus_test_array_output_char_option_0:test_array_output_char_option_0",
        "__modus_test_array_output_char_option_1_none:test_array_output_char_option_1_none",
        "__modus_test_array_output_char_option_1_some:test_array_output_char_option_1_some",
        "__modus_test_array_output_char_option_2:test_array_output_char_option_2",
        "__modus_test_array_output_char_option_3:test_array_output_char_option_3",
        "__modus_test_array_output_char_option_4:test_array_output_char_option_4",
        "__modus_test_array_output_double_0:test_array_output_double_0",
        "__modus_test_array_output_double_1:test_array_output_double_1",
        "__modus_test_array_output_double_2:test_array_output_double_2",
        "__modus_test_array_output_double_3:test_array_output_double_3",
        "__modus_test_array_output_double_4:test_array_output_double_4",
        "__modus_test_array_output_double_option:test_array_output_double_option",
        "__modus_test_array_output_double_option_0:test_array_output_double_option_0",
        "__modus_test_array_output_double_option_1_none:test_array_output_double_option_1_none",
        "__modus_test_array_output_double_option_1_some:test_array_output_double_option_1_some",
        "__modus_test_array_output_double_option_2:test_array_output_double_option_2",
        "__modus_test_array_output_double_option_3:test_array_output_double_option_3",
        "__modus_test_array_output_double_option_4:test_array_output_double_option_4",
        "__modus_test_array_output_float_0:test_array_output_float_0",
        "__modus_test_array_output_float_1:test_array_output_float_1",
        "__modus_test_array_output_float_2:test_array_output_float_2",
        "__modus_test_array_output_float_3:test_array_output_float_3",
        "__modus_test_array_output_float_4:test_array_output_float_4",
        "__modus_test_array_output_float_option:test_array_output_float_option",
        "__modus_test_array_output_float_option_0:test_array_output_float_option_0",
        "__modus_test_array_output_float_option_1_none:test_array_output_float_option_1_none",
        "__modus_test_array_output_float_option_1_some:test_array_output_float_option_1_some",
        "__modus_test_array_output_float_option_2:test_array_output_float_option_2",
        "__modus_test_array_output_float_option_3:test_array_output_float_option_3",
        "__modus_test_array_output_float_option_4:test_array_output_float_option_4",
        "__modus_test_array_output_int16_0:test_array_output_int16_0",
        "__modus_test_array_output_int16_1:test_array_output_int16_1",
        "__modus_test_array_output_int16_1_max:test_array_output_int16_1_max",
        "__modus_test_array_output_int16_1_min:test_array_output_int16_1_min",
        "__modus_test_array_output_int16_2:test_array_output_int16_2",
        "__modus_test_array_output_int16_3:test_array_output_int16_3",
        "__modus_test_array_output_int16_4:test_array_output_int16_4",
        "__modus_test_array_output_int16_option:test_array_output_int16_option",
        "__modus_test_array_output_int16_option_0:test_array_output_int16_option_0",
        "__modus_test_array_output_int16_option_1_max:test_array_output_int16_option_1_max",
        "__modus_test_array_output_int16_option_1_min:test_array_output_int16_option_1_min",
        "__modus_test_array_output_int16_option_1_none:test_array_output_int16_option_1_none",
        "__modus_test_array_output_int16_option_2:test_array_output_int16_option_2",
        "__modus_test_array_output_int16_option_3:test_array_output_int16_option_3",
        "__modus_test_array_output_int16_option_4:test_array_output_int16_option_4",
        "__modus_test_array_output_int64_0:test_array_output_int64_0",
        "__modus_test_array_output_int64_1:test_array_output_int64_1",
        "__modus_test_array_output_int64_1_max:test_array_output_int64_1_max",
        "__modus_test_array_output_int64_1_min:test_array_output_int64_1_min",
        "__modus_test_array_output_int64_2:test_array_output_int64_2",
        "__modus_test_array_output_int64_3:test_array_output_int64_3",
        "__modus_test_array_output_int64_4:test_array_output_int64_4",
        "__modus_test_array_output_int64_option_0:test_array_output_int64_option_0",
        "__modus_test_array_output_int64_option_1_max:test_array_output_int64_option_1_max",
        "__modus_test_array_output_int64_option_1_min:test_array_output_int64_option_1_min",
        "__modus_test_array_output_int64_option_1_none:test_array_output_int64_option_1_none",
        "__modus_test_array_output_int64_option_2:test_array_output_int64_option_2",
        "__modus_test_array_output_int64_option_3:test_array_output_int64_option_3",
        "__modus_test_array_output_int64_option_4:test_array_output_int64_option_4",
        "__modus_test_array_output_int_0:test_array_output_int_0",
        "__modus_test_array_output_int_1:test_array_output_int_1",
        "__modus_test_array_output_int_1_max:test_array_output_int_1_max",
        "__modus_test_array_output_int_1_min:test_array_output_int_1_min",
        "__modus_test_array_output_int_2:test_array_output_int_2",
        "__modus_test_array_output_int_3:test_array_output_int_3",
        "__modus_test_array_output_int_4:test_array_output_int_4",
        "__modus_test_array_output_int_option:test_array_output_int_option",
        "__modus_test_array_output_int_option_0:test_array_output_int_option_0",
        "__modus_test_array_output_int_option_1_max:test_array_output_int_option_1_max",
        "__modus_test_array_output_int_option_1_min:test_array_output_int_option_1_min",
        "__modus_test_array_output_int_option_1_none:test_array_output_int_option_1_none",
        "__modus_test_array_output_int_option_2:test_array_output_int_option_2",
        "__modus_test_array_output_int_option_3:test_array_output_int_option_3",
        "__modus_test_array_output_int_option_4:test_array_output_int_option_4",
        "__modus_test_array_output_string:test_array_output_string",
        "__modus_test_array_output_string_empty:test_array_output_string_empty",
        "__modus_test_array_output_string_none:test_array_output_string_none",
        "__modus_test_array_output_string_option:test_array_output_string_option",
        "__modus_test_array_output_uint16_0:test_array_output_uint16_0",
        "__modus_test_array_output_uint16_1:test_array_output_uint16_1",
        "__modus_test_array_output_uint16_1_max:test_array_output_uint16_1_max",
        "__modus_test_array_output_uint16_1_min:test_array_output_uint16_1_min",
        "__modus_test_array_output_uint16_2:test_array_output_uint16_2",
        "__modus_test_array_output_uint16_3:test_array_output_uint16_3",
        "__modus_test_array_output_uint16_4:test_array_output_uint16_4",
        "__modus_test_array_output_uint16_option:test_array_output_uint16_option",
        "__modus_test_array_output_uint16_option_0:test_array_output_uint16_option_0",
        "__modus_test_array_output_uint16_option_1_max:test_array_output_uint16_option_1_max",
        "__modus_test_array_output_uint16_option_1_min:test_array_output_uint16_option_1_min",
        "__modus_test_array_output_uint16_option_1_none:test_array_output_uint16_option_1_none",
        "__modus_test_array_output_uint16_option_2:test_array_output_uint16_option_2",
        "__modus_test_array_output_uint16_option_3:test_array_output_uint16_option_3",
        "__modus_test_array_output_uint16_option_4:test_array_output_uint16_option_4",
        "__modus_test_array_output_uint64_0:test_array_output_uint64_0",
        "__modus_test_array_output_uint64_1:test_array_output_uint64_1",
        "__modus_test_array_output_uint64_1_max:test_array_output_uint64_1_max",
        "__modus_test_array_output_uint64_1_min:test_array_output_uint64_1_min",
        "__modus_test_array_output_uint64_2:test_array_output_uint64_2",
        "__modus_test_array_output_uint64_3:test_array_output_uint64_3",
        "__modus_test_array_output_uint64_4:test_array_output_uint64_4",
        "__modus_test_array_output_uint64_option_0:test_array_output_uint64_option_0",
        "__modus_test_array_output_uint64_option_1_max:test_array_output_uint64_option_1_max",
        "__modus_test_array_output_uint64_option_1_min:test_array_output_uint64_option_1_min",
        "__modus_test_array_output_uint64_option_1_none:test_array_output_uint64_option_1_none",
        "__modus_test_array_output_uint64_option_2:test_array_output_uint64_option_2",
        "__modus_test_array_output_uint64_option_3:test_array_output_uint64_option_3",
        "__modus_test_array_output_uint64_option_4:test_array_output_uint64_option_4",
        "__modus_test_array_output_uint_0:test_array_output_uint_0",
        "__modus_test_array_output_uint_1:test_array_output_uint_1",
        "__modus_test_array_output_uint_1_max:test_array_output_uint_1_max",
        "__modus_test_array_output_uint_1_min:test_array_output_uint_1_min",
        "__modus_test_array_output_uint_2:test_array_output_uint_2",
        "__modus_test_array_output_uint_3:test_array_output_uint_3",
        "__modus_test_array_output_uint_4:test_array_output_uint_4",
        "__modus_test_array_output_uint_option:test_array_output_uint_option",
        "__modus_test_array_output_uint_option_0:test_array_output_uint_option_0",
        "__modus_test_array_output_uint_option_1_max:test_array_output_uint_option_1_max",
        "__modus_test_array_output_uint_option_1_min:test_array_output_uint_option_1_min",
        "__modus_test_array_output_uint_option_1_none:test_array_output_uint_option_1_none",
        "__modus_test_array_output_uint_option_2:test_array_output_uint_option_2",
        "__modus_test_array_output_uint_option_3:test_array_output_uint_option_3",
        "__modus_test_array_output_uint_option_4:test_array_output_uint_option_4",
        "__modus_test_bool_input_false:test_bool_input_false",
        "__modus_test_bool_input_true:test_bool_input_true",
        "__modus_test_bool_option_input_false:test_bool_option_input_false",
        "__modus_test_bool_option_input_none:test_bool_option_input_none",
        "__modus_test_bool_option_input_true:test_bool_option_input_true",
        "__modus_test_bool_option_output_false:test_bool_option_output_false",
        "__modus_test_bool_option_output_none:test_bool_option_output_none",
        "__modus_test_bool_option_output_true:test_bool_option_output_true",
        "__modus_test_bool_output_false:test_bool_output_false",
        "__modus_test_bool_output_true:test_bool_output_true",
        "__modus_test_byte_input_max:test_byte_input_max",
        "__modus_test_byte_input_min:test_byte_input_min",
        "__modus_test_byte_option_input_max:test_byte_option_input_max",
        "__modus_test_byte_option_input_min:test_byte_option_input_min",
        "__modus_test_byte_option_input_none:test_byte_option_input_none",
        "__modus_test_byte_option_output_max:test_byte_option_output_max",
        "__modus_test_byte_option_output_min:test_byte_option_output_min",
        "__modus_test_byte_option_output_none:test_byte_option_output_none",
        "__modus_test_byte_output_max:test_byte_output_max",
        "__modus_test_byte_output_min:test_byte_output_min",
        "__modus_test_char_input_max:test_char_input_max",
        "__modus_test_char_input_min:test_char_input_min",
        "__modus_test_char_option_input_max:test_char_option_input_max",
        "__modus_test_char_option_input_min:test_char_option_input_min",
        "__modus_test_char_option_input_none:test_char_option_input_none",
        "__modus_test_char_option_output_max:test_char_option_output_max",
        "__modus_test_char_option_output_min:test_char_option_output_min",
        "__modus_test_char_option_output_none:test_char_option_output_none",
        "__modus_test_char_output_max:test_char_output_max",
        "__modus_test_char_output_min:test_char_output_min",
        "__modus_test_double_input_max:test_double_input_max",
        "__modus_test_double_input_min:test_double_input_min",
        "__modus_test_double_option_input_max:test_double_option_input_max",
        "__modus_test_double_option_input_min:test_double_option_input_min",
        "__modus_test_double_option_input_none:test_double_option_input_none",
        "__modus_test_double_option_output_max:test_double_option_output_max",
        "__modus_test_double_option_output_min:test_double_option_output_min",
        "__modus_test_double_option_output_none:test_double_option_output_none",
        "__modus_test_double_output_max:test_double_output_max",
        "__modus_test_double_output_min:test_double_output_min",
        "__modus_test_duration_input:test_duration_input",
        "__modus_test_duration_option_input:test_duration_option_input",
        "__modus_test_duration_option_input_none:test_duration_option_input_none",
        "__modus_test_duration_option_input_none_style2:test_duration_option_input_none_style2",
        "__modus_test_duration_option_input_style2:test_duration_option_input_style2",
        "__modus_test_duration_option_output:test_duration_option_output",
        "__modus_test_duration_option_output_none:test_duration_option_output_none",
        "__modus_test_duration_output:test_duration_output",
        "__modus_test_fixedarray_input0_byte:test_fixedarray_input0_byte",
        "__modus_test_fixedarray_input0_int_option:test_fixedarray_input0_int_option",
        "__modus_test_fixedarray_input0_string:test_fixedarray_input0_string",
        "__modus_test_fixedarray_input0_string_option:test_fixedarray_input0_string_option",
        "__modus_test_fixedarray_input1_byte:test_fixedarray_input1_byte",
        "__modus_test_fixedarray_input1_int_option:test_fixedarray_input1_int_option",
        "__modus_test_fixedarray_input1_string:test_fixedarray_input1_string",
        "__modus_test_fixedarray_input1_string_option:test_fixedarray_input1_string_option",
        "__modus_test_fixedarray_input2_byte:test_fixedarray_input2_byte",
        "__modus_test_fixedarray_input2_int_option:test_fixedarray_input2_int_option",
        "__modus_test_fixedarray_input2_map:test_fixedarray_input2_map",
        "__modus_test_fixedarray_input2_map_option:test_fixedarray_input2_map_option",
        "__modus_test_fixedarray_input2_string:test_fixedarray_input2_string",
        "__modus_test_fixedarray_input2_string_option:test_fixedarray_input2_string_option",
        "__modus_test_fixedarray_input2_struct:test_fixedarray_input2_struct",
        "__modus_test_fixedarray_input2_struct_option:test_fixedarray_input2_struct_option",
        "__modus_test_fixedarray_input_bool_0:test_fixedarray_input_bool_0",
        "__modus_test_fixedarray_input_bool_1:test_fixedarray_input_bool_1",
        "__modus_test_fixedarray_input_bool_2:test_fixedarray_input_bool_2",
        "__modus_test_fixedarray_input_bool_3:test_fixedarray_input_bool_3",
        "__modus_test_fixedarray_input_bool_4:test_fixedarray_input_bool_4",
        "__modus_test_fixedarray_input_bool_option_0:test_fixedarray_input_bool_option_0",
        "__modus_test_fixedarray_input_bool_option_1_false:test_fixedarray_input_bool_option_1_false",
        "__modus_test_fixedarray_input_bool_option_1_none:test_fixedarray_input_bool_option_1_none",
        "__modus_test_fixedarray_input_bool_option_1_true:test_fixedarray_input_bool_option_1_true",
        "__modus_test_fixedarray_input_bool_option_2:test_fixedarray_input_bool_option_2",
        "__modus_test_fixedarray_input_bool_option_3:test_fixedarray_input_bool_option_3",
        "__modus_test_fixedarray_input_bool_option_4:test_fixedarray_input_bool_option_4",
        "__modus_test_fixedarray_input_byte_0:test_fixedarray_input_byte_0",
        "__modus_test_fixedarray_input_byte_1:test_fixedarray_input_byte_1",
        "__modus_test_fixedarray_input_byte_2:test_fixedarray_input_byte_2",
        "__modus_test_fixedarray_input_byte_3:test_fixedarray_input_byte_3",
        "__modus_test_fixedarray_input_byte_4:test_fixedarray_input_byte_4",
        "__modus_test_fixedarray_input_byte_option_0:test_fixedarray_input_byte_option_0",
        "__modus_test_fixedarray_input_byte_option_1:test_fixedarray_input_byte_option_1",
        "__modus_test_fixedarray_input_byte_option_2:test_fixedarray_input_byte_option_2",
        "__modus_test_fixedarray_input_byte_option_3:test_fixedarray_input_byte_option_3",
        "__modus_test_fixedarray_input_byte_option_4:test_fixedarray_input_byte_option_4",
        "__modus_test_fixedarray_input_char_0:test_fixedarray_input_char_0",
        "__modus_test_fixedarray_input_char_1:test_fixedarray_input_char_1",
        "__modus_test_fixedarray_input_char_2:test_fixedarray_input_char_2",
        "__modus_test_fixedarray_input_char_3:test_fixedarray_input_char_3",
        "__modus_test_fixedarray_input_char_4:test_fixedarray_input_char_4",
        "__modus_test_fixedarray_input_char_option_0:test_fixedarray_input_char_option_0",
        "__modus_test_fixedarray_input_char_option_1_none:test_fixedarray_input_char_option_1_none",
        "__modus_test_fixedarray_input_char_option_1_some:test_fixedarray_input_char_option_1_some",
        "__modus_test_fixedarray_input_char_option_2:test_fixedarray_input_char_option_2",
        "__modus_test_fixedarray_input_char_option_3:test_fixedarray_input_char_option_3",
        "__modus_test_fixedarray_input_char_option_4:test_fixedarray_input_char_option_4",
        "__modus_test_fixedarray_input_double_0:test_fixedarray_input_double_0",
        "__modus_test_fixedarray_input_double_1:test_fixedarray_input_double_1",
        "__modus_test_fixedarray_input_double_2:test_fixedarray_input_double_2",
        "__modus_test_fixedarray_input_double_3:test_fixedarray_input_double_3",
        "__modus_test_fixedarray_input_double_4:test_fixedarray_input_double_4",
        "__modus_test_fixedarray_input_double_option_0:test_fixedarray_input_double_option_0",
        "__modus_test_fixedarray_input_double_option_1_none:test_fixedarray_input_double_option_1_none",
        "__modus_test_fixedarray_input_double_option_1_some:test_fixedarray_input_double_option_1_some",
        "__modus_test_fixedarray_input_double_option_2:test_fixedarray_input_double_option_2",
        "__modus_test_fixedarray_input_double_option_3:test_fixedarray_input_double_option_3",
        "__modus_test_fixedarray_input_double_option_4:test_fixedarray_input_double_option_4",
        "__modus_test_fixedarray_input_float_0:test_fixedarray_input_float_0",
        "__modus_test_fixedarray_input_float_1:test_fixedarray_input_float_1",
        "__modus_test_fixedarray_input_float_2:test_fixedarray_input_float_2",
        "__modus_test_fixedarray_input_float_3:test_fixedarray_input_float_3",
        "__modus_test_fixedarray_input_float_4:test_fixedarray_input_float_4",
        "__modus_test_fixedarray_input_float_option_0:test_fixedarray_input_float_option_0",
        "__modus_test_fixedarray_input_float_option_1_none:test_fixedarray_input_float_option_1_none",
        "__modus_test_fixedarray_input_float_option_1_some:test_fixedarray_input_float_option_1_some",
        "__modus_test_fixedarray_input_float_option_2:test_fixedarray_input_float_option_2",
        "__modus_test_fixedarray_input_float_option_3:test_fixedarray_input_float_option_3",
        "__modus_test_fixedarray_input_float_option_4:test_fixedarray_input_float_option_4",
        "__modus_test_fixedarray_input_int16_0:test_fixedarray_input_int16_0",
        "__modus_test_fixedarray_input_int16_1:test_fixedarray_input_int16_1",
        "__modus_test_fixedarray_input_int16_1_max:test_fixedarray_input_int16_1_max",
        "__modus_test_fixedarray_input_int16_1_min:test_fixedarray_input_int16_1_min",
        "__modus_test_fixedarray_input_int16_2:test_fixedarray_input_int16_2",
        "__modus_test_fixedarray_input_int16_3:test_fixedarray_input_int16_3",
        "__modus_test_fixedarray_input_int16_4:test_fixedarray_input_int16_4",
        "__modus_test_fixedarray_input_int16_option_0:test_fixedarray_input_int16_option_0",
        "__modus_test_fixedarray_input_int16_option_1_max:test_fixedarray_input_int16_option_1_max",
        "__modus_test_fixedarray_input_int16_option_1_min:test_fixedarray_input_int16_option_1_min",
        "__modus_test_fixedarray_input_int16_option_1_none:test_fixedarray_input_int16_option_1_none",
        "__modus_test_fixedarray_input_int16_option_2:test_fixedarray_input_int16_option_2",
        "__modus_test_fixedarray_input_int16_option_3:test_fixedarray_input_int16_option_3",
        "__modus_test_fixedarray_input_int16_option_4:test_fixedarray_input_int16_option_4",
        "__modus_test_fixedarray_input_int64_0:test_fixedarray_input_int64_0",
        "__modus_test_fixedarray_input_int64_1:test_fixedarray_input_int64_1",
        "__modus_test_fixedarray_input_int64_1_max:test_fixedarray_input_int64_1_max",
        "__modus_test_fixedarray_input_int64_1_min:test_fixedarray_input_int64_1_min",
        "__modus_test_fixedarray_input_int64_2:test_fixedarray_input_int64_2",
        "__modus_test_fixedarray_input_int64_3:test_fixedarray_input_int64_3",
        "__modus_test_fixedarray_input_int64_4:test_fixedarray_input_int64_4",
        "__modus_test_fixedarray_input_int64_option_0:test_fixedarray_input_int64_option_0",
        "__modus_test_fixedarray_input_int64_option_1_max:test_fixedarray_input_int64_option_1_max",
        "__modus_test_fixedarray_input_int64_option_1_min:test_fixedarray_input_int64_option_1_min",
        "__modus_test_fixedarray_input_int64_option_1_none:test_fixedarray_input_int64_option_1_none",
        "__modus_test_fixedarray_input_int64_option_2:test_fixedarray_input_int64_option_2",
        "__modus_test_fixedarray_input_int64_option_3:test_fixedarray_input_int64_option_3",
        "__modus_test_fixedarray_input_int64_option_4:test_fixedarray_input_int64_option_4",
        "__modus_test_fixedarray_input_int_0:test_fixedarray_input_int_0",
        "__modus_test_fixedarray_input_int_1:test_fixedarray_input_int_1",
        "__modus_test_fixedarray_input_int_1_max:test_fixedarray_input_int_1_max",
        "__modus_test_fixedarray_input_int_1_min:test_fixedarray_input_int_1_min",
        "__modus_test_fixedarray_input_int_2:test_fixedarray_input_int_2",
        "__modus_test_fixedarray_input_int_3:test_fixedarray_input_int_3",
        "__modus_test_fixedarray_input_int_4:test_fixedarray_input_int_4",
        "__modus_test_fixedarray_input_int_option_0:test_fixedarray_input_int_option_0",
        "__modus_test_fixedarray_input_int_option_1_max:test_fixedarray_input_int_option_1_max",
        "__modus_test_fixedarray_input_int_option_1_min:test_fixedarray_input_int_option_1_min",
        "__modus_test_fixedarray_input_int_option_1_none:test_fixedarray_input_int_option_1_none",
        "__modus_test_fixedarray_input_int_option_2:test_fixedarray_input_int_option_2",
        "__modus_test_fixedarray_input_int_option_3:test_fixedarray_input_int_option_3",
        "__modus_test_fixedarray_input_int_option_4:test_fixedarray_input_int_option_4",
        "__modus_test_fixedarray_input_string_0:test_fixedarray_input_string_0",
        "__modus_test_fixedarray_input_string_1:test_fixedarray_input_string_1",
        "__modus_test_fixedarray_input_string_2:test_fixedarray_input_string_2",
        "__modus_test_fixedarray_input_string_3:test_fixedarray_input_string_3",
        "__modus_test_fixedarray_input_string_4:test_fixedarray_input_string_4",
        "__modus_test_fixedarray_input_string_option_0:test_fixedarray_input_string_option_0",
        "__modus_test_fixedarray_input_string_option_1_none:test_fixedarray_input_string_option_1_none",
        "__modus_test_fixedarray_input_string_option_1_some:test_fixedarray_input_string_option_1_some",
        "__modus_test_fixedarray_input_string_option_2:test_fixedarray_input_string_option_2",
        "__modus_test_fixedarray_input_string_option_3:test_fixedarray_input_string_option_3",
        "__modus_test_fixedarray_input_string_option_4:test_fixedarray_input_string_option_4",
        "__modus_test_fixedarray_input_uint16_0:test_fixedarray_input_uint16_0",
        "__modus_test_fixedarray_input_uint16_1:test_fixedarray_input_uint16_1",
        "__modus_test_fixedarray_input_uint16_1_max:test_fixedarray_input_uint16_1_max",
        "__modus_test_fixedarray_input_uint16_1_min:test_fixedarray_input_uint16_1_min",
        "__modus_test_fixedarray_input_uint16_2:test_fixedarray_input_uint16_2",
        "__modus_test_fixedarray_input_uint16_3:test_fixedarray_input_uint16_3",
        "__modus_test_fixedarray_input_uint16_4:test_fixedarray_input_uint16_4",
        "__modus_test_fixedarray_input_uint16_option_0:test_fixedarray_input_uint16_option_0",
        "__modus_test_fixedarray_input_uint16_option_1_max:test_fixedarray_input_uint16_option_1_max",
        "__modus_test_fixedarray_input_uint16_option_1_min:test_fixedarray_input_uint16_option_1_min",
        "__modus_test_fixedarray_input_uint16_option_1_none:test_fixedarray_input_uint16_option_1_none",
        "__modus_test_fixedarray_input_uint16_option_2:test_fixedarray_input_uint16_option_2",
        "__modus_test_fixedarray_input_uint16_option_3:test_fixedarray_input_uint16_option_3",
        "__modus_test_fixedarray_input_uint16_option_4:test_fixedarray_input_uint16_option_4",
        "__modus_test_fixedarray_input_uint64_0:test_fixedarray_input_uint64_0",
        "__modus_test_fixedarray_input_uint64_1:test_fixedarray_input_uint64_1",
        "__modus_test_fixedarray_input_uint64_1_max:test_fixedarray_input_uint64_1_max",
        "__modus_test_fixedarray_input_uint64_1_min:test_fixedarray_input_uint64_1_min",
        "__modus_test_fixedarray_input_uint64_2:test_fixedarray_input_uint64_2",
        "__modus_test_fixedarray_input_uint64_3:test_fixedarray_input_uint64_3",
        "__modus_test_fixedarray_input_uint64_4:test_fixedarray_input_uint64_4",
        "__modus_test_fixedarray_input_uint64_option_0:test_fixedarray_input_uint64_option_0",
        "__modus_test_fixedarray_input_uint64_option_1_max:test_fixedarray_input_uint64_option_1_max",
        "__modus_test_fixedarray_input_uint64_option_1_min:test_fixedarray_input_uint64_option_1_min",
        "__modus_test_fixedarray_input_uint64_option_1_none:test_fixedarray_input_uint64_option_1_none",
        "__modus_test_fixedarray_input_uint64_option_2:test_fixedarray_input_uint64_option_2",
        "__modus_test_fixedarray_input_uint64_option_3:test_fixedarray_input_uint64_option_3",
        "__modus_test_fixedarray_input_uint64_option_4:test_fixedarray_input_uint64_option_4",
        "__modus_test_fixedarray_input_uint_0:test_fixedarray_input_uint_0",
        "__modus_test_fixedarray_input_uint_1:test_fixedarray_input_uint_1",
        "__modus_test_fixedarray_input_uint_1_max:test_fixedarray_input_uint_1_max",
        "__modus_test_fixedarray_input_uint_1_min:test_fixedarray_input_uint_1_min",
        "__modus_test_fixedarray_input_uint_2:test_fixedarray_input_uint_2",
        "__modus_test_fixedarray_input_uint_3:test_fixedarray_input_uint_3",
        "__modus_test_fixedarray_input_uint_4:test_fixedarray_input_uint_4",
        "__modus_test_fixedarray_input_uint_option_0:test_fixedarray_input_uint_option_0",
        "__modus_test_fixedarray_input_uint_option_1_max:test_fixedarray_input_uint_option_1_max",
        "__modus_test_fixedarray_input_uint_option_1_min:test_fixedarray_input_uint_option_1_min",
        "__modus_test_fixedarray_input_uint_option_1_none:test_fixedarray_input_uint_option_1_none",
        "__modus_test_fixedarray_input_uint_option_2:test_fixedarray_input_uint_option_2",
        "__modus_test_fixedarray_input_uint_option_3:test_fixedarray_input_uint_option_3",
        "__modus_test_fixedarray_input_uint_option_4:test_fixedarray_input_uint_option_4",
        "__modus_test_fixedarray_output0_byte:test_fixedarray_output0_byte",
        "__modus_test_fixedarray_output0_int_option:test_fixedarray_output0_int_option",
        "__modus_test_fixedarray_output0_string:test_fixedarray_output0_string",
        "__modus_test_fixedarray_output0_string_option:test_fixedarray_output0_string_option",
        "__modus_test_fixedarray_output1_byte:test_fixedarray_output1_byte",
        "__modus_test_fixedarray_output1_int_option:test_fixedarray_output1_int_option",
        "__modus_test_fixedarray_output1_string:test_fixedarray_output1_string",
        "__modus_test_fixedarray_output1_string_option:test_fixedarray_output1_string_option",
        "__modus_test_fixedarray_output2_byte:test_fixedarray_output2_byte",
        "__modus_test_fixedarray_output2_int_option:test_fixedarray_output2_int_option",
        "__modus_test_fixedarray_output2_map:test_fixedarray_output2_map",
        "__modus_test_fixedarray_output2_map_option:test_fixedarray_output2_map_option",
        "__modus_test_fixedarray_output2_string:test_fixedarray_output2_string",
        "__modus_test_fixedarray_output2_string_option:test_fixedarray_output2_string_option",
        "__modus_test_fixedarray_output2_struct:test_fixedarray_output2_struct",
        "__modus_test_fixedarray_output2_struct_option:test_fixedarray_output2_struct_option",
        "__modus_test_fixedarray_output_bool_0:test_fixedarray_output_bool_0",
        "__modus_test_fixedarray_output_bool_1:test_fixedarray_output_bool_1",
        "__modus_test_fixedarray_output_bool_2:test_fixedarray_output_bool_2",
        "__modus_test_fixedarray_output_bool_3:test_fixedarray_output_bool_3",
        "__modus_test_fixedarray_output_bool_4:test_fixedarray_output_bool_4",
        "__modus_test_fixedarray_output_bool_option_0:test_fixedarray_output_bool_option_0",
        "__modus_test_fixedarray_output_bool_option_1_false:test_fixedarray_output_bool_option_1_false",
        "__modus_test_fixedarray_output_bool_option_1_none:test_fixedarray_output_bool_option_1_none",
        "__modus_test_fixedarray_output_bool_option_1_true:test_fixedarray_output_bool_option_1_true",
        "__modus_test_fixedarray_output_bool_option_2:test_fixedarray_output_bool_option_2",
        "__modus_test_fixedarray_output_bool_option_3:test_fixedarray_output_bool_option_3",
        "__modus_test_fixedarray_output_bool_option_4:test_fixedarray_output_bool_option_4",
        "__modus_test_fixedarray_output_byte_0:test_fixedarray_output_byte_0",
        "__modus_test_fixedarray_output_byte_1:test_fixedarray_output_byte_1",
        "__modus_test_fixedarray_output_byte_2:test_fixedarray_output_byte_2",
        "__modus_test_fixedarray_output_byte_3:test_fixedarray_output_byte_3",
        "__modus_test_fixedarray_output_byte_4:test_fixedarray_output_byte_4",
        "__modus_test_fixedarray_output_byte_option_0:test_fixedarray_output_byte_option_0",
        "__modus_test_fixedarray_output_byte_option_1:test_fixedarray_output_byte_option_1",
        "__modus_test_fixedarray_output_byte_option_2:test_fixedarray_output_byte_option_2",
        "__modus_test_fixedarray_output_byte_option_3:test_fixedarray_output_byte_option_3",
        "__modus_test_fixedarray_output_byte_option_4:test_fixedarray_output_byte_option_4",
        "__modus_test_fixedarray_output_char_0:test_fixedarray_output_char_0",
        "__modus_test_fixedarray_output_char_1:test_fixedarray_output_char_1",
        "__modus_test_fixedarray_output_char_2:test_fixedarray_output_char_2",
        "__modus_test_fixedarray_output_char_3:test_fixedarray_output_char_3",
        "__modus_test_fixedarray_output_char_4:test_fixedarray_output_char_4",
        "__modus_test_fixedarray_output_char_option_0:test_fixedarray_output_char_option_0",
        "__modus_test_fixedarray_output_char_option_1_none:test_fixedarray_output_char_option_1_none",
        "__modus_test_fixedarray_output_char_option_1_some:test_fixedarray_output_char_option_1_some",
        "__modus_test_fixedarray_output_char_option_2:test_fixedarray_output_char_option_2",
        "__modus_test_fixedarray_output_char_option_3:test_fixedarray_output_char_option_3",
        "__modus_test_fixedarray_output_char_option_4:test_fixedarray_output_char_option_4",
        "__modus_test_fixedarray_output_double_0:test_fixedarray_output_double_0",
        "__modus_test_fixedarray_output_double_1:test_fixedarray_output_double_1",
        "__modus_test_fixedarray_output_double_2:test_fixedarray_output_double_2",
        "__modus_test_fixedarray_output_double_3:test_fixedarray_output_double_3",
        "__modus_test_fixedarray_output_double_4:test_fixedarray_output_double_4",
        "__modus_test_fixedarray_output_double_option_0:test_fixedarray_output_double_option_0",
        "__modus_test_fixedarray_output_double_option_1_none:test_fixedarray_output_double_option_1_none",
        "__modus_test_fixedarray_output_double_option_1_some:test_fixedarray_output_double_option_1_some",
        "__modus_test_fixedarray_output_double_option_2:test_fixedarray_output_double_option_2",
        "__modus_test_fixedarray_output_double_option_3:test_fixedarray_output_double_option_3",
        "__modus_test_fixedarray_output_double_option_4:test_fixedarray_output_double_option_4",
        "__modus_test_fixedarray_output_float_0:test_fixedarray_output_float_0",
        "__modus_test_fixedarray_output_float_1:test_fixedarray_output_float_1",
        "__modus_test_fixedarray_output_float_2:test_fixedarray_output_float_2",
        "__modus_test_fixedarray_output_float_3:test_fixedarray_output_float_3",
        "__modus_test_fixedarray_output_float_4:test_fixedarray_output_float_4",
        "__modus_test_fixedarray_output_float_option_0:test_fixedarray_output_float_option_0",
        "__modus_test_fixedarray_output_float_option_1_none:test_fixedarray_output_float_option_1_none",
        "__modus_test_fixedarray_output_float_option_1_some:test_fixedarray_output_float_option_1_some",
        "__modus_test_fixedarray_output_float_option_2:test_fixedarray_output_float_option_2",
        "__modus_test_fixedarray_output_float_option_3:test_fixedarray_output_float_option_3",
        "__modus_test_fixedarray_output_float_option_4:test_fixedarray_output_float_option_4",
        "__modus_test_fixedarray_output_int16_0:test_fixedarray_output_int16_0",
        "__modus_test_fixedarray_output_int16_1:test_fixedarray_output_int16_1",
        "__modus_test_fixedarray_output_int16_1_max:test_fixedarray_output_int16_1_max",
        "__modus_test_fixedarray_output_int16_1_min:test_fixedarray_output_int16_1_min",
        "__modus_test_fixedarray_output_int16_2:test_fixedarray_output_int16_2",
        "__modus_test_fixedarray_output_int16_3:test_fixedarray_output_int16_3",
        "__modus_test_fixedarray_output_int16_4:test_fixedarray_output_int16_4",
        "__modus_test_fixedarray_output_int16_option_0:test_fixedarray_output_int16_option_0",
        "__modus_test_fixedarray_output_int16_option_1_max:test_fixedarray_output_int16_option_1_max",
        "__modus_test_fixedarray_output_int16_option_1_min:test_fixedarray_output_int16_option_1_min",
        "__modus_test_fixedarray_output_int16_option_1_none:test_fixedarray_output_int16_option_1_none",
        "__modus_test_fixedarray_output_int16_option_2:test_fixedarray_output_int16_option_2",
        "__modus_test_fixedarray_output_int16_option_3:test_fixedarray_output_int16_option_3",
        "__modus_test_fixedarray_output_int16_option_4:test_fixedarray_output_int16_option_4",
        "__modus_test_fixedarray_output_int64_0:test_fixedarray_output_int64_0",
        "__modus_test_fixedarray_output_int64_1:test_fixedarray_output_int64_1",
        "__modus_test_fixedarray_output_int64_1_max:test_fixedarray_output_int64_1_max",
        "__modus_test_fixedarray_output_int64_1_min:test_fixedarray_output_int64_1_min",
        "__modus_test_fixedarray_output_int64_2:test_fixedarray_output_int64_2",
        "__modus_test_fixedarray_output_int64_3:test_fixedarray_output_int64_3",
        "__modus_test_fixedarray_output_int64_4:test_fixedarray_output_int64_4",
        "__modus_test_fixedarray_output_int64_option_0:test_fixedarray_output_int64_option_0",
        "__modus_test_fixedarray_output_int64_option_1_max:test_fixedarray_output_int64_option_1_max",
        "__modus_test_fixedarray_output_int64_option_1_min:test_fixedarray_output_int64_option_1_min",
        "__modus_test_fixedarray_output_int64_option_1_none:test_fixedarray_output_int64_option_1_none",
        "__modus_test_fixedarray_output_int64_option_2:test_fixedarray_output_int64_option_2",
        "__modus_test_fixedarray_output_int64_option_3:test_fixedarray_output_int64_option_3",
        "__modus_test_fixedarray_output_int64_option_4:test_fixedarray_output_int64_option_4",
        "__modus_test_fixedarray_output_int_0:test_fixedarray_output_int_0",
        "__modus_test_fixedarray_output_int_1:test_fixedarray_output_int_1",
        "__modus_test_fixedarray_output_int_1_max:test_fixedarray_output_int_1_max",
        "__modus_test_fixedarray_output_int_1_min:test_fixedarray_output_int_1_min",
        "__modus_test_fixedarray_output_int_2:test_fixedarray_output_int_2",
        "__modus_test_fixedarray_output_int_3:test_fixedarray_output_int_3",
        "__modus_test_fixedarray_output_int_4:test_fixedarray_output_int_4",
        "__modus_test_fixedarray_output_int_option_0:test_fixedarray_output_int_option_0",
        "__modus_test_fixedarray_output_int_option_1_max:test_fixedarray_output_int_option_1_max",
        "__modus_test_fixedarray_output_int_option_1_min:test_fixedarray_output_int_option_1_min",
        "__modus_test_fixedarray_output_int_option_1_none:test_fixedarray_output_int_option_1_none",
        "__modus_test_fixedarray_output_int_option_2:test_fixedarray_output_int_option_2",
        "__modus_test_fixedarray_output_int_option_3:test_fixedarray_output_int_option_3",
        "__modus_test_fixedarray_output_int_option_4:test_fixedarray_output_int_option_4",
        "__modus_test_fixedarray_output_string_0:test_fixedarray_output_string_0",
        "__modus_test_fixedarray_output_string_1:test_fixedarray_output_string_1",
        "__modus_test_fixedarray_output_string_2:test_fixedarray_output_string_2",
        "__modus_test_fixedarray_output_string_3:test_fixedarray_output_string_3",
        "__modus_test_fixedarray_output_string_4:test_fixedarray_output_string_4",
        "__modus_test_fixedarray_output_string_option_0:test_fixedarray_output_string_option_0",
        "__modus_test_fixedarray_output_string_option_1_none:test_fixedarray_output_string_option_1_none",
        "__modus_test_fixedarray_output_string_option_1_some:test_fixedarray_output_string_option_1_some",
        "__modus_test_fixedarray_output_string_option_2:test_fixedarray_output_string_option_2",
        "__modus_test_fixedarray_output_string_option_3:test_fixedarray_output_string_option_3",
        "__modus_test_fixedarray_output_string_option_4:test_fixedarray_output_string_option_4",
        "__modus_test_fixedarray_output_uint16_0:test_fixedarray_output_uint16_0",
        "__modus_test_fixedarray_output_uint16_1:test_fixedarray_output_uint16_1",
        "__modus_test_fixedarray_output_uint16_1_max:test_fixedarray_output_uint16_1_max",
        "__modus_test_fixedarray_output_uint16_1_min:test_fixedarray_output_uint16_1_min",
        "__modus_test_fixedarray_output_uint16_2:test_fixedarray_output_uint16_2",
        "__modus_test_fixedarray_output_uint16_3:test_fixedarray_output_uint16_3",
        "__modus_test_fixedarray_output_uint16_4:test_fixedarray_output_uint16_4",
        "__modus_test_fixedarray_output_uint16_option_0:test_fixedarray_output_uint16_option_0",
        "__modus_test_fixedarray_output_uint16_option_1_max:test_fixedarray_output_uint16_option_1_max",
        "__modus_test_fixedarray_output_uint16_option_1_min:test_fixedarray_output_uint16_option_1_min",
        "__modus_test_fixedarray_output_uint16_option_1_none:test_fixedarray_output_uint16_option_1_none",
        "__modus_test_fixedarray_output_uint16_option_2:test_fixedarray_output_uint16_option_2",
        "__modus_test_fixedarray_output_uint16_option_3:test_fixedarray_output_uint16_option_3",
        "__modus_test_fixedarray_output_uint16_option_4:test_fixedarray_output_uint16_option_4",
        "__modus_test_fixedarray_output_uint64_0:test_fixedarray_output_uint64_0",
        "__modus_test_fixedarray_output_uint64_1:test_fixedarray_output_uint64_1",
        "__modus_test_fixedarray_output_uint64_1_max:test_fixedarray_output_uint64_1_max",
        "__modus_test_fixedarray_output_uint64_1_min:test_fixedarray_output_uint64_1_min",
        "__modus_test_fixedarray_output_uint64_2:test_fixedarray_output_uint64_2",
        "__modus_test_fixedarray_output_uint64_3:test_fixedarray_output_uint64_3",
        "__modus_test_fixedarray_output_uint64_4:test_fixedarray_output_uint64_4",
        "__modus_test_fixedarray_output_uint64_option_0:test_fixedarray_output_uint64_option_0",
        "__modus_test_fixedarray_output_uint64_option_1_max:test_fixedarray_output_uint64_option_1_max",
        "__modus_test_fixedarray_output_uint64_option_1_min:test_fixedarray_output_uint64_option_1_min",
        "__modus_test_fixedarray_output_uint64_option_1_none:test_fixedarray_output_uint64_option_1_none",
        "__modus_test_fixedarray_output_uint64_option_2:test_fixedarray_output_uint64_option_2",
        "__modus_test_fixedarray_output_uint64_option_3:test_fixedarray_output_uint64_option_3",
        "__modus_test_fixedarray_output_uint64_option_4:test_fixedarray_output_uint64_option_4",
        "__modus_test_fixedarray_output_uint_0:test_fixedarray_output_uint_0",
        "__modus_test_fixedarray_output_uint_1:test_fixedarray_output_uint_1",
        "__modus_test_fixedarray_output_uint_1_max:test_fixedarray_output_uint_1_max",
        "__modus_test_fixedarray_output_uint_1_min:test_fixedarray_output_uint_1_min",
        "__modus_test_fixedarray_output_uint_2:test_fixedarray_output_uint_2",
        "__modus_test_fixedarray_output_uint_3:test_fixedarray_output_uint_3",
        "__modus_test_fixedarray_output_uint_4:test_fixedarray_output_uint_4",
        "__modus_test_fixedarray_output_uint_option_0:test_fixedarray_output_uint_option_0",
        "__modus_test_fixedarray_output_uint_option_1_max:test_fixedarray_output_uint_option_1_max",
        "__modus_test_fixedarray_output_uint_option_1_min:test_fixedarray_output_uint_option_1_min",
        "__modus_test_fixedarray_output_uint_option_1_none:test_fixedarray_output_uint_option_1_none",
        "__modus_test_fixedarray_output_uint_option_2:test_fixedarray_output_uint_option_2",
        "__modus_test_fixedarray_output_uint_option_3:test_fixedarray_output_uint_option_3",
        "__modus_test_fixedarray_output_uint_option_4:test_fixedarray_output_uint_option_4",
        "__modus_test_float_input_max:test_float_input_max",
        "__modus_test_float_input_min:test_float_input_min",
        "__modus_test_float_option_input_max:test_float_option_input_max",
        "__modus_test_float_option_input_min:test_float_option_input_min",
        "__modus_test_float_option_input_none:test_float_option_input_none",
        "__modus_test_float_option_output_max:test_float_option_output_max",
        "__modus_test_float_option_output_min:test_float_option_output_min",
        "__modus_test_float_option_output_none:test_float_option_output_none",
        "__modus_test_float_output_max:test_float_output_max",
        "__modus_test_float_output_min:test_float_output_min",
        "__modus_test_generate_map_string_string_output:test_generate_map_string_string_output",
        "__modus_test_http_header:test_http_header",
        "__modus_test_http_header_map:test_http_header_map",
        "__modus_test_http_headers:test_http_headers",
        "__modus_test_http_response_headers:test_http_response_headers",
        "__modus_test_http_response_headers_output:test_http_response_headers_output",
        "__modus_test_int16_input_max:test_int16_input_max",
        "__modus_test_int16_input_min:test_int16_input_min",
        "__modus_test_int16_option_input_max:test_int16_option_input_max",
        "__modus_test_int16_option_input_min:test_int16_option_input_min",
        "__modus_test_int16_option_input_none:test_int16_option_input_none",
        "__modus_test_int16_option_output_max:test_int16_option_output_max",
        "__modus_test_int16_option_output_min:test_int16_option_output_min",
        "__modus_test_int16_option_output_none:test_int16_option_output_none",
        "__modus_test_int16_output_max:test_int16_output_max",
        "__modus_test_int16_output_min:test_int16_output_min",
        "__modus_test_int64_input_max:test_int64_input_max",
        "__modus_test_int64_input_min:test_int64_input_min",
        "__modus_test_int64_option_input_max:test_int64_option_input_max",
        "__modus_test_int64_option_input_min:test_int64_option_input_min",
        "__modus_test_int64_option_input_none:test_int64_option_input_none",
        "__modus_test_int64_option_output_max:test_int64_option_output_max",
        "__modus_test_int64_option_output_min:test_int64_option_output_min",
        "__modus_test_int64_option_output_none:test_int64_option_output_none",
        "__modus_test_int64_output_max:test_int64_output_max",
        "__modus_test_int64_output_min:test_int64_output_min",
        "__modus_test_int_input_max:test_int_input_max",
        "__modus_test_int_input_min:test_int_input_min",
        "__modus_test_int_option_input_max:test_int_option_input_max",
        "__modus_test_int_option_input_min:test_int_option_input_min",
        "__modus_test_int_option_input_none:test_int_option_input_none",
        "__modus_test_int_option_output_max:test_int_option_output_max",
        "__modus_test_int_option_output_min:test_int_option_output_min",
        "__modus_test_int_option_output_none:test_int_option_output_none",
        "__modus_test_int_output_max:test_int_output_max",
        "__modus_test_int_output_min:test_int_output_min",
        "__modus_test_iterate_map_string_string:test_iterate_map_string_string",
        "__modus_test_map_input_int_double:test_map_input_int_double",
        "__modus_test_map_input_int_float:test_map_input_int_float",
        "__modus_test_map_input_string_string:test_map_input_string_string",
        "__modus_test_map_lookup_string_string:test_map_lookup_string_string",
        "__modus_test_map_option_input_string_string:test_map_option_input_string_string",
        "__modus_test_map_option_output_string_string:test_map_option_output_string_string",
        "__modus_test_map_output_int_double:test_map_output_int_double",
        "__modus_test_map_output_int_float:test_map_output_int_float",
        "__modus_test_map_output_string_string:test_map_output_string_string",
        "__modus_test_option_fixedarray_input1_int:test_option_fixedarray_input1_int",
        "__modus_test_option_fixedarray_input1_string:test_option_fixedarray_input1_string",
        "__modus_test_option_fixedarray_input2_int:test_option_fixedarray_input2_int",
        "__modus_test_option_fixedarray_input2_string:test_option_fixedarray_input2_string",
        "__modus_test_option_fixedarray_output1_int:test_option_fixedarray_output1_int",
        "__modus_test_option_fixedarray_output1_string:test_option_fixedarray_output1_string",
        "__modus_test_option_fixedarray_output2_int:test_option_fixedarray_output2_int",
        "__modus_test_option_fixedarray_output2_string:test_option_fixedarray_output2_string",
        "__modus_test_recursive_struct_input:test_recursive_struct_input",
        "__modus_test_recursive_struct_option_input:test_recursive_struct_option_input",
        "__modus_test_recursive_struct_option_input_none:test_recursive_struct_option_input_none",
        "__modus_test_recursive_struct_option_output:test_recursive_struct_option_output",
        "__modus_test_recursive_struct_option_output_map:test_recursive_struct_option_output_map",
        "__modus_test_recursive_struct_option_output_none:test_recursive_struct_option_output_none",
        "__modus_test_recursive_struct_output:test_recursive_struct_output",
        "__modus_test_recursive_struct_output_map:test_recursive_struct_output_map",
        "__modus_test_smorgasbord_struct_input:test_smorgasbord_struct_input",
        "__modus_test_smorgasbord_struct_option_input:test_smorgasbord_struct_option_input",
        "__modus_test_smorgasbord_struct_option_input_none:test_smorgasbord_struct_option_input_none",
        "__modus_test_smorgasbord_struct_option_output:test_smorgasbord_struct_option_output",
        "__modus_test_smorgasbord_struct_option_output_map:test_smorgasbord_struct_option_output_map",
        "__modus_test_smorgasbord_struct_option_output_map_none:test_smorgasbord_struct_option_output_map_none",
        "__modus_test_smorgasbord_struct_option_output_none:test_smorgasbord_struct_option_output_none",
        "__modus_test_smorgasbord_struct_output:test_smorgasbord_struct_output",
        "__modus_test_smorgasbord_struct_output_map:test_smorgasbord_struct_output_map",
        "__modus_test_string_input:test_string_input",
        "__modus_test_string_option_input:test_string_option_input",
        "__modus_test_string_option_input_none:test_string_option_input_none",
        "__modus_test_string_option_output:test_string_option_output",
        "__modus_test_string_option_output_none:test_string_option_output_none",
        "__modus_test_string_output:test_string_output",
        "__modus_test_struct_containing_map_input_string_string:test_struct_containing_map_input_string_string",
        "__modus_test_struct_containing_map_output_string_string:test_struct_containing_map_output_string_string",
        "__modus_test_struct_input1:test_struct_input1",
        "__modus_test_struct_input2:test_struct_input2",
        "__modus_test_struct_input3:test_struct_input3",
        "__modus_test_struct_input4:test_struct_input4",
        "__modus_test_struct_input4_with_none:test_struct_input4_with_none",
        "__modus_test_struct_input5:test_struct_input5",
        "__modus_test_struct_option_input1:test_struct_option_input1",
        "__modus_test_struct_option_input1_none:test_struct_option_input1_none",
        "__modus_test_struct_option_input2:test_struct_option_input2",
        "__modus_test_struct_option_input2_none:test_struct_option_input2_none",
        "__modus_test_struct_option_input3:test_struct_option_input3",
        "__modus_test_struct_option_input3_none:test_struct_option_input3_none",
        "__modus_test_struct_option_input4:test_struct_option_input4",
        "__modus_test_struct_option_input4_none:test_struct_option_input4_none",
        "__modus_test_struct_option_input4_with_none:test_struct_option_input4_with_none",
        "__modus_test_struct_option_input5:test_struct_option_input5",
        "__modus_test_struct_option_input5_none:test_struct_option_input5_none",
        "__modus_test_struct_option_output1:test_struct_option_output1",
        "__modus_test_struct_option_output1_map:test_struct_option_output1_map",
        "__modus_test_struct_option_output1_none:test_struct_option_output1_none",
        "__modus_test_struct_option_output2:test_struct_option_output2",
        "__modus_test_struct_option_output2_map:test_struct_option_output2_map",
        "__modus_test_struct_option_output2_none:test_struct_option_output2_none",
        "__modus_test_struct_option_output3:test_struct_option_output3",
        "__modus_test_struct_option_output3_map:test_struct_option_output3_map",
        "__modus_test_struct_option_output3_none:test_struct_option_output3_none",
        "__modus_test_struct_option_output4:test_struct_option_output4",
        "__modus_test_struct_option_output4_map:test_struct_option_output4_map",
        "__modus_test_struct_option_output4_map_with_none:test_struct_option_output4_map_with_none",
        "__modus_test_struct_option_output4_none:test_struct_option_output4_none",
        "__modus_test_struct_option_output4_with_none:test_struct_option_output4_with_none",
        "__modus_test_struct_option_output5:test_struct_option_output5",
        "__modus_test_struct_option_output5_map:test_struct_option_output5_map",
        "__modus_test_struct_option_output5_none:test_struct_option_output5_none",
        "__modus_test_struct_output1:test_struct_output1",
        "__modus_test_struct_output1_map:test_struct_output1_map",
        "__modus_test_struct_output2:test_struct_output2",
        "__modus_test_struct_output2_map:test_struct_output2_map",
        "__modus_test_struct_output3:test_struct_output3",
        "__modus_test_struct_output3_map:test_struct_output3_map",
        "__modus_test_struct_output4:test_struct_output4",
        "__modus_test_struct_output4_map:test_struct_output4_map",
        "__modus_test_struct_output4_map_with_none:test_struct_output4_map_with_none",
        "__modus_test_struct_output4_with_none:test_struct_output4_with_none",
        "__modus_test_struct_output5:test_struct_output5",
        "__modus_test_struct_output5_map:test_struct_output5_map",
        "__modus_test_time_input:test_time_input",
        "__modus_test_time_option_input:test_time_option_input",
        "__modus_test_time_option_input_none:test_time_option_input_none",
        "__modus_test_time_option_input_none_style2:test_time_option_input_none_style2",
        "__modus_test_time_option_input_style2:test_time_option_input_style2",
        "__modus_test_time_option_output:test_time_option_output",
        "__modus_test_time_option_output_none:test_time_option_output_none",
        "__modus_test_time_output:test_time_output",
        "__modus_test_tuple_output:test_tuple_output",
        "__modus_test_uint16_input_max:test_uint16_input_max",
        "__modus_test_uint16_input_min:test_uint16_input_min",
        "__modus_test_uint16_option_input_max:test_uint16_option_input_max",
        "__modus_test_uint16_option_input_min:test_uint16_option_input_min",
        "__modus_test_uint16_option_input_none:test_uint16_option_input_none",
        "__modus_test_uint16_option_output_max:test_uint16_option_output_max",
        "__modus_test_uint16_option_output_min:test_uint16_option_output_min",
        "__modus_test_uint16_option_output_none:test_uint16_option_output_none",
        "__modus_test_uint16_output_max:test_uint16_output_max",
        "__modus_test_uint16_output_min:test_uint16_output_min",
        "__modus_test_uint64_input_max:test_uint64_input_max",
        "__modus_test_uint64_input_min:test_uint64_input_min",
        "__modus_test_uint64_option_input_max:test_uint64_option_input_max",
        "__modus_test_uint64_option_input_min:test_uint64_option_input_min",
        "__modus_test_uint64_option_input_none:test_uint64_option_input_none",
        "__modus_test_uint64_option_output_max:test_uint64_option_output_max",
        "__modus_test_uint64_option_output_min:test_uint64_option_output_min",
        "__modus_test_uint64_option_output_none:test_uint64_option_output_none",
        "__modus_test_uint64_output_max:test_uint64_output_max",
        "__modus_test_uint64_output_min:test_uint64_output_min",
        "__modus_test_uint_input_max:test_uint_input_max",
        "__modus_test_uint_input_min:test_uint_input_min",
        "__modus_test_uint_option_input_max:test_uint_option_input_max",
        "__modus_test_uint_option_input_min:test_uint_option_input_min",
        "__modus_test_uint_option_input_none:test_uint_option_input_none",
        "__modus_test_uint_option_output_max:test_uint_option_output_max",
        "__modus_test_uint_option_output_min:test_uint_option_output_min",
        "__modus_test_uint_option_output_none:test_uint_option_output_none",
        "__modus_test_uint_output_max:test_uint_output_max",
        "__modus_test_uint_output_min:test_uint_output_min",
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
