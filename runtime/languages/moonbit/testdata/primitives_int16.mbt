// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_int16_input_min(n : Int16) -> Unit!Error {
  assert_eq!(n, -32768) // TODO: @int16.min_value)
}

///|
pub fn test_int16_input_max(n : Int16) -> Unit!Error {
  assert_eq!(n, 32767) // TODO: @int16.max_value)
}

///|
pub fn test_int16_output_min() -> Int16 {
  -32768 // TODO: @int16.min_value
}

///|
pub fn test_int16_output_max() -> Int16 {
  32767 // TODO: @int16.max_value
}

///|
pub fn test_int16_option_input_min(n : Int16?) -> Unit!Error {
  assert_eq!(n, Some(-32768)) // TODO: @int16.min_value)
}

///|
pub fn test_int16_option_input_max(n : Int16?) -> Unit!Error {
  assert_eq!(n, Some(32767)) // TODO: @int16.max_value)
}

///|
pub fn test_int16_option_input_none(n : Int16?) -> Unit!Error {
  assert_eq!(n, None)
}

///|
pub fn test_int16_option_output_min() -> Int16? {
  Some(-32768) // TODO: @int16.min_value
}

///|
pub fn test_int16_option_output_max() -> Int16? {
  Some(32767) // TODO: @int16.max_value
}

///|
pub fn test_int16_option_output_none() -> Int16? {
  None
}
