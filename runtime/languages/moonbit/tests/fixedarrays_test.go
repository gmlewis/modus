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

func TestFixedArrayInput0_string_option(t *testing.T) {
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

	expected := []string{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput0_string_option(t *testing.T) {
	fnName := "test_fixedarray_output0_string_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*string{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayInput0_int_option(t *testing.T) {
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

func TestFixedArrayOutput0_int_option(t *testing.T) {
	fnName := "test_fixedarray_output0_int_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*int{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int); !ok {
		t.Errorf("expected a []*int, got %T", result)
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

func TestFixedArrayInput1_string_option(t *testing.T) {
	fnName := "test_fixedarray_input1_string_option"
	arr := getStringOptionFixedArray1()

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

	expected := []string{"abc"}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput1_string_option(t *testing.T) {
	fnName := "test_fixedarray_output1_string_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getStringOptionFixedArray1()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayInput1_int_option(t *testing.T) {
	fnName := "test_fixedarray_input1_int_option"
	arr := getIntOptionFixedArray1()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayOutput1_int_option(t *testing.T) {
	fnName := "test_fixedarray_output1_int_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getIntOptionFixedArray1()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int); !ok {
		t.Errorf("expected a []*int, got %T", result)
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

func TestFixedArrayInput2_string_option(t *testing.T) {
	fnName := "test_fixedarray_input2_string_option"
	arr := getStringOptionFixedArray2()

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

func TestFixedArrayInput2_struct_option(t *testing.T) {
	fnName := "test_fixedarray_input2_struct_option"
	arr := getStructOptionFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

// func TestFixedArrayInput2_map(t *testing.T) {
// 	fnName := "test_fixedarray_input2_map"
// 	arr := getMapFixedArray2()

// 	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
// 		t.Error(err)
// 	}
// 	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
// 		t.Error("failed conversion to interface slice")
// 	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestFixedArrayInput2_map_option(t *testing.T) {
// 	fnName := "test_fixedarray_input2_map_option"
// 	arr := getMapOptionFixedArray2()

// 	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
// 		t.Error(err)
// 	}
// 	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
// 		t.Error("failed conversion to interface slice")
// 	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
// 		t.Error(err)
// 	}
// }

func TestFixedArrayInput2_int_option(t *testing.T) {
	fnName := "test_fixedarray_input2_int_option"
	arr := getIntOptionFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayOutput2_int_option(t *testing.T) {
	fnName := "test_fixedarray_output2_int_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getIntOptionFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int); !ok {
		t.Errorf("expected a []*int, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

// func TestFixedArrayOutput2_string(t *testing.T) {
// 	fnName := "test_fixedarray_output2_string"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := [2]string{"abc", "def"}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([2]string); !ok {
// 		t.Errorf("expected a [2]string, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestFixedArrayOutput2_string_option(t *testing.T) {
// 	fnName := "test_fixedarray_output2_string_option"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getStringOptionFixedArray2()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([2]*string); !ok {
// 		t.Errorf("expected a [2]*string, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestFixedArrayOutput2_struct(t *testing.T) {
// 	fnName := "test_fixedarray_output2_struct"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getStructFixedArray2()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([2]TestStruct2); !ok {
// 		t.Errorf("expected a [2]TestStruct2, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestFixedArrayOutput2_struct_option(t *testing.T) {
// 	fnName := "test_fixedarray_output2_struct_option"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getStructOptionFixedArray2()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([2]*TestStruct2); !ok {
// 		t.Errorf("expected a [2]*TestStruct2, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestFixedArrayOutput2_map(t *testing.T) {
// 	fnName := "test_fixedarray_output2_map"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getMapFixedArray2()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([2]map[string]string); !ok {
// 		t.Errorf("expected a [2]map[string]string, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestFixedArrayOutput2_map_option(t *testing.T) {
// 	fnName := "test_fixedarray_output2_map_option"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getMapOptionFixedArray2()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([2]*map[string]string); !ok {
// 		t.Errorf("expected a [2]*map[string]string, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestOptionFixedArrayInput1_int(t *testing.T) {
// 	fnName := "test_option_fixedarray_input1_int"
// 	arr := getOptionIntFixedArray1()

// 	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
// 		t.Error(err)
// 	}
// 	if arr, ok := utils.ConvertToSliceOf[any](*arr); !ok {
// 		t.Error("failed conversion to interface slice")
// 	} else if _, err := fixture.CallFunction(t, fnName, &arr); err != nil {
// 		t.Error(err)
// 	}
// }

func TestOptionFixedArrayInput2_int(t *testing.T) {
	fnName := "test_option_fixedarray_input2_int"
	arr := getOptionIntFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](*arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, &arr); err != nil {
		t.Error(err)
	}
}

func TestOptionFixedArrayInput1_string(t *testing.T) {
	fnName := "test_option_fixedarray_input1_string"
	arr := getOptionStringFixedArray1()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](*arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, &arr); err != nil {
		t.Error(err)
	}
}

func TestOptionFixedArrayInput2_string(t *testing.T) {
	fnName := "test_option_fixedarray_input2_string"
	arr := getOptionStringFixedArray2()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](*arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, &arr); err != nil {
		t.Error(err)
	}
}

// func TestOptionFixedArrayOutput1_int(t *testing.T) {
// 	fnName := "test_option_fixedarray_output1_int"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getOptionIntFixedArray1()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.(*[1]int); !ok {
// 		t.Errorf("expected a *[1]int, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestOptionFixedArrayOutput2_int(t *testing.T) {
// 	fnName := "test_option_fixedarray_output2_int"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getOptionIntFixedArray2()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.(*[]int); !ok {
// 		t.Errorf("expected a *[]int, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestOptionFixedArrayOutput1_string(t *testing.T) {
// 	fnName := "test_option_fixedarray_output1_string"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getOptionStringFixedArray1()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.(*[1]string); !ok {
// 		t.Errorf("expected a *[1]string, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestOptionFixedArrayOutput2_string(t *testing.T) {
// 	fnName := "test_option_fixedarray_output2_string"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getOptionStringFixedArray2()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.(*[2]string); !ok {
// 		t.Errorf("expected a *[2]string, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

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

// func TestFixedArrayOutput0_byte(t *testing.T) {
// 	fnName := "test_fixedarray_output0_byte"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := [0]byte{}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([0]byte); !ok {
// 		t.Errorf("expected a [0]byte, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestFixedArrayOutput1_byte(t *testing.T) {
// 	fnName := "test_fixedarray_output1_byte"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := [1]byte{1}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([1]byte); !ok {
// 		t.Errorf("expected a [1]byte, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

// func TestFixedArrayOutput2_byte(t *testing.T) {
// 	fnName := "test_fixedarray_output2_byte"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := [2]byte{1, 2}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([2]byte); !ok {
// 		t.Errorf("expected a [2]byte, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

func getIntOptionFixedArray1() []*int {
	a := 11
	return []*int{&a}
}

func getIntOptionFixedArray2() []*int {
	a := 11
	b := 22
	return []*int{&a, &b}
}

func getStringOptionFixedArray1() []*string {
	a := "abc"
	return []*string{&a}
}

func getStringOptionFixedArray2() []*string {
	a := "abc"
	b := "def"
	return []*string{&a, &b}
}

func getStructFixedArray2() []TestStruct2 {
	return []TestStruct2{
		{A: true, B: 123},
		{A: false, B: 456},
	}
}

func getStructOptionFixedArray2() []*TestStruct2 {
	return []*TestStruct2{
		{A: true, B: 123},
		{A: false, B: 456},
	}
}

// func getMapFixedArray2() []map[string]string {
// 	return []map[string]string{
// 		{"A": "true", "B": "123"},
// 		{"C": "false", "D": "456"},
// 	}
// }

// func getMapOptionFixedArray2() []*map[string]string {
// 	return []*map[string]string{
// 		{"A": "true", "B": "123"},
// 		{"C": "false", "D": "456"},
// 	}
// }

// func getOptionIntFixedArray1() *[]int {
// 	a := 11
// 	return &[]int{a}
// }

func getOptionIntFixedArray2() *[2]int {
	a := 11
	b := 22
	return &[2]int{a, b}
}

func getOptionStringFixedArray1() *[1]string {
	a := "abc"
	return &[1]string{a}
}

func getOptionStringFixedArray2() *[2]string {
	a := "abc"
	b := "def"
	return &[2]string{a, b}
}
