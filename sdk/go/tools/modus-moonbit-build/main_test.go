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
	"log"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metagen"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
	"github.com/google/go-cmp/cmp"
)

//go:embed metadata/testdata/simple-example-metadata.json
var simpleExampleMetadataJSON []byte

func TestAllTestDataIsIdentical(t *testing.T) {
	t.Parallel()

	compareTwoTestSetups(t, codegenTestData, metagenTestData)
	compareTwoTestSetups(t, codegenTestData, packagesTestData)

	// Also test that the generated JSON in the metadata/testdata dir is up-to-date:
	var metaJSON *metadata.Metadata
	if err := json.Unmarshal(simpleExampleMetadataJSON, &metaJSON); err != nil || metaJSON == nil || metaJSON.FnExports == nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	// This JSON format strips out the `Name` fields, so fill them back in to match the live version:
	for name, fn := range metaJSON.FnExports {
		fn.Name = name
	}
	for name, fn := range metaJSON.FnImports {
		fn.Name = name
	}
	for name, typ := range metaJSON.Types {
		typ.Name = name
	}
	if metaJSON.FnImports == nil {
		metaJSON.FnImports = metadata.FunctionMap{}
	}

	metaLive := genMetadata("metagen/testdata/simple-example")
	// config := &config.Config{
	// 	SourceDir:
	// }
	// mod := &modinfo.ModuleInfo{
	// 	ModulePath:      "github.com/gmlewis/modus/examples/simple-example",
	// 	ModusSDKVersion: version.Must(version.NewVersion("40.11.0")),
	// }
	//
	// metaLive, err := metagen.GenerateMetadata(config, mod)
	// if err != nil {
	// 	t.Fatalf("GenerateMetadata returned an error: %v", err)
	// }

	// Make the "BuildID" and "BuildTime" fields match the JSON version:
	metaLive.BuildID = metaJSON.BuildID
	metaLive.BuildTime = metaJSON.BuildTime
	metaLive.GitCommit = metaJSON.GitCommit

	if diff := cmp.Diff(metaJSON, metaLive); diff != "" {
		t.Errorf("metadata/testdata/simple-example-metadata.json != live metadata:\n%v", diff)
	}
}

func compareTwoTestSetups(t *testing.T, a, b *testData) {
	t.Helper()

	if diff := cmp.Diff(a.simpleExample, b.simpleExample); diff != "" {
		t.Errorf("%v/testdata/simple-example.mbt != %v/testdata/simple-example.mbt:\n%v", a.name, b.name, diff)
	}

	// With the incorporation of parsing all MoonBit source code dependencies,
	// the path to "gmlewis/modus" must be accurate, so the relative paths differ.
	// if diff := cmp.Diff(a.moonModJSON, b.moonModJSON); diff != "" {
	// 	t.Errorf("%v/testdata/moon.mod.json != %v/testdata/moon.mod.json:\n%v", a.name, b.name, diff)
	// }

	if diff := cmp.Diff(a.moonPkgJSON, b.moonPkgJSON); diff != "" {
		t.Errorf("%v/testdata/moon.pkg.json != %v/testdata/moon.pkg.json:\n%v", a.name, b.name, diff)
	}
}

// genMetadata is used during 'go generate' and also during testing.
// To avoid cyclic imports, multiple copies of this function exist. :-(
func genMetadata(sourceDir string) *metadata.Metadata {
	config := &config.Config{
		SourceDir: sourceDir,
	}

	// Make sure the ".mooncakes" directory is initialized before running the test.
	mooncakesDir := path.Join(config.SourceDir, ".mooncakes")
	if _, err := os.Stat(mooncakesDir); err != nil {
		// run the "moon check" command in that directory to initialize it.
		args := []string{"moon", "check", "--directory", config.SourceDir}
		buf, err := exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			log.Fatalf("error running %q: %v\n%s", args, err, buf)
		}
	}

	mod, err := modinfo.CollectModuleInfo(config)
	if err != nil {
		log.Fatalf("CollectModuleInfo returned an error: %v", err)
	}

	meta, err := metagen.GenerateMetadata(config, mod)
	if err != nil {
		log.Fatalf("GenerateMetadata returned an error: %v", err)
	}

	return meta
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

//go:embed codegen/testdata/simple-example/simple-example.mbt
var codegenSimpleExample string

//go:embed codegen/testdata/simple-example/moon.mod.json
var codegenMoonModJSON string

//go:embed codegen/testdata/simple-example/moon.pkg.json
var codegenMoonPkgJSON string

//go:embed metagen/testdata/simple-example/simple-example.mbt
var metagenSimpleExample string

//go:embed metagen/testdata/simple-example/moon.mod.json
var metagenMoonModJSON string

//go:embed metagen/testdata/simple-example/moon.pkg.json
var metagenMoonPkgJSON string

//go:embed packages/testdata/simple-example/simple-example.mbt
var packagesSimpleExample string

//go:embed packages/testdata/simple-example/moon.mod.json
var packagesMoonModJSON string

//go:embed packages/testdata/simple-example/moon.pkg.json
var packagesMoonPkgJSON string
