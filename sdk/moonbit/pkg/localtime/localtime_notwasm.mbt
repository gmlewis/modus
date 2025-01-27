// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| now returns the current local time.  The time zone is determined in the following order of precedence:
/// - If the X-Time-Zone header is present in the request, the time zone is set to the value of the header.
/// - If the TZ environment variable is set on the host, the time zone is set to the value of the variable.
/// - Otherwise, the time zone is set to the host's local time zone.
pub fn now() -> @time.ZonedDateTime!Error {
  @time.unix!(@wallClock.now().seconds.reinterpret_as_int64())
}

///| now_in_zone returns the current time in the given time zone.
/// The time zone should be a valid IANA time zone identifier, such as "America/New_York".
pub fn now_in_zone(_timezone : String) -> @time.ZonedDateTime!Error {
  @time.unix!(@wallClock.now().seconds.reinterpret_as_int64())
}

///| get_time_zone returns the local time zone identifier, in IANA format.
pub fn get_time_zone() -> String {
  @sys.get_env_vars()["TZ"].or("")
}

// get_location returns the time.Location for the given time zone.
// The time zone should be a valid IANA time zone identifier, such as "America/New_York".
///|
pub fn get_location(_timezone : String) -> @time.Zone!Error {
  fail!("Not implemented")
  @time.utc_zone
}

// is_valid_time_zone returns true if the given time zone is valid.
///|
pub fn is_valid_time_zone(_timezone : String) -> Bool {
  true
}
