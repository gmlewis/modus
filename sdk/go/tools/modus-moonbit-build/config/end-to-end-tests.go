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
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type EndToEndTests struct {
	Plugins []*Plugin `json:"plugins"`
}

type Endpoint struct {
	Name   string          `json:"name"`
	Query  json.RawMessage `json:"query"`
	Expect json.RawMessage `json:"expect"`
	Regexp json.RawMessage `json:"regexp"`
	Errors []*Error        `json:"errors"`
}

type Error struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}

func (e *Endpoint) QueryBody() (string, error) {
	s, err := json.Marshal(e.Query)
	return string(s), err
}

func (e *Endpoint) ExpectBody() (any, error) {
	s, err := json.Marshal(e.Expect)
	return s, err
}

func (e *Endpoint) RegexpBody() (*regexp.Regexp, string, error) {
	if len(e.Regexp) == 0 {
		return nil, "", nil
	}
	buf, err := json.Marshal(e.Regexp)
	if err != nil {
		return nil, "", err
	}
	rs := strings.Trim(string(buf), `"`)
	rs = strings.ReplaceAll(rs, `\\`, `\`)
	rs = strings.ReplaceAll(rs, `\"`, `"`)
	r, err := regexp.Compile(rs)
	if err != nil {
		return nil, "", err
	}
	return r, rs, nil
}

type Plugin struct {
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	Endpoints []*Endpoint `json:"endpoints"`
}

func (c *Config) loadEndToEndTests(filename string, pluginsFlag string) error {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	var tests EndToEndTests
	if err := json.Unmarshal(buf, &tests); err != nil {
		return err
	}

	pluginsToTest := map[string]bool{}
	for _, pluginName := range strings.Split(pluginsFlag, ",") {
		pluginName = strings.TrimSpace(pluginName)
		if pluginName == "" {
			continue
		}
		pluginsToTest[pluginName] = true
		// A common mistake is to copy-paste the endpoint name instead of the plugin name.
		pluginName = strings.ReplaceAll(pluginName, "_", "-")
		pluginsToTest[pluginName] = true
	}

	// Run a quick sanity check to make sure every test has a name
	// and a GraphQL query.
	end2endTests := make([]*Plugin, 0, len(tests.Plugins))
	for _, plugin := range tests.Plugins {
		if pluginsFlag != "" && !pluginsToTest[plugin.Name] {
			continue
		}
		end2endTests = append(end2endTests, plugin)
		if plugin.Name == "" {
			return fmt.Errorf("file %q: plugin name is missing", filename)
		}
		if plugin.Path == "" {
			return fmt.Errorf("file %q: plugin path is missing", filename)
		}
		for _, endpoint := range plugin.Endpoints {
			if endpoint.Name == "" {
				return fmt.Errorf("file %q, plugin %q: endpoint name is missing", filename, plugin.Name)
			}
			if endpoint.Query == nil {
				return fmt.Errorf("file %q, plugin %q: endpoint %q: query is missing", filename, plugin.Name, endpoint.Name)
			}
		}
	}

	if len(end2endTests) == 0 {
		return errors.New("No plugins to test")
	}

	c.EndToEndTests = end2endTests
	return nil
}
