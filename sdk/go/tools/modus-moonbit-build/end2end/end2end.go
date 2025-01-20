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

func RunTest(config *config.Config, repoAbsPath string, start time.Time, trace bool, plugin *config.Plugin) {
	cwd, err := os.Getwd()
	must(err)
	defer func() {
		must(os.Chdir(cwd))
	}()

	log.Printf("Changing directory to %q", config.SourceDir)
	must(os.Chdir(config.SourceDir))

	log.Printf("Waiting for Modus CLI to terminate and release its port")
	if !waitForPortToBeFree(graphqlPort, 20*time.Second) {
		log.Printf("Start of main loop: port %v is still in use after waiting", graphqlPort)
	} else {
		log.Printf("Start of main loop: port %v is now free", graphqlPort)
	}

	log.Printf("Running Modus CLI for %q", plugin.Name)
	cliCmd := filepath.Join(repoAbsPath, "cli/bin/modus.js")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, cliCmd, "dev")
	// Set up stdout and stderr to tee to the current process' stdout and stderr
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	go func() {
		must(cmd.Start())

		pid := cmd.Process.Pid
		// Wait for port 8686 to be in use:
		if !waitForPortToBeInUse(graphqlPort, 20*time.Second) {
			log.Printf("WARNING: Modus CLI runner: port %v is not in use after waiting - did the Modus CLI start?", graphqlPort)
		} else {
			log.Printf("\n\nModus CLI runner: port %v is now ready for use.", graphqlPort)
		}

		// Get child PIDs
		childPIDs, err := getChildPIDs(pid)
		if err != nil {
			fmt.Printf("Error getting child PIDs: %v\n", err)
			return
		}
		defer func() {
			for _, childPID := range childPIDs {
				log.Printf("Killing child PID %v", childPID)
				if err := killProcess(childPID); err != nil {
					log.Printf("Failed to kill child PID %v: %v", childPID, err)
				}
			}
		}()

		log.Printf("Waiting for Modus CLI to finish... (pid: %v, child pids: %+v)", pid, childPIDs)
		if err := cmd.Wait(); err != nil {
			log.Printf("cmd.Wait: %v", err)
		}
	}()

	if !waitForPortToBeInUse(graphqlPort, 20*time.Second) {
		log.Printf("WARNING: Prior to testing endpoints: port %v is not in use after waiting - did the Modus CLI start?", graphqlPort)
	} else {
		log.Printf("\n\nPrior to testing endpoints: port %v is now ready for use.", graphqlPort)
	}

	// Test each endpoint
	for _, endpoint := range plugin.Endpoints {
		must(testEndpoint(ctx, endpoint))
	}

	// Kill the Modus CLI
	log.Printf("Terminating Modus CLI")
	must(cmd.Cancel())
	cancel() // Cancel the context

	log.Printf("Waiting for Modus CLI to terminate and release its port")
	if !waitForPortToBeFree(graphqlPort, 20*time.Second) {
		log.Printf("End of main loop: port %v is still in use after waiting", graphqlPort)
	} else {
		log.Printf("End of main loop: port %v is now free", graphqlPort)
	}
}
