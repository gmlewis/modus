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

	// memoryBlockAtOffset(offset: 48064=0x0000BBC0=[192 187 0 0], size: 12=8+words*4), moonBitType=0(Tuple), words=1, memBlock=[2 0 0 0 0 1 0 0 1 0 0 0]
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
		t.Error("expected a result")
	} else if r, ok := result.(TestStruct4); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
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

	// memoryBlockAtOffset(offset: 93344=0x00016CA0=[160 108 1 0], size: 40=8+words*4), classID=0(Tuple), words=8, memBlock=[2 0 0 0 0 8 0 0 48 56 0 0 200 59 0 0 104 161 0 0 128 108 1 0 124 242 176 80 107 154 191 63 238 148 14 214 255 255 88 64]
	// GML: handler_structs.go: structHandler.Read(offset: 93344=0x00016CA0=[160 108 1 0])
	// memoryBlockAtOffset(offset: 93344=0x00016CA0=[160 108 1 0], size: 40=8+words*4), classID=0(Tuple), words=8, memBlock=[2 0 0 0 0 8 0 0 48 56 0 0 200 59 0 0 104 161 0 0 128 108 1 0 124 242 176 80 107 154 191 63 238 148 14 214 255 255 88 64]
	// GML: handler_structs.go: structHandler.Read: field.Name: 'a', type: 'String', fieldOffset: 0, uint32 ptr: 14384=0x00003830=[48 56 0 0]
	// GML: handler_strings.go: stringHandler.Decode(vals: [14384])
	// memoryBlockAtOffset(offset: 14384=0x00003830=[48 56 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
	// GML: handler_structs.go: structHandler.Read: field.Name: 'b', type: 'String', fieldOffset: 4, uint32 ptr: 15304=0x00003BC8=[200 59 0 0]
	// GML: handler_strings.go: stringHandler.Decode(vals: [15304])
	// memoryBlockAtOffset(offset: 15304=0x00003BC8=[200 59 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 100 0 101 0 102 0 0 1] = 'def'
	// GML: handler_structs.go: structHandler.Read: field.Name: 'c', type: 'String', fieldOffset: 8, uint32 ptr: 41320=0x0000A168=[104 161 0 0]
	// GML: handler_strings.go: stringHandler.Decode(vals: [41320])
	// memoryBlockAtOffset(offset: 41320=0x0000A168=[104 161 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 103 0 104 0 105 0 0 1] = 'ghi'
	// GML: handler_structs.go: structHandler.Read: field.Name: 'd', type: 'Array[String]', fieldOffset: 12, uint32 ptr: 93312=0x00016C80=[128 108 1 0]
	// GML: handler_slices.go: sliceHandler.Decode(vals: [93312])
	// memoryBlockAtOffset(offset: 93312=0x00016C80=[128 108 1 0], size: 16=8+words*4), classID=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 96 108 1 0 3 0 0 0]
	// memoryBlockAtOffset(offset: 93280=0x00016C60=[96 108 1 0], size: 20=8+words*4), classID=242(FixedArray[String]), words=3, memBlock=[1 0 0 0 242 3 0 0 120 165 0 0 96 165 0 0 72 165 0 0]
	// GML: handler_slices.go: sliceHandler.Decode: sliceOffset=93280=0x00016C60=[96 108 1 0], numElements=3, size=12
	// memoryBlockAtOffset(offset: 93280=0x00016C60=[96 108 1 0], size: 20=8+words*4), classID=242(FixedArray[String]), words=3, memBlock=[1 0 0 0 242 3 0 0 120 165 0 0 96 165 0 0 72 165 0 0]
	// GML: handler_strings.go: stringHandler.Read(offset: 42360=0x0000A578=[120 165 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [42360])
	// memoryBlockAtOffset(offset: 42360=0x0000A578=[120 165 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 106 0 107 0 108 0 0 1] = 'jkl'
	// GML: handler_strings.go: stringHandler.Read(offset: 42336=0x0000A560=[96 165 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [42336])
	// memoryBlockAtOffset(offset: 42336=0x0000A560=[96 165 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 109 0 110 0 111 0 0 1] = 'mno'
	// GML: handler_strings.go: stringHandler.Read(offset: 42312=0x0000A548=[72 165 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [42312])
	// memoryBlockAtOffset(offset: 42312=0x0000A548=[72 165 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 112 0 113 0 114 0 0 1] = 'pqr'
	// GML: handler_structs.go: structHandler.Read: field.Name: 'e', type: 'Double', fieldOffset: 16, uint64 ptr: 4593559930647147132
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4593559930647147132])
	// GML: handler_structs.go: structHandler.Read: field.Name: 'f', type: 'Double', fieldOffset: 24, uint64 ptr: 4636737290650948846
	// GML: handler_primitives.go: primitiveHandler.Decode(vals: [4636737290650948846])
	// GML: handler_structs.go: getStructOutput: rt.Kind()=struct
	// GML: handler_pointers.go: pointerHandler.readData: data=moonbit_test.TestStruct5{A:"abc", B:"def", C:"ghi", D:[]string{"jkl", "mno", "pqr"}, E:0.12345, F:99.99999}
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

func TestStructOptionOutput4_with_none(t *testing.T) {
	fnName := "test_struct_option_output4_with_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct4_with_none

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

func TestStructOutput4_map_with_none(t *testing.T) {
	fnName := "test_struct_output4_map_with_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := testStruct4AsMap_with_none

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

func TestStructOptionOutput4_map_with_none(t *testing.T) {
	fnName := "test_struct_option_output4_map_with_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := &testStruct4AsMap_with_none

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
