// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) type! Neo4jError String derive(Show, Eq)

///|
pub(all) type Neo4jOption (Neo4jOptions) -> Unit

///|
pub typealias EagerResult = @host.EagerResult

///|
struct Neo4jOptions {
  mut db_name : String
}

///|
pub fn with_db_name(db_name : String) -> Neo4jOption {
  fn(options) { options.db_name = db_name }
}

///|
// pub(all) struct EagerResult {
//   keys : Array[String]
//   records : Array[Record]
// }

///|
pub(all) struct Record {
  values : Array[String]
  keys : Array[String]
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) enum RecordValue {
  Bool(Bool)
  Int64(Int64)
  Double(Double)
  String(String)
  // TODO: Time(Time)
  // TODO: Bytes(Bytes)
  // TODO: AnyArray(Array[Any])
  Map(Json)
  Node(Node)
  Relationship(Relationship)
  Path(Path)
  Point2D(Point2D)
  Point3D(Point3D)
} derive(Show, Eq, FromJson, ToJson)

///|
pub(open) trait Entity {
  get_element_id(Self) -> String
  get_properties(Self) -> Map[String, Json]
}

///|
pub(all) struct Node {
  element_id : String // `json:"ElementId"`
  labels : Array[String] // `json:"Labels"`
  props : Map[String, Json] // `json:"Props"`
} derive(Show, Eq, FromJson, ToJson)

///|
pub fn get_element_id(self : Node) -> String {
  self.element_id
}

///|
pub fn get_properties(self : Node) -> Map[String, Json] {
  self.props
}

///|
pub(all) struct Relationship {
  element_id : String // `json:"ElementId"`
  start_element_id : String // `json:"StartElementId"`
  end_element_id : String // `json:"EndElementId"`
  type_ : String // `json:"Type"`
  props : Json // `json:"Props"`
} derive(Show, Eq, FromJson, ToJson)

///|
pub fn get_element_id(self : Relationship) -> String {
  self.element_id
}

///|
pub fn get_properties(self : Relationship) -> Json {
  self.props
}

///|
pub(all) struct Path {
  nodes : Array[Node] // `json:"Nodes"`
  relationships : Array[Relationship] // `json:"Relationships"`
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) enum PropertyValue {
  Bool(Bool)
  Int64(Int64)
  Double(Double)
  String(String)
  // TODO: Time(Time)
  // TODO: Bytes(Bytes)
  // TODO: AnyArray(Array[Any])
  Point2D(Point2D)
  Point3D(Point3D)
} derive(Show, Eq, FromJson, ToJson)

///| Point2D represents a two dimensional point in a particular coordinate reference system.
pub(all) struct Point2D {
  x : Double
  y : Double
  spatial_ref_id : Int // TODO: UInt // Id of coordinate reference system.
} derive(Show, Eq, FromJson, ToJson)

///| Point3D represents a three dimensional point in a particular coordinate reference system.
pub(all) struct Point3D {
  x : Double
  y : Double
  z : Double
  spatial_ref_id : Int // TODO: UInt // Id of coordinate reference system.
} derive(Show, Eq, FromJson, ToJson)

///| String returns string representation of this point.
pub fn to_string(self : Point2D) -> String {
  "Point{SpatialRefId=\{self.spatial_ref_id}, X=\{self.x}, Y=\{self.y}}"
}

///| String returns string representation of this point.
pub fn to_string(self : Point3D) -> String {
  "Point{SpatialRefId=\{self.spatial_ref_id}, X=\{self.x}, Y=\{self.y}, Z=\{self.z}}"
}

///| ExecuteQuery executes a query or mutation on the Neo4j database.
pub fn execute_query(
  host_name : String,
  query : String,
  parameters : Json,
  opts~ : Array[Neo4jOption] = []
) -> EagerResult!Error {
  let db_opts = { db_name: "neo4j" }
  for opt in opts {
    (opt._)(db_opts)
  }

  //
  let parameters_json = @json.stringify(parameters)
  let response = @host.execute_query(
    host_name,
    db_opts.db_name,
    query,
    parameters_json,
  )
  if response.is_empty() {
    raise Neo4jError("Failed to execute the Neo4j query")
  }
  let response = response.unwrap()

  //
  response
}

///|
pub fn get_record_value(self : Record, key : String) -> RecordValue!Error {
  match self.get(key) {
    (value, true) => @json.from_json!(@json.parse!(value))
    (_, false) => raise Neo4jError("Key not found in record")
  }
}

///|
pub fn get(self : Record, key : String) -> (String, Bool) {
  for i in 0..<self.keys.length() {
    if self.keys[i] == key {
      return (self.values[i], true)
    }
  }
  ("", false)
}

///|
pub fn as_map(self : Record) -> Map[String, String] {
  let result = {}
  for i in 0..<self.keys.length() {
    result[self.keys[i]] = self.values[i]
  }
  result
}

///|
pub fn get_property(self : &Entity, key : String) -> PropertyValue!Error {
  let raw_val = self.get_properties().get(key)
  if raw_val.is_empty() {
    raise Neo4jError("Key not found in node")
  }
  let val = raw_val.unwrap()
  @json.from_json!(val)
}
