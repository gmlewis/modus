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
	"os"
	"strings"
	"sync"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/packages"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
)

// TODO: Remove debugging
var gmlDebugEnv bool

func gmlPrintf(fmtStr string, args ...any) {
	sync.OnceFunc(func() {
		log.SetFlags(0)
		if os.Getenv("GML_DEBUG") == "true" {
			gmlDebugEnv = true
		}
	})
	if gmlDebugEnv {
		log.Printf(fmtStr, args...)
	}
}

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
	// * If a function returns any `!Error`, the `logMessage` import will be added.
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
					// gmlPrintf("GML: Adding PossiblyMissingUnderlyingType '%v' to requiredTypes from pkg '%v'", k, pkg.PkgPath)
					// requiredTypes[k] = nil // make an empty entry for it.
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
			gmlPrintf("GML: extractor.go: CollectProgramInfo: A: meta.Types[%q] = %+v\n", name, meta.Types[name])
			// // now that this struct has resolved all its fields, make sure
			// // that the fields are also added to the meta.Types.
			// for _, field := range t.Fields {
			// 	_, acc := utils.FullyQualifyTypeName(t.Name, field.Type)
			// 	for k := range acc {
			// 		if _, ok := meta.Types[k]; !ok {
			// 			id++
			// 			log.Printf("GML: Adding id=%v Name='%v' to metadata.TypeDefinition", id, k)
			// 			meta.Types[k] = &metadata.TypeDefinition{
			// 				Id:   id,
			// 				Name: k,
			// 			}
			// 		}
			// 	}
			// }
		} else {
			meta.Types[name] = &metadata.TypeDefinition{
				Id:   id,
				Name: name,
			}
		}
		id++
	}

	// resolveForwardTypeRefs(meta)

	return nil
}

// func resolveForwardTypeRefs(meta *metadata.Metadata) {
// 	processed := map[string]bool{}
// 	for key, origTyp := range meta.Types {
// 		if processed[key] {
// 			continue
// 		}
// 		baseTypeName, _, hasOption := utils.StripErrorAndOption(key)
// 		if hasOption {
// 			gmlPrintf("GML: extractor.go: resolveForwardTypeRefs: key=%q, origTyp=%#v, baseTypeName=%q, hasOption=%v\n", key, origTyp, baseTypeName, hasOption)
// 			if baseTyp, ok := meta.Types[baseTypeName]; ok {
// 				if len(baseTyp.Fields) == 0 {
// 					baseTyp.Fields = origTyp.Fields
// 					processed[baseTypeName] = true
// 					processed[key] = true
// 					continue
// 				}
// 				if len(origTyp.Fields) == 0 {
// 					origTyp.Fields = baseTyp.Fields
// 					processed[baseTypeName] = true
// 					processed[key] = true
// 				}
// 			}
// 			continue
// 		}
// 		optionTypeName := key + "?"
// 		if optionTyp, ok := meta.Types[optionTypeName]; ok {
// 			if len(optionTyp.Fields) == 0 {
// 				optionTyp.Fields = origTyp.Fields
// 				processed[optionTypeName] = true
// 				processed[key] = true
// 				continue
// 			}
// 			if len(origTyp.Fields) == 0 {
// 				origTyp.Fields = optionTyp.Fields
// 				processed[optionTypeName] = true
// 				processed[key] = true
// 			}
// 		}
// 	}
// }
