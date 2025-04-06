/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package moonbit_test

import (
	"strings"
	"testing"
)

// Ptr is a helper routine that allocates a new T value
// to store v and returns a pointer to it.
func Ptr[T any](v T) *T {
	return &v
}

// testInputSide is a helper function that calls the corresponding
// input function with the expected value.
func testInputSide(t *testing.T, fnName string, expected any) {
	t.Helper()
	fnName = strings.Replace(fnName, "_output_", "_input_", 1)
	if _, err := fixture.CallFunction(t, fnName, expected); err != nil {
		t.Error(err)
	}
}
