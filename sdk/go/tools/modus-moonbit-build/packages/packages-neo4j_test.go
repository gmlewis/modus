// -*- compile-command: "go test -run ^TestPackage_Neo4j$ ."; -*-

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

func TestPackage_Neo4j(t *testing.T) {
	t.Parallel()
	dir := "../testdata/neo4j-example"
	testPackageLoadHelper(t, "neo4j", dir, wantPackageNeo4j)
}

var wantPackageNeo4j = &Package{
	MoonPkgJSON: MoonPkgJSON{
		IsMain: false,
		Imports: []json.RawMessage{
			json.RawMessage(`"gmlewis/modus/pkg/neo4j"`),
			json.RawMessage(`"gmlewis/modus/pkg/console"`),
		},
		Targets: map[string][]string{
			"modus_post_generated.mbt": {"wasm"},
		},
		LinkTargets: map[string]*LinkTarget{
			"wasm": {
				Exports: []string{
					"__modus_create_people_and_relationships:create_people_and_relationships",
					"__modus_delete_all_nodes:delete_all_nodes",
					"__modus_get_alice_friends_under_40:get_alice_friends_under_40",
					"__modus_get_alice_friends_under_40_ages:get_alice_friends_under_40_ages",
					"cabi_realloc",
					"copy",
					"free",
					"load32",
					"malloc",
					"ptr2str",
					"ptr_to_none",
					"read_map",
					"store32",
					"store8",
					"write_map",
				},
				ExportMemoryName: "memory",
			},
		},
	},
	ID:           "moonbit-main",
	Name:         "main",
	PkgPath:      "",
	MoonBitFiles: []string{"../testdata/neo4j-example/main.mbt"},
	StructLookup: map[string]*ast.TypeSpec{
		"Person": {Name: &ast.Ident{Name: "Person"}, Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{{Name: "name"}}, Type: &ast.Ident{Name: "String"}},
					{Names: []*ast.Ident{{Name: "age"}}, Type: &ast.Ident{Name: "Int"}},
				},
			},
		}},
	},
	PossiblyMissingUnderlyingTypes: map[string]struct{}{
		"Int":    {},
		"String": {},
	},
	Syntax: []*ast.File{
		{
			Name: &ast.Ident{Name: "../testdata/neo4j-example/main.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "Person"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "name"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "age"}}, Type: &ast.Ident{Name: "Int"}},
									},
								},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "create_people_and_relationships"},
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
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_alice_friends_under_40"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Person]!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_alice_friends_under_40_ages"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "delete_all_nodes"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/neo4j"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
			},
		},
	},
	TypesInfo: &types.Info{
		Defs: map[*ast.Ident]types.Object{
			// Using &moonType{} and &moonFunc{} is a hack to fake a struct/func for testing purposes only:
			{Name: "Person"}: types.NewTypeName(0, nil, "Person", &moonType{typeName: "struct{name String; age Int}"}),
			{Name: "create_people_and_relationships"}: &moonFunc{funcName: "func create_people_and_relationships() String!Error"},
			{Name: "get_alice_friends_under_40"}:      &moonFunc{funcName: "func get_alice_friends_under_40() Array[Person]!Error"},
			{Name: "get_alice_friends_under_40_ages"}: &moonFunc{funcName: "func get_alice_friends_under_40_ages() Array[Int]!Error"},
			{Name: "delete_all_nodes"}:                &moonFunc{funcName: "func delete_all_nodes() String!Error"},
		},
	},
}
