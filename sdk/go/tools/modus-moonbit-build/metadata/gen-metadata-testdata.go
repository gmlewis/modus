//go:build ignore

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// gen-metadata-testdata.go is used to marshal testdata/simple-example-metadata.json
// from actual parsing of the MoonBit code so that all testing stays in sync.
// It is needed to prevent an import cycle, but is also useful to test the JSON serialization.
package main

func main() {
}
