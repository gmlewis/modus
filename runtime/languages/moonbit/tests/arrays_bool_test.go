// -*- compile-command: "NO_COLOR=1 go test -timeout 30s -tags integration -run ^TestArrayOutput_bool ."; -*-

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
	"reflect"
	"testing"
)

func TestArrayOutput_bool_option_none(t *testing.T) {
	fnName := "test_array_option_output_bool_none"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if result != nil {
		t.Error("expected a nil result")
	}
}

func TestArrayOutput_bool_option_some_0(t *testing.T) {
	fnName := "test_array_option_output_bool_some_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &[]bool{}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*[]bool); !ok {
		t.Errorf("expected a *[]bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_0(t *testing.T) {
	fnName := "test_array_output_bool_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []bool{}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]bool); !ok {
		t.Errorf("expected a []bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_1(t *testing.T) {
	fnName := "test_array_output_bool_1"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []bool{true}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]bool); !ok {
		t.Errorf("expected a []bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_2(t *testing.T) {
	fnName := "test_array_output_bool_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []bool{false, true}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]bool); !ok {
		t.Errorf("expected a []bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_3(t *testing.T) {
	fnName := "test_array_output_bool_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []bool{true, true, true}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]bool); !ok {
		t.Errorf("expected a []bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_4(t *testing.T) {
	fnName := "test_array_output_bool_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []bool{false, false, false, false}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]bool); !ok {
		t.Errorf("expected a []bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_option_0(t *testing.T) {
	fnName := "test_array_output_bool_option_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*bool{}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_option_1_none(t *testing.T) {
	fnName := "test_array_output_bool_option_1_none"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*bool{nil}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_option_1_false(t *testing.T) {
	fnName := "test_array_output_bool_option_1_false"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0 := bool(false)
	expected := []*bool{&b0}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_option_1_true(t *testing.T) {
	fnName := "test_array_output_bool_option_1_true"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0 := bool(true)
	expected := []*bool{&b0}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_option_2(t *testing.T) {
	fnName := "test_array_output_bool_option_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0 := bool(true)
	expected := []*bool{&b0, nil}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_option_3(t *testing.T) {
	fnName := "test_array_output_bool_option_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b1 := bool(true)
	expected := []*bool{nil, &b1, &b1}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestArrayOutput_bool_option_4(t *testing.T) {
	fnName := "test_array_output_bool_option_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0, b1, b2 := bool(false), bool(true), bool(false)
	expected := []*bool{&b0, &b1, &b2, nil}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}
