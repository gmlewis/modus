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
  // The Modus MoonBit Runtime returns Go []byte as a MoonBit String
  // even though it is not value UTF-16. Therefore, this data needs
  // to be translated. This should probably be fixed. TODO.
  // Therefore do not use `body` directly, but use the helper methods instead.
  body : String
} derive(Show, Eq)

///|
pub fn ok(self : Response) -> Bool {
  self.status >= 200 && self.status < 300
}

///|
pub fn Response::text(self : Response) -> String {
  println("GML: Response::text: self.body.length(): \{self.body.length()}")
  // let bytes = self.body.map(fn(b) { b.to_byte() })
  // @base64.array2str(bytes[:])
  self.body
}

///| This converts the JSON response body to a MoonBit struct of type `T`.
pub fn Response::json[T : @json.FromJson](self : Response) -> T!Error {
  @json.from_json!(@json.parse!(self.text()))
}

///| This returns the body as raw bytes.
/// Array[Int] is used instead of Array[Byte] since MoonBit supports FromJson for Array[Int].
pub fn Response::bytes(self : Response) -> Array[Int] {
  println("GML: Response::bytes: self.body.length(): \{self.body.length()}")
  self.body.to_array().map(fn(c) { c.to_int() })
}
