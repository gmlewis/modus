/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package wasm

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/gmlewis/modus/lib/wasmextractor"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
)

func FilterMetadata(config *config.Config, meta *metadata.Metadata) error {
	wasmFilePath := filepath.Join(config.OutputDir, config.WasmFileName)

	bytes, err := wasmextractor.ReadWasmFile(wasmFilePath)
	if err != nil {
		return err
	}

	info, err := wasmextractor.ExtractWasmInfo(bytes)
	if err != nil {
		return err
	}

	buf, _ := json.MarshalIndent(info, "", "  ")
	gmlPrintf("GML: wasm/filters.go: FilterMetadata: wasm info:\n%s", buf)

	filterExportsImportsAndTypes(info, meta)

	return nil
}

func filterExportsImportsAndTypes(info *wasmextractor.WasmInfo, meta *metadata.Metadata) {
	// Remove unused imports (can easily happen when user code doesn't use all imports from a package)
	imports := make(map[string]bool, len(info.Imports))
	for _, i := range info.Imports {
		imports[i.Name] = true
	}
	for name := range meta.FnImports {
		if _, ok := imports[name]; !ok {
			gmlPrintf("GML: wasm/filters.go: FilterMetadata: removing unused FnImport: %v", name)
			delete(meta.FnImports, name)
		}
	}

	// Remove unused exports (less likely to happen, but still check)
	exports := make(map[string]bool, len(info.Exports))
	for _, e := range info.Exports {
		exports[e.Name] = true
	}
	for name := range meta.FnExports {
		if _, ok := exports[name]; !ok {
			//TODO: delete(meta.FnExports, name)
			if strings.HasPrefix(name, "__modus_") {
				gmlPrintf("GML: wasm/filters.go: FilterMetadata: removing special modus export: %v", name)
				delete(meta.FnExports, name)
			}
		}
	}

	// For now, leave all types as-is. This continues to be a source of great confusion
	// as we try to figure out what types are actually needed for MoonBit.
	/*
		// Remove unused types (they might not be needed now, due to removing functions)
		var keptTypes = make(metadata.TypeMap, len(meta.Types))
		for _, fn := range append(utils.MapValues(meta.FnImports), utils.MapValues(meta.FnExports)...) {
			for _, param := range fn.Parameters {
				// paramType := stripError(param.Type)
				if _, ok := meta.Types[param.Type]; ok {
					keptTypes[param.Type] = meta.Types[param.Type]
					//TODO: delete(meta.Types, paramType)
				} else {
					gmlPrintf("GML: wasm/filters.go: FilterMetadata: removing param type: %v", param.Type)
				}
			}
			for _, result := range fn.Results {
				// resultType := stripError(result.Type)
				if _, ok := meta.Types[result.Type]; ok {
					keptTypes[result.Type] = meta.Types[result.Type]
					//TODO: delete(meta.Types, resultType)
				} else {
					gmlPrintf("GML: wasm/filters.go: FilterMetadata: removing result type: %v", result.Type)
				}
			}
		}

		// ensure types used by kept types are also kept
		for dirty := true; len(meta.Types) > 0 && dirty; {
			dirty = false

			keep := func(t string) {
				if _, ok := meta.Types[t]; ok {
					if _, ok := keptTypes[t]; !ok {
						keptTypes[t] = meta.Types[t]
						//TODO: delete(meta.Types, t)
						//TODO: dirty = true
					}
				}
			}

			for _, t := range keptTypes {
				// if utils.IsPointerType(t.Name) {
				if utils.IsOptionType(t.Name) {
					keep(utils.GetUnderlyingType(t.Name))
				} else if utils.IsListType(t.Name) {
					keep(utils.GetListSubtype(t.Name))
				} else if utils.IsMapType(t.Name) {
					kt, vt := utils.GetMapSubtypes(t.Name)
					keep(kt)
					keep(vt)
					keep(fmt.Sprintf("Array[%v]", kt))
					keep(fmt.Sprintf("Array[%v]", vt))
				}

				for _, field := range t.Fields {
					keep(field.Type)
				}
			}
		}
		meta.Types = keptTypes
	*/
}
