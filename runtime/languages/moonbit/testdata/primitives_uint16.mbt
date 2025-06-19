// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_uint16_input_min(n : UInt16) -> Unit raise Error {
  assert_eq(n, 0)
}

///|
pub fn test_uint16_input_max(n : UInt16) -> Unit raise Error {
  assert_eq(n, 65535) // TODO: @uint16.max_value)
}

///|
pub fn test_uint16_output_min() -> UInt16 {
  0
}

///|
pub fn test_uint16_output_max() -> UInt16 {
  65535 // TODO: @uint16.max_value
}

///|
pub fn test_uint16_option_input_min(n : UInt16?) -> Unit raise Error {
  assert_eq(n, Some(0))
}

///|
pub fn test_uint16_option_input_max(n : UInt16?) -> Unit raise Error {
  assert_eq(n, Some(65535)) // TODO: @uint16.max_value)
}

///|
pub fn test_uint16_option_input_none(n : UInt16?) -> Unit raise Error {
  assert_eq(n, None)
}

///|
pub fn test_uint16_option_output_min() -> UInt16? {
  Some(0)
}

///|
pub fn test_uint16_option_output_max() -> UInt16? {
  Some(65535) // TODO: @uint16.max_value
}

///|
pub fn test_uint16_option_output_none() -> UInt16? {
  None
}
