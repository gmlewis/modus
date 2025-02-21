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
	"bytes"
	"math"
	"reflect"
	"slices"
	"testing"
)

// Ptr is a helper routine that allocates a new T value
// to store v and returns a pointer to it.
func Ptr[T any](v T) *T {
	return &v
}

func TestArrayOutput_bool_0(t *testing.T) {
	fnName := "test_array_output_bool_0"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []bool{}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]bool); !ok {
		t.Errorf("expected a []bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_1(t *testing.T) {
	fnName := "test_array_output_bool_1"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []bool{true}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]bool); !ok {
		t.Errorf("expected a []bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_2(t *testing.T) {
	fnName := "test_array_output_bool_2"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 0 0 0 0 1 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []bool{false, true}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]bool); !ok {
		t.Errorf("expected a []bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_3(t *testing.T) {
	fnName := "test_array_output_bool_3"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 1 0 0 0 1 0 0 0 1 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []bool{true, true, true}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]bool); !ok {
		t.Errorf("expected a []bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_4(t *testing.T) {
	fnName := "test_array_output_bool_4"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 4 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []bool{false, false, false, false}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]bool); !ok {
		t.Errorf("expected a []bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_option_0(t *testing.T) {
	fnName := "test_array_output_bool_option_0"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*bool{}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_option_1_none(t *testing.T) {
	fnName := "test_array_output_bool_option_1_none"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 255]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*bool{nil}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_option_1_false(t *testing.T) {
	fnName := "test_array_output_bool_option_1_false"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0 := bool(false)
	expected := []*bool{&b0}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_option_1_true(t *testing.T) {
	fnName := "test_array_output_bool_option_1_true"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0 := bool(true)
	expected := []*bool{&b0}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_option_2(t *testing.T) {
	fnName := "test_array_output_bool_option_2"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 255 255 255 255]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 255 255 255 255]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [1])
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4294967295])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0 := bool(true)
	expected := []*bool{&b0, nil}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_option_3(t *testing.T) {
	fnName := "test_array_output_bool_option_3"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 255 255 255 255 1 0 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 255 255 255 255 1 0 0 0 1 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4294967295])
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [1])
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [1])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b1 := bool(true)
	expected := []*bool{nil, &b1, &b1}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_bool_option_4(t *testing.T) {
	fnName := "test_array_output_bool_option_4"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0, b1, b2 := bool(false), bool(true), bool(false)
	expected := []*bool{&b0, &b1, &b2, nil}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*bool); !ok {
		t.Errorf("expected a []*bool, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

// func TestArrayInput_byte(t *testing.T) {
// 	fnName := "test_array_input_byte"
// 	s := []byte{1, 2, 3, 4}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func TestArrayInput_int_option(t *testing.T) {
	fnName := "test_array_input_int_option"
	s := getIntOptionArray()

	// println: val: [Some(11), None, Some(33)]
	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[Int]), words=6, memBlock=[1 0 0 0 241 6 0 0 11 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 33 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 32 191 0 0 3 0 0 0]
	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayInput_string(t *testing.T) {
	fnName := "test_array_input_string"
	s := []string{"abc", "def", "ghi"}

	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 32 191 0 0 64 191 0 0 96 191 0 0]
	// memoryBlockAtOffset(offset: 49024=0x0000BF80=[128 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayInput_string_option(t *testing.T) {
	fnName := "test_array_input_string_option"
	s := getStringOptionArray()

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[1 0 0 0 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
	// memoryBlockAtOffset(offset: 48960=0x0000BF40=[64 191 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[1 0 0 0 243 2 0 0 103 0 104 0 105 0 0 1] = 'ghi'
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 32 191 0 0 0 0 0 0 64 191 0 0]
	// memoryBlockAtOffset(offset: 48992=0x0000BF60=[96 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_byte_0(t *testing.T) {
	fnName := "test_array_output_byte_0"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 0 0 0 3]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]byte); !ok {
		t.Errorf("expected a []byte, got %T", result)
	} else if !bytes.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_byte_1(t *testing.T) {
	fnName := "test_array_output_byte_1"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 9=8+words*4), moonBitType=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 1]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x01}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]byte); !ok {
		t.Errorf("expected a []byte, got %T", result)
	} else if !bytes.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_byte_2(t *testing.T) {
	fnName := "test_array_output_byte_2"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 2 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 10=8+words*4), moonBitType=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 1 2]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x01, 0x02}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]byte); !ok {
		t.Errorf("expected a []byte, got %T", result)
	} else if !bytes.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_byte_3(t *testing.T) {
	fnName := "test_array_output_byte_3"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 11=8+words*4), moonBitType=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 1 2 3]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x01, 0x02, 0x03}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]byte); !ok {
		t.Errorf("expected a []byte, got %T", result)
	} else if !bytes.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_byte_4(t *testing.T) {
	fnName := "test_array_output_byte_4"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 4 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=246(FixedArray[Byte]), words=2, memBlock=[1 0 0 0 246 2 0 0 1 2 3 4]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x01, 0x02, 0x03, 0x04}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]byte); !ok {
		t.Errorf("expected a []byte, got %T", result)
	} else if !bytes.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_byte_option_0(t *testing.T) {
	fnName := "test_array_output_byte_option_0"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*byte{}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*byte); !ok {
		t.Errorf("expected a []*byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_byte_option_1(t *testing.T) {
	fnName := "test_array_output_byte_option_1"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0 := byte(0x01)
	expected := []*byte{&b0}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*byte); !ok {
		t.Errorf("expected a []*byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_byte_option_2(t *testing.T) {
	fnName := "test_array_output_byte_option_2"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 255 255 255 255]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0 := byte(0x01)
	expected := []*byte{&b0, nil}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*byte); !ok {
		t.Errorf("expected a []*byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_byte_option_3(t *testing.T) {
	fnName := "test_array_output_byte_option_3"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 255 255 255 255 2 0 0 0 3 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b1, b2 := byte(0x02), byte(0x03)
	expected := []*byte{nil, &b1, &b2}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*byte); !ok {
		t.Errorf("expected a []*byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_byte_option_4(t *testing.T) {
	fnName := "test_array_output_byte_option_4"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 4 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 2 0 0 0 3 0 0 0 255 255 255 255]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	b0, b1, b2 := byte(0x01), byte(0x02), byte(0x03)
	expected := []*byte{&b0, &b1, &b2, nil}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*byte); !ok {
		t.Errorf("expected a []*byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int_option(t *testing.T) {
	fnName := "test_array_output_int_option"

	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 32 191 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 11 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 33 0 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getIntOptionArray()
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

func TestArrayOutput_string(t *testing.T) {
	fnName := "test_array_output_string"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 32 59 0 0 208 62 0 0 104 75 0 0]
	// memoryBlockAtOffset(offset: 15136=0x00003B20=[32 59 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
	// memoryBlockAtOffset(offset: 16080=0x00003ED0=[208 62 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 100 0 101 0 102 0 0 1] = 'def'
	// memoryBlockAtOffset(offset: 19304=0x00004B68=[104 75 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 103 0 104 0 105 0 0 1] = 'ghi'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"abc", "def", "ghi"}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_string_option(t *testing.T) {
	fnName := "test_array_output_string_option"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 32 59 0 0 0 0 0 0 104 75 0 0]
	// memoryBlockAtOffset(offset: 15136=0x00003B20=[32 59 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
	// memoryBlockAtOffset(offset: 19304=0x00004B68=[104 75 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 103 0 104 0 105 0 0 1] = 'ghi'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getStringOptionArray()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !slices.EqualFunc(expected, r, func(a, b *string) bool { return (a == nil && b == nil) || (a != nil && b != nil && *a == *b) }) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func getIntOptionArray() []*int32 {
	a := int32(11)
	// b := 22
	c := int32(33)
	return []*int32{&a, nil, &c}
}

func getStringOptionArray() []*string {
	a := "abc"
	// b := "def"
	c := "ghi"
	return []*string{&a, nil, &c}
}

func TestArrayInput_string_none(t *testing.T) {
	fnName := "test_array_input_string_none"

	// memoryBlockAtOffset(offset: 0) = (data=0, size=0)
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_string_none(t *testing.T) {
	fnName := "test_array_output_string_none"

	// pointerHandler.Decode(vals: [0])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

func TestArrayInput_string_empty(t *testing.T) {
	fnName := "test_array_input_string_empty"
	s := []string{}

	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_string_empty(t *testing.T) {
	fnName := "test_array_output_string_empty"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=242(), words=0, memBlock=[1 0 0 0 242 0 0 0]
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
}

// func TestArrayInput_int_empty(t *testing.T) {
// 	fnName := "test_array_input_int_empty"
// 	s := []int32{}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func TestArrayOutput_int_0(t *testing.T) {
	fnName := "test_array_output_int_0"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
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

func TestArrayOutput_int_1(t *testing.T) {
	fnName := "test_array_output_int_1"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
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

func TestArrayOutput_int_1_min(t *testing.T) {
	fnName := "test_array_output_int_1_min"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 128]
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

func TestArrayOutput_int_1_max(t *testing.T) {
	fnName := "test_array_output_int_1_max"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 127]
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

func TestArrayOutput_int_2(t *testing.T) {
	fnName := "test_array_output_int_2"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 2 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []int32{1, 2}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]int32); !ok {
		t.Errorf("expected a []int32, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_int_3(t *testing.T) {
	fnName := "test_array_output_int_3"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 1 0 0 0 2 0 0 0 3 0 0 0]
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

func TestArrayOutput_int_4(t *testing.T) {
	fnName := "test_array_output_int_4"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 4 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 2 0 0 0 3 0 0 0 4 0 0 0]
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

func TestArrayOutput_int_option_0(t *testing.T) {
	fnName := "test_array_output_int_option_0"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
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

func TestArrayOutput_int_option_1_none(t *testing.T) {
	fnName := "test_array_output_int_option_1_none"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0 1 0 0 0]
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

func TestArrayOutput_int_option_1_min(t *testing.T) {
	fnName := "test_array_output_int_option_1_min"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 128]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 128 255 255 255 255]
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

func TestArrayOutput_int_option_1_max(t *testing.T) {
	fnName := "test_array_output_int_option_1_max"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 127]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 255 255 255 127 0 0 0 0]
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

func TestArrayOutput_int_option_2(t *testing.T) {
	fnName := "test_array_output_int_option_2"

	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 190 0 0 2 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0]
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

func TestArrayOutput_int_option_3(t *testing.T) {
	fnName := "test_array_output_int_option_3"

	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 20=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 0 0 0 0 1 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 32=8+words*4), moonBitType=241(FixedArray[Int]), words=3, memBlock=[1 0 0 0 241 3 0 0 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0]
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

func TestArrayOutput_int_option_4(t *testing.T) {
	fnName := "test_array_output_int_option_4"

	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 160 191 0 0 4 0 0 0]
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 40=8+words*4), moonBitType=241(FixedArray[Int]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0 3 0 0 0 0 0 0 0 4 0 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int32{Ptr(int32(1)), Ptr(int32(2)), Ptr(int32(3)), Ptr(int32(4))}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

// func Test2DArrayInput_string(t *testing.T) {
// 	fnName := "test2d_array_input_string"
// 	s := [][]string{
// 		{"abc", "def", "ghi"},
// 		{"jkl", "mno", "pqr"},
// 		{"stu", "vwx", "yz"},
// 	}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func Test2DArrayOutput_string(t *testing.T) {
	fnName := "test2d_array_output_string"

	// memoryBlockAtOffset(offset: 49120=0x0000BFE0=[224 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 192 191 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 49088=0x0000BFC0=[192 191 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 32 191 0 0 96 191 0 0 160 191 0 0]
	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 32 59 0 0 208 62 0 0 104 75 0 0]
	// memoryBlockAtOffset(offset: 15136=0x00003B20=[32 59 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
	// memoryBlockAtOffset(offset: 16080=0x00003ED0=[208 62 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 100 0 101 0 102 0 0 1] = 'def'
	// memoryBlockAtOffset(offset: 19304=0x00004B68=[104 75 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 103 0 104 0 105 0 0 1] = 'ghi'
	// memoryBlockAtOffset(offset: 48992=0x0000BF60=[96 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 191 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 48960=0x0000BF40=[64 191 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 48 81 0 0 24 81 0 0 0 81 0 0]
	// memoryBlockAtOffset(offset: 20784=0x00005130=[48 81 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 106 0 107 0 108 0 0 1] = 'jkl'
	// memoryBlockAtOffset(offset: 20760=0x00005118=[24 81 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 109 0 110 0 111 0 0 1] = 'mno'
	// memoryBlockAtOffset(offset: 20736=0x00005100=[0 81 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 112 0 113 0 114 0 0 1] = 'pqr'
	// memoryBlockAtOffset(offset: 49056=0x0000BFA0=[160 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 128 191 0 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 49024=0x0000BF80=[128 191 0 0], size: 20=8+words*4), moonBitType=242(), words=3, memBlock=[1 0 0 0 242 3 0 0 232 80 0 0 208 80 0 0 184 80 0 0]
	// memoryBlockAtOffset(offset: 20712=0x000050E8=[232 80 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 115 0 116 0 117 0 0 1] = 'stu'
	// memoryBlockAtOffset(offset: 20688=0x000050D0=[208 80 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 118 0 119 0 120 0 0 1] = 'vwx'
	// memoryBlockAtOffset(offset: 20664=0x000050B8=[184 80 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 121 0 122 0 0 0 0 3] = 'yz'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [][]string{
		{"abc", "def", "ghi"},
		{"jkl", "mno", "pqr"},
		{"stu", "vwx", "yz"},
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([][]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func Test2DArrayInput_string_none(t *testing.T) {
	fnName := "test2d_array_input_string_none"

	// memoryBlockAtOffset(offset: 0) = (data=0, size=0)
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func Test2DArrayOutput_string_none(t *testing.T) {
	fnName := "test2d_array_output_string_none"

	// pointerHandler.Decode(vals: [0])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

func Test2DArrayInput_string_empty(t *testing.T) {
	fnName := "test2d_array_input_string_empty"
	s := [][]string{}

	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func Test2DArrayOutput_string_empty(t *testing.T) {
	fnName := "test2d_array_output_string_empty"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=242(), words=0, memBlock=[1 0 0 0 242 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [][]string{}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([][]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func Test2DArrayInput_string_inner_empty(t *testing.T) {
	fnName := "test2d_array_input_string_inner_empty"
	s := [][]string{{}}

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 8=8+words*4), moonBitType=241(FixedArray[Int]), words=0, memBlock=[1 0 0 0 241 0 0 0]
	// memoryBlockAtOffset(offset: 48944=0x0000BF30=[48 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 32 191 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 48 191 0 0]
	// memoryBlockAtOffset(offset: 48976=0x0000BF50=[80 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 1 0 0 0]
	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func Test2DArrayOutput_string_inner_empty(t *testing.T) {
	fnName := "test2d_array_output_string_inner_empty"

	// memoryBlockAtOffset(offset: 48960=0x0000BF40=[64 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 80 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48720=0x0000BE50=[80 190 0 0], size: 12=8+words*4), moonBitType=242(), words=1, memBlock=[1 0 0 0 242 1 0 0 32 191 0 0]
	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 8=8+words*4), moonBitType=242(), words=0, memBlock=[1 0 0 0 242 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [][]string{{}}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([][]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func Test2DArrayInput_string_inner_none(t *testing.T) {
	fnName := "test2d_array_input_string_inner_none"
	s := []*[]string{nil}

	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 1 0 0 0]
	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func Test2DArrayOutput_string_inner_none(t *testing.T) {
	fnName := "test2d_array_output_string_inner_none"

	// memoryBlockAtOffset(offset: 48928=0x0000BF20=[32 191 0 0], size: 16=8+words*4), moonBitType=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 64 190 0 0 1 0 0 0]
	// memoryBlockAtOffset(offset: 48704=0x0000BE40=[64 190 0 0], size: 12=8+words*4), moonBitType=242(), words=1, memBlock=[1 0 0 0 242 1 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*[]string{nil}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*[]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}
