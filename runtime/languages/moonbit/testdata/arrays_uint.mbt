// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_uint_empty(val : Array[UInt]) -> Unit!Error {
  let expected = []
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_uint_0() -> Array[UInt] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[UInt]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  []
}

///|
pub fn test_array_output_uint_1() -> Array[UInt] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
  [1]
}

///|
pub fn test_array_output_uint_1_min() -> Array[UInt] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 128]
  [@uint.min_value]
}

///|
pub fn test_array_output_uint_1_max() -> Array[UInt] {
  [@uint.max_value]
}

///|
pub fn test_array_output_uint_2() -> Array[UInt] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[UInt]), words=2, memBlock=[1 0 0 0 241 2 0 0 0 0 0 128 255 255 255 127]
  [@uint.min_value, @uint.max_value]
}

///|
pub fn test_array_output_uint_3() -> Array[UInt] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[UInt]), words=3, memBlock=[1 0 0 0 241 3 0 0 1 0 0 0 2 0 0 0 3 0 0 0]
  [1, 2, 3]
}

///|
pub fn test_array_output_uint_4() -> Array[UInt] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[UInt]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 2 0 0 0 3 0 0 0 4 0 0 0]
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_uint_option_0() -> Array[UInt?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[UInt]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  []
}

///|
pub fn test_array_output_uint_option_1_none() -> Array[UInt?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0 1 0 0 0]
  [None]
}

///|
pub fn test_array_output_uint_option_1_min() -> Array[UInt?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 128]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 128 255 255 255 255]
  [Some(@uint.min_value)]
}

///|
pub fn test_array_output_uint_option_1_max() -> Array[UInt?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 127]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 127 0 0 0 0]
  [Some(@uint.max_value)]
}

///|
pub fn test_array_output_uint_option_2() -> Array[UInt?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[UInt]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[UInt]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0]
  [Some(1), Some(2)]
}

///|
pub fn test_array_output_uint_option_3() -> Array[UInt?] {
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[UInt]), words=3, memBlock=[1 0 0 0 241 3 0 0 0 0 0 0 1 0 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[UInt]), words=3, memBlock=[1 0 0 0 241 3 0 0 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0]
  [None, None, None]
}

///|
pub fn test_array_output_uint_option_4() -> Array[UInt?] {
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[UInt]), words=4, memBlock=[1 0 0 0 241 4 0 0 0 0 0 0 1 0 0 0 0 0 0 128 255 255 255 255]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 40=8+words*4), moonBitType=241(FixedArray[UInt]), words=4, memBlock=[1 0 0 0 241 4 0 0 0 0 0 0 1 0 0 0 0 0 0 128 255 255 255 255 0 0 0 0 0 0 0 0 255 255 255 127 0 0 0 0]
  [None, Some(@uint.min_value), Some(0), Some(@uint.max_value)]
}

///|
pub fn test_array_input_uint_option(val : Array[UInt?]) -> Unit!Error {
  let expected = get_uint_option_array()
  // pruintln("expected: \{expected}")
  // pruintln("val: \{val}")
  // pruintln: val: [Some(11), None, Some(33)]
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[UInt]), words=6, memBlock=[1 0 0 0 241 6 0 0 11 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 33 0 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 32 191 0 0 3 0 0 0]
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_uint_option() -> Array[UInt?] {
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 32 191 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[UInt]), words=3, memBlock=[1 0 0 0 241 3 0 0 11 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 33 0 0 0 0 0 0 0]
  get_uint_option_array()
}

///|
fn get_uint_option_array() -> Array[UInt?] {
  let a = Some(11U)
  let b = None
  let c = Some(33U)
  [a, b, c]
}
