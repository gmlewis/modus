// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

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
  [Some(11), None, Some(33)]
}

///|
pub fn test_fixedarray_output_int16_option_4() -> FixedArray[Int16?] {
  [None, Some(@int16.min_value), Some(0), Some(@int16.max_value)]
}

///|
pub fn test_fixedarray_input_int16_0(val : FixedArray[Int16]) -> Unit!Error {
  let expected = test_fixedarray_output_int16_0()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_1(val : FixedArray[Int16]) -> Unit!Error {
  let expected = test_fixedarray_output_int16_1()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_1_min(val : FixedArray[Int16]) -> Unit!Error {
  let expected = test_fixedarray_output_int16_1_min()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_1_max(val : FixedArray[Int16]) -> Unit!Error {
  let expected = test_fixedarray_output_int16_1_max()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_2(val : FixedArray[Int16]) -> Unit!Error {
  let expected = test_fixedarray_output_int16_2()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_3(val : FixedArray[Int16]) -> Unit!Error {
  let expected = test_fixedarray_output_int16_3()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_4(val : FixedArray[Int16]) -> Unit!Error {
  let expected = test_fixedarray_output_int16_4()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_option_0(
  val : FixedArray[Int16?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int16_option_0()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_option_1_none(
  val : FixedArray[Int16?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int16_option_1_none()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_option_1_min(
  val : FixedArray[Int16?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int16_option_1_min()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_option_1_max(
  val : FixedArray[Int16?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int16_option_1_max()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_option_2(
  val : FixedArray[Int16?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int16_option_2()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_option_3(
  val : FixedArray[Int16?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int16_option_3()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int16_option_4(
  val : FixedArray[Int16?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int16_option_4()
  assert_eq!(val, expected)
}
