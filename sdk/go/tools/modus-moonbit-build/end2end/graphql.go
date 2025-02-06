/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package end2end

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/google/go-cmp/cmp"
)

type GraphQLResponse struct {
	Data   map[string]any  `json:"data"`
	Errors []*config.Error `json:"errors"`
}

func testEndpoint(ctx context.Context, endpoint *config.Endpoint) error {
	if endpoint == nil || endpoint.Name == "" || endpoint.Query == nil {
		buf, _ := json.Marshal(endpoint)
		return fmt.Errorf("testEndpoint: invalid endpoint specification (missing 'name' or 'query'): %s", buf)
	}

	query, err := endpoint.QueryBody()
	if err != nil {
		return fmt.Errorf("endpoint.QueryBody: %w", err)
	}
	log.Printf("\n\n*** Testing endpoint with query body: '%v'", query)
	expect, err := endpoint.ExpectBody()
	if err != nil {
		return fmt.Errorf("endpoint.ExpectBody: %w", err)
	}
	regexpMatch, regexpStr, err := endpoint.RegexpBody()
	if err != nil {
		return fmt.Errorf("endpoint.RegexpBody: %w", err)
	}
	errors := endpoint.Errors

	// Create a new HTTP POST request with the context
	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8686/graphql", bytes.NewBufferString(query))
	if err != nil {
		return fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()

	// Read and print the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("io.ReadAll: %w", err)
	}

	fmt.Printf("\n\n***Response: %s\n", respBody)
	var res GraphQLResponse
	if err := json.Unmarshal(respBody, &res); err != nil {
		return fmt.Errorf("json.Unmarshal: %w\n%s", err, respBody)
	}
	got, err := json.Marshal(res.Data[endpoint.Name])
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	// Check that the response `data` matches the expected value or regexp:
	if regexpStr != "" {
		if !regexpMatch.MatchString(string(got)) {
			return fmt.Errorf("test: FAIL: Response from '%v' endpoint =\n%s\nfailed regexp:\n%s", endpoint.Name, got, regexpStr)
		}
	} else if diff := cmp.Diff(expect, got); diff != "" {
		return fmt.Errorf("test: FAIL: Response from '%v' endpoint %T=\n%s\nwant: %T\n%s\ndiff:\n%v", endpoint.Name, got, got, expect, expect, diff)
	}

	// Check that the errors match:
	if diff := cmp.Diff(errors, res.Errors); diff != "" {
		return fmt.Errorf("test: FAIL: Errors from '%v' endpoint mismatch (-want +got):\n%v", endpoint.Name, diff)
	}

	log.Printf("Test: OK passed for endpoint '%v'", endpoint.Name)

	return nil
}
