// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

// NOTES:
//
// - To get the local time, you should use the `localtime` package provided by Modus - as shown in the examples below.
//   If you actually want UTC, you should prefer `@time.unix!(@wasi.now())`.
//
// - If you return a `@time.ZonedDateTime` object to Modus, it will be converted to UTC - regardless of the time zone.
//   Thus, if you need to return a time object to Modus, you should return a formatted string instead - Preferably in RFC3339 (ISO 8601) format.

///|
/// Returns the current time in UTC.
pub fn get_utc_time() -> @time.ZonedDateTime!Error {
  println("GML: main.mbt: get_utc_time")
  @time.unix!(@wasi.now())
}

///|
/// Returns the current local time.
pub fn get_local_time() -> String!Error {
  println("GML: main.mbt: get_local_time")
  let now = @localtime.now!()
  // now.format(@localtime.RFC3339)
  now.to_string()
}

///|
/// Returns the current time in a specified time zone.
pub fn get_time_in_zone(tz : String) -> String!Error {
  println("GML: main.mbt: get_time_in_zone: tz: \{tz}")
  let now = @localtime.now_in_zone!(tz)
  // now.format(@localtime.RFC3339)
  now.to_string()
}

///|
/// Returns the local time zone identifier.
pub fn get_local_time_zone() -> String {
  println("GML: main.mbt: get_local_time_zone")
  @localtime.get_time_zone()
}

///|
pub struct TimeZoneInfo {
  standard_name : String
  standard_offset : String
  daylight_name : String
  daylight_offset : String
}

///| Returns some basic information about the time zone specified.
pub fn get_time_zone_info(tz : String) -> TimeZoneInfo!Error {
  let loc = @localtime.get_location!(tz)
  let jan1 = @time.date_time!(2024, 1, 1, zone=loc)
  let jul1 = @time.date_time!(2024, 7, 1, zone=loc)
  let jan_name = jan1.zone().to_string()
  let jul_name = jul1.zone().to_string()
  let jan_offset = jan1.offset()
  let jul_offset = jul1.offset()
  let (std_name, std_offset, dlt_name, dlt_offset) = if jan_offset <= jul_offset {
    (jan_name, jan_offset, jul_name, jul_offset)
  } else {
    (jul_name, jul_offset, jan_name, jan_offset)
  }
  {
    standard_name: std_name,
    // standard_offset: format_offset(std_offset),
    standard_offset: std_offset.to_string(),
    daylight_name: dlt_name,
    // daylight_offset: format_offset(dlt_offset),
    daylight_offset: dlt_offset.to_string(),
  }
}

///| Formats the offset in hours and minutes (used by get_time_zone_info).
fn format_offset(offset : Int) -> String {
  let mut sign = "+"
  let mut offset = offset
  if offset < 0 {
    sign = "-"
    offset = -offset
  }
  offset /= 60
  let hours = offset / 60
  let minutes = offset % 60
  "\{sign}\{hours}:\{minutes}"
}
