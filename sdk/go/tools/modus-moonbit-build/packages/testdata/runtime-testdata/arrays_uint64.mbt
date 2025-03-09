// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_output_uint64_0() -> Array[UInt64] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[UInt]), words=0, memBlock=[1 0 0 0 241 0 0 0]
  []
}

///|
pub fn test_array_output_uint64_1() -> Array[UInt64] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0 0 0 0 0]
  [1]
}

///|
pub fn test_array_output_uint64_1_min() -> Array[UInt64] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0 0 0 0 128]
  [@uint64.min_value]
}

///|
pub fn test_array_output_uint64_1_max() -> Array[UInt64] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[UInt]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 255 255 255 255 127]
  [@uint64.max_value]
}

///|
pub fn test_array_output_uint64_2() -> Array[UInt64] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[UInt]), words=2, memBlock=[1 0 0 0 241 2 0 0 0 0 0 0 0 0 0 128 255 255 255 255 255 255 255 127]
  [@uint64.min_value, @uint64.max_value]
}

///|
pub fn test_array_output_uint64_3() -> Array[UInt64] {
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[UInt]), words=3, memBlock=[1 0 0 0 241 3 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0 3 0 0 0 0 0 0 0]
  [1, 2, 3]
}

///|
pub fn test_array_output_uint64_4() -> Array[UInt64] {
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 40=8+words*4), moonBitType=241(FixedArray[UInt]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0 3 0 0 0 0 0 0 0 4 0 0 0 0 0 0 0]
  [1, 2, 3, 4]
}

///|
pub fn test_array_output_uint64_option_0() -> Array[UInt64?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=242(FixedArray[String]), words=0, memBlock=[1 0 0 0 242 0 0 0]
  []
}

///|
pub fn test_array_output_uint64_option_1_none() -> Array[UInt64?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 152 41 0 0]
  // memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
  [None]
}

///|
pub fn test_array_output_uint64_option_1_min() -> Array[UInt64?] {
  // memoryBlockAtOffset(offset: 49072=0x0000BFB0=[176 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 12=8+words*4), moonBitType=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 192 190 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 128]
  [Some(@uint64.min_value)]
}

///|
pub fn test_array_output_uint64_option_1_max() -> Array[UInt64?] {
  // memoryBlockAtOffset(offset: 49072=0x0000BFB0=[176 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 1 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 12=8+words*4), moonBitType=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 192 190 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 255 255 255 255 255 255 255 127]
  [Some(@uint64.max_value)]
}

///|
pub fn test_array_output_uint64_option_2() -> Array[UInt64?] {
  // memoryBlockAtOffset(offset: 49120=0x0000BFE0=[224 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 191 0 0 2 0 0 0]
  // memoryBlockAtOffset(offset: 49088=0x0000BFC0=[192 191 0 0], size: 16=8+words*4), moonBitType=242(FixedArray[String]), words=2, memBlock=[1 0 0 0 242 2 0 0 192 190 0 0 160 191 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 1 0 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 2 0 0 0 0 0 0 0]
  [Some(1), Some(2)]
}

///|
pub fn test_array_output_uint64_option_3() -> Array[UInt64?] {
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 3 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=242(FixedArray[String]), words=3, memBlock=[1 0 0 0 242 3 0 0 152 41 0 0 152 41 0 0 152 41 0 0]
  // memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
  // memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
  // memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
  [None, None, None]
}

///|
pub fn test_array_output_uint64_option_4() -> Array[UInt64?] {
  // memoryBlockAtOffset(offset: 49152=0x0000C000=[0 192 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 224 191 0 0 4 0 0 0]
  // memoryBlockAtOffset(offset: 49120=0x0000BFE0=[224 191 0 0], size: 24=8+words*4), moonBitType=242(FixedArray[String]), words=4, memBlock=[1 0 0 0 242 4 0 0 152 41 0 0 192 190 0 0 160 191 0 0 192 191 0 0]
  // memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
  // memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 128]
  // memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 0]
  // memoryBlockAtOffset(offset: 49088=0x0000BFC0=[192 191 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 255 255 255 255 255 255 255 127]
  [None, Some(@uint64.min_value), Some(0), Some(@uint64.max_value)]
}
