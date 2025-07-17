// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_output_uint16_0() -> FixedArray[UInt16] {
  []
}

///|
pub fn test_fixedarray_output_uint16_1() -> FixedArray[UInt16] {
  [1]
}

///|
pub fn test_fixedarray_output_uint16_1_min() -> FixedArray[UInt16] {
  [@uint16.min_value]
}

///|
pub fn test_fixedarray_output_uint16_1_max() -> FixedArray[UInt16] {
  [@uint16.max_value]
}

///|
pub fn test_fixedarray_output_uint16_2() -> FixedArray[UInt16] {
  [@uint16.min_value, @uint16.max_value]
}

///|
pub fn test_fixedarray_output_uint16_3() -> FixedArray[UInt16] {
  [1, 2, 3]
}

///|
pub fn test_fixedarray_output_uint16_4() -> FixedArray[UInt16] {
  [1, 2, 3, 4]
}

///|
pub fn test_fixedarray_output_uint16_option_0() -> FixedArray[UInt16?] {
  []
}

///|
pub fn test_fixedarray_output_uint16_option_1_none() -> FixedArray[UInt16?] {
  [None]
}

///|
pub fn test_fixedarray_output_uint16_option_1_min() -> FixedArray[UInt16?] {
  [Some(@uint16.min_value)]
}

///|
pub fn test_fixedarray_output_uint16_option_1_max() -> FixedArray[UInt16?] {
  [Some(@uint16.max_value)]
}

///|
pub fn test_fixedarray_output_uint16_option_2() -> FixedArray[UInt16?] {
  [Some(1), Some(2)]
}

///|
pub fn test_fixedarray_output_uint16_option_3() -> FixedArray[UInt16?] {
  [Some(11), None, Some(33)]
}

///|
pub fn test_fixedarray_output_uint16_option_4() -> FixedArray[UInt16?] {
  [None, Some(@uint16.min_value), Some(0), Some(@uint16.max_value)]
}

///|
pub fn test_fixedarray_input_uint16_0(
  val : FixedArray[UInt16]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_1(
  val : FixedArray[UInt16]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_1()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_1_min(
  val : FixedArray[UInt16]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_1_min()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_1_max(
  val : FixedArray[UInt16]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_1_max()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_2(
  val : FixedArray[UInt16]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_3(
  val : FixedArray[UInt16]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_4(
  val : FixedArray[UInt16]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_4()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_option_0(
  val : FixedArray[UInt16?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_option_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_option_1_none(
  val : FixedArray[UInt16?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_option_1_none()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_option_1_min(
  val : FixedArray[UInt16?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_option_1_min()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_option_1_max(
  val : FixedArray[UInt16?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_option_1_max()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_option_2(
  val : FixedArray[UInt16?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_option_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_option_3(
  val : FixedArray[UInt16?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_option_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint16_option_4(
  val : FixedArray[UInt16?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint16_option_4()
  assert_eq(val, expected)
}
