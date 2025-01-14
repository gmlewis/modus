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
func Load(cfg *Config, patterns ...string) ([]*Package, error) {
	// process imports
	var imports []*ast.ImportSpec
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
		default:
			log.Printf("Warning: unexpected import type: %T; skipping", jsonMsg)
		}
	}

	sourceFiles, err := filepath.Glob(filepath.Join(cfg.Dir, "*.mbt"))
	if err != nil {
		return nil, err
	}

	result := &Package{
		ID:      "moonbit-main", // unused?
		Name:    "main",
		PkgPath: "@" + filepath.Base(cfg.Dir),
		TypesInfo: &types.Info{
			Defs: map[*ast.Ident]types.Object{},
		},
		MoonPkgJSON: moonPkgJSON,
	}
	pkg := types.NewPackage(result.PkgPath, "main")
	for _, sourceFile := range sourceFiles {
		if strings.HasSuffix(sourceFile, "_test.mbt") { // ignore test files
			continue
		}
		result.MoonBitFiles = append(result.MoonBitFiles, sourceFile)
		buf, err := os.ReadFile(sourceFile)
		if err != nil {
			return nil, err
		}
		result.processSourceFile(pkg, sourceFile, buf, imports)
	}

	return []*Package{result}, nil
}
