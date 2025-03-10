// -*- compile-command: "go generate ./..."; -*-
//go:build ignore

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// gen-metadata-testdata.go is used to marshal testdata actual
// parsing of the MoonBit code so that all testing stays in sync.
// It is needed to prevent an import cycle, but is also useful
// to test the JSON serialization.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metagen"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
	"github.com/hashicorp/go-version"
)

var (
	genTestCases = []string{
		"neo4j-example",
		"runtime-testdata",
		"simple-example",
		"test-suite",
		"time-example",
	}
)

func main() {
	log.SetFlags(0)
	_, thisFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(thisFile)

	for _, name := range genTestCases {
		genJSON(name, baseDir)
	}

	log.Printf("Done.")
}

func genJSON(name, baseDir string) {
	filename := fmt.Sprintf("testdata/%v-metadata.json", name)
	log.Printf("gen-metadata-testdata.go: Generating %v ...", filename)
	sourceDir := filepath.Join(baseDir, "../testdata", name)

	config := &config.Config{
		SourceDir: sourceDir,
	}
	mod := &modinfo.ModuleInfo{
		ModulePath:      fmt.Sprintf("github.com/gmlewis/modus/examples/%v", name),
		ModusSDKVersion: version.Must(version.NewVersion("0.16.5")),
	}

	meta, err := metagen.GenerateMetadata(config, mod)
	if err != nil {
		log.Fatalf("GenerateMetadata returned an error: %v", err)
	}

	buf, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		log.Fatalf("json.MarshalIndent: %v\n%s", err, buf)
	}

	if err := os.WriteFile(filename, buf, 0644); err != nil {
		log.Fatalf("os.WriteFile: %v", err)
	}
}
