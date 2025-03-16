// -*- compile-command: "go test -run ^TestTestablePreProcess_Time ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package codegen

import (
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
)

func TestTestablePreProcess_Time(t *testing.T) {
	t.Parallel()

	config := &config.Config{
		SourceDir: "../testdata/time-example",
	}

	mod := preProcessTestSetup(t, config)

	body, header, moonPkgJSON, err := testablePreProcess(config, mod)
	if err != nil {
		t.Fatal(err)
	}

	wg := &preProcessDiffs{
		wantPreProcessBody:        wantTimePreProcessBody,
		gotPreProcessBody:         body.String(),
		wantPreProcessHeader:      wantTimePreProcessHeader,
		gotPreProcessHeader:       header.String(),
		wantPreProcessMoonPkgJSON: wantTimePreProcessMoonPkgJSON,
		gotPreProcessMoonPkgJSON:  moonPkgJSON.String(),
	}
	reportPreProcessDiffs(t, "time", wg)
}

var wantTimePreProcessBody = `pub fn __modus_get_utc_time() -> @time.ZonedDateTime!Error {
  try get_utc_time!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_local_time_moonbit() -> String!Error {
  try get_local_time_moonbit!() {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_local_time_modus() -> String {
  get_local_time_modus()
}

pub fn __modus_get_time_in_zone_moonbit(tz : String) -> String!Error {
  try get_time_in_zone_moonbit!(tz) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

pub fn __modus_get_time_in_zone_modus(tz : String) -> String {
  get_time_in_zone_modus(tz)
}

pub fn __modus_get_local_time_zone_id() -> String {
  get_local_time_zone_id()
}

pub fn __modus_get_time_zone_info(tz : String) -> TimeZoneInfo!Error {
  try get_time_zone_info!(tz) {
    e => {
      @console.error(e.to_string())
      raise e
    }
  }
}

`

var wantTimePreProcessHeader = `// Code generated by modus-moonbit-build. DO NOT EDIT.

`

var wantTimePreProcessMoonPkgJSON = `{
  "import": [
    "gmlewis/modus/pkg/console",
    "gmlewis/modus/pkg/localtime",
    "gmlewis/modus/wit/interface/wasi",
    "moonbitlang/x/time"
  ],
  "targets": {
    "modus_post_generated.mbt": [
      "wasm"
    ]
  },
  "link": {
    "wasm": {
      "exports": [
        "__modus_get_local_time_modus:get_local_time_modus",
        "__modus_get_local_time_moonbit:get_local_time_moonbit",
        "__modus_get_local_time_zone_id:get_local_time_zone_id",
        "__modus_get_time_in_zone_modus:get_time_in_zone_modus",
        "__modus_get_time_in_zone_moonbit:get_time_in_zone_moonbit",
        "__modus_get_time_zone_info:get_time_zone_info",
        "__modus_get_utc_time:get_utc_time",
        "cabi_realloc",
        "copy",
        "duration_from_nanos",
        "free",
        "load32",
        "malloc",
        "ptr2str",
        "ptr_to_none",
        "read_map",
        "store32",
        "store8",
        "write_map",
        "zoned_date_time_from_unix_seconds_and_nanos"
      ],
      "export-memory-name": "memory"
    }
  }
}`
