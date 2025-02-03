// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
type! LocaltimeError String derive(Show, Eq)

///| now returns the current time in the given time zone (or the local time zone if none is provided).
///
/// The time zone is determined in the following order of precedence:
/// - If the X-Time-Zone header is present in the request, the time zone is set to the value of the header.
/// - If the TZ environment variable is set on the host, the time zone is set to the value of the variable.
/// - Otherwise, the time zone is set to the host's local time zone.
pub fn now(timezone~ : String = "") -> @time.ZonedDateTime!Error {
  let timezone = if timezone == "" {
    get_local_time_zone_id()
  } else {
    timezone
  }
  println("pkg/localtime: now: timezone: \{timezone}")

  //
  if timezone == "" {
    // If still empty, let the host get the local time.
    let result = @host.get_local_time!()
    println("pkg/localtime: now: @host.get_local_time!(): result: \{result}")
    return result
  }
  let now = @wasi.now() / 1_000_000_000
  let zone = get_time_zone!(timezone~)
  println("pkg/localtime: now: now: \{now}, zone: \{zone}")
  let result = @time.unix!(now, zone~)
  println(
    "pkg/localtime: now: timezone: \{timezone}, now: \{now}, zone: \{zone}, result: \{result}",
  )
  result
}

///| get_local_time_zone_id returns the local time zone identifier, in IANA format
/// (e.g., "America/Los_Angeles").
pub fn get_local_time_zone_id() -> String {
  let result = @wasi.environ_get("TZ")
  println("pkg/localtime: get_local_time_zone_id: result: \{result}")
  result
}

///| get_time_zone returns the @time.Zone for the given time zone (or the local
/// time zone if none is provided).
/// `timezone` should be a valid IANA time zone identifier, such as "America/New_York".
pub fn get_time_zone(timezone~ : String = "") -> @time.Zone!Error {
  let timezone = if timezone == "" {
    get_local_time_zone_id()
  } else {
    timezone
  }
  if timezone == "" {
    raise LocaltimeError("time zone is required")
  }
  let result = @host.get_time_zone_data(timezone, "tzif")
  if result.length() == 0 {
    raise LocaltimeError("time zone '\{timezone}' not found")
  }
  let result = @time.Zone::from_tzif2!(timezone, result)
  println(
    "pkg/localtime: get_time_zone: timezone: \{timezone}, result: \{result}",
  )
  result
}

///| is_valid_time_zone returns true if the given time zone is valid.
pub fn is_valid_time_zone(timezone : String) -> Bool {
  let t = @host.get_time_in_zone(timezone)
  let result = t != ""
  println(
    "pkg/localtime: is_valid_time_zone: timezone: \{timezone}, result: \{result}",
  )
  result
}
