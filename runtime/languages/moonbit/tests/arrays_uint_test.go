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

func TestArrayInput_uint_option(t *testing.T) {
	fnName := "test_array_input_uint_option"
	s := getUIntOptionArray()

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_uint_option(t *testing.T) {
	fnName := "test_array_output_uint_option"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getUIntOptionArray()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint32); !ok {
		t.Errorf("expected a []*uint32, got %T", result)
	} else if !slices.EqualFunc(expected, r, func(a, b *uint32) bool { return (a == nil && b == nil) || (a != nil && b != nil && *a == *b) }) {
		for i, v := range expected {
			t.Logf("expected[%v]: %v", i, *v)
		}
		for i, v := range r {
			t.Logf("r[%v]: %v", i, *v)
		}
		t.Errorf("expected %#v, got %#v", expected, r)
	}
}

func getUIntOptionArray() []*uint32 {
	a := uint32(11)
	// b := 22
	c := uint32(33)
	return []*uint32{&a, nil, &c}
}

func TestArrayInput_uint_empty(t *testing.T) {
	fnName := "test_array_input_uint_empty"
	s := []uint32{}

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_uint_0(t *testing.T) {
	fnName := "test_array_output_uint_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint32{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint32); !ok {
		t.Errorf("expected a []uint32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_1(t *testing.T) {
	fnName := "test_array_output_uint_1"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint32{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint32); !ok {
		t.Errorf("expected a []uint32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_1_min(t *testing.T) {
	fnName := "test_array_output_uint_1_min"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint32{0}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint32); !ok {
		t.Errorf("expected a []uint32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_1_max(t *testing.T) {
	fnName := "test_array_output_uint_1_max"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint32{math.MaxUint32}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint32); !ok {
		t.Errorf("expected a []uint32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_2(t *testing.T) {
	fnName := "test_array_output_uint_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint32{0, math.MaxUint32}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint32); !ok {
		t.Errorf("expected a []uint32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_3(t *testing.T) {
	fnName := "test_array_output_uint_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint32{1, 2, 3}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint32); !ok {
		t.Errorf("expected a []uint32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_4(t *testing.T) {
	fnName := "test_array_output_uint_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint32{1, 2, 3, 4}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint32); !ok {
		t.Errorf("expected a []uint32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_option_0(t *testing.T) {
	fnName := "test_array_output_uint_option_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint32{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint32); !ok {
		t.Errorf("expected a []*uint32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_option_1_none(t *testing.T) {
	fnName := "test_array_output_uint_option_1_none"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint32{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint32); !ok {
		t.Errorf("expected a []*uint32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_option_1_min(t *testing.T) {
	fnName := "test_array_output_uint_option_1_min"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint32{Ptr(uint32(0))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint32); !ok {
		t.Errorf("expected a []*uint32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_option_1_max(t *testing.T) {
	fnName := "test_array_output_uint_option_1_max"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint32{Ptr(uint32(math.MaxUint32))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint32); !ok {
		t.Errorf("expected a []*uint32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_option_2(t *testing.T) {
	fnName := "test_array_output_uint_option_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint32{Ptr(uint32(1)), Ptr(uint32(2))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint32); !ok {
		t.Errorf("expected a []*uint32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_option_3(t *testing.T) {
	fnName := "test_array_output_uint_option_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint32{nil, nil, nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint32); !ok {
		t.Errorf("expected a []*uint32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint_option_4(t *testing.T) {
	fnName := "test_array_output_uint_option_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint32{nil, Ptr(uint32(0)), Ptr(uint32(0)), Ptr(uint32(math.MaxUint32))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint32); !ok {
		t.Errorf("expected a []*uint32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}
