///|
pub fn hello_maps_n_items(n~ : Int = 3) -> Map[String, Int] {
  let map = {}
  for i in 1..=n {
    let key = "key\{i}"
    map[key] = i
  }
  map
}
