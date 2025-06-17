// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_uint_input_min(n : UInt) -> Unit raise Error {
  assert_eq(n, 0)
}

///|
pub fn test_uint_input_max(n : UInt) -> Unit raise Error {
  assert_eq(n, @uint.max_value)
}

///|
pub fn test_uint_output_min() -> UInt {
  0
}

///|
pub fn test_uint_output_max() -> UInt {
  @uint.max_value
}

///|
pub fn test_uint_option_input_min(n : UInt?) -> Unit raise Error {
  assert_eq(n, Some(0))
}

///|
pub fn test_uint_option_input_max(n : UInt?) -> Unit raise Error {
  assert_eq(n, Some(@uint.max_value))
}

///|
pub fn test_uint_option_input_none(n : UInt?) -> Unit raise Error {
  assert_eq(n, None)
}

///|
pub fn test_uint_option_output_min() -> UInt? {
  Some(0)
}

///|
pub fn test_uint_option_output_max() -> UInt? {
  Some(@uint.max_value)
}

///|
pub fn test_uint_option_output_none() -> UInt? {
  None
}
