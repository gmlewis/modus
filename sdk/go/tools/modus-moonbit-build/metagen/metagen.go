/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package metagen

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"os"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/extractor"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/gitinfo"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
)

const sdkName = "modus-sdk-mbt"

func GenerateMetadata(config *config.Config, mod *modinfo.ModuleInfo) (*metadata.Metadata, error) {
	if _, err := os.Stat(config.SourceDir); err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	meta := metadata.NewMetadata()
	// meta.Module = mod.ModulePath  // "gmlewis/modus/examples/simple-example"
	meta.Module = "@" + filepath.Base(config.SourceDir) // "@simple" - should this be empty for the top-level?
	meta.Plugin = path.Base(mod.ModulePath)

	meta.SDK = sdkName
	if mod.ModusSDKVersion != nil {
		meta.SDK += "@" + mod.ModusSDKVersion.String()
	}

	if err := extractor.CollectProgramInfo(config, meta, mod); err != nil {
		return nil, fmt.Errorf("error collecting program info: %w", err)
	}

	// remove external funcs from FnExports
	for name := range meta.FnExports {
		if strings.HasPrefix(name, "@") {
			delete(meta.FnExports, name)
		}
	}

	gitinfo.TryCollectGitInfo(config, meta)

	return meta, nil
}
