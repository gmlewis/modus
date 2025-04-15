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

	typ, _, _ := stripErrorAndOption(ti.Name())

	switch typ {
	case "Bool":
		return newPrimitiveHandler[bool](ti), nil
		// https://docs.moonbitlang.com/en/latest/language/fundamentals.html#number
	case "Int16": // 16-bit signed integer, e.g. `(42 : Int16)`
		return newPrimitiveHandler[int16](ti), nil
	case "Int", "Unit": // 32-bit signed integer, e.g. `42` (or `Unit!Error`)
		return newPrimitiveHandler[int32](ti), nil
	case "Int64": // 64-bit signed integer, e.g. `1000L`
		return newPrimitiveHandler[int64](ti), nil
	case "UInt16": // 16-bit unsigned integer, e.g. `(14 : UInt16)`
		return newPrimitiveHandler[uint16](ti), nil
	case "UInt": // 32-bit unsigned integer, e.g. `14U`
		return newPrimitiveHandler[uint32](ti), nil
	case "UInt64": // 64-bit unsigned integer, e.g. `14UL`
		return newPrimitiveHandler[uint64](ti), nil
	case "Double": // 64-bit floating point, defined by IEEE754, e.g. `3.14`
		return newPrimitiveHandler[float64](ti), nil
	case "Float": // 32-bit floating point, defined by IEEE754, e.g. `(3.14 : Float)`
		return newPrimitiveHandler[float32](ti), nil
	case "Char": // represents a Unicode code point, e.g. `'a'`, `'\x41'`, `'\u{30}'`, `'\u03B1'`,
		return newPrimitiveHandler[int16](ti), nil
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
	val, ok := h.converter.Read(wa.Memory(), offset)
	if !ok {
		return 0, fmt.Errorf("failed to read %s from memory", h.typeInfo.Name())
	}

	if h.typeInfo.IsPointer() {
		return &val, nil
	}

	return val, nil
}

func (h *primitiveHandler[T]) Write(ctx context.Context, wasmadapter langsupport.WasmAdapter, offset uint32, obj any) (utils.Cleaner, error) {
	wa, ok := wasmadapter.(wasmMemoryWriter)
	if !ok {
		return nil, fmt.Errorf("expected a wasmMemoryWriter, got %T", wasmadapter)
	}

	if utils.HasNil(obj) {
		// TODO: Improve this to avoid string comparison.
		switch {
		case h.typeInfo.IsBoolean(),
			h.typeInfo.Name() == "Byte?",
			h.typeInfo.Name() == "Char?",
			h.typeInfo.Name() == "Int16?",
			h.typeInfo.Name() == "UInt16?":
			if ok := wa.Memory().WriteUint32Le(offset, 0xffffffff); !ok {
				return nil, fmt.Errorf("failed to write %v 'None' to memory offset %v", h.typeInfo.Name(), offset)
			}
			return nil, nil
		case h.typeInfo.Name() == "Int?", h.typeInfo.Name() == "UInt?":
			if ok := wa.Memory().WriteUint64Le(offset, 0x100000000); !ok {
				return nil, fmt.Errorf("failed to write %v 'None' to memory offset %v", h.typeInfo.Name(), offset)
			}
			return nil, nil
		case h.typeInfo.Name() == "Int64?", h.typeInfo.Name() == "UInt64?",
			h.typeInfo.Name() == "Float?", h.typeInfo.Name() == "Double?":
			var res []uint64
			var err error
			if t, ok := wasmadapter.(*wasmAdapter); ok {
				res, err = t.fnPtrToNone.Call(ctx)
				if err != nil {
					return nil, fmt.Errorf("failed to call ptr_to_none for type %v to memory offset %v", h.typeInfo.Name(), offset)
				}
			} else { // used only during unit testing since we can't call the plugin here.
				noneBlock, c, err := wa.allocateAndPinMemory(ctx, 1, 0) // cannot allocate 0 bytes
				if err != nil {
					return c, err
				}
				// innerCln.AddCleaner(c)
				wa.Memory().Write(noneBlock-8, []byte{255, 255, 255, 255, 0, 0, 0, 0}) // None in memBlock header
				res = []uint64{uint64(noneBlock - 8)}
			}
			if ok := wa.Memory().WriteUint32Le(offset, uint32(res[0])); !ok {
				return nil, fmt.Errorf("failed to write %v 'None' to memory offset %v", h.typeInfo.Name(), offset)
			}
			return nil, nil
		default:
			return nil, fmt.Errorf("unhandled write nil primitive '%v'", h.typeInfo.Name())
		}
	}

	val, err := utils.Cast[T](obj)
	if err != nil {
		return nil, err
	}

	if h.typeInfo.IsPointer() {
		elementSize := h.typeInfo.Size() // BUG?!?  Int64?/UInt64? are 8 bytes, not 4
		if h.typeInfo.Name() == "Int64?" || h.typeInfo.Name() == "UInt64?" {
			elementSize = 8
		} else if elementSize > 4 {
			// Note that the element size itself may be larger than 4, but we are writing a pointer to it, so make it 4.
			elementSize = 4
		}
		switch h.typeInfo.Name() {
		case "Int64?",
			"UInt64?",
			"Float?",
			"Double?":
			// TODO: Would it be easier/safer/better to write helper functions and never call allocateAndPinMemory directly?
			ptr, cln, err := wa.allocateAndPinMemory(ctx, elementSize, OptionBlockType)
			if err != nil {
				return cln, err
			}
			if ok := h.converter.Write(wa.Memory(), ptr, val); !ok {
				return cln, fmt.Errorf("failed to write %s to memory", h.typeInfo.Name())
			}
			if ok := wa.Memory().WriteUint32Le(offset, ptr-8); !ok {
				return cln, fmt.Errorf("failed to write %v 'None' to memory offset %v = %v", h.typeInfo.Name(), offset, ptr-8)
			}
			return cln, nil
		}
		// For all others, first write the data as usual, but then write the Some/None for pointers afterward
	}

	if ok := h.converter.Write(wa.Memory(), offset, val); !ok {
		return nil, fmt.Errorf("failed to write %s to memory", h.typeInfo.Name())
	}

	if h.typeInfo.IsPointer() {
		switch h.typeInfo.Name() {
		case "Byte?", "Bool?", "Char?", "String?": // "UInt?" in MoonBit for some reason uses sign extending below.
			wa.Memory().Write(offset+4, []byte{0, 0, 0, 0})
		case "Int16?":
			highByte, ok := wa.Memory().ReadByte(offset + 1)
			if ok && highByte&0x80 != 0 {
				wa.Memory().Write(offset+2, []byte{255, 255}) // Some(negative number)
			} else {
				wa.Memory().Write(offset+2, []byte{0, 0}) // Some(positive number)
			}
		case "UInt16?":
			wa.Memory().Write(offset+2, []byte{0, 0}) // Some(positive number)
		case "Int?", "UInt?": // I don't know why uses sign extending for "UInt?", but seems to.
			highByte, ok := wa.Memory().ReadByte(offset + 3)
			if ok && highByte&0x80 != 0 {
				wa.Memory().Write(offset+4, []byte{255, 255, 255, 255}) // Some(negative number)
			} else {
				wa.Memory().Write(offset+4, []byte{0, 0, 0, 0}) // Some(positive number)
			}
		default:
			return nil, fmt.Errorf("unsupported nullable primitive type: %v", h.typeInfo.Name())
		}
	}

	return nil, nil
}

// Unlike all the other handlers, the primitiveHandler[T].Decode method
// should always be passed an actual value, not a pointer to a value except for
// the four cases that don't fit in a single 64-bit value: Int64?, UInt64?, Float?, Double?.
func (h *primitiveHandler[T]) Decode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, vals []uint64) (any, error) {
	wa, ok := wasmAdapter.(wasmMemoryReader)
	if !ok {
		return nil, fmt.Errorf("expected a wasmMemoryReader, got %T", wasmAdapter)
	}

	if len(vals) != 1 {
		return nil, fmt.Errorf("expected 1 value, got %v", len(vals))
	}

	if h.TypeInfo().IsNullable() && h.typeInfo.IsPointer() {
		// TODO: Improve this to avoid string comparison.
		switch {
		case h.typeInfo.IsBoolean() && vals[0] == 0xffffffff:
			return nil, nil
		case h.typeInfo.Name() == "Byte?" && vals[0] == 0xffffffff:
			return nil, nil
		// Char?==None==0xffffffff but Array[Char?]==[None]==0x100000000
		case h.typeInfo.Name() == "Char?" && vals[0] >= 0xffffffff:
			return nil, nil
		// Int16?==None==0xffffffff but Array[Int16?]==[None]==0x100000000
		case h.typeInfo.Name() == "Int16?" && vals[0] >= 0xffffffff:
			return nil, nil
		// UInt16?==None==0xffffffff but Array[UInt16?]==[None]==0x100000000
		case h.typeInfo.Name() == "UInt16?" && vals[0] >= 0xffffffff:
			return nil, nil
		case h.typeInfo.Name() == "Int?" && vals[0] == 0x100000000:
			return nil, nil
		case h.typeInfo.Name() == "UInt?" && vals[0] == 0x100000000:
			return nil, nil
		case h.typeInfo.Name() == "Int64?", h.typeInfo.Name() == "UInt64?",
			h.typeInfo.Name() == "Float?", h.typeInfo.Name() == "Double?":
			memBlockHeader, ok := wa.Memory().ReadUint64Le(uint32(vals[0]))
			if !ok {
				return nil, fmt.Errorf("failed to read pointer %v from memory", vals[0]+8)
			}
			if memBlockHeader == 0x00000000ffffffff { // constant 'None' in MoonBit
				return nil, nil
			}
			result, ok := h.converter.Read(wa.Memory(), uint32(vals[0]+8))
			if !ok {
				return nil, fmt.Errorf("failed to read pointer %v from memory", vals[0]+8)
			}
			return &result, nil
		}
	}

	result := h.converter.Decode(vals[0])
	if h.typeInfo.IsPointer() {
		return &result, nil
	}
	return result, nil
}

func (h *primitiveHandler[T]) Encode(ctx context.Context, wasmAdapter langsupport.WasmAdapter, obj any) ([]uint64, utils.Cleaner, error) {
	wa, ok := wasmAdapter.(wasmMemoryWriter)
	if !ok {
		return nil, nil, fmt.Errorf("expected a wasmMemoryWriter, got %T", wasmAdapter)
	}

	val, err := utils.Cast[T](obj)
	if err != nil {
		return nil, nil, err
	}

	result := h.converter.Encode(val)

	if h.TypeInfo().IsNullable() && h.typeInfo.IsPointer() {
		// TODO: Improve this to avoid string comparison.
		if obj == nil {
			switch {
			case h.typeInfo.IsBoolean(),
				h.typeInfo.Name() == "Byte?":
				return []uint64{0xffffffff}, nil, nil
			case h.typeInfo.Name() == "Char?",
				h.typeInfo.Name() == "Int?",
				h.typeInfo.Name() == "Int16?",
				h.typeInfo.Name() == "UInt?",
				h.typeInfo.Name() == "UInt16?":
				return []uint64{0x100000000}, nil, nil
			case h.typeInfo.Name() == "Int64?", h.typeInfo.Name() == "UInt64?",
				h.typeInfo.Name() == "Float?", h.typeInfo.Name() == "Double?":
				ptr, cln, err := wa.allocateAndPinMemory(ctx, 1, 0) // cannot allocate 0 bytes
				if err != nil {
					return nil, cln, err
				}
				wa.Memory().WriteUint64Le(ptr-8, 0x00000000ffffffff) // None
				return []uint64{uint64(ptr - 8)}, cln, nil
			}
		}

		switch {
		case h.typeInfo.Name() == "Int64?", h.typeInfo.Name() == "UInt64?":
			ptr, cln, err := wa.allocateAndPinMemory(ctx, 8, 1)
			if err != nil {
				return nil, cln, err
			}
			wa.Memory().WriteUint64Le(ptr, result)
			return []uint64{uint64(ptr - 8)}, cln, nil
		case h.typeInfo.Name() == "Float?":
			ptr, cln, err := wa.allocateAndPinMemory(ctx, 4, 1)
			if err != nil {
				return nil, cln, err
			}
			wa.Memory().WriteUint32Le(ptr, uint32(result))
			return []uint64{uint64(ptr - 8)}, cln, nil
		case h.typeInfo.Name() == "Double?":
			ptr, cln, err := wa.allocateAndPinMemory(ctx, 8, 1)
			if err != nil {
				return nil, cln, err
			}
			wa.Memory().WriteUint64Le(ptr, result)
			return []uint64{uint64(ptr - 8)}, cln, nil
		}
	}

	return []uint64{result}, nil, nil
}
