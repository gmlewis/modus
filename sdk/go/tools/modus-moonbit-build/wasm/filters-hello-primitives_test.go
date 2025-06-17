// -*- compile-command: "go test -run ^TestFilterMetadata.*HelloPrimitives$ ."; -*-

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
	"testing"

	"github.com/gmlewis/modus/lib/wasmextractor"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/google/go-cmp/cmp"
)

const (
	helloPrimitivesPath      = "../../../../moonbit/examples/hello-primitives"
	helloPrimitivesBuildPath = helloPrimitivesPath + "/build"
)

func TestFilterMetadata_HelloPrimitives(t *testing.T) {
	config := &config.Config{
		SourceDir:    helloPrimitivesPath,
		OutputDir:    helloPrimitivesBuildPath,
		WasmFileName: "hello-primitives.wasm",
	}
	copyOfBefore := deepCopyMetadata(t, wantHelloPrimitivesBeforeFilter)
	testFilterMetadataHelper(t, config, copyOfBefore, wantHelloPrimitivesAfterFilter)
}

// The Go `dlv` debugger has problems running the previous test
// so this test provides all the necessary data without requiring
// compilation of the wasm plugin.
func TestFilterMetadataNoCompilation_HelloPrimitives(t *testing.T) {
	// make a copy of the "BEFORE" metadata since it is modified in-place.
	copyOfBefore := deepCopyMetadata(t, wantHelloPrimitivesBeforeFilter)
	filterExportsImportsAndTypes(helloPrimitivesWasmInfo, copyOfBefore)

	if diff := cmp.Diff(wantHelloPrimitivesAfterFilter, copyOfBefore); diff != "" {
		t.Fatalf("filterExportsImportsAndTypes meta AFTER filter mismatch (-want +got):\n%v", diff)
	}
}

var wantHelloPrimitivesBeforeFilter = &metadata.Metadata{
	Plugin: "hello-primitives",
	Module: "@hello-primitives",
	FnExports: metadata.FunctionMap{
		"__modus_hello_primitive_time_duration_max": {
			Name:    "__modus_hello_primitive_time_duration_max",
			Results: []*metadata.Result{{Type: "@time.Duration"}},
		},
		"__modus_hello_primitive_time_duration_min": {
			Name:    "__modus_hello_primitive_time_duration_min",
			Results: []*metadata.Result{{Type: "@time.Duration"}},
		},
		"cabi_realloc": {
			Name: "cabi_realloc",
			Parameters: []*metadata.Parameter{
				{Name: "src_offset", Type: "Int"}, {Name: "src_size", Type: "Int"},
				{Name: "_dst_alignment", Type: "Int"}, {Name: "dst_size", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
		"copy": {
			Name:       "copy",
			Parameters: []*metadata.Parameter{{Name: "dest", Type: "Int"}, {Name: "src", Type: "Int"}},
		},
		"duration_from_nanos": {
			Name:       "duration_from_nanos",
			Parameters: []*metadata.Parameter{{Name: "nanoseconds", Type: "Int64"}},
			Results:    []*metadata.Result{{Type: "@time.Duration raise Error"}},
		},
		"hello_primitive_time_duration_max": {
			Name:    "hello_primitive_time_duration_max",
			Results: []*metadata.Result{{Type: "@time.Duration"}},
		},
		"hello_primitive_time_duration_min": {
			Name:    "hello_primitive_time_duration_min",
			Results: []*metadata.Result{{Type: "@time.Duration"}},
		},
		"malloc": {
			Name:       "malloc",
			Parameters: []*metadata.Parameter{{Name: "size", Type: "Int"}},
			Results:    []*metadata.Result{{Type: "Int"}},
		},
		"ptr_to_none": {Name: "ptr_to_none", Results: []*metadata.Result{{Type: "Int"}}},
		"zoned_date_time_from_unix_seconds_and_nanos": {
			Name:       "zoned_date_time_from_unix_seconds_and_nanos",
			Parameters: []*metadata.Parameter{{Name: "second", Type: "Int64"}, {Name: "nanos", Type: "Int64"}},
			Results:    []*metadata.Result{{Type: "@time.ZonedDateTime raise Error"}},
		},
	},
	FnImports: metadata.FunctionMap{
		"modus_system.logMessage": {
			Name:       "modus_system.logMessage",
			Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
		},
	},
	Types: metadata.TypeMap{
		"(String)":                        {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
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
		"Array[Byte]":                     {Name: "Array[Byte]"},
		"Bool":                            {Name: "Bool"},
		"Byte":                            {Name: "Byte"},
		"FixedArray[Byte]":                {Name: "FixedArray[Byte]"},
		"Int":                             {Name: "Int"},
		"Int64":                           {Name: "Int64"},
		"String":                          {Name: "String"},
	},
}

var wantHelloPrimitivesAfterFilter = &metadata.Metadata{
	Plugin: "hello-primitives",
	Module: "@hello-primitives",
	FnExports: metadata.FunctionMap{
		"cabi_realloc": {
			Name: "cabi_realloc",
			Parameters: []*metadata.Parameter{
				{Name: "src_offset", Type: "Int"}, {Name: "src_size", Type: "Int"},
				{Name: "_dst_alignment", Type: "Int"}, {Name: "dst_size", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
		"copy": {
			Name:       "copy",
			Parameters: []*metadata.Parameter{{Name: "dest", Type: "Int"}, {Name: "src", Type: "Int"}},
		},
		"duration_from_nanos": {
			Name:       "duration_from_nanos",
			Parameters: []*metadata.Parameter{{Name: "nanoseconds", Type: "Int64"}},
			Results:    []*metadata.Result{{Type: "@time.Duration raise Error"}},
		},
		"hello_primitive_time_duration_max": {
			Name:    "hello_primitive_time_duration_max",
			Results: []*metadata.Result{{Type: "@time.Duration"}},
		},
		"hello_primitive_time_duration_min": {
			Name:    "hello_primitive_time_duration_min",
			Results: []*metadata.Result{{Type: "@time.Duration"}},
		},
		"malloc": {
			Name:       "malloc",
			Parameters: []*metadata.Parameter{{Name: "size", Type: "Int"}},
			Results:    []*metadata.Result{{Type: "Int"}},
		},
		"ptr_to_none": {Name: "ptr_to_none", Results: []*metadata.Result{{Type: "Int"}}},
		"zoned_date_time_from_unix_seconds_and_nanos": {
			Name:       "zoned_date_time_from_unix_seconds_and_nanos",
			Parameters: []*metadata.Parameter{{Name: "second", Type: "Int64"}, {Name: "nanos", Type: "Int64"}},
			Results:    []*metadata.Result{{Type: "@time.ZonedDateTime raise Error"}},
		},
	},
	FnImports: metadata.FunctionMap{},
	Types: metadata.TypeMap{
		"(String)":                        {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
		"@time.Duration":                  {Name: "@time.Duration"},
		"@time.Duration raise Error":      {Name: "@time.Duration raise Error"},
		"@time.ZonedDateTime":             {Name: "@time.ZonedDateTime"},
		"@time.ZonedDateTime raise Error": {Name: "@time.ZonedDateTime raise Error"},
		"Int":                             {Name: "Int"},
		"Int64":                           {Name: "Int64"},
		"String":                          {Name: "String"},
	},
}

var helloPrimitivesWasmInfo = &wasmextractor.WasmInfo{
	Exports: []wasmextractor.WasmItem{
		{Name: "cabi_realloc"},
		{Name: "copy"},
		{Name: "duration_from_nanos"},
		{Name: "hello_primitive_time_duration_max"},
		{Name: "hello_primitive_time_duration_min"},
		{Name: "malloc"},
		{Name: "ptr_to_none"},
		{Name: "zoned_date_time_from_unix_seconds_and_nanos"},
	},
	Imports: []wasmextractor.WasmItem{
		{Name: "spectest.print_char"},
	},
}
