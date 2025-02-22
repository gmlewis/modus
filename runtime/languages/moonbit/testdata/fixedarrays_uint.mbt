// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_input_uint_empty(val : FixedArray[UInt]) -> Unit!Error {
  let expected : FixedArray[UInt] = []
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_uint_0() -> FixedArray[UInt] {
  []
}

///|
pub fn test_fixedarray_output_uint_1() -> FixedArray[UInt] {
  [1]
}

///|
pub fn test_fixedarray_output_uint_1_min() -> FixedArray[UInt] {
  [@uint.min_value]
}

///|
pub fn test_fixedarray_output_uint_1_max() -> FixedArray[UInt] {
  [@uint.max_value]
}

///|
pub fn test_fixedarray_output_uint_2() -> FixedArray[UInt] {
  [@uint.min_value, @uint.max_value]
}

///|
pub fn test_fixedarray_output_uint_3() -> FixedArray[UInt] {
  [1, 2, 3]
}

///|
pub fn test_fixedarray_output_uint_4() -> FixedArray[UInt] {
  [1, 2, 3, 4]
}

///|
pub fn test_fixedarray_output_uint_option_0() -> FixedArray[UInt?] {
  []
}

///|
pub fn test_fixedarray_output_uint_option_1_none() -> FixedArray[UInt?] {
  [None]
}

///|
pub fn test_fixedarray_output_uint_option_1_min() -> FixedArray[UInt?] {
  [Some(@uint.min_value)]
}

///|
pub fn test_fixedarray_output_uint_option_1_max() -> FixedArray[UInt?] {
  [Some(@uint.max_value)]
}

///|
pub fn test_fixedarray_output_uint_option_2() -> FixedArray[UInt?] {
  [Some(1), Some(2)]
}

///|
pub fn test_fixedarray_output_uint_option_3() -> FixedArray[UInt?] {
  [None, None, None]
}

///|
pub fn test_fixedarray_output_uint_option_4() -> FixedArray[UInt?] {
  [None, Some(@uint.min_value), Some(0), Some(@uint.max_value)]
}

///|
pub fn test_fixedarray_input_uint_option(val : FixedArray[UInt?]) -> Unit!Error {
  let expected = get_uint_option_fixedarray()
  // pruintln("expected: \{expected}")
  // pruintln("val: \{val}")
  // pruintln: val: [Some(11), None, Some(33)]
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_uint_option() -> FixedArray[UInt?] {
  get_uint_option_fixedarray()
}

///|
fn get_uint_option_fixedarray() -> FixedArray[UInt?] {
  let a = Some(11U)
  let b = None
  let c = Some(33U)
  [a, b, c]
}
