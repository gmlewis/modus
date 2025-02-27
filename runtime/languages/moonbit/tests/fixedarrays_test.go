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

	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=242(), words=0, memBlock=[1 0 0 0 242 0 0 0]
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

	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 8=8+words*4), moonBitType=242(), words=0, memBlock=[1 0 0 0 242 0 0 0]
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
	arr := [0]*int32{}

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

	expected := []*int32{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
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

	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=242(), words=1, memBlock=[1 0 0 0 242 1 0 0 200 59 0 0]
	// memoryBlockAtOffset(offset: 15304=0x00003BC8=[200 59 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
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

	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=242(), words=1, memBlock=[1 0 0 0 242 1 0 0 200 59 0 0]
	// memoryBlockAtOffset(offset: 15304=0x00003BC8=[200 59 0 0], size: 16=8+words*4), moonBitType=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 97 0 98 0 99 0 0 1] = 'abc'
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

	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 12=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 11 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=1, memBlock=[1 0 0 0 241 1 0 0 11 0 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getIntOptionFixedArray1()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
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

	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 16=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 11 0 0 0 0 0 0 0]
	// memoryBlockAtOffset(offset: 48832=0x0000BEC0=[192 190 0 0], size: 24=8+words*4), moonBitType=241(FixedArray[Int]), words=2, memBlock=[1 0 0 0 241 2 0 0 11 0 0 0 0 0 0 0 22 0 0 0 0 0 0 0]
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getIntOptionFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*int32); !ok {
		t.Errorf("expected a []*int32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_string(t *testing.T) {
	fnName := "test_fixedarray_output2_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"abc", "def"}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]string); !ok {
		t.Errorf("expected a []string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_string_option(t *testing.T) {
	fnName := "test_fixedarray_output2_string_option"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getStringOptionFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]*string); !ok {
		t.Errorf("expected a []*string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_struct(t *testing.T) {
	fnName := "test_fixedarray_output2_struct"

	// memoryBlockAtOffset(offset: 92112=0x000167D0=[208 103 1 0], size: 16=8+words*4), classID=242(FixedArray[String]), words=2, memBlock=[1 0 0 0 242 2 0 0 208 102 1 0 176 103 1 0]
	// GML: handler_structs.go: structHandler.Read(offset: 91856=0x000166D0=[208 102 1 0])
	// memoryBlockAtOffset(offset: 91856=0x000166D0=[208 102 1 0], size: 16=8+words*4), classID=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 1 0 0 0 123 0 0 0]
	// GML: handler_structs.go: structHandler.Read: field.Name: 'a', fieldOffset: 91864=0x000166D8=[216 102 1 0]
	// GML: handler_primitives.go: primitiveHandler[[]bool].Read(offset: 91864=0x000166D8=[216 102 1 0])
	// GML: handler_structs.go: structHandler.Read: field.Name: 'b', fieldOffset: 91868=0x000166DC=[220 102 1 0]
	// GML: handler_primitives.go: primitiveHandler[[]int32].Read(offset: 91868=0x000166DC=[220 102 1 0])
	// GML: handler_structs.go: getStructOutput: rt.Kind()=struct
	// GML: handler_structs.go: structHandler.Read(offset: 92080=0x000167B0=[176 103 1 0])
	// memoryBlockAtOffset(offset: 92080=0x000167B0=[176 103 1 0], size: 16=8+words*4), classID=0(Tuple), words=2, memBlock=[1 0 0 0 0 2 0 0 0 0 0 0 200 1 0 0]
	// GML: handler_structs.go: structHandler.Read: field.Name: 'a', fieldOffset: 92088=0x000167B8=[184 103 1 0]
	// GML: handler_primitives.go: primitiveHandler[[]bool].Read(offset: 92088=0x000167B8=[184 103 1 0])
	// GML: handler_structs.go: structHandler.Read: field.Name: 'b', fieldOffset: 92092=0x000167BC=[188 103 1 0]
	// GML: handler_primitives.go: primitiveHandler[[]int32].Read(offset: 92092=0x000167BC=[188 103 1 0])
	// GML: handler_structs.go: getStructOutput: rt.Kind()=struct
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getStructFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]TestStruct2); !ok {
		t.Errorf("expected a []TestStruct2, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

// func TestFixedArrayOutput2_struct_option(t *testing.T) {
// 	fnName := "test_fixedarray_output2_struct_option"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getStructOptionFixedArray2()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]*TestStruct2); !ok {
// 		t.Errorf("expected a []*TestStruct2, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

func TestFixedArrayOutput2_map(t *testing.T) {
	fnName := "test_fixedarray_output2_map"

	// memoryBlockAtOffset(offset: 91856=0x000166D0=[208 102 1 0], size: 16=8+words*4), classID=242(FixedArray[String]), words=2, memBlock=[1 0 0 0 242 2 0 0 144 104 1 0 224 105 1 0]
	// GML: handler_maps.go: mapHandler.Read(offset: 92304)
	// GML: handler_maps.go: DEBUG mapHandler.Decode(vals: [92304])
	// memoryBlockAtOffset(offset: 92304=0x00016890=[144 104 1 0], size: 40=8+words*4), classID=0(Tuple), words=8, memBlock=[1 0 0 0 0 8 0 0 16 104 1 0 96 104 1 0 2 0 0 0 8 0 0 0 7 0 0 0 6 0 0 0 224 104 1 0 48 105 1 0]
	// GML: handler_maps.go: mapHandler.Decode: sliceOffset=[16 104 1 0]=92176
	// GML: handler_maps.go: mapHandler.Decode: numElements=[2 0 0 0]=2
	// memoryBlockAtOffset(offset: 92176=0x00016810=[16 104 1 0], size: 40=8+words*4), classID=242(FixedArray[String]), words=8, memBlock=[1 0 0 0 242 8 0 0 0 0 0 0 0 0 0 0 224 104 1 0 0 0 0 0 0 0 0 0 0 0 0 0 48 105 1 0 0 0 0 0]
	// GML: handler_maps.go: mapHandler.Decode: (sliceOffset: 92176, numElements: 2), sliceMemBlock([16 104 1 0]=@92176)=(40 bytes)=[1 0 0 0 242 8 0 0 0 0 0 0 0 0 0 0 224 104 1 0 0 0 0 0 0 0 0 0 0 0 0 0 48 105 1 0 0 0 0 0]
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[8:12]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[12:16]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[16:20]=[224 104 1 0]=92384
	// memoryBlockAtOffset(offset: 92384=0x000168E0=[224 104 1 0], size: 28=8+words*4), classID=0(Tuple), words=5, memBlock=[3 0 0 0 0 5 0 0 2 0 0 0 0 0 0 0 18 232 28 159 136 40 0 0 248 41 0 0]
	// GML: handler_strings.go: stringHandler.Read(offset: 10376=0x00002888=[136 40 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [10376])
	// memoryBlockAtOffset(offset: 10376=0x00002888=[136 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 65 0 0 1] = 'A'
	// GML: handler_strings.go: stringHandler.Read(offset: 10744=0x000029F8=[248 41 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [10744])
	// memoryBlockAtOffset(offset: 10744=0x000029F8=[248 41 0 0], size: 20=8+words*4), classID=243(String), words=3, memBlock=[255 255 255 255 243 3 0 0 116 0 114 0 117 0 101 0 0 0 0 3] = 'true'
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[20:24]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[24:28]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[28:32]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[32:36]=[48 105 1 0]=92464
	// memoryBlockAtOffset(offset: 92464=0x00016930=[48 105 1 0], size: 28=8+words*4), classID=0(Tuple), words=5, memBlock=[3 0 0 0 0 5 0 0 6 0 0 0 0 0 0 0 110 166 1 68 152 40 0 0 120 63 0 0]
	// GML: handler_strings.go: stringHandler.Read(offset: 10392=0x00002898=[152 40 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [10392])
	// memoryBlockAtOffset(offset: 10392=0x00002898=[152 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 66 0 0 1] = 'B'
	// GML: handler_strings.go: stringHandler.Read(offset: 16248=0x00003F78=[120 63 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [16248])
	// memoryBlockAtOffset(offset: 16248=0x00003F78=[120 63 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 49 0 50 0 51 0 0 1] = '123'
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[36:40]=[0 0 0 0]=0
	// GML: handler_maps.go: mapHandler.Read(offset: 92640)
	// GML: handler_maps.go: DEBUG mapHandler.Decode(vals: [92640])
	// memoryBlockAtOffset(offset: 92640=0x000169E0=[224 105 1 0], size: 40=8+words*4), classID=0(Tuple), words=8, memBlock=[1 0 0 0 0 8 0 0 128 105 1 0 176 105 1 0 2 0 0 0 8 0 0 0 7 0 0 0 6 0 0 0 48 106 1 0 128 106 1 0]
	// GML: handler_maps.go: mapHandler.Decode: sliceOffset=[128 105 1 0]=92544
	// GML: handler_maps.go: mapHandler.Decode: numElements=[2 0 0 0]=2
	// memoryBlockAtOffset(offset: 92544=0x00016980=[128 105 1 0], size: 40=8+words*4), classID=242(FixedArray[String]), words=8, memBlock=[1 0 0 0 242 8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 48 106 1 0 128 106 1 0]
	// GML: handler_maps.go: mapHandler.Decode: (sliceOffset: 92544, numElements: 2), sliceMemBlock([128 105 1 0]=@92544)=(40 bytes)=[1 0 0 0 242 8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 48 106 1 0 128 106 1 0]
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[8:12]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[12:16]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[16:20]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[20:24]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[24:28]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[28:32]=[0 0 0 0]=0
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[32:36]=[48 106 1 0]=92720
	// memoryBlockAtOffset(offset: 92720=0x00016A30=[48 106 1 0], size: 28=8+words*4), classID=0(Tuple), words=5, memBlock=[3 0 0 0 0 5 0 0 6 0 0 0 0 0 0 0 174 32 220 23 168 40 0 0 16 42 0 0]
	// GML: handler_strings.go: stringHandler.Read(offset: 10408=0x000028A8=[168 40 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [10408])
	// memoryBlockAtOffset(offset: 10408=0x000028A8=[168 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 67 0 0 1] = 'C'
	// GML: handler_strings.go: stringHandler.Read(offset: 10768=0x00002A10=[16 42 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [10768])
	// memoryBlockAtOffset(offset: 10768=0x00002A10=[16 42 0 0], size: 20=8+words*4), classID=243(String), words=3, memBlock=[255 255 255 255 243 3 0 0 102 0 97 0 108 0 115 0 101 0 0 1] = 'false'
	// GML: handler_maps.go: readMapKeysAndValues: tupleOffset=sliceMemBlock[36:40]=[128 106 1 0]=92800
	// memoryBlockAtOffset(offset: 92800=0x00016A80=[128 106 1 0], size: 28=8+words*4), classID=0(Tuple), words=5, memBlock=[3 0 0 0 0 5 0 0 7 0 0 0 1 0 0 0 102 35 205 149 184 40 0 0 96 63 0 0]
	// GML: handler_strings.go: stringHandler.Read(offset: 10424=0x000028B8=[184 40 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [10424])
	// memoryBlockAtOffset(offset: 10424=0x000028B8=[184 40 0 0], size: 12=8+words*4), classID=243(String), words=1, memBlock=[255 255 255 255 243 1 0 0 68 0 0 1] = 'D'
	// GML: handler_strings.go: stringHandler.Read(offset: 16224=0x00003F60=[96 63 0 0])
	// GML: handler_strings.go: stringHandler.Decode(vals: [16224])
	// memoryBlockAtOffset(offset: 16224=0x00003F60=[96 63 0 0], size: 16=8+words*4), classID=243(String), words=2, memBlock=[255 255 255 255 243 2 0 0 52 0 53 0 54 0 0 1] = '456'
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getMapFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]map[string]string); !ok {
		t.Errorf("expected a []map[string]string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

// func TestFixedArrayOutput2_map_option(t *testing.T) {
// 	fnName := "test_fixedarray_output2_map_option"
// 	result, err := fixture.CallFunction(t, fnName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expected := getMapOptionFixedArray2()
// 	if result == nil {
// 		t.Error("expected a result")
// 	} else if r, ok := result.([]*map[string]string); !ok {
// 		t.Errorf("expected a []*map[string]string, got %T", result)
// 	} else if !reflect.DeepEqual(expected, r) {
// 		t.Errorf("expected %v, got %v", expected, r)
// 	}
// }

func TestOptionFixedArrayInput1_int(t *testing.T) {
	fnName := "test_option_fixedarray_input1_int"
	arr := getOptionIntFixedArray1()

	if _, err := fixture.CallFunction(t, fnName, arr); err != nil {
		t.Error(err)
	}
	if arr, ok := utils.ConvertToSliceOf[any](*arr); !ok {
		t.Error("failed conversion to interface slice")
	} else if _, err := fixture.CallFunction(t, fnName, &arr); err != nil {
		t.Error(err)
	}
}

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

func TestOptionFixedArrayOutput1_int(t *testing.T) {
	fnName := "test_option_fixedarray_output1_int"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getOptionIntFixedArray1()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*[]int32); !ok {
		t.Errorf("expected a *[]int32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestOptionFixedArrayOutput2_int(t *testing.T) {
	fnName := "test_option_fixedarray_output2_int"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getOptionIntFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*[]int32); !ok {
		t.Errorf("expected a *[]int32, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestOptionFixedArrayOutput1_string(t *testing.T) {
	fnName := "test_option_fixedarray_output1_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getOptionStringFixedArray1()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*[]string); !ok {
		t.Errorf("expected a *[]string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestOptionFixedArrayOutput2_string(t *testing.T) {
	fnName := "test_option_fixedarray_output2_string"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := getOptionStringFixedArray2()
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*[]string); !ok {
		t.Errorf("expected a *[]string, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

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

func TestFixedArrayOutput0_byte(t *testing.T) {
	fnName := "test_fixedarray_output0_byte"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]byte); !ok {
		t.Errorf("expected a []byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput1_byte(t *testing.T) {
	fnName := "test_fixedarray_output1_byte"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{1}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]byte); !ok {
		t.Errorf("expected a []byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestFixedArrayOutput2_byte(t *testing.T) {
	fnName := "test_fixedarray_output2_byte"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{1, 2}
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.([]byte); !ok {
		t.Errorf("expected a []byte, got %T", result)
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func getIntOptionFixedArray1() []*int32 {
	a := int32(11)
	return []*int32{&a}
}

func getIntOptionFixedArray2() []*int32 {
	a := int32(11)
	b := int32(22)
	return []*int32{&a, &b}
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

func getMapFixedArray2() []map[string]string {
	return []map[string]string{
		{"A": "true", "B": "123"},
		{"C": "false", "D": "456"},
	}
}

func getMapOptionFixedArray2() []*map[string]string {
	return []*map[string]string{
		{"A": "true", "B": "123"},
		{"C": "false", "D": "456"},
	}
}

func getOptionIntFixedArray1() *[]int32 {
	return &[]int32{11}
}

func getOptionIntFixedArray2() *[]int32 {
	return &[]int32{11, 22}
}

func getOptionStringFixedArray1() *[]string {
	a := "abc"
	return &[]string{a}
}

func getOptionStringFixedArray2() *[]string {
	a := "abc"
	b := "def"
	return &[]string{a, b}
}
