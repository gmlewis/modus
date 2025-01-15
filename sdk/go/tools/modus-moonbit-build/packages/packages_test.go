package packages

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPackage(t *testing.T) {
	t.Parallel()
	mode := NeedName | NeedImports | NeedDeps | NeedTypes | NeedSyntax | NeedTypesInfo
	dir := "testdata"
	cfg := &Config{Mode: mode, Dir: dir}
	got, err := Load(cfg, ".")
	if err != nil || len(got) != 1 {
		t.Fatal(err)
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
			t.Logf("gotDef.Type(): %#v", gotDef.Type())
			t.Logf("gotDef.Type().Underlying(): %#v", gotDef.Type().Underlying())
			t.Errorf("Load() mismatch for TypesInfo.Defs[%q] (-want +got):\n%v", k, diff)
		}
	}

	got[0].TypesInfo = nil
	wantPackages[0].TypesInfo = nil

	if diff := cmp.Diff(wantPackages, got); diff != "" {
		t.Errorf("Load() mismatch (-want +got):\n%v", diff)
	}
}

var testdataPkg = types.NewPackage("@testdata", "main")
var wantPackages = []*Package{
	{
		MoonPkgJSON: MoonPkgJSON{
			IsMain: false,
			Imports: []json.RawMessage{
				json.RawMessage(`"gmlewis/modus/pkg/console"`),
				json.RawMessage(`"gmlewis/modus/wit/interface/imports/wasi/clocks/wallClock"`),
				json.RawMessage(`"moonbitlang/x/sys"`),
				json.RawMessage(`"moonbitlang/x/time"`),
			},
			TestImport: []json.RawMessage{json.RawMessage(`"gmlewis/modus/pkg/testutils"`)},
		},
		MoonBitFiles: []string{"testdata/simple-example.mbt"},
		ID:           "moonbit-main",
		Name:         "main",
		PkgPath:      "@testdata",
		Syntax: []*ast.File{
			{
				Name: &ast.Ident{Name: "testdata/simple-example.mbt"},
				Decls: []ast.Decl{
					&ast.GenDecl{
						Tok: token.TYPE,
						Specs: []ast.Spec{&ast.TypeSpec{Name: &ast.Ident{Name: "@testdata.Person"}, Type: &ast.StructType{
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
						Name: &ast.Ident{Name: "add3"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Int"}},
									{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
									{Names: []*ast.Ident{{Name: "c~"}}, Type: &ast.Ident{Name: "Int = 0"}},
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
						Name: &ast.Ident{Name: "get_current_time"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "now~"}}, Type: &ast.Ident{Name: "@wallClock.Datetime = @wallClock.now()"}},
								},
							},
							Results: &ast.FieldList{
								List: []*ast.Field{
									{Type: &ast.Ident{Name: "@time.PlainDateTime!Error"}},
								},
							},
						},
					},
					&ast.FuncDecl{
						Name: &ast.Ident{Name: "get_current_time_formatted"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "now~"}}, Type: &ast.Ident{Name: "@wallClock.Datetime = @wallClock.now()"}},
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
						Name: &ast.Ident{Name: "say_hello"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{Names: []*ast.Ident{{Name: "name~"}}, Type: &ast.Ident{Name: "String? = None"}},
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
						Name: &ast.Ident{Name: "get_name_and_age"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{},
							Results: &ast.FieldList{
								List: []*ast.Field{
									{Type: &ast.Ident{Name: "(String, Int)"}},
								},
							},
						},
					},
					&ast.FuncDecl{
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
						Name: &ast.Ident{Name: "test_abort"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{},
						},
					},
					&ast.FuncDecl{
						Name: &ast.Ident{Name: "test_exit"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{},
						},
					},
					&ast.FuncDecl{
						Name: &ast.Ident{Name: "test_logging"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{},
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
				{Name: "@testdata.Person"}: types.NewTypeName(0, testdataPkg, "Person", &moonType{typeName: "struct{firstName String; lastName String; age Int}"}), // hack
				{Name: "add"}: types.NewFunc(0, testdataPkg, "add", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "x", &moonType{typeName: "Int"}),
					types.NewVar(0, nil, "y", &moonType{typeName: "Int"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "Int"})), false)),
				{Name: "add3"}: types.NewFunc(0, testdataPkg, "add3", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "a", &moonType{typeName: "Int"}),
					types.NewVar(0, nil, "b", &moonType{typeName: "Int"}),
					types.NewVar(0, nil, "c~", &moonType{typeName: "Int = 0"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "Int"})), false)),
				{Name: "add_n"}: types.NewFunc(0, testdataPkg, "add_n", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "args", &moonType{typeName: "Array[Int]"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "Int"})), false)),
				{Name: "get_current_time"}: types.NewFunc(0, testdataPkg, "get_current_time", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "now~", &moonType{typeName: "@wallClock.Datetime = @wallClock.now()"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "@time.PlainDateTime!Error"})), false)),
				{Name: "get_current_time_formatted"}: types.NewFunc(0, testdataPkg, "get_current_time_formatted", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "now~", &moonType{typeName: "@wallClock.Datetime = @wallClock.now()"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String!Error"})), false)),
				{Name: "get_full_name"}: types.NewFunc(0, testdataPkg, "get_full_name", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "first_name", &moonType{typeName: "String"}),
					types.NewVar(0, nil, "last_name", &moonType{typeName: "String"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String"})), false)),
				{Name: "get_name_and_age"}: types.NewFunc(0, testdataPkg, "get_name_and_age", types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "(String, Int)"})), false)),
				{Name: "get_people"}: types.NewFunc(0, testdataPkg, "get_people", types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "Array[Person]"})), false)),
				{Name: "get_person"}: types.NewFunc(0, testdataPkg, "get_person", types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "@testdata.Person"})), false)),
				{Name: "get_random_person"}: types.NewFunc(0, testdataPkg, "get_random_person", types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(types.NewVar(0, testdataPkg, "", &moonType{typeName: "@testdata.Person"})), false)),
				{Name: "log_message"}: types.NewFunc(0, testdataPkg, "log_message", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "message", &moonType{typeName: "String"}),
				), nil, false)),
				{Name: "say_hello"}: types.NewFunc(0, testdataPkg, "say_hello", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "name~", &moonType{typeName: "String? = None"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String"})), false)),
				{Name: "test_abort"}: types.NewFunc(0, testdataPkg, "test_abort", types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String"})), false)),
				{Name: "test_alternative_error"}: types.NewFunc(0, testdataPkg, "test_alternative_error", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "input", &moonType{typeName: "String"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String"})), false)),
				{Name: "test_exit"}: types.NewFunc(0, testdataPkg, "test_exit", types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String"})), false)),
				{Name: "test_logging"}: types.NewFunc(0, testdataPkg, "test_logging", types.NewSignatureType(nil, nil, nil, nil,
					types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String"})), false)),
				{Name: "test_normal_error"}: types.NewFunc(0, testdataPkg, "test_normal_error", types.NewSignatureType(nil, nil, nil, types.NewTuple(
					types.NewVar(0, nil, "input", &moonType{typeName: "String"}),
				), types.NewTuple(types.NewVar(0, nil, "", &moonType{typeName: "String!Error"})), false)),
			},
		},
	},
}
