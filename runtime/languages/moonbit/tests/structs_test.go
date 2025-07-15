// -*- compile-command: "NO_COLOR=1 go test -timeout 30s -tags integration -run '^(TestStruct|TestSmorg)' ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// Tests FAIL with moonc v0.6.20

package moonbit_test

import (
	"testing"

	"github.com/gmlewis/modus/runtime/utils"
	"github.com/google/go-cmp/cmp"
)

type TestStruct1 struct {
	A bool
}

type TestStruct2 struct {
	A bool
	B int32
}

type TestStruct3 struct {
	A bool
	B int32
	C string
}

type TestStruct4 struct {
	A bool
	B int32
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

type TestSmorgasbordStruct struct {
	Bool       bool
	Byte       byte
	C          int16
	F          float32
	D          float64
	I16        int16
	I32        int32
	I64        int64
	S          string
	U16        uint16
	U32        uint32
	U64        uint64
	SomeBool   *bool
	NoneBool   *bool
	SomeByte   *byte
	NoneByte   *byte
	SomeChar   *int16
	NoneChar   *int16
	SomeFloat  *float32
	NoneFloat  *float32
	SomeDouble *float64
	NoneDouble *float64
	SomeI16    *int16
	NoneI16    *int16
	SomeI32    *int32
	NoneI32    *int32
	SomeI64    *int64
	NoneI64    *int64
	SomeString *string
	NoneString *string
	SomeU16    *uint16
	NoneU16    *uint16
	SomeU32    *uint32
	NoneU32    *uint32
	SomeU64    *uint64
	NoneU64    *uint64
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

var testStruct4_with_none = TestStruct4{
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

var testSmorgasbordStruct = TestSmorgasbordStruct{
	Bool:       true,
	Byte:       0x12,
	C:          'c',
	F:          1.23,
	D:          4.56,
	I16:        123,
	I32:        456,
	I64:        789,
	S:          "abc",
	U16:        123,
	U32:        456,
	U64:        789,
	SomeBool:   Ptr(true),
	NoneBool:   nil,
	SomeByte:   Ptr(byte(0x34)),
	NoneByte:   nil,
	SomeChar:   Ptr(int16('d')),
	NoneChar:   nil,
	SomeFloat:  Ptr(float32(7.89)),
	NoneFloat:  nil,
	SomeDouble: Ptr(float64(0.12)),
	NoneDouble: nil,
	SomeI16:    Ptr(int16(234)),
	NoneI16:    nil,
	SomeI32:    Ptr(int32(567)),
	NoneI32:    nil,
	SomeI64:    Ptr(int64(890)),
	NoneI64:    nil,
	SomeString: Ptr("def"),
	NoneString: nil,
	SomeU16:    Ptr(uint16(234)),
	NoneU16:    nil,
	SomeU32:    Ptr(uint32(567)),
	NoneU32:    nil,
	SomeU64:    Ptr(uint64(890)),
	NoneU64:    nil,
}

var testStruct1AsMap = map[string]any{
	"a": true,
}

var testStruct2AsMap = map[string]any{
	"a": true,
	"b": int32(123),
}

var testStruct3AsMap = map[string]any{
	"a": true,
	"b": int32(123),
	"c": "abc",
}

var testStruct4AsMap = map[string]any{
	"a": true,
	"b": int32(123),
	"c": func() *string { s := "abc"; return &s }(),
}

var testStruct4AsMap_with_none = map[string]any{
	"a": true,
	"b": int32(123),
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

var testSmorgasbordStructAsMap = map[string]any{
	"bool":       true,
	"byte":       byte(0x12),
	"c":          int16('c'),
	"f":          float32(1.23),
	"d":          float64(4.56),
	"i16":        int16(123),
	"i32":        int32(456),
	"i64":        int64(789),
	"s":          "abc",
	"u16":        uint16(123),
	"u32":        uint32(456),
	"u64":        uint64(789),
	"someBool":   Ptr(true),
	"noneBool":   nil,
	"someByte":   Ptr(byte(0x34)),
	"noneByte":   nil,
	"someChar":   Ptr(int16('d')),
	"noneChar":   nil,
	"someFloat":  Ptr(float32(7.89)),
	"noneFloat":  nil,
	"someDouble": Ptr(float64(0.12)),
	"noneDouble": nil,
	"someI16":    Ptr(int16(234)),
	"noneI16":    nil,
	"someI32":    Ptr(int32(567)),
	"noneI32":    nil,
	"someI64":    Ptr(int64(890)),
	"noneI64":    nil,
	"someString": Ptr("def"),
	"noneString": nil,
	"someU16":    Ptr(uint16(234)),
	"noneU16":    nil,
	"someU32":    Ptr(uint32(567)),
	"noneU32":    nil,
	"someU64":    Ptr(uint64(890)),
	"noneU64":    nil,
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
	t.Skip("TODO: fix this")
	fnName := "test_struct_input5"
	if _, err := fixture.CallFunction(t, fnName, testStruct5); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct5AsMap); err != nil {
		t.Error(err)
	}
}

func TestStructInput4_with_none(t *testing.T) {
	fnName := "test_struct_input4_with_none"
	if _, err := fixture.CallFunction(t, fnName, testStruct4_with_none); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct4AsMap_with_none); err != nil {
		t.Error(err)
	}
}

func TestSmorgasbordStructInput(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_smorgasbord_struct_input"
	if _, err := fixture.CallFunction(t, fnName, testSmorgasbordStruct); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testSmorgasbordStructAsMap); err != nil {
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
		t.Error("test_struct_option_input4(testStruct4): %w", err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct4); err != nil {
		t.Error("test_struct_option_input4(&testStruct4): %w", err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct4AsMap); err != nil {
		t.Error("test_struct_option_input4(testStruct4AsMap): %w", err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct4AsMap); err != nil {
		t.Error("test_struct_option_input4(&testStruct4AsMap): %w", err)
	}
}

func TestStructOptionInput5(t *testing.T) {
	t.Skip("TODO: fix this")
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

func TestStructOptionInput4_with_none(t *testing.T) {
	fnName := "test_struct_option_input4_with_none"
	if _, err := fixture.CallFunction(t, fnName, testStruct4_with_none); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct4_with_none); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, testStruct4AsMap_with_none); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &testStruct4AsMap_with_none); err != nil {
		t.Error(err)
	}
}

func TestSmorgasbordStructOptionInput(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_smorgasbord_struct_option_input"

	if _, err := fixture.CallFunction(t, fnName, testSmorgasbordStruct); err != nil {
		t.Error(err)
	}
	// if _, err := fixture.CallFunction(t, fnName, &testSmorgasbordStruct); err != nil {
	// 	t.Error(err)
	// }
	// if _, err := fixture.CallFunction(t, fnName, testSmorgasbordStructAsMap); err != nil {
	// 	t.Error(err)
	// }
	// if _, err := fixture.CallFunction(t, fnName, &testSmorgasbordStructAsMap); err != nil {
	// 	t.Error(err)
	// }
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

func TestSmorgasbordStructOptionInput_none(t *testing.T) {
	fnName := "test_smorgasbord_struct_option_input_none"
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
		t.Fatal("expected a result")
	}
	r, ok := result.(TestStruct1)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(TestStruct2)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(TestStruct3)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(TestStruct4)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestStructOutput5(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_struct_output5"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct5

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(TestStruct5)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestStructOutput4_with_none(t *testing.T) {
	fnName := "test_struct_output4_with_none"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct4_with_none

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(TestStruct4)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestSmorgasbordStructOutput(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_smorgasbord_struct_output"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testSmorgasbordStruct

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(TestSmorgasbordStruct)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(*TestStruct1)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(*TestStruct2)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(*TestStruct3)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(*TestStruct4)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestStructOptionOutput5(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_struct_option_output5"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct5

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(*TestStruct5)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestStructOptionOutput4_with_none(t *testing.T) {
	fnName := "test_struct_option_output4_with_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct4_with_none

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(*TestStruct4)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestSmorgasbordStructOptionOutput(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_smorgasbord_struct_option_output"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testSmorgasbordStruct

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(*TestSmorgasbordStruct)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestStructOutput5_map(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_struct_output5_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct5AsMap

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestStructOutput4_map_with_none(t *testing.T) {
	fnName := "test_struct_output4_map_with_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct4AsMap_with_none

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestSmorgasbordStructOutput_map(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_smorgasbord_struct_output_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testSmorgasbordStructAsMap

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(*map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(*map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(*map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
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
		t.Fatal("expected a result")
	}
	r, ok := result.(*map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestStructOptionOutput5_map(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_struct_option_output5_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct5AsMap

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(*map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestStructOptionOutput4_map_with_none(t *testing.T) {
	fnName := "test_struct_option_output4_map_with_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct4AsMap_with_none

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(*map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("mismatch (-want +got):\n%v", diff)
	}
}

func TestSmorgasbordStructOptionOutput_map(t *testing.T) {
	t.Skip("TODO: fix this")
	fnName := "test_smorgasbord_struct_option_output_map"

	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testSmorgasbordStructAsMap

	if result == nil {
		t.Fatal("expected a result")
	}
	r, ok := result.(*map[string]any)
	if !ok {
		t.Fatalf("expected %T, got %T", expected, result)
	}
	if diff := cmp.Diff(expected, r); diff != "" {
		t.Errorf("unexpected result (-want +got):\n%v", diff)
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

func TestSmorgasbordStructOptionOutput_none(t *testing.T) {
	fnName := "test_smorgasbord_struct_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
