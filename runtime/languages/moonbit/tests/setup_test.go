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
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/gmlewis/modus/runtime/httpclient"
	"github.com/gmlewis/modus/runtime/testutils"
)

var basePath = func() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}()

var fixture *testutils.WasmTestFixture

func TestMain(m *testing.M) {
	log.SetFlags(0)
	path := filepath.Join(basePath, "..", "testdata", "build", "testdata.wasm")

	customTypes := make(map[string]reflect.Type)
	customTypes["TestStruct1"] = reflect.TypeFor[TestStruct1]()
	customTypes["TestStruct2"] = reflect.TypeFor[TestStruct2]()
	customTypes["TestStruct3"] = reflect.TypeFor[TestStruct3]()
	customTypes["TestStruct4"] = reflect.TypeFor[TestStruct4]()
	customTypes["TestStruct5"] = reflect.TypeFor[TestStruct5]()
	customTypes["TestStructWithMap"] = reflect.TypeFor[TestStructWithMap1]()
	customTypes["TestSmorgasbordStruct"] = reflect.TypeFor[TestSmorgasbordStruct]()
	customTypes["HttpResponse"] = reflect.TypeFor[httpclient.HttpResponse]()
	customTypes["HttpHeaders"] = reflect.TypeFor[httpclient.HttpHeaders]()
	customTypes["HttpHeader"] = reflect.TypeFor[httpclient.HttpHeader]()

	registrations := getTestHostFunctionRegistrations()
	fixture = testutils.NewWasmTestFixture(path, customTypes, registrations)

	exitVal := m.Run()

	fixture.Close()
	os.Exit(exitVal)
}
