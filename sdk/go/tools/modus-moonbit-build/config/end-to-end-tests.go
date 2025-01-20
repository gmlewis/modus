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
	"fmt"
	"os"
	"strings"
)

type EndToEndTests struct {
	Plugins []*Plugin `json:"plugins"`
}

type Arg struct {
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
}

func (a *Arg) String() string {
	val, _ := a.Value.MarshalJSON()
	return fmt.Sprintf("%q:%s", a.Name, val)
}

type Endpoint struct {
	Name   string          `json:"name"`
	Args   []*Arg          `json:"args,omitempty"`
	Expect json.RawMessage `json:"expect"`
}

func (e *Endpoint) QueryBody() string {
	if len(e.Args) == 0 {
		return fmt.Sprintf(`{"query": "query {\n  %v\n}", "variables": {}}`, e.Name)
	}
	return fmt.Sprintf(`{"query": "query {\n  %v\n}", "variables": {%v}}`, e.Name, e.Params())
}

func (e *Endpoint) Params() string {
	var args []string
	for _, arg := range e.Args {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("%v", strings.Join(args, ","))
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
