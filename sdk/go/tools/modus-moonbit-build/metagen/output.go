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
	"io"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
)

func LogToConsole(meta *metadata.Metadata) {

	// FORCE_COLOR is set by Modus CLI
	forceColor := os.Getenv("FORCE_COLOR")
	if forceColor != "" && forceColor != "0" {
		color.NoColor = false
	}
	w := color.Output

	fmt.Fprintln(w)
	writeHeader(w, "Metadata:")
	writeTable(w, [][]string{
		{"Plugin Name", meta.Plugin},
		{"MoonBit Module", meta.Module},
		{"Modus SDK", meta.SDK},
		{"Build ID", meta.BuildID},
		{"Build Timestamp", meta.BuildTime},
		{"Git Repository", meta.GitRepo},
		{"Git Commit", meta.GitCommit},
	})
	fmt.Fprintln(w)

	if len(meta.FnExports) > 0 {
		writeHeader(w, "Functions:")
		for _, k := range meta.FnExports.SortedKeys() {
			fn := meta.FnExports[k]
			name := fn.String(meta)
			if strings.Contains(name, "__modus_") {
				continue
			}
			writeItem(w, name)
		}
		fmt.Fprintln(w)
	}

	types := make([]string, 0, len(meta.Types))
	for _, k := range meta.Types.SortedKeys(meta.Module) {
		t := meta.Types[k]
		// Local custom types no longer have the module prefix.
		if len(t.Fields) > 0 {
			types = append(types, k)
		}
	}

	if len(types) > 0 {
		alreadyWritten := map[string]bool{}
		writeHeader(w, "Custom Types:")
		for _, t := range types {
			typ, _, _ := utils.StripErrorAndOption(t)
			if alreadyWritten[typ] {
				continue
			}
			alreadyWritten[typ] = true
			s := meta.Types[t].String(meta)
			if strings.HasPrefix(s, "(") {
				continue // skip tuples
			}
			writeItem(w, s)
		}
		fmt.Fprintln(w)
	}

	if utils.IsDebugModeEnabled() {
		writeHeader(w, "Metadata JSON:")
		metaJson, _ := utils.JsonSerialize(meta, true)
		fmt.Fprintln(w, string(metaJson))
		fmt.Fprintln(w)
	}
}

func writeHeader(w io.Writer, text string) {
	color.Set(color.FgBlue, color.Bold)
	fmt.Fprintln(w, text)
	color.Unset()
}

func writeItem(w io.Writer, text string) {
	color.Set(color.FgCyan)
	fmt.Fprint(w, "  "+text)
	color.Unset()
	fmt.Fprintln(w)
}

func writeTable(w io.Writer, rows [][]string) {
	pad := make([]int, len(rows))
	for _, row := range rows {
		for i, cell := range row {
			if len(cell) > pad[i] {
				pad[i] = len(cell)
			}
		}
	}

	for _, row := range rows {
		if len(row) != 2 || len(row[0]) == 0 || len(row[1]) == 0 {
			continue
		}

		padding := strings.Repeat(" ", pad[0]-len(row[0]))

		fmt.Fprint(w, "  ")
		color.Set(color.FgCyan)
		fmt.Fprintf(w, "%s:%s ", row[0], padding)
		color.Set(color.FgBlue)
		fmt.Fprint(w, row[1])
		color.Unset()
		fmt.Fprintln(w)
	}
}
