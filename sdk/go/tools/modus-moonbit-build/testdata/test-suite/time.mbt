// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

// NOTES:
//
// - To get the local time, you should use the `localtime` package provided by Modus - as shown in the examples below.
//   If you actually want UTC, you should prefer `@time.unix!(@wasi.now()/1_000_000_000)`.
//
// - If you return a `@time.ZonedDateTime` object to Modus, it will be converted to UTC - regardless of the time zone.
//   Thus, if you need to return a time object to Modus, you should return a formatted string instead - Preferably in RFC3339 (ISO 8601) format.

///| Returns the current time in UTC.
pub fn get_utc_time() -> @time.ZonedDateTime raise Error {
  let now_nanos = @wasi.now()
  let seconds = now_nanos / 1_000_000_000
  let nanosecond = (now_nanos % 1_000_000_000).to_int()
  let result = @time.unix(seconds, nanosecond~)
  println(
    "GML: main.mbt: get_utc_time: now_nanos=\{now_nanos}, result=\{result}",
  )
  result
}

///| Returns the current local time.
pub fn get_local_time() -> String raise Error {
  let now = @localtime.now()
  println("GML: main.mbt: get_local_time: now=\{now}")
  now.to_string()
}

///| Returns the current time in a specified time zone.
pub fn get_time_in_zone(tz : String) -> String raise Error {
  let now = @localtime.now(timezone=tz)
  println("GML: main.mbt: get_time_in_zone: tz: \{tz}, now: \{now}")
  now.to_string()
}

///| Returns the local time zone identifier.
pub fn get_local_time_zone_id() -> String {
  @localtime.get_local_time_zone_id()
}

///|
pub struct TimeZoneInfo {
  standard_name : String
  standard_offset : String
  daylight_name : String
  daylight_offset : String
}

///| Returns some basic information about the time zone specified.
pub fn get_time_zone_info(tz : String) -> TimeZoneInfo raise Error {
  let zone = @localtime.get_time_zone(timezone=tz)
  let jan1 = @time.date_time(2024, 1, 1, zone~)
  let jul1 = @time.date_time(2024, 7, 1, zone~)
  let jan_name = jan1.offset().abbreviation()
  let jul_name = jul1.offset().abbreviation()
  let jan_offset = jan1.offset()
  let jul_offset = jul1.offset()
  let (std_name, std_offset, dlt_name, dlt_offset) = if jul_offset <= jan_offset {
    (jan_name, jan_offset, jul_name, jul_offset)
  } else {
    (jul_name, jul_offset, jan_name, jan_offset)
  }
  {
    standard_name: std_name,
    standard_offset: std_offset.to_string(),
    daylight_name: dlt_name,
    daylight_offset: dlt_offset.to_string(),
  }
}
