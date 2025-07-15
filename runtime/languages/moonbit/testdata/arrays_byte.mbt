// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_output_byte_0() -> Array[Byte] {
  []
}

///|
pub fn test_array_output_byte_1() -> Array[Byte] {
  [1]
}

///|
pub fn test_array_output_byte_2() -> Array[Byte] {
  [1, 2]
}

///|
pub fn test_array_output_byte_3() -> Array[Byte] {
  [1, 2, 3]
}

///|
pub fn test_array_output_byte_4() -> Array[Byte] {
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_byte_option_0() -> Array[Byte?] {
  []
}

///|
pub fn test_array_output_byte_option_1() -> Array[Byte?] {
  [Some(1)]
}

///|
pub fn test_array_output_byte_option_2() -> Array[Byte?] {
  [Some(1), None]
}

///|
pub fn test_array_output_byte_option_3() -> Array[Byte?] {
  [None, Some(2), Some(3)]
}

///|
pub fn test_array_output_byte_option_4() -> Array[Byte?] {
  [Some(1), Some(2), Some(3), None]
}

///|
pub fn test_array_input_byte_0(val : Array[Byte]) -> Unit raise Error {
  let expected = test_array_output_byte_0()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_1(val : Array[Byte]) -> Unit raise Error {
  let expected = test_array_output_byte_1()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_2(val : Array[Byte]) -> Unit raise Error {
  let expected = test_array_output_byte_2()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_3(val : Array[Byte]) -> Unit raise Error {
  let expected = test_array_output_byte_3()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_4(val : Array[Byte]) -> Unit raise Error {
  let expected = test_array_output_byte_4()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_option_0(val : Array[Byte?]) -> Unit raise Error {
  let expected = test_array_output_byte_option_0()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_option_1(val : Array[Byte?]) -> Unit raise Error {
  let expected = test_array_output_byte_option_1()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_option_2(val : Array[Byte?]) -> Unit raise Error {
  let expected = test_array_output_byte_option_2()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_option_3(val : Array[Byte?]) -> Unit raise Error {
  let expected = test_array_output_byte_option_3()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_option_4(val : Array[Byte?]) -> Unit raise Error {
  let expected = test_array_output_byte_option_4()
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}
