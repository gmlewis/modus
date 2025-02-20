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
)

func TestPrimitivesEncodeDecode_Bool(t *testing.T) {
	boolHandler := newPrimitiveHandler[bool](mustGetTypeInfo(t, "Bool"))
	boolOptionHandler := newPrimitiveHandler[bool](mustGetTypeInfo(t, "Bool?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[bool]
		value   any
	}{
		{
			name:    "Bool: false",
			handler: boolHandler,
			value:   false,
		},
		{
			name:    "Bool: true",
			handler: boolHandler,
			value:   true,
		},
		{
			name:    "Bool?: None",
			handler: boolOptionHandler,
			value:   nil,
		},
		{
			name:    "Bool?: Some(false)",
			handler: boolOptionHandler,
			value:   Ptr(false),
		},
		{
			name:    "Bool?: Some(true)",
			handler: boolOptionHandler,
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
	byteHandler := newPrimitiveHandler[uint8](mustGetTypeInfo(t, "Byte"))
	byteOptionHandler := newPrimitiveHandler[uint8](mustGetTypeInfo(t, "Byte?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[uint8]
		value   any
	}{
		{
			name:    "Byte: 0",
			handler: byteHandler,
			value:   byte(0),
		},
		{
			name:    "Byte: 255",
			handler: byteHandler,
			value:   byte(255),
		},
		{
			name:    "Byte?: None",
			handler: byteOptionHandler,
			value:   nil,
		},
		{
			name:    "Byte?: Some(0)",
			handler: byteOptionHandler,
			value:   Ptr(byte(0)),
		},
		{
			name:    "Byte?: Some(255)",
			handler: byteOptionHandler,
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
	charHandler := newPrimitiveHandler[int16](mustGetTypeInfo(t, "Char"))
	charOptionHandler := newPrimitiveHandler[int16](mustGetTypeInfo(t, "Char?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[int16]
		value   any
	}{
		{
			name:    "Char: -32768",
			handler: charHandler,
			value:   int16(-32768),
		},
		{
			name:    "Char: 0",
			handler: charHandler,
			value:   int16(0),
		},
		{
			name:    "Char: 32767",
			handler: charHandler,
			value:   int16(32767),
		},
		{
			name:    "Char?: None",
			handler: charOptionHandler,
			value:   nil,
		},
		{
			name:    "Char?: Some(-32768)",
			handler: charOptionHandler,
			value:   Ptr(int16(-32768)),
		},
		{
			name:    "Char?: Some(0)",
			handler: charOptionHandler,
			value:   Ptr(int16(0)),
		},
		{
			name:    "Char?: Some(32767)",
			handler: charOptionHandler,
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
	int16Handler := newPrimitiveHandler[int16](mustGetTypeInfo(t, "Int16"))
	int16OptionHandler := newPrimitiveHandler[int16](mustGetTypeInfo(t, "Int16?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[int16]
		value   any
	}{
		{
			name:    "Int16: -32768",
			handler: int16Handler,
			value:   int16(-32768),
		},
		{
			name:    "Int16: 0",
			handler: int16Handler,
			value:   int16(0),
		},
		{
			name:    "Int16: 32767",
			handler: int16Handler,
			value:   int16(32767),
		},
		{
			name:    "Int16?: None",
			handler: int16OptionHandler,
			value:   nil,
		},
		{
			name:    "Int16?: Some(-32768)",
			handler: int16OptionHandler,
			value:   Ptr(int16(-32768)),
		},
		{
			name:    "Int16?: Some(0)",
			handler: int16OptionHandler,
			value:   Ptr(int16(0)),
		},
		{
			name:    "Int16?: Some(32767)",
			handler: int16OptionHandler,
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
	uint16Handler := newPrimitiveHandler[uint16](mustGetTypeInfo(t, "UInt16"))
	uint16OptionHandler := newPrimitiveHandler[uint16](mustGetTypeInfo(t, "UInt16?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[uint16]
		value   any
	}{
		{
			name:    "UInt16: 0",
			handler: uint16Handler,
			value:   uint16(0),
		},
		{
			name:    "UInt16: 65535",
			handler: uint16Handler,
			value:   uint16(65535),
		},
		{
			name:    "UInt16?: None",
			handler: uint16OptionHandler,
			value:   nil,
		},
		{
			name:    "UInt16?: Some(0)",
			handler: uint16OptionHandler,
			value:   Ptr(uint16(0)),
		},
		{
			name:    "UInt16?: Some(65535)",
			handler: uint16OptionHandler,
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
	intHandler := newPrimitiveHandler[int](mustGetTypeInfo(t, "Int"))
	intOptionHandler := newPrimitiveHandler[int](mustGetTypeInfo(t, "Int?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[int]
		value   any
	}{
		{
			name:    "Int: -2147483648",
			handler: intHandler,
			value:   int(-2147483648),
		},
		{
			name:    "Int: 0",
			handler: intHandler,
			value:   int(0),
		},
		{
			name:    "Int: 2147483647",
			handler: intHandler,
			value:   int(2147483647),
		},
		{
			name:    "Int?: None",
			handler: intOptionHandler,
			value:   nil,
		},
		{
			name:    "Int?: Some(-2147483648)",
			handler: intOptionHandler,
			value:   Ptr(int(-2147483648)),
		},
		{
			name:    "Int?: Some(0)",
			handler: intOptionHandler,
			value:   Ptr(int(0)),
		},
		{
			name:    "Int?: Some(2147483647)",
			handler: intOptionHandler,
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
	uintHandler := newPrimitiveHandler[uint](mustGetTypeInfo(t, "UInt"))
	uintOptionHandler := newPrimitiveHandler[uint](mustGetTypeInfo(t, "UInt?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[uint]
		value   any
	}{
		{
			name:    "UInt: 0",
			handler: uintHandler,
			value:   uint(0),
		},
		{
			name:    "UInt: 4294967295",
			handler: uintHandler,
			value:   uint(4294967295),
		},
		{
			name:    "UInt?: None",
			handler: uintOptionHandler,
			value:   nil,
		},
		{
			name:    "UInt?: Some(0)",
			handler: uintOptionHandler,
			value:   Ptr(uint(0)),
		},
		{
			name:    "UInt?: Some(4294967295)",
			handler: uintOptionHandler,
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
	int64Handler := newPrimitiveHandler[int64](mustGetTypeInfo(t, "Int64"))
	int64OptionHandler := newPrimitiveHandler[int64](mustGetTypeInfo(t, "Int64?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[int64]
		value   any
	}{
		{
			name:    "Int64: -9223372036854775808",
			handler: int64Handler,
			value:   int64(-9223372036854775808),
		},
		{
			name:    "Int64: 0",
			handler: int64Handler,
			value:   int64(0),
		},
		{
			name:    "Int64: 9223372036854775807",
			handler: int64Handler,
			value:   int64(9223372036854775807),
		},
		{
			name:    "Int64?: None",
			handler: int64OptionHandler,
			value:   nil,
		},
		{
			name:    "Int64?: Some(-9223372036854775808)",
			handler: int64OptionHandler,
			value:   Ptr(int64(-9223372036854775808)),
		},
		{
			name:    "Int64?: Some(0)",
			handler: int64OptionHandler,
			value:   Ptr(int64(0)),
		},
		{
			name:    "Int64?: Some(9223372036854775807)",
			handler: int64OptionHandler,
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
	uint64Handler := newPrimitiveHandler[uint64](mustGetTypeInfo(t, "UInt64"))
	uint64OptionHandler := newPrimitiveHandler[uint64](mustGetTypeInfo(t, "UInt64?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[uint64]
		value   any
	}{
		{
			name:    "UInt64: 0",
			handler: uint64Handler,
			value:   uint64(0),
		},
		{
			name:    "UInt64: 18446744073709551615",
			handler: uint64Handler,
			value:   uint64(18446744073709551615),
		},
		{
			name:    "UInt64?: None",
			handler: uint64OptionHandler,
			value:   nil,
		},
		{
			name:    "UInt64?: Some(0)",
			handler: uint64OptionHandler,
			value:   Ptr(uint64(0)),
		},
		{
			name:    "UInt64?: Some(18446744073709551615)",
			handler: uint64OptionHandler,
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
	floatHandler := newPrimitiveHandler[float32](mustGetTypeInfo(t, "Float"))
	floatOptionHandler := newPrimitiveHandler[float32](mustGetTypeInfo(t, "Float?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[float32]
		value   any
	}{
		{
			name:    "Float: math.SmallestNonzeroFloat32",
			handler: floatHandler,
			value:   float32(math.SmallestNonzeroFloat32),
		},
		{
			name:    "Float: 0",
			handler: floatHandler,
			value:   float32(0),
		},
		{
			name:    "Float: math.MaxFloat32",
			handler: floatHandler,
			value:   float32(math.MaxFloat32),
		},
		{
			name:    "Float?: None",
			handler: floatOptionHandler,
			value:   nil,
		},
		{
			name:    "Float?: Some(math.SmallestNonzeroFloat32)",
			handler: floatOptionHandler,
			value:   Ptr(float32(math.SmallestNonzeroFloat32)),
		},
		{
			name:    "Float?: Some(0)",
			handler: floatOptionHandler,
			value:   Ptr(float32(0)),
		},
		{
			name:    "Float?: Some(math.MaxFloat32)",
			handler: floatOptionHandler,
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
	doubleHandler := newPrimitiveHandler[float64](mustGetTypeInfo(t, "Double"))
	doubleOptionHandler := newPrimitiveHandler[float64](mustGetTypeInfo(t, "Double?"))
	tests := []struct {
		name    string
		handler *primitiveHandler[float64]
		value   any
	}{
		{
			name:    "Double: math.SmallestNonzeroFloat64",
			handler: doubleHandler,
			value:   float64(math.SmallestNonzeroFloat64),
		},
		{
			name:    "Double: 0",
			handler: doubleHandler,
			value:   float64(0),
		},
		{
			name:    "Double: math.MaxFloat64",
			handler: doubleHandler,
			value:   float64(math.MaxFloat64),
		},
		{
			name:    "Double?: None",
			handler: doubleOptionHandler,
			value:   nil,
		},
		{
			name:    "Double?: Some(math.SmallestNonzeroFloat64)",
			handler: doubleOptionHandler,
			value:   Ptr(float64(math.SmallestNonzeroFloat64)),
		},
		{
			name:    "Double?: Some(0)",
			handler: doubleOptionHandler,
			value:   Ptr(float64(0)),
		},
		{
			name:    "Double?: Some(math.MaxFloat64)",
			handler: doubleOptionHandler,
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
