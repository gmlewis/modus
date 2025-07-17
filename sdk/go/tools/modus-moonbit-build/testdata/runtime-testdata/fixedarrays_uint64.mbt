// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_output_uint64_0() -> FixedArray[UInt64] {
  []
}

///|
pub fn test_fixedarray_output_uint64_1() -> FixedArray[UInt64] {
  [1]
}

///|
pub fn test_fixedarray_output_uint64_1_min() -> FixedArray[UInt64] {
  [@uint64.min_value]
}

///|
pub fn test_fixedarray_output_uint64_1_max() -> FixedArray[UInt64] {
  [@uint64.max_value]
}

///|
pub fn test_fixedarray_output_uint64_2() -> FixedArray[UInt64] {
  [@uint64.min_value, @uint64.max_value]
}

///|
pub fn test_fixedarray_output_uint64_3() -> FixedArray[UInt64] {
  [1, 2, 3]
}

///|
pub fn test_fixedarray_output_uint64_4() -> FixedArray[UInt64] {
  [1, 2, 3, 4]
}

///|
pub fn test_fixedarray_output_uint64_option_0() -> FixedArray[UInt64?] {
  []
}

///|
pub fn test_fixedarray_output_uint64_option_1_none() -> FixedArray[UInt64?] {
  [None]
}

///|
pub fn test_fixedarray_output_uint64_option_1_min() -> FixedArray[UInt64?] {
  [Some(@uint64.min_value)]
}

///|
pub fn test_fixedarray_output_uint64_option_1_max() -> FixedArray[UInt64?] {
  [Some(@uint64.max_value)]
}

///|
pub fn test_fixedarray_output_uint64_option_2() -> FixedArray[UInt64?] {
  [Some(1), Some(2)]
}

///|
pub fn test_fixedarray_output_uint64_option_3() -> FixedArray[UInt64?] {
  [Some(11), None, Some(33)]
}

///|
pub fn test_fixedarray_output_uint64_option_4() -> FixedArray[UInt64?] {
  [None, Some(@uint64.min_value), Some(0), Some(@uint64.max_value)]
}

///|
pub fn test_fixedarray_input_uint64_0(
  val : FixedArray[UInt64]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_1(
  val : FixedArray[UInt64]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_1()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_1_min(
  val : FixedArray[UInt64]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_1_min()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_1_max(
  val : FixedArray[UInt64]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_1_max()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_2(
  val : FixedArray[UInt64]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_3(
  val : FixedArray[UInt64]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_4(
  val : FixedArray[UInt64]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_4()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_option_0(
  val : FixedArray[UInt64?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_option_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_option_1_none(
  val : FixedArray[UInt64?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_option_1_none()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_option_1_min(
  val : FixedArray[UInt64?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_option_1_min()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_option_1_max(
  val : FixedArray[UInt64?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_option_1_max()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_option_2(
  val : FixedArray[UInt64?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_option_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_option_3(
  val : FixedArray[UInt64?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_option_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_uint64_option_4(
  val : FixedArray[UInt64?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_uint64_option_4()
  assert_eq(val, expected)
}
