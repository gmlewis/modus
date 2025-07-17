// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_int16_empty(val : Array[Int16]) -> Unit raise Error {
  let expected = []
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 0 0 0 0]
  assert_eq(expected, val)
}

///|
pub fn test_array_output_int16_0() -> Array[Int16] {
  []
}

///|
pub fn test_array_output_int16_1() -> Array[Int16] {
  [1]
}

///|
pub fn test_array_output_int16_1_min() -> Array[Int16] {
  [@int16.min_value]
}

///|
pub fn test_array_output_int16_1_max() -> Array[Int16] {
  [@int16.max_value]
}

///|
pub fn test_array_output_int16_2() -> Array[Int16] {
  [@int16.min_value, @int16.max_value]
}

///|
pub fn test_array_output_int16_3() -> Array[Int16] {
  [1, 2, 3]
}

///|
pub fn test_array_output_int16_4() -> Array[Int16] {
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_int16_option_0() -> Array[Int16?] {
  []
}

///|
pub fn test_array_output_int16_option_1_none() -> Array[Int16?] {
  [None]
}

///|
pub fn test_array_output_int16_option_1_min() -> Array[Int16?] {
  [Some(@int16.min_value)]
}

///|
pub fn test_array_output_int16_option_1_max() -> Array[Int16?] {
  [Some(@int16.max_value)]
}

///|
pub fn test_array_output_int16_option_2() -> Array[Int16?] {
  [Some(1), Some(2)]
}

///|
pub fn test_array_output_int16_option_3() -> Array[Int16?] {
  [None, None, None]
}

///|
pub fn test_array_output_int16_option_4() -> Array[Int16?] {
  [None, Some(@int16.min_value), Some(0), Some(@int16.max_value)]
}

///|
pub fn test_array_input_int16_option(val : Array[Int16?]) -> Unit raise Error {
  let expected = get_int16_option_array()
  assert_eq(expected, val)
}

///|
pub fn test_array_output_int16_option() -> Array[Int16?] {
  // memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 11 0 0 0 255 255 255 255 33 0 0 0]
  get_int16_option_array()
}

///|
fn get_int16_option_array() -> Array[Int16?] {
  [Some(11), None, Some(33)]
}
