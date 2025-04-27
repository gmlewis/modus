/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package extractor

import "github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"

func moonBitReturnType(export *metadata.Function) string {
	if len(export.Results) == 0 {
		return "Unit"
	}
	if len(export.Results) == 1 {
		return export.Results[0].Type
	}
	return "Tuple"
}

var moonBitFnImports = metadata.FunctionMap{
	// @console - always provide this metadata so that error messages can be printed.
	"modus_system.logMessage": {
		Name: "modus_system.logMessage",
		Parameters: []*metadata.Parameter{
			{Name: "level", Type: "String"},
			{Name: "message", Type: "String"},
		},
	},
}
