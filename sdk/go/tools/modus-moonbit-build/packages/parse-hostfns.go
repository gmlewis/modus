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
	"regexp"
	"strings"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
)

func (p *Package) processImportedHostFns(typesPkg *types.Package, decls []ast.Decl, m [][]string, fullSrc string) []ast.Decl {
	for _, match := range m {
		if len(match) != 5 {
			log.Fatalf("PROGRAMMING ERROR: len(match) != 5: %+v", match)
			continue
		}

		fnDetails := extractFnDetails(match, fullSrc)

		goDocs := &ast.CommentGroup{ // trigger this to be a host import function
			List: []*ast.Comment{
				{Text: "//go:wasmimport " + fnDetails.HostEnv + " " + fnDetails.HostFn},
			},
		}

		resultsList, resultsTuple := p.processReturnSignature(typesPkg, fnDetails.ReturnSig)
		paramsList, paramsVars := p.processParameters(typesPkg, fnDetails.AllArgs)

		newParamsList, newParamsVars := p.stripDefaultValues(typesPkg, paramsList, paramsVars)
		decls = p.addExportedFunctionDecls(typesPkg, decls, fnDetails.MethodName, newParamsList, newParamsVars, resultsList, resultsTuple, goDocs)

		// if functionParamsHasDefaultValue(paramsList) {
		// 	newMethodName, newParamsList, newParamsVars := duplicateMethodWithoutDefaultParams(methodName, paramsList, paramsVars)
		// 	decls = p.addExportedFunctionDecls(typesPkg, decls, newMethodName, newParamsList, newParamsVars, resultsList, resultsTuple, goDocs)
		// }
	}

	return decls
}

type fnDetailsT struct {
	ReturnSig  string
	HostEnv    string
	HostFn     string
	MethodName string
	AllArgs    string
}

var modusImportHostFnRE = regexp.MustCompile(`(?m)^// modus:import (.*?)\s+(.*?)\((.*?)\) (.*?)$`)

func extractFnDetails(match []string, fullSrc string) (fnDetails fnDetailsT) {
	origDetails := parseMatchFnDetails(match)
	docs := getDocsForFunction(match[0], fullSrc)
	if docs != nil {
		for _, line := range docs.List {
			m := modusImportHostFnRE.FindStringSubmatch(line.Text)
			if len(m) != 5 {
				continue
			}
			parts := []string{"", m[3], m[4], m[1], m[2]}
			fnDetails := parseMatchFnDetails(parts)
			if fnDetails.MethodName != origDetails.MethodName {
				log.Printf("WARNING: method names differ: %v != %v", fnDetails.MethodName, origDetails.MethodName)
			}
			if fnDetails.AllArgs != origDetails.AllArgs {
				fnDetails.AllArgs = resolveArgsDiffs(origDetails.AllArgs, fnDetails.AllArgs)
			}
			return fnDetails
		}
	}
	log.Printf("WARNING: no modus:import comment found for %v(%v) -> %v - skipping", origDetails.MethodName, origDetails.AllArgs, origDetails.ReturnSig)
	return origDetails
}

func parseMatchFnDetails(match []string) (fnDetails fnDetailsT) {
	fnDetails.ReturnSig = strings.TrimSpace(strings.TrimPrefix(match[2], "->"))
	if fnDetails.ReturnSig == "" {
		fnDetails.ReturnSig = "Unit"
	}
	fnDetails.HostEnv = match[3]
	fnDetails.HostFn = match[4]
	fnDetails.MethodName = fnDetails.HostEnv + "." + fnDetails.HostFn
	fnDetails.AllArgs = strings.TrimSpace(match[1])
	return fnDetails
}

func resolveArgsDiffs(origArgs, newArgs string) string {
	orig := utils.SplitParamsWithBrackets(origArgs)
	newA := utils.SplitParamsWithBrackets(newArgs)
	if len(orig) != len(newA) {
		log.Printf("WARNING: unable to resolve old (%v) and new (%v) args", orig, newA)
		return newArgs
	}
	result := make([]string, 0, len(orig))
	for i, s := range orig {
		newParts := strings.Split(newA[i], ":")
		if len(newParts) == 2 {
			result = append(result, newA[i])
			continue
		}
		origParts := strings.Split(s, ":")
		result = append(result, strings.TrimSpace(origParts[0])+" : "+strings.TrimSpace(newParts[0]))
	}
	return strings.Join(result, ", ")
}
