// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_bool(val : Array[Bool]) -> Unit!Error {
  let expected : Array[Bool] = [false, true, true, false]
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_bool_0() -> Array[Bool] {
  []
}

///|
pub fn test_array_output_bool_1() -> Array[Bool] {
  [true]
}

///|
pub fn test_array_output_bool_2() -> Array[Bool] {
  [false, true]
}

///|
pub fn test_array_output_bool_3() -> Array[Bool] {
  [true, true, true]
}

///|
pub fn test_array_output_bool_4() -> Array[Bool] {
  [false, false, false, false]
}

///|
pub fn test_array_output_bool_option_0() -> Array[Bool?] {
  []
}

///|
pub fn test_array_output_bool_option_1() -> Array[Bool?] {
  [Some(false)]
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
  []
}

///|
pub fn test_array_output_byte_1() -> Array[Byte] {
  [1]
}

///|
pub fn test_array_output_byte_2() -> Array[Byte] {
  [1, 2]
}

///|
pub fn test_array_output_byte_3() -> Array[Byte] {
  [1, 2, 3]
}

///|
pub fn test_array_output_byte_4() -> Array[Byte] {
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_byte_option_0() -> Array[Byte?] {
  []
}

///|
pub fn test_array_output_byte_option_1() -> Array[Byte?] {
  [Some(1)]
}

///|
pub fn test_array_output_byte_option_2() -> Array[Byte?] {
  [Some(1), None]
}

///|
pub fn test_array_output_byte_option_3() -> Array[Byte?] {
  [None, Some(2), Some(3)]
}

///|
pub fn test_array_output_byte_option_4() -> Array[Byte?] {
  [Some(1), Some(2), Some(3), None]
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
pub fn test_array_input_int_empty(val : Array[Int]) -> Unit!Error {
  let expected = []
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_int_0() -> Array[Int] {
  []
}

///|
pub fn test_array_output_int_1() -> Array[Int] {
  [1]
}

///|
pub fn test_array_output_int_2() -> Array[Int] {
  [1, 2]
}

///|
pub fn test_array_output_int_3() -> Array[Int] {
  [1, 2, 3]
}

///|
pub fn test_array_output_int_4() -> Array[Int] {
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_int_option_0() -> Array[Int?] {
  []
}

///|
pub fn test_array_output_int_option_1() -> Array[Int?] {
  [None]
}

///|
pub fn test_array_output_int_option_2() -> Array[Int?] {
  [Some(1), Some(2)]
}

///|
pub fn test_array_output_int_option_3() -> Array[Int?] {
  [None, None, None]
}

///|
pub fn test_array_output_int_option_4() -> Array[Int?] {
  [Some(1), Some(2), Some(3), Some(4)]
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
  let b = None
  let c = Some("ghi")
  [a, b, c]
}

///|
pub fn test_array_input_int_option(val : Array[Int?]) -> Unit!Error {
  let expected = get_int_option_array()
  println("expected: \{expected}")
  println("val: \{val}")
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_int_option() -> Array[Int?] {
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
