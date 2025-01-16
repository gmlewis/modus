// This file is based on: https://cs.opensource.google/go/x/tools/+/refs/tags/v0.28.0:go/packages/packages.go
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packages

// See doc.go for package documentation and implementation notes.

import (
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"regexp"
	"strings"
)

var argsRE = regexp.MustCompile(`^(.*?)\((.*)\)$`)
var commentRE = regexp.MustCompile(`(?m)^\s*//.*$`)
var pubStructRE = regexp.MustCompile(`(?ms)\npub\(.*?\) struct\s+(.*?)\s+{(.*?)\n}`)
var pubFnRE = regexp.MustCompile(`(?ms)\npub fn\s+(.*?)\s+{`)
var whiteSpaceRE = regexp.MustCompile(`(?ms)\s+`)

func (p *Package) processSourceFile(typesPkg *types.Package, filename string, buf []byte, imports []*ast.ImportSpec) {
	src := "\n" + string(buf)
	src = commentRE.ReplaceAllString(src, "")

	var decls []ast.Decl

	m := pubStructRE.FindAllStringSubmatch(src, -1)
	decls = p.processPubStructs(typesPkg, decls, m)

	m = pubFnRE.FindAllStringSubmatch(src, -1)
	decls = p.processPubFns(typesPkg, decls, m)

	p.Syntax = append(p.Syntax, &ast.File{
		Name:    &ast.Ident{Name: filename},
		Decls:   decls,
		Imports: imports,
	})
}

func (p *Package) processPubStructs(typesPkg *types.Package, decls []ast.Decl, m [][]string) []ast.Decl {
	for _, match := range m {
		name := match[1]
		var fields []*ast.Field
		var fieldVars []*types.Var
		allFields := strings.Split(match[2], "\n")
		for _, field := range allFields {
			field = strings.TrimSpace(field)
			if field == "" {
				continue
			}
			fieldParts := strings.Split(field, ":")
			if len(fieldParts) != 2 {
				log.Printf("Warning: invalid field: '%v'; skipping", field)
				continue
			}
			fieldName := strings.TrimSpace(fieldParts[0])
			fieldType := strings.TrimSpace(fieldParts[1])
			fields = append(fields, &ast.Field{
				Names: []*ast.Ident{{Name: fieldName}},
				Type:  &ast.Ident{Name: fieldType},
			})
			fieldVars = append(fieldVars, types.NewVar(0, nil, fieldName, &moonType{typeName: fieldType}))
		}

		typeSpec := &ast.TypeSpec{
			Name: &ast.Ident{Name: typesPkg.Path() + "." + name},
			Type: &ast.StructType{
				Fields: &ast.FieldList{
					List: fields,
				},
			},
		}
		decl := &ast.GenDecl{
			Tok:   token.TYPE,
			Specs: []ast.Spec{typeSpec},
		}

		decls = append(decls, decl)
		underlying := types.NewStruct(fieldVars, nil)
		p.TypesInfo.Defs[typeSpec.Name] = types.NewTypeName(0, typesPkg, name, underlying) // TODO
	}

	return decls
}

func (p *Package) processPubFns(typesPkg *types.Package, decls []ast.Decl, m [][]string) []ast.Decl {
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
		var resultsTuple *types.Tuple
		if returnSig != "Unit" {
			resultsList = &ast.FieldList{
				List: []*ast.Field{
					{
						Type: &ast.Ident{Name: returnSig},
					},
				},
			}
			var resultType *types.Named
			customName := typesPkg.Path() + "." + returnSig
			for ident, customType := range p.TypesInfo.Defs {
				if ident.Name == customName {
					if customType, ok := customType.(*types.TypeName); ok {
						underlying := customType.Type().Underlying()
						resultType = types.NewNamed(customType, underlying, nil)
						break
					}
				}
			}
			if resultType == nil {
				resultType = types.NewNamed(types.NewTypeName(0, nil, returnSig, nil), nil, nil) // &moonType{typeName: returnSig}
			}
			// resultType := types.NewNamed(types.NewTypeName(0, typesPkg, returnSig, nil), nil, nil) // &moonType{typeName: returnSig}
			resultsTuple = types.NewTuple(types.NewVar(0, nil, "", resultType))
		}

		var paramsList []*ast.Field
		var paramsVars []*types.Var

		allArgs := strings.TrimSpace(ma[2])
		// log.Printf("GML: %v(%v) -> %v", methodName, allArgs, returnSig)
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
			paramsVars = append(paramsVars, types.NewVar(0, nil, argName, &moonType{typeName: argType}))
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
		var paramsTuple *types.Tuple
		if len(paramsVars) > 0 {
			paramsTuple = types.NewTuple(paramsVars...)
		}
		p.TypesInfo.Defs[decl.Name] = types.NewFunc(0, typesPkg, methodName,
			types.NewSignatureType(nil, nil, nil, paramsTuple, resultsTuple, false)) // TODO
	}

	return decls
}

type moonType struct {
	typeName string
}

func (t *moonType) Underlying() types.Type {
	return t // TODO
}

func (t *moonType) String() string {
	return t.typeName
}
