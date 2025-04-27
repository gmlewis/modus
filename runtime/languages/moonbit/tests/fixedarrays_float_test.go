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
	"reflect"
	"slices"
	"testing"
)

func TestFixedArrayOutput_float_0(t *testing.T) {
	fnName := "test_fixedarray_output_float_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=241(FixedArray[Primitive]), words=0, memBlock=[1 0 0 0 241 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float32{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float32); !ok {
		t.Errorf("expected a []float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_1(t *testing.T) {
	fnName := "test_fixedarray_output_float_1"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 128 63]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float32{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float32); !ok {
		t.Errorf("expected a []float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_2(t *testing.T) {
	fnName := "test_fixedarray_output_float_2"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=2, memBlock=[1 0 0 0 241 2 0 0 0 0 128 63 0 0 0 64]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float32{1, 2}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float32); !ok {
		t.Errorf("expected a []float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_3(t *testing.T) {
	fnName := "test_fixedarray_output_float_3"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 20=8+words*4), classID=241(FixedArray[Primitive]), words=3, memBlock=[1 0 0 0 241 3 0 0 0 0 128 63 0 0 0 64 0 0 64 64]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float32{1, 2, 3}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float32); !ok {
		t.Errorf("expected a []float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_4(t *testing.T) {
	fnName := "test_fixedarray_output_float_4"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 24=8+words*4), classID=241(FixedArray[Primitive]), words=4, memBlock=[1 0 0 0 241 4 0 0 0 0 128 63 0 0 0 64 0 0 64 64 0 0 128 64]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []float32{1, 2, 3, 4}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]float32); !ok {
		t.Errorf("expected a []float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=242(FixedArray[String]), words=0, memBlock=[1 0 0 0 242 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float32{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_1_none(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_1_none"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 152 41 0 0]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), classID=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float32{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_1_some(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_1_some"

	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 12=8+words*4), classID=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 32 103 1 0]
	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 0 0 128 63]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float32{Ptr(float32(1))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_2"

	// memoryBlockAtOffset(offset: 92192=0x00016820=[32 104 1 0], size: 16=8+words*4), classID=242(FixedArray[String]), words=2, memBlock=[1 0 0 0 242 2 0 0 32 103 1 0 0 104 1 0]
	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 0 0 128 63]
	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 12=8+words*4), classID=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 0 0 0 64]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float32{Ptr(float32(1)), Ptr(float32(2))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_3(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_3"

	// memoryBlockAtOffset(offset: 92192=0x00016820=[32 104 1 0], size: 20=8+words*4), classID=242(FixedArray[String]), words=3, memBlock=[1 0 0 0 242 3 0 0 32 103 1 0 152 41 0 0 0 104 1 0]
	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 0 0 48 65]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), classID=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 12=8+words*4), classID=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 0 0 4 66]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	a := float32(11)
	c := float32(33)
	expected := []*float32{&a, nil, &c}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_float_option_4(t *testing.T) {
	fnName := "test_fixedarray_output_float_option_4"

	// memoryBlockAtOffset(offset: 92224=0x00016840=[64 104 1 0], size: 24=8+words*4), classID=242(FixedArray[String]), words=4, memBlock=[1 0 0 0 242 4 0 0 152 41 0 0 32 103 1 0 0 104 1 0 32 104 1 0]
	// memoryBlockAtOffset(offset: 10648=0x00002998=[152 41 0 0], size: 8=8+words*4), classID=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 0 0 0 64]
	// memoryBlockAtOffset(offset: 92160=0x00016800=[0 104 1 0], size: 12=8+words*4), classID=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 92192=0x00016820=[32 104 1 0], size: 12=8+words*4), classID=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 0 0 128 64]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*float32{nil, Ptr(float32(2)), Ptr(float32(0)), Ptr(float32(4))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*float32); !ok {
		t.Errorf("expected a []*float32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %+v, got %+v", expected, r)
	}

	testInputSide(t, fnName, expected)
}
