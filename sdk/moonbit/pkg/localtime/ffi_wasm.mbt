///|modus:import modus_system getTimeInZone(tz : String) -> String
fn wasmImportGetTimeInZone(tz : Int) -> Int = "modus_system" "getTimeInZone"

///|modus:import modus_system getTimeZoneData(tz : String, format : String) -> Array[Byte]
fn wasmImportGetTimeZoneData(tz : Int, format : Int) -> Int = "modus_system" "getTimeZoneData"

///|
fn host_get_time_in_zone_modus(timezone : String) -> String {
  let timezone_ptr = @ffi.str2ptr(timezone) - 8
  let ptr = wasmImportGetTimeInZone(timezone_ptr)
  @ffi.ptr2str(ptr + 8)
}

///|
fn host_get_time_zone_data(
  timezone : String,
  format : String
) -> FixedArray[Byte] {
  let timezone_ptr = @ffi.str2ptr(timezone) - 8
  let format_ptr = @ffi.str2ptr(format) - 8

  // Call the host function with the pointers to the strings
  let ptr = wasmImportGetTimeZoneData(timezone_ptr, format_ptr)
  if ptr == 0 {
    return []
  }

  // The host function actually returns `Array[Byte]`, not `FixedArray[Byte]`.
  let result : Array[Byte] = cast_from_ptr(ptr)
  FixedArray::from_array(result)
}

///|
fn[A] cast_from_ptr(a : Int) -> A = "%identity"

///|
fn host_get_local_time_modus() -> String {
  let ptr = wasmImportGetTimeInZone(0)
  @ffi.ptr2str(ptr + 8)
}
