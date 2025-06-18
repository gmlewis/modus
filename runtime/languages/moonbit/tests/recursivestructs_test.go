/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// Tests pass with moonc v0.6.18+8382ed77e

package moonbit_test

import (
	"testing"

	"github.com/gmlewis/modus/runtime/utils"
)

type TestRecursiveStruct struct {
	A bool
	B *TestRecursiveStruct
}

var testRecursiveStruct = func() *TestRecursiveStruct {
	r1 := &TestRecursiveStruct{
		A: true,
	}
	r2 := &TestRecursiveStruct{
		A: false,
	}
	r1.B = r2
	r2.B = r1
	return r1
}()

// Note that testRecursiveStruct and testRecursiveStructAsMap must
// both represents two nodes that point to each other for the MoonBit
// tests to pass. The first struct must have A=true and the second must
// have A=false.
var testRecursiveStructAsMap = func() map[string]any {
	r1 := map[string]any{
		"a": true,
	}
	r2 := map[string]any{
		"a": false,
	}
	r1["b"] = r2
	r2["b"] = r1
	return r1
}()

func TestRecursiveStructInput(t *testing.T) {
	fnName := "test_recursive_struct_input"
	if _, err := fixture.CallFunction(t, fnName, testRecursiveStruct); err != nil {
		t.Error(err)
	}
}

func TestRecursiveStructOptionInput(t *testing.T) {
	fnName := "test_recursive_struct_option_input"
	if _, err := fixture.CallFunction(t, fnName, testRecursiveStruct); err != nil {
		t.Error(err)
	}
}

func TestRecursiveStructInput_map(t *testing.T) {
	fnName := "test_recursive_struct_input"

	// log.Printf("GML: TestRecursiveStructInput_map: testRecursiveStructAsMap    =%v=0x%[1]x", reflect.ValueOf(testRecursiveStructAsMap).Pointer())
	// log.Printf("GML: TestRecursiveStructInput_map: testRecursiveStructAsMap.b  =%v=0x%[1]x", reflect.ValueOf(testRecursiveStructAsMap["b"]).Pointer())
	// log.Printf("GML: TestRecursiveStructInput_map: testRecursiveStructAsMap.b.b=%v=0x%[1]x", reflect.ValueOf(testRecursiveStructAsMap["b"].(map[string]any)["b"]).Pointer())

	if _, err := fixture.CallFunction(t, fnName, testRecursiveStructAsMap); err != nil {
		t.Error(err)
	}
}

func TestRecursiveStructOptionInput_map(t *testing.T) {
	fnName := "test_recursive_struct_input"
	if _, err := fixture.CallFunction(t, fnName, testRecursiveStructAsMap); err != nil {
		t.Fatal(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testRecursiveStructAsMap); err != nil {
		t.Fatal(err)
	}
}

func TestRecursiveStructOptionInput_none(t *testing.T) {
	fnName := "test_recursive_struct_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestRecursiveStructOutput(t *testing.T) {
	fnName := "test_recursive_struct_output"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("expected a result")
	}
	r1, ok := result.(TestRecursiveStruct)
	if !ok {
		t.Fatalf("expected a struct, got %T", result)
	}
	r2 := r1.B
	if r2 == nil {
		t.Fatalf("expected a struct, got %T", r2)
	}
	r3 := r2.B
	if r3 == nil {
		t.Fatalf("expected a struct, got %T", r2)
	}
	if !r1.A {
		t.Errorf("expected r1.A=true, got %v", r1.A)
	}
	if r2.A {
		t.Errorf("expected r2.A=false, got %v", r2.A)
	}
	if !r3.A {
		t.Errorf("expected r3.A=true, got %v", r3.A)
	}
}

func TestRecursiveStructOptionOutput(t *testing.T) {
	fnName := "test_recursive_struct_option_output"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Fatal("expected a result")
	}
	r1, ok := result.(*TestRecursiveStruct)
	if !ok {
		t.Fatalf("expected a pointer to a struct, got %T", result)
	}
	r2 := r1.B
	if r2 == nil {
		t.Fatalf("expected a struct, got %T", r2)
	}
	r3 := r2.B
	if r3 == nil {
		t.Fatalf("expected a struct, got %T", r2)
	}
	if !r1.A {
		t.Errorf("expected r1.A=true, got %v", r1.A)
	}
	if r2.A {
		t.Errorf("expected r2.A=false, got %v", r2.A)
	}
	if !r3.A {
		t.Errorf("expected r3.A=true, got %v", r3.A)
	}
}

func TestRecursiveStructOutput_map(t *testing.T) {
	fnName := "test_recursive_struct_output_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Fatal("expected a result")
	}
	r1, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected a map[string]any, got %T", result)
		// reflect.DeepEqual does not work here with two self-referencing maps.
		// } else if !reflect.DeepEqual(testRecursiveStructAsMap, r) {
	}
	if r1["a"] != true {
		t.Errorf("expected r1.a=true, got %v", r1["a"])
	}
	r2p, ok := r1["b"].(*map[string]any)
	if !ok {
		t.Fatalf("expected a map[string]any, got %T", r1["b"])
	}
	r2 := *r2p
	if r2["a"] != false {
		t.Errorf("expected r2.a=true, got %v", r2["a"])
	}
	r3, ok := r2["b"].(map[string]any)
	if !ok {
		t.Fatalf("expected a map[string]any, got %T", r2["b"])
	}
	if r3["a"] != true {
		t.Errorf("expected r3.a=true, got %v", r3["a"])
	}
}

func TestRecursiveStructOptionOutput_map(t *testing.T) {
	fnName := "test_recursive_struct_option_output_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Fatal("expected a result")
	}
	r1p, ok := result.(*map[string]any)
	if !ok {
		t.Fatalf("expected a map[string]any, got %T", result)
	}
	r1 := *r1p
	if r1["a"] != true {
		t.Errorf("expected r1.a=true, got %v", r1["a"])
	}
	r2p, ok := r1["b"].(*map[string]any)
	if !ok {
		t.Fatalf("expected a map[string]any, got %T", r1["b"])
	}
	r2 := *r2p
	if r2["a"] != false {
		t.Errorf("expected r2.a=true, got %v", r2["a"])
	}
	r3, ok := r2["b"].(map[string]any)
	if !ok {
		t.Fatalf("expected a map[string]any, got %T", r2["b"])
	}
	if r3["a"] != true {
		t.Errorf("expected r3.a=true, got %v", r3["a"])
	}
}

func TestRecursiveStructOptionOutput_none(t *testing.T) {
	fnName := "test_recursive_struct_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
