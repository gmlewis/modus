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
var importedHostFnRE = regexp.MustCompile(`(?m)^fn .*?\((.*?)\) (.*?)= "(modus_.*?)" "(.*?)"$`)
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
			fieldTypeName := strings.TrimSpace(fieldParts[1])
			fieldType := p.getMoonBitNamedType(typesPkg, fieldTypeName) // TODO: How to handle forward references? Resolve in 2nd pass?!?
			fullyQualifiedFieldSig := fieldType.String()
			fields = append(fields, &ast.Field{
				Names: []*ast.Ident{{Name: fieldName}},
				Type:  &ast.Ident{Name: fullyQualifiedFieldSig},
			})
			fieldVars = append(fieldVars, types.NewVar(0, nil, fieldName, fieldType))
		}

		typeSpec := &ast.TypeSpec{
			// Name: &ast.Ident{Name: typesPkg.Path() + "." + name},  // Do NOT add typesPkg.
			Name: &ast.Ident{Name: name},
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
		namedType := types.NewTypeName(0, nil, name, underlying) // Do NOT add typesPkg.

		// // In Go, the `types.Struct` type is immutable, meaning that all old references
		// // to an underlying `*types.Struct` need to be replaced with a pointer to a
		// // new `*types.Struct` that has the new fields added.
		// if emptyTypeSpec, ok := p.StructLookup[typeSpec.Name.Name]; ok {
		// 	if emptyNamedType, ok := p.TypesInfo.Defs[emptyTypeSpec.Name].(*types.TypeName); ok {
		// 		emptyStruct, ok := emptyNamedType.Type().Underlying().(*types.Struct)
		// 		if ok {
		// 			log.Printf("GML: processPubStructs: updating existing struct: emptyNamedType: %p, emptyStruct: %p, emptyTypeSpec: %p, namedType: %p, typeSpec: %p", emptyNamedType, emptyStruct, emptyTypeSpec, namedType, typeSpec)
		// 			// Now we have found the old, empty struct. Find all references to it and replace them.
		// 			p.replaceEmptyStructReferences(emptyNamedType, emptyTypeSpec, emptyStruct, namedType, typeSpec)
		// 		} else {
		// 			log.Printf("PROGRAMMING ERROR: expected *types.Struct, got %T", emptyNamedType.Type().Underlying())
		// 		}
		// 	} else {
		// 		log.Printf("PROGRAMMING ERROR: expected *types.TypeName, got %T", p.TypesInfo.Defs[emptyTypeSpec.Name])
		// 	}
		// }

		log.Printf("GML: packages/process-source.go: processPubStructs: CREATING OFFICIAL Struct type: p.StructLookup[%q]=%p, underlying: %+v", typeSpec.Name.Name, typeSpec, underlying)
		p.TypesInfo.Defs[typeSpec.Name] = namedType
		p.StructLookup[typeSpec.Name.Name] = typeSpec
	}

	return decls
}

// func (p *Package) replaceEmptyStructReferences(emptyNamedType *types.TypeName, emptyTypeSpec *ast.TypeSpec, emptyStruct *types.Struct, namedType *types.TypeName, typeSpec *ast.TypeSpec) {
// 	for typeSpecName, nt := range p.TypesInfo.Defs {
// 		if oldStruct, ok := nt.Type().Underlying().(*types.Struct); ok {
// 			if oldStruct != emptyStruct {
// 				continue
// 			}
// 			log.Printf("GML: packages/process-source.go: replaceEmptyStructReferences: typeSpecName: %v, typeSpecName: %p, nt: %p, oldStruct: %p", typeSpecName.Name, typeSpecName, nt, oldStruct)
// 			p.TypesInfo.Defs[typeSpecName] = namedType
// 			p.StructLookup[typeSpecName.Name] = typeSpec
// 		}
// 	}
// }

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
				paramTypeName = strings.Split(paramType.Obj().Name(), " = ")[0]
			default:
				log.Fatalf("PROGRAMMING ERROR: expected *types.Named, got %T", paramType)
			}
			// newName := strings.TrimSuffix(param.Name(), "~")  // still needed
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
		types.NewSignatureType(nil, nil, nil, paramsTuple, resultsTuple, false)) // TODO

	return decls
}

func (p *Package) processParameters(typesPkg *types.Package, allArgs string) (paramsList []*ast.Field, paramsVars []*types.Var) {
	allArgParts := splitParamsWithBrackets(allArgs)
	for _, arg := range allArgParts {
		argParts := strings.Split(arg, ":")
		if len(argParts) != 2 {
			log.Printf("Warning: invalid argument: '%v'; skipping", arg)
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

	// fullyQualifiedReturnSig := returnSig
	// if !strings.Contains(returnSig, ".") {
	// 	typ, _, _ := utils.utils.StripErrorAndOption(returnSig)
	// 	if utils.IsStructType(typ) {
	// 		pkgName := typesPkg.Path()
	// 		baseTypeName := pkgName + "." + typ
	// 		if _, ok := p.StructLookup[baseTypeName]; ok {
	// 			fullyQualifiedReturnSig = pkgName + "." + returnSig
	// 		}
	// 	}
	// }

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
			fullCustomType := types.NewTypeName(0, nil, typeSignature, nil) // do NOT add typesPkg to custom types!
			log.Printf("GML: packages/process-source.go: checkCustomMoonBitType: FOUND FULLY-SPECIFIED Struct DEFINITION for p.StructLookup[%q]=%p: underlying: %+v", typ, typeSpec, underlying)
			return types.NewNamed(fullCustomType, underlying, nil)
		}
		log.Printf("PROGRAMMING ERROR: checkCustomMoonBitType(typeSignature='%v'): typ '%v' missing from p.TypesInfo.Defs", typeSignature, typ)
	}

	if utils.IsStructType(typ) {
		// This could possibly be a forward reference or a recursive reference within a struct.
		// Create an entry for it and point to it as the underlying type.
		// Its fields will be filled in later.
		underlying := types.NewStruct(nil, nil)
		customType := types.NewTypeName(0, nil, typ, underlying) // do NOT add typesPkg to custom types!
		typeSpec := &ast.TypeSpec{Name: &ast.Ident{Name: typ}}
		log.Printf("GML: packages/process-source.go: checkCustomMoonBitType: CREATING FORWARD REFERENCE for Struct type: p.StructLookup[%q]=%p", typ, typeSpec)
		p.StructLookup[typ] = typeSpec
		p.TypesInfo.Defs[typeSpec.Name] = customType
		if typ == typeSignature {
			return types.NewNamed(customType, underlying, nil)
		}
		log.Printf("GML: packages/process-source.go: checkCustomMoonBitType: CREATING FORWARD REFERENCE for Struct type: p.StructLookup[%q]=%p", typeSignature, typeSpec)
		p.StructLookup[typeSignature] = typeSpec
		p.TypesInfo.Defs[typeSpec.Name] = customType
		fullCustomType := types.NewTypeName(0, nil, typeSignature, underlying)
		return types.NewNamed(fullCustomType, underlying, nil)
	}

	return resultType
}

func (p *Package) getMoonBitNamedType(typesPkg *types.Package, typeSignature string) (resultType types.Type) { // (resultType *types.Named) {
	log.Printf("GML: getMoonBitNamedType(typeSignature=%q)", typeSignature)
	// TODO: write a parser that can handle any MoonBit type.
	// if utils.IsListType(typeSignature) {
	// 	innerTypeSignature := utils.GetListSubtype(typeSignature)
	// 	resultType = p.checkCustomMoonBitType(typesPkg, innerTypeSignature)
	// 	if resultType != nil {
	// 		arrayTypeName := typeSignature[:strings.Index(typeSignature, "[")]
	// 		tmpSignature := fmt.Sprintf("%v[%v.%v]", arrayTypeName, typesPkg.Path(), innerTypeSignature)
	// 		return types.NewNamed(types.NewTypeName(0, nil, tmpSignature, nil), nil, nil)
	// 	}
	// }

	// if utils.IsMapType(typeSignature) {
	// 	ktSig, vtSig := utils.GetMapSubtypes(typeSignature)
	// 	// TODO: For now, assume that the Map key is a primitive type.
	// 	// kt := p.getMoonBitNamedType(typesPkg, ktSig)
	// 	vt := p.checkCustomMoonBitType(typesPkg, vtSig)
	// 	if vt != nil {
	// 		mapSignature := fmt.Sprintf("Map[%v, %v.%v]", ktSig, typesPkg.Path(), vtSig)
	// 		return types.NewNamed(types.NewTypeName(0, nil, mapSignature, nil), nil, nil)
	// 	}
	// }

	// Treat a MoonBit tuple like a struct whose field names are "0", "1", etc.
	if strings.HasPrefix(typeSignature, "(") && strings.HasSuffix(typeSignature, ")") {
		allArgs := typeSignature[1 : len(typeSignature)-1]
		allTupleParts := splitParamsWithBrackets(allArgs)
		// var fields []*ast.Field
		var fieldVars []*types.Var
		for i, field := range allTupleParts {
			fieldType := strings.TrimSpace(commentRE.ReplaceAllString(field, ""))
			if fieldType == "" {
				continue
			}
			fieldName := fmt.Sprintf("%v", i)
			// fields = append(fields, &ast.Field{
			// 	Names: []*ast.Ident{{Name: fieldName}},
			// 	Type:  &ast.Ident{Name: fieldType},
			// })
			underlying := p.getMoonBitNamedType(typesPkg, fieldType)
			fieldVars = append(fieldVars, types.NewVar(0, nil, fieldName, underlying)) // &moonType{typeName: fieldType}))
		}

		tupleStruct := types.NewStruct(fieldVars, nil)
		return types.NewNamed(types.NewTypeName(0, nil, typeSignature, tupleStruct), nil, nil) // &moonType{typeName: typeSignature}
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
