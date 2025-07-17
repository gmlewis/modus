// TODO: Why is the runtime failing on these functions?

///|
pub fn hello_primitive_time_duration_min() -> @time.Duration {
  let result = @time.Duration::zero()
  println("hello_primitive_time_duration_min: returning \{result}")
  result
}

///|
pub fn hello_primitive_time_duration_max() -> @time.Duration {
  let result = @time.Duration::of(seconds=@int64.max_value) catch {
    _ => @time.Duration::zero()
  }
  println("hello_primitive_time_duration_max: returning \{result}")
  result
}
