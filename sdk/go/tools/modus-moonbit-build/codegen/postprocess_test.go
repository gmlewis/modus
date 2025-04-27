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

type postProcessDiffs struct {
	wantPostProcessBody   string
	gotPostProcessBody    string
	wantPostProcessHeader string
	gotPostProcessHeader  string
}

func postProcessTestSetup(t *testing.T, config *config.Config) *metadata.Metadata {
	t.Helper()

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
		t.Fatalf("CollectModuleInfo failed: %v", err)
	}

	meta, err := metagen.GenerateMetadata(config, mod)
	if err != nil {
		t.Fatalf("GenerateMetadata returned an error: %v", err)
	}

	return meta
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
