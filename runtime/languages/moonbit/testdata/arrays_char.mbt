// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_array_input_char_empty(val : Array[Char]) -> Unit!Error {
  let expected = []
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_char_0() -> Array[Char] {
  []
}

///|
pub fn test_array_output_char_1() -> Array[Char] {
  ['1']
}

///|
pub fn test_array_output_char_2() -> Array[Char] {
  ['1', '2']
}

///|
pub fn test_array_output_char_3() -> Array[Char] {
  ['1', '2', '3']
}

///|
pub fn test_array_output_char_4() -> Array[Char] {
  ['1', '2', '3', '4']
}

///|
pub fn test_array_output_char_option_0() -> Array[Char?] {
  []
}

///|
pub fn test_array_output_char_option_1_none() -> Array[Char?] {
  [None]
}

///|
pub fn test_array_output_char_option_1_some() -> Array[Char?] {
  [Some('1')]
}

///|
pub fn test_array_output_char_option_2() -> Array[Char?] {
  [Some('1'), Some('2')]
}

///|
pub fn test_array_output_char_option_3() -> Array[Char?] {
  [None, None, None]
}

///|
pub fn test_array_output_char_option_4() -> Array[Char?] {
  [None, Some('2'), Some(Char::from_int(0)), Some('4')]
}

///|
pub fn test_array_input_char_option(val : Array[Char?]) -> Unit!Error {
  let expected = get_char_option_array()
  // println("expected: \{expected}")
  // println("val: \{val}")
  assert_eq!(expected, val)
}

///|
pub fn test_array_output_char_option() -> Array[Char?] {
  get_char_option_array()
}

///|
fn get_char_option_array() -> Array[Char?] {
  let a = Some('1')
  let b = None
  let c = Some('3')
  [a, b, c]
}
