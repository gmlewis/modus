// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_int64_input_min(n : Int64) -> Unit raise Error {
  assert_eq(n, @int64.min_value)
}

///|
pub fn test_int64_input_max(n : Int64) -> Unit raise Error {
  assert_eq(n, @int64.max_value)
}

///|
pub fn test_int64_output_min() -> Int64 {
  @int64.min_value
}

///|
pub fn test_int64_output_max() -> Int64 {
  @int64.max_value
}

///|
pub fn test_int64_option_input_min(n : Int64?) -> Unit raise Error {
  assert_eq(n, Some(@int64.min_value))
}

///|
pub fn test_int64_option_input_max(n : Int64?) -> Unit raise Error {
  assert_eq(n, Some(@int64.max_value))
}

///|
pub fn test_int64_option_input_none(n : Int64?) -> Unit raise Error {
  assert_eq(n, None)
}

///|
pub fn test_int64_option_output_min() -> Int64? {
  Some(@int64.min_value)
}

///|
pub fn test_int64_option_output_max() -> Int64? {
  Some(@int64.max_value)
}

///|
pub fn test_int64_option_output_none() -> Int64? {
  None
}
