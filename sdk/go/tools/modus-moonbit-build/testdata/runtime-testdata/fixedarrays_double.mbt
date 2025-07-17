// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

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
  [Some(11), None, Some(33)]
}

///|
pub fn test_fixedarray_output_double_option_4() -> FixedArray[Double?] {
  [None, Some(2), Some(0), Some(4)]
}

///|
pub fn test_fixedarray_input_double_0(
  val : FixedArray[Double]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_double_1(
  val : FixedArray[Double]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_1()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_double_2(
  val : FixedArray[Double]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_double_3(
  val : FixedArray[Double]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_double_4(
  val : FixedArray[Double]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_4()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_double_option_0(
  val : FixedArray[Double?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_option_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_double_option_1_none(
  val : FixedArray[Double?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_option_1_none()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_double_option_1_some(
  val : FixedArray[Double?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_option_1_some()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_double_option_2(
  val : FixedArray[Double?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_option_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_double_option_3(
  val : FixedArray[Double?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_option_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_double_option_4(
  val : FixedArray[Double?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_double_option_4()
  assert_eq(val, expected)
}
