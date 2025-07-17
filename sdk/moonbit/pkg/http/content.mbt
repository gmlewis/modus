// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) struct Content {
  data : String // TODO: Bytes
} derive(Show, Eq, FromJson, ToJson)

///|
pub fn Content::new(value : Json) -> Content {
  let data = value.stringify()
  { data, }
}

///|
pub fn text(self : Content) -> String {
  self.data
}

///|
pub fn[T : @json.FromJson] json(self : Content) -> T raise Error {
  @json.from_json(@json.parse(self.data))
}
