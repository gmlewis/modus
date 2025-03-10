/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package packages

import (
	"encoding/json"
	"fmt"
	"go/token"
	"go/types"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testPackageLoadHelper(t *testing.T, name, dir string, wantPackage *Package) {
	t.Helper()

	mode := NeedName | NeedImports | NeedDeps | NeedTypes | NeedSyntax | NeedTypesInfo
	cfg := &Config{Mode: mode, Dir: dir}
	got, err := Load(cfg, ".")
	if err != nil || len(got) != 1 {
		t.Fatal(err)
	}

	// Manually test the contents of TypesInfo since the map[*ast.Ident]types.Object
	// is not directly comparable.
	if len(got[0].TypesInfo.Defs) != len(wantPackage.TypesInfo.Defs) {
		gotKeys := make([]string, 0, len(got[0].TypesInfo.Defs))
		for k, v := range got[0].TypesInfo.Defs {
			fmt.Printf(`      {Name: %q}: &moonFunc{funcName: "%v"},`+"\n", k.Name, v)
			gotKeys = append(gotKeys, k.Name)
		}
		sort.Strings(gotKeys)
		wantKeys := make([]string, 0, len(wantPackage.TypesInfo.Defs))
		for k := range wantPackage.TypesInfo.Defs {
			wantKeys = append(wantKeys, k.Name)
		}
		sort.Strings(wantKeys)
		t.Errorf("%v Load() TypesInfo.Defs mismatch: got %v types, want %v", name, len(got[0].TypesInfo.Defs), len(wantPackage.TypesInfo.Defs))
		if diff := cmp.Diff(wantKeys, gotKeys); diff != "" {
			t.Errorf("%v Load() TypesInfo.Defs keys mismatch (-want +got):\n%v", name, diff)
		}
	}
	// Note that a single Def name can map to multiple definitions due to the pointer nature of the key.
	// Practically what this means is that the MoonBit makeshift parser can first create a forward reference
	// to an unknown `struct` and then later fill in the definition of the struct. One of the defs will be
	// empty and the other will be populated.
	gotTypesInfoDefs := make(map[string][]types.Object)
	for k, v := range got[0].TypesInfo.Defs {
		gotTypesInfoDefs[k.Name] = append(gotTypesInfoDefs[k.Name], v)
	}
	for k, wantDef := range wantPackage.TypesInfo.Defs {
		wantStr := fmt.Sprintf("%v", wantDef)
		gotDefs := gotTypesInfoDefs[k.Name]
		var matches int
		var diff string
		for _, gotDef := range gotDefs {
			gotStr := fmt.Sprintf("%v", gotDef)
			if diff = cmp.Diff(wantStr, gotStr); diff == "" {
				matches++
			}
		}
		if matches == 0 {
			t.Errorf("%v Load() mismatch for TypesInfo.Defs[%q] (-want +got):\n%v", name, k, diff)
		}
	}
	got[0].TypesInfo = nil
	wantPackage.TypesInfo = nil

	gotMoonPkgJSON, err := json.MarshalIndent(got[0].MoonPkgJSON, "", "  ")
	if err != nil {
		t.Fatalf("json.MarshalIndent(got[0]) failed: %v", err)
	}
	wantMoonPkgJSON, err := json.MarshalIndent(wantPackage.MoonPkgJSON, "", "  ")
	if err != nil {
		t.Fatalf("json.MarshalIndent(wantPackage) failed: %v", err)
	}
	if diff := cmp.Diff(string(wantMoonPkgJSON), string(gotMoonPkgJSON)); diff != "" {
		t.Errorf("%v Load() Package.MoonPkgJSON mismatch (-want +got):\n%v", name, diff)
	}
	got[0].MoonPkgJSON = MoonPkgJSON{}
	wantPackage.MoonPkgJSON = MoonPkgJSON{}

	if diff := cmp.Diff(wantPackage, got[0]); diff != "" {
		t.Errorf("%v Load() Package mismatch (-want +got):\n%v", name, diff)
	}
}

// moonFunc is a helper function to create a *types.Func.
type moonFunc struct {
	funcName string
	types.Object
}

var _ types.Object = &moonFunc{}

func (m *moonFunc) Exported() bool       { return false }
func (m *moonFunc) Id() string           { return "" }
func (m *moonFunc) Name() string         { return "" }
func (m *moonFunc) Parent() *types.Scope { return nil }
func (m *moonFunc) Pkg() *types.Package  { return nil }
func (m *moonFunc) Pos() token.Pos       { return 0 }
func (m *moonFunc) String() string       { return m.funcName }
func (m *moonFunc) Type() types.Type     { return nil }
