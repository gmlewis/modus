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
	"reflect"
	"testing"

	"github.com/gmlewis/modus/lib/metadata"
)

func TestPrimitiveSlicesEncodeDecode_Bool(t *testing.T) {
	metadata := &metadata.Metadata{
		Types: map[string]*metadata.TypeDefinition{
			"Array[Bool]":  {Name: "Array[Bool]", Id: 4},
			"Array[Bool?]": {Name: "Array[Bool?]", Id: 5},
		},
	}
	boolSliceHandler := newPrimitiveSliceHandler[bool](mustGetTypeInfo(t, "Array[Bool]"), mustGetTypeDef(t, metadata, "Array[Bool]"))
	boolOptionSliceHandler := newPrimitiveSliceHandler[bool](mustGetTypeInfo(t, "Array[Bool?]"), mustGetTypeDef(t, metadata, "Array[Bool?]"))
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[bool]
		value   any
	}{
		{
			name:    "Array[Bool]: nil",
			handler: boolSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[Bool]: []",
			handler: boolSliceHandler,
			value:   []bool{},
		},
		{
			name:    "Array[Bool]: [false]",
			handler: boolSliceHandler,
			value:   []bool{false},
		},
		{
			name:    "Array[Bool]: [false, true]",
			handler: boolSliceHandler,
			value:   []bool{false, true},
		},
		{
			name:    "Array[Bool]: [false, true, false]",
			handler: boolSliceHandler,
			value:   []bool{false, true, false},
		},
		{
			name:    "Array[Bool]: [false, true, false, true]",
			handler: boolSliceHandler,
			value:   []bool{false, true, false, true},
		},
		{
			name:    "Array[Bool?]: nil",
			handler: boolOptionSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[Bool?]: []",
			handler: boolOptionSliceHandler,
			value:   []*bool{},
		},
		{
			name:    "Array[Bool?]: [false]",
			handler: boolOptionSliceHandler,
			value:   []*bool{Ptr(false)},
		},
		{
			name:    "Array[Bool?]: [true, true]",
			handler: boolOptionSliceHandler,
			value:   []*bool{Ptr(true), Ptr(true)},
		},
		{
			name:    "Array[Bool?]: [false, true, false]",
			handler: boolOptionSliceHandler,
			value:   []*bool{Ptr(false), Ptr(true), Ptr(false)},
		},
		{
			name:    "Array[Bool?]: [false, true, false, true]",
			handler: boolOptionSliceHandler,
			value:   []*bool{Ptr(false), Ptr(true), Ptr(false), Ptr(true)},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

/*
func TestPrimitiveSlicesEncodeDecode_Byte(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[uint8]
		value   any
	}{
		{
			name:    "Array[Byte]: 0",
			handler: newPrimitiveSliceHandler[uint8](mustGetTypeInfo(t, "Byte")),
			value:   byte(0),
		},
		{
			name:    "Array[Byte]: 255",
			handler: newPrimitiveSliceHandler[uint8](mustGetTypeInfo(t, "Byte")),
			value:   byte(255),
		},
		{
			name:    "Array[Byte]?: None",
			handler: newPrimitiveSliceHandler[uint8](mustGetTypeInfo(t, "Byte?")),
			value:   nil,
		},
		{
			name:    "Array[Byte]?: Some(0)",
			handler: newPrimitiveSliceHandler[uint8](mustGetTypeInfo(t, "Byte?")),
			value:   Ptr(byte(0)),
		},
		{
			name:    "Array[Byte]?: Some(255)",
			handler: newPrimitiveSliceHandler[uint8](mustGetTypeInfo(t, "Byte?")),
			value:   Ptr(byte(255)),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

func TestPrimitiveSlicesEncodeDecode_Char(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[int16]
		value   any
	}{
		{
			name:    "Array[Char]: -32768",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Char")),
			value:   int16(-32768),
		},
		{
			name:    "Array[Char]: 0",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Char")),
			value:   int16(0),
		},
		{
			name:    "Array[Char]: 32767",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Char")),
			value:   int16(32767),
		},
		{
			name:    "Array[Char]?: None",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Char?")),
			value:   nil,
		},
		{
			name:    "Array[Char]?: Some(-32768)",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Char?")),
			value:   Ptr(int16(-32768)),
		},
		{
			name:    "Array[Char]?: Some(0)",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Char?")),
			value:   Ptr(int16(0)),
		},
		{
			name:    "Array[Char]?: Some(32767)",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Char?")),
			value:   Ptr(int16(32767)),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

func TestPrimitiveSlicesEncodeDecode_Int16(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[int16]
		value   any
	}{
		{
			name:    "Array[Int16]: -32768",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Int16")),
			value:   int16(-32768),
		},
		{
			name:    "Array[Int16]: 0",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Int16")),
			value:   int16(0),
		},
		{
			name:    "Array[Int16]: 32767",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Int16")),
			value:   int16(32767),
		},
		{
			name:    "Array[Int16]?: None",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Int16?")),
			value:   nil,
		},
		{
			name:    "Array[Int16]?: Some(-32768)",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Int16?")),
			value:   Ptr(int16(-32768)),
		},
		{
			name:    "Array[Int16]?: Some(0)",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Int16?")),
			value:   Ptr(int16(0)),
		},
		{
			name:    "Array[Int16]?: Some(32767)",
			handler: newPrimitiveSliceHandler[int16](mustGetTypeInfo(t, "Int16?")),
			value:   Ptr(int16(32767)),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

func TestPrimitiveSlicesEncodeDecode_UInt16(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[uint16]
		value   any
	}{
		{
			name:    "Array[UInt16]: 0",
			handler: newPrimitiveSliceHandler[uint16](mustGetTypeInfo(t, "UInt16")),
			value:   uint16(0),
		},
		{
			name:    "Array[UInt16]: 65535",
			handler: newPrimitiveSliceHandler[uint16](mustGetTypeInfo(t, "UInt16")),
			value:   uint16(65535),
		},
		{
			name:    "Array[UInt16]?: None",
			handler: newPrimitiveSliceHandler[uint16](mustGetTypeInfo(t, "UInt16?")),
			value:   nil,
		},
		{
			name:    "Array[UInt16]?: Some(0)",
			handler: newPrimitiveSliceHandler[uint16](mustGetTypeInfo(t, "UInt16?")),
			value:   Ptr(uint16(0)),
		},
		{
			name:    "Array[UInt16]?: Some(65535)",
			handler: newPrimitiveSliceHandler[uint16](mustGetTypeInfo(t, "UInt16?")),
			value:   Ptr(uint16(65535)),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

func TestPrimitiveSlicesEncodeDecode_Int(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[int]
		value   any
	}{
		{
			name:    "Array[Int]: -2147483648",
			handler: newPrimitiveSliceHandler[int](mustGetTypeInfo(t, "Int")),
			value:   int(-2147483648),
		},
		{
			name:    "Array[Int]: 0",
			handler: newPrimitiveSliceHandler[int](mustGetTypeInfo(t, "Int")),
			value:   int(0),
		},
		{
			name:    "Array[Int]: 2147483647",
			handler: newPrimitiveSliceHandler[int](mustGetTypeInfo(t, "Int")),
			value:   int(2147483647),
		},
		{
			name:    "Array[Int]?: None",
			handler: newPrimitiveSliceHandler[int](mustGetTypeInfo(t, "Int?")),
			value:   nil,
		},
		{
			name:    "Array[Int]?: Some(-2147483648)",
			handler: newPrimitiveSliceHandler[int](mustGetTypeInfo(t, "Int?")),
			value:   Ptr(int(-2147483648)),
		},
		{
			name:    "Array[Int]?: Some(0)",
			handler: newPrimitiveSliceHandler[int](mustGetTypeInfo(t, "Int?")),
			value:   Ptr(int(0)),
		},
		{
			name:    "Array[Int]?: Some(2147483647)",
			handler: newPrimitiveSliceHandler[int](mustGetTypeInfo(t, "Int?")),
			value:   Ptr(int(2147483647)),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

func TestPrimitiveSlicesEncodeDecode_UInt(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[uint]
		value   any
	}{
		{
			name:    "Array[UInt]: 0",
			handler: newPrimitiveSliceHandler[uint](mustGetTypeInfo(t, "UInt")),
			value:   uint(0),
		},
		{
			name:    "Array[UInt]: 4294967295",
			handler: newPrimitiveSliceHandler[uint](mustGetTypeInfo(t, "UInt")),
			value:   uint(4294967295),
		},
		{
			name:    "Array[UInt]?: None",
			handler: newPrimitiveSliceHandler[uint](mustGetTypeInfo(t, "UInt?")),
			value:   nil,
		},
		{
			name:    "Array[UInt]?: Some(0)",
			handler: newPrimitiveSliceHandler[uint](mustGetTypeInfo(t, "UInt?")),
			value:   Ptr(uint(0)),
		},
		{
			name:    "Array[UInt]?: Some(4294967295)",
			handler: newPrimitiveSliceHandler[uint](mustGetTypeInfo(t, "UInt?")),
			value:   Ptr(uint(4294967295)),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

func TestPrimitiveSlicesEncodeDecode_Int64(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[int64]
		value   any
	}{
		{
			name:    "Array[Int64]: -9223372036854775808",
			handler: newPrimitiveSliceHandler[int64](mustGetTypeInfo(t, "Int64")),
			value:   int64(-9223372036854775808),
		},
		{
			name:    "Array[Int64]: 0",
			handler: newPrimitiveSliceHandler[int64](mustGetTypeInfo(t, "Int64")),
			value:   int64(0),
		},
		{
			name:    "Array[Int64]: 9223372036854775807",
			handler: newPrimitiveSliceHandler[int64](mustGetTypeInfo(t, "Int64")),
			value:   int64(9223372036854775807),
		},
		{
			name:    "Array[Int64]?: None",
			handler: newPrimitiveSliceHandler[int64](mustGetTypeInfo(t, "Int64?")),
			value:   nil,
		},
		{
			name:    "Array[Int64]?: Some(-9223372036854775808)",
			handler: newPrimitiveSliceHandler[int64](mustGetTypeInfo(t, "Int64?")),
			value:   Ptr(int64(-9223372036854775808)),
		},
		{
			name:    "Array[Int64]?: Some(0)",
			handler: newPrimitiveSliceHandler[int64](mustGetTypeInfo(t, "Int64?")),
			value:   Ptr(int64(0)),
		},
		{
			name:    "Array[Int64]?: Some(9223372036854775807)",
			handler: newPrimitiveSliceHandler[int64](mustGetTypeInfo(t, "Int64?")),
			value:   Ptr(int64(9223372036854775807)),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

func TestPrimitiveSlicesEncodeDecode_UInt64(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[uint64]
		value   any
	}{
		{
			name:    "Array[UInt64]: 0",
			handler: newPrimitiveSliceHandler[uint64](mustGetTypeInfo(t, "UInt64")),
			value:   uint64(0),
		},
		{
			name:    "Array[UInt64]: 18446744073709551615",
			handler: newPrimitiveSliceHandler[uint64](mustGetTypeInfo(t, "UInt64")),
			value:   uint64(18446744073709551615),
		},
		{
			name:    "Array[UInt64]?: None",
			handler: newPrimitiveSliceHandler[uint64](mustGetTypeInfo(t, "UInt64?")),
			value:   nil,
		},
		{
			name:    "Array[UInt64]?: Some(0)",
			handler: newPrimitiveSliceHandler[uint64](mustGetTypeInfo(t, "UInt64?")),
			value:   Ptr(uint64(0)),
		},
		{
			name:    "Array[UInt64]?: Some(18446744073709551615)",
			handler: newPrimitiveSliceHandler[uint64](mustGetTypeInfo(t, "UInt64?")),
			value:   Ptr(uint64(18446744073709551615)),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

func TestPrimitiveSlicesEncodeDecode_Float(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[float32]
		value   any
	}{
		{
			name:    "Array[Float]: math.SmallestNonzeroFloat32",
			handler: newPrimitiveSliceHandler[float32](mustGetTypeInfo(t, "Float")),
			value:   float32(math.SmallestNonzeroFloat32),
		},
		{
			name:    "Array[Float]: 0",
			handler: newPrimitiveSliceHandler[float32](mustGetTypeInfo(t, "Float")),
			value:   float32(0),
		},
		{
			name:    "Array[Float]: math.MaxFloat32",
			handler: newPrimitiveSliceHandler[float32](mustGetTypeInfo(t, "Float")),
			value:   float32(math.MaxFloat32),
		},
		{
			name:    "Array[Float]?: None",
			handler: newPrimitiveSliceHandler[float32](mustGetTypeInfo(t, "Float?")),
			value:   nil,
		},
		{
			name:    "Array[Float]?: Some(math.SmallestNonzeroFloat32)",
			handler: newPrimitiveSliceHandler[float32](mustGetTypeInfo(t, "Float?")),
			value:   Ptr(float32(math.SmallestNonzeroFloat32)),
		},
		{
			name:    "Array[Float]?: Some(0)",
			handler: newPrimitiveSliceHandler[float32](mustGetTypeInfo(t, "Float?")),
			value:   Ptr(float32(0)),
		},
		{
			name:    "Array[Float]?: Some(math.MaxFloat32)",
			handler: newPrimitiveSliceHandler[float32](mustGetTypeInfo(t, "Float?")),
			value:   Ptr(float32(math.MaxFloat32)),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

func TestPrimitiveSlicesEncodeDecode_Double(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveSliceHandler[float64]
		value   any
	}{
		{
			name:    "Array[Double]: math.SmallestNonzeroFloat64",
			handler: newPrimitiveSliceHandler[float64](mustGetTypeInfo(t, "Double")),
			value:   float64(math.SmallestNonzeroFloat64),
		},
		{
			name:    "Array[Double]: 0",
			handler: newPrimitiveSliceHandler[float64](mustGetTypeInfo(t, "Double")),
			value:   float64(0),
		},
		{
			name:    "Array[Double]: math.MaxFloat64",
			handler: newPrimitiveSliceHandler[float64](mustGetTypeInfo(t, "Double")),
			value:   float64(math.MaxFloat64),
		},
		{
			name:    "Array[Double]?: None",
			handler: newPrimitiveSliceHandler[float64](mustGetTypeInfo(t, "Double?")),
			value:   nil,
		},
		{
			name:    "Array[Double]?: Some(math.SmallestNonzeroFloat64)",
			handler: newPrimitiveSliceHandler[float64](mustGetTypeInfo(t, "Double?")),
			value:   Ptr(float64(math.SmallestNonzeroFloat64)),
		},
		{
			name:    "Array[Double]?: Some(0)",
			handler: newPrimitiveSliceHandler[float64](mustGetTypeInfo(t, "Double?")),
			value:   Ptr(float64(0)),
		},
		{
			name:    "Array[Double]?: Some(math.MaxFloat64)",
			handler: newPrimitiveSliceHandler[float64](mustGetTypeInfo(t, "Double?")),
			value:   Ptr(float64(math.MaxFloat64)),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}
			h := tt.handler
			res, _, err := h.encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.encode() returned an error: %v", err)
			}

			got, err := h.decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}
*/
