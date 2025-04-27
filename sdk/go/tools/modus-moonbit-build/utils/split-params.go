/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import "strings"

func SplitParamsWithBrackets(allArgs string) []string {
	var result []string
	var n int     // count bracket pairs
	var start int // start of current arg
	for i := 0; i < len(allArgs); i++ {
		switch allArgs[i] {
		case '[', '(':
			n++
		case ']', ')':
			n--
		case ',':
			if n == 0 {
				arg := strings.TrimSpace(allArgs[start:i])
				if arg != "" {
					result = append(result, arg)
				}
				start = i + 1
			}
		}
	}
	if start < len(allArgs) {
		arg := strings.TrimSpace(allArgs[start:])
		if arg != "" {
			result = append(result, arg)
		}
	}

	return result
}
