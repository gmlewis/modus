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
	Bool:       true,               // 1 0 0 0
	Byte:       0x12,               // 18 0 0 0
	C:          'c',                // 99 0 0 0
	F:          1.23,               // 164 112 157 63
	D:          4.56,               // 61 10 215 163 112 61 18 64
	I16:        123,                // 123 0 0 0
	I32:        456,                // 200 1 0 0
	I64:        789,                // 21 3 0 0 0 0 0 0
	S:          "abc",              // ptr, e.g. 136 56 0 0
	U16:        123,                // 123 0 0 0
	U32:        456,                // 200 1 0 0
	U64:        789,                // 21 3 0 0 0 0 0 0
	SomeBool:   Ptr(true),          // 1 0 0 0
	NoneBool:   nil,                // 255 255 255 255
	SomeByte:   Ptr(byte(0x34)),    // 52 0 0 0
	NoneByte:   nil,                // 255 255 255 255
	SomeChar:   Ptr(int16('d')),    // 100 0 0 0
	NoneChar:   nil,                // 255 255 255 255
	SomeFloat:  Ptr(float32(7.89)), // ptr, e.g. 0 117 1 0
	NoneFloat:  nil,                // ptr to None, e.g. 240 41 0 0
	SomeDouble: Ptr(float64(0.12)), // ptr, e.g. 32 117 1 0
	NoneDouble: nil,                // ptr to None, e.g. 240 41 0 0
	SomeI16:    Ptr(int16(234)),    // 234 0 0 0
	NoneI16:    nil,                // 255 255 255 255
	SomeI32:    Ptr(int32(567)),    // 55 2 0 0 0 0 0 0
	NoneI32:    nil,                // 0 0 0 0 1 0 0 0
	SomeI64:    Ptr(int64(890)),    // ptr, e.g. 64 117 1 0
	NoneI64:    nil,                // ptr to None, e.g. 240 41 0 0
	SomeString: Ptr("def"),         // ptr, e.g. 32 60 0 0
	NoneString: nil,                // 0 0 0 0
	SomeU16:    Ptr(uint16(234)),   // 234 0 0 0
	NoneU16:    nil,                // 255 255 255 255
	SomeU32:    Ptr(uint32(567)),   // 55 2 0 0 0 0 0 0
	NoneU32:    nil,                // 0 0 0 0 1 0 0 0
	SomeU64:    Ptr(uint64(890)),   // ptr, e.g. 96 117 1 0
	NoneU64:    nil,                // ptr to None, e.g. 240 41 0 0
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
	fnName := "test_smorgasbord_struct_option_input"

	// GOT:  memoryBlockAtOffset(offset: 96912=0x00017A90=[144 122 1 0], size: 180=8+words*4), classID=0(Tuple), words=43, memBlock=[1 0 0 0 0 43 0 0 1 0 0 0 18 0 0 0 99 0 0 0 164 112 157 63 61 10 215 163 112 61 18 64 123 0 0 0 200 1 0 0 21 3 0 0 0 0 0 0 176 121 1 0 123 0 0 0 200 1 0 0 21 3 0 0 0 0 0 0 1 0 0 0 255 255 255 255 52 0 0 0 255 255 255 255 100 0 0 0 255 255 255 255 88 123 1 0 240 41 0 0 120 123 1 0 240 41 0 0 234 0 0 0 255 255 255 255 55 2 0 0 0 0 0 0 0 0 0 0 1 0 0 0 152 123 1 0 240 41 0 0 176 123 1 0 0 0 0 0 234 0 0 0 255 255 255 255 55 2 0 0 0 0 0 0 0 0 0 0 1 0 0 0 216 123 1 0 240 41 0 0]
	// WANT: memoryBlockAtOffset(offset: 95616=0x00017580=[128 117 1 0], size: 180=8+words*4), classID=0(Tuple), words=43, memBlock=[2 0 0 0 0 43 0 0 1 0 0 0 18 0 0 0 99 0 0 0 164 112 157 63 61 10 215 163 112 61 18 64 123 0 0 0 200 1 0 0 21 3 0 0 0 0 0 0 136 56 0 0 123 0 0 0 200 1 0 0 21 3 0 0 0 0 0 0 1 0 0 0 255 255 255 255 52 0 0 0 255 255 255 255 100 0 0 0 255 255 255 255 0 117 1 0 240 41 0 0 32 117 1 0 240 41 0 0 234 0 0 0 255 255 255 255 55 2 0 0 0 0 0 0 0 0 0 0 1 0 0 0 64 117 1 0 240 41 0 0 32 60 0 0 0 0 0 0 234 0 0 0 255 255 255 255 55 2 0 0 0 0 0 0 0 0 0 0 1 0 0 0 96 117 1 0 240 41 0 0]
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

	// memoryBlockAtOffset(offset: 48064=0x0000BBC0=[192 187 0 0], size: 12=8+words*4), moonBitType=0(Tuple), words=1, memBlock=[2 0 0 0 0 1 0 0 1 0 0 0]
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

	// memoryBlockAtOffset(offset: 91184=0x00016430=[48 100 1 0], size: 20=8+words*4), classID=0(Tuple), words=3, memBlock=[2 0 0 0 0 3 0 0 1 0 0 0 123 0 0 0 48 56 0 0]
	// GML: handler_primitives.go: primitiveHandler[[]bool].Read(offset: 91192=0x00016438=[56 100 1 0])
	// GML: handler_structs.go: structHandler.Decode: field.Name: 'a', Alignment: 4, DataSize: 4, EncodingLength: 1, Size: 4
	// GML: handler_primitives.go: primitiveHandler[[]int32].Read(offset: 91196=0x0001643C=[60 100 1 0])
	// GML: handler_structs.go: structHandler.Decode: field.Name: 'b', Alignment: 4, DataSize: 4, EncodingLength: 1, Size: 4
	// GML: handler_strings.go: stringHandler.Decode(vals: [14384])
	// memoryBlockAtOffset(offset: 14384=0x00003830=[48 56 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
	// GML: handler_structs.go: structHandler.Decode: field.Name: 'c', Alignment: 4, DataSize: 4, EncodingLength: 1, Size: 4
	// GML: handler_structs.go: getStructOutput: rt.Kind()=struct
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

	// memoryBlockAtOffset(offset: 91328=0x000164C0=[192 100 1 0], size: 20=8+words*4), classID=0(Tuple), words=3, memBlock=[2 0 0 0 0 3 0 0 1 0 0 0 123 0 0 0 0 0 0 0]
	// GML: handler_primitives.go: primitiveHandler[[]bool].Read(offset: 91336=0x000164C8=[200 100 1 0])
	// GML: handler_structs.go: structHandler.Decode: field.Name: 'a', Alignment: 4, DataSize: 4, EncodingLength: 1, Size: 4
	// GML: handler_primitives.go: primitiveHandler[[]int32].Read(offset: 91340=0x000164CC=[204 100 1 0])
	// GML: handler_structs.go: structHandler.Decode: field.Name: 'b', Alignment: 4, DataSize: 4, EncodingLength: 1, Size: 4
	// GML: handler_strings.go: stringHandler.Decode(vals: [0])
	// GML: handler_structs.go: structHandler.Decode: field.Name: 'c', Alignment: 4, DataSize: 4, EncodingLength: 1, Size: 4
	// GML: handler_structs.go: getStructOutput: rt.Kind()=struct
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
	fnName := "test_smorgasbord_struct_output"

	// memoryBlockAtOffset(offset: 95616=0x00017580=[128 117 1 0], size: 180=8+words*4), classID=0(Tuple), words=43, memBlock=[2 0 0 0 0 43 0 0 1 0 0 0 18 0 0 0 99 0 0 0 164 112 157 63 61 10 215 163 112 61 18 64 123 0 0 0 200 1 0 0 21 3 0 0 0 0 0 0 136 56 0 0 123 0 0 0 200 1 0 0 21 3 0 0 0 0 0 0 1 0 0 0 255 255 255 255 52 0 0 0 255 255 255 255 100 0 0 0 255 255 255 255 0 117 1 0 240 41 0 0 32 117 1 0 240 41 0 0 234 0 0 0 255 255 255 255 55 2 0 0 0 0 0 0 0 0 0 0 1 0 0 0 64 117 1 0 240 41 0 0 32 60 0 0 0 0 0 0 234 0 0 0 255 255 255 255 55 2 0 0 0 0 0 0 0 0 0 0 1 0 0 0 96 117 1 0 240 41 0 0]
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

	// GML: handler_structs.go: structHandler.Read(offset: 91088=0x000163D0=[208 99 1 0])
	// memoryBlockAtOffset(offset: 91088=0x000163D0=[208 99 1 0], size: 12=8+words*4), classID=0(Tuple), words=1, memBlock=[2 0 0 0 0 1 0 0 1 0 0 0]
	// GML: handler_structs.go: structHandler.Read: field.Name: 'a', fieldOffset: 91096=0x000163D8=[216 99 1 0]
	// GML: handler_primitives.go: primitiveHandler[[]bool].Read(offset: 91096=0x000163D8=[216 99 1 0])
	// GML: handler_structs.go: getStructOutput: rt.Kind()=struct
	// GML: handler_pointers.go: pointerHandler.readData: data=moonbit_test.TestStruct1{A:true}
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
	fnName := "test_struct_option_output5"

	// memoryBlockAtOffset(offset: 93344=0x00016CA0=[160 108 1 0], size: 40=8+words*4), classID=0(Tuple), words=8, memBlock=[2 0 0 0 0 8 0 0 48 56 0 0 200 59 0 0 104 161 0 0 128 108 1 0 124 242 176 80 107 154 191 63 238 148 14 214 255 255 88 64]
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
	fnName := "test_smorgasbord_struct_option_output_map"

	// memoryBlockAtOffset(offset: 96272=0x00017810=[16 120 1 0], size: 180=8+words*4), classID=0(Tuple), words=43, memBlock=[2 0 0 0 0 43 0 0 1 0 0 0 18 0 0 0 99 0 0 0 164 112 157 63 61 10 215 163 112 61 18 64 123 0 0 0 200 1 0 0 21 3 0 0 0 0 0 0 136 56 0 0 123 0 0 0 200 1 0 0 21 3 0 0 0 0 0 0 1 0 0 0 255 255 255 255 52 0 0 0 255 255 255 255 100 0 0 0 255 255 255 255 144 119 1 0 240 41 0 0 176 119 1 0 240 41 0 0 234 0 0 0 255 255 255 255 55 2 0 0 0 0 0 0 0 0 0 0 1 0 0 0 208 119 1 0 240 41 0 0 32 60 0 0 0 0 0 0 234 0 0 0 255 255 255 255 55 2 0 0 0 0 0 0 0 0 0 0 1 0 0 0 240 119 1 0 240 41 0 0]
	// GML: handler_structs.go: structHandler.Read: fieldName: 'bool', type: 'Bool', fieldOffset: 0, uint32 ptr: 1=0x00000001=[1 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [1])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'byte', type: 'Byte', fieldOffset: 4, uint32 ptr: 18=0x00000012=[18 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [18])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'c', type: 'Char', fieldOffset: 8, uint32 ptr: 99=0x00000063=[99 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [99])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'f', type: 'Float', fieldOffset: 12, uint32 ptr: 1067282596=0x3F9D70A4=[164 112 157 63]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [1067282596])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'd', type: 'Double', fieldOffset: 16, uint64 ptr: 4616820122002590269
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4616820122002590269])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'i16', type: 'Int16', fieldOffset: 24, uint32 ptr: 123=0x0000007B=[123 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [123])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'i32', type: 'Int', fieldOffset: 28, uint32 ptr: 456=0x000001C8=[200 1 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [456])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'i64', type: 'Int64', fieldOffset: 32, uint64 ptr: 789
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [789])
	// GML: handler_structs.go: structHandler.Read: fieldName: 's', type: 'String', fieldOffset: 40, uint32 ptr: 14472=0x00003888=[136 56 0 0]
	// GML: handler_strings.go: stringHandler.Decode(vals: [14472])
	// memoryBlockAtOffset(offset: 14472=0x00003888=[136 56 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
	// GML: handler_structs.go: structHandler.Read: fieldName: 'u16', type: 'UInt16', fieldOffset: 44, uint32 ptr: 123=0x0000007B=[123 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [123])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'u32', type: 'UInt', fieldOffset: 48, uint32 ptr: 456=0x000001C8=[200 1 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [456])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'u64', type: 'UInt64', fieldOffset: 52, uint64 ptr: 789
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [789])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someBool', type: 'Bool?', fieldOffset: 60, uint32 ptr: 1=0x00000001=[1 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [1])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneBool', type: 'Bool?', fieldOffset: 64, uint32 ptr: 4294967295=0xFFFFFFFF=[255 255 255 255]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4294967295])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someByte', type: 'Byte?', fieldOffset: 68, uint32 ptr: 52=0x00000034=[52 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [52])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneByte', type: 'Byte?', fieldOffset: 72, uint32 ptr: 4294967295=0xFFFFFFFF=[255 255 255 255]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4294967295])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someChar', type: 'Char?', fieldOffset: 76, uint32 ptr: 100=0x00000064=[100 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [100])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneChar', type: 'Char?', fieldOffset: 80, uint32 ptr: 4294967295=0xFFFFFFFF=[255 255 255 255]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4294967295])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someFloat', type: 'Float?', fieldOffset: 84, uint32 ptr: 96144=0x00017790=[144 119 1 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [96144])
	// memoryBlockAtOffset(offset: 96144=0x00017790=[144 119 1 0], size: 12=8+words*4), classID=1(), words=1, memBlock=[1 0 0 0 1 1 0 0 225 122 252 64]
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneFloat', type: 'Float?', fieldOffset: 88, uint32 ptr: 10736=0x000029F0=[240 41 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [10736])
	// memoryBlockAtOffset(offset: 10736=0x000029F0=[240 41 0 0], size: 8=8+words*4), classID=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someDouble', type: 'Double?', fieldOffset: 92, uint32 ptr: 96176=0x000177B0=[176 119 1 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [96176])
	// memoryBlockAtOffset(offset: 96176=0x000177B0=[176 119 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 184 30 133 235 81 184 190 63]
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneDouble', type: 'Double?', fieldOffset: 96, uint32 ptr: 10736=0x000029F0=[240 41 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [10736])
	// memoryBlockAtOffset(offset: 10736=0x000029F0=[240 41 0 0], size: 8=8+words*4), classID=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someI16', type: 'Int16?', fieldOffset: 100, uint32 ptr: 234=0x000000EA=[234 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [234])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneI16', type: 'Int16?', fieldOffset: 104, uint32 ptr: 4294967295=0xFFFFFFFF=[255 255 255 255]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4294967295])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someI32', type: 'Int?', fieldOffset: 108, uint64 ptr: 567
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [567])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneI32', type: 'Int?', fieldOffset: 116, uint64 ptr: 4294967296
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4294967296])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someI64', type: 'Int64?', fieldOffset: 124, uint32 ptr: 96208=0x000177D0=[208 119 1 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [96208])
	// memoryBlockAtOffset(offset: 96208=0x000177D0=[208 119 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 122 3 0 0 0 0 0 0]
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneI64', type: 'Int64?', fieldOffset: 128, uint32 ptr: 10736=0x000029F0=[240 41 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [10736])
	// memoryBlockAtOffset(offset: 10736=0x000029F0=[240 41 0 0], size: 8=8+words*4), classID=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someString', type: 'String?', fieldOffset: 132, uint32 ptr: 15392=0x00003C20=[32 60 0 0]
	// GML: handler_strings.go: stringHandler.Decode(vals: [15392])
	// memoryBlockAtOffset(offset: 15392=0x00003C20=[32 60 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 100 0 101 0 102 0 0 1] = 'def'
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneString', type: 'String?', fieldOffset: 136, uint32 ptr: 0=0x00000000=[0 0 0 0]
	// GML: handler_strings.go: stringHandler.Decode(vals: [0])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someU16', type: 'UInt16?', fieldOffset: 140, uint32 ptr: 234=0x000000EA=[234 0 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [234])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneU16', type: 'UInt16?', fieldOffset: 144, uint32 ptr: 4294967295=0xFFFFFFFF=[255 255 255 255]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4294967295])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someU32', type: 'UInt?', fieldOffset: 148, uint64 ptr: 567
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [567])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneU32', type: 'UInt?', fieldOffset: 156, uint64 ptr: 4294967296
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4294967296])
	// GML: handler_structs.go: structHandler.Read: fieldName: 'someU64', type: 'UInt64?', fieldOffset: 164, uint32 ptr: 96240=0x000177F0=[240 119 1 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [96240])
	// memoryBlockAtOffset(offset: 96240=0x000177F0=[240 119 1 0], size: 16=8+words*4), classID=1(), words=2, memBlock=[1 0 0 0 1 2 0 0 122 3 0 0 0 0 0 0]
	// GML: handler_structs.go: structHandler.Read: fieldName: 'noneU64', type: 'UInt64?', fieldOffset: 168, uint32 ptr: 10736=0x000029F0=[240 41 0 0]
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [10736])
	// memoryBlockAtOffset(offset: 10736=0x000029F0=[240 41 0 0], size: 8=8+words*4), classID=0(Tuple), words=0, memBlock=[255 255 255 255 0 0 0 0]
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
