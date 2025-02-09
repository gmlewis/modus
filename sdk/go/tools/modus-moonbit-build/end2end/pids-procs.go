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
	"net"
	"os"
	"os/exec"
	"regexp"
	"runtime"
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

func waitForPortToBeInUse(port int, timeout time.Duration) bool {
	start := time.Now()
	for time.Since(start) < timeout {
		if isPortInUse(port) {
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

func getPIDUsingPort(port int) (int, error) {
	os := runtime.GOOS

	switch os {
	case "darwin", "linux": // macOS and Linux
		return getPIDUsingPortUnix(port)
	case "windows":
		return getPIDUsingPortWindows(port)
	default:
		return 0, fmt.Errorf("unsupported operating system: %s", os)
	}
}

// macOS and Linux
func getPIDUsingPortUnix(port int) (int, error) {
	args := []string{"-i", fmt.Sprintf(":%v", port)}
	gmlPrintf("getPIDUsingPortUnix: running command: lsof %v", strings.Join(args, " "))
	cmd := exec.Command("lsof", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("error executing lsof: %w, Output: %s", err, string(output))
	}
	gmlPrintf("getPIDUsingPortUnix: lsof output:\n%s", output)

	// output might contain multiple lines starting with 'p'. We only want the first one
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 2 || parts[0] != "modus_run" || parts[1] == "PID" {
			continue
		}
		pid, err := strconv.Atoi(parts[1])
		if err != nil {
			continue // ignore this line, try next one
		}
		return pid, nil
	}

	return 0, fmt.Errorf("no process found using port %v", port)
}

// Windows
func getPIDUsingPortWindows(port int) (int, error) {
	cmd := exec.Command("netstat", "-ano", "-p", "TCP")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("error executing netstat: %w, Output: %s", err, string(output))
	}

	re := regexp.MustCompile(fmt.Sprintf(`\s+TCP\s+\S+:%d\s+\S+\s+LISTENING\s+(\d+)\s*`, port))
	matches := re.FindStringSubmatch(string(output))

	if len(matches) < 2 {
		return 0, fmt.Errorf("no process found using port %d", port)
	}

	pid, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("failed to convert pid: %w", err)
	}

	return pid, nil
}

func killAllPIDs(pids []int) {
	for _, pid := range pids {
		if err := killProcess(pid); err != nil {
			gmlPrintf("WARNING: failed to kill process with pid %v: %v", pid, err)
		}
	}
}
