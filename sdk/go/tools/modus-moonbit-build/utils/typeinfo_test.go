/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFullyQualifyTypeName(t *testing.T) {
	tests := []struct {
		name string
		want string
		acc  map[string]struct{}
	}{
		{
			name: "Array[Person]",
			want: "Array[@pkg.Person]",
			acc: map[string]struct{}{
				"@pkg.Person":        {},
				"Array[@pkg.Person]": {},
			},
		},
		{
			name: "Array[(String, String)]",
			want: "Array[(String, String)]",
			acc: map[string]struct{}{
				"(String, String)":        {},
				"Array[(String, String)]": {},
				"String":                  {},
			},
		},
		{
			name: "Array[(String, (Byte, Person?))]",
			want: "Array[(String, (Byte, @pkg.Person?))]",
			acc: map[string]struct{}{
				"(Byte, @pkg.Person?)":                  {},
				"(String, (Byte, @pkg.Person?))":        {},
				"@pkg.Person?":                          {},
				"Array[(String, (Byte, @pkg.Person?))]": {},
				"Byte":                                  {},
				"String":                                {},
			},
		},
		{
			name: "Array[(String, (Byte, @http.Person?))]",
			want: "Array[(String, (Byte, @http.Person?))]",
			acc: map[string]struct{}{
				"(Byte, @http.Person?)":                  {},
				"(String, (Byte, @http.Person?))":        {},
				"@http.Person?":                          {},
				"Array[(String, (Byte, @http.Person?))]": {},
				"Byte":                                   {},
				"String":                                 {},
			},
		},
		{
			// A `Map[K, V]` also needs to add `Array[K]` and `Array[V]`.
			name: "Map[K, V]",
			want: "Map[@pkg.K, @pkg.V]",
			acc: map[string]struct{}{
				"@pkg.K":              {},
				"@pkg.V":              {},
				"Array[@pkg.K]":       {},
				"Array[@pkg.V]":       {},
				"Map[@pkg.K, @pkg.V]": {},
			},
		},
		{
			name: "Map[@http.K, @http.V]",
			want: "Map[@http.K, @http.V]",
			acc: map[string]struct{}{
				"@http.K":               {},
				"@http.V":               {},
				"Array[@http.K]":        {},
				"Array[@http.V]":        {},
				"Map[@http.K, @http.V]": {},
			},
		},
		{
			name: "Map[String, Bool]",
			want: "Map[String, Bool]",
			acc: map[string]struct{}{
				"Array[Bool]":       {},
				"Array[String]":     {},
				"Bool":              {},
				"Map[String, Bool]": {},
				"String":            {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, acc := FullyQualifyTypeName("@pkg", tt.name)
			if got != tt.want {
				t.Errorf("FullyQualifyTypeName = '%v', want '%v'", got, tt.want)
			}
			if diff := cmp.Diff(tt.acc, acc); diff != "" {
				t.Errorf("FullyQualifyTypeName accumulator mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
