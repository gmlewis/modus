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

	// Ensure that `warn-list` and `supported-target` have been set.
	pkg.MoonPkgJSON.WarnList = "-44"
	pkg.MoonPkgJSON.SupportedTargets = []string{"wasm"}

	// TODO: Only include the exports that are actually needed.
	overrides := []string{ // clear out existing exports
		"cabi_realloc",
		"copy",
		"free",
		"load32",
		"malloc",
		"moonbit_array_bool_from_fixed",
		"moonbit_array_byte_from_fixed",
		"moonbit_array_char_from_fixed",
		"moonbit_array_double_from_fixed",
		"moonbit_array_float_from_fixed",
		"moonbit_array_int16_from_fixed",
		"moonbit_array_int64_from_fixed",
		"moonbit_array_int_from_fixed",
		"moonbit_array_string_from_fixed",
		"moonbit_array_uint16_from_fixed",
		"moonbit_array_uint64_from_fixed",
		"moonbit_array_uint_from_fixed",
		"moonbit_bytes_make",
		"moonbit_bytes_to_array",
		"moonbit_float32_array_make",
		"moonbit_float_array_make",
		"moonbit_i32_array_make",
		"moonbit_int16_array_make",
		"moonbit_int64_array_make",
		"moonbit_ref_array_make",
		"ptr2double_array",
		"ptr2float_array",
		"ptr2int64_array",
		"ptr2int_array",
		"ptr2str",
		"ptr2uint64_array",
		"ptr2uint_array",
		"ptr_to_none",
		"read_map",
		"store32",
		"store8",
		"write_map",
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
