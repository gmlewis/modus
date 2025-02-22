// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_input_bool(val : FixedArray[Bool]) -> Unit!Error {
  let expected : FixedArray[Bool] = [false, true, true, false]
  println("val: \{dump_mem_block(cast_to_ptr(val))}")
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_bool_0() -> FixedArray[Bool] {
  []
}

///|
pub fn test_fixedarray_output_bool_1() -> FixedArray[Bool] {
  [true]
}

///|
pub fn test_fixedarray_output_bool_2() -> FixedArray[Bool] {
  [false, true]
}

///|
pub fn test_fixedarray_output_bool_3() -> FixedArray[Bool] {
  [true, true, true]
}

///|
pub fn test_fixedarray_output_bool_4() -> FixedArray[Bool] {
  [false, false, false, false]
}

///|
pub fn test_fixedarray_output_bool_option_0() -> FixedArray[Bool?] {
  []
}

///|
pub fn test_fixedarray_output_bool_option_1_none() -> FixedArray[Bool?] {
  [None]
}

///|
pub fn test_fixedarray_output_bool_option_1_false() -> FixedArray[Bool?] {
  [Some(false)]
}

///|
pub fn test_fixedarray_output_bool_option_1_true() -> FixedArray[Bool?] {
  [Some(true)]
}

///|
pub fn test_fixedarray_output_bool_option_2() -> FixedArray[Bool?] {
  [Some(true), None]
}

///|
pub fn test_fixedarray_output_bool_option_3() -> FixedArray[Bool?] {
  [None, Some(true), Some(true)]
}

///|
pub fn test_fixedarray_output_bool_option_4() -> FixedArray[Bool?] {
  [Some(false), Some(true), Some(false), None]
}
