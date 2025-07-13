// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_option_output_bool_none() -> Array[Bool]? {
  None
}

///|
pub fn test_array_option_output_bool_some_0() -> Array[Bool]? {
  Some([])
}

///|
pub fn test_array_output_bool_0() -> Array[Bool] {
  []
}

///|
pub fn test_array_output_bool_1() -> Array[Bool] {
  [true]
}

///|
pub fn test_array_output_bool_2() -> Array[Bool] {
  [false, true]
}

///|
pub fn test_array_output_bool_3() -> Array[Bool] {
  [true, true, true]
}

///|
pub fn test_array_output_bool_4() -> Array[Bool] {
  [false, false, false, false]
}

///|
pub fn test_array_output_bool_option_0() -> Array[Bool?] {
  []
}

///|
pub fn test_array_output_bool_option_1_none() -> Array[Bool?] {
  [None]
}

///|
pub fn test_array_output_bool_option_1_false() -> Array[Bool?] {
  [Some(false)]
}

///|
pub fn test_array_output_bool_option_1_true() -> Array[Bool?] {
  [Some(true)]
}

///|
pub fn test_array_output_bool_option_2() -> Array[Bool?] {
  [Some(true), None]
}

///|
pub fn test_array_output_bool_option_3() -> Array[Bool?] {
  [None, Some(true), Some(true)]
}

///|
pub fn test_array_output_bool_option_4() -> Array[Bool?] {
  [Some(false), Some(true), Some(false), None]
}

///|
pub fn test_array_input_bool_0(val : Array[Bool]) -> Unit raise Error {
  let expected = test_array_output_bool_0()
  // println("expected: \{cast_to_ptr(expected)}")
  // println("expected: \{expected}")
  // println("val: \{cast_to_ptr(val)}")
  // println("val: \{val}")
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_1(val : Array[Bool]) -> Unit raise Error {
  let expected = test_array_output_bool_1()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_2(val : Array[Bool]) -> Unit raise Error {
  let expected = test_array_output_bool_2()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_3(val : Array[Bool]) -> Unit raise Error {
  let expected = test_array_output_bool_3()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_4(val : Array[Bool]) -> Unit raise Error {
  let expected = test_array_output_bool_4()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_0(val : Array[Bool?]) -> Unit raise Error {
  let expected = test_array_output_bool_option_0()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_1_none(
  val : Array[Bool?]
) -> Unit raise Error {
  let expected = test_array_output_bool_option_1_none()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_1_false(
  val : Array[Bool?]
) -> Unit raise Error {
  let expected = test_array_output_bool_option_1_false()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_1_true(
  val : Array[Bool?]
) -> Unit raise Error {
  let expected = test_array_output_bool_option_1_true()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_2(val : Array[Bool?]) -> Unit raise Error {
  let expected = test_array_output_bool_option_2()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_3(val : Array[Bool?]) -> Unit raise Error {
  let expected = test_array_output_bool_option_3()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_4(val : Array[Bool?]) -> Unit raise Error {
  let expected = test_array_output_bool_option_4()
  assert_eq(expected, val)
}
