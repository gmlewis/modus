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
	"reflect"
	"slices"
	"testing"
)

func TestFixedArrayOutput_int16_0(t *testing.T) {
	fnName := "test_fixedarray_output_int16_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_1(t *testing.T) {
	fnName := "test_fixedarray_output_int16_1"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_1_min(t *testing.T) {
	fnName := "test_fixedarray_output_int16_1_min"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{math.MinInt16}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_1_max(t *testing.T) {
	fnName := "test_fixedarray_output_int16_1_max"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{math.MaxInt16}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_2(t *testing.T) {
	fnName := "test_fixedarray_output_int16_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{math.MinInt16, math.MaxInt16}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_3(t *testing.T) {
	fnName := "test_fixedarray_output_int16_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{1, 2, 3}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_4(t *testing.T) {
	fnName := "test_fixedarray_output_int16_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{1, 2, 3, 4}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_int16_option_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_option_1_none(t *testing.T) {
	fnName := "test_fixedarray_output_int16_option_1_none"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_option_1_min(t *testing.T) {
	fnName := "test_fixedarray_output_int16_option_1_min"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{Ptr(int16(math.MinInt16))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_option_1_max(t *testing.T) {
	fnName := "test_fixedarray_output_int16_option_1_max"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{Ptr(int16(math.MaxInt16))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_int16_option_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{Ptr(int16(1)), Ptr(int16(2))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_option_3(t *testing.T) {
	fnName := "test_fixedarray_output_int16_option_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	a := int16(11)
	c := int16(33)
	expected := []*int16{&a, nil, &c}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_int16_option_4(t *testing.T) {
	fnName := "test_fixedarray_output_int16_option_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{nil, Ptr(int16(math.MinInt16)), Ptr(int16(0)), Ptr(int16(math.MaxInt16))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %+v, got %+v", expected, r)
	}

	testInputSide(t, fnName, expected)
}
