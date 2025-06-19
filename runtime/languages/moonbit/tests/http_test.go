/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// Tests FAIL with moonc v0.6.18+8382ed77e

package moonbit_test

import (
	"testing"

	"github.com/gmlewis/modus/runtime/httpclient"
	"github.com/google/go-cmp/cmp"
)

func TestHttpResponse(t *testing.T) {
	fnName := "test_http_response"
	r := getTestHttpResponse()

	if _, err := fixture.CallFunction(t, fnName, r); err != nil {
		t.Error(err)
	}
}

func TestHttpResponseOutput(t *testing.T) {
	fnName := "test_http_response_output"

	got, err := fixture.CallFunction(t, fnName)
	if err != nil {
		t.Error(err)
	}

	want := getTestHttpResponse()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("%v: unexpected response (-want +got):\n%v", fnName, diff)
	}
}

func TestHttpResponseInput(t *testing.T) {
	fnName := "test_http_response_input"
	r := getTestHttpResponse()
	if _, err := fixture.CallFunction(t, fnName, r); err != nil {
		t.Error(err)
	}
}

func getTestHttpResponse() *httpclient.HttpResponse {
	return &httpclient.HttpResponse{
		Status:     200,
		StatusText: "OK",
		Headers: &httpclient.HttpHeaders{
			Data: map[string]*httpclient.HttpHeader{
				"content-type": {
					Name:   "Content-Type",
					Values: []string{"text/plain"},
				},
			},
		},
		Body: []byte("Hello, world!"),
	}
}

func TestHttpHeaders(t *testing.T) {
	fnName := "test_http_headers"
	h := &httpclient.HttpHeaders{
		Data: map[string]*httpclient.HttpHeader{
			"content-type": {
				Name:   "Content-Type",
				Values: []string{"text/plain"},
			},
		},
	}

	if _, err := fixture.CallFunction(t, fnName, h); err != nil {
		t.Error(err)
	}
}

func TestHttpHeaderMap(t *testing.T) {
	fnName := "test_http_header_map"
	m := map[string]*httpclient.HttpHeader{
		"content-type": {
			Name:   "Content-Type",
			Values: []string{"text/plain"},
		},
	}

	if _, err := fixture.CallFunction(t, fnName, m); err != nil {
		t.Error(err)
	}
}

func TestHttpHeader(t *testing.T) {
	fnName := "test_http_header"
	h := httpclient.HttpHeader{
		Name:   "Content-Type",
		Values: []string{"text/plain"},
	}

	if _, err := fixture.CallFunction(t, fnName, h); err != nil {
		t.Error(err)
	}
}
