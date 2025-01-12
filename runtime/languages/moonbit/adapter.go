/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package moonbit

import (
	"context"
	"fmt"

	"github.com/hypermodeinc/modus/runtime/langsupport"
	"github.com/hypermodeinc/modus/runtime/utils"

	wasm "github.com/tetratelabs/wazero/api"
)

func NewWasmAdapter(mod wasm.Module) langsupport.WasmAdapter {
	return &wasmAdapter{
		mod:         mod,
		visitedPtrs: make(map[uint32]int),
		// TODO(gmlewis)
	}
}

type wasmAdapter struct {
	mod         wasm.Module
	visitedPtrs map[uint32]int
	// TODO(gmlewis)
}

func (*wasmAdapter) TypeInfo() langsupport.LanguageTypeInfo {
	return _langTypeInfo
}

func (wa *wasmAdapter) Memory() wasm.Memory {
	return wa.mod.Memory()
}

func (wa *wasmAdapter) GetFunction(name string) wasm.Function {
	return wa.mod.ExportedFunction(name)
}

func (wa *wasmAdapter) PreInvoke(ctx context.Context, plan langsupport.ExecutionPlan) error {
	return nil
}

func (wa *wasmAdapter) AllocateMemory(ctx context.Context, size uint32) (uint32, utils.Cleaner, error) {
	// TODO(gmlewis)
	return 0, nil, fmt.Errorf("wasmAdapter.AllocateMemory not implemented yet for MoonBit")
}
