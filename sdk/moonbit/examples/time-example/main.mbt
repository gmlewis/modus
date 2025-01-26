// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

// NOTES:
//
// - To get the local time, you should use the `localtime` package provided by Modus - as shown in the examples below.
//   If you actually want UTC, you should prefer `@time.now()`.
//
// - If you return a `@time.ZonedDateTime` object to Modus, it will be converted to UTC - regardless of the time zone.
//   Thus, if you need to return a time object to Modus, you should return a formatted string instead - Preferably in RFC3339 (ISO 8601) format.

///|
// Returns the current time in UTC.
pub fn get_utc_time() -> @time.ZonedDateTime {
  @time.now()
}

///|
// Returns the current local time.
pub fn get_local_time() -> String!Error {
  let now = @localtime.now!()
  now.format(@localtime.RFC3339)
}

///|
// Returns the current time in a specified time zone.
pub fn get_time_in_zone(tz : String) -> String!Error {
  let now = @localtime.now_in_zone!(tz)
  now.format(@localtime.RFC3339)
}

///|
// Returns the local time zone identifier.
pub fn get_local_time_zone() -> String {
  @localtime.get_time_zone()
}

///|
pub struct TimeZoneInfo {
  standard_name : String
  standard_offset : String
  daylight_name : String
  daylight_offset : String
}

// Returns some basic information about the time zone specified.
pub fn get_time_zone_info(tz : String) -> TimeZoneInfo!Error {
	let loc = @localtime.get_location!(tz)

	let (jan_name, jan_offset) = @time.date(2024, 1, 1, 0, 0, 0, 0, loc).zone()
	let (jul_name, jul_offset) = @time.date(2024, 7, 1, 0, 0, 0, 0, loc).zone()

  let (std_name, std_offset, dlt_name, dlt_offset) = if jan_offset <= jul_offset {
    (jan_name, jan_offset, jul_name, jul_offset)
  } else {
    (jul_name, jul_offset, jan_name, jan_offset)
  }

	TimeZoneInfo{
		standard_name:   std_name,
		standard_offset: format_offset(std_offset),
		daylight_name:   dlt_name,
		daylight_offset: format_offset(dlt_offset),
	}
}

///|
// Formats the offset in hours and minutes (used by get_time_zone_info).
fn format_offset(offset : Int) -> String {
  let mut sign = "+"
  if offset < 0 {
    sign = "-"
    offset = -offset
  }
  offset /= 60
  let hours = offset / 60
  let minutes = offset % 60
  "\{sign}\{hours}:\{minutes}"
}
