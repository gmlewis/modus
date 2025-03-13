// -*- compile-command: "go test -run ^TestFilterMetadata_Neo4j$ ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package wasm

import (
	"testing"

	"github.com/gmlewis/modus/lib/wasmextractor"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/google/go-cmp/cmp"
)

const (
	neo4jPath      = "../../../../moonbit/examples/neo4j"
	neo4jBuildPath = neo4jPath + "/build"
)

func TestFilterMetadata_Neo4j(t *testing.T) {
	config := &config.Config{
		SourceDir:    neo4jPath,
		OutputDir:    neo4jBuildPath,
		WasmFileName: "neo4j.wasm",
	}
	copyOfBefore := deepCopyMetadata(t, wantNeo4jBeforeFilter)
	testFilterMetadataHelper(t, config, copyOfBefore, wantNeo4jAfterFilter)
}

// The Go `dlv` debugger has problems running the previous test
// so this test provides all the necessary data without requiring
// compilation of the wasm plugin.
func TestFilterMetadataNoCompilation_Neo4j(t *testing.T) {
	// make a copy of the "BEFORE" metadata since it is modified in-place.
	copyOfBefore := deepCopyMetadata(t, wantNeo4jBeforeFilter)
	filterExportsImportsAndTypes(neo4jWasmInfo, copyOfBefore)

	if diff := cmp.Diff(wantNeo4jAfterFilter, copyOfBefore); diff != "" {
		t.Fatalf("filterExportsImportsAndTypes meta AFTER filter mismatch (-want +got):\n%v", diff)
	}
}

var wantNeo4jBeforeFilter = &metadata.Metadata{
	Plugin: "neo4j",
	Module: "@neo4j",
	FnExports: metadata.FunctionMap{
		"__modus_create_people_and_relationships": {
			Name:    "__modus_create_people_and_relationships",
			Results: []*metadata.Result{{Type: "String!Error"}},
		},
		"__modus_delete_all_nodes": {
			Name:    "__modus_delete_all_nodes",
			Results: []*metadata.Result{{Type: "String!Error"}},
		},
		"__modus_get_alice_friends_under_40": {
			Name:    "__modus_get_alice_friends_under_40",
			Results: []*metadata.Result{{Type: "Array[Person]!Error"}},
		},
		"__modus_get_alice_friends_under_40_ages": {
			Name:    "__modus_get_alice_friends_under_40_ages",
			Results: []*metadata.Result{{Type: "Array[Int]!Error"}},
		},
		"cabi_realloc": {
			Name: "cabi_realloc",
			Parameters: []*metadata.Parameter{
				{Name: "src_offset", Type: "Int"}, {Name: "src_size", Type: "Int"},
				{Name: "_dst_alignment", Type: "Int"}, {Name: "dst_size", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
		"copy": {
			Name:       "copy",
			Parameters: []*metadata.Parameter{{Name: "dest", Type: "Int"}, {Name: "src", Type: "Int"}},
		},
		"create_people_and_relationships": {
			Name:    "create_people_and_relationships",
			Results: []*metadata.Result{{Type: "String!Error"}},
		},
		"delete_all_nodes": {Name: "delete_all_nodes", Results: []*metadata.Result{{Type: "String!Error"}}},
		"get_alice_friends_under_40": {
			Name:    "get_alice_friends_under_40",
			Results: []*metadata.Result{{Type: "Array[Person]!Error"}},
		},
		"get_alice_friends_under_40_ages": {
			Name:    "get_alice_friends_under_40_ages",
			Results: []*metadata.Result{{Type: "Array[Int]!Error"}},
		},
		"malloc": {
			Name:       "malloc",
			Parameters: []*metadata.Parameter{{Name: "size", Type: "Int"}},
			Results:    []*metadata.Result{{Type: "Int"}},
		},
		"ptr_to_none": {Name: "ptr_to_none", Results: []*metadata.Result{{Type: "Int"}}},
		"read_map": {
			Name: "read_map",
			Parameters: []*metadata.Parameter{
				{Name: "key_type_name_ptr", Type: "Int"},
				{Name: "value_type_name_ptr", Type: "Int"}, {Name: "map_ptr", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int64"}},
		},
		"write_map": {
			Name: "write_map",
			Parameters: []*metadata.Parameter{
				{Name: "key_type_name_ptr", Type: "Int"},
				{Name: "value_type_name_ptr", Type: "Int"}, {Name: "keys_ptr", Type: "Int"},
				{Name: "values_ptr", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
	},
	FnImports: metadata.FunctionMap{
		"modus_neo4j_client.executeQuery": {
			Name: "modus_neo4j_client.executeQuery",
			Parameters: []*metadata.Parameter{
				{Name: "host_name", Type: "String"},
				{Name: "db_name", Type: "String"},
				{Name: "query", Type: "String"},
				{Name: "parameters_json", Type: "Map[String, Json]"},
			},
			Results: []*metadata.Result{{Type: "@neo4j.EagerResult?!Error"}},
		},
		"modus_system.logMessage": {
			Name:       "modus_system.logMessage",
			Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
		},
	},
	Types: metadata.TypeMap{
		"(String)": {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
		"(String, Bool)": {
			Name:   "(String, Bool)",
			Fields: []*metadata.Field{{Name: "0", Type: "String"}, {Name: "1", Type: "Bool"}},
		},
		"@neo4j.&Entity": {Name: "@neo4j.&Entity"},
		"@neo4j.EagerResult": {
			Name: "@neo4j.EagerResult",
			Fields: []*metadata.Field{
				{Name: "keys", Type: "Array[String]"},
				{Name: "records", Type: "Array[@neo4j.Record]"},
			},
		},
		"@neo4j.EagerResult!Error": {
			Name: "@neo4j.EagerResult!Error",
			Fields: []*metadata.Field{
				{Name: "keys", Type: "Array[String]"},
				{Name: "records", Type: "Array[@neo4j.Record]"},
			},
		},
		"@neo4j.EagerResult?": {
			Name: "@neo4j.EagerResult?",
			Fields: []*metadata.Field{
				{Name: "keys", Type: "Array[String]"},
				{Name: "records", Type: "Array[@neo4j.Record]"},
			},
		},
		"@neo4j.Neo4jOption": {Name: "@neo4j.Neo4jOption"},
		"@neo4j.Node": {
			Name: "@neo4j.Node",
			Fields: []*metadata.Field{
				{Name: "element_id", Type: "String"},
				{Name: "labels", Type: "Array[String]"},
				{Name: "props", Type: "Map[String, Json]"},
			},
		},
		"@neo4j.Point2D": {
			Name: "@neo4j.Point2D",
			Fields: []*metadata.Field{
				{Name: "x", Type: "Double"}, {Name: "y", Type: "Double"},
				{Name: "spatial_ref_id", Type: "UInt"},
			},
		},
		"@neo4j.Point3D": {
			Name: "@neo4j.Point3D",
			Fields: []*metadata.Field{
				{Name: "x", Type: "Double"}, {Name: "y", Type: "Double"},
				{Name: "z", Type: "Double"}, {Name: "spatial_ref_id", Type: "UInt"},
			},
		},
		"@neo4j.Record": {
			Name:   "@neo4j.Record",
			Fields: []*metadata.Field{{Name: "values", Type: "Array[String]"}, {Name: "keys", Type: "Array[String]"}},
		},
		"@neo4j.Relationship": {
			Name: "@neo4j.Relationship",
			Fields: []*metadata.Field{
				{Name: "element_id", Type: "String"},
				{Name: "start_element_id", Type: "String"},
				{Name: "end_element_id", Type: "String"}, {Name: "type_", Type: "String"},
				{Name: "props", Type: "Json"},
			},
		},
		"@neo4j.T":       {Name: "@neo4j.T"},
		"@neo4j.T!Error": {Name: "@neo4j.T!Error"},
		"@testutils.CallStack[T]": {
			Name:   "@testutils.CallStack[T]",
			Fields: []*metadata.Field{{Name: "items", Type: "Array[Array[@testutils.T]]"}},
		},
		"@testutils.T":               {Name: "@testutils.T"},
		"Array[@neo4j.Neo4jOption]":  {Name: "Array[@neo4j.Neo4jOption]"},
		"Array[@neo4j.Record]":       {Name: "Array[@neo4j.Record]"},
		"Array[@testutils.T]":        {Name: "Array[@testutils.T]"},
		"Array[Array[@testutils.T]]": {Name: "Array[Array[@testutils.T]]"},
		"Array[Int]":                 {Name: "Array[Int]"},
		"Array[Int]!Error":           {Name: "Array[Int]!Error"},
		"Array[Json]":                {Name: "Array[Json]"},
		"Array[Person]":              {Name: "Array[Person]"},
		"Array[Person]!Error":        {Name: "Array[Person]!Error"},
		"Array[String]":              {Name: "Array[String]"},
		"Bool":                       {Name: "Bool"},
		"Double":                     {Name: "Double"},
		"FixedArray[Double]":         {Name: "FixedArray[Double]"},
		"FixedArray[Float]":          {Name: "FixedArray[Float]"},
		"FixedArray[Int64]":          {Name: "FixedArray[Int64]"},
		"FixedArray[Int]":            {Name: "FixedArray[Int]"},
		"FixedArray[UInt64]":         {Name: "FixedArray[UInt64]"},
		"FixedArray[UInt]":           {Name: "FixedArray[UInt]"},
		"Float":                      {Name: "Float"},
		"Int":                        {Name: "Int"},
		"Int64":                      {Name: "Int64"},
		"Json":                       {Name: "Json"},
		"Map[String, Json]":          {Name: "Map[String, Json]"},
		"Map[String, String]":        {Name: "Map[String, String]"},
		"Person": {
			Name:   "Person",
			Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "age", Type: "Int"}},
		},
		"String":       {Name: "String"},
		"String!Error": {Name: "String!Error"},
		"UInt":         {Name: "UInt"},
		"UInt64":       {Name: "UInt64"},
	},
}

var wantNeo4jAfterFilter = &metadata.Metadata{
	Plugin: "neo4j",
	Module: "@neo4j",
	FnExports: metadata.FunctionMap{
		"cabi_realloc": {
			Name: "cabi_realloc",
			Parameters: []*metadata.Parameter{
				{Name: "src_offset", Type: "Int"}, {Name: "src_size", Type: "Int"},
				{Name: "_dst_alignment", Type: "Int"}, {Name: "dst_size", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
		"copy": {
			Name:       "copy",
			Parameters: []*metadata.Parameter{{Name: "dest", Type: "Int"}, {Name: "src", Type: "Int"}},
		},
		"create_people_and_relationships": {
			Name:    "create_people_and_relationships",
			Results: []*metadata.Result{{Type: "String!Error"}},
		},
		"delete_all_nodes": {Name: "delete_all_nodes", Results: []*metadata.Result{{Type: "String!Error"}}},
		"get_alice_friends_under_40": {
			Name:    "get_alice_friends_under_40",
			Results: []*metadata.Result{{Type: "Array[Person]!Error"}},
		},
		"get_alice_friends_under_40_ages": {
			Name:    "get_alice_friends_under_40_ages",
			Results: []*metadata.Result{{Type: "Array[Int]!Error"}},
		},
		"malloc": {
			Name:       "malloc",
			Parameters: []*metadata.Parameter{{Name: "size", Type: "Int"}},
			Results:    []*metadata.Result{{Type: "Int"}},
		},
		"ptr_to_none": {Name: "ptr_to_none", Results: []*metadata.Result{{Type: "Int"}}},
		"read_map": {
			Name: "read_map",
			Parameters: []*metadata.Parameter{
				{Name: "key_type_name_ptr", Type: "Int"},
				{Name: "value_type_name_ptr", Type: "Int"}, {Name: "map_ptr", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int64"}},
		},
		"write_map": {
			Name: "write_map",
			Parameters: []*metadata.Parameter{
				{Name: "key_type_name_ptr", Type: "Int"},
				{Name: "value_type_name_ptr", Type: "Int"}, {Name: "keys_ptr", Type: "Int"},
				{Name: "values_ptr", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
	},
	FnImports: metadata.FunctionMap{
		"modus_neo4j_client.executeQuery": {
			Name: "modus_neo4j_client.executeQuery",
			Parameters: []*metadata.Parameter{
				{Name: "host_name", Type: "String"},
				{Name: "db_name", Type: "String"},
				{Name: "query", Type: "String"},
				{Name: "parameters_json", Type: "Map[String, Json]"},
			},
			Results: []*metadata.Result{{Type: "@neo4j.EagerResult?!Error"}},
		},
		"modus_system.logMessage": {
			Name:       "modus_system.logMessage",
			Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
		},
	},
	Types: metadata.TypeMap{
		"(String)": {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
		"@neo4j.EagerResult": {
			Name: "@neo4j.EagerResult",
			Fields: []*metadata.Field{
				{Name: "keys", Type: "Array[String]"},
				{Name: "records", Type: "Array[@neo4j.Record]"},
			},
		},
		"@neo4j.Record": {
			Name:   "@neo4j.Record",
			Fields: []*metadata.Field{{Name: "values", Type: "Array[String]"}, {Name: "keys", Type: "Array[String]"}},
		},
		"Array[@neo4j.Record]": {Name: "Array[@neo4j.Record]"},
		"Array[Int]":           {Name: "Array[Int]"},
		"Array[Int]!Error":     {Name: "Array[Int]!Error"},
		"Array[Json]":          {Name: "Array[Json]"},
		"Array[Person]":        {Name: "Array[Person]"},
		"Array[Person]!Error":  {Name: "Array[Person]!Error"},
		"Array[String]":        {Name: "Array[String]"},
		"Int":                  {Name: "Int"},
		"Int64":                {Name: "Int64"},
		"Json":                 {Name: "Json"},
		"Map[String, Json]":    {Name: "Map[String, Json]"},
		"Person": {
			Name:   "Person",
			Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "age", Type: "Int"}},
		},
		"String":       {Name: "String"},
		"String!Error": {Name: "String!Error"},
	},
}

var neo4jWasmInfo = &wasmextractor.WasmInfo{
	Exports: []wasmextractor.WasmItem{
		{Name: "cabi_realloc"},
		{Name: "copy"},
		{Name: "create_people_and_relationships"},
		{Name: "delete_all_nodes"},
		{Name: "free"},
		{Name: "get_alice_friends_under_40"},
		{Name: "get_alice_friends_under_40_ages"},
		{Name: "load32"},
		{Name: "malloc"},
		{Name: "memory"},
		{Name: "ptr2str"},
		{Name: "ptr_to_none"},
		{Name: "read_map"},
		{Name: "store32"},
		{Name: "store8"},
		{Name: "write_map"},
	},
	Imports: []wasmextractor.WasmItem{
		{Name: "modus_neo4j_client.executeQuery"},
		{Name: "modus_system.logMessage"},
		{Name: "spectest.print_char"},
	},
}
