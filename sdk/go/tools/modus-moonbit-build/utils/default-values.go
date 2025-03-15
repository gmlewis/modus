/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import (
	"log"
	"strconv"
)

// GetDefaultValue returns the default value for a given type.
func GetDefaultValue(paramType, paramValue string) (any, bool) {
	switch paramType {
	case "Bool":
		if paramValue == "true" {
			return true, true
		}
		if paramValue == "false" {
			return false, true
		}
	case "Int":
		if v, err := strconv.Atoi(paramValue); err == nil {
			return int32(v), true
		}
	case "String":
		return paramValue, true
	}
	log.Printf("WARNING: GetDefaultValue: type %v = %v not yet supported", paramType, paramValue)
	return nil, false
}
