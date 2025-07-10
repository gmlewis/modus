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
		{name: "0", want: ""},
		{name: "1", want: "1"},
		{name: "2", want: "12"},
		{name: "3", want: "123"},
		{name: "4", want: "1234"},
		{name: "5", want: "12345"},
		{name: "6", want: "123456"},
		{name: "7", want: "1234567"},
		{name: "8", want: "12345678"},
		{name: "9", want: "123456789"},
		{name: "10", want: "1234567890"},
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
