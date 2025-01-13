/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package compiler

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hypermodeinc/modus/sdk/moonbit/tools/modus-moonbit-build/config"

	"github.com/hashicorp/go-version"
)

const (
	minMoonBitVersion = "0.1.20250107"
	wasmBuildDir      = "target/wasm/release/build"
)

func format(config *config.Config) error {
	args := []string{"fmt"}
	args = append(args, config.CompilerOptions...)

	log.Printf("GML: Running: %v '%v'", config.CompilerPath, strings.Join(args, "' '"))
	cmd := exec.Command(config.CompilerPath, args...)
	cmd.Dir = config.SourceDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = append(os.Environ())
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Compile(config *config.Config) error {
	if err := format(config); err != nil {
		return err
	}

	args := []string{"build"}
	args = append(args, "--target", "wasm")
	args = append(args, config.CompilerOptions...)

	log.Printf("GML: Running: %v '%v'", config.CompilerPath, strings.Join(args, "' '")) // TODO(gmlewis): remove
	cmd := exec.Command(config.CompilerPath, args...)
	cmd.Dir = config.SourceDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = append(os.Environ()) // , "GOOS=wasip1", "GOARCH=wasm")
	if err := cmd.Run(); err != nil {
		return err
	}

	// Copy the resulting .wasm file into the build directory
	wasmInFile := filepath.Join(config.SourceDir, wasmBuildDir, config.WasmFileName)
	buf, err := os.ReadFile(wasmInFile)
	if err != nil {
		return err
	}
	wasmOutFile := filepath.Join(config.OutputDir, config.WasmFileName)
	return os.WriteFile(wasmOutFile, buf, 0644)
}

func Validate(config *config.Config) error {
	if !strings.HasPrefix(filepath.Base(config.CompilerPath), "moon") {
		return fmt.Errorf("compiler must be moon")
	}

	ver, err := getCompilerVersion(config)
	if err != nil {
		return err
	}

	minVer, _ := version.NewVersion(minMoonBitVersion)
	if ver.LessThan(minVer) {
		return fmt.Errorf("found MoonBit version %s, but version %s or later is required", ver, minVer)
	}

	return nil
}

func getCompilerVersion(config *config.Config) (*version.Version, error) {
	cmd := exec.Command(config.CompilerPath, "version")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	parts := strings.Split(string(output), " ")
	if len(parts) < 3 {
		compiler := filepath.Base(config.CompilerPath)
		return nil, fmt.Errorf("unexpected output from '%s version': %s", compiler, output)
	}

	log.Printf("GML: parts=%+v", parts) // TODO(gmlewis): remove
	return version.NewVersion(parts[1])
}
