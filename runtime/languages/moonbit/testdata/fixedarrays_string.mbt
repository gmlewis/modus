// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_input_string(val : FixedArray[String]) -> Unit!Error {
  let expected : FixedArray[String] = ["abc", "def", "ghi"]
  // GML: handler_memory.go: memoryBlockAtOffset(offset: 49024=0x0000BF80=[128 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_string() -> FixedArray[String] {
  ["abc", "def", "ghi"]
}

///|
pub fn test_fixedarray_input_string_none(
  val : FixedArray[String]?
) -> Unit!Error {
  assert_eq!(val, None)
}

///|
pub fn test_fixedarray_output_string_none() -> FixedArray[String]? {
  // pointerHandler.Decode(vals: [0])
  None
}

///|
pub fn test_fixedarray_input_string_empty(
  val : FixedArray[String]
) -> Unit!Error {
  let expected : FixedArray[String] = []
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_string_empty() -> FixedArray[String] {
  []
}

///|
pub fn test_fixedarray_input_string_option(
  val : FixedArray[String?]
) -> Unit!Error {
  let expected = get_string_option_fixedarray()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_string_option() -> FixedArray[String?] {
  get_string_option_fixedarray()
}

///|
fn get_string_option_fixedarray() -> FixedArray[String?] {
  let a = Some("abc")
  let b = None
  let c = Some("ghi")
  [a, b, c]
}

///|
pub fn test2d_fixedarray_input_string(
  val : FixedArray[FixedArray[String]]
) -> Unit!Error {
  let expected : FixedArray[FixedArray[String]] = [
    ["abc", "def", "ghi"],
    ["jkl", "mno", "pqr"],
    ["stu", "vwx", "yz"],
  ]
  assert_eq!(expected, val)
}

///|
pub fn test2d_fixedarray_output_string() -> FixedArray[FixedArray[String]] {
  [["abc", "def", "ghi"], ["jkl", "mno", "pqr"], ["stu", "vwx", "yz"]]
}

///|
pub fn test2d_fixedarray_input_string_none(
  val : FixedArray[FixedArray[String]]?
) -> Unit!Error {
  assert_eq!(val, None)
}

///|
pub fn test2d_fixedarray_output_string_none() -> FixedArray[FixedArray[String]]? {
  // pointerHandler.Decode(vals: [0])
  None
}

///|
pub fn test2d_fixedarray_input_string_empty(
  val : FixedArray[FixedArray[String]]
) -> Unit!Error {
  let expected : FixedArray[FixedArray[String]] = []
  assert_eq!(expected, val)
}

///|
pub fn test2d_fixedarray_output_string_empty() -> FixedArray[FixedArray[String]] {
  []
}

///|
pub fn test2d_fixedarray_input_string_inner_empty(
  val : FixedArray[FixedArray[String]]
) -> Unit!Error {
  let expected : FixedArray[FixedArray[String]] = [[]]
  assert_eq!(expected, val)
}

///|
pub fn test2d_fixedarray_output_string_inner_empty() ->
     FixedArray[FixedArray[String]] {
  [[]]
}

///|
pub fn test2d_fixedarray_input_string_inner_none(
  val : FixedArray[FixedArray[String]?]
) -> Unit!Error {
  assert_eq!(val, [None])
}

///|
pub fn test2d_fixedarray_output_string_inner_none() ->
     FixedArray[FixedArray[String]?] {
  [None]
}
