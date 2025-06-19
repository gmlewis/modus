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
	"fmt"
	"strings"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"

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
	name := h.TypeInfo().Name()
	p.typeHandlers[name] = h
}

func (p *planner) AllHandlers() map[string]langsupport.TypeHandler {
	return p.typeHandlers
}

func NewTypeHandler(ti langsupport.TypeInfo) *typeHandler {
	return &typeHandler{
		typeInfo: ti,
	}
}

type typeHandler struct {
	typeInfo langsupport.TypeInfo
}

func (h *typeHandler) TypeInfo() langsupport.TypeInfo {
	return h.typeInfo
}

func (p *planner) GetHandler(ctx context.Context, typeName string) (langsupport.TypeHandler, error) {
	// Strip MoonBit error type suffix.
	if i := strings.Index(typeName, "!"); i >= 0 {
		typeName = typeName[:i]
	} else if i := strings.Index(typeName, " raise "); i >= 0 { // as of 'moonc v0.6.18+8382ed77e'
		typeName = typeName[:i]
	}

	if handler, ok := p.typeHandlers[typeName]; ok {
		return handler, nil
	}

	ti, err := GetTypeInfo(ctx, typeName, p.typeCache)
	if err != nil {
		if strings.Contains(err.Error(), "info for type") {
			return nil, fmt.Errorf("failed to get type info for %v: %w - did you remember to export ('pub') the struct?", typeName, err)
		}
		return nil, fmt.Errorf("failed to get type info for %v: %w", typeName, err)
	}

	if ti.IsPrimitive() {
		return p.NewPrimitiveHandler(ti)
	} else if ti.IsString() {
		return p.NewStringHandler(ti)
	} else if ti.IsPointer() {
		return p.NewPointerHandler(ctx, ti)
	} else if ti.IsList() {
		if _langTypeInfo.IsSliceType(typeName) { // A MoonBit `Array[]` is a Go slice type.
			elemType := ti.ListElementType()
			if !elemType.IsNullable() && elemType.IsPrimitive() {
				return p.NewPrimitiveSliceHandler(ti)
			} else {
				return p.NewSliceHandler(ctx, ti)
			}
			// MoonBit has _NO_ concept of a Go (fixed-length) "array" type.
			// Even a `FixedArray` in MoonBit is similar to a slice in Go because
			// its length is not encoded as part of its type.
		}
	} else if ti.IsMap() {
		return p.NewMapHandler(ctx, ti)
	} else if ti.IsTimestamp() {
		return p.NewTimeHandler(ti)
	} else if ti.IsObject() {
		if strings.HasPrefix(typeName, "@time.Duration") {
			return p.NewDurationHandler(ti)
		}
		return p.NewStructHandler(ctx, ti)
	}

	return nil, fmt.Errorf("can't determine plan for type: %s", typeName)
}

func (p *planner) GetPlan(ctx context.Context, fnMeta *metadata.Function, fnDef wasm.FunctionDefinition) (langsupport.ExecutionPlan, error) {
	span, ctx := utils.NewSentrySpanForCurrentFunc(ctx)
	defer span.Finish()

	paramHandlers := make([]langsupport.TypeHandler, len(fnMeta.Parameters))
	for i, param := range fnMeta.Parameters {
		handler, err := p.GetHandler(ctx, param.Type)
		if err != nil {
			return nil, err
		}
		paramHandlers[i] = handler
	}

	resultHandlers := make([]langsupport.TypeHandler, 0, len(fnMeta.Results)+1)
	var errorType string
	for _, result := range fnMeta.Results {
		typeName := result.Type
		if i := strings.Index(typeName, "!"); i >= 0 {
			errorType = typeName[i+1:]
			typeName = typeName[:i]
		} else if i := strings.Index(typeName, " raise "); i >= 0 { // as of 'moonc v0.6.18+8382ed77e'
			errorType = typeName[i+len(" raise "):]
			typeName = typeName[:i]
		}

		handler, err := p.GetHandler(ctx, typeName)
		if err != nil {
			return nil, err
		}
		resultHandlers = append(resultHandlers, handler)
	}
	switch errorType {
	case "": // no-op
	case "Error": // Treat as a struct with a single String field.
		errorHandler, err := p.GetHandler(ctx, "(String)")
		if err != nil {
			return nil, fmt.Errorf("failed to get handler for '(String)': %w", err)
		}
		resultHandlers = append(resultHandlers, errorHandler)
	default:
		return nil, fmt.Errorf("unsupported error type: !%v", errorType)
	}

	indirectResultSize, err := p.getIndirectResultSize(ctx, fnMeta, fnDef)
	if err != nil {
		return nil, err
	}

	plan := langsupport.NewExecutionPlan(fnDef, fnMeta, paramHandlers, resultHandlers, indirectResultSize)
	return plan, nil
}

func (p *planner) getIndirectResultSize(ctx context.Context, fnMeta *metadata.Function, fnDef wasm.FunctionDefinition) (uint32, error) {
	// If no results are expected, then we don't need to use indirection.
	if len(fnMeta.Results) == 0 {
		return 0, nil
	}

	// If the function definition has results, then we don't need to use indirection.
	if len(fnDef.ResultTypes()) > 0 {
		return 0, nil
	}

	totalSize := uint32(0)
	for _, r := range fnMeta.Results {
		size, err := _langTypeInfo.GetSizeOfType(ctx, r.Type)
		if err != nil {
			return 0, err
		}
		totalSize += size
	}
	return totalSize, nil
}
