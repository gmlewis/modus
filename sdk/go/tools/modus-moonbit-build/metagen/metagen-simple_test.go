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

func TestGenerateMetadata_Simple(t *testing.T) {
	config := &config.Config{
		SourceDir: "testdata/simple-example",
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

	if got, want := meta.Module, "@simple-example"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if got, want := meta.SDK, "modus-sdk-mbt@40.11.0"; got != want {
		t.Errorf("meta.SDK = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantSimpleFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantSimpleFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantSimpleTypes, meta.Types); diff != "" {
		t.Errorf("meta.Types mismatch (-want +got):\n%v", diff)
	}
}

var wantSimpleFnExports = metadata.FunctionMap{
	"add": {
		Name:       "add",
		Parameters: []*metadata.Parameter{{Name: "x", Type: "Int"}, {Name: "y", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs:       &metadata.Docs{Lines: []string{"Adds two integers together and returns the result."}},
	},
	"add3": {
		Name:       "add3",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}, {Name: "c~", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs: &metadata.Docs{Lines: []string{
			"Adds three integers together and returns the result.",
			"The third integer is optional.",
		}},
	},
	"add3_WithDefaults": {
		Name:       "add3_WithDefaults",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs: &metadata.Docs{Lines: []string{
			"Adds three integers together and returns the result.",
			"The third integer is optional.",
		}},
	},
	"add_n": {
		Name:       "add_n",
		Parameters: []*metadata.Parameter{{Name: "args", Type: "Array[Int]"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs:       &metadata.Docs{Lines: []string{"Adds any number of integers together and returns the result."}},
	},
	"get_current_time": {
		Name:       "get_current_time",
		Parameters: []*metadata.Parameter{{Name: "now~", Type: "@wallClock.Datetime"}},
		Results:    []*metadata.Result{{Type: "@time.ZonedDateTime!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Returns the current time."}},
	},
	"get_current_time_WithDefaults": {
		Name:    "get_current_time_WithDefaults",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time."}},
	},
	"get_current_time_formatted": {
		Name:       "get_current_time_formatted",
		Parameters: []*metadata.Parameter{{Name: "now~", Type: "@wallClock.Datetime"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Returns the current time formatted as a string."}},
	},
	"get_current_time_formatted_WithDefaults": {
		Name:    "get_current_time_formatted_WithDefaults",
		Results: []*metadata.Result{{Type: "String!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time formatted as a string."}},
	},
	"get_full_name": {
		Name:       "get_full_name",
		Parameters: []*metadata.Parameter{{Name: "first_name", Type: "String"}, {Name: "last_name", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Combines the first and last name of a person, and returns the full name."}},
	},
	"get_name_and_age": {
		Name:    "get_name_and_age",
		Results: []*metadata.Result{{Type: "(String, Int)"}},
		Docs:    &metadata.Docs{Lines: []string{"Gets the name and age of a person."}},
	},
	"get_people": {
		Name:    "get_people",
		Results: []*metadata.Result{{Type: "Array[@simple-example.Person]"}},
		Docs:    &metadata.Docs{Lines: []string{"Gets a list of people."}},
	},
	"get_person": {
		Name:    "get_person",
		Results: []*metadata.Result{{Type: "@simple-example.Person"}},
		Docs:    &metadata.Docs{Lines: []string{"Gets a person object."}},
	},
	"get_random_person": {
		Name:    "get_random_person",
		Results: []*metadata.Result{{Type: "@simple-example.Person"}},
		Docs:    &metadata.Docs{Lines: []string{"Gets a random person object from a list of people."}},
	},
	"log_message": {
		Name:       "log_message",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Logs a message."}},
	},
	"say_hello": {
		Name:       "say_hello",
		Parameters: []*metadata.Parameter{{Name: "name~", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "String"}},
		Docs: &metadata.Docs{Lines: []string{
			"Says hello to a person by name.",
			"If the name is not provided, it will say hello without a name.",
		}},
	},
	"say_hello_WithDefaults": {
		Name:    "say_hello_WithDefaults",
		Results: []*metadata.Result{{Type: "String"}},
		Docs: &metadata.Docs{Lines: []string{
			"Says hello to a person by name.",
			"If the name is not provided, it will say hello without a name.",
		}},
	},
	"test_abort": {
		Name: "test_abort",
		Docs: &metadata.Docs{Lines: []string{"Tests an abort."}},
	},
	"test_alternative_error": {
		Name:       "test_alternative_error",
		Parameters: []*metadata.Parameter{{Name: "input", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Tests an alternative way to handle errors in functions."}},
	},
	"test_exit": {
		Name: "test_exit",
		Docs: &metadata.Docs{Lines: []string{"Tests an exit with a non-zero exit code."}},
	},
	"test_logging": {
		Name: "test_logging",
		Docs: &metadata.Docs{Lines: []string{"Tests logging at different levels."}},
	},
	"test_normal_error": {
		Name:       "test_normal_error",
		Parameters: []*metadata.Parameter{{Name: "input", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Tests returning an error."}},
	},
}

var wantSimpleFnImports = metadata.FunctionMap{
	"modus_system.getTimeInZone": {
		Name:       "modus_system.getTimeInZone",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
	},
	"modus_system.getTimeZoneData": {
		Name:       "modus_system.getTimeZoneData",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}, {Name: "format", Type: "String"}},
		Results:    []*metadata.Result{{Type: "Array[Byte]"}},
	},
	"modus_system.logMessage": {
		Name:       "modus_system.logMessage",
		Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
	},
}

var wantSimpleTypes = metadata.TypeMap{
	"(String, Int)": {Id: 4, Name: "(String, Int)"},
	"@simple-example.Person": {
		Id:   5,
		Name: "@simple-example.Person",
		Fields: []*metadata.Field{
			{Name: "firstName", Type: "String"}, {Name: "lastName", Type: "String"},
			{Name: "age", Type: "Int"},
		},
	},
	"@time.ZonedDateTime":           {Id: 6, Name: "@time.ZonedDateTime"},
	"@wallClock.Datetime":           {Id: 7, Name: "@wallClock.Datetime"},
	"Array[@simple-example.Person]": {Id: 8, Name: "Array[@simple-example.Person]"},
	"Array[Byte]":                   {Id: 9, Name: "Array[Byte]"},
	"Array[Int]":                    {Id: 10, Name: "Array[Int]"},
	"Int":                           {Id: 11, Name: "Int"},
	"String":                        {Id: 12, Name: "String"},
	"String?":                       {Id: 13, Name: "String?"},
}
