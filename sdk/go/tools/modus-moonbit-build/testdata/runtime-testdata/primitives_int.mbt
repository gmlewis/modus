// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_int_input_min(n : Int) -> Unit raise Error {
  assert_eq(@int.min_value, n)
}

///|
pub fn test_int_input_max(n : Int) -> Unit raise Error {
  assert_eq(@int.max_value, n)
}

///|
pub fn test_int_output_min() -> Int {
  @int.min_value
}

///|
pub fn test_int_output_max() -> Int {
  @int.max_value
}

///|
pub fn test_int_option_input_min(n : Int?) -> Unit raise Error {
  assert_eq(n, Some(@int.min_value))
}

///|
pub fn test_int_option_input_max(n : Int?) -> Unit raise Error {
  assert_eq(n, Some(@int.max_value))
}

///|
pub fn test_int_option_input_none(n : Int?) -> Unit raise Error {
  assert_eq(n, None)
}

///|
pub fn test_int_option_output_min() -> Int? {
  Some(@int.min_value)
}

///|
pub fn test_int_option_output_max() -> Int? {
  Some(@int.max_value)
}

///|
pub fn test_int_option_output_none() -> Int? {
  None
}
