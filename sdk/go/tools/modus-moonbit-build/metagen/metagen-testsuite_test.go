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
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateMetadata_Testsuite(t *testing.T) {
	meta := setupTestConfig(t, "../testdata/test-suite")
	removeExternalFuncsForComparison(t, meta)

	if got, want := meta.Plugin, "test-suite"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@test-suite"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	// if got, want := meta.SDK, "modus-sdk-mbt@40.11.0"; got != want {
	// 	t.Errorf("meta.SDK = %q, want %q", got, want)
	// }

	if diff := cmp.Diff(wantTestsuiteFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantTestsuiteFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	diffMetaTypes(t, wantTestsuiteTypes, meta.Types)
}

var wantTestsuiteFnExports = metadata.FunctionMap{
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
	"call_test_time_option_input_some": {
		Name:    "call_test_time_option_input_some",
		Results: []*metadata.Result{{Type: "Unit!Error"}},
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
	"make_test_map": {
		Name:       "make_test_map",
		Parameters: []*metadata.Parameter{{Name: "size", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Map[String, String]"}},
		Docs:       &metadata.Docs{Lines: []string{"This generated map mimics the test map created on the Go side."}},
	},
	"test_abort": {Name: "test_abort", Docs: &metadata.Docs{Lines: []string{"Tests an abort."}}},
	"test_alternative_error": {
		Name:       "test_alternative_error",
		Parameters: []*metadata.Parameter{{Name: "input", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Tests an alternative way to handle errors in functions."}},
	},
	"test_array_input_string": {
		Name:       "test_array_input_string",
		Parameters: []*metadata.Parameter{{Name: "val", Type: "Array[String]"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_array_output_int_option": {
		Name:    "test_array_output_int_option",
		Results: []*metadata.Result{{Type: "Array[Int?]"}},
	},
	"test_array_output_string_option": {
		Name:    "test_array_output_string_option",
		Results: []*metadata.Result{{Type: "Array[String?]"}},
	},
	"test_exit": {
		Name: "test_exit",
		Docs: &metadata.Docs{Lines: []string{"Tests an exit with a non-zero exit code."}},
	},
	"test_iterate_map_string_string": {
		Name:       "test_iterate_map_string_string",
		Parameters: []*metadata.Parameter{{Name: "m", Type: "Map[String, String]"}},
	},
	"test_logging": {
		Name: "test_logging",
		Docs: &metadata.Docs{Lines: []string{"Tests logging at different levels."}},
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
	"test_normal_error": {
		Name:       "test_normal_error",
		Parameters: []*metadata.Parameter{{Name: "input", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String!Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Tests returning an error."}},
	},
	"test_struct_containing_map_input_string_string": {
		Name:       "test_struct_containing_map_input_string_string",
		Parameters: []*metadata.Parameter{{Name: "s", Type: "TestStructWithMap"}},
		Results:    []*metadata.Result{{Type: "Unit!Error"}},
	},
	"test_struct_containing_map_output_string_string": {
		Name:    "test_struct_containing_map_output_string_string",
		Results: []*metadata.Result{{Type: "TestStructWithMap"}},
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
	"test_tuple_output": {
		Name:    "test_tuple_output",
		Results: []*metadata.Result{{Type: "(Int, Bool, String)"}},
	},
	"test_tuple_simulator": {
		Name:    "test_tuple_simulator",
		Results: []*metadata.Result{{Type: "TupleSimulator"}},
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
	"(Int, Bool, String)": {
		Name: "(Int, Bool, String)",
		Fields: []*metadata.Field{
			{Name: "0", Type: "Int"}, {Name: "1", Type: "Bool"},
			{Name: "2", Type: "String"},
		},
	},
	},
	"(String)":                {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
	"@ffi.XExternByteArray":   {Name: "@ffi.XExternByteArray"},
	"@ffi.XExternString":      {Name: "@ffi.XExternString"},
	"@ffi.XExternStringArray": {Name: "@ffi.XExternStringArray"},
	"@testutils.CallStack[T]": {
		Name:   "@testutils.CallStack[T]",
		Fields: []*metadata.Field{{Name: "items", Type: "Array[Array[@testutils.T]]"}},
	},
	"@testutils.T":              {Name: "@testutils.T"},
	"@time.Duration":            {Name: "@time.Duration"},
	"@time.Duration!Error":      {Name: "@time.Duration!Error"},
	"@time.Period":              {Name: "@time.Period"},
	"@time.Period!Error":        {Name: "@time.Period!Error"},
	"@time.PlainDate":           {Name: "@time.PlainDate"},
	"@time.PlainDate!Error":     {Name: "@time.PlainDate!Error"},
	"@time.PlainDateTime":       {Name: "@time.PlainDateTime"},
	"@time.PlainDateTime!Error": {Name: "@time.PlainDateTime!Error"},
	"@time.PlainTime":           {Name: "@time.PlainTime"},
	"@time.PlainTime!Error":     {Name: "@time.PlainTime!Error"},
	"@time.Weekday":             {Name: "@time.Weekday"},
	"@time.Zone":                {Name: "@time.Zone"},
	"@time.Zone!Error":          {Name: "@time.Zone!Error"},
	"@time.ZoneOffset":          {Name: "@time.ZoneOffset"},
	"@time.ZoneOffset!Error":    {Name: "@time.ZoneOffset!Error"},
	"@time.ZonedDateTime":       {Name: "@time.ZonedDateTime"},
	"@time.ZonedDateTime!Error": {Name: "@time.ZonedDateTime!Error"},
	"ArrayView[Byte]":           {Name: "ArrayView[Byte]"},
	"Array[@testutils.T]":       {Name: "Array[@testutils.T]"},
	"Array[Byte]":               {Name: "Array[Byte]"},
	"Array[Int]":                {Name: "Array[Int]"},
	"Array[Int]?":               {Name: "Array[Int]?"},
	"Array[Person]":             {Name: "Array[Person]"},
	"Array[String]":             {Name: "Array[String]"},
	"Array[String]?":            {Name: "Array[String]?"},
	"Bool":                      {Name: "Bool"},
	"Byte":                      {Name: "Byte"},
	"Bytes":                     {Name: "Bytes"},
	"Bytes!Error":               {Name: "Bytes!Error"},
	"Char":                      {Name: "Char"},
	"Double":                    {Name: "Double"},
	"FixedArray[Byte]":          {Name: "FixedArray[Byte]"},
	"FixedArray[Double]":        {Name: "FixedArray[Double]"},
	"FixedArray[Int64]":         {Name: "FixedArray[Int64]"},
	"FixedArray[UInt64]":        {Name: "FixedArray[UInt64]"},
	"Float":                     {Name: "Float"},
	"Int":                       {Name: "Int"},
	"Int16":                     {Name: "Int16"},
	"Int64":                     {Name: "Int64"},
	"Iter[Byte]":                {Name: "Iter[Byte]"},
	"Iter[Char]":                {Name: "Iter[Char]"},
	"Map[String, String]":       {Name: "Map[String, String]"},
	"Map[String, String]?":      {Name: "Map[String, String]?"},
	"Person": {
		Name: "Person",
		Fields: []*metadata.Field{
			{Name: "firstName", Type: "String"}, {Name: "lastName", Type: "String"},
			{Name: "age", Type: "Int"},
		},
	},
	"Result[UInt64, UInt]": {Name: "Result[UInt64, UInt]"},
	"Result[Unit, UInt]":   {Name: "Result[Unit, UInt]"},
	"String":               {Name: "String"},
	"String!Error":         {Name: "String!Error"},
	"String?":              {Name: "String?"},
	"TimeZoneInfo": {
	},
		Name: "TimeZoneInfo",
		Fields: []*metadata.Field{
			{Name: "standard_name", Type: "String"},
			{Name: "standard_offset", Type: "String"},
			{Name: "daylight_name", Type: "String"},
			{Name: "daylight_offset", Type: "String"},
		},
	},
	"TimeZoneInfo!Error": {
		Name: "TimeZoneInfo!Error",
		Fields: []*metadata.Field{
			{Name: "standard_name", Type: "String"},
			{Name: "standard_offset", Type: "String"},
			{Name: "daylight_name", Type: "String"},
			{Name: "daylight_offset", Type: "String"},
		},
	},
	"TupleSimulator": {
		Name: "TupleSimulator",
		Fields: []*metadata.Field{
			{Name: "t0", Type: "Int"}, {Name: "t1", Type: "Bool"},
			{Name: "t2", Type: "String"},
		},
	},
	"UInt":       {Id: 32, Name: "UInt"},
	"UInt16":     {Id: 33, Name: "UInt16"},
	"UInt64":     {Id: 34, Name: "UInt64"},
	"Unit":       {Name: "Unit"},
	"Unit!Error": {Name: "Unit!Error"},
}
