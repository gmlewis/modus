// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
let test_time : Ref[@time.ZonedDateTime?] = { val: None }

///|
let test_duration : Ref[@time.Duration] = { val: @time.Duration::zero() }

///|
fn init {
  test_time.val = try {
    Some(
      @time.date_time!(
        2024,
        12,
        31,
        hour=23,
        minute=59,
        second=59,
        nanosecond=999_999_999,
      ),
    )
  } catch {
    _ => abort("Failed to create test time")
  }
  test_duration.val = try {
    @time.Duration::of!(seconds=5)
  } catch {
    _ => abort("Failed to create test duration")
  }
}

///|
pub fn test_time_input(t : @time.ZonedDateTime) -> Unit!Error {
  guard let Some(test_time) = test_time.val
  assert_eq!(t.to_unix_second(), test_time.to_unix_second())
}

///|
pub fn test_time_option_input(t : @time.ZonedDateTime?) -> Unit!Error {
  guard let Some(test_time) = test_time.val
  guard let Some(t) = t
  assert_eq!(t.to_unix_second(), test_time.to_unix_second())
}

///|
type! TestError String derive(Show, Eq)

///|
pub fn test_time_option_input_nil(t : @time.ZonedDateTime?) -> Unit!Error {
  guard let None = t else { _ => raise TestError("expected None") }

}

///|
pub fn test_time_output() -> @time.ZonedDateTime {
  guard let Some(test_time) = test_time.val
  test_time
}

///|
pub fn test_time_option_output() -> @time.ZonedDateTime? {
  test_time.val
}

///|
pub fn test_time_option_output_nil() -> @time.ZonedDateTime? {
  None
}

///|
pub fn test_duration_input(d : @time.Duration) -> Unit!Error {
  assert_eq!(d, test_duration.val)
}

///|
pub fn test_duration_option_input(d : @time.Duration?) -> Unit!Error {
  assert_eq!(d, Some(test_duration.val))
}

///|
pub fn test_duration_option_input_nil(d : @time.Duration?) -> Unit!Error {
  assert_eq!(d, None)
}

///|
pub fn test_duration_output() -> @time.Duration {
  test_duration.val
}

///|
pub fn test_duration_option_output() -> @time.Duration? {
  Some(test_duration.val)
}

///|
pub fn test_duration_option_output_nil() -> @time.Duration? {
  None
}
