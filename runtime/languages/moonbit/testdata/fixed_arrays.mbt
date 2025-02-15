// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_input0_byte(val : FixedArray[Byte]) -> Unit!Error {
  let expected : FixedArray[Byte] = []
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output0_byte() -> FixedArray[Byte] {
  []
}

///|
pub fn test_fixedarray_input0_string(val : FixedArray[String]) -> Unit!Error {
  let expected : FixedArray[String] = []
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output0_string() -> FixedArray[String] {
  []
}

///|
pub fn test_fixedarray_input0_string_option(
  val : FixedArray[String?]
) -> Unit!Error {
  let expected : FixedArray[String?] = []
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output0_string_option() -> FixedArray[String?] {
  []
}

///|
pub fn test_fixedarray_input0_int_option(val : FixedArray[Int?]) -> Unit!Error {
  let expected = get_int_option_fixedarray()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output0_int_option() -> FixedArray[Int?] {
  get_int_option_fixedarray()
}

///|
fn get_int_option_fixedarray() -> FixedArray[Int?] {
  []
}

///|
pub fn test_fixedarray_input1_byte(val : FixedArray[Byte]) -> Unit!Error {
  let expected : FixedArray[Byte] = [1]
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output1_byte() -> FixedArray[Byte] {
  [1]
}

///|
pub fn test_fixedarray_input1_string(val : FixedArray[String]) -> Unit!Error {
  let expected : FixedArray[String] = ["abc"]
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output1_string() -> FixedArray[String] {
  ["abc"]
}

///|
pub fn test_fixedarray_input1_string_option(
  val : FixedArray[String?]
) -> Unit!Error {
  let expected = get_string_option_fixedarray1()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output1_string_option() -> FixedArray[String?] {
  get_string_option_fixedarray1()
}

///|
pub fn get_string_option_fixedarray1() -> FixedArray[String?] {
  [Some("abc")]
}

///|
pub fn test_fixedarray_input1_int_option(val : FixedArray[Int?]) -> Unit!Error {
  let expected = get_int_option_fixedarray1()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output1_int_option() -> FixedArray[Int?] {
  get_int_option_fixedarray1()
}

///|
pub fn get_int_option_fixedarray1() -> FixedArray[Int?] {
  [Some(11)]
}

///|
pub fn test_fixedarray_input2_byte(val : FixedArray[Byte]) -> Unit!Error {
  let expected : FixedArray[Byte] = [1, 2]
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output2_byte() -> FixedArray[Byte] {
  [1, 2]
}

///|
pub fn test_fixedarray_input2_string(val : FixedArray[String]) -> Unit!Error {
  let expected : FixedArray[String] = ["abc", "def"]
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output2_string() -> FixedArray[String] {
  ["abc", "def"]
}

///|
pub fn test_fixedarray_input2_string_option(
  val : FixedArray[String?]
) -> Unit!Error {
  let expected = get_string_option_array2()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output2_string_option() -> FixedArray[String?] {
  get_string_option_array2()
}

///|
pub fn get_string_option_array2() -> FixedArray[String?] {
  [Some("abc"), Some("def")]
}

///|
pub fn test_fixedarray_input2_int_option(val : FixedArray[Int?]) -> Unit!Error {
  let expected = get_int_ptr_array2()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output2_int_option() -> FixedArray[Int?] {
  get_int_ptr_array2()
}

///|
pub fn get_int_ptr_array2() -> FixedArray[Int?] {
  [Some(11), Some(22)]
}

///|
pub fn test_fixedarray_input2_struct(
  val : FixedArray[TestStruct2]
) -> Unit!Error {
  let expected = get_struct_fixedarray2()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output2_struct() -> FixedArray[TestStruct2] {
  get_struct_fixedarray2()
}

///|
pub fn test_fixedarray_input2_struct_option(
  val : FixedArray[TestStruct2?]
) -> Unit!Error {
  let expected = get_struct_option_fixedarray2()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output2_struct_option() -> FixedArray[TestStruct2?] {
  get_struct_option_fixedarray2()
}

///|
fn get_struct_fixedarray2() -> FixedArray[TestStruct2] {
  [{ a: true, b: 123 }, { a: false, b: 456 }]
}

///|
fn get_struct_option_fixedarray2() -> FixedArray[TestStruct2?] {
  [Some({ a: true, b: 123 }), Some({ a: false, b: 456 })]
}

///|
pub fn test_fixedarray_input2_map(
  val : FixedArray[Map[String, String]]
) -> Unit!Error {
  let expected = get_map_array2()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output2_map() -> FixedArray[Map[String, String]] {
  get_map_array2()
}

///|
pub fn test_fixedarray_input2_map_option(
  val : FixedArray[Map[String, String]?]
) -> Unit!Error {
  let expected = getMapPtrArray2()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output2_map_option() -> FixedArray[Map[String, String]?] {
  getMapPtrArray2()
}

///|
pub fn get_map_array2() -> FixedArray[Map[String, String]] {
  [{ "A": "true", "B": "123" }, { "C": "false", "D": "456" }]
}

///|
pub fn getMapPtrArray2() -> FixedArray[Map[String, String]?] {
  [Some({ "A": "true", "B": "123" }), Some({ "C": "false", "D": "456" })]
}

///|
pub fn test_option_fixedarray_input1_int(val : FixedArray[Int]?) -> Unit!Error {
  let expected = get_option_int_fixedarray1()
  assert_eq!(expected, val)
}

///|
pub fn test_option_fixedarray_output1_int() -> FixedArray[Int]? {
  get_option_int_fixedarray1()
}

///|
pub fn get_option_int_fixedarray1() -> FixedArray[Int]? {
  Some([11])
}

///|
pub fn test_option_fixedarray_input2_int(val : FixedArray[Int]?) -> Unit!Error {
  let expected = get_option_int_fixedarray2()
  assert_eq!(expected, val)
}

///|
pub fn test_option_fixedarray_output2_int() -> FixedArray[Int]? {
  get_option_int_fixedarray2()
}

///|
pub fn get_option_int_fixedarray2() -> FixedArray[Int]? {
  Some([11, 22])
}

///|
pub fn test_option_fixedarray_input1_string(
  val : FixedArray[String]?
) -> Unit!Error {
  let expected = get_option_string_fixedarray1()
  assert_eq!(expected, val)
}

///|
pub fn test_option_fixedarray_output1_string() -> FixedArray[String]? {
  get_option_string_fixedarray1()
}

///|
pub fn get_option_string_fixedarray1() -> FixedArray[String]? {
  Some(["abc"])
}

///|
pub fn get_option_string_fixedarray2() -> FixedArray[String]? {
  Some(["abc", "def"])
}

///|
pub fn test_option_fixedarray_input2_string(
  val : FixedArray[String]?
) -> Unit!Error {
  let expected = get_option_string_fixedarray2()
  assert_eq!(expected, val)
}

///|
pub fn test_option_fixedarray_output2_string() -> FixedArray[String]? {
  get_option_string_fixedarray2()
}
