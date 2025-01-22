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
	"log"
	"testing"

	"github.com/gmlewis/modus/lib/manifest"
	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/manifestdata"

	"github.com/stretchr/testify/require"
)

func Test_GetGraphQLSchema_MoonBit(t *testing.T) {
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
		WithResult("@time.PlainDateTime")

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
		WithResult("@testdata.Person")

	md.FnExports.AddFunction("list_people").
		WithResult("Array[@testdata.Person]")

	md.FnExports.AddFunction("add_person").
		WithParameter("person", "@testdata.Person")

	md.FnExports.AddFunction("get_product_map").
		WithResult("Map[String,@testdata.Product]")

	md.FnExports.AddFunction("do_nothing")

	md.FnExports.AddFunction("test_options").
		WithParameter("a", "Int?").
		WithParameter("b", "Array[Int?]").
		WithParameter("c", "Array[Int]?").
		WithParameter("d", "Array[@testdata.Person?]").
		WithParameter("e", "Array[@testdata.Person]?").
		WithResult("@testdata.Person?").
		WithDocs(metadata.Docs{
			Lines: []string{
				"This function tests that options are working correctly",
			},
		})

	md.Types.AddType("Int?")
	md.Types.AddType("Array[Int?]")
	md.Types.AddType("Array[Int]?")
	md.Types.AddType("Array[@testdata.Person?]")
	md.Types.AddType("Array[@testdata.Person]?")
	md.Types.AddType("@testdata.Person?")

	// This should be excluded from the final schema
	md.FnExports.AddFunction("my_embedder").
		WithParameter("text", "String").
		WithResult("Array[Double]")

	// Generated input object from the output object
	md.FnExports.AddFunction("test_obj1").
		WithParameter("obj", "@testdata.Obj1").
		WithResult("@testdata.Obj1")

	md.Types.AddType("@testdata.Obj1").
		WithField("id", "Int").
		WithField("name", "String")

	// Separate input and output objects defined
	md.FnExports.AddFunction("test_obj2").
		WithParameter("obj", "@testdata.Obj2Input").
		WithResult("@testdata.Obj2")
	md.Types.AddType("@testdata.Obj2").
		WithField("id", "Int").
		WithField("name", "String")
	md.Types.AddType("@testdata.Obj2Input").
		WithField("name", "String")

	// Generated input object without output object
	md.FnExports.AddFunction("test_obj3").
		WithParameter("obj", "@testdata.Obj3")
	md.Types.AddType("@testdata.Obj3").
		WithField("name", "String")

	// Single input object defined without output object
	md.FnExports.AddFunction("test_obj4").
		WithParameter("obj", "@testdata.Obj4Input")
	md.Types.AddType("@testdata.Obj4Input").
		WithField("name", "String")

	// Test slice of pointers to structs
	md.FnExports.AddFunction("test_obj5").
		WithResult("Array[@testdata.Obj5?]")
	md.Types.AddType("Array[@testdata.Obj5?]")
	md.Types.AddType("@testdata.Obj5?")
	md.Types.AddType("@testdata.Obj5").
		WithField("name", "String")

	// Test slice of structs
	md.FnExports.AddFunction("test_obj6").
		WithResult("Array[@testdata.Obj6]")
	md.Types.AddType("Array[@testdata.Obj6]")
	md.Types.AddType("@testdata.Obj6").
		WithField("name", "String")

	md.Types.AddType("Array[Int]")
	md.Types.AddType("Array[Double]")
	md.Types.AddType("Array[@testdata.Person]")
	md.Types.AddType("Map[String,String]")
	md.Types.AddType("Map[String,@testdata.Product]")

	md.Types.AddType("@testdata.Company").
		WithField("name", "String")

	md.Types.AddType("@testdata.Product").
		WithField("name", "String").
		WithField("price", "Double").
		WithField("manufacturer", "@testdata.Company").
		WithField("components", "Array[@testdata.Product]")

	md.Types.AddType("@testdata.Person").
		WithField("name", "String").
		WithField("age", "Int").
		WithField("addresses", "Array[@testdata.Address]")

	md.Types.AddType("@testdata.Address").
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
		WithField("location", "@testdata.Coordinates").
		WithDocs(metadata.Docs{
			Lines: []string{
				"Address represents a physical address.",
				"Each field corresponds to a specific part of the address.",
				"The location field stores geospatial coordinates.",
			},
		})

	md.Types.AddType("@testdata.Coordinates").
		WithField("lat", "Double").
		WithField("lon", "Double")

	// This should be excluded from the final schema
	md.Types.AddType("@testdata.Header").
		WithField("name", "String").
		WithField("values", "Array[String]")

	result, err := GetGraphQLSchema(context.Background(), md)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result.Schema)

	expectedSchema := `
# Modus GraphQL Schema (auto-generated)

type Query {
  current_time: Timestamp!
  do_nothing: Void
  people: [Person!]!
  person: Person!
  product_map: [StringProductPair!]!
  say_hello(name~: String! = "World"): String!
  test_default_array_params(a: [Int!]!, b~: [Int!]! = [], c~: [Int!]! = [1,2,3], d~: [Int!], e~: [Int!] = null, f~: [Int!] = [], g~: [Int!] = [1,2,3]): Void
  test_default_int_params(a: Int!, b~: Int! = 0, c~: Int! = 1): Void
  test_default_string_params(a: String!, b~: String! = "", c~: String! = "a\"b", d~: String, e~: String = null, f~: String = "", g~: String = "test"): Void
  test_obj1(obj: Obj1Input!): Obj1!
  test_obj2(obj: Obj2Input!): Obj2!
  test_obj3(obj: Obj3Input!): Void
  test_obj4(obj: Obj4Input!): Void
  test_obj5: [Obj5]!
  test_obj6: [Obj6!]!
  """
  This function tests that options are working correctly
  """
  test_options(a: Int, b: [Int]!, c: [Int!], d: [PersonInput]!, e: [PersonInput!]): Person
  transform(items: [StringStringPairInput!]!): [StringStringPair!]!
}

type Mutation {
  add(a: Int!, b: Int!): Int!
  add3(a: Int!, b: Int!, c~: Int! = 0): Int!
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
  addresses: [AddressInput!]!
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
  addresses: [Address!]!
}

type Product {
  name: String!
  price: Float!
  manufacturer: Company!
  components: [Product!]!
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

	require.Nil(t, err)
	require.Equal(t, expectedSchema, result.Schema)
}
