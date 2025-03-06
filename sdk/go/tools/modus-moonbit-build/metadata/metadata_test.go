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
		{name: "add", want: "(x : Int, y : Int) -> Int"},
		{name: "add3", want: "(a : Int, b : Int, c~ : Int) -> Int"},
		{name: "add_n", want: "(args : Array[Int]) -> Int"},
		{name: "get_current_time", want: "() -> @time.ZonedDateTime!Error"},
		{name: "get_current_time_formatted", want: "() -> String!Error"},
		{name: "get_full_name", want: "(first_name : String, last_name : String) -> String"},
		{name: "get_people", want: "() -> Array[Person]"},
		{name: "get_person", want: "() -> Person"},
		{name: "get_random_person", want: "() -> Person"},
		{name: "log_message", want: "(message : String) -> Unit"},
		{name: "test_abort", want: "() -> Unit"},
		{name: "test_alternative_error", want: "(input : String) -> String"},
		{name: "test_exit", want: "() -> Unit"},
		{name: "test_logging", want: "() -> Unit"},
		{name: "test_normal_error", want: "(input : String) -> String!Error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := meta.FnExports[tt.name]
			if f == nil {
				t.Fatalf("meta.FnExports missing tt.name=%q", tt.name)
			}

			got := f.String(meta)
			if got != tt.want {
				t.Errorf("function[%q].String = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
