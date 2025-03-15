// -*- compile-command: "go test -run ^TestPackage_Runtime$ ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package packages

import (
	"encoding/json"
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestPackage_Runtime(t *testing.T) {
	t.Parallel()
	dir := "../testdata/runtime-testdata"
	testPackageLoadHelper(t, "runtime", dir, wantPackageRuntime)
}

var wantPackageRuntime = &Package{
	MoonPkgJSON: MoonPkgJSON{
		IsMain: false,
		Imports: []json.RawMessage{
			json.RawMessage(`"gmlewis/modus/pkg/console"`),
			json.RawMessage(`"gmlewis/modus/wit/ffi"`),
			json.RawMessage(`"moonbitlang/x/sys"`),
			json.RawMessage(`"moonbitlang/x/time"`),
		},
		Targets: map[string][]string{
			"debug-memory_notwasm.mbt": {"not", "wasm"},
			"debug-memory_wasm.mbt":    {"wasm"},
			"imports_notwasm.mbt":      {"not", "wasm"},
			"imports_wasm.mbt":         {"wasm"},
			"modus_post_generated.mbt": {"wasm"},
			"modus_pre_generated.mbt":  {"wasm"},
		},
		LinkTargets: map[string]*LinkTarget{
			"wasm": {
				Exports: []string{
					"__modus_add:add",
					"__modus_call_test_time_option_input_none:call_test_time_option_input_none",
					"__modus_call_test_time_option_input_some:call_test_time_option_input_some",
					"__modus_echo1:echo1",
					"__modus_echo2:echo2",
					"__modus_echo3:echo3",
					"__modus_echo4:echo4",
					"__modus_encode_strings1:encode_strings1",
					"__modus_encode_strings2:encode_strings2",
					"__modus_get_int_option_fixedarray1:get_int_option_fixedarray1",
					"__modus_get_int_ptr_array2:get_int_ptr_array2",
					"__modus_get_map_array2:get_map_array2",
					"__modus_get_map_ptr_array2:get_map_ptr_array2",
					"__modus_get_option_int_fixedarray1:get_option_int_fixedarray1",
					"__modus_get_option_int_fixedarray2:get_option_int_fixedarray2",
					"__modus_get_option_string_fixedarray1:get_option_string_fixedarray1",
					"__modus_get_option_string_fixedarray2:get_option_string_fixedarray2",
					"__modus_get_string_option_array2:get_string_option_array2",
					"__modus_get_string_option_fixedarray1:get_string_option_fixedarray1",
					"__modus_host_echo1:host_echo1",
					"__modus_host_echo2:host_echo2",
					"__modus_host_echo3:host_echo3",
					"__modus_host_echo4:host_echo4",
					"__modus_modus_test.add:modus_test.add",
					"__modus_modus_test.echo1:modus_test.echo1",
					"__modus_modus_test.echo2:modus_test.echo2",
					"__modus_modus_test.echo3:modus_test.echo3",
					"__modus_modus_test.echo4:modus_test.echo4",
					"__modus_modus_test.encodeStrings1:modus_test.encodeStrings1",
					"__modus_modus_test.encodeStrings2:modus_test.encodeStrings2",
					"__modus_test2d_array_input_string:test2d_array_input_string",
					"__modus_test2d_array_input_string_empty:test2d_array_input_string_empty",
					"__modus_test2d_array_input_string_inner_empty:test2d_array_input_string_inner_empty",
					"__modus_test2d_array_input_string_inner_none:test2d_array_input_string_inner_none",
					"__modus_test2d_array_input_string_none:test2d_array_input_string_none",
					"__modus_test2d_array_output_string:test2d_array_output_string",
					"__modus_test2d_array_output_string_empty:test2d_array_output_string_empty",
					"__modus_test2d_array_output_string_inner_empty:test2d_array_output_string_inner_empty",
					"__modus_test2d_array_output_string_inner_none:test2d_array_output_string_inner_none",
					"__modus_test2d_array_output_string_none:test2d_array_output_string_none",
					"__modus_test_array_input_bool_0:test_array_input_bool_0",
					"__modus_test_array_input_bool_1:test_array_input_bool_1",
					"__modus_test_array_input_bool_2:test_array_input_bool_2",
					"__modus_test_array_input_bool_3:test_array_input_bool_3",
					"__modus_test_array_input_bool_4:test_array_input_bool_4",
					"__modus_test_array_input_bool_option_0:test_array_input_bool_option_0",
					"__modus_test_array_input_bool_option_1_false:test_array_input_bool_option_1_false",
					"__modus_test_array_input_bool_option_1_none:test_array_input_bool_option_1_none",
					"__modus_test_array_input_bool_option_1_true:test_array_input_bool_option_1_true",
					"__modus_test_array_input_bool_option_2:test_array_input_bool_option_2",
					"__modus_test_array_input_bool_option_3:test_array_input_bool_option_3",
					"__modus_test_array_input_bool_option_4:test_array_input_bool_option_4",
					"__modus_test_array_input_byte_0:test_array_input_byte_0",
					"__modus_test_array_input_byte_1:test_array_input_byte_1",
					"__modus_test_array_input_byte_2:test_array_input_byte_2",
					"__modus_test_array_input_byte_3:test_array_input_byte_3",
					"__modus_test_array_input_byte_4:test_array_input_byte_4",
					"__modus_test_array_input_byte_option_0:test_array_input_byte_option_0",
					"__modus_test_array_input_byte_option_1:test_array_input_byte_option_1",
					"__modus_test_array_input_byte_option_2:test_array_input_byte_option_2",
					"__modus_test_array_input_byte_option_3:test_array_input_byte_option_3",
					"__modus_test_array_input_byte_option_4:test_array_input_byte_option_4",
					"__modus_test_array_input_char_empty:test_array_input_char_empty",
					"__modus_test_array_input_char_option:test_array_input_char_option",
					"__modus_test_array_input_double_empty:test_array_input_double_empty",
					"__modus_test_array_input_double_option:test_array_input_double_option",
					"__modus_test_array_input_float_empty:test_array_input_float_empty",
					"__modus_test_array_input_float_option:test_array_input_float_option",
					"__modus_test_array_input_int16_empty:test_array_input_int16_empty",
					"__modus_test_array_input_int16_option:test_array_input_int16_option",
					"__modus_test_array_input_int_empty:test_array_input_int_empty",
					"__modus_test_array_input_int_option:test_array_input_int_option",
					"__modus_test_array_input_string:test_array_input_string",
					"__modus_test_array_input_string_empty:test_array_input_string_empty",
					"__modus_test_array_input_string_none:test_array_input_string_none",
					"__modus_test_array_input_string_option:test_array_input_string_option",
					"__modus_test_array_input_uint16_empty:test_array_input_uint16_empty",
					"__modus_test_array_input_uint16_option:test_array_input_uint16_option",
					"__modus_test_array_input_uint_empty:test_array_input_uint_empty",
					"__modus_test_array_input_uint_option:test_array_input_uint_option",
					"__modus_test_array_output_bool_0:test_array_output_bool_0",
					"__modus_test_array_output_bool_1:test_array_output_bool_1",
					"__modus_test_array_output_bool_2:test_array_output_bool_2",
					"__modus_test_array_output_bool_3:test_array_output_bool_3",
					"__modus_test_array_output_bool_4:test_array_output_bool_4",
					"__modus_test_array_output_bool_option_0:test_array_output_bool_option_0",
					"__modus_test_array_output_bool_option_1_false:test_array_output_bool_option_1_false",
					"__modus_test_array_output_bool_option_1_none:test_array_output_bool_option_1_none",
					"__modus_test_array_output_bool_option_1_true:test_array_output_bool_option_1_true",
					"__modus_test_array_output_bool_option_2:test_array_output_bool_option_2",
					"__modus_test_array_output_bool_option_3:test_array_output_bool_option_3",
					"__modus_test_array_output_bool_option_4:test_array_output_bool_option_4",
					"__modus_test_array_output_byte_0:test_array_output_byte_0",
					"__modus_test_array_output_byte_1:test_array_output_byte_1",
					"__modus_test_array_output_byte_2:test_array_output_byte_2",
					"__modus_test_array_output_byte_3:test_array_output_byte_3",
					"__modus_test_array_output_byte_4:test_array_output_byte_4",
					"__modus_test_array_output_byte_option_0:test_array_output_byte_option_0",
					"__modus_test_array_output_byte_option_1:test_array_output_byte_option_1",
					"__modus_test_array_output_byte_option_2:test_array_output_byte_option_2",
					"__modus_test_array_output_byte_option_3:test_array_output_byte_option_3",
					"__modus_test_array_output_byte_option_4:test_array_output_byte_option_4",
					"__modus_test_array_output_char_0:test_array_output_char_0",
					"__modus_test_array_output_char_1:test_array_output_char_1",
					"__modus_test_array_output_char_2:test_array_output_char_2",
					"__modus_test_array_output_char_3:test_array_output_char_3",
					"__modus_test_array_output_char_4:test_array_output_char_4",
					"__modus_test_array_output_char_option:test_array_output_char_option",
					"__modus_test_array_output_char_option_0:test_array_output_char_option_0",
					"__modus_test_array_output_char_option_1_none:test_array_output_char_option_1_none",
					"__modus_test_array_output_char_option_1_some:test_array_output_char_option_1_some",
					"__modus_test_array_output_char_option_2:test_array_output_char_option_2",
					"__modus_test_array_output_char_option_3:test_array_output_char_option_3",
					"__modus_test_array_output_char_option_4:test_array_output_char_option_4",
					"__modus_test_array_output_double_0:test_array_output_double_0",
					"__modus_test_array_output_double_1:test_array_output_double_1",
					"__modus_test_array_output_double_2:test_array_output_double_2",
					"__modus_test_array_output_double_3:test_array_output_double_3",
					"__modus_test_array_output_double_4:test_array_output_double_4",
					"__modus_test_array_output_double_option:test_array_output_double_option",
					"__modus_test_array_output_double_option_0:test_array_output_double_option_0",
					"__modus_test_array_output_double_option_1_none:test_array_output_double_option_1_none",
					"__modus_test_array_output_double_option_1_some:test_array_output_double_option_1_some",
					"__modus_test_array_output_double_option_2:test_array_output_double_option_2",
					"__modus_test_array_output_double_option_3:test_array_output_double_option_3",
					"__modus_test_array_output_double_option_4:test_array_output_double_option_4",
					"__modus_test_array_output_float_0:test_array_output_float_0",
					"__modus_test_array_output_float_1:test_array_output_float_1",
					"__modus_test_array_output_float_2:test_array_output_float_2",
					"__modus_test_array_output_float_3:test_array_output_float_3",
					"__modus_test_array_output_float_4:test_array_output_float_4",
					"__modus_test_array_output_float_option:test_array_output_float_option",
					"__modus_test_array_output_float_option_0:test_array_output_float_option_0",
					"__modus_test_array_output_float_option_1_none:test_array_output_float_option_1_none",
					"__modus_test_array_output_float_option_1_some:test_array_output_float_option_1_some",
					"__modus_test_array_output_float_option_2:test_array_output_float_option_2",
					"__modus_test_array_output_float_option_3:test_array_output_float_option_3",
					"__modus_test_array_output_float_option_4:test_array_output_float_option_4",
					"__modus_test_array_output_int16_0:test_array_output_int16_0",
					"__modus_test_array_output_int16_1:test_array_output_int16_1",
					"__modus_test_array_output_int16_1_max:test_array_output_int16_1_max",
					"__modus_test_array_output_int16_1_min:test_array_output_int16_1_min",
					"__modus_test_array_output_int16_2:test_array_output_int16_2",
					"__modus_test_array_output_int16_3:test_array_output_int16_3",
					"__modus_test_array_output_int16_4:test_array_output_int16_4",
					"__modus_test_array_output_int16_option:test_array_output_int16_option",
					"__modus_test_array_output_int16_option_0:test_array_output_int16_option_0",
					"__modus_test_array_output_int16_option_1_max:test_array_output_int16_option_1_max",
					"__modus_test_array_output_int16_option_1_min:test_array_output_int16_option_1_min",
					"__modus_test_array_output_int16_option_1_none:test_array_output_int16_option_1_none",
					"__modus_test_array_output_int16_option_2:test_array_output_int16_option_2",
					"__modus_test_array_output_int16_option_3:test_array_output_int16_option_3",
					"__modus_test_array_output_int16_option_4:test_array_output_int16_option_4",
					"__modus_test_array_output_int64_0:test_array_output_int64_0",
					"__modus_test_array_output_int64_1:test_array_output_int64_1",
					"__modus_test_array_output_int64_1_max:test_array_output_int64_1_max",
					"__modus_test_array_output_int64_1_min:test_array_output_int64_1_min",
					"__modus_test_array_output_int64_2:test_array_output_int64_2",
					"__modus_test_array_output_int64_3:test_array_output_int64_3",
					"__modus_test_array_output_int64_4:test_array_output_int64_4",
					"__modus_test_array_output_int64_option_0:test_array_output_int64_option_0",
					"__modus_test_array_output_int64_option_1_max:test_array_output_int64_option_1_max",
					"__modus_test_array_output_int64_option_1_min:test_array_output_int64_option_1_min",
					"__modus_test_array_output_int64_option_1_none:test_array_output_int64_option_1_none",
					"__modus_test_array_output_int64_option_2:test_array_output_int64_option_2",
					"__modus_test_array_output_int64_option_3:test_array_output_int64_option_3",
					"__modus_test_array_output_int64_option_4:test_array_output_int64_option_4",
					"__modus_test_array_output_int_0:test_array_output_int_0",
					"__modus_test_array_output_int_1:test_array_output_int_1",
					"__modus_test_array_output_int_1_max:test_array_output_int_1_max",
					"__modus_test_array_output_int_1_min:test_array_output_int_1_min",
					"__modus_test_array_output_int_2:test_array_output_int_2",
					"__modus_test_array_output_int_3:test_array_output_int_3",
					"__modus_test_array_output_int_4:test_array_output_int_4",
					"__modus_test_array_output_int_option:test_array_output_int_option",
					"__modus_test_array_output_int_option_0:test_array_output_int_option_0",
					"__modus_test_array_output_int_option_1_max:test_array_output_int_option_1_max",
					"__modus_test_array_output_int_option_1_min:test_array_output_int_option_1_min",
					"__modus_test_array_output_int_option_1_none:test_array_output_int_option_1_none",
					"__modus_test_array_output_int_option_2:test_array_output_int_option_2",
					"__modus_test_array_output_int_option_3:test_array_output_int_option_3",
					"__modus_test_array_output_int_option_4:test_array_output_int_option_4",
					"__modus_test_array_output_string:test_array_output_string",
					"__modus_test_array_output_string_empty:test_array_output_string_empty",
					"__modus_test_array_output_string_none:test_array_output_string_none",
					"__modus_test_array_output_string_option:test_array_output_string_option",
					"__modus_test_array_output_uint16_0:test_array_output_uint16_0",
					"__modus_test_array_output_uint16_1:test_array_output_uint16_1",
					"__modus_test_array_output_uint16_1_max:test_array_output_uint16_1_max",
					"__modus_test_array_output_uint16_1_min:test_array_output_uint16_1_min",
					"__modus_test_array_output_uint16_2:test_array_output_uint16_2",
					"__modus_test_array_output_uint16_3:test_array_output_uint16_3",
					"__modus_test_array_output_uint16_4:test_array_output_uint16_4",
					"__modus_test_array_output_uint16_option:test_array_output_uint16_option",
					"__modus_test_array_output_uint16_option_0:test_array_output_uint16_option_0",
					"__modus_test_array_output_uint16_option_1_max:test_array_output_uint16_option_1_max",
					"__modus_test_array_output_uint16_option_1_min:test_array_output_uint16_option_1_min",
					"__modus_test_array_output_uint16_option_1_none:test_array_output_uint16_option_1_none",
					"__modus_test_array_output_uint16_option_2:test_array_output_uint16_option_2",
					"__modus_test_array_output_uint16_option_3:test_array_output_uint16_option_3",
					"__modus_test_array_output_uint16_option_4:test_array_output_uint16_option_4",
					"__modus_test_array_output_uint64_0:test_array_output_uint64_0",
					"__modus_test_array_output_uint64_1:test_array_output_uint64_1",
					"__modus_test_array_output_uint64_1_max:test_array_output_uint64_1_max",
					"__modus_test_array_output_uint64_1_min:test_array_output_uint64_1_min",
					"__modus_test_array_output_uint64_2:test_array_output_uint64_2",
					"__modus_test_array_output_uint64_3:test_array_output_uint64_3",
					"__modus_test_array_output_uint64_4:test_array_output_uint64_4",
					"__modus_test_array_output_uint64_option_0:test_array_output_uint64_option_0",
					"__modus_test_array_output_uint64_option_1_max:test_array_output_uint64_option_1_max",
					"__modus_test_array_output_uint64_option_1_min:test_array_output_uint64_option_1_min",
					"__modus_test_array_output_uint64_option_1_none:test_array_output_uint64_option_1_none",
					"__modus_test_array_output_uint64_option_2:test_array_output_uint64_option_2",
					"__modus_test_array_output_uint64_option_3:test_array_output_uint64_option_3",
					"__modus_test_array_output_uint64_option_4:test_array_output_uint64_option_4",
					"__modus_test_array_output_uint_0:test_array_output_uint_0",
					"__modus_test_array_output_uint_1:test_array_output_uint_1",
					"__modus_test_array_output_uint_1_max:test_array_output_uint_1_max",
					"__modus_test_array_output_uint_1_min:test_array_output_uint_1_min",
					"__modus_test_array_output_uint_2:test_array_output_uint_2",
					"__modus_test_array_output_uint_3:test_array_output_uint_3",
					"__modus_test_array_output_uint_4:test_array_output_uint_4",
					"__modus_test_array_output_uint_option:test_array_output_uint_option",
					"__modus_test_array_output_uint_option_0:test_array_output_uint_option_0",
					"__modus_test_array_output_uint_option_1_max:test_array_output_uint_option_1_max",
					"__modus_test_array_output_uint_option_1_min:test_array_output_uint_option_1_min",
					"__modus_test_array_output_uint_option_1_none:test_array_output_uint_option_1_none",
					"__modus_test_array_output_uint_option_2:test_array_output_uint_option_2",
					"__modus_test_array_output_uint_option_3:test_array_output_uint_option_3",
					"__modus_test_array_output_uint_option_4:test_array_output_uint_option_4",
					"__modus_test_bool_input_false:test_bool_input_false",
					"__modus_test_bool_input_true:test_bool_input_true",
					"__modus_test_bool_option_input_false:test_bool_option_input_false",
					"__modus_test_bool_option_input_none:test_bool_option_input_none",
					"__modus_test_bool_option_input_true:test_bool_option_input_true",
					"__modus_test_bool_option_output_false:test_bool_option_output_false",
					"__modus_test_bool_option_output_none:test_bool_option_output_none",
					"__modus_test_bool_option_output_true:test_bool_option_output_true",
					"__modus_test_bool_output_false:test_bool_output_false",
					"__modus_test_bool_output_true:test_bool_output_true",
					"__modus_test_byte_input_max:test_byte_input_max",
					"__modus_test_byte_input_min:test_byte_input_min",
					"__modus_test_byte_option_input_max:test_byte_option_input_max",
					"__modus_test_byte_option_input_min:test_byte_option_input_min",
					"__modus_test_byte_option_input_none:test_byte_option_input_none",
					"__modus_test_byte_option_output_max:test_byte_option_output_max",
					"__modus_test_byte_option_output_min:test_byte_option_output_min",
					"__modus_test_byte_option_output_none:test_byte_option_output_none",
					"__modus_test_byte_output_max:test_byte_output_max",
					"__modus_test_byte_output_min:test_byte_output_min",
					"__modus_test_char_input_max:test_char_input_max",
					"__modus_test_char_input_min:test_char_input_min",
					"__modus_test_char_option_input_max:test_char_option_input_max",
					"__modus_test_char_option_input_min:test_char_option_input_min",
					"__modus_test_char_option_input_none:test_char_option_input_none",
					"__modus_test_char_option_output_max:test_char_option_output_max",
					"__modus_test_char_option_output_min:test_char_option_output_min",
					"__modus_test_char_option_output_none:test_char_option_output_none",
					"__modus_test_char_output_max:test_char_output_max",
					"__modus_test_char_output_min:test_char_output_min",
					"__modus_test_double_input_max:test_double_input_max",
					"__modus_test_double_input_min:test_double_input_min",
					"__modus_test_double_option_input_max:test_double_option_input_max",
					"__modus_test_double_option_input_min:test_double_option_input_min",
					"__modus_test_double_option_input_none:test_double_option_input_none",
					"__modus_test_double_option_output_max:test_double_option_output_max",
					"__modus_test_double_option_output_min:test_double_option_output_min",
					"__modus_test_double_option_output_none:test_double_option_output_none",
					"__modus_test_double_output_max:test_double_output_max",
					"__modus_test_double_output_min:test_double_output_min",
					"__modus_test_duration_input:test_duration_input",
					"__modus_test_duration_option_input:test_duration_option_input",
					"__modus_test_duration_option_input_none:test_duration_option_input_none",
					"__modus_test_duration_option_input_none_style2:test_duration_option_input_none_style2",
					"__modus_test_duration_option_input_style2:test_duration_option_input_style2",
					"__modus_test_duration_option_output:test_duration_option_output",
					"__modus_test_duration_option_output_none:test_duration_option_output_none",
					"__modus_test_duration_output:test_duration_output",
					"__modus_test_fixedarray_input0_byte:test_fixedarray_input0_byte",
					"__modus_test_fixedarray_input0_int_option:test_fixedarray_input0_int_option",
					"__modus_test_fixedarray_input0_string:test_fixedarray_input0_string",
					"__modus_test_fixedarray_input0_string_option:test_fixedarray_input0_string_option",
					"__modus_test_fixedarray_input1_byte:test_fixedarray_input1_byte",
					"__modus_test_fixedarray_input1_int_option:test_fixedarray_input1_int_option",
					"__modus_test_fixedarray_input1_string:test_fixedarray_input1_string",
					"__modus_test_fixedarray_input1_string_option:test_fixedarray_input1_string_option",
					"__modus_test_fixedarray_input2_byte:test_fixedarray_input2_byte",
					"__modus_test_fixedarray_input2_int_option:test_fixedarray_input2_int_option",
					"__modus_test_fixedarray_input2_map:test_fixedarray_input2_map",
					"__modus_test_fixedarray_input2_map_option:test_fixedarray_input2_map_option",
					"__modus_test_fixedarray_input2_string:test_fixedarray_input2_string",
					"__modus_test_fixedarray_input2_string_option:test_fixedarray_input2_string_option",
					"__modus_test_fixedarray_input2_struct:test_fixedarray_input2_struct",
					"__modus_test_fixedarray_input2_struct_option:test_fixedarray_input2_struct_option",
					"__modus_test_fixedarray_input_bool_0:test_fixedarray_input_bool_0",
					"__modus_test_fixedarray_input_bool_1:test_fixedarray_input_bool_1",
					"__modus_test_fixedarray_input_bool_2:test_fixedarray_input_bool_2",
					"__modus_test_fixedarray_input_bool_3:test_fixedarray_input_bool_3",
					"__modus_test_fixedarray_input_bool_4:test_fixedarray_input_bool_4",
					"__modus_test_fixedarray_input_bool_option_0:test_fixedarray_input_bool_option_0",
					"__modus_test_fixedarray_input_bool_option_1_false:test_fixedarray_input_bool_option_1_false",
					"__modus_test_fixedarray_input_bool_option_1_none:test_fixedarray_input_bool_option_1_none",
					"__modus_test_fixedarray_input_bool_option_1_true:test_fixedarray_input_bool_option_1_true",
					"__modus_test_fixedarray_input_bool_option_2:test_fixedarray_input_bool_option_2",
					"__modus_test_fixedarray_input_bool_option_3:test_fixedarray_input_bool_option_3",
					"__modus_test_fixedarray_input_bool_option_4:test_fixedarray_input_bool_option_4",
					"__modus_test_fixedarray_input_byte_0:test_fixedarray_input_byte_0",
					"__modus_test_fixedarray_input_byte_1:test_fixedarray_input_byte_1",
					"__modus_test_fixedarray_input_byte_2:test_fixedarray_input_byte_2",
					"__modus_test_fixedarray_input_byte_3:test_fixedarray_input_byte_3",
					"__modus_test_fixedarray_input_byte_4:test_fixedarray_input_byte_4",
					"__modus_test_fixedarray_input_byte_option_0:test_fixedarray_input_byte_option_0",
					"__modus_test_fixedarray_input_byte_option_1:test_fixedarray_input_byte_option_1",
					"__modus_test_fixedarray_input_byte_option_2:test_fixedarray_input_byte_option_2",
					"__modus_test_fixedarray_input_byte_option_3:test_fixedarray_input_byte_option_3",
					"__modus_test_fixedarray_input_byte_option_4:test_fixedarray_input_byte_option_4",
					"__modus_test_fixedarray_input_char_0:test_fixedarray_input_char_0",
					"__modus_test_fixedarray_input_char_1:test_fixedarray_input_char_1",
					"__modus_test_fixedarray_input_char_2:test_fixedarray_input_char_2",
					"__modus_test_fixedarray_input_char_3:test_fixedarray_input_char_3",
					"__modus_test_fixedarray_input_char_4:test_fixedarray_input_char_4",
					"__modus_test_fixedarray_input_char_option_0:test_fixedarray_input_char_option_0",
					"__modus_test_fixedarray_input_char_option_1_none:test_fixedarray_input_char_option_1_none",
					"__modus_test_fixedarray_input_char_option_1_some:test_fixedarray_input_char_option_1_some",
					"__modus_test_fixedarray_input_char_option_2:test_fixedarray_input_char_option_2",
					"__modus_test_fixedarray_input_char_option_3:test_fixedarray_input_char_option_3",
					"__modus_test_fixedarray_input_char_option_4:test_fixedarray_input_char_option_4",
					"__modus_test_fixedarray_input_double_0:test_fixedarray_input_double_0",
					"__modus_test_fixedarray_input_double_1:test_fixedarray_input_double_1",
					"__modus_test_fixedarray_input_double_2:test_fixedarray_input_double_2",
					"__modus_test_fixedarray_input_double_3:test_fixedarray_input_double_3",
					"__modus_test_fixedarray_input_double_4:test_fixedarray_input_double_4",
					"__modus_test_fixedarray_input_double_option_0:test_fixedarray_input_double_option_0",
					"__modus_test_fixedarray_input_double_option_1_none:test_fixedarray_input_double_option_1_none",
					"__modus_test_fixedarray_input_double_option_1_some:test_fixedarray_input_double_option_1_some",
					"__modus_test_fixedarray_input_double_option_2:test_fixedarray_input_double_option_2",
					"__modus_test_fixedarray_input_double_option_3:test_fixedarray_input_double_option_3",
					"__modus_test_fixedarray_input_double_option_4:test_fixedarray_input_double_option_4",
					"__modus_test_fixedarray_input_float_0:test_fixedarray_input_float_0",
					"__modus_test_fixedarray_input_float_1:test_fixedarray_input_float_1",
					"__modus_test_fixedarray_input_float_2:test_fixedarray_input_float_2",
					"__modus_test_fixedarray_input_float_3:test_fixedarray_input_float_3",
					"__modus_test_fixedarray_input_float_4:test_fixedarray_input_float_4",
					"__modus_test_fixedarray_input_float_option_0:test_fixedarray_input_float_option_0",
					"__modus_test_fixedarray_input_float_option_1_none:test_fixedarray_input_float_option_1_none",
					"__modus_test_fixedarray_input_float_option_1_some:test_fixedarray_input_float_option_1_some",
					"__modus_test_fixedarray_input_float_option_2:test_fixedarray_input_float_option_2",
					"__modus_test_fixedarray_input_float_option_3:test_fixedarray_input_float_option_3",
					"__modus_test_fixedarray_input_float_option_4:test_fixedarray_input_float_option_4",
					"__modus_test_fixedarray_input_int16_0:test_fixedarray_input_int16_0",
					"__modus_test_fixedarray_input_int16_1:test_fixedarray_input_int16_1",
					"__modus_test_fixedarray_input_int16_1_max:test_fixedarray_input_int16_1_max",
					"__modus_test_fixedarray_input_int16_1_min:test_fixedarray_input_int16_1_min",
					"__modus_test_fixedarray_input_int16_2:test_fixedarray_input_int16_2",
					"__modus_test_fixedarray_input_int16_3:test_fixedarray_input_int16_3",
					"__modus_test_fixedarray_input_int16_4:test_fixedarray_input_int16_4",
					"__modus_test_fixedarray_input_int16_option_0:test_fixedarray_input_int16_option_0",
					"__modus_test_fixedarray_input_int16_option_1_max:test_fixedarray_input_int16_option_1_max",
					"__modus_test_fixedarray_input_int16_option_1_min:test_fixedarray_input_int16_option_1_min",
					"__modus_test_fixedarray_input_int16_option_1_none:test_fixedarray_input_int16_option_1_none",
					"__modus_test_fixedarray_input_int16_option_2:test_fixedarray_input_int16_option_2",
					"__modus_test_fixedarray_input_int16_option_3:test_fixedarray_input_int16_option_3",
					"__modus_test_fixedarray_input_int16_option_4:test_fixedarray_input_int16_option_4",
					"__modus_test_fixedarray_input_int64_0:test_fixedarray_input_int64_0",
					"__modus_test_fixedarray_input_int64_1:test_fixedarray_input_int64_1",
					"__modus_test_fixedarray_input_int64_1_max:test_fixedarray_input_int64_1_max",
					"__modus_test_fixedarray_input_int64_1_min:test_fixedarray_input_int64_1_min",
					"__modus_test_fixedarray_input_int64_2:test_fixedarray_input_int64_2",
					"__modus_test_fixedarray_input_int64_3:test_fixedarray_input_int64_3",
					"__modus_test_fixedarray_input_int64_4:test_fixedarray_input_int64_4",
					"__modus_test_fixedarray_input_int64_option_0:test_fixedarray_input_int64_option_0",
					"__modus_test_fixedarray_input_int64_option_1_max:test_fixedarray_input_int64_option_1_max",
					"__modus_test_fixedarray_input_int64_option_1_min:test_fixedarray_input_int64_option_1_min",
					"__modus_test_fixedarray_input_int64_option_1_none:test_fixedarray_input_int64_option_1_none",
					"__modus_test_fixedarray_input_int64_option_2:test_fixedarray_input_int64_option_2",
					"__modus_test_fixedarray_input_int64_option_3:test_fixedarray_input_int64_option_3",
					"__modus_test_fixedarray_input_int64_option_4:test_fixedarray_input_int64_option_4",
					"__modus_test_fixedarray_input_int_0:test_fixedarray_input_int_0",
					"__modus_test_fixedarray_input_int_1:test_fixedarray_input_int_1",
					"__modus_test_fixedarray_input_int_1_max:test_fixedarray_input_int_1_max",
					"__modus_test_fixedarray_input_int_1_min:test_fixedarray_input_int_1_min",
					"__modus_test_fixedarray_input_int_2:test_fixedarray_input_int_2",
					"__modus_test_fixedarray_input_int_3:test_fixedarray_input_int_3",
					"__modus_test_fixedarray_input_int_4:test_fixedarray_input_int_4",
					"__modus_test_fixedarray_input_int_option_0:test_fixedarray_input_int_option_0",
					"__modus_test_fixedarray_input_int_option_1_max:test_fixedarray_input_int_option_1_max",
					"__modus_test_fixedarray_input_int_option_1_min:test_fixedarray_input_int_option_1_min",
					"__modus_test_fixedarray_input_int_option_1_none:test_fixedarray_input_int_option_1_none",
					"__modus_test_fixedarray_input_int_option_2:test_fixedarray_input_int_option_2",
					"__modus_test_fixedarray_input_int_option_3:test_fixedarray_input_int_option_3",
					"__modus_test_fixedarray_input_int_option_4:test_fixedarray_input_int_option_4",
					"__modus_test_fixedarray_input_string_0:test_fixedarray_input_string_0",
					"__modus_test_fixedarray_input_string_1:test_fixedarray_input_string_1",
					"__modus_test_fixedarray_input_string_2:test_fixedarray_input_string_2",
					"__modus_test_fixedarray_input_string_3:test_fixedarray_input_string_3",
					"__modus_test_fixedarray_input_string_4:test_fixedarray_input_string_4",
					"__modus_test_fixedarray_input_string_option_0:test_fixedarray_input_string_option_0",
					"__modus_test_fixedarray_input_string_option_1_none:test_fixedarray_input_string_option_1_none",
					"__modus_test_fixedarray_input_string_option_1_some:test_fixedarray_input_string_option_1_some",
					"__modus_test_fixedarray_input_string_option_2:test_fixedarray_input_string_option_2",
					"__modus_test_fixedarray_input_string_option_3:test_fixedarray_input_string_option_3",
					"__modus_test_fixedarray_input_string_option_4:test_fixedarray_input_string_option_4",
					"__modus_test_fixedarray_input_uint16_0:test_fixedarray_input_uint16_0",
					"__modus_test_fixedarray_input_uint16_1:test_fixedarray_input_uint16_1",
					"__modus_test_fixedarray_input_uint16_1_max:test_fixedarray_input_uint16_1_max",
					"__modus_test_fixedarray_input_uint16_1_min:test_fixedarray_input_uint16_1_min",
					"__modus_test_fixedarray_input_uint16_2:test_fixedarray_input_uint16_2",
					"__modus_test_fixedarray_input_uint16_3:test_fixedarray_input_uint16_3",
					"__modus_test_fixedarray_input_uint16_4:test_fixedarray_input_uint16_4",
					"__modus_test_fixedarray_input_uint16_option_0:test_fixedarray_input_uint16_option_0",
					"__modus_test_fixedarray_input_uint16_option_1_max:test_fixedarray_input_uint16_option_1_max",
					"__modus_test_fixedarray_input_uint16_option_1_min:test_fixedarray_input_uint16_option_1_min",
					"__modus_test_fixedarray_input_uint16_option_1_none:test_fixedarray_input_uint16_option_1_none",
					"__modus_test_fixedarray_input_uint16_option_2:test_fixedarray_input_uint16_option_2",
					"__modus_test_fixedarray_input_uint16_option_3:test_fixedarray_input_uint16_option_3",
					"__modus_test_fixedarray_input_uint16_option_4:test_fixedarray_input_uint16_option_4",
					"__modus_test_fixedarray_input_uint64_0:test_fixedarray_input_uint64_0",
					"__modus_test_fixedarray_input_uint64_1:test_fixedarray_input_uint64_1",
					"__modus_test_fixedarray_input_uint64_1_max:test_fixedarray_input_uint64_1_max",
					"__modus_test_fixedarray_input_uint64_1_min:test_fixedarray_input_uint64_1_min",
					"__modus_test_fixedarray_input_uint64_2:test_fixedarray_input_uint64_2",
					"__modus_test_fixedarray_input_uint64_3:test_fixedarray_input_uint64_3",
					"__modus_test_fixedarray_input_uint64_4:test_fixedarray_input_uint64_4",
					"__modus_test_fixedarray_input_uint64_option_0:test_fixedarray_input_uint64_option_0",
					"__modus_test_fixedarray_input_uint64_option_1_max:test_fixedarray_input_uint64_option_1_max",
					"__modus_test_fixedarray_input_uint64_option_1_min:test_fixedarray_input_uint64_option_1_min",
					"__modus_test_fixedarray_input_uint64_option_1_none:test_fixedarray_input_uint64_option_1_none",
					"__modus_test_fixedarray_input_uint64_option_2:test_fixedarray_input_uint64_option_2",
					"__modus_test_fixedarray_input_uint64_option_3:test_fixedarray_input_uint64_option_3",
					"__modus_test_fixedarray_input_uint64_option_4:test_fixedarray_input_uint64_option_4",
					"__modus_test_fixedarray_input_uint_0:test_fixedarray_input_uint_0",
					"__modus_test_fixedarray_input_uint_1:test_fixedarray_input_uint_1",
					"__modus_test_fixedarray_input_uint_1_max:test_fixedarray_input_uint_1_max",
					"__modus_test_fixedarray_input_uint_1_min:test_fixedarray_input_uint_1_min",
					"__modus_test_fixedarray_input_uint_2:test_fixedarray_input_uint_2",
					"__modus_test_fixedarray_input_uint_3:test_fixedarray_input_uint_3",
					"__modus_test_fixedarray_input_uint_4:test_fixedarray_input_uint_4",
					"__modus_test_fixedarray_input_uint_option_0:test_fixedarray_input_uint_option_0",
					"__modus_test_fixedarray_input_uint_option_1_max:test_fixedarray_input_uint_option_1_max",
					"__modus_test_fixedarray_input_uint_option_1_min:test_fixedarray_input_uint_option_1_min",
					"__modus_test_fixedarray_input_uint_option_1_none:test_fixedarray_input_uint_option_1_none",
					"__modus_test_fixedarray_input_uint_option_2:test_fixedarray_input_uint_option_2",
					"__modus_test_fixedarray_input_uint_option_3:test_fixedarray_input_uint_option_3",
					"__modus_test_fixedarray_input_uint_option_4:test_fixedarray_input_uint_option_4",
					"__modus_test_fixedarray_output0_byte:test_fixedarray_output0_byte",
					"__modus_test_fixedarray_output0_int_option:test_fixedarray_output0_int_option",
					"__modus_test_fixedarray_output0_string:test_fixedarray_output0_string",
					"__modus_test_fixedarray_output0_string_option:test_fixedarray_output0_string_option",
					"__modus_test_fixedarray_output1_byte:test_fixedarray_output1_byte",
					"__modus_test_fixedarray_output1_int_option:test_fixedarray_output1_int_option",
					"__modus_test_fixedarray_output1_string:test_fixedarray_output1_string",
					"__modus_test_fixedarray_output1_string_option:test_fixedarray_output1_string_option",
					"__modus_test_fixedarray_output2_byte:test_fixedarray_output2_byte",
					"__modus_test_fixedarray_output2_int_option:test_fixedarray_output2_int_option",
					"__modus_test_fixedarray_output2_map:test_fixedarray_output2_map",
					"__modus_test_fixedarray_output2_map_option:test_fixedarray_output2_map_option",
					"__modus_test_fixedarray_output2_string:test_fixedarray_output2_string",
					"__modus_test_fixedarray_output2_string_option:test_fixedarray_output2_string_option",
					"__modus_test_fixedarray_output2_struct:test_fixedarray_output2_struct",
					"__modus_test_fixedarray_output2_struct_option:test_fixedarray_output2_struct_option",
					"__modus_test_fixedarray_output_bool_0:test_fixedarray_output_bool_0",
					"__modus_test_fixedarray_output_bool_1:test_fixedarray_output_bool_1",
					"__modus_test_fixedarray_output_bool_2:test_fixedarray_output_bool_2",
					"__modus_test_fixedarray_output_bool_3:test_fixedarray_output_bool_3",
					"__modus_test_fixedarray_output_bool_4:test_fixedarray_output_bool_4",
					"__modus_test_fixedarray_output_bool_option_0:test_fixedarray_output_bool_option_0",
					"__modus_test_fixedarray_output_bool_option_1_false:test_fixedarray_output_bool_option_1_false",
					"__modus_test_fixedarray_output_bool_option_1_none:test_fixedarray_output_bool_option_1_none",
					"__modus_test_fixedarray_output_bool_option_1_true:test_fixedarray_output_bool_option_1_true",
					"__modus_test_fixedarray_output_bool_option_2:test_fixedarray_output_bool_option_2",
					"__modus_test_fixedarray_output_bool_option_3:test_fixedarray_output_bool_option_3",
					"__modus_test_fixedarray_output_bool_option_4:test_fixedarray_output_bool_option_4",
					"__modus_test_fixedarray_output_byte_0:test_fixedarray_output_byte_0",
					"__modus_test_fixedarray_output_byte_1:test_fixedarray_output_byte_1",
					"__modus_test_fixedarray_output_byte_2:test_fixedarray_output_byte_2",
					"__modus_test_fixedarray_output_byte_3:test_fixedarray_output_byte_3",
					"__modus_test_fixedarray_output_byte_4:test_fixedarray_output_byte_4",
					"__modus_test_fixedarray_output_byte_option_0:test_fixedarray_output_byte_option_0",
					"__modus_test_fixedarray_output_byte_option_1:test_fixedarray_output_byte_option_1",
					"__modus_test_fixedarray_output_byte_option_2:test_fixedarray_output_byte_option_2",
					"__modus_test_fixedarray_output_byte_option_3:test_fixedarray_output_byte_option_3",
					"__modus_test_fixedarray_output_byte_option_4:test_fixedarray_output_byte_option_4",
					"__modus_test_fixedarray_output_char_0:test_fixedarray_output_char_0",
					"__modus_test_fixedarray_output_char_1:test_fixedarray_output_char_1",
					"__modus_test_fixedarray_output_char_2:test_fixedarray_output_char_2",
					"__modus_test_fixedarray_output_char_3:test_fixedarray_output_char_3",
					"__modus_test_fixedarray_output_char_4:test_fixedarray_output_char_4",
					"__modus_test_fixedarray_output_char_option_0:test_fixedarray_output_char_option_0",
					"__modus_test_fixedarray_output_char_option_1_none:test_fixedarray_output_char_option_1_none",
					"__modus_test_fixedarray_output_char_option_1_some:test_fixedarray_output_char_option_1_some",
					"__modus_test_fixedarray_output_char_option_2:test_fixedarray_output_char_option_2",
					"__modus_test_fixedarray_output_char_option_3:test_fixedarray_output_char_option_3",
					"__modus_test_fixedarray_output_char_option_4:test_fixedarray_output_char_option_4",
					"__modus_test_fixedarray_output_double_0:test_fixedarray_output_double_0",
					"__modus_test_fixedarray_output_double_1:test_fixedarray_output_double_1",
					"__modus_test_fixedarray_output_double_2:test_fixedarray_output_double_2",
					"__modus_test_fixedarray_output_double_3:test_fixedarray_output_double_3",
					"__modus_test_fixedarray_output_double_4:test_fixedarray_output_double_4",
					"__modus_test_fixedarray_output_double_option_0:test_fixedarray_output_double_option_0",
					"__modus_test_fixedarray_output_double_option_1_none:test_fixedarray_output_double_option_1_none",
					"__modus_test_fixedarray_output_double_option_1_some:test_fixedarray_output_double_option_1_some",
					"__modus_test_fixedarray_output_double_option_2:test_fixedarray_output_double_option_2",
					"__modus_test_fixedarray_output_double_option_3:test_fixedarray_output_double_option_3",
					"__modus_test_fixedarray_output_double_option_4:test_fixedarray_output_double_option_4",
					"__modus_test_fixedarray_output_float_0:test_fixedarray_output_float_0",
					"__modus_test_fixedarray_output_float_1:test_fixedarray_output_float_1",
					"__modus_test_fixedarray_output_float_2:test_fixedarray_output_float_2",
					"__modus_test_fixedarray_output_float_3:test_fixedarray_output_float_3",
					"__modus_test_fixedarray_output_float_4:test_fixedarray_output_float_4",
					"__modus_test_fixedarray_output_float_option_0:test_fixedarray_output_float_option_0",
					"__modus_test_fixedarray_output_float_option_1_none:test_fixedarray_output_float_option_1_none",
					"__modus_test_fixedarray_output_float_option_1_some:test_fixedarray_output_float_option_1_some",
					"__modus_test_fixedarray_output_float_option_2:test_fixedarray_output_float_option_2",
					"__modus_test_fixedarray_output_float_option_3:test_fixedarray_output_float_option_3",
					"__modus_test_fixedarray_output_float_option_4:test_fixedarray_output_float_option_4",
					"__modus_test_fixedarray_output_int16_0:test_fixedarray_output_int16_0",
					"__modus_test_fixedarray_output_int16_1:test_fixedarray_output_int16_1",
					"__modus_test_fixedarray_output_int16_1_max:test_fixedarray_output_int16_1_max",
					"__modus_test_fixedarray_output_int16_1_min:test_fixedarray_output_int16_1_min",
					"__modus_test_fixedarray_output_int16_2:test_fixedarray_output_int16_2",
					"__modus_test_fixedarray_output_int16_3:test_fixedarray_output_int16_3",
					"__modus_test_fixedarray_output_int16_4:test_fixedarray_output_int16_4",
					"__modus_test_fixedarray_output_int16_option_0:test_fixedarray_output_int16_option_0",
					"__modus_test_fixedarray_output_int16_option_1_max:test_fixedarray_output_int16_option_1_max",
					"__modus_test_fixedarray_output_int16_option_1_min:test_fixedarray_output_int16_option_1_min",
					"__modus_test_fixedarray_output_int16_option_1_none:test_fixedarray_output_int16_option_1_none",
					"__modus_test_fixedarray_output_int16_option_2:test_fixedarray_output_int16_option_2",
					"__modus_test_fixedarray_output_int16_option_3:test_fixedarray_output_int16_option_3",
					"__modus_test_fixedarray_output_int16_option_4:test_fixedarray_output_int16_option_4",
					"__modus_test_fixedarray_output_int64_0:test_fixedarray_output_int64_0",
					"__modus_test_fixedarray_output_int64_1:test_fixedarray_output_int64_1",
					"__modus_test_fixedarray_output_int64_1_max:test_fixedarray_output_int64_1_max",
					"__modus_test_fixedarray_output_int64_1_min:test_fixedarray_output_int64_1_min",
					"__modus_test_fixedarray_output_int64_2:test_fixedarray_output_int64_2",
					"__modus_test_fixedarray_output_int64_3:test_fixedarray_output_int64_3",
					"__modus_test_fixedarray_output_int64_4:test_fixedarray_output_int64_4",
					"__modus_test_fixedarray_output_int64_option_0:test_fixedarray_output_int64_option_0",
					"__modus_test_fixedarray_output_int64_option_1_max:test_fixedarray_output_int64_option_1_max",
					"__modus_test_fixedarray_output_int64_option_1_min:test_fixedarray_output_int64_option_1_min",
					"__modus_test_fixedarray_output_int64_option_1_none:test_fixedarray_output_int64_option_1_none",
					"__modus_test_fixedarray_output_int64_option_2:test_fixedarray_output_int64_option_2",
					"__modus_test_fixedarray_output_int64_option_3:test_fixedarray_output_int64_option_3",
					"__modus_test_fixedarray_output_int64_option_4:test_fixedarray_output_int64_option_4",
					"__modus_test_fixedarray_output_int_0:test_fixedarray_output_int_0",
					"__modus_test_fixedarray_output_int_1:test_fixedarray_output_int_1",
					"__modus_test_fixedarray_output_int_1_max:test_fixedarray_output_int_1_max",
					"__modus_test_fixedarray_output_int_1_min:test_fixedarray_output_int_1_min",
					"__modus_test_fixedarray_output_int_2:test_fixedarray_output_int_2",
					"__modus_test_fixedarray_output_int_3:test_fixedarray_output_int_3",
					"__modus_test_fixedarray_output_int_4:test_fixedarray_output_int_4",
					"__modus_test_fixedarray_output_int_option_0:test_fixedarray_output_int_option_0",
					"__modus_test_fixedarray_output_int_option_1_max:test_fixedarray_output_int_option_1_max",
					"__modus_test_fixedarray_output_int_option_1_min:test_fixedarray_output_int_option_1_min",
					"__modus_test_fixedarray_output_int_option_1_none:test_fixedarray_output_int_option_1_none",
					"__modus_test_fixedarray_output_int_option_2:test_fixedarray_output_int_option_2",
					"__modus_test_fixedarray_output_int_option_3:test_fixedarray_output_int_option_3",
					"__modus_test_fixedarray_output_int_option_4:test_fixedarray_output_int_option_4",
					"__modus_test_fixedarray_output_string_0:test_fixedarray_output_string_0",
					"__modus_test_fixedarray_output_string_1:test_fixedarray_output_string_1",
					"__modus_test_fixedarray_output_string_2:test_fixedarray_output_string_2",
					"__modus_test_fixedarray_output_string_3:test_fixedarray_output_string_3",
					"__modus_test_fixedarray_output_string_4:test_fixedarray_output_string_4",
					"__modus_test_fixedarray_output_string_option_0:test_fixedarray_output_string_option_0",
					"__modus_test_fixedarray_output_string_option_1_none:test_fixedarray_output_string_option_1_none",
					"__modus_test_fixedarray_output_string_option_1_some:test_fixedarray_output_string_option_1_some",
					"__modus_test_fixedarray_output_string_option_2:test_fixedarray_output_string_option_2",
					"__modus_test_fixedarray_output_string_option_3:test_fixedarray_output_string_option_3",
					"__modus_test_fixedarray_output_string_option_4:test_fixedarray_output_string_option_4",
					"__modus_test_fixedarray_output_uint16_0:test_fixedarray_output_uint16_0",
					"__modus_test_fixedarray_output_uint16_1:test_fixedarray_output_uint16_1",
					"__modus_test_fixedarray_output_uint16_1_max:test_fixedarray_output_uint16_1_max",
					"__modus_test_fixedarray_output_uint16_1_min:test_fixedarray_output_uint16_1_min",
					"__modus_test_fixedarray_output_uint16_2:test_fixedarray_output_uint16_2",
					"__modus_test_fixedarray_output_uint16_3:test_fixedarray_output_uint16_3",
					"__modus_test_fixedarray_output_uint16_4:test_fixedarray_output_uint16_4",
					"__modus_test_fixedarray_output_uint16_option_0:test_fixedarray_output_uint16_option_0",
					"__modus_test_fixedarray_output_uint16_option_1_max:test_fixedarray_output_uint16_option_1_max",
					"__modus_test_fixedarray_output_uint16_option_1_min:test_fixedarray_output_uint16_option_1_min",
					"__modus_test_fixedarray_output_uint16_option_1_none:test_fixedarray_output_uint16_option_1_none",
					"__modus_test_fixedarray_output_uint16_option_2:test_fixedarray_output_uint16_option_2",
					"__modus_test_fixedarray_output_uint16_option_3:test_fixedarray_output_uint16_option_3",
					"__modus_test_fixedarray_output_uint16_option_4:test_fixedarray_output_uint16_option_4",
					"__modus_test_fixedarray_output_uint64_0:test_fixedarray_output_uint64_0",
					"__modus_test_fixedarray_output_uint64_1:test_fixedarray_output_uint64_1",
					"__modus_test_fixedarray_output_uint64_1_max:test_fixedarray_output_uint64_1_max",
					"__modus_test_fixedarray_output_uint64_1_min:test_fixedarray_output_uint64_1_min",
					"__modus_test_fixedarray_output_uint64_2:test_fixedarray_output_uint64_2",
					"__modus_test_fixedarray_output_uint64_3:test_fixedarray_output_uint64_3",
					"__modus_test_fixedarray_output_uint64_4:test_fixedarray_output_uint64_4",
					"__modus_test_fixedarray_output_uint64_option_0:test_fixedarray_output_uint64_option_0",
					"__modus_test_fixedarray_output_uint64_option_1_max:test_fixedarray_output_uint64_option_1_max",
					"__modus_test_fixedarray_output_uint64_option_1_min:test_fixedarray_output_uint64_option_1_min",
					"__modus_test_fixedarray_output_uint64_option_1_none:test_fixedarray_output_uint64_option_1_none",
					"__modus_test_fixedarray_output_uint64_option_2:test_fixedarray_output_uint64_option_2",
					"__modus_test_fixedarray_output_uint64_option_3:test_fixedarray_output_uint64_option_3",
					"__modus_test_fixedarray_output_uint64_option_4:test_fixedarray_output_uint64_option_4",
					"__modus_test_fixedarray_output_uint_0:test_fixedarray_output_uint_0",
					"__modus_test_fixedarray_output_uint_1:test_fixedarray_output_uint_1",
					"__modus_test_fixedarray_output_uint_1_max:test_fixedarray_output_uint_1_max",
					"__modus_test_fixedarray_output_uint_1_min:test_fixedarray_output_uint_1_min",
					"__modus_test_fixedarray_output_uint_2:test_fixedarray_output_uint_2",
					"__modus_test_fixedarray_output_uint_3:test_fixedarray_output_uint_3",
					"__modus_test_fixedarray_output_uint_4:test_fixedarray_output_uint_4",
					"__modus_test_fixedarray_output_uint_option_0:test_fixedarray_output_uint_option_0",
					"__modus_test_fixedarray_output_uint_option_1_max:test_fixedarray_output_uint_option_1_max",
					"__modus_test_fixedarray_output_uint_option_1_min:test_fixedarray_output_uint_option_1_min",
					"__modus_test_fixedarray_output_uint_option_1_none:test_fixedarray_output_uint_option_1_none",
					"__modus_test_fixedarray_output_uint_option_2:test_fixedarray_output_uint_option_2",
					"__modus_test_fixedarray_output_uint_option_3:test_fixedarray_output_uint_option_3",
					"__modus_test_fixedarray_output_uint_option_4:test_fixedarray_output_uint_option_4",
					"__modus_test_float_input_max:test_float_input_max",
					"__modus_test_float_input_min:test_float_input_min",
					"__modus_test_float_option_input_max:test_float_option_input_max",
					"__modus_test_float_option_input_min:test_float_option_input_min",
					"__modus_test_float_option_input_none:test_float_option_input_none",
					"__modus_test_float_option_output_max:test_float_option_output_max",
					"__modus_test_float_option_output_min:test_float_option_output_min",
					"__modus_test_float_option_output_none:test_float_option_output_none",
					"__modus_test_float_output_max:test_float_output_max",
					"__modus_test_float_output_min:test_float_output_min",
					"__modus_test_generate_map_string_string_output:test_generate_map_string_string_output",
					"__modus_test_http_header:test_http_header",
					"__modus_test_http_header_map:test_http_header_map",
					"__modus_test_http_headers:test_http_headers",
					"__modus_test_http_response_headers:test_http_response_headers",
					"__modus_test_http_response_headers_output:test_http_response_headers_output",
					"__modus_test_int16_input_max:test_int16_input_max",
					"__modus_test_int16_input_min:test_int16_input_min",
					"__modus_test_int16_option_input_max:test_int16_option_input_max",
					"__modus_test_int16_option_input_min:test_int16_option_input_min",
					"__modus_test_int16_option_input_none:test_int16_option_input_none",
					"__modus_test_int16_option_output_max:test_int16_option_output_max",
					"__modus_test_int16_option_output_min:test_int16_option_output_min",
					"__modus_test_int16_option_output_none:test_int16_option_output_none",
					"__modus_test_int16_output_max:test_int16_output_max",
					"__modus_test_int16_output_min:test_int16_output_min",
					"__modus_test_int64_input_max:test_int64_input_max",
					"__modus_test_int64_input_min:test_int64_input_min",
					"__modus_test_int64_option_input_max:test_int64_option_input_max",
					"__modus_test_int64_option_input_min:test_int64_option_input_min",
					"__modus_test_int64_option_input_none:test_int64_option_input_none",
					"__modus_test_int64_option_output_max:test_int64_option_output_max",
					"__modus_test_int64_option_output_min:test_int64_option_output_min",
					"__modus_test_int64_option_output_none:test_int64_option_output_none",
					"__modus_test_int64_output_max:test_int64_output_max",
					"__modus_test_int64_output_min:test_int64_output_min",
					"__modus_test_int_input_max:test_int_input_max",
					"__modus_test_int_input_min:test_int_input_min",
					"__modus_test_int_option_input_max:test_int_option_input_max",
					"__modus_test_int_option_input_min:test_int_option_input_min",
					"__modus_test_int_option_input_none:test_int_option_input_none",
					"__modus_test_int_option_output_max:test_int_option_output_max",
					"__modus_test_int_option_output_min:test_int_option_output_min",
					"__modus_test_int_option_output_none:test_int_option_output_none",
					"__modus_test_int_output_max:test_int_output_max",
					"__modus_test_int_output_min:test_int_output_min",
					"__modus_test_iterate_map_string_string:test_iterate_map_string_string",
					"__modus_test_map_input_int_double:test_map_input_int_double",
					"__modus_test_map_input_int_float:test_map_input_int_float",
					"__modus_test_map_input_string_string:test_map_input_string_string",
					"__modus_test_map_lookup_string_string:test_map_lookup_string_string",
					"__modus_test_map_option_input_string_string:test_map_option_input_string_string",
					"__modus_test_map_option_output_string_string:test_map_option_output_string_string",
					"__modus_test_map_output_int_double:test_map_output_int_double",
					"__modus_test_map_output_int_float:test_map_output_int_float",
					"__modus_test_map_output_string_string:test_map_output_string_string",
					"__modus_test_option_fixedarray_input1_int:test_option_fixedarray_input1_int",
					"__modus_test_option_fixedarray_input1_string:test_option_fixedarray_input1_string",
					"__modus_test_option_fixedarray_input2_int:test_option_fixedarray_input2_int",
					"__modus_test_option_fixedarray_input2_string:test_option_fixedarray_input2_string",
					"__modus_test_option_fixedarray_output1_int:test_option_fixedarray_output1_int",
					"__modus_test_option_fixedarray_output1_string:test_option_fixedarray_output1_string",
					"__modus_test_option_fixedarray_output2_int:test_option_fixedarray_output2_int",
					"__modus_test_option_fixedarray_output2_string:test_option_fixedarray_output2_string",
					"__modus_test_recursive_struct_input:test_recursive_struct_input",
					"__modus_test_recursive_struct_option_input:test_recursive_struct_option_input",
					"__modus_test_recursive_struct_option_input_none:test_recursive_struct_option_input_none",
					"__modus_test_recursive_struct_option_output:test_recursive_struct_option_output",
					"__modus_test_recursive_struct_option_output_map:test_recursive_struct_option_output_map",
					"__modus_test_recursive_struct_option_output_none:test_recursive_struct_option_output_none",
					"__modus_test_recursive_struct_output:test_recursive_struct_output",
					"__modus_test_recursive_struct_output_map:test_recursive_struct_output_map",
					"__modus_test_smorgasbord_struct_input:test_smorgasbord_struct_input",
					"__modus_test_smorgasbord_struct_option_input:test_smorgasbord_struct_option_input",
					"__modus_test_smorgasbord_struct_option_input_none:test_smorgasbord_struct_option_input_none",
					"__modus_test_smorgasbord_struct_option_output:test_smorgasbord_struct_option_output",
					"__modus_test_smorgasbord_struct_option_output_map:test_smorgasbord_struct_option_output_map",
					"__modus_test_smorgasbord_struct_option_output_map_none:test_smorgasbord_struct_option_output_map_none",
					"__modus_test_smorgasbord_struct_option_output_none:test_smorgasbord_struct_option_output_none",
					"__modus_test_smorgasbord_struct_output:test_smorgasbord_struct_output",
					"__modus_test_smorgasbord_struct_output_map:test_smorgasbord_struct_output_map",
					"__modus_test_string_input:test_string_input",
					"__modus_test_string_option_input:test_string_option_input",
					"__modus_test_string_option_input_none:test_string_option_input_none",
					"__modus_test_string_option_output:test_string_option_output",
					"__modus_test_string_option_output_none:test_string_option_output_none",
					"__modus_test_string_output:test_string_output",
					"__modus_test_struct_containing_map_input_string_string:test_struct_containing_map_input_string_string",
					"__modus_test_struct_containing_map_output_string_string:test_struct_containing_map_output_string_string",
					"__modus_test_struct_input1:test_struct_input1",
					"__modus_test_struct_input2:test_struct_input2",
					"__modus_test_struct_input3:test_struct_input3",
					"__modus_test_struct_input4:test_struct_input4",
					"__modus_test_struct_input4_with_none:test_struct_input4_with_none",
					"__modus_test_struct_input5:test_struct_input5",
					"__modus_test_struct_option_input1:test_struct_option_input1",
					"__modus_test_struct_option_input1_none:test_struct_option_input1_none",
					"__modus_test_struct_option_input2:test_struct_option_input2",
					"__modus_test_struct_option_input2_none:test_struct_option_input2_none",
					"__modus_test_struct_option_input3:test_struct_option_input3",
					"__modus_test_struct_option_input3_none:test_struct_option_input3_none",
					"__modus_test_struct_option_input4:test_struct_option_input4",
					"__modus_test_struct_option_input4_none:test_struct_option_input4_none",
					"__modus_test_struct_option_input4_with_none:test_struct_option_input4_with_none",
					"__modus_test_struct_option_input5:test_struct_option_input5",
					"__modus_test_struct_option_input5_none:test_struct_option_input5_none",
					"__modus_test_struct_option_output1:test_struct_option_output1",
					"__modus_test_struct_option_output1_map:test_struct_option_output1_map",
					"__modus_test_struct_option_output1_none:test_struct_option_output1_none",
					"__modus_test_struct_option_output2:test_struct_option_output2",
					"__modus_test_struct_option_output2_map:test_struct_option_output2_map",
					"__modus_test_struct_option_output2_none:test_struct_option_output2_none",
					"__modus_test_struct_option_output3:test_struct_option_output3",
					"__modus_test_struct_option_output3_map:test_struct_option_output3_map",
					"__modus_test_struct_option_output3_none:test_struct_option_output3_none",
					"__modus_test_struct_option_output4:test_struct_option_output4",
					"__modus_test_struct_option_output4_map:test_struct_option_output4_map",
					"__modus_test_struct_option_output4_map_with_none:test_struct_option_output4_map_with_none",
					"__modus_test_struct_option_output4_none:test_struct_option_output4_none",
					"__modus_test_struct_option_output4_with_none:test_struct_option_output4_with_none",
					"__modus_test_struct_option_output5:test_struct_option_output5",
					"__modus_test_struct_option_output5_map:test_struct_option_output5_map",
					"__modus_test_struct_option_output5_none:test_struct_option_output5_none",
					"__modus_test_struct_output1:test_struct_output1",
					"__modus_test_struct_output1_map:test_struct_output1_map",
					"__modus_test_struct_output2:test_struct_output2",
					"__modus_test_struct_output2_map:test_struct_output2_map",
					"__modus_test_struct_output3:test_struct_output3",
					"__modus_test_struct_output3_map:test_struct_output3_map",
					"__modus_test_struct_output4:test_struct_output4",
					"__modus_test_struct_output4_map:test_struct_output4_map",
					"__modus_test_struct_output4_map_with_none:test_struct_output4_map_with_none",
					"__modus_test_struct_output4_with_none:test_struct_output4_with_none",
					"__modus_test_struct_output5:test_struct_output5",
					"__modus_test_struct_output5_map:test_struct_output5_map",
					"__modus_test_time_input:test_time_input",
					"__modus_test_time_option_input:test_time_option_input",
					"__modus_test_time_option_input_none:test_time_option_input_none",
					"__modus_test_time_option_input_none_style2:test_time_option_input_none_style2",
					"__modus_test_time_option_input_style2:test_time_option_input_style2",
					"__modus_test_time_option_output:test_time_option_output",
					"__modus_test_time_option_output_none:test_time_option_output_none",
					"__modus_test_time_output:test_time_output",
					"__modus_test_tuple_output:test_tuple_output",
					"__modus_test_uint16_input_max:test_uint16_input_max",
					"__modus_test_uint16_input_min:test_uint16_input_min",
					"__modus_test_uint16_option_input_max:test_uint16_option_input_max",
					"__modus_test_uint16_option_input_min:test_uint16_option_input_min",
					"__modus_test_uint16_option_input_none:test_uint16_option_input_none",
					"__modus_test_uint16_option_output_max:test_uint16_option_output_max",
					"__modus_test_uint16_option_output_min:test_uint16_option_output_min",
					"__modus_test_uint16_option_output_none:test_uint16_option_output_none",
					"__modus_test_uint16_output_max:test_uint16_output_max",
					"__modus_test_uint16_output_min:test_uint16_output_min",
					"__modus_test_uint64_input_max:test_uint64_input_max",
					"__modus_test_uint64_input_min:test_uint64_input_min",
					"__modus_test_uint64_option_input_max:test_uint64_option_input_max",
					"__modus_test_uint64_option_input_min:test_uint64_option_input_min",
					"__modus_test_uint64_option_input_none:test_uint64_option_input_none",
					"__modus_test_uint64_option_output_max:test_uint64_option_output_max",
					"__modus_test_uint64_option_output_min:test_uint64_option_output_min",
					"__modus_test_uint64_option_output_none:test_uint64_option_output_none",
					"__modus_test_uint64_output_max:test_uint64_output_max",
					"__modus_test_uint64_output_min:test_uint64_output_min",
					"__modus_test_uint_input_max:test_uint_input_max",
					"__modus_test_uint_input_min:test_uint_input_min",
					"__modus_test_uint_option_input_max:test_uint_option_input_max",
					"__modus_test_uint_option_input_min:test_uint_option_input_min",
					"__modus_test_uint_option_input_none:test_uint_option_input_none",
					"__modus_test_uint_option_output_max:test_uint_option_output_max",
					"__modus_test_uint_option_output_min:test_uint_option_output_min",
					"__modus_test_uint_option_output_none:test_uint_option_output_none",
					"__modus_test_uint_output_max:test_uint_output_max",
					"__modus_test_uint_output_min:test_uint_output_min",
					"cabi_realloc",
					"copy",
					"duration_from_nanos",
					"free",
					"load32",
					"malloc",
					"ptr2str",
					"ptr_to_none",
					"read_map",
					"store32",
					"store8",
					"write_map",
					"zoned_date_time_from_unix_seconds_and_nanos",
				},
				ExportMemoryName: "memory",
			},
		},
	},
	ID:      "moonbit-main",
	Name:    "main",
	PkgPath: "",
	MoonBitFiles: []string{
		"../testdata/runtime-testdata/arrays_bool.mbt",
		"../testdata/runtime-testdata/arrays_byte.mbt",
		"../testdata/runtime-testdata/arrays_char.mbt",
		"../testdata/runtime-testdata/arrays_double.mbt",
		"../testdata/runtime-testdata/arrays_float.mbt",
		"../testdata/runtime-testdata/arrays_int.mbt",
		"../testdata/runtime-testdata/arrays_int16.mbt",
		"../testdata/runtime-testdata/arrays_int64.mbt",
		"../testdata/runtime-testdata/arrays_string.mbt",
		"../testdata/runtime-testdata/arrays_uint.mbt",
		"../testdata/runtime-testdata/arrays_uint16.mbt",
		"../testdata/runtime-testdata/arrays_uint64.mbt",
		"../testdata/runtime-testdata/debug-memory_notwasm.mbt",
		"../testdata/runtime-testdata/debug-memory_wasm.mbt",
		"../testdata/runtime-testdata/fixedarrays.mbt",
		"../testdata/runtime-testdata/fixedarrays_bool.mbt",
		"../testdata/runtime-testdata/fixedarrays_byte.mbt",
		"../testdata/runtime-testdata/fixedarrays_char.mbt",
		"../testdata/runtime-testdata/fixedarrays_double.mbt",
		"../testdata/runtime-testdata/fixedarrays_float.mbt",
		"../testdata/runtime-testdata/fixedarrays_int.mbt",
		"../testdata/runtime-testdata/fixedarrays_int16.mbt",
		"../testdata/runtime-testdata/fixedarrays_int64.mbt",
		"../testdata/runtime-testdata/fixedarrays_string.mbt",
		"../testdata/runtime-testdata/fixedarrays_uint.mbt",
		"../testdata/runtime-testdata/fixedarrays_uint16.mbt",
		"../testdata/runtime-testdata/fixedarrays_uint64.mbt",
		"../testdata/runtime-testdata/hostfns.mbt",
		"../testdata/runtime-testdata/http.mbt",
		"../testdata/runtime-testdata/imports_notwasm.mbt",
		"../testdata/runtime-testdata/imports_wasm.mbt",
		"../testdata/runtime-testdata/maps.mbt",
		"../testdata/runtime-testdata/primitives_bool.mbt",
		"../testdata/runtime-testdata/primitives_byte.mbt",
		"../testdata/runtime-testdata/primitives_char.mbt",
		"../testdata/runtime-testdata/primitives_double.mbt",
		"../testdata/runtime-testdata/primitives_float.mbt",
		"../testdata/runtime-testdata/primitives_int.mbt",
		"../testdata/runtime-testdata/primitives_int16.mbt",
		"../testdata/runtime-testdata/primitives_int64.mbt",
		"../testdata/runtime-testdata/primitives_uint.mbt",
		"../testdata/runtime-testdata/primitives_uint16.mbt",
		"../testdata/runtime-testdata/primitives_uint64.mbt",
		"../testdata/runtime-testdata/strings.mbt",
		"../testdata/runtime-testdata/structs.mbt",
		"../testdata/runtime-testdata/time.mbt",
		"../testdata/runtime-testdata/tuples.mbt",
		"../testdata/runtime-testdata/utils.mbt",
	},
	StructLookup: map[string]*ast.TypeSpec{
		"HttpHeader": {
			Name: &ast.Ident{Name: "HttpHeader"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "name"}}, Type: &ast.Ident{Name: "String"}},
				{Names: []*ast.Ident{{Name: "values"}}, Type: &ast.Ident{Name: "Array[String]"}},
			}}},
		},
		"HttpHeaders": {
			Name: &ast.Ident{Name: "HttpHeaders"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "data"}}, Type: &ast.Ident{Name: "Map[String, HttpHeader?]"}},
			}}},
		},
		"HttpHeaders?": {Name: &ast.Ident{Name: "HttpHeaders?"}},
		"HttpResponse": {
			Name: &ast.Ident{Name: "HttpResponse"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "status"}}, Type: &ast.Ident{Name: "UInt16"}},
				{Names: []*ast.Ident{{Name: "statusText"}}, Type: &ast.Ident{Name: "String"}},
				{Names: []*ast.Ident{{Name: "headers"}}, Type: &ast.Ident{Name: "HttpHeaders?"}},
				{Names: []*ast.Ident{{Name: "body"}}, Type: &ast.Ident{Name: "Array[Byte]"}},
			}}},
		},
		"TestRecursiveStruct": {
			Name: &ast.Ident{Name: "TestRecursiveStruct"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "mut b"}}, Type: &ast.Ident{Name: "TestRecursiveStruct?"}},
			}}},
		},
		"TestRecursiveStruct?": {Name: &ast.Ident{Name: "TestRecursiveStruct?"}},
		"TestRecursiveStruct_map": {
			Name: &ast.Ident{Name: "TestRecursiveStruct_map"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "mut b"}}, Type: &ast.Ident{Name: "TestRecursiveStruct_map?"}},
			}}},
		},
		"TestRecursiveStruct_map?": {Name: &ast.Ident{Name: "TestRecursiveStruct_map?"}},
		"TestSmorgasbordStruct": {
			Name: &ast.Ident{Name: "TestSmorgasbordStruct"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "bool"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "byte"}}, Type: &ast.Ident{Name: "Byte"}},
				{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "Char"}},
				{Names: []*ast.Ident{{Name: "f"}}, Type: &ast.Ident{Name: "Float"}},
				{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "Double"}},
				{Names: []*ast.Ident{{Name: "i16"}}, Type: &ast.Ident{Name: "Int16"}},
				{Names: []*ast.Ident{{Name: "i32"}}, Type: &ast.Ident{Name: "Int"}},
				{Names: []*ast.Ident{{Name: "i64"}}, Type: &ast.Ident{Name: "Int64"}},
				{Names: []*ast.Ident{{Name: "s"}}, Type: &ast.Ident{Name: "String"}},
				{Names: []*ast.Ident{{Name: "u16"}}, Type: &ast.Ident{Name: "UInt16"}},
				{Names: []*ast.Ident{{Name: "u32"}}, Type: &ast.Ident{Name: "UInt"}},
				{Names: []*ast.Ident{{Name: "u64"}}, Type: &ast.Ident{Name: "UInt64"}},
				{Names: []*ast.Ident{{Name: "someBool"}}, Type: &ast.Ident{Name: "Bool?"}},
				{Names: []*ast.Ident{{Name: "noneBool"}}, Type: &ast.Ident{Name: "Bool?"}},
				{Names: []*ast.Ident{{Name: "someByte"}}, Type: &ast.Ident{Name: "Byte?"}},
				{Names: []*ast.Ident{{Name: "noneByte"}}, Type: &ast.Ident{Name: "Byte?"}},
				{Names: []*ast.Ident{{Name: "someChar"}}, Type: &ast.Ident{Name: "Char?"}},
				{Names: []*ast.Ident{{Name: "noneChar"}}, Type: &ast.Ident{Name: "Char?"}},
				{Names: []*ast.Ident{{Name: "someFloat"}}, Type: &ast.Ident{Name: "Float?"}},
				{Names: []*ast.Ident{{Name: "noneFloat"}}, Type: &ast.Ident{Name: "Float?"}},
				{Names: []*ast.Ident{{Name: "someDouble"}}, Type: &ast.Ident{Name: "Double?"}},
				{Names: []*ast.Ident{{Name: "noneDouble"}}, Type: &ast.Ident{Name: "Double?"}},
				{Names: []*ast.Ident{{Name: "someI16"}}, Type: &ast.Ident{Name: "Int16?"}},
				{Names: []*ast.Ident{{Name: "noneI16"}}, Type: &ast.Ident{Name: "Int16?"}},
				{Names: []*ast.Ident{{Name: "someI32"}}, Type: &ast.Ident{Name: "Int?"}},
				{Names: []*ast.Ident{{Name: "noneI32"}}, Type: &ast.Ident{Name: "Int?"}},
				{Names: []*ast.Ident{{Name: "someI64"}}, Type: &ast.Ident{Name: "Int64?"}},
				{Names: []*ast.Ident{{Name: "noneI64"}}, Type: &ast.Ident{Name: "Int64?"}},
				{Names: []*ast.Ident{{Name: "someString"}}, Type: &ast.Ident{Name: "String?"}},
				{Names: []*ast.Ident{{Name: "noneString"}}, Type: &ast.Ident{Name: "String?"}},
				{Names: []*ast.Ident{{Name: "someU16"}}, Type: &ast.Ident{Name: "UInt16?"}},
				{Names: []*ast.Ident{{Name: "noneU16"}}, Type: &ast.Ident{Name: "UInt16?"}},
				{Names: []*ast.Ident{{Name: "someU32"}}, Type: &ast.Ident{Name: "UInt?"}},
				{Names: []*ast.Ident{{Name: "noneU32"}}, Type: &ast.Ident{Name: "UInt?"}},
				{Names: []*ast.Ident{{Name: "someU64"}}, Type: &ast.Ident{Name: "UInt64?"}},
				{Names: []*ast.Ident{{Name: "noneU64"}}, Type: &ast.Ident{Name: "UInt64?"}},
			}}},
		},
		"TestSmorgasbordStruct_map": {
			Name: &ast.Ident{Name: "TestSmorgasbordStruct_map"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "bool"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "byte"}}, Type: &ast.Ident{Name: "Byte"}},
				{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "Char"}},
				{Names: []*ast.Ident{{Name: "f"}}, Type: &ast.Ident{Name: "Float"}},
				{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "Double"}},
				{Names: []*ast.Ident{{Name: "i16"}}, Type: &ast.Ident{Name: "Int16"}},
				{Names: []*ast.Ident{{Name: "i32"}}, Type: &ast.Ident{Name: "Int"}},
				{Names: []*ast.Ident{{Name: "i64"}}, Type: &ast.Ident{Name: "Int64"}},
				{Names: []*ast.Ident{{Name: "s"}}, Type: &ast.Ident{Name: "String"}},
				{Names: []*ast.Ident{{Name: "u16"}}, Type: &ast.Ident{Name: "UInt16"}},
				{Names: []*ast.Ident{{Name: "u32"}}, Type: &ast.Ident{Name: "UInt"}},
				{Names: []*ast.Ident{{Name: "u64"}}, Type: &ast.Ident{Name: "UInt64"}},
				{Names: []*ast.Ident{{Name: "someBool"}}, Type: &ast.Ident{Name: "Bool?"}},
				{Names: []*ast.Ident{{Name: "noneBool"}}, Type: &ast.Ident{Name: "Bool?"}},
				{Names: []*ast.Ident{{Name: "someByte"}}, Type: &ast.Ident{Name: "Byte?"}},
				{Names: []*ast.Ident{{Name: "noneByte"}}, Type: &ast.Ident{Name: "Byte?"}},
				{Names: []*ast.Ident{{Name: "someChar"}}, Type: &ast.Ident{Name: "Char?"}},
				{Names: []*ast.Ident{{Name: "noneChar"}}, Type: &ast.Ident{Name: "Char?"}},
				{Names: []*ast.Ident{{Name: "someFloat"}}, Type: &ast.Ident{Name: "Float?"}},
				{Names: []*ast.Ident{{Name: "noneFloat"}}, Type: &ast.Ident{Name: "Float?"}},
				{Names: []*ast.Ident{{Name: "someDouble"}}, Type: &ast.Ident{Name: "Double?"}},
				{Names: []*ast.Ident{{Name: "noneDouble"}}, Type: &ast.Ident{Name: "Double?"}},
				{Names: []*ast.Ident{{Name: "someI16"}}, Type: &ast.Ident{Name: "Int16?"}},
				{Names: []*ast.Ident{{Name: "noneI16"}}, Type: &ast.Ident{Name: "Int16?"}},
				{Names: []*ast.Ident{{Name: "someI32"}}, Type: &ast.Ident{Name: "Int?"}},
				{Names: []*ast.Ident{{Name: "noneI32"}}, Type: &ast.Ident{Name: "Int?"}},
				{Names: []*ast.Ident{{Name: "someI64"}}, Type: &ast.Ident{Name: "Int64?"}},
				{Names: []*ast.Ident{{Name: "noneI64"}}, Type: &ast.Ident{Name: "Int64?"}},
				{Names: []*ast.Ident{{Name: "someString"}}, Type: &ast.Ident{Name: "String?"}},
				{Names: []*ast.Ident{{Name: "noneString"}}, Type: &ast.Ident{Name: "String?"}},
				{Names: []*ast.Ident{{Name: "someU16"}}, Type: &ast.Ident{Name: "UInt16?"}},
				{Names: []*ast.Ident{{Name: "noneU16"}}, Type: &ast.Ident{Name: "UInt16?"}},
				{Names: []*ast.Ident{{Name: "someU32"}}, Type: &ast.Ident{Name: "UInt?"}},
				{Names: []*ast.Ident{{Name: "noneU32"}}, Type: &ast.Ident{Name: "UInt?"}},
				{Names: []*ast.Ident{{Name: "someU64"}}, Type: &ast.Ident{Name: "UInt64?"}},
				{Names: []*ast.Ident{{Name: "noneU64"}}, Type: &ast.Ident{Name: "UInt64?"}},
			}}},
		},
		"TestStruct1": {
			Name: &ast.Ident{Name: "TestStruct1"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
			}}},
		},
		"TestStruct1_map": {
			Name: &ast.Ident{Name: "TestStruct1_map"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
			}}},
		},
		"TestStruct2": {
			Name: &ast.Ident{Name: "TestStruct2"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
			}}},
		},
		"TestStruct2_map": {
			Name: &ast.Ident{Name: "TestStruct2_map"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
			}}},
		},
		"TestStruct3": {
			Name: &ast.Ident{Name: "TestStruct3"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
				{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String"}},
			}}},
		},
		"TestStruct3_map": {
			Name: &ast.Ident{Name: "TestStruct3_map"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
				{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String"}},
			}}},
		},
		"TestStruct4": {
			Name: &ast.Ident{Name: "TestStruct4"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
				{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String?"}},
			}}},
		},
		"TestStruct4_map": {
			Name: &ast.Ident{Name: "TestStruct4_map"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
				{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
				{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String?"}},
			}}},
		},
		"TestStruct5": {
			Name: &ast.Ident{Name: "TestStruct5"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "String"}},
				{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "String"}},
				{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String"}},
				{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "Array[String]"}},
				{Names: []*ast.Ident{{Name: "e"}}, Type: &ast.Ident{Name: "Double"}},
				{Names: []*ast.Ident{{Name: "f"}}, Type: &ast.Ident{Name: "Double"}},
			}}},
		},
		"TestStruct5_map": {
			Name: &ast.Ident{Name: "TestStruct5_map"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "String"}},
				{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "String"}},
				{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String"}},
				{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "Array[String]"}},
				{Names: []*ast.Ident{{Name: "e"}}, Type: &ast.Ident{Name: "Double"}},
				{Names: []*ast.Ident{{Name: "f"}}, Type: &ast.Ident{Name: "Double"}},
			}}},
		},
		"TestStructWithMap": {
			Name: &ast.Ident{Name: "TestStructWithMap"},
			Type: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]"}},
			}}},
		},
	},
	PossiblyMissingUnderlyingTypes: map[string]struct{}{
		"Array[Byte]":              {},
		"Array[String]":            {},
		"Bool":                     {},
		"Bool?":                    {},
		"Byte":                     {},
		"Byte?":                    {},
		"Char":                     {},
		"Char?":                    {},
		"Double":                   {},
		"Double?":                  {},
		"Float":                    {},
		"Float?":                   {},
		"HttpHeaders":              {},
		"HttpHeaders?":             {},
		"Int":                      {},
		"Int16":                    {},
		"Int16?":                   {},
		"Int64":                    {},
		"Int64?":                   {},
		"Int?":                     {},
		"Map[String, HttpHeader?]": {},
		"Map[String, String]":      {},
		"String":                   {},
		"String?":                  {},
		"TestRecursiveStruct":      {},
		"TestRecursiveStruct?":     {},
		"TestRecursiveStruct_map":  {},
		"TestRecursiveStruct_map?": {},
		"UInt":                     {},
		"UInt16":                   {},
		"UInt16?":                  {},
		"UInt64":                   {},
		"UInt64?":                  {},
		"UInt?":                    {},
	},
	Syntax: []*ast.File{
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_bool.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_option_1_false"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_option_1_true"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_bool_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_option_1_false"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_option_1_true"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_bool_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_byte.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_byte_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_byte_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_byte_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_byte_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_byte_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_byte_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Byte?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_byte_option_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Byte?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_byte_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Byte?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_byte_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Byte?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_byte_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Byte?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_byte_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_byte_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_byte_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_byte_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_byte_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_byte_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Byte?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_byte_option_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Byte?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_byte_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Byte?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_byte_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Byte?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_byte_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Byte?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_char.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_char_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Char]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_char_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Char?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_char_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Char?]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_double.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_double_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Double]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_double_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Double?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_double_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Double?]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_float.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_float_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Float]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_float_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Float?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_float_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Float?]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_int.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_int_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Int]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_int_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int?]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_int16.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_int16_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Int16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_int16_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Int16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int16_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int16?]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_int64.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_int64_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Int64?]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_string.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_string_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[String]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_string_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[String]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_string_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_string_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_string_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[String?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_string_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test2d_array_input_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Array[String]]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test2d_array_output_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Array[String]]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test2d_array_input_string_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Array[String]]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test2d_array_output_string_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Array[String]]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test2d_array_input_string_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Array[String]]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test2d_array_output_string_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Array[String]]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test2d_array_input_string_inner_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Array[String]]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test2d_array_output_string_inner_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Array[String]]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test2d_array_input_string_inner_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[Array[String]?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test2d_array_output_string_inner_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[Array[String]?]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_uint.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_uint_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[UInt]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_uint_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[UInt?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt?]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_uint16.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_uint16_empty"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[UInt16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_input_uint16_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "Array[UInt16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint16_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt16?]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/arrays_uint64.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_array_output_uint64_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Array[UInt64?]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/debug-memory_notwasm.mbt"},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/debug-memory_wasm.mbt"},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input0_byte"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output0_byte"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input0_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output0_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input0_string_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output0_string_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input0_int_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output0_int_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input1_byte"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output1_byte"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input1_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output1_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input1_string_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output1_string_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_string_option_fixedarray1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input1_int_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output1_int_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_int_option_fixedarray1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input2_byte"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output2_byte"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input2_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output2_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input2_string_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output2_string_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_string_option_array2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input2_int_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output2_int_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_int_ptr_array2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input2_struct"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[TestStruct2]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output2_struct"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[TestStruct2]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input2_struct_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[TestStruct2?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output2_struct_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[TestStruct2?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input2_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Map[String, String]]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output2_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Map[String, String]]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input2_map_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Map[String, String]?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output2_map_option"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Map[String, String]?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_map_array2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Map[String, String]]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_map_ptr_array2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Map[String, String]?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_option_fixedarray_input1_int"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_option_fixedarray_output1_int"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_option_int_fixedarray1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_option_fixedarray_input2_int"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_option_fixedarray_output2_int"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_option_int_fixedarray2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_option_fixedarray_input1_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_option_fixedarray_output1_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_option_string_fixedarray1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "get_option_string_fixedarray2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_option_fixedarray_input2_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_option_fixedarray_output2_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_bool.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_option_1_false"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_option_1_true"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_bool_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_option_1_false"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_option_1_true"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_bool_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Bool?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_byte.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_byte_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_byte_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_byte_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_byte_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_byte_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_byte_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_byte_option_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_byte_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_byte_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_byte_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Byte?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_byte_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_byte_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_byte_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_byte_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_byte_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_byte_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_byte_option_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_byte_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_byte_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_byte_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Byte?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_char.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_char_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_char_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Char?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_double.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_double_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_double_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Double?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_float.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_float_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_float_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Float?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_int.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_int16.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int16_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int16_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_int64.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_int64_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_int64_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[Int64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_string.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_string_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_option_1_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_string_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[String?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_uint.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_uint16.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint16_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint16_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt16?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/fixedarrays_uint64.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_output_uint64_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_option_0"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_option_1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_option_1_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_option_1_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_option_2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_option_3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_fixedarray_input_uint64_option_4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "val"}}, Type: &ast.Ident{Name: "FixedArray[UInt64?]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/hostfns.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "add"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Int"}},
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "echo1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "echo2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "echo3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "echo4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "encode_strings1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "items"}}, Type: &ast.Ident{Name: "Array[String]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "encode_strings2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "items"}}, Type: &ast.Ident{Name: "Array[String?]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/http.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "HttpResponse"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "status"}}, Type: &ast.Ident{Name: "UInt16"}},
										{Names: []*ast.Ident{{Name: "statusText"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "headers"}}, Type: &ast.Ident{Name: "HttpHeaders?"}},
										{Names: []*ast.Ident{{Name: "body"}}, Type: &ast.Ident{Name: "Array[Byte]"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "HttpHeaders"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "data"}}, Type: &ast.Ident{Name: "Map[String, HttpHeader?]"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "HttpHeader"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "name"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "values"}}, Type: &ast.Ident{Name: "Array[String]"}},
									},
								},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_http_response_headers"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "r"}}, Type: &ast.Ident{Name: "HttpResponse?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_http_response_headers_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "HttpResponse?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_http_headers"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "h"}}, Type: &ast.Ident{Name: "HttpHeaders"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_http_header_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, HttpHeader?]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_http_header"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "h"}}, Type: &ast.Ident{Name: "HttpHeader?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/imports_notwasm.mbt"},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/imports_wasm.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "host_echo1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "host_echo2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "host_echo3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "host_echo4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "//go:wasmimport modus_test add"},
						},
					},
					Name: &ast.Ident{Name: "modus_test.add"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Int"}},
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "//go:wasmimport modus_test echo1"},
						},
					},
					Name: &ast.Ident{Name: "modus_test.echo1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "//go:wasmimport modus_test echo2"},
						},
					},
					Name: &ast.Ident{Name: "modus_test.echo2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "//go:wasmimport modus_test echo3"},
						},
					},
					Name: &ast.Ident{Name: "modus_test.echo3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "//go:wasmimport modus_test echo4"},
						},
					},
					Name: &ast.Ident{Name: "modus_test.echo4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "message"}}, Type: &ast.Ident{Name: "String?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "//go:wasmimport modus_test encodeStrings1"},
						},
					},
					Name: &ast.Ident{Name: "modus_test.encodeStrings1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "items"}}, Type: &ast.Ident{Name: "Array[String]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc: &ast.CommentGroup{
						List: []*ast.Comment{
							{Text: "//go:wasmimport modus_test encodeStrings2"},
						},
					},
					Name: &ast.Ident{Name: "modus_test.encodeStrings2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "items"}}, Type: &ast.Ident{Name: "Array[String?]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/maps.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStructWithMap"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]"}},
									},
								},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_input_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_option_input_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_output_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_option_output_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_iterate_map_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_lookup_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[String, String]"}},
								{Names: []*ast.Ident{{Name: "key"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_containing_map_input_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "s"}}, Type: &ast.Ident{Name: "TestStructWithMap"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_containing_map_output_string_string"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStructWithMap"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_input_int_float"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[Int, Float]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_output_int_float"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[Int, Float]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_input_int_double"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "m"}}, Type: &ast.Ident{Name: "Map[Int, Double]"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_map_output_int_double"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[Int, Double]"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_generate_map_string_string_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Map[String, String]"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_bool.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_bool_input_false"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Bool"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_bool_input_true"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Bool"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_bool_output_false"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Bool"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_bool_output_true"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Bool"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_bool_option_input_false"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Bool?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_bool_option_input_true"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Bool?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_bool_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Bool?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_bool_option_output_false"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Bool?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_bool_option_output_true"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Bool?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_bool_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Bool?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_byte.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_byte_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Byte"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_byte_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Byte"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_byte_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Byte"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_byte_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Byte"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_byte_option_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Byte?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_byte_option_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Byte?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_byte_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Byte?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_byte_option_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Byte?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_byte_option_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Byte?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_byte_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Byte?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_char.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_char_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "Char"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_char_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "Char"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_char_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Char"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_char_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Char"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_char_option_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "Char?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_char_option_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "Char?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_char_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "Char?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_char_option_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Char?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_char_option_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Char?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_char_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Char?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_double.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_double_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Double"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_double_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Double"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_double_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Double"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_double_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Double"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_double_option_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Double?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_double_option_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Double?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_double_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Double?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_double_option_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Double?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_double_option_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Double?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_double_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Double?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_float.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_float_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Float"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_float_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Float"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_float_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Float"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_float_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Float"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_float_option_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Float?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_float_option_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Float?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_float_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Float?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_float_option_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Float?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_float_option_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Float?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_float_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Float?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_int.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int_option_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int_option_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int_option_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int_option_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_int16.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int16_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int16"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int16_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int16"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int16_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int16"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int16_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int16"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int16_option_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int16?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int16_option_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int16?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int16_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int16?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int16_option_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int16?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int16_option_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int16?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int16_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int16?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_int64.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int64_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int64"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int64_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int64"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int64_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int64"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int64_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int64"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int64_option_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int64?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int64_option_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int64?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int64_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "Int64?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int64_option_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int64?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int64_option_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int64?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_int64_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Int64?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_uint.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint_option_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint_option_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint_option_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint_option_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_uint16.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint16_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt16"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint16_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt16"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint16_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt16"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint16_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt16"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint16_option_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt16?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint16_option_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt16?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint16_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt16?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint16_option_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt16?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint16_option_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt16?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint16_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt16?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/primitives_uint64.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint64_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt64"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint64_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt64"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint64_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt64"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint64_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt64"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint64_option_input_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt64?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint64_option_input_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt64?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint64_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "n"}}, Type: &ast.Ident{Name: "UInt64?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint64_option_output_min"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt64?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint64_option_output_max"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt64?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_uint64_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "UInt64?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/strings.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_string_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "s"}}, Type: &ast.Ident{Name: "String"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_string_option_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "s"}}, Type: &ast.Ident{Name: "String?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_string_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "s"}}, Type: &ast.Ident{Name: "String?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_string_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_string_option_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_string_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "String?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/structs.mbt"},
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStruct1"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStruct1_map"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStruct2"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStruct2_map"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStruct3"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
										{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStruct3_map"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
										{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStruct4"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
										{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String?"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStruct4_map"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "Int"}},
										{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String?"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStruct5"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "Array[String]"}},
										{Names: []*ast.Ident{{Name: "e"}}, Type: &ast.Ident{Name: "Double"}},
										{Names: []*ast.Ident{{Name: "f"}}, Type: &ast.Ident{Name: "Double"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestStruct5_map"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "b"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "Array[String]"}},
										{Names: []*ast.Ident{{Name: "e"}}, Type: &ast.Ident{Name: "Double"}},
										{Names: []*ast.Ident{{Name: "f"}}, Type: &ast.Ident{Name: "Double"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestRecursiveStruct"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "mut b"}}, Type: &ast.Ident{Name: "TestRecursiveStruct?"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestRecursiveStruct_map"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "a"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "mut b"}}, Type: &ast.Ident{Name: "TestRecursiveStruct_map?"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestSmorgasbordStruct"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "bool"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "byte"}}, Type: &ast.Ident{Name: "Byte"}},
										{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "Char"}},
										{Names: []*ast.Ident{{Name: "f"}}, Type: &ast.Ident{Name: "Float"}},
										{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "Double"}},
										{Names: []*ast.Ident{{Name: "i16"}}, Type: &ast.Ident{Name: "Int16"}},
										{Names: []*ast.Ident{{Name: "i32"}}, Type: &ast.Ident{Name: "Int"}},
										{Names: []*ast.Ident{{Name: "i64"}}, Type: &ast.Ident{Name: "Int64"}},
										{Names: []*ast.Ident{{Name: "s"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "u16"}}, Type: &ast.Ident{Name: "UInt16"}},
										{Names: []*ast.Ident{{Name: "u32"}}, Type: &ast.Ident{Name: "UInt"}},
										{Names: []*ast.Ident{{Name: "u64"}}, Type: &ast.Ident{Name: "UInt64"}},
										{Names: []*ast.Ident{{Name: "someBool"}}, Type: &ast.Ident{Name: "Bool?"}},
										{Names: []*ast.Ident{{Name: "noneBool"}}, Type: &ast.Ident{Name: "Bool?"}},
										{Names: []*ast.Ident{{Name: "someByte"}}, Type: &ast.Ident{Name: "Byte?"}},
										{Names: []*ast.Ident{{Name: "noneByte"}}, Type: &ast.Ident{Name: "Byte?"}},
										{Names: []*ast.Ident{{Name: "someChar"}}, Type: &ast.Ident{Name: "Char?"}},
										{Names: []*ast.Ident{{Name: "noneChar"}}, Type: &ast.Ident{Name: "Char?"}},
										{Names: []*ast.Ident{{Name: "someFloat"}}, Type: &ast.Ident{Name: "Float?"}},
										{Names: []*ast.Ident{{Name: "noneFloat"}}, Type: &ast.Ident{Name: "Float?"}},
										{Names: []*ast.Ident{{Name: "someDouble"}}, Type: &ast.Ident{Name: "Double?"}},
										{Names: []*ast.Ident{{Name: "noneDouble"}}, Type: &ast.Ident{Name: "Double?"}},
										{Names: []*ast.Ident{{Name: "someI16"}}, Type: &ast.Ident{Name: "Int16?"}},
										{Names: []*ast.Ident{{Name: "noneI16"}}, Type: &ast.Ident{Name: "Int16?"}},
										{Names: []*ast.Ident{{Name: "someI32"}}, Type: &ast.Ident{Name: "Int?"}},
										{Names: []*ast.Ident{{Name: "noneI32"}}, Type: &ast.Ident{Name: "Int?"}},
										{Names: []*ast.Ident{{Name: "someI64"}}, Type: &ast.Ident{Name: "Int64?"}},
										{Names: []*ast.Ident{{Name: "noneI64"}}, Type: &ast.Ident{Name: "Int64?"}},
										{Names: []*ast.Ident{{Name: "someString"}}, Type: &ast.Ident{Name: "String?"}},
										{Names: []*ast.Ident{{Name: "noneString"}}, Type: &ast.Ident{Name: "String?"}},
										{Names: []*ast.Ident{{Name: "someU16"}}, Type: &ast.Ident{Name: "UInt16?"}},
										{Names: []*ast.Ident{{Name: "noneU16"}}, Type: &ast.Ident{Name: "UInt16?"}},
										{Names: []*ast.Ident{{Name: "someU32"}}, Type: &ast.Ident{Name: "UInt?"}},
										{Names: []*ast.Ident{{Name: "noneU32"}}, Type: &ast.Ident{Name: "UInt?"}},
										{Names: []*ast.Ident{{Name: "someU64"}}, Type: &ast.Ident{Name: "UInt64?"}},
										{Names: []*ast.Ident{{Name: "noneU64"}}, Type: &ast.Ident{Name: "UInt64?"}},
									},
								},
							},
						},
					},
				},
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{Name: &ast.Ident{Name: "TestSmorgasbordStruct_map"},
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										{Names: []*ast.Ident{{Name: "bool"}}, Type: &ast.Ident{Name: "Bool"}},
										{Names: []*ast.Ident{{Name: "byte"}}, Type: &ast.Ident{Name: "Byte"}},
										{Names: []*ast.Ident{{Name: "c"}}, Type: &ast.Ident{Name: "Char"}},
										{Names: []*ast.Ident{{Name: "f"}}, Type: &ast.Ident{Name: "Float"}},
										{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "Double"}},
										{Names: []*ast.Ident{{Name: "i16"}}, Type: &ast.Ident{Name: "Int16"}},
										{Names: []*ast.Ident{{Name: "i32"}}, Type: &ast.Ident{Name: "Int"}},
										{Names: []*ast.Ident{{Name: "i64"}}, Type: &ast.Ident{Name: "Int64"}},
										{Names: []*ast.Ident{{Name: "s"}}, Type: &ast.Ident{Name: "String"}},
										{Names: []*ast.Ident{{Name: "u16"}}, Type: &ast.Ident{Name: "UInt16"}},
										{Names: []*ast.Ident{{Name: "u32"}}, Type: &ast.Ident{Name: "UInt"}},
										{Names: []*ast.Ident{{Name: "u64"}}, Type: &ast.Ident{Name: "UInt64"}},
										{Names: []*ast.Ident{{Name: "someBool"}}, Type: &ast.Ident{Name: "Bool?"}},
										{Names: []*ast.Ident{{Name: "noneBool"}}, Type: &ast.Ident{Name: "Bool?"}},
										{Names: []*ast.Ident{{Name: "someByte"}}, Type: &ast.Ident{Name: "Byte?"}},
										{Names: []*ast.Ident{{Name: "noneByte"}}, Type: &ast.Ident{Name: "Byte?"}},
										{Names: []*ast.Ident{{Name: "someChar"}}, Type: &ast.Ident{Name: "Char?"}},
										{Names: []*ast.Ident{{Name: "noneChar"}}, Type: &ast.Ident{Name: "Char?"}},
										{Names: []*ast.Ident{{Name: "someFloat"}}, Type: &ast.Ident{Name: "Float?"}},
										{Names: []*ast.Ident{{Name: "noneFloat"}}, Type: &ast.Ident{Name: "Float?"}},
										{Names: []*ast.Ident{{Name: "someDouble"}}, Type: &ast.Ident{Name: "Double?"}},
										{Names: []*ast.Ident{{Name: "noneDouble"}}, Type: &ast.Ident{Name: "Double?"}},
										{Names: []*ast.Ident{{Name: "someI16"}}, Type: &ast.Ident{Name: "Int16?"}},
										{Names: []*ast.Ident{{Name: "noneI16"}}, Type: &ast.Ident{Name: "Int16?"}},
										{Names: []*ast.Ident{{Name: "someI32"}}, Type: &ast.Ident{Name: "Int?"}},
										{Names: []*ast.Ident{{Name: "noneI32"}}, Type: &ast.Ident{Name: "Int?"}},
										{Names: []*ast.Ident{{Name: "someI64"}}, Type: &ast.Ident{Name: "Int64?"}},
										{Names: []*ast.Ident{{Name: "noneI64"}}, Type: &ast.Ident{Name: "Int64?"}},
										{Names: []*ast.Ident{{Name: "someString"}}, Type: &ast.Ident{Name: "String?"}},
										{Names: []*ast.Ident{{Name: "noneString"}}, Type: &ast.Ident{Name: "String?"}},
										{Names: []*ast.Ident{{Name: "someU16"}}, Type: &ast.Ident{Name: "UInt16?"}},
										{Names: []*ast.Ident{{Name: "noneU16"}}, Type: &ast.Ident{Name: "UInt16?"}},
										{Names: []*ast.Ident{{Name: "someU32"}}, Type: &ast.Ident{Name: "UInt?"}},
										{Names: []*ast.Ident{{Name: "noneU32"}}, Type: &ast.Ident{Name: "UInt?"}},
										{Names: []*ast.Ident{{Name: "someU64"}}, Type: &ast.Ident{Name: "UInt64?"}},
										{Names: []*ast.Ident{{Name: "noneU64"}}, Type: &ast.Ident{Name: "UInt64?"}},
									},
								},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_input1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct1"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_input2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct2"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_input3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct3"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_input4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct4"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_input5"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct5"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_input4_with_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct4"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_smorgasbord_struct_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestSmorgasbordStruct"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_recursive_struct_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "r1"}}, Type: &ast.Ident{Name: "TestRecursiveStruct"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct1?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct2?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct3?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct4?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input5"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct5?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input4_with_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct4?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_recursive_struct_option_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestRecursiveStruct?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_smorgasbord_struct_option_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestSmorgasbordStruct?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct1?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input2_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct2?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input3_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct3?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input4_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct4?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_input5_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestStruct5?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_recursive_struct_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestRecursiveStruct?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_smorgasbord_struct_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "o"}}, Type: &ast.Ident{Name: "TestSmorgasbordStruct?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct1"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output1_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct1_map"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct2"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output2_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct2_map"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct3"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output3_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct3_map"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct4"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output4_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct4_map"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output5"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct5"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output5_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct5_map"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output4_with_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct4"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_output4_map_with_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct4_map"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_recursive_struct_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestRecursiveStruct"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_recursive_struct_output_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestRecursiveStruct_map"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_smorgasbord_struct_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestSmorgasbordStruct"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_smorgasbord_struct_output_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestSmorgasbordStruct_map"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output1"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct1?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output1_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct1_map?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct2?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output2_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct2_map?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output3"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct3?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output3_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct3_map?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output4"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct4?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output4_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct4_map?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output5"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct5?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output5_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct5_map?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output4_with_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct4?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output4_map_with_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct4_map?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_recursive_struct_option_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestRecursiveStruct?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_recursive_struct_option_output_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestRecursiveStruct_map?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_smorgasbord_struct_option_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestSmorgasbordStruct?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_smorgasbord_struct_option_output_map"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestSmorgasbordStruct_map?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output1_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct1?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output2_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct2?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output3_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct3?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output4_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct4?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_struct_option_output5_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestStruct5?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_recursive_struct_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestRecursiveStruct?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_smorgasbord_struct_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestSmorgasbordStruct?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_smorgasbord_struct_option_output_map_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "TestSmorgasbordStruct_map?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/time.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_time_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "t"}}, Type: &ast.Ident{Name: "@time.ZonedDateTime"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_time_option_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "t"}}, Type: &ast.Ident{Name: "@time.ZonedDateTime?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "call_test_time_option_input_some"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "@time.ZonedDateTime?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "call_test_time_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "@time.ZonedDateTime?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_time_option_input_style2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "t?"}}, Type: &ast.Ident{Name: "@time.ZonedDateTime"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_time_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "t"}}, Type: &ast.Ident{Name: "@time.ZonedDateTime?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_time_option_input_none_style2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "t?"}}, Type: &ast.Ident{Name: "@time.ZonedDateTime"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_time_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "@time.ZonedDateTime"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_time_option_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "@time.ZonedDateTime?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_time_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "@time.ZonedDateTime?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_duration_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "@time.Duration"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_duration_option_input"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "@time.Duration?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_duration_option_input_style2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "d?"}}, Type: &ast.Ident{Name: "@time.Duration"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_duration_option_input_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "d"}}, Type: &ast.Ident{Name: "@time.Duration?"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_duration_option_input_none_style2"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								{Names: []*ast.Ident{{Name: "d?"}}, Type: &ast.Ident{Name: "@time.Duration"}},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "Unit!Error"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_duration_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "@time.Duration"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_duration_option_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "@time.Duration?"}},
							},
						},
					},
				},
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_duration_option_output_none"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "@time.Duration?"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/tuples.mbt"},
			Decls: []ast.Decl{
				&ast.FuncDecl{
					Doc:  &ast.CommentGroup{},
					Name: &ast.Ident{Name: "test_tuple_output"},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{Type: &ast.Ident{Name: "(Int, Bool, String)"}},
							},
						},
					},
				},
			},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
		{
			Name: &ast.Ident{Name: "../testdata/runtime-testdata/utils.mbt"},
			Imports: []*ast.ImportSpec{
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/pkg/console"`}},
				{Path: &ast.BasicLit{Value: `"gmlewis/modus/wit/ffi"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/sys"`}},
				{Path: &ast.BasicLit{Value: `"moonbitlang/x/time"`}},
			},
		},
	},
	TypesInfo: &types.Info{
		Defs: map[*ast.Ident]types.Object{
			// Using &moonType{} and &moonFunc{} is a hack to fake a struct/func for testing purposes only:
			{Name: "HttpHeader"}:                                      &moonFunc{funcName: "type HttpHeader = struct{name String; values Array[String]}"},
			{Name: "HttpHeaders"}:                                     &moonFunc{funcName: "type HttpHeaders = struct{data Map[String, HttpHeader?]}"},
			{Name: "HttpHeaders"}:                                     &moonFunc{funcName: "type HttpHeaders = struct{data Map[String, HttpHeader?]}"},
			{Name: "HttpHeaders?"}:                                    &moonFunc{funcName: "type HttpHeaders? = struct{}"},
			{Name: "HttpResponse"}:                                    &moonFunc{funcName: "type HttpResponse = struct{status UInt16; statusText String; headers HttpHeaders?; body Array[Byte]}"},
			{Name: "TestRecursiveStruct"}:                             &moonFunc{funcName: "type TestRecursiveStruct = struct{a Bool; mut b TestRecursiveStruct?}"},
			{Name: "TestRecursiveStruct"}:                             &moonFunc{funcName: "type TestRecursiveStruct = struct{a Bool; mut b TestRecursiveStruct?}"},
			{Name: "TestRecursiveStruct?"}:                            &moonFunc{funcName: "type TestRecursiveStruct? = struct{}"},
			{Name: "TestRecursiveStruct_map"}:                         &moonFunc{funcName: "type TestRecursiveStruct_map = struct{a Bool; mut b TestRecursiveStruct_map?}"},
			{Name: "TestRecursiveStruct_map"}:                         &moonFunc{funcName: "type TestRecursiveStruct_map = struct{a Bool; mut b TestRecursiveStruct_map?}"},
			{Name: "TestRecursiveStruct_map?"}:                        &moonFunc{funcName: "type TestRecursiveStruct_map? = struct{}"},
			{Name: "TestSmorgasbordStruct"}:                           &moonFunc{funcName: "type TestSmorgasbordStruct = struct{bool Bool; byte Byte; c Char; f Float; d Double; i16 Int16; i32 Int; i64 Int64; s String; u16 UInt16; u32 UInt; u64 UInt64; someBool Bool?; noneBool Bool?; someByte Byte?; noneByte Byte?; someChar Char?; noneChar Char?; someFloat Float?; noneFloat Float?; someDouble Double?; noneDouble Double?; someI16 Int16?; noneI16 Int16?; someI32 Int?; noneI32 Int?; someI64 Int64?; noneI64 Int64?; someString String?; noneString String?; someU16 UInt16?; noneU16 UInt16?; someU32 UInt?; noneU32 UInt?; someU64 UInt64?; noneU64 UInt64?}"},
			{Name: "TestSmorgasbordStruct_map"}:                       &moonFunc{funcName: "type TestSmorgasbordStruct_map = struct{bool Bool; byte Byte; c Char; f Float; d Double; i16 Int16; i32 Int; i64 Int64; s String; u16 UInt16; u32 UInt; u64 UInt64; someBool Bool?; noneBool Bool?; someByte Byte?; noneByte Byte?; someChar Char?; noneChar Char?; someFloat Float?; noneFloat Float?; someDouble Double?; noneDouble Double?; someI16 Int16?; noneI16 Int16?; someI32 Int?; noneI32 Int?; someI64 Int64?; noneI64 Int64?; someString String?; noneString String?; someU16 UInt16?; noneU16 UInt16?; someU32 UInt?; noneU32 UInt?; someU64 UInt64?; noneU64 UInt64?}"},
			{Name: "TestStruct1"}:                                     &moonFunc{funcName: "type TestStruct1 = struct{a Bool}"},
			{Name: "TestStruct1_map"}:                                 &moonFunc{funcName: "type TestStruct1_map = struct{a Bool}"},
			{Name: "TestStruct2"}:                                     &moonFunc{funcName: "type TestStruct2 = struct{a Bool; b Int}"},
			{Name: "TestStruct2_map"}:                                 &moonFunc{funcName: "type TestStruct2_map = struct{a Bool; b Int}"},
			{Name: "TestStruct3"}:                                     &moonFunc{funcName: "type TestStruct3 = struct{a Bool; b Int; c String}"},
			{Name: "TestStruct3_map"}:                                 &moonFunc{funcName: "type TestStruct3_map = struct{a Bool; b Int; c String}"},
			{Name: "TestStruct4"}:                                     &moonFunc{funcName: "type TestStruct4 = struct{a Bool; b Int; c String?}"},
			{Name: "TestStruct4_map"}:                                 &moonFunc{funcName: "type TestStruct4_map = struct{a Bool; b Int; c String?}"},
			{Name: "TestStruct5"}:                                     &moonFunc{funcName: "type TestStruct5 = struct{a String; b String; c String; d Array[String]; e Double; f Double}"},
			{Name: "TestStruct5_map"}:                                 &moonFunc{funcName: "type TestStruct5_map = struct{a String; b String; c String; d Array[String]; e Double; f Double}"},
			{Name: "TestStructWithMap"}:                               &moonFunc{funcName: "type TestStructWithMap = struct{m Map[String, String]}"},
			{Name: "add"}:                                             &moonFunc{funcName: "func add(a Int, b Int) Int"},
			{Name: "call_test_time_option_input_none"}:                &moonFunc{funcName: "func call_test_time_option_input_none() @time.ZonedDateTime?"},
			{Name: "call_test_time_option_input_some"}:                &moonFunc{funcName: "func call_test_time_option_input_some() @time.ZonedDateTime?"},
			{Name: "echo1"}:                                           &moonFunc{funcName: "func echo1(message String) String"},
			{Name: "echo2"}:                                           &moonFunc{funcName: "func echo2(message String?) String"},
			{Name: "echo3"}:                                           &moonFunc{funcName: "func echo3(message String) String?"},
			{Name: "echo4"}:                                           &moonFunc{funcName: "func echo4(message String?) String?"},
			{Name: "encode_strings1"}:                                 &moonFunc{funcName: "func encode_strings1(items Array[String]?) String?"},
			{Name: "encode_strings2"}:                                 &moonFunc{funcName: "func encode_strings2(items Array[String?]?) String?"},
			{Name: "get_int_option_fixedarray1"}:                      &moonFunc{funcName: "func get_int_option_fixedarray1() FixedArray[Int?]"},
			{Name: "get_int_ptr_array2"}:                              &moonFunc{funcName: "func get_int_ptr_array2() FixedArray[Int?]"},
			{Name: "get_map_array2"}:                                  &moonFunc{funcName: "func get_map_array2() FixedArray[Map[String, String]]"},
			{Name: "get_map_ptr_array2"}:                              &moonFunc{funcName: "func get_map_ptr_array2() FixedArray[Map[String, String]?]"},
			{Name: "get_option_int_fixedarray1"}:                      &moonFunc{funcName: "func get_option_int_fixedarray1() FixedArray[Int]?"},
			{Name: "get_option_int_fixedarray2"}:                      &moonFunc{funcName: "func get_option_int_fixedarray2() FixedArray[Int]?"},
			{Name: "get_option_string_fixedarray1"}:                   &moonFunc{funcName: "func get_option_string_fixedarray1() FixedArray[String]?"},
			{Name: "get_option_string_fixedarray2"}:                   &moonFunc{funcName: "func get_option_string_fixedarray2() FixedArray[String]?"},
			{Name: "get_string_option_array2"}:                        &moonFunc{funcName: "func get_string_option_array2() FixedArray[String?]"},
			{Name: "get_string_option_fixedarray1"}:                   &moonFunc{funcName: "func get_string_option_fixedarray1() FixedArray[String?]"},
			{Name: "host_echo1"}:                                      &moonFunc{funcName: "func host_echo1(message String) String"},
			{Name: "host_echo2"}:                                      &moonFunc{funcName: "func host_echo2(message String?) String"},
			{Name: "host_echo3"}:                                      &moonFunc{funcName: "func host_echo3(message String) String?"},
			{Name: "host_echo4"}:                                      &moonFunc{funcName: "func host_echo4(message String?) String?"},
			{Name: "modus_test.add"}:                                  &moonFunc{funcName: "func modus_test.add(a Int, b Int) Int"},
			{Name: "modus_test.echo1"}:                                &moonFunc{funcName: "func modus_test.echo1(message String) String"},
			{Name: "modus_test.echo2"}:                                &moonFunc{funcName: "func modus_test.echo2(message String?) String"},
			{Name: "modus_test.echo3"}:                                &moonFunc{funcName: "func modus_test.echo3(message String) String?"},
			{Name: "modus_test.echo4"}:                                &moonFunc{funcName: "func modus_test.echo4(message String?) String?"},
			{Name: "modus_test.encodeStrings1"}:                       &moonFunc{funcName: "func modus_test.encodeStrings1(items Array[String]?) String?"},
			{Name: "modus_test.encodeStrings2"}:                       &moonFunc{funcName: "func modus_test.encodeStrings2(items Array[String?]?) String?"},
			{Name: "test2d_array_input_string"}:                       &moonFunc{funcName: "func test2d_array_input_string(val Array[Array[String]]) Unit!Error"},
			{Name: "test2d_array_input_string_empty"}:                 &moonFunc{funcName: "func test2d_array_input_string_empty(val Array[Array[String]]) Unit!Error"},
			{Name: "test2d_array_input_string_inner_empty"}:           &moonFunc{funcName: "func test2d_array_input_string_inner_empty(val Array[Array[String]]) Unit!Error"},
			{Name: "test2d_array_input_string_inner_none"}:            &moonFunc{funcName: "func test2d_array_input_string_inner_none(val Array[Array[String]?]) Unit!Error"},
			{Name: "test2d_array_input_string_none"}:                  &moonFunc{funcName: "func test2d_array_input_string_none(val Array[Array[String]]?) Unit!Error"},
			{Name: "test2d_array_output_string"}:                      &moonFunc{funcName: "func test2d_array_output_string() Array[Array[String]]"},
			{Name: "test2d_array_output_string_empty"}:                &moonFunc{funcName: "func test2d_array_output_string_empty() Array[Array[String]]"},
			{Name: "test2d_array_output_string_inner_empty"}:          &moonFunc{funcName: "func test2d_array_output_string_inner_empty() Array[Array[String]]"},
			{Name: "test2d_array_output_string_inner_none"}:           &moonFunc{funcName: "func test2d_array_output_string_inner_none() Array[Array[String]?]"},
			{Name: "test2d_array_output_string_none"}:                 &moonFunc{funcName: "func test2d_array_output_string_none() Array[Array[String]]?"},
			{Name: "test_array_input_bool_0"}:                         &moonFunc{funcName: "func test_array_input_bool_0(val Array[Bool]) Unit!Error"},
			{Name: "test_array_input_bool_1"}:                         &moonFunc{funcName: "func test_array_input_bool_1(val Array[Bool]) Unit!Error"},
			{Name: "test_array_input_bool_2"}:                         &moonFunc{funcName: "func test_array_input_bool_2(val Array[Bool]) Unit!Error"},
			{Name: "test_array_input_bool_3"}:                         &moonFunc{funcName: "func test_array_input_bool_3(val Array[Bool]) Unit!Error"},
			{Name: "test_array_input_bool_4"}:                         &moonFunc{funcName: "func test_array_input_bool_4(val Array[Bool]) Unit!Error"},
			{Name: "test_array_input_bool_option_0"}:                  &moonFunc{funcName: "func test_array_input_bool_option_0(val Array[Bool?]) Unit!Error"},
			{Name: "test_array_input_bool_option_1_false"}:            &moonFunc{funcName: "func test_array_input_bool_option_1_false(val Array[Bool?]) Unit!Error"},
			{Name: "test_array_input_bool_option_1_none"}:             &moonFunc{funcName: "func test_array_input_bool_option_1_none(val Array[Bool?]) Unit!Error"},
			{Name: "test_array_input_bool_option_1_true"}:             &moonFunc{funcName: "func test_array_input_bool_option_1_true(val Array[Bool?]) Unit!Error"},
			{Name: "test_array_input_bool_option_2"}:                  &moonFunc{funcName: "func test_array_input_bool_option_2(val Array[Bool?]) Unit!Error"},
			{Name: "test_array_input_bool_option_3"}:                  &moonFunc{funcName: "func test_array_input_bool_option_3(val Array[Bool?]) Unit!Error"},
			{Name: "test_array_input_bool_option_4"}:                  &moonFunc{funcName: "func test_array_input_bool_option_4(val Array[Bool?]) Unit!Error"},
			{Name: "test_array_input_byte_0"}:                         &moonFunc{funcName: "func test_array_input_byte_0(val Array[Byte]) Unit!Error"},
			{Name: "test_array_input_byte_1"}:                         &moonFunc{funcName: "func test_array_input_byte_1(val Array[Byte]) Unit!Error"},
			{Name: "test_array_input_byte_2"}:                         &moonFunc{funcName: "func test_array_input_byte_2(val Array[Byte]) Unit!Error"},
			{Name: "test_array_input_byte_3"}:                         &moonFunc{funcName: "func test_array_input_byte_3(val Array[Byte]) Unit!Error"},
			{Name: "test_array_input_byte_4"}:                         &moonFunc{funcName: "func test_array_input_byte_4(val Array[Byte]) Unit!Error"},
			{Name: "test_array_input_byte_option_0"}:                  &moonFunc{funcName: "func test_array_input_byte_option_0(val Array[Byte?]) Unit!Error"},
			{Name: "test_array_input_byte_option_1"}:                  &moonFunc{funcName: "func test_array_input_byte_option_1(val Array[Byte?]) Unit!Error"},
			{Name: "test_array_input_byte_option_2"}:                  &moonFunc{funcName: "func test_array_input_byte_option_2(val Array[Byte?]) Unit!Error"},
			{Name: "test_array_input_byte_option_3"}:                  &moonFunc{funcName: "func test_array_input_byte_option_3(val Array[Byte?]) Unit!Error"},
			{Name: "test_array_input_byte_option_4"}:                  &moonFunc{funcName: "func test_array_input_byte_option_4(val Array[Byte?]) Unit!Error"},
			{Name: "test_array_input_char_empty"}:                     &moonFunc{funcName: "func test_array_input_char_empty(val Array[Char]) Unit!Error"},
			{Name: "test_array_input_char_option"}:                    &moonFunc{funcName: "func test_array_input_char_option(val Array[Char?]) Unit!Error"},
			{Name: "test_array_input_double_empty"}:                   &moonFunc{funcName: "func test_array_input_double_empty(val Array[Double]) Unit!Error"},
			{Name: "test_array_input_double_option"}:                  &moonFunc{funcName: "func test_array_input_double_option(val Array[Double?]) Unit!Error"},
			{Name: "test_array_input_float_empty"}:                    &moonFunc{funcName: "func test_array_input_float_empty(val Array[Float]) Unit!Error"},
			{Name: "test_array_input_float_option"}:                   &moonFunc{funcName: "func test_array_input_float_option(val Array[Float?]) Unit!Error"},
			{Name: "test_array_input_int16_empty"}:                    &moonFunc{funcName: "func test_array_input_int16_empty(val Array[Int16]) Unit!Error"},
			{Name: "test_array_input_int16_option"}:                   &moonFunc{funcName: "func test_array_input_int16_option(val Array[Int16?]) Unit!Error"},
			{Name: "test_array_input_int_empty"}:                      &moonFunc{funcName: "func test_array_input_int_empty(val Array[Int]) Unit!Error"},
			{Name: "test_array_input_int_option"}:                     &moonFunc{funcName: "func test_array_input_int_option(val Array[Int?]) Unit!Error"},
			{Name: "test_array_input_string"}:                         &moonFunc{funcName: "func test_array_input_string(val Array[String]) Unit!Error"},
			{Name: "test_array_input_string_empty"}:                   &moonFunc{funcName: "func test_array_input_string_empty(val Array[String]) Unit!Error"},
			{Name: "test_array_input_string_none"}:                    &moonFunc{funcName: "func test_array_input_string_none(val Array[String]?) Unit!Error"},
			{Name: "test_array_input_string_option"}:                  &moonFunc{funcName: "func test_array_input_string_option(val Array[String?]) Unit!Error"},
			{Name: "test_array_input_uint16_empty"}:                   &moonFunc{funcName: "func test_array_input_uint16_empty(val Array[UInt16]) Unit!Error"},
			{Name: "test_array_input_uint16_option"}:                  &moonFunc{funcName: "func test_array_input_uint16_option(val Array[UInt16?]) Unit!Error"},
			{Name: "test_array_input_uint_empty"}:                     &moonFunc{funcName: "func test_array_input_uint_empty(val Array[UInt]) Unit!Error"},
			{Name: "test_array_input_uint_option"}:                    &moonFunc{funcName: "func test_array_input_uint_option(val Array[UInt?]) Unit!Error"},
			{Name: "test_array_output_bool_0"}:                        &moonFunc{funcName: "func test_array_output_bool_0() Array[Bool]"},
			{Name: "test_array_output_bool_1"}:                        &moonFunc{funcName: "func test_array_output_bool_1() Array[Bool]"},
			{Name: "test_array_output_bool_2"}:                        &moonFunc{funcName: "func test_array_output_bool_2() Array[Bool]"},
			{Name: "test_array_output_bool_3"}:                        &moonFunc{funcName: "func test_array_output_bool_3() Array[Bool]"},
			{Name: "test_array_output_bool_4"}:                        &moonFunc{funcName: "func test_array_output_bool_4() Array[Bool]"},
			{Name: "test_array_output_bool_option_0"}:                 &moonFunc{funcName: "func test_array_output_bool_option_0() Array[Bool?]"},
			{Name: "test_array_output_bool_option_1_false"}:           &moonFunc{funcName: "func test_array_output_bool_option_1_false() Array[Bool?]"},
			{Name: "test_array_output_bool_option_1_none"}:            &moonFunc{funcName: "func test_array_output_bool_option_1_none() Array[Bool?]"},
			{Name: "test_array_output_bool_option_1_true"}:            &moonFunc{funcName: "func test_array_output_bool_option_1_true() Array[Bool?]"},
			{Name: "test_array_output_bool_option_2"}:                 &moonFunc{funcName: "func test_array_output_bool_option_2() Array[Bool?]"},
			{Name: "test_array_output_bool_option_3"}:                 &moonFunc{funcName: "func test_array_output_bool_option_3() Array[Bool?]"},
			{Name: "test_array_output_bool_option_4"}:                 &moonFunc{funcName: "func test_array_output_bool_option_4() Array[Bool?]"},
			{Name: "test_array_output_byte_0"}:                        &moonFunc{funcName: "func test_array_output_byte_0() Array[Byte]"},
			{Name: "test_array_output_byte_1"}:                        &moonFunc{funcName: "func test_array_output_byte_1() Array[Byte]"},
			{Name: "test_array_output_byte_2"}:                        &moonFunc{funcName: "func test_array_output_byte_2() Array[Byte]"},
			{Name: "test_array_output_byte_3"}:                        &moonFunc{funcName: "func test_array_output_byte_3() Array[Byte]"},
			{Name: "test_array_output_byte_4"}:                        &moonFunc{funcName: "func test_array_output_byte_4() Array[Byte]"},
			{Name: "test_array_output_byte_option_0"}:                 &moonFunc{funcName: "func test_array_output_byte_option_0() Array[Byte?]"},
			{Name: "test_array_output_byte_option_1"}:                 &moonFunc{funcName: "func test_array_output_byte_option_1() Array[Byte?]"},
			{Name: "test_array_output_byte_option_2"}:                 &moonFunc{funcName: "func test_array_output_byte_option_2() Array[Byte?]"},
			{Name: "test_array_output_byte_option_3"}:                 &moonFunc{funcName: "func test_array_output_byte_option_3() Array[Byte?]"},
			{Name: "test_array_output_byte_option_4"}:                 &moonFunc{funcName: "func test_array_output_byte_option_4() Array[Byte?]"},
			{Name: "test_array_output_char_0"}:                        &moonFunc{funcName: "func test_array_output_char_0() Array[Char]"},
			{Name: "test_array_output_char_1"}:                        &moonFunc{funcName: "func test_array_output_char_1() Array[Char]"},
			{Name: "test_array_output_char_2"}:                        &moonFunc{funcName: "func test_array_output_char_2() Array[Char]"},
			{Name: "test_array_output_char_3"}:                        &moonFunc{funcName: "func test_array_output_char_3() Array[Char]"},
			{Name: "test_array_output_char_4"}:                        &moonFunc{funcName: "func test_array_output_char_4() Array[Char]"},
			{Name: "test_array_output_char_option"}:                   &moonFunc{funcName: "func test_array_output_char_option() Array[Char?]"},
			{Name: "test_array_output_char_option_0"}:                 &moonFunc{funcName: "func test_array_output_char_option_0() Array[Char?]"},
			{Name: "test_array_output_char_option_1_none"}:            &moonFunc{funcName: "func test_array_output_char_option_1_none() Array[Char?]"},
			{Name: "test_array_output_char_option_1_some"}:            &moonFunc{funcName: "func test_array_output_char_option_1_some() Array[Char?]"},
			{Name: "test_array_output_char_option_2"}:                 &moonFunc{funcName: "func test_array_output_char_option_2() Array[Char?]"},
			{Name: "test_array_output_char_option_3"}:                 &moonFunc{funcName: "func test_array_output_char_option_3() Array[Char?]"},
			{Name: "test_array_output_char_option_4"}:                 &moonFunc{funcName: "func test_array_output_char_option_4() Array[Char?]"},
			{Name: "test_array_output_double_0"}:                      &moonFunc{funcName: "func test_array_output_double_0() Array[Double]"},
			{Name: "test_array_output_double_1"}:                      &moonFunc{funcName: "func test_array_output_double_1() Array[Double]"},
			{Name: "test_array_output_double_2"}:                      &moonFunc{funcName: "func test_array_output_double_2() Array[Double]"},
			{Name: "test_array_output_double_3"}:                      &moonFunc{funcName: "func test_array_output_double_3() Array[Double]"},
			{Name: "test_array_output_double_4"}:                      &moonFunc{funcName: "func test_array_output_double_4() Array[Double]"},
			{Name: "test_array_output_double_option"}:                 &moonFunc{funcName: "func test_array_output_double_option() Array[Double?]"},
			{Name: "test_array_output_double_option_0"}:               &moonFunc{funcName: "func test_array_output_double_option_0() Array[Double?]"},
			{Name: "test_array_output_double_option_1_none"}:          &moonFunc{funcName: "func test_array_output_double_option_1_none() Array[Double?]"},
			{Name: "test_array_output_double_option_1_some"}:          &moonFunc{funcName: "func test_array_output_double_option_1_some() Array[Double?]"},
			{Name: "test_array_output_double_option_2"}:               &moonFunc{funcName: "func test_array_output_double_option_2() Array[Double?]"},
			{Name: "test_array_output_double_option_3"}:               &moonFunc{funcName: "func test_array_output_double_option_3() Array[Double?]"},
			{Name: "test_array_output_double_option_4"}:               &moonFunc{funcName: "func test_array_output_double_option_4() Array[Double?]"},
			{Name: "test_array_output_float_0"}:                       &moonFunc{funcName: "func test_array_output_float_0() Array[Float]"},
			{Name: "test_array_output_float_1"}:                       &moonFunc{funcName: "func test_array_output_float_1() Array[Float]"},
			{Name: "test_array_output_float_2"}:                       &moonFunc{funcName: "func test_array_output_float_2() Array[Float]"},
			{Name: "test_array_output_float_3"}:                       &moonFunc{funcName: "func test_array_output_float_3() Array[Float]"},
			{Name: "test_array_output_float_4"}:                       &moonFunc{funcName: "func test_array_output_float_4() Array[Float]"},
			{Name: "test_array_output_float_option"}:                  &moonFunc{funcName: "func test_array_output_float_option() Array[Float?]"},
			{Name: "test_array_output_float_option_0"}:                &moonFunc{funcName: "func test_array_output_float_option_0() Array[Float?]"},
			{Name: "test_array_output_float_option_1_none"}:           &moonFunc{funcName: "func test_array_output_float_option_1_none() Array[Float?]"},
			{Name: "test_array_output_float_option_1_some"}:           &moonFunc{funcName: "func test_array_output_float_option_1_some() Array[Float?]"},
			{Name: "test_array_output_float_option_2"}:                &moonFunc{funcName: "func test_array_output_float_option_2() Array[Float?]"},
			{Name: "test_array_output_float_option_3"}:                &moonFunc{funcName: "func test_array_output_float_option_3() Array[Float?]"},
			{Name: "test_array_output_float_option_4"}:                &moonFunc{funcName: "func test_array_output_float_option_4() Array[Float?]"},
			{Name: "test_array_output_int16_0"}:                       &moonFunc{funcName: "func test_array_output_int16_0() Array[Int16]"},
			{Name: "test_array_output_int16_1"}:                       &moonFunc{funcName: "func test_array_output_int16_1() Array[Int16]"},
			{Name: "test_array_output_int16_1_max"}:                   &moonFunc{funcName: "func test_array_output_int16_1_max() Array[Int16]"},
			{Name: "test_array_output_int16_1_min"}:                   &moonFunc{funcName: "func test_array_output_int16_1_min() Array[Int16]"},
			{Name: "test_array_output_int16_2"}:                       &moonFunc{funcName: "func test_array_output_int16_2() Array[Int16]"},
			{Name: "test_array_output_int16_3"}:                       &moonFunc{funcName: "func test_array_output_int16_3() Array[Int16]"},
			{Name: "test_array_output_int16_4"}:                       &moonFunc{funcName: "func test_array_output_int16_4() Array[Int16]"},
			{Name: "test_array_output_int16_option"}:                  &moonFunc{funcName: "func test_array_output_int16_option() Array[Int16?]"},
			{Name: "test_array_output_int16_option_0"}:                &moonFunc{funcName: "func test_array_output_int16_option_0() Array[Int16?]"},
			{Name: "test_array_output_int16_option_1_max"}:            &moonFunc{funcName: "func test_array_output_int16_option_1_max() Array[Int16?]"},
			{Name: "test_array_output_int16_option_1_min"}:            &moonFunc{funcName: "func test_array_output_int16_option_1_min() Array[Int16?]"},
			{Name: "test_array_output_int16_option_1_none"}:           &moonFunc{funcName: "func test_array_output_int16_option_1_none() Array[Int16?]"},
			{Name: "test_array_output_int16_option_2"}:                &moonFunc{funcName: "func test_array_output_int16_option_2() Array[Int16?]"},
			{Name: "test_array_output_int16_option_3"}:                &moonFunc{funcName: "func test_array_output_int16_option_3() Array[Int16?]"},
			{Name: "test_array_output_int16_option_4"}:                &moonFunc{funcName: "func test_array_output_int16_option_4() Array[Int16?]"},
			{Name: "test_array_output_int64_0"}:                       &moonFunc{funcName: "func test_array_output_int64_0() Array[Int64]"},
			{Name: "test_array_output_int64_1"}:                       &moonFunc{funcName: "func test_array_output_int64_1() Array[Int64]"},
			{Name: "test_array_output_int64_1_max"}:                   &moonFunc{funcName: "func test_array_output_int64_1_max() Array[Int64]"},
			{Name: "test_array_output_int64_1_min"}:                   &moonFunc{funcName: "func test_array_output_int64_1_min() Array[Int64]"},
			{Name: "test_array_output_int64_2"}:                       &moonFunc{funcName: "func test_array_output_int64_2() Array[Int64]"},
			{Name: "test_array_output_int64_3"}:                       &moonFunc{funcName: "func test_array_output_int64_3() Array[Int64]"},
			{Name: "test_array_output_int64_4"}:                       &moonFunc{funcName: "func test_array_output_int64_4() Array[Int64]"},
			{Name: "test_array_output_int64_option_0"}:                &moonFunc{funcName: "func test_array_output_int64_option_0() Array[Int64?]"},
			{Name: "test_array_output_int64_option_1_max"}:            &moonFunc{funcName: "func test_array_output_int64_option_1_max() Array[Int64?]"},
			{Name: "test_array_output_int64_option_1_min"}:            &moonFunc{funcName: "func test_array_output_int64_option_1_min() Array[Int64?]"},
			{Name: "test_array_output_int64_option_1_none"}:           &moonFunc{funcName: "func test_array_output_int64_option_1_none() Array[Int64?]"},
			{Name: "test_array_output_int64_option_2"}:                &moonFunc{funcName: "func test_array_output_int64_option_2() Array[Int64?]"},
			{Name: "test_array_output_int64_option_3"}:                &moonFunc{funcName: "func test_array_output_int64_option_3() Array[Int64?]"},
			{Name: "test_array_output_int64_option_4"}:                &moonFunc{funcName: "func test_array_output_int64_option_4() Array[Int64?]"},
			{Name: "test_array_output_int_0"}:                         &moonFunc{funcName: "func test_array_output_int_0() Array[Int]"},
			{Name: "test_array_output_int_1"}:                         &moonFunc{funcName: "func test_array_output_int_1() Array[Int]"},
			{Name: "test_array_output_int_1_max"}:                     &moonFunc{funcName: "func test_array_output_int_1_max() Array[Int]"},
			{Name: "test_array_output_int_1_min"}:                     &moonFunc{funcName: "func test_array_output_int_1_min() Array[Int]"},
			{Name: "test_array_output_int_2"}:                         &moonFunc{funcName: "func test_array_output_int_2() Array[Int]"},
			{Name: "test_array_output_int_3"}:                         &moonFunc{funcName: "func test_array_output_int_3() Array[Int]"},
			{Name: "test_array_output_int_4"}:                         &moonFunc{funcName: "func test_array_output_int_4() Array[Int]"},
			{Name: "test_array_output_int_option"}:                    &moonFunc{funcName: "func test_array_output_int_option() Array[Int?]"},
			{Name: "test_array_output_int_option_0"}:                  &moonFunc{funcName: "func test_array_output_int_option_0() Array[Int?]"},
			{Name: "test_array_output_int_option_1_max"}:              &moonFunc{funcName: "func test_array_output_int_option_1_max() Array[Int?]"},
			{Name: "test_array_output_int_option_1_min"}:              &moonFunc{funcName: "func test_array_output_int_option_1_min() Array[Int?]"},
			{Name: "test_array_output_int_option_1_none"}:             &moonFunc{funcName: "func test_array_output_int_option_1_none() Array[Int?]"},
			{Name: "test_array_output_int_option_2"}:                  &moonFunc{funcName: "func test_array_output_int_option_2() Array[Int?]"},
			{Name: "test_array_output_int_option_3"}:                  &moonFunc{funcName: "func test_array_output_int_option_3() Array[Int?]"},
			{Name: "test_array_output_int_option_4"}:                  &moonFunc{funcName: "func test_array_output_int_option_4() Array[Int?]"},
			{Name: "test_array_output_string"}:                        &moonFunc{funcName: "func test_array_output_string() Array[String]"},
			{Name: "test_array_output_string_empty"}:                  &moonFunc{funcName: "func test_array_output_string_empty() Array[String]"},
			{Name: "test_array_output_string_none"}:                   &moonFunc{funcName: "func test_array_output_string_none() Array[String]?"},
			{Name: "test_array_output_string_option"}:                 &moonFunc{funcName: "func test_array_output_string_option() Array[String?]"},
			{Name: "test_array_output_uint16_0"}:                      &moonFunc{funcName: "func test_array_output_uint16_0() Array[UInt16]"},
			{Name: "test_array_output_uint16_1"}:                      &moonFunc{funcName: "func test_array_output_uint16_1() Array[UInt16]"},
			{Name: "test_array_output_uint16_1_max"}:                  &moonFunc{funcName: "func test_array_output_uint16_1_max() Array[UInt16]"},
			{Name: "test_array_output_uint16_1_min"}:                  &moonFunc{funcName: "func test_array_output_uint16_1_min() Array[UInt16]"},
			{Name: "test_array_output_uint16_2"}:                      &moonFunc{funcName: "func test_array_output_uint16_2() Array[UInt16]"},
			{Name: "test_array_output_uint16_3"}:                      &moonFunc{funcName: "func test_array_output_uint16_3() Array[UInt16]"},
			{Name: "test_array_output_uint16_4"}:                      &moonFunc{funcName: "func test_array_output_uint16_4() Array[UInt16]"},
			{Name: "test_array_output_uint16_option"}:                 &moonFunc{funcName: "func test_array_output_uint16_option() Array[UInt16?]"},
			{Name: "test_array_output_uint16_option_0"}:               &moonFunc{funcName: "func test_array_output_uint16_option_0() Array[UInt16?]"},
			{Name: "test_array_output_uint16_option_1_max"}:           &moonFunc{funcName: "func test_array_output_uint16_option_1_max() Array[UInt16?]"},
			{Name: "test_array_output_uint16_option_1_min"}:           &moonFunc{funcName: "func test_array_output_uint16_option_1_min() Array[UInt16?]"},
			{Name: "test_array_output_uint16_option_1_none"}:          &moonFunc{funcName: "func test_array_output_uint16_option_1_none() Array[UInt16?]"},
			{Name: "test_array_output_uint16_option_2"}:               &moonFunc{funcName: "func test_array_output_uint16_option_2() Array[UInt16?]"},
			{Name: "test_array_output_uint16_option_3"}:               &moonFunc{funcName: "func test_array_output_uint16_option_3() Array[UInt16?]"},
			{Name: "test_array_output_uint16_option_4"}:               &moonFunc{funcName: "func test_array_output_uint16_option_4() Array[UInt16?]"},
			{Name: "test_array_output_uint64_0"}:                      &moonFunc{funcName: "func test_array_output_uint64_0() Array[UInt64]"},
			{Name: "test_array_output_uint64_1"}:                      &moonFunc{funcName: "func test_array_output_uint64_1() Array[UInt64]"},
			{Name: "test_array_output_uint64_1_max"}:                  &moonFunc{funcName: "func test_array_output_uint64_1_max() Array[UInt64]"},
			{Name: "test_array_output_uint64_1_min"}:                  &moonFunc{funcName: "func test_array_output_uint64_1_min() Array[UInt64]"},
			{Name: "test_array_output_uint64_2"}:                      &moonFunc{funcName: "func test_array_output_uint64_2() Array[UInt64]"},
			{Name: "test_array_output_uint64_3"}:                      &moonFunc{funcName: "func test_array_output_uint64_3() Array[UInt64]"},
			{Name: "test_array_output_uint64_4"}:                      &moonFunc{funcName: "func test_array_output_uint64_4() Array[UInt64]"},
			{Name: "test_array_output_uint64_option_0"}:               &moonFunc{funcName: "func test_array_output_uint64_option_0() Array[UInt64?]"},
			{Name: "test_array_output_uint64_option_1_max"}:           &moonFunc{funcName: "func test_array_output_uint64_option_1_max() Array[UInt64?]"},
			{Name: "test_array_output_uint64_option_1_min"}:           &moonFunc{funcName: "func test_array_output_uint64_option_1_min() Array[UInt64?]"},
			{Name: "test_array_output_uint64_option_1_none"}:          &moonFunc{funcName: "func test_array_output_uint64_option_1_none() Array[UInt64?]"},
			{Name: "test_array_output_uint64_option_2"}:               &moonFunc{funcName: "func test_array_output_uint64_option_2() Array[UInt64?]"},
			{Name: "test_array_output_uint64_option_3"}:               &moonFunc{funcName: "func test_array_output_uint64_option_3() Array[UInt64?]"},
			{Name: "test_array_output_uint64_option_4"}:               &moonFunc{funcName: "func test_array_output_uint64_option_4() Array[UInt64?]"},
			{Name: "test_array_output_uint_0"}:                        &moonFunc{funcName: "func test_array_output_uint_0() Array[UInt]"},
			{Name: "test_array_output_uint_1"}:                        &moonFunc{funcName: "func test_array_output_uint_1() Array[UInt]"},
			{Name: "test_array_output_uint_1_max"}:                    &moonFunc{funcName: "func test_array_output_uint_1_max() Array[UInt]"},
			{Name: "test_array_output_uint_1_min"}:                    &moonFunc{funcName: "func test_array_output_uint_1_min() Array[UInt]"},
			{Name: "test_array_output_uint_2"}:                        &moonFunc{funcName: "func test_array_output_uint_2() Array[UInt]"},
			{Name: "test_array_output_uint_3"}:                        &moonFunc{funcName: "func test_array_output_uint_3() Array[UInt]"},
			{Name: "test_array_output_uint_4"}:                        &moonFunc{funcName: "func test_array_output_uint_4() Array[UInt]"},
			{Name: "test_array_output_uint_option"}:                   &moonFunc{funcName: "func test_array_output_uint_option() Array[UInt?]"},
			{Name: "test_array_output_uint_option_0"}:                 &moonFunc{funcName: "func test_array_output_uint_option_0() Array[UInt?]"},
			{Name: "test_array_output_uint_option_1_max"}:             &moonFunc{funcName: "func test_array_output_uint_option_1_max() Array[UInt?]"},
			{Name: "test_array_output_uint_option_1_min"}:             &moonFunc{funcName: "func test_array_output_uint_option_1_min() Array[UInt?]"},
			{Name: "test_array_output_uint_option_1_none"}:            &moonFunc{funcName: "func test_array_output_uint_option_1_none() Array[UInt?]"},
			{Name: "test_array_output_uint_option_2"}:                 &moonFunc{funcName: "func test_array_output_uint_option_2() Array[UInt?]"},
			{Name: "test_array_output_uint_option_3"}:                 &moonFunc{funcName: "func test_array_output_uint_option_3() Array[UInt?]"},
			{Name: "test_array_output_uint_option_4"}:                 &moonFunc{funcName: "func test_array_output_uint_option_4() Array[UInt?]"},
			{Name: "test_bool_input_false"}:                           &moonFunc{funcName: "func test_bool_input_false(b Bool) Unit!Error"},
			{Name: "test_bool_input_true"}:                            &moonFunc{funcName: "func test_bool_input_true(b Bool) Unit!Error"},
			{Name: "test_bool_option_input_false"}:                    &moonFunc{funcName: "func test_bool_option_input_false(b Bool?) Unit!Error"},
			{Name: "test_bool_option_input_none"}:                     &moonFunc{funcName: "func test_bool_option_input_none(b Bool?) Unit!Error"},
			{Name: "test_bool_option_input_true"}:                     &moonFunc{funcName: "func test_bool_option_input_true(b Bool?) Unit!Error"},
			{Name: "test_bool_option_output_false"}:                   &moonFunc{funcName: "func test_bool_option_output_false() Bool?"},
			{Name: "test_bool_option_output_none"}:                    &moonFunc{funcName: "func test_bool_option_output_none() Bool?"},
			{Name: "test_bool_option_output_true"}:                    &moonFunc{funcName: "func test_bool_option_output_true() Bool?"},
			{Name: "test_bool_output_false"}:                          &moonFunc{funcName: "func test_bool_output_false() Bool"},
			{Name: "test_bool_output_true"}:                           &moonFunc{funcName: "func test_bool_output_true() Bool"},
			{Name: "test_byte_input_max"}:                             &moonFunc{funcName: "func test_byte_input_max(b Byte) Unit!Error"},
			{Name: "test_byte_input_min"}:                             &moonFunc{funcName: "func test_byte_input_min(b Byte) Unit!Error"},
			{Name: "test_byte_option_input_max"}:                      &moonFunc{funcName: "func test_byte_option_input_max(b Byte?) Unit!Error"},
			{Name: "test_byte_option_input_min"}:                      &moonFunc{funcName: "func test_byte_option_input_min(b Byte?) Unit!Error"},
			{Name: "test_byte_option_input_none"}:                     &moonFunc{funcName: "func test_byte_option_input_none(b Byte?) Unit!Error"},
			{Name: "test_byte_option_output_max"}:                     &moonFunc{funcName: "func test_byte_option_output_max() Byte?"},
			{Name: "test_byte_option_output_min"}:                     &moonFunc{funcName: "func test_byte_option_output_min() Byte?"},
			{Name: "test_byte_option_output_none"}:                    &moonFunc{funcName: "func test_byte_option_output_none() Byte?"},
			{Name: "test_byte_output_max"}:                            &moonFunc{funcName: "func test_byte_output_max() Byte"},
			{Name: "test_byte_output_min"}:                            &moonFunc{funcName: "func test_byte_output_min() Byte"},
			{Name: "test_char_input_max"}:                             &moonFunc{funcName: "func test_char_input_max(c Char) Unit!Error"},
			{Name: "test_char_input_min"}:                             &moonFunc{funcName: "func test_char_input_min(c Char) Unit!Error"},
			{Name: "test_char_option_input_max"}:                      &moonFunc{funcName: "func test_char_option_input_max(c Char?) Unit!Error"},
			{Name: "test_char_option_input_min"}:                      &moonFunc{funcName: "func test_char_option_input_min(c Char?) Unit!Error"},
			{Name: "test_char_option_input_none"}:                     &moonFunc{funcName: "func test_char_option_input_none(c Char?) Unit!Error"},
			{Name: "test_char_option_output_max"}:                     &moonFunc{funcName: "func test_char_option_output_max() Char?"},
			{Name: "test_char_option_output_min"}:                     &moonFunc{funcName: "func test_char_option_output_min() Char?"},
			{Name: "test_char_option_output_none"}:                    &moonFunc{funcName: "func test_char_option_output_none() Char?"},
			{Name: "test_char_output_max"}:                            &moonFunc{funcName: "func test_char_output_max() Char"},
			{Name: "test_char_output_min"}:                            &moonFunc{funcName: "func test_char_output_min() Char"},
			{Name: "test_double_input_max"}:                           &moonFunc{funcName: "func test_double_input_max(n Double) Unit!Error"},
			{Name: "test_double_input_min"}:                           &moonFunc{funcName: "func test_double_input_min(n Double) Unit!Error"},
			{Name: "test_double_option_input_max"}:                    &moonFunc{funcName: "func test_double_option_input_max(n Double?) Unit!Error"},
			{Name: "test_double_option_input_min"}:                    &moonFunc{funcName: "func test_double_option_input_min(n Double?) Unit!Error"},
			{Name: "test_double_option_input_none"}:                   &moonFunc{funcName: "func test_double_option_input_none(n Double?) Unit!Error"},
			{Name: "test_double_option_output_max"}:                   &moonFunc{funcName: "func test_double_option_output_max() Double?"},
			{Name: "test_double_option_output_min"}:                   &moonFunc{funcName: "func test_double_option_output_min() Double?"},
			{Name: "test_double_option_output_none"}:                  &moonFunc{funcName: "func test_double_option_output_none() Double?"},
			{Name: "test_double_output_max"}:                          &moonFunc{funcName: "func test_double_output_max() Double"},
			{Name: "test_double_output_min"}:                          &moonFunc{funcName: "func test_double_output_min() Double"},
			{Name: "test_duration_input"}:                             &moonFunc{funcName: "func test_duration_input(d @time.Duration) Unit!Error"},
			{Name: "test_duration_option_input"}:                      &moonFunc{funcName: "func test_duration_option_input(d @time.Duration?) Unit!Error"},
			{Name: "test_duration_option_input_none"}:                 &moonFunc{funcName: "func test_duration_option_input_none(d @time.Duration?) Unit!Error"},
			{Name: "test_duration_option_input_none_style2"}:          &moonFunc{funcName: "func test_duration_option_input_none_style2(d? @time.Duration) Unit!Error"},
			{Name: "test_duration_option_input_style2"}:               &moonFunc{funcName: "func test_duration_option_input_style2(d? @time.Duration) Unit!Error"},
			{Name: "test_duration_option_output"}:                     &moonFunc{funcName: "func test_duration_option_output() @time.Duration?"},
			{Name: "test_duration_option_output_none"}:                &moonFunc{funcName: "func test_duration_option_output_none() @time.Duration?"},
			{Name: "test_duration_output"}:                            &moonFunc{funcName: "func test_duration_output() @time.Duration"},
			{Name: "test_fixedarray_input0_byte"}:                     &moonFunc{funcName: "func test_fixedarray_input0_byte(val FixedArray[Byte]) Unit!Error"},
			{Name: "test_fixedarray_input0_int_option"}:               &moonFunc{funcName: "func test_fixedarray_input0_int_option(val FixedArray[Int?]) Unit!Error"},
			{Name: "test_fixedarray_input0_string"}:                   &moonFunc{funcName: "func test_fixedarray_input0_string(val FixedArray[String]) Unit!Error"},
			{Name: "test_fixedarray_input0_string_option"}:            &moonFunc{funcName: "func test_fixedarray_input0_string_option(val FixedArray[String?]) Unit!Error"},
			{Name: "test_fixedarray_input1_byte"}:                     &moonFunc{funcName: "func test_fixedarray_input1_byte(val FixedArray[Byte]) Unit!Error"},
			{Name: "test_fixedarray_input1_int_option"}:               &moonFunc{funcName: "func test_fixedarray_input1_int_option(val FixedArray[Int?]) Unit!Error"},
			{Name: "test_fixedarray_input1_string"}:                   &moonFunc{funcName: "func test_fixedarray_input1_string(val FixedArray[String]) Unit!Error"},
			{Name: "test_fixedarray_input1_string_option"}:            &moonFunc{funcName: "func test_fixedarray_input1_string_option(val FixedArray[String?]) Unit!Error"},
			{Name: "test_fixedarray_input2_byte"}:                     &moonFunc{funcName: "func test_fixedarray_input2_byte(val FixedArray[Byte]) Unit!Error"},
			{Name: "test_fixedarray_input2_int_option"}:               &moonFunc{funcName: "func test_fixedarray_input2_int_option(val FixedArray[Int?]) Unit!Error"},
			{Name: "test_fixedarray_input2_map"}:                      &moonFunc{funcName: "func test_fixedarray_input2_map(val FixedArray[Map[String, String]]) Unit!Error"},
			{Name: "test_fixedarray_input2_map_option"}:               &moonFunc{funcName: "func test_fixedarray_input2_map_option(val FixedArray[Map[String, String]?]) Unit!Error"},
			{Name: "test_fixedarray_input2_string"}:                   &moonFunc{funcName: "func test_fixedarray_input2_string(val FixedArray[String]) Unit!Error"},
			{Name: "test_fixedarray_input2_string_option"}:            &moonFunc{funcName: "func test_fixedarray_input2_string_option(val FixedArray[String?]) Unit!Error"},
			{Name: "test_fixedarray_input2_struct"}:                   &moonFunc{funcName: "func test_fixedarray_input2_struct(val FixedArray[TestStruct2]) Unit!Error"},
			{Name: "test_fixedarray_input2_struct_option"}:            &moonFunc{funcName: "func test_fixedarray_input2_struct_option(val FixedArray[TestStruct2?]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_0"}:                    &moonFunc{funcName: "func test_fixedarray_input_bool_0(val FixedArray[Bool]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_1"}:                    &moonFunc{funcName: "func test_fixedarray_input_bool_1(val FixedArray[Bool]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_2"}:                    &moonFunc{funcName: "func test_fixedarray_input_bool_2(val FixedArray[Bool]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_3"}:                    &moonFunc{funcName: "func test_fixedarray_input_bool_3(val FixedArray[Bool]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_4"}:                    &moonFunc{funcName: "func test_fixedarray_input_bool_4(val FixedArray[Bool]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_option_0"}:             &moonFunc{funcName: "func test_fixedarray_input_bool_option_0(val FixedArray[Bool?]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_option_1_false"}:       &moonFunc{funcName: "func test_fixedarray_input_bool_option_1_false(val FixedArray[Bool?]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_option_1_none"}:        &moonFunc{funcName: "func test_fixedarray_input_bool_option_1_none(val FixedArray[Bool?]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_option_1_true"}:        &moonFunc{funcName: "func test_fixedarray_input_bool_option_1_true(val FixedArray[Bool?]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_option_2"}:             &moonFunc{funcName: "func test_fixedarray_input_bool_option_2(val FixedArray[Bool?]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_option_3"}:             &moonFunc{funcName: "func test_fixedarray_input_bool_option_3(val FixedArray[Bool?]) Unit!Error"},
			{Name: "test_fixedarray_input_bool_option_4"}:             &moonFunc{funcName: "func test_fixedarray_input_bool_option_4(val FixedArray[Bool?]) Unit!Error"},
			{Name: "test_fixedarray_input_byte_0"}:                    &moonFunc{funcName: "func test_fixedarray_input_byte_0(val FixedArray[Byte]) Unit!Error"},
			{Name: "test_fixedarray_input_byte_1"}:                    &moonFunc{funcName: "func test_fixedarray_input_byte_1(val FixedArray[Byte]) Unit!Error"},
			{Name: "test_fixedarray_input_byte_2"}:                    &moonFunc{funcName: "func test_fixedarray_input_byte_2(val FixedArray[Byte]) Unit!Error"},
			{Name: "test_fixedarray_input_byte_3"}:                    &moonFunc{funcName: "func test_fixedarray_input_byte_3(val FixedArray[Byte]) Unit!Error"},
			{Name: "test_fixedarray_input_byte_4"}:                    &moonFunc{funcName: "func test_fixedarray_input_byte_4(val FixedArray[Byte]) Unit!Error"},
			{Name: "test_fixedarray_input_byte_option_0"}:             &moonFunc{funcName: "func test_fixedarray_input_byte_option_0(val FixedArray[Byte?]) Unit!Error"},
			{Name: "test_fixedarray_input_byte_option_1"}:             &moonFunc{funcName: "func test_fixedarray_input_byte_option_1(val FixedArray[Byte?]) Unit!Error"},
			{Name: "test_fixedarray_input_byte_option_2"}:             &moonFunc{funcName: "func test_fixedarray_input_byte_option_2(val FixedArray[Byte?]) Unit!Error"},
			{Name: "test_fixedarray_input_byte_option_3"}:             &moonFunc{funcName: "func test_fixedarray_input_byte_option_3(val FixedArray[Byte?]) Unit!Error"},
			{Name: "test_fixedarray_input_byte_option_4"}:             &moonFunc{funcName: "func test_fixedarray_input_byte_option_4(val FixedArray[Byte?]) Unit!Error"},
			{Name: "test_fixedarray_input_char_0"}:                    &moonFunc{funcName: "func test_fixedarray_input_char_0(val FixedArray[Char]) Unit!Error"},
			{Name: "test_fixedarray_input_char_1"}:                    &moonFunc{funcName: "func test_fixedarray_input_char_1(val FixedArray[Char]) Unit!Error"},
			{Name: "test_fixedarray_input_char_2"}:                    &moonFunc{funcName: "func test_fixedarray_input_char_2(val FixedArray[Char]) Unit!Error"},
			{Name: "test_fixedarray_input_char_3"}:                    &moonFunc{funcName: "func test_fixedarray_input_char_3(val FixedArray[Char]) Unit!Error"},
			{Name: "test_fixedarray_input_char_4"}:                    &moonFunc{funcName: "func test_fixedarray_input_char_4(val FixedArray[Char]) Unit!Error"},
			{Name: "test_fixedarray_input_char_option_0"}:             &moonFunc{funcName: "func test_fixedarray_input_char_option_0(val FixedArray[Char?]) Unit!Error"},
			{Name: "test_fixedarray_input_char_option_1_none"}:        &moonFunc{funcName: "func test_fixedarray_input_char_option_1_none(val FixedArray[Char?]) Unit!Error"},
			{Name: "test_fixedarray_input_char_option_1_some"}:        &moonFunc{funcName: "func test_fixedarray_input_char_option_1_some(val FixedArray[Char?]) Unit!Error"},
			{Name: "test_fixedarray_input_char_option_2"}:             &moonFunc{funcName: "func test_fixedarray_input_char_option_2(val FixedArray[Char?]) Unit!Error"},
			{Name: "test_fixedarray_input_char_option_3"}:             &moonFunc{funcName: "func test_fixedarray_input_char_option_3(val FixedArray[Char?]) Unit!Error"},
			{Name: "test_fixedarray_input_char_option_4"}:             &moonFunc{funcName: "func test_fixedarray_input_char_option_4(val FixedArray[Char?]) Unit!Error"},
			{Name: "test_fixedarray_input_double_0"}:                  &moonFunc{funcName: "func test_fixedarray_input_double_0(val FixedArray[Double]) Unit!Error"},
			{Name: "test_fixedarray_input_double_1"}:                  &moonFunc{funcName: "func test_fixedarray_input_double_1(val FixedArray[Double]) Unit!Error"},
			{Name: "test_fixedarray_input_double_2"}:                  &moonFunc{funcName: "func test_fixedarray_input_double_2(val FixedArray[Double]) Unit!Error"},
			{Name: "test_fixedarray_input_double_3"}:                  &moonFunc{funcName: "func test_fixedarray_input_double_3(val FixedArray[Double]) Unit!Error"},
			{Name: "test_fixedarray_input_double_4"}:                  &moonFunc{funcName: "func test_fixedarray_input_double_4(val FixedArray[Double]) Unit!Error"},
			{Name: "test_fixedarray_input_double_option_0"}:           &moonFunc{funcName: "func test_fixedarray_input_double_option_0(val FixedArray[Double?]) Unit!Error"},
			{Name: "test_fixedarray_input_double_option_1_none"}:      &moonFunc{funcName: "func test_fixedarray_input_double_option_1_none(val FixedArray[Double?]) Unit!Error"},
			{Name: "test_fixedarray_input_double_option_1_some"}:      &moonFunc{funcName: "func test_fixedarray_input_double_option_1_some(val FixedArray[Double?]) Unit!Error"},
			{Name: "test_fixedarray_input_double_option_2"}:           &moonFunc{funcName: "func test_fixedarray_input_double_option_2(val FixedArray[Double?]) Unit!Error"},
			{Name: "test_fixedarray_input_double_option_3"}:           &moonFunc{funcName: "func test_fixedarray_input_double_option_3(val FixedArray[Double?]) Unit!Error"},
			{Name: "test_fixedarray_input_double_option_4"}:           &moonFunc{funcName: "func test_fixedarray_input_double_option_4(val FixedArray[Double?]) Unit!Error"},
			{Name: "test_fixedarray_input_float_0"}:                   &moonFunc{funcName: "func test_fixedarray_input_float_0(val FixedArray[Float]) Unit!Error"},
			{Name: "test_fixedarray_input_float_1"}:                   &moonFunc{funcName: "func test_fixedarray_input_float_1(val FixedArray[Float]) Unit!Error"},
			{Name: "test_fixedarray_input_float_2"}:                   &moonFunc{funcName: "func test_fixedarray_input_float_2(val FixedArray[Float]) Unit!Error"},
			{Name: "test_fixedarray_input_float_3"}:                   &moonFunc{funcName: "func test_fixedarray_input_float_3(val FixedArray[Float]) Unit!Error"},
			{Name: "test_fixedarray_input_float_4"}:                   &moonFunc{funcName: "func test_fixedarray_input_float_4(val FixedArray[Float]) Unit!Error"},
			{Name: "test_fixedarray_input_float_option_0"}:            &moonFunc{funcName: "func test_fixedarray_input_float_option_0(val FixedArray[Float?]) Unit!Error"},
			{Name: "test_fixedarray_input_float_option_1_none"}:       &moonFunc{funcName: "func test_fixedarray_input_float_option_1_none(val FixedArray[Float?]) Unit!Error"},
			{Name: "test_fixedarray_input_float_option_1_some"}:       &moonFunc{funcName: "func test_fixedarray_input_float_option_1_some(val FixedArray[Float?]) Unit!Error"},
			{Name: "test_fixedarray_input_float_option_2"}:            &moonFunc{funcName: "func test_fixedarray_input_float_option_2(val FixedArray[Float?]) Unit!Error"},
			{Name: "test_fixedarray_input_float_option_3"}:            &moonFunc{funcName: "func test_fixedarray_input_float_option_3(val FixedArray[Float?]) Unit!Error"},
			{Name: "test_fixedarray_input_float_option_4"}:            &moonFunc{funcName: "func test_fixedarray_input_float_option_4(val FixedArray[Float?]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_0"}:                   &moonFunc{funcName: "func test_fixedarray_input_int16_0(val FixedArray[Int16]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_1"}:                   &moonFunc{funcName: "func test_fixedarray_input_int16_1(val FixedArray[Int16]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_1_max"}:               &moonFunc{funcName: "func test_fixedarray_input_int16_1_max(val FixedArray[Int16]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_1_min"}:               &moonFunc{funcName: "func test_fixedarray_input_int16_1_min(val FixedArray[Int16]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_2"}:                   &moonFunc{funcName: "func test_fixedarray_input_int16_2(val FixedArray[Int16]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_3"}:                   &moonFunc{funcName: "func test_fixedarray_input_int16_3(val FixedArray[Int16]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_4"}:                   &moonFunc{funcName: "func test_fixedarray_input_int16_4(val FixedArray[Int16]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_option_0"}:            &moonFunc{funcName: "func test_fixedarray_input_int16_option_0(val FixedArray[Int16?]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_option_1_max"}:        &moonFunc{funcName: "func test_fixedarray_input_int16_option_1_max(val FixedArray[Int16?]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_option_1_min"}:        &moonFunc{funcName: "func test_fixedarray_input_int16_option_1_min(val FixedArray[Int16?]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_option_1_none"}:       &moonFunc{funcName: "func test_fixedarray_input_int16_option_1_none(val FixedArray[Int16?]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_option_2"}:            &moonFunc{funcName: "func test_fixedarray_input_int16_option_2(val FixedArray[Int16?]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_option_3"}:            &moonFunc{funcName: "func test_fixedarray_input_int16_option_3(val FixedArray[Int16?]) Unit!Error"},
			{Name: "test_fixedarray_input_int16_option_4"}:            &moonFunc{funcName: "func test_fixedarray_input_int16_option_4(val FixedArray[Int16?]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_0"}:                   &moonFunc{funcName: "func test_fixedarray_input_int64_0(val FixedArray[Int64]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_1"}:                   &moonFunc{funcName: "func test_fixedarray_input_int64_1(val FixedArray[Int64]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_1_max"}:               &moonFunc{funcName: "func test_fixedarray_input_int64_1_max(val FixedArray[Int64]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_1_min"}:               &moonFunc{funcName: "func test_fixedarray_input_int64_1_min(val FixedArray[Int64]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_2"}:                   &moonFunc{funcName: "func test_fixedarray_input_int64_2(val FixedArray[Int64]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_3"}:                   &moonFunc{funcName: "func test_fixedarray_input_int64_3(val FixedArray[Int64]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_4"}:                   &moonFunc{funcName: "func test_fixedarray_input_int64_4(val FixedArray[Int64]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_option_0"}:            &moonFunc{funcName: "func test_fixedarray_input_int64_option_0(val FixedArray[Int64?]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_option_1_max"}:        &moonFunc{funcName: "func test_fixedarray_input_int64_option_1_max(val FixedArray[Int64?]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_option_1_min"}:        &moonFunc{funcName: "func test_fixedarray_input_int64_option_1_min(val FixedArray[Int64?]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_option_1_none"}:       &moonFunc{funcName: "func test_fixedarray_input_int64_option_1_none(val FixedArray[Int64?]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_option_2"}:            &moonFunc{funcName: "func test_fixedarray_input_int64_option_2(val FixedArray[Int64?]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_option_3"}:            &moonFunc{funcName: "func test_fixedarray_input_int64_option_3(val FixedArray[Int64?]) Unit!Error"},
			{Name: "test_fixedarray_input_int64_option_4"}:            &moonFunc{funcName: "func test_fixedarray_input_int64_option_4(val FixedArray[Int64?]) Unit!Error"},
			{Name: "test_fixedarray_input_int_0"}:                     &moonFunc{funcName: "func test_fixedarray_input_int_0(val FixedArray[Int]) Unit!Error"},
			{Name: "test_fixedarray_input_int_1"}:                     &moonFunc{funcName: "func test_fixedarray_input_int_1(val FixedArray[Int]) Unit!Error"},
			{Name: "test_fixedarray_input_int_1_max"}:                 &moonFunc{funcName: "func test_fixedarray_input_int_1_max(val FixedArray[Int]) Unit!Error"},
			{Name: "test_fixedarray_input_int_1_min"}:                 &moonFunc{funcName: "func test_fixedarray_input_int_1_min(val FixedArray[Int]) Unit!Error"},
			{Name: "test_fixedarray_input_int_2"}:                     &moonFunc{funcName: "func test_fixedarray_input_int_2(val FixedArray[Int]) Unit!Error"},
			{Name: "test_fixedarray_input_int_3"}:                     &moonFunc{funcName: "func test_fixedarray_input_int_3(val FixedArray[Int]) Unit!Error"},
			{Name: "test_fixedarray_input_int_4"}:                     &moonFunc{funcName: "func test_fixedarray_input_int_4(val FixedArray[Int]) Unit!Error"},
			{Name: "test_fixedarray_input_int_option_0"}:              &moonFunc{funcName: "func test_fixedarray_input_int_option_0(val FixedArray[Int?]) Unit!Error"},
			{Name: "test_fixedarray_input_int_option_1_max"}:          &moonFunc{funcName: "func test_fixedarray_input_int_option_1_max(val FixedArray[Int?]) Unit!Error"},
			{Name: "test_fixedarray_input_int_option_1_min"}:          &moonFunc{funcName: "func test_fixedarray_input_int_option_1_min(val FixedArray[Int?]) Unit!Error"},
			{Name: "test_fixedarray_input_int_option_1_none"}:         &moonFunc{funcName: "func test_fixedarray_input_int_option_1_none(val FixedArray[Int?]) Unit!Error"},
			{Name: "test_fixedarray_input_int_option_2"}:              &moonFunc{funcName: "func test_fixedarray_input_int_option_2(val FixedArray[Int?]) Unit!Error"},
			{Name: "test_fixedarray_input_int_option_3"}:              &moonFunc{funcName: "func test_fixedarray_input_int_option_3(val FixedArray[Int?]) Unit!Error"},
			{Name: "test_fixedarray_input_int_option_4"}:              &moonFunc{funcName: "func test_fixedarray_input_int_option_4(val FixedArray[Int?]) Unit!Error"},
			{Name: "test_fixedarray_input_string_0"}:                  &moonFunc{funcName: "func test_fixedarray_input_string_0(val FixedArray[String]) Unit!Error"},
			{Name: "test_fixedarray_input_string_1"}:                  &moonFunc{funcName: "func test_fixedarray_input_string_1(val FixedArray[String]) Unit!Error"},
			{Name: "test_fixedarray_input_string_2"}:                  &moonFunc{funcName: "func test_fixedarray_input_string_2(val FixedArray[String]) Unit!Error"},
			{Name: "test_fixedarray_input_string_3"}:                  &moonFunc{funcName: "func test_fixedarray_input_string_3(val FixedArray[String]) Unit!Error"},
			{Name: "test_fixedarray_input_string_4"}:                  &moonFunc{funcName: "func test_fixedarray_input_string_4(val FixedArray[String]) Unit!Error"},
			{Name: "test_fixedarray_input_string_option_0"}:           &moonFunc{funcName: "func test_fixedarray_input_string_option_0(val FixedArray[String?]) Unit!Error"},
			{Name: "test_fixedarray_input_string_option_1_none"}:      &moonFunc{funcName: "func test_fixedarray_input_string_option_1_none(val FixedArray[String?]) Unit!Error"},
			{Name: "test_fixedarray_input_string_option_1_some"}:      &moonFunc{funcName: "func test_fixedarray_input_string_option_1_some(val FixedArray[String?]) Unit!Error"},
			{Name: "test_fixedarray_input_string_option_2"}:           &moonFunc{funcName: "func test_fixedarray_input_string_option_2(val FixedArray[String?]) Unit!Error"},
			{Name: "test_fixedarray_input_string_option_3"}:           &moonFunc{funcName: "func test_fixedarray_input_string_option_3(val FixedArray[String?]) Unit!Error"},
			{Name: "test_fixedarray_input_string_option_4"}:           &moonFunc{funcName: "func test_fixedarray_input_string_option_4(val FixedArray[String?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_0"}:                  &moonFunc{funcName: "func test_fixedarray_input_uint16_0(val FixedArray[UInt16]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_1"}:                  &moonFunc{funcName: "func test_fixedarray_input_uint16_1(val FixedArray[UInt16]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_1_max"}:              &moonFunc{funcName: "func test_fixedarray_input_uint16_1_max(val FixedArray[UInt16]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_1_min"}:              &moonFunc{funcName: "func test_fixedarray_input_uint16_1_min(val FixedArray[UInt16]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_2"}:                  &moonFunc{funcName: "func test_fixedarray_input_uint16_2(val FixedArray[UInt16]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_3"}:                  &moonFunc{funcName: "func test_fixedarray_input_uint16_3(val FixedArray[UInt16]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_4"}:                  &moonFunc{funcName: "func test_fixedarray_input_uint16_4(val FixedArray[UInt16]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_option_0"}:           &moonFunc{funcName: "func test_fixedarray_input_uint16_option_0(val FixedArray[UInt16?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_option_1_max"}:       &moonFunc{funcName: "func test_fixedarray_input_uint16_option_1_max(val FixedArray[UInt16?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_option_1_min"}:       &moonFunc{funcName: "func test_fixedarray_input_uint16_option_1_min(val FixedArray[UInt16?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_option_1_none"}:      &moonFunc{funcName: "func test_fixedarray_input_uint16_option_1_none(val FixedArray[UInt16?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_option_2"}:           &moonFunc{funcName: "func test_fixedarray_input_uint16_option_2(val FixedArray[UInt16?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_option_3"}:           &moonFunc{funcName: "func test_fixedarray_input_uint16_option_3(val FixedArray[UInt16?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint16_option_4"}:           &moonFunc{funcName: "func test_fixedarray_input_uint16_option_4(val FixedArray[UInt16?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_0"}:                  &moonFunc{funcName: "func test_fixedarray_input_uint64_0(val FixedArray[UInt64]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_1"}:                  &moonFunc{funcName: "func test_fixedarray_input_uint64_1(val FixedArray[UInt64]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_1_max"}:              &moonFunc{funcName: "func test_fixedarray_input_uint64_1_max(val FixedArray[UInt64]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_1_min"}:              &moonFunc{funcName: "func test_fixedarray_input_uint64_1_min(val FixedArray[UInt64]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_2"}:                  &moonFunc{funcName: "func test_fixedarray_input_uint64_2(val FixedArray[UInt64]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_3"}:                  &moonFunc{funcName: "func test_fixedarray_input_uint64_3(val FixedArray[UInt64]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_4"}:                  &moonFunc{funcName: "func test_fixedarray_input_uint64_4(val FixedArray[UInt64]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_option_0"}:           &moonFunc{funcName: "func test_fixedarray_input_uint64_option_0(val FixedArray[UInt64?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_option_1_max"}:       &moonFunc{funcName: "func test_fixedarray_input_uint64_option_1_max(val FixedArray[UInt64?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_option_1_min"}:       &moonFunc{funcName: "func test_fixedarray_input_uint64_option_1_min(val FixedArray[UInt64?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_option_1_none"}:      &moonFunc{funcName: "func test_fixedarray_input_uint64_option_1_none(val FixedArray[UInt64?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_option_2"}:           &moonFunc{funcName: "func test_fixedarray_input_uint64_option_2(val FixedArray[UInt64?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_option_3"}:           &moonFunc{funcName: "func test_fixedarray_input_uint64_option_3(val FixedArray[UInt64?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint64_option_4"}:           &moonFunc{funcName: "func test_fixedarray_input_uint64_option_4(val FixedArray[UInt64?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_0"}:                    &moonFunc{funcName: "func test_fixedarray_input_uint_0(val FixedArray[UInt]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_1"}:                    &moonFunc{funcName: "func test_fixedarray_input_uint_1(val FixedArray[UInt]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_1_max"}:                &moonFunc{funcName: "func test_fixedarray_input_uint_1_max(val FixedArray[UInt]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_1_min"}:                &moonFunc{funcName: "func test_fixedarray_input_uint_1_min(val FixedArray[UInt]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_2"}:                    &moonFunc{funcName: "func test_fixedarray_input_uint_2(val FixedArray[UInt]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_3"}:                    &moonFunc{funcName: "func test_fixedarray_input_uint_3(val FixedArray[UInt]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_4"}:                    &moonFunc{funcName: "func test_fixedarray_input_uint_4(val FixedArray[UInt]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_option_0"}:             &moonFunc{funcName: "func test_fixedarray_input_uint_option_0(val FixedArray[UInt?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_option_1_max"}:         &moonFunc{funcName: "func test_fixedarray_input_uint_option_1_max(val FixedArray[UInt?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_option_1_min"}:         &moonFunc{funcName: "func test_fixedarray_input_uint_option_1_min(val FixedArray[UInt?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_option_1_none"}:        &moonFunc{funcName: "func test_fixedarray_input_uint_option_1_none(val FixedArray[UInt?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_option_2"}:             &moonFunc{funcName: "func test_fixedarray_input_uint_option_2(val FixedArray[UInt?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_option_3"}:             &moonFunc{funcName: "func test_fixedarray_input_uint_option_3(val FixedArray[UInt?]) Unit!Error"},
			{Name: "test_fixedarray_input_uint_option_4"}:             &moonFunc{funcName: "func test_fixedarray_input_uint_option_4(val FixedArray[UInt?]) Unit!Error"},
			{Name: "test_fixedarray_output0_byte"}:                    &moonFunc{funcName: "func test_fixedarray_output0_byte() FixedArray[Byte]"},
			{Name: "test_fixedarray_output0_int_option"}:              &moonFunc{funcName: "func test_fixedarray_output0_int_option() FixedArray[Int?]"},
			{Name: "test_fixedarray_output0_string"}:                  &moonFunc{funcName: "func test_fixedarray_output0_string() FixedArray[String]"},
			{Name: "test_fixedarray_output0_string_option"}:           &moonFunc{funcName: "func test_fixedarray_output0_string_option() FixedArray[String?]"},
			{Name: "test_fixedarray_output1_byte"}:                    &moonFunc{funcName: "func test_fixedarray_output1_byte() FixedArray[Byte]"},
			{Name: "test_fixedarray_output1_int_option"}:              &moonFunc{funcName: "func test_fixedarray_output1_int_option() FixedArray[Int?]"},
			{Name: "test_fixedarray_output1_string"}:                  &moonFunc{funcName: "func test_fixedarray_output1_string() FixedArray[String]"},
			{Name: "test_fixedarray_output1_string_option"}:           &moonFunc{funcName: "func test_fixedarray_output1_string_option() FixedArray[String?]"},
			{Name: "test_fixedarray_output2_byte"}:                    &moonFunc{funcName: "func test_fixedarray_output2_byte() FixedArray[Byte]"},
			{Name: "test_fixedarray_output2_int_option"}:              &moonFunc{funcName: "func test_fixedarray_output2_int_option() FixedArray[Int?]"},
			{Name: "test_fixedarray_output2_map"}:                     &moonFunc{funcName: "func test_fixedarray_output2_map() FixedArray[Map[String, String]]"},
			{Name: "test_fixedarray_output2_map_option"}:              &moonFunc{funcName: "func test_fixedarray_output2_map_option() FixedArray[Map[String, String]?]"},
			{Name: "test_fixedarray_output2_string"}:                  &moonFunc{funcName: "func test_fixedarray_output2_string() FixedArray[String]"},
			{Name: "test_fixedarray_output2_string_option"}:           &moonFunc{funcName: "func test_fixedarray_output2_string_option() FixedArray[String?]"},
			{Name: "test_fixedarray_output2_struct"}:                  &moonFunc{funcName: "func test_fixedarray_output2_struct() FixedArray[TestStruct2]"},
			{Name: "test_fixedarray_output2_struct_option"}:           &moonFunc{funcName: "func test_fixedarray_output2_struct_option() FixedArray[TestStruct2?]"},
			{Name: "test_fixedarray_output_bool_0"}:                   &moonFunc{funcName: "func test_fixedarray_output_bool_0() FixedArray[Bool]"},
			{Name: "test_fixedarray_output_bool_1"}:                   &moonFunc{funcName: "func test_fixedarray_output_bool_1() FixedArray[Bool]"},
			{Name: "test_fixedarray_output_bool_2"}:                   &moonFunc{funcName: "func test_fixedarray_output_bool_2() FixedArray[Bool]"},
			{Name: "test_fixedarray_output_bool_3"}:                   &moonFunc{funcName: "func test_fixedarray_output_bool_3() FixedArray[Bool]"},
			{Name: "test_fixedarray_output_bool_4"}:                   &moonFunc{funcName: "func test_fixedarray_output_bool_4() FixedArray[Bool]"},
			{Name: "test_fixedarray_output_bool_option_0"}:            &moonFunc{funcName: "func test_fixedarray_output_bool_option_0() FixedArray[Bool?]"},
			{Name: "test_fixedarray_output_bool_option_1_false"}:      &moonFunc{funcName: "func test_fixedarray_output_bool_option_1_false() FixedArray[Bool?]"},
			{Name: "test_fixedarray_output_bool_option_1_none"}:       &moonFunc{funcName: "func test_fixedarray_output_bool_option_1_none() FixedArray[Bool?]"},
			{Name: "test_fixedarray_output_bool_option_1_true"}:       &moonFunc{funcName: "func test_fixedarray_output_bool_option_1_true() FixedArray[Bool?]"},
			{Name: "test_fixedarray_output_bool_option_2"}:            &moonFunc{funcName: "func test_fixedarray_output_bool_option_2() FixedArray[Bool?]"},
			{Name: "test_fixedarray_output_bool_option_3"}:            &moonFunc{funcName: "func test_fixedarray_output_bool_option_3() FixedArray[Bool?]"},
			{Name: "test_fixedarray_output_bool_option_4"}:            &moonFunc{funcName: "func test_fixedarray_output_bool_option_4() FixedArray[Bool?]"},
			{Name: "test_fixedarray_output_byte_0"}:                   &moonFunc{funcName: "func test_fixedarray_output_byte_0() FixedArray[Byte]"},
			{Name: "test_fixedarray_output_byte_1"}:                   &moonFunc{funcName: "func test_fixedarray_output_byte_1() FixedArray[Byte]"},
			{Name: "test_fixedarray_output_byte_2"}:                   &moonFunc{funcName: "func test_fixedarray_output_byte_2() FixedArray[Byte]"},
			{Name: "test_fixedarray_output_byte_3"}:                   &moonFunc{funcName: "func test_fixedarray_output_byte_3() FixedArray[Byte]"},
			{Name: "test_fixedarray_output_byte_4"}:                   &moonFunc{funcName: "func test_fixedarray_output_byte_4() FixedArray[Byte]"},
			{Name: "test_fixedarray_output_byte_option_0"}:            &moonFunc{funcName: "func test_fixedarray_output_byte_option_0() FixedArray[Byte?]"},
			{Name: "test_fixedarray_output_byte_option_1"}:            &moonFunc{funcName: "func test_fixedarray_output_byte_option_1() FixedArray[Byte?]"},
			{Name: "test_fixedarray_output_byte_option_2"}:            &moonFunc{funcName: "func test_fixedarray_output_byte_option_2() FixedArray[Byte?]"},
			{Name: "test_fixedarray_output_byte_option_3"}:            &moonFunc{funcName: "func test_fixedarray_output_byte_option_3() FixedArray[Byte?]"},
			{Name: "test_fixedarray_output_byte_option_4"}:            &moonFunc{funcName: "func test_fixedarray_output_byte_option_4() FixedArray[Byte?]"},
			{Name: "test_fixedarray_output_char_0"}:                   &moonFunc{funcName: "func test_fixedarray_output_char_0() FixedArray[Char]"},
			{Name: "test_fixedarray_output_char_1"}:                   &moonFunc{funcName: "func test_fixedarray_output_char_1() FixedArray[Char]"},
			{Name: "test_fixedarray_output_char_2"}:                   &moonFunc{funcName: "func test_fixedarray_output_char_2() FixedArray[Char]"},
			{Name: "test_fixedarray_output_char_3"}:                   &moonFunc{funcName: "func test_fixedarray_output_char_3() FixedArray[Char]"},
			{Name: "test_fixedarray_output_char_4"}:                   &moonFunc{funcName: "func test_fixedarray_output_char_4() FixedArray[Char]"},
			{Name: "test_fixedarray_output_char_option_0"}:            &moonFunc{funcName: "func test_fixedarray_output_char_option_0() FixedArray[Char?]"},
			{Name: "test_fixedarray_output_char_option_1_none"}:       &moonFunc{funcName: "func test_fixedarray_output_char_option_1_none() FixedArray[Char?]"},
			{Name: "test_fixedarray_output_char_option_1_some"}:       &moonFunc{funcName: "func test_fixedarray_output_char_option_1_some() FixedArray[Char?]"},
			{Name: "test_fixedarray_output_char_option_2"}:            &moonFunc{funcName: "func test_fixedarray_output_char_option_2() FixedArray[Char?]"},
			{Name: "test_fixedarray_output_char_option_3"}:            &moonFunc{funcName: "func test_fixedarray_output_char_option_3() FixedArray[Char?]"},
			{Name: "test_fixedarray_output_char_option_4"}:            &moonFunc{funcName: "func test_fixedarray_output_char_option_4() FixedArray[Char?]"},
			{Name: "test_fixedarray_output_double_0"}:                 &moonFunc{funcName: "func test_fixedarray_output_double_0() FixedArray[Double]"},
			{Name: "test_fixedarray_output_double_1"}:                 &moonFunc{funcName: "func test_fixedarray_output_double_1() FixedArray[Double]"},
			{Name: "test_fixedarray_output_double_2"}:                 &moonFunc{funcName: "func test_fixedarray_output_double_2() FixedArray[Double]"},
			{Name: "test_fixedarray_output_double_3"}:                 &moonFunc{funcName: "func test_fixedarray_output_double_3() FixedArray[Double]"},
			{Name: "test_fixedarray_output_double_4"}:                 &moonFunc{funcName: "func test_fixedarray_output_double_4() FixedArray[Double]"},
			{Name: "test_fixedarray_output_double_option_0"}:          &moonFunc{funcName: "func test_fixedarray_output_double_option_0() FixedArray[Double?]"},
			{Name: "test_fixedarray_output_double_option_1_none"}:     &moonFunc{funcName: "func test_fixedarray_output_double_option_1_none() FixedArray[Double?]"},
			{Name: "test_fixedarray_output_double_option_1_some"}:     &moonFunc{funcName: "func test_fixedarray_output_double_option_1_some() FixedArray[Double?]"},
			{Name: "test_fixedarray_output_double_option_2"}:          &moonFunc{funcName: "func test_fixedarray_output_double_option_2() FixedArray[Double?]"},
			{Name: "test_fixedarray_output_double_option_3"}:          &moonFunc{funcName: "func test_fixedarray_output_double_option_3() FixedArray[Double?]"},
			{Name: "test_fixedarray_output_double_option_4"}:          &moonFunc{funcName: "func test_fixedarray_output_double_option_4() FixedArray[Double?]"},
			{Name: "test_fixedarray_output_float_0"}:                  &moonFunc{funcName: "func test_fixedarray_output_float_0() FixedArray[Float]"},
			{Name: "test_fixedarray_output_float_1"}:                  &moonFunc{funcName: "func test_fixedarray_output_float_1() FixedArray[Float]"},
			{Name: "test_fixedarray_output_float_2"}:                  &moonFunc{funcName: "func test_fixedarray_output_float_2() FixedArray[Float]"},
			{Name: "test_fixedarray_output_float_3"}:                  &moonFunc{funcName: "func test_fixedarray_output_float_3() FixedArray[Float]"},
			{Name: "test_fixedarray_output_float_4"}:                  &moonFunc{funcName: "func test_fixedarray_output_float_4() FixedArray[Float]"},
			{Name: "test_fixedarray_output_float_option_0"}:           &moonFunc{funcName: "func test_fixedarray_output_float_option_0() FixedArray[Float?]"},
			{Name: "test_fixedarray_output_float_option_1_none"}:      &moonFunc{funcName: "func test_fixedarray_output_float_option_1_none() FixedArray[Float?]"},
			{Name: "test_fixedarray_output_float_option_1_some"}:      &moonFunc{funcName: "func test_fixedarray_output_float_option_1_some() FixedArray[Float?]"},
			{Name: "test_fixedarray_output_float_option_2"}:           &moonFunc{funcName: "func test_fixedarray_output_float_option_2() FixedArray[Float?]"},
			{Name: "test_fixedarray_output_float_option_3"}:           &moonFunc{funcName: "func test_fixedarray_output_float_option_3() FixedArray[Float?]"},
			{Name: "test_fixedarray_output_float_option_4"}:           &moonFunc{funcName: "func test_fixedarray_output_float_option_4() FixedArray[Float?]"},
			{Name: "test_fixedarray_output_int16_0"}:                  &moonFunc{funcName: "func test_fixedarray_output_int16_0() FixedArray[Int16]"},
			{Name: "test_fixedarray_output_int16_1"}:                  &moonFunc{funcName: "func test_fixedarray_output_int16_1() FixedArray[Int16]"},
			{Name: "test_fixedarray_output_int16_1_max"}:              &moonFunc{funcName: "func test_fixedarray_output_int16_1_max() FixedArray[Int16]"},
			{Name: "test_fixedarray_output_int16_1_min"}:              &moonFunc{funcName: "func test_fixedarray_output_int16_1_min() FixedArray[Int16]"},
			{Name: "test_fixedarray_output_int16_2"}:                  &moonFunc{funcName: "func test_fixedarray_output_int16_2() FixedArray[Int16]"},
			{Name: "test_fixedarray_output_int16_3"}:                  &moonFunc{funcName: "func test_fixedarray_output_int16_3() FixedArray[Int16]"},
			{Name: "test_fixedarray_output_int16_4"}:                  &moonFunc{funcName: "func test_fixedarray_output_int16_4() FixedArray[Int16]"},
			{Name: "test_fixedarray_output_int16_option_0"}:           &moonFunc{funcName: "func test_fixedarray_output_int16_option_0() FixedArray[Int16?]"},
			{Name: "test_fixedarray_output_int16_option_1_max"}:       &moonFunc{funcName: "func test_fixedarray_output_int16_option_1_max() FixedArray[Int16?]"},
			{Name: "test_fixedarray_output_int16_option_1_min"}:       &moonFunc{funcName: "func test_fixedarray_output_int16_option_1_min() FixedArray[Int16?]"},
			{Name: "test_fixedarray_output_int16_option_1_none"}:      &moonFunc{funcName: "func test_fixedarray_output_int16_option_1_none() FixedArray[Int16?]"},
			{Name: "test_fixedarray_output_int16_option_2"}:           &moonFunc{funcName: "func test_fixedarray_output_int16_option_2() FixedArray[Int16?]"},
			{Name: "test_fixedarray_output_int16_option_3"}:           &moonFunc{funcName: "func test_fixedarray_output_int16_option_3() FixedArray[Int16?]"},
			{Name: "test_fixedarray_output_int16_option_4"}:           &moonFunc{funcName: "func test_fixedarray_output_int16_option_4() FixedArray[Int16?]"},
			{Name: "test_fixedarray_output_int64_0"}:                  &moonFunc{funcName: "func test_fixedarray_output_int64_0() FixedArray[Int64]"},
			{Name: "test_fixedarray_output_int64_1"}:                  &moonFunc{funcName: "func test_fixedarray_output_int64_1() FixedArray[Int64]"},
			{Name: "test_fixedarray_output_int64_1_max"}:              &moonFunc{funcName: "func test_fixedarray_output_int64_1_max() FixedArray[Int64]"},
			{Name: "test_fixedarray_output_int64_1_min"}:              &moonFunc{funcName: "func test_fixedarray_output_int64_1_min() FixedArray[Int64]"},
			{Name: "test_fixedarray_output_int64_2"}:                  &moonFunc{funcName: "func test_fixedarray_output_int64_2() FixedArray[Int64]"},
			{Name: "test_fixedarray_output_int64_3"}:                  &moonFunc{funcName: "func test_fixedarray_output_int64_3() FixedArray[Int64]"},
			{Name: "test_fixedarray_output_int64_4"}:                  &moonFunc{funcName: "func test_fixedarray_output_int64_4() FixedArray[Int64]"},
			{Name: "test_fixedarray_output_int64_option_0"}:           &moonFunc{funcName: "func test_fixedarray_output_int64_option_0() FixedArray[Int64?]"},
			{Name: "test_fixedarray_output_int64_option_1_max"}:       &moonFunc{funcName: "func test_fixedarray_output_int64_option_1_max() FixedArray[Int64?]"},
			{Name: "test_fixedarray_output_int64_option_1_min"}:       &moonFunc{funcName: "func test_fixedarray_output_int64_option_1_min() FixedArray[Int64?]"},
			{Name: "test_fixedarray_output_int64_option_1_none"}:      &moonFunc{funcName: "func test_fixedarray_output_int64_option_1_none() FixedArray[Int64?]"},
			{Name: "test_fixedarray_output_int64_option_2"}:           &moonFunc{funcName: "func test_fixedarray_output_int64_option_2() FixedArray[Int64?]"},
			{Name: "test_fixedarray_output_int64_option_3"}:           &moonFunc{funcName: "func test_fixedarray_output_int64_option_3() FixedArray[Int64?]"},
			{Name: "test_fixedarray_output_int64_option_4"}:           &moonFunc{funcName: "func test_fixedarray_output_int64_option_4() FixedArray[Int64?]"},
			{Name: "test_fixedarray_output_int_0"}:                    &moonFunc{funcName: "func test_fixedarray_output_int_0() FixedArray[Int]"},
			{Name: "test_fixedarray_output_int_1"}:                    &moonFunc{funcName: "func test_fixedarray_output_int_1() FixedArray[Int]"},
			{Name: "test_fixedarray_output_int_1_max"}:                &moonFunc{funcName: "func test_fixedarray_output_int_1_max() FixedArray[Int]"},
			{Name: "test_fixedarray_output_int_1_min"}:                &moonFunc{funcName: "func test_fixedarray_output_int_1_min() FixedArray[Int]"},
			{Name: "test_fixedarray_output_int_2"}:                    &moonFunc{funcName: "func test_fixedarray_output_int_2() FixedArray[Int]"},
			{Name: "test_fixedarray_output_int_3"}:                    &moonFunc{funcName: "func test_fixedarray_output_int_3() FixedArray[Int]"},
			{Name: "test_fixedarray_output_int_4"}:                    &moonFunc{funcName: "func test_fixedarray_output_int_4() FixedArray[Int]"},
			{Name: "test_fixedarray_output_int_option_0"}:             &moonFunc{funcName: "func test_fixedarray_output_int_option_0() FixedArray[Int?]"},
			{Name: "test_fixedarray_output_int_option_1_max"}:         &moonFunc{funcName: "func test_fixedarray_output_int_option_1_max() FixedArray[Int?]"},
			{Name: "test_fixedarray_output_int_option_1_min"}:         &moonFunc{funcName: "func test_fixedarray_output_int_option_1_min() FixedArray[Int?]"},
			{Name: "test_fixedarray_output_int_option_1_none"}:        &moonFunc{funcName: "func test_fixedarray_output_int_option_1_none() FixedArray[Int?]"},
			{Name: "test_fixedarray_output_int_option_2"}:             &moonFunc{funcName: "func test_fixedarray_output_int_option_2() FixedArray[Int?]"},
			{Name: "test_fixedarray_output_int_option_3"}:             &moonFunc{funcName: "func test_fixedarray_output_int_option_3() FixedArray[Int?]"},
			{Name: "test_fixedarray_output_int_option_4"}:             &moonFunc{funcName: "func test_fixedarray_output_int_option_4() FixedArray[Int?]"},
			{Name: "test_fixedarray_output_string_0"}:                 &moonFunc{funcName: "func test_fixedarray_output_string_0() FixedArray[String]"},
			{Name: "test_fixedarray_output_string_1"}:                 &moonFunc{funcName: "func test_fixedarray_output_string_1() FixedArray[String]"},
			{Name: "test_fixedarray_output_string_2"}:                 &moonFunc{funcName: "func test_fixedarray_output_string_2() FixedArray[String]"},
			{Name: "test_fixedarray_output_string_3"}:                 &moonFunc{funcName: "func test_fixedarray_output_string_3() FixedArray[String]"},
			{Name: "test_fixedarray_output_string_4"}:                 &moonFunc{funcName: "func test_fixedarray_output_string_4() FixedArray[String]"},
			{Name: "test_fixedarray_output_string_option_0"}:          &moonFunc{funcName: "func test_fixedarray_output_string_option_0() FixedArray[String?]"},
			{Name: "test_fixedarray_output_string_option_1_none"}:     &moonFunc{funcName: "func test_fixedarray_output_string_option_1_none() FixedArray[String?]"},
			{Name: "test_fixedarray_output_string_option_1_some"}:     &moonFunc{funcName: "func test_fixedarray_output_string_option_1_some() FixedArray[String?]"},
			{Name: "test_fixedarray_output_string_option_2"}:          &moonFunc{funcName: "func test_fixedarray_output_string_option_2() FixedArray[String?]"},
			{Name: "test_fixedarray_output_string_option_3"}:          &moonFunc{funcName: "func test_fixedarray_output_string_option_3() FixedArray[String?]"},
			{Name: "test_fixedarray_output_string_option_4"}:          &moonFunc{funcName: "func test_fixedarray_output_string_option_4() FixedArray[String?]"},
			{Name: "test_fixedarray_output_uint16_0"}:                 &moonFunc{funcName: "func test_fixedarray_output_uint16_0() FixedArray[UInt16]"},
			{Name: "test_fixedarray_output_uint16_1"}:                 &moonFunc{funcName: "func test_fixedarray_output_uint16_1() FixedArray[UInt16]"},
			{Name: "test_fixedarray_output_uint16_1_max"}:             &moonFunc{funcName: "func test_fixedarray_output_uint16_1_max() FixedArray[UInt16]"},
			{Name: "test_fixedarray_output_uint16_1_min"}:             &moonFunc{funcName: "func test_fixedarray_output_uint16_1_min() FixedArray[UInt16]"},
			{Name: "test_fixedarray_output_uint16_2"}:                 &moonFunc{funcName: "func test_fixedarray_output_uint16_2() FixedArray[UInt16]"},
			{Name: "test_fixedarray_output_uint16_3"}:                 &moonFunc{funcName: "func test_fixedarray_output_uint16_3() FixedArray[UInt16]"},
			{Name: "test_fixedarray_output_uint16_4"}:                 &moonFunc{funcName: "func test_fixedarray_output_uint16_4() FixedArray[UInt16]"},
			{Name: "test_fixedarray_output_uint16_option_0"}:          &moonFunc{funcName: "func test_fixedarray_output_uint16_option_0() FixedArray[UInt16?]"},
			{Name: "test_fixedarray_output_uint16_option_1_max"}:      &moonFunc{funcName: "func test_fixedarray_output_uint16_option_1_max() FixedArray[UInt16?]"},
			{Name: "test_fixedarray_output_uint16_option_1_min"}:      &moonFunc{funcName: "func test_fixedarray_output_uint16_option_1_min() FixedArray[UInt16?]"},
			{Name: "test_fixedarray_output_uint16_option_1_none"}:     &moonFunc{funcName: "func test_fixedarray_output_uint16_option_1_none() FixedArray[UInt16?]"},
			{Name: "test_fixedarray_output_uint16_option_2"}:          &moonFunc{funcName: "func test_fixedarray_output_uint16_option_2() FixedArray[UInt16?]"},
			{Name: "test_fixedarray_output_uint16_option_3"}:          &moonFunc{funcName: "func test_fixedarray_output_uint16_option_3() FixedArray[UInt16?]"},
			{Name: "test_fixedarray_output_uint16_option_4"}:          &moonFunc{funcName: "func test_fixedarray_output_uint16_option_4() FixedArray[UInt16?]"},
			{Name: "test_fixedarray_output_uint64_0"}:                 &moonFunc{funcName: "func test_fixedarray_output_uint64_0() FixedArray[UInt64]"},
			{Name: "test_fixedarray_output_uint64_1"}:                 &moonFunc{funcName: "func test_fixedarray_output_uint64_1() FixedArray[UInt64]"},
			{Name: "test_fixedarray_output_uint64_1_max"}:             &moonFunc{funcName: "func test_fixedarray_output_uint64_1_max() FixedArray[UInt64]"},
			{Name: "test_fixedarray_output_uint64_1_min"}:             &moonFunc{funcName: "func test_fixedarray_output_uint64_1_min() FixedArray[UInt64]"},
			{Name: "test_fixedarray_output_uint64_2"}:                 &moonFunc{funcName: "func test_fixedarray_output_uint64_2() FixedArray[UInt64]"},
			{Name: "test_fixedarray_output_uint64_3"}:                 &moonFunc{funcName: "func test_fixedarray_output_uint64_3() FixedArray[UInt64]"},
			{Name: "test_fixedarray_output_uint64_4"}:                 &moonFunc{funcName: "func test_fixedarray_output_uint64_4() FixedArray[UInt64]"},
			{Name: "test_fixedarray_output_uint64_option_0"}:          &moonFunc{funcName: "func test_fixedarray_output_uint64_option_0() FixedArray[UInt64?]"},
			{Name: "test_fixedarray_output_uint64_option_1_max"}:      &moonFunc{funcName: "func test_fixedarray_output_uint64_option_1_max() FixedArray[UInt64?]"},
			{Name: "test_fixedarray_output_uint64_option_1_min"}:      &moonFunc{funcName: "func test_fixedarray_output_uint64_option_1_min() FixedArray[UInt64?]"},
			{Name: "test_fixedarray_output_uint64_option_1_none"}:     &moonFunc{funcName: "func test_fixedarray_output_uint64_option_1_none() FixedArray[UInt64?]"},
			{Name: "test_fixedarray_output_uint64_option_2"}:          &moonFunc{funcName: "func test_fixedarray_output_uint64_option_2() FixedArray[UInt64?]"},
			{Name: "test_fixedarray_output_uint64_option_3"}:          &moonFunc{funcName: "func test_fixedarray_output_uint64_option_3() FixedArray[UInt64?]"},
			{Name: "test_fixedarray_output_uint64_option_4"}:          &moonFunc{funcName: "func test_fixedarray_output_uint64_option_4() FixedArray[UInt64?]"},
			{Name: "test_fixedarray_output_uint_0"}:                   &moonFunc{funcName: "func test_fixedarray_output_uint_0() FixedArray[UInt]"},
			{Name: "test_fixedarray_output_uint_1"}:                   &moonFunc{funcName: "func test_fixedarray_output_uint_1() FixedArray[UInt]"},
			{Name: "test_fixedarray_output_uint_1_max"}:               &moonFunc{funcName: "func test_fixedarray_output_uint_1_max() FixedArray[UInt]"},
			{Name: "test_fixedarray_output_uint_1_min"}:               &moonFunc{funcName: "func test_fixedarray_output_uint_1_min() FixedArray[UInt]"},
			{Name: "test_fixedarray_output_uint_2"}:                   &moonFunc{funcName: "func test_fixedarray_output_uint_2() FixedArray[UInt]"},
			{Name: "test_fixedarray_output_uint_3"}:                   &moonFunc{funcName: "func test_fixedarray_output_uint_3() FixedArray[UInt]"},
			{Name: "test_fixedarray_output_uint_4"}:                   &moonFunc{funcName: "func test_fixedarray_output_uint_4() FixedArray[UInt]"},
			{Name: "test_fixedarray_output_uint_option_0"}:            &moonFunc{funcName: "func test_fixedarray_output_uint_option_0() FixedArray[UInt?]"},
			{Name: "test_fixedarray_output_uint_option_1_max"}:        &moonFunc{funcName: "func test_fixedarray_output_uint_option_1_max() FixedArray[UInt?]"},
			{Name: "test_fixedarray_output_uint_option_1_min"}:        &moonFunc{funcName: "func test_fixedarray_output_uint_option_1_min() FixedArray[UInt?]"},
			{Name: "test_fixedarray_output_uint_option_1_none"}:       &moonFunc{funcName: "func test_fixedarray_output_uint_option_1_none() FixedArray[UInt?]"},
			{Name: "test_fixedarray_output_uint_option_2"}:            &moonFunc{funcName: "func test_fixedarray_output_uint_option_2() FixedArray[UInt?]"},
			{Name: "test_fixedarray_output_uint_option_3"}:            &moonFunc{funcName: "func test_fixedarray_output_uint_option_3() FixedArray[UInt?]"},
			{Name: "test_fixedarray_output_uint_option_4"}:            &moonFunc{funcName: "func test_fixedarray_output_uint_option_4() FixedArray[UInt?]"},
			{Name: "test_float_input_max"}:                            &moonFunc{funcName: "func test_float_input_max(n Float) Unit!Error"},
			{Name: "test_float_input_min"}:                            &moonFunc{funcName: "func test_float_input_min(n Float) Unit!Error"},
			{Name: "test_float_option_input_max"}:                     &moonFunc{funcName: "func test_float_option_input_max(n Float?) Unit!Error"},
			{Name: "test_float_option_input_min"}:                     &moonFunc{funcName: "func test_float_option_input_min(n Float?) Unit!Error"},
			{Name: "test_float_option_input_none"}:                    &moonFunc{funcName: "func test_float_option_input_none(n Float?) Unit!Error"},
			{Name: "test_float_option_output_max"}:                    &moonFunc{funcName: "func test_float_option_output_max() Float?"},
			{Name: "test_float_option_output_min"}:                    &moonFunc{funcName: "func test_float_option_output_min() Float?"},
			{Name: "test_float_option_output_none"}:                   &moonFunc{funcName: "func test_float_option_output_none() Float?"},
			{Name: "test_float_output_max"}:                           &moonFunc{funcName: "func test_float_output_max() Float"},
			{Name: "test_float_output_min"}:                           &moonFunc{funcName: "func test_float_output_min() Float"},
			{Name: "test_generate_map_string_string_output"}:          &moonFunc{funcName: "func test_generate_map_string_string_output() Map[String, String]"},
			{Name: "test_http_header"}:                                &moonFunc{funcName: "func test_http_header(h HttpHeader?)"},
			{Name: "test_http_header_map"}:                            &moonFunc{funcName: "func test_http_header_map(m Map[String, HttpHeader?])"},
			{Name: "test_http_headers"}:                               &moonFunc{funcName: "func test_http_headers(h HttpHeaders)"},
			{Name: "test_http_response_headers"}:                      &moonFunc{funcName: "func test_http_response_headers(r HttpResponse?)"},
			{Name: "test_http_response_headers_output"}:               &moonFunc{funcName: "func test_http_response_headers_output() HttpResponse?"},
			{Name: "test_int16_input_max"}:                            &moonFunc{funcName: "func test_int16_input_max(n Int16) Unit!Error"},
			{Name: "test_int16_input_min"}:                            &moonFunc{funcName: "func test_int16_input_min(n Int16) Unit!Error"},
			{Name: "test_int16_option_input_max"}:                     &moonFunc{funcName: "func test_int16_option_input_max(n Int16?) Unit!Error"},
			{Name: "test_int16_option_input_min"}:                     &moonFunc{funcName: "func test_int16_option_input_min(n Int16?) Unit!Error"},
			{Name: "test_int16_option_input_none"}:                    &moonFunc{funcName: "func test_int16_option_input_none(n Int16?) Unit!Error"},
			{Name: "test_int16_option_output_max"}:                    &moonFunc{funcName: "func test_int16_option_output_max() Int16?"},
			{Name: "test_int16_option_output_min"}:                    &moonFunc{funcName: "func test_int16_option_output_min() Int16?"},
			{Name: "test_int16_option_output_none"}:                   &moonFunc{funcName: "func test_int16_option_output_none() Int16?"},
			{Name: "test_int16_output_max"}:                           &moonFunc{funcName: "func test_int16_output_max() Int16"},
			{Name: "test_int16_output_min"}:                           &moonFunc{funcName: "func test_int16_output_min() Int16"},
			{Name: "test_int64_input_max"}:                            &moonFunc{funcName: "func test_int64_input_max(n Int64) Unit!Error"},
			{Name: "test_int64_input_min"}:                            &moonFunc{funcName: "func test_int64_input_min(n Int64) Unit!Error"},
			{Name: "test_int64_option_input_max"}:                     &moonFunc{funcName: "func test_int64_option_input_max(n Int64?) Unit!Error"},
			{Name: "test_int64_option_input_min"}:                     &moonFunc{funcName: "func test_int64_option_input_min(n Int64?) Unit!Error"},
			{Name: "test_int64_option_input_none"}:                    &moonFunc{funcName: "func test_int64_option_input_none(n Int64?) Unit!Error"},
			{Name: "test_int64_option_output_max"}:                    &moonFunc{funcName: "func test_int64_option_output_max() Int64?"},
			{Name: "test_int64_option_output_min"}:                    &moonFunc{funcName: "func test_int64_option_output_min() Int64?"},
			{Name: "test_int64_option_output_none"}:                   &moonFunc{funcName: "func test_int64_option_output_none() Int64?"},
			{Name: "test_int64_output_max"}:                           &moonFunc{funcName: "func test_int64_output_max() Int64"},
			{Name: "test_int64_output_min"}:                           &moonFunc{funcName: "func test_int64_output_min() Int64"},
			{Name: "test_int_input_max"}:                              &moonFunc{funcName: "func test_int_input_max(n Int) Unit!Error"},
			{Name: "test_int_input_min"}:                              &moonFunc{funcName: "func test_int_input_min(n Int) Unit!Error"},
			{Name: "test_int_option_input_max"}:                       &moonFunc{funcName: "func test_int_option_input_max(n Int?) Unit!Error"},
			{Name: "test_int_option_input_min"}:                       &moonFunc{funcName: "func test_int_option_input_min(n Int?) Unit!Error"},
			{Name: "test_int_option_input_none"}:                      &moonFunc{funcName: "func test_int_option_input_none(n Int?) Unit!Error"},
			{Name: "test_int_option_output_max"}:                      &moonFunc{funcName: "func test_int_option_output_max() Int?"},
			{Name: "test_int_option_output_min"}:                      &moonFunc{funcName: "func test_int_option_output_min() Int?"},
			{Name: "test_int_option_output_none"}:                     &moonFunc{funcName: "func test_int_option_output_none() Int?"},
			{Name: "test_int_output_max"}:                             &moonFunc{funcName: "func test_int_output_max() Int"},
			{Name: "test_int_output_min"}:                             &moonFunc{funcName: "func test_int_output_min() Int"},
			{Name: "test_iterate_map_string_string"}:                  &moonFunc{funcName: "func test_iterate_map_string_string(m Map[String, String])"},
			{Name: "test_map_input_int_double"}:                       &moonFunc{funcName: "func test_map_input_int_double(m Map[Int, Double]) Unit!Error"},
			{Name: "test_map_input_int_float"}:                        &moonFunc{funcName: "func test_map_input_int_float(m Map[Int, Float]) Unit!Error"},
			{Name: "test_map_input_string_string"}:                    &moonFunc{funcName: "func test_map_input_string_string(m Map[String, String]) Unit!Error"},
			{Name: "test_map_lookup_string_string"}:                   &moonFunc{funcName: "func test_map_lookup_string_string(m Map[String, String], key String) String"},
			{Name: "test_map_option_input_string_string"}:             &moonFunc{funcName: "func test_map_option_input_string_string(m Map[String, String]?) Unit!Error"},
			{Name: "test_map_option_output_string_string"}:            &moonFunc{funcName: "func test_map_option_output_string_string() Map[String, String]?"},
			{Name: "test_map_output_int_double"}:                      &moonFunc{funcName: "func test_map_output_int_double() Map[Int, Double]"},
			{Name: "test_map_output_int_float"}:                       &moonFunc{funcName: "func test_map_output_int_float() Map[Int, Float]"},
			{Name: "test_map_output_string_string"}:                   &moonFunc{funcName: "func test_map_output_string_string() Map[String, String]"},
			{Name: "test_option_fixedarray_input1_int"}:               &moonFunc{funcName: "func test_option_fixedarray_input1_int(val FixedArray[Int]?) Unit!Error"},
			{Name: "test_option_fixedarray_input1_string"}:            &moonFunc{funcName: "func test_option_fixedarray_input1_string(val FixedArray[String]?) Unit!Error"},
			{Name: "test_option_fixedarray_input2_int"}:               &moonFunc{funcName: "func test_option_fixedarray_input2_int(val FixedArray[Int]?) Unit!Error"},
			{Name: "test_option_fixedarray_input2_string"}:            &moonFunc{funcName: "func test_option_fixedarray_input2_string(val FixedArray[String]?) Unit!Error"},
			{Name: "test_option_fixedarray_output1_int"}:              &moonFunc{funcName: "func test_option_fixedarray_output1_int() FixedArray[Int]?"},
			{Name: "test_option_fixedarray_output1_string"}:           &moonFunc{funcName: "func test_option_fixedarray_output1_string() FixedArray[String]?"},
			{Name: "test_option_fixedarray_output2_int"}:              &moonFunc{funcName: "func test_option_fixedarray_output2_int() FixedArray[Int]?"},
			{Name: "test_option_fixedarray_output2_string"}:           &moonFunc{funcName: "func test_option_fixedarray_output2_string() FixedArray[String]?"},
			{Name: "test_recursive_struct_input"}:                     &moonFunc{funcName: "func test_recursive_struct_input(r1 TestRecursiveStruct) Unit!Error"},
			{Name: "test_recursive_struct_option_input"}:              &moonFunc{funcName: "func test_recursive_struct_option_input(o TestRecursiveStruct?) Unit!Error"},
			{Name: "test_recursive_struct_option_input_none"}:         &moonFunc{funcName: "func test_recursive_struct_option_input_none(o TestRecursiveStruct?) Unit!Error"},
			{Name: "test_recursive_struct_option_output"}:             &moonFunc{funcName: "func test_recursive_struct_option_output() TestRecursiveStruct?"},
			{Name: "test_recursive_struct_option_output_map"}:         &moonFunc{funcName: "func test_recursive_struct_option_output_map() TestRecursiveStruct_map?"},
			{Name: "test_recursive_struct_option_output_none"}:        &moonFunc{funcName: "func test_recursive_struct_option_output_none() TestRecursiveStruct?"},
			{Name: "test_recursive_struct_output"}:                    &moonFunc{funcName: "func test_recursive_struct_output() TestRecursiveStruct"},
			{Name: "test_recursive_struct_output_map"}:                &moonFunc{funcName: "func test_recursive_struct_output_map() TestRecursiveStruct_map"},
			{Name: "test_smorgasbord_struct_input"}:                   &moonFunc{funcName: "func test_smorgasbord_struct_input(o TestSmorgasbordStruct) Unit!Error"},
			{Name: "test_smorgasbord_struct_option_input"}:            &moonFunc{funcName: "func test_smorgasbord_struct_option_input(o TestSmorgasbordStruct?) Unit!Error"},
			{Name: "test_smorgasbord_struct_option_input_none"}:       &moonFunc{funcName: "func test_smorgasbord_struct_option_input_none(o TestSmorgasbordStruct?) Unit!Error"},
			{Name: "test_smorgasbord_struct_option_output"}:           &moonFunc{funcName: "func test_smorgasbord_struct_option_output() TestSmorgasbordStruct?"},
			{Name: "test_smorgasbord_struct_option_output_map"}:       &moonFunc{funcName: "func test_smorgasbord_struct_option_output_map() TestSmorgasbordStruct_map?"},
			{Name: "test_smorgasbord_struct_option_output_map_none"}:  &moonFunc{funcName: "func test_smorgasbord_struct_option_output_map_none() TestSmorgasbordStruct_map?"},
			{Name: "test_smorgasbord_struct_option_output_none"}:      &moonFunc{funcName: "func test_smorgasbord_struct_option_output_none() TestSmorgasbordStruct?"},
			{Name: "test_smorgasbord_struct_output"}:                  &moonFunc{funcName: "func test_smorgasbord_struct_output() TestSmorgasbordStruct"},
			{Name: "test_smorgasbord_struct_output_map"}:              &moonFunc{funcName: "func test_smorgasbord_struct_output_map() TestSmorgasbordStruct_map"},
			{Name: "test_string_input"}:                               &moonFunc{funcName: "func test_string_input(s String) Unit!Error"},
			{Name: "test_string_option_input"}:                        &moonFunc{funcName: "func test_string_option_input(s String?) Unit!Error"},
			{Name: "test_string_option_input_none"}:                   &moonFunc{funcName: "func test_string_option_input_none(s String?) Unit!Error"},
			{Name: "test_string_option_output"}:                       &moonFunc{funcName: "func test_string_option_output() String?"},
			{Name: "test_string_option_output_none"}:                  &moonFunc{funcName: "func test_string_option_output_none() String?"},
			{Name: "test_string_output"}:                              &moonFunc{funcName: "func test_string_output() String"},
			{Name: "test_struct_containing_map_input_string_string"}:  &moonFunc{funcName: "func test_struct_containing_map_input_string_string(s TestStructWithMap) Unit!Error"},
			{Name: "test_struct_containing_map_output_string_string"}: &moonFunc{funcName: "func test_struct_containing_map_output_string_string() TestStructWithMap"},
			{Name: "test_struct_input1"}:                              &moonFunc{funcName: "func test_struct_input1(o TestStruct1) Unit!Error"},
			{Name: "test_struct_input2"}:                              &moonFunc{funcName: "func test_struct_input2(o TestStruct2) Unit!Error"},
			{Name: "test_struct_input3"}:                              &moonFunc{funcName: "func test_struct_input3(o TestStruct3) Unit!Error"},
			{Name: "test_struct_input4"}:                              &moonFunc{funcName: "func test_struct_input4(o TestStruct4) Unit!Error"},
			{Name: "test_struct_input4_with_none"}:                    &moonFunc{funcName: "func test_struct_input4_with_none(o TestStruct4) Unit!Error"},
			{Name: "test_struct_input5"}:                              &moonFunc{funcName: "func test_struct_input5(o TestStruct5) Unit!Error"},
			{Name: "test_struct_option_input1"}:                       &moonFunc{funcName: "func test_struct_option_input1(o TestStruct1?) Unit!Error"},
			{Name: "test_struct_option_input1_none"}:                  &moonFunc{funcName: "func test_struct_option_input1_none(o TestStruct1?) Unit!Error"},
			{Name: "test_struct_option_input2"}:                       &moonFunc{funcName: "func test_struct_option_input2(o TestStruct2?) Unit!Error"},
			{Name: "test_struct_option_input2_none"}:                  &moonFunc{funcName: "func test_struct_option_input2_none(o TestStruct2?) Unit!Error"},
			{Name: "test_struct_option_input3"}:                       &moonFunc{funcName: "func test_struct_option_input3(o TestStruct3?) Unit!Error"},
			{Name: "test_struct_option_input3_none"}:                  &moonFunc{funcName: "func test_struct_option_input3_none(o TestStruct3?) Unit!Error"},
			{Name: "test_struct_option_input4"}:                       &moonFunc{funcName: "func test_struct_option_input4(o TestStruct4?) Unit!Error"},
			{Name: "test_struct_option_input4_none"}:                  &moonFunc{funcName: "func test_struct_option_input4_none(o TestStruct4?) Unit!Error"},
			{Name: "test_struct_option_input4_with_none"}:             &moonFunc{funcName: "func test_struct_option_input4_with_none(o TestStruct4?) Unit!Error"},
			{Name: "test_struct_option_input5"}:                       &moonFunc{funcName: "func test_struct_option_input5(o TestStruct5?) Unit!Error"},
			{Name: "test_struct_option_input5_none"}:                  &moonFunc{funcName: "func test_struct_option_input5_none(o TestStruct5?) Unit!Error"},
			{Name: "test_struct_option_output1"}:                      &moonFunc{funcName: "func test_struct_option_output1() TestStruct1?"},
			{Name: "test_struct_option_output1_map"}:                  &moonFunc{funcName: "func test_struct_option_output1_map() TestStruct1_map?"},
			{Name: "test_struct_option_output1_none"}:                 &moonFunc{funcName: "func test_struct_option_output1_none() TestStruct1?"},
			{Name: "test_struct_option_output2"}:                      &moonFunc{funcName: "func test_struct_option_output2() TestStruct2?"},
			{Name: "test_struct_option_output2_map"}:                  &moonFunc{funcName: "func test_struct_option_output2_map() TestStruct2_map?"},
			{Name: "test_struct_option_output2_none"}:                 &moonFunc{funcName: "func test_struct_option_output2_none() TestStruct2?"},
			{Name: "test_struct_option_output3"}:                      &moonFunc{funcName: "func test_struct_option_output3() TestStruct3?"},
			{Name: "test_struct_option_output3_map"}:                  &moonFunc{funcName: "func test_struct_option_output3_map() TestStruct3_map?"},
			{Name: "test_struct_option_output3_none"}:                 &moonFunc{funcName: "func test_struct_option_output3_none() TestStruct3?"},
			{Name: "test_struct_option_output4"}:                      &moonFunc{funcName: "func test_struct_option_output4() TestStruct4?"},
			{Name: "test_struct_option_output4_map"}:                  &moonFunc{funcName: "func test_struct_option_output4_map() TestStruct4_map?"},
			{Name: "test_struct_option_output4_map_with_none"}:        &moonFunc{funcName: "func test_struct_option_output4_map_with_none() TestStruct4_map?"},
			{Name: "test_struct_option_output4_none"}:                 &moonFunc{funcName: "func test_struct_option_output4_none() TestStruct4?"},
			{Name: "test_struct_option_output4_with_none"}:            &moonFunc{funcName: "func test_struct_option_output4_with_none() TestStruct4?"},
			{Name: "test_struct_option_output5"}:                      &moonFunc{funcName: "func test_struct_option_output5() TestStruct5?"},
			{Name: "test_struct_option_output5_map"}:                  &moonFunc{funcName: "func test_struct_option_output5_map() TestStruct5_map?"},
			{Name: "test_struct_option_output5_none"}:                 &moonFunc{funcName: "func test_struct_option_output5_none() TestStruct5?"},
			{Name: "test_struct_output1"}:                             &moonFunc{funcName: "func test_struct_output1() TestStruct1"},
			{Name: "test_struct_output1_map"}:                         &moonFunc{funcName: "func test_struct_output1_map() TestStruct1_map"},
			{Name: "test_struct_output2"}:                             &moonFunc{funcName: "func test_struct_output2() TestStruct2"},
			{Name: "test_struct_output2_map"}:                         &moonFunc{funcName: "func test_struct_output2_map() TestStruct2_map"},
			{Name: "test_struct_output3"}:                             &moonFunc{funcName: "func test_struct_output3() TestStruct3"},
			{Name: "test_struct_output3_map"}:                         &moonFunc{funcName: "func test_struct_output3_map() TestStruct3_map"},
			{Name: "test_struct_output4"}:                             &moonFunc{funcName: "func test_struct_output4() TestStruct4"},
			{Name: "test_struct_output4_map"}:                         &moonFunc{funcName: "func test_struct_output4_map() TestStruct4_map"},
			{Name: "test_struct_output4_map_with_none"}:               &moonFunc{funcName: "func test_struct_output4_map_with_none() TestStruct4_map"},
			{Name: "test_struct_output4_with_none"}:                   &moonFunc{funcName: "func test_struct_output4_with_none() TestStruct4"},
			{Name: "test_struct_output5"}:                             &moonFunc{funcName: "func test_struct_output5() TestStruct5"},
			{Name: "test_struct_output5_map"}:                         &moonFunc{funcName: "func test_struct_output5_map() TestStruct5_map"},
			{Name: "test_time_input"}:                                 &moonFunc{funcName: "func test_time_input(t @time.ZonedDateTime) Unit!Error"},
			{Name: "test_time_option_input"}:                          &moonFunc{funcName: "func test_time_option_input(t @time.ZonedDateTime?) Unit!Error"},
			{Name: "test_time_option_input_none"}:                     &moonFunc{funcName: "func test_time_option_input_none(t @time.ZonedDateTime?) Unit!Error"},
			{Name: "test_time_option_input_none_style2"}:              &moonFunc{funcName: "func test_time_option_input_none_style2(t? @time.ZonedDateTime) Unit!Error"},
			{Name: "test_time_option_input_style2"}:                   &moonFunc{funcName: "func test_time_option_input_style2(t? @time.ZonedDateTime) Unit!Error"},
			{Name: "test_time_option_output"}:                         &moonFunc{funcName: "func test_time_option_output() @time.ZonedDateTime?"},
			{Name: "test_time_option_output_none"}:                    &moonFunc{funcName: "func test_time_option_output_none() @time.ZonedDateTime?"},
			{Name: "test_time_output"}:                                &moonFunc{funcName: "func test_time_output() @time.ZonedDateTime"},
			{Name: "test_tuple_output"}:                               &moonFunc{funcName: "func test_tuple_output() (Int, Bool, String)"},
			{Name: "test_uint16_input_max"}:                           &moonFunc{funcName: "func test_uint16_input_max(n UInt16) Unit!Error"},
			{Name: "test_uint16_input_min"}:                           &moonFunc{funcName: "func test_uint16_input_min(n UInt16) Unit!Error"},
			{Name: "test_uint16_option_input_max"}:                    &moonFunc{funcName: "func test_uint16_option_input_max(n UInt16?) Unit!Error"},
			{Name: "test_uint16_option_input_min"}:                    &moonFunc{funcName: "func test_uint16_option_input_min(n UInt16?) Unit!Error"},
			{Name: "test_uint16_option_input_none"}:                   &moonFunc{funcName: "func test_uint16_option_input_none(n UInt16?) Unit!Error"},
			{Name: "test_uint16_option_output_max"}:                   &moonFunc{funcName: "func test_uint16_option_output_max() UInt16?"},
			{Name: "test_uint16_option_output_min"}:                   &moonFunc{funcName: "func test_uint16_option_output_min() UInt16?"},
			{Name: "test_uint16_option_output_none"}:                  &moonFunc{funcName: "func test_uint16_option_output_none() UInt16?"},
			{Name: "test_uint16_output_max"}:                          &moonFunc{funcName: "func test_uint16_output_max() UInt16"},
			{Name: "test_uint16_output_min"}:                          &moonFunc{funcName: "func test_uint16_output_min() UInt16"},
			{Name: "test_uint64_input_max"}:                           &moonFunc{funcName: "func test_uint64_input_max(n UInt64) Unit!Error"},
			{Name: "test_uint64_input_min"}:                           &moonFunc{funcName: "func test_uint64_input_min(n UInt64) Unit!Error"},
			{Name: "test_uint64_option_input_max"}:                    &moonFunc{funcName: "func test_uint64_option_input_max(n UInt64?) Unit!Error"},
			{Name: "test_uint64_option_input_min"}:                    &moonFunc{funcName: "func test_uint64_option_input_min(n UInt64?) Unit!Error"},
			{Name: "test_uint64_option_input_none"}:                   &moonFunc{funcName: "func test_uint64_option_input_none(n UInt64?) Unit!Error"},
			{Name: "test_uint64_option_output_max"}:                   &moonFunc{funcName: "func test_uint64_option_output_max() UInt64?"},
			{Name: "test_uint64_option_output_min"}:                   &moonFunc{funcName: "func test_uint64_option_output_min() UInt64?"},
			{Name: "test_uint64_option_output_none"}:                  &moonFunc{funcName: "func test_uint64_option_output_none() UInt64?"},
			{Name: "test_uint64_output_max"}:                          &moonFunc{funcName: "func test_uint64_output_max() UInt64"},
			{Name: "test_uint64_output_min"}:                          &moonFunc{funcName: "func test_uint64_output_min() UInt64"},
			{Name: "test_uint_input_max"}:                             &moonFunc{funcName: "func test_uint_input_max(n UInt) Unit!Error"},
			{Name: "test_uint_input_min"}:                             &moonFunc{funcName: "func test_uint_input_min(n UInt) Unit!Error"},
			{Name: "test_uint_option_input_max"}:                      &moonFunc{funcName: "func test_uint_option_input_max(n UInt?) Unit!Error"},
			{Name: "test_uint_option_input_min"}:                      &moonFunc{funcName: "func test_uint_option_input_min(n UInt?) Unit!Error"},
			{Name: "test_uint_option_input_none"}:                     &moonFunc{funcName: "func test_uint_option_input_none(n UInt?) Unit!Error"},
			{Name: "test_uint_option_output_max"}:                     &moonFunc{funcName: "func test_uint_option_output_max() UInt?"},
			{Name: "test_uint_option_output_min"}:                     &moonFunc{funcName: "func test_uint_option_output_min() UInt?"},
			{Name: "test_uint_option_output_none"}:                    &moonFunc{funcName: "func test_uint_option_output_none() UInt?"},
			{Name: "test_uint_output_max"}:                            &moonFunc{funcName: "func test_uint_output_max() UInt"},
			{Name: "test_uint_output_min"}:                            &moonFunc{funcName: "func test_uint_output_min() UInt"},
		},
	},
}
