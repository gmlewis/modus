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
		Results: []*metadata.Result{{Type: "String raise Error"}},
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
		Results:    []*metadata.Result{{Type: "String raise Error"}},
		Docs: &metadata.Docs{
			Lines: []string{"Returns the current time in a specified time zone using", "the moonbitlang/x/time package."},
		},
	},
	"get_time_zone_info": {
		Name:       "get_time_zone_info",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "TimeZoneInfo raise Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Returns some basic information about the time zone specified."}},
	},
	"get_utc_time": {
		Name:    "get_utc_time",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime raise Error"}},
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
	"(String)": {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
	"@testutils.CallStack[T]": {
		Name:   "@testutils.CallStack[T]",
		Fields: []*metadata.Field{{Name: "items", Type: "Array[Array[@testutils.T]]"}},
	},
	"@testutils.T":                    {Name: "@testutils.T"},
	"@time.Duration":                  {Name: "@time.Duration"},
	"@time.Duration raise Error":      {Name: "@time.Duration raise Error"},
	"@time.Period":                    {Name: "@time.Period"},
	"@time.Period raise Error":        {Name: "@time.Period raise Error"},
	"@time.PlainDate":                 {Name: "@time.PlainDate"},
	"@time.PlainDate raise Error":     {Name: "@time.PlainDate raise Error"},
	"@time.PlainDateTime":             {Name: "@time.PlainDateTime"},
	"@time.PlainDateTime raise Error": {Name: "@time.PlainDateTime raise Error"},
	"@time.PlainTime":                 {Name: "@time.PlainTime"},
	"@time.PlainTime raise Error":     {Name: "@time.PlainTime raise Error"},
	"@time.Weekday":                   {Name: "@time.Weekday"},
	"@time.Zone":                      {Name: "@time.Zone"},
	"@time.Zone raise Error":          {Name: "@time.Zone raise Error"},
	"@time.ZoneOffset":                {Name: "@time.ZoneOffset"},
	"@time.ZoneOffset raise Error":    {Name: "@time.ZoneOffset raise Error"},
	"@time.ZonedDateTime":             {Name: "@time.ZonedDateTime"},
	"@time.ZonedDateTime raise Error": {Name: "@time.ZonedDateTime raise Error"},
	"ArrayView[Byte]":                 {Name: "ArrayView[Byte]"},
	"Array[@testutils.T]":             {Name: "Array[@testutils.T]"},
	"Array[Array[@testutils.T]]":      {Name: "Array[Array[@testutils.T]]"},
	"Array[Byte]":                     {Name: "Array[Byte]"},
	"Bool":                            {Name: "Bool"},
	"Byte":                            {Name: "Byte"},
	"Bytes":                           {Name: "Bytes"},
	"Bytes raise Error":               {Name: "Bytes raise Error"},
	"Char":                            {Name: "Char"},
	"Double":                          {Name: "Double"},
	"FixedArray[Byte]":                {Name: "FixedArray[Byte]"},
	"FixedArray[Double]":              {Name: "FixedArray[Double]"},
	"FixedArray[Float]":               {Name: "FixedArray[Float]"},
	"FixedArray[Int64]":               {Name: "FixedArray[Int64]"},
	"FixedArray[Int]":                 {Name: "FixedArray[Int]"},
	"FixedArray[UInt64]":              {Name: "FixedArray[UInt64]"},
	"FixedArray[UInt]":                {Name: "FixedArray[UInt]"},
	"Float":                           {Name: "Float"},
	"Int":                             {Name: "Int"},
	"Int64":                           {Name: "Int64"},
	"Iter[Byte]":                      {Name: "Iter[Byte]"},
	"Iter[Char]":                      {Name: "Iter[Char]"},
	"Result[UInt64, UInt]":            {Name: "Result[UInt64, UInt]"},
	"Result[Unit, UInt]":              {Name: "Result[Unit, UInt]"},
	"String":                          {Name: "String"},
	"String raise Error":              {Name: "String raise Error"},
	"TimeZoneInfo": {
		Name: "TimeZoneInfo",
		Fields: []*metadata.Field{
			{Name: "standard_name", Type: "String"},
			{Name: "standard_offset", Type: "String"},
			{Name: "daylight_name", Type: "String"},
			{Name: "daylight_offset", Type: "String"},
		},
	},
	"TimeZoneInfo raise Error": {
		Name: "TimeZoneInfo raise Error",
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
