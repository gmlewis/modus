// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_output_uint_0() -> FixedArray[UInt] {
  []
}

///|
pub fn test_fixedarray_output_uint_1() -> FixedArray[UInt] {
  [1]
}

///|
pub fn test_fixedarray_output_uint_1_min() -> FixedArray[UInt] {
  [@uint.min_value]
}

///|
pub fn test_fixedarray_output_uint_1_max() -> FixedArray[UInt] {
  [@uint.max_value]
}

///|
pub fn test_fixedarray_output_uint_2() -> FixedArray[UInt] {
  [@uint.min_value, @uint.max_value]
}

///|
pub fn test_fixedarray_output_uint_3() -> FixedArray[UInt] {
  [1, 2, 3]
}

///|
pub fn test_fixedarray_output_uint_4() -> FixedArray[UInt] {
  [1, 2, 3, 4]
}

///|
pub fn test_fixedarray_output_uint_option_0() -> FixedArray[UInt?] {
  []
}

///|
pub fn test_fixedarray_output_uint_option_1_none() -> FixedArray[UInt?] {
  [None]
}

///|
pub fn test_fixedarray_output_uint_option_1_min() -> FixedArray[UInt?] {
  [Some(@uint.min_value)]
}

///|
pub fn test_fixedarray_output_uint_option_1_max() -> FixedArray[UInt?] {
  [Some(@uint.max_value)]
}

///|
pub fn test_fixedarray_output_uint_option_2() -> FixedArray[UInt?] {
  [Some(1), Some(2)]
}

///|
pub fn test_fixedarray_output_uint_option_3() -> FixedArray[UInt?] {
  [Some(11), None, Some(33)]
}

///|
pub fn test_fixedarray_output_uint_option_4() -> FixedArray[UInt?] {
  [None, Some(@uint.min_value), Some(0), Some(@uint.max_value)]
}

///|
pub fn test_fixedarray_input_uint_0(val : FixedArray[UInt]) -> Unit!Error {
  let expected = test_fixedarray_output_uint_0()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_1(val : FixedArray[UInt]) -> Unit!Error {
  let expected = test_fixedarray_output_uint_1()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_1_min(val : FixedArray[UInt]) -> Unit!Error {
  let expected = test_fixedarray_output_uint_1_min()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_1_max(val : FixedArray[UInt]) -> Unit!Error {
  let expected = test_fixedarray_output_uint_1_max()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_2(val : FixedArray[UInt]) -> Unit!Error {
  let expected = test_fixedarray_output_uint_2()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_3(val : FixedArray[UInt]) -> Unit!Error {
  let expected = test_fixedarray_output_uint_3()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_4(val : FixedArray[UInt]) -> Unit!Error {
  let expected = test_fixedarray_output_uint_4()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_option_0(
  val : FixedArray[UInt?]
) -> Unit!Error {
  let expected = test_fixedarray_output_uint_option_0()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_option_1_none(
  val : FixedArray[UInt?]
) -> Unit!Error {
  let expected = test_fixedarray_output_uint_option_1_none()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_option_1_min(
  val : FixedArray[UInt?]
) -> Unit!Error {
  let expected = test_fixedarray_output_uint_option_1_min()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_option_1_max(
  val : FixedArray[UInt?]
) -> Unit!Error {
  let expected = test_fixedarray_output_uint_option_1_max()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_option_2(
  val : FixedArray[UInt?]
) -> Unit!Error {
  let expected = test_fixedarray_output_uint_option_2()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_option_3(
  val : FixedArray[UInt?]
) -> Unit!Error {
  let expected = test_fixedarray_output_uint_option_3()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_uint_option_4(
  val : FixedArray[UInt?]
) -> Unit!Error {
  let expected = test_fixedarray_output_uint_option_4()
  assert_eq!(val, expected)
}
