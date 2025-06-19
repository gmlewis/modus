// -*- compile-command: "go test -timeout 30s -tags integration -run ^TestPrimitivesByte github.com/gmlewis/modus/runtime/languages/moonbit/tests"; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// Tests pass with moonc v0.6.18+8382ed77e

package moonbit_test

import (
	"math"
	"testing"

	"github.com/gmlewis/modus/runtime/utils"
)

func TestPrimitivesByteInput_min(t *testing.T) {
	fnName := "test_byte_input_min"
	if _, err := fixture.CallFunction(t, fnName, byte(0)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesByteOutput_min(t *testing.T) {
	fnName := "test_byte_output_min"

	// primitiveHandler.Decode(vals: [0])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := byte(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(byte); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesByteInput_max(t *testing.T) {
	fnName := "test_byte_input_max"
	if _, err := fixture.CallFunction(t, fnName, byte(math.MaxUint8)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesByteOutput_max(t *testing.T) {
	fnName := "test_byte_output_max"

	// primitiveHandler.Decode(vals: [255])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := byte(math.MaxUint8)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(byte); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesByteOptionInput_min(t *testing.T) {
	fnName := "test_byte_option_input_min"
	b := byte(0)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesByteOptionOutput_min(t *testing.T) {
	fnName := "test_byte_option_output_min"

	// primitiveHandler.Decode(vals: [0])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := byte(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*byte); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesByteOptionInput_max(t *testing.T) {
	fnName := "test_byte_option_input_max"
	b := byte(math.MaxUint8)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesByteOptionOutput_max(t *testing.T) {
	fnName := "test_byte_option_output_max"

	// primitiveHandler.Decode(vals: [255])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := byte(math.MaxUint8)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*byte); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesByteOptionInput_none(t *testing.T) {
	fnName := "test_byte_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesByteOptionOutput_none(t *testing.T) {
	fnName := "test_byte_option_output_none"

	// primitiveHandler.Decode(vals: [4294967295])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
