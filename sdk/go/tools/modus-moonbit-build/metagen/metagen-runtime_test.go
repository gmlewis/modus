// -*- compile-command: "go test -run ^TestGenerateMetadata_Runtime$ ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package metagen

import (
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateMetadata_Runtime(t *testing.T) {
	meta := setupTestConfig(t, "testdata/runtime-testdata")
	removeExternalFuncsForComparison(t, meta)

	if got, want := meta.Plugin, "testdata"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@runtime-testdata"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	// if got, want := meta.SDK, "modus-sdk-mbt@0.16.5"; got != want {
	// 	t.Errorf("meta.SDK = %q, want %q", got, want)
	// }

	if diff := cmp.Diff(wantRuntimeFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantRuntimeFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	diffMetaTypes(t, wantRuntimeTypes, meta.Types)
	// if diff := cmp.Diff(wantRuntimeTypes, meta.Types); diff != "" {
	// 	t.Errorf("meta.Types mismatch (-want +got):\n%v", diff)
	// }

	// This call makes it easy to step through the code with a debugger:
	// LogToConsole(meta)
}

var wantRuntimeFnExports = metadata.FunctionMap{
	"add": {
		Name:       "add",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"call_test_time_option_input_none": {
		Name:    "call_test_time_option_input_none",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime?"}},
	},
	"call_test_time_option_input_some": {
		Name:    "call_test_time_option_input_some",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime?"}},
	},
	"echo1": {
		Name:       "echo1",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"echo2": {
		Name:       "echo2",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"echo3": {
		Name:       "echo3",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"echo4": {
		Name:       "echo4",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"encode_strings1": {
		Name:       "encode_strings1",
		Parameters: []*metadata.Parameter{{Name: "items", Type: "Array[String]?"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"encode_strings2": {
		Name:       "encode_strings2",
		Parameters: []*metadata.Parameter{{Name: "items", Type: "Array[String?]?"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"get_int_option_fixedarray1": {
		Name:    "get_int_option_fixedarray1",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"get_int_ptr_array2": {
		Name:    "get_int_ptr_array2",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"get_map_ptr_array2": {
		Name:    "get_map_ptr_array2",
		Results: []*metadata.Result{{Type: "FixedArray[Map[String, String]?]"}},
	},
	"get_map_array2": {
		Name:    "get_map_array2",
		Results: []*metadata.Result{{Type: "FixedArray[Map[String, String]]"}},
	},
	"get_option_int_fixedarray1": {
		Name:    "get_option_int_fixedarray1",
		Results: []*metadata.Result{{Type: "FixedArray[Int]?"}},
	},
	"get_option_int_fixedarray2": {
		Name:    "get_option_int_fixedarray2",
		Results: []*metadata.Result{{Type: "FixedArray[Int]?"}},
	},
	"get_option_string_fixedarray1": {
		Name:    "get_option_string_fixedarray1",
		Results: []*metadata.Result{{Type: "FixedArray[String]?"}},
	},
	"get_option_string_fixedarray2": {
		Name:    "get_option_string_fixedarray2",
		Results: []*metadata.Result{{Type: "FixedArray[String]?"}},
	},
	"get_string_option_array2": {
		Name:    "get_string_option_array2",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"get_string_option_fixedarray1": {
		Name:    "get_string_option_fixedarray1",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"host_echo1": {
		Name:       "host_echo1",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"host_echo2": {
		Name:       "host_echo2",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"host_echo3": {
		Name:       "host_echo3",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"host_echo4": {
		Name:       "host_echo4",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"test2d_array_input_string": {
		Name:       "test2d_array_input_string",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Array[String]]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test2d_array_input_string_empty": {
		Name:       "test2d_array_input_string_empty",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Array[String]]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test2d_array_input_string_inner_empty": {
		Name:       "test2d_array_input_string_inner_empty",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Array[String]]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test2d_array_input_string_inner_none": {
		Name:       "test2d_array_input_string_inner_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Array[String]?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test2d_array_input_string_none": {
		Name:       "test2d_array_input_string_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Array[String]]?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test2d_array_output_string": {
		Name:    "test2d_array_output_string",
		Results: []*metadata.Result{{Type: "Array[Array[String]]"}},
	},
	"test2d_array_output_string_empty": {
		Name:    "test2d_array_output_string_empty",
		Results: []*metadata.Result{{Type: "Array[Array[String]]"}},
	},
	"test2d_array_output_string_inner_empty": {
		Name:    "test2d_array_output_string_inner_empty",
		Results: []*metadata.Result{{Type: "Array[Array[String]]"}},
	},
	"test2d_array_output_string_inner_none": {
		Name:    "test2d_array_output_string_inner_none",
		Results: []*metadata.Result{{Type: "Array[Array[String]?]"}},
	},
	"test2d_array_output_string_none": {
		Name:    "test2d_array_output_string_none",
		Results: []*metadata.Result{{Type: "Array[Array[String]]?"}},
	},
	"test_array_input_bool_0": {
		Name:       "test_array_input_bool_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_1": {
		Name:       "test_array_input_bool_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_2": {
		Name:       "test_array_input_bool_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_3": {
		Name:       "test_array_input_bool_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_4": {
		Name:       "test_array_input_bool_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_option_0": {
		Name:       "test_array_input_bool_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_option_1_false": {
		Name:       "test_array_input_bool_option_1_false",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_option_1_none": {
		Name:       "test_array_input_bool_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_option_1_true": {
		Name:       "test_array_input_bool_option_1_true",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_option_2": {
		Name:       "test_array_input_bool_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_option_3": {
		Name:       "test_array_input_bool_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_bool_option_4": {
		Name:       "test_array_input_bool_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_byte_0": {
		Name:       "test_array_input_byte_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_byte_1": {
		Name:       "test_array_input_byte_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_byte_2": {
		Name:       "test_array_input_byte_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_byte_3": {
		Name:       "test_array_input_byte_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_byte_4": {
		Name:       "test_array_input_byte_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_byte_option_0": {
		Name:       "test_array_input_byte_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_byte_option_1": {
		Name:       "test_array_input_byte_option_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_byte_option_2": {
		Name:       "test_array_input_byte_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_byte_option_3": {
		Name:       "test_array_input_byte_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_byte_option_4": {
		Name:       "test_array_input_byte_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_char_empty": {
		Name:       "test_array_input_char_empty",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Char]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_char_option": {
		Name:       "test_array_input_char_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Char?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_double_empty": {
		Name:       "test_array_input_double_empty",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Double]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_double_option": {
		Name:       "test_array_input_double_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Double?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_float_empty": {
		Name:       "test_array_input_float_empty",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Float]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_float_option": {
		Name:       "test_array_input_float_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Float?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_int16_empty": {
		Name:       "test_array_input_int16_empty",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Int16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_int16_option": {
		Name:       "test_array_input_int16_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Int16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_int_empty": {
		Name:       "test_array_input_int_empty",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Int]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_int_option": {
		Name:       "test_array_input_int_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_string": {
		Name:       "test_array_input_string",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_string_empty": {
		Name:       "test_array_input_string_empty",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_string_none": {
		Name:       "test_array_input_string_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[String]?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_string_option": {
		Name:       "test_array_input_string_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[String?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_uint16_empty": {
		Name:       "test_array_input_uint16_empty",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[UInt16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_uint16_option": {
		Name:       "test_array_input_uint16_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[UInt16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_uint_empty": {
		Name:       "test_array_input_uint_empty",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[UInt]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_uint_option": {
		Name:       "test_array_input_uint_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[UInt?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_output_bool_0": {
		Name:    "test_array_output_bool_0",
		Results: []*metadata.Result{{Type: "Array[Bool]"}},
	},
	"test_array_output_bool_1": {
		Name:    "test_array_output_bool_1",
		Results: []*metadata.Result{{Type: "Array[Bool]"}},
	},
	"test_array_output_bool_2": {
		Name:    "test_array_output_bool_2",
		Results: []*metadata.Result{{Type: "Array[Bool]"}},
	},
	"test_array_output_bool_3": {
		Name:    "test_array_output_bool_3",
		Results: []*metadata.Result{{Type: "Array[Bool]"}},
	},
	"test_array_output_bool_4": {
		Name:    "test_array_output_bool_4",
		Results: []*metadata.Result{{Type: "Array[Bool]"}},
	},
	"test_array_output_bool_option_0": {
		Name:    "test_array_output_bool_option_0",
		Results: []*metadata.Result{{Type: "Array[Bool?]"}},
	},
	"test_array_output_bool_option_1_false": {
		Name:    "test_array_output_bool_option_1_false",
		Results: []*metadata.Result{{Type: "Array[Bool?]"}},
	},
	"test_array_output_bool_option_1_none": {
		Name:    "test_array_output_bool_option_1_none",
		Results: []*metadata.Result{{Type: "Array[Bool?]"}},
	},
	"test_array_output_bool_option_1_true": {
		Name:    "test_array_output_bool_option_1_true",
		Results: []*metadata.Result{{Type: "Array[Bool?]"}},
	},
	"test_array_output_bool_option_2": {
		Name:    "test_array_output_bool_option_2",
		Results: []*metadata.Result{{Type: "Array[Bool?]"}},
	},
	"test_array_output_bool_option_3": {
		Name:    "test_array_output_bool_option_3",
		Results: []*metadata.Result{{Type: "Array[Bool?]"}},
	},
	"test_array_output_bool_option_4": {
		Name:    "test_array_output_bool_option_4",
		Results: []*metadata.Result{{Type: "Array[Bool?]"}},
	},
	"test_array_output_byte_0": {
		Name:    "test_array_output_byte_0",
		Results: []*metadata.Result{{Type: "Array[Byte]"}},
	},
	"test_array_output_byte_1": {
		Name:    "test_array_output_byte_1",
		Results: []*metadata.Result{{Type: "Array[Byte]"}},
	},
	"test_array_output_byte_2": {
		Name:    "test_array_output_byte_2",
		Results: []*metadata.Result{{Type: "Array[Byte]"}},
	},
	"test_array_output_byte_3": {
		Name:    "test_array_output_byte_3",
		Results: []*metadata.Result{{Type: "Array[Byte]"}},
	},
	"test_array_output_byte_4": {
		Name:    "test_array_output_byte_4",
		Results: []*metadata.Result{{Type: "Array[Byte]"}},
	},
	"test_array_output_byte_option_0": {
		Name:    "test_array_output_byte_option_0",
		Results: []*metadata.Result{{Type: "Array[Byte?]"}},
	},
	"test_array_output_byte_option_1": {
		Name:    "test_array_output_byte_option_1",
		Results: []*metadata.Result{{Type: "Array[Byte?]"}},
	},
	"test_array_output_byte_option_2": {
		Name:    "test_array_output_byte_option_2",
		Results: []*metadata.Result{{Type: "Array[Byte?]"}},
	},
	"test_array_output_byte_option_3": {
		Name:    "test_array_output_byte_option_3",
		Results: []*metadata.Result{{Type: "Array[Byte?]"}},
	},
	"test_array_output_byte_option_4": {
		Name:    "test_array_output_byte_option_4",
		Results: []*metadata.Result{{Type: "Array[Byte?]"}},
	},
	"test_array_output_char_0": {
		Name:    "test_array_output_char_0",
		Results: []*metadata.Result{{Type: "Array[Char]"}},
	},
	"test_array_output_char_1": {
		Name:    "test_array_output_char_1",
		Results: []*metadata.Result{{Type: "Array[Char]"}},
	},
	"test_array_output_char_2": {
		Name:    "test_array_output_char_2",
		Results: []*metadata.Result{{Type: "Array[Char]"}},
	},
	"test_array_output_char_3": {
		Name:    "test_array_output_char_3",
		Results: []*metadata.Result{{Type: "Array[Char]"}},
	},
	"test_array_output_char_4": {
		Name:    "test_array_output_char_4",
		Results: []*metadata.Result{{Type: "Array[Char]"}},
	},
	"test_array_output_char_option": {
		Name:    "test_array_output_char_option",
		Results: []*metadata.Result{{Type: "Array[Char?]"}},
	},
	"test_array_output_char_option_0": {
		Name:    "test_array_output_char_option_0",
		Results: []*metadata.Result{{Type: "Array[Char?]"}},
	},
	"test_array_output_char_option_1_none": {
		Name:    "test_array_output_char_option_1_none",
		Results: []*metadata.Result{{Type: "Array[Char?]"}},
	},
	"test_array_output_char_option_1_some": {
		Name:    "test_array_output_char_option_1_some",
		Results: []*metadata.Result{{Type: "Array[Char?]"}},
	},
	"test_array_output_char_option_2": {
		Name:    "test_array_output_char_option_2",
		Results: []*metadata.Result{{Type: "Array[Char?]"}},
	},
	"test_array_output_char_option_3": {
		Name:    "test_array_output_char_option_3",
		Results: []*metadata.Result{{Type: "Array[Char?]"}},
	},
	"test_array_output_char_option_4": {
		Name:    "test_array_output_char_option_4",
		Results: []*metadata.Result{{Type: "Array[Char?]"}},
	},
	"test_array_output_double_0": {
		Name:    "test_array_output_double_0",
		Results: []*metadata.Result{{Type: "Array[Double]"}},
	},
	"test_array_output_double_1": {
		Name:    "test_array_output_double_1",
		Results: []*metadata.Result{{Type: "Array[Double]"}},
	},
	"test_array_output_double_2": {
		Name:    "test_array_output_double_2",
		Results: []*metadata.Result{{Type: "Array[Double]"}},
	},
	"test_array_output_double_3": {
		Name:    "test_array_output_double_3",
		Results: []*metadata.Result{{Type: "Array[Double]"}},
	},
	"test_array_output_double_4": {
		Name:    "test_array_output_double_4",
		Results: []*metadata.Result{{Type: "Array[Double]"}},
	},
	"test_array_output_double_option": {
		Name:    "test_array_output_double_option",
		Results: []*metadata.Result{{Type: "Array[Double?]"}},
	},
	"test_array_output_double_option_0": {
		Name:    "test_array_output_double_option_0",
		Results: []*metadata.Result{{Type: "Array[Double?]"}},
	},
	"test_array_output_double_option_1_none": {
		Name:    "test_array_output_double_option_1_none",
		Results: []*metadata.Result{{Type: "Array[Double?]"}},
	},
	"test_array_output_double_option_1_some": {
		Name:    "test_array_output_double_option_1_some",
		Results: []*metadata.Result{{Type: "Array[Double?]"}},
	},
	"test_array_output_double_option_2": {
		Name:    "test_array_output_double_option_2",
		Results: []*metadata.Result{{Type: "Array[Double?]"}},
	},
	"test_array_output_double_option_3": {
		Name:    "test_array_output_double_option_3",
		Results: []*metadata.Result{{Type: "Array[Double?]"}},
	},
	"test_array_output_double_option_4": {
		Name:    "test_array_output_double_option_4",
		Results: []*metadata.Result{{Type: "Array[Double?]"}},
	},
	"test_array_output_float_0": {
		Name:    "test_array_output_float_0",
		Results: []*metadata.Result{{Type: "Array[Float]"}},
	},
	"test_array_output_float_1": {
		Name:    "test_array_output_float_1",
		Results: []*metadata.Result{{Type: "Array[Float]"}},
	},
	"test_array_output_float_2": {
		Name:    "test_array_output_float_2",
		Results: []*metadata.Result{{Type: "Array[Float]"}},
	},
	"test_array_output_float_3": {
		Name:    "test_array_output_float_3",
		Results: []*metadata.Result{{Type: "Array[Float]"}},
	},
	"test_array_output_float_4": {
		Name:    "test_array_output_float_4",
		Results: []*metadata.Result{{Type: "Array[Float]"}},
	},
	"test_array_output_float_option": {
		Name:    "test_array_output_float_option",
		Results: []*metadata.Result{{Type: "Array[Float?]"}},
	},
	"test_array_output_float_option_0": {
		Name:    "test_array_output_float_option_0",
		Results: []*metadata.Result{{Type: "Array[Float?]"}},
	},
	"test_array_output_float_option_1_none": {
		Name:    "test_array_output_float_option_1_none",
		Results: []*metadata.Result{{Type: "Array[Float?]"}},
	},
	"test_array_output_float_option_1_some": {
		Name:    "test_array_output_float_option_1_some",
		Results: []*metadata.Result{{Type: "Array[Float?]"}},
	},
	"test_array_output_float_option_2": {
		Name:    "test_array_output_float_option_2",
		Results: []*metadata.Result{{Type: "Array[Float?]"}},
	},
	"test_array_output_float_option_3": {
		Name:    "test_array_output_float_option_3",
		Results: []*metadata.Result{{Type: "Array[Float?]"}},
	},
	"test_array_output_float_option_4": {
		Name:    "test_array_output_float_option_4",
		Results: []*metadata.Result{{Type: "Array[Float?]"}},
	},
	"test_array_output_int16_0": {
		Name:    "test_array_output_int16_0",
		Results: []*metadata.Result{{Type: "Array[Int16]"}},
	},
	"test_array_output_int16_1": {
		Name:    "test_array_output_int16_1",
		Results: []*metadata.Result{{Type: "Array[Int16]"}},
	},
	"test_array_output_int16_1_max": {
		Name:    "test_array_output_int16_1_max",
		Results: []*metadata.Result{{Type: "Array[Int16]"}},
	},
	"test_array_output_int16_1_min": {
		Name:    "test_array_output_int16_1_min",
		Results: []*metadata.Result{{Type: "Array[Int16]"}},
	},
	"test_array_output_int16_2": {
		Name:    "test_array_output_int16_2",
		Results: []*metadata.Result{{Type: "Array[Int16]"}},
	},
	"test_array_output_int16_3": {
		Name:    "test_array_output_int16_3",
		Results: []*metadata.Result{{Type: "Array[Int16]"}},
	},
	"test_array_output_int16_4": {
		Name:    "test_array_output_int16_4",
		Results: []*metadata.Result{{Type: "Array[Int16]"}},
	},
	"test_array_output_int16_option": {
		Name:    "test_array_output_int16_option",
		Results: []*metadata.Result{{Type: "Array[Int16?]"}},
	},
	"test_array_output_int16_option_0": {
		Name:    "test_array_output_int16_option_0",
		Results: []*metadata.Result{{Type: "Array[Int16?]"}},
	},
	"test_array_output_int16_option_1_max": {
		Name:    "test_array_output_int16_option_1_max",
		Results: []*metadata.Result{{Type: "Array[Int16?]"}},
	},
	"test_array_output_int16_option_1_min": {
		Name:    "test_array_output_int16_option_1_min",
		Results: []*metadata.Result{{Type: "Array[Int16?]"}},
	},
	"test_array_output_int16_option_1_none": {
		Name:    "test_array_output_int16_option_1_none",
		Results: []*metadata.Result{{Type: "Array[Int16?]"}},
	},
	"test_array_output_int16_option_2": {
		Name:    "test_array_output_int16_option_2",
		Results: []*metadata.Result{{Type: "Array[Int16?]"}},
	},
	"test_array_output_int16_option_3": {
		Name:    "test_array_output_int16_option_3",
		Results: []*metadata.Result{{Type: "Array[Int16?]"}},
	},
	"test_array_output_int16_option_4": {
		Name:    "test_array_output_int16_option_4",
		Results: []*metadata.Result{{Type: "Array[Int16?]"}},
	},
	"test_array_output_int64_0": {
		Name:    "test_array_output_int64_0",
		Results: []*metadata.Result{{Type: "Array[Int64]"}},
	},
	"test_array_output_int64_1": {
		Name:    "test_array_output_int64_1",
		Results: []*metadata.Result{{Type: "Array[Int64]"}},
	},
	"test_array_output_int64_1_max": {
		Name:    "test_array_output_int64_1_max",
		Results: []*metadata.Result{{Type: "Array[Int64]"}},
	},
	"test_array_output_int64_1_min": {
		Name:    "test_array_output_int64_1_min",
		Results: []*metadata.Result{{Type: "Array[Int64]"}},
	},
	"test_array_output_int64_2": {
		Name:    "test_array_output_int64_2",
		Results: []*metadata.Result{{Type: "Array[Int64]"}},
	},
	"test_array_output_int64_3": {
		Name:    "test_array_output_int64_3",
		Results: []*metadata.Result{{Type: "Array[Int64]"}},
	},
	"test_array_output_int64_4": {
		Name:    "test_array_output_int64_4",
		Results: []*metadata.Result{{Type: "Array[Int64]"}},
	},
	"test_array_output_int64_option_0": {
		Name:    "test_array_output_int64_option_0",
		Results: []*metadata.Result{{Type: "Array[Int64?]"}},
	},
	"test_array_output_int64_option_1_max": {
		Name:    "test_array_output_int64_option_1_max",
		Results: []*metadata.Result{{Type: "Array[Int64?]"}},
	},
	"test_array_output_int64_option_1_min": {
		Name:    "test_array_output_int64_option_1_min",
		Results: []*metadata.Result{{Type: "Array[Int64?]"}},
	},
	"test_array_output_int64_option_1_none": {
		Name:    "test_array_output_int64_option_1_none",
		Results: []*metadata.Result{{Type: "Array[Int64?]"}},
	},
	"test_array_output_int64_option_2": {
		Name:    "test_array_output_int64_option_2",
		Results: []*metadata.Result{{Type: "Array[Int64?]"}},
	},
	"test_array_output_int64_option_3": {
		Name:    "test_array_output_int64_option_3",
		Results: []*metadata.Result{{Type: "Array[Int64?]"}},
	},
	"test_array_output_int64_option_4": {
		Name:    "test_array_output_int64_option_4",
		Results: []*metadata.Result{{Type: "Array[Int64?]"}},
	},
	"test_array_output_int_0": {
		Name:    "test_array_output_int_0",
		Results: []*metadata.Result{{Type: "Array[Int]"}},
	},
	"test_array_output_int_1": {
		Name:    "test_array_output_int_1",
		Results: []*metadata.Result{{Type: "Array[Int]"}},
	},
	"test_array_output_int_1_max": {
		Name:    "test_array_output_int_1_max",
		Results: []*metadata.Result{{Type: "Array[Int]"}},
	},
	"test_array_output_int_1_min": {
		Name:    "test_array_output_int_1_min",
		Results: []*metadata.Result{{Type: "Array[Int]"}},
	},
	"test_array_output_int_2": {
		Name:    "test_array_output_int_2",
		Results: []*metadata.Result{{Type: "Array[Int]"}},
	},
	"test_array_output_int_3": {
		Name:    "test_array_output_int_3",
		Results: []*metadata.Result{{Type: "Array[Int]"}},
	},
	"test_array_output_int_4": {
		Name:    "test_array_output_int_4",
		Results: []*metadata.Result{{Type: "Array[Int]"}},
	},
	"test_array_output_int_option": {
		Name:    "test_array_output_int_option",
		Results: []*metadata.Result{{Type: "Array[Int?]"}},
	},
	"test_array_output_int_option_0": {
		Name:    "test_array_output_int_option_0",
		Results: []*metadata.Result{{Type: "Array[Int?]"}},
	},
	"test_array_output_int_option_1_max": {
		Name:    "test_array_output_int_option_1_max",
		Results: []*metadata.Result{{Type: "Array[Int?]"}},
	},
	"test_array_output_int_option_1_min": {
		Name:    "test_array_output_int_option_1_min",
		Results: []*metadata.Result{{Type: "Array[Int?]"}},
	},
	"test_array_output_int_option_1_none": {
		Name:    "test_array_output_int_option_1_none",
		Results: []*metadata.Result{{Type: "Array[Int?]"}},
	},
	"test_array_output_int_option_2": {
		Name:    "test_array_output_int_option_2",
		Results: []*metadata.Result{{Type: "Array[Int?]"}},
	},
	"test_array_output_int_option_3": {
		Name:    "test_array_output_int_option_3",
		Results: []*metadata.Result{{Type: "Array[Int?]"}},
	},
	"test_array_output_int_option_4": {
		Name:    "test_array_output_int_option_4",
		Results: []*metadata.Result{{Type: "Array[Int?]"}},
	},
	"test_array_output_string": {
		Name:    "test_array_output_string",
		Results: []*metadata.Result{{Type: "Array[String]"}},
	},
	"test_array_output_string_empty": {
		Name:    "test_array_output_string_empty",
		Results: []*metadata.Result{{Type: "Array[String]"}},
	},
	"test_array_output_string_none": {
		Name:    "test_array_output_string_none",
		Results: []*metadata.Result{{Type: "Array[String]?"}},
	},
	"test_array_output_string_option": {
		Name:    "test_array_output_string_option",
		Results: []*metadata.Result{{Type: "Array[String?]"}},
	},
	"test_array_output_uint16_0": {
		Name:    "test_array_output_uint16_0",
		Results: []*metadata.Result{{Type: "Array[UInt16]"}},
	},
	"test_array_output_uint16_1": {
		Name:    "test_array_output_uint16_1",
		Results: []*metadata.Result{{Type: "Array[UInt16]"}},
	},
	"test_array_output_uint16_1_max": {
		Name:    "test_array_output_uint16_1_max",
		Results: []*metadata.Result{{Type: "Array[UInt16]"}},
	},
	"test_array_output_uint16_1_min": {
		Name:    "test_array_output_uint16_1_min",
		Results: []*metadata.Result{{Type: "Array[UInt16]"}},
	},
	"test_array_output_uint16_2": {
		Name:    "test_array_output_uint16_2",
		Results: []*metadata.Result{{Type: "Array[UInt16]"}},
	},
	"test_array_output_uint16_3": {
		Name:    "test_array_output_uint16_3",
		Results: []*metadata.Result{{Type: "Array[UInt16]"}},
	},
	"test_array_output_uint16_4": {
		Name:    "test_array_output_uint16_4",
		Results: []*metadata.Result{{Type: "Array[UInt16]"}},
	},
	"test_array_output_uint16_option": {
		Name:    "test_array_output_uint16_option",
		Results: []*metadata.Result{{Type: "Array[UInt16?]"}},
	},
	"test_array_output_uint16_option_0": {
		Name:    "test_array_output_uint16_option_0",
		Results: []*metadata.Result{{Type: "Array[UInt16?]"}},
	},
	"test_array_output_uint16_option_1_max": {
		Name:    "test_array_output_uint16_option_1_max",
		Results: []*metadata.Result{{Type: "Array[UInt16?]"}},
	},
	"test_array_output_uint16_option_1_min": {
		Name:    "test_array_output_uint16_option_1_min",
		Results: []*metadata.Result{{Type: "Array[UInt16?]"}},
	},
	"test_array_output_uint16_option_1_none": {
		Name:    "test_array_output_uint16_option_1_none",
		Results: []*metadata.Result{{Type: "Array[UInt16?]"}},
	},
	"test_array_output_uint16_option_2": {
		Name:    "test_array_output_uint16_option_2",
		Results: []*metadata.Result{{Type: "Array[UInt16?]"}},
	},
	"test_array_output_uint16_option_3": {
		Name:    "test_array_output_uint16_option_3",
		Results: []*metadata.Result{{Type: "Array[UInt16?]"}},
	},
	"test_array_output_uint16_option_4": {
		Name:    "test_array_output_uint16_option_4",
		Results: []*metadata.Result{{Type: "Array[UInt16?]"}},
	},
	"test_array_output_uint64_0": {
		Name:    "test_array_output_uint64_0",
		Results: []*metadata.Result{{Type: "Array[UInt64]"}},
	},
	"test_array_output_uint64_1": {
		Name:    "test_array_output_uint64_1",
		Results: []*metadata.Result{{Type: "Array[UInt64]"}},
	},
	"test_array_output_uint64_1_max": {
		Name:    "test_array_output_uint64_1_max",
		Results: []*metadata.Result{{Type: "Array[UInt64]"}},
	},
	"test_array_output_uint64_1_min": {
		Name:    "test_array_output_uint64_1_min",
		Results: []*metadata.Result{{Type: "Array[UInt64]"}},
	},
	"test_array_output_uint64_2": {
		Name:    "test_array_output_uint64_2",
		Results: []*metadata.Result{{Type: "Array[UInt64]"}},
	},
	"test_array_output_uint64_3": {
		Name:    "test_array_output_uint64_3",
		Results: []*metadata.Result{{Type: "Array[UInt64]"}},
	},
	"test_array_output_uint64_4": {
		Name:    "test_array_output_uint64_4",
		Results: []*metadata.Result{{Type: "Array[UInt64]"}},
	},
	"test_array_output_uint64_option_0": {
		Name:    "test_array_output_uint64_option_0",
		Results: []*metadata.Result{{Type: "Array[UInt64?]"}},
	},
	"test_array_output_uint64_option_1_max": {
		Name:    "test_array_output_uint64_option_1_max",
		Results: []*metadata.Result{{Type: "Array[UInt64?]"}},
	},
	"test_array_output_uint64_option_1_min": {
		Name:    "test_array_output_uint64_option_1_min",
		Results: []*metadata.Result{{Type: "Array[UInt64?]"}},
	},
	"test_array_output_uint64_option_1_none": {
		Name:    "test_array_output_uint64_option_1_none",
		Results: []*metadata.Result{{Type: "Array[UInt64?]"}},
	},
	"test_array_output_uint64_option_2": {
		Name:    "test_array_output_uint64_option_2",
		Results: []*metadata.Result{{Type: "Array[UInt64?]"}},
	},
	"test_array_output_uint64_option_3": {
		Name:    "test_array_output_uint64_option_3",
		Results: []*metadata.Result{{Type: "Array[UInt64?]"}},
	},
	"test_array_output_uint64_option_4": {
		Name:    "test_array_output_uint64_option_4",
		Results: []*metadata.Result{{Type: "Array[UInt64?]"}},
	},
	"test_array_output_uint_0": {
		Name:    "test_array_output_uint_0",
		Results: []*metadata.Result{{Type: "Array[UInt]"}},
	},
	"test_array_output_uint_1": {
		Name:    "test_array_output_uint_1",
		Results: []*metadata.Result{{Type: "Array[UInt]"}},
	},
	"test_array_output_uint_1_max": {
		Name:    "test_array_output_uint_1_max",
		Results: []*metadata.Result{{Type: "Array[UInt]"}},
	},
	"test_array_output_uint_1_min": {
		Name:    "test_array_output_uint_1_min",
		Results: []*metadata.Result{{Type: "Array[UInt]"}},
	},
	"test_array_output_uint_2": {
		Name:    "test_array_output_uint_2",
		Results: []*metadata.Result{{Type: "Array[UInt]"}},
	},
	"test_array_output_uint_3": {
		Name:    "test_array_output_uint_3",
		Results: []*metadata.Result{{Type: "Array[UInt]"}},
	},
	"test_array_output_uint_4": {
		Name:    "test_array_output_uint_4",
		Results: []*metadata.Result{{Type: "Array[UInt]"}},
	},
	"test_array_output_uint_option": {
		Name:    "test_array_output_uint_option",
		Results: []*metadata.Result{{Type: "Array[UInt?]"}},
	},
	"test_array_output_uint_option_0": {
		Name:    "test_array_output_uint_option_0",
		Results: []*metadata.Result{{Type: "Array[UInt?]"}},
	},
	"test_array_output_uint_option_1_max": {
		Name:    "test_array_output_uint_option_1_max",
		Results: []*metadata.Result{{Type: "Array[UInt?]"}},
	},
	"test_array_output_uint_option_1_min": {
		Name:    "test_array_output_uint_option_1_min",
		Results: []*metadata.Result{{Type: "Array[UInt?]"}},
	},
	"test_array_output_uint_option_1_none": {
		Name:    "test_array_output_uint_option_1_none",
		Results: []*metadata.Result{{Type: "Array[UInt?]"}},
	},
	"test_array_output_uint_option_2": {
		Name:    "test_array_output_uint_option_2",
		Results: []*metadata.Result{{Type: "Array[UInt?]"}},
	},
	"test_array_output_uint_option_3": {
		Name:    "test_array_output_uint_option_3",
		Results: []*metadata.Result{{Type: "Array[UInt?]"}},
	},
	"test_array_output_uint_option_4": {
		Name:    "test_array_output_uint_option_4",
		Results: []*metadata.Result{{Type: "Array[UInt?]"}},
	},
	"test_bool_input_false": {
		Name:       "test_bool_input_false",
		Parameters: []*metadata.Parameter{{Name: "b", Type: "Bool"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_bool_input_true": {
		Name:       "test_bool_input_true",
		Parameters: []*metadata.Parameter{{Name: "b", Type: "Bool"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_bool_option_input_false": {
		Name:       "test_bool_option_input_false",
		Parameters: []*metadata.Parameter{{Name: "b", Type: "Bool?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_bool_option_input_none": {
		Name:       "test_bool_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "b", Type: "Bool?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_bool_option_input_true": {
		Name:       "test_bool_option_input_true",
		Parameters: []*metadata.Parameter{{Name: "b", Type: "Bool?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_bool_option_output_false": {
		Name:    "test_bool_option_output_false",
		Results: []*metadata.Result{{Type: "Bool?"}},
	},
	"test_bool_option_output_none": {
		Name:    "test_bool_option_output_none",
		Results: []*metadata.Result{{Type: "Bool?"}},
	},
	"test_bool_option_output_true": {
		Name:    "test_bool_option_output_true",
		Results: []*metadata.Result{{Type: "Bool?"}},
	},
	"test_bool_output_false": {Name: "test_bool_output_false", Results: []*metadata.Result{{Type: "Bool"}}},
	"test_bool_output_true":  {Name: "test_bool_output_true", Results: []*metadata.Result{{Type: "Bool"}}},
	"test_byte_input_max": {
		Name:       "test_byte_input_max",
		Parameters: []*metadata.Parameter{{Name: "b", Type: "Byte"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_byte_input_min": {
		Name:       "test_byte_input_min",
		Parameters: []*metadata.Parameter{{Name: "b", Type: "Byte"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_byte_option_input_max": {
		Name:       "test_byte_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "b", Type: "Byte?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_byte_option_input_min": {
		Name:       "test_byte_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "b", Type: "Byte?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_byte_option_input_none": {
		Name:       "test_byte_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "b", Type: "Byte?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_byte_option_output_max": {
		Name:    "test_byte_option_output_max",
		Results: []*metadata.Result{{Type: "Byte?"}},
	},
	"test_byte_option_output_min": {
		Name:    "test_byte_option_output_min",
		Results: []*metadata.Result{{Type: "Byte?"}},
	},
	"test_byte_option_output_none": {
		Name:    "test_byte_option_output_none",
		Results: []*metadata.Result{{Type: "Byte?"}},
	},
	"test_byte_output_max": {Name: "test_byte_output_max", Results: []*metadata.Result{{Type: "Byte"}}},
	"test_byte_output_min": {Name: "test_byte_output_min", Results: []*metadata.Result{{Type: "Byte"}}},
	"test_char_input_max": {
		Name:       "test_char_input_max",
		Parameters: []*metadata.Parameter{{Name: "c", Type: "Char"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_char_input_min": {
		Name:       "test_char_input_min",
		Parameters: []*metadata.Parameter{{Name: "c", Type: "Char"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_char_option_input_max": {
		Name:       "test_char_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "c", Type: "Char?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_char_option_input_min": {
		Name:       "test_char_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "c", Type: "Char?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_char_option_input_none": {
		Name:       "test_char_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "c", Type: "Char?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_char_option_output_max": {
		Name:    "test_char_option_output_max",
		Results: []*metadata.Result{{Type: "Char?"}},
	},
	"test_char_option_output_min": {
		Name:    "test_char_option_output_min",
		Results: []*metadata.Result{{Type: "Char?"}},
	},
	"test_char_option_output_none": {
		Name:    "test_char_option_output_none",
		Results: []*metadata.Result{{Type: "Char?"}},
	},
	"test_char_output_max": {Name: "test_char_output_max", Results: []*metadata.Result{{Type: "Char"}}},
	"test_char_output_min": {Name: "test_char_output_min", Results: []*metadata.Result{{Type: "Char"}}},
	"test_double_input_max": {
		Name:       "test_double_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Double"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_double_input_min": {
		Name:       "test_double_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Double"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_double_option_input_max": {
		Name:       "test_double_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Double?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_double_option_input_min": {
		Name:       "test_double_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Double?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_double_option_input_none": {
		Name:       "test_double_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Double?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_double_option_output_max": {
		Name:    "test_double_option_output_max",
		Results: []*metadata.Result{{Type: "Double?"}},
	},
	"test_double_option_output_min": {
		Name:    "test_double_option_output_min",
		Results: []*metadata.Result{{Type: "Double?"}},
	},
	"test_double_option_output_none": {
		Name:    "test_double_option_output_none",
		Results: []*metadata.Result{{Type: "Double?"}},
	},
	"test_double_output_max": {Name: "test_double_output_max", Results: []*metadata.Result{{Type: "Double"}}},
	"test_double_output_min": {Name: "test_double_output_min", Results: []*metadata.Result{{Type: "Double"}}},
	"test_duration_input": {
		Name:       "test_duration_input",
		Parameters: []*metadata.Parameter{{Name: "d", Type: "@time.Duration"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_duration_option_input": {
		Name:       "test_duration_option_input",
		Parameters: []*metadata.Parameter{{Name: "d", Type: "@time.Duration?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_duration_option_input_none": {
		Name:       "test_duration_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "d", Type: "@time.Duration?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_duration_option_input_none_style2": {
		Name:       "test_duration_option_input_none_style2",
		Parameters: []*metadata.Parameter{{Name: "d?", Type: "@time.Duration"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_duration_option_input_style2": {
		Name:       "test_duration_option_input_style2",
		Parameters: []*metadata.Parameter{{Name: "d?", Type: "@time.Duration"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_duration_option_output": {
		Name:    "test_duration_option_output",
		Results: []*metadata.Result{{Type: "@time.Duration?"}},
	},
	"test_duration_option_output_none": {
		Name:    "test_duration_option_output_none",
		Results: []*metadata.Result{{Type: "@time.Duration?"}},
	},
	"test_duration_output": {
		Name:    "test_duration_output",
		Results: []*metadata.Result{{Type: "@time.Duration"}},
	},
	"test_fixedarray_input0_byte": {
		Name:       "test_fixedarray_input0_byte",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input0_int_option": {
		Name:       "test_fixedarray_input0_int_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input0_string": {
		Name:       "test_fixedarray_input0_string",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input0_string_option": {
		Name:       "test_fixedarray_input0_string_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input1_byte": {
		Name:       "test_fixedarray_input1_byte",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input1_int_option": {
		Name:       "test_fixedarray_input1_int_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input1_string": {
		Name:       "test_fixedarray_input1_string",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input1_string_option": {
		Name:       "test_fixedarray_input1_string_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input2_byte": {
		Name:       "test_fixedarray_input2_byte",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input2_int_option": {
		Name:       "test_fixedarray_input2_int_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input2_map": {
		Name:       "test_fixedarray_input2_map",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Map[String, String]]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input2_map_option": {
		Name:       "test_fixedarray_input2_map_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Map[String, String]?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input2_string": {
		Name:       "test_fixedarray_input2_string",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input2_string_option": {
		Name:       "test_fixedarray_input2_string_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input2_struct": {
		Name:       "test_fixedarray_input2_struct",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[TestStruct2]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input2_struct_option": {
		Name:       "test_fixedarray_input2_struct_option",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[TestStruct2?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_0": {
		Name:       "test_fixedarray_input_bool_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_1": {
		Name:       "test_fixedarray_input_bool_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_2": {
		Name:       "test_fixedarray_input_bool_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_3": {
		Name:       "test_fixedarray_input_bool_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_4": {
		Name:       "test_fixedarray_input_bool_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_option_0": {
		Name:       "test_fixedarray_input_bool_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_option_1_false": {
		Name:       "test_fixedarray_input_bool_option_1_false",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_option_1_none": {
		Name:       "test_fixedarray_input_bool_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_option_1_true": {
		Name:       "test_fixedarray_input_bool_option_1_true",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_option_2": {
		Name:       "test_fixedarray_input_bool_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_option_3": {
		Name:       "test_fixedarray_input_bool_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_bool_option_4": {
		Name:       "test_fixedarray_input_bool_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Bool?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_byte_0": {
		Name:       "test_fixedarray_input_byte_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_byte_1": {
		Name:       "test_fixedarray_input_byte_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_byte_2": {
		Name:       "test_fixedarray_input_byte_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_byte_3": {
		Name:       "test_fixedarray_input_byte_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_byte_4": {
		Name:       "test_fixedarray_input_byte_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_byte_option_0": {
		Name:       "test_fixedarray_input_byte_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_byte_option_1": {
		Name:       "test_fixedarray_input_byte_option_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_byte_option_2": {
		Name:       "test_fixedarray_input_byte_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_byte_option_3": {
		Name:       "test_fixedarray_input_byte_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_byte_option_4": {
		Name:       "test_fixedarray_input_byte_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Byte?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_0": {
		Name:       "test_fixedarray_input_char_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_1": {
		Name:       "test_fixedarray_input_char_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_2": {
		Name:       "test_fixedarray_input_char_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_3": {
		Name:       "test_fixedarray_input_char_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_4": {
		Name:       "test_fixedarray_input_char_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_option_0": {
		Name:       "test_fixedarray_input_char_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_option_1_none": {
		Name:       "test_fixedarray_input_char_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_option_1_some": {
		Name:       "test_fixedarray_input_char_option_1_some",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_option_2": {
		Name:       "test_fixedarray_input_char_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_option_3": {
		Name:       "test_fixedarray_input_char_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_char_option_4": {
		Name:       "test_fixedarray_input_char_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Char?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_0": {
		Name:       "test_fixedarray_input_double_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_1": {
		Name:       "test_fixedarray_input_double_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_2": {
		Name:       "test_fixedarray_input_double_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_3": {
		Name:       "test_fixedarray_input_double_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_4": {
		Name:       "test_fixedarray_input_double_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_option_0": {
		Name:       "test_fixedarray_input_double_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_option_1_none": {
		Name:       "test_fixedarray_input_double_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_option_1_some": {
		Name:       "test_fixedarray_input_double_option_1_some",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_option_2": {
		Name:       "test_fixedarray_input_double_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_option_3": {
		Name:       "test_fixedarray_input_double_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_double_option_4": {
		Name:       "test_fixedarray_input_double_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Double?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_0": {
		Name:       "test_fixedarray_input_float_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_1": {
		Name:       "test_fixedarray_input_float_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_2": {
		Name:       "test_fixedarray_input_float_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_3": {
		Name:       "test_fixedarray_input_float_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_4": {
		Name:       "test_fixedarray_input_float_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_option_0": {
		Name:       "test_fixedarray_input_float_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_option_1_none": {
		Name:       "test_fixedarray_input_float_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_option_1_some": {
		Name:       "test_fixedarray_input_float_option_1_some",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_option_2": {
		Name:       "test_fixedarray_input_float_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_option_3": {
		Name:       "test_fixedarray_input_float_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_float_option_4": {
		Name:       "test_fixedarray_input_float_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Float?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_0": {
		Name:       "test_fixedarray_input_int16_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_1": {
		Name:       "test_fixedarray_input_int16_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_1_max": {
		Name:       "test_fixedarray_input_int16_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_1_min": {
		Name:       "test_fixedarray_input_int16_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_2": {
		Name:       "test_fixedarray_input_int16_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_3": {
		Name:       "test_fixedarray_input_int16_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_4": {
		Name:       "test_fixedarray_input_int16_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_option_0": {
		Name:       "test_fixedarray_input_int16_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_option_1_max": {
		Name:       "test_fixedarray_input_int16_option_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_option_1_min": {
		Name:       "test_fixedarray_input_int16_option_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_option_1_none": {
		Name:       "test_fixedarray_input_int16_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_option_2": {
		Name:       "test_fixedarray_input_int16_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_option_3": {
		Name:       "test_fixedarray_input_int16_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int16_option_4": {
		Name:       "test_fixedarray_input_int16_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_0": {
		Name:       "test_fixedarray_input_int64_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_1": {
		Name:       "test_fixedarray_input_int64_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_1_max": {
		Name:       "test_fixedarray_input_int64_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_1_min": {
		Name:       "test_fixedarray_input_int64_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_2": {
		Name:       "test_fixedarray_input_int64_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_3": {
		Name:       "test_fixedarray_input_int64_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_4": {
		Name:       "test_fixedarray_input_int64_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_option_0": {
		Name:       "test_fixedarray_input_int64_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_option_1_max": {
		Name:       "test_fixedarray_input_int64_option_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_option_1_min": {
		Name:       "test_fixedarray_input_int64_option_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_option_1_none": {
		Name:       "test_fixedarray_input_int64_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_option_2": {
		Name:       "test_fixedarray_input_int64_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_option_3": {
		Name:       "test_fixedarray_input_int64_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int64_option_4": {
		Name:       "test_fixedarray_input_int64_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_0": {
		Name:       "test_fixedarray_input_int_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_1": {
		Name:       "test_fixedarray_input_int_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_1_max": {
		Name:       "test_fixedarray_input_int_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_1_min": {
		Name:       "test_fixedarray_input_int_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_2": {
		Name:       "test_fixedarray_input_int_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_3": {
		Name:       "test_fixedarray_input_int_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_4": {
		Name:       "test_fixedarray_input_int_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_option_0": {
		Name:       "test_fixedarray_input_int_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_option_1_max": {
		Name:       "test_fixedarray_input_int_option_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_option_1_min": {
		Name:       "test_fixedarray_input_int_option_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_option_1_none": {
		Name:       "test_fixedarray_input_int_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_option_2": {
		Name:       "test_fixedarray_input_int_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_option_3": {
		Name:       "test_fixedarray_input_int_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_int_option_4": {
		Name:       "test_fixedarray_input_int_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_0": {
		Name:       "test_fixedarray_input_string_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_1": {
		Name:       "test_fixedarray_input_string_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_2": {
		Name:       "test_fixedarray_input_string_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_3": {
		Name:       "test_fixedarray_input_string_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_4": {
		Name:       "test_fixedarray_input_string_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_option_0": {
		Name:       "test_fixedarray_input_string_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_option_1_none": {
		Name:       "test_fixedarray_input_string_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_option_1_some": {
		Name:       "test_fixedarray_input_string_option_1_some",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_option_2": {
		Name:       "test_fixedarray_input_string_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_option_3": {
		Name:       "test_fixedarray_input_string_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_string_option_4": {
		Name:       "test_fixedarray_input_string_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_0": {
		Name:       "test_fixedarray_input_uint16_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_1": {
		Name:       "test_fixedarray_input_uint16_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_1_max": {
		Name:       "test_fixedarray_input_uint16_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_1_min": {
		Name:       "test_fixedarray_input_uint16_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_2": {
		Name:       "test_fixedarray_input_uint16_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_3": {
		Name:       "test_fixedarray_input_uint16_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_4": {
		Name:       "test_fixedarray_input_uint16_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_option_0": {
		Name:       "test_fixedarray_input_uint16_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_option_1_max": {
		Name:       "test_fixedarray_input_uint16_option_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_option_1_min": {
		Name:       "test_fixedarray_input_uint16_option_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_option_1_none": {
		Name:       "test_fixedarray_input_uint16_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_option_2": {
		Name:       "test_fixedarray_input_uint16_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_option_3": {
		Name:       "test_fixedarray_input_uint16_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint16_option_4": {
		Name:       "test_fixedarray_input_uint16_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt16?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_0": {
		Name:       "test_fixedarray_input_uint64_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_1": {
		Name:       "test_fixedarray_input_uint64_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_1_max": {
		Name:       "test_fixedarray_input_uint64_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_1_min": {
		Name:       "test_fixedarray_input_uint64_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_2": {
		Name:       "test_fixedarray_input_uint64_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_3": {
		Name:       "test_fixedarray_input_uint64_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_4": {
		Name:       "test_fixedarray_input_uint64_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_option_0": {
		Name:       "test_fixedarray_input_uint64_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_option_1_max": {
		Name:       "test_fixedarray_input_uint64_option_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_option_1_min": {
		Name:       "test_fixedarray_input_uint64_option_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_option_1_none": {
		Name:       "test_fixedarray_input_uint64_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_option_2": {
		Name:       "test_fixedarray_input_uint64_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_option_3": {
		Name:       "test_fixedarray_input_uint64_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint64_option_4": {
		Name:       "test_fixedarray_input_uint64_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt64?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_0": {
		Name:       "test_fixedarray_input_uint_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_1": {
		Name:       "test_fixedarray_input_uint_1",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_1_max": {
		Name:       "test_fixedarray_input_uint_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_1_min": {
		Name:       "test_fixedarray_input_uint_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_2": {
		Name:       "test_fixedarray_input_uint_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_3": {
		Name:       "test_fixedarray_input_uint_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_4": {
		Name:       "test_fixedarray_input_uint_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_option_0": {
		Name:       "test_fixedarray_input_uint_option_0",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_option_1_max": {
		Name:       "test_fixedarray_input_uint_option_1_max",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_option_1_min": {
		Name:       "test_fixedarray_input_uint_option_1_min",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_option_1_none": {
		Name:       "test_fixedarray_input_uint_option_1_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_option_2": {
		Name:       "test_fixedarray_input_uint_option_2",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_option_3": {
		Name:       "test_fixedarray_input_uint_option_3",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_input_uint_option_4": {
		Name:       "test_fixedarray_input_uint_option_4",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[UInt?]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_fixedarray_output0_byte": {
		Name:    "test_fixedarray_output0_byte",
		Results: []*metadata.Result{{Type: "FixedArray[Byte]"}},
	},
	"test_fixedarray_output0_int_option": {
		Name:    "test_fixedarray_output0_int_option",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"test_fixedarray_output0_string": {
		Name:    "test_fixedarray_output0_string",
		Results: []*metadata.Result{{Type: "FixedArray[String]"}},
	},
	"test_fixedarray_output0_string_option": {
		Name:    "test_fixedarray_output0_string_option",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"test_fixedarray_output1_byte": {
		Name:    "test_fixedarray_output1_byte",
		Results: []*metadata.Result{{Type: "FixedArray[Byte]"}},
	},
	"test_fixedarray_output1_int_option": {
		Name:    "test_fixedarray_output1_int_option",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"test_fixedarray_output1_string": {
		Name:    "test_fixedarray_output1_string",
		Results: []*metadata.Result{{Type: "FixedArray[String]"}},
	},
	"test_fixedarray_output1_string_option": {
		Name:    "test_fixedarray_output1_string_option",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"test_fixedarray_output2_byte": {
		Name:    "test_fixedarray_output2_byte",
		Results: []*metadata.Result{{Type: "FixedArray[Byte]"}},
	},
	"test_fixedarray_output2_int_option": {
		Name:    "test_fixedarray_output2_int_option",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"test_fixedarray_output2_map": {
		Name:    "test_fixedarray_output2_map",
		Results: []*metadata.Result{{Type: "FixedArray[Map[String, String]]"}},
	},
	"test_fixedarray_output2_map_option": {
		Name:    "test_fixedarray_output2_map_option",
		Results: []*metadata.Result{{Type: "FixedArray[Map[String, String]?]"}},
	},
	"test_fixedarray_output2_string": {
		Name:    "test_fixedarray_output2_string",
		Results: []*metadata.Result{{Type: "FixedArray[String]"}},
	},
	"test_fixedarray_output2_string_option": {
		Name:    "test_fixedarray_output2_string_option",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"test_fixedarray_output2_struct": {
		Name:    "test_fixedarray_output2_struct",
		Results: []*metadata.Result{{Type: "FixedArray[TestStruct2]"}},
	},
	"test_fixedarray_output2_struct_option": {
		Name:    "test_fixedarray_output2_struct_option",
		Results: []*metadata.Result{{Type: "FixedArray[TestStruct2?]"}},
	},
	"test_fixedarray_output_bool_0": {
		Name:    "test_fixedarray_output_bool_0",
		Results: []*metadata.Result{{Type: "FixedArray[Bool]"}},
	},
	"test_fixedarray_output_bool_1": {
		Name:    "test_fixedarray_output_bool_1",
		Results: []*metadata.Result{{Type: "FixedArray[Bool]"}},
	},
	"test_fixedarray_output_bool_2": {
		Name:    "test_fixedarray_output_bool_2",
		Results: []*metadata.Result{{Type: "FixedArray[Bool]"}},
	},
	"test_fixedarray_output_bool_3": {
		Name:    "test_fixedarray_output_bool_3",
		Results: []*metadata.Result{{Type: "FixedArray[Bool]"}},
	},
	"test_fixedarray_output_bool_4": {
		Name:    "test_fixedarray_output_bool_4",
		Results: []*metadata.Result{{Type: "FixedArray[Bool]"}},
	},
	"test_fixedarray_output_bool_option_0": {
		Name:    "test_fixedarray_output_bool_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[Bool?]"}},
	},
	"test_fixedarray_output_bool_option_1_false": {
		Name:    "test_fixedarray_output_bool_option_1_false",
		Results: []*metadata.Result{{Type: "FixedArray[Bool?]"}},
	},
	"test_fixedarray_output_bool_option_1_none": {
		Name:    "test_fixedarray_output_bool_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[Bool?]"}},
	},
	"test_fixedarray_output_bool_option_1_true": {
		Name:    "test_fixedarray_output_bool_option_1_true",
		Results: []*metadata.Result{{Type: "FixedArray[Bool?]"}},
	},
	"test_fixedarray_output_bool_option_2": {
		Name:    "test_fixedarray_output_bool_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[Bool?]"}},
	},
	"test_fixedarray_output_bool_option_3": {
		Name:    "test_fixedarray_output_bool_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[Bool?]"}},
	},
	"test_fixedarray_output_bool_option_4": {
		Name:    "test_fixedarray_output_bool_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[Bool?]"}},
	},
	"test_fixedarray_output_byte_0": {
		Name:    "test_fixedarray_output_byte_0",
		Results: []*metadata.Result{{Type: "FixedArray[Byte]"}},
	},
	"test_fixedarray_output_byte_1": {
		Name:    "test_fixedarray_output_byte_1",
		Results: []*metadata.Result{{Type: "FixedArray[Byte]"}},
	},
	"test_fixedarray_output_byte_2": {
		Name:    "test_fixedarray_output_byte_2",
		Results: []*metadata.Result{{Type: "FixedArray[Byte]"}},
	},
	"test_fixedarray_output_byte_3": {
		Name:    "test_fixedarray_output_byte_3",
		Results: []*metadata.Result{{Type: "FixedArray[Byte]"}},
	},
	"test_fixedarray_output_byte_4": {
		Name:    "test_fixedarray_output_byte_4",
		Results: []*metadata.Result{{Type: "FixedArray[Byte]"}},
	},
	"test_fixedarray_output_byte_option_0": {
		Name:    "test_fixedarray_output_byte_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[Byte?]"}},
	},
	"test_fixedarray_output_byte_option_1": {
		Name:    "test_fixedarray_output_byte_option_1",
		Results: []*metadata.Result{{Type: "FixedArray[Byte?]"}},
	},
	"test_fixedarray_output_byte_option_2": {
		Name:    "test_fixedarray_output_byte_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[Byte?]"}},
	},
	"test_fixedarray_output_byte_option_3": {
		Name:    "test_fixedarray_output_byte_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[Byte?]"}},
	},
	"test_fixedarray_output_byte_option_4": {
		Name:    "test_fixedarray_output_byte_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[Byte?]"}},
	},
	"test_fixedarray_output_char_0": {
		Name:    "test_fixedarray_output_char_0",
		Results: []*metadata.Result{{Type: "FixedArray[Char]"}},
	},
	"test_fixedarray_output_char_1": {
		Name:    "test_fixedarray_output_char_1",
		Results: []*metadata.Result{{Type: "FixedArray[Char]"}},
	},
	"test_fixedarray_output_char_2": {
		Name:    "test_fixedarray_output_char_2",
		Results: []*metadata.Result{{Type: "FixedArray[Char]"}},
	},
	"test_fixedarray_output_char_3": {
		Name:    "test_fixedarray_output_char_3",
		Results: []*metadata.Result{{Type: "FixedArray[Char]"}},
	},
	"test_fixedarray_output_char_4": {
		Name:    "test_fixedarray_output_char_4",
		Results: []*metadata.Result{{Type: "FixedArray[Char]"}},
	},
	"test_fixedarray_output_char_option_0": {
		Name:    "test_fixedarray_output_char_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[Char?]"}},
	},
	"test_fixedarray_output_char_option_1_none": {
		Name:    "test_fixedarray_output_char_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[Char?]"}},
	},
	"test_fixedarray_output_char_option_1_some": {
		Name:    "test_fixedarray_output_char_option_1_some",
		Results: []*metadata.Result{{Type: "FixedArray[Char?]"}},
	},
	"test_fixedarray_output_char_option_2": {
		Name:    "test_fixedarray_output_char_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[Char?]"}},
	},
	"test_fixedarray_output_char_option_3": {
		Name:    "test_fixedarray_output_char_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[Char?]"}},
	},
	"test_fixedarray_output_char_option_4": {
		Name:    "test_fixedarray_output_char_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[Char?]"}},
	},
	"test_fixedarray_output_double_0": {
		Name:    "test_fixedarray_output_double_0",
		Results: []*metadata.Result{{Type: "FixedArray[Double]"}},
	},
	"test_fixedarray_output_double_1": {
		Name:    "test_fixedarray_output_double_1",
		Results: []*metadata.Result{{Type: "FixedArray[Double]"}},
	},
	"test_fixedarray_output_double_2": {
		Name:    "test_fixedarray_output_double_2",
		Results: []*metadata.Result{{Type: "FixedArray[Double]"}},
	},
	"test_fixedarray_output_double_3": {
		Name:    "test_fixedarray_output_double_3",
		Results: []*metadata.Result{{Type: "FixedArray[Double]"}},
	},
	"test_fixedarray_output_double_4": {
		Name:    "test_fixedarray_output_double_4",
		Results: []*metadata.Result{{Type: "FixedArray[Double]"}},
	},
	"test_fixedarray_output_double_option_0": {
		Name:    "test_fixedarray_output_double_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[Double?]"}},
	},
	"test_fixedarray_output_double_option_1_none": {
		Name:    "test_fixedarray_output_double_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[Double?]"}},
	},
	"test_fixedarray_output_double_option_1_some": {
		Name:    "test_fixedarray_output_double_option_1_some",
		Results: []*metadata.Result{{Type: "FixedArray[Double?]"}},
	},
	"test_fixedarray_output_double_option_2": {
		Name:    "test_fixedarray_output_double_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[Double?]"}},
	},
	"test_fixedarray_output_double_option_3": {
		Name:    "test_fixedarray_output_double_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[Double?]"}},
	},
	"test_fixedarray_output_double_option_4": {
		Name:    "test_fixedarray_output_double_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[Double?]"}},
	},
	"test_fixedarray_output_float_0": {
		Name:    "test_fixedarray_output_float_0",
		Results: []*metadata.Result{{Type: "FixedArray[Float]"}},
	},
	"test_fixedarray_output_float_1": {
		Name:    "test_fixedarray_output_float_1",
		Results: []*metadata.Result{{Type: "FixedArray[Float]"}},
	},
	"test_fixedarray_output_float_2": {
		Name:    "test_fixedarray_output_float_2",
		Results: []*metadata.Result{{Type: "FixedArray[Float]"}},
	},
	"test_fixedarray_output_float_3": {
		Name:    "test_fixedarray_output_float_3",
		Results: []*metadata.Result{{Type: "FixedArray[Float]"}},
	},
	"test_fixedarray_output_float_4": {
		Name:    "test_fixedarray_output_float_4",
		Results: []*metadata.Result{{Type: "FixedArray[Float]"}},
	},
	"test_fixedarray_output_float_option_0": {
		Name:    "test_fixedarray_output_float_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[Float?]"}},
	},
	"test_fixedarray_output_float_option_1_none": {
		Name:    "test_fixedarray_output_float_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[Float?]"}},
	},
	"test_fixedarray_output_float_option_1_some": {
		Name:    "test_fixedarray_output_float_option_1_some",
		Results: []*metadata.Result{{Type: "FixedArray[Float?]"}},
	},
	"test_fixedarray_output_float_option_2": {
		Name:    "test_fixedarray_output_float_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[Float?]"}},
	},
	"test_fixedarray_output_float_option_3": {
		Name:    "test_fixedarray_output_float_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[Float?]"}},
	},
	"test_fixedarray_output_float_option_4": {
		Name:    "test_fixedarray_output_float_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[Float?]"}},
	},
	"test_fixedarray_output_int16_0": {
		Name:    "test_fixedarray_output_int16_0",
		Results: []*metadata.Result{{Type: "FixedArray[Int16]"}},
	},
	"test_fixedarray_output_int16_1": {
		Name:    "test_fixedarray_output_int16_1",
		Results: []*metadata.Result{{Type: "FixedArray[Int16]"}},
	},
	"test_fixedarray_output_int16_1_max": {
		Name:    "test_fixedarray_output_int16_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[Int16]"}},
	},
	"test_fixedarray_output_int16_1_min": {
		Name:    "test_fixedarray_output_int16_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[Int16]"}},
	},
	"test_fixedarray_output_int16_2": {
		Name:    "test_fixedarray_output_int16_2",
		Results: []*metadata.Result{{Type: "FixedArray[Int16]"}},
	},
	"test_fixedarray_output_int16_3": {
		Name:    "test_fixedarray_output_int16_3",
		Results: []*metadata.Result{{Type: "FixedArray[Int16]"}},
	},
	"test_fixedarray_output_int16_4": {
		Name:    "test_fixedarray_output_int16_4",
		Results: []*metadata.Result{{Type: "FixedArray[Int16]"}},
	},
	"test_fixedarray_output_int16_option_0": {
		Name:    "test_fixedarray_output_int16_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[Int16?]"}},
	},
	"test_fixedarray_output_int16_option_1_max": {
		Name:    "test_fixedarray_output_int16_option_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[Int16?]"}},
	},
	"test_fixedarray_output_int16_option_1_min": {
		Name:    "test_fixedarray_output_int16_option_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[Int16?]"}},
	},
	"test_fixedarray_output_int16_option_1_none": {
		Name:    "test_fixedarray_output_int16_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[Int16?]"}},
	},
	"test_fixedarray_output_int16_option_2": {
		Name:    "test_fixedarray_output_int16_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[Int16?]"}},
	},
	"test_fixedarray_output_int16_option_3": {
		Name:    "test_fixedarray_output_int16_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[Int16?]"}},
	},
	"test_fixedarray_output_int16_option_4": {
		Name:    "test_fixedarray_output_int16_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[Int16?]"}},
	},
	"test_fixedarray_output_int64_0": {
		Name:    "test_fixedarray_output_int64_0",
		Results: []*metadata.Result{{Type: "FixedArray[Int64]"}},
	},
	"test_fixedarray_output_int64_1": {
		Name:    "test_fixedarray_output_int64_1",
		Results: []*metadata.Result{{Type: "FixedArray[Int64]"}},
	},
	"test_fixedarray_output_int64_1_max": {
		Name:    "test_fixedarray_output_int64_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[Int64]"}},
	},
	"test_fixedarray_output_int64_1_min": {
		Name:    "test_fixedarray_output_int64_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[Int64]"}},
	},
	"test_fixedarray_output_int64_2": {
		Name:    "test_fixedarray_output_int64_2",
		Results: []*metadata.Result{{Type: "FixedArray[Int64]"}},
	},
	"test_fixedarray_output_int64_3": {
		Name:    "test_fixedarray_output_int64_3",
		Results: []*metadata.Result{{Type: "FixedArray[Int64]"}},
	},
	"test_fixedarray_output_int64_4": {
		Name:    "test_fixedarray_output_int64_4",
		Results: []*metadata.Result{{Type: "FixedArray[Int64]"}},
	},
	"test_fixedarray_output_int64_option_0": {
		Name:    "test_fixedarray_output_int64_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[Int64?]"}},
	},
	"test_fixedarray_output_int64_option_1_max": {
		Name:    "test_fixedarray_output_int64_option_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[Int64?]"}},
	},
	"test_fixedarray_output_int64_option_1_min": {
		Name:    "test_fixedarray_output_int64_option_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[Int64?]"}},
	},
	"test_fixedarray_output_int64_option_1_none": {
		Name:    "test_fixedarray_output_int64_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[Int64?]"}},
	},
	"test_fixedarray_output_int64_option_2": {
		Name:    "test_fixedarray_output_int64_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[Int64?]"}},
	},
	"test_fixedarray_output_int64_option_3": {
		Name:    "test_fixedarray_output_int64_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[Int64?]"}},
	},
	"test_fixedarray_output_int64_option_4": {
		Name:    "test_fixedarray_output_int64_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[Int64?]"}},
	},
	"test_fixedarray_output_int_0": {
		Name:    "test_fixedarray_output_int_0",
		Results: []*metadata.Result{{Type: "FixedArray[Int]"}},
	},
	"test_fixedarray_output_int_1": {
		Name:    "test_fixedarray_output_int_1",
		Results: []*metadata.Result{{Type: "FixedArray[Int]"}},
	},
	"test_fixedarray_output_int_1_max": {
		Name:    "test_fixedarray_output_int_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[Int]"}},
	},
	"test_fixedarray_output_int_1_min": {
		Name:    "test_fixedarray_output_int_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[Int]"}},
	},
	"test_fixedarray_output_int_2": {
		Name:    "test_fixedarray_output_int_2",
		Results: []*metadata.Result{{Type: "FixedArray[Int]"}},
	},
	"test_fixedarray_output_int_3": {
		Name:    "test_fixedarray_output_int_3",
		Results: []*metadata.Result{{Type: "FixedArray[Int]"}},
	},
	"test_fixedarray_output_int_4": {
		Name:    "test_fixedarray_output_int_4",
		Results: []*metadata.Result{{Type: "FixedArray[Int]"}},
	},
	"test_fixedarray_output_int_option_0": {
		Name:    "test_fixedarray_output_int_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"test_fixedarray_output_int_option_1_max": {
		Name:    "test_fixedarray_output_int_option_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"test_fixedarray_output_int_option_1_min": {
		Name:    "test_fixedarray_output_int_option_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"test_fixedarray_output_int_option_1_none": {
		Name:    "test_fixedarray_output_int_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"test_fixedarray_output_int_option_2": {
		Name:    "test_fixedarray_output_int_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"test_fixedarray_output_int_option_3": {
		Name:    "test_fixedarray_output_int_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"test_fixedarray_output_int_option_4": {
		Name:    "test_fixedarray_output_int_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"test_fixedarray_output_string_0": {
		Name:    "test_fixedarray_output_string_0",
		Results: []*metadata.Result{{Type: "FixedArray[String]"}},
	},
	"test_fixedarray_output_string_1": {
		Name:    "test_fixedarray_output_string_1",
		Results: []*metadata.Result{{Type: "FixedArray[String]"}},
	},
	"test_fixedarray_output_string_2": {
		Name:    "test_fixedarray_output_string_2",
		Results: []*metadata.Result{{Type: "FixedArray[String]"}},
	},
	"test_fixedarray_output_string_3": {
		Name:    "test_fixedarray_output_string_3",
		Results: []*metadata.Result{{Type: "FixedArray[String]"}},
	},
	"test_fixedarray_output_string_4": {
		Name:    "test_fixedarray_output_string_4",
		Results: []*metadata.Result{{Type: "FixedArray[String]"}},
	},
	"test_fixedarray_output_string_option_0": {
		Name:    "test_fixedarray_output_string_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"test_fixedarray_output_string_option_1_none": {
		Name:    "test_fixedarray_output_string_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"test_fixedarray_output_string_option_1_some": {
		Name:    "test_fixedarray_output_string_option_1_some",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"test_fixedarray_output_string_option_2": {
		Name:    "test_fixedarray_output_string_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"test_fixedarray_output_string_option_3": {
		Name:    "test_fixedarray_output_string_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"test_fixedarray_output_string_option_4": {
		Name:    "test_fixedarray_output_string_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[String?]"}},
	},
	"test_fixedarray_output_uint16_0": {
		Name:    "test_fixedarray_output_uint16_0",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16]"}},
	},
	"test_fixedarray_output_uint16_1": {
		Name:    "test_fixedarray_output_uint16_1",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16]"}},
	},
	"test_fixedarray_output_uint16_1_max": {
		Name:    "test_fixedarray_output_uint16_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16]"}},
	},
	"test_fixedarray_output_uint16_1_min": {
		Name:    "test_fixedarray_output_uint16_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16]"}},
	},
	"test_fixedarray_output_uint16_2": {
		Name:    "test_fixedarray_output_uint16_2",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16]"}},
	},
	"test_fixedarray_output_uint16_3": {
		Name:    "test_fixedarray_output_uint16_3",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16]"}},
	},
	"test_fixedarray_output_uint16_4": {
		Name:    "test_fixedarray_output_uint16_4",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16]"}},
	},
	"test_fixedarray_output_uint16_option_0": {
		Name:    "test_fixedarray_output_uint16_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16?]"}},
	},
	"test_fixedarray_output_uint16_option_1_max": {
		Name:    "test_fixedarray_output_uint16_option_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16?]"}},
	},
	"test_fixedarray_output_uint16_option_1_min": {
		Name:    "test_fixedarray_output_uint16_option_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16?]"}},
	},
	"test_fixedarray_output_uint16_option_1_none": {
		Name:    "test_fixedarray_output_uint16_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16?]"}},
	},
	"test_fixedarray_output_uint16_option_2": {
		Name:    "test_fixedarray_output_uint16_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16?]"}},
	},
	"test_fixedarray_output_uint16_option_3": {
		Name:    "test_fixedarray_output_uint16_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16?]"}},
	},
	"test_fixedarray_output_uint16_option_4": {
		Name:    "test_fixedarray_output_uint16_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[UInt16?]"}},
	},
	"test_fixedarray_output_uint64_0": {
		Name:    "test_fixedarray_output_uint64_0",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64]"}},
	},
	"test_fixedarray_output_uint64_1": {
		Name:    "test_fixedarray_output_uint64_1",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64]"}},
	},
	"test_fixedarray_output_uint64_1_max": {
		Name:    "test_fixedarray_output_uint64_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64]"}},
	},
	"test_fixedarray_output_uint64_1_min": {
		Name:    "test_fixedarray_output_uint64_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64]"}},
	},
	"test_fixedarray_output_uint64_2": {
		Name:    "test_fixedarray_output_uint64_2",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64]"}},
	},
	"test_fixedarray_output_uint64_3": {
		Name:    "test_fixedarray_output_uint64_3",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64]"}},
	},
	"test_fixedarray_output_uint64_4": {
		Name:    "test_fixedarray_output_uint64_4",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64]"}},
	},
	"test_fixedarray_output_uint64_option_0": {
		Name:    "test_fixedarray_output_uint64_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64?]"}},
	},
	"test_fixedarray_output_uint64_option_1_max": {
		Name:    "test_fixedarray_output_uint64_option_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64?]"}},
	},
	"test_fixedarray_output_uint64_option_1_min": {
		Name:    "test_fixedarray_output_uint64_option_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64?]"}},
	},
	"test_fixedarray_output_uint64_option_1_none": {
		Name:    "test_fixedarray_output_uint64_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64?]"}},
	},
	"test_fixedarray_output_uint64_option_2": {
		Name:    "test_fixedarray_output_uint64_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64?]"}},
	},
	"test_fixedarray_output_uint64_option_3": {
		Name:    "test_fixedarray_output_uint64_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64?]"}},
	},
	"test_fixedarray_output_uint64_option_4": {
		Name:    "test_fixedarray_output_uint64_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[UInt64?]"}},
	},
	"test_fixedarray_output_uint_0": {
		Name:    "test_fixedarray_output_uint_0",
		Results: []*metadata.Result{{Type: "FixedArray[UInt]"}},
	},
	"test_fixedarray_output_uint_1": {
		Name:    "test_fixedarray_output_uint_1",
		Results: []*metadata.Result{{Type: "FixedArray[UInt]"}},
	},
	"test_fixedarray_output_uint_1_max": {
		Name:    "test_fixedarray_output_uint_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[UInt]"}},
	},
	"test_fixedarray_output_uint_1_min": {
		Name:    "test_fixedarray_output_uint_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[UInt]"}},
	},
	"test_fixedarray_output_uint_2": {
		Name:    "test_fixedarray_output_uint_2",
		Results: []*metadata.Result{{Type: "FixedArray[UInt]"}},
	},
	"test_fixedarray_output_uint_3": {
		Name:    "test_fixedarray_output_uint_3",
		Results: []*metadata.Result{{Type: "FixedArray[UInt]"}},
	},
	"test_fixedarray_output_uint_4": {
		Name:    "test_fixedarray_output_uint_4",
		Results: []*metadata.Result{{Type: "FixedArray[UInt]"}},
	},
	"test_fixedarray_output_uint_option_0": {
		Name:    "test_fixedarray_output_uint_option_0",
		Results: []*metadata.Result{{Type: "FixedArray[UInt?]"}},
	},
	"test_fixedarray_output_uint_option_1_max": {
		Name:    "test_fixedarray_output_uint_option_1_max",
		Results: []*metadata.Result{{Type: "FixedArray[UInt?]"}},
	},
	"test_fixedarray_output_uint_option_1_min": {
		Name:    "test_fixedarray_output_uint_option_1_min",
		Results: []*metadata.Result{{Type: "FixedArray[UInt?]"}},
	},
	"test_fixedarray_output_uint_option_1_none": {
		Name:    "test_fixedarray_output_uint_option_1_none",
		Results: []*metadata.Result{{Type: "FixedArray[UInt?]"}},
	},
	"test_fixedarray_output_uint_option_2": {
		Name:    "test_fixedarray_output_uint_option_2",
		Results: []*metadata.Result{{Type: "FixedArray[UInt?]"}},
	},
	"test_fixedarray_output_uint_option_3": {
		Name:    "test_fixedarray_output_uint_option_3",
		Results: []*metadata.Result{{Type: "FixedArray[UInt?]"}},
	},
	"test_fixedarray_output_uint_option_4": {
		Name:    "test_fixedarray_output_uint_option_4",
		Results: []*metadata.Result{{Type: "FixedArray[UInt?]"}},
	},
	"test_float_input_max": {
		Name:       "test_float_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Float"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_float_input_min": {
		Name:       "test_float_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Float"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_float_option_input_max": {
		Name:       "test_float_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Float?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_float_option_input_min": {
		Name:       "test_float_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Float?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_float_option_input_none": {
		Name:       "test_float_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Float?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_float_option_output_max": {
		Name:    "test_float_option_output_max",
		Results: []*metadata.Result{{Type: "Float?"}},
	},
	"test_float_option_output_min": {
		Name:    "test_float_option_output_min",
		Results: []*metadata.Result{{Type: "Float?"}},
	},
	"test_float_option_output_none": {
		Name:    "test_float_option_output_none",
		Results: []*metadata.Result{{Type: "Float?"}},
	},
	"test_float_output_max": {Name: "test_float_output_max", Results: []*metadata.Result{{Type: "Float"}}},
	"test_float_output_min": {Name: "test_float_output_min", Results: []*metadata.Result{{Type: "Float"}}},
	"test_generate_map_string_string_output": {
		Name:    "test_generate_map_string_string_output",
		Results: []*metadata.Result{{Type: "Map[String, String]"}},
	},
	"test_http_header": {
		Name:       "test_http_header",
		Parameters: []*metadata.Parameter{{Name: "h", Type: "HttpHeader?"}},
	},
	"test_http_header_map": {
		Name:       "test_http_header_map",
		Parameters: []*metadata.Parameter{{Name: "m", Type: "Map[String, HttpHeader?]"}},
	},
	"test_http_headers": {
		Name:       "test_http_headers",
		Parameters: []*metadata.Parameter{{Name: "h", Type: "HttpHeaders"}},
	},
	"test_http_response_headers_output": {
		Name:    "test_http_response_headers_output",
		Results: []*metadata.Result{{Type: "HttpResponse?"}},
	},
	"test_http_response_headers": {
		Name:       "test_http_response_headers",
		Parameters: []*metadata.Parameter{{Name: "r", Type: "HttpResponse?"}},
	},
	"test_int16_input_max": {
		Name:       "test_int16_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int16"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int16_input_min": {
		Name:       "test_int16_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int16"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int16_option_input_max": {
		Name:       "test_int16_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int16?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int16_option_input_min": {
		Name:       "test_int16_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int16?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int16_option_input_none": {
		Name:       "test_int16_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int16?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int16_option_output_max": {
		Name:    "test_int16_option_output_max",
		Results: []*metadata.Result{{Type: "Int16?"}},
	},
	"test_int16_option_output_min": {
		Name:    "test_int16_option_output_min",
		Results: []*metadata.Result{{Type: "Int16?"}},
	},
	"test_int16_option_output_none": {
		Name:    "test_int16_option_output_none",
		Results: []*metadata.Result{{Type: "Int16?"}},
	},
	"test_int16_output_max": {Name: "test_int16_output_max", Results: []*metadata.Result{{Type: "Int16"}}},
	"test_int16_output_min": {Name: "test_int16_output_min", Results: []*metadata.Result{{Type: "Int16"}}},
	"test_int64_input_max": {
		Name:       "test_int64_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int64"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int64_input_min": {
		Name:       "test_int64_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int64"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int64_option_input_max": {
		Name:       "test_int64_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int64?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int64_option_input_min": {
		Name:       "test_int64_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int64?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int64_option_input_none": {
		Name:       "test_int64_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int64?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int64_option_output_max": {
		Name:    "test_int64_option_output_max",
		Results: []*metadata.Result{{Type: "Int64?"}},
	},
	"test_int64_option_output_min": {
		Name:    "test_int64_option_output_min",
		Results: []*metadata.Result{{Type: "Int64?"}},
	},
	"test_int64_option_output_none": {
		Name:    "test_int64_option_output_none",
		Results: []*metadata.Result{{Type: "Int64?"}},
	},
	"test_int64_output_max": {Name: "test_int64_output_max", Results: []*metadata.Result{{Type: "Int64"}}},
	"test_int64_output_min": {Name: "test_int64_output_min", Results: []*metadata.Result{{Type: "Int64"}}},
	"test_int_input_max": {
		Name:       "test_int_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int_input_min": {
		Name:       "test_int_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int_option_input_max": {
		Name:       "test_int_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int_option_input_min": {
		Name:       "test_int_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int_option_input_none": {
		Name:       "test_int_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int_option_output_max":  {Name: "test_int_option_output_max", Results: []*metadata.Result{{Type: "Int?"}}},
	"test_int_option_output_min":  {Name: "test_int_option_output_min", Results: []*metadata.Result{{Type: "Int?"}}},
	"test_int_option_output_none": {Name: "test_int_option_output_none", Results: []*metadata.Result{{Type: "Int?"}}},
	"test_int_output_max":         {Name: "test_int_output_max", Results: []*metadata.Result{{Type: "Int"}}},
	"test_int_output_min":         {Name: "test_int_output_min", Results: []*metadata.Result{{Type: "Int"}}},
	"test_iterate_map_string_string": {
		Name:       "test_iterate_map_string_string",
		Parameters: []*metadata.Parameter{{Name: "m", Type: "Map[String, String]"}},
	},
	"test_map_input_int_double": {
		Name:       "test_map_input_int_double",
		Parameters: []*metadata.Parameter{{Name: "m", Type: "Map[Int, Double]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_map_input_int_float": {
		Name:       "test_map_input_int_float",
		Parameters: []*metadata.Parameter{{Name: "m", Type: "Map[Int, Float]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_map_input_string_string": {
		Name:       "test_map_input_string_string",
		Parameters: []*metadata.Parameter{{Name: "m", Type: "Map[String, String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_map_lookup_string_string": {
		Name:       "test_map_lookup_string_string",
		Parameters: []*metadata.Parameter{{Name: "m", Type: "Map[String, String]"}, {Name: "key", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"test_map_option_input_string_string": {
		Name:       "test_map_option_input_string_string",
		Parameters: []*metadata.Parameter{{Name: "m", Type: "Map[String, String]?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_map_option_output_string_string": {
		Name:    "test_map_option_output_string_string",
		Results: []*metadata.Result{{Type: "Map[String, String]?"}},
	},
	"test_map_output_int_double": {
		Name:    "test_map_output_int_double",
		Results: []*metadata.Result{{Type: "Map[Int, Double]"}},
	},
	"test_map_output_int_float": {
		Name:    "test_map_output_int_float",
		Results: []*metadata.Result{{Type: "Map[Int, Float]"}},
	},
	"test_map_output_string_string": {
		Name:    "test_map_output_string_string",
		Results: []*metadata.Result{{Type: "Map[String, String]"}},
	},
	"test_option_fixedarray_input1_int": {
		Name:       "test_option_fixedarray_input1_int",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int]?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_option_fixedarray_input1_string": {
		Name:       "test_option_fixedarray_input1_string",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String]?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_option_fixedarray_input2_int": {
		Name:       "test_option_fixedarray_input2_int",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[Int]?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_option_fixedarray_input2_string": {
		Name:       "test_option_fixedarray_input2_string",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "FixedArray[String]?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_option_fixedarray_output1_int": {
		Name:    "test_option_fixedarray_output1_int",
		Results: []*metadata.Result{{Type: "FixedArray[Int]?"}},
	},
	"test_option_fixedarray_output1_string": {
		Name:    "test_option_fixedarray_output1_string",
		Results: []*metadata.Result{{Type: "FixedArray[String]?"}},
	},
	"test_option_fixedarray_output2_int": {
		Name:    "test_option_fixedarray_output2_int",
		Results: []*metadata.Result{{Type: "FixedArray[Int]?"}},
	},
	"test_option_fixedarray_output2_string": {
		Name:    "test_option_fixedarray_output2_string",
		Results: []*metadata.Result{{Type: "FixedArray[String]?"}},
	},
	"test_recursive_struct_input": {
		Name:       "test_recursive_struct_input",
		Parameters: []*metadata.Parameter{{Name: "r1", Type: "TestRecursiveStruct"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_recursive_struct_option_input": {
		Name:       "test_recursive_struct_option_input",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestRecursiveStruct?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_recursive_struct_option_input_none": {
		Name:       "test_recursive_struct_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestRecursiveStruct?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_recursive_struct_option_output": {
		Name:    "test_recursive_struct_option_output",
		Results: []*metadata.Result{{Type: "TestRecursiveStruct?"}},
	},
	"test_recursive_struct_option_output_map": {
		Name:    "test_recursive_struct_option_output_map",
		Results: []*metadata.Result{{Type: "TestRecursiveStruct_map?"}},
	},
	"test_recursive_struct_option_output_none": {
		Name:    "test_recursive_struct_option_output_none",
		Results: []*metadata.Result{{Type: "TestRecursiveStruct?"}},
	},
	"test_recursive_struct_output": {
		Name:    "test_recursive_struct_output",
		Results: []*metadata.Result{{Type: "TestRecursiveStruct"}},
	},
	"test_recursive_struct_output_map": {
		Name:    "test_recursive_struct_output_map",
		Results: []*metadata.Result{{Type: "TestRecursiveStruct_map"}},
	},
	"test_smorgasbord_struct_input": {
		Name:       "test_smorgasbord_struct_input",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestSmorgasbordStruct"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_smorgasbord_struct_option_input": {
		Name:       "test_smorgasbord_struct_option_input",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestSmorgasbordStruct?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_smorgasbord_struct_option_input_none": {
		Name:       "test_smorgasbord_struct_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestSmorgasbordStruct?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_smorgasbord_struct_option_output": {
		Name:    "test_smorgasbord_struct_option_output",
		Results: []*metadata.Result{{Type: "TestSmorgasbordStruct?"}},
	},
	"test_smorgasbord_struct_option_output_map": {
		Name:    "test_smorgasbord_struct_option_output_map",
		Results: []*metadata.Result{{Type: "TestSmorgasbordStruct_map?"}},
	},
	"test_smorgasbord_struct_option_output_map_none": {
		Name:    "test_smorgasbord_struct_option_output_map_none",
		Results: []*metadata.Result{{Type: "TestSmorgasbordStruct_map?"}},
	},
	"test_smorgasbord_struct_option_output_none": {
		Name:    "test_smorgasbord_struct_option_output_none",
		Results: []*metadata.Result{{Type: "TestSmorgasbordStruct?"}},
	},
	"test_smorgasbord_struct_output": {
		Name:    "test_smorgasbord_struct_output",
		Results: []*metadata.Result{{Type: "TestSmorgasbordStruct"}},
	},
	"test_smorgasbord_struct_output_map": {
		Name:    "test_smorgasbord_struct_output_map",
		Results: []*metadata.Result{{Type: "TestSmorgasbordStruct_map"}},
	},
	"test_string_input": {
		Name:       "test_string_input",
		Parameters: []*metadata.Parameter{{Name: "s", Type: "String"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_string_option_input": {
		Name:       "test_string_option_input",
		Parameters: []*metadata.Parameter{{Name: "s", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_string_option_input_none": {
		Name:       "test_string_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "s", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_string_option_output": {
		Name:    "test_string_option_output",
		Results: []*metadata.Result{{Type: "String?"}},
	},
	"test_string_option_output_none": {
		Name:    "test_string_option_output_none",
		Results: []*metadata.Result{{Type: "String?"}},
	},
	"test_string_output": {Name: "test_string_output", Results: []*metadata.Result{{Type: "String"}}},
	"test_struct_containing_map_input_string_string": {
		Name:       "test_struct_containing_map_input_string_string",
		Parameters: []*metadata.Parameter{{Name: "s", Type: "TestStructWithMap"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_containing_map_output_string_string": {
		Name:    "test_struct_containing_map_output_string_string",
		Results: []*metadata.Result{{Type: "TestStructWithMap"}},
	},
	"test_struct_input1": {
		Name:       "test_struct_input1",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct1"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_input2": {
		Name:       "test_struct_input2",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct2"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_input3": {
		Name:       "test_struct_input3",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct3"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_input4": {
		Name:       "test_struct_input4",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct4"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_input4_with_none": {
		Name:       "test_struct_input4_with_none",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct4"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_input5": {
		Name:       "test_struct_input5",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct5"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input1": {
		Name:       "test_struct_option_input1",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct1?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input1_none": {
		Name:       "test_struct_option_input1_none",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct1?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input2": {
		Name:       "test_struct_option_input2",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct2?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input2_none": {
		Name:       "test_struct_option_input2_none",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct2?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input3": {
		Name:       "test_struct_option_input3",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct3?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input3_none": {
		Name:       "test_struct_option_input3_none",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct3?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input4": {
		Name:       "test_struct_option_input4",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct4?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input4_none": {
		Name:       "test_struct_option_input4_none",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct4?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input4_with_none": {
		Name:       "test_struct_option_input4_with_none",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct4?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input5": {
		Name:       "test_struct_option_input5",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct5?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_input5_none": {
		Name:       "test_struct_option_input5_none",
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestStruct5?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_option_output1": {
		Name:    "test_struct_option_output1",
		Results: []*metadata.Result{{Type: "TestStruct1?"}},
	},
	"test_struct_option_output1_map": {
		Name:    "test_struct_option_output1_map",
		Results: []*metadata.Result{{Type: "TestStruct1_map?"}},
	},
	"test_struct_option_output1_none": {
		Name:    "test_struct_option_output1_none",
		Results: []*metadata.Result{{Type: "TestStruct1?"}},
	},
	"test_struct_option_output2": {
		Name:    "test_struct_option_output2",
		Results: []*metadata.Result{{Type: "TestStruct2?"}},
	},
	"test_struct_option_output2_map": {
		Name:    "test_struct_option_output2_map",
		Results: []*metadata.Result{{Type: "TestStruct2_map?"}},
	},
	"test_struct_option_output2_none": {
		Name:    "test_struct_option_output2_none",
		Results: []*metadata.Result{{Type: "TestStruct2?"}},
	},
	"test_struct_option_output3": {
		Name:    "test_struct_option_output3",
		Results: []*metadata.Result{{Type: "TestStruct3?"}},
	},
	"test_struct_option_output3_map": {
		Name:    "test_struct_option_output3_map",
		Results: []*metadata.Result{{Type: "TestStruct3_map?"}},
	},
	"test_struct_option_output3_none": {
		Name:    "test_struct_option_output3_none",
		Results: []*metadata.Result{{Type: "TestStruct3?"}},
	},
	"test_struct_option_output4": {
		Name:    "test_struct_option_output4",
		Results: []*metadata.Result{{Type: "TestStruct4?"}},
	},
	"test_struct_option_output4_map": {
		Name:    "test_struct_option_output4_map",
		Results: []*metadata.Result{{Type: "TestStruct4_map?"}},
	},
	"test_struct_option_output4_map_with_none": {
		Name:    "test_struct_option_output4_map_with_none",
		Results: []*metadata.Result{{Type: "TestStruct4_map?"}},
	},
	"test_struct_option_output4_none": {
		Name:    "test_struct_option_output4_none",
		Results: []*metadata.Result{{Type: "TestStruct4?"}},
	},
	"test_struct_option_output4_with_none": {
		Name:    "test_struct_option_output4_with_none",
		Results: []*metadata.Result{{Type: "TestStruct4?"}},
	},
	"test_struct_option_output5": {
		Name:    "test_struct_option_output5",
		Results: []*metadata.Result{{Type: "TestStruct5?"}},
	},
	"test_struct_option_output5_map": {
		Name:    "test_struct_option_output5_map",
		Results: []*metadata.Result{{Type: "TestStruct5_map?"}},
	},
	"test_struct_option_output5_none": {
		Name:    "test_struct_option_output5_none",
		Results: []*metadata.Result{{Type: "TestStruct5?"}},
	},
	"test_struct_output1": {
		Name:    "test_struct_output1",
		Results: []*metadata.Result{{Type: "TestStruct1"}},
	},
	"test_struct_output1_map": {
		Name:    "test_struct_output1_map",
		Results: []*metadata.Result{{Type: "TestStruct1_map"}},
	},
	"test_struct_output2": {
		Name:    "test_struct_output2",
		Results: []*metadata.Result{{Type: "TestStruct2"}},
	},
	"test_struct_output2_map": {
		Name:    "test_struct_output2_map",
		Results: []*metadata.Result{{Type: "TestStruct2_map"}},
	},
	"test_struct_output3": {
		Name:    "test_struct_output3",
		Results: []*metadata.Result{{Type: "TestStruct3"}},
	},
	"test_struct_output3_map": {
		Name:    "test_struct_output3_map",
		Results: []*metadata.Result{{Type: "TestStruct3_map"}},
	},
	"test_struct_output4": {
		Name:    "test_struct_output4",
		Results: []*metadata.Result{{Type: "TestStruct4"}},
	},
	"test_struct_output4_map": {
		Name:    "test_struct_output4_map",
		Results: []*metadata.Result{{Type: "TestStruct4_map"}},
	},
	"test_struct_output4_map_with_none": {
		Name:    "test_struct_output4_map_with_none",
		Results: []*metadata.Result{{Type: "TestStruct4_map"}},
	},
	"test_struct_output4_with_none": {
		Name:    "test_struct_output4_with_none",
		Results: []*metadata.Result{{Type: "TestStruct4"}},
	},
	"test_struct_output5": {
		Name:    "test_struct_output5",
		Results: []*metadata.Result{{Type: "TestStruct5"}},
	},
	"test_struct_output5_map": {
		Name:    "test_struct_output5_map",
		Results: []*metadata.Result{{Type: "TestStruct5_map"}},
	},
	"test_time_input": {
		Name:       "test_time_input",
		Parameters: []*metadata.Parameter{{Name: "t", Type: "@time.ZonedDateTime"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_time_option_input": {
		Name:       "test_time_option_input",
		Parameters: []*metadata.Parameter{{Name: "t", Type: "@time.ZonedDateTime?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_time_option_input_none": {
		Name:       "test_time_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "t", Type: "@time.ZonedDateTime?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_time_option_input_none_style2": {
		Name:       "test_time_option_input_none_style2",
		Parameters: []*metadata.Parameter{{Name: "t?", Type: "@time.ZonedDateTime"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_time_option_input_style2": {
		Name:       "test_time_option_input_style2",
		Parameters: []*metadata.Parameter{{Name: "t?", Type: "@time.ZonedDateTime"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_time_option_output": {
		Name:    "test_time_option_output",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime?"}},
	},
	"test_time_option_output_none": {
		Name:    "test_time_option_output_none",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime?"}},
	},
	"test_time_output": {
		Name:    "test_time_output",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime"}},
	},
	"test_tuple_output": {
		Name:    "test_tuple_output",
		Results: []*metadata.Result{{Type: "(Int, Bool, String)"}},
	},
	"test_uint16_input_max": {
		Name:       "test_uint16_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt16"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint16_input_min": {
		Name:       "test_uint16_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt16"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint16_option_input_max": {
		Name:       "test_uint16_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt16?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint16_option_input_min": {
		Name:       "test_uint16_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt16?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint16_option_input_none": {
		Name:       "test_uint16_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt16?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint16_option_output_max": {
		Name:    "test_uint16_option_output_max",
		Results: []*metadata.Result{{Type: "UInt16?"}},
	},
	"test_uint16_option_output_min": {
		Name:    "test_uint16_option_output_min",
		Results: []*metadata.Result{{Type: "UInt16?"}},
	},
	"test_uint16_option_output_none": {
		Name:    "test_uint16_option_output_none",
		Results: []*metadata.Result{{Type: "UInt16?"}},
	},
	"test_uint16_output_max": {Name: "test_uint16_output_max", Results: []*metadata.Result{{Type: "UInt16"}}},
	"test_uint16_output_min": {Name: "test_uint16_output_min", Results: []*metadata.Result{{Type: "UInt16"}}},
	"test_uint64_input_max": {
		Name:       "test_uint64_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt64"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint64_input_min": {
		Name:       "test_uint64_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt64"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint64_option_input_max": {
		Name:       "test_uint64_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt64?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint64_option_input_min": {
		Name:       "test_uint64_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt64?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint64_option_input_none": {
		Name:       "test_uint64_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt64?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint64_option_output_max": {
		Name:    "test_uint64_option_output_max",
		Results: []*metadata.Result{{Type: "UInt64?"}},
	},
	"test_uint64_option_output_min": {
		Name:    "test_uint64_option_output_min",
		Results: []*metadata.Result{{Type: "UInt64?"}},
	},
	"test_uint64_option_output_none": {
		Name:    "test_uint64_option_output_none",
		Results: []*metadata.Result{{Type: "UInt64?"}},
	},
	"test_uint64_output_max": {Name: "test_uint64_output_max", Results: []*metadata.Result{{Type: "UInt64"}}},
	"test_uint64_output_min": {Name: "test_uint64_output_min", Results: []*metadata.Result{{Type: "UInt64"}}},
	"test_uint_input_max": {
		Name:       "test_uint_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint_input_min": {
		Name:       "test_uint_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint_option_input_max": {
		Name:       "test_uint_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint_option_input_min": {
		Name:       "test_uint_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint_option_input_none": {
		Name:       "test_uint_option_input_none",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "UInt?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_uint_option_output_max": {
		Name:    "test_uint_option_output_max",
		Results: []*metadata.Result{{Type: "UInt?"}},
	},
	"test_uint_option_output_min": {
		Name:    "test_uint_option_output_min",
		Results: []*metadata.Result{{Type: "UInt?"}},
	},
	"test_uint_option_output_none": {
		Name:    "test_uint_option_output_none",
		Results: []*metadata.Result{{Type: "UInt?"}},
	},
	"test_uint_output_max": {Name: "test_uint_output_max", Results: []*metadata.Result{{Type: "UInt"}}},
	"test_uint_output_min": {Name: "test_uint_output_min", Results: []*metadata.Result{{Type: "UInt"}}},
}

var wantRuntimeFnImports = metadata.FunctionMap{
	// "modus_system.getTimeInZone": {
	// 	Name:       "modus_system.getTimeInZone",
	// 	Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
	// 	Results:    []*metadata.Result{{Type: "String"}},
	// },
	// "modus_system.getTimeZoneData": {
	// 	Name: "modus_system.getTimeZoneData",
	// 	Parameters: []*metadata.Parameter{
	// 		{Name: "tz", Type: "String"},
	// 		{Name: "format", Type: "String"},
	// 	},
	// 	Results: []*metadata.Result{{Type: "Array[Byte]"}},
	// },
	"modus_system.logMessage": {
		Name: "modus_system.logMessage",
		Parameters: []*metadata.Parameter{
			{Name: "level", Type: "String"},
			{Name: "message", Type: "String"},
		},
	},
	"modus_test.add": {
		Name:       "modus_test.add",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"modus_test.echo1": {
		Name:       "modus_test.echo1",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"modus_test.echo2": {
		Name:       "modus_test.echo2",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"modus_test.echo3": {
		Name:       "modus_test.echo3",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"modus_test.echo4": {
		Name:       "modus_test.echo4",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"modus_test.encodeStrings1": {
		Name:       "modus_test.encodeStrings1",
		Parameters: []*metadata.Parameter{{Name: "items", Type: "Array[String]?"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"modus_test.encodeStrings2": {
		Name:       "modus_test.encodeStrings2",
		Parameters: []*metadata.Parameter{{Name: "items", Type: "Array[String?]?"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
}

var wantRuntimeTypes = metadata.TypeMap{
	"(Int, Bool, String)": {
		Name: "(Int, Bool, String)",
		Fields: []*metadata.Field{
			{Name: "0", Type: "Int"}, {Name: "1", Type: "Bool"},
			{Name: "2", Type: "String"},
		},
	},
	"(String)": {
		Name:   "(String)",
		Fields: []*metadata.Field{{Name: "0", Type: "String"}},
	},
	"@ffi.XExternByteArray":   {Name: "@ffi.XExternByteArray"},
	"@ffi.XExternString":      {Name: "@ffi.XExternString"},
	"@ffi.XExternStringArray": {Name: "@ffi.XExternStringArray"},
	"@testutils.CallStack[T]": {Name: "@testutils.CallStack[T]",
		Fields: []*metadata.Field{{Name: "items", Type: "Array[Array[@testutils.T]]"}},
	},
	"@time.Duration":                   {Name: "@time.Duration"},
	"@time.Duration!Error":             {Name: "@time.Duration!Error"},
	"@time.Duration?":                  {Name: "@time.Duration?"},
	"@time.Period":                     {Name: "@time.Period"},
	"@time.Period!Error":               {Name: "@time.Period!Error"},
	"@time.PlainDate":                  {Name: "@time.PlainDate"},
	"@time.PlainDate!Error":            {Name: "@time.PlainDate!Error"},
	"@time.PlainDateTime":              {Name: "@time.PlainDateTime"},
	"@time.PlainDateTime!Error":        {Name: "@time.PlainDateTime!Error"},
	"@time.PlainTime":                  {Name: "@time.PlainTime"},
	"@time.PlainTime!Error":            {Name: "@time.PlainTime!Error"},
	"@time.Weekday":                    {Name: "@time.Weekday"},
	"@time.Zone":                       {Name: "@time.Zone"},
	"@time.Zone!Error":                 {Name: "@time.Zone!Error"},
	"@time.ZoneOffset":                 {Name: "@time.ZoneOffset"},
	"@time.ZoneOffset!Error":           {Name: "@time.ZoneOffset!Error"},
	"@time.ZonedDateTime":              {Name: "@time.ZonedDateTime"},
	"@time.ZonedDateTime!Error":        {Name: "@time.ZonedDateTime!Error"},
	"@time.ZonedDateTime?":             {Name: "@time.ZonedDateTime?"},
	"Array[@testutils.T]":              {Name: "Array[@testutils.T]"},
	"Array[Array[String]?]":            {Name: "Array[Array[String]?]"},
	"Array[Array[String]]":             {Name: "Array[Array[String]]"},
	"Array[Array[String]]?":            {Name: "Array[Array[String]]?"},
	"Array[Bool?]":                     {Name: "Array[Bool?]"},
	"Array[Bool]":                      {Name: "Array[Bool]"},
	"Array[Byte?]":                     {Name: "Array[Byte?]"},
	"Array[Byte]":                      {Name: "Array[Byte]"},
	"Array[Char?]":                     {Name: "Array[Char?]"},
	"Array[Char]":                      {Name: "Array[Char]"},
	"Array[Double?]":                   {Name: "Array[Double?]"},
	"Array[Double]":                    {Name: "Array[Double]"},
	"Array[Float?]":                    {Name: "Array[Float?]"},
	"Array[Float]":                     {Name: "Array[Float]"},
	"Array[HttpHeader?]":               {Name: "Array[HttpHeader?]"},
	"Array[Int16?]":                    {Name: "Array[Int16?]"},
	"Array[Int16]":                     {Name: "Array[Int16]"},
	"Array[Int64?]":                    {Name: "Array[Int64?]"},
	"Array[Int64]":                     {Name: "Array[Int64]"},
	"Array[Int?]":                      {Name: "Array[Int?]"},
	"Array[Int]":                       {Name: "Array[Int]"},
	"Array[String?]":                   {Name: "Array[String?]"},
	"Array[String?]?":                  {Name: "Array[String?]?"},
	"Array[String]":                    {Name: "Array[String]"},
	"Array[String]?":                   {Name: "Array[String]?"},
	"Array[UInt16?]":                   {Name: "Array[UInt16?]"},
	"Array[UInt16]":                    {Name: "Array[UInt16]"},
	"Array[UInt64?]":                   {Name: "Array[UInt64?]"},
	"Array[UInt64]":                    {Name: "Array[UInt64]"},
	"Array[UInt?]":                     {Name: "Array[UInt?]"},
	"Array[UInt]":                      {Name: "Array[UInt]"},
	"Bool":                             {Name: "Bool"},
	"Bool?":                            {Name: "Bool?"},
	"Byte":                             {Name: "Byte"},
	"Byte?":                            {Name: "Byte?"},
	"Char":                             {Name: "Char"},
	"Char?":                            {Name: "Char?"},
	"Double":                           {Name: "Double"},
	"Double?":                          {Name: "Double?"},
	"FixedArray[Bool?]":                {Name: "FixedArray[Bool?]"},
	"FixedArray[Bool]":                 {Name: "FixedArray[Bool]"},
	"FixedArray[Byte?]":                {Name: "FixedArray[Byte?]"},
	"FixedArray[Byte]":                 {Name: "FixedArray[Byte]"},
	"FixedArray[Char?]":                {Name: "FixedArray[Char?]"},
	"FixedArray[Char]":                 {Name: "FixedArray[Char]"},
	"FixedArray[Double?]":              {Name: "FixedArray[Double?]"},
	"FixedArray[Double]":               {Name: "FixedArray[Double]"},
	"FixedArray[Float?]":               {Name: "FixedArray[Float?]"},
	"FixedArray[Float]":                {Name: "FixedArray[Float]"},
	"FixedArray[Int16?]":               {Name: "FixedArray[Int16?]"},
	"FixedArray[Int16]":                {Name: "FixedArray[Int16]"},
	"FixedArray[Int64?]":               {Name: "FixedArray[Int64?]"},
	"FixedArray[Int64]":                {Name: "FixedArray[Int64]"},
	"FixedArray[Int?]":                 {Name: "FixedArray[Int?]"},
	"FixedArray[Int]":                  {Name: "FixedArray[Int]"},
	"FixedArray[Int]?":                 {Name: "FixedArray[Int]?"},
	"FixedArray[Map[String, String]?]": {Name: "FixedArray[Map[String, String]?]"},
	"FixedArray[Map[String, String]]":  {Name: "FixedArray[Map[String, String]]"},
	"FixedArray[String?]":              {Name: "FixedArray[String?]"},
	"FixedArray[String]":               {Name: "FixedArray[String]"},
	"FixedArray[String]?":              {Name: "FixedArray[String]?"},
	"FixedArray[TestStruct2?]":         {Name: "FixedArray[TestStruct2?]"},
	"FixedArray[TestStruct2]":          {Name: "FixedArray[TestStruct2]"},
	"FixedArray[UInt16?]":              {Name: "FixedArray[UInt16?]"},
	"FixedArray[UInt16]":               {Name: "FixedArray[UInt16]"},
	"FixedArray[UInt64?]":              {Name: "FixedArray[UInt64?]"},
	"FixedArray[UInt64]":               {Name: "FixedArray[UInt64]"},
	"FixedArray[UInt?]":                {Name: "FixedArray[UInt?]"},
	"FixedArray[UInt]":                 {Name: "FixedArray[UInt]"},
	"Float":                            {Name: "Float"},
	"Float?":                           {Name: "Float?"},
	"HttpHeader": {
		Name:   "HttpHeader",
		Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "values", Type: "Array[String]"}},
	},
	"HttpHeader?": {
		Name:   "HttpHeader?",
		Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "values", Type: "Array[String]"}},
	},
	"HttpHeaders": {
		Name:   "HttpHeaders",
		Fields: []*metadata.Field{{Name: "data", Type: "Map[String, HttpHeader?]"}},
	},
	"HttpHeaders?": {
		Name:   "HttpHeaders?",
		Fields: []*metadata.Field{{Name: "data", Type: "Map[String, HttpHeader?]"}},
	},
	"HttpResponse": {
		Name: "HttpResponse",
		Fields: []*metadata.Field{
			{Name: "status", Type: "UInt16"}, {Name: "statusText", Type: "String"},
			{Name: "headers", Type: "HttpHeaders?"}, {Name: "body", Type: "Array[Byte]"},
		},
	},
	"HttpResponse?": {
		Name: "HttpResponse?",
		Fields: []*metadata.Field{
			{Name: "status", Type: "UInt16"}, {Name: "statusText", Type: "String"},
			{Name: "headers", Type: "HttpHeaders?"}, {Name: "body", Type: "Array[Byte]"},
		},
	},
	"Int":                      {Name: "Int"},
	"Int16":                    {Name: "Int16"},
	"Int16?":                   {Name: "Int16?"},
	"Int64":                    {Name: "Int64"},
	"Int64?":                   {Name: "Int64?"},
	"Int?":                     {Name: "Int?"},
	"Map[Int, Double]":         {Name: "Map[Int, Double]"},
	"Map[Int, Float]":          {Name: "Map[Int, Float]"},
	"Map[String, HttpHeader?]": {Name: "Map[String, HttpHeader?]"},
	"Map[String, String]":      {Name: "Map[String, String]"},
	"Map[String, String]?":     {Name: "Map[String, String]?"},
	"String":                   {Name: "String"},
	"String?":                  {Name: "String?"},
	"TestRecursiveStruct": {
		Name:   "TestRecursiveStruct",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "mut b", Type: "TestRecursiveStruct?"}},
	},
	"TestRecursiveStruct?": {
		Name:   "TestRecursiveStruct?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "mut b", Type: "TestRecursiveStruct?"}},
	},
	"TestRecursiveStruct_map": {
		Name:   "TestRecursiveStruct_map",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "mut b", Type: "TestRecursiveStruct_map?"}},
	},
	"TestRecursiveStruct_map?": {
		Name:   "TestRecursiveStruct_map?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "mut b", Type: "TestRecursiveStruct_map?"}},
	},
	"TestSmorgasbordStruct": {
		Name: "TestSmorgasbordStruct",
		Fields: []*metadata.Field{
			{Name: "bool", Type: "Bool"},
			{Name: "byte", Type: "Byte"},
			{Name: "c", Type: "Char"},
			{Name: "f", Type: "Float"},
			{Name: "d", Type: "Double"},
			{Name: "i16", Type: "Int16"},
			{Name: "i32", Type: "Int"},
			{Name: "i64", Type: "Int64"},
			{Name: "s", Type: "String"},
			{Name: "u16", Type: "UInt16"},
			{Name: "u32", Type: "UInt"},
			{Name: "u64", Type: "UInt64"},
			{Name: "someBool", Type: "Bool?"},
			{Name: "noneBool", Type: "Bool?"},
			{Name: "someByte", Type: "Byte?"},
			{Name: "noneByte", Type: "Byte?"},
			{Name: "someChar", Type: "Char?"},
			{Name: "noneChar", Type: "Char?"},
			{Name: "someFloat", Type: "Float?"},
			{Name: "noneFloat", Type: "Float?"},
			{Name: "someDouble", Type: "Double?"},
			{Name: "noneDouble", Type: "Double?"},
			{Name: "someI16", Type: "Int16?"},
			{Name: "noneI16", Type: "Int16?"},
			{Name: "someI32", Type: "Int?"},
			{Name: "noneI32", Type: "Int?"},
			{Name: "someI64", Type: "Int64?"},
			{Name: "noneI64", Type: "Int64?"},
			{Name: "someString", Type: "String?"},
			{Name: "noneString", Type: "String?"},
			{Name: "someU16", Type: "UInt16?"},
			{Name: "noneU16", Type: "UInt16?"},
			{Name: "someU32", Type: "UInt?"},
			{Name: "noneU32", Type: "UInt?"},
			{Name: "someU64", Type: "UInt64?"},
			{Name: "noneU64", Type: "UInt64?"},
		},
	},
	"TestSmorgasbordStruct?": {
		Name: "TestSmorgasbordStruct?",
		Fields: []*metadata.Field{
			{Name: "bool", Type: "Bool"},
			{Name: "byte", Type: "Byte"},
			{Name: "c", Type: "Char"},
			{Name: "f", Type: "Float"},
			{Name: "d", Type: "Double"},
			{Name: "i16", Type: "Int16"},
			{Name: "i32", Type: "Int"},
			{Name: "i64", Type: "Int64"},
			{Name: "s", Type: "String"},
			{Name: "u16", Type: "UInt16"},
			{Name: "u32", Type: "UInt"},
			{Name: "u64", Type: "UInt64"},
			{Name: "someBool", Type: "Bool?"},
			{Name: "noneBool", Type: "Bool?"},
			{Name: "someByte", Type: "Byte?"},
			{Name: "noneByte", Type: "Byte?"},
			{Name: "someChar", Type: "Char?"},
			{Name: "noneChar", Type: "Char?"},
			{Name: "someFloat", Type: "Float?"},
			{Name: "noneFloat", Type: "Float?"},
			{Name: "someDouble", Type: "Double?"},
			{Name: "noneDouble", Type: "Double?"},
			{Name: "someI16", Type: "Int16?"},
			{Name: "noneI16", Type: "Int16?"},
			{Name: "someI32", Type: "Int?"},
			{Name: "noneI32", Type: "Int?"},
			{Name: "someI64", Type: "Int64?"},
			{Name: "noneI64", Type: "Int64?"},
			{Name: "someString", Type: "String?"},
			{Name: "noneString", Type: "String?"},
			{Name: "someU16", Type: "UInt16?"},
			{Name: "noneU16", Type: "UInt16?"},
			{Name: "someU32", Type: "UInt?"},
			{Name: "noneU32", Type: "UInt?"},
			{Name: "someU64", Type: "UInt64?"},
			{Name: "noneU64", Type: "UInt64?"},
		},
	},
	"TestSmorgasbordStruct_map": {
		Name: "TestSmorgasbordStruct_map",
		Fields: []*metadata.Field{
			{Name: "bool", Type: "Bool"},
			{Name: "byte", Type: "Byte"},
			{Name: "c", Type: "Char"},
			{Name: "f", Type: "Float"},
			{Name: "d", Type: "Double"},
			{Name: "i16", Type: "Int16"},
			{Name: "i32", Type: "Int"},
			{Name: "i64", Type: "Int64"},
			{Name: "s", Type: "String"},
			{Name: "u16", Type: "UInt16"},
			{Name: "u32", Type: "UInt"},
			{Name: "u64", Type: "UInt64"},
			{Name: "someBool", Type: "Bool?"},
			{Name: "noneBool", Type: "Bool?"},
			{Name: "someByte", Type: "Byte?"},
			{Name: "noneByte", Type: "Byte?"},
			{Name: "someChar", Type: "Char?"},
			{Name: "noneChar", Type: "Char?"},
			{Name: "someFloat", Type: "Float?"},
			{Name: "noneFloat", Type: "Float?"},
			{Name: "someDouble", Type: "Double?"},
			{Name: "noneDouble", Type: "Double?"},
			{Name: "someI16", Type: "Int16?"},
			{Name: "noneI16", Type: "Int16?"},
			{Name: "someI32", Type: "Int?"},
			{Name: "noneI32", Type: "Int?"},
			{Name: "someI64", Type: "Int64?"},
			{Name: "noneI64", Type: "Int64?"},
			{Name: "someString", Type: "String?"},
			{Name: "noneString", Type: "String?"},
			{Name: "someU16", Type: "UInt16?"},
			{Name: "noneU16", Type: "UInt16?"},
			{Name: "someU32", Type: "UInt?"},
			{Name: "noneU32", Type: "UInt?"},
			{Name: "someU64", Type: "UInt64?"},
			{Name: "noneU64", Type: "UInt64?"},
		},
	},
	"TestSmorgasbordStruct_map?": {
		Name: "TestSmorgasbordStruct_map?",
		Fields: []*metadata.Field{
			{Name: "bool", Type: "Bool"},
			{Name: "byte", Type: "Byte"},
			{Name: "c", Type: "Char"},
			{Name: "f", Type: "Float"},
			{Name: "d", Type: "Double"},
			{Name: "i16", Type: "Int16"},
			{Name: "i32", Type: "Int"},
			{Name: "i64", Type: "Int64"},
			{Name: "s", Type: "String"},
			{Name: "u16", Type: "UInt16"},
			{Name: "u32", Type: "UInt"},
			{Name: "u64", Type: "UInt64"},
			{Name: "someBool", Type: "Bool?"},
			{Name: "noneBool", Type: "Bool?"},
			{Name: "someByte", Type: "Byte?"},
			{Name: "noneByte", Type: "Byte?"},
			{Name: "someChar", Type: "Char?"},
			{Name: "noneChar", Type: "Char?"},
			{Name: "someFloat", Type: "Float?"},
			{Name: "noneFloat", Type: "Float?"},
			{Name: "someDouble", Type: "Double?"},
			{Name: "noneDouble", Type: "Double?"},
			{Name: "someI16", Type: "Int16?"},
			{Name: "noneI16", Type: "Int16?"},
			{Name: "someI32", Type: "Int?"},
			{Name: "noneI32", Type: "Int?"},
			{Name: "someI64", Type: "Int64?"},
			{Name: "noneI64", Type: "Int64?"},
			{Name: "someString", Type: "String?"},
			{Name: "noneString", Type: "String?"},
			{Name: "someU16", Type: "UInt16?"},
			{Name: "noneU16", Type: "UInt16?"},
			{Name: "someU32", Type: "UInt?"},
			{Name: "noneU32", Type: "UInt?"},
			{Name: "someU64", Type: "UInt64?"},
			{Name: "noneU64", Type: "UInt64?"},
		},
	},
	"TestStruct1": {
		Name:   "TestStruct1",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}},
	},
	"TestStruct1?": {
		Name:   "TestStruct1?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}},
	},
	"TestStruct1_map": {
		Name:   "TestStruct1_map",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}},
	},
	"TestStruct1_map?": {
		Name:   "TestStruct1_map?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}},
	},
	"TestStruct2": {
		Name:   "TestStruct2",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"}},
	},
	"TestStruct2?": {
		Name:   "TestStruct2?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"}},
	},
	"TestStruct2_map": {
		Name:   "TestStruct2_map",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"}},
	},
	"TestStruct2_map?": {
		Name:   "TestStruct2_map?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"}},
	},
	"TestStruct3": {
		Name: "TestStruct3",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String"},
		},
	},
	"TestStruct3?": {
		Name: "TestStruct3?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String"},
		},
	},
	"TestStruct3_map": {
		Name: "TestStruct3_map",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String"},
		},
	},
	"TestStruct3_map?": {
		Name: "TestStruct3_map?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String"},
		},
	},
	"TestStruct4": {
		Name: "TestStruct4",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String?"},
		},
	},
	"TestStruct4?": {
		Name: "TestStruct4?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String?"},
		},
	},
	"TestStruct4_map": {
		Name: "TestStruct4_map",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String?"},
		},
	},
	"TestStruct4_map?": {
		Name: "TestStruct4_map?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String?"},
		},
	},
	"TestStruct5": {
		Name: "TestStruct5",
		Fields: []*metadata.Field{
			{Name: "a", Type: "String"}, {Name: "b", Type: "String"},
			{Name: "c", Type: "String"}, {Name: "d", Type: "Array[String]"},
			{Name: "e", Type: "Double"}, {Name: "f", Type: "Double"},
		},
	},
	"TestStruct5?": {
		Name: "TestStruct5?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "String"}, {Name: "b", Type: "String"},
			{Name: "c", Type: "String"}, {Name: "d", Type: "Array[String]"},
			{Name: "e", Type: "Double"}, {Name: "f", Type: "Double"},
		},
	},
	"TestStruct5_map": {
		Name: "TestStruct5_map",
		Fields: []*metadata.Field{
			{Name: "a", Type: "String"}, {Name: "b", Type: "String"},
			{Name: "c", Type: "String"}, {Name: "d", Type: "Array[String]"},
			{Name: "e", Type: "Double"}, {Name: "f", Type: "Double"},
		},
	},
	"TestStruct5_map?": {
		Name: "TestStruct5_map?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "String"}, {Name: "b", Type: "String"},
			{Name: "c", Type: "String"}, {Name: "d", Type: "Array[String]"},
			{Name: "e", Type: "Double"}, {Name: "f", Type: "Double"},
		},
	},
	"TestStructWithMap": {
		Name:   "TestStructWithMap",
		Fields: []*metadata.Field{{Name: "m", Type: "Map[String, String]"}},
	},
	"UInt":       {Name: "UInt"},
	"UInt16":     {Name: "UInt16"},
	"UInt16?":    {Name: "UInt16?"},
	"UInt64":     {Name: "UInt64"},
	"UInt64?":    {Name: "UInt64?"},
	"UInt?":      {Name: "UInt?"},
	"Unit":       {Name: "Unit"},
	"Unit!Error": {Name: "Unit!Error"},
}
