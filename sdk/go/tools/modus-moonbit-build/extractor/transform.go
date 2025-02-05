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
	if s == nil {
		log.Printf("GML: transform.go: transformStruct(name=%q, s=nil)", name)
		return nil
	}

	structDecl, structType := getStructDeclarationAndType(name, pkgs)
	if structDecl == nil || structType == nil {
		log.Printf("GML: transform.go: transformStruct(name=%q): structDecl=%v, structType=%v", name, structDecl, structType)
		return nil
	}

	structDocs := getDocs(structDecl.Doc)

	fields := make([]*metadata.Field, s.NumFields())
	for i := 0; i < s.NumFields(); i++ {
		f := s.Field(i)

		fieldDocs := getDocs(structType.Fields.List[i].Doc)

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

func transformFunc(name string, f *types.Func, pkgs map[string]*packages.Package) *metadata.Function {
	if f == nil {
		log.Printf("GML: transform.go: transformFunc(name=%q, f=nil)", name)
		return nil
	}

	sig := f.Type().(*types.Signature)
	params := sig.Params()
	results := sig.Results()

	funcDecl := getFuncDeclaration(f, pkgs)
	if funcDecl == nil {
		log.Printf("GML: transform.go: transformFunc(name=%q, f=%#v): funcDecl=nil", name, f)
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
			param := &metadata.Parameter{
				Name: p.Name(),
				Type: p.Type().String(),
				// TODO: Add default value here if it is a constant literal
				// Default:
			}
			log.Printf("GML: transform.go: p.Type()=%#v\n", p.Type())
			ret.Parameters[i] = param
		}
	}

	if results != nil {
		ret.Results = make([]*metadata.Result, results.Len())
		for i := 0; i < results.Len(); i++ {
			r := results.At(i)
			ret.Results[i] = &metadata.Result{
				Name: r.Name(),
				Type: r.Type().String(),
			}
		}
	}

	return &ret
}

func getStructDeclarationAndType(name string, pkgs map[string]*packages.Package) (*ast.GenDecl, *ast.StructType) {
	objName := name[strings.LastIndex(name, ".")+1:]
	pkgName := utils.GetPackageNamesForType(name)[0]
	pkg := pkgs[pkgName]

	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if typeSpec.Name.Name == objName || typeSpec.Name.Name == name {
							if structType, ok := typeSpec.Type.(*ast.StructType); ok {
								log.Printf("GML: transform.go: getStructDeclarationAndType(name=%q): A", name)
								return genDecl, structType
							} else if ident, ok := typeSpec.Type.(*ast.Ident); ok {
								typePath := pkgName + "." + ident.Name
								log.Printf("GML: transform.go: getStructDeclarationAndType(name=%q): B: typePath=%q", name, typePath)
								return getStructDeclarationAndType(typePath, pkgs)
							} else if selExp, ok := typeSpec.Type.(*ast.SelectorExpr); ok {
								if pkgIdent, ok := selExp.X.(*ast.Ident); !ok {
									log.Printf("GML: transform.go: getStructDeclarationAndType(name=%q): C", name)
									return nil, nil
								} else {
									pkgPath := getFullImportPath(file, pkgIdent.Name)
									typePath := pkgPath + "." + selExp.Sel.Name
									log.Printf("GML: transform.go: getStructDeclarationAndType(name=%q): D: pkgPath=%q, typePath=%q", name, pkgPath, typePath)
									return getStructDeclarationAndType(typePath, pkgs)
								}
							}
						}
					}
				}
			}
		}
	}

	return nil, nil
}

func getFullImportPath(file *ast.File, pkgName string) string {
	for _, imp := range file.Imports {
		path := strings.Trim(imp.Path.Value, `"`)
		if imp.Name == nil {
			parts := strings.Split(path, "/")
			if len(parts) == 0 { // huh?!? is this possible?!?
				log.Printf("GML: transform.go: getFullImportPath(pkgName=%q): A: path=%q", pkgName, path)
				return ""
			}
			if parts[len(parts)-1] == pkgName {
				log.Printf("GML: transform.go: getFullImportPath(pkgName=%q): B: path=%q", pkgName, path)
				return path
			}
		} else if imp.Name.Name == pkgName {
			log.Printf("GML: transform.go: getFullImportPath(pkgName=%q): C: path=%q", pkgName, path)
			return path
		}
	}
	log.Printf("GML: transform.go: getFullImportPath(pkgName=%q): D: ''", pkgName)
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
