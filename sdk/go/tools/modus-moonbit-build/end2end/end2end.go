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
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
)

func RunTest(config *config.Config, repoAbsPath string, start time.Time, trace bool, plugin *config.Plugin) {
	cwd, err := os.Getwd()
	must(err)
	defer func() {
		must(os.Chdir(cwd))
	}()

	log.Printf("Changing directory to %q", config.SourceDir)
	must(os.Chdir(config.SourceDir))

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
		log.Printf("Waiting for Modus CLI to finish... (pid: %v)", pid)
		time.Sleep(5 * time.Second)

		// Get child PIDs
		childPIDs, err := getChildPIDs(pid)
		if err != nil {
			fmt.Printf("Error getting child PIDs: %v\n", err)
			return
		}

		fmt.Printf("Child PIDs: %+v\n", childPIDs)

		if err := cmd.Wait(); err != nil {
			log.Printf("cmd.Wait: %v", err)
		}

		for _, childPID := range childPIDs {
			log.Printf("Killing child PID %v", childPID)
			if err := killProcess(childPID); err != nil {
				log.Printf("Failed to kill child PID %v: %v", childPID, err)
			}
		}
	}()

	time.Sleep(20 * time.Second)
	// Kill the Modus CLI
	log.Printf("Terminating Modus CLI")
	must(cmd.Cancel())
	// if cmd.Process != nil {
	// 	if err := cmd.Process.Kill(); err != nil {
	// 		log.Printf("Failed to kill Modus CLI: %v", err)
	// 	}
	// }
	cancel() // Cancel the context

	log.Printf("Waiting for Modus CLI to terminate and release its port")
	if !waitForPortToBeFree(8686, 20*time.Second) {
		log.Printf("Port 8686 is still in use after waiting")
	} else {
		log.Printf("Port 8686 is now free")
	}
}

func isPortInUse(port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%v", port), time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func waitForPortToBeFree(port int, timeout time.Duration) bool {
	start := time.Now()
	for time.Since(start) < timeout {
		if !isPortInUse(port) {
			return true
		}
		time.Sleep(500 * time.Millisecond)
	}
	return false
}

// getChildPIDs returns the PIDs of all child processes of the given parent PID
func getChildPIDs(parentPID int) ([]int, error) {
	out, err := exec.Command("pgrep", "-P", strconv.Itoa(parentPID)).Output()
	if err != nil {
		return nil, fmt.Errorf("pgrep failed: %v", err)
	}

	// Parse the output of pgrep
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	childPIDs := make([]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		pid, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PID: %v", err)
		}
		childPIDs = append(childPIDs, pid)
	}

	return childPIDs, nil
}

// killProcess kills a process by PID
func killProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to find process: %v", err)
	}

	// Send a SIGKILL signal to the process
	err = process.Kill()
	if err != nil {
		return fmt.Errorf("failed to kill process: %v", err)
	}

	return nil
}

func must(arg0 any, args ...any) {
	switch t := arg0.(type) {
	case error:
		log.Fatal(t)
	case string:
		log.Fatalf(t, args...)
	}
}
