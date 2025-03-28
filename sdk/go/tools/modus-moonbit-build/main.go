/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/codegen"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/compiler"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/end2end"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metagen"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/wasm"
	"github.com/hypermodeinc/modus/lib/manifest"
)

func main() {
	log.SetFlags(0)

	start := time.Now()
	trace := utils.IsTraceModeEnabled()
	if trace {
		log.Println("Starting build process...")
	}

	config, err := config.GetConfig()
	if err != nil {
		exitWithError("Error", err)
	}

	if trace {
		log.Println("Configuration loaded.")
	}

	if len(config.EndToEndTests) > 0 {
		_, thisFilename, _, _ := runtime.Caller(0)
		repoAbsPath := filepath.Join(filepath.Dir(thisFilename), "../../../..")
		for _, plugin := range config.EndToEndTests {
			config.SourceDir = filepath.Join(repoAbsPath, plugin.Path)
			config.WasmFileName = plugin.Name + ".wasm"
			config.OutputDir = filepath.Join(config.SourceDir, "build")
			if err := end2end.RunTest(config, repoAbsPath, start, trace, plugin); err != nil {
				exitWithError("Error running end-to-end test", err)
			}
		}
		log.Printf("\n\n*** ALL END-TO-END TESTS PASSED ***\n\n")
	} else {
		buildPlugin(config, start, trace)
	}
}

func buildPlugin(config *config.Config, start time.Time, trace bool) {
	if err := compiler.Validate(config); err != nil {
		exitWithError("Error", err)
	}

	if trace {
		log.Println("Configuration validated.")
	}

	mod, err := modinfo.CollectModuleInfo(config)
	if err != nil {
		exitWithError("Error", err)
	}

	if trace {
		log.Println("Module info collected.")
	}

	if err := codegen.PreProcess(config, mod); err != nil {
		exitWithError("Error while pre-processing source files", err)
	}

	if trace {
		log.Println("Pre-processing done.")
	}

	meta, err := metagen.GenerateMetadata(config, mod)
	if err != nil {
		exitWithError("Error generating metadata", err)
	}

	if trace {
		log.Println("Metadata generated.")
	}

	if err := codegen.PostProcess(config, meta); err != nil {
		exitWithError("Error while post-processing source files", err)
	}

	if trace {
		log.Println("Post-processing done.")
	}

	if err := compiler.Compile(config); err != nil {
		exitWithError("Error building wasm", err)
	}

	if trace {
		log.Println("Wasm compiled.")
	}

	if err := wasm.FilterMetadata(config, meta); err != nil {
		exitWithError("Error filtering metadata", err)
	}

	if trace {
		log.Println("Metadata filtered.")
	}

	if err := wasm.WriteMetadata(config, meta); err != nil {
		exitWithError("Error writing metadata", err)
	}

	if trace {
		log.Println("Metadata written.")
	}

	if err := validateAndCopyManifestToOutput(config); err != nil {
		exitWithError("Manifest error", err)
	}

	if trace {
		log.Println("Manifest copied.")
	}

	if trace {
		log.Printf("Build completed in %.2f seconds.\n\n", time.Since(start).Seconds())
	}

	metagen.LogToConsole(meta)
}

func exitWithError(msg string, err error) {
	fmt.Fprintf(os.Stderr, msg+": %v\n", err)
	os.Exit(1)
}

func validateAndCopyManifestToOutput(config *config.Config) error {
	manifestFile := filepath.Join(config.SourceDir, "modus.json")
	if _, err := os.Stat(manifestFile); os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	data, err := os.ReadFile(manifestFile)
	if err != nil {
		return err
	}

	if err := validateManifest(data); err != nil {
		return err
	}

	outFile := filepath.Join(config.OutputDir, "modus.json")
	if err := os.WriteFile(outFile, data, 0644); err != nil {
		return err
	}

	return nil
}

func validateManifest(data []byte) error {
	// Make a copy of the data to avoid modifying the original
	// TODO: this should be fixed in the manifest library
	manData := make([]byte, len(data))
	copy(manData, data)
	return manifest.ValidateManifest(manData)
}
