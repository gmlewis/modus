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

func TestFixedArrayOutput_uint16_0(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[1 0 0 0 243 1 0 0 0 0 0 3] = ''
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint16{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint16); !ok {
		t.Errorf("expected a []uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_1(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_1"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[1 0 0 0 243 1 0 0 1 0 0 1] = ''
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint16{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint16); !ok {
		t.Errorf("expected a []uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_1_min(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_1_min"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[1 0 0 0 243 1 0 0 0 0 0 1] = ''
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint16{0}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint16); !ok {
		t.Errorf("expected a []uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_1_max(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_1_max"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[1 0 0 0 243 1 0 0 255 255 0 1] = '￿'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint16{math.MaxUint16}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint16); !ok {
		t.Errorf("expected a []uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_2(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_2"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[1 0 0 0 243 2 0 0 0 0 255 255 0 0 0 3] = '￿'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint16{0, math.MaxUint16}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint16); !ok {
		t.Errorf("expected a []uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_3(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_3"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[1 0 0 0 243 2 0 0 1 0 2 0 3 0 0 1] = ''
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint16{1, 2, 3}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint16); !ok {
		t.Errorf("expected a []uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_4(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_4"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 20=8+words*4), classID=243(String), words=3, memBlock=[1 0 0 0 243 3 0 0 1 0 2 0 3 0 4 0 0 0 0 3] = ''
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint16{1, 2, 3, 4}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint16); !ok {
		t.Errorf("expected a []uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_option_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=241(FixedArray[Primitive]), words=0, memBlock=[1 0 0 0 241 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint16{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint16); !ok {
		t.Errorf("expected a []*uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_option_1_none(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_option_1_none"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 255]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint16{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint16); !ok {
		t.Errorf("expected a []*uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_option_1_min(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_option_1_min"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint16{Ptr(uint16(0))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint16); !ok {
		t.Errorf("expected a []*uint16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_option_1_max(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_option_1_max"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint16{Ptr(uint16(math.MaxUint16))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint16); !ok {
		t.Errorf("expected a []*uint16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_option_2"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 2 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint16{Ptr(uint16(1)), Ptr(uint16(2))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint16); !ok {
		t.Errorf("expected a []*uint16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_option_3(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_option_3"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 20=8+words*4), classID=241(FixedArray[Primitive]), words=3, memBlock=[1 0 0 0 241 3 0 0 11 0 0 0 255 255 255 255 33 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	a := uint16(11)
	c := uint16(33)
	expected := []*uint16{&a, nil, &c}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint16); !ok {
		t.Errorf("expected a []*uint16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint16_option_4(t *testing.T) {
	fnName := "test_fixedarray_output_uint16_option_4"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 24=8+words*4), classID=241(FixedArray[Primitive]), words=4, memBlock=[1 0 0 0 241 4 0 0 255 255 255 255 0 0 0 0 0 0 0 0 255 255 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint16{nil, Ptr(uint16(0)), Ptr(uint16(0)), Ptr(uint16(math.MaxUint16))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint16); !ok {
		t.Errorf("expected a []*uint16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %+v, got %+v", expected, r)
	}

	testInputSide(t, fnName, expected)
}
