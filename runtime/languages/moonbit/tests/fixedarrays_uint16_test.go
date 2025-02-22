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
	"slices"
	"testing"
)

// func TestFixedArrayInput_uint16_option(t *testing.T) {
// 	fnName := "test_fixedfixedarray_input_uint16_option"
// 	s := getUint16OptionFixedArray()
//
// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestFixedArrayOutput_uint16_option(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_option"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := getUint16OptionFixedArray()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]*uint16); !ok {
// 		t.Errorf("expected a []*uint16, got %T", result)
// 	} else if !slices.EqualFunc(expected, r, func(a, b *uint16) bool { return (a == nil && b == nil) || (a != nil && b != nil && *a == *b) }) {
// 		t.Errorf("expected %#v, got %#v", expected, r)
// 	}
// }

func getUint16OptionFixedArray() []*uint16 {
	a := uint16(11)
	c := uint16(33)
	return []*uint16{&a, nil, &c}
}

// func TestFixedArrayInput_uint16_empty(t *testing.T) {
// 	fnName := "test_fixedfixedarray_input_uint16_empty"
// 	s := []uint16{}

// 	if _, err := fixture.CallFunction(t, fnName, s); err != nil {
// 		t.Error(err)
// 	}
// }

func TestFixedArrayOutput_uint16_0(t *testing.T) {
	fnName := "test_fixedfixedarray_output_uint16_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []uint16{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]uint16); !ok {
		t.Errorf("expected a []uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

// func TestFixedArrayOutput_uint16_1(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_1"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []uint16{1}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]uint16); !ok {
// 		t.Errorf("expected a []uint16, got %T", result)
// 	} else if !slices.Equal(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }
//
// func TestFixedArrayOutput_uint16_1_min(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_1_min"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []uint16{0}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]uint16); !ok {
// 		t.Errorf("expected a []uint16, got %T", result)
// 	} else if !slices.Equal(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }
//
// func TestFixedArrayOutput_uint16_1_max(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_1_max"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []uint16{math.MaxUint16}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]uint16); !ok {
// 		t.Errorf("expected a []uint16, got %T", result)
// 	} else if !slices.Equal(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }
//
// func TestFixedArrayOutput_uint16_2(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_2"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []uint16{0, math.MaxUint16}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]uint16); !ok {
// 		t.Errorf("expected a []uint16, got %T", result)
// 	} else if !slices.Equal(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }
//
// func TestFixedArrayOutput_uint16_3(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_3"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []uint16{1, 2, 3}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]uint16); !ok {
// 		t.Errorf("expected a []uint16, got %T", result)
// 	} else if !slices.Equal(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }
//
// func TestFixedArrayOutput_uint16_4(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_4"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []uint16{1, 2, 3, 4}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]uint16); !ok {
// 		t.Errorf("expected a []uint16, got %T", result)
// 	} else if !slices.Equal(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

func TestFixedArrayOutput_uint16_option_0(t *testing.T) {
	fnName := "test_fixedfixedarray_output_uint16_option_0"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint16{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint16); !ok {
		t.Errorf("expected a []*uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput_uint16_option_1_none(t *testing.T) {
	fnName := "test_fixedfixedarray_output_uint16_option_1_none"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*uint16{nil}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*uint16); !ok {
		t.Errorf("expected a []*uint16, got %T", result)
	} else if !slices.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

// func TestFixedArrayOutput_uint16_option_1_min(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_option_1_min"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []*uint16{Ptr(uint16(0))}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]*uint16); !ok {
// 		t.Errorf("expected a []*uint16, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }
//
// func TestFixedArrayOutput_uint16_option_1_max(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_option_1_max"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []*uint16{Ptr(uint16(math.MaxUint16))}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]*uint16); !ok {
// 		t.Errorf("expected a []*uint16, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }
//
// func TestFixedArrayOutput_uint16_option_2(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_option_2"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []*uint16{Ptr(uint16(1)), Ptr(uint16(2))}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]*uint16); !ok {
// 		t.Errorf("expected a []*uint16, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }
//
// func TestFixedArrayOutput_uint16_option_3(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_option_3"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []*uint16{nil, nil, nil}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]*uint16); !ok {
// 		t.Errorf("expected a []*uint16, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }
//
// func TestFixedArrayOutput_uint16_option_4(t *testing.T) {
// 	fnName := "test_fixedfixedarray_output_uint16_option_4"
//
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	expected := []*uint16{nil, Ptr(uint16(0)), Ptr(uint16(0)), Ptr(uint16(math.MaxUint16))}
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]*uint16); !ok {
// 		t.Errorf("expected a []*uint16, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }
