// -*- compile-command: "./build.sh"; -*-

// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| "Hello World" in Japanese
const TEST_STRING = "こんにちは、世界"

///|
pub fn test_string_input(s : String) -> Unit raise Error {
  assert_eq(TEST_STRING, s)
}

///|
pub fn test_string_option_input(s : String?) -> Unit raise Error {
  assert_eq(s, Some(TEST_STRING))
}

///|
pub fn test_string_option_input_none(s : String?) -> Unit raise Error {
  assert_eq(s, None)
}

///|
pub fn test_string_output_len_0() -> String {
  ""
}

///|
pub fn test_string_output_len_1() -> String {
  "1"
}

///|
pub fn test_string_output_len_2() -> String {
  "12"
}

///|
pub fn test_string_output_len_3() -> String {
  "123"
}

///|
pub fn test_string_output_len_4() -> String {
  "1234"
}

///|
pub fn test_string_output_len_5() -> String {
  "12345"
}

///|
pub fn test_string_output_len_6() -> String {
  "123456"
}

///|
pub fn test_string_output_len_7() -> String {
  "1234567"
}

///|
pub fn test_string_output_len_8() -> String {
  "12345678"
}

///|
pub fn test_string_output_len_9() -> String {
  "123456789"
}

///|
pub fn test_string_output_len_10() -> String {
  "1234567890"
}

///|
pub fn test_string_output() -> String {
  TEST_STRING
}

///|
pub fn test_string_option_output() -> String? {
  Some(TEST_STRING)
}

///|
pub fn test_string_option_output_none() -> String? {
  None
}
