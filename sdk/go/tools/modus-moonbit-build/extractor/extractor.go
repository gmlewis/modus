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
	"sort"
	"strings"
	"sync"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
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

func CollectProgramInfo(config *config.Config, meta *metadata.Metadata) error {
	pkgs, err := loadPackages(config.SourceDir)
	if err != nil {
		return err
	}

	return collectProgramInfoFromPkgs(pkgs, meta)
}

func collectProgramInfoFromPkgs(pkgs map[string]*packages.Package, meta *metadata.Metadata) error {
	requiredTypes := make(map[string]types.Type)

	for name, f := range getExportedFunctions(pkgs) {
		meta.FnExports[name] = transformFunc(name, f, pkgs)
		findRequiredTypes(f, requiredTypes)
	}

	// Since MoonBit currently does not give us AST information for imported functions,
	// and our `packages` MoonBit-as-Go-AST simulator does not parse all source files (yet),
	// we currently check all function signatures for the following:
	// * If a function returns any `@time.*`, the time-related imports will be added.
	// * If a function returns any `!Error`, the `logMessage` import will be added.
	// * If the @neo4j package is used, the corresponding modus:import will be added.
	for _, export := range meta.FnExports {
		returnType := moonBitReturnType(export)
		if strings.Contains(returnType, "@time.") {
			meta.FnImports["modus_system.getTimeInZone"] = moonBitFnImports["modus_system.getTimeInZone"]
			meta.FnImports["modus_system.getTimeZoneData"] = moonBitFnImports["modus_system.getTimeZoneData"]
			requiredTypes["Array[Byte]"] = types.NewNamed(types.NewTypeName(0, nil, "Array[Byte]", nil), nil, nil)
		}
		if strings.Contains(returnType, "!") {
			meta.FnImports["modus_system.logMessage"] = moonBitFnImports["modus_system.logMessage"]
		}
	}
	for _, pkg := range pkgs {
		for _, imp := range pkg.MoonPkgJSON.Imports {
			if strings.Contains(string(imp), "pkg/neo4j") {
				name := "modus_neo4j_client.executeQuery"
				meta.FnImports[name] = moonBitFnImports[name]
				requiredTypes["Array[Json]"] = types.NewNamed(types.NewTypeName(0, nil, "Array[Json]", nil), nil, nil)
				requiredTypes["Array[String]"] = types.NewNamed(types.NewTypeName(0, nil, "Array[String]", nil), nil, nil)
				requiredTypes["Json"] = types.NewNamed(types.NewTypeName(0, nil, "Json", nil), nil, nil)

				stringArray := types.NewNamed(types.NewTypeName(0, nil, "Array[String]", nil), nil, nil)
				recordsArray := types.NewNamed(types.NewTypeName(0, nil, "Array[@neo4j.Record?]", nil), nil, nil)
				resultFields := []*types.Var{types.NewVar(0, nil, "keys", stringArray), types.NewVar(0, nil, "records", recordsArray)}
				resultStruct := types.NewStruct(resultFields, nil)
				requiredTypes["@neo4j.EagerResult"] = types.NewNamed(types.NewTypeName(0, nil, "@neo4j.EagerResult", nil), resultStruct, nil)

				recordFields := []*types.Var{types.NewVar(0, nil, "values", stringArray), types.NewVar(0, nil, "keys", stringArray)}
				recordStruct := types.NewStruct(recordFields, nil)
				requiredTypes["@neo4j.Record"] = types.NewNamed(types.NewTypeName(0, nil, "@neo4j.Record", nil), recordStruct, nil)
				requiredTypes["@neo4j.Record?"] = types.NewNamed(types.NewTypeName(0, nil, "@neo4j.Record?", nil), nil, nil)
				requiredTypes["Array[@neo4j.Record?]"] = types.NewNamed(types.NewTypeName(0, nil, "Array[@neo4j.Record?]", nil), nil, nil)

				requiredTypes["@neo4j.EagerResult?"] = types.NewNamed(types.NewTypeName(0, nil, "@neo4j.EagerResult?", nil), nil, nil)
				requiredTypes["@neo4j.EagerResult?!Error"] = types.NewNamed(types.NewTypeName(0, nil, "@neo4j.EagerResult?!Error", nil), nil, nil)
				// requiredTypes["Map[String, Json]"] = types.NewNamed(types.NewTypeName(0, nil, "Map[String, Json]", nil), nil, nil)
			}
		}
	}

	for name, f := range getImportedFunctions(pkgs) {
		meta.FnImports[name] = transformFunc(name, f, pkgs)
		findRequiredTypes(f, requiredTypes)
	}

	// proxy imports overwrite regular imports
	for name, f := range getProxyImportFunctions(pkgs) {
		if _, ok := meta.FnImports[name]; ok {
			meta.FnImports[name] = transformFunc(name, f, pkgs)
			findRequiredTypes(f, requiredTypes)
		}
	}

	id := uint32(4) // 1-3 are reserved for Bytes, Array[Byte], and String
	keys := utils.MapKeys(requiredTypes)
	sort.Strings(keys)
	for _, name := range keys {
		t := requiredTypes[name]

		resolvedName := name
		if n, ok := t.(*types.Named); ok {
			resolvedName = n.String()
			underlying := n.Underlying()
			if underlying != nil {
				t = underlying
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
		} else {
			if name == "Person" || strings.HasPrefix(name, "TimeZoneInfo") {
				gmlPrintf("GML: extractor.go: CollectProgramInfo: B: t: %T, meta.Types[%q] = {ID:%q,Name:%q}\n", t, name, id, name)
			}
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
