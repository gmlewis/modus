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

const minMoonBitVersion = "0.1.20241218"

func Compile(config *config.Config) error {
	args := []string{"build"}
	args = append(args, "--target", "wasm")
	args = append(args, "-o", filepath.Join(config.OutputDir, config.WasmFileName))

	args = append(args, config.CompilerOptions...)
	args = append(args, ".")

	log.Printf("GML: Running: %v '%v'", config.CompilerPath, strings.Join(args, "' '"))
	cmd := exec.Command(config.CompilerPath, args...)
	cmd.Dir = config.SourceDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = append(os.Environ()) // , "GOOS=wasip1", "GOARCH=wasm")

	return cmd.Run()
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

	log.Printf("GML: parts=%+v", parts)
	return version.NewVersion(parts[1])
}