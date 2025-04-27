// -*- compile-command: "go test -timeout 30s -tags integration -run ^TestPrimitivesInt github.com/gmlewis/modus/runtime/languages/moonbit/tests"; -*-

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

func TestPrimitivesIntInput_min(t *testing.T) {
	fnName := "test_int_input_min"
	if _, err := fixture.CallFunction(t, fnName, int(math.MinInt32)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesIntOutput_min(t *testing.T) {
	fnName := "test_int_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int32(math.MinInt32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int32); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesIntInput_max(t *testing.T) {
	fnName := "test_int_input_max"
	if _, err := fixture.CallFunction(t, fnName, int(math.MaxInt32)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesIntOutput_max(t *testing.T) {
	fnName := "test_int_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int32(math.MaxInt32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int32); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesIntOptionInput_min(t *testing.T) {
	fnName := "test_int_option_input_min"
	b := int32(math.MinInt32)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesIntOptionOutput_min(t *testing.T) {
	fnName := "test_int_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int32(math.MinInt32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int32); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesIntOptionInput_max(t *testing.T) {
	fnName := "test_int_option_input_max"
	b := int32(math.MaxInt32)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesIntOptionOutput_max(t *testing.T) {
	fnName := "test_int_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int32(math.MaxInt32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int32); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesIntOptionInput_none(t *testing.T) {
	fnName := "test_int_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesIntOptionOutput_none(t *testing.T) {
	fnName := "test_int_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
