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

func TestAllTestDataIsIdentical(t *testing.T) {
	t.Parallel()

	compareTwo(t, codegenTestData, metagenTestData)
	compareTwo(t, codegenTestData, packagesTestData)
}

func compareTwo(t *testing.T, a, b *testData) {
	t.Helper()

	if diff := cmp.Diff(a.simpleExample, b.simpleExample); diff != "" {
		t.Errorf("%v/testdata/simple-example.mbt != %v/testdata/simple-example.mbt:\n%v", a.name, b.name, diff)
	}

	if diff := cmp.Diff(a.moonModJSON, b.moonModJSON); diff != "" {
		t.Errorf("%v/testdata/moon.mod.json != %v/testdata/moon.mod.json:\n%v", a.name, b.name, diff)
	}

	if diff := cmp.Diff(a.moonPkgJSON, b.moonPkgJSON); diff != "" {
		t.Errorf("%v/testdata/moon.pkg.json != %v/testdata/moon.pkg.json:\n%v", a.name, b.name, diff)
	}
}

type testData struct {
	name          string
	simpleExample string
	moonModJSON   string
	moonPkgJSON   string
}

var codegenTestData = &testData{
	name:          "codegen",
	simpleExample: codegenSimpleExample,
	moonModJSON:   codegenMoonModJSON,
	moonPkgJSON:   codegenMoonPkgJSON,
}

var metagenTestData = &testData{
	name:          "metagen",
	simpleExample: metagenSimpleExample,
	moonModJSON:   metagenMoonModJSON,
	moonPkgJSON:   metagenMoonPkgJSON,
}

var packagesTestData = &testData{
	name:          "packages",
	simpleExample: packagesSimpleExample,
	moonModJSON:   packagesMoonModJSON,
	moonPkgJSON:   packagesMoonPkgJSON,
}

//go:embed codegen/testdata/simple-example.mbt
var codegenSimpleExample string

//go:embed codegen/testdata/moon.mod.json
var codegenMoonModJSON string

//go:embed codegen/testdata/moon.pkg.json
var codegenMoonPkgJSON string

//go:embed metagen/testdata/simple-example.mbt
var metagenSimpleExample string

//go:embed metagen/testdata/moon.mod.json
var metagenMoonModJSON string

//go:embed metagen/testdata/moon.pkg.json
var metagenMoonPkgJSON string

//go:embed packages/testdata/simple-example.mbt
var packagesSimpleExample string

//go:embed packages/testdata/moon.mod.json
var packagesMoonModJSON string

//go:embed packages/testdata/moon.pkg.json
var packagesMoonPkgJSON string
