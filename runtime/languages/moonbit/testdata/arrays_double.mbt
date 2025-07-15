// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_output_double_0() -> Array[Double] {
  []
}

///|
pub fn test_array_output_double_1() -> Array[Double] {
  [1]
}

///|
pub fn test_array_output_double_2() -> Array[Double] {
  [1, 2]
}

///|
pub fn test_array_output_double_3() -> Array[Double] {
  [1, 2, 3]
}

///|
pub fn test_array_output_double_4() -> Array[Double] {
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_double_option_0() -> Array[Double?] {
  []
}

///|
pub fn test_array_output_double_option_1_none() -> Array[Double?] {
  [None]
}

///|
pub fn test_array_output_double_option_1_some() -> Array[Double?] {
  [Some(1)]
}

///|
pub fn test_array_output_double_option_2() -> Array[Double?] {
  [Some(1), Some(2)]
}

///|
pub fn test_array_output_double_option_3() -> Array[Double?] {
  [Some(11), None, Some(33)]
}

///|
pub fn test_array_output_double_option_4() -> Array[Double?] {
  [None, Some(2), Some(0), Some(4)]
}

///|
pub fn test_array_input_double_0(val : Array[Double]) -> Unit raise Error {
  let expected = test_array_output_double_0()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_double_1(val : Array[Double]) -> Unit raise Error {
  let expected = test_array_output_double_1()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_double_2(val : Array[Double]) -> Unit raise Error {
  let expected = test_array_output_double_2()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_double_3(val : Array[Double]) -> Unit raise Error {
  let expected = test_array_output_double_3()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_double_4(val : Array[Double]) -> Unit raise Error {
  let expected = test_array_output_double_4()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_double_option_0(
  val : Array[Double?]
) -> Unit raise Error {
  let expected = test_array_output_double_option_0()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_double_option_1_none(
  val : Array[Double?]
) -> Unit raise Error {
  let expected = test_array_output_double_option_1_none()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_double_option_1_some(
  val : Array[Double?]
) -> Unit raise Error {
  let expected = test_array_output_double_option_1_some()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_double_option_2(
  val : Array[Double?]
) -> Unit raise Error {
  let expected = test_array_output_double_option_2()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_double_option_3(
  val : Array[Double?]
) -> Unit raise Error {
  let expected = test_array_output_double_option_3()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_double_option_4(
  val : Array[Double?]
) -> Unit raise Error {
  let expected = test_array_output_double_option_4()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}
