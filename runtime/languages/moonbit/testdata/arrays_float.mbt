// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_output_float_0() -> Array[Float] {
  []
}

///|
pub fn test_array_output_float_1() -> Array[Float] {
  [1]
}

///|
pub fn test_array_output_float_2() -> Array[Float] {
  [1, 2]
}

///|
pub fn test_array_output_float_3() -> Array[Float] {
  [1, 2, 3]
}

///|
pub fn test_array_output_float_4() -> Array[Float] {
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_float_option_0() -> Array[Float?] {
  []
}

///|
pub fn test_array_output_float_option_1_none() -> Array[Float?] {
  [None]
}

///|
pub fn test_array_output_float_option_1_some() -> Array[Float?] {
  [Some(1)]
}

///|
pub fn test_array_output_float_option_2() -> Array[Float?] {
  [Some(1), Some(2)]
}

///|
pub fn test_array_output_float_option_3() -> Array[Float?] {
  [Some(11), None, Some(33)]
}

///|
pub fn test_array_output_float_option_4() -> Array[Float?] {
  [None, Some(2), Some(0), Some(4)]
}

///|
pub fn test_array_input_float_0(val : Array[Float]) -> Unit raise Error {
  let expected = test_array_output_float_0()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_float_1(val : Array[Float]) -> Unit raise Error {
  let expected = test_array_output_float_1()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_float_2(val : Array[Float]) -> Unit raise Error {
  let expected = test_array_output_float_2()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_float_3(val : Array[Float]) -> Unit raise Error {
  let expected = test_array_output_float_3()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_float_4(val : Array[Float]) -> Unit raise Error {
  let expected = test_array_output_float_4()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_float_option_0(val : Array[Float?]) -> Unit raise Error {
  let expected = test_array_output_float_option_0()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_float_option_1_none(
  val : Array[Float?]
) -> Unit raise Error {
  let expected = test_array_output_float_option_1_none()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_float_option_1_some(
  val : Array[Float?]
) -> Unit raise Error {
  let expected = test_array_output_float_option_1_some()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_float_option_2(val : Array[Float?]) -> Unit raise Error {
  let expected = test_array_output_float_option_2()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_float_option_3(val : Array[Float?]) -> Unit raise Error {
  let expected = test_array_output_float_option_3()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_float_option_4(val : Array[Float?]) -> Unit raise Error {
  let expected = test_array_output_float_option_4()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}
