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

func TestFixedArrayOutput_byte_0(t *testing.T) {
	fnName := "test_fixedarray_output_byte_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 0 0 0 3]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_byte_1(t *testing.T) {
	fnName := "test_fixedarray_output_byte_1"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 1 0 0 2]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_byte_2(t *testing.T) {
	fnName := "test_fixedarray_output_byte_2"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 1 2 0 1]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_byte_3(t *testing.T) {
	fnName := "test_fixedarray_output_byte_3"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=246(FixedArray[Byte]), words=1, memBlock=[1 0 0 0 246 1 0 0 1 2 3 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_byte_4(t *testing.T) {
	fnName := "test_fixedarray_output_byte_4"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=246(FixedArray[Byte]), words=2, memBlock=[1 0 0 0 246 2 0 0 1 2 3 4 0 0 0 3]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_byte_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_byte_option_0"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 8=8+words*4), classID=241(FixedArray[Primitive]), words=0, memBlock=[1 0 0 0 241 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_byte_option_1(t *testing.T) {
	fnName := "test_fixedarray_output_byte_option_1"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 12=8+words*4), classID=241(FixedArray[Primitive]), words=1, memBlock=[1 0 0 0 241 1 0 0 1 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_byte_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_byte_option_2"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 16=8+words*4), classID=241(FixedArray[Primitive]), words=2, memBlock=[1 0 0 0 241 2 0 0 1 0 0 0 255 255 255 255]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_byte_option_3(t *testing.T) {
	fnName := "test_fixedarray_output_byte_option_3"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 20=8+words*4), classID=241(FixedArray[Primitive]), words=3, memBlock=[1 0 0 0 241 3 0 0 255 255 255 255 2 0 0 0 3 0 0 0]
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

	testInputSide(t, fnName, expected)
}

func TestFixedArrayOutput_byte_option_4(t *testing.T) {
	fnName := "test_fixedarray_output_byte_option_4"

	// memoryBlockAtOffset(offset: 91936=0x00016720=[32 103 1 0], size: 24=8+words*4), classID=241(FixedArray[Primitive]), words=4, memBlock=[1 0 0 0 241 4 0 0 1 0 0 0 2 0 0 0 3 0 0 0 255 255 255 255]
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

	testInputSide(t, fnName, expected)
}
