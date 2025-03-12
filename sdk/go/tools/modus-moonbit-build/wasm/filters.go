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
	"fmt"
	"log"
	"path/filepath"

	"github.com/gmlewis/modus/lib/wasmextractor"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
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

	// for debugging purposes:
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
			gmlPrintf("GML: WARNING: wasm/filters.go: FilterMetadata: removing unused meta.FnImports['%v']", name)
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
			gmlPrintf("GML: WARNING: wasm/filters.go: FilterMetadata: removing unused meta.FnExports['%v']", name)
			delete(meta.FnExports, name)
			// if strings.HasPrefix(name, "__modus_") {
			// 	log.Printf("GML: wasm/filters.go: FilterMetadata: removing special modus export: %v", name)
			// 	delete(meta.FnExports, name)
			// }
		}
	}

	// Remove unused types (they might not be needed now, due to removing functions)
	var keptTypes = make(metadata.TypeMap, len(meta.Types))
	for _, fn := range append(utils.MapValues(meta.FnImports), utils.MapValues(meta.FnExports)...) {
		for _, param := range fn.Parameters {
			if _, ok := meta.Types[param.Type]; ok {
				keptTypes[param.Type] = meta.Types[param.Type]
				// gmlPrintf("GML: wasm/filters.go: FilterMetadata: keeping meta.Types[paramType='%v']", param.Type)
				delete(meta.Types, param.Type)
				// } else {
				// 	log.Printf("GML: WARNING: wasm/filters.go: FilterMetadata: NOT KEEPING param.Type: '%v'", param.Type)
			}
		}
		for _, result := range fn.Results {
			if _, ok := meta.Types[result.Type]; ok {
				keptTypes[result.Type] = meta.Types[result.Type]
				// gmlPrintf("GML: wasm/filters.go: FilterMetadata: keeping meta.Types[result.Type='%v']", result.Type)
				delete(meta.Types, result.Type)
				// } else {
				// 	log.Printf("GML: WARNING: wasm/filters.go: FilterMetadata: NOT KEEPING result.Type: '%v'", result.Type)
			}
		}
	}

	// ensure types used by kept types are also kept
	// Note that this currently only looks at types one-level deep, assuming that the source-code
	// processing in the `packages` package has already fully added all underlying types.
	for dirty := true; len(meta.Types) > 0 && dirty; {
		dirty = false

		keep := func(t string) {
			if t == "Unit" {
				return // no need to keep 'Unit'.
			}
			if _, ok := meta.Types[t]; ok {
				if _, ok := keptTypes[t]; !ok {
					keptTypes[t] = meta.Types[t]
					// gmlPrintf("GML: SECOND PASS: wasm/filters.go: FilterMetadata: keeping meta.Types['%v']", t)
					delete(meta.Types, t)
					dirty = true
				}
			} else if _, ok := keptTypes[t]; !ok {
				log.Printf("PROGRAMMING ERROR!!! wasm/filters.go: UNABLE TO KEEP type '%v' since it is missing from meta.Types!!!", t)
			}
		}

		for _, t := range keptTypes {
			// types should not have default values at this point.
			tName := utils.StripDefaultValue(t.Name) // TODO: verify and remove if true.
			typ, hasError, _ := utils.StripErrorAndOption(tName)
			if hasError {
				keep("(String)") // needed for all !Error processing
			}
			keep(typ)

			if utils.IsOptionType(tName) {
				keep(utils.GetUnderlyingType(tName))
			} else if utils.IsListType(tName) {
				keep(utils.GetListSubtype(tName))
			} else if utils.IsMapType(tName) {
				kt, vt := utils.GetMapSubtypes(tName)
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
}
