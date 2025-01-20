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
	"os"
	"time"

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
	// a newline is encountered, at which point we print the entire line to stderr.
	var printBuffer string
	printChar := func(ctx context.Context, stack []uint64) {
		char := rune(stack[0]) // TODO: Does this need UTF16 to UTF8 conversion?
		printBuffer += string(char)
		if char == '\n' {
			fmt.Fprintf(os.Stderr, "println: %v\n", printBuffer)
			printBuffer = ""
		}
	}

	_, err = r.NewHostModuleBuilder("spectest").
		NewFunctionBuilder().WithGoFunction(
		api.GoFunc(printChar), []api.ValueType{api.ValueTypeI32}, nil).Export("print_char").
		Instantiate(ctx)

	return err
}
