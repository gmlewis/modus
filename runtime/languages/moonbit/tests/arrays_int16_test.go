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
	"math"
	"reflect"
	"slices"
	"testing"
)

func TestArrayInput_int16_option(t *testing.T) {
	fnName := "test_array_input_int16_option"
	s := getInt16OptionArray()

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_int16_option(t *testing.T) {
	fnName := "test_array_output_int16_option"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 11 0 0 0 255 255 255 255 33 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getInt16OptionArray()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !slices.EqualFunc(expected, r, func(a, b *int16) bool { return (a == nil && b == nil) || (a != nil && b != nil && *a == *b) }) {
		t.Errorf("expected %#v, got %#v", expected, r)
	}
}

func getInt16OptionArray() []*int16 {
	a := int16(11)
	// b := 22
	c := int16(33)
	return []*int16{&a, nil, &c}
}

func TestArrayInput_int16_empty(t *testing.T) {
	fnName := "test_array_input_int16_empty"
	s := []int16{}

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_int16_0(t *testing.T) {
	fnName := "test_array_output_int16_0"

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
}

func TestArrayOutput_int16_1(t *testing.T) {
	fnName := "test_array_output_int16_1"

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
}

func TestArrayOutput_int16_1_min(t *testing.T) {
	fnName := "test_array_output_int16_1_min"

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
}

func TestArrayOutput_int16_1_max(t *testing.T) {
	fnName := "test_array_output_int16_1_max"

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
}

func TestArrayOutput_int16_2(t *testing.T) {
	fnName := "test_array_output_int16_2"

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
}

func TestArrayOutput_int16_3(t *testing.T) {
	fnName := "test_array_output_int16_3"

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
}

func TestArrayOutput_int16_4(t *testing.T) {
	fnName := "test_array_output_int16_4"

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
}

func TestArrayOutput_int16_option_0(t *testing.T) {
	fnName := "test_array_output_int16_option_0"

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
}

func TestArrayOutput_int16_option_1_none(t *testing.T) {
	fnName := "test_array_output_int16_option_1_none"

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
}

func TestArrayOutput_int16_option_1_min(t *testing.T) {
	fnName := "test_array_output_int16_option_1_min"

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
}

func TestArrayOutput_int16_option_1_max(t *testing.T) {
	fnName := "test_array_output_int16_option_1_max"

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
}

func TestArrayOutput_int16_option_2(t *testing.T) {
	fnName := "test_array_output_int16_option_2"

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
}

func TestArrayOutput_int16_option_3(t *testing.T) {
	fnName := "test_array_output_int16_option_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{nil, nil, nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int16_option_4(t *testing.T) {
	fnName := "test_array_output_int16_option_4"

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
		t.Errorf("expected %v, got %v", expected, r)
	}
}
