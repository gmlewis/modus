// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_output_int64_0() -> FixedArray[Int64] {
  []
}

///|
pub fn test_fixedarray_output_int64_1() -> FixedArray[Int64] {
  [1]
}

///|
pub fn test_fixedarray_output_int64_1_min() -> FixedArray[Int64] {
  [@int64.min_value]
}

///|
pub fn test_fixedarray_output_int64_1_max() -> FixedArray[Int64] {
  [@int64.max_value]
}

///|
pub fn test_fixedarray_output_int64_2() -> FixedArray[Int64] {
  [@int64.min_value, @int64.max_value]
}

///|
pub fn test_fixedarray_output_int64_3() -> FixedArray[Int64] {
  [1, 2, 3]
}

///|
pub fn test_fixedarray_output_int64_4() -> FixedArray[Int64] {
  [1, 2, 3, 4]
}

///|
pub fn test_fixedarray_output_int64_option_0() -> FixedArray[Int64?] {
  []
}

///|
pub fn test_fixedarray_output_int64_option_1_none() -> FixedArray[Int64?] {
  [None]
}

///|
pub fn test_fixedarray_output_int64_option_1_min() -> FixedArray[Int64?] {
  [Some(@int64.min_value)]
}

///|
pub fn test_fixedarray_output_int64_option_1_max() -> FixedArray[Int64?] {
  [Some(@int64.max_value)]
}

///|
pub fn test_fixedarray_output_int64_option_2() -> FixedArray[Int64?] {
  [Some(1), Some(2)]
}

///|
pub fn test_fixedarray_output_int64_option_3() -> FixedArray[Int64?] {
  [Some(11), None, Some(33)]
}

///|
pub fn test_fixedarray_output_int64_option_4() -> FixedArray[Int64?] {
  [None, Some(@int64.min_value), Some(0), Some(@int64.max_value)]
}

///|
pub fn test_fixedarray_input_int64_0(val : FixedArray[Int64]) -> Unit!Error {
  let expected = test_fixedarray_output_int64_0()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_1(val : FixedArray[Int64]) -> Unit!Error {
  let expected = test_fixedarray_output_int64_1()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_1_min(val : FixedArray[Int64]) -> Unit!Error {
  let expected = test_fixedarray_output_int64_1_min()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_1_max(val : FixedArray[Int64]) -> Unit!Error {
  let expected = test_fixedarray_output_int64_1_max()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_2(val : FixedArray[Int64]) -> Unit!Error {
  let expected = test_fixedarray_output_int64_2()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_3(val : FixedArray[Int64]) -> Unit!Error {
  let expected = test_fixedarray_output_int64_3()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_4(val : FixedArray[Int64]) -> Unit!Error {
  let expected = test_fixedarray_output_int64_4()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_option_0(
  val : FixedArray[Int64?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int64_option_0()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_option_1_none(
  val : FixedArray[Int64?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int64_option_1_none()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_option_1_min(
  val : FixedArray[Int64?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int64_option_1_min()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_option_1_max(
  val : FixedArray[Int64?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int64_option_1_max()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_option_2(
  val : FixedArray[Int64?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int64_option_2()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_option_3(
  val : FixedArray[Int64?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int64_option_3()
  assert_eq!(val, expected)
}

///|
pub fn test_fixedarray_input_int64_option_4(
  val : FixedArray[Int64?]
) -> Unit!Error {
  let expected = test_fixedarray_output_int64_option_4()
  assert_eq!(val, expected)
}
