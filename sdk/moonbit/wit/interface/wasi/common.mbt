///|
pub fn now() -> Int64 {
  match clock_time_get(0, 0) {
    Ok(time) => time.reinterpret_as_int64()
    Err(_) => 0
  }
}
