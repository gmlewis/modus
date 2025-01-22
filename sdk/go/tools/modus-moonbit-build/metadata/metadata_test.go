/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package metadata

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed testdata/simple-example-metadata.json
var simpleExampleMetadataJSON []byte

func TestFunction_String(t *testing.T) {
	t.Parallel()

	var meta *Metadata
	if err := json.Unmarshal(simpleExampleMetadataJSON, &meta); err != nil || meta == nil || meta.FnExports == nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	tests := []struct {
		name string
		want string
	}{
		{name: "add", want: "add(x : Int, y : Int) -> Int"},
		{name: "add3", want: "add3(a : Int, b : Int, c~ : Int = 0) -> Int"},
		{name: "add_n", want: "add_n(args : Array[Int]) -> Int"},
		{name: "get_current_time", want: "get_current_time(now~ : @wallClock.Datetime = @wallClock.now()) -> @time.PlainDateTime!Error"},
		{name: "get_current_time_formatted", want: "get_current_time_formatted(now~ : @wallClock.Datetime = @wallClock.now()) -> String!Error"},
		{name: "get_full_name", want: "get_full_name(first_name : String, last_name : String) -> String"},
		{name: "get_name_and_age", want: "get_name_and_age() -> (String, Int)"},
		{name: "get_people", want: "get_people() -> Array[@simple.Person]"},
		{name: "get_person", want: "get_person() -> @simple.Person"},
		{name: "get_random_person", want: "get_random_person() -> @simple.Person"},
		{name: "log_message", want: "log_message(message : String) -> Unit"},
		{name: "say_hello", want: "say_hello(name~ : String? = None) -> String"},
		{name: "test_abort", want: "test_abort() -> Unit"},
		{name: "test_alternative_error", want: "test_alternative_error(input : String) -> String"},
		{name: "test_exit", want: "test_exit() -> Unit"},
		{name: "test_logging", want: "test_logging() -> Unit"},
		{name: "test_normal_error", want: "test_normal_error(input : String) -> String!Error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := meta.FnExports[tt.name]
			got := f.String(meta)
			if got != tt.want {
				t.Errorf("function.String = %q, want %q", got, tt.want)
			}
		})
	}
}

/*
var testMetadata = &Metadata{
	FnExports: map[string]*Function{
		"add": {
			Name: "add",
			Parameters: []*Parameter{
				{Name: "x", Type: "Int"},
				{Name: "y", Type: "Int"},
			},
			Results: []*Result{{Type: "Int"}},
		},
		"add3": {
			Name: "add3",
			Parameters: []*Parameter{
				{Name: "a", Type: "Int"},
				{Name: "b", Type: "Int"},
				{Name: "c~", Type: "Int = 0"},
			},
			Results: []*Result{{Type: "Int"}},
		},
		"add_n": {
			Name: "add_n",
			Parameters: []*Parameter{
				{Name: "args", Type: "Array[Int]"},
			},
			Results: []*Result{{Type: "Int"}},
		},
		"get_current_time": {
			Name: "get_current_time",
			Parameters: []*Parameter{
				{Name: "now~", Type: "@wallClock.Datetime = @wallClock.now()"},
			},
			Results: []*Result{{Type: "@time.PlainDateTime!Error"}},
		},
		"get_current_time_formatted": {
			Name: "get_current_time_formatted",
			Parameters: []*Parameter{
				{Name: "now~", Type: "@wallClock.Datetime = @wallClock.now()"},
			},
			Results: []*Result{{Type: "String!Error"}},
		},
		"get_full_name": {
			Name: "get_full_name",
			Parameters: []*Parameter{
				{Name: "first_name", Type: "String"},
				{Name: "last_name", Type: "String"},
			},
			Results: []*Result{{Type: "String"}},
		},
		"get_name_and_age": {
			Name:    "get_name_and_age",
			Results: []*Result{{Type: "(String, Int)"}},
		},
		"get_people": {
			Name:    "get_people",
			Results: []*Result{{Type: "Array[@simple.Person]"}},
		},
		"get_person": {
			Name:    "get_person",
			Results: []*Result{{Type: "@simple.Person"}},
		},
		"get_random_person": {
			Name:    "get_random_person",
			Results: []*Result{{Type: "@simple.Person"}},
		},
		"log_message": {
			Name: "log_message",
			Parameters: []*Parameter{
				{Name: "message", Type: "String"}},
		},
		"say_hello": {
			Name: "say_hello",
			Parameters: []*Parameter{
				{Name: "name~", Type: "String? = None"},
			},
			Results: []*Result{{Type: "String"}},
		},
		"test_abort": {
			Name: "test_abort",
		},
		"test_alternative_error": {
			Name: "test_alternative_error",
			Parameters: []*Parameter{
				{Name: "input", Type: "String"},
			},
			Results: []*Result{{Type: "String"}},
		},
		"test_exit": {
			Name: "test_exit",
		},
		"test_logging": {
			Name: "test_logging",
		},
		"test_normal_error": {
			Name: "test_normal_error",
			Parameters: []*Parameter{
				{Name: "input", Type: "String"},
			},
			Results: []*Result{{Type: "String!Error"}},
		},
	},
}
*/
