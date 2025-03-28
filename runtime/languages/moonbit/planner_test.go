/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package moonbit

import (
	"context"
	"log"
	"testing"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/manifestdata"
	"github.com/gmlewis/modus/runtime/utils"
	"github.com/hypermodeinc/modus/lib/manifest"

	"github.com/stretchr/testify/assert"
	wasm "github.com/tetratelabs/wazero/api"
)

func mustGetHandler(t *testing.T, name string) langsupport.TypeHandler {
	t.Helper()
	metadata := &metadata.Metadata{
		Types: map[string]*metadata.TypeDefinition{
			name: {Name: name, Id: 4},
		},
	}

	p := NewPlanner(metadata)
	h, err := p.GetHandler(context.Background(), name)
	if err != nil {
		t.Fatalf("no handler found for %q", name)
	}
	return h
}

func TestGetPlan(t *testing.T) {
	// Copied from runtime/graphql/schemagen/schemagen_mbt_test.go
	log.SetFlags(0)
	manifest := &manifest.Manifest{
		Models:      map[string]manifest.ModelInfo{},
		Connections: map[string]manifest.ConnectionInfo{},
		Collections: map[string]manifest.CollectionInfo{
			"collection1": {
				SearchMethods: map[string]manifest.SearchMethodInfo{
					"search1": {
						Embedder: "my_embedder",
					},
				},
			},
		},
	}
	manifestdata.SetManifest(manifest)

	md := metadata.NewPluginMetadata()
	md.SDK = "modus-sdk-mbt"

	md.FnExports.AddFunction("add").
		WithParameter("a", "Int").
		WithParameter("b", "Int").
		WithResult("Int")

	md.FnExports.AddFunction("add3").
		WithParameter("a", "Int").
		WithParameter("b", "Int").
		WithParameter("c~", "Int", 0).
		WithResult("Int")

	md.FnExports.AddFunction("say_hello").
		WithParameter("name~", "String", "World").
		WithResult("String")

	md.FnExports.AddFunction("current_time").
		WithResult("@time.ZonedDateTime")

	md.FnExports.AddFunction("transform").
		WithParameter("items", "Map[String,String]").
		WithResult("Map[String,String]")

	md.FnExports.AddFunction("test_default_int_params").
		WithParameter("a", "Int").
		WithParameter("b~", "Int", 0).
		WithParameter("c~", "Int", 1)

	md.FnExports.AddFunction("test_default_string_params").
		WithParameter("a", "String").
		WithParameter("b~", "String", "").
		WithParameter("c~", "String", `a"b`).
		WithParameter("d~", "String?").
		WithParameter("e~", "String?", nil).
		WithParameter("f~", "String?", "").
		WithParameter("g~", "String?", "test")

	md.FnExports.AddFunction("test_default_array_params").
		WithParameter("a", "Array[Int]").
		WithParameter("b~", "Array[Int]", []int32{}).
		WithParameter("c~", "Array[Int]", []int32{1, 2, 3}).
		WithParameter("d~", "Array[Int]?").
		WithParameter("e~", "Array[Int]?", nil).
		WithParameter("f~", "Array[Int]?", []int32{}).
		WithParameter("g~", "Array[Int]?", []int32{1, 2, 3})

	md.FnExports.AddFunction("get_person").
		WithResult("Person")

	md.FnExports.AddFunction("list_people").
		WithResult("Array[Person]")

	md.FnExports.AddFunction("add_person").
		WithParameter("person", "Person")

	md.FnExports.AddFunction("get_product_map").
		WithResult("Map[String,Product]")

	md.FnExports.AddFunction("do_nothing")

	md.FnExports.AddFunction("test_options").
		WithParameter("a", "Int?").
		WithParameter("b", "Array[Int?]").
		WithParameter("c", "Array[Int]?").
		WithParameter("d", "Array[Person?]").
		WithParameter("e", "Array[Person]?").
		WithResult("Person?").
		WithDocs(metadata.Docs{
			Lines: []string{
				"This function tests that options are working correctly",
			},
		})

	md.Types.AddType("Int?")
	md.Types.AddType("Array[Int?]")
	md.Types.AddType("Array[Int]?")
	md.Types.AddType("Array[Person?]")
	md.Types.AddType("Array[Person]?")
	md.Types.AddType("Person?")
	md.Types.AddType("String?")
	md.Types.AddType("Array[Address]")
	md.Types.AddType("Array[String]")
	md.Types.AddType("Array[Product]")
	md.Types.AddType("@time.Duration")

	// This should be excluded from the final schema
	md.FnExports.AddFunction("my_embedder").
		WithParameter("text", "String").
		WithResult("Array[Double]")

	// Generated input object from the output object
	md.FnExports.AddFunction("test_obj1").
		WithParameter("obj", "Obj1").
		WithResult("Obj1")

	md.Types.AddType("Obj1").
		WithField("id", "Int").
		WithField("name", "String")

	// Separate input and output objects defined
	md.FnExports.AddFunction("test_obj2").
		WithParameter("obj", "Obj2Input").
		WithResult("Obj2")
	md.Types.AddType("Obj2").
		WithField("id", "Int").
		WithField("name", "String")
	md.Types.AddType("Obj2Input").
		WithField("name", "String")

	// Generated input object without output object
	md.FnExports.AddFunction("test_obj3").
		WithParameter("obj", "Obj3")
	md.Types.AddType("Obj3").
		WithField("name", "String")

	// Single input object defined without output object
	md.FnExports.AddFunction("test_obj4").
		WithParameter("obj", "Obj4Input")
	md.Types.AddType("Obj4Input").
		WithField("name", "String")

	// Test slice of pointers to structs
	md.FnExports.AddFunction("test_obj5").
		WithResult("Array[Obj5?]")
	md.Types.AddType("Array[Obj5?]")
	md.Types.AddType("Obj5?")
	md.Types.AddType("Obj5").
		WithField("name", "String")

	// Test slice of structs
	md.FnExports.AddFunction("test_obj6").
		WithResult("Array[Obj6]")
	md.Types.AddType("Array[Obj6]")
	md.Types.AddType("Obj6").
		WithField("name", "String")

	md.Types.AddType("Array[Int]")
	md.Types.AddType("Array[Double]")
	md.Types.AddType("Array[Person]")
	md.Types.AddType("Map[String,String]")
	md.Types.AddType("Map[String,Product]")

	md.Types.AddType("Company").
		WithField("name", "String")

	md.Types.AddType("Product").
		WithField("name", "String").
		WithField("price", "Double").
		WithField("manufacturer", "Company").
		WithField("components", "Array[Product]")

	md.Types.AddType("Person").
		WithField("name", "String").
		WithField("age", "Int").
		WithField("addresses", "Array[Address]")

	md.Types.AddType("Address").
		WithField("street", "String", &metadata.Docs{
			Lines: []string{
				"Street that the user lives on",
			},
		}).
		WithField("city", "String").
		WithField("state", "String").
		WithField("country", "String", &metadata.Docs{
			Lines: []string{
				"Country that the user is from",
			},
		}).
		WithField("postalCode", "String").
		WithField("location", "Coordinates").
		WithDocs(metadata.Docs{
			Lines: []string{
				"Address represents a physical address.",
				"Each field corresponds to a specific part of the address.",
				"The location field stores geospatial coordinates.",
			},
		})

	md.Types.AddType("Coordinates").
		WithField("lat", "Double").
		WithField("lon", "Double")

	md.FnExports.AddFunction("return_duration").
		WithResult("@time.Duration")

	// This should be excluded from the final schema
	md.Types.AddType("Header").
		WithField("name", "String").
		WithField("values", "Array[String]")

	planner := NewPlanner(md)

	// importsMap := make(map[string]wasm.FunctionDefinition, len(imports))
	// for _, fnDef := range imports {
	// 	if modName, fnName, ok := fnDef.Import(); ok {
	// 		importName := modName + "." + fnName
	// 		importsMap[importName] = fnDef
	// 	}
	// }
	exports := map[string]wasm.FunctionDefinition{
		"add_person":                 &fakeWasmFunc{},
		"add":                        &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI32}},
		"add3":                       &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI32}},
		"current_time":               &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"do_nothing":                 &fakeWasmFunc{},
		"get_person":                 &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"get_product_map":            &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"list_people":                &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"my_embedder":                &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"return_duration":            &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"say_hello":                  &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"test_default_array_params":  &fakeWasmFunc{},
		"test_default_int_params":    &fakeWasmFunc{},
		"test_default_string_params": &fakeWasmFunc{},
		"test_obj1":                  &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"test_obj2":                  &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"test_obj3":                  &fakeWasmFunc{},
		"test_obj4":                  &fakeWasmFunc{},
		"test_obj5":                  &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"test_obj6":                  &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"test_options":               &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"transform":                  &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
	}

	ctx := t.Context()
	ctx = context.WithValue(ctx, utils.MetadataContextKey, md)

	for fnName, fn := range md.FnExports {
		t.Run(fnName, func(t *testing.T) {
			fnDef, ok := exports[fnName]
			if !ok {
				t.Fatalf("no wasm function definition found for '%v'", fnName)
			}

			plan, err := planner.GetPlan(ctx, fn, fnDef)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, len(fn.Parameters), len(plan.ParamHandlers()))
			assert.Equal(t, len(fn.Results), len(plan.ResultHandlers()))
		})
	}
}

type fakeWasmFunc struct {
	wasm.FunctionDefinition
	resultTypes []wasm.ValueType
}

func (f *fakeWasmFunc) DebugName() string              { return "" }
func (f *fakeWasmFunc) ExportNames() []string          { return nil }
func (f *fakeWasmFunc) GoFunction() any                { return nil }
func (f *fakeWasmFunc) Import() (string, string, bool) { return "", "", false }
func (f *fakeWasmFunc) Index() uint32                  { return 0 }
func (f *fakeWasmFunc) ModuleName() string             { return "" }
func (f *fakeWasmFunc) Name() string                   { return "" }
func (f *fakeWasmFunc) ParamNames() []string           { return nil }
func (f *fakeWasmFunc) ParamTypes() []wasm.ValueType   { return nil }
func (f *fakeWasmFunc) ResultNames() []string          { return nil }
func (f *fakeWasmFunc) ResultTypes() []wasm.ValueType  { return f.resultTypes }
