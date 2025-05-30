/*
 * The code in this file originates from https://github.com/coder/hnsw
 * and is licensed under the terms of the Creative Commons Zero v1.0 Universal license
 * See the LICENSE file in the "hnsw" directory that accompanied this code for further details.
 * See also: https://github.com/coder/hnsw/blob/main/LICENSE
 *
 * SPDX-FileCopyrightText: Ammar Bandukwala <ammar@ammar.io>
 * SPDX-License-Identifier: CC0-1.0
 */

package hnsw

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEuclideanDistance(t *testing.T) {
	a := []float32{1, 2, 3}
	b := []float32{4, 5, 6}
	expected := float32(5.196152)
	actual, _ := EuclideanDistance(a, b)
	require.Equal(t, expected, actual)
}

func TestCosineSimilarity(t *testing.T) {
	var a, b []float32
	// Same magnitude, same direction.
	a = []float32{1, 1, 1}
	b = []float32{0.8, 0.8, 0.8}
	distance, _ := CosineDistance(a, b)
	require.InDelta(t, 0, distance, 0.000001)

	// Perpendicular vectors.
	a = []float32{1, 0}
	b = []float32{0, 1}
	distance, _ = CosineDistance(a, b)
	require.InDelta(t, 1, distance, 0.000001)

	// Equivalent vectors.
	a = []float32{1, 0}
	b = []float32{1, 0}
	distance, _ = CosineDistance(a, b)
	require.InDelta(t, 0, distance, 0.000001)
}

func BenchmarkCosineSimilarity(b *testing.B) {
	v1 := randFloats(1536)
	v2 := randFloats(1536)

	for b.Loop() {
		_, err := CosineDistance(v1, v2)
		if err != nil {
			b.Fatal(err)
		}
	}
}
