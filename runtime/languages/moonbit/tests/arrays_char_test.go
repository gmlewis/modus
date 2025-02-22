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

func TestArrayInput_char_option(t *testing.T) {
	fnName := "test_array_input_char_option"
	s := getCharOptionArray()

	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 49 0 0 0 255 255 255 255 51 0 0 0]
	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_char_option(t *testing.T) {
	fnName := "test_array_output_char_option"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 49 0 0 0 255 255 255 255 51 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getCharOptionArray()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !slices.EqualFunc(expected, r, func(a, b *int16) bool { return (a == nil && b == nil) || (a != nil && b != nil && *a == *b) }) {
		t.Errorf("expected %#v, got %#v", expected, r)
	}
}

func getCharOptionArray() []*int16 {
	a := int16('1')
	c := int16('3')
	return []*int16{&a, nil, &c}
}

func TestArrayInput_char_empty(t *testing.T) {
	fnName := "test_array_input_char_empty"
	s := []int16{}

	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 0 0 0 0]
	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_char_0(t *testing.T) {
	fnName := "test_array_output_char_0"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
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
}

func TestArrayOutput_char_1(t *testing.T) {
	fnName := "test_array_output_char_1"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 49 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{'1'}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_char_2(t *testing.T) {
	fnName := "test_array_output_char_2"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 2 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 49 0 0 0 50 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{'1', '2'}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_char_3(t *testing.T) {
	fnName := "test_array_output_char_3"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 49 0 0 0 50 0 0 0 51 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{'1', '2', '3'}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_char_4(t *testing.T) {
	fnName := "test_array_output_char_4"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 4 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 49 0 0 0 50 0 0 0 51 0 0 0 52 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int16{'1', '2', '3', '4'}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int16); !ok {
		t.Errorf("expected a []int16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_char_option_0(t *testing.T) {
	fnName := "test_array_output_char_option_0"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
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
}

func TestArrayOutput_char_option_1_none(t *testing.T) {
	fnName := "test_array_output_char_option_1_none"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 255]
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
}

func TestArrayOutput_char_option_1_some(t *testing.T) {
	fnName := "test_array_output_char_option_1_some"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 49 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{Ptr(int16('1'))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_char_option_2(t *testing.T) {
	fnName := "test_array_output_char_option_2"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 2 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 49 0 0 0 50 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{Ptr(int16('1')), Ptr(int16('2'))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_char_option_3(t *testing.T) {
	fnName := "test_array_output_char_option_3"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 255 255 255 255 255 255 255 255 255 255 255 255]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{nil, nil, nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_char_option_4(t *testing.T) {
	fnName := "test_array_output_char_option_4"

	// memoryBlockAtOffset(offset: 52320=0x0000CC60=[96 204 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 203 0 0 4 0 0 0]
	// memoryBlockAtOffset(offset: 52096=0x0000CB80=[128 203 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 255 255 255 255 50 0 0 0 0 0 0 0 52 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int16{nil, Ptr(int16('2')), Ptr(int16(0)), Ptr(int16('4'))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int16); !ok {
		t.Errorf("expected a []*int16, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}
