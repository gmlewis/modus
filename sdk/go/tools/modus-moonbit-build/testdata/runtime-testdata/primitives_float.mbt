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
  // primitiveHandler.Decode(vals: [4286578687])
  @float.min_value
}

///|
pub fn test_float_output_max() -> Float {
  // primitiveHandler.Decode(vals: [2139095039])
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
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 255 255 127 255]
  Some(@float.min_value)
}

///|
pub fn test_float_option_output_max() -> Float? {
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 255 255 127 127]
  Some(@float.max_value)
}

///|
pub fn test_float_option_output_none() -> Float? {
  // memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
  None
}
