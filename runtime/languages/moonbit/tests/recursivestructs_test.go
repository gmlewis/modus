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

type TestRecursiveStruct struct {
	A bool
	B *TestRecursiveStruct
}

var testRecursiveStruct = func() TestRecursiveStruct {
	r := TestRecursiveStruct{
		A: true,
	}
	r.B = &r
	return r
}()

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

// TODO:
// func TestRecursiveStructInput(t *testing.T) {
// 	fnName := "test_recursive_struct_input"
// 	if _, err := fixture.CallFunction(t, fnName, testRecursiveStruct); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestRecursiveStructOptionInput(t *testing.T) {
// 	fnName := "test_recursive_struct_option_input"
// 	if _, err := fixture.CallFunction(t, fnName, testRecursiveStruct); err != nil {
// 		t.Error(err)
// 	}
// 	fnName = "test_recursive_struct_option_input"
// 	if _, err := fixture.CallFunction(t, fnName, &testRecursiveStruct); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestRecursiveStructInput_map(t *testing.T) {
// 	fnName := "test_recursive_struct_input"
// 	if _, err := fixture.CallFunction(t, fnName, testRecursiveStructAsMap); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestRecursiveStructOptionInput_map(t *testing.T) {
// 	fnName := "test_recursive_struct_input"
// 	if _, err := fixture.CallFunction(t, fnName, testRecursiveStructAsMap); err != nil {
// 		t.Error(err)
// 	}
// 	fnName = "test_recursive_struct_input"
// 	if _, err := fixture.CallFunction(t, fnName, &testRecursiveStructAsMap); err != nil {
// 		t.Error(err)
// 	}
// }

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
		t.Error(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(TestRecursiveStruct); !ok {
		t.Errorf("expected a struct, got %T", result)
	} else if !reflect.DeepEqual(testRecursiveStruct, r) {
		t.Errorf("expected %v, got %v", testRecursiveStruct, r)
	}
}

func TestRecursiveStructOptionOutput(t *testing.T) {
	fnName := "test_recursive_struct_option_output"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*TestRecursiveStruct); !ok {
		t.Errorf("expected a pointer to a struct, got %T", result)
	} else if !reflect.DeepEqual(testRecursiveStruct, *r) {
		t.Errorf("expected %v, got %v", testRecursiveStruct, *r)
	}
}

func TestRecursiveStructOutput_map(t *testing.T) {
	fnName := "test_recursive_struct_output_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(map[string]any); !ok {
		t.Errorf("expected a map[string]any, got %T", result)
		// reflect.DeepEqual does not work here with two self-referencing maps.
		// } else if !reflect.DeepEqual(testRecursiveStructAsMap, r) {
	} else if r["a"] != true {
		t.Errorf("expected r.a=true, got %v", r["a"])
	} else if r["b"] == nil {
		t.Errorf("expected r.b!=nil, got %v", r["b"])
	}
}

func TestRecursiveStructOptionOutput_map(t *testing.T) {
	fnName := "test_recursive_struct_option_output_map"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*map[string]any); !ok {
		t.Errorf("expected a map[string]any, got %T", result)
		// reflect.DeepEqual does not work here with two self-referencing maps.
		// } else if !reflect.DeepEqual(testRecursiveStructAsMap, r) {
	} else if (*r)["a"] != true {
		t.Errorf("expected r.a=true, got %v", (*r)["a"])
	} else if (*r)["b"] == nil {
		t.Errorf("expected r.b!=nil, got %v", (*r)["b"])
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
