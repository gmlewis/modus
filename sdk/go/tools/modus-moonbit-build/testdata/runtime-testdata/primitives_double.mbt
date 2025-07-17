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
  // primitiveHandler.Decode(vals: [18442240474082181119])
  @double.min_value
}

///|
pub fn test_double_output_max() -> Double {
  // primitiveHandler.Decode(vals: [9218868437227405311])
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
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 255 255 255 255 255 255 239 255]
  Some(@double.min_value)
}

///|
pub fn test_double_option_output_max() -> Double? {
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 255 255 255 255 255 255 239 127]
  Some(@double.max_value)
}

///|
pub fn test_double_option_output_none() -> Double? {
  // memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
  None
}
