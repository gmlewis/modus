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
					if name := getExportedFuncName(fd); name != "" {
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

func getExportedFuncName(fn *ast.FuncDecl) string {
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

func findRequiredTypes(fpkg *funcWithPkg, m map[string]types.Type) {
	f := fpkg.fn
	pkg := fpkg.pkg
	sig := f.Type().(*types.Signature)

	if params := sig.Params(); params != nil {
		for i := 0; i < params.Len(); i++ {
			t := params.At(i).Type()
			addRequiredTypes(t, m, pkg)
		}
	}

	if results := sig.Results(); results != nil {
		for i := 0; i < results.Len(); i++ {
			t := results.At(i).Type()
			addRequiredTypes(t, m, pkg)
		}
	}
}

// During runtime, we sometimes get types that start with "@..".
// Remove them so that the types can be properly resolved.
func hackStripEmptyPackage(typ string) string {
	// if strings.HasPrefix(typ, "@..") { // TODO: Why is this seen during runtime?
	// 	gmlPrintf("GML: extractor/functions.go: STRIPPING '@..' from type=%q", typ)
	// 	return typ[3:]
	// }
	return typ
}

func addRequiredTypes(t types.Type, m map[string]types.Type, pkg *packages.Package) bool {
	name := hackStripEmptyPackage(t.String())

	// prevent infinite recursion
	if _, ok := m[name]; ok {
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
			m[fullName] = tmpType
			// Handle MoonBit error types as a tuple: `(String)`
			fullName = "(String)"
			if _, ok := m[fullName]; !ok {
				underlying := types.NewNamed(types.NewTypeName(0, nil, "String", nil), nil, nil)
				fieldVars := []*types.Var{types.NewVar(0, nil, "0", underlying)}
				tupleStruct := types.NewStruct(fieldVars, nil)
				tmpType = types.NewNamed(types.NewTypeName(0, nil, fullName, nil), tupleStruct, nil)
				m[fullName] = tmpType
			}
		}

		// Since the `modus_pre_generated.mbt` file handles all errors, strip error types here.
		if i := strings.Index(name, "!"); i >= 0 {
			name = name[:i]
		}

		u := t.Underlying()
		if u == nil {
			u = resolveMissingUnderlyingType(name, t, m, pkg)
		}
		m[name] = u
		gmlPrintf("GML: extractor/functions.go: addRequiredTypes: *types.Named: m[%q]=%T", name, u)
		if strings.Contains(name, "FixedArray[Int]") {
			log.Printf("GML: BREAKPOINT")
		}
		// Make sure that the underlying type is also added to the required types.
		if hasOption {
			m[typ] = u
		}
		// if hasError {
		// 	m[name] = u
		// }

		// Because the MoonBit SDK is currently using *types.Named for _ALL_ types, more processing needs to happen here.
		if strings.HasPrefix(name, "Map[") {
			keyType, valueType := GetMapSubtypes(name)
			// t := strings.TrimSuffix(strings.TrimSuffix(name[4:], "?"), "]")
			// parts := strings.Split(t, ",")
			// if len(parts) != 2 {
			// 	gmlPrintf("PROGRAMMING ERROR: GML: extractor/functions.go: addRequiredTypes: *types.Named: m[%q]=%T", name, u)
			// 	return false
			// }
			// Force the planner to make a plan for slices of the keys and values of the map.
			// keyType := strings.TrimSpace(parts[0])
			keyName := fmt.Sprintf("Array[%v]", keyType)
			m[keyName] = nil
			gmlPrintf("GML: extractor/functions.go: addRequiredTypes: *types.Named: m[%q]=nil", keyName)
			// valueType := strings.TrimSpace(parts[1])
			valueName := fmt.Sprintf("Array[%v]", valueType)
			m[valueName] = nil
			gmlPrintf("GML: extractor/functions.go: addRequiredTypes: *types.Named: m[%q]=nil", valueName)
			_, _, hasOption = utils.StripErrorAndOption(valueType)
			if hasOption {
				m[valueName] = nil
				gmlPrintf("GML: extractor/functions.go: addRequiredTypes: *types.Named: m[%q]=nil", valueName)
			}
			// TODO: This is not correct. May have to rethink how the MoonBit source code is parsed.
			// if utils.IsStructType(valueType) {
			// 	if _, ok := m[valueType]; !ok {
			// 		for k := range m {
			// 			if strings.Contains(k, valueType) {
			// 				gmlPrintf("GML: known type: %q", k)
			// 			}
			// 		}
			// 		log.Fatalf("ERROR: Required type '%v' is not exported. Please export it by adding the `pub` keyword and try again.\n", valueType)
			// 	}
			// }
		}

		// skip fields for some well-known types
		if wellKnownTypes[name] {
			return true
		}

		if s, ok := u.(*types.Struct); ok {
			for i := 0; i < s.NumFields(); i++ {
				addRequiredTypes(s.Field(i).Type(), m, pkg)
			}
		}

		return true

	case *types.Pointer:
		if addRequiredTypes(t.Elem(), m, pkg) {
			m[name] = t
			return true
		}
	case *types.Struct:
		// TODO: handle unnamed structs
	case *types.Slice:
		if addRequiredTypes(t.Elem(), m, pkg) {
			m[name] = t
			return true
		}
	case *types.Array:
		if addRequiredTypes(t.Elem(), m, pkg) {
			m[name] = t
			return true
		}
	case *types.Map:
		gmlPrintf("GML: extractor/functions.go: addRequiredTypes: *types.Map: A")
		if addRequiredTypes(t.Key(), m, pkg) {
			gmlPrintf("GML: extractor/functions.go: addRequiredTypes: *types.Map: B")
			if addRequiredTypes(t.Elem(), m, pkg) {
				gmlPrintf("GML: extractor/functions.go: addRequiredTypes: *types.Map: C")
				if addRequiredTypes(types.NewSlice(t.Key()), m, pkg) {
					gmlPrintf("GML: extractor/functions.go: addRequiredTypes: *types.Map: D")
					if addRequiredTypes(types.NewSlice(t.Elem()), m, pkg) {
						gmlPrintf("GML: extractor/functions.go: addRequiredTypes: *types.Map: E: m[%q]=%T", name, t)
						m[name] = t
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
		gmlPrintf("ERROR: functions.go: GetMapSubtypes('%v'): Bad map type!", typ)
		return "", ""
	}

	const prefix = "Map[" // e.g. Map[String, Int]
	if !strings.HasPrefix(typ, prefix) {
		gmlPrintf("GML: extractor/functions.go: A: GetMapSubtypes('%v') = ('', '')", typ)
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
				r1, r2 := strings.TrimSpace(typ[:i]), strings.TrimSpace(typ[i+1:])
				gmlPrintf("GML: extractor/functions.go: B: GetMapSubtypes('%v') = ('%v', '%v')", typ, r1, r2)
				return r1, r2
			}
		}
	}

	gmlPrintf("GML: extractor/functions.go: C: GetMapSubtypes('%v') = ('', '')", typ)
	return "", ""
}

// Due to simulating the MoonBit source code AST as Go AST nodes without fully parsing all source files
// that the MoonBit compiler reads, it is currently necessary to provide workarounds as bugs are found.
// This is one such workaround.
func resolveMissingUnderlyingType(name string, t *types.Named, m map[string]types.Type, pkg *packages.Package) types.Type {
	// Hack to make a tuple appear to have an underlying struct type for the metadata:
	gmlPrintf("GML: extractor/functions.go: addRequiredTypes: t.Obj().Type: %T=%+v", t.Obj().Type(), t.Obj().Type())
	if s, ok := t.Obj().Type().(*types.Struct); ok {
		return s
	}

	gmlPrintf("GML: extractor/functions: addRequiredTypes: p.StructLookup[%q]=%p", name, pkg.StructLookup[name])
	if typeSpec, ok := pkg.StructLookup[name]; ok {
		if customType, ok := pkg.TypesInfo.Defs[typeSpec.Name].(*types.TypeName); ok {
			u := customType.Type().Underlying()
			gmlPrintf("GML: extractor/functions: addRequiredTypes: typeSpec=%p, u=%p=%+v", typeSpec, u, u)
			return u
		}
		gmlPrintf("PROGRAMMING ERROR: extractor/functions.go: addRequiredTypes: *types.Named: pkg.TypesInfo.Defs[%q]=%T", typeSpec.Name.Name, pkg.TypesInfo.Defs[typeSpec.Name])
		return nil
	}

	if strings.HasPrefix(name, "Array[") {
		// Add the underlying struct, but don't make it the underlying type for the array.
		name = strings.TrimSuffix(strings.TrimPrefix(name, "Array["), "]")
		if u := lookupStruct(name, m, pkg); u != nil {
			m[name] = u
		}
		return nil
	}

	if strings.HasPrefix(name, "FixedArray[") {
		// Add the underlying struct, but don't make it the underlying type for the fixedarray.
		name = strings.TrimSuffix(strings.TrimPrefix(name, "FixedArray["), "]")
		if u := lookupStruct(name, m, pkg); u != nil {
			m[name] = u
		}
		return nil
	}

	return nil
}

func lookupStruct(name string, m map[string]types.Type, pkg *packages.Package) types.Type {
	if u, ok := m[name]; ok {
		return u
	}
	if typeSpec, ok := pkg.StructLookup[name]; ok {
		if customType, ok := pkg.TypesInfo.Defs[typeSpec.Name].(*types.TypeName); ok {
			u := customType.Type().Underlying()
			return u
		}
	}
	return nil
}
