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

func TestArrayInput_string_option(t *testing.T) {
	fnName := "test_array_input_string_option"
	s := getStringOptionArray()

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

// func TestArrayOutput_byte(t *testing.T) {
// 	fnName := "test_array_output_byte"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := []byte{0x01, 0x02, 0x03, 0x04}

// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]byte); !ok {
// 		t.Errorf("expected a []byte, got %T", result)
// 	} else if !bytes.Equal(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

func TestArrayOutput_int_option(t *testing.T) {
	fnName := "test_array_output_int_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getIntOptionArray()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int); !ok {
		t.Errorf("expected a []*int, got %T", result)
	} else if !slices.EqualFunc(expected, r, func(a, b *int) bool { return (a == nil && b == nil) || (a != nil && b != nil && *a == *b) }) {
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

func getIntOptionArray() []*int {
	a := 11
	// b := 22
	c := 33
	return []*int{&a, nil, &c}
}

func getStringOptionArray() []*string {
	a := "abc"
	// b := "def"
	c := "ghi"
	return []*string{&a, nil, &c}
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

// func TestArrayInput_int_empty(t *testing.T) {
// 	fnName := "test_array_input_int_empty"
// 	s := []int{}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestArrayOutput_int_empty(t *testing.T) {
// 	fnName := "test_array_output_int_empty"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := []int{}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]int); !ok {
// 		t.Errorf("expected a []int, got %T", result)
// 	} else if !slices.Equal(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

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

func Test2DArrayInput_string_inner_empty(t *testing.T) {
	fnName := "test2d_array_input_string_inner_empty"
	s := [][]string{{}}

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func Test2DArrayOutput_string_inner_empty(t *testing.T) {
	fnName := "test2d_array_output_string_inner_empty"
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
	s := []*[]string{nil} // [][]string{nil}

	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
		t.Error(err)
	}
}

func Test2DArrayOutput_string_inner_none(t *testing.T) {
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
