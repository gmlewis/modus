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
  @base64.array2str(self.body[:])
}

///| This converts the JSON response body to a MoonBit struct of type `T`.
pub fn[T : @json.FromJson] Response::json(self : Response) -> T raise Error {
  @json.from_json(@json.parse(self.text()))
}

///| This returns the body as an array of Ints.
/// Array[Int] is used instead of Array[Byte] since MoonBit supports FromJson for Array[Int] but not for Array[Byte].
pub fn Response::to_ints(self : Response) -> Array[Int] {
  self.body.map(fn(c) { c.to_int() })
}
