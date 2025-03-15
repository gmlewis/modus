/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package extractor

import (
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"strings"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/packages"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
)

func transformStruct(name string, s *types.Struct, pkgs map[string]*packages.Package) *metadata.TypeDefinition {
	structDecl, structType, typesStruct := getStructDeclarationAndType(name, pkgs)
	if typesStruct != nil {
		s = typesStruct // yeah, confusing variable names. This is the *types.Struct to get the fields.
	}

	var structDocs *metadata.Docs
	if structDecl != nil && structDecl.Doc != nil {
		structDocs = getDocs(structDecl.Doc)
	}

	var fields []*metadata.Field
	if s.NumFields() > 0 {
		fields = make([]*metadata.Field, s.NumFields())
	}
	for i := 0; i < s.NumFields(); i++ {
		f := s.Field(i)

		var fieldDocs *metadata.Docs
		if structType != nil && structType.Fields != nil && structType.Fields.List[i] != nil {
			fieldDocs = getDocs(structType.Fields.List[i].Doc)
		}

		fields[i] = &metadata.Field{
			Name: utils.CamelCase(f.Name()),
			Type: f.Type().String(),
			Docs: fieldDocs,
		}
	}

	return &metadata.TypeDefinition{
		Name:   name,
		Fields: fields,
		Docs:   structDocs,
	}
}

func transformFunc(name string, fpkg *funcWithPkg, pkgs map[string]*packages.Package, requiredTypes requiredTypesMap) *metadata.Function {
	f := fpkg.fn
	if f == nil {
		gmlPrintf("GML: extractor/transform.go: transformFunc(name=%q, f=nil)", name)
		return nil
	}

	sig := f.Type().(*types.Signature)
	params := sig.Params()
	results := sig.Results()

	funcDecl := getFuncDeclaration(f, pkgs)
	if funcDecl == nil {
		gmlPrintf("GML: extractor/transform.go: transformFunc(name=%q): sig: '%v', funcDecl=nil", sig.String(), name)
		return nil
	}

	ret := metadata.Function{
		Name: name,
		Docs: getDocs(funcDecl.Doc),
	}

	if params != nil {
		ret.Parameters = make([]*metadata.Parameter, params.Len())
		for i := 0; i < params.Len(); i++ {
			p := params.At(i)
			paramType, acc := utils.FullyQualifyTypeName(fpkg.pkg.PkgPath, p.Type().String())
			paramTypeParts := strings.Split(paramType, " = ")
			paramType = paramTypeParts[0]
			param := &metadata.Parameter{
				Name: p.Name(),
				Type: paramType,
			}
			if len(paramTypeParts) > 1 {
				if defaultValue, ok := utils.GetDefaultValue(paramType, paramTypeParts[1]); ok {
					param.Default = &defaultValue
				}
			}
			gmlPrintf("GML: extractor/transform.go: transformFunc(name=%q): sig: '%v', param[%v]: {Name: %q, Type: %q}", name, sig.String(), i, param.Name, param.Type)
			ret.Parameters[i] = param
			// Now make empty types for all underlying types needed by this parameter.
			delete(acc, paramType)
			for k := range acc {
				if _, ok := requiredTypes[k]; !ok {
					gmlPrintf("GML: extractor/transform.go: A: paramType ADDING m['%v']=nil", k)
					requiredTypes[k] = &typeWithPkgT{pkg: fpkg.pkg}
				}
			}
		}
	}

	if results != nil {
		ret.Results = make([]*metadata.Result, results.Len())
		for i := 0; i < results.Len(); i++ {
			r := results.At(i)
			resultType, acc := utils.FullyQualifyTypeName(fpkg.pkg.PkgPath, r.Type().String())
			result := &metadata.Result{
				Name: r.Name(),
				Type: resultType,
			}
			gmlPrintf("GML: extractor/transform.go: transformFunc(name=%q): sig: '%v', result[%v]: {Name: %q, Type: %q}", name, sig.String(), i, result.Name, result.Type)
			ret.Results[i] = result
			// Now make empty types for all underlying types needed by this parameter.
			delete(acc, resultType)
			for k := range acc {
				if _, ok := requiredTypes[k]; !ok {
					gmlPrintf("GML: extractor/transform.go: B: resultType ADDING m['%v']=nil", k)
					requiredTypes[k] = &typeWithPkgT{pkg: fpkg.pkg}
				}
			}
		}
	}

	return &ret
}

func getStructDeclarationAndType(name string, pkgs map[string]*packages.Package) (*ast.GenDecl, *ast.StructType, *types.Struct) {
	gmlPrintf("GML: extractor/transform.go: ENTER getStructDeclarationAndType(name=%q)", name)
	if utils.IsTupleType(name) {
		return nil, nil, nil
	}

	name, _, _ = utils.StripErrorAndOption(name)
	objName := name[strings.LastIndex(name, ".")+1:]
	pkgNames := utils.GetPackageNamesForType(name)
	if len(pkgNames) == 0 {
		// This struct does not have a package name, therefore it should be in the main package. Find it.
		typ, _, _ := utils.StripErrorAndOption(objName)
		for pkgName, pkg := range pkgs {
			gmlPrintf("GML: extractor/transform.go: getStructDeclarationAndType(name=%q): looking at pkgs[%q]", name, pkgName)
			if typeSpec, ok := pkg.StructLookup[typ]; ok {
				if structType, ok := typeSpec.Type.(*ast.StructType); ok {
					gmlPrintf("GML: extractor/transform.go: getStructDeclarationAndType(name=%q): FOUND: p.StructLookup[%q]=%p", name, typ, typeSpec)
					customType, ok := pkg.TypesInfo.Defs[typeSpec.Name]
					if !ok {
						log.Fatalf("PROGRAMMING ERROR: transform.go: getStructDeclarationAndType(name=%q): customType not found!", name)
					}
					underlying := customType.Type().Underlying()
					typesStruct, ok := underlying.(*types.Struct)
					if !ok {
						log.Fatalf("PROGRAMMING ERROR: transform.go: getStructDeclarationAndType(name=%q): typesStruct not found!", name)
					}
					genDecl := findGenDeclForTypeSpecName(pkg.Syntax, typeSpec.Name.Name)
					return genDecl, structType, typesStruct
				}
			}
		}
		// The struct is not found.
		log.Fatalf("PROGRAMMING ERROR: transform.go: getStructDeclarationAndType(name=%q): pkg not found!", name)
		return nil, nil, nil
	}

	pkgName := pkgNames[0]
	pkg := pkgs[pkgName]
	if pkg == nil {
		log.Fatalf("PROGRAMMING ERROR: transform.go: getStructDeclarationAndType(name=%q): pkg[%q] is nil", name, pkgName)
		for pkgName := range pkgs {
			gmlPrintf("GML: extractor/transform.go: getStructDeclarationAndType(name=%q): found pkgs[%q]", name, pkgName)
		}
		return nil, nil, nil
	}

	// Yes, the following is still necessary even with the advent of p.StructLookup
	// because the main package may refer to types from other packages
	// and will not have defined the structs locally in their own package.
	// Therefore, search and find them.

	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if typeSpec.Name.Name == objName || typeSpec.Name.Name == name {
							if structType, ok := typeSpec.Type.(*ast.StructType); ok {
								// Make sure we use the complete struct definition with all its fields and not an empty forward reference.
								if fullTypeSpec, ok := pkg.StructLookup[typeSpec.Name.Name]; ok {
									if fullStructType, ok := fullTypeSpec.Type.(*ast.StructType); ok {
										gmlPrintf("GML: extractor/transform.go: FOUND AND USING FULL DEFINITION getStructDeclarationAndType(name=%q): Z: fullStructType=%#v", name, fullStructType)
										customType, ok := pkg.TypesInfo.Defs[fullTypeSpec.Name]
										if !ok {
											log.Fatalf("PROGRAMMING ERROR: transform.go: getStructDeclarationAndType(name=%q): customType not found!", name)
										}
										underlying := customType.Type().Underlying()
										typesStruct, ok := underlying.(*types.Struct)
										if !ok {
											log.Fatalf("PROGRAMMING ERROR: transform.go: getStructDeclarationAndType(name=%q): typesStruct not found!", name)
										}
										genDecl := findGenDeclForTypeSpecName(pkg.Syntax, typeSpec.Name.Name)
										return genDecl, fullStructType, typesStruct
									}
								}
								gmlPrintf("GML: IS THIS NECESSARY?!? extractor/transform.go: getStructDeclarationAndType(name=%q): A", name)
								return genDecl, structType, nil
							} else if ident, ok := typeSpec.Type.(*ast.Ident); ok {
								typePath := pkgName + "." + ident.Name
								gmlPrintf("GML: IS THIS RECURSION NECESSARY?!? extractor/transform.go: getStructDeclarationAndType(name=%q): B: typePath=%q", name, typePath)
								return getStructDeclarationAndType(typePath, pkgs)
							} else if selExp, ok := typeSpec.Type.(*ast.SelectorExpr); ok {
								if pkgIdent, ok := selExp.X.(*ast.Ident); !ok {
									gmlPrintf("GML: IS THIS NECESSARY?!? extractor/transform.go: getStructDeclarationAndType(name=%q): C", name)
									return nil, nil, nil
								} else {
									pkgPath := getFullImportPath(file, pkgIdent.Name)
									typePath := pkgPath + "." + selExp.Sel.Name
									gmlPrintf("GML: IS THIS RECURSION NECESSARY?!? extractor/transform.go: getStructDeclarationAndType(name=%q): D: pkgPath=%q, typePath=%q", name, pkgPath, typePath)
									return getStructDeclarationAndType(typePath, pkgs)
								}
							}
						}
					}
				}
			}
		}
	}

	return nil, nil, nil
}

func findGenDeclForTypeSpecName(files []*ast.File, typeSpecName string) *ast.GenDecl {
	for _, file := range files {
		for _, decl := range file.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					if ts, ok := spec.(*ast.TypeSpec); ok {
						if ts.Name.Name == typeSpecName {
							return genDecl
						}
					}
				}
			}
		}
	}
	return nil
}

func getFullImportPath(file *ast.File, pkgName string) string {
	for _, imp := range file.Imports {
		path := strings.Trim(imp.Path.Value, `"`)
		if imp.Name == nil {
			parts := strings.Split(path, "/")
			// if len(parts) == 0 { // huh?!? is this even possible?!?  No.
			// 	gmlPrintf("GML: extractor/transform.go: getFullImportPath(pkgName=%q): A: path=%q", pkgName, path)
			// 	return ""
			// }
			if parts[len(parts)-1] == pkgName {
				gmlPrintf("GML: extractor/transform.go: getFullImportPath(pkgName=%q): B: path=%q", pkgName, path)
				return path
			}
		} else if imp.Name.Name == pkgName {
			gmlPrintf("GML: extractor/transform.go: getFullImportPath(pkgName=%q): C: path=%q", pkgName, path)
			return path
		}
	}
	gmlPrintf("GML: extractor/transform.go: getFullImportPath(pkgName=%q): D: ''", pkgName)
	return ""
}

func getDocs(comments *ast.CommentGroup) *metadata.Docs {
	if comments == nil {
		return nil
	}

	var lines []string
	for _, comment := range comments.List {
		txt := comment.Text
		if strings.HasPrefix(txt, "// ") {
			txt = strings.TrimPrefix(txt, "// ")
			txt = strings.TrimSpace(txt)
			lines = append(lines, txt)
		} else if strings.HasPrefix(txt, "/*") {
			txt = strings.TrimPrefix(txt, "/*")
			txt = strings.TrimSuffix(txt, "*/")
			txt = strings.TrimSpace(txt)
			lines = append(lines, strings.Split(txt, "\n")...)
		}
	}

	if len(lines) == 0 {
		return nil
	}

	return &metadata.Docs{
		Lines: lines,
	}
}
