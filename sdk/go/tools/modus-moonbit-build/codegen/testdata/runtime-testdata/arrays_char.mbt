// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_char_empty(val : Array[Char]) -> Unit!Error {
  let expected = []
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 0 0 0 0]
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_char_0() -> Array[Char] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  []
}

///|
pub fn test_array_output_char_1() -> Array[Char] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 49 0 0 0]
  ['1']
}

///|
pub fn test_array_output_char_2() -> Array[Char] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 49 0 0 0 50 0 0 0]
  ['1', '2']
}

///|
pub fn test_array_output_char_3() -> Array[Char] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 49 0 0 0 50 0 0 0 51 0 0 0]
  ['1', '2', '3']
}

///|
pub fn test_array_output_char_4() -> Array[Char] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 49 0 0 0 50 0 0 0 51 0 0 0 52 0 0 0]
  ['1', '2', '3', '4']
}

///|
pub fn test_array_output_char_option_0() -> Array[Char?] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  []
}

///|
pub fn test_array_output_char_option_1_none() -> Array[Char?] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 255]
  [None]
}

///|
pub fn test_array_output_char_option_1_some() -> Array[Char?] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 49 0 0 0]
  [Some('1')]
}

///|
pub fn test_array_output_char_option_2() -> Array[Char?] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 49 0 0 0 50 0 0 0]
  [Some('1'), Some('2')]
}

///|
pub fn test_array_output_char_option_3() -> Array[Char?] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 255 255 255 255 255 255 255 255 255 255 255 255]
  [None, None, None]
}

///|
pub fn test_array_output_char_option_4() -> Array[Char?] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 255 255 255 255 50 0 0 0 0 0 0 0 52 0 0 0]
  [None, Some('2'), Some(Char::from_int(0)), Some('4')]
}

///|
pub fn test_array_input_char_option(val : Array[Char?]) -> Unit!Error {
  let expected = get_char_option_array()
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 49 0 0 0 255 255 255 255 51 0 0 0]
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_char_option() -> Array[Char?] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 49 0 0 0 255 255 255 255 51 0 0 0]
  get_char_option_array()
}

///|
fn get_char_option_array() -> Array[Char?] {
  let a = Some('1')
  let b = None
  let c = Some('3')
  [a, b, c]
}
