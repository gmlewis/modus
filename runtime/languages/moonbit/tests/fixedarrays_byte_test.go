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

// func TestFixedArrayInput_byte(t *testing.T) {
// 	fnName := "test_fixedarray_input_byte"
// 	s := []byte{1, 2, 3, 4}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func TestFixedArrayOutput_byte_0(t *testing.T) {
	fnName := "test_fixedarray_output_byte_0"

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

func TestFixedArrayOutput_byte_1(t *testing.T) {
	fnName := "test_fixedarray_output_byte_1"

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

func TestFixedArrayOutput_byte_2(t *testing.T) {
	fnName := "test_fixedarray_output_byte_2"

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

func TestFixedArrayOutput_byte_3(t *testing.T) {
	fnName := "test_fixedarray_output_byte_3"

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

func TestFixedArrayOutput_byte_4(t *testing.T) {
	fnName := "test_fixedarray_output_byte_4"

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

func TestFixedArrayOutput_byte_option_0(t *testing.T) {
	fnName := "test_fixedarray_output_byte_option_0"

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

func TestFixedArrayOutput_byte_option_1(t *testing.T) {
	fnName := "test_fixedarray_output_byte_option_1"

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

func TestFixedArrayOutput_byte_option_2(t *testing.T) {
	fnName := "test_fixedarray_output_byte_option_2"

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

func TestFixedArrayOutput_byte_option_3(t *testing.T) {
	fnName := "test_fixedarray_output_byte_option_3"

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

func TestFixedArrayOutput_byte_option_4(t *testing.T) {
	fnName := "test_fixedarray_output_byte_option_4"

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
