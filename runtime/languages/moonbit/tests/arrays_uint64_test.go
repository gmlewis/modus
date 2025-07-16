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

func TestArrayOutput_uint64_0(t *testing.T) {
	fnName := "test_array_output_uint64_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint64{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint64); !ok {
		t.Errorf("expected a []uint64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_1(t *testing.T) {
	fnName := "test_array_output_uint64_1"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint64{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint64); !ok {
		t.Errorf("expected a []uint64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_1_min(t *testing.T) {
	fnName := "test_array_output_uint64_1_min"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint64{0}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint64); !ok {
		t.Errorf("expected a []uint64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_1_max(t *testing.T) {
	fnName := "test_array_output_uint64_1_max"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint64{math.MaxUint64}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint64); !ok {
		t.Errorf("expected a []uint64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_2(t *testing.T) {
	fnName := "test_array_output_uint64_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint64{0, math.MaxUint64}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint64); !ok {
		t.Errorf("expected a []uint64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_3(t *testing.T) {
	fnName := "test_array_output_uint64_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint64{1, 2, 3}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint64); !ok {
		t.Errorf("expected a []uint64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_4(t *testing.T) {
	fnName := "test_array_output_uint64_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint64{1, 2, 3, 4}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint64); !ok {
		t.Errorf("expected a []uint64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_option_0(t *testing.T) {
	fnName := "test_array_output_uint64_option_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint64{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint64); !ok {
		t.Errorf("expected a []*uint64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_option_1_none(t *testing.T) {
	fnName := "test_array_output_uint64_option_1_none"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint64{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint64); !ok {
		t.Errorf("expected a []*uint64, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_option_1_min(t *testing.T) {
	fnName := "test_array_output_uint64_option_1_min"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint64{Ptr(uint64(0))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint64); !ok {
		t.Errorf("expected a []*uint64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_option_1_max(t *testing.T) {
	fnName := "test_array_output_uint64_option_1_max"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint64{Ptr(uint64(math.MaxUint64))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint64); !ok {
		t.Errorf("expected a []*uint64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_option_2(t *testing.T) {
	fnName := "test_array_output_uint64_option_2"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint64{Ptr(uint64(1)), Ptr(uint64(2))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint64); !ok {
		t.Errorf("expected a []*uint64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_option_3(t *testing.T) {
	fnName := "test_array_output_uint64_option_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint64{nil, nil, nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint64); !ok {
		t.Errorf("expected a []*uint64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_uint64_option_4(t *testing.T) {
	fnName := "test_array_output_uint64_option_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint64{nil, Ptr(uint64(0)), Ptr(uint64(0)), Ptr(uint64(math.MaxUint64))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint64); !ok {
		t.Errorf("expected a []*uint64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}
