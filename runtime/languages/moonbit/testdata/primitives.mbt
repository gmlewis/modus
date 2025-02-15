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
  false
}

///|
pub fn test_bool_output_true() -> Bool {
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
  Some(false)
}

///|
pub fn test_bool_option_output_true() -> Bool? {
  Some(true)
}

///|
pub fn test_bool_option_output_none() -> Bool? {
  None
}

///|
pub fn test_byte_input_min(b : Byte) -> Unit!Error {
  assert_eq!(b, 0)
}

///|
pub fn test_byte_input_max(b : Byte) -> Unit!Error {
  assert_eq!(@byte.max_value, b)
}

///|
pub fn test_byte_output_min() -> Byte {
  0
}

///|
pub fn test_byte_output_max() -> Byte {
  @byte.max_value
}

///|
pub fn test_byte_option_input_min(b : Byte?) -> Unit!Error {
  assert_eq!(b, Some(0))
}

///|
pub fn test_byte_option_input_max(b : Byte?) -> Unit!Error {
  assert_eq!(b, Some(@byte.max_value))
}

///|
pub fn test_byte_option_input_none(b : Byte?) -> Unit!Error {
  assert_eq!(None, b)
}

///|
pub fn test_byte_option_output_min() -> Byte? {
  Some(0)
}

///|
pub fn test_byte_option_output_max() -> Byte? {
  Some(@byte.max_value)
}

///|
pub fn test_byte_option_output_none() -> Byte? {
  None
}

///|
pub fn test_char_input_min(c : Char) -> Unit!Error {
  // can't seem to pass in a char with a value of -32768
  assert_eq!(c, Char::from_int(0))
}

///|
pub fn test_char_input_max(c : Char) -> Unit!Error {
  assert_eq!(c, Char::from_int(32767))
}

///|
pub fn test_char_output_min() -> Char {
  Char::from_int(-32768)
}

///|
pub fn test_char_output_max() -> Char {
  Char::from_int(32767) // TODO: @int16.max_value))
}

///|
pub fn test_char_option_input_min(c : Char?) -> Unit!Error {
  assert_eq!(c, Some(Char::from_int(-32768))) // TODO: @int16.min_value)))
}

///|
pub fn test_char_option_input_max(c : Char?) -> Unit!Error {
  assert_eq!(c, Some(Char::from_int(32767))) // TODO: @int16.max_value)))
}

///|
pub fn test_char_option_input_none(c : Char?) -> Unit!Error {
  assert_eq!(c, None)
}

///|
pub fn test_char_option_output_min() -> Char? {
  Some(Char::from_int(-32768)) // TODO: @int16.min_value))
}

///|
pub fn test_char_option_output_max() -> Char? {
  Some(Char::from_int(32767)) // TODO: @int16.max_value))
}

///|
pub fn test_char_option_output_none() -> Char? {
  None
}

///|
pub fn test_int_input_min(n : Int) -> Unit!Error {
  assert_eq!(@int.min_value, n)
}

///|
pub fn test_int_input_max(n : Int) -> Unit!Error {
  assert_eq!(@int.max_value, n)
}

///|
pub fn test_int_output_min() -> Int {
  @int.min_value
}

///|
pub fn test_int_output_max() -> Int {
  @int.max_value
}

///|
pub fn test_int_option_input_min(n : Int?) -> Unit!Error {
  assert_eq!(n, Some(@int.min_value))
}

///|
pub fn test_int_option_input_max(n : Int?) -> Unit!Error {
  assert_eq!(n, Some(@int.max_value))
}

///|
pub fn test_int_option_input_none(n : Int?) -> Unit!Error {
  assert_eq!(n, None)
}

///|
pub fn test_int_option_output_min() -> Int? {
  Some(@int.min_value)
}

///|
pub fn test_int_option_output_max() -> Int? {
  Some(@int.max_value)
}

///|
pub fn test_int_option_output_none() -> Int? {
  None
}

///|
pub fn test_int16_input_min(n : Int16) -> Unit!Error {
  assert_eq!(n, -32768) // TODO: @int16.min_value)
}

///|
pub fn test_int16_input_max(n : Int16) -> Unit!Error {
  assert_eq!(n, 32767) // TODO: @int16.max_value)
}

///|
pub fn test_int16_output_min() -> Int16 {
  -32768 // TODO: @int16.min_value
}

///|
pub fn test_int16_output_max() -> Int16 {
  32767 // TODO: @int16.max_value
}

///|
pub fn test_int16_option_input_min(n : Int16?) -> Unit!Error {
  assert_eq!(n, Some(-32768)) // TODO: @int16.min_value)
}

///|
pub fn test_int16_option_input_max(n : Int16?) -> Unit!Error {
  assert_eq!(n, Some(32767)) // TODO: @int16.max_value)
}

///|
pub fn test_int16_option_input_none(n : Int16?) -> Unit!Error {
  assert_eq!(n, None)
}

///|
pub fn test_int16_option_output_min() -> Int16? {
  Some(-32768) // TODO: @int16.min_value
}

///|
pub fn test_int16_option_output_max() -> Int16? {
  Some(32767) // TODO: @int16.max_value
}

///|
pub fn test_int16_option_output_none() -> Int16? {
  None
}

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
  @int64.min_value
}

///|
pub fn test_int64_output_max() -> Int64 {
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
  Some(@int64.min_value)
}

///|
pub fn test_int64_option_output_max() -> Int64? {
  Some(@int64.max_value)
}

///|
pub fn test_int64_option_output_none() -> Int64? {
  None
}

///|
pub fn test_uint_input_min(n : UInt) -> Unit!Error {
  assert_eq!(n, 0)
}

///|
pub fn test_uint_input_max(n : UInt) -> Unit!Error {
  assert_eq!(n, @uint.max_value)
}

///|
pub fn test_uint_output_min() -> UInt {
  0
}

///|
pub fn test_uint_output_max() -> UInt {
  @uint.max_value
}

///|
pub fn test_uint_option_input_min(n : UInt?) -> Unit!Error {
  assert_eq!(n, Some(0))
}

///|
pub fn test_uint_option_input_max(n : UInt?) -> Unit!Error {
  assert_eq!(n, Some(@uint.max_value))
}

///|
pub fn test_uint_option_input_none(n : UInt?) -> Unit!Error {
  assert_eq!(n, None)
}

///|
pub fn test_uint_option_output_min() -> UInt? {
  Some(0)
}

///|
pub fn test_uint_option_output_max() -> UInt? {
  Some(@uint.max_value)
}

///|
pub fn test_uint_option_output_none() -> UInt? {
  None
}

///|
pub fn test_uint16_input_min(n : UInt16) -> Unit!Error {
  assert_eq!(n, 0)
}

///|
pub fn test_uint16_input_max(n : UInt16) -> Unit!Error {
  assert_eq!(n, 65535) // TODO: @uint16.max_value)
}

///|
pub fn test_uint16_output_min() -> UInt16 {
  0
}

///|
pub fn test_uint16_output_max() -> UInt16 {
  65535 // TODO: @uint16.max_value
}

///|
pub fn test_uint16_option_input_min(n : UInt16?) -> Unit!Error {
  assert_eq!(n, Some(0))
}

///|
pub fn test_uint16_option_input_max(n : UInt16?) -> Unit!Error {
  assert_eq!(n, Some(65535)) // TODO: @uint16.max_value)
}

///|
pub fn test_uint16_option_input_none(n : UInt16?) -> Unit!Error {
  assert_eq!(n, None)
}

///|
pub fn test_uint16_option_output_min() -> UInt16? {
  Some(0)
}

///|
pub fn test_uint16_option_output_max() -> UInt16? {
  Some(65535) // TODO: @uint16.max_value
}

///|
pub fn test_uint16_option_output_none() -> UInt16? {
  None
}

///|
pub fn test_uint64_input_min(n : UInt64) -> Unit!Error {
  assert_eq!(n, 0)
}

///|
pub fn test_uint64_input_max(n : UInt64) -> Unit!Error {
  assert_eq!(n, @uint64.max_value)
}

///|
pub fn test_uint64_output_min() -> UInt64 {
  0
}

///|
pub fn test_uint64_output_max() -> UInt64 {
  @uint64.max_value
}

///|
pub fn test_uint64_option_input_min(n : UInt64?) -> Unit!Error {
  assert_eq!(n, Some(0))
}

///|
pub fn test_uint64_option_input_max(n : UInt64?) -> Unit!Error {
  assert_eq!(n, Some(@uint64.max_value))
}

///|
pub fn test_uint64_option_input_none(n : UInt64?) -> Unit!Error {
  assert_eq!(n, None)
}

///|
pub fn test_uint64_option_output_min() -> UInt64? {
  Some(0)
}

///|
pub fn test_uint64_option_output_max() -> UInt64? {
  Some(@uint64.max_value)
}

///|
pub fn test_uint64_option_output_none() -> UInt64? {
  None
}

///|
pub fn test_float_input_min(n : Float) -> Unit!Error {
  assert_eq!(n, @float.min_value)
}

///|
pub fn test_float_input_max(n : Float) -> Unit!Error {
  assert_eq!(n, @float.max_value)
}

///|
pub fn test_float_output_min() -> Float {
  @float.min_value
}

///|
pub fn test_float_output_max() -> Float {
  @float.max_value
}

///|
pub fn test_float_option_input_min(n : Float?) -> Unit!Error {
  assert_eq!(n, Some(@float.min_value))
}

///|
pub fn test_float_option_input_max(n : Float?) -> Unit!Error {
  assert_eq!(n, Some(@float.max_value))
}

///|
pub fn test_float_option_input_none(n : Float?) -> Unit!Error {
  assert_eq!(n, None)
}

///|
pub fn test_float_option_output_min() -> Float? {
  Some(@float.min_value)
}

///|
pub fn test_float_option_output_max() -> Float? {
  Some(@float.max_value)
}

///|
pub fn test_float_option_output_none() -> Float? {
  None
}

///|
pub fn test_double_input_min(n : Double) -> Unit!Error {
  assert_eq!(n, @double.min_value)
}

///|
pub fn test_double_input_max(n : Double) -> Unit!Error {
  assert_eq!(n, @double.max_value)
}

///|
pub fn test_double_output_min() -> Double {
  @double.min_value
}

///|
pub fn test_double_output_max() -> Double {
  @double.max_value
}

///|
pub fn test_double_option_input_min(n : Double?) -> Unit!Error {
  assert_eq!(n, Some(@double.min_value))
}

///|
pub fn test_double_option_input_max(n : Double?) -> Unit!Error {
  assert_eq!(n, Some(@double.max_value))
}

///|
pub fn test_double_option_input_none(n : Double?) -> Unit!Error {
  assert_eq!(n, None)
}

///|
pub fn test_double_option_output_min() -> Double? {
  Some(@double.min_value)
}

///|
pub fn test_double_option_output_max() -> Double? {
  Some(@double.max_value)
}

///|
pub fn test_double_option_output_none() -> Double? {
  None
}
