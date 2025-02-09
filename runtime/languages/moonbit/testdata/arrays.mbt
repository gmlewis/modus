// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_byte(val : Array[Byte]) -> Unit!Error {
  let expected : Array[Byte] = [1, 2, 3, 4]
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_byte() -> Array[Byte] {
  [1, 2, 3, 4]
}

///|
pub fn test_array_input_string(val : Array[String]) -> Unit!Error {
  let expected = ["abc", "def", "ghi"]
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_string() -> Array[String] {
  ["abc", "def", "ghi"]
}

///|
pub fn test_array_input_string_none(val : Array[String]?) -> Unit!Error {
  assert_eq!(val, None)
}

///|
pub fn test_array_output_string_none() -> Array[String]? {
  None
}

///|
pub fn test_array_input_string_empty(val : Array[String]) -> Unit!Error {
  let expected = []
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_string_empty() -> Array[String] {
  []
}

///|
pub fn test_array_input_int32_empty(val : Array[Int]) -> Unit!Error {
  let expected = []
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_int32_empty() -> Array[Int] {
  []
}

///|
pub fn test_array_input_string_option(val : Array[String?]) -> Unit!Error {
  let expected = get_string_option_array()
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_string_option() -> Array[String?] {
  get_string_option_array()
}

///|
fn get_string_option_array() -> Array[String?] {
  let a = Some("abc")
  let b = Some("def")
  let c = Some("ghi")
  [a, b, c]
}

///|
pub fn test_array_input_int_option(val : Array[Int?]) -> Unit!Error {
  let expected = get_int_option_array()
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_int_option() -> Array[Int?] {
  get_int_option_array()
}

///|
fn get_int_option_array() -> Array[Int?] {
  let a = Some(11)
  let b = Some(22)
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
  [["abc", "def", "ghi"], ["jkl", "mno", "pqr"], ["stu", "vwx", "yz"]]
}

///|
pub fn test2d_array_input_string_none(
  val : Array[Array[String]]?
) -> Unit!Error {
  assert_eq!(val, None)
}

///|
pub fn test2d_array_output_string_none() -> Array[Array[String]]? {
  None
}

///|
pub fn test2d_array_input_string_empty(
  val : Array[Array[String]]
) -> Unit!Error {
  let expected = []
  assert_eq!(expected, val)
}

///|
pub fn test2d_array_output_string_empty() -> Array[Array[String]] {
  []
}

///|
pub fn test2d_array_input_string_inner_empty(
  val : Array[Array[String]]
) -> Unit!Error {
  let expected = [[]]
  assert_eq!(expected, val)
}

///|
pub fn test2d_array_output_string_inner_empty() -> Array[Array[String]] {
  [[]]
}

///|
pub fn test2d_array_input_string_inner_none(
  val : Array[Array[String]?]
) -> Unit!Error {
  assert_eq!(val, [None])
}

///|
pub fn test2d_array_output_string_inner_none() -> Array[Array[String]?] {
  [None]
}
