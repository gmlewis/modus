// -*- compile-command: "go test -run ^TestFunction_String_Testsuite$ ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package metadata

import (
	_ "embed"
	"testing"
)

//go:embed testdata/test-suite-metadata.json
var testsuiteMetadataJSON []byte

func TestFunction_String_Testsuite(t *testing.T) {
	t.Parallel()

	tests := []functionStringTest{
		{name: "add", want: "(x : Int, y : Int) -> Int"},
		{name: "add3", want: "(a : Int, b : Int, c~ : Int) -> Int"},
		{name: "add3_WithDefaults", want: "(a : Int, b : Int) -> Int"},
		{name: "call_test_time_option_input_some", want: "() -> Unit!Error"},
		{name: "get_current_time", want: "() -> @time.ZonedDateTime!Error"},
		{name: "get_current_time_formatted", want: "() -> String!Error"},
		{name: "get_full_name", want: "(first_name : String, last_name : String) -> String"},
		{name: "get_local_time", want: "() -> String!Error"},
		{name: "get_local_time_zone_id", want: "() -> String"},
		{name: "get_people", want: "() -> Array[Person]"},
		{name: "get_person", want: "() -> Person"},
		{name: "get_random_person", want: "() -> Person"},
		{name: "get_time_in_zone", want: "(tz : String) -> String!Error"},
		{name: "get_time_zone_info", want: "(tz : String) -> TimeZoneInfo!Error"},
		{name: "get_utc_time", want: "() -> @time.ZonedDateTime!Error"},
		{name: "hello_array_of_ints", want: "(n~ : Int) -> Array[Int]"},
		{name: "hello_array_of_ints_WithDefaults", want: "() -> Array[Int]"},
		{name: "hello_array_of_ints_option", want: "(n~ : Int) -> Array[Int]?"},
		{name: "hello_array_of_ints_option_WithDefaults", want: "() -> Array[Int]?"},
		{name: "hello_array_of_strings", want: "(n~ : Int) -> Array[String]"},
		{name: "hello_array_of_strings_WithDefaults", want: "() -> Array[String]"},
		{name: "hello_array_of_strings_option", want: "(n~ : Int) -> Array[String]?"},
		{name: "hello_array_of_strings_option_WithDefaults", want: "() -> Array[String]?"},
		{name: "hello_maps_n_items", want: "(n~ : Int) -> Map[String, String]"},
		{name: "hello_maps_n_items_WithDefaults", want: "() -> Map[String, String]"},
		{name: "hello_maps_n_items_option", want: "(n~ : Int) -> Map[String, String]?"},
		{name: "hello_maps_n_items_option_WithDefaults", want: "() -> Map[String, String]?"},
		{name: "hello_option_empty_string", want: "(some~ : Bool) -> String?"},
		{name: "hello_option_empty_string_WithDefaults", want: "() -> String?"},
		{name: "hello_option_none", want: "(some~ : Bool) -> String?"},
		{name: "hello_option_none_WithDefaults", want: "() -> String?"},
		{name: "hello_option_some_string", want: "(some~ : Bool) -> String?"},
		{name: "hello_option_some_string_WithDefaults", want: "() -> String?"},
		{name: "hello_primitive_bool_max", want: "() -> Bool"},
		{name: "hello_primitive_bool_min", want: "() -> Bool"},
		{name: "hello_primitive_byte_max", want: "() -> Byte"},
		{name: "hello_primitive_byte_min", want: "() -> Byte"},
		{name: "hello_primitive_char_max", want: "() -> Char"},
		{name: "hello_primitive_char_min", want: "() -> Char"},
		{name: "hello_primitive_double_max", want: "() -> Double"},
		{name: "hello_primitive_double_min", want: "() -> Double"},
		{name: "hello_primitive_float_max", want: "() -> Float"},
		{name: "hello_primitive_float_min", want: "() -> Float"},
		{name: "hello_primitive_int16_max", want: "() -> Int16"},
		{name: "hello_primitive_int16_min", want: "() -> Int16"},
		{name: "hello_primitive_int64_max", want: "() -> Int64"},
		{name: "hello_primitive_int64_min", want: "() -> Int64"},
		{name: "hello_primitive_int_max", want: "() -> Int"},
		{name: "hello_primitive_int_min", want: "() -> Int"},
		{name: "hello_primitive_uint16_max", want: "() -> UInt16"},
		{name: "hello_primitive_uint16_min", want: "() -> UInt16"},
		{name: "hello_primitive_uint64_max", want: "() -> UInt64"},
		{name: "hello_primitive_uint64_min", want: "() -> UInt64"},
		{name: "hello_primitive_uint_max", want: "() -> UInt"},
		{name: "hello_primitive_uint_min", want: "() -> UInt"},
		{name: "hello_world", want: "() -> String"},
		{name: "hello_world_with_arg", want: "(name : String) -> String"},
		{name: "hello_world_with_optional_arg", want: "(name~ : String) -> String"},
		{name: "hello_world_with_optional_arg_WithDefaults", want: "() -> String"},
		{name: "log_message", want: "(message : String) -> Unit"},
		{name: "make_test_map", want: "(size : Int) -> Map[String, String]"},
		{name: "test_abort", want: "() -> Unit"},
		{name: "test_alternative_error", want: "(input : String) -> String"},
		{name: "test_array_input_string", want: "(val : Array[String]) -> Unit!Error"},
		{name: "test_array_output_int_option", want: "() -> Array[Int?]"},
		{name: "test_array_output_string_option", want: "() -> Array[String?]"},
		{name: "test_exit", want: "() -> Unit"},
		{name: "test_iterate_map_string_string", want: "(m : Map[String, String]) -> Unit"},
		{name: "test_logging", want: "() -> Unit"},
		{name: "test_map_input_int_double", want: "(m : Map[Int, Double]) -> Unit!Error"},
		{name: "test_map_input_int_float", want: "(m : Map[Int, Float]) -> Unit!Error"},
		{name: "test_map_input_string_string", want: "(m : Map[String, String]) -> Unit!Error"},
		{name: "test_map_lookup_string_string", want: "(m : Map[String, String], key : String) -> String"},
		{name: "test_map_option_input_string_string", want: "(m : Map[String, String]?) -> Unit!Error"},
		{name: "test_map_option_output_string_string", want: "() -> Map[String, String]?"},
		{name: "test_map_output_int_double", want: "() -> Map[Int, Double]"},
		{name: "test_map_output_int_float", want: "() -> Map[Int, Float]"},
		{name: "test_map_output_string_string", want: "() -> Map[String, String]"},
		{name: "test_normal_error", want: "(input : String) -> String!Error"},
		{name: "test_struct_containing_map_input_string_string", want: "(s : TestStructWithMap) -> Unit!Error"},
		{name: "test_struct_containing_map_output_string_string", want: "() -> TestStructWithMap"},
		{name: "test_time_input", want: "(t : @time.ZonedDateTime) -> Unit!Error"},
		{name: "test_time_option_input", want: "(t : @time.ZonedDateTime?) -> Unit!Error"},
		{name: "test_tuple_output", want: "() -> (Int, Bool, String)"},
		{name: "test_tuple_simulator", want: "() -> TupleSimulator"},
	}

	testFunctionStringHelper(t, "testsuite", testsuiteMetadataJSON, tests)
}
