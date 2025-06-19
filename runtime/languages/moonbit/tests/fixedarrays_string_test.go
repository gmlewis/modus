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
	"reflect"
	"slices"
	"testing"
)

func TestFixedArrayOutput_string_0(t *testing.T) {
	fnName := "test_fixedarray_output_string_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=242(FixedArray[String]), words=0, memBlock=[1 0 0 0 242 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_string_1(t *testing.T) {
	fnName := "test_fixedarray_output_string_1"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 248 39 0 0]
	// memoryBlockAtOffset(offset: 10232=0x000027F8=[248 39 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 49 0 0 1] = '1'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"1"}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_string_2(t *testing.T) {
	fnName := "test_fixedarray_output_string_2"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=242(FixedArray[String]), words=2, memBlock=[1 0 0 0 242 2 0 0 248 39 0 0 8 40 0 0]
	// memoryBlockAtOffset(offset: 10232=0x000027F8=[248 39 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 49 0 0 1] = '1'
	// memoryBlockAtOffset(offset: 10248=0x00002808=[8 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 50 0 0 1] = '2'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"1", "2"}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_string_3(t *testing.T) {
	fnName := "test_fixedarray_output_string_3"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 20=8+words*4), classID=242(FixedArray[String]), words=3, memBlock=[1 0 0 0 242 3 0 0 248 39 0 0 8 40 0 0 24 40 0 0]
	// memoryBlockAtOffset(offset: 10232=0x000027F8=[248 39 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 49 0 0 1] = '1'
	// memoryBlockAtOffset(offset: 10248=0x00002808=[8 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 50 0 0 1] = '2'
	// memoryBlockAtOffset(offset: 10264=0x00002818=[24 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 51 0 0 1] = '3'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"1", "2", "3"}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_string_4(t *testing.T) {
	fnName := "test_fixedarray_output_string_4"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 24=8+words*4), classID=242(FixedArray[String]), words=4, memBlock=[1 0 0 0 242 4 0 0 248 39 0 0 8 40 0 0 24 40 0 0 40 40 0 0]
	// memoryBlockAtOffset(offset: 10232=0x000027F8=[248 39 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 49 0 0 1] = '1'
	// memoryBlockAtOffset(offset: 10248=0x00002808=[8 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 50 0 0 1] = '2'
	// memoryBlockAtOffset(offset: 10264=0x00002818=[24 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 51 0 0 1] = '3'
	// memoryBlockAtOffset(offset: 10280=0x00002828=[40 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 52 0 0 1] = '4'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"1", "2", "3", "4"}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_string_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_string_option_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=242(FixedArray[String]), words=0, memBlock=[1 0 0 0 242 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*string{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_string_option_1_none(t *testing.T) {
	fnName := "test_fixedarray_output_string_option_1_none"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*string{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_string_option_1_some(t *testing.T) {
	fnName := "test_fixedarray_output_string_option_1_some"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=242(FixedArray[String]), words=1, memBlock=[1 0 0 0 242 1 0 0 248 39 0 0]
	// memoryBlockAtOffset(offset: 10232=0x000027F8=[248 39 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 49 0 0 1] = '1'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*string{Ptr(string("1"))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_string_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_string_option_2"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=242(FixedArray[String]), words=2, memBlock=[1 0 0 0 242 2 0 0 248 39 0 0 8 40 0 0]
	// memoryBlockAtOffset(offset: 10232=0x000027F8=[248 39 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 49 0 0 1] = '1'
	// memoryBlockAtOffset(offset: 10248=0x00002808=[8 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 50 0 0 1] = '2'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*string{Ptr(string("1")), Ptr(string("2"))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_string_option_3(t *testing.T) {
	fnName := "test_fixedarray_output_string_option_3"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 20=8+words*4), classID=242(FixedArray[String]), words=3, memBlock=[1 0 0 0 242 3 0 0 128 250 0 0 0 0 0 0 104 250 0 0]
	// memoryBlockAtOffset(offset: 64128=0x0000FA80=[128 250 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 49 0 49 0 0 0 0 3] = '11'
	// stringHandler.Read(offset: 0=0x00000000=[0 0 0 0])
	// memoryBlockAtOffset(offset: 64104=0x0000FA68=[104 250 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 51 0 51 0 0 0 0 3] = '33'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	a := string("11")
	c := string("33")
	expected := []*string{&a, nil, &c}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_string_option_4(t *testing.T) {
	fnName := "test_fixedarray_output_string_option_4"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 24=8+words*4), classID=242(FixedArray[String]), words=4, memBlock=[1 0 0 0 242 4 0 0 0 0 0 0 8 40 0 0 192 39 0 0 40 40 0 0]
	// memoryBlockAtOffset(offset: 10248=0x00002808=[8 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 50 0 0 1] = '2'
	// memoryBlockAtOffset(offset: 10176=0x000027C0=[192 39 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 48 0 0 1] = '0'
	// memoryBlockAtOffset(offset: 10280=0x00002828=[40 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 52 0 0 1] = '4'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*string{nil, Ptr(string("2")), Ptr(string("0")), Ptr(string("4"))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}

	testInputSide(t, fnName, expected)
}
