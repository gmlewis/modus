/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// Package printer provides functionality to print Go AST nodes
// as MoonBit source code.
package printer

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"log"
	"strings"
)

// Fprint "pretty-prints" a Go AST node to w as MoonBit source code.
func Fprint(w io.Writer, fset *token.FileSet, node any) {
	var params []string
	resultType := "Unit"
	switch n := node.(type) {
	case *ast.FuncType:
		for _, param := range n.Params.List {
			params = append(params, fmt.Sprintf("%v : %v", param.Names[0].Name, param.Type))
		}
		if n.Results != nil {
			resultType = n.Results.List[0].Type.(*ast.Ident).Name
		}
		fmt.Fprintf(w, "(%v) -> %v", strings.Join(params, ", "), resultType)
	default:
		log.Printf("WARNING: printer.FPrint: unhandled node type %T\n", node)
	}
}
