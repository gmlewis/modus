// -*- compile-command: "NO_COLOR=1 go test -timeout 30s -tags integration -run ^TestDebugArrayBool1$ ."; -*-

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
	"log"
	"testing"
)

func TestDebugArrayBool1(t *testing.T) {
	t.Skip("TODO: fix this")
	result, err := fixture.CallFunction(t, "test_debug_array_bool_1", []bool{true})
	if err != nil {
		t.Fatal(err)
	}

	// The debug output should show us exactly what's happening
	log.Printf("TestDebugArrayBool1: result=%#v", result)
}
