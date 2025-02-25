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
	"context"
	"strings"
	"testing"

	"github.com/gmlewis/modus/runtime/hostfunctions"
	"github.com/gmlewis/modus/runtime/testutils"
	"github.com/gmlewis/modus/runtime/utils"
	"github.com/gmlewis/modus/runtime/wasmhost"
)

func getTestHostFunctionRegistrations() []func(wasmhost.WasmHost) error {
	return []func(wasmhost.WasmHost) error{
		func(host wasmhost.WasmHost) error {
			return host.RegisterHostFunction("modus_system", "logMessage", hostLog)
		},
		func(host wasmhost.WasmHost) error {
			return host.RegisterHostFunction("modus_test", "add", hostAdd)
		},
		func(host wasmhost.WasmHost) error {
			return host.RegisterHostFunction("modus_test", "echo1", hostEcho1)
		},
		func(host wasmhost.WasmHost) error {
			return host.RegisterHostFunction("modus_test", "echo2", hostEcho2)
		},
		func(host wasmhost.WasmHost) error {
			return host.RegisterHostFunction("modus_test", "echo3", hostEcho3)
		},
		func(host wasmhost.WasmHost) error {
			return host.RegisterHostFunction("modus_test", "echo4", hostEcho4)
		},
		func(host wasmhost.WasmHost) error {
			return host.RegisterHostFunction("modus_test", "encodeStrings1", hostEncodeStrings1)
		},
		func(host wasmhost.WasmHost) error {
			return host.RegisterHostFunction("modus_test", "encodeStrings2", hostEncodeStrings2)
		},
	}
}

func hostLog(ctx context.Context, level, message string) {
	if utils.DebugModeEnabled() {
		hostfunctions.LogMessage(ctx, level, message)
	}
	t := testutils.GetTestT(ctx)
	t.Logf("[%v] %v", level, message)
}

func hostAdd(a, b int32) int32 {
	return a + b
}

func hostEcho1(s string) string {
	return "echo: " + s
}

func hostEcho2(s *string) string {
	return "echo: " + *s
}

func hostEcho3(s string) *string {
	result := "echo: " + s
	return &result
}

func hostEcho4(s *string) *string {
	result := "echo: " + *s
	return &result
}

func hostEncodeStrings1(items *[]string) *string {
	out := strings.Builder{}
	out.WriteString("[")
	for i, item := range *items {
		if i > 0 {
			out.WriteString(",")
		}
		b, _ := utils.JsonSerialize(item)
		out.Write(b)
	}
	out.WriteString("]")

	result := out.String()
	return &result
}

func hostEncodeStrings2(items *[]*string) *string {
	out := strings.Builder{}
	out.WriteString("[")
	for i, item := range *items {
		if i > 0 {
			out.WriteString(",")
		}
		b, _ := utils.JsonSerialize(item)
		out.Write(b)
	}
	out.WriteString("]")

	result := out.String()
	return &result
}

func TestHostFn_add(t *testing.T) {
	fnName := "add"
	result, err := fixture.CallFunction(t, fnName, int32(1), int32(2))
	if err != nil {
		t.Fatal(err)
	}

	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(int32); !ok {
		t.Errorf("expected an int32, got %T", result)
	} else if r != 3 {
		t.Errorf("expected %v, got %v", 3, r)
	}
}

func TestHostFn_echo1_string(t *testing.T) {
	fnName := "echo1"
	result, err := fixture.CallFunction(t, fnName, "hello")
	if err != nil {
		t.Fatal(err)
	}

	expected := "echo: hello"
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestHostFn_echo1_string_option(t *testing.T) {
	fnName := "echo1"
	s := "hello"

	result, err := fixture.CallFunction(t, fnName, &s)
	if err != nil {
		t.Fatal(err)
	}

	expected := "echo: hello"
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestHostFn_echo2_string(t *testing.T) {
	fnName := "echo2"
	result, err := fixture.CallFunction(t, fnName, "hello")
	if err != nil {
		t.Fatal(err)
	}

	expected := "echo: hello"
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestHostFn_echo2_string_option(t *testing.T) {
	fnName := "echo2"
	s := "hello"

	result, err := fixture.CallFunction(t, fnName, &s)
	if err != nil {
		t.Fatal(err)
	}

	expected := "echo: hello"
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if r != expected {
		t.Errorf("expected %v, got %v", expected, r)
	}
}

func TestHostFn_echo3_string(t *testing.T) {
	fnName := "echo3"
	result, err := fixture.CallFunction(t, fnName, "hello")
	if err != nil {
		t.Fatal(err)
	}

	expected := "echo: hello"
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestHostFn_echo3_string_option(t *testing.T) {
	fnName := "echo3"
	s := "hello"

	result, err := fixture.CallFunction(t, fnName, &s)
	if err != nil {
		t.Fatal(err)
	}

	expected := "echo: hello"
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestHostFn_echo4_string(t *testing.T) {
	fnName := "echo4"
	result, err := fixture.CallFunction(t, fnName, "hello")
	if err != nil {
		t.Fatal(err)
	}

	expected := "echo: hello"
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestHostFn_echo4_string_option(t *testing.T) {
	fnName := "echo4"
	s := "hello"

	result, err := fixture.CallFunction(t, fnName, &s)
	if err != nil {
		t.Fatal(err)
	}

	expected := "echo: hello"
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestHostFn_encodeStrings1(t *testing.T) {
	fnName := "encode_strings1"
	s := []string{"hello", "world"}

	result, err := fixture.CallFunction(t, fnName, s)
	if err != nil {
		t.Fatal(err)
	}

	expected := `["hello","world"]`
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}

func TestHostFn_encodeStrings2(t *testing.T) {
	fnName := "encode_strings2"
	e0 := "hello"
	e1 := "world"
	s := []*string{&e0, &e1}

	result, err := fixture.CallFunction(t, fnName, s)
	if err != nil {
		t.Fatal(err)
	}

	expected := `["hello","world"]`
	if result == nil {
		t.Error("expected a result")
	} else if r, ok := result.(*string); !ok {
		t.Errorf("expected a string, got %T", result)
	} else if *r != expected {
		t.Errorf("expected %v, got %v", expected, *r)
	}
}
