/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package golang

import (
	"context"
	"log"
	"testing"

	"github.com/gmlewis/modus/lib/manifest"
	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/manifestdata"
	"github.com/gmlewis/modus/runtime/utils"

	"github.com/stretchr/testify/assert"
	wasm "github.com/tetratelabs/wazero/api"
)

func TestGetPlan(t *testing.T) {
	// Copied from runtime/graphql/schemagen/schemagen_go_test.go
	log.SetFlags(0)
	manifest := &manifest.Manifest{
		Models:      map[string]manifest.ModelInfo{},
		Connections: map[string]manifest.ConnectionInfo{},
		Collections: map[string]manifest.CollectionInfo{
			"collection1": {
				SearchMethods: map[string]manifest.SearchMethodInfo{
					"search1": {
						Embedder: "myEmbedder",
					},
				},
			},
		},
	}
	manifestdata.SetManifest(manifest)

	md := metadata.NewPluginMetadata()
	md.SDK = "modus-sdk-go"

	md.FnExports.AddFunction("add").
		WithParameter("a", "int32").
		WithParameter("b", "int32").
		WithResult("int32")

	md.FnExports.AddFunction("sayHello").
		WithParameter("name", "string").
		WithResult("string")

	md.FnExports.AddFunction("currentTime").
		WithResult("time.Time")

	md.FnExports.AddFunction("transform").
		WithParameter("items", "map[string]string").
		WithResult("map[string]string")

	md.FnExports.AddFunction("testDefaultIntParams").
		WithParameter("a", "int32").
		WithParameter("b", "int32", 0).
		WithParameter("c", "int32", 1)

	md.FnExports.AddFunction("testDefaultStringParams").
		WithParameter("a", "string").
		WithParameter("b", "string", "").
		WithParameter("c", "string", `a"b`).
		WithParameter("d", "*string").
		WithParameter("e", "*string", nil).
		WithParameter("f", "*string", "").
		WithParameter("g", "*string", "test")

	md.FnExports.AddFunction("testDefaultArrayParams").
		WithParameter("a", "[]int32").
		WithParameter("b", "[]int32", []int32{}).
		WithParameter("c", "[]int32", []int32{1, 2, 3}).
		WithParameter("d", "*[]int32").
		WithParameter("e", "*[]int32", nil).
		WithParameter("f", "*[]int32", []int32{}).
		WithParameter("g", "*[]int32", []int32{1, 2, 3})

	md.FnExports.AddFunction("getPerson").
		WithResult("testdata.Person")

	md.FnExports.AddFunction("listPeople").
		WithResult("[]testdata.Person")

	md.FnExports.AddFunction("addPerson").
		WithParameter("person", "testdata.Person")

	md.FnExports.AddFunction("getProductMap").
		WithResult("map[string]testdata.Product")

	md.FnExports.AddFunction("doNothing")

	md.FnExports.AddFunction("testPointers").
		WithParameter("a", "*int32").
		WithParameter("b", "[]*int32").
		WithParameter("c", "*[]int32").
		WithParameter("d", "[]*testdata.Person").
		WithParameter("e", "*[]testdata.Person").
		WithResult("*testdata.Person").
		WithDocs(metadata.Docs{
			Lines: []string{
				"This function tests that pointers are working correctly",
			},
		})

	md.Types.AddType("*int32")
	md.Types.AddType("[]*int32")
	md.Types.AddType("*[]int32")
	md.Types.AddType("[]*testdata.Person")
	md.Types.AddType("*[]testdata.Person")
	md.Types.AddType("*testdata.Person")
	md.Types.AddType("*string")
	md.Types.AddType("[]testdata.Address")
	md.Types.AddType("[]string")
	md.Types.AddType("[]testdata.Product")

	// This should be excluded from the final schema
	md.FnExports.AddFunction("myEmbedder").
		WithParameter("text", "string").
		WithResult("[]float64")

	// Generated input object from the output object
	md.FnExports.AddFunction("testObj1").
		WithParameter("obj", "testdata.Obj1").
		WithResult("testdata.Obj1")

	md.Types.AddType("testdata.Obj1").
		WithField("id", "int32").
		WithField("name", "string")

	// Separate input and output objects defined
	md.FnExports.AddFunction("testObj2").
		WithParameter("obj", "testdata.Obj2Input").
		WithResult("testdata.Obj2")
	md.Types.AddType("testdata.Obj2").
		WithField("id", "int32").
		WithField("name", "string")
	md.Types.AddType("testdata.Obj2Input").
		WithField("name", "string")

	// Generated input object without output object
	md.FnExports.AddFunction("testObj3").
		WithParameter("obj", "testdata.Obj3")
	md.Types.AddType("testdata.Obj3").
		WithField("name", "string")

	// Single input object defined without output object
	md.FnExports.AddFunction("testObj4").
		WithParameter("obj", "testdata.Obj4Input")
	md.Types.AddType("testdata.Obj4Input").
		WithField("name", "string")

	// Test slice of pointers to structs
	md.FnExports.AddFunction("testObj5").
		WithResult("[]*testdata.Obj5")
	md.Types.AddType("[]*testdata.Obj5")
	md.Types.AddType("*testdata.Obj5")
	md.Types.AddType("testdata.Obj5").
		WithField("name", "string")

	// Test slice of structs
	md.FnExports.AddFunction("testObj6").
		WithResult("[]testdata.Obj6")
	md.Types.AddType("[]testdata.Obj6")
	md.Types.AddType("testdata.Obj6").
		WithField("name", "string")

	md.Types.AddType("[]int32")
	md.Types.AddType("[]float64")
	md.Types.AddType("[]testdata.Person")
	md.Types.AddType("map[string]string")
	md.Types.AddType("map[string]testdata.Product")

	md.Types.AddType("testdata.Company").
		WithField("name", "string")

	md.Types.AddType("testdata.Product").
		WithField("name", "string").
		WithField("price", "float64").
		WithField("manufacturer", "testdata.Company").
		WithField("components", "[]testdata.Product")

	md.Types.AddType("testdata.Person").
		WithField("name", "string").
		WithField("age", "int32").
		WithField("addresses", "[]testdata.Address")

	md.Types.AddType("testdata.Address").
		WithField("street", "string", &metadata.Docs{
			Lines: []string{
				"Street that the user lives on",
			},
		}).
		WithField("city", "string").
		WithField("state", "string").
		WithField("country", "string", &metadata.Docs{
			Lines: []string{
				"Country that the user is from",
			},
		}).
		WithField("postalCode", "string").
		WithField("location", "testdata.Coordinates").
		WithDocs(metadata.Docs{
			Lines: []string{
				"Address represents a physical address.",
				"Each field corresponds to a specific part of the address.",
				"The location field stores geospatial coordinates.",
			},
		})

	md.Types.AddType("testdata.Coordinates").
		WithField("lat", "float64").
		WithField("lon", "float64")

	// This should be excluded from the final schema
	md.Types.AddType("testdata.Header").
		WithField("name", "string").
		WithField("values", "[]string")

	planner := NewPlanner(md)

	// importsMap := make(map[string]wasm.FunctionDefinition, len(imports))
	// for _, fnDef := range imports {
	// 	if modName, fnName, ok := fnDef.Import(); ok {
	// 		importName := modName + "." + fnName
	// 		importsMap[importName] = fnDef
	// 	}
	// }
	exports := map[string]wasm.FunctionDefinition{
		"addPerson":               &fakeWasmFunc{},
		"add":                     &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI32}},
		"add3":                    &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI32}},
		"currentTime":             &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"doNothing":               &fakeWasmFunc{},
		"getPerson":               &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"getProductMap":           &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"listPeople":              &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"myEmbedder":              &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"sayHello":                &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"testDefaultArrayParams":  &fakeWasmFunc{},
		"testDefaultIntParams":    &fakeWasmFunc{},
		"testDefaultStringParams": &fakeWasmFunc{},
		"testPointers":            &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"testObj1":                &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"testObj2":                &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"testObj3":                &fakeWasmFunc{},
		"testObj4":                &fakeWasmFunc{},
		"testObj5":                &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"testObj6":                &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"testOptions":             &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
		"transform":               &fakeWasmFunc{resultTypes: []wasm.ValueType{wasm.ValueTypeI64}},
	}

	ctx := context.Background()
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
