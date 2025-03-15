// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) struct Location {
  longitude : Double
  latitude : Double
} derive(Show, Eq, FromJson, ToJson)

///|
pub impl Show for Location with to_string(self) {
  "(\{self.longitude},\{self.latitude})"
}

///|
pub fn Location::new(longitude : Double, latitude : Double) -> Location {
  { longitude, latitude }
}

///|
pub fn Location::parse(s : String) -> Location!Error {
  let s = s.trim_space()
  let parts = s.split(",").collect()
  if parts.length() != 2 ||
    not(parts[0].starts_with("(")) ||
    not(parts[1].ends_with(")")) {
    raise PostgreSqlError("invalid location")
  }
  let s0 = parts[0].trim_start("(").trim_space()
  let s1 = parts[1].trim_end(")").trim_space()
  let longitude = @strconv.parse_decimal!(s0).to_double!()
  let latitude = @strconv.parse_decimal!(s1).to_double!()
  { longitude, latitude }
}
