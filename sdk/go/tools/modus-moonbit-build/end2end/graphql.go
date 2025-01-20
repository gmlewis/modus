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

func testEndpoint(ctx context.Context, endpoint *config.Endpoint) error {
	query := endpoint.QueryBody()
	log.Printf("\n\n*** Testing endpoint with query body: '%v'", query)

	// Create a new HTTP POST request with the context
	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8686/graphql", bytes.NewBufferString(query))
	if err != nil {
		return err
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read and print the response body
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("\n\n***Response: %s\n", got)
	expect, err := json.Marshal(endpoint.Expect)
	if err != nil {
		return err
	}

	if diff := cmp.Diff(expect, got); diff != "" {
		return fmt.Errorf("Response from '%v' endpoint =\n%s\nwant:\n%s", endpoint.Name, got, expect)
	}

	return nil
}
