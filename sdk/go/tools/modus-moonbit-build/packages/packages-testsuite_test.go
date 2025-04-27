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
	dir := "../testdata/test-suite"
	testPackageLoadHelper(t, "testsuite", dir, wantPackageTestsuite)
}

var wantPackageTestsuite = &Package{
	MoonPkgJSON: MoonPkgJSON{
		IsMain: false,
		Imports: []json.RawMessage{
			json.RawMessage(`"gmlewis/modus/pkg/console"`),
			json.RawMessage(`"gmlewis/modus/pkg/localtime"`),
			json.RawMessage(`"gmlewis/modus/wit/interface/wasi"`),
			json.RawMessage(`"moonbitlang/x/sys"`),
			json.RawMessage(`"moonbitlang/x/time"`),
		},
		Targets: map[string][]string{
			"modus_post_generated.mbt": {"wasm"},
		},
		LinkTargets: map[string]*LinkTarget{
			"wasm": {
				Exports: []string{
					"__modus_add3:add3",
					"__modus_add3_WithDefaults:add3_WithDefaults",
					"__modus_add:add",
					"__modus_call_test_time_option_input_some:call_test_time_option_input_some",
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
					"__modus_make_test_map:make_test_map",
					"__modus_test_abort:test_abort",
					"__modus_test_alternative_error:test_alternative_error",
					"__modus_test_array_input_string:test_array_input_string",
					"__modus_test_array_output_int_option:test_array_output_int_option",
					"__modus_test_array_output_string_option:test_array_output_string_option",
					"__modus_test_exit:test_exit",
					"__modus_test_iterate_map_string_string:test_iterate_map_string_string",
					"__modus_test_logging:test_logging",
					"__modus_test_map_input_int_double:test_map_input_int_double",
					"__modus_test_map_input_int_float:test_map_input_int_float",
					"__modus_test_map_input_string_string:test_map_input_string_string",
					"__modus_test_map_lookup_string_string:test_map_lookup_string_string",
					"__modus_test_map_option_input_string_string:test_map_option_input_string_string",
					"__modus_test_map_option_output_string_string:test_map_option_output_string_string",
					"__modus_test_map_output_int_double:test_map_output_int_double",
					"__modus_test_map_output_int_float:test_map_output_int_float",
					"__modus_test_map_output_string_string:test_map_output_string_string",
					"__modus_test_normal_error:test_normal_error",
					"__modus_test_struct_containing_map_input_string_string:test_struct_containing_map_input_string_string",
					"__modus_test_struct_containing_map_output_string_string:test_struct_containing_map_output_string_string",
					"__modus_test_time_input:test_time_input",
					"__modus_test_time_option_input:test_time_option_input",
					"__modus_test_tuple_output:test_tuple_output",
					"__modus_test_tuple_simulator:test_tuple_simulator",
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
	ID:      "moonbit-main",
	Name:    "main",
	PkgPath: "",
	MoonBitFiles: []string{
		"../testdata/test-suite/arrays.mbt",
		"../testdata/test-suite/maps.mbt",
		"../testdata/test-suite/primitives.mbt",
		"../testdata/test-suite/simple.mbt",
		"../testdata/test-suite/strings.mbt",
		"../testdata/test-suite/time.mbt",
		"../testdata/test-suite/time2.mbt",
		"../testdata/test-suite/tuples.mbt",
	},
	StructLookup: map[string]*ast.TypeSpec{
		"Person": {Name: &ast.Ident{Name: "Person"}, Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{{Name: "firstName"}}, Type: &ast.Ident{Name: "String"}},
					{Names: []*ast.Ident{{Name: "lastName"}}, Type: &ast.Ident{Name: "String"}},
					{Names: []*ast.Ident{{Name: "age"}}, Type: &ast.Ident{Name: "Int"}},
				},
			},
		}},
		"TestStructWithMap": {
			Name: &ast.Ident{Name: "TestStructWithMap"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]"}},
			}}},
		},
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
		"TupleSimulator": {
			Name: &ast.Ident{Name: "TupleSimulator"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "t0"}}, Type: &ast.Ident{Name: "Int"}},
				{Names: []*ast.Ident{{Name: "t1"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "t2"}}, Type: &ast.Ident{Name: "String"}},
			}}},
		},
	},
	PossiblyMissingUnderlyingTypes: map[string]struct{}{
		"Bool":                {},
		"Int":                 {},
		"Map[String, String]": {},
		"String":              {},
	},
	Syntax: []*ast.File{
		{
			Name: &ast.Ident{Name: "../testdata/test-suite/arrays.mbt"},
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
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_string_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/test-suite/maps.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStructWithMap"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]"}},
									},
								},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_maps_n_items"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n~"}}, Type: &ast.Ident{Name: "Int"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:3`"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
					},
				},
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
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n~"}}, Type: &ast.Ident{Name: "Int"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:3`"}},
							},
						},
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
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_input_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_option_input_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_output_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_option_output_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_iterate_map_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_lookup_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]"}},
								{Names: []*ast.Ident{{Name: "key"}}, Type: &ast.Ident{Name: "String"}},
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
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_containing_map_input_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "s"}}, Type: &ast.Ident{Name: "TestStructWithMap"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_containing_map_output_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStructWithMap"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_input_int_float"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[Int, Float]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_output_int_float"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[Int, Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_input_int_double"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[Int, Double]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_output_int_double"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[Int, Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// This generated map mimics the test map created on the Go side."},
						},
					},
					Name: &ast.Ident{Name: "make_test_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "size"}}, Type: &ast.Ident{Name: "Int"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/test-suite/primitives.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_bool_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Bool"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_bool_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Bool"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_byte_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Byte"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_byte_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Byte"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_char_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Char"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_char_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Char"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_double_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Double"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_double_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Double"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_float_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Float"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_float_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Float"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int16_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int16"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int16_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int16"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int64_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int64"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_int64_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int64"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint16_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt16"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint16_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt16"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint64_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt64"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_primitive_uint64_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt64"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/test-suite/simple.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "Person"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "firstName"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "lastName"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "age"}}, Type: &ast.Ident{Name: "Int"}},
									},
								},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Logs a message."},
						},
					},
					Name: &ast.Ident{Name: "log_message"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Adds two integers together and returns the result."},
						},
					},
					Name: &ast.Ident{Name: "add"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "x"}}, Type: &ast.Ident{Name: "Int"}},
								{Names: []*ast.Ident{{Name: "y"}}, Type: &ast.Ident{Name: "Int"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Adds three integers together and returns the result."},
							{Text: "// The third integer is optional."},
						},
					},
					Name: &ast.Ident{Name: "add3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Int"}},
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
								{Names: []*ast.Ident{{Name: "c~"}}, Type: &ast.Ident{Name: "Int"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:0`"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Adds three integers together and returns the result."},
							{Text: "// The third integer is optional."},
						},
					},
					Name: &ast.Ident{Name: "add3_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Int"}},
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Returns the current time."},
						},
					},
					Name: &ast.Ident{Name: "get_current_time"},
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
							{Text: "// Returns the current time formatted as a string."},
						},
					},
					Name: &ast.Ident{Name: "get_current_time_formatted"},
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
							{Text: "// Combines the first and last name of a person, and returns the full name."},
						},
					},
					Name: &ast.Ident{Name: "get_full_name"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "first_name"}}, Type: &ast.Ident{Name: "String"}},
								{Names: []*ast.Ident{{Name: "last_name"}}, Type: &ast.Ident{Name: "String"}},
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
							{Text: "// Gets a person object."},
						},
					},
					Name: &ast.Ident{Name: "get_person"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Person"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Gets a random person object from a list of people."},
						},
					},
					Name: &ast.Ident{Name: "get_random_person"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Person"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Gets a list of people."},
						},
					},
					Name: &ast.Ident{Name: "get_people"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Person]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Tests returning an error."},
						},
					},
					Name: &ast.Ident{Name: "test_normal_error"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "input"}}, Type: &ast.Ident{Name: "String"}},
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
							{Text: "// Tests an alternative way to handle errors in functions."},
						},
					},
					Name: &ast.Ident{Name: "test_alternative_error"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "input"}}, Type: &ast.Ident{Name: "String"}},
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
							{Text: "// Tests an abort."},
						},
					},
					Name: &ast.Ident{Name: "test_abort"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Tests an exit with a non-zero exit code."},
						},
					},
					Name: &ast.Ident{Name: "test_exit"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "// Tests logging at different levels."},
						},
					},
					Name: &ast.Ident{Name: "test_logging"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/test-suite/strings.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_empty_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "some~"}}, Type: &ast.Ident{Name: "Bool"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:true`"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_empty_string_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "some~"}}, Type: &ast.Ident{Name: "Bool"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:false`"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_none_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_some_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "some~"}}, Type: &ast.Ident{Name: "Bool"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:true`"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_option_some_string_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_world_with_arg"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "name"}}, Type: &ast.Ident{Name: "String"}},
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
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_world_with_optional_arg"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "name~"}}, Type: &ast.Ident{Name: "String"}, Tag: &ast.BasicLit{Kind: token.STRING, Value: "`default:\"世界 🌍 from MoonBit\"`"}},
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
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_world_with_optional_arg_WithDefaults"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{}},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "hello_world"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/test-suite/time.mbt"},
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
							{Text: "// Returns the current local time."},
						},
					},
					Name: &ast.Ident{Name: "get_local_time"},
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
							{Text: "// Returns the current time in a specified time zone."},
						},
					},
					Name: &ast.Ident{Name: "get_time_in_zone"},
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
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/test-suite/time2.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_time_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "t"}}, Type: &ast.Ident{Name: "@time.ZonedDateTime"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_time_option_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "t"}}, Type: &ast.Ident{Name: "@time.ZonedDateTime?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "call_test_time_option_input_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/test-suite/tuples.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TupleSimulator"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "t0"}}, Type: &ast.Ident{Name: "Int"}},
										{Names: []*ast.Ident{{Name: "t1"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "t2"}}, Type: &ast.Ident{Name: "String"}},
									},
								},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_tuple_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "(Int, Bool, String)"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_tuple_simulator"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TupleSimulator"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/localtime"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
	},
	TypesInfo: &types.Info{
		Defs: map[*ast.Ident]types.Object{
			// Using &moonType{} and &moonFunc{} is a hack to fake a struct/func for testing purposes only:
			{Name: "hello_primitive_uint16_min"}:                      &moonFunc{funcName: "func hello_primitive_uint16_min() UInt16"},
			{Name: "hello_primitive_uint16_max"}:                      &moonFunc{funcName: "func hello_primitive_uint16_max() UInt16"},
			{Name: "get_time_in_zone"}:                                &moonFunc{funcName: "func get_time_in_zone(tz String) String!Error"},
			{Name: "hello_array_of_ints_WithDefaults"}:                &moonFunc{funcName: "func hello_array_of_ints_WithDefaults() Array[Int]"},
			{Name: "hello_primitive_int64_max"}:                       &moonFunc{funcName: "func hello_primitive_int64_max() Int64"},
			{Name: "test_exit"}:                                       &moonFunc{funcName: "func test_exit()"},
			{Name: "hello_array_of_ints"}:                             &moonFunc{funcName: "func hello_array_of_ints(n~ Int = 3) Array[Int]"},
			{Name: "test_array_output_int_option"}:                    &moonFunc{funcName: "func test_array_output_int_option() Array[Int?]"},
			{Name: "make_test_map"}:                                   &moonFunc{funcName: "func make_test_map(size Int) Map[String, String]"},
			{Name: "hello_primitive_uint64_min"}:                      &moonFunc{funcName: "func hello_primitive_uint64_min() UInt64"},
			{Name: "get_current_time_formatted"}:                      &moonFunc{funcName: "func get_current_time_formatted() String!Error"},
			{Name: "get_utc_time"}:                                    &moonFunc{funcName: "func get_utc_time() @time.ZonedDateTime!Error"},
			{Name: "test_tuple_simulator"}:                            &moonFunc{funcName: "func test_tuple_simulator() TupleSimulator"},
			{Name: "hello_array_of_strings_option"}:                   &moonFunc{funcName: "func hello_array_of_strings_option(n~ Int = 3) Array[String]?"},
			{Name: "hello_primitive_float_min"}:                       &moonFunc{funcName: "func hello_primitive_float_min() Float"},
			{Name: "hello_primitive_int_max"}:                         &moonFunc{funcName: "func hello_primitive_int_max() Int"},
			{Name: "add3"}:                                            &moonFunc{funcName: "func add3(a Int, b Int, c~ Int = 0) Int"},
			{Name: "add3_WithDefaults"}:                               &moonFunc{funcName: "func add3_WithDefaults(a Int, b Int) Int"},
			{Name: "get_random_person"}:                               &moonFunc{funcName: "func get_random_person() Person"},
			{Name: "hello_option_none"}:                               &moonFunc{funcName: "func hello_option_none(some~ Bool = false) String?"},
			{Name: "hello_option_none_WithDefaults"}:                  &moonFunc{funcName: "func hello_option_none_WithDefaults() String?"},
			{Name: "hello_primitive_char_min"}:                        &moonFunc{funcName: "func hello_primitive_char_min() Char"},
			{Name: "test_time_option_input"}:                          &moonFunc{funcName: "func test_time_option_input(t @time.ZonedDateTime?) Unit!Error"},
			{Name: "hello_array_of_ints_option"}:                      &moonFunc{funcName: "func hello_array_of_ints_option(n~ Int = 3) Array[Int]?"},
			{Name: "test_array_output_string_option"}:                 &moonFunc{funcName: "func test_array_output_string_option() Array[String?]"},
			{Name: "test_map_output_int_double"}:                      &moonFunc{funcName: "func test_map_output_int_double() Map[Int, Double]"},
			{Name: "hello_primitive_float_max"}:                       &moonFunc{funcName: "func hello_primitive_float_max() Float"},
			{Name: "Person"}:                                          &moonFunc{funcName: "type Person = struct{firstName String; lastName String; age Int}"},
			{Name: "get_person"}:                                      &moonFunc{funcName: "func get_person() Person"},
			{Name: "TimeZoneInfo"}:                                    &moonFunc{funcName: "type TimeZoneInfo = struct{standard_name String; standard_offset String; daylight_name String; daylight_offset String}"},
			{Name: "hello_array_of_strings"}:                          &moonFunc{funcName: "func hello_array_of_strings(n~ Int = 3) Array[String]"},
			{Name: "test_map_input_string_string"}:                    &moonFunc{funcName: "func test_map_input_string_string(m Map[String, String]) Unit!Error"},
			{Name: "hello_primitive_bool_max"}:                        &moonFunc{funcName: "func hello_primitive_bool_max() Bool"},
			{Name: "hello_primitive_byte_min"}:                        &moonFunc{funcName: "func hello_primitive_byte_min() Byte"},
			{Name: "hello_primitive_char_max"}:                        &moonFunc{funcName: "func hello_primitive_char_max() Char"},
			{Name: "hello_primitive_uint_max"}:                        &moonFunc{funcName: "func hello_primitive_uint_max() UInt"},
			{Name: "get_people"}:                                      &moonFunc{funcName: "func get_people() Array[Person]"},
			{Name: "hello_maps_n_items"}:                              &moonFunc{funcName: "func hello_maps_n_items(n~ Int = 3) Map[String, String]"},
			{Name: "hello_maps_n_items_option_WithDefaults"}:          &moonFunc{funcName: "func hello_maps_n_items_option_WithDefaults() Map[String, String]?"},
			{Name: "test_struct_containing_map_output_string_string"}: &moonFunc{funcName: "func test_struct_containing_map_output_string_string() TestStructWithMap"},
			{Name: "hello_primitive_byte_max"}:                        &moonFunc{funcName: "func hello_primitive_byte_max() Byte"},
			{Name: "hello_primitive_int_min"}:                         &moonFunc{funcName: "func hello_primitive_int_min() Int"},
			{Name: "test_logging"}:                                    &moonFunc{funcName: "func test_logging()"},
			{Name: "hello_world_with_optional_arg_WithDefaults"}:      &moonFunc{funcName: "func hello_world_with_optional_arg_WithDefaults() String"},
			{Name: "hello_array_of_ints_option_WithDefaults"}:         &moonFunc{funcName: "func hello_array_of_ints_option_WithDefaults() Array[Int]?"},
			{Name: "hello_array_of_strings_WithDefaults"}:             &moonFunc{funcName: "func hello_array_of_strings_WithDefaults() Array[String]"},
			{Name: "test_time_input"}:                                 &moonFunc{funcName: "func test_time_input(t @time.ZonedDateTime) Unit!Error"},
			{Name: "hello_maps_n_items_WithDefaults"}:                 &moonFunc{funcName: "func hello_maps_n_items_WithDefaults() Map[String, String]"},
			{Name: "test_map_option_output_string_string"}:            &moonFunc{funcName: "func test_map_option_output_string_string() Map[String, String]?"},
			{Name: "test_map_input_int_float"}:                        &moonFunc{funcName: "func test_map_input_int_float(m Map[Int, Float]) Unit!Error"},
			{Name: "hello_primitive_uint64_max"}:                      &moonFunc{funcName: "func hello_primitive_uint64_max() UInt64"},
			{Name: "get_current_time"}:                                &moonFunc{funcName: "func get_current_time() @time.ZonedDateTime!Error"},
			{Name: "hello_option_empty_string"}:                       &moonFunc{funcName: "func hello_option_empty_string(some~ Bool = true) String?"},
			{Name: "get_local_time"}:                                  &moonFunc{funcName: "func get_local_time() String!Error"},
			{Name: "test_map_output_string_string"}:                   &moonFunc{funcName: "func test_map_output_string_string() Map[String, String]"},
			{Name: "get_full_name"}:                                   &moonFunc{funcName: "func get_full_name(first_name String, last_name String) String"},
			{Name: "test_abort"}:                                      &moonFunc{funcName: "func test_abort()"},
			{Name: "hello_option_some_string_WithDefaults"}:           &moonFunc{funcName: "func hello_option_some_string_WithDefaults() String?"},
			{Name: "get_time_zone_info"}:                              &moonFunc{funcName: "func get_time_zone_info(tz String) TimeZoneInfo!Error"},
			{Name: "TupleSimulator"}:                                  &moonFunc{funcName: "type TupleSimulator = struct{t0 Int; t1 Bool; t2 String}"},
			{Name: "test_tuple_output"}:                               &moonFunc{funcName: "func test_tuple_output() (Int, Bool, String)"},
			{Name: "test_array_input_string"}:                         &moonFunc{funcName: "func test_array_input_string(val Array[String]) Unit!Error"},
			{Name: "test_map_lookup_string_string"}:                   &moonFunc{funcName: "func test_map_lookup_string_string(m Map[String, String], key String) String"},
			{Name: "test_struct_containing_map_input_string_string"}:  &moonFunc{funcName: "func test_struct_containing_map_input_string_string(s TestStructWithMap) Unit!Error"},
			{Name: "hello_primitive_double_max"}:                      &moonFunc{funcName: "func hello_primitive_double_max() Double"},
			{Name: "hello_primitive_int16_max"}:                       &moonFunc{funcName: "func hello_primitive_int16_max() Int16"},
			{Name: "hello_primitive_uint_min"}:                        &moonFunc{funcName: "func hello_primitive_uint_min() UInt"},
			{Name: "test_normal_error"}:                               &moonFunc{funcName: "func test_normal_error(input String) String!Error"},
			{Name: "hello_option_some_string"}:                        &moonFunc{funcName: "func hello_option_some_string(some~ Bool = true) String?"},
			{Name: "test_iterate_map_string_string"}:                  &moonFunc{funcName: "func test_iterate_map_string_string(m Map[String, String])"},
			{Name: "hello_primitive_int64_min"}:                       &moonFunc{funcName: "func hello_primitive_int64_min() Int64"},
			{Name: "add"}:                                             &moonFunc{funcName: "func add(x Int, y Int) Int"},
			{Name: "hello_world_with_optional_arg"}:                   &moonFunc{funcName: `func hello_world_with_optional_arg(name~ String = "世界 🌍 from MoonBit") String`},
			{Name: "get_local_time_zone_id"}:                          &moonFunc{funcName: "func get_local_time_zone_id() String"},
			{Name: "call_test_time_option_input_some"}:                &moonFunc{funcName: "func call_test_time_option_input_some() Unit!Error"},
			{Name: "hello_primitive_int16_min"}:                       &moonFunc{funcName: "func hello_primitive_int16_min() Int16"},
			{Name: "hello_maps_n_items_option"}:                       &moonFunc{funcName: "func hello_maps_n_items_option(n~ Int = 3) Map[String, String]?"},
			{Name: "hello_primitive_bool_min"}:                        &moonFunc{funcName: "func hello_primitive_bool_min() Bool"},
			{Name: "hello_primitive_double_min"}:                      &moonFunc{funcName: "func hello_primitive_double_min() Double"},
			{Name: "log_message"}:                                     &moonFunc{funcName: "func log_message(message String)"},
			{Name: "test_alternative_error"}:                          &moonFunc{funcName: "func test_alternative_error(input String) String"},
			{Name: "hello_option_empty_string_WithDefaults"}:          &moonFunc{funcName: "func hello_option_empty_string_WithDefaults() String?"},
			{Name: "hello_world_with_arg"}:                            &moonFunc{funcName: "func hello_world_with_arg(name String) String"},
			{Name: "hello_world"}:                                     &moonFunc{funcName: "func hello_world() String"},
			{Name: "hello_array_of_strings_option_WithDefaults"}:      &moonFunc{funcName: "func hello_array_of_strings_option_WithDefaults() Array[String]?"},
			{Name: "TestStructWithMap"}:                               &moonFunc{funcName: "type TestStructWithMap = struct{m Map[String, String]}"},
			{Name: "test_map_option_input_string_string"}:             &moonFunc{funcName: "func test_map_option_input_string_string(m Map[String, String]?) Unit!Error"},
			{Name: "test_map_output_int_float"}:                       &moonFunc{funcName: "func test_map_output_int_float() Map[Int, Float]"},
			{Name: "test_map_input_int_double"}:                       &moonFunc{funcName: "func test_map_input_int_double(m Map[Int, Double]) Unit!Error"},
		},
	},
}
