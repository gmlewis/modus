// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_int64_input_min(n : Int64) -> Unit!Error {
  assert_eq!(n, @int64.min_value)
}

///|
pub fn test_int64_input_max(n : Int64) -> Unit!Error {
  assert_eq!(n, @int64.max_value)
}

///|
pub fn test_int64_output_min() -> Int64 {
  // primitiveHandler.Decode(vals: [9223372036854775808])
  @int64.min_value
}

///|
pub fn test_int64_output_max() -> Int64 {
  // primitiveHandler.Decode(vals: [9223372036854775807])
  @int64.max_value
}

///|
pub fn test_int64_option_input_min(n : Int64?) -> Unit!Error {
  assert_eq!(n, Some(@int64.min_value))
}

///|
pub fn test_int64_option_input_max(n : Int64?) -> Unit!Error {
  assert_eq!(n, Some(@int64.max_value))
}

///|
pub fn test_int64_option_input_none(n : Int64?) -> Unit!Error {
  assert_eq!(n, None)
}

///|
pub fn test_int64_option_output_min() -> Int64? {
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 128]
  Some(@int64.min_value)
}

///|
pub fn test_int64_option_output_max() -> Int64? {
  // memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 255 255 255 255 255 255 255 127]
  Some(@int64.max_value)
}

///|
pub fn test_int64_option_output_none() -> Int64? {
  // memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
  None
}
