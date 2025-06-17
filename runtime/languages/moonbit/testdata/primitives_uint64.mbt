// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_uint64_input_min(n : UInt64) -> Unit raise Error {
  assert_eq(n, 0)
}

///|
pub fn test_uint64_input_max(n : UInt64) -> Unit raise Error {
  assert_eq(n, @uint64.max_value)
}

///|
pub fn test_uint64_output_min() -> UInt64 {
  0
}

///|
pub fn test_uint64_output_max() -> UInt64 {
  @uint64.max_value
}

///|
pub fn test_uint64_option_input_min(n : UInt64?) -> Unit raise Error {
  assert_eq(n, Some(0))
}

///|
pub fn test_uint64_option_input_max(n : UInt64?) -> Unit raise Error {
  assert_eq(n, Some(@uint64.max_value))
}

///|
pub fn test_uint64_option_input_none(n : UInt64?) -> Unit raise Error {
  assert_eq(n, None)
}

///|
pub fn test_uint64_option_output_min() -> UInt64? {
  Some(0)
}

///|
pub fn test_uint64_option_output_max() -> UInt64? {
  Some(@uint64.max_value)
}

///|
pub fn test_uint64_option_output_none() -> UInt64? {
  None
}
