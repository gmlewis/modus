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

type preProcessDiffs struct {
	wantPreProcessBody        string
	gotPreProcessBody         string
	wantPreProcessHeader      string
	gotPreProcessHeader       string
	wantPreProcessMoonPkgJSON string
	gotPreProcessMoonPkgJSON  string
}

func reportPreProcessDiffs(t *testing.T, name string, wg *preProcessDiffs) {
	t.Helper()

	if diff := cmp.Diff(wg.wantPreProcessBody, wg.gotPreProcessBody); diff != "" {
		t.Log(wg.gotPreProcessBody)
		t.Errorf("%v body mismatch (-want +got):\n%v", name, diff)
	}

	if diff := cmp.Diff(wg.wantPreProcessHeader, wg.gotPreProcessHeader); diff != "" {
		t.Log(wg.gotPreProcessHeader)
		t.Errorf("%v header mismatch (-want +got):\n%v", name, diff)
	}

	if diff := cmp.Diff(wg.wantPreProcessMoonPkgJSON, wg.gotPreProcessMoonPkgJSON); diff != "" {
		t.Log(wg.gotPreProcessMoonPkgJSON)
		t.Errorf("%v moonPkgJSON mismatch (-want +got):\n%v", name, diff)
	}
}
