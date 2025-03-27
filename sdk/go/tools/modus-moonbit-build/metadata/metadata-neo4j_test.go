// -*- compile-command: "go test -run ^TestFunction_String_Neo4j$ ."; -*-

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

//go:embed testdata/neo4j-example-metadata.json
var neo4jExampleMetadataJSON []byte

func TestFunction_String_Neo4j(t *testing.T) {
	t.Parallel()

	tests := []functionStringTest{
		{name: "create_people_and_relationships", want: "() -> String!Error"},
		{name: "delete_all_nodes", want: "() -> String!Error"},
		{name: "get_alice_friends_under_40", want: "() -> Array[Person]!Error"},
		{name: "get_alice_friends_under_40_ages", want: "() -> Array[Int]!Error"},
	}

	testFunctionStringHelper(t, "neo4j", neo4jExampleMetadataJSON, tests)
}
