// -*- compile-command: "go test -timeout 30s -tags integration -run ^TestPrimitivesInt16 github.com/gmlewis/modus/runtime/languages/moonbit/tests"; -*-

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

func TestPrimitivesInt16Input_min(t *testing.T) {
	fnName := "test_int16_input_min"
	if _, err := fixture.CallFunction(t, fnName, int16(math.MinInt16)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt16Output_min(t *testing.T) {
	fnName := "test_int16_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MinInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int16); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesInt16Input_max(t *testing.T) {
	fnName := "test_int16_input_max"
	if _, err := fixture.CallFunction(t, fnName, int16(math.MaxInt16)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt16Output_max(t *testing.T) {
	fnName := "test_int16_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MaxInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int16); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesInt16OptionInput_min(t *testing.T) {
	fnName := "test_int16_option_input_min"
	b := int16(math.MinInt16)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt16OptionOutput_min(t *testing.T) {
	fnName := "test_int16_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MinInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int16); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesInt16OptionInput_max(t *testing.T) {
	fnName := "test_int16_option_input_max"
	b := int16(math.MaxInt16)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt16OptionOutput_max(t *testing.T) {
	fnName := "test_int16_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MaxInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int16); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesInt16OptionInput_none(t *testing.T) {
	fnName := "test_int16_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt16OptionOutput_none(t *testing.T) {
	fnName := "test_int16_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
