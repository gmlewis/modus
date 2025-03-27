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

func TestFixedArrayOutput_uint64_0(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=241(FixedArray[Primitive]), words=0, memBlock=[1 0 0 0 241 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_1(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_1"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0 0 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_1_min(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_1_min"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0 0 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_1_max(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_1_max"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 255 255 255 255 255]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_2(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_2"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 24=8+words*4), classID=241(FixedArray[Primitive]), words=2, memBlock=[1 0 0 0 241 2 0 0 0 0 0 0 0 0 0 0 255 255 255 255 255 255 255 255]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_3(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_3"

	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 32=8+words*4), classID=241(FixedArray[Primitive]), words=3, memBlock=[1 0 0 0 241 3 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0 3 0 0 0 0 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_4(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_4"

	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 40=8+words*4), classID=241(FixedArray[Primitive]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0 3 0 0 0 0 0 0 0 4 0 0 0 0 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_option_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=242(FixedArray[String]), words=0, memBlock=[1 0 0 0 242 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_option_1_none(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_option_1_none"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 152 41 0 0]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), classID=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_option_1_min(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_option_1_min"

	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 12=8+words*4), classID=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 32 103 1 0]
	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_option_1_max(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_option_1_max"

	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 12=8+words*4), classID=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 32 103 1 0]
	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 255 255 255 255 255 255 255 255]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_option_2"

	// memoryBlockAtOffset(offset: 92192=0x00016820=[32 104 1 0], size: 16=8+words*4), classID=242(FixedArray[String]), words=2, memBlock=[1 0 0 0 242 2 0 0 32 103 1 0 0 104 1 0]
	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 1 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 2 0 0 0 0 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_option_3(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_option_3"

	// memoryBlockAtOffset(offset: 92192=0x00016820=[32 104 1 0], size: 20=8+words*4), classID=242(FixedArray[String]), words=3, memBlock=[1 0 0 0 242 3 0 0 32 103 1 0 152 41 0 0 0 104 1 0]
	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 11 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), classID=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 33 0 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	a := uint64(11)
	c := uint64(33)
	expected := []*uint64{&a, nil, &c}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint64); !ok {
		t.Errorf("expected a []*uint64, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_uint64_option_4(t *testing.T) {
	fnName := "test_fixedarray_output_uint64_option_4"

	// memoryBlockAtOffset(offset: 92224=0x00016840=[64 104 1 0], size: 24=8+words*4), classID=242(FixedArray[String]), words=4, memBlock=[1 0 0 0 242 4 0 0 152 41 0 0 32 103 1 0 0 104 1 0 32 104 1 0]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), classID=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 0 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 92192=0x00016820=[32 104 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 255 255 255 255 255 255 255 255]
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
		t.Errorf("expected %+v, got %+v", expected, r)
	}

	testInputSide(t, fnName, expected)
}
