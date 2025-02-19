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
	"math"
	"reflect"
	"testing"

	"github.com/gmlewis/modus/runtime/langsupport"
)

func mustGetTypeInfo(t *testing.T, name string) langsupport.TypeInfo {
	t.Helper()
	cache := map[string]langsupport.TypeInfo{}
	ti, err := GetTypeInfo(context.Background(), name, cache)
	if err != nil {
		t.Fatalf("GetTypeInfo(%q) returned nil", name)
	}
	return ti
}

func TestPrimitivesEncodeDecode_Bool(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[bool]
		value   any
	}{
		{
			name:    "Bool: false",
			handler: newPrimitiveHandler[bool](mustGetTypeInfo(t, "Bool")),
			value:   false,
		},
		{
			name:    "Bool: true",
			handler: newPrimitiveHandler[bool](mustGetTypeInfo(t, "Bool")),
			value:   true,
		},
		{
			name:    "Bool?: None",
			handler: newPrimitiveHandler[bool](mustGetTypeInfo(t, "Bool?")),
			value:   nil,
		},
		{
			name:    "Bool?: Some(false)",
			handler: newPrimitiveHandler[bool](mustGetTypeInfo(t, "Bool?")),
			value:   Ptr(false),
		},
		{
			name:    "Bool?: Some(true)",
			handler: newPrimitiveHandler[bool](mustGetTypeInfo(t, "Bool?")),
			value:   Ptr(true),
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

func TestPrimitivesEncodeDecode_Byte(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[uint8]
		value   any
	}{
		{
			name:    "Byte: 0",
			handler: newPrimitiveHandler[uint8](mustGetTypeInfo(t, "Byte")),
			value:   byte(0),
		},
		{
			name:    "Byte: 255",
			handler: newPrimitiveHandler[uint8](mustGetTypeInfo(t, "Byte")),
			value:   byte(255),
		},
		{
			name:    "Byte?: None",
			handler: newPrimitiveHandler[uint8](mustGetTypeInfo(t, "Byte?")),
			value:   nil,
		},
		{
			name:    "Byte?: Some(0)",
			handler: newPrimitiveHandler[uint8](mustGetTypeInfo(t, "Byte?")),
			value:   Ptr(byte(0)),
		},
		{
			name:    "Byte?: Some(255)",
			handler: newPrimitiveHandler[uint8](mustGetTypeInfo(t, "Byte?")),
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

func TestPrimitivesEncodeDecode_Char(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[int16]
		value   any
	}{
		{
			name:    "Char: -32768",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Char")),
			value:   int16(-32768),
		},
		{
			name:    "Char: 0",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Char")),
			value:   int16(0),
		},
		{
			name:    "Char: 32767",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Char")),
			value:   int16(32767),
		},
		{
			name:    "Char?: None",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Char?")),
			value:   nil,
		},
		{
			name:    "Char?: Some(-32768)",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Char?")),
			value:   Ptr(int16(-32768)),
		},
		{
			name:    "Char?: Some(0)",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Char?")),
			value:   Ptr(int16(0)),
		},
		{
			name:    "Char?: Some(32767)",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Char?")),
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

func TestPrimitivesEncodeDecode_Int16(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[int16]
		value   any
	}{
		{
			name:    "Int16: -32768",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Int16")),
			value:   int16(-32768),
		},
		{
			name:    "Int16: 0",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Int16")),
			value:   int16(0),
		},
		{
			name:    "Int16: 32767",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Int16")),
			value:   int16(32767),
		},
		{
			name:    "Int16?: None",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Int16?")),
			value:   nil,
		},
		{
			name:    "Int16?: Some(-32768)",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Int16?")),
			value:   Ptr(int16(-32768)),
		},
		{
			name:    "Int16?: Some(0)",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Int16?")),
			value:   Ptr(int16(0)),
		},
		{
			name:    "Int16?: Some(32767)",
			handler: newPrimitiveHandler[int16](mustGetTypeInfo(t, "Int16?")),
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

func TestPrimitivesEncodeDecode_UInt16(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[uint16]
		value   any
	}{
		{
			name:    "UInt16: 0",
			handler: newPrimitiveHandler[uint16](mustGetTypeInfo(t, "UInt16")),
			value:   uint16(0),
		},
		{
			name:    "UInt16: 65535",
			handler: newPrimitiveHandler[uint16](mustGetTypeInfo(t, "UInt16")),
			value:   uint16(65535),
		},
		{
			name:    "UInt16?: None",
			handler: newPrimitiveHandler[uint16](mustGetTypeInfo(t, "UInt16?")),
			value:   nil,
		},
		{
			name:    "UInt16?: Some(0)",
			handler: newPrimitiveHandler[uint16](mustGetTypeInfo(t, "UInt16?")),
			value:   Ptr(uint16(0)),
		},
		{
			name:    "UInt16?: Some(65535)",
			handler: newPrimitiveHandler[uint16](mustGetTypeInfo(t, "UInt16?")),
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

func TestPrimitivesEncodeDecode_Int(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[int]
		value   any
	}{
		{
			name:    "Int: -2147483648",
			handler: newPrimitiveHandler[int](mustGetTypeInfo(t, "Int")),
			value:   int(-2147483648),
		},
		{
			name:    "Int: 0",
			handler: newPrimitiveHandler[int](mustGetTypeInfo(t, "Int")),
			value:   int(0),
		},
		{
			name:    "Int: 2147483647",
			handler: newPrimitiveHandler[int](mustGetTypeInfo(t, "Int")),
			value:   int(2147483647),
		},
		{
			name:    "Int?: None",
			handler: newPrimitiveHandler[int](mustGetTypeInfo(t, "Int?")),
			value:   nil,
		},
		{
			name:    "Int?: Some(-2147483648)",
			handler: newPrimitiveHandler[int](mustGetTypeInfo(t, "Int?")),
			value:   Ptr(int(-2147483648)),
		},
		{
			name:    "Int?: Some(0)",
			handler: newPrimitiveHandler[int](mustGetTypeInfo(t, "Int?")),
			value:   Ptr(int(0)),
		},
		{
			name:    "Int?: Some(2147483647)",
			handler: newPrimitiveHandler[int](mustGetTypeInfo(t, "Int?")),
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

func TestPrimitivesEncodeDecode_UInt(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[uint]
		value   any
	}{
		{
			name:    "UInt: 0",
			handler: newPrimitiveHandler[uint](mustGetTypeInfo(t, "UInt")),
			value:   uint(0),
		},
		{
			name:    "UInt: 4294967295",
			handler: newPrimitiveHandler[uint](mustGetTypeInfo(t, "UInt")),
			value:   uint(4294967295),
		},
		{
			name:    "UInt?: None",
			handler: newPrimitiveHandler[uint](mustGetTypeInfo(t, "UInt?")),
			value:   nil,
		},
		{
			name:    "UInt?: Some(0)",
			handler: newPrimitiveHandler[uint](mustGetTypeInfo(t, "UInt?")),
			value:   Ptr(uint(0)),
		},
		{
			name:    "UInt?: Some(4294967295)",
			handler: newPrimitiveHandler[uint](mustGetTypeInfo(t, "UInt?")),
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

func TestPrimitivesEncodeDecode_Int64(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[int64]
		value   any
	}{
		{
			name:    "Int64: -9223372036854775808",
			handler: newPrimitiveHandler[int64](mustGetTypeInfo(t, "Int64")),
			value:   int64(-9223372036854775808),
		},
		{
			name:    "Int64: 0",
			handler: newPrimitiveHandler[int64](mustGetTypeInfo(t, "Int64")),
			value:   int64(0),
		},
		{
			name:    "Int64: 9223372036854775807",
			handler: newPrimitiveHandler[int64](mustGetTypeInfo(t, "Int64")),
			value:   int64(9223372036854775807),
		},
		{
			name:    "Int64?: None",
			handler: newPrimitiveHandler[int64](mustGetTypeInfo(t, "Int64?")),
			value:   nil,
		},
		{
			name:    "Int64?: Some(-9223372036854775808)",
			handler: newPrimitiveHandler[int64](mustGetTypeInfo(t, "Int64?")),
			value:   Ptr(int64(-9223372036854775808)),
		},
		{
			name:    "Int64?: Some(0)",
			handler: newPrimitiveHandler[int64](mustGetTypeInfo(t, "Int64?")),
			value:   Ptr(int64(0)),
		},
		{
			name:    "Int64?: Some(9223372036854775807)",
			handler: newPrimitiveHandler[int64](mustGetTypeInfo(t, "Int64?")),
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

func TestPrimitivesEncodeDecode_UInt64(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[uint64]
		value   any
	}{
		{
			name:    "UInt64: 0",
			handler: newPrimitiveHandler[uint64](mustGetTypeInfo(t, "UInt64")),
			value:   uint64(0),
		},
		{
			name:    "UInt64: 18446744073709551615",
			handler: newPrimitiveHandler[uint64](mustGetTypeInfo(t, "UInt64")),
			value:   uint64(18446744073709551615),
		},
		{
			name:    "UInt64?: None",
			handler: newPrimitiveHandler[uint64](mustGetTypeInfo(t, "UInt64?")),
			value:   nil,
		},
		{
			name:    "UInt64?: Some(0)",
			handler: newPrimitiveHandler[uint64](mustGetTypeInfo(t, "UInt64?")),
			value:   Ptr(uint64(0)),
		},
		{
			name:    "UInt64?: Some(18446744073709551615)",
			handler: newPrimitiveHandler[uint64](mustGetTypeInfo(t, "UInt64?")),
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

func TestPrimitivesEncodeDecode_Float(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[float32]
		value   any
	}{
		{
			name:    "Float: math.SmallestNonzeroFloat32",
			handler: newPrimitiveHandler[float32](mustGetTypeInfo(t, "Float")),
			value:   float32(math.SmallestNonzeroFloat32),
		},
		{
			name:    "Float: 0",
			handler: newPrimitiveHandler[float32](mustGetTypeInfo(t, "Float")),
			value:   float32(0),
		},
		{
			name:    "Float: math.MaxFloat32",
			handler: newPrimitiveHandler[float32](mustGetTypeInfo(t, "Float")),
			value:   float32(math.MaxFloat32),
		},
		{
			name:    "Float?: None",
			handler: newPrimitiveHandler[float32](mustGetTypeInfo(t, "Float?")),
			value:   nil,
		},
		{
			name:    "Float?: Some(math.SmallestNonzeroFloat32)",
			handler: newPrimitiveHandler[float32](mustGetTypeInfo(t, "Float?")),
			value:   Ptr(float32(math.SmallestNonzeroFloat32)),
		},
		{
			name:    "Float?: Some(0)",
			handler: newPrimitiveHandler[float32](mustGetTypeInfo(t, "Float?")),
			value:   Ptr(float32(0)),
		},
		{
			name:    "Float?: Some(math.MaxFloat32)",
			handler: newPrimitiveHandler[float32](mustGetTypeInfo(t, "Float?")),
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

func TestPrimitivesEncodeDecode_Double(t *testing.T) {
	tests := []struct {
		name    string
		handler *primitiveHandler[float64]
		value   any
	}{
		{
			name:    "Double: math.SmallestNonzeroFloat64",
			handler: newPrimitiveHandler[float64](mustGetTypeInfo(t, "Double")),
			value:   float64(math.SmallestNonzeroFloat64),
		},
		{
			name:    "Double: 0",
			handler: newPrimitiveHandler[float64](mustGetTypeInfo(t, "Double")),
			value:   float64(0),
		},
		{
			name:    "Double: math.MaxFloat64",
			handler: newPrimitiveHandler[float64](mustGetTypeInfo(t, "Double")),
			value:   float64(math.MaxFloat64),
		},
		{
			name:    "Double?: None",
			handler: newPrimitiveHandler[float64](mustGetTypeInfo(t, "Double?")),
			value:   nil,
		},
		{
			name:    "Double?: Some(math.SmallestNonzeroFloat64)",
			handler: newPrimitiveHandler[float64](mustGetTypeInfo(t, "Double?")),
			value:   Ptr(float64(math.SmallestNonzeroFloat64)),
		},
		{
			name:    "Double?: Some(0)",
			handler: newPrimitiveHandler[float64](mustGetTypeInfo(t, "Double?")),
			value:   Ptr(float64(0)),
		},
		{
			name:    "Double?: Some(math.MaxFloat64)",
			handler: newPrimitiveHandler[float64](mustGetTypeInfo(t, "Double?")),
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
