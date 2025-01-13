// This file is based on: https://cs.opensource.google/go/x/tools/+/refs/tags/v0.28.0:go/packages/packages.go
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packages

// See doc.go for package documentation and implementation notes.

import (
	"fmt"
	"go/ast"
	"go/types"
	"log"
	"regexp"
	"strings"
)

var argsRE = regexp.MustCompile(`^(.*?)\((.*)\)$`)
var commentRE = regexp.MustCompile(`(?m)^\s+//.*$`)
var pubFnRE = regexp.MustCompile(`(?ms)\npub fn\s+(.*?)\s+{`)
var whiteSpaceRE = regexp.MustCompile(`(?ms)\s+`)

func (p *Package) processSourceFile(typesPkg *types.Package, filename string, buf []byte, imports []*ast.ImportSpec) {
	src := "\n" + string(buf)
	src = commentRE.ReplaceAllString(src, "")
	m := pubFnRE.FindAllStringSubmatch(src, -1)

	var decls []ast.Decl
	var resultsTuple *types.Tuple
	for _, match := range m {
		fnSig := whiteSpaceRE.ReplaceAllString(match[1], " ")
		parts := strings.Split(fnSig, " -> ")
		if len(parts) != 2 {
			log.Printf("Warning: invalid function signature: '%v'; skipping", fnSig)
			continue
		}
		fnSig = strings.TrimSpace(parts[0])
		returnSig := strings.TrimSpace(parts[1])
		ma := argsRE.FindStringSubmatch(fnSig)
		if len(ma) != 3 {
			log.Printf("Warning: invalid function signature: '%v'; skipping", fnSig)
			continue
		}
		methodName := strings.TrimSpace(ma[1])

		var resultsList *ast.FieldList
		if returnSig != "Unit" {
			resultsList = &ast.FieldList{
				List: []*ast.Field{
					{
						Type: &ast.Ident{Name: returnSig},
					},
				},
			}
			resultType := &moonType{argType: returnSig}
			resultsTuple = types.NewTuple(types.NewVar(0, nil, "", resultType))
		}

		var paramsList []*ast.Field
		var paramsVars []*types.Var

		allArgs := strings.TrimSpace(ma[2])
		log.Printf("GML: %v(%v) -> %v", methodName, allArgs, returnSig) // TODO(gmlewis): remove
		allArgParts := strings.Split(allArgs, ",")
		for _, arg := range allArgParts {
			arg = strings.TrimSpace(arg)
			if arg == "" {
				continue
			}
			argParts := strings.Split(arg, ":")
			if len(argParts) != 2 {
				log.Printf("Warning: invalid argument: '%v'; skipping", arg)
				continue
			}
			argName := strings.TrimSpace(argParts[0])
			argType := strings.TrimSpace(argParts[1])
			paramsList = append(paramsList, &ast.Field{
				Names: []*ast.Ident{{Name: argName}},
				Type:  &ast.Ident{Name: argType},
			})
			paramsVars = append(paramsVars, types.NewVar(0, nil, argName, &moonType{argName: argName, argType: argType}))
		}

		decl := &ast.FuncDecl{
			Name: &ast.Ident{Name: methodName},
			Type: &ast.FuncType{
				Params: &ast.FieldList{
					List: paramsList,
				},
				Results: resultsList,
			},
		}

		decls = append(decls, decl)
		paramsTuple := types.NewTuple(paramsVars...)
		p.TypesInfo.Defs[decl.Name] = types.NewFunc(0, typesPkg, methodName,
			types.NewSignatureType(nil, nil, nil, paramsTuple, resultsTuple, false)) // TODO
	}

	p.Syntax = append(p.Syntax, &ast.File{
		Name:    &ast.Ident{Name: filename},
		Decls:   decls,
		Imports: imports,
	})
}

type moonType struct {
	argName string
	argType string
}

func (t *moonType) Underlying() types.Type {
	return t // TODO
}

func (t *moonType) String() string {
	if t.argName == "" {
		return t.argType
	}
	return fmt.Sprintf("%v: %v", t.argName, t.argType)
}
