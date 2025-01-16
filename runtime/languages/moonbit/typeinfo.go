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
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
	"unsafe"

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
	result := typ[strings.Index(typ, "]")+1:] // TODO(gmlewis)
	log.Printf("GML: typeinfo.go: GetListSubtype('%v') = '%v'", typ, result)
	return result
}

func (lti *langTypeInfo) GetMapSubtypes(typ string) (string, string) {
	log.Printf("GML: typeinfo.go: GetMapSubtypes('%v') = '','' - not implemented yet", typ)
	return "", "" // TODO(gmlewis)
}

func (lti *langTypeInfo) GetNameForType(typ string) string {
	log.Printf("GML: typeinfo.go: GetNameForType('%v') = '' - not implemented yet", typ)
	return "" // TODO(gmlewis)
}

func (lti *langTypeInfo) IsObjectType(typ string) bool {
	result := !lti.IsPrimitiveType(typ) &&
		!lti.IsListType(typ) &&
		!lti.IsMapType(typ) &&
		!lti.IsStringType(typ) &&
		!lti.IsTimestampType(typ) &&
		!lti.IsPointerType(typ)
	log.Printf("GML: typeinfo.go: IsObjectType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) GetUnderlyingType(typ string) string {
	result := strings.TrimPrefix(typ, "*") // TODO(gmlewis)
	log.Printf("GML: typeinfo.go: GetUnderlyingType('%v') = '%v' - TODO", typ, result)
	return result
}

func (lti *langTypeInfo) IsListType(typ string) bool {
	result := len(typ) > 2 && typ[0] == '[' // TODO(gmlewis)
	log.Printf("GML: typeinfo.go: IsListType('%v') = %v - TODO", typ, result)
	return result
}

func (lti *langTypeInfo) IsSliceType(typ string) bool {
	result := len(typ) > 2 && typ[0] == '[' && typ[1] == ']' // TODO(gmlewis)
	log.Printf("GML: typeinfo.go: IsSliceType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsArrayType(typ string) bool {
	result := strings.HasPrefix(typ, "Array[") && strings.HasSuffix(typ, "]")
	log.Printf("GML: typeinfo.go: IsArrayType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsBooleanType(typ string) bool {
	result := typ == "Bool"
	log.Printf("GML: typeinfo.go: IsBooleanType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsByteSequenceType(typ string) bool {
	switch typ {
	case "Array[Byte]", "Bytes", "ArrayView[Byte]", "BytesView":
		log.Printf("GML: typeinfo.go: IsByteSequenceType('%v') = true", typ)
		return true
	}

	if lti.IsArrayType(typ) {
		switch lti.GetListSubtype(typ) {
		case "Byte":
			log.Printf("GML: B: typeinfo.go: IsByteSequenceType('%v') = true", typ)
			return true
		}
	}

	log.Printf("GML: typeinfo.go: IsByteSequenceType('%v') = false", typ)
	return false
}

func (lti *langTypeInfo) IsFloatType(typ string) bool {
	switch typ {
	case "Float", "Double":
		log.Printf("GML: typeinfo.go: IsFloatType('%v') = true", typ)
		return true
	default:
		log.Printf("GML: typeinfo.go: IsFloatType('%v') = false", typ)
		return false
	}
}

func (lti *langTypeInfo) IsIntegerType(typ string) bool {
	switch typ {
	case "Int", "Int64",
		"Uint", "Uint64",
		"Byte", "Char": // TODO(gmlewis)
		log.Printf("GML: typeinfo.go: IsIntegerType('%v') = true - TODO", typ)
		return true
	default:
		log.Printf("GML: typeinfo.go: IsIntegerType('%v') = false - TODO", typ)
		return false
	}
}

func (lti *langTypeInfo) IsMapType(typ string) bool {
	result := strings.HasPrefix(typ, "Map[")
	log.Printf("GML: typeinfo.go: IsMapType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsNullableType(typ string) bool {
	result := lti.IsPointerType(typ) || lti.IsSliceType(typ) || lti.IsMapType(typ) // TODO(gmlewis)
	log.Printf("GML: typeinfo.go: IsNullableType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsPointerType(typ string) bool {
	result := strings.HasPrefix(typ, "*") // TODO(gmlewis)
	log.Printf("GML: typeinfo.go: IsPointerType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsPrimitiveType(typ string) bool {
	result := lti.IsBooleanType(typ) || lti.IsIntegerType(typ) || lti.IsFloatType(typ) // TODO(gmlewis)
	log.Printf("GML: typeinfo.go: IsPrimitiveType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsSignedIntegerType(typ string) bool {
	switch typ {
	case "Int", "Int64": // TODO(gmlewis)
		log.Printf("GML: typeinfo.go: IsSignedIntegerType('%v') = true", typ)
		return true
	default:
		log.Printf("GML: typeinfo.go: IsSignedIntegerType('%v') = false", typ)
		return false
	}
}

func (lti *langTypeInfo) IsStringType(typ string) bool {
	result := typ == "String"
	log.Printf("GML: typeinfo.go: IsStringType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsTimestampType(typ string) bool {
	// Special case for MoonBit wasi "timestamp"-like struct.
	// Could also have an initializer, like: '@wallClock.Datetime = @wallClock.now()'
	result := strings.HasPrefix(typ, "@wallClock.Datetime")
	log.Printf("GML: typeinfo.go: IsTimestampType('%v') = %v", typ, result)
	return result
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
		result, err := lti.getReflectedType(typ, customTypes)
		log.Printf("GML: typeinfo.go: A: GetReflectedType('%v') = %v", typ, result)
		return result, err
	} else {
		result, err := lti.getReflectedType(typ, nil)
		log.Printf("GML: typeinfo.go: B: GetReflectedType('%v') = %v", typ, result)
		return result, err
	}
}

func (lti *langTypeInfo) getReflectedType(typ string, customTypes map[string]reflect.Type) (reflect.Type, error) {
	// special case - '@wallClock.Datetime' is handled by `time.Time`
	if strings.HasPrefix(typ, "@wallClock.Datetime") {
		log.Printf("GML: typeinfo.go: Z: getReflectedType('%v') - translating to 'time.Time'", typ)
		typ = "time.Time"
	}

	if customTypes != nil {
		if rt, ok := customTypes[typ]; ok {
			log.Printf("GML: typeinfo.go: A: getReflectedType('%v') = %v", typ, rt)
			return rt, nil
		}
	}

	if rt, ok := reflectedTypeMap[typ]; ok {
		log.Printf("GML: typeinfo.go: B: getReflectedType('%v') = %v", typ, rt)
		return rt, nil
	}

	if lti.IsPointerType(typ) {
		tt := lti.GetUnderlyingType(typ)
		targetType, err := lti.getReflectedType(tt, customTypes)
		if err != nil {
			return nil, err
		}
		result := reflect.PointerTo(targetType)
		log.Printf("GML: typeinfo.go: C: getReflectedType('%v') = %v", typ, result)
		return result, nil
	}

	if lti.IsSliceType(typ) {
		et := lti.GetListSubtype(typ)
		if et == "" {
			return nil, fmt.Errorf("invalid slice type: %s", typ)
		}

		elementType, err := lti.getReflectedType(et, customTypes)
		if err != nil {
			return nil, err
		}
		result := reflect.SliceOf(elementType)
		log.Printf("GML: typeinfo.go: D: getReflectedType('%v') = %v", typ, result)
		return result, nil
	}

	if lti.IsArrayType(typ) {
		et := lti.GetListSubtype(typ)
		if et == "" {
			return nil, fmt.Errorf("invalid array type: %s", typ)
		}

		size, err := lti.ArrayLength(typ)
		if err != nil {
			return nil, err
		}

		elementType, err := lti.getReflectedType(et, customTypes)
		if err != nil {
			return nil, err
		}
		result := reflect.ArrayOf(size, elementType)
		log.Printf("GML: typeinfo.go: E: getReflectedType('%v') = %v", typ, result)
		return result, nil
	}

	if lti.IsMapType(typ) {
		kt, vt := lti.GetMapSubtypes(typ)
		if kt == "" || vt == "" {
			return nil, fmt.Errorf("invalid map type: %s", typ)
		}

		keyType, err := lti.getReflectedType(kt, customTypes)
		if err != nil {
			return nil, err
		}
		valType, err := lti.getReflectedType(vt, customTypes)
		if err != nil {
			return nil, err
		}

		result := reflect.MapOf(keyType, valType)
		log.Printf("GML: typeinfo.go: F: getReflectedType('%v') = %v", typ, result)
		return result, nil
	}

	// All other types are custom classes, which are represented as a map[string]any
	log.Printf("GML: typeinfo.go: G: getReflectedType('%v') = %v", typ, rtMapStringAny)
	return rtMapStringAny, nil
}

var rtMapStringAny = reflect.TypeFor[map[string]any]()
var reflectedTypeMap = map[string]reflect.Type{
	"bool":           reflect.TypeFor[bool](),
	"byte":           reflect.TypeFor[byte](),
	"uint":           reflect.TypeFor[uint](),
	"uint8":          reflect.TypeFor[uint8](),
	"uint16":         reflect.TypeFor[uint16](),
	"uint32":         reflect.TypeFor[uint32](),
	"uint64":         reflect.TypeFor[uint64](),
	"uintptr":        reflect.TypeFor[uintptr](),
	"int":            reflect.TypeFor[int](),
	"int8":           reflect.TypeFor[int8](),
	"int16":          reflect.TypeFor[int16](),
	"int32":          reflect.TypeFor[int32](),
	"int64":          reflect.TypeFor[int64](),
	"rune":           reflect.TypeFor[rune](),
	"float32":        reflect.TypeFor[float32](),
	"float64":        reflect.TypeFor[float64](),
	"string":         reflect.TypeFor[string](),
	"time.Time":      reflect.TypeFor[time.Time](),
	"time.Duration":  reflect.TypeFor[time.Duration](),
	"unsafe.Pointer": reflect.TypeFor[unsafe.Pointer](),
}
