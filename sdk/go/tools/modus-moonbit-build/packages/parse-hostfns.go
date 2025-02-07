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
	"go/ast"
	"go/types"
	"log"
	"strings"
)

func (p *Package) processImportedHostFns(typesPkg *types.Package, decls []ast.Decl, m [][]string) []ast.Decl {
	for _, match := range m {
		log.Printf("GML: processImportedHostFns: processing: %+v", match)
		if len(match) != 5 {
			log.Printf("PROGRAMMING ERROR: len(match) != 5: %+v", match)
			continue
		}
		returnSig := strings.TrimSpace(strings.TrimPrefix(match[2], "->"))
		if returnSig == "" {
			returnSig = "Unit"
		}
		hostEnv := match[3]
		hostFn := match[4]
		methodName := hostEnv + "." + hostFn
		allArgs := strings.TrimSpace(match[1])
		log.Printf("GML: imported host function: %v(%v) -> %v", methodName, allArgs, returnSig)

		docs := &ast.CommentGroup{ // trigger this to be a host import function
			List: []*ast.Comment{
				{Text: "//go:wasmimport " + hostEnv + " " + hostFn},
			},
		}

		resultsList, resultsTuple := p.processReturnSignature(typesPkg, returnSig)
		paramsList, paramsVars := p.processParameters(typesPkg, allArgs)

		newParamsList, newParamsVars := p.stripDefaultValues(typesPkg, paramsList, paramsVars)
		decls = p.addExportedFunctionDecls(typesPkg, decls, methodName, newParamsList, newParamsVars, resultsList, resultsTuple, docs)

		// if functionParamsHasDefaultValue(paramsList) {
		// 	newMethodName, newParamsList, newParamsVars := duplicateMethodWithoutDefaultParams(methodName, paramsList, paramsVars)
		// 	decls = p.addExportedFunctionDecls(typesPkg, decls, newMethodName, newParamsList, newParamsVars, resultsList, resultsTuple, docs)
		// }
	}

	return decls
}
