/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package wasm

import (
	"encoding/json"
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/codegen"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/compiler"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metagen"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"

	"github.com/google/go-cmp/cmp"
)

func testFilterMetadataHelper(t *testing.T, config *config.Config, wantBeforeFilter, wantAfterFilter *metadata.Metadata) {
	t.Helper()
	config.CompilerPath = "moon"

	mod, err := modinfo.CollectModuleInfo(config)
	if err != nil {
		t.Fatal(err)
	}
	meta, err := metagen.GenerateMetadata(config, mod)
	if err != nil {
		t.Fatal(err)
	}

	if config.WasmFileName != "testdata.wasm" {
		// Need to build the wasm file
		if err := codegen.PostProcess(config, meta); err != nil {
			t.Fatal(err)
		}
		if err := compiler.Compile(config); err != nil {
			t.Fatal(err)
		}
	}

	// zero out fields that change regularly
	meta.SDK = ""
	meta.BuildID = ""
	meta.BuildTime = ""
	meta.GitRepo = ""
	meta.GitCommit = ""

	// zero out all the IDs because they are horribly distracting in the diffs
	for k := range meta.Types {
		meta.Types[k].Id = 0
	}

	if diff := cmp.Diff(wantBeforeFilter, meta); diff != "" {
		t.Fatalf("GenerateMetadata meta BEFORE filter mismatch (-want +got):\n%v", diff)
	}

	if err := FilterMetadata(config, meta); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(wantAfterFilter, meta); diff != "" {
		t.Fatalf("GenerateMetadata meta AFTER filter mismatch (-want +got):\n%v", diff)
	}
}

func deepCopyMetadata(t *testing.T, m *metadata.Metadata) *metadata.Metadata {
	t.Helper()
	buf, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	var deepCopy *metadata.Metadata
	if err := json.Unmarshal(buf, &deepCopy); err != nil {
		t.Fatal(err)
	}
	for name, v := range deepCopy.FnExports {
		v.Name = name
	}
	for name, v := range deepCopy.FnImports {
		v.Name = name
	}
	for name, v := range deepCopy.Types {
		v.Name = name
	}
	return deepCopy
}
