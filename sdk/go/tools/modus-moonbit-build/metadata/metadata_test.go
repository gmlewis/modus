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
	"encoding/json"
	"fmt"
	"testing"
)

type functionStringTest struct {
	name string
	want string
}

func testFunctionStringHelper(t *testing.T, name string, metadataJSON []byte, tests []functionStringTest) {
	t.Helper()

	var meta *Metadata
	if err := json.Unmarshal(metadataJSON, &meta); err != nil || meta == nil || meta.FnExports == nil {
		t.Fatalf("%v json.Unmarshal: %v", name, err)
	}

	if len(tests) != len(meta.FnExports) {
		for k, v := range meta.FnExports {
			fmt.Printf("    {name: %q, want: %q},\n", k, v.String(meta))
		}
		t.Fatalf("%v FnExports length mismatch: got %v, want %v", name, len(meta.FnExports), len(tests))
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
				t.Errorf("%v function.String = %q, want %q", name, got, tt.want)
			}
		})
	}

}
