// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
pub fn hello_option_empty_string(some~ : Bool = true) -> String? {
  println("hello_option_empty_string: Got 'some' param with value: \{some}")
  if some {
    Some("")
  } else {
    None
  }
}

///|
pub fn hello_option_none(some~ : Bool = false) -> String? {
  println("hello_option_none: Got 'some' param with value: \{some}")
  if some {
    Some("")
  } else {
    None
  }
}

///|
pub fn hello_option_some_string(some~ : Bool = true) -> String? {
  println("hello_option_some_string: Got 'some' param with value: \{some}")
  if some {
    Some("Hello, World!")
  } else {
    None
  }
}

///|
pub fn hello_world_with_arg(name : String) -> String {
  println(
    "hello_world_with_arg: Got 'name' param with length: \{name.length()}: '\{name}'",
  )
  "Hello, \{name}!"
}

///|
pub fn hello_world_with_optional_arg(
  name~ : String = "世界 🌍 from MoonBit"
) -> String {
  println(
    "hello_world_with_optional_arg: Got 'name' param with length: \{name.length()}: '\{name}'",
  )
  "Hello, \{name}!"
}

///|
pub fn hello_world() -> String {
  println("hello_world: No params")
  "Hello, World!"
}

///|
pub fn hello_array_of_ints(n~ : Int = 3) -> Array[Int] {
  if n < 0 {
    println("hello_array_of_ints(\{n}): returning: []")
    return []
  }
  let arr = Array::new(capacity=n)
  for i in 1..=n {
    arr.push(i)
  }
  println("hello_array_of_ints(\{n}): returning array of ints: \{arr}")
  arr
}

///|
pub fn hello_array_of_ints_option(n~ : Int = 3) -> Array[Int]? {
  if n < 0 {
    println("hello_array_of_ints_option(\{n}): returning: None")
    return None
  }
  let arr = Array::new(capacity=n)
  for i in 1..=n {
    arr.push(i)
  }
  println("hello_array_of_ints_option(\{n}): returning array of ints: \{arr}")
  Some(arr)
}

///|
pub fn hello_array_of_strings(n~ : Int = 3) -> Array[String] {
  if n < 0 {
    println("hello_array_of_strings(\{n}): returning: []")
    return []
  }
  let arr = Array::new(capacity=n)
  for i in 1..=n {
    arr.push("\{i}")
  }
  println("hello_array_of_strings(\{n}): returning array of strings: \{arr}")
  arr
}

///|
pub fn hello_array_of_strings_option(n~ : Int = 3) -> Array[String]? {
  if n < 0 {
    println("hello_array_of_strings_option(\{n}): returning: None")
    return None
  }
  let arr = Array::new(capacity=n)
  for i in 1..=n {
    arr.push("\{i}")
  }
  println(
    "hello_array_of_strings_option(\{n}): returning array of strings: \{arr}",
  )
  Some(arr)
}

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
  let result = Char::from_int(0) // Char::min_value
  println("hello_primitive_char_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_char_max() -> Char {
  let result = Char::from_int(65535) // Char::max_value
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
  let result : Int16 = -32768 // @int16.min_value
  println("hello_primitive_int16_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_int16_max() -> Int16 {
  let result : Int16 = 32767 // @int16.max_value
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
  let result : UInt16 = 0 // @uint16.min_value
  println("hello_primitive_uint16_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_uint16_max() -> UInt16 {
  let result : UInt16 = 65535 // @uint16.max_value
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

///|
pub fn hello_maps_n_items(n~ : Int = 3) -> Map[String, String] {
  if n < 0 {
    println("hello_maps_n_items(\{n}): returning: {}")
    return {}
  }
  let map = {}
  for i in 1..=n {
    let key = "key\{i}"
    let value = "value\{i}"
    map[key] = value
  }
  println("hello_maps_n_items(\{n}): returning: \{map}")
  map
}

///|
pub fn hello_maps_n_items_option(n~ : Int = 3) -> Map[String, String]? {
  if n < 0 {
    println("hello_maps_n_items_option(\{n}): returning: None")
    return None
  }
  let map = {}
  for i in 1..=n {
    let key = "key\{i}"
    let value = "value\{i}"
    map[key] = value
  }
  println("hello_maps_n_items_option(\{n}): returning: Some(\{map})")
  Some(map)
}
