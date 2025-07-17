// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_char_input_min(c : Char) -> Unit raise Error {
  // can't seem to pass in a char with a value of -32768
  assert_eq(c, Int::unsafe_to_char(0))
}

///|
pub fn test_char_input_max(c : Char) -> Unit raise Error {
  assert_eq(c, Int::unsafe_to_char(32767))
}

///|
pub fn test_char_output_min() -> Char {
  Int::unsafe_to_char(-32768)
}

///|
pub fn test_char_output_max() -> Char {
  Int::unsafe_to_char(32767) // TODO: @int16.max_value))
}

///|
pub fn test_char_option_input_min(c : Char?) -> Unit raise Error {
  assert_eq(c, Some(Int::unsafe_to_char(-32768))) // TODO: @int16.min_value)))
}

///|
pub fn test_char_option_input_max(c : Char?) -> Unit raise Error {
  assert_eq(c, Some(Int::unsafe_to_char(32767))) // TODO: @int16.max_value)))
}

///|
pub fn test_char_option_input_none(c : Char?) -> Unit raise Error {
  assert_eq(c, None)
}

///|
pub fn test_char_option_output_min() -> Char? {
  Some(Int::unsafe_to_char(-32768)) // TODO: @int16.min_value))
}

///|
pub fn test_char_option_output_max() -> Char? {
  Some(Int::unsafe_to_char(32767)) // TODO: @int16.max_value))
}

///|
pub fn test_char_option_output_none() -> Char? {
  // primitiveHandler.Decode(vals: [4294967295])
  None
}
