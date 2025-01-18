/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package codegen

import (
	"bytes"
	"go/ast"
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/packages"
	"github.com/google/go-cmp/cmp"
)

func TestWriteFuncWrappers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		pkg       *packages.Package
		imports   map[string]string
		functions []*funcInfo
		expected  string
	}{
		{
			name: "single function without default params",
			pkg: &packages.Package{
				Name: "main",
				Syntax: []*ast.File{
					{
						Decls: []ast.Decl{
							&ast.FuncDecl{
								Name: &ast.Ident{Name: "test_func"},
								Type: &ast.FuncType{
									Params: &ast.FieldList{
										List: []*ast.Field{
											{
												Names: []*ast.Ident{{Name: "param1"}},
												Type:  &ast.Ident{Name: "Int"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			imports: map[string]string{},
			functions: []*funcInfo{
				{
					function: &ast.FuncDecl{
						Name: &ast.Ident{Name: "test_func"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{{Name: "param1"}},
										Type:  &ast.Ident{Name: "Int"},
									},
								},
							},
						},
					},
					imports: map[string]string{},
					aliases: map[string]string{},
				},
			},
			expected: `pub fn __modus_test_func(param1 : Int) -> Unit {
  test_func(param1)
}

`,
		},
		{
			name: "single function with default params",
			pkg: &packages.Package{
				Name: "main",
				Syntax: []*ast.File{
					{
						Decls: []ast.Decl{
							&ast.FuncDecl{
								Name: &ast.Ident{Name: "test_func_WithDefaults"},
								Type: &ast.FuncType{
									Params: &ast.FieldList{
										List: []*ast.Field{
											{
												Names: []*ast.Ident{{Name: "param1"}},
												Type:  &ast.Ident{Name: "Int"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			imports: map[string]string{},
			functions: []*funcInfo{
				{
					function: &ast.FuncDecl{
						Name: &ast.Ident{Name: "test_func_WithDefaults"},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{{Name: "param1"}},
										Type:  &ast.Ident{Name: "Int"},
									},
								},
							},
						},
					},
					imports: map[string]string{},
					aliases: map[string]string{},
				},
			},
			expected: `pub fn __modus_test_func_WithDefaults(param1 : Int) -> Unit {
  test_func(param1)
}

`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			buf := &bytes.Buffer{}
			writeFuncWrappers(buf, tt.pkg, tt.imports, tt.functions)
			got := buf.String()
			if diff := cmp.Diff(tt.expected, got); diff != "" {
				t.Errorf("writeFuncWrappers() mismatch (-want +got):\n%v", diff)
			}
		})
	}
}
