// -*- compile-command: "go test -run ^TestPackage_Time$ ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package packages

import (
	"encoding/json"
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestPackage_Time(t *testing.T) {
	t.Parallel()
	dir := "testdata/time-example"
	testPackageLoadHelper(t, "time", dir, wantPackageTime)
}

var wantPackageTime = &Package{
	MoonPkgJSON: MoonPkgJSON{
		IsMain: false,
		Imports: []json.RawMessage{
			json.RawMessage(`"gmlewis/modus/pkg/console"`),
			json.RawMessage(`"gmlewis/modus/wit/interface/imports/wasi/clocks/wallClock"`),
			json.RawMessage(`"moonbitlang/x/sys"`),
			json.RawMessage(`"moonbitlang/x/time"`),
		},
		TestImport: []json.RawMessage{
			[]byte(`"gmlewis/modus/pkg/testutils"`),
		},
	},
	MoonBitFiles: []string{"testdata/time-example/time-example.mbt"},
	ID:           "moonbit-main",
	Name:         "main",
	PkgPath:      "@time-example",
	StructLookup: map[string]*ast.TypeSpec{
		"TimeZoneInfo": {Name: &ast.Ident{Name: "TimeZoneInfo"}, Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{{Name: "standard_name"}}, Type: &ast.Ident{Name: "String"}},
					{Names: []*ast.Ident{{Name: "standard_offset"}}, Type: &ast.Ident{Name: "String"}},
					{Names: []*ast.Ident{{Name: "daylight_name"}}, Type: &ast.Ident{Name: "String"}},
					{Names: []*ast.Ident{{Name: "daylight_offset"}}, Type: &ast.Ident{Name: "String"}},
				},
			},
		}},
	},
	Syntax: []*ast.File{
		{
			Name: &ast.Ident{Name: "testdata/time-example/time-example.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{&ast.TypeSpec{Name: &ast.Ident{Name: "TimeZoneInfo"}, Type: &ast.StructType{
						Fields: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "standard_name"}}, Type: &ast.Ident{Name: "String"}},
								{Names: []*ast.Ident{{Name: "standard_offset"}}, Type: &ast.Ident{Name: "String"}},
								{Names: []*ast.Ident{{Name: "daylight_name"}}, Type: &ast.Ident{Name: "String"}},
								{Names: []*ast.Ident{{Name: "daylight_offset"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
					}}},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Returns the current time in UTC."}}},
					Name: &ast.Ident{Name: "get_utc_time"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{{Type: &ast.Ident{Name: "@time.ZonedDateTime!Error"}}},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Returns the current local time."}}},
					Name: &ast.Ident{Name: "get_local_time"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{{Type: &ast.Ident{Name: "String!Error"}}},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Returns the current time in a specified time zone."}}},
					Name: &ast.Ident{Name: "get_time_in_zone"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{{Names: []*ast.Ident{{Name: "tz"}}, Type: &ast.Ident{Name: "String"}}},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{{Type: &ast.Ident{Name: "String!Error"}}},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Returns the local time zone identifier."}}},
					Name: &ast.Ident{Name: "get_local_time_zone"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Returns some basic information about the time zone specified."}}},
					Name: &ast.Ident{Name: "get_time_zone_info"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{{Names: []*ast.Ident{{Name: "tz"}}, Type: &ast.Ident{Name: "String"}}},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{{Type: &ast.Ident{Name: "TimeZoneInfo!Error"}}},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/imports/wasi/clocks/wallClock"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
	},
	TypesInfo: &types.Info{
		Defs: map[*ast.Ident]types.Object{
			// Using &moonType{} and &moonFunc{} is a hack to fake a struct/func for testing purposes only:
			{Name: "TimeZoneInfo"}:        &moonFunc{funcName: "type TimeZoneInfo = struct{standard_name String; standard_offset String; daylight_name String; daylight_offset String}"},
			{Name: "get_local_time"}:      &moonFunc{funcName: "func @time-example.get_local_time() String!Error"},
			{Name: "get_local_time_zone"}: &moonFunc{funcName: "func @time-example.get_local_time_zone() String"},
			{Name: "get_time_in_zone"}:    &moonFunc{funcName: "func @time-example.get_time_in_zone(tz String) String!Error"},
			{Name: "get_time_zone_info"}:  &moonFunc{funcName: "func @time-example.get_time_zone_info(tz String) TimeZoneInfo!Error"},
			{Name: "get_utc_time"}:        &moonFunc{funcName: "func @time-example.get_utc_time() @time.ZonedDateTime!Error"},
		},
	},
}
