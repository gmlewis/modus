/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"encoding/json"
	"os"
)

type EndToEndTests struct {
	Plugins []*Plugin `json:"plugins"`
}

type Arg struct {
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
}

type Endpoint struct {
	Name   string          `json:"name"`
	Args   []*Arg          `json:"args,omitempty"`
	Expect json.RawMessage `json:"expect"`
}

type Plugin struct {
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	Endpoints []*Endpoint `json:"endpoints"`
}

func (c *Config) loadEndToEndTests(filename string) error {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	var tests EndToEndTests
	if err := json.Unmarshal(buf, &tests); err != nil {
		return err
	}
	c.EndToEndTests = tests.Plugins
	return nil
}
