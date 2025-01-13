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
	"go/token"
	"io"
)

// Fprint "pretty-prints" a Go AST node to w as MoonBit source code.
func Fprint(w io.Writer, fset *token.FileSet, node any) {
	fmt.Fprintf(w, "printer.Fprint: TODO: %T", node)
}
