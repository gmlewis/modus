// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
type! LocaltimeError String derive(Show, Eq)

///| now returns the current local time.  The time zone is determined in the following order of precedence:
/// - If the X-Time-Zone header is present in the request, the time zone is set to the value of the header.
/// - If the TZ environment variable is set on the host, the time zone is set to the value of the variable.
/// - Otherwise, the time zone is set to the host's local time zone.
pub fn now() -> @time.ZonedDateTime!Error {
  println("GML: pkg/localtime: now")

  // Try to return a time with the proper full time zone.
  let timezone = get_time_zone()
  println("GML: pkg/localtime: now: timezone: \{timezone}")
  if timezone != "" {
    let result = @host.get_time_in_zone(timezone)
    println("GML: pkg/localtime: now: result: \{result}")
    let zoneIndex = result.length() - 6
    let noZone = result.substring(end=zoneIndex)
    let zoneOffset = result.substring(start=zoneIndex)
    let zoneOffsetHours = @strconv.parse_int!(
      zoneOffset.substring(end=3),
      base=10,
    )
    let zoneOffsetMinutes = @strconv.parse_int!(
      zoneOffset.substring(start=4),
      base=10,
    )
    let zoneOffsetSeconds = zoneOffsetHours * 3600 + zoneOffsetMinutes * 60
    println(
      "GML: pkg/localtime: now: noZone: \{noZone}, zoneOffset: \{zoneOffset}, zoneOffsetHours: \{zoneOffsetHours}, zoneOffsetMinutes: \{zoneOffsetMinutes}, zoneOffsetSeconds: \{zoneOffsetSeconds}",
    )
    let plainDateTime = @time.PlainDateTime::from_string!(noZone)
    println("GML: pkg/localtime: now: plainDateTime: \{plainDateTime}")
    let zone = @time.fixed_zone!(timezone, zoneOffsetSeconds)
    return @time.ZonedDateTime::from_plain_datetime(plainDateTime, zone~)
  }

  // Otherwise, let the host get the local time with the current offset applied.
  @host.get_local_time!()
  // @time.unix!(@wallClock.now().seconds.reinterpret_as_int64()) // TODO
  // @time.unix!(@wasi.now())
}

///| now_in_zone returns the current time in the given time zone.
/// The time zone should be a valid IANA time zone identifier, such as "America/New_York".
pub fn now_in_zone(timezone : String) -> @time.ZonedDateTime!Error {
  if timezone == "" {
    raise LocaltimeError("a time zone is required")
  }

  // Try to return a time with the proper full time zone.
  // @host.get_time_location(timezone)

  // Otherwise, let the host get the local time with the current offset applied.
  let result = @host.get_time_in_zone(timezone)
  println(
    "GML: pkg/localtime: now_in_zone: timezone: \{timezone}, result: \{result}",
  )
  @time.unix!(@wasi.now())
}

///| get_time_zone returns the local time zone identifier, in IANA format.
pub fn get_time_zone() -> String {
  @wasi.environ_get("TZ")
}

///| get_location returns the @time.Zone for the given time zone.
/// `timezone` should be a valid IANA time zone identifier, such as "America/New_York".
pub fn get_location(timezone : String) -> (@time.Zone, @time.ZoneOffset)!Error {
  if timezone == "" {
    raise LocaltimeError("time zone is required")
  }
  // @host.get_time_location(timezone)
  // @time.zone!(timezone)
  let result = @host.get_time_zone_data(timezone, "tzif")
  println(
    "GML: pkg/localtime: get_location: timezone: \{timezone}, result.length(): \{result.length()}",
  )
  if result.length() == 0 {
    raise LocaltimeError("time zone '\{timezone}' not found")
  }
  let current_time_seconds = (@wasi.now() / 1_000_000_000).to_int()
  load_location_from_tzdata!(timezone, result, current_time_seconds)
}

///| is_valid_time_zone returns true if the given time zone is valid.
pub fn is_valid_time_zone(timezone : String) -> Bool {
  let t = @host.get_time_in_zone(timezone)
  t != ""
}
