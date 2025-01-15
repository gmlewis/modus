/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

// The models package provides a generic interface for invoking various AI models.
package models

import (
	"fmt"

	"github.com/gmlewis/modus/sdk/go/pkg/console"
	"github.com/gmlewis/modus/sdk/go/pkg/utils"
)

// Provides information about a model.
type ModelInfo struct {

	// The name of the model, as specified in the modus.json manifest file.
	Name string

	// The full name of the model that the provider uses to identify the model.
	FullName string
}

// Provides a generic interface for interacting with an AI model.
type Model[TIn, TOut any] interface {
	Info() *ModelInfo
	Invoke(input *TIn) (*TOut, error)
}

// Provides a generic interface for setting the model information (used internally).
type modelPtr[TModel any] interface {
	*TModel
	setInfo(info *ModelInfo)
}

// Provides a base implementation for all models.
type ModelBase[TIn, TOut any] struct {
	info  *ModelInfo
	Debug bool
}

// Gets the model information.
func (m ModelBase[TIn, TOut]) Info() *ModelInfo {
	return m.info
}

// Sets the model information (used internally).
func (m *ModelBase[TIn, TOut]) setInfo(info *ModelInfo) {
	m.info = info
}

// Gets a model object instance, which can be used to interact with the model.
//   - The generic type parameter [TModel] is used to specify the API that the model adheres to.
//   - The name parameter is used to identify the model in the modus.json manifest file.
//
// Note that of the generic type parameters, only [TModel] needs to be specified.  The others are inferred automatically.
func GetModel[TModel Model[TIn, TOut], TIn, TOut any, P modelPtr[TModel]](name string) (*TModel, error) {
	info := hostGetModelInfo(&name)
	if info == nil {
		return nil, fmt.Errorf("model not found: %s", name)
	}

	model := P(new(TModel))
	model.setInfo(info)

	return model, nil
}

// Invokes the model with the specified input and returns the output generated by the model.
func (m ModelBase[TIn, TOut]) Invoke(input *TIn) (*TOut, error) {
	if m.info == nil {
		return nil, fmt.Errorf("model info is not set (use GetModel to create a model instance)")
	}

	modelName := m.info.Name
	inputJson, err := utils.JsonSerialize(input)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize model input for %s: %w", modelName, err)
	}

	if m.Debug {
		console.Logf("Invoking model %s with input: %s", modelName, inputJson)
	}

	sInputJson := string(inputJson)
	sOutputJson := hostInvokeModel(&modelName, &sInputJson)
	if sOutputJson == nil {
		return nil, fmt.Errorf("failed to invoke model %s", modelName)
	}

	if m.Debug {
		console.Logf("Received output for model %s: %s", modelName, *sOutputJson)
	}

	var result TOut
	err = utils.JsonDeserialize([]byte(*sOutputJson), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize model output for %s: %w", modelName, err)
	}

	return &result, nil
}
