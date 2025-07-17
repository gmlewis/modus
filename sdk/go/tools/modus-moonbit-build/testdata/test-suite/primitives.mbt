// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
pub fn hello_primitive_bool_min() -> Bool {
  let result = false
  println("hello_primitive_bool_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_bool_max() -> Bool {
  let result = true
  println("hello_primitive_bool_max: returning \{result}")
  result
}

///|
pub fn hello_primitive_byte_min() -> Byte {
  let result = @byte.min_value
  println("hello_primitive_byte_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_byte_max() -> Byte {
  let result = @byte.max_value
  println("hello_primitive_byte_max: returning \{result}")
  result
}

///|
pub fn hello_primitive_char_min() -> Char {
  let result = Int::unsafe_to_char(0) // Char::min_value
  println("hello_primitive_char_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_char_max() -> Char {
  let result = Int::unsafe_to_char(65535) // Char::max_value
  println("hello_primitive_char_max: returning \{result}")
  result
}

///|
pub fn hello_primitive_double_min() -> Double {
  let result = @double.min_value
  println("hello_primitive_double_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_double_max() -> Double {
  let result = @double.max_value
  println("hello_primitive_double_max: returning \{result}")
  result
}

///|
pub fn hello_primitive_float_min() -> Float {
  let result = @float.min_value
  println("hello_primitive_float_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_float_max() -> Float {
  let result = @float.max_value
  println("hello_primitive_float_max: returning \{result}")
  result
}

///|
pub fn hello_primitive_int_min() -> Int {
  let result = @int.min_value
  println("hello_primitive_int_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_int_max() -> Int {
  let result = @int.max_value
  println("hello_primitive_int_max: returning \{result}")
  result
}

///|
pub fn hello_primitive_int16_min() -> Int16 {
  let result : Int16 = -32768 // TODO: @int16.min_value
  println("hello_primitive_int16_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_int16_max() -> Int16 {
  let result : Int16 = 32767 // TODO: @int16.max_value
  println("hello_primitive_int16_max: returning \{result}")
  result
}

///|
pub fn hello_primitive_int64_min() -> Int64 {
  let result = @int64.min_value
  println("hello_primitive_int64_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_int64_max() -> Int64 {
  let result = @int64.max_value
  println("hello_primitive_int64_max: returning \{result}")
  result
}

///|
pub fn hello_primitive_uint_min() -> UInt {
  let result = @uint.min_value
  println("hello_primitive_uint_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_uint_max() -> UInt {
  let result = @uint.max_value
  println("hello_primitive_uint_max: returning \{result}")
  result
}

///|
pub fn hello_primitive_uint16_min() -> UInt16 {
  let result : UInt16 = 0 // TODO: @uint16.min_value
  println("hello_primitive_uint16_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_uint16_max() -> UInt16 {
  let result : UInt16 = 65535 // TODO: @uint16.max_value
  println("hello_primitive_uint16_max: returning \{result}")
  result
}

///|
pub fn hello_primitive_uint64_min() -> UInt64 {
  let result = @uint64.min_value
  println("hello_primitive_uint64_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_uint64_max() -> UInt64 {
  let result = @uint64.max_value
  println("hello_primitive_uint64_max: returning \{result}")
  result
}
