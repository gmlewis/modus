// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_input_double_empty(
  val : FixedArray[Double]
) -> Unit!Error {
  let expected : FixedArray[Double] = []
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_double_0() -> FixedArray[Double] {
  []
}

///|
pub fn test_fixedarray_output_double_1() -> FixedArray[Double] {
  [1]
}

///|
pub fn test_fixedarray_output_double_2() -> FixedArray[Double] {
  [1, 2]
}

///|
pub fn test_fixedarray_output_double_3() -> FixedArray[Double] {
  [1, 2, 3]
}

///|
pub fn test_fixedarray_output_double_4() -> FixedArray[Double] {
  [1, 2, 3, 4]
}

///|
pub fn test_fixedarray_output_double_option_0() -> FixedArray[Double?] {
  []
}

///|
pub fn test_fixedarray_output_double_option_1_none() -> FixedArray[Double?] {
  [None]
}

///|
pub fn test_fixedarray_output_double_option_1_some() -> FixedArray[Double?] {
  [Some(1)]
}

///|
pub fn test_fixedarray_output_double_option_2() -> FixedArray[Double?] {
  [Some(1), Some(2)]
}

///|
pub fn test_fixedarray_output_double_option_3() -> FixedArray[Double?] {
  [None, None, None]
}

///|
pub fn test_fixedarray_output_double_option_4() -> FixedArray[Double?] {
  [None, Some(2), Some(0), Some(4)]
}

///|
pub fn test_fixedarray_input_double_option(
  val : FixedArray[Double?]
) -> Unit!Error {
  let expected = get_double_option_fixedarray()
  // println("expected: \{expected}")
  // println("val: \{val}")
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_double_option() -> FixedArray[Double?] {
  get_double_option_fixedarray()
}

///|
fn get_double_option_fixedarray() -> FixedArray[Double?] {
  [Some(1), None, Some(3)]
}
