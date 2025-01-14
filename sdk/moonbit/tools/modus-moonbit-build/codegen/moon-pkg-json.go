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
	"io"
	"log"

	"github.com/hypermodeinc/modus/sdk/moonbit/tools/modus-moonbit-build/packages"
)

// updateMoonPkgJSON updates the moon.pkg.json file with the new imports
// and functions and writes it to the given writer.
func updateMoonPkgJSON(w io.Writer, pkg *packages.Package, imports map[string]string, functions []*funcInfo) error {
	// currentImports := map[string]bool{}
	// for _, imp := range pkg.MoonPkgJSON.Imports {
	// currentImports[imp.Path] = true
	// }
	for k, v := range imports {
		log.Printf("GML: k=%q, v=%q", k, v)
	}
	for i, fn := range functions {
		log.Printf("GML: i=%v, fn=%#v", i, fn)
	}

	buf, err := json.MarshalIndent(pkg.MoonPkgJSON, "", "  ")
	if err != nil {
		return err
	}
	if _, err := io.Copy(w, bytes.NewReader(buf)); err != nil {
		return err
	}

	return nil
}
