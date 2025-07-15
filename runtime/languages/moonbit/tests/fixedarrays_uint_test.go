// -*- compile-command: "NO_COLOR=1 go test -timeout 30s -tags integration -run '^(TestFixedArrayOutput_uint_|TestFixedArrayInput_uint_)' ."; -*-

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

func TestFixedArrayOutput_uint_0(t *testing.T) {
	fnName := "test_fixedarray_output_uint_0"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_1(t *testing.T) {
	fnName := "test_fixedarray_output_uint_1"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_1_min(t *testing.T) {
	fnName := "test_fixedarray_output_uint_1_min"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_1_max(t *testing.T) {
	fnName := "test_fixedarray_output_uint_1_max"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_2(t *testing.T) {
	fnName := "test_fixedarray_output_uint_2"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_3(t *testing.T) {
	fnName := "test_fixedarray_output_uint_3"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_4(t *testing.T) {
	fnName := "test_fixedarray_output_uint_4"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_uint_option_0"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_option_1_none(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_fixedarray_output_uint_option_1_none"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_option_1_min(t *testing.T) {
	fnName := "test_fixedarray_output_uint_option_1_min"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_option_1_max(t *testing.T) {
	fnName := "test_fixedarray_output_uint_option_1_max"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_uint_option_2"

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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_option_3(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_fixedarray_output_uint_option_3"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	a := uint32(11)
	c := uint32(33)
	expected := []*uint32{&a, nil, &c}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint32); !ok {
		t.Errorf("expected a []*uint32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint_option_4(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_fixedarray_output_uint_option_4"

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
		t.Errorf("expected %+v, got %+v", expected, r)
	}

	testInputSide(t, fnName, expected)
}
