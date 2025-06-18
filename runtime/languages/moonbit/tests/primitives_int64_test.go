// -*- compile-command: "go test -timeout 30s -tags integration -run ^TestPrimitivesInt64 github.com/gmlewis/modus/runtime/languages/moonbit/tests"; -*-

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

func TestPrimitivesInt64Input_min(t *testing.T) {
	fnName := "test_int64_input_min"
	if _, err := fixture.CallFunction(t, fnName, int64(math.MinInt64)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt64Output_min(t *testing.T) {
	fnName := "test_int64_output_min"

	// primitiveHandler.Decode(vals: [9223372036854775808])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int64(math.MinInt64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesInt64Input_max(t *testing.T) {
	fnName := "test_int64_input_max"
	if _, err := fixture.CallFunction(t, fnName, int64(math.MaxInt64)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt64Output_max(t *testing.T) {
	fnName := "test_int64_output_max"

	// primitiveHandler.Decode(vals: [9223372036854775807])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int64(math.MaxInt64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesInt64OptionInput_min(t *testing.T) {
	fnName := "test_int64_option_input_min"
	b := int64(math.MinInt64)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt64OptionOutput_min(t *testing.T) {
	fnName := "test_int64_option_output_min"

	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 128]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int64(math.MinInt64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int64); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesInt64OptionInput_max(t *testing.T) {
	fnName := "test_int64_option_input_max"
	b := int64(math.MaxInt64)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt64OptionOutput_max(t *testing.T) {
	fnName := "test_int64_option_output_max"

	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 255 255 255 255 255 255 255 127]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int64(math.MaxInt64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int64); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

//	func TestPrimitivesInt64OptionInput_none(t *testing.T) {
//		fnName := "test_int64_option_input_none"
//		if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
//			t.Error(err)
//		}
//	}

func TestPrimitivesInt64OptionOutput_none(t *testing.T) {
	fnName := "test_int64_option_output_none"

	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), moonBitType=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
