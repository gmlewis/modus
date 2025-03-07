// -*- compile-command: "go test -run ^TestPackage_Testsuite$ ."; -*-

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

func TestPackage_Testsuite(t *testing.T) {
	t.Parallel()
	dir := "testdata/test-suite"
	testPackageLoadHelper(t, "testsuite", dir, wantPackageTestsuite)
}

var wantPackageTestsuite = &Package{
	MoonPkgJSON: MoonPkgJSON{
		IsMain: false,
		Imports: []json.RawMessage{
			json.RawMessage(`"gmlewis/modus/pkg/console"`),
			json.RawMessage(`"gmlewis/modus/pkg/localtime"`),
			json.RawMessage(`"gmlewis/modus/wit/interface/wasi"`),
			json.RawMessage(`"gmlewis/modus/pkg/time"`),
			json.RawMessage(`"moonbitlang/x/sys"`),
		},
		Targets: map[string][]string{
			"modus_post_generated.mbt": {"wasm"},
			"modus_pre_generated.mbt":  {"wasm"},
		},
		LinkTargets: map[string]*LinkTarget{
			"wasm": {
				Exports: []string{
					"__modus_add3:add3",
					"__modus_add3_WithDefaults:add3_WithDefaults",
					"__modus_add:add",
					"__modus_get_current_time:get_current_time",
					"__modus_get_current_time_formatted:get_current_time_formatted",
					"__modus_get_full_name:get_full_name",
					"__modus_get_local_time:get_local_time",
					"__modus_get_local_time_zone_id:get_local_time_zone_id",
					"__modus_get_people:get_people",
					"__modus_get_person:get_person",
					"__modus_get_random_person:get_random_person",
					"__modus_get_time_in_zone:get_time_in_zone",
					"__modus_get_time_zone_info:get_time_zone_info",
					"__modus_get_utc_time:get_utc_time",
					"__modus_hello_array_of_ints:hello_array_of_ints",
					"__modus_hello_array_of_ints_WithDefaults:hello_array_of_ints_WithDefaults",
					"__modus_hello_array_of_ints_option:hello_array_of_ints_option",
					"__modus_hello_array_of_ints_option_WithDefaults:hello_array_of_ints_option_WithDefaults",
					"__modus_hello_array_of_strings:hello_array_of_strings",
					"__modus_hello_array_of_strings_WithDefaults:hello_array_of_strings_WithDefaults",
					"__modus_hello_array_of_strings_option:hello_array_of_strings_option",
					"__modus_hello_array_of_strings_option_WithDefaults:hello_array_of_strings_option_WithDefaults",
					"__modus_hello_maps_n_items:hello_maps_n_items",
					"__modus_hello_maps_n_items_WithDefaults:hello_maps_n_items_WithDefaults",
					"__modus_hello_maps_n_items_option:hello_maps_n_items_option",
					"__modus_hello_maps_n_items_option_WithDefaults:hello_maps_n_items_option_WithDefaults",
					"__modus_hello_option_empty_string:hello_option_empty_string",
					"__modus_hello_option_empty_string_WithDefaults:hello_option_empty_string_WithDefaults",
					"__modus_hello_option_none:hello_option_none",
					"__modus_hello_option_none_WithDefaults:hello_option_none_WithDefaults",
					"__modus_hello_option_some_string:hello_option_some_string",
					"__modus_hello_option_some_string_WithDefaults:hello_option_some_string_WithDefaults",
					"__modus_hello_primitive_bool_max:hello_primitive_bool_max",
					"__modus_hello_primitive_bool_min:hello_primitive_bool_min",
					"__modus_hello_primitive_byte_max:hello_primitive_byte_max",
					"__modus_hello_primitive_byte_min:hello_primitive_byte_min",
					"__modus_hello_primitive_char_max:hello_primitive_char_max",
					"__modus_hello_primitive_char_min:hello_primitive_char_min",
					"__modus_hello_primitive_double_max:hello_primitive_double_max",
					"__modus_hello_primitive_double_min:hello_primitive_double_min",
					"__modus_hello_primitive_float_max:hello_primitive_float_max",
					"__modus_hello_primitive_float_min:hello_primitive_float_min",
					"__modus_hello_primitive_int16_max:hello_primitive_int16_max",
					"__modus_hello_primitive_int16_min:hello_primitive_int16_min",
					"__modus_hello_primitive_int64_max:hello_primitive_int64_max",
					"__modus_hello_primitive_int64_min:hello_primitive_int64_min",
					"__modus_hello_primitive_int_max:hello_primitive_int_max",
					"__modus_hello_primitive_int_min:hello_primitive_int_min",
					"__modus_hello_primitive_uint16_max:hello_primitive_uint16_max",
					"__modus_hello_primitive_uint16_min:hello_primitive_uint16_min",
					"__modus_hello_primitive_uint64_max:hello_primitive_uint64_max",
					"__modus_hello_primitive_uint64_min:hello_primitive_uint64_min",
					"__modus_hello_primitive_uint_max:hello_primitive_uint_max",
					"__modus_hello_primitive_uint_min:hello_primitive_uint_min",
					"__modus_hello_world:hello_world",
					"__modus_hello_world_with_arg:hello_world_with_arg",
					"__modus_hello_world_with_optional_arg:hello_world_with_optional_arg",
					"__modus_hello_world_with_optional_arg_WithDefaults:hello_world_with_optional_arg_WithDefaults",
					"__modus_log_message:log_message",
					"__modus_test_abort:test_abort",
					"__modus_test_alternative_error:test_alternative_error",
					"__modus_test_exit:test_exit",
					"__modus_test_logging:test_logging",
					"__modus_test_normal_error:test_normal_error",
					"bytes2ptr",
					"cabi_realloc",
					"copy",
					"double_array2ptr",
					"extend16",
					"extend8",
					"f32_to_i32",
					"f32_to_i64",
					"float_array2ptr",
					"free",
					"int64_array2ptr",
					"int_array2ptr",
					"load16",
					"load16_u",
					"load32",
					"load64",
					"load8",
					"load8_u",
					"loadf32",
					"loadf64",
					"malloc",
					"ptr2bytes",
					"ptr2double_array",
					"ptr2float_array",
					"ptr2int64_array",
					"ptr2int_array",
					"ptr2str",
					"ptr2uint64_array",
					"ptr2uint_array",
					"store16",
					"store32",
					"store64",
					"store8",
					"storef32",
					"storef64",
					"str2ptr",
					"uint64_array2ptr",
					"uint_array2ptr",
				},
				ExportMemoryName: "memory",
			},
		},
	},
	MoonBitFiles: []string{
		"testdata/test-suite/arrays.mbt",
		"testdata/test-suite/maps.mbt",
		"testdata/test-suite/modus_post_generated.mbt",
		"testdata/test-suite/modus_pre_generated.mbt",
		"testdata/test-suite/primitives.mbt",
		"testdata/test-suite/simple.mbt",
		"testdata/test-suite/strings.mbt",
		"testdata/test-suite/time.mbt",
	},
	ID:      "moonbit-main",
	Name:    "main",
	PkgPath: "@test-suite",
	StructLookup: map[string]*ast.TypeSpec{
		"Cleanup": {Name: &ast.Ident{Name: "Cleanup"}, Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{{Name: "address"}}, Type: &ast.Ident{Name: "Int"}},
					{Names: []*ast.Ident{{Name: "size"}}, Type: &ast.Ident{Name: "Int"}},
					{Names: []*ast.Ident{{Name: "align"}}, Type: &ast.Ident{Name: "Int"}},
				},
			},
		}},
		"Person": {Name: &ast.Ident{Name: "Person"}, Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{{Name: "firstName"}}, Type: &ast.Ident{Name: "String"}},
					{Names: []*ast.Ident{{Name: "lastName"}}, Type: &ast.Ident{Name: "String"}},
					{Names: []*ast.Ident{{Name: "age"}}, Type: &ast.Ident{Name: "Int"}},
				},
			},
		}},
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
		"TimeZoneInfo!Error": {Name: &ast.Ident{Name: "TimeZoneInfo"}},
	},
	Syntax: []*ast.File{
		{
			Name: &ast.Ident{Name: "testdata/test-suite/arrays.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_array_of_ints"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n~"}}, Type: &ast.Ident{Name: "Int"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:3`"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_array_of_ints_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_array_of_ints_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n~"}}, Type: &ast.Ident{Name: "Int"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:3`"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_array_of_ints_option_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_array_of_strings"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n~"}}, Type: &ast.Ident{Name: "Int"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:3`"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_array_of_strings_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_array_of_strings_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n~"}}, Type: &ast.Ident{Name: "Int"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:3`"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[String]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_array_of_strings_option_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[String]?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/time"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "testdata/test-suite/maps.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_maps_n_items"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "n~"}}, Type: &ast.Ident{Name: "Int"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:3`"}},
						}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
					}},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_maps_n_items_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_maps_n_items_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "n~"}}, Type: &ast.Ident{Name: "Int"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:3`"}},
						}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_maps_n_items_option_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/time"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "testdata/test-suite/modus_post_generated.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: &ast.Ident{Name: "Cleanup"}, Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "address"}}, Type: &ast.Ident{Name: "Int"}},
									{Names: []*ast.Ident{{Name: "size"}}, Type: &ast.Ident{Name: "Int"}},
									{Names: []*ast.Ident{{Name: "align"}}, Type: &ast.Ident{Name: "Int"}},
								},
							},
						}},
				}},
				&ast.FuncDecl{Doc: &ast.CommentGroup{}, Name: &ast.Ident{Name: "cabi_realloc"}, Type: &ast.FuncType{
					Params: &ast.FieldList{List: []*ast.Field{
						{Names: []*ast.Ident{{Name: "src_offset"}}, Type: &ast.Ident{Name: "Int"}},
						{Names: []*ast.Ident{{Name: "src_size"}}, Type: &ast.Ident{Name: "Int"}},
						{Names: []*ast.Ident{{Name: "_dst_alignment"}}, Type: &ast.Ident{Name: "Int"}},
						{Names: []*ast.Ident{{Name: "dst_size"}}, Type: &ast.Ident{Name: "Int"}},
					}},
					Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
				}},
				&ast.FuncDecl{Doc: &ast.CommentGroup{}, Name: &ast.Ident{Name: "malloc"}, Type: &ast.FuncType{
					Params: &ast.FieldList{List: []*ast.Field{
						{Names: []*ast.Ident{{Name: "size"}}, Type: &ast.Ident{Name: "Int"}},
					}},
					Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
				}},
				&ast.FuncDecl{Doc: &ast.CommentGroup{}, Name: &ast.Ident{Name: "copy"}, Type: &ast.FuncType{
					Params: &ast.FieldList{List: []*ast.Field{
						{Names: []*ast.Ident{{Name: "dest"}}, Type: &ast.Ident{Name: "Int"}},
						{Names: []*ast.Ident{{Name: "src"}}, Type: &ast.Ident{Name: "Int"}},
					}},
				}},
				&ast.FuncDecl{Doc: &ast.CommentGroup{}, Name: &ast.Ident{Name: "ptr2uint64_array"}, Type: &ast.FuncType{
					Params: &ast.FieldList{List: []*ast.Field{
						{Names: []*ast.Ident{{Name: "ptr"}}, Type: &ast.Ident{Name: "Int"}},
					}},
					Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "FixedArray[UInt64]"}}}},
				}},
				&ast.FuncDecl{Doc: &ast.CommentGroup{}, Name: &ast.Ident{Name: "ptr2int64_array"}, Type: &ast.FuncType{
					Params: &ast.FieldList{List: []*ast.Field{
						{Names: []*ast.Ident{{Name: "ptr"}}, Type: &ast.Ident{Name: "Int"}},
					}},
					Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "FixedArray[Int64]"}}}},
				}},
				&ast.FuncDecl{Doc: &ast.CommentGroup{}, Name: &ast.Ident{Name: "ptr2double_array"}, Type: &ast.FuncType{
					Params: &ast.FieldList{List: []*ast.Field{
						{Names: []*ast.Ident{{Name: "ptr"}}, Type: &ast.Ident{Name: "Int"}},
					}},
					Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "FixedArray[Double]"}}}},
				}},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/time"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "testdata/test-suite/modus_pre_generated.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_array_of_ints"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Array[Int]"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_array_of_ints_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Array[Int]"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_array_of_ints_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Array[Int]?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_array_of_ints_option_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Array[Int]?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_array_of_strings"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Array[String]"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_array_of_strings_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Array[String]"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_array_of_strings_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Array[String]?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_array_of_strings_option_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Array[String]?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_maps_n_items"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Map[String, String]"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_maps_n_items_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Map[String, String]"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_maps_n_items_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Map[String, String]?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_maps_n_items_option_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Map[String, String]?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_bool_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Bool"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_bool_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Bool"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_byte_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Byte"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_byte_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Byte"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_char_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Char"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_char_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Char"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_double_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Double"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_double_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Double"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_float_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Float"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_float_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Float"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_int_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_int_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_int16_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int16"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_int16_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int16"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_int64_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int64"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_int64_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int64"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_uint_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_uint_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_uint16_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt16"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_uint16_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt16"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_uint64_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt64"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_primitive_uint64_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt64"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_log_message"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String"}},
						}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_add"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "x"}}, Type: &ast.Ident{Name: "Int"}},
							{Names: []*ast.Ident{{Name: "y"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_add3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Int"}},
							{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
							{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_add3_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Int"}},
							{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_current_time"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "@time.ZonedDateTime!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_current_time_formatted"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_full_name"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "first_name"}}, Type: &ast.Ident{Name: "String"}},
							{Names: []*ast.Ident{{Name: "last_name"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_person"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Person"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_random_person"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Person"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_people"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Array[Person]"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_test_normal_error"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "input"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_test_alternative_error"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "input"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_test_abort"},
					Type: &ast.FuncType{Params: &ast.FieldList{}},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_test_exit"},
					Type: &ast.FuncType{Params: &ast.FieldList{}},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_test_logging"},
					Type: &ast.FuncType{Params: &ast.FieldList{}},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_option_empty_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "some"}}, Type: &ast.Ident{Name: "Bool"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_option_empty_string_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_option_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "some"}}, Type: &ast.Ident{Name: "Bool"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_option_none_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_option_some_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "some"}}, Type: &ast.Ident{Name: "Bool"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_option_some_string_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_world_with_arg"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "name"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_world_with_optional_arg"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "name"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_world_with_optional_arg_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_hello_world"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_utc_time"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "@time.ZonedDateTime!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_local_time"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_time_in_zone"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "tz"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_local_time_zone_id"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "__modus_get_time_zone_info"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "tz"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "TimeZoneInfo!Error"}}}},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/time"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "testdata/test-suite/primitives.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_bool_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Bool"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_bool_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Bool"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_byte_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Byte"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_byte_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Byte"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_char_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Char"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_char_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Char"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_double_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Double"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_double_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Double"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_float_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Float"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_float_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Float"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int16_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int16"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int16_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int16"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int64_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int64"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int64_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int64"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint16_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt16"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint16_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt16"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint64_min"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt64"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint64_max"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "UInt64"}}}},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/time"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "testdata/test-suite/simple.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: &ast.Ident{Name: "Person"}, Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "firstName"}}, Type: &ast.Ident{Name: "String"}},
									{Names: []*ast.Ident{{Name: "lastName"}}, Type: &ast.Ident{Name: "String"}},
									{Names: []*ast.Ident{{Name: "age"}}, Type: &ast.Ident{Name: "Int"}},
								},
							},
						}},
				}},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Logs a message."}}},
					Name: &ast.Ident{Name: "log_message"}, Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String"}},
						}},
					}},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Adds two integers together and returns the result."}}},
					Name: &ast.Ident{Name: "add"}, Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "x"}}, Type: &ast.Ident{Name: "Int"}},
							{Names: []*ast.Ident{{Name: "y"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
					}},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Adds three integers together and returns the result."},
						{Text: "// The third integer is optional."},
					}},
					Name: &ast.Ident{Name: "add3"}, Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Int"}},
							{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
							{Names: []*ast.Ident{{Name: "c~"}}, Type: &ast.Ident{Name: "Int"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:0`"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
					}},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Adds three integers together and returns the result."},
						{Text: "// The third integer is optional."},
					}},
					Name: &ast.Ident{Name: "add3_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Int"}},
							{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Int"}}}},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Returns the current time."},
					}},
					Name: &ast.Ident{Name: "get_current_time"}, Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "@time.ZonedDateTime!Error"}}}},
					}},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Returns the current time formatted as a string."},
					}},
					Name: &ast.Ident{Name: "get_current_time_formatted"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Combines the first and last name of a person, and returns the full name."},
					}},
					Name: &ast.Ident{Name: "get_full_name"}, Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "first_name"}}, Type: &ast.Ident{Name: "String"}},
							{Names: []*ast.Ident{{Name: "last_name"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					}},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Gets a person object."},
					}},
					Name: &ast.Ident{Name: "get_person"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Person"}}}},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Gets a random person object from a list of people."},
					}},
					Name: &ast.Ident{Name: "get_random_person"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Person"}}}},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Gets a list of people."},
					}},
					Name: &ast.Ident{Name: "get_people"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "Array[Person]"}}}},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Tests returning an error."},
					}},
					Name: &ast.Ident{Name: "test_normal_error"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "input"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Tests an alternative way to handle errors in functions."},
					}},
					Name: &ast.Ident{Name: "test_alternative_error"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "input"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Tests an abort."},
					}},
					Name: &ast.Ident{Name: "test_abort"},
					Type: &ast.FuncType{Params: &ast.FieldList{}},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Tests an exit with a non-zero exit code."},
					}},
					Name: &ast.Ident{Name: "test_exit"},
					Type: &ast.FuncType{Params: &ast.FieldList{}},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{List: []*ast.Comment{
						{Text: "// Tests logging at different levels."},
					}},
					Name: &ast.Ident{Name: "test_logging"},
					Type: &ast.FuncType{Params: &ast.FieldList{}},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/time"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "testdata/test-suite/strings.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_empty_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "some~"}}, Type: &ast.Ident{Name: "Bool"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:true`"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_empty_string_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "some~"}}, Type: &ast.Ident{Name: "Bool"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:false`"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_none_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_some_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "some~"}}, Type: &ast.Ident{Name: "Bool"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:true`"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_some_string_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String?"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_world_with_arg"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "name"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_world_with_optional_arg"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "name~"}}, Type: &ast.Ident{Name: "String"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:\"世界 🌍 from MoonBit\"`"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_world_with_optional_arg_WithDefaults"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_world"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/time"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "testdata/test-suite/time.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: &ast.Ident{Name: "TimeZoneInfo"}, Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "standard_name"}}, Type: &ast.Ident{Name: "String"}},
									{Names: []*ast.Ident{{Name: "standard_offset"}}, Type: &ast.Ident{Name: "String"}},
									{Names: []*ast.Ident{{Name: "daylight_name"}}, Type: &ast.Ident{Name: "String"}},
									{Names: []*ast.Ident{{Name: "daylight_offset"}}, Type: &ast.Ident{Name: "String"}},
								},
							},
						}},
				}},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Returns the current time in UTC."}}},
					Name: &ast.Ident{Name: "get_utc_time"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "@time.ZonedDateTime!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Returns the current local time."}}},
					Name: &ast.Ident{Name: "get_local_time"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{{Text: "// Returns the current time in a specified time zone."}},
					},
					Name: &ast.Ident{Name: "get_time_in_zone"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "tz"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String!Error"}}}},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Returns the local time zone identifier."}}},
					Name: &ast.Ident{Name: "get_local_time_zone_id"},
					Type: &ast.FuncType{
						Params:  &ast.FieldList{},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "String"}}}},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{{Text: "// Returns some basic information about the time zone specified."}},
					},
					Name: &ast.Ident{Name: "get_time_zone_info"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{Names: []*ast.Ident{{Name: "tz"}}, Type: &ast.Ident{Name: "String"}},
						}},
						Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.Ident{Name: "TimeZoneInfo!Error"}}}},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/time"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
			},
		},
	},
	TypesInfo: &types.Info{
		Defs: map[*ast.Ident]types.Object{
			// Using &moonType{} and &moonFunc{} is a hack to fake a struct/func for testing purposes only:
			{Name: "Cleanup"}: &moonFunc{funcName: "type Cleanup = struct{address Int; size Int; align Int}"},
			{Name: "Person"}:  &moonFunc{funcName: "type Person = struct{firstName String; lastName String; age Int}"},
			// Note that this is a forward reference that is later ignored thanks to the StructLookup map.
			{Name: "Person"}:       &moonFunc{funcName: "type Person = struct{}"},
			{Name: "TimeZoneInfo"}: &moonFunc{funcName: "type TimeZoneInfo = struct{standard_name String; standard_offset String; daylight_name String; daylight_offset String}"},
			// Note that this is a forward reference that is later ignored thanks to the StructLookup map.
			{Name: "TimeZoneInfo"}:                                       &moonFunc{funcName: "type TimeZoneInfo = struct{}"},
			{Name: "__modus_add"}:                                        &moonFunc{funcName: "func @test-suite.__modus_add(x Int, y Int) Int"},
			{Name: "__modus_add3"}:                                       &moonFunc{funcName: "func @test-suite.__modus_add3(a Int, b Int, c Int) Int"},
			{Name: "__modus_add3_WithDefaults"}:                          &moonFunc{funcName: "func @test-suite.__modus_add3_WithDefaults(a Int, b Int) Int"},
			{Name: "__modus_get_current_time"}:                           &moonFunc{funcName: "func @test-suite.__modus_get_current_time() @time.ZonedDateTime!Error"},
			{Name: "__modus_get_current_time_formatted"}:                 &moonFunc{funcName: "func @test-suite.__modus_get_current_time_formatted() String!Error"},
			{Name: "__modus_get_full_name"}:                              &moonFunc{funcName: "func @test-suite.__modus_get_full_name(first_name String, last_name String) String"},
			{Name: "__modus_get_local_time"}:                             &moonFunc{funcName: "func @test-suite.__modus_get_local_time() String!Error"},
			{Name: "__modus_get_local_time_zone_id"}:                     &moonFunc{funcName: "func @test-suite.__modus_get_local_time_zone_id() String"},
			{Name: "__modus_get_people"}:                                 &moonFunc{funcName: "func @test-suite.__modus_get_people() Array[Person]"},
			{Name: "__modus_get_person"}:                                 &moonFunc{funcName: "func @test-suite.__modus_get_person() Person"},
			{Name: "__modus_get_random_person"}:                          &moonFunc{funcName: "func @test-suite.__modus_get_random_person() Person"},
			{Name: "__modus_get_time_in_zone"}:                           &moonFunc{funcName: "func @test-suite.__modus_get_time_in_zone(tz String) String!Error"},
			{Name: "__modus_get_time_zone_info"}:                         &moonFunc{funcName: "func @test-suite.__modus_get_time_zone_info(tz String) TimeZoneInfo!Error"},
			{Name: "__modus_get_utc_time"}:                               &moonFunc{funcName: "func @test-suite.__modus_get_utc_time() @time.ZonedDateTime!Error"},
			{Name: "__modus_hello_array_of_ints"}:                        &moonFunc{funcName: "func @test-suite.__modus_hello_array_of_ints(n Int) Array[Int]"},
			{Name: "__modus_hello_array_of_ints_WithDefaults"}:           &moonFunc{funcName: "func @test-suite.__modus_hello_array_of_ints_WithDefaults() Array[Int]"},
			{Name: "__modus_hello_array_of_ints_option"}:                 &moonFunc{funcName: "func @test-suite.__modus_hello_array_of_ints_option(n Int) Array[Int]?"},
			{Name: "__modus_hello_array_of_ints_option_WithDefaults"}:    &moonFunc{funcName: "func @test-suite.__modus_hello_array_of_ints_option_WithDefaults() Array[Int]?"},
			{Name: "__modus_hello_array_of_strings"}:                     &moonFunc{funcName: "func @test-suite.__modus_hello_array_of_strings(n Int) Array[String]"},
			{Name: "__modus_hello_array_of_strings_WithDefaults"}:        &moonFunc{funcName: "func @test-suite.__modus_hello_array_of_strings_WithDefaults() Array[String]"},
			{Name: "__modus_hello_array_of_strings_option"}:              &moonFunc{funcName: "func @test-suite.__modus_hello_array_of_strings_option(n Int) Array[String]?"},
			{Name: "__modus_hello_array_of_strings_option_WithDefaults"}: &moonFunc{funcName: "func @test-suite.__modus_hello_array_of_strings_option_WithDefaults() Array[String]?"},
			{Name: "__modus_hello_maps_n_items"}:                         &moonFunc{funcName: "func @test-suite.__modus_hello_maps_n_items(n Int) Map[String, String]"},
			{Name: "__modus_hello_maps_n_items_WithDefaults"}:            &moonFunc{funcName: "func @test-suite.__modus_hello_maps_n_items_WithDefaults() Map[String, String]"},
			{Name: "__modus_hello_maps_n_items_option"}:                  &moonFunc{funcName: "func @test-suite.__modus_hello_maps_n_items_option(n Int) Map[String, String]?"},
			{Name: "__modus_hello_maps_n_items_option_WithDefaults"}:     &moonFunc{funcName: "func @test-suite.__modus_hello_maps_n_items_option_WithDefaults() Map[String, String]?"},
			{Name: "__modus_hello_option_empty_string"}:                  &moonFunc{funcName: "func @test-suite.__modus_hello_option_empty_string(some Bool) String?"},
			{Name: "__modus_hello_option_empty_string_WithDefaults"}:     &moonFunc{funcName: "func @test-suite.__modus_hello_option_empty_string_WithDefaults() String?"},
			{Name: "__modus_hello_option_none"}:                          &moonFunc{funcName: "func @test-suite.__modus_hello_option_none(some Bool) String?"},
			{Name: "__modus_hello_option_none_WithDefaults"}:             &moonFunc{funcName: "func @test-suite.__modus_hello_option_none_WithDefaults() String?"},
			{Name: "__modus_hello_option_some_string"}:                   &moonFunc{funcName: "func @test-suite.__modus_hello_option_some_string(some Bool) String?"},
			{Name: "__modus_hello_option_some_string_WithDefaults"}:      &moonFunc{funcName: "func @test-suite.__modus_hello_option_some_string_WithDefaults() String?"},
			{Name: "__modus_hello_primitive_bool_max"}:                   &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_bool_max() Bool"},
			{Name: "__modus_hello_primitive_bool_min"}:                   &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_bool_min() Bool"},
			{Name: "__modus_hello_primitive_byte_max"}:                   &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_byte_max() Byte"},
			{Name: "__modus_hello_primitive_byte_min"}:                   &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_byte_min() Byte"},
			{Name: "__modus_hello_primitive_char_max"}:                   &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_char_max() Char"},
			{Name: "__modus_hello_primitive_char_min"}:                   &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_char_min() Char"},
			{Name: "__modus_hello_primitive_double_max"}:                 &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_double_max() Double"},
			{Name: "__modus_hello_primitive_double_min"}:                 &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_double_min() Double"},
			{Name: "__modus_hello_primitive_float_max"}:                  &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_float_max() Float"},
			{Name: "__modus_hello_primitive_float_min"}:                  &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_float_min() Float"},
			{Name: "__modus_hello_primitive_int16_max"}:                  &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_int16_max() Int16"},
			{Name: "__modus_hello_primitive_int16_min"}:                  &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_int16_min() Int16"},
			{Name: "__modus_hello_primitive_int64_max"}:                  &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_int64_max() Int64"},
			{Name: "__modus_hello_primitive_int64_min"}:                  &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_int64_min() Int64"},
			{Name: "__modus_hello_primitive_int_max"}:                    &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_int_max() Int"},
			{Name: "__modus_hello_primitive_int_min"}:                    &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_int_min() Int"},
			{Name: "__modus_hello_primitive_uint16_max"}:                 &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_uint16_max() UInt16"},
			{Name: "__modus_hello_primitive_uint16_min"}:                 &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_uint16_min() UInt16"},
			{Name: "__modus_hello_primitive_uint64_max"}:                 &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_uint64_max() UInt64"},
			{Name: "__modus_hello_primitive_uint64_min"}:                 &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_uint64_min() UInt64"},
			{Name: "__modus_hello_primitive_uint_max"}:                   &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_uint_max() UInt"},
			{Name: "__modus_hello_primitive_uint_min"}:                   &moonFunc{funcName: "func @test-suite.__modus_hello_primitive_uint_min() UInt"},
			{Name: "__modus_hello_world"}:                                &moonFunc{funcName: "func @test-suite.__modus_hello_world() String"},
			{Name: "__modus_hello_world_with_arg"}:                       &moonFunc{funcName: "func @test-suite.__modus_hello_world_with_arg(name String) String"},
			{Name: "__modus_hello_world_with_optional_arg"}:              &moonFunc{funcName: "func @test-suite.__modus_hello_world_with_optional_arg(name String) String"},
			{Name: "__modus_hello_world_with_optional_arg_WithDefaults"}: &moonFunc{funcName: "func @test-suite.__modus_hello_world_with_optional_arg_WithDefaults() String"},
			{Name: "__modus_log_message"}:                                &moonFunc{funcName: "func @test-suite.__modus_log_message(message String)"},
			{Name: "__modus_test_abort"}:                                 &moonFunc{funcName: "func @test-suite.__modus_test_abort()"},
			{Name: "__modus_test_alternative_error"}:                     &moonFunc{funcName: "func @test-suite.__modus_test_alternative_error(input String) String"},
			{Name: "__modus_test_exit"}:                                  &moonFunc{funcName: "func @test-suite.__modus_test_exit()"},
			{Name: "__modus_test_logging"}:                               &moonFunc{funcName: "func @test-suite.__modus_test_logging()"},
			{Name: "__modus_test_normal_error"}:                          &moonFunc{funcName: "func @test-suite.__modus_test_normal_error(input String) String!Error"},
			{Name: "add"}:                                                &moonFunc{funcName: "func @test-suite.add(x Int, y Int) Int"},
			{Name: "add3"}:                                               &moonFunc{funcName: "func @test-suite.add3(a Int, b Int, c~ Int) Int"},
			{Name: "add3_WithDefaults"}:                                  &moonFunc{funcName: "func @test-suite.add3_WithDefaults(a Int, b Int) Int"},
			{Name: "cabi_realloc"}:                                       &moonFunc{funcName: "func @test-suite.cabi_realloc(src_offset Int, src_size Int, _dst_alignment Int, dst_size Int) Int"},
			{Name: "copy"}:                                               &moonFunc{funcName: "func @test-suite.copy(dest Int, src Int)"},
			{Name: "get_current_time"}:                                   &moonFunc{funcName: "func @test-suite.get_current_time() @time.ZonedDateTime!Error"},
			{Name: "get_current_time_formatted"}:                         &moonFunc{funcName: "func @test-suite.get_current_time_formatted() String!Error"},
			{Name: "get_full_name"}:                                      &moonFunc{funcName: "func @test-suite.get_full_name(first_name String, last_name String) String"},
			{Name: "get_local_time"}:                                     &moonFunc{funcName: "func @test-suite.get_local_time() String!Error"},
			{Name: "get_local_time_zone_id"}:                             &moonFunc{funcName: "func @test-suite.get_local_time_zone_id() String"},
			{Name: "get_people"}:                                         &moonFunc{funcName: "func @test-suite.get_people() Array[Person]"},
			{Name: "get_person"}:                                         &moonFunc{funcName: "func @test-suite.get_person() Person"},
			{Name: "get_random_person"}:                                  &moonFunc{funcName: "func @test-suite.get_random_person() Person"},
			{Name: "get_time_in_zone"}:                                   &moonFunc{funcName: "func @test-suite.get_time_in_zone(tz String) String!Error"},
			{Name: "get_time_zone_info"}:                                 &moonFunc{funcName: "func @test-suite.get_time_zone_info(tz String) TimeZoneInfo!Error"},
			{Name: "get_utc_time"}:                                       &moonFunc{funcName: "func @test-suite.get_utc_time() @time.ZonedDateTime!Error"},
			{Name: "hello_array_of_ints"}:                                &moonFunc{funcName: "func @test-suite.hello_array_of_ints(n~ Int) Array[Int]"},
			{Name: "hello_array_of_ints_WithDefaults"}:                   &moonFunc{funcName: "func @test-suite.hello_array_of_ints_WithDefaults() Array[Int]"},
			{Name: "hello_array_of_ints_option"}:                         &moonFunc{funcName: "func @test-suite.hello_array_of_ints_option(n~ Int) Array[Int]?"},
			{Name: "hello_array_of_ints_option_WithDefaults"}:            &moonFunc{funcName: "func @test-suite.hello_array_of_ints_option_WithDefaults() Array[Int]?"},
			{Name: "hello_array_of_strings"}:                             &moonFunc{funcName: "func @test-suite.hello_array_of_strings(n~ Int) Array[String]"},
			{Name: "hello_array_of_strings_WithDefaults"}:                &moonFunc{funcName: "func @test-suite.hello_array_of_strings_WithDefaults() Array[String]"},
			{Name: "hello_array_of_strings_option"}:                      &moonFunc{funcName: "func @test-suite.hello_array_of_strings_option(n~ Int) Array[String]?"},
			{Name: "hello_array_of_strings_option_WithDefaults"}:         &moonFunc{funcName: "func @test-suite.hello_array_of_strings_option_WithDefaults() Array[String]?"},
			{Name: "hello_maps_n_items"}:                                 &moonFunc{funcName: "func @test-suite.hello_maps_n_items(n~ Int) Map[String, String]"},
			{Name: "hello_maps_n_items_WithDefaults"}:                    &moonFunc{funcName: "func @test-suite.hello_maps_n_items_WithDefaults() Map[String, String]"},
			{Name: "hello_maps_n_items_option"}:                          &moonFunc{funcName: "func @test-suite.hello_maps_n_items_option(n~ Int) Map[String, String]?"},
			{Name: "hello_maps_n_items_option_WithDefaults"}:             &moonFunc{funcName: "func @test-suite.hello_maps_n_items_option_WithDefaults() Map[String, String]?"},
			{Name: "hello_option_empty_string"}:                          &moonFunc{funcName: "func @test-suite.hello_option_empty_string(some~ Bool) String?"},
			{Name: "hello_option_empty_string_WithDefaults"}:             &moonFunc{funcName: "func @test-suite.hello_option_empty_string_WithDefaults() String?"},
			{Name: "hello_option_none"}:                                  &moonFunc{funcName: "func @test-suite.hello_option_none(some~ Bool) String?"},
			{Name: "hello_option_none_WithDefaults"}:                     &moonFunc{funcName: "func @test-suite.hello_option_none_WithDefaults() String?"},
			{Name: "hello_option_some_string"}:                           &moonFunc{funcName: "func @test-suite.hello_option_some_string(some~ Bool) String?"},
			{Name: "hello_option_some_string_WithDefaults"}:              &moonFunc{funcName: "func @test-suite.hello_option_some_string_WithDefaults() String?"},
			{Name: "hello_primitive_bool_max"}:                           &moonFunc{funcName: "func @test-suite.hello_primitive_bool_max() Bool"},
			{Name: "hello_primitive_bool_min"}:                           &moonFunc{funcName: "func @test-suite.hello_primitive_bool_min() Bool"},
			{Name: "hello_primitive_byte_max"}:                           &moonFunc{funcName: "func @test-suite.hello_primitive_byte_max() Byte"},
			{Name: "hello_primitive_byte_min"}:                           &moonFunc{funcName: "func @test-suite.hello_primitive_byte_min() Byte"},
			{Name: "hello_primitive_char_max"}:                           &moonFunc{funcName: "func @test-suite.hello_primitive_char_max() Char"},
			{Name: "hello_primitive_char_min"}:                           &moonFunc{funcName: "func @test-suite.hello_primitive_char_min() Char"},
			{Name: "hello_primitive_double_max"}:                         &moonFunc{funcName: "func @test-suite.hello_primitive_double_max() Double"},
			{Name: "hello_primitive_double_min"}:                         &moonFunc{funcName: "func @test-suite.hello_primitive_double_min() Double"},
			{Name: "hello_primitive_float_max"}:                          &moonFunc{funcName: "func @test-suite.hello_primitive_float_max() Float"},
			{Name: "hello_primitive_float_min"}:                          &moonFunc{funcName: "func @test-suite.hello_primitive_float_min() Float"},
			{Name: "hello_primitive_int16_max"}:                          &moonFunc{funcName: "func @test-suite.hello_primitive_int16_max() Int16"},
			{Name: "hello_primitive_int16_min"}:                          &moonFunc{funcName: "func @test-suite.hello_primitive_int16_min() Int16"},
			{Name: "hello_primitive_int64_max"}:                          &moonFunc{funcName: "func @test-suite.hello_primitive_int64_max() Int64"},
			{Name: "hello_primitive_int64_min"}:                          &moonFunc{funcName: "func @test-suite.hello_primitive_int64_min() Int64"},
			{Name: "hello_primitive_int_max"}:                            &moonFunc{funcName: "func @test-suite.hello_primitive_int_max() Int"},
			{Name: "hello_primitive_int_min"}:                            &moonFunc{funcName: "func @test-suite.hello_primitive_int_min() Int"},
			{Name: "hello_primitive_uint16_max"}:                         &moonFunc{funcName: "func @test-suite.hello_primitive_uint16_max() UInt16"},
			{Name: "hello_primitive_uint16_min"}:                         &moonFunc{funcName: "func @test-suite.hello_primitive_uint16_min() UInt16"},
			{Name: "hello_primitive_uint64_max"}:                         &moonFunc{funcName: "func @test-suite.hello_primitive_uint64_max() UInt64"},
			{Name: "hello_primitive_uint64_min"}:                         &moonFunc{funcName: "func @test-suite.hello_primitive_uint64_min() UInt64"},
			{Name: "hello_primitive_uint_max"}:                           &moonFunc{funcName: "func @test-suite.hello_primitive_uint_max() UInt"},
			{Name: "hello_primitive_uint_min"}:                           &moonFunc{funcName: "func @test-suite.hello_primitive_uint_min() UInt"},
			{Name: "hello_world"}:                                        &moonFunc{funcName: "func @test-suite.hello_world() String"},
			{Name: "hello_world_with_arg"}:                               &moonFunc{funcName: "func @test-suite.hello_world_with_arg(name String) String"},
			{Name: "hello_world_with_optional_arg"}:                      &moonFunc{funcName: "func @test-suite.hello_world_with_optional_arg(name~ String) String"},
			{Name: "hello_world_with_optional_arg_WithDefaults"}:         &moonFunc{funcName: "func @test-suite.hello_world_with_optional_arg_WithDefaults() String"},
			{Name: "log_message"}:                                        &moonFunc{funcName: "func @test-suite.log_message(message String)"},
			{Name: "malloc"}:                                             &moonFunc{funcName: "func @test-suite.malloc(size Int) Int"},
			{Name: "ptr2double_array"}:                                   &moonFunc{funcName: "func @test-suite.ptr2double_array(ptr Int) FixedArray[Double]"},
			{Name: "ptr2int64_array"}:                                    &moonFunc{funcName: "func @test-suite.ptr2int64_array(ptr Int) FixedArray[Int64]"},
			{Name: "ptr2uint64_array"}:                                   &moonFunc{funcName: "func @test-suite.ptr2uint64_array(ptr Int) FixedArray[UInt64]"},
			{Name: "test_abort"}:                                         &moonFunc{funcName: "func @test-suite.test_abort()"},
			{Name: "test_alternative_error"}:                             &moonFunc{funcName: "func @test-suite.test_alternative_error(input String) String"},
			{Name: "test_exit"}:                                          &moonFunc{funcName: "func @test-suite.test_exit()"},
			{Name: "test_logging"}:                                       &moonFunc{funcName: "func @test-suite.test_logging()"},
			{Name: "test_normal_error"}:                                  &moonFunc{funcName: "func @test-suite.test_normal_error(input String) String!Error"},
		},
	},
}
