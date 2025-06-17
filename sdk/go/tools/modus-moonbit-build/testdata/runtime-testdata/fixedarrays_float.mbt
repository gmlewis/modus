// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_output_float_0() -> FixedArray[Float] {
  []
}

///|
pub fn test_fixedarray_output_float_1() -> FixedArray[Float] {
  [1]
}

///|
pub fn test_fixedarray_output_float_2() -> FixedArray[Float] {
  [1, 2]
}

///|
pub fn test_fixedarray_output_float_3() -> FixedArray[Float] {
  [1, 2, 3]
}

///|
pub fn test_fixedarray_output_float_4() -> FixedArray[Float] {
  [1, 2, 3, 4]
}

///|
pub fn test_fixedarray_output_float_option_0() -> FixedArray[Float?] {
  []
}

///|
pub fn test_fixedarray_output_float_option_1_none() -> FixedArray[Float?] {
  [None]
}

///|
pub fn test_fixedarray_output_float_option_1_some() -> FixedArray[Float?] {
  [Some(1)]
}

///|
pub fn test_fixedarray_output_float_option_2() -> FixedArray[Float?] {
  [Some(1), Some(2)]
}

///|
pub fn test_fixedarray_output_float_option_3() -> FixedArray[Float?] {
  [Some(11), None, Some(33)]
}

///|
pub fn test_fixedarray_output_float_option_4() -> FixedArray[Float?] {
  [None, Some(2), Some(0), Some(4)]
}

///|
pub fn test_fixedarray_input_float_0(
  val : FixedArray[Float]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_float_1(
  val : FixedArray[Float]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_1()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_float_2(
  val : FixedArray[Float]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_float_3(
  val : FixedArray[Float]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_float_4(
  val : FixedArray[Float]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_4()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_float_option_0(
  val : FixedArray[Float?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_option_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_float_option_1_none(
  val : FixedArray[Float?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_option_1_none()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_float_option_1_some(
  val : FixedArray[Float?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_option_1_some()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_float_option_2(
  val : FixedArray[Float?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_option_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_float_option_3(
  val : FixedArray[Float?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_option_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_float_option_4(
  val : FixedArray[Float?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_float_option_4()
  assert_eq(val, expected)
}
