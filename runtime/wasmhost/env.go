/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package wasmhost

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"unicode/utf16"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func instantiateEnvHostFunctions(ctx context.Context, r wazero.Runtime) error {

	// AssemblyScript date import uses:
	//
	//   @external("env", "Date.now")
	//   export function now(): f64;
	//
	// We use this instead of relying on the date override in the AssemblyScript wasi shim,
	// because the shim's approach leads to problems such as:
	//   ERROR TS2322: Type '~lib/date/Date' is not assignable to type '~lib/wasi_date/wasi_Date'.
	//
	dateNow := func(ctx context.Context, stack []uint64) {
		now := time.Now().UnixMilli()
		stack[0] = api.EncodeF64(float64(now))
	}

	_, err := r.NewHostModuleBuilder("env").
		NewFunctionBuilder().WithGoFunction(
		api.GoFunc(dateNow), nil, []api.ValueType{api.ValueTypeF64}).Export("Date.now").
		Instantiate(ctx)
	if err != nil {
		return err
	}

	// MoonBit has a function `println` that is used to print to the console.
	// It is implemented in the host environment with a function called `spectest.print_char`
	// that prints a single character at a time to stderr. However, we buffer it until
	// a newline is encountered, at which point we print the entire line to stderr
	// (along with an extra newline just for visibility).
	var printBuffer string
	var highSurrogate uint16 // Temporary storage for a high surrogate
	printChar := func(ctx context.Context, stack []uint64) {
		utf16CodeUnit := uint16(stack[0])

		// Check if the code unit is a high surrogate
		if utf16CodeUnit >= 0xD800 && utf16CodeUnit <= 0xDBFF {
			// Store the high surrogate and wait for the low surrogate
			highSurrogate = utf16CodeUnit
			return
		}

		// Check if the code unit is a low surrogate
		if utf16CodeUnit >= 0xDC00 && utf16CodeUnit <= 0xDFFF {
			// If we have a stored high surrogate, decode the surrogate pair
			if highSurrogate != 0 {
				// Decode the surrogate pair into a UTF-8 string
				decodedString := string(utf16.Decode([]uint16{highSurrogate, utf16CodeUnit}))
				printBuffer += decodedString
				highSurrogate = 0 // Reset the high surrogate
			} else {
				log.Printf("WARNING: Invalid UTF-16 low surrogate without a high surrogate: %v", utf16CodeUnit)
				printBuffer += string(rune(utf16CodeUnit))
			}
			return
		}
		// Regular UTF-16 code unit (not part of a surrogate pair)
		printBuffer += string(rune(utf16CodeUnit))

		if utf16CodeUnit == '\n' {
			fmt.Fprintf(os.Stderr, "println: %v\n", printBuffer) // yes, two newlines here for visibility.
			printBuffer = ""
		}
	}

	_, err = r.NewHostModuleBuilder("spectest").
		NewFunctionBuilder().WithGoFunction(
		api.GoFunc(printChar), []api.ValueType{api.ValueTypeI32}, nil).Export("print_char").
		Instantiate(ctx)

	return err
}
