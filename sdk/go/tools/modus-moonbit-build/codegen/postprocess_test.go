/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package codegen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type postProcessDiffs struct {
	wantPostProcessBody   string
	gotPostProcessBody    string
	wantPostProcessHeader string
	gotPostProcessHeader  string
}

func reportPostProcessDiffs(t *testing.T, name string, wg *postProcessDiffs) {
	t.Helper()

	if diff := cmp.Diff(wg.wantPostProcessBody, wg.gotPostProcessBody); diff != "" {
		t.Log(wg.gotPostProcessBody)
		t.Errorf("%v body mismatch (-want +got):\n%v", name, diff)
	}

	if diff := cmp.Diff(wg.wantPostProcessHeader, wg.gotPostProcessHeader); diff != "" {
		t.Log(wg.gotPostProcessHeader)
		t.Errorf("%v header mismatch (-want +got):\n%v", name, diff)
	}
}
