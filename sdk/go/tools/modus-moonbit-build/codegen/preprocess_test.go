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

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
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

/*
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
	pkg, err := getMainPackage(config.SourceDir, mod)
		t.Log(body.String())
		t.Log(header.String())
		t.Log(moonPkgJSON.String())
}
*/

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
