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
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/go-version"
)

func TestGenerateMetadata_Time(t *testing.T) {
	config := &config.Config{
		SourceDir: "testdata/time-example",
	}
	mod := &modinfo.ModuleInfo{
		ModulePath:      "github.com/gmlewis/modus/examples/time-example",
		ModusSDKVersion: version.Must(version.NewVersion("40.11.0")),
	}

	meta, err := GenerateMetadata(config, mod)
	if err != nil {
		t.Fatalf("GenerateMetadata returned an error: %v", err)
	}

	if got, want := meta.Plugin, "time-example"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@time-example"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if got, want := meta.SDK, "modus-sdk-mbt@40.11.0"; got != want {
		t.Errorf("meta.SDK = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantTimeFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantTimeFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantTimeTypes, meta.Types); diff != "" {
		t.Errorf("meta.Types mismatch (-want +got):\n%v", diff)
	}
}

var wantTimeFnExports = metadata.FunctionMap{
	"get_local_time": {
		Name:    "get_local_time",
		Results: []*metadata.Result{{Type: "String!Error"}},
	},
	"get_local_time_zone": {
		Name:    "get_local_time_zone",
		Results: []*metadata.Result{{Type: "String"}},
	},
	"get_time_in_zone": {
		Name:       "get_time_in_zone",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
	},
	"get_utc_time": {
		Name:    "get_utc_time",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime!Error"}},
	},
}

var wantTimeFnImports = metadata.FunctionMap{}

var wantTimeTypes = metadata.TypeMap{
	"@time.ZonedDateTime": {Id: 4, Name: "@time.ZonedDateTime"},
	"String":              {Id: 5, Name: "String"},
}
