/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package codegen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"slices"
	"strings"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/packages"
)

// updateMoonPkgJSON updates the moon.pkg.json file with the new imports
// and functions and writes it to the given writer.
func updateMoonPkgJSON(w io.Writer, pkg *packages.Package, imports map[string]string, functions []*funcInfo) error {
	currentImports := map[string]bool{}
	for _, imp := range pkg.MoonPkgJSON.Imports {
		var value any
		if err := json.Unmarshal(imp, &value); err != nil {
			return fmt.Errorf("updateMoonPkgJSON: %w", err)
		}
		switch v := value.(type) {
		case string:
			currentImports[v] = true
		default:
			log.Printf("WARNING: updateMoonPkgJSON: unexpected import type: %T, ignored.", value)
		}
	}
	for k := range imports {
		k = strings.TrimSpace(k)
		if k == "" {
			continue
		}
		if _, ok := currentImports[k]; !ok {
			pkg.MoonPkgJSON.Imports = append(pkg.MoonPkgJSON.Imports, json.RawMessage(`"`+k+`"`))
		}
	}

	if pkg.MoonPkgJSON.Targets == nil {
		pkg.MoonPkgJSON.Targets = map[string][]string{}
	}
	pkg.MoonPkgJSON.Targets["modus_post_generated.mbt"] = []string{"wasm"}
	pkg.MoonPkgJSON.Targets["modus_pre_generated.mbt"] = []string{"wasm"}

	// no need to preserve current exports.
	if pkg.MoonPkgJSON.LinkTargets == nil {
		pkg.MoonPkgJSON.LinkTargets = map[string]*packages.LinkTarget{
			"wasm": {ExportMemoryName: "memory"},
		}
	}
	wasmLinkTarget, ok := pkg.MoonPkgJSON.LinkTargets["wasm"]
	if ok {
		wasmLinkTarget.ExportMemoryName = "memory"
	} else {
		wasmLinkTarget = &packages.LinkTarget{ExportMemoryName: "memory"}
		pkg.MoonPkgJSON.LinkTargets["wasm"] = wasmLinkTarget
	}

	// TODO: Only include the exports that are actually needed.
	overrides := []string{ // clear out existing exports
		"cabi_realloc",
		"store8",
		"store32",
		"load32",
		"malloc",
		"free",
		"copy",
		"ptr2str",
		"read_map",
		"write_map",
		"ptr_to_none",
	}
	for _, v := range imports {
		if v == "@time" {
			overrides = append(overrides,
				"zoned_date_time_from_unix_seconds_and_nanos",
				"duration_from_nanos",
			)
			break
		}
	}
	wasmLinkTarget.Exports = overrides

	for _, fn := range functions {
		modusName := fmt.Sprintf("__modus_%v:%[1]v", fn.function.Name.Name)
		pkg.MoonPkgJSON.LinkTargets["wasm"].Exports = append(pkg.MoonPkgJSON.LinkTargets["wasm"].Exports, modusName)
	}
	slices.Sort(pkg.MoonPkgJSON.LinkTargets["wasm"].Exports)

	buf, err := json.MarshalIndent(pkg.MoonPkgJSON, "", "  ")
	if err != nil {
		return err
	}
	if _, err := io.Copy(w, bytes.NewReader(buf)); err != nil {
		return err
	}

	return nil
}
