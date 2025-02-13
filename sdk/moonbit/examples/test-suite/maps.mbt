// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn hello_maps_n_items(n~ : Int = 3) -> Map[String, String] {
  if n < 0 {
    println("hello_maps_n_items(\{n}): returning: {}")
    return {}
  }
  let map = {}
  for i in 1..=n {
    let key = "key\{i}"
    let value = "value\{i}"
    map[key] = value
  }
  println("hello_maps_n_items(\{n}): returning: \{map}")
  map
}

///|
pub fn hello_maps_n_items_option(n~ : Int = 3) -> Map[String, String]? {
  if n < 0 {
    println("hello_maps_n_items_option(\{n}): returning: None")
    return None
  }
  let map = {}
  for i in 1..=n {
    let key = "key\{i}"
    let value = "value\{i}"
    map[key] = value
  }
  println("hello_maps_n_items_option(\{n}): returning: Some(\{map})")
  Some(map)
}

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

///| This generated map mimics the test map created on the Go side.
pub fn make_test_map(size : Int) -> Map[String, String] {
  let m = {}
  for i in 0..<size {
    let s = "\{1000+i}".substring(start=1)
    let key = "key_\{s}"
    let val = "val_\{s}"
    m[key] = val
  }
  m
}
