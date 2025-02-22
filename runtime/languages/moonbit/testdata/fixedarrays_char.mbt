// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_fixedarray_input_char_empty(val : FixedArray[Char]) -> Unit!Error {
  let expected : FixedArray[Char] = []
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_char_0() -> FixedArray[Char] {
  []
}

///|
pub fn test_fixedarray_output_char_1() -> FixedArray[Char] {
  ['1']
}

///|
pub fn test_fixedarray_output_char_2() -> FixedArray[Char] {
  ['1', '2']
}

///|
pub fn test_fixedarray_output_char_3() -> FixedArray[Char] {
  ['1', '2', '3']
}

///|
pub fn test_fixedarray_output_char_4() -> FixedArray[Char] {
  ['1', '2', '3', '4']
}

///|
pub fn test_fixedarray_output_char_option_0() -> FixedArray[Char?] {
  []
}

///|
pub fn test_fixedarray_output_char_option_1_none() -> FixedArray[Char?] {
  [None]
}

///|
pub fn test_fixedarray_output_char_option_1_some() -> FixedArray[Char?] {
  [Some('1')]
}

///|
pub fn test_fixedarray_output_char_option_2() -> FixedArray[Char?] {
  [Some('1'), Some('2')]
}

///|
pub fn test_fixedarray_output_char_option_3() -> FixedArray[Char?] {
  [None, None, None]
}

///|
pub fn test_fixedarray_output_char_option_4() -> FixedArray[Char?] {
  [None, Some('2'), Some(Char::from_int(0)), Some('4')]
}

///|
pub fn test_fixedarray_input_char_option(val : FixedArray[Char?]) -> Unit!Error {
  let expected = get_char_option_fixedarray()
  assert_eq!(expected, val)
}

///|
pub fn test_fixedarray_output_char_option() -> FixedArray[Char?] {
  get_char_option_fixedarray()
}

///|
fn get_char_option_fixedarray() -> FixedArray[Char?] {
  let a = Some('1')
  let b = None
  let c = Some('3')
  [a, b, c]
}
