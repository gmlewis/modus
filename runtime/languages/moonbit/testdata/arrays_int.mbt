// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_int_empty(val : Array[Int]) -> Unit raise Error {
  let expected = []
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_output_int_0() -> Array[Int] {
  []
}

///|
pub fn test_array_output_int_1() -> Array[Int] {
  [1]
}

///|
pub fn test_array_output_int_1_min() -> Array[Int] {
  [@int.min_value]
}

///|
pub fn test_array_output_int_1_max() -> Array[Int] {
  [@int.max_value]
}

///|
pub fn test_array_output_int_2() -> Array[Int] {
  [@int.min_value, @int.max_value]
}

///|
pub fn test_array_output_int_3() -> Array[Int] {
  [1, 2, 3]
}

///|
pub fn test_array_output_int_4() -> Array[Int] {
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_int_option_0() -> Array[Int?] {
  []
}

///|
pub fn test_array_output_int_option_1_none() -> Array[Int?] {
  [None]
}

///|
pub fn test_array_output_int_option_1_min() -> Array[Int?] {
  [Some(@int.min_value)]
}

///|
pub fn test_array_output_int_option_1_max() -> Array[Int?] {
  [Some(@int.max_value)]
}

///|
pub fn test_array_output_int_option_2() -> Array[Int?] {
  [Some(1), Some(2)]
}

///|
pub fn test_array_output_int_option_3() -> Array[Int?] {
  [None, None, None]
}

///|
pub fn test_array_output_int_option_4() -> Array[Int?] {
  [None, Some(@int.min_value), Some(0), Some(@int.max_value)]
}

///|
pub fn test_array_input_int_option(val : Array[Int?]) -> Unit raise Error {
  let expected = get_int_option_array()
  // println("expected: \{expected}")
  // println("val: \{val}")
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}

///|
pub fn test_array_output_int_option() -> Array[Int?] {
  get_int_option_array()
}

///|
fn get_int_option_array() -> Array[Int?] {
  let a = Some(11)
  let b = None
  let c = Some(33)
  [a, b, c]
}
