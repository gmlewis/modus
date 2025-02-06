/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	SourceDir       string
	CompilerPath    string
	CompilerOptions []string
	WasmFileName    string
	OutputDir       string
	EndToEndTests   []*Plugin
	Timeout         time.Duration
}

func GetConfig() (*Config, error) {
	c := &Config{}

	defaultCompilerPath, err := getCompilerDefaultPath()
	if err != nil {
		return nil, err
	}

	flag.StringVar(&c.CompilerPath, "compiler", defaultCompilerPath, "Path to the compiler to use.")
	flag.StringVar(&c.OutputDir, "output", "build", "Output directory for the generated files. Relative paths are resolved relative to the source directory.")
	var endToEndTestsFilename string
	flag.StringVar(&endToEndTestsFilename, "test", "", "Path to the end-to-end tests JSON file (e.g. '-test end-to-end-tests.json').")
	var noTimeout bool
	flag.BoolVar(&noTimeout, "notimeout", false, "Disable the timeout for the end-to-end tests (useful for debugging).")
	var pluginsFlag string
	flag.StringVar(&pluginsFlag, "plugins", "", "Comma-separated list of plugin names to test. Must use with -test flag.")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "modus-moonbit-build - The build tool for MoonBit-based Modus apps")
		fmt.Fprintln(os.Stderr, "Usage: modus-moonbit-build [options] <source directory> [... additional compiler options ...]")
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if !noTimeout {
		c.Timeout = 30 * time.Second
	}

	if endToEndTestsFilename != "" {
		if err := c.loadEndToEndTests(endToEndTestsFilename, pluginsFlag); err != nil {
			log.Printf("Error loading end-to-end tests JSON file: %v", err)
			flag.Usage()
			os.Exit(1)
		}
		return c, nil
	}

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	c.SourceDir = flag.Arg(0)

	if flag.NArg() > 1 {
		c.CompilerOptions = flag.Args()[1:]
	}

	if _, err := os.Stat(c.SourceDir); err != nil {
		return nil, fmt.Errorf("error locating source directory: %w", err)
	}

	wasmFileName, err := getWasmFileName(c.SourceDir)
	if err != nil {
		return nil, fmt.Errorf("error determining output filename: %w", err)
	}
	c.WasmFileName = wasmFileName

	if !filepath.IsAbs(c.OutputDir) {
		c.OutputDir = filepath.Join(c.SourceDir, c.OutputDir)
	}

	if err := os.MkdirAll(c.OutputDir, 0755); err != nil {
		return nil, fmt.Errorf("error creating output directory: %w", err)
	}

	return c, nil
}

func getWasmFileName(sourceDir string) (string, error) {
	sourceDir, err := filepath.Abs(sourceDir)
	if err != nil {
		return "", err
	}
	return filepath.Base(sourceDir) + ".wasm", nil
}

func getCompilerDefaultPath() (string, error) {
	for _, arg := range os.Args {
		if strings.TrimLeft(arg, "-") == "compiler" {
			return "", nil
		}
	}

	path, err := exec.LookPath("moon")
	if err != nil {
		return "", fmt.Errorf("moon not found in PATH.\nSee https://docs.moonbitlang.com/en/latest/tutorial/tour.html#installation for installation instructions.")
	}

	return path, nil
}
