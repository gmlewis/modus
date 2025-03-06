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
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"

	"github.com/google/go-cmp/cmp"
)

func TestPackage(t *testing.T) {
	t.Parallel()
	mode := NeedName | NeedImports | NeedDeps | NeedTypes | NeedSyntax | NeedTypesInfo
	dir := "testdata"
	cfg := &Config{Mode: mode, Dir: dir}
	mod := genModuleInfo(t, dir)
	got, err := Load(cfg, mod, ".")
	if err != nil || len(got) != 1 {
		t.Fatalf("len(got)=%v, want 1, err=%v", len(got), err)
	}

	// Manually test the contents of TypesInfo since the map[*ast.Ident]types.Object
	// is not directly comparable.
	if len(got[0].TypesInfo.Defs) != len(wantPackages[0].TypesInfo.Defs) {
		keys := make([]string, 0, len(got[0].TypesInfo.Defs))
		for k := range got[0].TypesInfo.Defs {
			keys = append(keys, k.Name)
		}
		sort.Strings(keys)
		t.Errorf("Load() mismatch: got %v types, want %v:\n%#v", len(got[0].TypesInfo.Defs), len(wantPackages[0].TypesInfo.Defs), keys)
	}
	gotTypesInfoDefs := make(map[string]types.Object)
	for k, v := range got[0].TypesInfo.Defs {
		gotTypesInfoDefs[k.Name] = v
	}
	for k, wantDef := range wantPackages[0].TypesInfo.Defs {
		gotDef := gotTypesInfoDefs[k.Name]
		gotStr := fmt.Sprintf("%v", gotDef)
		wantStr := fmt.Sprintf("%v", wantDef)
		if diff := cmp.Diff(wantStr, gotStr); diff != "" {
			t.Logf("gotDef: %#v", gotDef)
			if gotDef != nil {
				t.Logf("gotDef.Type(): %#v", gotDef.Type())
			}
			// t.Logf("gotDef.Type().Underlying(): %#v", gotDef.Type().Underlying())
			t.Errorf("Load() mismatch for TypesInfo.Defs[%q] (-want +got):\n%v", k, diff)
		}
	}

	got[0].TypesInfo = nil
	wantPackages[0].TypesInfo = nil

	if diff := cmp.Diff(wantPackages, got); diff != "" {
		t.Errorf("Load() mismatch (-want +got):\n%v", diff)
	}
}

var testdataPkg *types.Package // nil - was: types.NewPackage("@testdata", "main")
var wantPackages = []*Package{
	{
		MoonPkgJSON: MoonPkgJSON{
			IsMain: false,
			Imports: []json.RawMessage{
				json.RawMessage(`"gmlewis/modus/pkg/console"`),
				json.RawMessage(`"gmlewis/modus/wit/interface/wasi"`),
				json.RawMessage(`"moonbitlang/x/sys"`),
				json.RawMessage(`"moonbitlang/x/time"`),
			},
			TestImport: []json.RawMessage{json.RawMessage(`"gmlewis/modus/pkg/testutils"`)},
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
						"__modus_add_n:add_n",
						"__modus_get_current_time:get_current_time",
						"__modus_get_current_time_formatted:get_current_time_formatted",
						"__modus_get_full_name:get_full_name",
						"__modus_get_people:get_people",
						"__modus_get_person:get_person",
						"__modus_get_random_person:get_random_person",
						"__modus_log_message:log_message",
						"__modus_test_abort:test_abort",
						"__modus_test_alternative_error:test_alternative_error",
						"__modus_test_exit:test_exit",
						"__modus_test_logging:test_logging",
						"__modus_test_normal_error:test_normal_error",
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
		MoonBitFiles: []string{"testdata/simple-example.mbt"},
		ID:           "moonbit-main",
		// Name:         "main",
		// PkgPath:      "@testdata",
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
		},
		Syntax: []*ast.File{
			{
				Name: &ast.Ident{Name: "testdata/simple-example.mbt"},
				Decls: []ast.Decl{
					&ast.GenDecl{
						Tok: token.TYPE,
						Specs: []ast.Spec{&ast.TypeSpec{Name: &ast.Ident{Name: "Person"}, Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "firstName"}}, Type: &ast.Ident{Name: "String"}},
									{Names: []*ast.Ident{{Name: "lastName"}}, Type: &ast.Ident{Name: "String"}},
									{Names: []*ast.Ident{{Name: "age"}}, Type: &ast.Ident{Name: "Int"}},
								},
							},
						}}},
					},
					&ast.FuncDecl{
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Logs a message."}}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Adds two integers together and returns the result."}}},
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
						Doc: &ast.CommentGroup{List: []*ast.Comment{
							{Text: "// Adds three integers together and returns the result."},
							{Text: "// The third integer is optional."},
						}},
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
						Doc: &ast.CommentGroup{List: []*ast.Comment{
							{Text: "// Adds three integers together and returns the result."},
							{Text: "// The third integer is optional."},
						}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Adds any number of integers together and returns the result."}}},
						Name: &ast.Ident{Name: "add_n"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "args"}}, Type: &ast.Ident{Name: "Array[Int]"}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Returns the current time."}}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Returns the current time formatted as a string."}}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Combines the first and last name of a person, and returns the full name."}}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Gets a person object."}}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Gets a random person object from a list of people."}}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Gets a list of people."}}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Tests returning an error."}}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Tests an alternative way to handle errors in functions."}}},
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
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Tests an abort."}}},
						Name: &ast.Ident{Name: "test_abort"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{},
						},
					},
					&ast.FuncDecl{
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Tests an exit with a non-zero exit code."}}},
						Name: &ast.Ident{Name: "test_exit"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{},
						},
					},
					&ast.FuncDecl{
						Doc:  &ast.CommentGroup{List: []*ast.Comment{{Text: "// Tests logging at different levels."}}},
						Name: &ast.Ident{Name: "test_logging"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{},
						},
					},
				},
				Imports: []*ast.ImportSpec{
					{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
					{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/interface/wasi"`}},
					{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
					{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
				},
			},
		},
		TypesInfo: &types.Info{
			Defs: map[*ast.Ident]types.Object{
				// Using &moonType{} is a hack to fake a struct for testing purposes only:
				{Name: "Person"}: types.NewTypeName(0, nil, "Person", &moonType{typeName: "struct{firstName String; lastName String; age Int}"}),
				{Name: "add"}: types.NewFunc(0, testdataPkg, "add", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "x", &moonType{typeName: "Int"}),
					types.NewVar(0, nil, "y", &moonType{typeName: "Int"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "Int"})), false)),
				{Name: "add3"}: types.NewFunc(0, testdataPkg, "add3", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "a", &moonType{typeName: "Int"}),
					types.NewVar(0, nil, "b", &moonType{typeName: "Int"}),
					types.NewVar(0, nil, "c~", &moonType{typeName: "Int"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "Int"})), false)),
				{Name: "add3_WithDefaults"}: types.NewFunc(0, testdataPkg, "add3_WithDefaults", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "a", &moonType{typeName: "Int"}),
					types.NewVar(0, nil, "b", &moonType{typeName: "Int"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "Int"})), false)),
				{Name: "add_n"}: types.NewFunc(0, testdataPkg, "add_n", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "args", &moonType{typeName: "Array[Int]"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "Int"})), false)),
				{Name: "get_current_time"}:           types.NewFunc(0, testdataPkg, "get_current_time", types.NewSignatureType(nil, nil, nil, types.NewTuple(), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "@time.ZonedDateTime!Error"})), false)),
				{Name: "get_current_time_formatted"}: types.NewFunc(0, testdataPkg, "get_current_time_formatted", types.NewSignatureType(nil, nil, nil, types.NewTuple(), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String!Error"})), false)),
				{Name: "get_full_name"}: types.NewFunc(0, testdataPkg, "get_full_name", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "first_name", &moonType{typeName: "String"}),
					types.NewVar(0, nil, "last_name", &moonType{typeName: "String"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String"})), false)),
				{Name: "get_people"}: types.NewFunc(0, testdataPkg, "get_people", types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "Array[Person]"})), false)),
				{Name: "get_person"}: types.NewFunc(0, testdataPkg, "get_person", types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "Person"})), false)),
				{Name: "get_random_person"}: types.NewFunc(0, testdataPkg, "get_random_person", types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(types.NewVar(0, testdataPkg, "", &moonType{typeName: "Person"})), false)),
				{Name: "log_message"}: types.NewFunc(0, testdataPkg, "log_message", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "message", &moonType{typeName: "String"}),
				), nil, false)),
				{Name: "test_abort"}: types.NewFunc(0, testdataPkg, "test_abort", types.NewSignatureType(nil, nil, nil, nil, nil, false)),
				{Name: "test_alternative_error"}: types.NewFunc(0, testdataPkg, "test_alternative_error", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "input", &moonType{typeName: "String"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String"})), false)),
				{Name: "test_exit"}:    types.NewFunc(0, testdataPkg, "test_exit", types.NewSignatureType(nil, nil, nil, nil, nil, false)),
				{Name: "test_logging"}: types.NewFunc(0, testdataPkg, "test_logging", types.NewSignatureType(nil, nil, nil, nil, nil, false)),
				{Name: "test_normal_error"}: types.NewFunc(0, testdataPkg, "test_normal_error", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "input", &moonType{typeName: "String"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String!Error"})), false)),
			},
		},
	},
}

// genModuleInfo is used during 'go generate' and also during testing.
// To avoid cyclic imports, multiple copies of this function exist. :-(
func genModuleInfo(t *testing.T, sourceDir string) *modinfo.ModuleInfo {
	t.Helper()
	log.SetFlags(0)
	config := &config.Config{
		SourceDir: sourceDir,
	}

	// Make sure the ".mooncakes" directory is initialized before running the test.
	mooncakesDir := path.Join(config.SourceDir, ".mooncakes")
	if _, err := os.Stat(mooncakesDir); err != nil {
		// run the "moon check" command in that directory to initialize it.
		args := []string{"moon", "check", "--directory", config.SourceDir}
		buf, err := exec.Command(args[0], args[1:]...).CombinedOutput()
		if err != nil {
			log.Fatalf("error running %q: %v\n%s", args, err, buf)
		}
	}

	mod, err := modinfo.CollectModuleInfo(config)
	if err != nil {
		log.Fatalf("CollectModuleInfo returned an error: %v", err)
	}

	return mod
}
