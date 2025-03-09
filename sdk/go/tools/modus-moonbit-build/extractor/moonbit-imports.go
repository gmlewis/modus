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
	"modus_neo4j_client.executeQuery": {
		Name: "modus_neo4j_client.executeQuery",
		Parameters: []*metadata.Parameter{
			{Name: "host_name", Type: "String"},
			{Name: "db_name", Type: "String"},
			{Name: "query", Type: "String"},
			{Name: "parameters_json", Type: "String"},
		},
		Results: []*metadata.Result{{Type: "@neo4j.EagerResult?"}},
	},
	"modus_system.getTimeInZone": {
		Name:       "modus_system.getTimeInZone",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"modus_system.getTimeZoneData": {
		Name: "modus_system.getTimeZoneData",
		Parameters: []*metadata.Parameter{
			{Name: "tz", Type: "String"},
			{Name: "format", Type: "String"},
		},
		Results: []*metadata.Result{{Type: "Array[Byte]"}},
	},
	"modus_system.logMessage": {
		Name: "modus_system.logMessage",
		Parameters: []*metadata.Parameter{
			{Name: "level", Type: "String"},
			{Name: "message", Type: "String"},
		},
	},
}
