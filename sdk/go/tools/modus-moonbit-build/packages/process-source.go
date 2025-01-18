// This file is based on: https://cs.opensource.google/go/x/tools/+/refs/tags/v0.28.0:go/packages/packages.go
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packages

// See doc.go for package documentation and implementation notes.

import (
	"fmt"
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
		allArgs := strings.TrimSpace(ma[2])
		log.Printf("GML: %v(%v) -> %v", methodName, allArgs, returnSig)

		resultsList, resultsTuple := p.processReturnSignature(typesPkg, returnSig)
		paramsList, paramsVars := p.processParameters(typesPkg, allArgs)

		decls = p.addExportedFunctionDecls(typesPkg, decls, methodName, paramsList, paramsVars, resultsList, resultsTuple)
		if functionParamsHasDefaultValue(paramsList) {
			newMethodName, newParamsList, newParamsVars := duplicateMethodWithDefaultValues(methodName, paramsList, paramsVars)
			decls = p.addExportedFunctionDecls(typesPkg, decls, newMethodName, newParamsList, newParamsVars, resultsList, resultsTuple)
		}
	}

	return decls
}

func functionParamsHasDefaultValue(paramsList []*ast.Field) bool {
	for _, param := range paramsList {
		paramType, ok := param.Type.(*ast.Ident)
		if ok && strings.Contains(paramType.Name, "=") {
			return true
		}
	}
	return false
}

func duplicateMethodWithDefaultValues(methodName string, paramsList []*ast.Field, paramsVars []*types.Var) (string, []*ast.Field, []*types.Var) {
	newMethodName := methodName + "_WithDefaults"
	newParamsList := make([]*ast.Field, 0, len(paramsList))
	newParamsVars := make([]*types.Var, 0, len(paramsVars))
	for _, param := range paramsList {
		paramType, ok := param.Type.(*ast.Ident)
		if !ok {
			log.Fatalf("programming error: expected *ast.Ident, got %T", param.Type)
		}
		if strings.Contains(paramType.Name, "=") {
			continue
		}
		newParam := &ast.Field{
			Names: param.Names,
			Type:  param.Type,
		}
		newParamsList = append(newParamsList, newParam)
	}
	for _, param := range paramsVars {
		switch paramType := param.Type().(type) {
		case *types.Named:
			paramTypeName := paramType.Obj().Name()
			if strings.Contains(paramTypeName, "=") {
				continue
			}
		default:
			log.Fatalf("programming error: expected *moonType, got %T", paramType)
		}
		newParam := types.NewVar(0, nil, param.Name(), param.Type())
		newParamsVars = append(newParamsVars, newParam)
	}
	return newMethodName, newParamsList, newParamsVars
}

func (p *Package) addExportedFunctionDecls(typesPkg *types.Package, decls []ast.Decl, methodName string, paramsList []*ast.Field, paramsVars []*types.Var, resultsList *ast.FieldList, resultsTuple *types.Tuple) []ast.Decl {
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

	return decls
}

func (p *Package) processParameters(typesPkg *types.Package, allArgs string) (paramsList []*ast.Field, paramsVars []*types.Var) {
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
		paramType := p.getMoonBitNamedType(typesPkg, argType)
		paramsVars = append(paramsVars, types.NewVar(0, nil, argName, paramType)) // &moonType{typeName: argType}))
	}

	return paramsList, paramsVars
}

func (p *Package) processReturnSignature(typesPkg *types.Package, returnSig string) (resultsList *ast.FieldList, resultsTuple *types.Tuple) {
	if returnSig != "Unit" {
		resultsList = &ast.FieldList{
			List: []*ast.Field{{Type: &ast.Ident{Name: returnSig}}},
		}
		resultType := p.getMoonBitNamedType(typesPkg, returnSig)
		resultsTuple = types.NewTuple(types.NewVar(0, nil, "", resultType))
	}

	return resultsList, resultsTuple
}

func (p *Package) checkCustomMoonBitType(typesPkg *types.Package, typeSignature string) (resultType *types.Named) {
	customName := typesPkg.Path() + "." + typeSignature
	for ident, customType := range p.TypesInfo.Defs {
		if ident.Name == customName {
			if customType, ok := customType.(*types.TypeName); ok {
				underlying := customType.Type().Underlying()
				resultType = types.NewNamed(customType, underlying, nil)
				log.Printf("GML: checkCustomMoonBitNamedType: found custom type for typeSignature=%q", typeSignature)
				break
			}
		}
	}
	return resultType
}

func (p *Package) getMoonBitNamedType(typesPkg *types.Package, typeSignature string) (resultType *types.Named) {
	log.Printf("GML: getMoonBitNamedType(typeSignature=%q)", typeSignature)
	// TODO: write a parser that can handle any MoonBit type.
	if strings.HasPrefix(typeSignature, "Array[") && strings.HasSuffix(typeSignature, "]") {
		tmpSignature := typeSignature[6 : len(typeSignature)-1]
		resultType = p.checkCustomMoonBitType(typesPkg, tmpSignature)
		if resultType != nil {
			tmpSignature = fmt.Sprintf("Array[%v.%v]", typesPkg.Path(), tmpSignature)
			return types.NewNamed(types.NewTypeName(0, nil, tmpSignature, nil), nil, nil)
		}
	}

	resultType = p.checkCustomMoonBitType(typesPkg, typeSignature)
	if resultType == nil {
		resultType = types.NewNamed(types.NewTypeName(0, nil, typeSignature, nil), nil, nil) // &moonType{typeName: typeSignature}
	}
	return resultType
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
