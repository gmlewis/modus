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

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-version"
)

func TestGenerateMetadata_Runtime(t *testing.T) {
	config := &config.Config{
		SourceDir: "testdata/runtime-testdata",
	}
	mod := &modinfo.ModuleInfo{
		ModulePath:      "github.com/gmlewis/modus/runtime/testdata",
		ModusSDKVersion: version.Must(version.NewVersion("40.11.0")),
	}

	meta, err := GenerateMetadata(config, mod)
	if err != nil {
		t.Fatalf("GenerateMetadata returned an error: %v", err)
	}

	if got, want := meta.Plugin, "testdata"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@runtime-testdata"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if got, want := meta.SDK, "modus-sdk-mbt@40.11.0"; got != want {
		t.Errorf("meta.SDK = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantRuntimeFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantRuntimeFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantRuntimeTypes, meta.Types); diff != "" {
		t.Errorf("meta.Types mismatch (-want +got):\n%v", diff)
	}

	// This call makes it easy to step through the code with a debugger:
	// LogToConsole(meta)
}

var wantRuntimeFnExports = metadata.FunctionMap{
	"add": {
		Name:       "add",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"echo1": {
		Name:       "echo1",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"echo2": {
		Name:       "echo2",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String?"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"echo3": {
		Name:       "echo3",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String?"}},
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
	"getIntPtrArray2": {
		Name:    "getIntPtrArray2",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
	},
	"getMapPtrArray2": {
		Name:    "getMapPtrArray2",
		Results: []*metadata.Result{{Type: "FixedArray[Map[String, String]?]"}},
	},
	"get_int_option_fixedarray1": {
		Name:    "get_int_option_fixedarray1",
		Results: []*metadata.Result{{Type: "FixedArray[Int?]"}},
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
	"test2d_array_input_string_inner_none": {
		Name:       "test2d_array_input_string_inner_none",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Array[String]]?"}},
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
	"test2d_array_output_string_inner_none": {
		Name:    "test2d_array_output_string_inner_none",
		Results: []*metadata.Result{{Type: "Array[Array[String]]"}},
	},
	"test2d_array_output_string_none": {
		Name:    "test2d_array_output_string_none",
		Results: []*metadata.Result{{Type: "Array[Array[String]]?"}},
	},
	"test_array_input_byte": {
		Name:       "test_array_input_byte",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[Byte]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_input_int32_empty": {
		Name:       "test_array_input_int32_empty",
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
	"test_array_output_byte": {
		Name:    "test_array_output_byte",
		Results: []*metadata.Result{{Type: "Array[Byte]"}},
	},
	"test_array_output_int32_empty": {
		Name:    "test_array_output_int32_empty",
		Results: []*metadata.Result{{Type: "Array[Int]"}},
	},
	"test_array_output_int_option": {
		Name:    "test_array_output_int_option",
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
	"test_bool_option_input_nil": {
		Name:       "test_bool_option_input_nil",
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
	"test_bool_option_output_nil": {
		Name:    "test_bool_option_output_nil",
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
	"test_byte_option_input_nil": {
		Name:       "test_byte_option_input_nil",
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
	"test_byte_option_output_nil": {
		Name:    "test_byte_option_output_nil",
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
	"test_char_option_input_nil": {
		Name:       "test_char_option_input_nil",
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
	"test_char_option_output_nil": {
		Name:    "test_char_option_output_nil",
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
	"test_double_option_input_nil": {
		Name:       "test_double_option_input_nil",
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
	"test_double_option_output_nil": {
		Name:    "test_double_option_output_nil",
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
	"test_duration_option_input_nil": {
		Name:       "test_duration_option_input_nil",
		Parameters: []*metadata.Parameter{{Name: "d", Type: "@time.Duration?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_duration_option_output": {
		Name:    "test_duration_option_output",
		Results: []*metadata.Result{{Type: "@time.Duration?"}},
	},
	"test_duration_option_output_nil": {
		Name:    "test_duration_option_output_nil",
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
	"test_float32_input_max": {
		Name:       "test_float32_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Float"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_float32_input_min": {
		Name:       "test_float32_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Float"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_float32_option_input_max": {
		Name:       "test_float32_option_input_max",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Float?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_float32_option_input_min": {
		Name:       "test_float32_option_input_min",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Float?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_float32_option_input_nil": {
		Name:       "test_float32_option_input_nil",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Float?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_float32_option_output_max": {
		Name:    "test_float32_option_output_max",
		Results: []*metadata.Result{{Type: "Float?"}},
	},
	"test_float32_option_output_min": {
		Name:    "test_float32_option_output_min",
		Results: []*metadata.Result{{Type: "Float?"}},
	},
	"test_float32_option_output_nil": {
		Name:    "test_float32_option_output_nil",
		Results: []*metadata.Result{{Type: "Float?"}},
	},
	"test_float32_output_max": {Name: "test_float32_output_max", Results: []*metadata.Result{{Type: "Float"}}},
	"test_float32_output_min": {Name: "test_float32_output_min", Results: []*metadata.Result{{Type: "Float"}}},
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
	"test_int16_option_input_nil": {
		Name:       "test_int16_option_input_nil",
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
	"test_int16_option_output_nil": {
		Name:    "test_int16_option_output_nil",
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
	"test_int64_option_input_nil": {
		Name:       "test_int64_option_input_nil",
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
	"test_int64_option_output_nil": {
		Name:    "test_int64_option_output_nil",
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
	"test_int_option_input_nil": {
		Name:       "test_int_option_input_nil",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_int_option_output_max": {Name: "test_int_option_output_max", Results: []*metadata.Result{{Type: "Int?"}}},
	"test_int_option_output_min": {Name: "test_int_option_output_min", Results: []*metadata.Result{{Type: "Int?"}}},
	"test_int_option_output_nil": {Name: "test_int_option_output_nil", Results: []*metadata.Result{{Type: "Int?"}}},
	"test_int_output_max":        {Name: "test_int_output_max", Results: []*metadata.Result{{Type: "Int"}}},
	"test_int_output_min":        {Name: "test_int_output_min", Results: []*metadata.Result{{Type: "Int"}}},
	"test_iterate_map_string_string": {
		Name:       "test_iterate_map_string_string",
		Parameters: []*metadata.Parameter{{Name: "m", Type: "Map[String, String]"}},
	},
	"test_map_input_int_double": {
		Name:       "test_map_input_int_double",
		Parameters: []*metadata.Parameter{{Name: "m", Type: "Map[Int, Double]"}},
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
		Parameters: []*metadata.Parameter{{Name: "o", Type: "TestRecursiveStruct"}},
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
	"test_time_option_input_nil": {
		Name:       "test_time_option_input_nil",
		Parameters: []*metadata.Parameter{{Name: "t", Type: "@time.ZonedDateTime?"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_time_option_output": {
		Name:    "test_time_option_output",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime?"}},
	},
	"test_time_option_output_nil": {
		Name:    "test_time_option_output_nil",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime?"}},
	},
	"test_time_output": {
		Name:    "test_time_output",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime"}},
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
	"test_uint16_option_input_nil": {
		Name:       "test_uint16_option_input_nil",
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
	"test_uint16_option_output_nil": {
		Name:    "test_uint16_option_output_nil",
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
	"test_uint64_option_input_nil": {
		Name:       "test_uint64_option_input_nil",
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
	"test_uint64_option_output_nil": {
		Name:    "test_uint64_option_output_nil",
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
	"test_uint_option_input_nil": {
		Name:       "test_uint_option_input_nil",
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
	"test_uint_option_output_nil": {
		Name:    "test_uint_option_output_nil",
		Results: []*metadata.Result{{Type: "UInt?"}},
	},
	"test_uint_output_max": {Name: "test_uint_output_max", Results: []*metadata.Result{{Type: "UInt"}}},
	"test_uint_output_min": {Name: "test_uint_output_min", Results: []*metadata.Result{{Type: "UInt"}}},
}

var wantRuntimeFnImports = metadata.FunctionMap{
	"modus_system.getTimeInZone": {
		Name:       "modus_system.getTimeInZone",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
	},
	"modus_system.getTimeZoneData": {
		Name: "modus_system.getTimeZoneData",
		Parameters: []*metadata.Parameter{
			{Name: "tz", Type: "String"},
			{Name: "format", Type: "String"},
		},
		Results: []*metadata.Result{{Type: "Array[Byte]"}},
	},
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
		Parameters: []*metadata.Parameter{{Name: "message", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"modus_test.echo2": {
		Name:       "modus_test.echo2",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"modus_test.echo3": {
		Name:       "modus_test.echo3",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"modus_test.echo4": {
		Name:       "modus_test.echo4",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"modus_test.encodeStrings1": {
		Name:       "modus_test.encodeStrings1",
		Parameters: []*metadata.Parameter{{Name: "items", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"modus_test.encodeStrings2": {
		Name:       "modus_test.encodeStrings2",
		Parameters: []*metadata.Parameter{{Name: "items", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
}

var wantRuntimeTypes = metadata.TypeMap{
	"@time.Duration":                   {Id: 4, Name: "@time.Duration"},
	"@time.Duration?":                  {Id: 5, Name: "@time.Duration?"},
	"@time.ZonedDateTime":              {Id: 6, Name: "@time.ZonedDateTime", Fields: nil, Docs: nil},
	"@time.ZonedDateTime?":             {Id: 7, Name: "@time.ZonedDateTime?"},
	"Array[Array[String]]":             {Id: 8, Name: "Array[Array[String]]"},
	"Array[Array[String]]?":            {Id: 9, Name: "Array[Array[String]]?"},
	"Array[Byte]":                      {Id: 10, Name: "Array[Byte]", Fields: nil, Docs: nil},
	"Array[Double]":                    {Id: 11, Name: "Array[Double]"},
	"Array[HttpHeader?]":               {Id: 12, Name: "Array[HttpHeader?]"},
	"Array[Int?]":                      {Id: 13, Name: "Array[Int?]"},
	"Array[Int]":                       {Id: 14, Name: "Array[Int]"},
	"Array[String?]":                   {Id: 15, Name: "Array[String?]"},
	"Array[String?]?":                  {Id: 16, Name: "Array[String?]?"},
	"Array[String]":                    {Id: 17, Name: "Array[String]"},
	"Array[String]?":                   {Id: 18, Name: "Array[String]?"},
	"Bool":                             {Id: 19, Name: "Bool"},
	"Bool?":                            {Id: 20, Name: "Bool?"},
	"Byte":                             {Id: 21, Name: "Byte"},
	"Byte?":                            {Id: 22, Name: "Byte?"},
	"Char":                             {Id: 23, Name: "Char"},
	"Char?":                            {Id: 24, Name: "Char?"},
	"Double":                           {Id: 25, Name: "Double"},
	"Double?":                          {Id: 26, Name: "Double?"},
	"FixedArray[Byte]":                 {Id: 27, Name: "FixedArray[Byte]"},
	"FixedArray[Int?]":                 {Id: 28, Name: "FixedArray[Int?]"},
	"FixedArray[Int]?":                 {Id: 29, Name: "FixedArray[Int]?"},
	"FixedArray[Map[String, String]?]": {Id: 30, Name: "FixedArray[Map[String, String]?]"},
	"FixedArray[Map[String, String]]":  {Id: 31, Name: "FixedArray[Map[String, String]]"},
	"FixedArray[String?]":              {Id: 32, Name: "FixedArray[String?]"},
	"FixedArray[String]":               {Id: 33, Name: "FixedArray[String]"},
	"FixedArray[String]?":              {Id: 34, Name: "FixedArray[String]?"},
	"FixedArray[TestStruct2?]":         {Id: 35, Name: "FixedArray[TestStruct2?]"},
	"FixedArray[TestStruct2]":          {Id: 36, Name: "FixedArray[TestStruct2]"},
	"Float":                            {Id: 37, Name: "Float"},
	"Float?":                           {Id: 38, Name: "Float?"},
	"HttpHeader?": {Id: 39,
		Name:   "HttpHeader?",
		Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "values", Type: "Array[String]"}},
	},
	"HttpHeaders": {Id: 40,
		Name:   "HttpHeaders",
		Fields: []*metadata.Field{{Name: "data", Type: "Map[String, HttpHeader?]"}},
	},
	"HttpHeaders?": {Id: 41,
		Name:   "HttpHeaders?",
		Fields: []*metadata.Field{{Name: "data", Type: "Map[String, HttpHeader?]"}},
	},
	"HttpResponse?": {Id: 42,
		Name: "HttpResponse?",
		Fields: []*metadata.Field{
			{Name: "status", Type: "UInt16"}, {Name: "status_text", Type: "String"},
			{Name: "headers", Type: "HttpHeaders?"}, {Name: "body", Type: "Array[Byte]"},
		},
	},
	"Int":                      {Id: 43, Name: "Int"},
	"Int16":                    {Id: 44, Name: "Int16"},
	"Int16?":                   {Id: 45, Name: "Int16?"},
	"Int64":                    {Id: 46, Name: "Int64"},
	"Int64?":                   {Id: 47, Name: "Int64?"},
	"Int?":                     {Id: 48, Name: "Int?"},
	"Map[Int, Double]":         {Id: 49, Name: "Map[Int, Double]"},
	"Map[String, HttpHeader?]": {Id: 50, Name: "Map[String, HttpHeader?]"},
	"Map[String, String]":      {Id: 51, Name: "Map[String, String]"},
	"Map[String, String]?":     {Id: 52, Name: "Map[String, String]?"},
	"String?":                  {Id: 53, Name: "String?"},
	"TestRecursiveStruct": {Id: 54,
		Name:   "TestRecursiveStruct",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "mut b", Type: "TestRecursiveStruct?"}},
	},
	"TestRecursiveStruct?": {Id: 55,
		Name:   "TestRecursiveStruct?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "mut b", Type: "TestRecursiveStruct?"}},
	},
	"TestRecursiveStruct_map": {Id: 56,
		Name:   "TestRecursiveStruct_map",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "mut b", Type: "TestRecursiveStruct_map?"}},
	},
	"TestRecursiveStruct_map?": {Id: 57,
		Name:   "TestRecursiveStruct_map?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "mut b", Type: "TestRecursiveStruct_map?"}},
	},
	"TestStruct1": {Id: 58,
		Name:   "TestStruct1",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}},
	},
	"TestStruct1?": {Id: 59,
		Name:   "TestStruct1?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}},
	},
	"TestStruct1_map": {Id: 60,
		Name:   "TestStruct1_map",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}},
	},
	"TestStruct1_map?": {Id: 61,
		Name:   "TestStruct1_map?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}},
	},
	"TestStruct2": {Id: 62,
		Name:   "TestStruct2",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"}},
	},
	"TestStruct2?": {Id: 63,
		Name:   "TestStruct2?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"}},
	},
	"TestStruct2_map": {Id: 64,
		Name:   "TestStruct2_map",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"}},
	},
	"TestStruct2_map?": {Id: 65,
		Name:   "TestStruct2_map?",
		Fields: []*metadata.Field{{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"}},
	},
	"TestStruct3": {Id: 66,
		Name: "TestStruct3",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String"},
		},
	},
	"TestStruct3?": {Id: 67,
		Name: "TestStruct3?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String"},
		},
	},
	"TestStruct3_map": {Id: 68,
		Name: "TestStruct3_map",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String"},
		},
	},
	"TestStruct3_map?": {Id: 69,
		Name: "TestStruct3_map?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String"},
		},
	},
	"TestStruct4": {Id: 70,
		Name: "TestStruct4",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String?"},
		},
	},
	"TestStruct4?": {Id: 71,
		Name: "TestStruct4?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String?"},
		},
	},
	"TestStruct4_map": {Id: 72,
		Name: "TestStruct4_map",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String?"},
		},
	},
	"TestStruct4_map?": {Id: 73,
		Name: "TestStruct4_map?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "Bool"}, {Name: "b", Type: "Int"},
			{Name: "c", Type: "String?"},
		},
	},
	"TestStruct5": {Id: 74,
		Name: "TestStruct5",
		Fields: []*metadata.Field{
			{Name: "a", Type: "String"}, {Name: "b", Type: "String"},
			{Name: "c", Type: "String"}, {Name: "d", Type: "Array[String]"},
			{Name: "e", Type: "Double"}, {Name: "f", Type: "Double"},
		},
	},
	"TestStruct5?": {Id: 75,
		Name: "TestStruct5?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "String"}, {Name: "b", Type: "String"},
			{Name: "c", Type: "String"}, {Name: "d", Type: "Array[String]"},
			{Name: "e", Type: "Double"}, {Name: "f", Type: "Double"},
		},
	},
	"TestStruct5_map": {Id: 76,
		Name: "TestStruct5_map",
		Fields: []*metadata.Field{
			{Name: "a", Type: "String"}, {Name: "b", Type: "String"},
			{Name: "c", Type: "String"}, {Name: "d", Type: "Array[String]"},
			{Name: "e", Type: "Double"}, {Name: "f", Type: "Double"},
		},
	},
	"TestStruct5_map?": {Id: 77,
		Name: "TestStruct5_map?",
		Fields: []*metadata.Field{
			{Name: "a", Type: "String"}, {Name: "b", Type: "String"},
			{Name: "c", Type: "String"}, {Name: "d", Type: "Array[String]"},
			{Name: "e", Type: "Double"}, {Name: "f", Type: "Double"},
		},
	},
	"TestStructWithMap": {Id: 78,
		Name:   "TestStructWithMap",
		Fields: []*metadata.Field{{Name: "m", Type: "Map[String, String]"}},
	},
	"UInt":       {Id: 79, Name: "UInt"},
	"UInt16":     {Id: 80, Name: "UInt16"},
	"UInt16?":    {Id: 81, Name: "UInt16?"},
	"UInt64":     {Id: 82, Name: "UInt64"},
	"UInt64?":    {Id: 83, Name: "UInt64?"},
	"UInt?":      {Id: 84, Name: "UInt?"},
	"Unit":       {Id: 85, Name: "Unit"},
	"Unit!Error": {Id: 86, Name: "Unit!Error"},
}
