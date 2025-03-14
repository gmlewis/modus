// -*- compile-command: "go test -run ^TestFilterMetadata_YoutubeWalkthrough$ ."; -*-

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
	youtubePath      = "../../../../moonbit/examples/youtube-walkthrough"
	youtubeBuildPath = youtubePath + "/build"
)

func TestFilterMetadata_YoutubeWalkthrough(t *testing.T) {
	config := &config.Config{
		SourceDir:    youtubePath,
		OutputDir:    youtubeBuildPath,
		WasmFileName: "youtube-walkthrough.wasm",
	}
	copyOfBefore := deepCopyMetadata(t, wantYoutubeWalkthroughBeforeFilter)
	testFilterMetadataHelper(t, config, copyOfBefore, wantYoutubeWalkthroughAfterFilter)
}

// The Go `dlv` debugger has problems running the previous test
// so this test provides all the necessary data without requiring
// compilation of the wasm plugin.
func TestFilterMetadataNoCompilation_YoutubeWalkthrough(t *testing.T) {
	// make a copy of the "BEFORE" metadata since it is modified in-place.
	copyOfBefore := deepCopyMetadata(t, wantYoutubeWalkthroughBeforeFilter)
	filterExportsImportsAndTypes(youtubeWasmInfo, copyOfBefore)

	if diff := cmp.Diff(wantYoutubeWalkthroughAfterFilter, copyOfBefore); diff != "" {
		t.Fatalf("filterExportsImportsAndTypes meta AFTER filter mismatch (-want +got):\n%v", diff)
	}
}

var wantYoutubeWalkthroughBeforeFilter = &metadata.Metadata{
	Plugin: "youtube-walkthrough",
	Module: "@youtube-walkthrough",
	FnExports: metadata.FunctionMap{
		"__modus_get_random_quote": {Name: "__modus_get_random_quote", Results: []*metadata.Result{{Type: "Quote!Error"}}},
		"__modus_say_hello": {
			Name:       "__modus_say_hello",
			Parameters: []*metadata.Parameter{{Name: "name", Type: "String"}},
			Results:    []*metadata.Result{{Type: "String"}},
		},
		"__modus_say_hello_WithDefaults": {Name: "__modus_say_hello_WithDefaults", Results: []*metadata.Result{{Type: "String"}}},
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
		"get_random_quote": {Name: "get_random_quote", Results: []*metadata.Result{{Type: "Quote!Error"}}},
		"malloc": {
			Name:       "malloc",
			Parameters: []*metadata.Parameter{{Name: "size", Type: "Int"}},
			Results:    []*metadata.Result{{Type: "Int"}},
		},
		"ptr_to_none": {Name: "ptr_to_none", Results: []*metadata.Result{{Type: "Int"}}},
		"read_map": {
			Name: "read_map",
			Parameters: []*metadata.Parameter{
				{Name: "key_type_name_ptr", Type: "Int"},
				{Name: "value_type_name_ptr", Type: "Int"}, {Name: "map_ptr", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int64"}},
		},
		"say_hello": {
			Name:       "say_hello",
			Parameters: []*metadata.Parameter{{Name: "name~", Type: "String"}},
			Results:    []*metadata.Result{{Type: "String"}},
		},
		"say_hello_WithDefaults": {Name: "say_hello_WithDefaults", Results: []*metadata.Result{{Type: "String"}}},
		"write_map": {
			Name: "write_map",
			Parameters: []*metadata.Parameter{
				{Name: "key_type_name_ptr", Type: "Int"},
				{Name: "value_type_name_ptr", Type: "Int"}, {Name: "keys_ptr", Type: "Int"},
				{Name: "values_ptr", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
	},
	FnImports: metadata.FunctionMap{
		"modus_http_client.fetch": {
			Name:       "modus_http_client.fetch",
			Parameters: []*metadata.Parameter{{Name: "request", Type: "@http.Request?"}},
			Results:    []*metadata.Result{{Type: "@http.Response?"}},
		},
		"modus_system.logMessage": {
			Name:       "modus_system.logMessage",
			Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
		},
	},
	Types: metadata.TypeMap{
		"(String)":         {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
		"(String, String)": {Name: "(String, String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}, {Name: "1", Type: "String"}}},
		"@http.Content":    {Name: "@http.Content", Fields: []*metadata.Field{{Name: "data", Type: "String"}}},
		"@http.Header": {Name: "@http.Header", Fields: []*metadata.Field{
			{Name: "name", Type: "String"},
			{Name: "values", Type: "Array[String]"},
		}},
		"@http.Header?":  {Name: "@http.Header?"},
		"@http.Headers":  {Name: "@http.Headers", Fields: []*metadata.Field{{Name: "data", Type: "Map[String, @http.Header?]"}}},
		"@http.Headers?": {Name: "@http.Headers?"},
		"@http.Request": {Name: "@http.Request", Fields: []*metadata.Field{
			{Name: "url", Type: "String"},
			{Name: "method_", Type: "String"},
			{Name: "headers", Type: "@http.Headers?"},
			{Name: "body", Type: "Array[Byte]"},
		}},
		"@http.Request?": {
			Name: "@http.Request?",
			Fields: []*metadata.Field{
				{Name: "url", Type: "String"},
				{Name: "method_", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		"@http.RequestOptions": {Name: "@http.RequestOptions"},
		"@http.Response": {
			Name: "@http.Response",
			Fields: []*metadata.Field{
				{Name: "status", Type: "UInt16"},
				{Name: "status_text", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		"@http.Response!Error": {
			Name: "@http.Response!Error",
			Fields: []*metadata.Field{
				{Name: "status", Type: "UInt16"},
				{Name: "status_text", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		"@http.Response?": {
			Name: "@http.Response?",
			Fields: []*metadata.Field{
				{Name: "status", Type: "UInt16"},
				{Name: "status_text", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		"@http.T":                     {Name: "@http.T"},
		"@http.T!Error":               {Name: "@http.T!Error"},
		"@testutils.CallStack[T]":     {Name: "@testutils.CallStack[T]", Fields: []*metadata.Field{{Name: "items", Type: "Array[Array[@testutils.T]]"}}},
		"@testutils.T":                {Name: "@testutils.T"},
		"Array[@http.Header?]":        {Name: "Array[@http.Header?]"},
		"Array[@http.RequestOptions]": {Name: "Array[@http.RequestOptions]"},
		"Array[@testutils.T]":         {Name: "Array[@testutils.T]"},
		"Array[Array[@testutils.T]]":  {Name: "Array[Array[@testutils.T]]"},
		"Array[Array[String]]":        {Name: "Array[Array[String]]"},
		"Array[Byte]":                 {Name: "Array[Byte]"},
		"Array[String]":               {Name: "Array[String]"},
		"Bool":                        {Name: "Bool"},
		"Byte":                        {Name: "Byte"},
		"FixedArray[Double]":          {Name: "FixedArray[Double]"},
		"FixedArray[Float]":           {Name: "FixedArray[Float]"},
		"FixedArray[Int64]":           {Name: "FixedArray[Int64]"},
		"FixedArray[Int]":             {Name: "FixedArray[Int]"},
		"FixedArray[UInt64]":          {Name: "FixedArray[UInt64]"},
		"FixedArray[UInt]":            {Name: "FixedArray[UInt]"},
		"Float":                       {Name: "Float"},
		"Double":                      {Name: "Double"},
		"Int":                         {Name: "Int"},
		"Int64":                       {Name: "Int64"},
		"Json":                        {Name: "Json"},
		"Map[String, @http.Header?]":  {Name: "Map[String, @http.Header?]"},
		"Map[String, Array[String]]":  {Name: "Map[String, Array[String]]"},
		"Map[String, String]":         {Name: "Map[String, String]"},
		"Quote": {
			Name:   "Quote",
			Fields: []*metadata.Field{{Name: "quote", Type: "String"}, {Name: "author", Type: "String"}},
		},
		"Quote!Error": {
			Name:   "Quote!Error",
			Fields: []*metadata.Field{{Name: "quote", Type: "String"}, {Name: "author", Type: "String"}},
		},
		"String": {Name: "String"},
		"UInt":   {Name: "UInt"},
		"UInt16": {Name: "UInt16"},
		"UInt64": {Name: "UInt64"},
	},
}

var wantYoutubeWalkthroughAfterFilter = &metadata.Metadata{
	Plugin:    "youtube-walkthrough",
	Module:    "@youtube-walkthrough",
	FnExports: metadata.FunctionMap{},
	FnImports: metadata.FunctionMap{},
	Types:     metadata.TypeMap{},
}

var youtubeWasmInfo = &wasmextractor.WasmInfo{
	Exports: []wasmextractor.WasmItem{
		{Name: "memory"},
		{Name: "get_random_quote"},
		{Name: "say_hello"},
		{Name: "say_hello_WithDefaults"},
		{Name: "load32"},
		{Name: "copy"},
		{Name: "free"},
		{Name: "store32"},
		{Name: "store8"},
		{Name: "malloc"},
		{Name: "cabi_realloc"},
		{Name: "ptr2str"},
		{Name: "ptr_to_none"},
		{Name: "read_map"},
		{Name: "write_map"},
	},
	Imports: []wasmextractor.WasmItem{
		{Name: "modus_http_client.fetch"},
		{Name: "modus_system.logMessage"},
	},
}
