// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| HTTP Request record
pub(all) struct Request {
  url : String
  method_ : String
  headers : Headers?
  body : Array[Byte]
} derive(Show, Eq)

///|
pub(all) enum RequestOptions {
  Method(String)
  HeaderArray(Array[Array[String]])
  HeaderMap(Map[String, String])
  HeaderMapArray(Map[String, Array[String]])
  HeaderTuples(Array[(String, String)])
  BodyString(String)
  BodyJson(Json)
}

///|
pub fn Request::new(
  url : String,
  options~ : Array[RequestOptions] = []
) -> Request {
  let req = { url, method_: "GET", headers: None, body: [] }
  req.clone(options)
}

///|
pub fn Request::clone(
  self : Request,
  options : Array[RequestOptions]
) -> Request {
  let mut method_ = self.method_
  let mut headers = match self.headers {
    Some(value) => value.clone()
    None => Headers::new()
  }
  let mut body = self.body
  for option in options {
    match option {
      RequestOptions::Method(value) => method_ = value
      RequestOptions::HeaderArray(value) => headers = Headers::from_array(value)
      RequestOptions::HeaderMap(value) => headers = Headers::from_map(value)
      RequestOptions::HeaderMapArray(value) =>
        headers = Headers::from_map_array(value)
      RequestOptions::HeaderTuples(value) =>
        headers = Headers::from_tuples(value)
      RequestOptions::BodyString(value) =>
        body = @string.to_bytes(value).to_array() // TODO: utf8/utf16
      RequestOptions::BodyJson(value) =>
        body = @string.to_bytes(@json.stringify(value)).to_array()
    }
  }
  let data = headers.entries()
  let headers = if data.length() > 0 { Some(headers) } else { None }
  { url: self.url, method_, headers, body }
}

///|
pub fn Request::text(self : Request) -> String {
  self.body.to_string() // TODO
}

///|
pub fn Request::json[T : @json.FromJson](self : Request) -> T!Error {
  @json.from_json!(@json.parse!(self.body.to_string())) // TODO
}
