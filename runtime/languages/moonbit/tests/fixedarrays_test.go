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
	"testing"

	"github.com/gmlewis/modus/runtime/utils"
)

func TestFixedArrayInput0_string(t *testing.T) {
	fnName := "test_fixedarray_input0_string"
	arr := [0]string{}

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayInput0_stringPtr(t *testing.T) {
	fnName := "test_fixedarray_input0_string_option"
	arr := [0]*string{}

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayOutput0_string(t *testing.T) {
	fnName := "test_fixedarray_output0_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [0]string{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([0]string); !ok {
		t.Errorf("expected a [0]string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput0_stringPtr(t *testing.T) {
	fnName := "test_fixedarray_output0_string_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [0]*string{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([0]*string); !ok {
		t.Errorf("expected a [0]*string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayInput0_intPtr(t *testing.T) {
	fnName := "test_fixedarray_input0_int_option"
	arr := [0]*int{}

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayOutput0_intPtr(t *testing.T) {
	fnName := "test_fixedarray_output0_int_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [0]*int{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([0]*int); !ok {
		t.Errorf("expected a [0]*int, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayInput1_string(t *testing.T) {
	fnName := "test_fixedarray_input1_string"
	arr := [1]string{"abc"}

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayInput1_stringPtr(t *testing.T) {
	fnName := "test_fixedarray_input1_string_option"
	arr := getStringPtrFixedArray1()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayOutput1_string(t *testing.T) {
	fnName := "test_fixedarray_output1_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [1]string{"abc"}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([1]string); !ok {
		t.Errorf("expected a [1]string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput1_stringPtr(t *testing.T) {
	fnName := "test_fixedarray_output1_string_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getStringPtrFixedArray1()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([1]*string); !ok {
		t.Errorf("expected a [1]*string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayInput1_intPtr(t *testing.T) {
	fnName := "test_fixedarray_input1_int_option"
	arr := getIntPtrFixedArray1()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayOutput1_intPtr(t *testing.T) {
	fnName := "test_fixedarray_output1_int_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getIntPtrFixedArray1()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([1]*int); !ok {
		t.Errorf("expected a [1]*int, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayInput2_string(t *testing.T) {
	fnName := "test_fixedarray_input2_string"
	arr := [2]string{"abc", "def"}

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayInput2_stringPtr(t *testing.T) {
	fnName := "test_fixedarray_input2_string_option"
	arr := getStringPtrFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayInput2_struct(t *testing.T) {
	fnName := "test_fixedarray_input2_struct"
	arr := getStructFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayInput2_structPtr(t *testing.T) {
	fnName := "test_fixedarray_input2_struct_option"
	arr := getStructPtrFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayInput2_map(t *testing.T) {
	fnName := "test_fixedarray_input2_map"
	arr := getMapFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayInput2_mapPtr(t *testing.T) {
	fnName := "test_fixedarray_input2_map_option"
	arr := getMapPtrFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayInput2_intPtr(t *testing.T) {
	fnName := "test_fixedarray_input2_int_option"
	arr := getIntPtrFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayOutput2_intPtr(t *testing.T) {
	fnName := "test_fixedarray_output2_int_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getIntPtrFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([2]*int); !ok {
		t.Errorf("expected a [2]*int, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_string(t *testing.T) {
	fnName := "test_fixedarray_output2_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [2]string{"abc", "def"}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([2]string); !ok {
		t.Errorf("expected a [2]string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_stringPtr(t *testing.T) {
	fnName := "test_fixedarray_output2_string_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getStringPtrFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([2]*string); !ok {
		t.Errorf("expected a [2]*string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_struct(t *testing.T) {
	fnName := "test_fixedarray_output2_struct"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getStructFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([2]TestStruct2); !ok {
		t.Errorf("expected a [2]TestStruct2, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_structPtr(t *testing.T) {
	fnName := "test_fixedarray_output2_struct_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getStructPtrFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([2]*TestStruct2); !ok {
		t.Errorf("expected a [2]*TestStruct2, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_map(t *testing.T) {
	fnName := "test_fixedarray_output2_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getMapFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([2]map[string]string); !ok {
		t.Errorf("expected a [2]map[string]string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_mapPtr(t *testing.T) {
	fnName := "test_fixedarray_output2_map_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getMapPtrFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([2]*map[string]string); !ok {
		t.Errorf("expected a [2]*map[string]string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPtrFixedArrayInput1_int(t *testing.T) {
	fnName := "test_option_fixedarray_input1_int"
	arr := getPtrIntFixedArray1()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](*arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, &arr); err != nil {
		t.Error(err)
	}
}

func TestPtrFixedArrayInput2_int(t *testing.T) {
	fnName := "test_option_fixedarray_input2_int"
	arr := getPtrIntFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](*arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, &arr); err != nil {
		t.Error(err)
	}
}

func TestPtrFixedArrayInput1_string(t *testing.T) {
	fnName := "test_option_fixedarray_input1_string"
	arr := getPtrStringFixedArray1()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](*arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, &arr); err != nil {
		t.Error(err)
	}
}

func TestPtrFixedArrayInput2_string(t *testing.T) {
	fnName := "test_option_fixedarray_input2_string"
	arr := getPtrStringFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](*arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, &arr); err != nil {
		t.Error(err)
	}
}

func TestPtrFixedArrayOutput1_int(t *testing.T) {
	fnName := "test_option_fixedarray_output1_int"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getPtrIntFixedArray1()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*[1]int); !ok {
		t.Errorf("expected a *[1]int, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPtrFixedArrayOutput2_int(t *testing.T) {
	fnName := "test_option_fixedarray_output2_int"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getPtrIntFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*[2]int); !ok {
		t.Errorf("expected a *[2]int, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPtrFixedArrayOutput1_string(t *testing.T) {
	fnName := "test_option_fixedarray_output1_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getPtrStringFixedArray1()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*[1]string); !ok {
		t.Errorf("expected a *[1]string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPtrFixedArrayOutput2_string(t *testing.T) {
	fnName := "test_option_fixedarray_output2_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getPtrStringFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*[2]string); !ok {
		t.Errorf("expected a *[2]string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayInput0_byte(t *testing.T) {
	fnName := "test_fixedarray_input0_byte"
	arr := [0]byte{}

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayInput1_byte(t *testing.T) {
	fnName := "test_fixedarray_input1_byte"
	arr := [1]byte{1}

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayInput2_byte(t *testing.T) {
	fnName := "test_fixedarray_input2_byte"
	arr := [2]byte{1, 2}

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayOutput0_byte(t *testing.T) {
	fnName := "test_fixedarray_output0_byte"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [0]byte{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([0]byte); !ok {
		t.Errorf("expected a [0]byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput1_byte(t *testing.T) {
	fnName := "test_fixedarray_output1_byte"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [1]byte{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([1]byte); !ok {
		t.Errorf("expected a [1]byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_byte(t *testing.T) {
	fnName := "test_fixedarray_output2_byte"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := [2]byte{1, 2}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([2]byte); !ok {
		t.Errorf("expected a [2]byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func getIntPtrFixedArray1() [1]*int {
	a := 11
	return [1]*int{&a}
}

func getIntPtrFixedArray2() [2]*int {
	a := 11
	b := 22
	return [2]*int{&a, &b}
}

func getStringPtrFixedArray1() [1]*string {
	a := "abc"
	return [1]*string{&a}
}

func getStringPtrFixedArray2() [2]*string {
	a := "abc"
	b := "def"
	return [2]*string{&a, &b}
}

func getStructFixedArray2() [2]TestStruct2 {
	return [2]TestStruct2{
		{A: true, B: 123},
		{A: false, B: 456},
	}
}

func getStructPtrFixedArray2() [2]*TestStruct2 {
	return [2]*TestStruct2{
		{A: true, B: 123},
		{A: false, B: 456},
	}
}

func getMapFixedArray2() [2]map[string]string {
	return [2]map[string]string{
		{"A": "true", "B": "123"},
		{"C": "false", "D": "456"},
	}
}

func getMapPtrFixedArray2() [2]*map[string]string {
	return [2]*map[string]string{
		{"A": "true", "B": "123"},
		{"C": "false", "D": "456"},
	}
}

func getPtrIntFixedArray1() *[1]int {
	a := 11
	return &[1]int{a}
}

func getPtrIntFixedArray2() *[2]int {
	a := 11
	b := 22
	return &[2]int{a, b}
}

func getPtrStringFixedArray1() *[1]string {
	a := "abc"
	return &[1]string{a}
}

func getPtrStringFixedArray2() *[2]string {
	a := "abc"
	b := "def"
	return &[2]string{a, b}
}
