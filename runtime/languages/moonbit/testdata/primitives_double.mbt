// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_double_input_min(n : Double) -> Unit raise Error {
  assert_eq(n, @double.min_value)
}

///|
pub fn test_double_input_max(n : Double) -> Unit raise Error {
  assert_eq(n, @double.max_value)
}

///|
pub fn test_double_output_min() -> Double {
  @double.min_value
}

///|
pub fn test_double_output_max() -> Double {
  @double.max_value
}

///|
pub fn test_double_option_input_min(n : Double?) -> Unit raise Error {
  assert_eq(n, Some(@double.min_value))
}

///|
pub fn test_double_option_input_max(n : Double?) -> Unit raise Error {
  assert_eq(n, Some(@double.max_value))
}

///|
pub fn test_double_option_input_none(n : Double?) -> Unit raise Error {
  assert_eq(n, None)
}

///|
pub fn test_double_option_output_min() -> Double? {
  Some(@double.min_value)
}

///|
pub fn test_double_option_output_max() -> Double? {
  Some(@double.max_value)
}

///|
pub fn test_double_option_output_none() -> Double? {
  None
}
