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
	"time"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/langsupport/primitives"
	"github.com/gmlewis/modus/runtime/utils"

	"golang.org/x/exp/constraints"
)

type primitive interface {
	constraints.Integer | constraints.Float | ~bool
}

func (p *planner) NewPrimitiveHandler(ti langsupport.TypeInfo) (h langsupport.TypeHandler, err error) {
	defer func() {
		if err == nil {
			p.typeHandlers[ti.Name()] = h
		}
	}()

	typ, hasError, hasOption := stripErrorAndOption(ti.Name())
	gmlPrintf("GML: handler_primitives.go: NewPrimitiveHandler('%v'): '%v', hasError=%v, hasOption=%v", ti.Name(), typ, hasError, hasOption)

	switch typ {
	case "Bool":
		return newPrimitiveHandler[bool](ti), nil
		// https://docs.moonbitlang.com/en/latest/language/fundamentals.html#number
	case "Int16": // 16-bit signed integer, e.g. `(42 : Int16)`
		return newPrimitiveHandler[int16](ti), nil
	case "Int", "Unit": // 32-bit signed integer, e.g. `42` (or `Unit!Error`)
		return newPrimitiveHandler[int](ti), nil
	case "Int64": // 64-bit signed integer, e.g. `1000L`
		return newPrimitiveHandler[int64](ti), nil
	case "UInt16": // 16-bit unsigned integer, e.g. `(14 : UInt16)`
		return newPrimitiveHandler[uint16](ti), nil
	case "UInt": // 32-bit unsigned integer, e.g. `14U`
		return newPrimitiveHandler[uint](ti), nil
	case "UInt64": // 64-bit unsigned integer, e.g. `14UL`
		return newPrimitiveHandler[uint64](ti), nil
	case "Double": // 64-bit floating point, defined by IEEE754, e.g. `3.14`
		return newPrimitiveHandler[float64](ti), nil
	case "Float": // 32-bit floating point, defined by IEEE754, e.g. `(3.14 : Float)`
		return newPrimitiveHandler[float32](ti), nil
	case "Char": // represents a Unicode code point, e.g. `'a'`, `'\x41'`, `'\u{30}'`, `'\u03B1'`,
		return newPrimitiveHandler[uint16](ti), nil
	case "Byte": // either a single ASCII character, e.g. `b'a'`, `b'\xff'`
		return newPrimitiveHandler[uint8](ti), nil
	// case "BigInt": // represents numeric values larger than other types, e.g. `10000000000000000000000N`
	// case "String": // holds a sequence of UTF-16 code units, e.g. `"Hello, World!"`
	case "@time.Duration":
		return newPrimitiveHandler[time.Duration](ti), nil
	default:
		return nil, fmt.Errorf("unsupported primitive MoonBit type: %s", ti.Name())
	}
}

func newPrimitiveHandler[T primitive](ti langsupport.TypeInfo) *primitiveHandler[T] {
	return &primitiveHandler[T]{
		*NewTypeHandler(ti),
		primitives.NewPrimitiveTypeConverter[T](),
	}
}

type primitiveHandler[T primitive] struct {
	typeHandler
	converter primitives.TypeConverter[T]
}

func (h *primitiveHandler[T]) Read(ctx context.Context, wa langsupport.WasmAdapter, offset uint32) (any, error) {
	gmlPrintf("GML: handler_primitives.go: primitiveHandler[%T].Read(offset: %v)", []T{}, offset)

	val, ok := h.converter.Read(wa.Memory(), offset)
	if !ok {
		return 0, fmt.Errorf("failed to read %s from memory", h.typeInfo.Name())
	}

	return val, nil
}

func (h *primitiveHandler[T]) Write(ctx context.Context, wa langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	val, err := utils.Cast[T](obj)
	if err != nil {
		return nil, err
	}

	if ok := h.converter.Write(wa.Memory(), offset, val); !ok {
		return nil, fmt.Errorf("failed to write %s to memory", h.typeInfo.Name())
	}

	return nil, nil
}

func (h *primitiveHandler[T]) Decode(ctx context.Context, wa langsupport.WasmAdapter, vals []uint64) (any, error) {
	gmlPrintf("GML: handler_primitives.go: primitiveHandler.Decode(vals: %+v)", vals)

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value, got %d", len(vals))
	}

	result := h.converter.Decode(vals[0])
	if h.typeInfo.IsPointer() {
		return &result, nil
	}
	return result, nil
}

func (h *primitiveHandler[T]) Encode(ctx context.Context, wa langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	val, err := utils.Cast[T](obj)
	if err != nil {
		return nil, nil, err
	}

	return []uint64{h.converter.Encode(val)}, nil, nil
}
