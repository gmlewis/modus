/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package neo4jclient

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gmlewis/modus/runtime/manifestdata"
	"github.com/gmlewis/modus/runtime/utils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func Initialize() {
	manifestdata.RegisterManifestLoadedCallback(func(ctx context.Context) error {
		CloseDrivers(ctx)
		return nil
	})
}

func ExecuteQuery(ctx context.Context, hostName, dbName, query string, parametersJson string) (*EagerResult, error) {
	log.Printf("GML: neo4jclient.go: ExecuteQuery:\nhostName='%v'\ndbName='%v'\nquery='%v'\nparametersJson='%v'", hostName, dbName, query, parametersJson)
	driver, err := n4j.getDriver(ctx, hostName)
	if err != nil {
		return nil, err
	}

	parameters := make(map[string]any)
	if err := json.Unmarshal([]byte(parametersJson), &parameters); err != nil {
		return nil, err
	}

	res, err := neo4j.ExecuteQuery(ctx, driver, query, parameters, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(dbName))
	if err != nil {
		return nil, err
	}

	records := make([]*Record, len(res.Records))
	for i, record := range res.Records {
		vals := make([]string, len(record.Values))
		for j, val := range record.Values {
			valBytes, err := utils.JsonSerialize(val)
			if err != nil {
				return nil, err
			}
			vals[j] = string(valBytes)
		}
		log.Printf("GML: neo4jclient.go: ExecuteQuery:\nrecords[%v]={\n  Values: %+v\n  Keys: %+v\n}", i, vals, record.Keys)
		records[i] = &Record{
			Values: vals,
			Keys:   record.Keys,
		}
	}
	log.Printf("GML: neo4jclient.go: ExecuteQuery:\nKeys=%+v\n", res.Keys)
	return &EagerResult{
		Keys:    res.Keys,
		Records: records,
	}, nil
}
