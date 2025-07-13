// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_uint_empty(val : Array[UInt]) -> Unit raise Error {
  let expected = []
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_output_uint_0() -> Array[UInt] {
  []
}

///|
pub fn test_array_output_uint_1() -> Array[UInt] {
  [1]
}

///|
pub fn test_array_output_uint_1_min() -> Array[UInt] {
  [@uint.min_value]
}

///|
pub fn test_array_output_uint_1_max() -> Array[UInt] {
  [@uint.max_value]
}

///|
pub fn test_array_output_uint_2() -> Array[UInt] {
  [@uint.min_value, @uint.max_value]
}

///|
pub fn test_array_output_uint_3() -> Array[UInt] {
  [1, 2, 3]
}

///|
pub fn test_array_output_uint_4() -> Array[UInt] {
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_uint_option_0() -> Array[UInt?] {
  []
}

///|
pub fn test_array_output_uint_option_1_none() -> Array[UInt?] {
  [None]
}

///|
pub fn test_array_output_uint_option_1_min() -> Array[UInt?] {
  [Some(@uint.min_value)]
}

///|
pub fn test_array_output_uint_option_1_max() -> Array[UInt?] {
  [Some(@uint.max_value)]
}

///|
pub fn test_array_output_uint_option_2() -> Array[UInt?] {
  [Some(1), Some(2)]
}

///|
pub fn test_array_output_uint_option_3() -> Array[UInt?] {
  [None, None, None]
}

///|
pub fn test_array_output_uint_option_4() -> Array[UInt?] {
  [None, Some(@uint.min_value), Some(0), Some(@uint.max_value)]
}

///|
pub fn test_array_input_uint_option(val : Array[UInt?]) -> Unit raise Error {
  let expected = get_uint_option_array()
  // println("expected: \{expected}")
  // println("val: \{val}")
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_output_uint_option() -> Array[UInt?] {
  get_uint_option_array()
}

///|
fn get_uint_option_array() -> Array[UInt?] {
  let a = Some(11U)
  let b = None
  let c = Some(33U)
  [a, b, c]
}
