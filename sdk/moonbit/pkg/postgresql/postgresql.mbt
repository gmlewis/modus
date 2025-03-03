// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) type! PostgreSqlError String derive(Show, Eq)

///|
const DB_TYPE = "postgresql"

///|
pub fn query[T : @json.FromJson](
  host_name : String,
  statement : String,
  params : Json?
) -> (Array[T], UInt)!Error {
  @db.query!(host_name, DB_TYPE, statement, params)
}

///|
pub fn query_scalar[T : @json.FromJson](
  host_name : String,
  statement : String,
  params : Json?
) -> (T, UInt)!Error {
  @db.query_scalar!(host_name, DB_TYPE, statement, params)
}

///|
pub fn execute(
  host_name : String,
  statement : String,
  params : Json?
) -> UInt!Error {
  @db.execute!(host_name, DB_TYPE, statement, params)
}
