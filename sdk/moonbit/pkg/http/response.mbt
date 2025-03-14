// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| HTTP Response record
pub(all) struct Response {
  status : UInt16
  status_text : String
  headers : Headers?
  body : Array[Byte]
} derive(Show, Eq)

///|
pub fn ok(self : Response) -> Bool {
  self.status >= 200 && self.status < 300
}

///|
pub fn Response::text(self : Response) -> String {
  self.body.to_string() // TODO
}

///|
pub fn Response::json[T : @json.FromJson](self : Response) -> T!Error {
  @json.from_json!(@json.parse!(self.body.to_string())) // TODO
}
