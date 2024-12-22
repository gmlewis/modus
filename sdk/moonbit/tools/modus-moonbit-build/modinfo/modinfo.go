/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package modinfo

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/hypermodeinc/modus/sdk/moonbit/tools/modus-moonbit-build/config"

	"github.com/hashicorp/go-version"
)

const sdkModulePath = "gmlewis/modus"

type ModuleInfo struct {
	ModulePath      string
	ModusSDKVersion *version.Version
}

type MoonModJSON struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Deps        map[string]string `json:"deps"`
	Readme      string            `json:"readme"`
	Repository  string            `json:"repository"`
	License     string            `json:"license"`
	Keywords    []string          `json:"keywords"`
	Description string            `json:"description"`
}

type MoonPkgJSON struct {
	Imports []json.RawMessage `json:"import"`
}

func CollectModuleInfo(config *config.Config) (*ModuleInfo, error) {
	modFilePath := filepath.Join(config.SourceDir, "moon.mod.json")
	data, err := os.ReadFile(modFilePath)
	if err != nil {
		return nil, fmt.Errorf("moon.mod.json not found: %w", err)
	}

	var moonMod MoonModJSON
	if err := json.Unmarshal(data, &moonMod); err != nil {
		return nil, fmt.Errorf("moon.mod.json json.Unmarshal: %w", err)
	}

	modPath := moonMod.Name
	config.WasmFileName = path.Base(modPath) + ".wasm"

	result := &ModuleInfo{
		ModulePath: modPath,
	}

	for modName, modVer := range moonMod.Deps {
		if modName == sdkModulePath {
			if ver, err := version.NewVersion(modVer); err == nil {
				result.ModusSDKVersion = ver
			}
		}
	}

	return result, nil
}
