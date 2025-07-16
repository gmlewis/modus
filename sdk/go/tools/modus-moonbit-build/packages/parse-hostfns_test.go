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

func TestExtractFnDetails(t *testing.T) {
	match := []string{
		`fn _host_echo1(message : Int) -> Int = "modus_test" "echo1"`,
		"message : Int",
		"-> Int ",
		"modus_test",
		"echo1",
	}
	fullSrc := "\n\n" + `
///|modus:import modus_test echo1(String) -> String
fn _host_echo1(message : Int) -> Int = "modus_test" "echo1"
`
	got := extractFnDetails(match, fullSrc)
	want := fnDetailsT{
		ReturnSig:  "String",
		HostEnv:    "modus_test",
		HostFn:     "echo1",
		MethodName: "modus_test.echo1",
		AllArgs:    "message : String",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("extractFnDetails mismatch (-want +got):\n%v", diff)
	}
}
