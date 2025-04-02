// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub struct HttpResponse {
  status : UInt16
  statusText : String
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
  println("test_http_response_headers")
  guard r is Some(r) else { fail("Response is None") }
  println("Status: \{r.status}")
  println("StatusText: \{r.statusText}")
  guard r.headers is Some(h) else { fail("Headers is None") }
  dump_header_map(h.data)
  println("Body: \{r.body}")
}

///|
pub fn test_http_response_headers_output() -> HttpResponse? {
  Some({
    status: 200,
    statusText: "OK",
    headers: Some({
      data: {
        "content-type": Some({ name: "Content-Type", values: ["text/plain"] }),
      },
    }),
    body: b"Hello, world!".to_array(),
  })
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
  guard h is Some(h) else { fail("Header is None") }
  if h.name == "" {
    fail("Header.Name is empty")
  }
  if h.values.is_empty() {
    fail("Header.Values is None")
  }
  println("Header: \{h}\n")
}
