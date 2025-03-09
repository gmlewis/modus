/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import "testing"

func TestFullyQualifyTypeName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Array[Person]",
			want: "Array[@pkg.Person]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FullyQualifyTypeName("@pkg", tt.name)
			if got != tt.want {
				t.Errorf("FullyQualifyTypeName = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
