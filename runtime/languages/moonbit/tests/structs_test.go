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

type TestStruct1 struct {
	A bool
}

type TestStruct2 struct {
	A bool
	B int
}

type TestStruct3 struct {
	A bool
	B int
	C string
}

type TestStruct4 struct {
	A bool
	B int
	C *string
}

type TestStruct5 struct {
	A string
	B string
	C string
	D []string
	E float64
	F float64
}

var testStruct1 = TestStruct1{
	A: true,
}

var testStruct2 = TestStruct2{
	A: true,
	B: 123,
}

var testStruct3 = TestStruct3{
	A: true,
	B: 123,
	C: "abc",
}

var testStruct4 = TestStruct4{
	A: true,
	B: 123,
	C: func() *string { s := "abc"; return &s }(),
}

var testStruct4_withNil = TestStruct4{
	A: true,
	B: 123,
	C: nil,
}

var testStruct5 = TestStruct5{
	A: "abc",
	B: "def",
	C: "ghi",
	D: []string{
		"jkl",
		"mno",
		"pqr",
	},
	E: 0.12345,
	F: 99.99999,
}

var testStruct1AsMap = map[string]any{
	"a": true,
}

var testStruct2AsMap = map[string]any{
	"a": true,
	"b": 123,
}

var testStruct3AsMap = map[string]any{
	"a": true,
	"b": 123,
	"c": "abc",
}

var testStruct4AsMap = map[string]any{
	"a": true,
	"b": 123,
	"c": func() *string { s := "abc"; return &s }(),
}

var testStruct4AsMap_withNil = map[string]any{
	"a": true,
	"b": 123,
	"c": nil,
}

var testStruct5AsMap = map[string]any{
	"a": "abc",
	"b": "def",
	"c": "ghi",
	"d": []string{
		"jkl",
		"mno",
		"pqr",
	},
	"e": 0.12345,
	"f": 99.99999,
}

func TestStructInput1(t *testing.T) {
	fnName := "test_struct_input1"
	if _, err := fixture.CallFunction(t, fnName, testStruct1); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct1AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructInput2(t *testing.T) {
	fnName := "test_struct_input2"
	if _, err := fixture.CallFunction(t, fnName, testStruct2); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct2AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructInput3(t *testing.T) {
	fnName := "test_struct_input3"
	if _, err := fixture.CallFunction(t, fnName, testStruct3); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct3AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructInput4(t *testing.T) {
	fnName := "test_struct_input4"
	if _, err := fixture.CallFunction(t, fnName, testStruct4); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct4AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructInput5(t *testing.T) {
	fnName := "test_struct_input5"
	if _, err := fixture.CallFunction(t, fnName, testStruct5); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct5AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructInput4_withNil(t *testing.T) {
	fnName := "test_struct_input4_with_none"
	if _, err := fixture.CallFunction(t, fnName, testStruct4_withNil); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct4AsMap_withNil); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput1(t *testing.T) {
	fnName := "test_struct_option_input1"
	if _, err := fixture.CallFunction(t, fnName, testStruct1); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct1); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct1AsMap); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct1AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput2(t *testing.T) {
	fnName := "test_struct_option_input2"
	if _, err := fixture.CallFunction(t, fnName, testStruct2); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct2); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct2AsMap); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct2AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput3(t *testing.T) {
	fnName := "test_struct_option_input3"
	if _, err := fixture.CallFunction(t, fnName, testStruct3); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct3); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct3AsMap); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct3AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput4(t *testing.T) {
	fnName := "test_struct_option_input4"
	if _, err := fixture.CallFunction(t, fnName, testStruct4); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct4); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct4AsMap); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct4AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput5(t *testing.T) {
	fnName := "test_struct_option_input5"
	if _, err := fixture.CallFunction(t, fnName, testStruct5); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct5); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct5AsMap); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct5AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput4_withNil(t *testing.T) {
	fnName := "test_struct_option_input4_with_none"
	if _, err := fixture.CallFunction(t, fnName, testStruct4_withNil); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct4_withNil); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct4AsMap_withNil); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct4AsMap_withNil); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput1_none(t *testing.T) {
	fnName := "test_struct_option_input1_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput2_none(t *testing.T) {
	fnName := "test_struct_option_input2_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput3_none(t *testing.T) {
	fnName := "test_struct_option_input3_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput4_none(t *testing.T) {
	fnName := "test_struct_option_input4_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestStructOptionInput5_none(t *testing.T) {
	fnName := "test_struct_option_input5_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestStructOutput1(t *testing.T) {
	fnName := "test_struct_output1"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct1

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(TestStruct1); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput2(t *testing.T) {
	fnName := "test_struct_output2"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct2

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(TestStruct2); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput3(t *testing.T) {
	fnName := "test_struct_output3"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct3

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(TestStruct3); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput4(t *testing.T) {
	fnName := "test_struct_output4"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct4

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(TestStruct4); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput5(t *testing.T) {
	fnName := "test_struct_output5"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct5

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(TestStruct5); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput4_withNil(t *testing.T) {
	fnName := "test_struct_output4_with_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct4_withNil

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(TestStruct4); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput1(t *testing.T) {
	fnName := "test_struct_option_output1"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct1

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*TestStruct1); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput2(t *testing.T) {
	fnName := "test_struct_option_output2"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct2

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*TestStruct2); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput3(t *testing.T) {
	fnName := "test_struct_option_output3"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct3

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*TestStruct3); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput4(t *testing.T) {
	fnName := "test_struct_option_output4"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct4

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*TestStruct4); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput5(t *testing.T) {
	fnName := "test_struct_option_output5"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct5

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*TestStruct5); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput4_withNil(t *testing.T) {
	fnName := "test_struct_option_output4_with_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct4_withNil

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*TestStruct4); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput1_map(t *testing.T) {
	fnName := "test_struct_output1_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct1AsMap

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput2_map(t *testing.T) {
	fnName := "test_struct_output2_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct2AsMap

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput3_map(t *testing.T) {
	fnName := "test_struct_output3_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct3AsMap

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput4_map(t *testing.T) {
	fnName := "test_struct_output4_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct4AsMap

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput5_map(t *testing.T) {
	fnName := "test_struct_output5_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct5AsMap

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOutput4_map_withNil(t *testing.T) {
	fnName := "test_struct_output4_map_with_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct4AsMap_withNil

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput1_map(t *testing.T) {
	fnName := "test_struct_option_output1_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct1AsMap

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput2_map(t *testing.T) {
	fnName := "test_struct_option_output2_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct2AsMap

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput3_map(t *testing.T) {
	fnName := "test_struct_option_output3_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct3AsMap

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput4_map(t *testing.T) {
	fnName := "test_struct_option_output4_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct4AsMap

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput5_map(t *testing.T) {
	fnName := "test_struct_option_output5_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct5AsMap

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput4_map_withNil(t *testing.T) {
	fnName := "test_struct_option_output4_map_with_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct4AsMap_withNil

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*map[string]any); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestStructOptionOutput1_none(t *testing.T) {
	fnName := "test_struct_option_output1_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestStructOptionOutput2_none(t *testing.T) {
	fnName := "test_struct_option_output2_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestStructOptionOutput3_none(t *testing.T) {
	fnName := "test_struct_option_output3_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestStructOptionOutput4_none(t *testing.T) {
	fnName := "test_struct_option_output4_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestStructOptionOutput5_none(t *testing.T) {
	fnName := "test_struct_option_output5_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
