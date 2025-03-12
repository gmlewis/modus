// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
priv type! LocaltimeError String derive(Show, Eq)

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

  //
  if timezone == "" {
    // If still empty, let the host get the local time.
    let s = @host.get_local_time_modus()
    let pdt = @time.PlainDateTime::from_string!(s)
    return @time.ZonedDateTime::from_plain_datetime(pdt)
  }
  let now = @wasi.now() / 1_000_000_000
  let zone = get_time_zone!(timezone~)
  @time.unix!(now, zone~)
}

///| get_local_time_zone_id returns the local time zone identifier, in IANA format
/// (e.g., "America/Los_Angeles").
pub fn get_local_time_zone_id() -> String {
  @wasi.environ_get("TZ")
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
  @time.Zone::from_tzif2!(timezone, result)
}

///| is_valid_time_zone returns true if the given time zone is valid.
pub fn is_valid_time_zone(timezone : String) -> Bool {
  @host.get_time_in_zone_modus(timezone) != ""
}

///|
pub fn get_local_time_modus() -> String {
  @host.get_local_time_modus()
}

///|
pub fn get_time_in_zone_modus(timezone : String) -> String {
  @host.get_time_in_zone_modus(timezone)
}
