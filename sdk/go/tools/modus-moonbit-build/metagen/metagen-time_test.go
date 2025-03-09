// -*- compile-command: "go test -run ^TestGenerateMetadata_Time$ ."; -*-

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

func TestGenerateMetadata_Time(t *testing.T) {
	meta := setupTestConfig(t, "../testdata/time-example")

	if got, want := meta.Plugin, "time-example"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@time-example"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if got, want := meta.SDK, "modus-sdk-mbt@0.16.5"; got != want {
		t.Errorf("meta.SDK = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantTimeFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantTimeFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	diffMetaTypes(t, wantTimeTypes, meta.Types)
}

var wantTimeFnExports = metadata.FunctionMap{
	"get_local_time_modus": {
		Name:    "get_local_time_modus",
		Results: []*metadata.Result{{Type: "String"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current local time using the Modus host Go function."}},
	},
	"get_local_time_moonbit": {
		Name:    "get_local_time_moonbit",
		Results: []*metadata.Result{{Type: "String!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current local time using the moonbitlang/x/time package."}},
	},
	"get_local_time_zone_id": {
		Name:    "get_local_time_zone_id",
		Results: []*metadata.Result{{Type: "String"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the local time zone identifier."}},
	},
	"get_time_in_zone_modus": {
		Name:       "get_time_in_zone_modus",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
		Docs: &metadata.Docs{
			Lines: []string{"Returns the current time in a specified time zone using", "the Modus host Go function."},
		},
	},
	"get_time_in_zone_moonbit": {
		Name:       "get_time_in_zone_moonbit",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
		Docs: &metadata.Docs{
			Lines: []string{"Returns the current time in a specified time zone using", "the moonbitlang/x/time package."},
		},
	},
	"get_time_zone_info": {
		Name:       "get_time_zone_info",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "TimeZoneInfo!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Returns some basic information about the time zone specified."}},
	},
	"get_utc_time": {
		Name:    "get_utc_time",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time in UTC."}},
	},
}

var wantTimeFnImports = metadata.FunctionMap{
	"modus_system.getTimeInZone": {
		Name:       "modus_system.getTimeInZone",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"modus_system.getTimeZoneData": {
		Name: "modus_system.getTimeZoneData",
		Parameters: []*metadata.Parameter{
			{Name: "tz", Type: "String"},
			{Name: "format", Type: "String"},
		},
		Results: []*metadata.Result{{Type: "Array[Byte]"}},
	},
	"modus_system.logMessage": {
		Name: "modus_system.logMessage",
		Parameters: []*metadata.Parameter{
			{Name: "level", Type: "String"},
			{Name: "message", Type: "String"},
		},
	},
}

var wantTimeTypes = metadata.TypeMap{
	"(String)": {Id: 4,
		Name:   "(String)",
		Fields: []*metadata.Field{{Name: "0", Type: "String"}},
	},
	"@time.ZonedDateTime":       {Id: 5, Name: "@time.ZonedDateTime"},
	"@time.ZonedDateTime!Error": {Id: 6, Name: "@time.ZonedDateTime!Error"},
	"Array[Byte]":               {Id: 7, Name: "Array[Byte]"},
	"String":                    {Id: 8, Name: "String"},
	"String!Error":              {Id: 9, Name: "String!Error"},
	"TimeZoneInfo": {Id: 10,
		Name: "TimeZoneInfo",
		Fields: []*metadata.Field{
			{Name: "standard_name", Type: "String"},
			{Name: "standard_offset", Type: "String"},
			{Name: "daylight_name", Type: "String"},
			{Name: "daylight_offset", Type: "String"},
		},
	},
	"TimeZoneInfo!Error": {Id: 11,
		Name: "TimeZoneInfo!Error",
		Fields: []*metadata.Field{
			{Name: "standard_name", Type: "String"},
			{Name: "standard_offset", Type: "String"},
			{Name: "daylight_name", Type: "String"},
			{Name: "daylight_offset", Type: "String"},
		},
	},
}
