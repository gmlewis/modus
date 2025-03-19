// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| Record type for query response
pub(all) struct HostQueryResponse {
  error : String?
  result_json : String?
  rows_affected : UInt
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) type! DbError String derive(Show, Eq)

///|
pub fn execute(
  host_name : String,
  db_type : String,
  statement : String,
  params : Json?
) -> UInt!Error {
  let (_, affected) = do_query!(host_name, db_type, statement, params)
  affected
}

///|
pub fn query[T : @json.FromJson](
  host_name : String,
  db_type : String,
  statement : String,
  params : Json?
) -> (Array[T], UInt)!Error {
  let (result_json, affected) = do_query!(host_name, db_type, statement, params)
  if result_json.is_empty() {
    return ([], affected)
  }
  let result_json = result_json.unwrap()

  //
  let rows : Array[T] = try {
    @json.from_json!(result_json)
  } catch {
    e => raise DbError("could not JSON deserialize database response: \{e}")
  }

  //
  (rows, affected)
}

///|
pub fn query_scalar[T : @json.FromJson](
  host_name : String,
  db_type : String,
  statement : String,
  params : Json?
) -> (T, UInt)!Error {
  let (rows, affected) = query!(host_name, db_type, statement, params)
  if rows.length() == 0 {
    raise DbError("no result returned from database query")
  }
  if rows.length() > 1 {
    raise DbError(
      "expected a single row from a scalar database query, but received \{rows.length()}",
    )
  }
  let row = rows[0]
  (row, affected)
}

///|
pub fn do_query(
  host_name : String,
  db_type : String,
  statement : String,
  params : Json?
) -> (Json?, UInt)!Error {
  let params_json = match params {
    Some(p) => @json.stringify(p)
    None => "[]"
  }

  //
  let statement = statement.trim_space()
  let response = host_execute_query(host_name, db_type, statement, params_json)
  if response.is_empty() {
    raise DbError("no response received from database query")
  }
  let response = response.unwrap()

  //
  // TODO: let affected = response.rows_affected
  let affected = response.rows_affected
  if not(response.error.is_empty()) {
    raise DbError("database returned an error: \{response.error.unwrap()}")
  }
  match response.result_json {
    Some(e) => (Some(@json.parse!(e)), affected)
    None => (None, affected)
  }
}
