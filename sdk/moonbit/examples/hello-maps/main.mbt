///|
pub fn hello_maps_0_items() -> Map[String, String] {
  {}
}

///|
pub fn hello_maps_1_item() -> Map[String, String] {
  { "a": "apple" }
}

///|
pub fn hello_maps_2_items() -> Map[String, String] {
  { "b": "banana", "c": "cherry" }
}

///|
pub fn hello_maps_3_items() -> Map[String, String] {
  { "d": "date", "e": "elderberry", "f": "fig" }
}

///|
pub fn hello_maps_4_items() -> Map[String, String] {
  { "g": "grape", "h": "honeydew", "i": "imbe", "j": "jackfruit" }
}

///|
pub fn hello_maps_5_items() -> Map[String, String] {
  { "k": "kiwi", "l": "lemon", "m": "mango", "n": "nectarine", "o": "orange" }
}

///|
pub fn hello_maps_6_items() -> Map[String, String] {
  {
    "p": "papaya",
    "q": "quince",
    "r": "raspberry",
    "s": "strawberry",
    "t": "tangerine",
    "u": "ugli",
  }
}

///|
pub fn hello_maps_n_items(n~ : Int = 3) -> Map[String, String] {
  let map = {}
  for i in 1..=n {
    let key = "key\{i}"
    let value = "value\{i}"
    map[key] = value
  }
  map
}

// ///|
// pub fn hack() -> Array[String] {
//   ["d", "e", "f"]
// }
