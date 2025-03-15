/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// Package end2end contains the end-to-end test runner for modus-moonbit-build.
// Each test started the Modus runtime and runs a series of requests against it.
package end2end

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
)

const graphqlPort = 8686

func RunTest(config *config.Config, repoAbsPath string, start time.Time, trace bool, plugin *config.Plugin) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer func() {
		if err := os.Chdir(cwd); err != nil {
			log.Printf("os.Chdir(%q): %v", cwd, err)
		}
	}()

	if err := os.Chdir(config.SourceDir); err != nil {
		return err
	}

	if pid, err := getPIDUsingPort(graphqlPort); err != nil {
		log.Printf("Error getting PID using port %v: %v", graphqlPort, err)
	} else {
		if err := killProcess(pid); err != nil {
			log.Printf("Failed to kill process with pid %v: %v", pid, err)
		}
	}

	if !waitForPortToBeFree(graphqlPort, 20*time.Second) {
		log.Printf("Start of main loop: port %v is still in use after waiting", graphqlPort)
	}

	log.Printf("Running Modus CLI for %q", plugin.Name)
	cliCmd := filepath.Join(repoAbsPath, "cli/bin/modus.js")
	ctx := context.Background()
	cancel := func() {} // No-op cancel function
	if config.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, config.Timeout)
	}
	defer cancel()

	cmd := exec.CommandContext(ctx, cliCmd, "dev")
	// Set up stdout and stderr to tee to the current process' stdout and stderr
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}

	pid := cmd.Process.Pid
	// Wait for port 8686 to be in use:
	if !waitForPortToBeInUse(graphqlPort, 20*time.Second) {
		log.Printf("WARNING: Modus CLI runner: port %v is not in use after waiting - did the Modus CLI start?", graphqlPort)
	}

	// Get child PIDs
	childPIDs, err := getChildPIDs(pid)
	if err != nil {
		return fmt.Errorf("error getting child PIDs: %w\n", err)
	}
	allPIDs := append([]int{pid}, childPIDs...)
	defer killAllPIDs(allPIDs)

	go func() {
		if err := cmd.Wait(); err != nil {
			log.Printf("cmd.Wait: %v", err)
		}
	}()

	// Test each endpoint
	for _, endpoint := range plugin.Endpoints {
		if err := testEndpoint(ctx, endpoint); err != nil {
			return err
		}
	}

	// Kill the Modus CLI
	if err := cmd.Cancel(); err != nil {
		log.Printf("Failed to cancel Modus CLI: %v", err)
	}
	cancel() // Cancel the context
	killAllPIDs(allPIDs)

	if !waitForPortToBeFree(graphqlPort, 20*time.Second) {
		log.Printf("End of main loop: port %v is still in use after waiting", graphqlPort)
	}

	return nil
}
