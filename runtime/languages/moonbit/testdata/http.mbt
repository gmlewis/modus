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
pub fn test_http_response(r : HttpResponse?) -> Unit!Error {
  println("test_http_response")
  guard r is Some(r) else { raise TestError("Response is None") }
  println("Status: \{r.status}")
  println("StatusText: \{r.statusText}")
  guard r.headers is Some(h) else { raise TestError("Headers is None") }
  dump_header_map!(h.data)
  println("Body: \{r.body}")
}

///|
pub fn test_http_response_output() -> HttpResponse? {
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
pub fn test_http_response_input(response : HttpResponse) -> Unit!Error {
  guard test_http_response_output() is Some(expected) else {
    raise TestError("invalid response")
  }
  assert_eq!(expected, response)
}

///|
fn dump_header_map(m : Map[String, HttpHeader?]) -> Unit!Error {
  if m.size() == 0 {
    raise TestError("expected headers > 0, but got none")
  }
  println("Headers:")
  for k, v in m {
    println("  \{k}: \{v}")
    test_http_header!(v)
  }
  let expected = {
    "content-type": Some({ name: "Content-Type", values: ["text/plain"] }),
  }
  assert_eq!(expected, m)
}

///|
pub fn test_http_headers(h : HttpHeaders) -> Unit!Error {
  dump_header_map!(h.data)
}

///|
pub fn test_http_header_map(m : Map[String, HttpHeader?]) -> Unit!Error {
  dump_header_map!(m)
}

///|
pub fn test_http_header(h : HttpHeader?) -> Unit!Error {
  guard h is Some(h) else { raise TestError("Header is None") }
  if h.name == "" {
    raise TestError("Header.Name is empty")
  }
  if h.values.is_empty() {
    raise TestError("Header.Values is None")
  }
  println("Header: \{h}\n")
  let expected = { name: "Content-Type", values: ["text/plain"] }
  assert_eq!(expected, h)
}
