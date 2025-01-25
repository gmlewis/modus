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
)

var wellKnownTypes = map[string]bool{
	"Bytes":       true, // 1
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

func getExportedFunctions(pkgs map[string]*packages.Package) map[string]*types.Func {
	results := make(map[string]*types.Func)
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				if fd, ok := decl.(*ast.FuncDecl); ok {
					if name := getExportedFuncName(fd); name != "" {
						if f, ok := pkg.TypesInfo.Defs[fd.Name].(*types.Func); ok {
							results[name] = f
						}
					}
				}
			}
		}
	}
	return results
}

func getProxyImportFunctions(pkgs map[string]*packages.Package) map[string]*types.Func {
	results := make(map[string]*types.Func)
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				if fd, ok := decl.(*ast.FuncDecl); ok {
					if name := getProxyImportFuncName(fd); name != "" {
						if f, ok := pkg.TypesInfo.Defs[fd.Name].(*types.Func); ok {
							results[name] = f
						}
					}
				}
			}
		}
	}
	return results
}

func getImportedFunctions(pkgs map[string]*packages.Package) map[string]*types.Func {
	results := make(map[string]*types.Func)
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				if fd, ok := decl.(*ast.FuncDecl); ok {
					if name := getImportedFuncName(fd); name != "" {
						if f, ok := pkg.TypesInfo.Defs[fd.Name].(*types.Func); ok {
							// we only care about imported modus host functions
							if strings.HasPrefix(name, "modus_") {
								results[name] = f
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
	// TODO
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

func findRequiredTypes(f *types.Func, m map[string]types.Type) {
	sig := f.Type().(*types.Signature)

	if params := sig.Params(); params != nil {
		for i := 0; i < params.Len(); i++ {
			t := params.At(i).Type()
			addRequiredTypes(t, m)
		}
	}

	if results := sig.Results(); results != nil {
		for i := 0; i < results.Len(); i++ {
			t := results.At(i).Type()
			addRequiredTypes(t, m)
		}
	}
}

func addRequiredTypes(t types.Type, m map[string]types.Type) bool {
	name := t.String()

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
		// 	fmt.Fprintf(os.Stderr, "ERROR: Required type %s is not exported. Rename it to start with a capital letter and try again.\n", name)
		// 	os.Exit(1)
		// }

		u := t.Underlying()
		m[name] = u
		log.Printf("GML: extractor/functions.go: addRequiredTypes: *types.Named: m[%q]=%T", name, u)

		// Because the MoonBit SDK is currently using *types.Named for _ALL_ types, more processing needs to happen here.
		if strings.HasPrefix(name, "Map[") && strings.HasSuffix(name, "]") {
			parts := strings.Split(name[4:len(name)-1], ",")
			// Force the planner to make a plan for slices of the keys and values of the map.
			keyType := strings.TrimSpace(parts[0])
			keyName := fmt.Sprintf("Array[%v]", keyType)
			m[keyName] = nil
			log.Printf("GML: extractor/functions.go: addRequiredTypes: *types.Named: m[%q]=nil", keyName)
			valueType := strings.TrimSpace(parts[1])
			valueName := fmt.Sprintf("Array[%v]", valueType)
			m[valueName] = nil
			log.Printf("GML: extractor/functions.go: addRequiredTypes: *types.Named: m[%q]=nil", valueName)
		}

		// skip fields for some well-known types
		if wellKnownTypes[name] {
			return true
		}

		if s, ok := u.(*types.Struct); ok {
			for i := 0; i < s.NumFields(); i++ {
				addRequiredTypes(s.Field(i).Type(), m)
			}
		}

		return true

	case *types.Pointer:
		if addRequiredTypes(t.Elem(), m) {
			m[name] = t
			return true
		}
	case *types.Struct:
		// TODO: handle unnamed structs
	case *types.Slice:
		if addRequiredTypes(t.Elem(), m) {
			m[name] = t
			return true
		}
	case *types.Array:
		if addRequiredTypes(t.Elem(), m) {
			m[name] = t
			return true
		}
	case *types.Map:
		log.Printf("GML: extractor/functions.go: addRequiredTypes: *types.Map: A")
		if addRequiredTypes(t.Key(), m) {
			log.Printf("GML: extractor/functions.go: addRequiredTypes: *types.Map: B")
			if addRequiredTypes(t.Elem(), m) {
				log.Printf("GML: extractor/functions.go: addRequiredTypes: *types.Map: C")
				if addRequiredTypes(types.NewSlice(t.Key()), m) {
					log.Printf("GML: extractor/functions.go: addRequiredTypes: *types.Map: D")
					if addRequiredTypes(types.NewSlice(t.Elem()), m) {
						log.Printf("GML: extractor/functions.go: addRequiredTypes: *types.Map: E: m[%q]=%T", name, t)
						m[name] = t
						return true
					}
				}
			}
		}
	}

	return false
}
