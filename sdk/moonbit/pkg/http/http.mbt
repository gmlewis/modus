// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) type! HttpError String derive(Show, Eq)

///|
pub fn fetch_url(
  url : String,
  options~ : Array[RequestOptions] = []
) -> Response!Error {
  let request = Request::new(url, options~)
  request.fetch!()
}

///|
pub fn fetch(self : Request) -> Response!Error {
  let response = host_fetch!(self)
  println("fetch: got \{response.body.length()} body bytes:\n\{response.body}")
  response
}
