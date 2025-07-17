// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_output_byte_0() -> Array[Byte] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 0 0 0 3]
  []
}

///|
pub fn test_array_output_byte_1() -> Array[Byte] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 9=8+words*4), moonBitType=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 1]
  [1]
}

///|
pub fn test_array_output_byte_2() -> Array[Byte] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 10=8+words*4), moonBitType=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 1 2]
  [1, 2]
}

///|
pub fn test_array_output_byte_3() -> Array[Byte] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 11=8+words*4), moonBitType=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 1 2 3]
  [1, 2, 3]
}

///|
pub fn test_array_output_byte_4() -> Array[Byte] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=246(FixedArray[Byte]), words=2, memBlock=[1 0 0 0 246 2 0 0 1 2 3 4]
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_byte_option_0() -> Array[Byte?] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  []
}

///|
pub fn test_array_output_byte_option_1() -> Array[Byte?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
  [Some(1)]
}

///|
pub fn test_array_output_byte_option_2() -> Array[Byte?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 255 255 255 255]
  [Some(1), None]
}

///|
pub fn test_array_output_byte_option_3() -> Array[Byte?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 255 255 255 255 2 0 0 0 3 0 0 0]
  [None, Some(2), Some(3)]
}

///|
pub fn test_array_output_byte_option_4() -> Array[Byte?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 2 0 0 0 3 0 0 0 255 255 255 255]
  [Some(1), Some(2), Some(3), None]
}

///|
pub fn test_array_input_byte_0(val : Array[Byte]) -> Unit raise Error {
  let expected = test_array_output_byte_0()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_1(val : Array[Byte]) -> Unit raise Error {
  let expected = test_array_output_byte_1()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_2(val : Array[Byte]) -> Unit raise Error {
  let expected = test_array_output_byte_2()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_3(val : Array[Byte]) -> Unit raise Error {
  let expected = test_array_output_byte_3()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_4(val : Array[Byte]) -> Unit raise Error {
  let expected = test_array_output_byte_4()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_option_0(val : Array[Byte?]) -> Unit raise Error {
  let expected = test_array_output_byte_option_0()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_option_1(val : Array[Byte?]) -> Unit raise Error {
  let expected = test_array_output_byte_option_1()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_option_2(val : Array[Byte?]) -> Unit raise Error {
  let expected = test_array_output_byte_option_2()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_option_3(val : Array[Byte?]) -> Unit raise Error {
  let expected = test_array_output_byte_option_3()
  assert_eq(expected, val)
}

///|
pub fn test_array_input_byte_option_4(val : Array[Byte?]) -> Unit raise Error {
  let expected = test_array_output_byte_option_4()
  assert_eq(expected, val)
}
