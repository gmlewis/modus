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
	meta := setupTestConfig(t, "testdata/time-example")
	removeExternalFuncsForComparison(t, meta)

	if got, want := meta.Plugin, "time-example"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@time-example"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	// if got, want := meta.SDK, "modus-sdk-mbt@40.11.0"; got != want {
	// 	t.Errorf("meta.SDK = %q, want %q", got, want)
	// }

	if diff := cmp.Diff(wantTimeFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantTimeFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	diffMetaTypes(t, wantTimeTypes, meta.Types)
	// if diff := cmp.Diff(wantTimeTypes, meta.Types); diff != "" {
	// 	t.Errorf("meta.Types mismatch (-want +got):\n%v", diff)
	// }

	// This call makes it easy to step through the code with a debugger:
	// LogToConsole(meta)
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
	"(Int, Int, Int)": {
		Name:   "(Int, Int, Int)",
		Fields: []*metadata.Field{{Name: "0", Type: "Int"}, {Name: "1", Type: "Int"}, {Name: "2", Type: "Int"}},
	},
	"(String)":                  {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
	"@ffi.XExternByteArray":     {Name: "@ffi.XExternByteArray"},
	"@ffi.XExternString":        {Name: "@ffi.XExternString"},
	"@ffi.XExternStringArray":   {Name: "@ffi.XExternStringArray"},
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
	"ArrayView[Byte]":           {Name: "ArrayView[Byte]"},
	"Array[Byte]":               {Name: "Array[Byte]"},
	"Array[String]":             {Name: "Array[String]"},
	"Bool":                      {Name: "Bool"},
	"Byte":                      {Name: "Byte"},
	"Bytes":                     {Name: "Bytes"},
	"Bytes!Error":               {Name: "Bytes!Error"},
	"FixedArray[Byte]":          {Name: "FixedArray[Byte]"},
	"Int":                       {Name: "Int"},
	"Int64":                     {Name: "Int64"},
	"Iter[Byte]":                {Name: "Iter[Byte]"},
	"Iter[Char]":                {Name: "Iter[Char]"},
	"Map[String, String]":       {Name: "Map[String, String]"},
	"Result[UInt64, UInt]":      {Name: "Result[UInt64, UInt]"},
	"Result[Unit, UInt]":        {Name: "Result[Unit, UInt]"},
	"String":                    {Name: "String"},
	"String!Error":              {Name: "String!Error"},
	"TimeZoneInfo": {
		Name: "TimeZoneInfo",
		Fields: []*metadata.Field{
			{Name: "standard_name", Type: "String"},
			{Name: "standard_offset", Type: "String"},
			{Name: "daylight_name", Type: "String"},
			{Name: "daylight_offset", Type: "String"},
		},
	},
	"TimeZoneInfo!Error": {
		Name: "TimeZoneInfo!Error",
		Fields: []*metadata.Field{
			{Name: "standard_name", Type: "String"},
			{Name: "standard_offset", Type: "String"},
			{Name: "daylight_name", Type: "String"},
			{Name: "daylight_offset", Type: "String"},
		},
	},
	"UInt":   {Name: "UInt"},
	"UInt64": {Name: "UInt64"},
}
