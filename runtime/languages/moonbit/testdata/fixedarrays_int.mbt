// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_input_int_empty(val : FixedArray[Int]) -> Unit!Error {
  let expected : FixedArray[Int] = []
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_int_0() -> FixedArray[Int] {
  []
}

///|
pub fn test_fixedarray_output_int_1() -> FixedArray[Int] {
  [1]
}

///|
pub fn test_fixedarray_output_int_1_min() -> FixedArray[Int] {
  [@int.min_value]
}

///|
pub fn test_fixedarray_output_int_1_max() -> FixedArray[Int] {
  [@int.max_value]
}

///|
pub fn test_fixedarray_output_int_2() -> FixedArray[Int] {
  [@int.min_value, @int.max_value]
}

///|
pub fn test_fixedarray_output_int_3() -> FixedArray[Int] {
  [1, 2, 3]
}

///|
pub fn test_fixedarray_output_int_4() -> FixedArray[Int] {
  [1, 2, 3, 4]
}

///|
pub fn test_fixedarray_output_int_option_0() -> FixedArray[Int?] {
  []
}

///|
pub fn test_fixedarray_output_int_option_1_none() -> FixedArray[Int?] {
  [None]
}

///|
pub fn test_fixedarray_output_int_option_1_min() -> FixedArray[Int?] {
  [Some(@int.min_value)]
}

///|
pub fn test_fixedarray_output_int_option_1_max() -> FixedArray[Int?] {
  [Some(@int.max_value)]
}

///|
pub fn test_fixedarray_output_int_option_2() -> FixedArray[Int?] {
  [Some(1), Some(2)]
}

///|
pub fn test_fixedarray_output_int_option_3() -> FixedArray[Int?] {
  [None, None, None]
}

///|
pub fn test_fixedarray_output_int_option_4() -> FixedArray[Int?] {
  [None, Some(@int.min_value), Some(0), Some(@int.max_value)]
}

///|
pub fn test_fixedarray_input_int_option(val : FixedArray[Int?]) -> Unit!Error {
  let expected = get_int_option_fixedarray()
  // println("expected: \{expected}")
  // println("val: \{val}")
  // println: val: [Some(11), None, Some(33)]
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_int_option() -> FixedArray[Int?] {
  get_int_option_fixedarray()
}

///|
fn get_int_option_fixedarray() -> FixedArray[Int?] {
  let a = Some(11)
  let b = None
  let c = Some(33)
  [a, b, c]
}
