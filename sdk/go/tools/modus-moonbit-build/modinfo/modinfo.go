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
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"

	"github.com/hashicorp/go-version"
)

const sdkModulePath = "gmlewis/modus"

type ModuleInfo struct {
	ModulePath      string
	ModusSDKVersion *version.Version
	// DepPaths is a map of module names to their absolute paths on disk.
	DepPaths map[string]string

	// alreadyProcessed is a map of package dirs that have already been processed.
	alreadyProcessed map[string]bool
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
		DepPaths:   map[string]string{},
	}

	for modName, modVer := range moonMod.Deps {
		ver, absPath, err := processDependency(cfg.SourceDir, modName, modVer)
		if err != nil {
			return nil, err
		}
		result.DepPaths[modName] = absPath
		if modName == sdkModulePath {
			result.ModusSDKVersion = ver
		}
	}

	return result, nil
}

// GetModuleAbsPath returns the absolute path of the requested module directory.
func (m *ModuleInfo) GetModuleAbsPath(rootAbsPath, modName string) (string, error) {
	for name, path := range m.DepPaths {
		if strings.HasPrefix(modName, name) {
			fullPath := filepath.Join(path, strings.TrimPrefix(modName, name))
			return fullPath, nil
		}
	}
	// Could not find import. Check the ".mooncakes" subdir.
	mooncakesPath := filepath.Join(rootAbsPath, ".mooncakes")
	parts := strings.Split(modName, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("module '%v' not found", modName)
	}
	name := filepath.Join(parts[0], parts[1])
	path := filepath.Join(mooncakesPath, name)
	if _, err := os.Stat(path); err == nil {
		m.DepPaths[name] = path
		fullPath := filepath.Join(path, strings.TrimPrefix(modName, name))
		return fullPath, nil
	}
	return "", fmt.Errorf("module '%v' not found - missing dir: '%v'", modName, path)
}

func (m *ModuleInfo) ResetAlreadyProcessed() {
	m.alreadyProcessed = nil
}

func (m *ModuleInfo) AlreadyProcessed(dir string) bool {
	if m.alreadyProcessed == nil {
		m.alreadyProcessed = map[string]bool{}
	}
	return m.alreadyProcessed[dir]
}

func (m *ModuleInfo) MarkProcessed(dir string) {
	if m.alreadyProcessed == nil {
		m.alreadyProcessed = map[string]bool{}
	}
	m.alreadyProcessed[dir] = true
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

func getModuleAbsPath(dirName string) (string, error) {
	fs, err := os.Stat(dirName)
	if err != nil || !fs.IsDir() {
		return "", err
	}
	absPath, err := filepath.Abs(dirName)
	if err != nil {
		return "", fmt.Errorf("filepath.Abs('%v'): %w", dirName, err)
	}
	return absPath, nil
}

func processDependency(sourceDir, modName string, modVer json.RawMessage) (ver *version.Version, absPath string, err error) {
	var dst any
	if err := json.Unmarshal(modVer, &dst); err != nil {
		return nil, "", fmt.Errorf("moon.mod.json json.Unmarshal: %w", err)
	}
	switch v := dst.(type) {
	case string:
		ver, err = version.NewVersion(v)
		if err != nil {
			return nil, "", fmt.Errorf("bad version '%v': %w", v, err)
		}
		sdkPath := filepath.Join(sourceDir, ".mooncakes", modName)
		absPath, err = getModuleAbsPath(sdkPath)
		if err != nil {
			log.Printf("WARNING: unable to load module '%v' 'moon.mod.json' by looking in dir '%v': %v", modName, sdkPath, err)
		}
	case map[string]interface{}:
		if path, ok := v["path"].(string); ok {
			sdkPath := filepath.Join(sourceDir, path)
			absPath, err = getModuleAbsPath(sdkPath)
			if err != nil {
				log.Printf("WARNING: unable to load module '%v' 'moon.mod.json' by looking in dir '%v/%v': %v", modName, sourceDir, path, err)
				return ver, absPath, nil
			}
			sdkMoonMod, err := parseMoonModJSON(sdkPath)
			if err != nil {
				log.Printf("WARNING: unable to parse module '%v' 'moon.mod.json' dir '%v/%v': %v", modName, sourceDir, path, err)
				return ver, absPath, nil
			}
			ver, err = version.NewVersion(sdkMoonMod.Version)
			if err != nil {
				return nil, "", fmt.Errorf("bad version '%v': %w", v, err)
			}
		}
	default:
		return nil, "", fmt.Errorf("unexpected type for modus version: %T", dst)
	}
	return ver, absPath, nil
}
