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
	"reflect"
	"testing"
)

// func TestArrayInput_byte(t *testing.T) {
// 	fnName := "test_array_input_byte"
// 	s := []byte{1, 2, 3, 4}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

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
