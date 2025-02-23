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

// func TestFixedArrayInput_int_option(t *testing.T) {
// 	fnName := "test_fixedarray_input_int_option"
// 	s := getIntOptionFixedArray()

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func TestFixedArrayOutput_int_option(t *testing.T) {
	fnName := "test_fixedarray_output_int_option"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getIntOptionFixedArray()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
	} else if !slices.EqualFunc(expected, r, func(a, b *int32) bool { return (a == nil && b == nil) || (a != nil && b != nil && *a == *b) }) {
		for i, v := range expected {
			t.Logf("expected[%v]: %v", i, *v)
		}
		for i, v := range r {
			t.Logf("r[%v]: %v", i, *v)
		}
		t.Errorf("expected %#v, got %#v", expected, r)
	}
}

func getIntOptionFixedArray() []*int32 {
	a := int32(11)
	c := int32(33)
	return []*int32{&a, nil, &c}
}

// func TestFixedArrayInput_int_empty(t *testing.T) {
// 	fnName := "test_fixedarray_input_int_empty"
// 	s := []int32{}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func TestFixedArrayOutput_int_0(t *testing.T) {
	fnName := "test_fixedarray_output_int_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int32{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int32); !ok {
		t.Errorf("expected a []int32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_1(t *testing.T) {
	fnName := "test_fixedarray_output_int_1"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int32{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int32); !ok {
		t.Errorf("expected a []int32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_1_min(t *testing.T) {
	fnName := "test_fixedarray_output_int_1_min"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int32{math.MinInt32}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int32); !ok {
		t.Errorf("expected a []int32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_1_max(t *testing.T) {
	fnName := "test_fixedarray_output_int_1_max"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int32{math.MaxInt32}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int32); !ok {
		t.Errorf("expected a []int32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_2(t *testing.T) {
	fnName := "test_fixedarray_output_int_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int32{math.MinInt32, math.MaxInt32}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int32); !ok {
		t.Errorf("expected a []int32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_3(t *testing.T) {
	fnName := "test_fixedarray_output_int_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int32{1, 2, 3}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int32); !ok {
		t.Errorf("expected a []int32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_4(t *testing.T) {
	fnName := "test_fixedarray_output_int_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int32{1, 2, 3, 4}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int32); !ok {
		t.Errorf("expected a []int32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_int_option_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int32{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_option_1_none(t *testing.T) {
	fnName := "test_fixedarray_output_int_option_1_none"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int32{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_option_1_min(t *testing.T) {
	fnName := "test_fixedarray_output_int_option_1_min"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int32{Ptr(int32(math.MinInt32))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_option_1_max(t *testing.T) {
	fnName := "test_fixedarray_output_int_option_1_max"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int32{Ptr(int32(math.MaxInt32))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_int_option_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int32{Ptr(int32(1)), Ptr(int32(2))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_option_3(t *testing.T) {
	fnName := "test_fixedarray_output_int_option_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int32{nil, nil, nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_int_option_4(t *testing.T) {
	fnName := "test_fixedarray_output_int_option_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int32{nil, Ptr(int32(math.MinInt32)), Ptr(int32(0)), Ptr(int32(math.MaxInt32))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}
