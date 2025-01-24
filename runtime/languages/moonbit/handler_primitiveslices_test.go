/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package moonbit

/*
func TestDoReadSlice(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		data []byte
		want any
	}{
		{
			name: "static array of 0 ints",
			data: []byte{112, 47, 0, 0, 0, 0, 0}, // also: [64 45 0 0 0 0 0]
			want: []int{},
		},
		{
			name: "static array of 1 int: [333]",
			data: []byte{112, 47, 0, 0, 1, 0, 0}, // also: [64 45 0 0 1 0 0]
			want: []int{333},
		},
		{
			name: "static array of 2 ints: [333, 444]",
			data: []byte{112, 47, 0, 0, 2, 0, 0}, // also: [64 45 0 0 2 0 0]
		},
		{
			name: "static array of 3 ints: [333, 444, 555]",
			data: []byte{112, 47, 0, 0, 3, 0, 0}, // also: [64 45 0 0 3 0 0]
		},
		// 112 decimal is 70 hex
		// 64 decimal is 40 hex
		// 47 decimal is 2F hex
		// 45 decimal is 2D hex
		// static array of 2 ints: [333, 444]
		// also: [64 45 0 0 2 0 0]
		// static array of 3 ints: [333, 444, 555]
		// also: [64 45 0 0 3 0 0]
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := doReadString(tt.data)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("doReadString() mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
*/
