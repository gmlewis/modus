// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| "Hello World" in Japanese
const TEST_STRING = "こんにちは、世界"

///|
pub fn test_string_input(s : String) -> Unit!Error {
  assert_eq!(TEST_STRING, s)
}

///|
pub fn test_string_option_input(s : String?) -> Unit!Error {
  assert_eq!(s, Some(TEST_STRING))
}

///|
pub fn test_string_option_input_none(s : String?) -> Unit!Error {
  assert_eq!(s, None)
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
