// -*- compile-command: "go test -run ^TestFunction_String_Time$ ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package metadata

import (
	_ "embed"
	"testing"
)

//go:embed testdata/time-example-metadata.json
var timeExampleMetadataJSON []byte

func TestFunction_String_Time(t *testing.T) {
	t.Parallel()

	tests := []functionStringTest{
		{name: "get_local_time_modus", want: "() -> String"},
		{name: "get_local_time_moonbit", want: "() -> String!Error"},
		{name: "get_local_time_zone_id", want: "() -> String"},
		{name: "get_time_in_zone_modus", want: "(tz : String) -> String"},
		{name: "get_time_in_zone_moonbit", want: "(tz : String) -> String!Error"},
		{name: "get_time_zone_info", want: "(tz : String) -> TimeZoneInfo!Error"},
		{name: "get_utc_time", want: "() -> @time.ZonedDateTime!Error"},
	}

	testFunctionStringHelper(t, "time", timeExampleMetadataJSON, tests)
}
