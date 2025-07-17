// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) suberror GraphQLError String derive(Show, Eq)

///|
pub fn[T : @json.FromJson] execute(
  host_name : String,
  statement : String,
  vars~ : Map[String, Json] = {}
) -> Response[T] raise Error {
  let vars_str = vars.to_json().stringify()
  let response = host_execute_query(host_name, statement, vars_str)
  //
  let result : Response[T] = @json.from_json(@json.parse(response))
  if result.errors is Some(errors) && errors.length() > 0 {
    let err_str = result.errors.to_json().stringify()
    @console.error("GraphQL API Errors: \{err_str}")
  }

  //
  result
}

///|
pub(all) struct Response[T] {
  errors : Array[ErrorResult]?
  data : T?
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct ErrorResult {
  message : String
  locations : Array[CodeLocation]
  path : Array[String]
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct CodeLocation {
  line : UInt
  column : UInt
} derive(Show, Eq, FromJson, ToJson)
