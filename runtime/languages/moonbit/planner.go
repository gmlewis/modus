/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// Package moonbit adds modus runtime support for the MoonBit programming language.
package moonbit

import (
	"context"
	"errors"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"

	wasm "github.com/tetratelabs/wazero/api"
)

func NewPlanner(metadata *metadata.Metadata) langsupport.Planner {
	return &planner{
		typeCache:    make(map[string]langsupport.TypeInfo),
		typeHandlers: make(map[string]langsupport.TypeHandler),
		metadata:     metadata,
	}
}

type planner struct {
	typeCache    map[string]langsupport.TypeInfo
	typeHandlers map[string]langsupport.TypeHandler
	metadata     *metadata.Metadata
}

func (p *planner) AddHandler(h langsupport.TypeHandler) {
	p.typeHandlers[h.TypeInfo().Name()] = h
}

func (p *planner) AllHandlers() map[string]langsupport.TypeHandler {
	return p.typeHandlers
}

func (p *planner) GetHandler(ctx context.Context, typeName string) (langsupport.TypeHandler, error) {
	return nil, errors.New("planner.GetHandler not implemented yet for MoonBit")
}

func (p *planner) GetPlan(ctx context.Context, fnMeta *metadata.Function, fnDef wasm.FunctionDefinition) (langsupport.ExecutionPlan, error) {
	return nil, errors.New("planner.GetPlan not implemented yet for MoonBit")
}
