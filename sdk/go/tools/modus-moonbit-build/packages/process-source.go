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
var importedHostFnRE = regexp.MustCompile(`(?m)^fn .*?\((.*?)\) (.*?)= "(modus_.*?)" "(.*?)"$`)
var pubStructRE = regexp.MustCompile(`(?ms)\npub.*? struct\s+(.*?)\s+{(.*?)\n}`)
var pubFnRE = regexp.MustCompile(`(?ms)\npub fn\s+(.*?)\s+{`)
var pubFnPrefixRE = regexp.MustCompile(`(?ms)\npub fn\s+(.*?)\(`)
var whiteSpaceRE = regexp.MustCompile(`(?ms)\s+`)

func (p *Package) processSourceFile(typesPkg *types.Package, filename string, buf []byte, imports []*ast.ImportSpec) {
	// add newlines to simplify regexp matching and normalize line endings
	fullSrc := "\n\n" + strings.ReplaceAll(string(buf), "\r\n", "\n")
	src := commentRE.ReplaceAllString(fullSrc, "")

	var decls []ast.Decl

	m := pubStructRE.FindAllStringSubmatch(src, -1)
	decls = p.processPubStructs(typesPkg, decls, m)

	m = pubFnRE.FindAllStringSubmatch(src, -1)
	decls = p.processPubFns(typesPkg, decls, m, fullSrc)

	m = importedHostFnRE.FindAllStringSubmatch(src, -1)
	decls = p.processImportedHostFns(typesPkg, decls, m)

	p.Syntax = append(p.Syntax, &ast.File{
		Name:    &ast.Ident{Name: filename},
		Decls:   decls,
		Imports: imports,
	})
}

func (p *Package) processPubStructs(typesPkg *types.Package, decls []ast.Decl, m [][]string) []ast.Decl {
	for _, match := range m {
		log.Printf("GML: processPubStructs: processing: %+v", match)
		name := match[1]
		var fields []*ast.Field
		var fieldVars []*types.Var
		allFields := strings.Split(match[2], "\n")
		for _, field := range allFields {
			field = commentRE.ReplaceAllString(field, "")
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

func (p *Package) processPubFns(typesPkg *types.Package, decls []ast.Decl, m [][]string, fullSrc string) []ast.Decl {
	for _, match := range m {
		docs := getDocsForFunction(match[0], fullSrc)
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

		newParamsList, newParamsVars := stripDefaultValues(paramsList, paramsVars)
		decls = p.addExportedFunctionDecls(typesPkg, decls, methodName, newParamsList, newParamsVars, resultsList, resultsTuple, docs)

		if functionParamsHasDefaultValue(paramsList) {
			newMethodName, newParamsList, newParamsVars := duplicateMethodWithoutDefaultParams(methodName, paramsList, paramsVars)
			decls = p.addExportedFunctionDecls(typesPkg, decls, newMethodName, newParamsList, newParamsVars, resultsList, resultsTuple, docs)
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

func stripDefaultValues(paramsList []*ast.Field, paramsVars []*types.Var) ([]*ast.Field, []*types.Var) {
	if len(paramsList) == 0 {
		return nil, nil
	}

	newParamsList := make([]*ast.Field, 0, len(paramsList))
	for _, param := range paramsList {
		if paramType, ok := param.Type.(*ast.Ident); ok {
			if strings.Contains(paramType.Name, "=") {
				// newName := strings.TrimSuffix(param.Names[0].Name, "~")  // cannot remove "~" yet, as it is needed in modus_pre_generated.mbt.
				paramTypeParts := strings.Split(paramType.Name, " = ")
				newType := paramTypeParts[0]
				field := &ast.Field{
					Names: []*ast.Ident{{Name: param.Names[0].Name}},
					Type:  &ast.Ident{Name: newType},
				}
				if len(paramTypeParts) > 1 {
					defaultValue := paramTypeParts[1]
					field.Tag = &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("`default:%v`", defaultValue)}
				}
				newParamsList = append(newParamsList, field)
			} else {
				newParamsList = append(newParamsList, param)
			}
		}
	}

	newParamsVars := make([]*types.Var, 0, len(paramsVars))
	for _, param := range paramsVars {
		var paramTypeName string
		switch paramType := param.Type().(type) {
		case *types.Named:
			paramTypeName = strings.Split(paramType.Obj().Name(), " = ")[0]
		default:
			log.Fatalf("programming error: expected *moonType, got %T", paramType)
		}
		// newName := strings.TrimSuffix(param.Name(), "~")  // still needed
		newParam := types.NewVar(0, nil, param.Name(), types.NewNamed(types.NewTypeName(0, nil, paramTypeName, nil), nil, nil))
		newParamsVars = append(newParamsVars, newParam)
	}

	return newParamsList, newParamsVars
}

func duplicateMethodWithoutDefaultParams(methodName string, paramsList []*ast.Field, paramsVars []*types.Var) (string, []*ast.Field, []*types.Var) {
	newMethodName := methodName + "_WithDefaults"
	if len(paramsList) == 0 {
		return newMethodName, nil, nil
	}
	newParamsList := make([]*ast.Field, 0, len(paramsList))
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
	newParamsVars := make([]*types.Var, 0, len(paramsVars))
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

func (p *Package) addExportedFunctionDecls(typesPkg *types.Package, decls []ast.Decl, methodName string, paramsList []*ast.Field, paramsVars []*types.Var, resultsList *ast.FieldList, resultsTuple *types.Tuple, docs *ast.CommentGroup) []ast.Decl {
	decl := &ast.FuncDecl{
		Name: &ast.Ident{Name: methodName},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: paramsList,
			},
			Results: resultsList,
		},
		Doc: docs,
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
	allArgParts := splitFunctionParameters(allArgs)
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

func stripError(typeSignature string) string {
	if i := strings.Index(typeSignature, "!"); i >= 0 {
		return typeSignature[:i]
	}
	return typeSignature
}

func (p *Package) checkCustomMoonBitType(typesPkg *types.Package, typeSignature string) (resultType *types.Named) {
	fullCustomName := typesPkg.Path() + "." + typeSignature
	customName := stripError(fullCustomName)
	for ident, customType := range p.TypesInfo.Defs {
		if ident.Name == customName {
			if customType, ok := customType.(*types.TypeName); ok {
				underlying := customType.Type().Underlying()
				if customName == fullCustomName {
					resultType = types.NewNamed(customType, underlying, nil)
				} else {
					fullCustomType := types.NewTypeName(0, typesPkg, typeSignature, nil)
					resultType = types.NewNamed(fullCustomType, underlying, nil)
				}
				// log.Printf("GML: checkCustomMoonBitNamedType: found custom type for typeSignature=%q", typeSignature)
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
