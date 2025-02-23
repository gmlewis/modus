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
	"reflect"
	"slices"
	"testing"
)

func TestFixedArrayOutput_double_0(t *testing.T) {
	fnName := "test_fixedarray_output_double_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float64{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float64); !ok {
		t.Errorf("expected a []float64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_double_1(t *testing.T) {
	fnName := "test_fixedarray_output_double_1"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float64{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float64); !ok {
		t.Errorf("expected a []float64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_double_2(t *testing.T) {
	fnName := "test_fixedarray_output_double_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float64{1, 2}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float64); !ok {
		t.Errorf("expected a []float64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_double_3(t *testing.T) {
	fnName := "test_fixedarray_output_double_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float64{1, 2, 3}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float64); !ok {
		t.Errorf("expected a []float64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_double_4(t *testing.T) {
	fnName := "test_fixedarray_output_double_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float64{1, 2, 3, 4}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float64); !ok {
		t.Errorf("expected a []float64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_double_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_double_option_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float64{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float64); !ok {
		t.Errorf("expected a []*float64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_double_option_1_none(t *testing.T) {
	fnName := "test_fixedarray_output_double_option_1_none"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float64{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float64); !ok {
		t.Errorf("expected a []*float64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_double_option_1_some(t *testing.T) {
	fnName := "test_fixedarray_output_double_option_1_some"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float64{Ptr(float64(1))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float64); !ok {
		t.Errorf("expected a []*float64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_double_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_double_option_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float64{Ptr(float64(1)), Ptr(float64(2))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float64); !ok {
		t.Errorf("expected a []*float64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_double_option_3(t *testing.T) {
	fnName := "test_fixedarray_output_double_option_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	a := float64(11)
	c := float64(33)
	expected := []*float64{&a, nil, &c}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float64); !ok {
		t.Errorf("expected a []*float64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_double_option_4(t *testing.T) {
	fnName := "test_fixedarray_output_double_option_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float64{nil, Ptr(float64(2)), Ptr(float64(0)), Ptr(float64(4))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float64); !ok {
		t.Errorf("expected a []*float64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}
