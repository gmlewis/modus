// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

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
