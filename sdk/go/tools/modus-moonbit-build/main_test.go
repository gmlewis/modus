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
	"encoding/json"
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metagen"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-version"
)

//go:embed metadata/testdata/simple-example-metadata.json
var simpleExampleMetadataJSON []byte

func TestAllTestDataIsIdentical(t *testing.T) {
	t.Parallel()

	compareTwo(t, codegenTestData, metagenTestData)
	compareTwo(t, codegenTestData, packagesTestData)

	// Also test that the generated JSON in the metadata/testdata dir is up-to-date:
	var metaJSON *metadata.Metadata
	if err := json.Unmarshal(simpleExampleMetadataJSON, &metaJSON); err != nil || metaJSON == nil || metaJSON.FnExports == nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	// This JSON format strips out the `Name` fields, so fill them back in to match the live version:
	for name, fn := range metaJSON.FnExports {
		fn.Name = name
	}
	for name, typ := range metaJSON.Types {
		typ.Name = name
	}
	if metaJSON.FnImports == nil {
		metaJSON.FnImports = metadata.FunctionMap{}
	}

	config := &config.Config{
		SourceDir: "metagen/testdata",
	}
	mod := &modinfo.ModuleInfo{
		ModulePath:      "github.com/gmlewis/modus/examples/simple-example",
		ModusSDKVersion: version.Must(version.NewVersion("40.11.0")),
	}

	metaLive, err := metagen.GenerateMetadata(config, mod)
	if err != nil {
		t.Fatalf("GenerateMetadata returned an error: %v", err)
	}
	// Make the "BuildID" and "BuildTime" fields match the JSON version:
	metaLive.BuildID = metaJSON.BuildID
	metaLive.BuildTime = metaJSON.BuildTime

	if diff := cmp.Diff(metaJSON, metaLive); diff != "" {
		t.Errorf("metadata/testdata/simple-example-metadata.json != live metadata:\n%v", diff)
	}
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
