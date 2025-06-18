// -*- compile-command: "go test -run ^TestGenerateMetadata_Neo4j$ ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package metagen

import (
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateMetadata_Neo4j(t *testing.T) {
	meta := setupTestConfig(t, "../testdata/neo4j-example")

	if got, want := meta.Plugin, "neo4j"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@neo4j-example"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantNeo4jFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantNeo4jFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	diffMetaTypes(t, wantNeo4jTypes, meta.Types)
}

var wantNeo4jFnExports = metadata.FunctionMap{
	"create_people_and_relationships": {
		Name:    "create_people_and_relationships",
		Results: []*metadata.Result{{Type: "String raise Error"}},
	},
	"delete_all_nodes": {Name: "delete_all_nodes", Results: []*metadata.Result{{Type: "String raise Error"}}},
	"get_alice_friends_under_40": {
		Name:    "get_alice_friends_under_40",
		Results: []*metadata.Result{{Type: "Array[Person] raise Error"}},
	},
	"get_alice_friends_under_40_ages": {
		Name:    "get_alice_friends_under_40_ages",
		Results: []*metadata.Result{{Type: "Array[Int] raise Error"}},
	},
}

var wantNeo4jFnImports = metadata.FunctionMap{
	"modus_neo4j_client.executeQuery": {
		Name: "modus_neo4j_client.executeQuery",
		Parameters: []*metadata.Parameter{
			{Name: "host_name", Type: "String"},
			{Name: "db_name", Type: "String"},
			{Name: "query", Type: "String"},
			{Name: "parameters_json", Type: "String"},
		},
		Results: []*metadata.Result{{Type: "@neo4j.EagerResult?"}},
	},
	"modus_system.logMessage": {
		Name:       "modus_system.logMessage",
		Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
	},
}

var wantNeo4jTypes = metadata.TypeMap{
	"(String)": {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
	"(String, Bool)": {
		Name:   "(String, Bool)",
		Fields: []*metadata.Field{{Name: "0", Type: "String"}, {Name: "1", Type: "Bool"}},
	},
	// "@neo4j.&Entity": {Name: "@neo4j.&Entity"},
	"@neo4j.EagerResult": {
		Name:   "@neo4j.EagerResult",
		Fields: []*metadata.Field{{Name: "keys", Type: "Array[String]"}, {Name: "records", Type: "Array[@neo4j.Record]"}},
	},
	"@neo4j.EagerResult?": {
		Name:   "@neo4j.EagerResult?",
		Fields: []*metadata.Field{{Name: "keys", Type: "Array[String]"}, {Name: "records", Type: "Array[@neo4j.Record]"}},
	},
	"@neo4j.EagerResult raise Error": {
		Name: "@neo4j.EagerResult raise Error",
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
	// "@neo4j.T":             {Name: "@neo4j.T"},
	// "@neo4j.T raise Error": {Name: "@neo4j.T raise Error"},
	// "@testutils.CallStack[T]": {
	// 	Name:   "@testutils.CallStack[T]",
	// 	Fields: []*metadata.Field{{Name: "items", Type: "Array[Array[@testutils.T]]"}},
	// },
	"@testutils.T":               {Name: "@testutils.T"},
	"Array[@neo4j.Neo4jOption]":  {Name: "Array[@neo4j.Neo4jOption]"},
	"Array[@neo4j.Node]":         {Name: "Array[@neo4j.Node]"},
	"Array[@neo4j.Record]":       {Name: "Array[@neo4j.Record]"},
	"Array[@neo4j.Relationship]": {Name: "Array[@neo4j.Relationship]"},
	"Array[@testutils.T]":        {Name: "Array[@testutils.T]"},
	"Array[Array[@testutils.T]]": {Name: "Array[Array[@testutils.T]]"},
	"Array[Int]":                 {Name: "Array[Int]"},
	"Array[Int] raise Error":     {Name: "Array[Int] raise Error"},
	"Array[Json]":                {Name: "Array[Json]"},
	"Array[Person]":              {Name: "Array[Person]"},
	"Array[Person] raise Error":  {Name: "Array[Person] raise Error"},
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
	"String":             {Name: "String"},
	"String raise Error": {Name: "String raise Error"},
	"UInt":               {Name: "UInt"},
	"UInt64":             {Name: "UInt64"},
}
