// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

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
pub fn test_array_output_int_option() -> Array[Int?] {
  get_int_option_array()
}

///|
fn get_int_option_array() -> Array[Int?] {
  let a = Some(11)
  let b = None
  let c = Some(33)
  [a, b, c]
}

///|
pub fn test_array_output_string_option() -> Array[String?] {
  get_string_option_array()
}

///|
fn get_string_option_array() -> Array[String?] {
  let a = Some("abc")
  let b = None
  let c = Some("ghi")
  [a, b, c]
}

///|
pub fn test_array_input_string(val : Array[String]) -> Unit!Error {
  let expected = ["abc", "def", "ghi"]
  assert_eq!(expected, val)
}
