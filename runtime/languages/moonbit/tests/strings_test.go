/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package moonbit_test

import (
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
