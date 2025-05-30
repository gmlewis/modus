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
	"fmt"
	"maps"
	"reflect"
	"testing"

	"github.com/gmlewis/modus/runtime/utils"
	"github.com/google/go-cmp/cmp"
)

func TestMapInput_string_string(t *testing.T) {
	fnName := "test_map_input_string_string"
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	if _, err := fixture.CallFunction(t, fnName, m); err != nil {
		t.Error(err)
	}
	if m, err := utils.ConvertToMap(m); err != nil {
		t.Error(fmt.Errorf("failed conversion to interface map: %w", err))
	} else if _, err := fixture.CallFunction(t, fnName, m); err != nil {
		t.Error(err)
	}
}

func TestMapOptionInput_string_string(t *testing.T) {
	fnName := "test_map_option_input_string_string"
	m := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	if _, err := fixture.CallFunction(t, fnName, m); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &m); err != nil {
		t.Error(err)
	}
	if m, err := utils.ConvertToMap(m); err != nil {
		t.Error(fmt.Errorf("failed conversion to interface map: %w", err))
	} else if _, err := fixture.CallFunction(t, fnName, m); err != nil {
		t.Error(err)
	} else if _, err := fixture.CallFunction(t, fnName, &m); err != nil {
		t.Error(err)
	}
}

func TestMapOutput_string_string(t *testing.T) {
	fnName := "test_map_output_string_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(map[string]string); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !maps.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestMapOptionOutput_string_string(t *testing.T) {
	fnName := "test_map_option_output_string_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*map[string]string); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !maps.Equal(expected, *r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestIterateMap_string_string(t *testing.T) {
	fnName := "test_iterate_map_string_string"
	m := makeTestMap(100)

	if _, err := fixture.CallFunction(t, fnName, m); err != nil {
		t.Error(err)
	}
}

func TestGenerateMap_string_string_output(t *testing.T) {
	fnName := "test_generate_map_string_string_output"
	want := makeTestMap(100)
	got, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("expected %T, got %T", want, got)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("%v: mismatch (-want +got):\n%v", fnName, diff)
	}
}

func TestMapLookup_string_string(t *testing.T) {
	fnName := "test_map_lookup_string_string"
	m := makeTestMap(100)

	result, err := fixture.CallFunction(t, fnName, m, "key_047")
	if err != nil {
		t.Fatal(err)
	}

	expected := "val_047"

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(string); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if expected != r {
		t.Errorf("expected %s, got %s", expected, r)
	}
}

type TestStructWithMap1 struct {
	M map[string]string
}

type TestStructWithMap2 struct {
	M map[string]any
}

func TestStructContainingMapInput_string_string(t *testing.T) {
	fnName := "test_struct_containing_map_input_string_string"
	s1 := TestStructWithMap1{M: map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}}
	if _, err := fixture.CallFunction(t, fnName, s1); err != nil {
		t.Error(err)
	}

	s2 := TestStructWithMap2{M: map[string]any{
		"a": any("1"),
		"b": any("2"),
		"c": any("3"),
	}}

	if _, err := fixture.CallFunction(t, fnName, s2); err != nil {
		t.Error(err)
	}
}

func TestStructContainingMapOutput_string_string(t *testing.T) {
	fnName := "test_struct_containing_map_output_string_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := TestStructWithMap1{M: map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(TestStructWithMap1); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func makeTestMap(size int) map[string]string {
	m := make(map[string]string, size)
	for i := 0; i < size; i++ {
		key := fmt.Sprintf("key_%03d", i)
		val := fmt.Sprintf("val_%03d", i)
		m[key] = val
	}
	return m
}

func TestMapInput_int_float(t *testing.T) {
	fnName := "test_map_input_int_float"
	m := map[int]float32{
		1: 1.1,
		2: 2.2,
		3: 3.3,
	}

	if _, err := fixture.CallFunction(t, fnName, m); err != nil {
		t.Error(err)
	}
	if m, err := utils.ConvertToMap(m); err != nil {
		t.Error(fmt.Errorf("failed conversion to interface map: %w", err))
	} else if _, err := fixture.CallFunction(t, fnName, m); err != nil {
		t.Error(err)
	}
}

func TestMapOutput_int_float(t *testing.T) {
	fnName := "test_map_output_int_float"

	// memoryBlockAtOffset(offset: 94304=0x00017060=[96 112 1 0], size: 40=8+words*4), classID=0(Tuple), words=8, memBlock=[1 0 0 0 0 8 0 0 224 111 1 0 48 112 1 0 3 0 0 0 8 0 0 0 7 0 0 0 6 0 0 0 176 112 1 0 80 113 1 0]
	// GML: handler_maps.go: mapHandler.Decode: sliceOffset=[224 111 1 0]=94176
	// GML: handler_maps.go: mapHandler.Decode: numElements=[3 0 0 0]=3
	// memoryBlockAtOffset(offset: 94176=0x00016FE0=[224 111 1 0], size: 40=8+words*4), classID=242(FixedArray[String]), words=8, memBlock=[1 0 0 0 242 8 0 0 0 0 0 0 0 113 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 176 112 1 0 80 113 1 0]
	// GML: handler_maps.go: mapHandler.Decode: (sliceOffset: 94176, numElements: 3), sliceMemBlock([224 111 1 0]=@94176)=(40 bytes)=[1 0 0 0 242 8 0 0 0 0 0 0 0 113 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 176 112 1 0 80 113 1 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[int32]float32{
		1: 1.1,
		2: 2.2,
		3: 3.3,
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(map[int32]float32); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !maps.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestMapInput_int_double(t *testing.T) {
	fnName := "test_map_input_int_double"
	m := map[int]float64{
		1: 1.1,
		2: 2.2,
		3: 3.3,
	}

	if _, err := fixture.CallFunction(t, fnName, m); err != nil {
		t.Error(err)
	}
	if m, err := utils.ConvertToMap(m); err != nil {
		t.Error(fmt.Errorf("failed conversion to interface map: %w", err))
	} else if _, err := fixture.CallFunction(t, fnName, m); err != nil {
		t.Error(err)
	}
}

func TestMapOutput_int_double(t *testing.T) {
	fnName := "test_map_output_int_double"

	// memoryBlockAtOffset(offset: 49312=0x0000C0A0=[160 192 0 0], size: 40=8+words*4), moonBitType=0(Tuple), words=8, memBlock=[1 0 0 0 0 8 0 0 32 192 0 0 112 192 0 0 3 0 0 0 8 0 0 0 7 0 0 0 6 0 0 0 240 192 0 0 144 193 0 0]
	// memoryBlockAtOffset(offset: 49184=0x0000C020=[32 192 0 0], size: 40=8+words*4), moonBitType=242(), words=8, memBlock=[1 0 0 0 242 8 0 0 0 0 0 0 64 193 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 240 192 0 0 144 193 0 0]
	// memoryBlockAtOffset(offset: 49472=0x0000C140=[64 193 0 0], size: 32=8+words*4), moonBitType=0(Tuple), words=6, memBlock=[3 0 0 0 0 6 0 0 1 0 0 0 0 0 0 0 233 232 223 241 2 0 0 0 154 153 153 153 153 153 1 64]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[int32]float64{
		1: 1.1,
		2: 2.2,
		3: 3.3,
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(map[int32]float64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !maps.Equal(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}
