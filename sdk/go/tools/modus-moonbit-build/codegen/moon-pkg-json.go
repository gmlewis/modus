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
		if _, ok := currentImports[k]; !ok {
			log.Printf("adding import %q to moon.pkg.json", k)
			pkg.MoonPkgJSON.Imports = append(pkg.MoonPkgJSON.Imports, json.RawMessage(k))
		}
	}

	if pkg.MoonPkgJSON.Targets == nil {
		pkg.MoonPkgJSON.Targets = map[string][]string{}
	}
	pkg.MoonPkgJSON.Targets["modus_post_generated.mbt"] = []string{"wasm"}
	pkg.MoonPkgJSON.Targets["modus_pre_generated.mbt"] = []string{"wasm"}

	currentExports := map[string]bool{}
	if pkg.MoonPkgJSON.LinkTargets == nil {
		pkg.MoonPkgJSON.LinkTargets = map[string]*packages.LinkTarget{
			"wasm": {ExportMemoryName: "memory"},
		}
	} else if wasmLinkTarget, ok := pkg.MoonPkgJSON.LinkTargets["wasm"]; ok {
		wasmLinkTarget.ExportMemoryName = "memory"
		for _, export := range wasmLinkTarget.Exports {
			currentExports[export] = true
		}
	}
	for _, fn := range functions {
		modusName := fmt.Sprintf("__modus_%v", fn.function.Name.Name)
		if _, ok := currentExports[modusName]; !ok {
			log.Printf("adding link.wasm.export %q to moon.pkg.json", modusName)
			pkg.MoonPkgJSON.LinkTargets["wasm"].Exports = append(pkg.MoonPkgJSON.LinkTargets["wasm"].Exports, modusName)
		}
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
