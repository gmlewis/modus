//go:build !wasip1

/*
 * This example is part of the Modus project, licensed under the Apache License 2.0.
 * You may modify and use this example in accordance with the license.
 * See the LICENSE file that accompanied this code for further details.
 */

package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHelloArrayOfInts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		nilSlice bool
		numInts  int
		want     []int
	}{
		{
			name:     "nil slice",
			nilSlice: true,
			numInts:  5,
			want:     nil,
		},
		{
			name:     "empty slice",
			nilSlice: false,
			numInts:  0,
			want:     []int{},
		},
		{
			name:     "non-empty slice",
			nilSlice: false,
			numInts:  3,
			want:     []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := HelloArrayOfInts(tt.nilSlice, tt.numInts)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("HelloArrayOfInts() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHelloOptionString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		nilString   bool
		emptyString bool
		expected    *string
	}{
		{
			name:      "nil string",
			nilString: true,
			expected:  nil,
		},
		{
			name:        "empty string",
			emptyString: true,
			expected:    ptr(""),
		},
		{
			name:     "default string",
			expected: ptr("Hello, World!"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := HelloOptionString(tt.nilString, tt.emptyString)
			if (result == nil && tt.expected != nil) || (result != nil && tt.expected == nil) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			} else if result != nil && *result != *tt.expected {
				t.Errorf("expected %v, got %v", *tt.expected, *result)
			}
		})
	}
}
