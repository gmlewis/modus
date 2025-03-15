// -*- compile-command: "go test -run ^TestGenerateMetadata_YoutubeWalkthrough$ ."; -*-

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

func TestGenerateMetadata_YoutubeWalkthrough(t *testing.T) {
	meta := setupTestConfig(t, "../testdata/youtube-walkthrough")

	if got, want := meta.Plugin, "youtube-walkthrough"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@youtube-walkthrough"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantYoutubeWalkthroughFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantYoutubeWalkthroughFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	diffMetaTypes(t, wantYoutubeWalkthroughTypes, meta.Types)
}

var wantYoutubeWalkthroughFnExports = metadata.FunctionMap{
	"get_random_quote": {Name: "get_random_quote", Results: []*metadata.Result{{Type: "Quote!Error"}}},
	"say_hello": {
		Name:       "say_hello",
		Parameters: []*metadata.Parameter{{Name: "name~", Type: "String", Default: AnyPtr(`"World"`)}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"say_hello_WithDefaults": {Name: "say_hello_WithDefaults", Results: []*metadata.Result{{Type: "String"}}},
}

var wantYoutubeWalkthroughFnImports = metadata.FunctionMap{
	"modus_http_client.fetch": {
		Name:       "modus_http_client.fetch",
		Parameters: []*metadata.Parameter{{Name: "request", Type: "@http.Request?"}},
		Results:    []*metadata.Result{{Type: "@http.Response?"}},
	},
	"modus_system.logMessage": {
		Name:       "modus_system.logMessage",
		Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
	},
}

var wantYoutubeWalkthroughTypes = metadata.TypeMap{
	"(String)":         {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
	"(String, String)": {Name: "(String, String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}, {Name: "1", Type: "String"}}},
	"@http.Content":    {Name: "@http.Content", Fields: []*metadata.Field{{Name: "data", Type: "String"}}},
	"@http.Header": {Name: "@http.Header", Fields: []*metadata.Field{
		{Name: "name", Type: "String"},
		{Name: "values", Type: "Array[String]"},
	}},
	"@http.Header?":  {Name: "@http.Header?"},
	"@http.Headers":  {Name: "@http.Headers", Fields: []*metadata.Field{{Name: "data", Type: "Map[String, @http.Header?]"}}},
	"@http.Headers?": {Name: "@http.Headers?", Fields: []*metadata.Field{{Name: "data", Type: "Map[String, @http.Header?]"}}},
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
			{Name: "body", Type: "String"},
		},
	},
	"@http.Response!Error": {
		Name: "@http.Response!Error",
		Fields: []*metadata.Field{
			{Name: "status", Type: "UInt16"},
			{Name: "status_text", Type: "String"},
			{Name: "headers", Type: "@http.Headers?"},
			{Name: "body", Type: "String"},
		},
	},
	"@http.Response?": {
		Name: "@http.Response?",
		Fields: []*metadata.Field{
			{Name: "status", Type: "UInt16"},
			{Name: "status_text", Type: "String"},
			{Name: "headers", Type: "@http.Headers?"},
			{Name: "body", Type: "String"},
		},
	},
	"@http.T":                     {Name: "@http.T"},
	"@http.T!Error":               {Name: "@http.T!Error"},
	"@testutils.CallStack[T]":     {Name: "@testutils.CallStack[T]", Fields: []*metadata.Field{{Name: "items", Type: "Array[Array[@testutils.T]]"}}},
	"@testutils.T":                {Name: "@testutils.T"},
	"Array[(String, String)]":     {Name: "Array[(String, String)]"},
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
}
