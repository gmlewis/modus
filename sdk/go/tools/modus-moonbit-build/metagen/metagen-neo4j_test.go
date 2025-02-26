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

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-version"
)

func TestGenerateMetadata_Neo4j(t *testing.T) {
	config := &config.Config{
		SourceDir: "testdata/neo4j-example",
	}
	mod := &modinfo.ModuleInfo{
		ModulePath:      "github.com/gmlewis/modus/examples/neo4j-example",
		ModusSDKVersion: version.Must(version.NewVersion("40.11.0")),
	}

	meta, err := GenerateMetadata(config, mod)
	if err != nil {
		t.Fatalf("GenerateMetadata returned an error: %v", err)
	}

	if got, want := meta.Plugin, "neo4j-example"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@neo4j-example"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if got, want := meta.SDK, "modus-sdk-mbt@40.11.0"; got != want {
		t.Errorf("meta.SDK = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantNeo4jFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantNeo4jFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantNeo4jTypes, meta.Types); diff != "" {
		t.Errorf("meta.Types mismatch (-want +got):\n%v", diff)
	}
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
		Results: []*metadata.Result{{Type: "Array[Int64]!Error"}},
	},
}

var wantNeo4jFnImports = metadata.FunctionMap{
	"modus_neo4j_client.executeQuery": {
		Name: "modus_neo4j_client.executeQuery",
		Parameters: []*metadata.Parameter{
			{Name: "host_name", Type: "String"}, {Name: "db_name", Type: "String"},
			{Name: "query", Type: "String"},
			{Name: "parameters_json", Type: "Map[String, Json]"},
		},
		Results: []*metadata.Result{{Type: "EagerResult?!Error"}},
	},
	"modus_system.logMessage": {
		Name:       "modus_system.logMessage",
		Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
	},
}

var wantNeo4jTypes = metadata.TypeMap{
	"(String)":            {Id: 4, Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
	"Array[Int64]":        {Id: 5, Name: "Array[Int64]"},
	"Array[Int64]!Error":  {Id: 6, Name: "Array[Int64]!Error"},
	"Array[Json]":         {Id: 7, Name: "Array[Json]"},
	"Array[Person]":       {Id: 8, Name: "Array[Person]"},
	"Array[Person]!Error": {Id: 9, Name: "Array[Person]!Error"},
	"Array[String]":       {Id: 10, Name: "Array[String]"},
	"EagerResult":         {Id: 11, Name: "EagerResult"},
	"EagerResult?":        {Id: 12, Name: "EagerResult?"},
	"EagerResult?!Error":  {Id: 13, Name: "EagerResult?!Error"},
	"Json":                {Id: 14, Name: "Json"},
	"Map[String, Json]":   {Id: 15, Name: "Map[String, Json]"},
	"Person": {Id: 16,
		Name:   "Person",
		Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "age", Type: "Int64"}},
	},
	"String":       {Id: 17, Name: "String"},
	"String!Error": {Id: 18, Name: "String!Error"},
}
