// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_output_bool_0() -> FixedArray[Bool] {
  // memoryBlockAtOffset(offset: 57888=0x0000E220=[32 226 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  []
}

///|
pub fn test_fixedarray_output_bool_1() -> FixedArray[Bool] {
  // memoryBlockAtOffset(offset: 57888=0x0000E220=[32 226 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
  [true]
}

///|
pub fn test_fixedarray_output_bool_2() -> FixedArray[Bool] {
  [false, true]
}

///|
pub fn test_fixedarray_output_bool_3() -> FixedArray[Bool] {
  [true, true, true]
}

///|
pub fn test_fixedarray_output_bool_4() -> FixedArray[Bool] {
  [false, false, false, false]
}

///|
pub fn test_fixedarray_output_bool_option_0() -> FixedArray[Bool?] {
  []
}

///|
pub fn test_fixedarray_output_bool_option_1_none() -> FixedArray[Bool?] {
  [None]
}

///|
pub fn test_fixedarray_output_bool_option_1_false() -> FixedArray[Bool?] {
  [Some(false)]
}

///|
pub fn test_fixedarray_output_bool_option_1_true() -> FixedArray[Bool?] {
  [Some(true)]
}

///|
pub fn test_fixedarray_output_bool_option_2() -> FixedArray[Bool?] {
  [Some(true), None]
}

///|
pub fn test_fixedarray_output_bool_option_3() -> FixedArray[Bool?] {
  [None, Some(true), Some(true)]
}

///|
pub fn test_fixedarray_output_bool_option_4() -> FixedArray[Bool?] {
  [Some(false), Some(true), Some(false), None]
}

///|
pub fn test_fixedarray_input_bool_0(val : FixedArray[Bool]) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_1(val : FixedArray[Bool]) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_1()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_2(val : FixedArray[Bool]) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_3(val : FixedArray[Bool]) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_4(val : FixedArray[Bool]) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_4()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_option_0(
  val : FixedArray[Bool?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_option_0()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_option_1_none(
  val : FixedArray[Bool?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_option_1_none()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_option_1_false(
  val : FixedArray[Bool?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_option_1_false()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_option_1_true(
  val : FixedArray[Bool?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_option_1_true()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_option_2(
  val : FixedArray[Bool?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_option_2()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_option_3(
  val : FixedArray[Bool?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_option_3()
  assert_eq(val, expected)
}

///|
pub fn test_fixedarray_input_bool_option_4(
  val : FixedArray[Bool?]
) -> Unit raise Error {
  let expected = test_fixedarray_output_bool_option_4()
  assert_eq(val, expected)
}
