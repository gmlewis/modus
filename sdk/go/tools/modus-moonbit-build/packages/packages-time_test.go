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
	dir := "../testdata/time-example"
	testPackageLoadHelper(t, "time", dir, wantPackageTime)
}

var wantPackageTime = &Package{
	MoonPkgJSON: MoonPkgJSON{
		IsMain: false,
		Imports: []json.RawMessage{
			json.RawMessage(`"gmlewis/modus/pkg/console"`),
			json.RawMessage(`"gmlewis/modus/pkg/localtime"`),
			json.RawMessage(`"gmlewis/modus/wit/interface/wasi"`),
			json.RawMessage(`"moonbitlang/x/time"`),
		},
		Targets: map[string][]string{
			"modus_post_generated.mbt": []string{"wasm"},
		},
		LinkTargets: map[string]*LinkTarget{
			"wasm": {
				Exports: []string{
					"__modus_get_local_time_modus:get_local_time_modus",
					"__modus_get_local_time_moonbit:get_local_time_moonbit",
					"__modus_get_local_time_zone_id:get_local_time_zone_id",
					"__modus_get_time_in_zone_modus:get_time_in_zone_modus",
					"__modus_get_time_in_zone_moonbit:get_time_in_zone_moonbit",
					"__modus_get_time_zone_info:get_time_zone_info",
					"__modus_get_utc_time:get_utc_time",
					"cabi_realloc",
					"copy",
					"duration_from_nanos",
					"free",
					"load32",
					"malloc",
					"ptr2str",
					"ptr_to_none",
					"read_map",
					"store32",
					"store8",
					"write_map",
					"zoned_date_time_from_unix_seconds_and_nanos",
				},
				ExportMemoryName: "memory",
			},
		},
	},
	ID:           "moonbit-main",
	Name:         "main",
	PkgPath:      "",
	MoonBitFiles: []string{"../testdata/time-example/main.mbt"},
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
	PossiblyMissingUnderlyingTypes: map[string]struct{}{
		"String": {},
	},
	Syntax: []*ast.File{
		{
			Name: &ast.Ident{Name: "../testdata/time-example/main.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TimeZoneInfo"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "standard_name"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "standard_offset"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "daylight_name"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "daylight_offset"}}, Type: &ast.Ident{Name: "String"}},
									},
								},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Returns the current time in UTC."},
						},
					},
					Name: &ast.Ident{Name: "get_utc_time"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "@time.ZonedDateTime!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Returns the current local time using the moonbitlang/x/time package."},
						},
					},
					Name: &ast.Ident{Name: "get_local_time_moonbit"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Returns the current local time using the Modus host Go function."},
						},
					},
					Name: &ast.Ident{Name: "get_local_time_modus"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Returns the current time in a specified time zone using"},
							{Text: "// the moonbitlang/x/time package."},
						},
					},
					Name: &ast.Ident{Name: "get_time_in_zone_moonbit"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "tz"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Returns the current time in a specified time zone using"},
							{Text: "// the Modus host Go function."},
						},
					},
					Name: &ast.Ident{Name: "get_time_in_zone_modus"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "tz"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Returns the local time zone identifier."},
						},
					},
					Name: &ast.Ident{Name: "get_local_time_zone_id"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Returns some basic information about the time zone specified."},
						},
					},
					Name: &ast.Ident{Name: "get_time_zone_info"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "tz"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TimeZoneInfo!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
	},
	TypesInfo: &types.Info{
		Defs: map[*ast.Ident]types.Object{
			// Using &moonType{} and &moonFunc{} is a hack to fake a struct/func for testing purposes only:
			{Name: "TimeZoneInfo"}:             &moonFunc{funcName: "type TimeZoneInfo = struct{standard_name String; standard_offset String; daylight_name String; daylight_offset String}"},
			{Name: "get_local_time_modus"}:     &moonFunc{funcName: "func get_local_time_modus() String"},
			{Name: "get_local_time_moonbit"}:   &moonFunc{funcName: "func get_local_time_moonbit() String!Error"},
			{Name: "get_local_time_zone_id"}:   &moonFunc{funcName: "func get_local_time_zone_id() String"},
			{Name: "get_time_in_zone_modus"}:   &moonFunc{funcName: "func get_time_in_zone_modus(tz String) String"},
			{Name: "get_time_in_zone_moonbit"}: &moonFunc{funcName: "func get_time_in_zone_moonbit(tz String) String!Error"},
			{Name: "get_time_zone_info"}:       &moonFunc{funcName: "func get_time_zone_info(tz String) TimeZoneInfo!Error"},
			{Name: "get_utc_time"}:             &moonFunc{funcName: "func get_utc_time() @time.ZonedDateTime!Error"},
		},
	},
}
