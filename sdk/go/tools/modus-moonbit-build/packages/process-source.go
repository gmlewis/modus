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

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
)

var argsRE = regexp.MustCompile(`^(.*?)\((.*)\)$`)
var commentRE = regexp.MustCompile(`(?m)^\s*//.*$`)
var importedHostFnRE = regexp.MustCompile(`(?ms)^fn [^\{]*?\(([^\{]*?)\) ([^\{]*?)= "(modus_.*?)" "(.*?)"$`)
var pubStructRE = regexp.MustCompile(`(?ms)\npub(?:\([a-z]+\))? struct\s+(.*?)\s+{(.*?)\n}`)
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
	decls = p.processImportedHostFns(typesPkg, decls, m, fullSrc)

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
			field = commentRE.ReplaceAllString(field, "")
			if i := strings.Index(field, "//"); i >= 0 { // strip comments
				field = field[:i]
			}
			field = strings.TrimSpace(field)
			if field == "" {
				continue
			}
			fieldParts := strings.Split(field, ":")
			if len(fieldParts) != 2 {
				continue
			}
			fieldName := strings.TrimSpace(fieldParts[0])
			fieldTypeName := strings.TrimSpace(fieldParts[1])
			fieldType := p.getMoonBitNamedType(typesPkg, fieldTypeName) // resolve forward references later.
			fullyQualifiedFieldSig := fieldType.String()
			fields = append(fields, &ast.Field{
				Names: []*ast.Ident{{Name: fieldName}},
				Type:  &ast.Ident{Name: fullyQualifiedFieldSig},
			})
			fieldVars = append(fieldVars, types.NewVar(0, nil, fieldName, fieldType))
			// Add types for all field subtypes:
			_, acc := utils.FullyQualifyTypeName(typesPkg.Path(), fieldTypeName)
			for k := range acc {
				p.AddPossiblyMissingUnderlyingType(k)
			}
		}

		typeSpec := &ast.TypeSpec{
			Name: &ast.Ident{Name: fullyQualifiedName(typesPkg, name)},
			// Name: &ast.Ident{Name: name},  // Do NOT add typesPkg.
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

		// This is the new `types.Struct` with the new fields added.
		decls = append(decls, decl)
		underlying := types.NewStruct(fieldVars, nil)
		newNamedTypeName, namedType := fullyQualifiedNewTypeName(typesPkg, name, underlying)

		p.replaceForwardReferences(newNamedTypeName, namedType)
		p.StructLookup[newNamedTypeName] = typeSpec
		p.TypesInfo.Defs[typeSpec.Name] = namedType
	}

	return decls
}

func (p *Package) processPubFns(typesPkg *types.Package, decls []ast.Decl, m [][]string, fullSrc string) []ast.Decl {
	for _, match := range m {
		docs := getDocsForFunction(match[0], fullSrc)
		fnSig := whiteSpaceRE.ReplaceAllString(match[1], " ")
		parts := strings.Split(fnSig, " -> ")
		if len(parts) != 2 {
			continue
		}
		fnSig = strings.TrimSpace(parts[0])
		returnSig := strings.TrimSpace(parts[1])
		ma := argsRE.FindStringSubmatch(fnSig)
		if len(ma) != 3 {
			continue
		}
		methodName := strings.TrimSpace(ma[1])
		allArgs := strings.TrimSpace(ma[2])

		resultsList, resultsTuple := p.processReturnSignature(typesPkg, returnSig)
		paramsList, paramsVars := p.processParameters(typesPkg, allArgs)

		newParamsList, newParamsVars := p.stripDefaultValues(typesPkg, paramsList, paramsVars)
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

func (p *Package) stripDefaultValues(typesPkg *types.Package, paramsList []*ast.Field, paramsVars []*types.Var) ([]*ast.Field, []*types.Var) {
	if len(paramsList) == 0 {
		return nil, nil
	}

	newParamsList := make([]*ast.Field, 0, len(paramsList))
	for _, param := range paramsList {
		if paramType, ok := param.Type.(*ast.Ident); ok {
			if strings.Contains(paramType.Name, "=") {
				// newName := strings.TrimSuffix(param.Names[0].Name, "~")  // cannot remove "~" yet, as it is needed in modus_pre_generated.mbt.
				paramTypeParts := strings.Split(paramType.Name, " = ")
				newTypeName := strings.TrimSpace(paramTypeParts[0])
				newType := p.getMoonBitNamedType(typesPkg, newTypeName)
				fullyQualifiedNewSig := newType.String()
				field := &ast.Field{
					Names: []*ast.Ident{{Name: param.Names[0].Name}},
					Type:  &ast.Ident{Name: fullyQualifiedNewSig},
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
		paramTypeName := param.Type().String()
		if strings.Contains(paramTypeName, "=") {
			switch paramType := param.Type().(type) {
			case *types.Named:
				// In Go, a function parameter cannot have a default value
				// so there is no way in the Go AST to represent this.
				// Therefore, pass the default value as part of the type
				// of this parameter, then convert it later in extractor/transform.go.
				paramTypeName = paramType.Obj().Name()
			default:
				log.Fatalf("PROGRAMMING ERROR: expected *types.Named, got %T", paramType)
			}
			// newName := strings.TrimSuffix(param.Name(), "~")  // still needed - do not remove
			newParam := types.NewVar(0, nil, param.Name(), types.NewNamed(types.NewTypeName(0, nil, paramTypeName, nil), nil, nil))
			newParamsVars = append(newParamsVars, newParam)
		} else {
			// Make a copy:
			newParam := types.NewVar(0, param.Pkg(), param.Name(), param.Type())
			newParamsVars = append(newParamsVars, newParam)
		}
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
		types.NewSignatureType(nil, nil, nil, paramsTuple, resultsTuple, false))

	return decls
}

func (p *Package) processParameters(typesPkg *types.Package, allArgs string) (paramsList []*ast.Field, paramsVars []*types.Var) {
	allArgParts := utils.SplitParamsWithBrackets(allArgs)
	for _, arg := range allArgParts {
		argParts := strings.Split(arg, ":")
		if len(argParts) != 2 {
			continue
		}
		argName := strings.TrimSpace(argParts[0])
		argTypeName := strings.TrimSpace(argParts[1])
		paramType := p.getMoonBitNamedType(typesPkg, argTypeName)
		fullyQualifiedParamSig := paramType.String()
		paramsList = append(paramsList, &ast.Field{
			Names: []*ast.Ident{{Name: argName}},
			Type:  &ast.Ident{Name: fullyQualifiedParamSig},
		})
		paramsVars = append(paramsVars, types.NewVar(0, nil, argName, paramType))
	}

	return paramsList, paramsVars
}

func (p *Package) processReturnSignature(typesPkg *types.Package, returnSig string) (resultsList *ast.FieldList, resultsTuple *types.Tuple) {
	if returnSig == "Unit" {
		return nil, nil
	}

	resultType := p.getMoonBitNamedType(typesPkg, returnSig)
	fullyQualifiedReturnSig := resultType.String()

	resultsList = &ast.FieldList{
		List: []*ast.Field{{Type: &ast.Ident{Name: fullyQualifiedReturnSig}}},
	}

	resultsTuple = types.NewTuple(types.NewVar(0, nil, "", resultType))

	return resultsList, resultsTuple
}

func (p *Package) checkCustomMoonBitType(typesPkg *types.Package, typeSignature string) (resultType types.Type) {
	typeSignature = utils.StripDefaultValue(typeSignature)
	if strings.Contains(typeSignature, ".") { // already has package name - must be external, so ignore.
		return nil
	}
	typ, _, _ := utils.StripErrorAndOption(typeSignature)

	if typeSpec, ok := p.StructLookup[typ]; ok {
		if customType, ok := p.TypesInfo.Defs[typeSpec.Name].(*types.TypeName); ok {
			underlying := customType.Type().Underlying()
			if typ == typeSignature {
				return types.NewNamed(customType, underlying, nil)
			}
			_, fullCustomType := fullyQualifiedNewTypeName(typesPkg, typeSignature, nil)
			return types.NewNamed(fullCustomType, underlying, nil)
		}
		log.Fatalf("PROGRAMMING ERROR: checkCustomMoonBitType(typeSignature='%v'): typ '%v' missing from p.TypesInfo.Defs", typeSignature, typ)
	}

	if utils.IsStructType(typ) {
		// This could possibly be a forward reference or a recursive reference within a struct.
		// Create an entry for it and point to it as the underlying type.
		// Its fields will be filled in later.
		underlying := types.NewStruct(nil, nil)
		{
			newCustomTypeName, customType := fullyQualifiedNewTypeName(typesPkg, typ, underlying)
			if v, ok := p.StructLookup[newCustomTypeName]; ok {
				customType, ok = p.TypesInfo.Defs[v.Name].(*types.TypeName)
				if !ok {
					log.Fatalf("PROGRAMMING ERROR! p.StructLookup and p.TypesInfo.Def out-of-sync!")
				}
			} else {
				typeSpec := &ast.TypeSpec{Name: &ast.Ident{Name: newCustomTypeName}}
				p.StructLookup[newCustomTypeName] = typeSpec
				p.TypesInfo.Defs[typeSpec.Name] = customType
			}
			if typ == typeSignature {
				return types.NewNamed(customType, underlying, nil)
			}
		}
		{
			newFullCustomTypeName, fullCustomType := fullyQualifiedNewTypeName(typesPkg, typeSignature, underlying)
			if v, ok := p.StructLookup[newFullCustomTypeName]; ok {
				fullCustomType, ok = p.TypesInfo.Defs[v.Name].(*types.TypeName)
				if !ok {
					log.Fatalf("PROGRAMMING ERROR! p.StructLookup and p.TypesInfo.Def out-of-sync!")
				}
			} else {
				fullTypeSpec := &ast.TypeSpec{Name: &ast.Ident{Name: newFullCustomTypeName}}
				p.StructLookup[newFullCustomTypeName] = fullTypeSpec
				p.TypesInfo.Defs[fullTypeSpec.Name] = fullCustomType
			}
			return types.NewNamed(fullCustomType, underlying, nil)
		}
	}

	return nil
}

func (p *Package) GetMoonBitNamedType(typeSignature string) (resultType types.Type) {
	typesPkg := types.NewPackage(p.PkgPath, p.Name)
	return p.getMoonBitNamedType(typesPkg, typeSignature)
}

func (p *Package) getMoonBitNamedType(typesPkg *types.Package, typeSignature string) (resultType types.Type) {
	// Treat a MoonBit tuple like a struct whose field names are "0", "1", etc.
	if strings.HasPrefix(typeSignature, "(") && strings.HasSuffix(typeSignature, ")") {
		allArgs := typeSignature[1 : len(typeSignature)-1]
		allTupleParts := utils.SplitParamsWithBrackets(allArgs)
		var fieldVars []*types.Var
		for i, field := range allTupleParts {
			fieldType := strings.TrimSpace(commentRE.ReplaceAllString(field, ""))
			if fieldType == "" {
				continue
			}
			fieldName := fmt.Sprintf("%v", i)
			underlying := p.getMoonBitNamedType(typesPkg, fieldType)
			fieldVars = append(fieldVars, types.NewVar(0, nil, fieldName, underlying))
		}

		tupleStruct := types.NewStruct(fieldVars, nil)
		return types.NewNamed(types.NewTypeName(0, nil, typeSignature, nil), tupleStruct, nil)
	}

	if v, _ := utils.FullyQualifyTypeName(typesPkg.Path(), typeSignature); !strings.HasPrefix(v, "@") {
		// Do not use the generated name if it is a package-level struct, as the next line will add
		// the full typesPkg qualified name to the prefix. However, if it is something like `Array[@pkg.T]`
		// then go ahead and use it.
		typeSignature = v
	}

	resultType = p.checkCustomMoonBitType(typesPkg, typeSignature)
	if resultType == nil {
		resultType = types.NewNamed(types.NewTypeName(0, nil, typeSignature, nil), nil, nil)
	}
	return resultType
}

type moonType struct {
	typeName string
}

func (t *moonType) Underlying() types.Type {
	return t
}

func (t *moonType) String() string {
	return t.typeName
}

func (p *Package) replaceForwardReferences(newNamedTypeName string, namedType *types.TypeName) {
	for k, v := range p.StructLookup {
		if k != newNamedTypeName {
			continue
		}
		// Update the old references to the new namedType.
		p.TypesInfo.Defs[v.Name] = namedType
	}
}
