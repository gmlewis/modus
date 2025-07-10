// -*- compile-command: "NO_COLOR=1 go test -timeout 30s -tags integration -run ^TestString github.com/gmlewis/modus/runtime/languages/moonbit/tests"; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// Tests FAIL with moonc v0.6.18+8382ed77e

package moonbit_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gmlewis/modus/runtime/utils"
)

// "Hello World" in Japanese
const testString = "こんにちは、世界"

func TestStringInput(t *testing.T) {
	fnName := "test_string_input"
	if _, err := fixture.CallFunction(t, fnName, testString); err != nil {
		t.Error(err)
	}
}

func TestStringOptionInput(t *testing.T) {
	fnName := "test_string_option_input"
	s := testString

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &s); err != nil {
		t.Error(err)
	}
}

func TestStringOptionInput_none(t *testing.T) {
	fnName := "test_string_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestStringOutputLengths(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// NEW: memoryBlockAtOffset(offset: 10408=0x000028A8=[168 40 0 0]): classID: 80, words: 0, size: 16, memBlockHeader: [255 255 255 255 0 0 0 80 0 0 0 0 0 0 0 0]
		{name: "0", want: ""},
		// NEW: memoryBlockAtOffset(offset: 10392=0x00002898=[152 40 0 0]): classID: 80, words: 1, size: 16, memBlockHeader: [255 255 255 255 1 0 0 80 49 0 0 0 0 0 0 0]
		{name: "1", want: "1"},
		// NEW: memoryBlockAtOffset(offset: 10376=0x00002888=[136 40 0 0]): classID: 80, words: 2, size: 16, memBlockHeader: [255 255 255 255 2 0 0 80 49 0 50 0 0 0 0 0]
		{name: "2", want: "12"},
		// NEW: memoryBlockAtOffset(offset: 10360=0x00002878=[120 40 0 0]): classID: 80, words: 3, size: 16, memBlockHeader: [255 255 255 255 3 0 0 80 49 0 50 0 51 0 0 0]
		{name: "3", want: "123"},
		// NEW: memoryBlockAtOffset(offset: 10336=0x00002860=[96 40 0 0]): classID: 80, words: 4, size: 24, memBlockHeader: [255 255 255 255 4 0 0 80 49 0 50 0 51 0 52 0]
		{name: "4", want: "1234"},
		// NEW: memoryBlockAtOffset(offset: 10312=0x00002848=[72 40 0 0]): classID: 80, words: 5, size: 24, memBlockHeader: [255 255 255 255 5 0 0 80 49 0 50 0 51 0 52 0]
		{name: "5", want: "12345"},
		// NEW: memoryBlockAtOffset(offset: 10288=0x00002830=[48 40 0 0]): classID: 80, words: 6, size: 24, memBlockHeader: [255 255 255 255 6 0 0 80 49 0 50 0 51 0 52 0]
		{name: "6", want: "123456"},
		// NEW: memoryBlockAtOffset(offset: 10264=0x00002818=[24 40 0 0]): classID: 80, words: 7, size: 24, memBlockHeader: [255 255 255 255 7 0 0 80 49 0 50 0 51 0 52 0]
		{name: "7", want: "1234567"},
		// NEW: memoryBlockAtOffset(offset: 10232=0x000027F8=[248 39 0 0]): classID: 80, words: 8, size: 32, memBlockHeader: [255 255 255 255 8 0 0 80 49 0 50 0 51 0 52 0]
		{name: "8", want: "12345678"},
		// NEW: memoryBlockAtOffset(offset: 10200=0x000027D8=[216 39 0 0]): classID: 80, words: 9, size: 32, memBlockHeader: [255 255 255 255 9 0 0 80 49 0 50 0 51 0 52 0]
		{name: "9", want: "123456789"},
		// NEW: memoryBlockAtOffset(offset: 10168=0x000027B8=[184 39 0 0]): classID: 80, words: 10, size: 32, memBlockHeader: [255 255 255 255 10 0 0 80 49 0 50 0 51 0 52 0]
		{name: "10", want: "1234567890"},
		// NEW: memoryBlockAtOffset(offset: 396432=0x00060C90=[144 12 6 0]): classID: 80, words: 100000, size: 200016, memBlockHeader: [1 0 0 0 160 134 1 80 49 0 50 0 51 0 52 0]
		{name: "100000", want: strings.Repeat("1234567890", 10000)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fnName := fmt.Sprintf("test_string_output_len_%v", tt.name)
			result, err := fixture.CallFunction(t, fnName)
			if err != nil {
				t.Fatal(err)
			}

			if result == nil {
				t.Error("expected a result")
			} else if got, ok := result.(string); !ok {
				t.Errorf("expected a string, got %T", result)
			} else if got != tt.want {
				t.Errorf("%v = %q, want %q", fnName, got, tt.want)
			}
		})
	}
}

func TestStringOutputRepeat(t *testing.T) {
	fnName := "test_string_output_repeat"
	tests := []struct {
		name string
		num  int
	}{
		// NEW: memoryBlockAtOffset(offset: 10408=0x000028A8=[168 40 0 0]): classID: 80, words: 0, size: 16, memBlockHeader: [255 255 255 255 0 0 0 80 0 0 0 0 0 0 0 0]
		{name: "0"},
		// NEW: memoryBlockAtOffset(offset: 98864=0x00018230=[48 130 1 0]): classID: 80, words: 800, size: 1616, memBlockHeader: [1 0 0 0 32 3 0 80 49 0 48 0 13 0 0 0 0 0 0 0 0 0 0 0]
		{name: "10", num: 10}, // string length 20
		// NEW: memoryBlockAtOffset(offset: 120464=0x0001D690=[144 214 1 0]): classID: 80, words: 8000, size: 16016, memBlockHeader: [1 0 0 0 64 31 0 80 49 0 48 0 48 0 0 0 0 0 0 0 0 0 0 0]
		{name: "100", num: 100}, // string length 300
		// NEW: memoryBlockAtOffset(offset: 1104464=0x0010DA50=[80 218 16 0]): classID: 80, words: 336000, size: 672016, memBlockHeader: [1 0 0 0 128 32 5 80 49 0 48 0 48 0 48 0 0 0 0 0 0 0 0 0]
		{name: "1000", num: 1000}, // string length 4000
		// NEW: memoryBlockAtOffset(offset: 10176464=0x009B47D0=[208 71 155 0]): classID: 80, words: 3360000, size: 6720016, memBlockHeader: [1 0 0 0 0 69 51 80 49 0 48 0 48 0 48 0 48 0 0 0 0 0 0 0]
		{name: "10000", num: 10000}, // string length 50000
		// NEW: memoryBlockAtOffset(offset: 96160=0x000177A0=[160 119 1 0]): classID: 0, words: 80, size: 176, memBlockHeader: [1 0 0 0 80 0 0 0 49 0 0 2 13 0 0 0 0 0 0 0 0 0 0 0]
		{name: "1", num: 1},
		// NEW: memoryBlockAtOffset(offset: 96160=0x000177A0=[160 119 1 0]): classID: 0, words: 80, size: 176, memBlockHeader: [1 0 0 0 80 0 0 0 49 0 50 0 13 0 0 0 0 0 0 0 0 0 0 0]
		{name: "12", num: 1},
		// NEW: memoryBlockAtOffset(offset: 96160=0x000177A0=[160 119 1 0]): classID: 0, words: 80, size: 176, memBlockHeader: [1 0 0 0 80 0 0 0 49 0 50 0 51 0 0 0 0 0 0 0 0 0 0 0]
		{name: "123", num: 1},
		// NEW: memoryBlockAtOffset(offset: 96160=0x000177A0=[160 119 1 0]): classID: 0, words: 336, size: 688, memBlockHeader: [1 0 0 0 80 1 0 0 49 0 50 0 51 0 52 0 0 0 0 0 0 0 0 0]
		{name: "1234", num: 1},
		// NEW: memoryBlockAtOffset(offset: 96160=0x000177A0=[160 119 1 0]): classID: 0, words: 336, size: 688, memBlockHeader: [1 0 0 0 80 1 0 0 49 0 50 0 51 0 52 0 53 0 0 0 0 0 0 0]
		{name: "12345", num: 1},
		// NEW: memoryBlockAtOffset(offset: 96160=0x000177A0=[160 119 1 0]): classID: 0, words: 336, size: 688, memBlockHeader: [1 0 0 0 80 1 0 0 49 0 50 0 51 0 52 0 53 0 54 0 0 0 0 0]
		{name: "123456", num: 1},
		// NEW: memoryBlockAtOffset(offset: 96160=0x000177A0=[160 119 1 0]): classID: 0, words: 336, size: 688, memBlockHeader: [1 0 0 0 80 1 0 0 49 0 50 0 51 0 52 0 53 0 54 0 55 0 0 0]
		{name: "1234567", num: 1},
		// NEW: memoryBlockAtOffset(offset: 96160=0x000177A0=[160 119 1 0]): classID: 0, words: 592, size: 1200, memBlockHeader: [1 0 0 0 80 2 0 0 49 0 50 0 51 0 52 0 53 0 54 0 55 0 56 0]
		{name: "12345678", num: 1},
		// NEW: memoryBlockAtOffset(offset: 96160=0x000177A0=[160 119 1 0]): classID: 0, words: 592, size: 1200, memBlockHeader: [1 0 0 0 80 2 0 0 49 0 50 0 51 0 52 0 53 0 54 0 55 0 56 0]
		{name: "123456789", num: 1},
		// NEW: memoryBlockAtOffset(offset: 96160=0x000177A0=[160 119 1 0]): classID: 0, words: 592, size: 1200, memBlockHeader: [1 0 0 0 80 2 0 0 49 0 50 0 51 0 52 0 53 0 54 0 55 0 56 0]
		{name: "1234567890", num: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := fixture.CallFunction(t, fnName, tt.name, tt.num)
			if err != nil {
				t.Fatal(err)
			}

			want := strings.Repeat(tt.name, tt.num)

			if result == nil {
				t.Error("expected a result")
			} else if got, ok := result.(string); !ok {
				t.Errorf("expected a string, got %T", result)
			} else if got != want {
				t.Errorf("%v = %q, want %q", fnName, got, want)
			}
		})
	}
}

func TestStringOutput(t *testing.T) {
	fnName := "test_string_output"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if r != testString {
		t.Errorf("expected %s, got %s", testString, r)
	}
}

func TestStringOptionOutput(t *testing.T) {
	fnName := "test_string_option_output"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*string); !ok {
		t.Errorf("expected a *string, got %T", result)
	} else if *r != testString {
		t.Errorf("expected %s, got %s", testString, *r)
	}
}

func TestStringOptionOutput_none(t *testing.T) {
	fnName := "test_string_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
