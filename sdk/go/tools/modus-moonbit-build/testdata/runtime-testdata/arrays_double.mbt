// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_double_empty(val : Array[Double]) -> Unit raise Error {
  let expected = []
  assert_eq(expected, val)
}

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
  [None, None, None]
}

///|
pub fn test_array_output_double_option_4() -> Array[Double?] {
  [None, Some(2), Some(0), Some(4)]
}

///|
pub fn test_array_input_double_option(val : Array[Double?]) -> Unit raise Error {
  let expected = get_double_option_array()
  // println("expected: \{expected}")
  // println("val: \{val}")
  assert_eq(expected, val)
}

///|
pub fn test_array_output_double_option() -> Array[Double?] {
  get_double_option_array()
}

///|
fn get_double_option_array() -> Array[Double?] {
  [Some(1), None, Some(3)]
}
