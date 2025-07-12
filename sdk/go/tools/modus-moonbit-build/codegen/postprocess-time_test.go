// -*- compile-command: "NO_COLOR=1 go test -run ^TestTestablePostProcess_Time ."; -*-

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

func TestTestablePostProcess_Time(t *testing.T) {
	t.Parallel()

	config := &config.Config{
		SourceDir: "../testdata/time-example",
	}

	meta := postProcessTestSetup(t, config)

	body, header := testablePostProcess(meta)

	wg := &postProcessDiffs{
		wantPostProcessBody:   wantTimePostProcessBody,
		gotPostProcessBody:    body.String(),
		wantPostProcessHeader: postProcessHeader + wantTimePostProcessHeader,
		gotPostProcessHeader:  header.String(),
	}
	reportPostProcessDiffs(t, "time", wg)
}

var wantTimePostProcessBody = ``

var wantTimePostProcessHeader = `
///|
pub fn zoned_date_time_from_unix_seconds_and_nanos(second : Int64, nanos : Int64) -> @time.ZonedDateTime raise Error {
  let nanosecond = (nanos % 1_000_000_000).to_int()
  @time.unix!(second, nanosecond~)
}

///|
pub fn duration_from_nanos(nanoseconds : Int64) -> @time.Duration raise Error {
  @time.Duration::of!(nanoseconds~)
}
`
