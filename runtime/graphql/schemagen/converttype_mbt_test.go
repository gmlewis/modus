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
	"strings"
	"testing"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/languages"
	"github.com/gmlewis/modus/runtime/utils"

	"github.com/stretchr/testify/require"
)

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
		{"Bool", false, "Boolean!", nil, nil},
		{"Bool", true, "Boolean!", nil, nil},
		{"Byte", false, "Int!", nil, nil},
		{"Byte", true, "Int!", nil, nil},
		{"Char", false, "Int!", nil, nil},
		{"Char", true, "Int!", nil, nil},
		{"Double", false, "Float!", nil, nil},
		{"Double", false, "Float!", nil, nil},
		{"Double", true, "Float!", nil, nil},
		{"Double", true, "Float!", nil, nil},
		{"Float", false, "Float!", nil, nil},
		{"Float", false, "Float!", nil, nil},
		{"Float", true, "Float!", nil, nil},
		{"Float", true, "Float!", nil, nil},
		{"Int", false, "Int!", nil, nil},
		{"Int", true, "Int!", nil, nil},
		{"Int16", false, "Int!", nil, nil},
		{"Int16", true, "Int!", nil, nil},
		{"String", false, "String!", nil, nil},
		{"String", true, "String!", nil, nil},
		{"UInt16", false, "Int!", nil, nil},
		{"UInt16", true, "Int!", nil, nil},

		// Option[Plain] (nullable) types
		{"Bool?", false, "Boolean", nil, nil},
		{"Bool?", true, "Boolean", nil, nil},
		{"Byte?", false, "Int", nil, nil},
		{"Byte?", true, "Int", nil, nil},
		{"Char?", false, "Int", nil, nil},
		{"Char?", true, "Int", nil, nil},
		{"Double?", false, "Float", nil, nil},
		{"Double?", false, "Float", nil, nil},
		{"Double?", true, "Float", nil, nil},
		{"Double?", true, "Float", nil, nil},
		{"Float?", false, "Float", nil, nil},
		{"Float?", false, "Float", nil, nil},
		{"Float?", true, "Float", nil, nil},
		{"Float?", true, "Float", nil, nil},
		{"Int?", false, "Int", nil, nil},
		{"Int?", true, "Int", nil, nil},
		{"Int16?", false, "Int", nil, nil},
		{"Int16?", true, "Int", nil, nil},
		{"String?", false, "String", nil, nil},
		{"String?", true, "String", nil, nil},
		{"UInt16?", false, "Int", nil, nil},
		{"UInt16?", true, "Int", nil, nil},

		// Option[Array[...]] ("slice") types are nullable:
		// can return either `[]` (with `Some(...)`) or `null` (with `None`).
		{"Array[Int]?", false, "[Int!]", nil, nil},
		{"Array[Int]?", true, "[Int!]", nil, nil},
		{"Array[Array[Int]?]?", false, "[[Int!]]", nil, nil},
		{"Array[Array[Int]?]?", true, "[[Int!]]", nil, nil},
		{"Array[String]?", false, "[String!]", nil, nil},
		{"Array[String]?", true, "[String!]", nil, nil},
		{"Array[Array[String]?]?", false, "[[String!]]", nil, nil},
		{"Array[Array[String]?]?", true, "[[String!]]", nil, nil},
		// Option[Array[Option[...]]] types with nullable element types:
		{"Array[Int?]?", false, "[Int]", nil, nil},
		{"Array[Int?]?", true, "[Int]", nil, nil},
		{"Array[Array[Int?]?]?", false, "[[Int]]", nil, nil},
		{"Array[Array[Int?]?]?", true, "[[Int]]", nil, nil},
		{"Array[String?]?", false, "[String]", nil, nil},
		{"Array[String?]?", true, "[String]", nil, nil},
		{"Array[Array[String?]?]?", false, "[[String]]", nil, nil},
		{"Array[Array[String?]?]?", true, "[[String]]", nil, nil},

		// Array[...] types (non-nullable: can return `[]` but not `null`)
		{"Array[Int]", false, "[Int!]!", nil, nil},
		{"Array[Int]", true, "[Int!]!", nil, nil},
		{"Array[Array[Int]]", false, "[[Int!]!]!", nil, nil},
		{"Array[Array[Int]]", true, "[[Int!]!]!", nil, nil},
		{"Array[String]", false, "[String!]!", nil, nil},
		{"Array[String]", true, "[String!]!", nil, nil},
		{"Array[Array[String]]", false, "[[String!]!]!", nil, nil},
		{"Array[Array[String]]", true, "[[String!]!]!", nil, nil},
		// Array[Option[...]] types with nullable element types:
		{"Array[Int?]", false, "[Int]!", nil, nil},
		{"Array[Int?]", true, "[Int]!", nil, nil},
		{"Array[Array[Int?]]", false, "[[Int]!]!", nil, nil},
		{"Array[Array[Int?]]", true, "[[Int]!]!", nil, nil},
		{"Array[String?]", false, "[String]!", nil, nil},
		{"Array[String?]", true, "[String]!", nil, nil},
		{"Array[Array[String?]]", false, "[[String]!]!", nil, nil},
		{"Array[Array[String?]]", true, "[[String]!]!", nil, nil},

		// Custom scalar types (non-nullable)
		{"Int64", false, "Int64!", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"Int64", true, "Int64!", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"UInt", false, "UInt!", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"UInt", true, "UInt!", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"UInt64", false, "UInt64!", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"UInt64", true, "UInt64!", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"@time.ZonedDateTime", false, "Timestamp!", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		{"@time.ZonedDateTime", true, "Timestamp!", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		// Option[...] variations:
		{"Int64?", false, "Int64", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"Int64?", true, "Int64", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"UInt?", false, "UInt", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"UInt?", true, "UInt", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"UInt64?", false, "UInt64", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"UInt64?", true, "UInt64", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"@time.ZonedDateTime?", false, "Timestamp", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		{"@time.ZonedDateTime?", true, "Timestamp", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		// Option[Array[...]] variations:
		{"Array[Int64]?", false, "[Int64!]", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"Array[Int64]?", true, "[Int64!]", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"Array[UInt]?", false, "[UInt!]", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"Array[UInt]?", true, "[UInt!]", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"Array[UInt64]?", false, "[UInt64!]", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"Array[UInt64]?", true, "[UInt64!]", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"Array[@time.ZonedDateTime]?", false, "[Timestamp!]", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		{"Array[@time.ZonedDateTime]?", true, "[Timestamp!]", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		// Option[Array[Option[...]]] variations:
		{"Array[Int64?]?", false, "[Int64]", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"Array[Int64?]?", true, "[Int64]", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"Array[UInt?]?", false, "[UInt]", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"Array[UInt?]?", true, "[UInt]", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"Array[UInt64?]?", false, "[UInt64]", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"Array[UInt64?]?", true, "[UInt64]", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"Array[@time.ZonedDateTime?]?", false, "[Timestamp]", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		{"Array[@time.ZonedDateTime?]?", true, "[Timestamp]", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		// Array[...] variations:
		{"Array[Int64]", false, "[Int64!]!", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"Array[Int64]", true, "[Int64!]!", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"Array[UInt]", false, "[UInt!]!", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"Array[UInt]", true, "[UInt!]!", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"Array[UInt64]", false, "[UInt64!]!", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"Array[UInt64]", true, "[UInt64!]!", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"Array[@time.ZonedDateTime]", false, "[Timestamp!]!", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		{"Array[@time.ZonedDateTime]", true, "[Timestamp!]!", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		// Array[Option[...]] variations:
		{"Array[Int64?]", false, "[Int64]!", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"Array[Int64?]", true, "[Int64]!", nil, []*TypeDefinition{{Name: "Int64"}}},
		{"Array[UInt?]", false, "[UInt]!", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"Array[UInt?]", true, "[UInt]!", nil, []*TypeDefinition{{Name: "UInt"}}},
		{"Array[UInt64?]", false, "[UInt64]!", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"Array[UInt64?]", true, "[UInt64]!", nil, []*TypeDefinition{{Name: "UInt64"}}},
		{"Array[@time.ZonedDateTime?]", false, "[Timestamp]!", nil, []*TypeDefinition{{Name: "Timestamp"}}},
		{"Array[@time.ZonedDateTime?]", true, "[Timestamp]!", nil, []*TypeDefinition{{Name: "Timestamp"}}},

		// Custom types
		{"@testdata.User", false, "User!",
			[]*metadata.TypeDefinition{{
				Name: "User",
				Fields: []*metadata.Field{
					{Name: "firstName", Type: "String"},
					{Name: "lastName", Type: "String"},
					{Name: "age", Type: "Byte"},
				},
			}},
			[]*TypeDefinition{{
				Name: "User",
				Fields: []*FieldDefinition{
					{Name: "firstName", Type: "String!"},
					{Name: "lastName", Type: "String!"},
					{Name: "age", Type: "Int!"},
				},
			}}},
		{"testdata.User", true, "UserInput!",
			[]*metadata.TypeDefinition{{
				Name: "User",
				Fields: []*metadata.Field{
					{Name: "firstName", Type: "String"},
					{Name: "lastName", Type: "String"},
					{Name: "age", Type: "Byte"},
				},
			}},
			[]*TypeDefinition{{
				Name: "UserInput",
				Fields: []*FieldDefinition{
					{Name: "firstName", Type: "String!"},
					{Name: "lastName", Type: "String!"},
					{Name: "age", Type: "Int!"},
				},
			}}},

		{"@testdata.Foo?", false, "Foo", // Option[Foo]
			[]*metadata.TypeDefinition{{Name: "testdata.Foo"}},
			[]*TypeDefinition{{Name: "Foo"}}},
		{"@testdata.Foo?", true, "Foo", // Option[Foo]
			[]*metadata.TypeDefinition{{Name: "testdata.Foo"}},
			[]*TypeDefinition{{Name: "Foo"}}},

		// Map types
		{"Map[String,String]", false, "[StringStringPair!]!", nil, []*TypeDefinition{{
			Name: "StringStringPair",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "String!"},
				{Name: "value", Type: "String!"},
			},
			IsMapType: true,
		}}},
		{"Map[String,String]", true, "[StringStringPairInput!]!", nil, []*TypeDefinition{{
			Name: "StringStringPairInput",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "String!"},
				{Name: "value", Type: "String!"},
			},
			IsMapType: true,
		}}},
		{"Map[String,String]?", false, "[StringStringPair!]", nil, []*TypeDefinition{{
			Name: "StringStringPair",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "String!"},
				{Name: "value", Type: "String!"},
			},
			IsMapType: true,
		}}},
		{"Map[String,String]?", true, "[StringStringPairInput!]", nil, []*TypeDefinition{{
			Name: "StringStringPairInput",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "String!"},
				{Name: "value", Type: "String!"},
			},
			IsMapType: true,
		}}},
		{"Map[String,String?]?", false, "[StringNullableStringPair!]", nil, []*TypeDefinition{{
			Name: "StringNullableStringPair",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "String!"},
				{Name: "value", Type: "String"},
			},
			IsMapType: true,
		}}},
		{"Map[String,String?]?", true, "[StringNullableStringPairInput!]", nil, []*TypeDefinition{{
			Name: "StringNullableStringPairInput",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "String!"},
				{Name: "value", Type: "String"},
			},
			IsMapType: true,
		}}},
		{"Map[Int,String]", false, "[IntStringPair!]!", nil, []*TypeDefinition{{
			Name: "IntStringPair",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "Int!"},
				{Name: "value", Type: "String!"},
			},
			IsMapType: true,
		}}},
		{"Map[Int,String]", true, "[IntStringPairInput!]!", nil, []*TypeDefinition{{
			Name: "IntStringPairInput",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "Int!"},
				{Name: "value", Type: "String!"},
			},
			IsMapType: true,
		}}},
		{"Map[Int,String]?", false, "[IntStringPair!]", nil, []*TypeDefinition{{
			Name: "IntStringPair",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "Int!"},
				{Name: "value", Type: "String!"},
			},
			IsMapType: true,
		}}},
		{"Map[Int,String]?", true, "[IntStringPairInput!]", nil, []*TypeDefinition{{
			Name: "IntStringPairInput",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "Int!"},
				{Name: "value", Type: "String!"},
			},
			IsMapType: true,
		}}},
		{"Map[Int,String?]?", false, "[IntNullableStringPair!]", nil, []*TypeDefinition{{
			Name: "IntNullableStringPair",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "Int!"},
				{Name: "value", Type: "String"},
			},
			IsMapType: true,
		}}},
		{"Map[Int,String?]?", true, "[IntNullableStringPairInput!]", nil, []*TypeDefinition{{
			Name: "IntNullableStringPairInput",
			Fields: []*FieldDefinition{
				{Name: "key", Type: "Int!"},
				{Name: "value", Type: "String"},
			},
			IsMapType: true,
		}}},
		{"Map[String,Map[String,Float]?]?", false, "[StringNullableStringFloatPairListPair!]", nil, []*TypeDefinition{
			{
				Name: "StringNullableStringFloatPairListPair",
				Fields: []*FieldDefinition{
					{Name: "key", Type: "String!"},
					{Name: "value", Type: "[StringFloatPair!]"},
				},
				IsMapType: true,
			},
			{
				Name: "StringFloatPair",
				Fields: []*FieldDefinition{
					{Name: "key", Type: "String!"},
					{Name: "value", Type: "Float!"},
				},
				IsMapType: true,
			},
		}},
		{"Map[String,Map[String,Float]?]?", true, "[StringNullableStringFloatPairListPairInput!]", nil, []*TypeDefinition{
			{
				Name: "StringNullableStringFloatPairListPairInput",
				Fields: []*FieldDefinition{
					{Name: "key", Type: "String!"},
					{Name: "value", Type: "[StringFloatPairInput!]"},
				},
				IsMapType: true,
			},
			{
				Name: "StringFloatPairInput",
				Fields: []*FieldDefinition{
					{Name: "key", Type: "String!"},
					{Name: "value", Type: "Float!"},
				},
				IsMapType: true,
			},
		}},
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
