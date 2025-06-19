// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_output_string_0() -> FixedArray[String] {
  []
}

///|
pub fn test_fixedarray_output_string_1() -> FixedArray[String] {
  ["1"]
}

///|
pub fn test_fixedarray_output_string_2() -> FixedArray[String] {
  ["1", "2"]
}

///|
pub fn test_fixedarray_output_string_3() -> FixedArray[String] {
  ["1", "2", "3"]
}

///|
pub fn test_fixedarray_output_string_4() -> FixedArray[String] {
  ["1", "2", "3", "4"]
}

///|
pub fn test_fixedarray_output_string_option_0() -> FixedArray[String?] {
  []
}

///|
pub fn test_fixedarray_output_string_option_1_none() -> FixedArray[String?] {
  [None]
}

///|
pub fn test_fixedarray_output_string_option_1_some() -> FixedArray[String?] {
  [Some("1")]
}

///|
pub fn test_fixedarray_output_string_option_2() -> FixedArray[String?] {
  [Some("1"), Some("2")]
}

///|
pub fn test_fixedarray_output_string_option_3() -> FixedArray[String?] {
  [Some("11"), None, Some("33")]
}

///|
pub fn test_fixedarray_output_string_option_4() -> FixedArray[String?] {
  [None, Some("2"), Some("0"), Some("4")]
}

///|
pub fn test_fixedarray_input_string_0(
  val : FixedArray[String]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_string_1(
  val : FixedArray[String]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_1()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_string_2(
  val : FixedArray[String]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_string_3(
  val : FixedArray[String]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_string_4(
  val : FixedArray[String]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_4()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_string_option_0(
  val : FixedArray[String?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_option_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_string_option_1_none(
  val : FixedArray[String?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_option_1_none()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_string_option_1_some(
  val : FixedArray[String?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_option_1_some()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_string_option_2(
  val : FixedArray[String?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_option_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_string_option_3(
  val : FixedArray[String?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_option_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_string_option_4(
  val : FixedArray[String?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_string_option_4()
  assert_eq(val, expected)
}
