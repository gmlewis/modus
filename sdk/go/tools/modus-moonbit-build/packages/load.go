// This file is based on: https://cs.opensource.google/go/x/tools/+/refs/tags/v0.28.0:go/packages/packages.go
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packages

// See doc.go for package documentation and implementation notes.

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
)

// Load loads and returns the MoonBit packages named by the given patterns.
//
// The cfg parameter specifies loading options; nil behaves the same as an empty [Config].
//
// The [Config.Mode] field is a set of bits that determine what kinds
// of information should be computed and returned. Modes that require
// more information tend to be slower. See [LoadMode] for details
// and important caveats. Its zero value is equivalent to
// [NeedName] | [NeedFiles] | [NeedCompiledMoonBitFiles].
//
// Each call to Load returns a new set of [Package] instances.
// The Packages and their Imports form a directed acyclic graph.
//
// If the [NeedTypes] mode flag was set, each call to Load uses a new
// [types.Importer], so [types.Object] and [types.Type] values from
// different calls to Load must not be mixed as they will have
// inconsistent notions of type identity.
//
// If any of the patterns was invalid as defined by the
// underlying build system, Load returns an error.
// It may return an empty list of packages without an error,
// for instance for an empty expansion of a valid wildcard.
// Errors associated with a particular package are recorded in the
// corresponding Package's Errors list, and do not cause Load to
// return an error. Clients may need to handle such errors before
// proceeding with further analysis. The [PrintErrors] function is
// provided for convenient display of all errors.
func Load(cfg *Config, mod *modinfo.ModuleInfo, patterns ...string) ([]*Package, error) {
	if mod.AlreadyProcessed(cfg.Dir) {
		return nil, nil
	}
	mod.MarkProcessed(cfg.Dir)

	// Note that if patterns=="." then only the main package is loaded.
	// Also note that all MoonBit dependencies must be parsed in either
	// case in order to record all the imported host functions.
	if len(patterns) == 0 {
		return nil, fmt.Errorf("no patterns provided")
	}

	returnMainPkgOnly := patterns[0] == "."
	if cfg.RootAbsPath == "" {
		var err error
		cfg.RootAbsPath, err = filepath.Abs(cfg.Dir)
		if err != nil {
			return nil, fmt.Errorf("failed to get absolute path for %q: %w", cfg.Dir, err)
		}
	}

	// process imports
	var imports []*ast.ImportSpec
	var subPackages []*Package
	gmlPrintf("GML: packages/load.go: Parsing '%v/moon.pkg.json'", cfg.Dir)
	buf, err := os.ReadFile(filepath.Join(cfg.Dir, "moon.pkg.json"))
	if err != nil {
		return nil, err
	}
	var moonPkgJSON MoonPkgJSON
	if err := json.Unmarshal(buf, &moonPkgJSON); err != nil {
		return nil, err
	}
	for _, imp := range moonPkgJSON.Imports {
		var jsonMsg any
		if err := json.Unmarshal(imp, &jsonMsg); err != nil {
			return nil, err
		}
		switch v := jsonMsg.(type) {
		case string:
			imports = append(imports, &ast.ImportSpec{Path: &ast.BasicLit{Value: fmt.Sprintf("%q", v)}})
			dir, err := mod.GetModuleAbsPath(cfg.RootAbsPath, v)
			if err != nil {
				return nil, err
			}
			subCfg := &Config{
				Mode:        cfg.Mode,
				RootAbsPath: cfg.RootAbsPath,
				Dir:         dir,
				PackageName: v,
				Env:         cfg.Env,
			}
			// recurse into sub-packages
			subPkgs, err := Load(subCfg, mod, patterns...)
			if err != nil {
				return nil, err
			}
			subPackages = append(subPackages, subPkgs...)
		case map[string]any:
			path, ok1 := v["path"].(string)
			alias, ok2 := v["alias"].(string)
			if !ok1 || !ok2 {
				log.Printf("Warning: unexpected import map: %+v; skipping", v)
				continue
			}
			dir, err := mod.GetModuleAbsPath(cfg.RootAbsPath, path)
			if err != nil {
				return nil, err
			}
			subCfg := &Config{
				Mode:         cfg.Mode,
				RootAbsPath:  cfg.RootAbsPath,
				Dir:          dir,
				PackageName:  path,
				PackageAlias: alias,
				Env:          cfg.Env,
			}
			// recurse into sub-packages
			subPkgs, err := Load(subCfg, mod, patterns...)
			if err != nil {
				return nil, err
			}
			subPackages = append(subPackages, subPkgs...)
		default:
			gmlPrintf("Warning: unexpected import type: %T; skipping", jsonMsg)
		}
	}

	// ensure that "is-main" is false:
	moonPkgJSON.IsMain = false

	sourceFiles, err := filepath.Glob(filepath.Join(cfg.Dir, "*.mbt"))
	if err != nil {
		return nil, err
	}

	pkgPath := "@" + filepath.Base(cfg.Dir)
	log.Printf("\n\n*** GML: cfg.PackageName=%q ***\n\n", cfg.PackageName)
	if cfg.PackageName == "" {
		// The top-level package must be "main" for the runtime to load it properly.
		cfg.PackageName = "main"
		// However, its PkgPath must be empty.
		pkgPath = ""
	}

	result := &Package{
		ID:      "moonbit-main", // this appears to be unused
		Name:    cfg.PackageName,
		PkgPath: pkgPath,
		TypesInfo: &types.Info{
			Defs: map[*ast.Ident]types.Object{},
		},
		MoonPkgJSON:  moonPkgJSON,
		StructLookup: map[string]*ast.TypeSpec{},
	}
	pkg := types.NewPackage(result.PkgPath, cfg.PackageName)
	for _, sourceFile := range sourceFiles {
		// TODO: Support "target" MoonPkgJSON field with ["wasm"] and ["not","wasm"].
		if strings.HasSuffix(sourceFile, "_test.mbt") { // ignore test files
			continue
		}
		result.MoonBitFiles = append(result.MoonBitFiles, sourceFile)
		gmlPrintf("GML: packages/load.go: Parsing '%v'", sourceFile)
		buf, err := os.ReadFile(sourceFile)
		if err != nil {
			return nil, err
		}
		result.processSourceFile(pkg, sourceFile, buf, imports)
	}

	if returnMainPkgOnly {
		return []*Package{result}, nil
	}
	return append([]*Package{result}, subPackages...), nil
}
