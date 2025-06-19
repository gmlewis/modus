// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_output_char_0() -> FixedArray[Char] {
  []
}

///|
pub fn test_fixedarray_output_char_1() -> FixedArray[Char] {
  ['1']
}

///|
pub fn test_fixedarray_output_char_2() -> FixedArray[Char] {
  ['1', '2']
}

///|
pub fn test_fixedarray_output_char_3() -> FixedArray[Char] {
  ['1', '2', '3']
}

///|
pub fn test_fixedarray_output_char_4() -> FixedArray[Char] {
  ['1', '2', '3', '4']
}

///|
pub fn test_fixedarray_output_char_option_0() -> FixedArray[Char?] {
  []
}

///|
pub fn test_fixedarray_output_char_option_1_none() -> FixedArray[Char?] {
  [None]
}

///|
pub fn test_fixedarray_output_char_option_1_some() -> FixedArray[Char?] {
  [Some('1')]
}

///|
pub fn test_fixedarray_output_char_option_2() -> FixedArray[Char?] {
  [Some('1'), Some('2')]
}

///|
pub fn test_fixedarray_output_char_option_3() -> FixedArray[Char?] {
  [Some('1'), None, Some('3')]
}

///|
pub fn test_fixedarray_output_char_option_4() -> FixedArray[Char?] {
  [None, Some('2'), Some(Int::unsafe_to_char(0)), Some('4')]
}

///|
pub fn test_fixedarray_input_char_0(val : FixedArray[Char]) -> Unit raise Error {
  let expected = test_fixedarray_output_char_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_char_1(val : FixedArray[Char]) -> Unit raise Error {
  let expected = test_fixedarray_output_char_1()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_char_2(val : FixedArray[Char]) -> Unit raise Error {
  let expected = test_fixedarray_output_char_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_char_3(val : FixedArray[Char]) -> Unit raise Error {
  let expected = test_fixedarray_output_char_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_char_4(val : FixedArray[Char]) -> Unit raise Error {
  let expected = test_fixedarray_output_char_4()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_char_option_0(
  val : FixedArray[Char?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_char_option_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_char_option_1_none(
  val : FixedArray[Char?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_char_option_1_none()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_char_option_1_some(
  val : FixedArray[Char?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_char_option_1_some()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_char_option_2(
  val : FixedArray[Char?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_char_option_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_char_option_3(
  val : FixedArray[Char?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_char_option_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_char_option_4(
  val : FixedArray[Char?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_char_option_4()
  assert_eq(val, expected)
}
