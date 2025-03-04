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
	meta := setupTestConfig(t, "testdata/neo4j-example")
	removeExternalFuncsForComparison(t, meta)

	if got, want := meta.Plugin, "neo4j"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@neo4j-example"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	// if got, want := meta.SDK, "modus-sdk-mbt@0.16.5"; got != want {
	// 	t.Errorf("meta.SDK = %q, want %q", got, want)
	// }

	if diff := cmp.Diff(wantNeo4jFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantNeo4jFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	diffMetaTypes(t, wantNeo4jTypes, meta.Types)
	// if diff := cmp.Diff(wantNeo4jTypes, meta.Types); diff != "" {
	// 	t.Errorf("meta.Types mismatch (-want +got):\n%v", diff)
	// }
}

var wantNeo4jFnExports = metadata.FunctionMap{
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
}

var wantNeo4jFnImports = metadata.FunctionMap{
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
}

var wantNeo4jTypes = metadata.TypeMap{
	"(String)": {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
	"@neo4j.EagerResult": {
		Name:   "@neo4j.EagerResult",
		Fields: []*metadata.Field{{Name: "keys", Type: "Array[String]"}, {Name: "records", Type: "Array[@neo4j.Record?]"}},
	},
	"@neo4j.EagerResult?":       {Name: "@neo4j.EagerResult?"},
	"@neo4j.EagerResult?!Error": {Name: "@neo4j.EagerResult?!Error"},
	"@neo4j.Record": {
		Name:   "@neo4j.Record",
		Fields: []*metadata.Field{{Name: "values", Type: "Array[String]"}, {Name: "keys", Type: "Array[String]"}},
	},
	"@neo4j.Record?":        {Name: "@neo4j.Record?"},
	"Array[@neo4j.Record?]": {Name: "Array[@neo4j.Record?]"},
	"Array[Int]":            {Name: "Array[Int]"},
	"Array[Int]!Error":      {Name: "Array[Int]!Error"},
	"Array[Json]":           {Name: "Array[Json]"},
	"Array[Person]":         {Name: "Array[Person]"},
	"Array[Person]!Error":   {Name: "Array[Person]!Error"},
	"Array[String]":         {Name: "Array[String]"},
	"Json":                  {Name: "Json"},
	"Person": {
		Name:   "Person",
		Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "age", Type: "Int"}},
	},
	"String":       {Name: "String"},
	"String!Error": {Name: "String!Error"},
}
