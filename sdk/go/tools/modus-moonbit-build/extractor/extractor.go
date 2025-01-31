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
	"go/types"
	"log"
	"sort"
	"strings"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/packages"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
)

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
	// we currently check all function signatures for two things:
	// If a function returns any `@time.*`, the time-related imports will be added.
	// If a function returns any `!Error`, the `logMessage` import will be added.
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
		if s, ok := t.(*types.Struct); ok && !wellKnownTypes[name] {
			t := transformStruct(name, s, pkgs)
			t.Id = id
			meta.Types[name] = t
		} else {
			meta.Types[name] = &metadata.TypeDefinition{
				Id:   id,
				Name: name,
			}
			log.Printf("GML: extractor.go: CollectProgramInfo: meta.Types[%q] = %#v\n", name, meta.Types[name])
		}
		id++
	}

	return nil
}
