// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_input_float_empty(val : FixedArray[Float]) -> Unit!Error {
  let expected : FixedArray[Float] = []
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_float_0() -> FixedArray[Float] {
  []
}

///|
pub fn test_fixedarray_output_float_1() -> FixedArray[Float] {
  [1]
}

///|
pub fn test_fixedarray_output_float_2() -> FixedArray[Float] {
  [1, 2]
}

///|
pub fn test_fixedarray_output_float_3() -> FixedArray[Float] {
  [1, 2, 3]
}

///|
pub fn test_fixedarray_output_float_4() -> FixedArray[Float] {
  [1, 2, 3, 4]
}

///|
pub fn test_fixedarray_output_float_option_0() -> FixedArray[Float?] {
  []
}

///|
pub fn test_fixedarray_output_float_option_1_none() -> FixedArray[Float?] {
  [None]
}

///|
pub fn test_fixedarray_output_float_option_1_some() -> FixedArray[Float?] {
  [Some(1)]
}

///|
pub fn test_fixedarray_output_float_option_2() -> FixedArray[Float?] {
  [Some(1), Some(2)]
}

///|
pub fn test_fixedarray_output_float_option_3() -> FixedArray[Float?] {
  [None, None, None]
}

///|
pub fn test_fixedarray_output_float_option_4() -> FixedArray[Float?] {
  [None, Some(2), Some(0), Some(4)]
}

///|
pub fn test_fixedarray_input_float_option(
  val : FixedArray[Float?]
) -> Unit!Error {
  let expected = get_float_option_fixedarray()
  // println("expected: \{expected}")
  // println("val: \{val}")
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_float_option() -> FixedArray[Float?] {
  get_float_option_fixedarray()
}

///|
fn get_float_option_fixedarray() -> FixedArray[Float?] {
  [Some(1), None, Some(3)]
}
