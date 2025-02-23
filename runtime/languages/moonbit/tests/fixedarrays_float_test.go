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

func TestFixedArrayOutput_float_0(t *testing.T) {
	fnName := "test_fixedarray_output_float_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float32{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float32); !ok {
		t.Errorf("expected a []float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_1(t *testing.T) {
	fnName := "test_fixedarray_output_float_1"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float32{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float32); !ok {
		t.Errorf("expected a []float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_2(t *testing.T) {
	fnName := "test_fixedarray_output_float_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float32{1, 2}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float32); !ok {
		t.Errorf("expected a []float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_3(t *testing.T) {
	fnName := "test_fixedarray_output_float_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float32{1, 2, 3}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float32); !ok {
		t.Errorf("expected a []float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_4(t *testing.T) {
	fnName := "test_fixedarray_output_float_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float32{1, 2, 3, 4}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float32); !ok {
		t.Errorf("expected a []float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float32{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_1_none(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_1_none"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float32{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_1_some(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_1_some"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float32{Ptr(float32(1))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float32{Ptr(float32(1)), Ptr(float32(2))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_3(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	a := float32(1)
	c := float32(3)
	expected := []*float32{&a, nil, &c}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_4(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float32{nil, Ptr(float32(2)), Ptr(float32(0)), Ptr(float32(4))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %+v, got %+v", expected, r)
	}

	testInputSide(t, fnName, expected)
}
