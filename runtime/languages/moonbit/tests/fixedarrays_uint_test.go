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

func TestFixedArrayOutput_uint_0(t *testing.T) {
	fnName := "test_fixedarray_output_uint_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=241(FixedArray[Primitive]), words=0, memBlock=[1 0 0 0 241 0 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 255]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=2, memBlock=[1 0 0 0 241 2 0 0 0 0 0 0 255 255 255 255]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 20=8+words*4), classID=241(FixedArray[Primitive]), words=3, memBlock=[1 0 0 0 241 3 0 0 1 0 0 0 2 0 0 0 3 0 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 24=8+words*4), classID=241(FixedArray[Primitive]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 2 0 0 0 3 0 0 0 4 0 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=241(FixedArray[Primitive]), words=0, memBlock=[1 0 0 0 241 0 0 0]
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
	fnName := "test_fixedarray_output_uint_option_1_none"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0 1 0 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0 0 0 0 0]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 255 255 255 255 255]
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

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 24=8+words*4), classID=241(FixedArray[Primitive]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0]
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
	fnName := "test_fixedarray_output_uint_option_3"

	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 32=8+words*4), classID=241(FixedArray[Primitive]), words=3, memBlock=[1 0 0 0 241 3 0 0 11 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 33 0 0 0 0 0 0 0]
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
	fnName := "test_fixedarray_output_uint_option_4"

	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 40=8+words*4), classID=241(FixedArray[Primitive]), words=4, memBlock=[1 0 0 0 241 4 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 255 255 255 255 255 255 255 255]
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
