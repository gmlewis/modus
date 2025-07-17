// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_byte_input_min(b : Byte) -> Unit raise Error {
  assert_eq(b, 0)
}

///|
pub fn test_byte_input_max(b : Byte) -> Unit raise Error {
  assert_eq(@byte.max_value, b)
}

///|
pub fn test_byte_output_min() -> Byte {
  // primitiveHandler.Decode(vals: [0])
  0
}

///|
pub fn test_byte_output_max() -> Byte {
  // primitiveHandler.Decode(vals: [255])
  @byte.max_value
}

///|
pub fn test_byte_option_input_min(b : Byte?) -> Unit raise Error {
  assert_eq(b, Some(0))
}

///|
pub fn test_byte_option_input_max(b : Byte?) -> Unit raise Error {
  assert_eq(b, Some(@byte.max_value))
}

///|
pub fn test_byte_option_input_none(b : Byte?) -> Unit raise Error {
  assert_eq(None, b)
}

///|
pub fn test_byte_option_output_min() -> Byte? {
  // primitiveHandler.Decode(vals: [0])
  Some(0)
}

///|
pub fn test_byte_option_output_max() -> Byte? {
  // primitiveHandler.Decode(vals: [255])
  Some(@byte.max_value)
}

///|
pub fn test_byte_option_output_none() -> Byte? {
  // primitiveHandler.Decode(vals: [4294967295])
  None
}
