/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRegexp(t *testing.T) {
	tests := []struct {
		name    string
		reg     json.RawMessage
		wantReg string
		got     json.RawMessage
		want    bool
	}{
		{
			name:    "local_time",
			reg:     json.RawMessage(`"^\"\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}[-+]\\d{2}:\\d{2}\\[.*\\]\"$"`),
			wantReg: `^"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}\[.*\]"$`,
			got:     json.RawMessage(`"2025-02-05T21:51:16-05:00[EST]"`),
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Endpoint{Name: tt.name, Regexp: tt.reg}
			r, rs, err := e.RegexpBody()
			if err != nil || rs == "" {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.wantReg, rs); diff != "" {
				t.Errorf("rs mismatch (-want +got):\n%v", diff)
			}

			result, err := json.Marshal(tt.got)
			if err != nil {
				t.Fatal(err)
			}

			got := r.MatchString(string(result))
			if got != tt.want {
				t.Errorf("MatchString('%v') = %v, want %v", string(tt.got), got, tt.want)
			}
		})
	}
}
