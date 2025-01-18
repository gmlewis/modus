/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

//go:embed packages/testdata/simple-example.mbt
var packagesSimpleExample string

//go:embed packages/testdata/moon.mod.json
var packagesMoonModJSON string

//go:embed packages/testdata/moon.pkg.json
var packagesMoonPkgJSON string

//go:embed metagen/testdata/simple-example.mbt
var metagenSimpleExample string

//go:embed metagen/testdata/moon.mod.json
var metagenMoonModJSON string

//go:embed metagen/testdata/moon.pkg.json
var metagenMoonPkgJSON string

func TestAllTestDataIdentical(t *testing.T) {
	t.Parallel()

	if diff := cmp.Diff(packagesSimpleExample, metagenSimpleExample); diff != "" {
		t.Errorf("packages/testdata/simple-example.mbt != metagen/testdata/simple-example.mbt:\n%v", diff)
	}

	if diff := cmp.Diff(packagesMoonModJSON, metagenMoonModJSON); diff != "" {
		t.Errorf("packages/testdata/moon.mod.json != metagen/testdata/moon.mod.json:\n%v", diff)
	}

	if diff := cmp.Diff(packagesMoonPkgJSON, metagenMoonPkgJSON); diff != "" {
		t.Errorf("packages/testdata/moon.pkg.json != metagen/testdata/moon.pkg.json:\n%v", diff)
	}
}
