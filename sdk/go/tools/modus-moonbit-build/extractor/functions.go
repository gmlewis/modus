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
	"fmt"
	"go/ast"
	"go/types"
	"log"
	"strings"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/packages"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
)

var wellKnownTypes = map[string]bool{
	"Bytes":       true, // 1 - TODO: How should these wellKnownTypes be handled?
	"Array[Byte]": true, // 2
	"String":      true, // 3
	// "time.Time":     true,
	// "time.Duration": true,
}

func getFuncDeclaration(fn *types.Func, pkgs map[string]*packages.Package) *ast.FuncDecl {
	fnName := strings.TrimPrefix(fn.Name(), "__modus_")
	pkg := pkgs[fn.Pkg().Path()]

	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			if fd, ok := decl.(*ast.FuncDecl); ok {
				if fd.Name.Name == fnName {
					return fd
				}
			}
		}
	}

	return nil
}

type funcWithPkg struct {
	fn  *types.Func
	pkg *packages.Package
}

func getExportedFunctions(pkgs map[string]*packages.Package) map[string]*funcWithPkg {
	results := make(map[string]*funcWithPkg)
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				if fd, ok := decl.(*ast.FuncDecl); ok {
					if name := getImportedFuncName(fd); name != "" {
						continue
					}
					if name := getExportedFuncName(pkg, fd); name != "" {
						if f, ok := pkg.TypesInfo.Defs[fd.Name].(*types.Func); ok {
							results[name] = &funcWithPkg{fn: f, pkg: pkg}
						}
					}
				}
			}
		}
	}
	return results
}

func getProxyImportFunctions(pkgs map[string]*packages.Package) map[string]*funcWithPkg {
	results := make(map[string]*funcWithPkg)
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				if fd, ok := decl.(*ast.FuncDecl); ok {
					if name := getProxyImportFuncName(fd); name != "" {
						if f, ok := pkg.TypesInfo.Defs[fd.Name].(*types.Func); ok {
							results[name] = &funcWithPkg{fn: f, pkg: pkg}
						}
					}
				}
			}
		}
	}
	return results
}

func getImportedFunctions(pkgs map[string]*packages.Package) map[string]*funcWithPkg {
	results := make(map[string]*funcWithPkg)
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				if fd, ok := decl.(*ast.FuncDecl); ok {
					if name := getImportedFuncName(fd); name != "" {
						if f, ok := pkg.TypesInfo.Defs[fd.Name].(*types.Func); ok {
							// we only care about imported modus host functions
							if strings.HasPrefix(name, "modus_") {
								results[name] = &funcWithPkg{fn: f, pkg: pkg}
							}
						}
					}
				}
			}
		}
	}
	return results
}

func getExportedFuncName(pkg *packages.Package, fn *ast.FuncDecl) string {
	if pkg.PkgPath != "" {
		return pkg.PkgPath + "." + fn.Name.Name
	}
	return fn.Name.Name
}

func getImportedFuncName(fn *ast.FuncDecl) string {
	// TODO: Allow the MoonBit compiler to return all information needed here.
	// Imported functions have no body, and are decorated as follows:

	if fn.Body == nil && fn.Doc != nil {
		for _, c := range fn.Doc.List {
			parts := strings.Split(c.Text, " ")
			// TODO
			if len(parts) == 3 && parts[0] == "//go:wasmimport" {
				return parts[1] + "." + parts[2]
			}
		}
	}
	return ""
}

func getProxyImportFuncName(fn *ast.FuncDecl) string {
	/*
		A proxy import is a function wrapper that is decorated as follows:

		//modus:import <module> <function>

		Its definition will be used in lieu of the original function that matches the same wasm module and function name.
	*/

	if fn.Body != nil && fn.Doc != nil {
		for _, c := range fn.Doc.List {
			parts := strings.Split(c.Text, " ")
			if len(parts) == 3 && parts[0] == "//modus:import" {
				return parts[1] + "." + parts[2]
			}
		}
	}
	return ""
}

func findRequiredTypes(fpkg *funcWithPkg, requiredTypes requiredTypesMap) {
	f := fpkg.fn
	pkg := fpkg.pkg
	sig := f.Type().(*types.Signature)

	if params := sig.Params(); params != nil {
		for i := 0; i < params.Len(); i++ {
			t := params.At(i).Type()
			addRequiredTypes(t, requiredTypes, pkg)
		}
	}

	if results := sig.Results(); results != nil {
		for i := 0; i < results.Len(); i++ {
			t := results.At(i).Type()
			addRequiredTypes(t, requiredTypes, pkg)
		}
	}
}

func addRequiredTypes(t types.Type, requiredTypes requiredTypesMap, pkg *packages.Package) bool {
	name := strings.Split(t.String(), " = ")[0] // strip default values

	// prevent infinite recursion
	if _, ok := requiredTypes[name]; ok {
		return true
	}

	// skip Bytes, Arary[Byte], and String, because they're hardcoded as type id 1, 2, and 3
	switch name {
	case "Bytes", "Array[Byte]", "String":
		return true
	}

	switch t := t.(type) {
	case *types.Basic:
		// don't add basic types, but allow other objects to use them
		return true
	case *types.Named:
		// required types are required to be exported, so that we can use them in generated code
		// if !t.Obj().Exported() {
		// 	fmt.Fprintf(os.Stderr, "ERROR: Required type '%v' is not exported. Please export it by adding the `pub` keyword and try again.\n", name)
		// 	os.Exit(1)
		// }

		typ, hasError, hasOption := utils.StripErrorAndOption(name)
		if hasError {
			fullName := fmt.Sprintf("%v!Error", typ)
			tmpType := types.NewNamed(types.NewTypeName(0, nil, fullName, nil), t.Underlying(), nil) // do NOT add typesPkg to named types.
			// Do not recurse here as it would cause an infinite loop.
			requiredTypes[fullName] = &typeWithPkgT{t: tmpType, pkg: pkg}
			// Handle MoonBit error types as a tuple: `(String)`
			fullName = "(String)"
			if _, ok := requiredTypes[fullName]; !ok {
				underlying := types.NewNamed(types.NewTypeName(0, nil, "String", nil), nil, nil)
				fieldVars := []*types.Var{types.NewVar(0, nil, "0", underlying)}
				tupleStruct := types.NewStruct(fieldVars, nil)
				tmpType = types.NewNamed(types.NewTypeName(0, nil, fullName, nil), tupleStruct, nil)
				requiredTypes[fullName] = &typeWithPkgT{t: tmpType, pkg: pkg}
			}
		}

		// Since the `modus_pre_generated.mbt` file handles all errors, strip error types here.
		if i := strings.Index(name, "!"); i >= 0 {
			name = name[:i]
		}

		u := t.Underlying()
		if u == nil {
			u = resolveMissingUnderlyingType(name, t, requiredTypes, pkg)
		}
		requiredTypes[name] = &typeWithPkgT{t: u, pkg: pkg}
		// Make sure that the underlying type is also added to the required types.
		if hasOption {
			requiredTypes[typ] = &typeWithPkgT{t: u, pkg: pkg}
		}

		// Because the MoonBit SDK is currently using *types.Named for _ALL_ types, more processing needs to happen here.
		if strings.HasPrefix(name, "Map[") {
			keyType, valueType := GetMapSubtypes(name)
			// Force the planner to make a plan for slices of the keys and values of the map.
			keyName := fmt.Sprintf("Array[%v]", keyType)
			requiredTypes[keyName] = &typeWithPkgT{pkg: pkg}
			valueName := fmt.Sprintf("Array[%v]", valueType)
			requiredTypes[valueName] = &typeWithPkgT{pkg: pkg}
			_, _, hasOption = utils.StripErrorAndOption(valueType)
			if hasOption {
				requiredTypes[valueName] = &typeWithPkgT{pkg: pkg}
			}
		}

		// skip fields for some well-known types
		if wellKnownTypes[name] {
			return true
		}

		if s, ok := u.(*types.Struct); ok {
			for i := 0; i < s.NumFields(); i++ {
				addRequiredTypes(s.Field(i).Type(), requiredTypes, pkg)
			}
		}

		return true

	case *types.Pointer:
		if addRequiredTypes(t.Elem(), requiredTypes, pkg) {
			requiredTypes[name] = &typeWithPkgT{t: t, pkg: pkg}
			return true
		}
	case *types.Struct:
		// TODO: handle unnamed structs
	case *types.Slice:
		if addRequiredTypes(t.Elem(), requiredTypes, pkg) {
			requiredTypes[name] = &typeWithPkgT{t: t, pkg: pkg}
			return true
		}
	case *types.Array:
		if addRequiredTypes(t.Elem(), requiredTypes, pkg) {
			requiredTypes[name] = &typeWithPkgT{t: t, pkg: pkg}
			return true
		}
	case *types.Map:
		if addRequiredTypes(t.Key(), requiredTypes, pkg) {
			if addRequiredTypes(t.Elem(), requiredTypes, pkg) {
				if addRequiredTypes(types.NewSlice(t.Key()), requiredTypes, pkg) {
					if addRequiredTypes(types.NewSlice(t.Elem()), requiredTypes, pkg) {
						requiredTypes[name] = &typeWithPkgT{t: t, pkg: pkg}
						return true
					}
				}
			}
		}
	}

	return false
}

func GetMapSubtypes(typ string) (string, string) {
	typ, _, _ = utils.StripErrorAndOption(typ)

	if !strings.HasSuffix(typ, "]") && !strings.HasSuffix(typ, "]?") {
		return "", ""
	}

	const prefix = "Map[" // e.g. Map[String, Int]
	if !strings.HasPrefix(typ, prefix) {
		return "", ""
	}
	typ = strings.TrimSuffix(typ, "?")
	typ = strings.TrimSuffix(typ, "]")
	typ = strings.TrimPrefix(typ, prefix)

	n := 1
	for i := 0; i < len(typ); i++ {
		switch typ[i] {
		case '[':
			n++
		case ']':
			n--
		case ',':
			if n == 1 {
				return strings.TrimSpace(typ[:i]), strings.TrimSpace(typ[i+1:])
			}
		}
	}

	return "", ""
}

// Due to simulating the MoonBit source code AST as Go AST nodes without fully parsing all source files
// that the MoonBit compiler reads, it is currently necessary to provide workarounds as bugs are found.
// This is one such workaround.
func resolveMissingUnderlyingType(name string, t *types.Named, requiredTypes requiredTypesMap, pkg *packages.Package) types.Type {
	// Hack to make a tuple appear to have an underlying struct type for the metadata:
	if s, ok := t.Obj().Type().(*types.Struct); ok {
		return s
	}

	if typeSpec, ok := pkg.StructLookup[name]; ok {
		if customType, ok := pkg.TypesInfo.Defs[typeSpec.Name].(*types.TypeName); ok {
			u := customType.Type().Underlying()
			return u
		}
		log.Fatalf("PROGRAMMING ERROR: extractor/functions.go: resolveMissingUnderlyingType: *types.Named: pkg.TypesInfo.Defs[%q]=%T", typeSpec.Name.Name, pkg.TypesInfo.Defs[typeSpec.Name])
		return nil
	}

	if strings.HasPrefix(name, "Array[") {
		typ := utils.StripDefaultValue(name)
		typ, _, _ = utils.StripErrorAndOption(typ)
		// Add the underlying struct, but don't make it the underlying type for the array.
		typ = strings.TrimSuffix(strings.TrimPrefix(typ, "Array["), "]")
		// This is an ugly workaround - find a better solution
		pkg.AddPossiblyMissingUnderlyingType(typ)
		if _, ok := requiredTypes[typ]; !ok {
			requiredTypes[typ] = lookupStruct(typ, requiredTypes, pkg)
		}
		return nil
	}

	if strings.HasPrefix(name, "FixedArray[") {
		typ := utils.StripDefaultValue(name)
		typ, _, _ = utils.StripErrorAndOption(typ)
		// Add the underlying struct, but don't make it the underlying type for the fixedarray.
		typ = strings.TrimSuffix(strings.TrimPrefix(typ, "FixedArray["), "]")
		// This is an ugly workaround - find a better solution
		pkg.AddPossiblyMissingUnderlyingType(typ)
		if _, ok := requiredTypes[typ]; !ok {
			requiredTypes[typ] = lookupStruct(typ, requiredTypes, pkg)
		}
		return nil
	}

	return nil
}

func lookupStruct(name string, requiredTypes requiredTypesMap, pkg *packages.Package) *typeWithPkgT {
	if u, ok := requiredTypes[name]; ok {
		return u
	}
	if typeSpec, ok := pkg.StructLookup[name]; ok {
		if customType, ok := pkg.TypesInfo.Defs[typeSpec.Name].(*types.TypeName); ok {
			u := customType.Type().Underlying()
			return &typeWithPkgT{t: u, pkg: pkg}
		}
	}
	return &typeWithPkgT{t: nil, pkg: pkg}
}
