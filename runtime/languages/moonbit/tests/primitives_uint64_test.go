// -*- compile-command: "go test -timeout 30s -tags integration -run ^TestPrimitivesUInt64 github.com/gmlewis/modus/runtime/languages/moonbit/tests"; -*-

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
	"math"
	"testing"

	"github.com/gmlewis/modus/runtime/utils"
)

func TestPrimitivesUInt64Input_min(t *testing.T) {
	fnName := "test_uint64_input_min"
	if _, err := fixture.CallFunction(t, fnName, uint64(0)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUInt64Output_min(t *testing.T) {
	fnName := "test_uint64_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint64(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(uint64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesUInt64Input_max(t *testing.T) {
	fnName := "test_uint64_input_max"
	if _, err := fixture.CallFunction(t, fnName, uint64(math.MaxUint64)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUInt64Output_max(t *testing.T) {
	fnName := "test_uint64_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint64(math.MaxUint64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(uint64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesUInt64OptionInput_min(t *testing.T) {
	fnName := "test_uint64_option_input_min"
	b := uint64(0)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUInt64OptionOutput_min(t *testing.T) {
	fnName := "test_uint64_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint64(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*uint64); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesUInt64OptionInput_max(t *testing.T) {
	fnName := "test_uint64_option_input_max"
	b := uint64(math.MaxUint64)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUInt64OptionOutput_max(t *testing.T) {
	fnName := "test_uint64_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint64(math.MaxUint64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*uint64); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesUInt64OptionInput_none(t *testing.T) {
	fnName := "test_uint64_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUInt64OptionOutput_none(t *testing.T) {
	fnName := "test_uint64_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
