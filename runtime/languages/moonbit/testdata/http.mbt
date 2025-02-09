// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub struct HttpResponse {
  status : UInt16
  status_text : String
  headers : HttpHeaders?
  body : Array[Byte]
} derive(Show, Eq)

///|
pub struct HttpHeaders {
  data : Map[String, HttpHeader?]
} derive(Show, Eq)

///|
pub struct HttpHeader {
  name : String
  values : Array[String]
} derive(Show, Eq)

///|
pub fn test_http_response_headers(r : HttpResponse?) -> Unit {
  guard let Some(r) = r else { _ => fail("Response is None") }
  guard let Some(h) = r.headers else { _ => fail("Headers is None") }
  dump_header_map(h.data)
}

///|
fn dump_header_map(m : Map[String, HttpHeader?]) -> Unit {
  if m.size() == 0 {
    fail("expected headers > 0, but got none")
  }
  println("Headers:")
  for k, v in m {
    println("  \{k}: \{v}")
    test_http_header(v)
  }
}

///|
pub fn test_http_headers(h : HttpHeaders) -> Unit {
  dump_header_map(h.data)
}

///|
pub fn test_http_header_map(m : Map[String, HttpHeader?]) -> Unit {
  dump_header_map(m)
}

///|
pub fn test_http_header(h : HttpHeader?) -> Unit {
  guard let Some(h) = h else { _ => fail("Header is None") }
  if h.name == "" {
    fail("Header.Name is empty")
  }
  if h.values.is_empty() {
    fail("Header.Values is None")
  }
  println("Header: \{h}\n")
}
