/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package schemagen

import (
	"context"
	"strings"
	"testing"

	"github.com/gmlewis/modus/lib/manifest"
	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/languages"
	"github.com/gmlewis/modus/runtime/manifestdata"
	"github.com/gmlewis/modus/runtime/utils"

	"github.com/stretchr/testify/require"
)

func Test_GetGraphQLSchema_MoonBit(t *testing.T) {

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

	// TODO:
	// md.FnExports.AddFunction("add").
	// 	WithParameter("a", "Int").
	// 	WithParameter("b", "Int").
	// 	WithResult("Int")

	// md.FnExports.AddFunction("add3").
	// 	WithParameter("a", "Int").
	// 	WithParameter("b", "Int").
	// 	WithParameter("c~", "Int").
	// 	WithResult("Int")

	// md.FnExports.AddFunction("say_hello").
	// 	WithParameter("name~", "String").
	// 	WithResult("String")

	// md.FnExports.AddFunction("current_time").
	// 	WithResult("time.Time")

	// md.FnExports.AddFunction("transform").
	// 	WithParameter("items", "map[string]string").
	// 	WithResult("map[string]string")

	// md.FnExports.AddFunction("test_default_int_params").
	// 	WithParameter("a", "Int").
	// 	WithParameter("b", "Int", 0).
	// 	WithParameter("c", "Int", 1)

	// md.FnExports.AddFunction("test_default_string_params").
	// 	WithParameter("a", "String").
	// 	WithParameter("b", "String", "").
	// 	WithParameter("c", "String", `a"b`).
	// 	WithParameter("d", "*string").
	// 	WithParameter("e", "*string", nil).
	// 	WithParameter("f", "*string", "").
	// 	WithParameter("g", "*string", "test")

	// md.FnExports.AddFunction("test_default_array_params").
	// 	WithParameter("a", "[]Int").
	// 	WithParameter("b", "[]Int", []int32{}).
	// 	WithParameter("c", "[]Int", []int32{1, 2, 3}).
	// 	WithParameter("d", "*[]Int").
	// 	WithParameter("e", "*[]Int", nil).
	// 	WithParameter("f", "*[]Int", []int32{}).
	// 	WithParameter("g", "*[]Int", []int32{1, 2, 3})

	// md.FnExports.AddFunction("get_person").
	// 	WithResult("testdata.Person")

	// md.FnExports.AddFunction("list_people").
	// 	WithResult("[]testdata.Person")

	// md.FnExports.AddFunction("add_person").
	// 	WithParameter("person", "testdata.Person")

	// md.FnExports.AddFunction("get_product_map").
	// 	WithResult("map[string]testdata.Product")

	// md.FnExports.AddFunction("do_nothing")

	// md.FnExports.AddFunction("test_pointers").
	// 	WithParameter("a", "*Int").
	// 	WithParameter("b", "[]*Int").
	// 	WithParameter("c", "*[]Int").
	// 	WithParameter("d", "[]*testdata.Person").
	// 	WithParameter("e", "*[]testdata.Person").
	// 	WithResult("*testdata.Person").
	// 	WithDocs(metadata.Docs{
	// 		Lines: []string{
	// 			"This function tests that pointers are working correctly",
	// 		},
	// 	})

	// md.Types.AddType("*Int")
	// md.Types.AddType("[]*Int")
	// md.Types.AddType("*[]Int")
	// md.Types.AddType("[]*testdata.Person")
	// md.Types.AddType("*[]testdata.Person")
	// md.Types.AddType("*testdata.Person")

	// // This should be excluded from the final schema
	// md.FnExports.AddFunction("my_embedder").
	// 	WithParameter("text", "String").
	// 	WithResult("[]float64")

	// // Generated input object from the output object
	// md.FnExports.AddFunction("test_obj1").
	// 	WithParameter("obj", "testdata.Obj1").
	// 	WithResult("testdata.Obj1")

	// md.Types.AddType("testdata.Obj1").
	// 	WithField("id", "Int").
	// 	WithField("name", "String")

	// // Separate input and output objects defined
	// md.FnExports.AddFunction("test_obj2").
	// 	WithParameter("obj", "testdata.Obj2Input").
	// 	WithResult("testdata.Obj2")
	// md.Types.AddType("testdata.Obj2").
	// 	WithField("id", "Int").
	// 	WithField("name", "String")
	// md.Types.AddType("testdata.Obj2Input").
	// 	WithField("name", "String")

	// // Generated input object without output object
	// md.FnExports.AddFunction("test_obj3").
	// 	WithParameter("obj", "testdata.Obj3")
	// md.Types.AddType("testdata.Obj3").
	// 	WithField("name", "String")

	// // Single input object defined without output object
	// md.FnExports.AddFunction("test_obj4").
	// 	WithParameter("obj", "testdata.Obj4Input")
	// md.Types.AddType("testdata.Obj4Input").
	// 	WithField("name", "String")

	// // Test slice of pointers to structs
	// md.FnExports.AddFunction("test_obj5").
	// 	WithResult("[]*testdata.Obj5")
	// md.Types.AddType("[]*testdata.Obj5")
	// md.Types.AddType("*testdata.Obj5")
	// md.Types.AddType("testdata.Obj5").
	// 	WithField("name", "String")

	// // Test slice of structs
	// md.FnExports.AddFunction("test_obj6").
	// 	WithResult("[]testdata.Obj6")
	// md.Types.AddType("[]testdata.Obj6")
	// md.Types.AddType("testdata.Obj6").
	// 	WithField("name", "String")

	// md.Types.AddType("[]Int")
	// md.Types.AddType("[]float64")
	// md.Types.AddType("[]testdata.Person")
	// md.Types.AddType("map[string]string")
	// md.Types.AddType("map[string]testdata.Product")

	// md.Types.AddType("testdata.Company").
	// 	WithField("name", "String")

	// md.Types.AddType("testdata.Product").
	// 	WithField("name", "String").
	// 	WithField("price", "float64").
	// 	WithField("manufacturer", "testdata.Company").
	// 	WithField("components", "[]testdata.Product")

	// md.Types.AddType("testdata.Person").
	// 	WithField("name", "String").
	// 	WithField("age", "Int").
	// 	WithField("addresses", "[]testdata.Address")

	// md.Types.AddType("testdata.Address").
	// 	WithField("street", "String", &metadata.Docs{
	// 		Lines: []string{
	// 			"Street that the user lives on",
	// 		},
	// 	}).
	// 	WithField("city", "String").
	// 	WithField("state", "String").
	// 	WithField("country", "String", &metadata.Docs{
	// 		Lines: []string{
	// 			"Country that the user is from",
	// 		},
	// 	}).
	// 	WithField("postalCode", "String").
	// 	WithField("location", "testdata.Coordinates").
	// 	WithDocs(metadata.Docs{
	// 		Lines: []string{
	// 			"Address represents a physical address.",
	// 			"Each field corresponds to a specific part of the address.",
	// 			"The location field stores geospatial coordinates.",
	// 		},
	// 	})

	// md.Types.AddType("testdata.Coordinates").
	// 	WithField("lat", "float64").
	// 	WithField("lon", "float64")

	// // This should be excluded from the final schema
	// md.Types.AddType("testdata.Header").
	// 	WithField("name", "String").
	// 	WithField("values", "[]string")

	result, err := GetGraphQLSchema(context.Background(), md)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result.Schema)

	expectedSchema := `
# Modus GraphQL Schema (auto-generated)
`[1:]

	/* TODO:
	   	expectedSchema := `
	   # Modus GraphQL Schema (auto-generated)

	   type Query {
	     current_time: Timestamp!
	     do_nothing: Void
	     people: [Person!]
	     person: Person!
	     product_map: [StringProductPair!]
	     say_hello(name: String!): String!
	     test_default_array_params(a: [Int!], b: [Int!] = [], c: [Int!] = [1,2,3], d: [Int!], e: [Int!] = null, f: [Int!] = [], g: [Int!] = [1,2,3]): Void
	     test_default_int_params(a: Int!, b: Int! = 0, c: Int! = 1): Void
	     test_default_string_params(a: String!, b: String! = "", c: String! = "a\"b", d: String, e: String = null, f: String = "", g: String = "test"): Void
	     test_obj1(obj: Obj1Input!): Obj1!
	     test_obj2(obj: Obj2Input!): Obj2!
	     test_obj3(obj: Obj3Input!): Void
	     test_obj4(obj: Obj4Input!): Void
	     test_obj5: [Obj5]
	     test_obj6: [Obj6!]
	     """
	     This function tests that pointers are working correctly
	     """
	     test_pointers(a: Int, b: [Int], c: [Int!], d: [PersonInput], e: [PersonInput!]): Person
	     transform(items: [StringStringPairInput!]): [StringStringPair!]
	   }

	   type Mutation {
	     add(a: Int!, b: Int!): Int!
	     add_person(person: PersonInput!): Void
	   }

	   scalar Timestamp
	   scalar Void

	   """
	   Address represents a physical address.
	   Each field corresponds to a specific part of the address.
	   The location field stores geospatial coordinates.
	   """
	   input AddressInput {
	     """
	     Street that the user lives on
	     """
	     street: String!
	     city: String!
	     state: String!
	     """
	     Country that the user is from
	     """
	     country: String!
	     postalCode: String!
	     location: CoordinatesInput!
	   }

	   input CoordinatesInput {
	     lat: Float!
	     lon: Float!
	   }

	   input Obj1Input {
	     id: Int!
	     name: String!
	   }

	   input Obj2Input {
	     name: String!
	   }

	   input Obj3Input {
	     name: String!
	   }

	   input Obj4Input {
	     name: String!
	   }

	   input PersonInput {
	     name: String!
	     age: Int!
	     addresses: [AddressInput!]
	   }

	   input StringStringPairInput {
	     key: String!
	     value: String!
	   }

	   """
	   Address represents a physical address.
	   Each field corresponds to a specific part of the address.
	   The location field stores geospatial coordinates.
	   """
	   type Address {
	     """
	     Street that the user lives on
	     """
	     street: String!
	     city: String!
	     state: String!
	     """
	     Country that the user is from
	     """
	     country: String!
	     postalCode: String!
	     location: Coordinates!
	   }

	   type Company {
	     name: String!
	   }

	   type Coordinates {
	     lat: Float!
	     lon: Float!
	   }

	   type Obj1 {
	     id: Int!
	     name: String!
	   }

	   type Obj2 {
	     id: Int!
	     name: String!
	   }

	   type Obj5 {
	     name: String!
	   }

	   type Obj6 {
	     name: String!
	   }

	   type Person {
	     name: String!
	     age: Int!
	     addresses: [Address!]
	   }

	   type Product {
	     name: String!
	     price: Float!
	     manufacturer: Company!
	     components: [Product!]
	   }

	   type StringProductPair {
	     key: String!
	     value: Product!
	   }

	   type StringStringPair {
	     key: String!
	     value: String!
	   }
	   `[1:]
	*/

	require.Nil(t, err)
	require.Equal(t, expectedSchema, result.Schema)
}

func Test_ConvertType_MoonBit(t *testing.T) {

	lti := languages.MoonBit().TypeInfo()

	testCases := []struct {
		sourceType          string
		forInput            bool
		expectedGraphQLType string
		sourceTypeDefs      []*metadata.TypeDefinition
		expectedTypeDefs    []*TypeDefinition
	}{
		// Plain non-nullable types
		{"String", false, "String!", nil, nil},
		{"String", true, "String!", nil, nil},
		{"Bool", false, "Boolean!", nil, nil},
		{"Bool", true, "Boolean!", nil, nil},
		// {"Int8", false, "Int!", nil, nil},
		// {"Int8", true, "Int!", nil, nil},
		// {"Int16", false, "Int!", nil, nil},
		// {"Int16", true, "Int!", nil, nil},
		// {"Int", false, "Int!", nil, nil},
		// {"Int", true, "Int!", nil, nil},
		// {"Byte", false, "Int!", nil, nil},
		// {"Byte", true, "Int!", nil, nil},
		// {"Uint16", false, "Int!", nil, nil},
		// {"Uint16", true, "Int!", nil, nil},
		{"Float", false, "Float!", nil, nil},
		{"Float", true, "Float!", nil, nil},
		{"Double", false, "Float!", nil, nil},
		{"Double", true, "Float!", nil, nil},

		// TODO:
		// // Slice types
		// {"[]string", false, "[String!]", nil, nil},
		// {"[]string", true, "[String!]", nil, nil},
		// {"[][]string", false, "[[String!]]", nil, nil},
		// {"[][]string", true, "[[String!]]", nil, nil},
		// {"[]*string", false, "[String]", nil, nil},
		// {"[]*string", true, "[String]", nil, nil},

		// // Custom scalar types
		// {"time.Time", false, "Timestamp!", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		// {"time.Time", true, "Timestamp!", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		// {"int64", false, "Int64!", nil, []*TypeDefinition{{Name: "Int64"}}},
		// {"int64", true, "Int64!", nil, []*TypeDefinition{{Name: "Int64"}}},
		// {"uInt", false, "UInt!", nil, []*TypeDefinition{{Name: "UInt"}}},
		// {"uInt", true, "UInt!", nil, []*TypeDefinition{{Name: "UInt"}}},
		// {"uint64", false, "UInt64!", nil, []*TypeDefinition{{Name: "UInt64"}}},
		// {"uint64", true, "UInt64!", nil, []*TypeDefinition{{Name: "UInt64"}}},

		// // Custom types
		// {"testdata.User", false, "User!",
		// 	[]*metadata.TypeDefinition{{
		// 		Name: "User",
		// 		Fields: []*metadata.Field{
		// 			{Name: "firstName", Type: "String"},
		// 			{Name: "lastName", Type: "String"},
		// 			{Name: "age", Type: "uint8"},
		// 		},
		// 	}},
		// 	[]*TypeDefinition{{
		// 		Name: "User",
		// 		Fields: []*FieldDefinition{
		// 			{Name: "firstName", Type: "String!"},
		// 			{Name: "lastName", Type: "String!"},
		// 			{Name: "age", Type: "Int!"},
		// 		},
		// 	}}},
		// {"testdata.User", true, "UserInput!",
		// 	[]*metadata.TypeDefinition{{
		// 		Name: "User",
		// 		Fields: []*metadata.Field{
		// 			{Name: "firstName", Type: "String"},
		// 			{Name: "lastName", Type: "String"},
		// 			{Name: "age", Type: "uint8"},
		// 		},
		// 	}},
		// 	[]*TypeDefinition{{
		// 		Name: "UserInput",
		// 		Fields: []*FieldDefinition{
		// 			{Name: "firstName", Type: "String!"},
		// 			{Name: "lastName", Type: "String!"},
		// 			{Name: "age", Type: "Int!"},
		// 		},
		// 	}}},

		// {"*bool", false, "Boolean", nil, nil},
		// {"*bool", true, "Boolean", nil, nil},
		// {"*int", false, "Int", nil, nil},
		// {"*int", true, "Int", nil, nil},
		// {"*float64", false, "Float", nil, nil},
		// {"*float64", true, "Float", nil, nil},
		// {"*string", false, "String", nil, nil},
		// {"*string", true, "String", nil, nil},
		// {"*testdata.Foo", false, "Foo", // scalar
		// 	[]*metadata.TypeDefinition{{Name: "testdata.Foo"}},
		// 	[]*TypeDefinition{{Name: "Foo"}}},
		// {"*testdata.Foo", true, "Foo", // scalar
		// 	[]*metadata.TypeDefinition{{Name: "testdata.Foo"}},
		// 	[]*TypeDefinition{{Name: "Foo"}}},

		// // Map types
		// {"map[string]string", false, "[StringStringPair!]", nil, []*TypeDefinition{{
		// 	Name: "StringStringPair",
		// 	Fields: []*FieldDefinition{
		// 		{Name: "key", Type: "String!"},
		// 		{Name: "value", Type: "String!"},
		// 	},
		// 	IsMapType: true,
		// }}},
		// {"map[string]string", true, "[StringStringPairInput!]", nil, []*TypeDefinition{{
		// 	Name: "StringStringPairInput",
		// 	Fields: []*FieldDefinition{
		// 		{Name: "key", Type: "String!"},
		// 		{Name: "value", Type: "String!"},
		// 	},
		// 	IsMapType: true,
		// }}},
		// {"map[string]*string", false, "[StringNullableStringPair!]", nil, []*TypeDefinition{{
		// 	Name: "StringNullableStringPair",
		// 	Fields: []*FieldDefinition{
		// 		{Name: "key", Type: "String!"},
		// 		{Name: "value", Type: "String"},
		// 	},
		// 	IsMapType: true,
		// }}},
		// {"map[string]*string", true, "[StringNullableStringPairInput!]", nil, []*TypeDefinition{{
		// 	Name: "StringNullableStringPairInput",
		// 	Fields: []*FieldDefinition{
		// 		{Name: "key", Type: "String!"},
		// 		{Name: "value", Type: "String"},
		// 	},
		// 	IsMapType: true,
		// }}},
		// {"map[Int]string", false, "[IntStringPair!]", nil, []*TypeDefinition{{
		// 	Name: "IntStringPair",
		// 	Fields: []*FieldDefinition{
		// 		{Name: "key", Type: "Int!"},
		// 		{Name: "value", Type: "String!"},
		// 	},
		// 	IsMapType: true,
		// }}},
		// {"map[Int]string", true, "[IntStringPairInput!]", nil, []*TypeDefinition{{
		// 	Name: "IntStringPairInput",
		// 	Fields: []*FieldDefinition{
		// 		{Name: "key", Type: "Int!"},
		// 		{Name: "value", Type: "String!"},
		// 	},
		// 	IsMapType: true,
		// }}},
		// {"map[string]map[string]float32", false, "[StringNullableStringFloatPairListPair!]", nil, []*TypeDefinition{
		// 	{
		// 		Name: "StringNullableStringFloatPairListPair",
		// 		Fields: []*FieldDefinition{
		// 			{Name: "key", Type: "String!"},
		// 			{Name: "value", Type: "[StringFloatPair!]"},
		// 		},
		// 		IsMapType: true,
		// 	},
		// 	{
		// 		Name: "StringFloatPair",
		// 		Fields: []*FieldDefinition{
		// 			{Name: "key", Type: "String!"},
		// 			{Name: "value", Type: "Float!"},
		// 		},
		// 		IsMapType: true,
		// 	},
		// }},
		// {"map[string]map[string]float32", true, "[StringNullableStringFloatPairListPairInput!]", nil, []*TypeDefinition{
		// 	{
		// 		Name: "StringNullableStringFloatPairListPairInput",
		// 		Fields: []*FieldDefinition{
		// 			{Name: "key", Type: "String!"},
		// 			{Name: "value", Type: "[StringFloatPairInput!]"},
		// 		},
		// 		IsMapType: true,
		// 	},
		// 	{
		// 		Name: "StringFloatPairInput",
		// 		Fields: []*FieldDefinition{
		// 			{Name: "key", Type: "String!"},
		// 			{Name: "value", Type: "Float!"},
		// 		},
		// 		IsMapType: true,
		// 	},
		// }},
	}

	for _, tc := range testCases {
		testName := strings.ReplaceAll(tc.sourceType, " ", "")
		if tc.forInput {
			testName += "_input"
		}
		t.Run(testName, func(t *testing.T) {

			types := make(metadata.TypeMap, len(tc.sourceTypeDefs))
			for _, td := range tc.sourceTypeDefs {
				types[td.Name] = td
			}

			typeDefs, errors := transformTypes(types, lti, tc.forInput)
			require.Empty(t, errors)

			result, err := convertType(tc.sourceType, lti, typeDefs, false, tc.forInput)

			require.Nil(t, err)
			require.Equal(t, tc.expectedGraphQLType, result)

			if tc.expectedTypeDefs == nil {
				require.Empty(t, typeDefs)
			} else {
				require.ElementsMatch(t, tc.expectedTypeDefs, utils.MapValues(typeDefs))
			}
		})
	}
}
