// -*- compile-command: "NO_COLOR=1 go test -timeout 30s -tags integration -run '^TestStrings' ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// Tests pass with moonc v0.6.20

package moonbit

import (
	"context"
	"encoding/binary"
	"errors"
	"log"
	"testing"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
	wasm "github.com/tetratelabs/wazero/api"

	"github.com/stretchr/testify/mock"
)

func TestStrings_ConvertMoonBitUTF16ToUTF8(t *testing.T) {
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

func TestStrings_ConvertGoUTF8ToUTF16(t *testing.T) {
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
	uint32Val, ok := args.Get(0).(uint32)
	if !ok {
		log.Printf("mockWasmAdapter.allocateAndPinMemory() FAILURE: expected uint32, got %T", args.Get(0))
		return 0, nil, errors.New("mockWasmAdapter.allocateAndPinMemory() expected uint32 return value")
	}
	return uint32Val, nil, args.Error(2)
}

type mockMemory struct {
	mock.Mock
	wasm.Memory
}

func (m *mockMemory) Read(offset, size uint32) ([]byte, bool) {
	if m == nil {
		log.Printf("mockMemory.Read() FAILURE: mockMemory is nil")
		return nil, false
	}
	args := m.Called(offset, size)
	byteSlice, ok := args.Get(0).([]byte)
	if !ok {
		log.Printf("mockMemory.Read() FAILURE: expected []byte, got %T", args.Get(0))
		return nil, false
	}
	boolVal, ok := args.Get(1).(bool)
	if !ok {
		log.Printf("mockMemory.Read() FAILURE: expected bool, got %T", args.Get(1))
		return nil, false
	}
	return byteSlice, boolVal
}

func TestStrings_StringDataAtOffset(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		memBlock       []byte
		expectedLength int
		want           string
		expectedErr    error
	}{
		{
			name:           "empty string",
			memBlock:       []byte("\xff\xff\xff\xff\x00\x00\x00\xf3\x00\x00\x00\x00\x00\x00\x00\x00"),
			expectedLength: 0,
		},
		{
			name:           "length 1 string",
			memBlock:       []byte("\xff\xff\xff\xff\x01\x00\x00\xf31\x00\x00\x00\x00\x00\x00\x00"),
			expectedLength: 2,
			want:           "1",
		},
		{
			name:           "length 2 string",
			memBlock:       []byte("\xff\xff\xff\xff\x02\x00\x00\xf31\x002\x00\x00\x00\x00\x00"),
			expectedLength: 4,
			want:           "12",
		},
		{
			name:           "length 3 string",
			memBlock:       []byte("\xff\xff\xff\xff\x03\x00\x00\xf31\x002\x003\x00\x00\x00"),
			expectedLength: 6,
			want:           "123",
		},
		{
			name:           "length 4 string",
			memBlock:       []byte("\xff\xff\xff\xff\x04\x00\x00\xf31\x002\x003\x004\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
			expectedLength: 8,
			want:           "1234",
		},
		{
			name:           "length 5 string",
			memBlock:       []byte("\xff\xff\xff\xff\x05\x00\x00\xf31\x002\x003\x004\x005\x00\x00\x00\x00\x00\x00\x00"),
			expectedLength: 10,
			want:           "12345",
		},
		{
			name:           "length 6 string",
			memBlock:       []byte("\xff\xff\xff\xff\x06\x00\x00\xf31\x002\x003\x004\x005\x006\x00\x00\x00\x00\x00"),
			expectedLength: 12,
			want:           "123456",
		},
		{
			name:           "length 7 string",
			memBlock:       []byte("\xff\xff\xff\xff\x07\x00\x00\xf31\x002\x003\x004\x005\x006\x007\x00\x00\x00"),
			expectedLength: 14,
			want:           "1234567",
		},
		{
			name:           "length 8 string",
			memBlock:       []byte("\xff\xff\xff\xff\x08\x00\x00\xf31\x002\x003\x004\x005\x006\x007\x008\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
			expectedLength: 16,
			want:           "12345678",
		},
		{
			name:           "length 9 string",
			memBlock:       []byte("\xff\xff\xff\xff\x09\x00\x00\xf31\x002\x003\x004\x005\x006\x007\x008\x009\x00\x00\x00\x00\x00\x00\x00"),
			expectedLength: 18,
			want:           "123456789",
		},
		{
			name:           "length 10 string",
			memBlock:       []byte("\xff\xff\xff\xff\x0a\x00\x00\xf31\x002\x003\x004\x005\x006\x007\x008\x009\x000\x00\x00\x00\x00\x00"),
			expectedLength: 20,
			want:           "1234567890",
		},
		{
			name:           "Valid memory block, UTF-16 String 'Hello, ...0!'",
			memBlock:       []byte{1, 0, 0, 0, 12, 0, 0, 243, 72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 0, 46, 0, 46, 0, 46, 0, 48, 0, 33, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedLength: 24,
			want:           "Hello, ...0!",
		},
		{
			name:           "Valid memory block, UTF-16 String 'Hello, 2!'",
			memBlock:       []byte{1, 0, 0, 0, 9, 0, 0, 243, 72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 0, 50, 0, 33, 0, 0, 0, 0, 0, 0, 0},
			expectedLength: 18,
			want:           "Hello, 2!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			offset := uint32(100)
			mockMem := new(mockMemory)
			mockWA := new(mockWasmAdapter)
			mockWA.On("Memory").Return(mockMem)
			// Calculate expected memory block size based on string length
			// Parse header to get words count using new memory format
			part2 := binary.LittleEndian.Uint32(tt.memBlock[4:8])
			words := part2 & 0xffffff
			// For string tests, use actual string data size (words*2 for UTF-16)
			expectedSize := uint32(8 + words*2)
			// First call: read header (8 bytes)
			mockMem.On("Read", offset, uint32(8)).Return(tt.memBlock[:8], true)
			// Second call: read full memory block (calculated size)
			mockMem.On("Read", offset, expectedSize).Return(tt.memBlock[:expectedSize], true)

			data, err := stringDataAtOffset(mockWA, offset)
			size := len(data)
			if size != tt.expectedLength || !errorsEqual(err, tt.expectedErr) {
				t.Errorf("stringDataAtOffset() = (data: %v, size: %v, err: %v), want (size: %v, err: %v)",
					data, size, err, tt.expectedLength, tt.expectedErr)
			}

			s, err := doReadString(data)
			if err != nil {
				t.Fatal(err)
			}
			if s != tt.want {
				t.Errorf("doReadString() = '%v', want '%v'", s, tt.want)
			}
		})
	}
}

func TestStrings_DoWriteStringBytes(t *testing.T) {
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

	ctx := t.Context()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			mockWA := &myWasmMock{}

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

// Helper function to compare errors safely
func errorsEqual(a, b error) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Error() == b.Error()
}