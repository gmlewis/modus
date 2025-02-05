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
	"slices"
	"testing"
)

func TestArrayInput_byte(t *testing.T) {
	fnName := "test_array_input_byte"
	s := []byte{1, 2, 3, 4}

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayInput_intPtr(t *testing.T) {
	fnName := "test_array_input_int_ptr"
	s := getIntPtrArray()

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayInput_string(t *testing.T) {
	fnName := "test_array_input_string"
	s := []string{"abc", "def", "ghi"}

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayInput_stringPtr(t *testing.T) {
	fnName := "test_array_input_string_ptr"
	s := getStringPtrArray()

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_byte(t *testing.T) {
	fnName := "test_array_output_byte"
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

func TestArrayOutput_intPtr(t *testing.T) {
	fnName := "test_array_output_int_ptr"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getIntPtrArray()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int); !ok {
		t.Errorf("expected a []*int, got %T", result)
	} else if !slices.EqualFunc(expected, r, func(a, b *int) bool { return *a == *b }) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestArrayOutput_string(t *testing.T) {
	fnName := "test_array_output_string"
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

func TestArrayOutput_stringPtr(t *testing.T) {
	fnName := "test_array_output_string_ptr"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getStringPtrArray()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !slices.EqualFunc(expected, r, func(a, b *string) bool { return *a == *b }) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func getIntPtrArray() []*int {
	a := 11
	b := 22
	c := 33
	return []*int{&a, &b, &c}
}

func getStringPtrArray() []*string {
	a := "abc"
	b := "def"
	c := "ghi"
	return []*string{&a, &b, &c}
}

func TestArrayInput_string_none(t *testing.T) {
	fnName := "test_array_input_string_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_string_none(t *testing.T) {
	fnName := "test_array_output_string_none"
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

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_string_empty(t *testing.T) {
	fnName := "test_array_output_string_empty"
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

func TestArrayInput_int32_empty(t *testing.T) {
	fnName := "test_array_input_int32_empty"
	s := []int32{}

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func TestArrayOutput_int32_empty(t *testing.T) {
	fnName := "test_array_output_int32_empty"
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

func Test2DArrayInput_string(t *testing.T) {
	fnName := "test2d_array_input_string"
	s := [][]string{
		{"abc", "def", "ghi"},
		{"jkl", "mno", "pqr"},
		{"stu", "vwx", "yz"},
	}

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func Test2DArrayOutput_string(t *testing.T) {
	fnName := "test2d_array_output_string"
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
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func Test2DArrayOutput_string_none(t *testing.T) {
	fnName := "test2d_array_output_string_none"
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

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func Test2DArrayOutput_string_empty(t *testing.T) {
	fnName := "test2d_array_output_string_empty"
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

func Test2DArrayInput_string_innerNil(t *testing.T) {
	fnName := "test2d_array_input_string_inner_none"
	s := [][]string{nil} // []*[]string{nil}  ?

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func Test2DArrayOutput_string_innerNil(t *testing.T) {
	fnName := "test2d_array_output_string_inner_none"
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
