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
	"bytes"
	"context"
	"encoding/binary"
	"math"
	"reflect"
	"testing"

	"github.com/gmlewis/modus/runtime/langsupport"
)

func TestPrimitiveSlicesEncodeDecode_Bool(t *testing.T) {
	boolSliceHandler := mustGetHandler(t, "Array[Bool]")
	boolOptionSliceHandler := mustGetHandler(t, "Array[Bool?]")
	tests := []struct {
		name    string
		handler langsupport.TypeHandler
		value   any
		want    []byte
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
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[Bool]: [false]",
			handler: boolSliceHandler,
			value:   []bool{false},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Bool]: [true]",
			handler: boolSliceHandler,
			value:   []bool{true},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 1, 0, 0, 0},
		},
		{
			name:    "Array[Bool]: [false, true]",
			handler: boolSliceHandler,
			value:   []bool{false, true},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		},
		{
			name:    "Array[Bool]: [false, true, false]",
			handler: boolSliceHandler,
			value:   []bool{false, true, false},
			want:    []byte{1, 0, 0, 0, 241, 3, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Bool]: [false, true, false, true]",
			handler: boolSliceHandler,
			value:   []bool{false, true, false, true},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
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
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[Bool?]: [nil]",
			handler: boolOptionSliceHandler,
			value:   []*bool{nil},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 255, 255, 255, 255},
		},
		{
			name:    "Array[Bool?]: [false]",
			handler: boolOptionSliceHandler,
			value:   []*bool{Ptr(false)},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Bool?]: [true]",
			handler: boolOptionSliceHandler,
			value:   []*bool{Ptr(true)},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 1, 0, 0, 0},
		},
		{
			name:    "Array[Bool?]: [true, true]",
			handler: boolOptionSliceHandler,
			value:   []*bool{Ptr(true), Ptr(true)},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		},
		{
			name:    "Array[Bool?]: [false, true, false]",
			handler: boolOptionSliceHandler,
			value:   []*bool{Ptr(false), Ptr(true), Ptr(false)},
			want:    []byte{1, 0, 0, 0, 241, 3, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Bool?]: [false, true, false, true]",
			handler: boolOptionSliceHandler,
			value:   []*bool{Ptr(false), Ptr(true), Ptr(false), Ptr(true)},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
			}

			if res[0] != 0 {
				memBlock, _, _ := memoryBlockAtOffset(mockWA, uint32(res[0]), 0, true)
				sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
				numElements := binary.LittleEndian.Uint32(memBlock[12:16])
				var wantNumElements uint32
				switch slice := tt.value.(type) {
				case []bool:
					wantNumElements = uint32(len(slice))
				case []*bool:
					wantNumElements = uint32(len(slice))
				default:
					t.Fatalf("tt.value is not a slice: %T", tt.value)
				}
				if wantNumElements != numElements {
					t.Errorf("numElements = %v, want = %v", numElements, wantNumElements)
				}
				sliceMemBlock, _, _ := memoryBlockAtOffset(mockWA, sliceOffset, 0, true)
				if !bytes.Equal(sliceMemBlock, tt.want) {
					t.Errorf("\ngot  = %v\nwant = %v", sliceMemBlock, tt.want)
				}
			}

			got, err := h.Decode(ctx, mockWA, res)
			if err != nil {
				t.Fatalf("h.decode() returned an error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("encode/decode round trip conversion failed: got = %v, want = %v", got, tt.value)
			}
		})
	}
}

func TestPrimitiveSlicesEncodeDecode_Byte(t *testing.T) {
	byteSliceHandler := mustGetHandler(t, "Array[Byte]")
	byteOptionSliceHandler := mustGetHandler(t, "Array[Byte?]")
	tests := []struct {
		name    string
		handler langsupport.TypeHandler
		value   any
		want    []byte
	}{
		{
			name:    "Array[Byte]: nil",
			handler: byteSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[Byte]: []",
			handler: byteSliceHandler,
			value:   []byte{},
			want:    []byte{1, 0, 0, 0, 246, 0, 0, 0},
		},
		{
			name:    "Array[Byte]: [0]",
			handler: byteSliceHandler,
			value:   []byte{0},
			want:    []byte{1, 0, 0, 0, 246, 1, 0, 0, 0, 0, 0, 2},
		},
		{
			name:    "Array[Byte]: [255]",
			handler: byteSliceHandler,
			value:   []byte{255},
			want:    []byte{1, 0, 0, 0, 246, 1, 0, 0, 255, 0, 0, 2},
		},
		{
			name:    "Array[Byte]: [0, 255]",
			handler: byteSliceHandler,
			value:   []byte{0, 255},
			want:    []byte{1, 0, 0, 0, 246, 1, 0, 0, 0, 255, 0, 1},
		},
		{
			name:    "Array[Byte]: [0, 255, 0]",
			handler: byteSliceHandler,
			value:   []byte{0, 255, 0},
			want:    []byte{1, 0, 0, 0, 246, 1, 0, 0, 0, 255, 0, 0},
		},
		{
			name:    "Array[Byte]: [0, 255, 0, 255]",
			handler: byteSliceHandler,
			value:   []byte{0, 255, 0, 255},
			want:    []byte{1, 0, 0, 0, 246, 1, 0, 0, 0, 255, 0, 255},
		},
		{
			name:    "Array[Byte?]: nil",
			handler: byteOptionSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[Byte?]: []",
			handler: byteOptionSliceHandler,
			value:   []*byte{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[Byte?]: [nil]",
			handler: byteOptionSliceHandler,
			value:   []*byte{nil},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 255, 255, 255, 255},
		},
		{
			name:    "Array[Byte?]: [0]",
			handler: byteOptionSliceHandler,
			value:   []*byte{Ptr(byte(0))},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Byte?]: [255]",
			handler: byteOptionSliceHandler,
			value:   []*byte{Ptr(byte(255))},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 255, 0, 0, 0},
		},
		{
			name:    "Array[Byte?]: [255, 255]",
			handler: byteOptionSliceHandler,
			value:   []*byte{Ptr(byte(255)), Ptr(byte(255))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 255, 0, 0, 0, 255, 0, 0, 0},
		},
		{
			name:    "Array[Byte?]: [0, 255, 0]",
			handler: byteOptionSliceHandler,
			value:   []*byte{Ptr(byte(0)), Ptr(byte(255)), Ptr(byte(0))},
			want:    []byte{1, 0, 0, 0, 241, 3, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Byte?]: [0, 255, 0, 255]",
			handler: byteOptionSliceHandler,
			value:   []*byte{Ptr(byte(0)), Ptr(byte(255)), Ptr(byte(0)), Ptr(byte(255))},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0},
		},
		{
			name:    "Array[Byte?]: [nil, 128, nil, 128]",
			handler: byteOptionSliceHandler,
			value:   []*byte{nil, Ptr(byte(128)), nil, Ptr(byte(128))},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 255, 255, 255, 255, 128, 0, 0, 0, 255, 255, 255, 255, 128, 0, 0, 0},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
			}

			if res[0] != 0 {
				memBlock, _, _ := memoryBlockAtOffset(mockWA, uint32(res[0]), 0, true)
				sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
				numElements := binary.LittleEndian.Uint32(memBlock[12:16])
				var wantNumElements uint32
				switch slice := tt.value.(type) {
				case []byte:
					wantNumElements = uint32(len(slice))
				case []*byte:
					wantNumElements = uint32(len(slice))
				default:
					t.Fatalf("tt.value is not a slice: %T", tt.value)
				}
				if wantNumElements != numElements {
					t.Errorf("numElements = %v, want = %v", numElements, wantNumElements)
				}
				sliceMemBlock, _, _ := memoryBlockAtOffset(mockWA, sliceOffset, 0, true)
				if !bytes.Equal(sliceMemBlock, tt.want) {
					t.Errorf("\ngot  = %v\nwant = %v", sliceMemBlock, tt.want)
				}
			}

			got, err := h.Decode(ctx, mockWA, res)
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
	charSliceHandler := mustGetHandler(t, "Array[Char]")
	charOptionSliceHandler := mustGetHandler(t, "Array[Char?]")
	tests := []struct {
		name    string
		handler langsupport.TypeHandler
		value   any
		want    []byte
	}{
		{
			name:    "Array[Char]: nil",
			handler: charSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[Char]: []",
			handler: charSliceHandler,
			value:   []int16{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[Char]: [0]",
			handler: charSliceHandler,
			value:   []int16{0},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Char]: [255]",
			handler: charSliceHandler,
			value:   []int16{255},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 255, 0, 0, 0},
		},
		{
			name:    "Array[Char]: [0, 255]",
			handler: charSliceHandler,
			value:   []int16{0, 255},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Char]: [0, 255, 0]",
			handler: charSliceHandler,
			value:   []int16{0, 255, 0},
			want:    []byte{1, 0, 0, 0, 241, 3, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Char]: [0, 255, 0, 255]",
			handler: charSliceHandler,
			value:   []int16{0, 255, 0, 255},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 0, 0, 255, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Char?]: nil",
			handler: charOptionSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[Char?]: []",
			handler: charOptionSliceHandler,
			value:   []*int16{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[Char?]: [nil]",
			handler: charOptionSliceHandler,
			value:   []*int16{nil},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		},
		{
			name:    "Array[Char?]: [0]",
			handler: charOptionSliceHandler,
			value:   []*int16{Ptr(int16(0))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Char?]: [255]",
			handler: charOptionSliceHandler,
			value:   []*int16{Ptr(int16(255))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Char?]: [255, 255]",
			handler: charOptionSliceHandler,
			value:   []*int16{Ptr(int16(255)), Ptr(int16(255))},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Char?]: [0, 255, 0]",
			handler: charOptionSliceHandler,
			value:   []*int16{Ptr(int16(0)), Ptr(int16(255)), Ptr(int16(0))},
			want:    []byte{1, 0, 0, 0, 241, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Char?]: [0, 255, 0, 255]",
			handler: charOptionSliceHandler,
			value:   []*int16{Ptr(int16(0)), Ptr(int16(255)), Ptr(int16(0)), Ptr(int16(255))},
			want:    []byte{1, 0, 0, 0, 241, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Char?]: [nil, 128, nil, 128]",
			handler: charOptionSliceHandler,
			value:   []*int16{nil, Ptr(int16(128)), nil, Ptr(int16(128))},
			want:    []byte{1, 0, 0, 0, 241, 8, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
			}

			if res[0] != 0 {
				memBlock, _, _ := memoryBlockAtOffset(mockWA, uint32(res[0]), 0, true)
				sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
				numElements := binary.LittleEndian.Uint32(memBlock[12:16])
				var wantNumElements uint32
				switch slice := tt.value.(type) {
				case []int16:
					wantNumElements = uint32(len(slice))
				case []*int16:
					wantNumElements = uint32(len(slice))
				default:
					t.Fatalf("tt.value is not a slice: %T", tt.value)
				}
				if wantNumElements != numElements {
					t.Errorf("numElements = %v, want = %v", numElements, wantNumElements)
				}
				sliceMemBlock, _, _ := memoryBlockAtOffset(mockWA, sliceOffset, 0, true)
				if !bytes.Equal(sliceMemBlock, tt.want) {
					t.Errorf("\ngot  = %v\nwant = %v", sliceMemBlock, tt.want)
				}
			}

			got, err := h.Decode(ctx, mockWA, res)
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
	int16SliceHandler := mustGetHandler(t, "Array[Int16]")
	int16OptionSliceHandler := mustGetHandler(t, "Array[Int16?]")
	tests := []struct {
		name    string
		handler langsupport.TypeHandler
		value   any
		want    []byte
	}{
		{
			name:    "Array[Int16]: nil",
			handler: int16SliceHandler,
			value:   nil,
		},
		{
			name:    "Array[Int16]: []",
			handler: int16SliceHandler,
			value:   []int16{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[Int16]: [-32768]",
			handler: int16SliceHandler,
			value:   []int16{-32768},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 0, 128, 0, 0},
		},
		{
			name:    "Array[Int16]: [32767]",
			handler: int16SliceHandler,
			value:   []int16{32767},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 255, 127, 0, 0},
		},
		{
			name:    "Array[Int16]: [-32768, 32767]",
			handler: int16SliceHandler,
			value:   []int16{-32768, 32767},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 128, 255, 127, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int16]: [-32768, 32767, -32768]",
			handler: int16SliceHandler,
			value:   []int16{-32768, 32767, -32768},
			want:    []byte{1, 0, 0, 0, 241, 3, 0, 0, 0, 128, 255, 127, 0, 128, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int16]: [-32768, 32767, -32768, 32767]",
			handler: int16SliceHandler,
			value:   []int16{-32768, 32767, -32768, 32767},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 0, 128, 255, 127, 0, 128, 255, 127, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int16?]: nil",
			handler: int16OptionSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[Int16?]: []",
			handler: int16OptionSliceHandler,
			value:   []*int16{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[Int16?]: [nil]",
			handler: int16OptionSliceHandler,
			value:   []*int16{nil},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		},
		{
			name:    "Array[Int16?]: [-32768]",
			handler: int16OptionSliceHandler,
			value:   []*int16{Ptr(int16(-32768))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int16?]: [32767]",
			handler: int16OptionSliceHandler,
			value:   []*int16{Ptr(int16(32767))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int16?]: [32767, 32767]",
			handler: int16OptionSliceHandler,
			value:   []*int16{Ptr(int16(32767)), Ptr(int16(32767))},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int16?]: [-32768, 32767, -32768]",
			handler: int16OptionSliceHandler,
			value:   []*int16{Ptr(int16(-32768)), Ptr(int16(32767)), Ptr(int16(-32768))},
			want:    []byte{1, 0, 0, 0, 241, 6, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int16?]: [-32768, 32767, -32768, 32767]",
			handler: int16OptionSliceHandler,
			value:   []*int16{Ptr(int16(-32768)), Ptr(int16(32767)), Ptr(int16(-32768)), Ptr(int16(32767))},
			want:    []byte{1, 0, 0, 0, 241, 8, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int16?]: [nil, 128, nil, 128]",
			handler: int16OptionSliceHandler,
			value:   []*int16{nil, Ptr(int16(128)), nil, Ptr(int16(128))},
			want:    []byte{1, 0, 0, 0, 241, 8, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
			}

			if res[0] != 0 {
				memBlock, _, _ := memoryBlockAtOffset(mockWA, uint32(res[0]), 0, true)
				sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
				numElements := binary.LittleEndian.Uint32(memBlock[12:16])
				var wantNumElements uint32
				switch slice := tt.value.(type) {
				case []int16:
					wantNumElements = uint32(len(slice))
				case []*int16:
					wantNumElements = uint32(len(slice))
				default:
					t.Fatalf("tt.value is not a slice: %T", tt.value)
				}
				if wantNumElements != numElements {
					t.Errorf("numElements = %v, want = %v", numElements, wantNumElements)
				}
				sliceMemBlock, _, _ := memoryBlockAtOffset(mockWA, sliceOffset, 0, true)
				if !bytes.Equal(sliceMemBlock, tt.want) {
					t.Errorf("\ngot  = %v\nwant = %v", sliceMemBlock, tt.want)
				}
			}

			got, err := h.Decode(ctx, mockWA, res)
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
	uint16SliceHandler := mustGetHandler(t, "Array[UInt16]")
	uint16OptionSliceHandler := mustGetHandler(t, "Array[UInt16?]")
	tests := []struct {
		name    string
		handler langsupport.TypeHandler
		value   any
		want    []byte
	}{
		{
			name:    "Array[UInt16]: nil",
			handler: uint16SliceHandler,
			value:   nil,
		},
		{
			name:    "Array[UInt16]: []",
			handler: uint16SliceHandler,
			value:   []uint16{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[UInt16]: [0]",
			handler: uint16SliceHandler,
			value:   []uint16{0},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt16]: [65535]",
			handler: uint16SliceHandler,
			value:   []uint16{65535},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 255, 255, 0, 0},
		},
		{
			name:    "Array[UInt16]: [0, 65535]",
			handler: uint16SliceHandler,
			value:   []uint16{0, 65535},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt16]: [0, 65535, 0]",
			handler: uint16SliceHandler,
			value:   []uint16{0, 65535, 0},
			want:    []byte{1, 0, 0, 0, 241, 3, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt16]: [0, 65535, 0, 65535]",
			handler: uint16SliceHandler,
			value:   []uint16{0, 65535, 0, 65535},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 0, 0, 255, 255, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt16?]: nil",
			handler: uint16OptionSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[UInt16?]: []",
			handler: uint16OptionSliceHandler,
			value:   []*uint16{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[UInt16?]: [nil]",
			handler: uint16OptionSliceHandler,
			value:   []*uint16{nil},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		},
		{
			name:    "Array[UInt16?]: [0]",
			handler: uint16OptionSliceHandler,
			value:   []*uint16{Ptr(uint16(0))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt16?]: [65535]",
			handler: uint16OptionSliceHandler,
			value:   []*uint16{Ptr(uint16(65535))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt16?]: [65535, 65535]",
			handler: uint16OptionSliceHandler,
			value:   []*uint16{Ptr(uint16(65535)), Ptr(uint16(65535))},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt16?]: [0, 65535, 0]",
			handler: uint16OptionSliceHandler,
			value:   []*uint16{Ptr(uint16(0)), Ptr(uint16(65535)), Ptr(uint16(0))},
			want:    []byte{1, 0, 0, 0, 241, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt16?]: [0, 65535, 0, 65535]",
			handler: uint16OptionSliceHandler,
			value:   []*uint16{Ptr(uint16(0)), Ptr(uint16(65535)), Ptr(uint16(0)), Ptr(uint16(65535))},
			want:    []byte{1, 0, 0, 0, 241, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt16?]: [nil, 128, nil, 128]",
			handler: uint16OptionSliceHandler,
			value:   []*uint16{nil, Ptr(uint16(128)), nil, Ptr(uint16(128))},
			want:    []byte{1, 0, 0, 0, 241, 8, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
			}

			if res[0] != 0 {
				memBlock, _, _ := memoryBlockAtOffset(mockWA, uint32(res[0]), 0, true)
				sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
				numElements := binary.LittleEndian.Uint32(memBlock[12:16])
				var wantNumElements uint32
				switch slice := tt.value.(type) {
				case []uint16:
					wantNumElements = uint32(len(slice))
				case []*uint16:
					wantNumElements = uint32(len(slice))
				default:
					t.Fatalf("tt.value is not a slice: %T", tt.value)
				}
				if wantNumElements != numElements {
					t.Errorf("numElements = %v, want = %v", numElements, wantNumElements)
				}
				sliceMemBlock, _, _ := memoryBlockAtOffset(mockWA, sliceOffset, 0, true)
				if !bytes.Equal(sliceMemBlock, tt.want) {
					t.Errorf("\ngot  = %v\nwant = %v", sliceMemBlock, tt.want)
				}
			}

			got, err := h.Decode(ctx, mockWA, res)
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
	intSliceHandler := mustGetHandler(t, "Array[Int]")
	intOptionSliceHandler := mustGetHandler(t, "Array[Int?]")
	tests := []struct {
		name    string
		handler langsupport.TypeHandler
		value   any
		want    []byte
	}{
		{
			name:    "Array[Int]: nil",
			handler: intSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[Int]: []",
			handler: intSliceHandler,
			value:   []int32{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[Int]: [-2147483648]",
			handler: intSliceHandler,
			value:   []int32{math.MinInt32},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 0, 0, 0, 128},
		},
		{
			name:    "Array[Int]: [-32768]",
			handler: intSliceHandler,
			value:   []int32{-32768},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 0, 128, 255, 255},
		},
		{
			name:    "Array[Int]: [32767]",
			handler: intSliceHandler,
			value:   []int32{32767},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 255, 127, 0, 0},
		},
		{
			name:    "Array[Int]: [2147483647]",
			handler: intSliceHandler,
			value:   []int32{math.MaxInt32},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 255, 255, 255, 127},
		},
		{
			name:    "Array[Int]: [-32768, 32767]",
			handler: intSliceHandler,
			value:   []int32{-32768, 32767},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 128, 255, 255, 255, 127, 0, 0},
		},
		{
			name:    "Array[Int]: [-32768, 32767, -32768]",
			handler: intSliceHandler,
			value:   []int32{-32768, 32767, -32768},
			want:    []byte{1, 0, 0, 0, 241, 3, 0, 0, 0, 128, 255, 255, 255, 127, 0, 0, 0, 128, 255, 255},
		},
		{
			name:    "Array[Int]: [-32768, 32767, -32768, 32767]",
			handler: intSliceHandler,
			value:   []int32{-32768, 32767, -32768, 32767},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 0, 128, 255, 255, 255, 127, 0, 0, 0, 128, 255, 255, 255, 127, 0, 0},
		},
		{
			name:    "Array[Int?]: nil",
			handler: intOptionSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[Int?]: []",
			handler: intOptionSliceHandler,
			value:   []*int32{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[Int?]: [nil]",
			handler: intOptionSliceHandler,
			value:   []*int32{nil},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		},
		{
			name:    "Array[Int?]: [-2147483648]",
			handler: intOptionSliceHandler,
			value:   []*int32{Ptr(int32(math.MinInt32))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 128, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int?]: [-32768]",
			handler: intOptionSliceHandler,
			value:   []*int32{Ptr(int32(-32768))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 128, 255, 255, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int?]: [32767]",
			handler: intOptionSliceHandler,
			value:   []*int32{Ptr(int32(32767))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int?]: [2147483647]",
			handler: intOptionSliceHandler,
			value:   []*int32{Ptr(int32(math.MaxInt32))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 255, 255, 255, 127, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int?]: [32767, 32767]",
			handler: intOptionSliceHandler,
			value:   []*int32{Ptr(int32(32767)), Ptr(int32(32767))},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int?]: [-32768, 32767, -32768]",
			handler: intOptionSliceHandler,
			value:   []*int32{Ptr(int32(-32768)), Ptr(int32(32767)), Ptr(int32(-32768))},
			want:    []byte{1, 0, 0, 0, 241, 6, 0, 0, 0, 128, 255, 255, 0, 0, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0, 0, 128, 255, 255, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int?]: [-32768, 32767, -32768, 32767]",
			handler: intOptionSliceHandler,
			value:   []*int32{Ptr(int32(-32768)), Ptr(int32(32767)), Ptr(int32(-32768)), Ptr(int32(32767))},
			want:    []byte{1, 0, 0, 0, 241, 8, 0, 0, 0, 128, 255, 255, 0, 0, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0, 0, 128, 255, 255, 0, 0, 0, 0, 255, 127, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[Int?]: [nil, 128, nil, 128]",
			handler: intOptionSliceHandler,
			value:   []*int32{nil, Ptr(int32(128)), nil, Ptr(int32(128))},
			want:    []byte{1, 0, 0, 0, 241, 8, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
			}

			if res[0] != 0 {
				memBlock, _, _ := memoryBlockAtOffset(mockWA, uint32(res[0]), 0, true)
				sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
				numElements := binary.LittleEndian.Uint32(memBlock[12:16])
				var wantNumElements uint32
				switch slice := tt.value.(type) {
				case []int32:
					wantNumElements = uint32(len(slice))
				case []*int32:
					wantNumElements = uint32(len(slice))
				default:
					t.Fatalf("tt.value is not a slice: %T", tt.value)
				}
				if wantNumElements != numElements {
					t.Errorf("numElements = %v, want = %v", numElements, wantNumElements)
				}
				sliceMemBlock, _, _ := memoryBlockAtOffset(mockWA, sliceOffset, 0, true)
				if !bytes.Equal(sliceMemBlock, tt.want) {
					t.Errorf("\ngot  = %v\nwant = %v", sliceMemBlock, tt.want)
				}
			}

			got, err := h.Decode(ctx, mockWA, res)
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
	uintSliceHandler := mustGetHandler(t, "Array[UInt]")
	uintOptionSliceHandler := mustGetHandler(t, "Array[UInt?]")
	tests := []struct {
		name    string
		handler langsupport.TypeHandler
		value   any
		want    []byte
	}{
		{
			name:    "Array[UInt]: nil",
			handler: uintSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[UInt]: []",
			handler: uintSliceHandler,
			value:   []uint32{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[UInt]: [0]",
			handler: uintSliceHandler,
			value:   []uint32{0},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt]: [65535]",
			handler: uintSliceHandler,
			value:   []uint32{65535},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 255, 255, 0, 0},
		},
		{
			name:    "Array[UInt]: [4294967295]",
			handler: uintSliceHandler,
			value:   []uint32{math.MaxUint32},
			want:    []byte{1, 0, 0, 0, 241, 1, 0, 0, 255, 255, 255, 255},
		},
		{
			name:    "Array[UInt]: [0, 65535]",
			handler: uintSliceHandler,
			value:   []uint32{0, 65535},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0},
		},
		{
			name:    "Array[UInt]: [0, 65535, 0]",
			handler: uintSliceHandler,
			value:   []uint32{0, 65535, 0},
			want:    []byte{1, 0, 0, 0, 241, 3, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt]: [0, 65535, 0, 65535]",
			handler: uintSliceHandler,
			value:   []uint32{0, 65535, 0, 65535},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0},
		},
		{
			name:    "Array[UInt?]: nil",
			handler: uintOptionSliceHandler,
			value:   nil,
		},
		{
			name:    "Array[UInt?]: []",
			handler: uintOptionSliceHandler,
			value:   []*uint32{},
			want:    []byte{1, 0, 0, 0, 241, 0, 0, 0},
		},
		{
			name:    "Array[UInt?]: [nil]",
			handler: uintOptionSliceHandler,
			value:   []*uint32{nil},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		},
		{
			name:    "Array[UInt?]: [0]",
			handler: uintOptionSliceHandler,
			value:   []*uint32{Ptr(uint32(0))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt?]: [65535]",
			handler: uintOptionSliceHandler,
			value:   []*uint32{Ptr(uint32(65535))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt?]: [4294967295]",
			handler: uintOptionSliceHandler,
			value:   []*uint32{Ptr(uint32(math.MaxUint32))},
			want:    []byte{1, 0, 0, 0, 241, 2, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt?]: [65535, 65535]",
			handler: uintOptionSliceHandler,
			value:   []*uint32{Ptr(uint32(65535)), Ptr(uint32(65535))},
			want:    []byte{1, 0, 0, 0, 241, 4, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt?]: [0, 65535, 0]",
			handler: uintOptionSliceHandler,
			value:   []*uint32{Ptr(uint32(0)), Ptr(uint32(65535)), Ptr(uint32(0))},
			want:    []byte{1, 0, 0, 0, 241, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt?]: [0, 65535, 0, 65535]",
			handler: uintOptionSliceHandler,
			value:   []*uint32{Ptr(uint32(0)), Ptr(uint32(65535)), Ptr(uint32(0)), Ptr(uint32(65535))},
			want:    []byte{1, 0, 0, 0, 241, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0},
		},
		{
			name:    "Array[UInt?]: [nil, 128, nil, 128]",
			handler: uintOptionSliceHandler,
			value:   []*uint32{nil, Ptr(uint32(128)), nil, Ptr(uint32(128))},
			want:    []byte{1, 0, 0, 0, 241, 8, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
			}

			if res[0] != 0 {
				memBlock, _, _ := memoryBlockAtOffset(mockWA, uint32(res[0]), 0, true)
				sliceOffset := binary.LittleEndian.Uint32(memBlock[8:12])
				numElements := binary.LittleEndian.Uint32(memBlock[12:16])
				var wantNumElements uint32
				switch slice := tt.value.(type) {
				case []uint32:
					wantNumElements = uint32(len(slice))
				case []*uint32:
					wantNumElements = uint32(len(slice))
				default:
					t.Fatalf("tt.value is not a slice: %T", tt.value)
				}
				if wantNumElements != numElements {
					t.Errorf("numElements = %v, want = %v", numElements, wantNumElements)
				}
				sliceMemBlock, _, _ := memoryBlockAtOffset(mockWA, sliceOffset, 0, true)
				if !bytes.Equal(sliceMemBlock, tt.want) {
					t.Errorf("\ngot  = %v\nwant = %v", sliceMemBlock, tt.want)
				}
			}

			got, err := h.Decode(ctx, mockWA, res)
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
func TestPrimitiveSlicesEncodeDecode_Int64(t *testing.T) {
	tests := []struct {
		name    string
		handler langsupport.TypeHandler
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
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
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
		handler langsupport.TypeHandler
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
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
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
		handler langsupport.TypeHandler
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
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
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
		handler langsupport.TypeHandler
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
			mockWA := &myWasmMock{}
			h := tt.handler
			res, _, err := h.Encode(ctx, mockWA, tt.value)
			if err != nil {
				t.Fatalf("h.Encode() returned an error: %v", err)
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
