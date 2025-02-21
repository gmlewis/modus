// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_bool(val : Array[Bool]) -> Unit!Error {
  let expected : Array[Bool] = [false, true, true, false]
  println("val: \{dump_mem_block(cast_to_ptr(val))}")
  assert_eq!(expected, val)
}

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
pub fn test_array_input_byte(val : Array[Byte]) -> Unit!Error {
  let expected : Array[Byte] = [1, 2, 3, 4]
  assert_eq!(expected, val)
}

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
pub fn test_array_input_string(val : Array[String]) -> Unit!Error {
  let expected = ["abc", "def", "ghi"]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 32 191 0 0 64 191 0 0 96 191 0 0]
  // GML: handler_memory.go: memoryBlockAtOffset(offset: 49024=0x0000BF80=[128 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_string() -> Array[String] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 32 59 0 0 208 62 0 0 104 75 0 0]
  // memoryBlockAtOffset(offset: 15136=0x00003B20=[32 59 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
  // memoryBlockAtOffset(offset: 16080=0x00003ED0=[208 62 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 100 0 101 0 102 0 0 1] = 'def'
  // memoryBlockAtOffset(offset: 19304=0x00004B68=[104 75 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 103 0 104 0 105 0 0 1] = 'ghi'
  ["abc", "def", "ghi"]
}

///|
pub fn test_array_input_string_none(val : Array[String]?) -> Unit!Error {
  // memoryBlockAtOffset(offset: 0) = (data=0, size=0)
  assert_eq!(val, None)
}

///|
pub fn test_array_output_string_none() -> Array[String]? {
  // pointerHandler.Decode(vals: [0])
  None
}

///|
pub fn test_array_input_string_empty(val : Array[String]) -> Unit!Error {
  let expected = []
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_string_empty() -> Array[String] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=242(), words=0, memBlock=[1 0 0 0 242 0 0 0]
  []
}

///|
pub fn test_array_input_int_empty(val : Array[Int]) -> Unit!Error {
  let expected = []
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_int_0() -> Array[Int] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  []
}

///|
pub fn test_array_output_int_1() -> Array[Int] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
  [1]
}

///|
pub fn test_array_output_int_1_min() -> Array[Int] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 128]
  [@int.min_value]
}

///|
pub fn test_array_output_int_1_max() -> Array[Int] {
  [@int.max_value]
}

///|
pub fn test_array_output_int_2() -> Array[Int] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 0 0 0 128 255 255 255 127]
  [@int.min_value, @int.max_value]
}

///|
pub fn test_array_output_int_3() -> Array[Int] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 1 0 0 0 2 0 0 0 3 0 0 0]
  [1, 2, 3]
}

///|
pub fn test_array_output_int_4() -> Array[Int] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 2 0 0 0 3 0 0 0 4 0 0 0]
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_int_option_0() -> Array[Int?] {

  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  []
}

///|
pub fn test_array_output_int_option_1_none() -> Array[Int?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0 1 0 0 0]
  [None]
}

///|
pub fn test_array_output_int_option_1_min() -> Array[Int?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 128]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 128 255 255 255 255]
  [Some(@int.min_value)]
}

///|
pub fn test_array_output_int_option_1_max() -> Array[Int?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 127]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 127 0 0 0 0]
  [Some(@int.max_value)]
}

///|
pub fn test_array_output_int_option_2() -> Array[Int?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0]
  [Some(1), Some(2)]
}

///|
pub fn test_array_output_int_option_3() -> Array[Int?] {
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 0 0 0 0 1 0 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0]
  [None, None, None]
}

///|
pub fn test_array_output_int_option_4() -> Array[Int?] {
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 0 0 0 0 1 0 0 0 0 0 0 128 255 255 255 255]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 40=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 0 0 0 0 1 0 0 0 0 0 0 128 255 255 255 255 0 0 0 0 0 0 0 0 255 255 255 127 0 0 0 0]
  // GML: handler_primitives.go: primitiveHandler.Decode(vals: [4294967296])
  // GML: handler_primitives.go: primitiveHandler.Decode(vals: [18446744071562067968])
  // GML: handler_primitives.go: primitiveHandler.Decode(vals: [0])
  // GML: handler_primitives.go: primitiveHandler.Decode(vals: [2147483647])
  [None, Some(@int.min_value), Some(0), Some(@int.max_value)]
}

///|
pub fn test_array_input_string_option(val : Array[String?]) -> Unit!Error {
  let expected = get_string_option_array()
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[1 0 0 0 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
  // memoryBlockAtOffset(offset: 48960=0x0000BF40=[64 191 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[1 0 0 0 243 2 0 0 103 0 104 0 105 0 0 1] = 'ghi'
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 32 191 0 0 0 0 0 0 64 191 0 0]
  // memoryBlockAtOffset(offset: 48992=0x0000BF60=[96 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_string_option() -> Array[String?] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 32 59 0 0 0 0 0 0 104 75 0 0]
  // memoryBlockAtOffset(offset: 15136=0x00003B20=[32 59 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
  // memoryBlockAtOffset(offset: 19304=0x00004B68=[104 75 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 103 0 104 0 105 0 0 1] = 'ghi'
  get_string_option_array()
}

///|
fn get_string_option_array() -> Array[String?] {
  let a = Some("abc")
  let b = None
  let c = Some("ghi")
  [a, b, c]
}

///|
pub fn test_array_input_int_option(val : Array[Int?]) -> Unit!Error {
  let expected = get_int_option_array()
  // println("expected: \{expected}")
  // println("val: \{val}")
  // println: val: [Some(11), None, Some(33)]
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[Int]), words=6, memBlock=[1 0 0 0 241 6 0 0 11 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 33 0 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 32 191 0 0 3 0 0 0]
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_int_option() -> Array[Int?] {
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 32 191 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 11 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 33 0 0 0 0 0 0 0]
  get_int_option_array()
}

///|
fn get_int_option_array() -> Array[Int?] {
  let a = Some(11)
  let b = None
  let c = Some(33)
  [a, b, c]
}

///|
pub fn test2d_array_input_string(val : Array[Array[String]]) -> Unit!Error {
  let expected = [
    ["abc", "def", "ghi"],
    ["jkl", "mno", "pqr"],
    ["stu", "vwx", "yz"],
  ]
  assert_eq!(expected, val)
}

///|
pub fn test2d_array_output_string() -> Array[Array[String]] {
  // memoryBlockAtOffset(offset: 49120=0x0000BFE0=[224 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 191 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 49088=0x0000BFC0=[192 191 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 32 191 0 0 96 191 0 0 160 191 0 0]
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 32 59 0 0 208 62 0 0 104 75 0 0]
  // memoryBlockAtOffset(offset: 15136=0x00003B20=[32 59 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
  // memoryBlockAtOffset(offset: 16080=0x00003ED0=[208 62 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 100 0 101 0 102 0 0 1] = 'def'
  // memoryBlockAtOffset(offset: 19304=0x00004B68=[104 75 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 103 0 104 0 105 0 0 1] = 'ghi'
  // memoryBlockAtOffset(offset: 48992=0x0000BF60=[96 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 191 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48960=0x0000BF40=[64 191 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 48 81 0 0 24 81 0 0 0 81 0 0]
  // memoryBlockAtOffset(offset: 20784=0x00005130=[48 81 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 106 0 107 0 108 0 0 1] = 'jkl'
  // memoryBlockAtOffset(offset: 20760=0x00005118=[24 81 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 109 0 110 0 111 0 0 1] = 'mno'
  // memoryBlockAtOffset(offset: 20736=0x00005100=[0 81 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 112 0 113 0 114 0 0 1] = 'pqr'
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 191 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 49024=0x0000BF80=[128 191 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 232 80 0 0 208 80 0 0 184 80 0 0]
  // memoryBlockAtOffset(offset: 20712=0x000050E8=[232 80 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 115 0 116 0 117 0 0 1] = 'stu'
  // memoryBlockAtOffset(offset: 20688=0x000050D0=[208 80 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 118 0 119 0 120 0 0 1] = 'vwx'
  // memoryBlockAtOffset(offset: 20664=0x000050B8=[184 80 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 121 0 122 0 0 0 0 3] = 'yz'
  [["abc", "def", "ghi"], ["jkl", "mno", "pqr"], ["stu", "vwx", "yz"]]
}

///|
pub fn test2d_array_input_string_none(
  val : Array[Array[String]]?
) -> Unit!Error {
  // memoryBlockAtOffset(offset: 0) = (data=0, size=0)
  assert_eq!(val, None)
}

///|
pub fn test2d_array_output_string_none() -> Array[Array[String]]? {
  // pointerHandler.Decode(vals: [0])
  None
}

///|
pub fn test2d_array_input_string_empty(
  val : Array[Array[String]]
) -> Unit!Error {
  let expected = []
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
  assert_eq!(expected, val)
}

///|
pub fn test2d_array_output_string_empty() -> Array[Array[String]] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=242(), words=0, memBlock=[1 0 0 0 242 0 0 0]
  []
}

///|
pub fn test2d_array_input_string_inner_empty(
  val : Array[Array[String]]
) -> Unit!Error {
  let expected = [[]]
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  // memoryBlockAtOffset(offset: 48944=0x0000BF30=[48 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 32 191 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 48 191 0 0]
  // memoryBlockAtOffset(offset: 48976=0x0000BF50=[80 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 1 0 0 0]
  assert_eq!(expected, val)
}

///|
pub fn test2d_array_output_string_inner_empty() -> Array[Array[String]] {
  // memoryBlockAtOffset(offset: 48960=0x0000BF40=[64 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 80 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48720=0x0000BE50=[80 190 0 0], size: 12=8+words*4), moonBitType=242(), words=1, memBlock=[1 0 0 0 242 1 0 0 32 191 0 0]
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=242(), words=0, memBlock=[1 0 0 0 242 0 0 0]
  [[]]
}

///|
pub fn test2d_array_input_string_inner_none(
  val : Array[Array[String]?]
) -> Unit!Error {
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 1 0 0 0]
  assert_eq!(val, [None])
}

///|
pub fn test2d_array_output_string_inner_none() -> Array[Array[String]?] {
  // memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=242(), words=1, memBlock=[1 0 0 0 242 1 0 0 0 0 0 0]
  [None]
}
