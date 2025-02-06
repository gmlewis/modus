/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package packages

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplitFunctionParameters(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "simple",
			input: "a : Int, b : Bool, c : String",
			want:  []string{"a : Int", "b : Bool", "c : String"},
		},
		{
			name:  "simple map",
			input: "a : Array[String?]?, b : Map[String, String]",
			want:  []string{"a : Array[String?]?", "b : Map[String, String]"},
		},
		{
			name:  "complex map",
			input: "a : Map[String, Map[Int, Bool]?], b : Map[Char, Array[Int, Map[String, Bool]]]",
			want:  []string{"a : Map[String, Map[Int, Bool]?]", "b : Map[Char, Array[Int, Map[String, Bool]]]"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := splitParamsWithBrackets(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("splitFunctionParameters() mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
