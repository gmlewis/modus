/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package httpclient

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gmlewis/modus/runtime/secrets"
	"github.com/gmlewis/modus/runtime/utils"
)

func Fetch(ctx context.Context, request *HttpRequest) (*HttpResponse, error) {
	host, err := GetHttpConnectionForUrl(request.Url)
	if err != nil {
		return nil, err
	}

	log.Printf("Fetch: %v %v\n%s", request.Method, request.Url, request.Body)

	body := bytes.NewBuffer(request.Body)
	req, err := http.NewRequestWithContext(ctx, request.Method, request.Url, body)
	if err != nil {
		return nil, err
	}

	if request.Headers != nil {
		for _, header := range request.Headers.Data {
			req.Header[header.Name] = header.Values
		}
	}

	if err := secrets.ApplySecretsToHttpRequest(ctx, host, req); err != nil {
		return nil, err
	}

	resp, err := utils.HttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Don't check status code here, just pass it back to the caller.

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]*HttpHeader, len(resp.Header))
	for name, values := range resp.Header {
		header := &HttpHeader{
			Name:   name,
			Values: values,
		}
		headers[strings.ToLower(name)] = header
	}

	response := &HttpResponse{
		Status:     uint16(resp.StatusCode),
		StatusText: resp.Status[4:], // Remove the status code from the status text.
		Headers:    &HttpHeaders{Data: headers},
		Body:       content,
	}

	log.Printf("Response: %v %v\n%s", response.Status, response.StatusText, response.Body)

	return response, nil
}
