// -*- compile-command: "go test -run ^TestGenerateMetadata_HelloPrimitives$ ."; -*-

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

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateMetadata_HelloPrimitives(t *testing.T) {
	meta := setupTestConfig(t, "../testdata/hello-primitives")

	if got, want := meta.Plugin, "hello-primitives"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@hello-primitives"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantHelloPrimitivesFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantHelloPrimitivesFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	diffMetaTypes(t, wantHelloPrimitivesTypes, meta.Types)
}

var wantHelloPrimitivesFnExports = metadata.FunctionMap{
	"hello_primitive_time_duration_max": {
		Name:    "hello_primitive_time_duration_max",
		Results: []*metadata.Result{{Type: "@time.Duration"}},
	},
	"hello_primitive_time_duration_min": {
		Name:    "hello_primitive_time_duration_min",
		Results: []*metadata.Result{{Type: "@time.Duration"}},
	},
}

var wantHelloPrimitivesFnImports = metadata.FunctionMap{
	"modus_system.logMessage": {
		Name:       "modus_system.logMessage",
		Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
	},
}

var wantHelloPrimitivesTypes = metadata.TypeMap{
	"(String)":                  {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
	"@time.Duration":            {Name: "@time.Duration"},
	"@time.Duration!Error":      {Name: "@time.Duration!Error"},
	"@time.Period":              {Name: "@time.Period"},
	"@time.Period!Error":        {Name: "@time.Period!Error"},
	"@time.PlainDate":           {Name: "@time.PlainDate"},
	"@time.PlainDate!Error":     {Name: "@time.PlainDate!Error"},
	"@time.PlainDateTime":       {Name: "@time.PlainDateTime"},
	"@time.PlainDateTime!Error": {Name: "@time.PlainDateTime!Error"},
	"@time.PlainTime":           {Name: "@time.PlainTime"},
	"@time.PlainTime!Error":     {Name: "@time.PlainTime!Error"},
	"@time.Weekday":             {Name: "@time.Weekday"},
	"@time.Zone":                {Name: "@time.Zone"},
	"@time.Zone!Error":          {Name: "@time.Zone!Error"},
	"@time.ZoneOffset":          {Name: "@time.ZoneOffset"},
	"@time.ZoneOffset!Error":    {Name: "@time.ZoneOffset!Error"},
	"@time.ZonedDateTime":       {Name: "@time.ZonedDateTime"},
	"@time.ZonedDateTime!Error": {Name: "@time.ZonedDateTime!Error"},
	"Array[Byte]":               {Name: "Array[Byte]"},
	"Bool":                      {Name: "Bool"},
	"Byte":                      {Name: "Byte"},
	"FixedArray[Byte]":          {Name: "FixedArray[Byte]"},
	"Int":                       {Name: "Int"},
	"Int64":                     {Name: "Int64"},
}
