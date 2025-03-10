// -*- compile-command: "go run main.go ../../testdata/test-suite"; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"log"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/packages"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	for _, arg := range flag.Args() {
		log.Printf("Processing %v ...", arg)
		processDir(arg)
	}

	log.Printf("Done.")
}

func processDir(dir string) {
	mode := packages.NeedName | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo
	cfg := &packages.Config{Mode: mode, Dir: dir}
	got, err := packages.Load(cfg, ".")
	if err != nil || len(got) != 1 {
		log.Fatal(err)
	}

	dumpSyntax(got[0].Syntax)
}

func dumpSyntax(syntax []*ast.File) {
	fmt.Println("  Syntax: []*ast.File{")
	for _, file := range syntax {
		fmt.Printf("    {\n")
		fmt.Printf("      Name: &ast.Ident{Name: %q},\n", file.Name)
		if len(file.Decls) > 0 {
			fmt.Printf("      Decls: []ast.Decl{\n")
			for _, v := range file.Decls {
				printDecl(v)
			}
			fmt.Printf("      },\n") // Decls
		}
		fmt.Printf("      Imports: []*ast.ImportSpec{\n")
		for _, v := range file.Imports {
			fmt.Printf("        {Path: &ast.BasicLit{Value: `%v`}},\n", v.Path.Value)
		}
		fmt.Printf("      },\n") // Imports
		fmt.Printf("    },\n")   // File
	}
	fmt.Printf("  },\n") // Syntax
}

func printDecl(v ast.Decl) {
	switch t := v.(type) {
	case *ast.GenDecl:
		fmt.Printf("        &ast.GenDecl{\n")
		fmt.Printf("          Tok: token.TYPE,\n")
		fmt.Printf("          Specs: []ast.Spec{\n")
		for _, spec := range t.Specs {
			switch ts := spec.(type) {
			case *ast.TypeSpec:
				fmt.Printf("            &ast.TypeSpec{Name: &ast.Ident{Name: %q},\n", ts.Name.Name)
				s := ts.Type.(*ast.StructType)
				fmt.Printf("              Type: &ast.StructType{\n")
				fmt.Printf("                Fields: &ast.FieldList{\n")
				fmt.Printf("                  List: []*ast.Field{\n")
				for _, field := range s.Fields.List {
					fmt.Printf("                    {Names: []*ast.Ident{{Name: %q}}, Type: &ast.Ident{Name: %q}},\n", field.Names[0].Name, field.Type)
				}
				fmt.Printf("                  },\n")
				fmt.Printf("                },\n")
				fmt.Printf("              },\n")
				fmt.Printf("            },\n")
			default:
				log.Fatalf("unhandled ast.Spec type: %T", v)
			}
		}
		fmt.Printf("          },\n") // Specs
		fmt.Printf("        },\n")   // GenDecl
	case *ast.FuncDecl:
		fmt.Printf("        &ast.FuncDecl{\n")
		if t.Doc != nil {
			if len(t.Doc.List) > 0 {
				fmt.Printf("          Doc: &ast.CommentGroup{\n")
				fmt.Printf("            List: []*ast.Comment{\n")
				for _, comment := range t.Doc.List {
					fmt.Printf("              {Text: %q},\n", comment.Text)
				}
				fmt.Printf("            },\n") // List
				fmt.Printf("          },\n")   // Doc
			} else {
				fmt.Printf("          Doc: &ast.CommentGroup{},\n")
			}
		}
		fmt.Printf("          Name: &ast.Ident{Name: %q},\n", t.Name.Name)
		fmt.Printf("          Type: &ast.FuncType{\n")
		if len(t.Type.Params.List) > 0 {
			fmt.Printf("            Params: &ast.FieldList{\n")
			fmt.Printf("              List: []*ast.Field{\n")
			for _, field := range t.Type.Params.List {
				if field.Tag != nil {
					fmt.Printf("                {Names: []*ast.Ident{{Name: %q}}, Type: &ast.Ident{Name: %q}, Tag: &ast.BasicLit{Kind: token.STRING, Value: %q}},\n", field.Names[0].Name, field.Type.(*ast.Ident).Name, field.Tag.Value)
				} else {
					fmt.Printf("                {Names: []*ast.Ident{{Name: %q}}, Type: &ast.Ident{Name: %q}},\n", field.Names[0].Name, field.Type.(*ast.Ident).Name)
				}
			}
			fmt.Printf("              },\n") // List
			fmt.Printf("            },\n")   // Params
		} else {
			fmt.Printf("            Params: &ast.FieldList{List: []*ast.Field{}},\n")
		}
		if t.Type.Results != nil && len(t.Type.Results.List) > 0 {
			fmt.Printf("            Results: &ast.FieldList{\n")
			fmt.Printf("              List: []*ast.Field{\n")
			for _, field := range t.Type.Results.List {
				fmt.Printf("                {Type: &ast.Ident{Name: %q}},\n", field.Type.(*ast.Ident).Name)
			}
			fmt.Printf("              },\n") // List
			fmt.Printf("            },\n")   // Params
		}
		fmt.Printf("          },\n") // Type
		fmt.Printf("        },\n")   // FuncDecl
	default:
		log.Fatalf("unhandled ast.Decl type: %T", v)
	}
}
