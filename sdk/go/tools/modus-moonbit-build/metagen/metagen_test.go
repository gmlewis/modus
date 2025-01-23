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

func TestGenerateMetadata(t *testing.T) {
	config := &config.Config{
		SourceDir: "testdata",
	}
	mod := &modinfo.ModuleInfo{
		ModulePath:      "github.com/gmlewis/modus/examples/simple-example",
		ModusSDKVersion: version.Must(version.NewVersion("40.11.0")),
	}

	meta, err := GenerateMetadata(config, mod)
	if err != nil {
		t.Fatalf("GenerateMetadata returned an error: %v", err)
	}

	if got, want := meta.Plugin, "simple-example"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@testdata"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if got, want := meta.SDK, "modus-sdk-mbt@40.11.0"; got != want {
		t.Errorf("meta.SDK = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantTypes, meta.Types); diff != "" {
		t.Errorf("meta.Types mismatch (-want +got):\n%v", diff)
	}
}

var wantFnExports = metadata.FunctionMap{
	"add": {
		Name:       "add",
		Parameters: []*metadata.Parameter{{Name: "x", Type: "Int"}, {Name: "y", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"add3": {
		Name:       "add3",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}, {Name: "c~", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"add3_WithDefaults": {
		Name:       "add3_WithDefaults",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"add_n": {
		Name:       "add_n",
		Parameters: []*metadata.Parameter{{Name: "args", Type: "Array[Int]"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"get_current_time": {
		Name:       "get_current_time",
		Parameters: []*metadata.Parameter{{Name: "now~", Type: "@wallClock.Datetime"}},
		Results:    []*metadata.Result{{Type: "@time.ZonedDateTime!Error"}},
	},
	"get_current_time_WithDefaults": {
		Name:    "get_current_time_WithDefaults",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime!Error"}},
	},
	"get_current_time_formatted": {
		Name:       "get_current_time_formatted",
		Parameters: []*metadata.Parameter{{Name: "now~", Type: "@wallClock.Datetime"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
	},
	"get_current_time_formatted_WithDefaults": {
		Name:    "get_current_time_formatted_WithDefaults",
		Results: []*metadata.Result{{Type: "String!Error"}},
	},
	"get_full_name": {
		Name:       "get_full_name",
		Parameters: []*metadata.Parameter{{Name: "first_name", Type: "String"}, {Name: "last_name", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"get_name_and_age": {Name: "get_name_and_age", Results: []*metadata.Result{{Type: "(String, Int)"}}},
	"get_people": {
		Name:    "get_people",
		Results: []*metadata.Result{{Type: "Array[@testdata.Person]"}},
	},
	"get_person": {Name: "get_person", Results: []*metadata.Result{{Type: "@testdata.Person"}}},
	"get_random_person": {
		Name:    "get_random_person",
		Results: []*metadata.Result{{Type: "@testdata.Person"}},
	},
	"log_message": {
		Name:       "log_message",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
	},
	"say_hello": {
		Name:       "say_hello",
		Parameters: []*metadata.Parameter{{Name: "name~", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"say_hello_WithDefaults": {Name: "say_hello_WithDefaults", Results: []*metadata.Result{{Type: "String"}}},
	"test_abort":             {Name: "test_abort"},
	"test_alternative_error": {
		Name:       "test_alternative_error",
		Parameters: []*metadata.Parameter{{Name: "input", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"test_exit":    {Name: "test_exit"},
	"test_logging": {Name: "test_logging"},
	"test_normal_error": {
		Name:       "test_normal_error",
		Parameters: []*metadata.Parameter{{Name: "input", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
	},
}

var wantFnImports = metadata.FunctionMap{}

var wantTypes = metadata.TypeMap{
	"(String, Int)": {Id: 4, Name: "(String, Int)"},
	"@testdata.Person": {
		Id:   5,
		Name: "@testdata.Person",
		Fields: []*metadata.Field{
			{Name: "firstName", Type: "String"}, {Name: "lastName", Type: "String"},
			{Name: "age", Type: "Int"},
		},
	},
	"@time.ZonedDateTime!Error": {Id: 6, Name: "@time.ZonedDateTime!Error"},
	"@wallClock.Datetime":       {Id: 7, Name: "@wallClock.Datetime"},
	"Array[@testdata.Person]":   {Id: 8, Name: "Array[@testdata.Person]"},
	"Array[Int]":                {Id: 9, Name: "Array[Int]"},
	"Int":                       {Id: 10, Name: "Int"},
	"String!Error":              {Id: 11, Name: "String!Error"},
	"String?":                   {Id: 12, Name: "String?"},
}
