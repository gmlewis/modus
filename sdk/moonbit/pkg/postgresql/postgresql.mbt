// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) suberror PostgreSqlError String derive(Show, Eq)

///|
const DB_TYPE = "postgresql"

///|
pub fn[T : @json.FromJson] query(
  host_name : String,
  statement : String,
  params : Json?
) -> (Array[T], UInt) raise Error {
  @db.query(host_name, DB_TYPE, statement, params)
}

///|
pub fn[T : @json.FromJson] query_scalar(
  host_name : String,
  statement : String,
  params : Json?
) -> (T, UInt) raise Error {
  @db.query_scalar(host_name, DB_TYPE, statement, params)
}

///|
pub fn execute(
  host_name : String,
  statement : String,
  params : Json?
) -> UInt raise Error {
  @db.execute(host_name, DB_TYPE, statement, params)
}
