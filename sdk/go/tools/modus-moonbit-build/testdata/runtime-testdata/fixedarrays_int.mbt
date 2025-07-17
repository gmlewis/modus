// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

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
  [Some(11), None, Some(33)]
}

///|
pub fn test_fixedarray_output_int_option_4() -> FixedArray[Int?] {
  [None, Some(@int.min_value), Some(0), Some(@int.max_value)]
}

///|
pub fn test_fixedarray_input_int_0(val : FixedArray[Int]) -> Unit raise Error {
  let expected = test_fixedarray_output_int_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_1(val : FixedArray[Int]) -> Unit raise Error {
  let expected = test_fixedarray_output_int_1()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_1_min(
  val : FixedArray[Int]
) -> Unit raise Error {
  let expected = test_fixedarray_output_int_1_min()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_1_max(
  val : FixedArray[Int]
) -> Unit raise Error {
  let expected = test_fixedarray_output_int_1_max()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_2(val : FixedArray[Int]) -> Unit raise Error {
  let expected = test_fixedarray_output_int_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_3(val : FixedArray[Int]) -> Unit raise Error {
  let expected = test_fixedarray_output_int_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_4(val : FixedArray[Int]) -> Unit raise Error {
  let expected = test_fixedarray_output_int_4()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_option_0(
  val : FixedArray[Int?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_int_option_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_option_1_none(
  val : FixedArray[Int?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_int_option_1_none()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_option_1_min(
  val : FixedArray[Int?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_int_option_1_min()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_option_1_max(
  val : FixedArray[Int?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_int_option_1_max()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_option_2(
  val : FixedArray[Int?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_int_option_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_option_3(
  val : FixedArray[Int?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_int_option_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_int_option_4(
  val : FixedArray[Int?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_int_option_4()
  assert_eq(val, expected)
}
