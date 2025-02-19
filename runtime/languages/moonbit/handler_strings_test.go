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
	"testing"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
	wasm "github.com/tetratelabs/wazero/api"

	"github.com/stretchr/testify/mock"
)

func TestConvertMoonBitUTF16ToUTF8(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		data     []byte
		expected string
		wantErr  bool
	}{
		{
			name:     "Valid UTF-16 data",
			data:     []byte{0x48, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x6f, 0x00, 0x2c, 0x00, 0x20, 0x00, 0x57, 0x00, 0x6f, 0x00, 0x72, 0x00, 0x6c, 0x00, 0x64, 0x00, 0x21, 0x00},
			expected: "Hello, World!",
			wantErr:  false,
		},
		{
			name:     "UTF-16 with emojis",
			data:     []byte{0x48, 0x00, 0x65, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x6f, 0x00, 0x2c, 0x00, 0x20, 0x00, 0x3d, 0xd8, 0x0d, 0xde, 0x21, 0x00},
			expected: "Hello, 😍!",
			wantErr:  false,
		},
		{
			name:     "UTF-16 with non-Latin characters",
			data:     []byte{0x53, 0x30, 0x93, 0x30, 0x6b, 0x30, 0x6f, 0x30, 0x6b, 0x30, 0x59, 0x4e, 0x16, 0x75},
			expected: "こんにはに乙甖",
			wantErr:  false,
		},
		{
			name: "empty array",
			data: []byte{},
		},
		{
			name: "bad array",
			data: []byte{0},
		},
		{
			name:     "null",
			data:     []byte{0, 0},
			expected: "\x00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := convertMoonBitUTF16ToUTF8(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertMoonBitUTF16ToUTF8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("convertMoonBitUTF16ToUTF8() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestConvertGoUTF8ToUTF16(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Simple ASCII",
			input: "Hello, World!",
		},
		{
			name:  "UTF-8 with emojis",
			input: "Hello, 🌍!",
		},
		{
			name:  "UTF-8 with non-Latin characters",
			input: "こんにちは世界",
		},
		{
			name:  "UTF-8 with mixed characters",
			input: "Hello, 世界! 🌍",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			encoded := convertGoUTF8ToUTF16(tt.input)
			decoded, err := convertMoonBitUTF16ToUTF8(encoded)
			if err != nil {
				t.Errorf("convertMoonBitUTF16ToUTF8() error = %v", err)
				return
			}
			if decoded != tt.input {
				t.Errorf("Round trip conversion failed: got = %v, want = %v", decoded, tt.input)
			}
		})
	}
}

type mockWasmAdapter struct {
	mock.Mock
	langsupport.WasmAdapter
}

func (m *mockWasmAdapter) Memory() wasm.Memory {
	args := m.Called()
	return args.Get(0).(wasm.Memory)
}

func (m *mockWasmAdapter) allocateAndPinMemory(ctx context.Context, size, blockType uint32) (uint32, utils.Cleaner, error) {
	args := m.Called(ctx, size, blockType)
	return args.Get(0).(uint32), nil, args.Error(2)
}

type mockMemory struct {
	mock.Mock
	wasm.Memory
}

func (m *mockMemory) Read(offset, size uint32) ([]byte, bool) {
	args := m.Called(offset, size)
	return args.Get(0).([]byte), args.Bool(1)
}

func TestStringDataAtOffset(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		offset       uint32
		memBlock     []byte
		expectedSize int
		expectedErr  error
	}{
		{
			name:         "Valid memory block, UTF-16 String 'Hello, ...0!' with remainder 0",
			offset:       100,
			memBlock:     []byte{1, 0, 0, 0, 243, 7, 0, 0, 72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 0, 46, 0, 46, 0, 46, 0, 48, 0, 33, 0, 97, 0, 109, 3},
			expectedSize: 24,
			expectedErr:  nil,
		},
		{
			name:         "Valid memory block, UTF-16 String 'Hello, 2!' with remainder 2",
			offset:       100,
			memBlock:     []byte{1, 0, 0, 0, 243, 5, 0, 0, 72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 0, 50, 0, 33, 0, 32, 1},
			expectedSize: 18,
			expectedErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockMem := new(mockMemory)
			mockWA := new(mockWasmAdapter)
			mockWA.On("Memory").Return(mockMem)
			mockMem.On("Read", tt.offset, uint32(8)).Return(tt.memBlock[:8], true)
			mockMem.On("Read", tt.offset, uint32(len(tt.memBlock))).Return(tt.memBlock, true)

			data, err := stringDataAtOffset(mockWA, tt.offset)
			size := len(data)
			if size != tt.expectedSize || (err != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("stringDataAtOffset() = (data: %v, size: %v, err: %v), want (size: %v, err: %v)",
					data, size, err, tt.expectedSize, tt.expectedErr)
			}
		})
	}
}

func TestDoWriteStringBytes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         string
		wantTotalSize uint32
	}{
		{
			name:          "empty string",
			input:         "",
			wantTotalSize: 4,
		},
		{
			name:  "string length 1",
			input: "a",
		},
		{
			name:  "string length 2",
			input: "ab",
		},
		{
			name:  "string length 3",
			input: "abc",
		},
		{
			name:  "string length 4",
			input: "abcd",
		},
		{
			name:  "string length 5",
			input: "abcde",
		},
		{
			name:  "Simple ASCII",
			input: "Hello, World!",
		},
		{
			name:  "UTF-8 with emojis",
			input: "Hello, 🌍!",
		},
		{
			name:  "UTF-8 with non-Latin characters",
			input: "こんにちは世界",
		},
		{
			name:  "UTF-8 with mixed characters",
			input: "Hello, 世界! 🌍",
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{offset: 100}

			h := &stringHandler{}
			bytes := convertGoUTF8ToUTF16(tt.input)
			expectedSize := len(bytes)
			offset, _, err := h.doWriteStringBytes(ctx, mockWA, bytes)
			if err != nil {
				t.Fatal(err)
			}

			data, err := stringDataAtOffset(mockWA, offset)
			size := len(data)
			if size != expectedSize || err != nil {
				t.Errorf("stringDataAtOffset() = (data: %v, size: %v, err: %v), want (size: %v)",
					data, size, err, expectedSize)
			}
			encoded := convertGoUTF8ToUTF16(tt.input)
			decoded, err := convertMoonBitUTF16ToUTF8(encoded)
			if err != nil {
				t.Errorf("convertMoonBitUTF16ToUTF8() error = %v", err)
				return
			}
			if decoded != tt.input {
				t.Errorf("Round trip conversion failed: got = %v, want = %v", decoded, tt.input)
			}
		})
	}
}
