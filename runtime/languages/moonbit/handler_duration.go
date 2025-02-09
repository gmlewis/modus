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
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
)

func (p *planner) NewDurationHandler(ti langsupport.TypeInfo) (langsupport.TypeHandler, error) {
	handler := &durationHandler{*NewTypeHandler(ti)}
	p.AddHandler(handler)
	return handler, nil
}

type durationHandler struct {
	typeHandler
}

func (h *durationHandler) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	gmlPrintf("GML: handler_duration.go: durationHandler.Read(offset: %v)", offset)

	if offset == 0 {
		return nil, nil
	}

	return h.Decode(ctx, wa, []uint64{uint64(offset)})
}

func (h *durationHandler) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	gmlPrintf("GML: handler_duration.go: durationHandler.Write(offset: %v, obj: %v)", offset, obj)
	res, _, err := h.Encode(ctx, wa, obj)
	if err != nil || len(res) != 1 {
		return nil, err
	}

	if !wa.Memory().WriteUint64Le(offset, res[0]) {
		return nil, errors.New("failed to write @time.Duration to WASM memory")
	}

	return nil, nil
}

func (h *durationHandler) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	gmlPrintf("GML: handler_duration.go: durationHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("MoonBit: expected 1 value when decoding duration but got %v: %+v", len(vals), vals)
	}

	// MoonBit is configured to return a @time.Duration from the moonbitlang/x/time package:
	// struct Duration {
	//   secs : Int64
	//   nanos : Int
	// } derive(Eq, Compare)
	memBlock, _, err := memoryBlockAtOffset(wa, uint32(vals[0]), true)
	if err != nil {
		return nil, err
	}
	if len(memBlock) != 20 { // TODO
		return nil, fmt.Errorf("MoonBit: expected 20 bytes when decoding duration but got %v: %+v", len(memBlock), memBlock)
	}

	second := time.Duration(binary.LittleEndian.Uint32(memBlock[16:]))
	nanosecond := time.Duration(binary.LittleEndian.Uint32(memBlock[20:]))
	return time.Duration(second*time.Second + nanosecond), nil
}

// This converts a Go `time.Duration` to a MoonBit `@time.Duration`.
func (h *durationHandler) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	if obj == nil || utils.HasNil(obj) {
		return []uint64{0}, nil, nil
	}

	nanos, ok := obj.(int64)
	if !ok {
		gmlPrintf("GML: handler_duration.go: durationHandler.Encode(obj: %+v): unexpected type obj=%T", obj, obj)
		return []uint64{0}, nil, nil
	}

	gmlPrintf("GML: handler_duration.go: durationHandler.Encode(obj: %v): nanos=%v", obj, nanos)
	res, err := wa.(*wasmAdapter).durationFromNanos.Call(ctx, uint64(nanos))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert time.Duration to Duration: %w", err)
	}
	gmlPrintf("GML: handler_duration.go: durationHandler.Encode: res=%+v", res)
	if len(res) != 2 || res[0] == 0 {
		return nil, nil, fmt.Errorf("failed to convert time.Duration to Duration: %+v", res)
	}

	return res[1:], nil, nil
}
