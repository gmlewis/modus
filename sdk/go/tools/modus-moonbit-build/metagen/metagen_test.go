// -*- compile-command: "go test ./..."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package metagen

import (
	"log"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
	"github.com/google/go-cmp/cmp"
)

// AnyPtr is a helper routine that allocates a new T value
// to store v and returns a pointer to it.
func AnyPtr(v any) *any {
	return &v
}

func setupTestConfig(t *testing.T, sourceDir string) *metadata.Metadata {
	t.Helper()
	log.SetFlags(0)
	return genMetadata(sourceDir)
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

	meta, err := GenerateMetadata(config, mod)
	if err != nil {
		log.Fatalf("GenerateMetadata returned an error: %v", err)
	}

	return meta
}

func diffMetaTypes(t *testing.T, wantMetaTypes, gotMetaTypes metadata.TypeMap) {
	t.Helper()
	// zero out all the IDs because they are horribly distracting in the diffs
	for k := range wantMetaTypes {
		wantMetaTypes[k].Id = 0
	}
	for k := range gotMetaTypes {
		gotMetaTypes[k].Id = 0
	}
	if diff := cmp.Diff(wantMetaTypes, gotMetaTypes); diff != "" {
		t.Errorf("meta.Types mismatch (-want +got):\n%v", diff)
	}
}
