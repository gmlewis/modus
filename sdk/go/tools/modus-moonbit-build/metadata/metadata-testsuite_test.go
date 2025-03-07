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
		{name: "__modus_add", want: "(x : Int, y : Int) -> Int"},
		{name: "__modus_add3", want: "(a : Int, b : Int, c : Int) -> Int"},
		{name: "__modus_add3_WithDefaults", want: "(a : Int, b : Int) -> Int"},
		{name: "__modus_get_current_time", want: "() -> @time.ZonedDateTime!Error"},
		{name: "__modus_get_current_time_formatted", want: "() -> String!Error"},
		{name: "__modus_get_full_name", want: "(first_name : String, last_name : String) -> String"},
		{name: "__modus_get_local_time", want: "() -> String!Error"},
		{name: "__modus_get_local_time_zone_id", want: "() -> String"},
		{name: "__modus_get_people", want: "() -> Array[Person]"},
		{name: "__modus_get_person", want: "() -> Person"},
		{name: "__modus_get_random_person", want: "() -> Person"},
		{name: "__modus_get_time_in_zone", want: "(tz : String) -> String!Error"},
		{name: "__modus_get_time_zone_info", want: "(tz : String) -> TimeZoneInfo!Error"},
		{name: "__modus_get_utc_time", want: "() -> @time.ZonedDateTime!Error"},
		{name: "__modus_hello_array_of_ints", want: "(n : Int) -> Array[Int]"},
		{name: "__modus_hello_array_of_ints_WithDefaults", want: "() -> Array[Int]"},
		{name: "__modus_hello_array_of_ints_option", want: "(n : Int) -> Array[Int]?"},
		{name: "__modus_hello_array_of_ints_option_WithDefaults", want: "() -> Array[Int]?"},
		{name: "__modus_hello_array_of_strings", want: "(n : Int) -> Array[String]"},
		{name: "__modus_hello_array_of_strings_WithDefaults", want: "() -> Array[String]"},
		{name: "__modus_hello_array_of_strings_option", want: "(n : Int) -> Array[String]?"},
		{name: "__modus_hello_array_of_strings_option_WithDefaults", want: "() -> Array[String]?"},
		{name: "__modus_hello_maps_n_items", want: "(n : Int) -> Map[String, String]"},
		{name: "__modus_hello_maps_n_items_WithDefaults", want: "() -> Map[String, String]"},
		{name: "__modus_hello_maps_n_items_option", want: "(n : Int) -> Map[String, String]?"},
		{name: "__modus_hello_maps_n_items_option_WithDefaults", want: "() -> Map[String, String]?"},
		{name: "__modus_hello_option_empty_string", want: "(some : Bool) -> String?"},
		{name: "__modus_hello_option_empty_string_WithDefaults", want: "() -> String?"},
		{name: "__modus_hello_option_none", want: "(some : Bool) -> String?"},
		{name: "__modus_hello_option_none_WithDefaults", want: "() -> String?"},
		{name: "__modus_hello_option_some_string", want: "(some : Bool) -> String?"},
		{name: "__modus_hello_option_some_string_WithDefaults", want: "() -> String?"},
		{name: "__modus_hello_primitive_bool_max", want: "() -> Bool"},
		{name: "__modus_hello_primitive_bool_min", want: "() -> Bool"},
		{name: "__modus_hello_primitive_byte_max", want: "() -> Byte"},
		{name: "__modus_hello_primitive_byte_min", want: "() -> Byte"},
		{name: "__modus_hello_primitive_char_max", want: "() -> Char"},
		{name: "__modus_hello_primitive_char_min", want: "() -> Char"},
		{name: "__modus_hello_primitive_double_max", want: "() -> Double"},
		{name: "__modus_hello_primitive_double_min", want: "() -> Double"},
		{name: "__modus_hello_primitive_float_max", want: "() -> Float"},
		{name: "__modus_hello_primitive_float_min", want: "() -> Float"},
		{name: "__modus_hello_primitive_int16_max", want: "() -> Int16"},
		{name: "__modus_hello_primitive_int16_min", want: "() -> Int16"},
		{name: "__modus_hello_primitive_int64_max", want: "() -> Int64"},
		{name: "__modus_hello_primitive_int64_min", want: "() -> Int64"},
		{name: "__modus_hello_primitive_int_max", want: "() -> Int"},
		{name: "__modus_hello_primitive_int_min", want: "() -> Int"},
		{name: "__modus_hello_primitive_uint16_max", want: "() -> UInt16"},
		{name: "__modus_hello_primitive_uint16_min", want: "() -> UInt16"},
		{name: "__modus_hello_primitive_uint64_max", want: "() -> UInt64"},
		{name: "__modus_hello_primitive_uint64_min", want: "() -> UInt64"},
		{name: "__modus_hello_primitive_uint_max", want: "() -> UInt"},
		{name: "__modus_hello_primitive_uint_min", want: "() -> UInt"},
		{name: "__modus_hello_world", want: "() -> String"},
		{name: "__modus_hello_world_with_arg", want: "(name : String) -> String"},
		{name: "__modus_hello_world_with_optional_arg", want: "(name : String) -> String"},
		{name: "__modus_hello_world_with_optional_arg_WithDefaults", want: "() -> String"},
		{name: "__modus_log_message", want: "(message : String) -> Unit"},
		{name: "__modus_test_abort", want: "() -> Unit"},
		{name: "__modus_test_alternative_error", want: "(input : String) -> String"},
		{name: "__modus_test_exit", want: "() -> Unit"},
		{name: "__modus_test_logging", want: "() -> Unit"},
		{name: "__modus_test_normal_error", want: "(input : String) -> String!Error"},
		{name: "add", want: "(x : Int, y : Int) -> Int"},
		{name: "add3", want: "(a : Int, b : Int, c~ : Int) -> Int"},
		{name: "add3_WithDefaults", want: "(a : Int, b : Int) -> Int"},
		{name: "cabi_realloc", want: "(src_offset : Int, src_size : Int, _dst_alignment : Int, dst_size : Int) -> Int"},
		{name: "copy", want: "(dest : Int, src : Int) -> Unit"},
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
		{name: "malloc", want: "(size : Int) -> Int"},
		{name: "ptr2double_array", want: "(ptr : Int) -> FixedArray[Double]"},
		{name: "ptr2int64_array", want: "(ptr : Int) -> FixedArray[Int64]"},
		{name: "ptr2uint64_array", want: "(ptr : Int) -> FixedArray[UInt64]"},
		{name: "test_abort", want: "() -> Unit"},
		{name: "test_alternative_error", want: "(input : String) -> String"},
		{name: "test_exit", want: "() -> Unit"},
		{name: "test_logging", want: "() -> Unit"},
		{name: "test_normal_error", want: "(input : String) -> String!Error"},
	}

	testFunctionStringHelper(t, "testsuite", testsuiteMetadataJSON, tests)
}
