// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
typealias Request = @host.Request

///|
// typealias Query = @host.Query

///|
// typealias Mutation = @host.Mutation

///|
typealias Response = @host.Response

///|
type! DgraphError String derive(Show, Eq)

///| Executes a DQL query or mutation on the Dgraph database.
pub fn execute(host_name : String, request : Request) -> Response!Error {
  let response = @host.execute_query(host_name, request)
  if response.is_empty() {
    raise DgraphError("Failed to execute the DQL query.")
  }
  response.unwrap()
}

///| Alters the schema of the dgraph database
pub fn alter_schema(host_name : String, schema : String) -> Unit!Error {
  let response = @host.alter_schema(host_name, schema)
  if response.is_empty() {
    raise DgraphError("Failed to alter the schema.")
  }
}

///| Drops an attribute from the schema.
pub fn drop_attr(host_name : String, attr : String) -> Unit!Error {
  let response = @host.drop_attribute(host_name, attr)
  if response.is_empty() {
    raise DgraphError("Failed to drop the attribute.")
  }
}

///| Drops all data from the database.
pub fn drop_all(host_name : String) -> Unit!Error {
  let response = @host.drop_all_data(host_name)
  if response.is_empty() {
    raise DgraphError("Failed to drop all data.")
  }
}
