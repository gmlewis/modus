// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_map_input_string_string(m : Map[String, String]) -> Unit!Error {
  let expected = { "a": "1", "b": "2", "c": "3" }
  assert_eq!(expected, m)
}

///|
pub fn test_map_option_input_string_string(
  m : Map[String, String]?
) -> Unit!Error {
  let expected = Some({ "a": "1", "b": "2", "c": "3" })
  assert_eq!(expected, m)
}

///|
pub fn test_map_output_string_string() -> Map[String, String] {
  { "a": "1", "b": "2", "c": "3" }
}

///|
pub fn test_map_option_output_string_string() -> Map[String, String]? {
  Some({ "a": "1", "b": "2", "c": "3" })
}

///|
pub fn test_iterate_map_string_string(m : Map[String, String]) -> Unit {
  println("m.size(): \{m.size()}")
  for k, v in m {
    println("\{k},\{v}")
  }
}

///|
pub fn test_map_lookup_string_string(
  m : Map[String, String],
  key : String
) -> String {
  m[key].or_default()
}

///|
pub struct TestStructWithMap {
  m : Map[String, String]
}

///|
pub fn test_struct_containing_map_input_string_string(
  s : TestStructWithMap
) -> Unit!Error {
  let expected = { "a": "1", "b": "2", "c": "3" }
  assert_eq!(expected, s.m)
}

///|
pub fn test_struct_containing_map_output_string_string() -> TestStructWithMap {
  { m: { "a": "1", "b": "2", "c": "3" } }
}

///|
pub fn test_map_input_int_float(m : Map[Int, Float]) -> Unit!Error {
  let expected : Map[Int, Float] = { 1: 1.1, 2: 2.2, 3: 3.3 }
  assert_eq!(expected, m)
}

///|
pub fn test_map_output_int_float() -> Map[Int, Float] {
  { 1: 1.1, 2: 2.2, 3: 3.3 }
}

///|
pub fn test_map_input_int_double(m : Map[Int, Double]) -> Unit!Error {
  let expected = { 1: 1.1, 2: 2.2, 3: 3.3 }
  assert_eq!(expected, m)
}

///|
pub fn test_map_output_int_double() -> Map[Int, Double] {
  { 1: 1.1, 2: 2.2, 3: 3.3 }
}
