// This file is based on: https://cs.opensource.google/go/x/tools/+/refs/tags/v0.28.0:go/packages/packages.go
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packages

// See doc.go for package documentation and implementation notes.

import (
	"encoding/json"
)

// MoonPkgJSON represents the JSON format of the moon.pkg.json file.
type MoonPkgJSON struct {
	IsMain     bool              `json:"is-main,omitempty"`
	Imports    []json.RawMessage `json:"import,omitempty"`
	TestImport []json.RawMessage `json:"test-import,omitempty"`
	// Targets is a map of filenames (keys) followed by its compilation rules
	// (which is an array of strings).
	// e.g. "file_notwasm.mbt": ["not", "wasm"], "file_wasm.mbt": ["wasm"], etc.
	Targets map[string][]string `json:"targets,omitempty"`
	// LinkTargets is a map of link targets (e.g. "wasm", "js", etc.)
	// followed by its exports.
	LinkTargets map[string]*LinkTarget `json:"link,omitempty"`
}

// LinkTarget represents the JSON format of a link target in moon.pkg.json.
type LinkTarget struct {
	// Exports is a list of exported functions in the resulting wasm module.
	Exports []string `json:"exports,omitempty"`
	// ExportMemoryName is the name of the exported memory (e.g. "memory").
	ExportMemoryName string `json:"export-memory-name,omitempty"`
}
