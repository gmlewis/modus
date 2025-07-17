// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_float_input_min(n : Float) -> Unit raise Error {
  assert_eq(n, @float.min_value)
}

///|
pub fn test_float_input_max(n : Float) -> Unit raise Error {
  assert_eq(n, @float.max_value)
}

///|
pub fn test_float_output_min() -> Float {
  @float.min_value
}

///|
pub fn test_float_output_max() -> Float {
  @float.max_value
}

///|
pub fn test_float_option_input_min(n : Float?) -> Unit raise Error {
  assert_eq(n, Some(@float.min_value))
}

///|
pub fn test_float_option_input_max(n : Float?) -> Unit raise Error {
  assert_eq(n, Some(@float.max_value))
}

///|
pub fn test_float_option_input_none(n : Float?) -> Unit raise Error {
  assert_eq(n, None)
}

///|
pub fn test_float_option_output_min() -> Float? {
  Some(@float.min_value)
}

///|
pub fn test_float_option_output_max() -> Float? {
  Some(@float.max_value)
}

///|
pub fn test_float_option_output_none() -> Float? {
  None
}
