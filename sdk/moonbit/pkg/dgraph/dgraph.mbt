// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| Represents a Dgraph request.
pub(all) struct Request {
  query : Query?
  mutations : Array[Mutation]?
} derive(Show, Eq, FromJson, ToJson)

///| Represents a Dgraph query.
pub(all) struct Query {
  query : String
  variables : Map[String, String]
} derive(Show, Eq, FromJson, ToJson)

///| Represents a Dgraph mutation.
pub(all) struct Mutation {
  // JSON for setting data
  set_json : String
  // JSON for deleting data
  del_json : String
  // RDF N-Quads for setting data
  set_nquads : String
  // RDF N-Quads for deleting data
  del_nquads : String
  // Condition for the mutation, as a DQL @if directive
  condition : String
} derive(Show, Eq, FromJson, ToJson)

///|
pub fn set_json(set_json : String) -> Mutation {
  { set_json, del_json: "", set_nquads: "", del_nquads: "", condition: "" }
}

///|
pub fn del_json(del_json : String) -> Mutation {
  { set_json: "", del_json, set_nquads: "", del_nquads: "", condition: "" }
}

///|
pub fn set_nquads(set_nquads : String) -> Mutation {
  { set_json: "", del_json: "", set_nquads, del_nquads: "", condition: "" }
}

///|
pub fn del_nquads(del_nquads : String) -> Mutation {
  { set_json: "", del_json: "", set_nquads: "", del_nquads, condition: "" }
}

///|
pub fn condition(condition : String) -> Mutation {
  { set_json: "", del_json: "", set_nquads: "", del_nquads: "", condition }
}

///| Represents a Dgraph response.
pub(all) struct Response {
  json : String
  uids : Map[String, String]
} derive(Show, Eq, FromJson, ToJson)

///|
pub fn new_query(
  query : String,
  variables~ : Map[String, String] = {}
) -> Query {
  { query, variables }
}

///|
pub(all) type! DgraphError String derive(Show, Eq)

///| Executes a DQL query or mutation on the Dgraph database.
fn execute(connection : String, request : Request) -> Response!Error {
  let response = host_execute_query!(connection, request)
  if response.is_empty() {
    raise DgraphError("Failed to execute the DQL query.")
  }
  response.unwrap()
}

///| Executes a DQL query on the Dgraph database, optionally with mutations.
pub fn execute_query(
  connection : String,
  query : Query,
  mutations~ : Array[Mutation] = []
) -> Response!Error {
  let request = {
    query: Some(query),
    mutations: match mutations.is_empty() {
      true => None
      false => Some(mutations)
    },
  }
  execute!(connection, request)
}

///| Executes one or more mutations on the Dgraph database.
pub fn execute_mutations(
  connection : String,
  mutations : Array[Mutation]
) -> Response!Error {
  let request = { query: None, mutations: Some(mutations) }
  execute!(connection, request)
}

///| Alters the schema of the dgraph database
pub fn alter_schema(connection : String, schema : String) -> Unit!Error {
  host_alter_schema!(connection, schema)
}

///| Drops an attribute from the schema.
pub fn drop_attr(connection : String, attr : String) -> Unit!Error {
  host_drop_attribute!(connection, attr)
}

///| Drops all data from the database.
pub fn drop_all(connection : String) -> Unit!Error {
  host_drop_all_data!(connection)
}

///| Ensures proper escaping of RDF string literals
pub fn escape_rdf(value : String) -> String {
  let sb = StringBuilder::new()
  for r in value {
    match r {
      '\\' => sb.write_string("\\\\")
      '"' => sb.write_string("\\\"")
      '\n' => sb.write_string("\\n")
      '\r' => sb.write_string("\\r")
      '\t' => sb.write_string("\\t")
      _ => {
        // handle control characters
        let v = r.to_int()
        if v < 0x20 || (v >= 0x7F && v <= 0x9F) {
          sb.write_string("\\u")
          sb.write_string(int_to_hex_string(v, 4))
        } else {
          sb.write_char(r)
        }
      }
    }
  }
  sb.to_string()
}

///|
fn int_to_hex_string(n : Int, num_zeros : Int) -> String {
  let hex_str = n.to_string(radix=16)
  let zeros_needed = num_zeros - hex_str.length()
  if zeros_needed > 0 {
    let zeros = "0".repeat(zeros_needed)
    zeros + hex_str
  } else {
    hex_str
  }
}
