// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
type! GraphQLError String derive(Show, Eq)

///|
pub fn execute[T : @json.FromJson](
  host_name : String,
  statement : String,
  variables : Map[String, Json]
) -> Response[T]!Error {
  let vars_str = @json.stringify(variables.to_json())
  let response = @host.execute_query(host_name, statement, vars_str)
  if response.is_empty() {
    raise GraphQLError("Failed to execute the GQL query")
  }

  //
  let result : Response[T] = @json.from_json!(@json.parse!(response.unwrap()))
  result
}
// func Execute[T any](hostName, statement string, variables map[string]any) (*Response[T], error) {
// 	bytes, err := utils.JsonSerialize(variables)
// 	if err != nil {
// 		console.Error(err.Error())
// 		return nil, err
// 	}
//
// 	varsStr := string(bytes)
//
// 	response := hostExecuteQuery(&hostName, &statement, &varsStr)
//
// 	if response == nil {
// 		return nil, errors.New("Failed to execute the GQL query.")
// 	}
//
// 	var result Response[T]
// 	err = utils.JsonDeserialize([]byte(*response), &result)
// 	if err != nil {
// 		console.Error(err.Error())
// 		return nil, err
// 	}
//
// 	if len(result.Errors) > 0 {
// 		errBytes, err := utils.JsonSerialize(result.Errors)
// 		if err != nil {
// 			console.Error(err.Error())
// 			return nil, err
// 		}
// 		console.Error(fmt.Sprint("GraphQL API Errors:" + string(errBytes)))
// 	}
// 	return &result, nil
// }

///|
pub(all) struct Response[T] {
  errors : Array[ErrorResult]
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
