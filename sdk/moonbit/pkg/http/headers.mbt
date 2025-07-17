// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) struct Header {
  name : String
  values : Array[String]
} derive(Show, Eq, FromJson, ToJson)

///| Headers type representation
pub(all) struct Headers {
  data : Map[String, Header?]
} derive(Show, Eq, FromJson, ToJson)

///|
pub fn Headers::new() -> Headers {
  { data: {} }
}

///|
pub fn Headers::clone(self : Headers) -> Headers {
  let headers = Headers::new()
  for kv in self.data {
    let key = kv.0
    let header = kv.1.unwrap()
    headers.data[key] = Some(header)
  }
  headers
}

///|
pub fn Headers::from_array(value : Array[Array[String]]) -> Headers {
  let headers = Headers::new()
  for entry in value {
    for i in 1..<entry.length() {
      headers.append(entry[0], entry[i])
    }
  }
  headers
}

///|
pub fn Headers::from_tuples(value : Array[(String, String)]) -> Headers {
  let headers = Headers::new()
  for tup in value {
    headers.append(tup.0, tup.1)
  }
  headers
}

///|
pub fn Headers::from_map(value : Map[String, String]) -> Headers {
  let headers = Headers::new()
  for name, value in value {
    headers.append(name, value)
  }
  headers
}

///|
pub fn Headers::from_map_array(value : Map[String, Array[String]]) -> Headers {
  let headers = Headers::new()
  for name, values in value {
    for value in values {
      headers.append(name, value)
    }
  }
  headers
}

///|
pub fn append(self : Headers, name : String, value : String) -> Unit {
  let key = name.to_lower()
  if self.data.contains(key) {
    let header = self.data.get(key).unwrap().unwrap()
    header.values.push(value)
    self.data[key] = Some(header)
  } else {
    self.data[key] = Some({ name, values: [value] })
  }
}

///|
pub fn entries(self : Headers) -> Array[(String, String)] {
  let values = self.data.values().collect()
  // Note that this is optimized for a single value per key which is
  // the most common case.
  let entries = Array::new(capacity=values.length())
  for header in values {
    let header = header.unwrap()
    for value in header.values {
      entries.push((header.name, value))
    }
  }
  entries
}

///|
pub fn get(self : Headers, name : String) -> String {
  let key = name.to_lower()
  if self.data.contains(key) {
    let header = self.data.get(key).unwrap().unwrap()
    header.values.join(",")
  } else {
    ""
  }
}
