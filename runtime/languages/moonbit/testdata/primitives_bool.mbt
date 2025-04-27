// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_bool_input_false(b : Bool) -> Unit!Error {
  assert_eq!(false, b)
}

///|
pub fn test_bool_input_true(b : Bool) -> Unit!Error {
  assert_eq!(true, b)
}

///|
pub fn test_bool_output_false() -> Bool {
  // primitiveHandler.Decode(vals: [0])
  false
}

///|
pub fn test_bool_output_true() -> Bool {
  // primitiveHandler.Decode(vals: [1])
  true
}

///|
pub fn test_bool_option_input_false(b : Bool?) -> Unit!Error {
  assert_eq!(Some(false), b)
}

///|
pub fn test_bool_option_input_true(b : Bool?) -> Unit!Error {
  assert_eq!(Some(true), b)
}

///|
pub fn test_bool_option_input_none(b : Bool?) -> Unit!Error {
  assert_eq!(None, b)
}

///|
pub fn test_bool_option_output_false() -> Bool? {
  // primitiveHandler.Decode(vals: [0])
  Some(false)
}

///|
pub fn test_bool_option_output_true() -> Bool? {
  // primitiveHandler.Decode(vals: [1])
  Some(true)
}

///|
pub fn test_bool_option_output_none() -> Bool? {
  // primitiveHandler.Decode(vals: [4294967295])
  None
}
