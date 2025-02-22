// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_input_int16_empty(val : FixedArray[Int16]) -> Unit!Error {
  let expected : FixedArray[Int16] = []
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_int16_0() -> FixedArray[Int16] {
  []
}

///|
pub fn test_fixedarray_output_int16_1() -> FixedArray[Int16] {
  [1]
}

///|
pub fn test_fixedarray_output_int16_1_min() -> FixedArray[Int16] {
  [@int16.min_value]
}

///|
pub fn test_fixedarray_output_int16_1_max() -> FixedArray[Int16] {
  [@int16.max_value]
}

///|
pub fn test_fixedarray_output_int16_2() -> FixedArray[Int16] {
  [@int16.min_value, @int16.max_value]
}

///|
pub fn test_fixedarray_output_int16_3() -> FixedArray[Int16] {
  [1, 2, 3]
}

///|
pub fn test_fixedarray_output_int16_4() -> FixedArray[Int16] {
  [1, 2, 3, 4]
}

///|
pub fn test_fixedarray_output_int16_option_0() -> FixedArray[Int16?] {
  []
}

///|
pub fn test_fixedarray_output_int16_option_1_none() -> FixedArray[Int16?] {
  [None]
}

///|
pub fn test_fixedarray_output_int16_option_1_min() -> FixedArray[Int16?] {
  [Some(@int16.min_value)]
}

///|
pub fn test_fixedarray_output_int16_option_1_max() -> FixedArray[Int16?] {
  [Some(@int16.max_value)]
}

///|
pub fn test_fixedarray_output_int16_option_2() -> FixedArray[Int16?] {
  [Some(1), Some(2)]
}

///|
pub fn test_fixedarray_output_int16_option_3() -> FixedArray[Int16?] {
  [None, None, None]
}

///|
pub fn test_fixedarray_output_int16_option_4() -> FixedArray[Int16?] {
  [None, Some(@int16.min_value), Some(0), Some(@int16.max_value)]
}

///|
pub fn test_fixedarray_input_int16_option(
  val : FixedArray[Int16?]
) -> Unit!Error {
  let expected = get_int16_option_fixedarray()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_int16_option() -> FixedArray[Int16?] {
  get_int16_option_fixedarray()
}

///|
fn get_int16_option_fixedarray() -> FixedArray[Int16?] {
  [Some(11), None, Some(33)]
}
