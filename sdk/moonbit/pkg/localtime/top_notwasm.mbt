///|
pub fn host_get_time_in_zone_modus(_timezone : String) -> String {
  abort("not wasm")
}

///|
pub fn host_get_time_zone_data(
  _timezone : String,
  _format : String
) -> FixedArray[Byte] {
  abort("not wasm")
}

///|
pub fn host_get_local_time_modus() -> String {
  abort("not wasm")
}
