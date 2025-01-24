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
    println("returning: []")
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
    println("returning: None")
    return None
  }
  let arr = Array::new(capacity=n)
  for i in 1..=n {
    arr.push(i)
  }
  println("returning array of ints: \{arr}")
  Some(arr)
}
