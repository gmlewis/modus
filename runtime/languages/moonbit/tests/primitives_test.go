// -*- compile-command: "go test -timeout 30s -tags integration -run ^TestPrimitives github.com/gmlewis/modus/runtime/languages/moonbit/tests"; -*-

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
	"math"
	"testing"

	"github.com/gmlewis/modus/runtime/utils"
)

func TestPrimitivesBoolInput_false(t *testing.T) {
	fnName := "test_bool_input_false"
	if _, err := fixture.CallFunction(t, fnName, false); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBoolOutput_false(t *testing.T) {
	fnName := "test_bool_output_false"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := false
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(bool); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesBoolInput_true(t *testing.T) {
	fnName := "test_bool_input_true"
	if _, err := fixture.CallFunction(t, fnName, true); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBoolOutput_true(t *testing.T) {
	fnName := "test_bool_output_true"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := true
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(bool); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesBoolOptionInput_false(t *testing.T) {
	fnName := "test_bool_option_input_false"
	b := false

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBoolOptionOutput_false(t *testing.T) {
	fnName := "test_bool_option_output_false"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := false
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*bool); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesBoolOptionInput_true(t *testing.T) {
	fnName := "test_bool_option_input_true"
	b := true

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBoolOptionOutput_true(t *testing.T) {
	fnName := "test_bool_option_output_true"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := true
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*bool); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesBoolOptionInput_none(t *testing.T) {
	fnName := "test_bool_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesBoolOptionOutput_none(t *testing.T) {
	fnName := "test_bool_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if !utils.HasNil(result) {
		t.Errorf("expected a nil result, got %T=%[1]v", result)
	}
}

func TestPrimitivesByteInput_min(t *testing.T) {
	fnName := "test_byte_input_min"
	if _, err := fixture.CallFunction(t, fnName, byte(0)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesByteOutput_min(t *testing.T) {
	fnName := "test_byte_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := byte(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(byte); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesByteInput_max(t *testing.T) {
	fnName := "test_byte_input_max"
	if _, err := fixture.CallFunction(t, fnName, byte(math.MaxUint8)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesByteOutput_max(t *testing.T) {
	fnName := "test_byte_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := byte(math.MaxUint8)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(byte); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesByteOptionInput_min(t *testing.T) {
	fnName := "test_byte_option_input_min"
	b := byte(0)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesByteOptionOutput_min(t *testing.T) {
	fnName := "test_byte_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := byte(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*byte); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesByteOptionInput_max(t *testing.T) {
	fnName := "test_byte_option_input_max"
	b := byte(math.MaxUint8)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesByteOptionOutput_max(t *testing.T) {
	fnName := "test_byte_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := byte(math.MaxUint8)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*byte); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesByteOptionInput_none(t *testing.T) {
	fnName := "test_byte_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesByteOptionOutput_none(t *testing.T) {
	fnName := "test_byte_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestPrimitivesCharInput_min(t *testing.T) {
	fnName := "test_char_input_min"
	if _, err := fixture.CallFunction(t, fnName, 0); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesCharOutput_min(t *testing.T) {
	fnName := "test_char_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(-32768)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int16); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesCharInput_max(t *testing.T) {
	fnName := "test_char_input_max"
	if _, err := fixture.CallFunction(t, fnName, rune(math.MaxInt16)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesCharOutput_max(t *testing.T) {
	fnName := "test_char_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MaxInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int16); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesCharOptionInput_min(t *testing.T) {
	fnName := "test_char_option_input_min"
	b := rune(math.MinInt16)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesCharOptionOutput_min(t *testing.T) {
	fnName := "test_char_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MinInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int16); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesCharOptionInput_max(t *testing.T) {
	fnName := "test_char_option_input_max"
	b := rune(math.MaxInt16)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesCharOptionOutput_max(t *testing.T) {
	fnName := "test_char_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MaxInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int16); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesCharOptionInput_none(t *testing.T) {
	fnName := "test_char_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesCharOptionOutput_none(t *testing.T) {
	fnName := "test_char_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestPrimitivesIntInput_min(t *testing.T) {
	fnName := "test_int_input_min"
	if _, err := fixture.CallFunction(t, fnName, int(math.MinInt)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesIntOutput_min(t *testing.T) {
	fnName := "test_int_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int(math.MinInt32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesIntInput_max(t *testing.T) {
	fnName := "test_int_input_max"
	if _, err := fixture.CallFunction(t, fnName, int(math.MaxInt)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesIntOutput_max(t *testing.T) {
	fnName := "test_int_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int(math.MaxInt32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesIntOptionInput_min(t *testing.T) {
	fnName := "test_int_option_input_min"
	b := int(math.MinInt)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesIntOptionOutput_min(t *testing.T) {
	fnName := "test_int_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int(math.MinInt32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesIntOptionInput_max(t *testing.T) {
	fnName := "test_int_option_input_max"
	b := int(math.MaxInt)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesIntOptionOutput_max(t *testing.T) {
	fnName := "test_int_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int(math.MaxInt32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesIntOptionInput_none(t *testing.T) {
	fnName := "test_int_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesIntOptionOutput_none(t *testing.T) {
	fnName := "test_int_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestPrimitivesInt16Input_min(t *testing.T) {
	fnName := "test_int16_input_min"
	if _, err := fixture.CallFunction(t, fnName, int16(math.MinInt16)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt16Output_min(t *testing.T) {
	fnName := "test_int16_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MinInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int16); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesInt16Input_max(t *testing.T) {
	fnName := "test_int16_input_max"
	if _, err := fixture.CallFunction(t, fnName, int16(math.MaxInt16)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt16Output_max(t *testing.T) {
	fnName := "test_int16_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MaxInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int16); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesInt16OptionInput_min(t *testing.T) {
	fnName := "test_int16_option_input_min"
	b := int16(math.MinInt16)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt16OptionOutput_min(t *testing.T) {
	fnName := "test_int16_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MinInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int16); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesInt16OptionInput_max(t *testing.T) {
	fnName := "test_int16_option_input_max"
	b := int16(math.MaxInt16)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt16OptionOutput_max(t *testing.T) {
	fnName := "test_int16_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int16(math.MaxInt16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int16); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesInt16OptionInput_none(t *testing.T) {
	fnName := "test_int16_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt16OptionOutput_none(t *testing.T) {
	fnName := "test_int16_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestPrimitivesInt64Input_min(t *testing.T) {
	fnName := "test_int64_input_min"
	if _, err := fixture.CallFunction(t, fnName, int64(math.MinInt64)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt64Output_min(t *testing.T) {
	fnName := "test_int64_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int64(math.MinInt64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesInt64Input_max(t *testing.T) {
	fnName := "test_int64_input_max"
	if _, err := fixture.CallFunction(t, fnName, int64(math.MaxInt64)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt64Output_max(t *testing.T) {
	fnName := "test_int64_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int64(math.MaxInt64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesInt64OptionInput_min(t *testing.T) {
	fnName := "test_int64_option_input_min"
	b := int64(math.MinInt64)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt64OptionOutput_min(t *testing.T) {
	fnName := "test_int64_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int64(math.MinInt64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int64); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesInt64OptionInput_max(t *testing.T) {
	fnName := "test_int64_option_input_max"
	b := int64(math.MaxInt64)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesInt64OptionOutput_max(t *testing.T) {
	fnName := "test_int64_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := int64(math.MaxInt64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*int64); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

//	func TestPrimitivesInt64OptionInput_none(t *testing.T) {
//		fnName := "test_int64_option_input_none"
//		if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
//			t.Error(err)
//		}
//	}

func TestPrimitivesInt64OptionOutput_none(t *testing.T) {
	fnName := "test_int64_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestPrimitivesUintInput_min(t *testing.T) {
	fnName := "test_uint_input_min"
	if _, err := fixture.CallFunction(t, fnName, uint(0)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUintOutput_min(t *testing.T) {
	fnName := "test_uint_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(uint); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesUintInput_max(t *testing.T) {
	fnName := "test_uint_input_max"
	if _, err := fixture.CallFunction(t, fnName, uint(math.MaxUint)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUintOutput_max(t *testing.T) {
	fnName := "test_uint_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint(math.MaxUint32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(uint); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesUintOptionInput_min(t *testing.T) {
	fnName := "test_uint_option_input_min"
	b := uint(0)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUintOptionOutput_min(t *testing.T) {
	fnName := "test_uint_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*uint); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesUintOptionInput_max(t *testing.T) {
	fnName := "test_uint_option_input_max"
	b := uint(math.MaxUint)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUintOptionOutput_max(t *testing.T) {
	fnName := "test_uint_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint(math.MaxUint32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*uint); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesUintOptionInput_none(t *testing.T) {
	fnName := "test_uint_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUintOptionOutput_none(t *testing.T) {
	fnName := "test_uint_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestPrimitivesUint16Input_min(t *testing.T) {
	fnName := "test_uint16_input_min"
	if _, err := fixture.CallFunction(t, fnName, uint16(0)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUint16Output_min(t *testing.T) {
	fnName := "test_uint16_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint16(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(uint16); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesUint16Input_max(t *testing.T) {
	fnName := "test_uint16_input_max"
	if _, err := fixture.CallFunction(t, fnName, uint16(math.MaxUint16)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUint16Output_max(t *testing.T) {
	fnName := "test_uint16_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint16(math.MaxUint16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(uint16); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesUint16OptionInput_min(t *testing.T) {
	fnName := "test_uint16_option_input_min"
	b := uint16(0)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUint16OptionOutput_min(t *testing.T) {
	fnName := "test_uint16_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint16(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*uint16); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesUint16OptionInput_max(t *testing.T) {
	fnName := "test_uint16_option_input_max"
	b := uint16(math.MaxUint16)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUint16OptionOutput_max(t *testing.T) {
	fnName := "test_uint16_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint16(math.MaxUint16)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*uint16); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesUint16OptionInput_none(t *testing.T) {
	fnName := "test_uint16_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUint16OptionOutput_none(t *testing.T) {
	fnName := "test_uint16_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestPrimitivesUint64Input_min(t *testing.T) {
	fnName := "test_uint64_input_min"
	if _, err := fixture.CallFunction(t, fnName, uint64(0)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUint64Output_min(t *testing.T) {
	fnName := "test_uint64_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint64(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(uint64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesUint64Input_max(t *testing.T) {
	fnName := "test_uint64_input_max"
	if _, err := fixture.CallFunction(t, fnName, uint64(math.MaxUint64)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUint64Output_max(t *testing.T) {
	fnName := "test_uint64_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint64(math.MaxUint64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(uint64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesUint64OptionInput_min(t *testing.T) {
	fnName := "test_uint64_option_input_min"
	b := uint64(0)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUint64OptionOutput_min(t *testing.T) {
	fnName := "test_uint64_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint64(0)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*uint64); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesUint64OptionInput_max(t *testing.T) {
	fnName := "test_uint64_option_input_max"
	b := uint64(math.MaxUint64)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUint64OptionOutput_max(t *testing.T) {
	fnName := "test_uint64_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint64(math.MaxUint64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*uint64); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesUint64OptionInput_none(t *testing.T) {
	fnName := "test_uint64_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesUint64OptionOutput_none(t *testing.T) {
	fnName := "test_uint64_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestPrimitivesFloatInput_min(t *testing.T) {
	fnName := "test_float_input_min"
	if _, err := fixture.CallFunction(t, fnName, float32(math.SmallestNonzeroFloat32)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesFloatOutput_min(t *testing.T) {
	fnName := "test_float_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := float32(-3.4028235e+38)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(float32); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesFloatInput_max(t *testing.T) {
	fnName := "test_float_input_max"
	if _, err := fixture.CallFunction(t, fnName, float32(math.MaxFloat32)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesFloatOutput_max(t *testing.T) {
	fnName := "test_float_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := float32(math.MaxFloat32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(float32); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesFloatOptionInput_min(t *testing.T) {
	fnName := "test_float_option_input_min"
	b := float32(math.SmallestNonzeroFloat32)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesFloatOptionOutput_min(t *testing.T) {
	fnName := "test_float_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := float32(-3.4028235e+38)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*float32); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesFloatOptionInput_max(t *testing.T) {
	fnName := "test_float_option_input_max"
	b := float32(math.MaxFloat32)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesFloatOptionOutput_max(t *testing.T) {
	fnName := "test_float_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := float32(math.MaxFloat32)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*float32); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesFloatOptionInput_none(t *testing.T) {
	fnName := "test_float_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesFloatOptionOutput_none(t *testing.T) {
	fnName := "test_float_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}

func TestPrimitivesDoubleInput_min(t *testing.T) {
	fnName := "test_double_input_min"
	if _, err := fixture.CallFunction(t, fnName, float64(math.SmallestNonzeroFloat64)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesDoubleOutput_min(t *testing.T) {
	fnName := "test_double_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := -1.7976931348623157e+308
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(float64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesDoubleInput_max(t *testing.T) {
	fnName := "test_double_input_max"
	if _, err := fixture.CallFunction(t, fnName, float64(math.MaxFloat64)); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesDoubleOutput_max(t *testing.T) {
	fnName := "test_double_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := float64(math.MaxFloat64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(float64); !ok {
		t.Errorf("expected %T, got %T", expected, result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestPrimitivesDoubleOptionInput_min(t *testing.T) {
	fnName := "test_double_option_input_min"
	b := float64(math.SmallestNonzeroFloat64)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesDoubleOptionOutput_min(t *testing.T) {
	fnName := "test_double_option_output_min"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := -1.7976931348623157e+308
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*float64); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesDoubleOptionInput_max(t *testing.T) {
	fnName := "test_double_option_input_max"
	b := float64(math.MaxFloat64)

	if _, err := fixture.CallFunction(t, fnName, b); err != nil {
		t.Error(err)
	}
	if _, err := fixture.CallFunction(t, fnName, &b); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesDoubleOptionOutput_max(t *testing.T) {
	fnName := "test_double_option_output_max"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	expected := float64(math.MaxFloat64)
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*float64); !ok {
		t.Errorf("expected *%T, got %T", expected, result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestPrimitivesDoubleOptionInput_none(t *testing.T) {
	fnName := "test_double_option_input_none"
	if _, err := fixture.CallFunction(t, fnName, nil); err != nil {
		t.Error(err)
	}
}

func TestPrimitivesDoubleOptionOutput_none(t *testing.T) {
	fnName := "test_double_option_output_none"
	result, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Fatal(err)
	}

	if !utils.HasNil(result) {
		t.Error("expected a nil result")
	}
}
