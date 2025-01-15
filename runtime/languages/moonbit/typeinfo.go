/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package moonbit

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/utils"
)

var _langTypeInfo = &langTypeInfo{}

func LanguageTypeInfo() langsupport.LanguageTypeInfo {
	return _langTypeInfo
}

func GetTypeInfo(ctx context.Context, typeName string, typeCache map[string]langsupport.TypeInfo) (langsupport.TypeInfo, error) {
	return langsupport.GetTypeInfo(ctx, _langTypeInfo, typeName, typeCache)
}

type langTypeInfo struct{}

func (lti *langTypeInfo) GetListSubtype(typ string) string {
	return typ[strings.Index(typ, "]")+1:] // TODO(gmlewis)
}

func (lti *langTypeInfo) GetMapSubtypes(typ string) (string, string) {
	return "", "" // TODO(gmlewis)
}

func (lti *langTypeInfo) GetNameForType(typ string) string {
	return "" // TODO(gmlewis)
}

func (lti *langTypeInfo) IsObjectType(typ string) bool {
	return !lti.IsPrimitiveType(typ) &&
		!lti.IsListType(typ) &&
		!lti.IsMapType(typ) &&
		!lti.IsStringType(typ) &&
		!lti.IsTimestampType(typ) &&
		!lti.IsPointerType(typ)
}

func (lti *langTypeInfo) GetUnderlyingType(typ string) string {
	return strings.TrimPrefix(typ, "*") // TODO(gmlewis)
}

func (lti *langTypeInfo) IsListType(typ string) bool {
	return len(typ) > 2 && typ[0] == '[' // TODO(gmlewis)
}

func (lti *langTypeInfo) IsSliceType(typ string) bool {
	return len(typ) > 2 && typ[0] == '[' && typ[1] == ']' // TODO(gmlewis)
}

func (lti *langTypeInfo) IsArrayType(typ string) bool {
	return len(typ) > 2 && typ[0] == '[' && typ[1] != ']' // TODO(gmlewis)
}

func (lti *langTypeInfo) IsBooleanType(typ string) bool {
	return typ == "Bool"
}

func (lti *langTypeInfo) IsByteSequenceType(typ string) bool {
	switch typ {
	case "Array[Byte]", "Bytes", "ArrayView[Byte]", "BytesView":
		return true
	}

	if lti.IsArrayType(typ) {
		switch lti.GetListSubtype(typ) {
		case "Byte":
			return true
		}
	}

	return false
}

func (lti *langTypeInfo) IsFloatType(typ string) bool {
	switch typ {
	case "Float", "Double":
		return true
	default:
		return false
	}
}

func (lti *langTypeInfo) IsIntegerType(typ string) bool {
	switch typ {
	case "Int", "Int64",
		"Uint", "Uint64",
		"Byte", "Char": // TODO(gmlewis)
		return true
	default:
		return false
	}
}

func (lti *langTypeInfo) IsMapType(typ string) bool {
	return strings.HasPrefix(typ, "Map[")
}

func (lti *langTypeInfo) IsNullableType(typ string) bool {
	return lti.IsPointerType(typ) || lti.IsSliceType(typ) || lti.IsMapType(typ) // TODO(gmlewis)
}

func (lti *langTypeInfo) IsPointerType(typ string) bool {
	return strings.HasPrefix(typ, "*") // TODO(gmlewis)
}

func (lti *langTypeInfo) IsPrimitiveType(typ string) bool {
	return lti.IsBooleanType(typ) || lti.IsIntegerType(typ) || lti.IsFloatType(typ) // TODO(gmlewis)
}

func (lti *langTypeInfo) IsSignedIntegerType(typ string) bool {
	switch typ {
	case "Int", "Int64": // TODO(gmlewis)
		return true
	default:
		return false
	}
}

func (lti *langTypeInfo) IsStringType(typ string) bool {
	return typ == "String"
}

func (lti *langTypeInfo) IsTimestampType(typ string) bool {
	return typ == "time.Time" // TODO(gmlewis)
}

func (lti *langTypeInfo) ArrayLength(typ string) (int, error) {
	return 0, errors.New("langTypeInfo.ArrayLength not implemented yet for MoonBit")
}

func (lti *langTypeInfo) GetAlignmentOfType(ctx context.Context, typ string) (uint32, error) {
	return 0, errors.New("langTypeInfo.GetAlignmentOfType not implemented yet for MoonBit")
}

func (lti *langTypeInfo) ObjectsUseMaxFieldAlignment() bool {
	return true // TODO(gmlewis)
}
func (lti *langTypeInfo) GetDataSizeOfType(ctx context.Context, typ string) (uint32, error) {
	return lti.GetSizeOfType(ctx, typ)
}

func (lti *langTypeInfo) GetEncodingLengthOfType(ctx context.Context, typ string) (uint32, error) {
	return 0, errors.New("langTypeInfo.GetEncodingLengthOfType not implemented yet for MoonBit")
}

func (lti *langTypeInfo) GetSizeOfType(ctx context.Context, typ string) (uint32, error) {
	return 0, errors.New("langTypeInfo.GetSizeOfType not implemented yet for MoonBit")
}

func (lti *langTypeInfo) GetReflectedType(ctx context.Context, typ string) (reflect.Type, error) {
	if customTypes, ok := ctx.Value(utils.CustomTypesContextKey).(map[string]reflect.Type); ok {
		return lti.getReflectedType(typ, customTypes)
	} else {
		return lti.getReflectedType(typ, nil)
	}
}

func (lti *langTypeInfo) getReflectedType(typ string, customTypes map[string]reflect.Type) (reflect.Type, error) {
	return nil, errors.New("langTypeInfo.getReflectedType not implemented yet for MoonBit")
}
