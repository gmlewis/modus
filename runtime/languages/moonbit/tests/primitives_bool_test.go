// -*- compile-command: "go test -timeout 30s -tags integration -run ^TestPrimitivesBool github.com/gmlewis/modus/runtime/languages/moonbit/tests"; -*-

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

func TestPrimitivesBoolInput_false(t *testing.T) {
	fnName := "test_bool_input_false"
	if _, err := fixture.CallFunction(t, fnName, false); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBoolOutput_false(t *testing.T) {
	fnName := "test_bool_output_false"

	// primitiveHandler.Decode(vals: [0])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := false
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(bool); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesBoolInput_true(t *testing.T) {
	fnName := "test_bool_input_true"
	if _, err := fixture.CallFunction(t, fnName, true); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBoolOutput_true(t *testing.T) {
	fnName := "test_bool_output_true"

	// primitiveHandler.Decode(vals: [1])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := true
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(bool); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesBoolOptionInput_false(t *testing.T) {
	fnName := "test_bool_option_input_false"
	b := false

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBoolOptionOutput_false(t *testing.T) {
	fnName := "test_bool_option_output_false"

	// primitiveHandler.Decode(vals: [0])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := false
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*bool); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesBoolOptionInput_true(t *testing.T) {
	fnName := "test_bool_option_input_true"
	b := true

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBoolOptionOutput_true(t *testing.T) {
	fnName := "test_bool_option_output_true"

	// primitiveHandler.Decode(vals: [1])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := true
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*bool); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesBoolOptionInput_none(t *testing.T) {
	fnName := "test_bool_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBoolOptionOutput_none(t *testing.T) {
	fnName := "test_bool_option_output_none"

	// primitiveHandler.Decode(vals: [4294967295])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if !utils.HasNil(result) {
		t.Errorf("expected a nil result, got %T=%[1]v", result)
	}
}
