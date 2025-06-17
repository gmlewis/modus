// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_uint16_empty(val : Array[UInt16]) -> Unit raise Error {
  let expected = []
  assert_eq(expected, val)
}

///|
pub fn test_array_output_uint16_0() -> Array[UInt16] {
  []
}

///|
pub fn test_array_output_uint16_1() -> Array[UInt16] {
  [1]
}

///|
pub fn test_array_output_uint16_1_min() -> Array[UInt16] {
  [@uint16.min_value]
}

///|
pub fn test_array_output_uint16_1_max() -> Array[UInt16] {
  [@uint16.max_value]
}

///|
pub fn test_array_output_uint16_2() -> Array[UInt16] {
  [@uint16.min_value, @uint16.max_value]
}

///|
pub fn test_array_output_uint16_3() -> Array[UInt16] {
  [1, 2, 3]
}

///|
pub fn test_array_output_uint16_4() -> Array[UInt16] {
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_uint16_option_0() -> Array[UInt16?] {
  []
}

///|
pub fn test_array_output_uint16_option_1_none() -> Array[UInt16?] {
  [None]
}

///|
pub fn test_array_output_uint16_option_1_min() -> Array[UInt16?] {
  [Some(@uint16.min_value)]
}

///|
pub fn test_array_output_uint16_option_1_max() -> Array[UInt16?] {
  [Some(@uint16.max_value)]
}

///|
pub fn test_array_output_uint16_option_2() -> Array[UInt16?] {
  [Some(1), Some(2)]
}

///|
pub fn test_array_output_uint16_option_3() -> Array[UInt16?] {
  [None, None, None]
}

///|
pub fn test_array_output_uint16_option_4() -> Array[UInt16?] {
  [None, Some(@uint16.min_value), Some(0), Some(@uint16.max_value)]
}

///|
pub fn test_array_input_uint16_option(val : Array[UInt16?]) -> Unit raise Error {
  let expected = get_uint16_option_array()
  assert_eq(expected, val)
}

///|
pub fn test_array_output_uint16_option() -> Array[UInt16?] {
  get_uint16_option_array()
}

///|
fn get_uint16_option_array() -> Array[UInt16?] {
  [Some(11), None, Some(33)]
}
