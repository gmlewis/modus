// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) type Response @host.Response

///|
pub fn ok(self : Response) -> Bool {
  self.status >= 200 && self.status < 300
}

///|
pub fn text(self : Response) -> String {
  self.body
}

///|
pub fn json[T : @json.FromJson](self : Response) -> T!Error {
  @json.from_json!(@json.parse!(self.body))
}
