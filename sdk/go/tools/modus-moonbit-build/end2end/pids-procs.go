/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package end2end

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

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
