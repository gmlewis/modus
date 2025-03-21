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
	"os"
	"path/filepath"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/packages"
)

func loadPackages(dir string, mod *modinfo.ModuleInfo) (map[string]*packages.Package, error) {
	mode := packages.NeedName |
		packages.NeedImports |
		packages.NeedDeps |
		packages.NeedTypes |
		packages.NeedSyntax |
		packages.NeedTypesInfo

	rootAbsPath, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	cfg := &packages.Config{
		RootAbsPath: rootAbsPath,
		Mode:        mode,
		Dir:         rootAbsPath,
		Env:         os.Environ(),
	}

	mod.ResetAlreadyProcessed()

	pkgs, err := packages.Load(cfg, mod, "./...")
	if err != nil {
		return nil, err
	}

	pkgMap := make(map[string]*packages.Package)
	for _, pkg := range pkgs {
		expandPackages(pkg, pkgMap)
	}

	return pkgMap, nil
}

func expandPackages(pkg *packages.Package, pkgMap map[string]*packages.Package) {
	for _, imp := range pkg.Imports {
		if _, ok := pkgMap[imp.PkgPath]; ok {
			continue
		}
		expandPackages(imp, pkgMap)
	}
	pkgMap[pkg.PkgPath] = pkg
}
