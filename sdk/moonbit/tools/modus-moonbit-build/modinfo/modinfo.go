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
	Name        string                     `json:"name"`
	Version     string                     `json:"version"`
	Deps        map[string]json.RawMessage `json:"deps"`
	Readme      string                     `json:"readme"`
	Repository  string                     `json:"repository"`
	License     string                     `json:"license"`
	Keywords    []string                   `json:"keywords"`
	Description string                     `json:"description"`
}

// CollectModuleInfo collects module information from moon.mod.json.
func CollectModuleInfo(cfg *config.Config) (*ModuleInfo, error) {
	moonMod, err := parseMoonModJSON(cfg.SourceDir)
	if err != nil {
		return nil, fmt.Errorf("parseMoonModJSON: %w", err)
	}

	modPath := moonMod.Name
	cfg.WasmFileName = path.Base(modPath) + ".wasm"

	result := &ModuleInfo{
		ModulePath: modPath,
	}

	for modName, modVer := range moonMod.Deps {
		if modName == sdkModulePath {
			var dst any
			if err := json.Unmarshal(modVer, &dst); err != nil {
				return nil, fmt.Errorf("moon.mod.json json.Unmarshal: %w", err)
			}
			switch v := dst.(type) {
			case string:
				if ver, err := version.NewVersion(v); err == nil {
					result.ModusSDKVersion = ver
				} else {
					return nil, fmt.Errorf("version.NewVersion: %w", err)
				}
			case map[string]interface{}:
				if path, ok := v["path"].(string); ok {
					sdkPath := filepath.Join(cfg.SourceDir, path)
					sdkMoonMod, err := parseMoonModJSON(sdkPath)
					if err != nil {
						return nil, fmt.Errorf("parseMoonModJSON: %w", err)
					}
					if ver, err := version.NewVersion(sdkMoonMod.Version); err == nil {
						result.ModusSDKVersion = ver
					} else {
						return nil, fmt.Errorf("version.NewVersion: %w", err)
					}
				}
			default:
				return nil, fmt.Errorf("unexpected type for modus version: %T", dst)
			}
		}
	}

	return result, nil
}

func parseMoonModJSON(sourceDir string) (*MoonModJSON, error) {
	moonModPath := filepath.Join(sourceDir, "moon.mod.json")
	f, err := os.Open(moonModPath)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer f.Close()

	var moonMod MoonModJSON
	if err := json.NewDecoder(f).Decode(&moonMod); err != nil {
		return nil, fmt.Errorf("json.Decode: %w", err)
	}

	return &moonMod, nil
}
