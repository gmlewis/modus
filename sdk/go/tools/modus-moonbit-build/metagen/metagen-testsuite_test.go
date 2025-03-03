// -*- compile-command: "go test -run ^TestGenerateMetadata_Testsuite$ ."; -*-

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
	"log"
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/modinfo"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/go-version"
)

func TestGenerateMetadata_Testsuite(t *testing.T) {
	log.SetFlags(0)
	config := &config.Config{
		SourceDir: "testdata/test-suite",
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

	if got, want := meta.Module, "@test-suite"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if got, want := meta.SDK, "modus-sdk-mbt@40.11.0"; got != want {
		t.Errorf("meta.SDK = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantTestsuiteFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantTestsuiteFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantTestsuiteTypes, meta.Types); diff != "" {
		t.Errorf("meta.Types mismatch (-want +got):\n%v", diff)
	}

	// This call makes it easy to step through the code with a debugger:
	// LogToConsole(meta)
}

var wantTestsuiteFnExports = metadata.FunctionMap{
	"__modus_add": {
		Name:       "__modus_add",
		Parameters: []*metadata.Parameter{{Name: "x", Type: "Int"}, {Name: "y", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs:       &metadata.Docs{Lines: []string{"Adds two integers together and returns the result."}},
	},
	"__modus_add3": {
		Name:       "__modus_add3",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}, {Name: "c", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs: &metadata.Docs{
			Lines: []string{"Adds three integers together and returns the result.", "The third integer is optional."},
		},
	},
	"__modus_add3_WithDefaults": {
		Name:       "__modus_add3_WithDefaults",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs: &metadata.Docs{
			Lines: []string{"Adds three integers together and returns the result.", "The third integer is optional."},
		},
	},
	"__modus_get_current_time": {
		Name:    "__modus_get_current_time",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time."}},
	},
	"__modus_get_current_time_formatted": {
		Name:    "__modus_get_current_time_formatted",
		Results: []*metadata.Result{{Type: "String!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time formatted as a string."}},
	},
	"__modus_get_full_name": {
		Name:       "__modus_get_full_name",
		Parameters: []*metadata.Parameter{{Name: "first_name", Type: "String"}, {Name: "last_name", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Combines the first and last name of a person, and returns the full name."}},
	},
	"__modus_get_local_time": {
		Name:    "__modus_get_local_time",
		Results: []*metadata.Result{{Type: "String!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current local time."}},
	},
	"__modus_get_local_time_zone_id": {
		Name:    "__modus_get_local_time_zone_id",
		Results: []*metadata.Result{{Type: "String"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the local time zone identifier."}},
	},
	"__modus_get_people": {
		Name:    "__modus_get_people",
		Results: []*metadata.Result{{Type: "Array[Person]"}},
		Docs:    &metadata.Docs{Lines: []string{"Gets a list of people."}},
	},
	"__modus_get_person": {
		Name:    "__modus_get_person",
		Results: []*metadata.Result{{Type: "Person"}},
		Docs:    &metadata.Docs{Lines: []string{"Gets a person object."}},
	},
	"__modus_get_random_person": {
		Name:    "__modus_get_random_person",
		Results: []*metadata.Result{{Type: "Person"}},
		Docs:    &metadata.Docs{Lines: []string{"Gets a random person object from a list of people."}},
	},
	"__modus_get_time_in_zone": {
		Name:       "__modus_get_time_in_zone",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Returns the current time in a specified time zone."}},
	},
	"__modus_get_time_zone_info": {
		Name:       "__modus_get_time_zone_info",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "TimeZoneInfo!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Returns some basic information about the time zone specified."}},
	},
	"__modus_get_utc_time": {
		Name:    "__modus_get_utc_time",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time in UTC."}},
	},
	"__modus_hello_array_of_ints": {
		Name:       "__modus_hello_array_of_ints",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Array[Int]"}},
	},
	"__modus_hello_array_of_ints_WithDefaults": {
		Name:    "__modus_hello_array_of_ints_WithDefaults",
		Results: []*metadata.Result{{Type: "Array[Int]"}},
	},
	"__modus_hello_array_of_ints_option": {
		Name:       "__modus_hello_array_of_ints_option",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Array[Int]?"}},
	},
	"__modus_hello_array_of_ints_option_WithDefaults": {
		Name:    "__modus_hello_array_of_ints_option_WithDefaults",
		Results: []*metadata.Result{{Type: "Array[Int]?"}},
	},
	"__modus_hello_array_of_strings": {
		Name:       "__modus_hello_array_of_strings",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Array[String]"}},
	},
	"__modus_hello_array_of_strings_WithDefaults": {
		Name:    "__modus_hello_array_of_strings_WithDefaults",
		Results: []*metadata.Result{{Type: "Array[String]"}},
	},
	"__modus_hello_array_of_strings_option": {
		Name:       "__modus_hello_array_of_strings_option",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Array[String]?"}},
	},
	"__modus_hello_array_of_strings_option_WithDefaults": {
		Name:    "__modus_hello_array_of_strings_option_WithDefaults",
		Results: []*metadata.Result{{Type: "Array[String]?"}},
	},
	"__modus_hello_maps_n_items": {
		Name:       "__modus_hello_maps_n_items",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Map[String, String]"}},
	},
	"__modus_hello_maps_n_items_WithDefaults": {
		Name:    "__modus_hello_maps_n_items_WithDefaults",
		Results: []*metadata.Result{{Type: "Map[String, String]"}},
	},
	"__modus_hello_maps_n_items_option": {
		Name:       "__modus_hello_maps_n_items_option",
		Parameters: []*metadata.Parameter{{Name: "n", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Map[String, String]?"}},
	},
	"__modus_hello_maps_n_items_option_WithDefaults": {
		Name:    "__modus_hello_maps_n_items_option_WithDefaults",
		Results: []*metadata.Result{{Type: "Map[String, String]?"}},
	},
	"__modus_hello_option_empty_string": {
		Name:       "__modus_hello_option_empty_string",
		Parameters: []*metadata.Parameter{{Name: "some", Type: "Bool"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"__modus_hello_option_empty_string_WithDefaults": {
		Name:    "__modus_hello_option_empty_string_WithDefaults",
		Results: []*metadata.Result{{Type: "String?"}},
	},
	"__modus_hello_option_none": {
		Name:       "__modus_hello_option_none",
		Parameters: []*metadata.Parameter{{Name: "some", Type: "Bool"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"__modus_hello_option_none_WithDefaults": {
		Name:    "__modus_hello_option_none_WithDefaults",
		Results: []*metadata.Result{{Type: "String?"}},
	},
	"__modus_hello_option_some_string": {
		Name:       "__modus_hello_option_some_string",
		Parameters: []*metadata.Parameter{{Name: "some", Type: "Bool"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"__modus_hello_option_some_string_WithDefaults": {
		Name:    "__modus_hello_option_some_string_WithDefaults",
		Results: []*metadata.Result{{Type: "String?"}},
	},
	"__modus_hello_primitive_bool_max": {
		Name:    "__modus_hello_primitive_bool_max",
		Results: []*metadata.Result{{Type: "Bool"}},
	},
	"__modus_hello_primitive_bool_min": {
		Name:    "__modus_hello_primitive_bool_min",
		Results: []*metadata.Result{{Type: "Bool"}},
	},
	"__modus_hello_primitive_byte_max": {
		Name:    "__modus_hello_primitive_byte_max",
		Results: []*metadata.Result{{Type: "Byte"}},
	},
	"__modus_hello_primitive_byte_min": {
		Name:    "__modus_hello_primitive_byte_min",
		Results: []*metadata.Result{{Type: "Byte"}},
	},
	"__modus_hello_primitive_char_max": {
		Name:    "__modus_hello_primitive_char_max",
		Results: []*metadata.Result{{Type: "Char"}},
	},
	"__modus_hello_primitive_char_min": {
		Name:    "__modus_hello_primitive_char_min",
		Results: []*metadata.Result{{Type: "Char"}},
	},
	"__modus_hello_primitive_double_max": {
		Name:    "__modus_hello_primitive_double_max",
		Results: []*metadata.Result{{Type: "Double"}},
	},
	"__modus_hello_primitive_double_min": {
		Name:    "__modus_hello_primitive_double_min",
		Results: []*metadata.Result{{Type: "Double"}},
	},
	"__modus_hello_primitive_float_max": {
		Name:    "__modus_hello_primitive_float_max",
		Results: []*metadata.Result{{Type: "Float"}},
	},
	"__modus_hello_primitive_float_min": {
		Name:    "__modus_hello_primitive_float_min",
		Results: []*metadata.Result{{Type: "Float"}},
	},
	"__modus_hello_primitive_int16_max": {
		Name:    "__modus_hello_primitive_int16_max",
		Results: []*metadata.Result{{Type: "Int16"}},
	},
	"__modus_hello_primitive_int16_min": {
		Name:    "__modus_hello_primitive_int16_min",
		Results: []*metadata.Result{{Type: "Int16"}},
	},
	"__modus_hello_primitive_int64_max": {
		Name:    "__modus_hello_primitive_int64_max",
		Results: []*metadata.Result{{Type: "Int64"}},
	},
	"__modus_hello_primitive_int64_min": {
		Name:    "__modus_hello_primitive_int64_min",
		Results: []*metadata.Result{{Type: "Int64"}},
	},
	"__modus_hello_primitive_int_max": {
		Name:    "__modus_hello_primitive_int_max",
		Results: []*metadata.Result{{Type: "Int"}},
	},
	"__modus_hello_primitive_int_min": {
		Name:    "__modus_hello_primitive_int_min",
		Results: []*metadata.Result{{Type: "Int"}},
	},
	"__modus_hello_primitive_uint16_max": {
		Name:    "__modus_hello_primitive_uint16_max",
		Results: []*metadata.Result{{Type: "UInt16"}},
	},
	"__modus_hello_primitive_uint16_min": {
		Name:    "__modus_hello_primitive_uint16_min",
		Results: []*metadata.Result{{Type: "UInt16"}},
	},
	"__modus_hello_primitive_uint64_max": {
		Name:    "__modus_hello_primitive_uint64_max",
		Results: []*metadata.Result{{Type: "UInt64"}},
	},
	"__modus_hello_primitive_uint64_min": {
		Name:    "__modus_hello_primitive_uint64_min",
		Results: []*metadata.Result{{Type: "UInt64"}},
	},
	"__modus_hello_primitive_uint_max": {
		Name:    "__modus_hello_primitive_uint_max",
		Results: []*metadata.Result{{Type: "UInt"}},
	},
	"__modus_hello_primitive_uint_min": {
		Name:    "__modus_hello_primitive_uint_min",
		Results: []*metadata.Result{{Type: "UInt"}},
	},
	"__modus_hello_world": {Name: "__modus_hello_world", Results: []*metadata.Result{{Type: "String"}}},
	"__modus_hello_world_with_arg": {
		Name:       "__modus_hello_world_with_arg",
		Parameters: []*metadata.Parameter{{Name: "name", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"__modus_hello_world_with_optional_arg": {
		Name:       "__modus_hello_world_with_optional_arg",
		Parameters: []*metadata.Parameter{{Name: "name", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"__modus_hello_world_with_optional_arg_WithDefaults": {
		Name:    "__modus_hello_world_with_optional_arg_WithDefaults",
		Results: []*metadata.Result{{Type: "String"}},
	},
	"__modus_log_message": {
		Name:       "__modus_log_message",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Logs a message."}},
	},
	"__modus_test_abort": {
		Name: "__modus_test_abort",
		Docs: &metadata.Docs{Lines: []string{"Tests an abort."}},
	},
	"__modus_test_alternative_error": {
		Name:       "__modus_test_alternative_error",
		Parameters: []*metadata.Parameter{{Name: "input", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Tests an alternative way to handle errors in functions."}},
	},
	"__modus_test_exit": {
		Name: "__modus_test_exit",
		Docs: &metadata.Docs{Lines: []string{"Tests an exit with a non-zero exit code."}},
	},
	"__modus_test_logging": {
		Name: "__modus_test_logging",
		Docs: &metadata.Docs{Lines: []string{"Tests logging at different levels."}},
	},
	"__modus_test_normal_error": {
		Name:       "__modus_test_normal_error",
		Parameters: []*metadata.Parameter{{Name: "input", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Tests returning an error."}},
	},
	"add": {
		Name:       "add",
		Parameters: []*metadata.Parameter{{Name: "x", Type: "Int"}, {Name: "y", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs:       &metadata.Docs{Lines: []string{"Adds two integers together and returns the result."}},
	},
	"add3": {
		Name:       "add3",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}, {Name: "c~", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs: &metadata.Docs{
			Lines: []string{"Adds three integers together and returns the result.", "The third integer is optional."},
		},
	},
	"add3_WithDefaults": {
		Name:       "add3_WithDefaults",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs: &metadata.Docs{
			Lines: []string{"Adds three integers together and returns the result.", "The third integer is optional."},
		},
	},
	"cabi_realloc": {
		Name: "cabi_realloc",
		Parameters: []*metadata.Parameter{
			{Name: "src_offset", Type: "Int"}, {Name: "src_size", Type: "Int"},
			{Name: "_dst_alignment", Type: "Int"}, {Name: "dst_size", Type: "Int"},
		},
		Results: []*metadata.Result{{Type: "Int"}},
	},
	"copy": {
		Name:       "copy",
		Parameters: []*metadata.Parameter{{Name: "dest", Type: "Int"}, {Name: "src", Type: "Int"}},
	},
	"get_current_time": {
		Name:    "get_current_time",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time."}},
	},
	"get_current_time_formatted": {
		Name:    "get_current_time_formatted",
		Results: []*metadata.Result{{Type: "String!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time formatted as a string."}},
	},
	"get_full_name": {
		Name:       "get_full_name",
		Parameters: []*metadata.Parameter{{Name: "first_name", Type: "String"}, {Name: "last_name", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Combines the first and last name of a person, and returns the full name."}},
	},
	"get_local_time": {
		Name:    "get_local_time",
		Results: []*metadata.Result{{Type: "String!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current local time."}},
	},
	"get_local_time_zone_id": {
		Name:    "get_local_time_zone_id",
		Results: []*metadata.Result{{Type: "String"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the local time zone identifier."}},
	},
	"get_people": {
		Name:    "get_people",
		Results: []*metadata.Result{{Type: "Array[Person]"}},
		Docs:    &metadata.Docs{Lines: []string{"Gets a list of people."}},
	},
	"get_person": {
		Name:    "get_person",
		Results: []*metadata.Result{{Type: "Person"}},
		Docs:    &metadata.Docs{Lines: []string{"Gets a person object."}},
	},
	"get_random_person": {
		Name:    "get_random_person",
		Results: []*metadata.Result{{Type: "Person"}},
		Docs:    &metadata.Docs{Lines: []string{"Gets a random person object from a list of people."}},
	},
	"get_time_in_zone": {
		Name:       "get_time_in_zone",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Returns the current time in a specified time zone."}},
	},
	"get_time_zone_info": {
		Name:       "get_time_zone_info",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "TimeZoneInfo!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Returns some basic information about the time zone specified."}},
	},
	"get_utc_time": {
		Name:    "get_utc_time",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime!Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time in UTC."}},
	},
	"hello_array_of_ints": {
		Name:       "hello_array_of_ints",
		Parameters: []*metadata.Parameter{{Name: "n~", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Array[Int]"}},
	},
	"hello_array_of_ints_WithDefaults": {
		Name:    "hello_array_of_ints_WithDefaults",
		Results: []*metadata.Result{{Type: "Array[Int]"}},
	},
	"hello_array_of_ints_option": {
		Name:       "hello_array_of_ints_option",
		Parameters: []*metadata.Parameter{{Name: "n~", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Array[Int]?"}},
	},
	"hello_array_of_ints_option_WithDefaults": {
		Name:    "hello_array_of_ints_option_WithDefaults",
		Results: []*metadata.Result{{Type: "Array[Int]?"}},
	},
	"hello_array_of_strings": {
		Name:       "hello_array_of_strings",
		Parameters: []*metadata.Parameter{{Name: "n~", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Array[String]"}},
	},
	"hello_array_of_strings_WithDefaults": {
		Name:    "hello_array_of_strings_WithDefaults",
		Results: []*metadata.Result{{Type: "Array[String]"}},
	},
	"hello_array_of_strings_option": {
		Name:       "hello_array_of_strings_option",
		Parameters: []*metadata.Parameter{{Name: "n~", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Array[String]?"}},
	},
	"hello_array_of_strings_option_WithDefaults": {
		Name:    "hello_array_of_strings_option_WithDefaults",
		Results: []*metadata.Result{{Type: "Array[String]?"}},
	},
	"hello_maps_n_items": {
		Name:       "hello_maps_n_items",
		Parameters: []*metadata.Parameter{{Name: "n~", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Map[String, String]"}},
	},
	"hello_maps_n_items_WithDefaults": {
		Name:    "hello_maps_n_items_WithDefaults",
		Results: []*metadata.Result{{Type: "Map[String, String]"}},
	},
	"hello_maps_n_items_option": {
		Name:       "hello_maps_n_items_option",
		Parameters: []*metadata.Parameter{{Name: "n~", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Map[String, String]?"}},
	},
	"hello_maps_n_items_option_WithDefaults": {
		Name:    "hello_maps_n_items_option_WithDefaults",
		Results: []*metadata.Result{{Type: "Map[String, String]?"}},
	},
	"hello_option_empty_string": {
		Name:       "hello_option_empty_string",
		Parameters: []*metadata.Parameter{{Name: "some~", Type: "Bool"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"hello_option_empty_string_WithDefaults": {
		Name:    "hello_option_empty_string_WithDefaults",
		Results: []*metadata.Result{{Type: "String?"}},
	},
	"hello_option_none": {
		Name:       "hello_option_none",
		Parameters: []*metadata.Parameter{{Name: "some~", Type: "Bool"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"hello_option_none_WithDefaults": {
		Name:    "hello_option_none_WithDefaults",
		Results: []*metadata.Result{{Type: "String?"}},
	},
	"hello_option_some_string": {
		Name:       "hello_option_some_string",
		Parameters: []*metadata.Parameter{{Name: "some~", Type: "Bool"}},
		Results:    []*metadata.Result{{Type: "String?"}},
	},
	"hello_option_some_string_WithDefaults": {
		Name:    "hello_option_some_string_WithDefaults",
		Results: []*metadata.Result{{Type: "String?"}},
	},
	"hello_primitive_bool_max": {Name: "hello_primitive_bool_max", Results: []*metadata.Result{{Type: "Bool"}}},
	"hello_primitive_bool_min": {Name: "hello_primitive_bool_min", Results: []*metadata.Result{{Type: "Bool"}}},
	"hello_primitive_byte_max": {Name: "hello_primitive_byte_max", Results: []*metadata.Result{{Type: "Byte"}}},
	"hello_primitive_byte_min": {Name: "hello_primitive_byte_min", Results: []*metadata.Result{{Type: "Byte"}}},
	"hello_primitive_char_max": {Name: "hello_primitive_char_max", Results: []*metadata.Result{{Type: "Char"}}},
	"hello_primitive_char_min": {Name: "hello_primitive_char_min", Results: []*metadata.Result{{Type: "Char"}}},
	"hello_primitive_double_max": {
		Name:    "hello_primitive_double_max",
		Results: []*metadata.Result{{Type: "Double"}},
	},
	"hello_primitive_double_min": {
		Name:    "hello_primitive_double_min",
		Results: []*metadata.Result{{Type: "Double"}},
	},
	"hello_primitive_float_max": {Name: "hello_primitive_float_max", Results: []*metadata.Result{{Type: "Float"}}},
	"hello_primitive_float_min": {Name: "hello_primitive_float_min", Results: []*metadata.Result{{Type: "Float"}}},
	"hello_primitive_int16_max": {Name: "hello_primitive_int16_max", Results: []*metadata.Result{{Type: "Int16"}}},
	"hello_primitive_int16_min": {Name: "hello_primitive_int16_min", Results: []*metadata.Result{{Type: "Int16"}}},
	"hello_primitive_int64_max": {Name: "hello_primitive_int64_max", Results: []*metadata.Result{{Type: "Int64"}}},
	"hello_primitive_int64_min": {Name: "hello_primitive_int64_min", Results: []*metadata.Result{{Type: "Int64"}}},
	"hello_primitive_int_max":   {Name: "hello_primitive_int_max", Results: []*metadata.Result{{Type: "Int"}}},
	"hello_primitive_int_min":   {Name: "hello_primitive_int_min", Results: []*metadata.Result{{Type: "Int"}}},
	"hello_primitive_uint16_max": {
		Name:    "hello_primitive_uint16_max",
		Results: []*metadata.Result{{Type: "UInt16"}},
	},
	"hello_primitive_uint16_min": {
		Name:    "hello_primitive_uint16_min",
		Results: []*metadata.Result{{Type: "UInt16"}},
	},
	"hello_primitive_uint64_max": {
		Name:    "hello_primitive_uint64_max",
		Results: []*metadata.Result{{Type: "UInt64"}},
	},
	"hello_primitive_uint64_min": {
		Name:    "hello_primitive_uint64_min",
		Results: []*metadata.Result{{Type: "UInt64"}},
	},
	"hello_primitive_uint_max": {Name: "hello_primitive_uint_max", Results: []*metadata.Result{{Type: "UInt"}}},
	"hello_primitive_uint_min": {Name: "hello_primitive_uint_min", Results: []*metadata.Result{{Type: "UInt"}}},
	"hello_world":              {Name: "hello_world", Results: []*metadata.Result{{Type: "String"}}},
	"hello_world_with_arg": {
		Name:       "hello_world_with_arg",
		Parameters: []*metadata.Parameter{{Name: "name", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"hello_world_with_optional_arg": {
		Name:       "hello_world_with_optional_arg",
		Parameters: []*metadata.Parameter{{Name: "name~", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"hello_world_with_optional_arg_WithDefaults": {
		Name:    "hello_world_with_optional_arg_WithDefaults",
		Results: []*metadata.Result{{Type: "String"}},
	},
	"log_message": {
		Name:       "log_message",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Logs a message."}},
	},
	"malloc": {
		Name:       "malloc",
		Parameters: []*metadata.Parameter{{Name: "size", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
	},
	"ptr2double_array": {
		Name:       "ptr2double_array",
		Parameters: []*metadata.Parameter{{Name: "ptr", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "FixedArray[Double]"}},
	},
	"ptr2int64_array": {
		Name:       "ptr2int64_array",
		Parameters: []*metadata.Parameter{{Name: "ptr", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "FixedArray[Int64]"}},
	},
	"ptr2uint64_array": {
		Name:       "ptr2uint64_array",
		Parameters: []*metadata.Parameter{{Name: "ptr", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "FixedArray[UInt64]"}},
	},
	"test_abort": {Name: "test_abort", Docs: &metadata.Docs{Lines: []string{"Tests an abort."}}},
	"test_alternative_error": {
		Name:       "test_alternative_error",
		Parameters: []*metadata.Parameter{{Name: "input", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Tests an alternative way to handle errors in functions."}},
	},
	"test_exit": {
		Name: "test_exit",
		Docs: &metadata.Docs{Lines: []string{"Tests an exit with a non-zero exit code."}},
	},
	"test_logging": {
		Name: "test_logging",
		Docs: &metadata.Docs{Lines: []string{"Tests logging at different levels."}},
	},
	"test_normal_error": {
		Name:       "test_normal_error",
		Parameters: []*metadata.Parameter{{Name: "input", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Tests returning an error."}},
	},
}

var wantTestsuiteFnImports = metadata.FunctionMap{
	"modus_system.getTimeInZone": {
		Name:       "modus_system.getTimeInZone",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
	},
	"modus_system.getTimeZoneData": {
		Name:       "modus_system.getTimeZoneData",
		Parameters: []*metadata.Parameter{{Name: "tz", Type: "String"}, {Name: "format", Type: "String"}},
		Results:    []*metadata.Result{{Type: "Array[Byte]"}},
	},
	"modus_system.logMessage": {
		Name:       "modus_system.logMessage",
		Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
	},
}

var wantTestsuiteTypes = metadata.TypeMap{
	"(String)": {Id: 4,
		Name:   "(String)",
		Fields: []*metadata.Field{{Name: "0", Type: "String"}},
	},
	"@time.ZonedDateTime":       {Id: 5, Name: "@time.ZonedDateTime"},
	"@time.ZonedDateTime!Error": {Id: 6, Name: "@time.ZonedDateTime!Error"},
	"Array[Byte]":               {Id: 7, Name: "Array[Byte]"},
	"Array[Int]":                {Id: 8, Name: "Array[Int]"},
	"Array[Int]?":               {Id: 9, Name: "Array[Int]?"},
	"Array[Person]":             {Id: 10, Name: "Array[Person]"},
	"Array[String]":             {Id: 11, Name: "Array[String]"},
	"Array[String]?":            {Id: 12, Name: "Array[String]?"},
	"Bool":                      {Id: 13, Name: "Bool"},
	"Byte":                      {Id: 14, Name: "Byte"},
	"Char":                      {Id: 15, Name: "Char"},
	"Double":                    {Id: 16, Name: "Double"},
	"FixedArray[Double]":        {Id: 17, Name: "FixedArray[Double]"},
	"FixedArray[Int64]":         {Id: 18, Name: "FixedArray[Int64]"},
	"FixedArray[UInt64]":        {Id: 19, Name: "FixedArray[UInt64]"},
	"Float":                     {Id: 20, Name: "Float"},
	"Int":                       {Id: 21, Name: "Int"},
	"Int16":                     {Id: 22, Name: "Int16"},
	"Int64":                     {Id: 23, Name: "Int64"},
	"Map[String, String]":       {Id: 24, Name: "Map[String, String]"},
	"Map[String, String]?":      {Id: 25, Name: "Map[String, String]?"},
	"Person": {Id: 26,
		Name: "Person",
		Fields: []*metadata.Field{
			{Name: "firstName", Type: "String"}, {Name: "lastName", Type: "String"},
			{Name: "age", Type: "Int"},
		},
	},
	"String":       {Id: 27, Name: "String"},
	"String!Error": {Id: 28, Name: "String!Error"},
	"String?":      {Id: 29, Name: "String?"},
	"TimeZoneInfo": {Id: 30,
		Name: "TimeZoneInfo",
		Fields: []*metadata.Field{
			{Name: "standard_name", Type: "String"},
			{Name: "standard_offset", Type: "String"},
			{Name: "daylight_name", Type: "String"},
			{Name: "daylight_offset", Type: "String"},
		},
	},
	"TimeZoneInfo!Error": {Id: 31,
		Name: "TimeZoneInfo!Error",
		Fields: []*metadata.Field{
			{Name: "standard_name", Type: "String"},
			{Name: "standard_offset", Type: "String"},
			{Name: "daylight_name", Type: "String"},
			{Name: "daylight_offset", Type: "String"},
		},
	},
	"UInt":   {Id: 32, Name: "UInt"},
	"UInt16": {Id: 33, Name: "UInt16"},
	"UInt64": {Id: 34, Name: "UInt64"},
}
