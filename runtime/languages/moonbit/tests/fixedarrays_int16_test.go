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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[1 0 0 0 243 1 0 0 0 0 0 3]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[1 0 0 0 243 1 0 0 1 0 0 1]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[1 0 0 0 243 1 0 0 0 128 0 1]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[1 0 0 0 243 1 0 0 255 127 0 1]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[1 0 0 0 243 2 0 0 0 128 255 127 0 0 0 3]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[1 0 0 0 243 2 0 0 1 0 2 0 3 0 0 1]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 20=8+words*4), classID=243(String), words=3, memBlock=[1 0 0 0 243 3 0 0 1 0 2 0 3 0 4 0 0 0 0 3]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=241(FixedArray[Primitive]), words=0, memBlock=[1 0 0 0 241 0 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 255]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 128 255 255]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 127 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 2 0 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 20=8+words*4), classID=241(FixedArray[Primitive]), words=3, memBlock=[1 0 0 0 241 3 0 0 11 0 0 0 255 255 255 255 33 0 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 24=8+words*4), classID=241(FixedArray[Primitive]), words=4, memBlock=[1 0 0 0 241 4 0 0 255 255 255 255 0 128 255 255 0 0 0 0 255 127 0 0]
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
