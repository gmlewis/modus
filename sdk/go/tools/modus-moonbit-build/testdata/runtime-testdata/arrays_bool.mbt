// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_output_bool_0() -> Array[Bool] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  []
}

///|
pub fn test_array_output_bool_1() -> Array[Bool] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
  [true]
}

///|
pub fn test_array_output_bool_2() -> Array[Bool] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 0 0 0 0 1 0 0 0]
  [false, true]
}

///|
pub fn test_array_output_bool_3() -> Array[Bool] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 1 0 0 0 1 0 0 0 1 0 0 0]
  [true, true, true]
}

///|
pub fn test_array_output_bool_4() -> Array[Bool] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
  [false, false, false, false]
}

///|
pub fn test_array_output_bool_option_0() -> Array[Bool?] {
  []
}

///|
pub fn test_array_output_bool_option_1_none() -> Array[Bool?] {
  [None]
}

///|
pub fn test_array_output_bool_option_1_false() -> Array[Bool?] {
  [Some(false)]
}

///|
pub fn test_array_output_bool_option_1_true() -> Array[Bool?] {
  [Some(true)]
}

///|
pub fn test_array_output_bool_option_2() -> Array[Bool?] {
  [Some(true), None]
}

///|
pub fn test_array_output_bool_option_3() -> Array[Bool?] {
  [None, Some(true), Some(true)]
}

///|
pub fn test_array_output_bool_option_4() -> Array[Bool?] {
  [Some(false), Some(true), Some(false), None]
}

///|
pub fn test_array_input_bool_0(val : Array[Bool]) -> Unit raise Error {
  let expected = test_array_output_bool_0()
  println("expected: \{cast_to_ptr(expected)}")
  println("expected: \{dump_mem_block(cast_to_ptr(expected))}")
  println("expected: \{expected}")
  println("val: \{cast_to_ptr(val)}")
  println("val: \{dump_mem_block(cast_to_ptr(val))}")
  println("val: \{val}")
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_1(val : Array[Bool]) -> Unit raise Error {
  let expected = test_array_output_bool_1()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_2(val : Array[Bool]) -> Unit raise Error {
  let expected = test_array_output_bool_2()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_3(val : Array[Bool]) -> Unit raise Error {
  let expected = test_array_output_bool_3()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_4(val : Array[Bool]) -> Unit raise Error {
  let expected = test_array_output_bool_4()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_0(val : Array[Bool?]) -> Unit raise Error {
  let expected = test_array_output_bool_option_0()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_1_none(
  val : Array[Bool?]
) -> Unit raise Error {
  let expected = test_array_output_bool_option_1_none()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_1_false(
  val : Array[Bool?]
) -> Unit raise Error {
  let expected = test_array_output_bool_option_1_false()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_1_true(
  val : Array[Bool?]
) -> Unit raise Error {
  let expected = test_array_output_bool_option_1_true()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_2(val : Array[Bool?]) -> Unit raise Error {
  let expected = test_array_output_bool_option_2()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_3(val : Array[Bool?]) -> Unit raise Error {
  let expected = test_array_output_bool_option_3()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_bool_option_4(val : Array[Bool?]) -> Unit raise Error {
  let expected = test_array_output_bool_option_4()
  assert_eq(expected, val)
}
