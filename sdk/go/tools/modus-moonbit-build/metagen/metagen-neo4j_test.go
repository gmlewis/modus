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

	if got, want := meta.SDK, "modus-sdk-mbt@0.16.5"; got != want {
		t.Errorf("meta.SDK = %q, want %q", got, want)
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
			{Name: "host_name", Type: "String"}, {Name: "db_name", Type: "String"},
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
	"(String)": {Id: 4, Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
	"@neo4j.EagerResult": {Id: 5,
		Name:   "@neo4j.EagerResult",
		Fields: []*metadata.Field{{Name: "keys", Type: "Array[String]"}, {Name: "records", Type: "Array[@neo4j.Record?]"}},
	},
	"@neo4j.EagerResult?":       {Id: 6, Name: "@neo4j.EagerResult?"},
	"@neo4j.EagerResult?!Error": {Id: 7, Name: "@neo4j.EagerResult?!Error"},
	"@neo4j.Record": {Id: 8,
		Name:   "@neo4j.Record",
		Fields: []*metadata.Field{{Name: "values", Type: "Array[String]"}, {Name: "keys", Type: "Array[String]"}},
	},
	"@neo4j.Record?":        {Id: 9, Name: "@neo4j.Record?"},
	"Array[@neo4j.Record?]": {Id: 10, Name: "Array[@neo4j.Record?]"},
	"Array[Int]":            {Id: 11, Name: "Array[Int]"},
	"Array[Int]!Error":      {Id: 12, Name: "Array[Int]!Error"},
	"Array[Json]":           {Id: 13, Name: "Array[Json]"},
	"Array[Person]":         {Id: 14, Name: "Array[Person]"},
	"Array[Person]!Error":   {Id: 15, Name: "Array[Person]!Error"},
	"Array[String]":         {Id: 16, Name: "Array[String]"},
	"Json":                  {Id: 17, Name: "Json"},
	"Person": {Id: 18,
		Name:   "Person",
		Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "age", Type: "Int"}},
	},
	"String":       {Id: 19, Name: "String"},
	"String!Error": {Id: 20, Name: "String!Error"},
}
