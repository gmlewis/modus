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

// func TestFixedArrayInput_string(t *testing.T) {
// 	fnName := "test_fixedarray_input_string"
// 	s := []string{"abc", "def", "ghi"}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestFixedArrayInput_string_option(t *testing.T) {
// 	fnName := "test_fixedarray_input_string_option"
// 	s := getStringOptionFixedArray()

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func TestFixedArrayOutput_string(t *testing.T) {
	fnName := "test_fixedarray_output_string"

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

func TestFixedArrayOutput_string_option(t *testing.T) {
	fnName := "test_fixedarray_output_string_option"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getStringOptionFixedArray()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !slices.EqualFunc(expected, r, func(a, b *string) bool { return (a == nil && b == nil) || (a != nil && b != nil && *a == *b) }) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func getStringOptionFixedArray() []*string {
	a := "abc"
	// b := "def"
	c := "ghi"
	return []*string{&a, nil, &c}
}

func TestFixedArrayInput_string_none(t *testing.T) {
	fnName := "test_fixedarray_input_string_none"

	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestFixedArrayOutput_string_none(t *testing.T) {
	fnName := "test_fixedarray_output_string_none"

	// pointerHandler.Decode(vals: [0])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

// func TestFixedArrayInput_string_empty(t *testing.T) {
// 	fnName := "test_fixedarray_input_string_empty"
// 	s := []string{}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func TestFixedArrayOutput_string_empty(t *testing.T) {
	fnName := "test_fixedarray_output_string_empty"

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

// func Test2DFixedArrayInput_string(t *testing.T) {
// 	fnName := "test2d_fixedarray_input_string"
// 	s := [][]string{
// 		{"abc", "def", "ghi"},
// 		{"jkl", "mno", "pqr"},
// 		{"stu", "vwx", "yz"},
// 	}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func Test2DFixedArrayOutput_string(t *testing.T) {
	fnName := "test2d_fixedarray_output_string"

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

func Test2DFixedArrayInput_string_none(t *testing.T) {
	fnName := "test2d_fixedarray_input_string_none"

	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func Test2DFixedArrayOutput_string_none(t *testing.T) {
	fnName := "test2d_fixedarray_output_string_none"

	// pointerHandler.Decode(vals: [0])
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

// func Test2DFixedArrayInput_string_empty(t *testing.T) {
// 	fnName := "test2d_fixedarray_input_string_empty"
// 	s := [][]string{}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func Test2DFixedArrayOutput_string_empty(t *testing.T) {
	fnName := "test2d_fixedarray_output_string_empty"

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

// func Test2DFixedArrayInput_string_inner_empty(t *testing.T) {
// 	fnName := "test2d_fixedarray_input_string_inner_empty"
// 	s := [][]string{{}}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func Test2DFixedArrayOutput_string_inner_empty(t *testing.T) {
	fnName := "test2d_fixedarray_output_string_inner_empty"

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

// func Test2DFixedArrayInput_string_inner_none(t *testing.T) {
// 	fnName := "test2d_fixedarray_input_string_inner_none"
// 	s := []*[]string{nil}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func Test2DFixedArrayOutput_string_inner_none(t *testing.T) {
	fnName := "test2d_fixedarray_output_string_inner_none"

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
