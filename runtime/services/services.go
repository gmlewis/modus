/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package services

import (
	"context"

	"github.com/gmlewis/modus/runtime/aws"
	"github.com/gmlewis/modus/runtime/collections"
	"github.com/gmlewis/modus/runtime/db"
	"github.com/gmlewis/modus/runtime/dgraphclient"
	"github.com/gmlewis/modus/runtime/envfiles"
	"github.com/gmlewis/modus/runtime/graphql"
	"github.com/gmlewis/modus/runtime/hostfunctions"
	"github.com/gmlewis/modus/runtime/logger"
	"github.com/gmlewis/modus/runtime/manifestdata"
	"github.com/gmlewis/modus/runtime/middleware"
	"github.com/gmlewis/modus/runtime/neo4jclient"
	"github.com/gmlewis/modus/runtime/pluginmanager"
	"github.com/gmlewis/modus/runtime/secrets"
	"github.com/gmlewis/modus/runtime/sqlclient"
	"github.com/gmlewis/modus/runtime/storage"
	"github.com/gmlewis/modus/runtime/utils"
	"github.com/gmlewis/modus/runtime/wasmhost"
)

// Starts any services that need to be started when the runtime starts.
func Start(ctx context.Context) context.Context {

	// Note, we cannot start a Sentry transaction here, or it will also be used for the background services, post-initiation.

	// Init the wasm host and put it in context
	registrations := hostfunctions.GetRegistrations()
	host := wasmhost.InitWasmHost(ctx, registrations...)
	ctx = context.WithValue(ctx, utils.WasmHostContextKey, host)

	// Start the background services.  None of these should block. If they need to do background work, they should start a goroutine internally.

	// NOTE: Initialization order is important in some cases.
	// If you need to change the order or add new services, be sure to test thoroughly.
	// Generally, new services should be added to the end of the list, unless there is a specific reason to do otherwise.

	sqlclient.Initialize()
	dgraphclient.Initialize()
	neo4jclient.Initialize()
	aws.Initialize(ctx)
	secrets.Initialize(ctx)
	storage.Initialize(ctx)
	db.Initialize(ctx)
	collections.Initialize(ctx)
	manifestdata.MonitorManifestFile(ctx)
	envfiles.MonitorEnvFiles(ctx)
	pluginmanager.Initialize(ctx)
	graphql.Initialize()

	return ctx
}

// Stops any services that need to be stopped when the runtime stops.
func Stop(ctx context.Context) {

	// Stop the wasm host first
	wasmhost.GetWasmHost(ctx).Close(ctx)

	// Stop the rest of the background services.
	// NOTE: Stopping services also has an order dependency.
	// If you need to change the order or add new services, be sure to test thoroughly.
	// Unlike start, these should each block until they are fully stopped.

	collections.Shutdown(ctx)
	middleware.Shutdown()
	sqlclient.ShutdownPGPools()
	dgraphclient.ShutdownConns()
	neo4jclient.CloseDrivers(ctx)
	logger.Close()
	db.Stop(ctx)
}
