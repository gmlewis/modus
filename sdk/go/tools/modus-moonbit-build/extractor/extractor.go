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
	"go/types"
	"log"
	"strings"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/packages"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
)

func CollectProgramInfo(config *config.Config, meta *metadata.Metadata, mod *modinfo.ModuleInfo) error {
	pkgs, err := loadPackages(config.SourceDir, mod)
	if err != nil {
		return err
	}

	return collectProgramInfoFromPkgs(pkgs, meta)
}

type typeWithPkgT struct {
	t   types.Type
	pkg *packages.Package
}

type requiredTypesMap map[string]*typeWithPkgT

func collectProgramInfoFromPkgs(pkgs map[string]*packages.Package, meta *metadata.Metadata) error {
	requiredTypes := requiredTypesMap{}

	for name, f := range getExportedFunctions(pkgs) {
		meta.FnExports[name] = transformFunc(name, f, pkgs, requiredTypes)
		findRequiredTypes(f, requiredTypes)
	}

	// * If a function returns any `@time.*`, the time-related types will be added.
	// * If a function returns any `!Error` or ` raise Error`, the `logMessage` import will be added.
	for _, export := range meta.FnExports {
		returnType := moonBitReturnType(export)
		if strings.Contains(returnType, "@time.") {
			// TODO: See if this can be added by the parsing phase.
			requiredTypes["Array[Byte]"] = &typeWithPkgT{
				t:   types.NewNamed(types.NewTypeName(0, nil, "Array[Byte]", nil), nil, nil),
				pkg: pkgs["@time"],
			}
		}
		if strings.Contains(returnType, "!") {
			meta.FnImports["modus_system.logMessage"] = moonBitFnImports["modus_system.logMessage"]
		}
	}

	for name, f := range getImportedFunctions(pkgs) {
		meta.FnImports[name] = transformFunc(name, f, pkgs, requiredTypes)
		findRequiredTypes(f, requiredTypes)
	}

	// proxy imports overwrite regular imports
	for name, f := range getProxyImportFunctions(pkgs) {
		if _, ok := meta.FnImports[name]; ok {
			meta.FnImports[name] = transformFunc(name, f, pkgs, requiredTypes)
			findRequiredTypes(f, requiredTypes)
		}
	}

	// Now that all the packages have been processed, see if any underlying
	// types have not been added to the `requiredTypes`.
	for _, pkg := range pkgs {
		for typ := range pkg.PossiblyMissingUnderlyingTypes {
			_, acc := utils.FullyQualifyTypeName(pkg.PkgPath, typ)
			for k := range acc {
				if _, ok := requiredTypes[k]; !ok {
					requiredTypes[k] = &typeWithPkgT{
						t:   pkg.GetMoonBitNamedType(k),
						pkg: pkg,
					}
				}
			}
		}
	}

	id := uint32(4) // 1-3 are reserved for Bytes, Array[Byte], and String - WHY?!? Where is this used?
	for name, typeWithPkg := range requiredTypes {
		resolvedName := name
		if typeWithPkg == nil {
			log.Fatalf("PROGRAMMING ERROR: extractor.go requiredTypes['%v'] = nil", name)
		}
		t := typeWithPkg.t

		if name == "(String, String)" || name == "@http.Header" {
			log.Printf("GML: DEBUGGER")
		}

		if t == nil {
			// See if this can be found in pkgs.StructLookup map.
			pkg := typeWithPkg.pkg
			if typeSpec, ok := pkg.StructLookup[name]; ok {
				if customType, ok := pkg.TypesInfo.Defs[typeSpec.Name].(*types.TypeName); ok {
					t = customType.Type()
				} else {
					log.Fatalf("PROGRAMMING ERROR: extractor.go: missing entry for name='%v' in pkg.TypesInfo.Defs", name)
				}
			} else {
				t = pkg.GetMoonBitNamedType(name)
			}
		}

		if n, ok := t.(*types.Named); ok {
			resolvedName = n.String()
			underlying := n.Underlying()
			if underlying != nil {
				t = underlying
			} else {
				// This appears to be a forward reference.
				// See if this can be found in pkg.StructLookup map.
				pkg := typeWithPkg.pkg
				if typeSpec, ok := pkg.StructLookup[resolvedName]; ok {
					if customType, ok := pkg.TypesInfo.Defs[typeSpec.Name].(*types.TypeName); ok {
						t = customType.Type() // .Underlying()
					}
				}
			}
		}

		if s, ok := t.(*types.Struct); ok && !wellKnownTypes[name] {
			t := transformStruct(resolvedName, s, pkgs)
			if t == nil {
				return fmt.Errorf("failed to transform struct %v", resolvedName)
			}
			t.Id = id
			meta.Types[name] = t
		} else {
			meta.Types[name] = &metadata.TypeDefinition{
				Id:   id,
				Name: name,
			}
		}
		id++
	}

	return nil
}
