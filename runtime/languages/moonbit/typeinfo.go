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
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/neo4jclient"
	"github.com/gmlewis/modus/runtime/utils"
)

var _langTypeInfo = &langTypeInfo{}

func LanguageTypeInfo() langsupport.LanguageTypeInfo {
	return _langTypeInfo
}

func GetTypeInfo(ctx context.Context, typeName string, typeCache map[string]langsupport.TypeInfo) (langsupport.TypeInfo, error) {
	// DO NOT STRIP THE ERROR TYPE HERE! Strip it later.
	// When an "...!Error" is the return type, two values are returned.
	// The first value is 0 on failure, and the second value is the actual return type.
	return langsupport.GetTypeInfo(ctx, _langTypeInfo, typeName, typeCache)
}

func stripErrorAndOption(typeSignature string) (typ string, hasError, hasOption bool) {
	if i := strings.Index(typeSignature, "!"); i >= 0 {
		hasError = true
		typeSignature = typeSignature[:i]
	}
	hasOption = strings.HasSuffix(typeSignature, "?")
	return strings.TrimSuffix(typeSignature, "?"), hasError, hasOption
}

type langTypeInfo struct{}

func (lti *langTypeInfo) GetListSubtype(typ string) string {
	typ, _, _ = stripErrorAndOption(typ)

	if !strings.HasSuffix(typ, "]") {
		return ""
	}
	typ = strings.TrimSuffix(typ, "]")

	switch {
	case strings.HasPrefix(typ, "Array["):
		return strings.TrimPrefix(typ, "Array[")
	case strings.HasPrefix(typ, "FixedArray["):
		return strings.TrimPrefix(typ, "FixedArray[")
	default:
		return ""
	}
}

func (lti *langTypeInfo) GetMapSubtypes(typ string) (string, string) {
	typ, _, _ = stripErrorAndOption(typ)

	if !strings.HasSuffix(typ, "]") {
		return "", ""
	}

	const prefix = "Map[" // e.g. Map[String, Int]
	if !strings.HasPrefix(typ, prefix) {
		return "", ""
	}
	typ = strings.TrimSuffix(typ, "]")
	typ = strings.TrimPrefix(typ, prefix)

	n := 1
	for i := 0; i < len(typ); i++ {
		switch typ[i] {
		case '[':
			n++
		case ']':
			n--
		case ',':
			if n == 1 {
				return strings.TrimSpace(typ[:i]), strings.TrimSpace(typ[i+1:])
			}
		}
	}

	return "", ""
}

// Note that GetNameForType is used for GraphQL Schema generation and must
// strip all package, error, and option information.
// TODO: This needs to be kept in sync with sdk/go/tools/modus-moonbit-build/utils/typeinfo.go GetNameForType()
func (lti *langTypeInfo) GetNameForType(typ string) string {
	// "github.com/gmlewis/modus/sdk/go/examples/simple.Person" -> "Person"
	var hasError bool
	typ, hasError, _ = stripErrorAndOption(typ)

	if typ == "Unit" && hasError {
		return "Unit!Error" // Special case - used internally and not by GraphQL.
	} else if typ == "Unit" {
		return "Unit"
	}

	// Handle special cases that the GraphQL engine expects to see:
	switch typ {
	case "@time.Duration":
		return "time.Duration"
	}

	if lti.IsOptionType(typ) {
		return lti.GetNameForType(lti.GetUnderlyingType(typ)) + "?"
	}

	if lti.IsListType(typ) {
		switch {
		case strings.HasPrefix(typ, "Array["):
			return "Array[" + lti.GetNameForType(lti.GetListSubtype(typ)) + "]"
		case strings.HasPrefix(typ, "FixedArray["):
			return "FixedArray[" + lti.GetNameForType(lti.GetListSubtype(typ)) + "]"
		default:
			log.Fatalf("PROGRAMMING ERROR: moonbit/typeinfo.go: GetNameForType('%v'): Bad list type!", typ)
		}
	}

	if lti.IsMapType(typ) {
		kt, vt := lti.GetMapSubtypes(typ)
		return "Map[" + lti.GetNameForType(kt) + "," + lti.GetNameForType(vt) + "]"
	}

	return typ[strings.LastIndex(typ, ".")+1:] // strip package information
}

func (lti *langTypeInfo) IsObjectType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	return !lti.IsPrimitiveType(typ) &&
		!lti.IsListType(typ) &&
		!lti.IsMapType(typ) &&
		!lti.IsStringType(typ) &&
		!lti.IsTimestampType(typ) &&
		!lti.IsOptionType(typ)
}

func (lti *langTypeInfo) GetUnderlyingType(typ string) (result string) {
	var hasError bool
	typ, hasError, _ = stripErrorAndOption(typ)

	if typ == "Unit" || hasError {
		return "Int"
	}

	return typ
}

func (lti *langTypeInfo) IsListType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	if !strings.HasSuffix(typ, "]") {
		return false
	}
	return strings.HasPrefix(typ, "Array[") || strings.HasPrefix(typ, "FixedArray[")
}

func (lti *langTypeInfo) IsSliceType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	if !strings.HasSuffix(typ, "]") {
		return false
	}
	// MoonBit Arrays and FixedArrays are similar to Go slices.
	return strings.HasPrefix(typ, "Array[") || strings.HasPrefix(typ, "FixedArray[")
}

// MoonBit does not have an equivalent fixed-length array type where the
// length is declared in the type.  Instead, a MoonBit Array is a slice type.
func (lti *langTypeInfo) IsArrayType(typ string) bool {
	return false
}

func (lti *langTypeInfo) IsBooleanType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	return strings.HasPrefix(typ, "Bool")
}

func (lti *langTypeInfo) IsByteSequenceType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	switch {
	case
		strings.HasPrefix(typ, "Array[Byte]"),
		strings.HasPrefix(typ, "ArrayView[Byte]"),
		strings.HasPrefix(typ, "FixedArray[Byte]"),
		strings.HasPrefix(typ, "Bytes"):
		// strings.HasPrefix(typ, "BytesView"),  // covered by last case
		return true
	}

	if lti.IsArrayType(typ) {
		subtype := lti.GetListSubtype(typ)
		return strings.HasPrefix(subtype, "Byte")
	}

	return false
}

func (lti *langTypeInfo) IsFloatType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	return strings.HasPrefix(typ, "Float") || strings.HasPrefix(typ, "Double")
}

func (lti *langTypeInfo) IsIntegerType(typ string) bool {
	var hasError bool
	typ, hasError, _ = stripErrorAndOption(typ)

	if typ == "Unit" || hasError {
		return true
	}

	switch typ {
	case "Int", "Int16", "Int64",
		"UInt", "UInt16", "UInt64",
		"Byte", "Char":
		return true
	default:
		return false
	}
}

func (lti *langTypeInfo) IsMapType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	if !strings.HasSuffix(typ, "]") {
		return false
	}
	return strings.HasPrefix(typ, "Map[")
}

func (lti *langTypeInfo) IsNullableType(typ string) bool {
	typ, hasError, hasOption := stripErrorAndOption(typ)

	if typ == "Unit" && hasError {
		return true
	}

	return hasError || hasOption
}

func (lti *langTypeInfo) IsOptionType(typ string) bool {
	t, hasError, hasOption := stripErrorAndOption(typ)

	if t == "Unit" && hasError {
		return true
	}

	return hasOption
}

// NOTE! This is _NOT_ a pointer type in MoonBit! (but is the closest thing to a pointer type)
// To satisfy the languages.TypeInfo interface, we must implement this method!
// This is used within the handlers to determine, for example, if a Go-like `string`
// or a Go-like `*string` should be returned.
func (lti *langTypeInfo) IsPointerType(typ string) bool {
	return lti.IsOptionType(typ)
}

func (lti *langTypeInfo) IsPrimitiveType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	return lti.IsBooleanType(typ) || lti.IsIntegerType(typ) || lti.IsFloatType(typ)
}

func (lti *langTypeInfo) IsSignedIntegerType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	switch typ {
	case "Int", "Int16", "Int64":
		return true
	default:
		return false
	}
}

func (lti *langTypeInfo) IsStringType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	return strings.HasPrefix(typ, "String")
}

func (lti *langTypeInfo) IsTimestampType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	// Special case for MoonBit moonbitlang/x or wasi "timestamp"-like struct.
	return strings.HasPrefix(typ, "@time.ZonedDateTime") || strings.HasPrefix(typ, "@wallClock.Datetime")
}

func (lti *langTypeInfo) IsErrorType(typ string) (string, bool) {
	if i := strings.Index(typ, "!"); i >= 0 {
		return typ[:i], true
	}
	return typ, false
}

func (lti *langTypeInfo) IsTupleType(typ string) bool {
	return strings.HasPrefix(typ, "(")
}

func (lti *langTypeInfo) GetTupleSubtypes(typ string) []string {
	typ = typ[1 : len(typ)-1]
	return splitParamsWithBrackets(typ)
}

// TODO: DRY - copied from sdk/go/tools/modus-moonbit-build/utils/split-params.go
func splitParamsWithBrackets(allArgs string) []string {
	var result []string
	var n int     // count bracket pairs
	var start int // start of current arg
	for i := 0; i < len(allArgs); i++ {
		switch allArgs[i] {
		case '[':
			n++
		case ']':
			n--
		case ',':
			if n == 0 {
				arg := strings.TrimSpace(allArgs[start:i])
				if arg != "" {
					result = append(result, arg)
				}
				start = i + 1
			}
		}
	}
	if start < len(allArgs) {
		arg := strings.TrimSpace(allArgs[start:])
		if arg != "" {
			result = append(result, arg)
		}
	}

	return result
}

// FixedArrays in MoonBit do not declare their size.
func (lti *langTypeInfo) ArrayLength(typ string) (int, error) {
	log.Fatalf("PROGRAMMING ERROR: GML: moonbit/typeinfo.go: ArrayLength('%v'): Bad array type!", typ)
	return 0, nil
}

func (lti *langTypeInfo) getSizeOfStruct(ctx context.Context, typ string) (uint32, error) {
	def, err := lti.GetTypeDefinition(ctx, typ)
	if err != nil {
		return 0, err
	}
	if len(def.Fields) == 0 {
		return 0, nil
	}

	offset := uint32(0)
	maxAlign := uint32(0)
	for _, field := range def.Fields {
		size, err := lti.GetSizeOfType(ctx, field.Type)
		if err != nil {
			return 0, err
		}
		alignment, err := lti.GetAlignmentOfType(ctx, field.Type)
		if err != nil {
			return 0, err
		}
		if alignment > maxAlign {
			maxAlign = alignment
		}
		offset = langsupport.AlignOffset(offset, alignment) + size
	}

	size := langsupport.AlignOffset(offset, maxAlign)
	return size, nil
}

func (lti *langTypeInfo) GetAlignmentOfType(ctx context.Context, typ string) (uint32, error) {
	var hasOption bool
	typ, _, hasOption = stripErrorAndOption(typ)
	// reference: https://github.com/tinygo-org/tinygo/blob/release/compiler/sizes.go

	if hasOption {
		return 4, nil
	}

	// primitives align to their natural size
	if lti.IsPrimitiveType(typ) {
		return lti.GetSizeOfType(ctx, typ)
	}

	// arrays align to the alignment of their element type
	if lti.IsArrayType(typ) {
		t := lti.GetListSubtype(typ)
		return lti.GetAlignmentOfType(ctx, t)
	}

	// reference types align to the pointer size (4 bytes on 32-bit wasm)
	if lti.IsSliceType(typ) || lti.IsStringType(typ) || lti.IsMapType(typ) {
		return 4, nil
	}

	// time.Time has 3 fields, the maximum alignment is 8 bytes
	if lti.IsTimestampType(typ) {
		return 8, nil
	}

	// structs align to the maximum alignment of their fields
	return lti.getAlignmentOfStruct(ctx, typ)
}

func (lti *langTypeInfo) ObjectsUseMaxFieldAlignment() bool {
	return true
}

func (lti *langTypeInfo) getAlignmentOfStruct(ctx context.Context, typ string) (uint32, error) {
	def, err := lti.GetTypeDefinition(ctx, typ)
	if err != nil {
		return 0, err
	}

	max := uint32(4)
	for _, field := range def.Fields {
		align, err := lti.GetAlignmentOfType(ctx, field.Type)
		if err != nil {
			return 0, err
		}
		if align > max {
			max = align
		}
	}

	return max, nil
}

func (lti *langTypeInfo) GetDataSizeOfType(ctx context.Context, typ string) (uint32, error) {
	return lti.GetSizeOfType(ctx, typ)
}

func (lti *langTypeInfo) GetEncodingLengthOfType(ctx context.Context, typ string) (uint32, error) {
	// NOTE that the encoding length refers to the number of uint64 words on the wazero stack
	// needed to represent the type, which for MoonBit is typically 1 (or 2 if an error code is included).
	// This is used by the wasm runtime to determine how many words to pop off the stack after a function call.
	// The encoding length is not the same as the size of the type in memory.
	var hasError, hasOption bool
	typ, hasError, hasOption = stripErrorAndOption(typ)
	var errorSize uint32
	if hasError {
		errorSize = 1
	}

	if hasOption {
		return 1 + errorSize, nil
	}

	if lti.IsPrimitiveType(typ) || lti.IsMapType(typ) {
		return 1 + errorSize, nil
	} else if lti.IsStringType(typ) {
		return 1 + errorSize, nil
	} else if lti.IsSliceType(typ) || lti.IsTimestampType(typ) {
		return 3 + errorSize, nil
	} else if lti.IsArrayType(typ) {
		result, err := lti.getEncodingLengthOfArray(ctx, typ)
		return result + errorSize, err
	} else if lti.IsObjectType(typ) {
		result, err := lti.getEncodingLengthOfStruct(ctx, typ)
		return result + errorSize, err
	}

	return 0, fmt.Errorf("unable to determine encoding length for type: %s", typ)
}

func (lti *langTypeInfo) getEncodingLengthOfArray(ctx context.Context, typ string) (uint32, error) {
	log.Fatalf("PROGRAMMING ERROR: GML: moonbit/typeinfo.go: getEncodingLengthOfArray('%v'): Bad array type!", typ)
	return 0, nil
}

func (lti *langTypeInfo) getEncodingLengthOfStruct(ctx context.Context, typ string) (uint32, error) {
	def, err := lti.GetTypeDefinition(ctx, typ)
	if err != nil {
		return 0, err
	}

	total := uint32(0)
	for _, field := range def.Fields {
		len, err := lti.GetEncodingLengthOfType(ctx, field.Type)
		if err != nil {
			return 0, err
		}
		total += len
	}

	return total, nil
}

func (lti *langTypeInfo) GetSizeOfType(ctx context.Context, typ string) (uint32, error) {
	var hasError, hasOption bool
	typ, hasError, hasOption = stripErrorAndOption(typ)

	if hasOption {
		return 4, nil
	}

	if typ == "Unit" || hasError {
		return 1, nil
	}

	switch typ {
	case "Byte":
		return 1, nil
	case "Char", "Int16", "UInt16":
		return 2, nil
	case "Bool", "Int", "UInt", "Float": // we only support 32-bit wasm
		return 4, nil
	case "Int64", "UInt64", "Double", "@time.Duration":
		return 8, nil
	}

	if lti.IsStringType(typ) {
		return 4, nil
	}

	if lti.IsMapType(typ) {
		// maps are passed by reference using a 4 byte pointer
		return 4, nil
	}

	if lti.IsSliceType(typ) {
		return 4, nil
	}

	if lti.IsTimestampType(typ) {
		// time.Time has 3 fields: 8 byte uint64, 8 byte int64, 4 byte pointer
		return 20, nil
	}

	// MoonBit has _NO_ concept of a Go (fixed-length) "array" type.

	return lti.getSizeOfStruct(ctx, typ)
}

func (lti *langTypeInfo) GetTypeDefinition(ctx context.Context, typ string) (*metadata.TypeDefinition, error) {
	md, ok := ctx.Value(utils.MetadataContextKey).(*metadata.Metadata)
	if !ok {
		return nil, fmt.Errorf("moonbit/typeinfo.go: GetTypeDefinition('%v'): metadata not found in context", typ)
	}
	// Must use original type definition from SDK (with default values) to match the saved metadata.
	result, err := md.GetTypeDefinition(typ)
	if err != nil {
		typ, _, _ = stripErrorAndOption(typ) // try locating the type without the error or option suffix
		result, err = md.GetTypeDefinition(typ)
	}

	return result, err
}

func (lti *langTypeInfo) GetReflectedType(ctx context.Context, typ string) (reflect.Type, error) {
	// NOTE that the reflected type tells wazero how to interpret the handler's type with respect to Go's type system.
	var hasOption bool
	typ, _, hasOption = stripErrorAndOption(typ)

	// Strip !Error information, but add back in the option type.
	if hasOption {
		typ += "?"
	}

	if customTypes, ok := ctx.Value(utils.CustomTypesContextKey).(map[string]reflect.Type); ok {
		return lti.getReflectedType(typ, customTypes)
	}

	return lti.getReflectedType(typ, nil)
}

func (lti *langTypeInfo) getReflectedType(typ string, customTypes map[string]reflect.Type) (reflect.Type, error) {
	if customTypes != nil {
		if rt, ok := customTypes[typ]; ok {
			return rt, nil
		}
	}

	if rt, ok := reflectedTypeMap[typ]; ok {
		return rt, nil
	}

	if lti.IsOptionType(typ) {
		tt := lti.GetUnderlyingType(typ)
		targetType, err := lti.getReflectedType(tt, customTypes)
		if err != nil {
			return nil, err
		}

		return reflect.PointerTo(targetType), nil
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

		return reflect.SliceOf(elementType), nil
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

		return reflect.ArrayOf(size, elementType), nil
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

		return reflect.MapOf(keyType, valType), nil
	}

	// All other types are custom classes, which are represented as a map[string]any
	return rtMapStringAny, nil
}

var rtMapStringAny = reflect.TypeFor[map[string]any]()
var reflectedTypeMap = map[string]reflect.Type{
	"@neo4j.EagerResult":   reflect.TypeFor[neo4jclient.EagerResult](),
	"@neo4j.EagerResult?":  reflect.TypeFor[*neo4jclient.EagerResult](),
	"@neo4j.Record":        reflect.TypeFor[neo4jclient.Record](),
	"@neo4j.Record?":       reflect.TypeFor[*neo4jclient.Record](),
	"@time.Duration?":      reflect.TypeFor[*time.Duration](),
	"@time.Duration":       reflect.TypeFor[time.Duration](),
	"@time.ZonedDateTime?": reflect.TypeFor[*time.Time](),
	"@time.ZonedDateTime":  reflect.TypeFor[time.Time](),
	"@wallClock.Datetime?": reflect.TypeFor[*time.Time](),
	"@wallClock.Datetime":  reflect.TypeFor[time.Time](),
	"Bool?":                reflect.TypeFor[*bool](),
	"Bool":                 reflect.TypeFor[bool](),
	"Byte?":                reflect.TypeFor[*byte](),
	"Byte":                 reflect.TypeFor[byte](),
	"Char?":                reflect.TypeFor[*int16](),
	"Char":                 reflect.TypeFor[int16](),
	"Double?":              reflect.TypeFor[*float64](),
	"Double":               reflect.TypeFor[float64](),
	"Float?":               reflect.TypeFor[*float32](),
	"Float":                reflect.TypeFor[float32](),
	"Int?":                 reflect.TypeFor[*int32](),
	"Int":                  reflect.TypeFor[int32](),
	"Int16?":               reflect.TypeFor[*int16](),
	"Int16":                reflect.TypeFor[int16](),
	"Int64?":               reflect.TypeFor[*int64](),
	"Int64":                reflect.TypeFor[int64](),
	"String?":              reflect.TypeFor[*string](),
	"String":               reflect.TypeFor[string](),
	"UInt?":                reflect.TypeFor[*uint32](),
	"UInt":                 reflect.TypeFor[uint32](),
	"UInt16?":              reflect.TypeFor[*uint16](),
	"UInt16":               reflect.TypeFor[uint16](),
	"UInt64?":              reflect.TypeFor[*uint64](),
	"UInt64":               reflect.TypeFor[uint64](),
}
