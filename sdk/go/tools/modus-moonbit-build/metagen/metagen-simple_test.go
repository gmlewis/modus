// -*- compile-command: "go test -run ^TestGenerateMetadata_Simple$ ."; -*-

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

func TestGenerateMetadata_Simple(t *testing.T) {
	meta := setupTestConfig(t, "../testdata/simple-example")

	if got, want := meta.Plugin, "simple-example"; got != want {
		t.Errorf("meta.Plugin = %q, want %q", got, want)
	}

	if got, want := meta.Module, "@simple-example"; got != want {
		t.Errorf("meta.Module = %q, want %q", got, want)
	}

	if diff := cmp.Diff(wantSimpleFnExports, meta.FnExports); diff != "" {
		t.Errorf("meta.FnExports mismatch (-want +got):\n%v", diff)
	}

	if diff := cmp.Diff(wantSimpleFnImports, meta.FnImports); diff != "" {
		t.Errorf("meta.FnImports mismatch (-want +got):\n%v", diff)
	}

	diffMetaTypes(t, wantSimpleTypes, meta.Types)
}

var wantSimpleFnExports = metadata.FunctionMap{
	"add": {
		Name:       "add",
		Parameters: []*metadata.Parameter{{Name: "x", Type: "Int"}, {Name: "y", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs:       &metadata.Docs{Lines: []string{"Adds two integers together and returns the result."}},
	},
	"add3": {
		Name: "add3",
		Parameters: []*metadata.Parameter{
			{Name: "a", Type: "Int"},
			{Name: "b", Type: "Int"},
			{Name: "c~", Type: "Int", Default: AnyPtr(int32(0))},
		},
		Results: []*metadata.Result{{Type: "Int"}},
		Docs: &metadata.Docs{Lines: []string{
			"Adds three integers together and returns the result.",
			"The third integer is optional.",
		}},
	},
	"add3_WithDefaults": {
		Name:       "add3_WithDefaults",
		Parameters: []*metadata.Parameter{{Name: "a", Type: "Int"}, {Name: "b", Type: "Int"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs: &metadata.Docs{Lines: []string{
			"Adds three integers together and returns the result.",
			"The third integer is optional.",
		}},
	},
	"add_n": {
		Name:       "add_n",
		Parameters: []*metadata.Parameter{{Name: "args", Type: "Array[Int]"}},
		Results:    []*metadata.Result{{Type: "Int"}},
		Docs:       &metadata.Docs{Lines: []string{"Adds any number of integers together and returns the result."}},
	},
	"get_current_time": {
		Name:    "get_current_time",
		Results: []*metadata.Result{{Type: "@time.ZonedDateTime raise Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time."}},
	},
	"get_current_time_formatted": {
		Name:    "get_current_time_formatted",
		Results: []*metadata.Result{{Type: "String raise Error"}},
		Docs:    &metadata.Docs{Lines: []string{"Returns the current time formatted as a string."}},
	},
	"get_full_name": {
		Name:       "get_full_name",
		Parameters: []*metadata.Parameter{{Name: "first_name", Type: "String"}, {Name: "last_name", Type: "String"}},
		Results:    []*metadata.Result{{Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Combines the first and last name of a person, and returns the full name."}},
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
	"log_message": {
		Name:       "log_message",
		Parameters: []*metadata.Parameter{{Name: "message", Type: "String"}},
		Docs:       &metadata.Docs{Lines: []string{"Logs a message."}},
	},
	"test_abort": {
		Name: "test_abort",
		Docs: &metadata.Docs{Lines: []string{"Tests an abort."}},
	},
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
		Results:    []*metadata.Result{{Type: "String raise Error"}},
		Docs:       &metadata.Docs{Lines: []string{"Tests returning an error."}},
	},
}

var wantSimpleFnImports = metadata.FunctionMap{
	"modus_system.logMessage": {
		Name:       "modus_system.logMessage",
		Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
	},
}

var wantSimpleTypes = metadata.TypeMap{
	"(Int, Int, Int)": {
		Name:   "(Int, Int, Int)",
		Fields: []*metadata.Field{{Name: "0", Type: "Int"}, {Name: "1", Type: "Int"}, {Name: "2", Type: "Int"}},
	},
	"(String)": {
		Name:   "(String)",
		Fields: []*metadata.Field{{Name: "0", Type: "String"}},
	},
	"@ffi.XExternByteArray":   {Name: "@ffi.XExternByteArray"},
	"@ffi.XExternString":      {Name: "@ffi.XExternString"},
	"@ffi.XExternStringArray": {Name: "@ffi.XExternStringArray"},
	"@testutils.CallStack[T]": {
		Name:   "@testutils.CallStack[T]",
		Fields: []*metadata.Field{{Name: "items", Type: "Array[Array[@testutils.T]]"}},
	},
	"@testutils.T":                    {Name: "@testutils.T"},
	"@time.Duration":                  {Name: "@time.Duration"},
	"@time.Duration raise Error":      {Name: "@time.Duration raise Error"},
	"@time.Period":                    {Name: "@time.Period"},
	"@time.Period raise Error":        {Name: "@time.Period raise Error"},
	"@time.PlainDate":                 {Name: "@time.PlainDate"},
	"@time.PlainDate raise Error":     {Name: "@time.PlainDate raise Error"},
	"@time.PlainDateTime":             {Name: "@time.PlainDateTime"},
	"@time.PlainDateTime raise Error": {Name: "@time.PlainDateTime raise Error"},
	"@time.PlainTime":                 {Name: "@time.PlainTime"},
	"@time.PlainTime raise Error":     {Name: "@time.PlainTime raise Error"},
	"@time.Weekday":                   {Name: "@time.Weekday"},
	"@time.Zone":                      {Name: "@time.Zone"},
	"@time.Zone raise Error":          {Name: "@time.Zone raise Error"},
	"@time.ZoneOffset":                {Name: "@time.ZoneOffset"},
	"@time.ZoneOffset raise Error":    {Name: "@time.ZoneOffset raise Error"},
	"@time.ZonedDateTime":             {Name: "@time.ZonedDateTime"},
	"@time.ZonedDateTime raise Error": {Name: "@time.ZonedDateTime raise Error"},
	"ArrayView[Byte]":                 {Name: "ArrayView[Byte]"},
	"Array[@testutils.T]":             {Name: "Array[@testutils.T]"},
	"Array[Array[@testutils.T]]":      {Name: "Array[Array[@testutils.T]]"},
	"Array[Byte]":                     {Name: "Array[Byte]"},
	"Array[Int]":                      {Name: "Array[Int]"},
	"Array[Person]":                   {Name: "Array[Person]"},
	"Array[String]":                   {Name: "Array[String]"},
	"Bool":                            {Name: "Bool"},
	"Byte":                            {Name: "Byte"},
	"Bytes":                           {Name: "Bytes"},
	"Bytes raise Error":               {Name: "Bytes raise Error"},
	"Char":                            {Name: "Char"},
	"FixedArray[Byte]":                {Name: "FixedArray[Byte]"},
	"Int":                             {Name: "Int"},
	"Int64":                           {Name: "Int64"},
	"Iter[Byte]":                      {Name: "Iter[Byte]"},
	"Iter[Char]":                      {Name: "Iter[Char]"},
	"Map[String, String]":             {Name: "Map[String, String]"},
	"Person": {
		Name: "Person",
		Fields: []*metadata.Field{
			{Name: "firstName", Type: "String"}, {Name: "lastName", Type: "String"}, {Name: "age", Type: "Int"},
		},
	},
	"Result[UInt64, UInt]": {Name: "Result[UInt64, UInt]"},
	"Result[Unit, UInt]":   {Name: "Result[Unit, UInt]"},
	"String":               {Name: "String"},
	"String raise Error":   {Name: "String raise Error"},
	"UInt":                 {Name: "UInt"},
	"UInt64":               {Name: "UInt64"},
}
