// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_string(val : Array[String]) -> Unit!Error {
  let expected = ["abc", "def", "ghi"]
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 32 191 0 0 64 191 0 0 96 191 0 0]
  // memoryBlockAtOffset(offset: 49024=0x0000BF80=[128 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
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
