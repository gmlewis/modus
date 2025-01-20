// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
pub fn hello_array_of_ints(n~ : Int = 3) -> Array[Int] {
  let arr = Array::new(capacity=n)
  for i in 0..<n {
    arr.push(i)
  }
  arr
}

///|
pub fn hello_array_of_ints_option(n~ : Int = 3) -> Array[Int]? {
  if n < 0 {
    return None
  }
  let arr = Array::new(capacity=n)
  for i in 0..<n {
    arr.push(i)
  }
  Some(arr)
}
