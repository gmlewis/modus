/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package plugins

import (
	"context"
	"fmt"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/languages"
	"github.com/gmlewis/modus/runtime/logger"
	"github.com/gmlewis/modus/runtime/utils"

	"github.com/tetratelabs/wazero"
	wasm "github.com/tetratelabs/wazero/api"
)

type Plugin struct {
	Id             string
	Module         wazero.CompiledModule
	Metadata       *metadata.Metadata
	FileName       string
	Language       langsupport.Language
	ExecutionPlans map[string]langsupport.ExecutionPlan
}

func NewPlugin(ctx context.Context, cm wazero.CompiledModule, filename string, md *metadata.Metadata) (*Plugin, error) {
	span, ctx := utils.NewSentrySpanForCurrentFunc(ctx)
	defer span.Finish()

	language, err := languages.GetLanguageForSDK(md.SDK)
	if err != nil {
		return nil, err
	}

	planner := language.NewPlanner(md)
	imports := cm.ImportedFunctions()
	exports := cm.ExportedFunctions()
	plans := make(map[string]langsupport.ExecutionPlan, len(imports)+len(exports))

	ctx = context.WithValue(ctx, utils.MetadataContextKey, md)

	for fnName, fnMeta := range md.FnExports {
		fnDef, ok := exports[fnName]
		if !ok {
			return nil, fmt.Errorf("no wasm function definition found for %s", fnName)
		}

		plan, err := planner.GetPlan(ctx, fnMeta, fnDef)
		if err != nil {
			return nil, fmt.Errorf("failed to get execution plan for %v: %w", fnName, err)
		}
		plans[fnName] = plan
	}

	importsMap := make(map[string]wasm.FunctionDefinition, len(imports))
	for _, fnDef := range imports {
		if modName, fnName, ok := fnDef.Import(); ok {
			importName := modName + "." + fnName
			importsMap[importName] = fnDef
		}
	}

	for importName, fnMeta := range md.FnImports {
		fnDef, ok := importsMap[importName]
		if !ok {
			logger.Warn(ctx).Msgf("Unused import %s in plugin metadata. Please update your Modus SDK.", importName)
			continue
		}

		plan, err := planner.GetPlan(ctx, fnMeta, fnDef)
		if err != nil {
			return nil, fmt.Errorf("failed to get execution plan for %s: %w", importName, err)
		}
		plans[importName] = plan
	}

	plugin := &Plugin{
		Id:             utils.GenerateUUIDv7(),
		Module:         cm,
		Metadata:       md,
		FileName:       filename,
		Language:       language,
		ExecutionPlans: plans,
	}

	return plugin, nil
}

func (p *Plugin) NameAndVersion() (name string, version string) {
	return p.Metadata.NameAndVersion()
}

func (p *Plugin) Name() string {
	return p.Metadata.Name()
}

func (p *Plugin) Version() string {
	return p.Metadata.Version()
}

func (p *Plugin) BuildId() string {
	return p.Metadata.BuildId
}

func GetPluginFromContext(ctx context.Context) (*Plugin, bool) {
	p, ok := ctx.Value(utils.PluginContextKey).(*Plugin)
	return p, ok
}

// type customFnDef struct {
// 	moduleName  string
// 	debugName   string
// 	name        string
// 	exportNames []string
// 	paramNames  []string
// 	paramTypes  []wasm.ValueType
// 	resultNames []string
// 	resultTypes []wasm.ValueType
// 	fn          any
// }

// func (c *customFnDef) ModuleName() string             { return c.moduleName }
// func (c *customFnDef) Name() string                   { return c.name }
// func (c *customFnDef) DebugName() string              { return c.debugName }
// func (c *customFnDef) ExportNames() []string          { return c.exportNames }
// func (c *customFnDef) ParamNames() []string           { return c.paramNames }
// func (c *customFnDef) ParamTypes() []wasm.ValueType   { return c.paramTypes }
// func (c *customFnDef) ResultNames() []string          { return c.resultNames }
// func (c *customFnDef) ResultTypes() []wasm.ValueType  { return c.resultTypes }
// func (c *customFnDef) GoFunction() any                { return c.fn }
// func (c *customFnDef) Import() (string, string, bool) { return c.moduleName, c.name, true }
// func (c *customFnDef) Index() uint32                  { return 0 }
