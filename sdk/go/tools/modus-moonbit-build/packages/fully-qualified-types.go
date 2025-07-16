/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package packages

import (
	"go/types"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/utils"
)

func fullyQualifiedName(typesPkg *types.Package, name string) string {
	if typesPkg.Path() != "" && !utils.IsWellKnownType(name) {
		return typesPkg.Path() + "." + name
	}
	return name
}

func fullyQualifiedNewTypeName(typesPkg *types.Package, name string, underlying types.Type) (string, *types.TypeName) {
	if typesPkg.Path() != "" && !utils.IsWellKnownType(name) {
		newName := typesPkg.Path() + "." + name
		return newName, types.NewTypeName(0, typesPkg, name, underlying)
	}
	return name, types.NewTypeName(0, nil, name, underlying)
}
