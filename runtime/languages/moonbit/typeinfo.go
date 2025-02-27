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
	result, err := langsupport.GetTypeInfo(ctx, _langTypeInfo, typeName, typeCache)
	gmlPrintf("GML: moonbit/typeinfo.go: GetTypeInfo('%v') = %+v, err=%v", typeName, result, err)
	return result, err
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
		gmlPrintf("ERROR: moonbit/typeinfo.go: GetListSubtype('%v'): Bad list type!", typ)
		return ""
	}
	typ = strings.TrimSuffix(typ, "]")

	switch {
	case strings.HasPrefix(typ, "Array["):
		result := strings.TrimPrefix(typ, "Array[")
		gmlPrintf("GML: moonbit/typeinfo.go: GetListSubtype('%v') = '%v'", typ, result)
		return result
	case strings.HasPrefix(typ, "FixedArray["):
		result := strings.TrimPrefix(typ, "FixedArray[")
		gmlPrintf("GML: moonbit/typeinfo.go: GetListSubtype('%v') = '%v'", typ, result)
		return result
	default:
		gmlPrintf("ERROR: moonbit/typeinfo.go: GetListSubtype('%v'): Bad list type!", typ)
		return ""
	}
}

func (lti *langTypeInfo) GetMapSubtypes(typ string) (string, string) {
	typ, _, _ = stripErrorAndOption(typ)

	if !strings.HasSuffix(typ, "]") {
		gmlPrintf("ERROR: moonbit/typeinfo.go: GetMapSubtypes('%v'): Bad map type!", typ)
		return "", ""
	}

	const prefix = "Map[" // e.g. Map[String, Int]
	if !strings.HasPrefix(typ, prefix) {
		gmlPrintf("GML: moonbit/typeinfo.go: A: GetMapSubtypes('%v') = ('', '')", typ)
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
				r1, r2 := strings.TrimSpace(typ[:i]), strings.TrimSpace(typ[i+1:])
				gmlPrintf("GML: moonbit/typeinfo.go: B: GetMapSubtypes('%v') = ('%v', '%v')", typ, r1, r2)
				return r1, r2
			}
		}
	}

	gmlPrintf("GML: moonbit/typeinfo.go: C: GetMapSubtypes('%v') = ('', '')", typ)
	return "", ""
}

// Note that GetNameForType is used for GraphQL Schema generation and must
// strip all package, error, and option information.
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
	// case "@time.ZonedDateTime":
	// 	return "time.Time"
	// case "@wallClock.Datetime":
	// 	return "time.Time"
	case "@time.Duration":
		return "time.Duration"
	}

	if lti.IsOptionType(typ) { // TODO
		result := lti.GetNameForType(lti.GetUnderlyingType(typ)) + "?"
		gmlPrintf("GML: moonbit/typeinfo.go: A: GetNameForType('%v') = '%v'", typ, result)
		return result
	}

	// if lti.IsPointerType(typ) { // TODO
	// 	result := "*" + lti.GetNameForType(lti.GetUnderlyingType(typ))
	// 	gmlPrintf("GML: moonbit/typeinfo.go: A: GetNameForType('%v') = '%v'", typ, result)
	// 	return result
	// }

	if lti.IsListType(typ) {
		switch {
		case strings.HasPrefix(typ, "Array["):
			result := "Array[" + lti.GetNameForType(lti.GetListSubtype(typ)) + "]"
			gmlPrintf("GML: moonbit/typeinfo.go: B: GetNameForType('%v') = '%v'", typ, result)
			return result
		case strings.HasPrefix(typ, "FixedArray["):
			result := "FixedArray[" + lti.GetNameForType(lti.GetListSubtype(typ)) + "]"
			gmlPrintf("GML: moonbit/typeinfo.go: B: GetNameForType('%v') = '%v'", typ, result)
			return result
		default:
			gmlPrintf("PROGRAMMING ERROR: moonbit/typeinfo.go: GetNameForType('%v'): Bad list type!", typ)
		}
	}

	if lti.IsMapType(typ) {
		kt, vt := lti.GetMapSubtypes(typ)
		result := "Map[" + lti.GetNameForType(kt) + "," + lti.GetNameForType(vt) + "]"
		gmlPrintf("GML: moonbit/typeinfo.go: C: GetNameForType('%v') = '%v'", typ, result)
		return result
	}

	result := typ[strings.LastIndex(typ, ".")+1:] // strip package information
	gmlPrintf("GML: moonbit/typeinfo.go: D: GetNameForType('%v') = '%v'", typ, result)
	return result
}

func (lti *langTypeInfo) IsObjectType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	result := !lti.IsPrimitiveType(typ) &&
		!lti.IsListType(typ) &&
		!lti.IsMapType(typ) &&
		!lti.IsStringType(typ) &&
		!lti.IsTimestampType(typ) &&
		!lti.IsOptionType(typ)
		// !lti.IsPointerType(typ)
	gmlPrintf("GML: moonbit/typeinfo.go: IsObjectType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) GetUnderlyingType(typ string) (result string) {
	var hasError bool
	typ, hasError, _ = stripErrorAndOption(typ)

	if typ == "Unit" || hasError {
		gmlPrintf("GML: moonbit/typeinfo.go: GetUnderlyingType('Unit!Error') = 'Int'")
		return "Int" // GML: Experiment
	}

	// TODO
	// result := strings.TrimPrefix(typ, "*") // for Go
	switch typ {
	case "Bool", "Byte", "Char", "Int", "Int16", "Int64", "UInt", "UInt16", "UInt64", "Float", "Double", "String":
		result = typ
	default:
		result = typ
		gmlPrintf("GML: moonbit/typeinfo.go: GetUnderlyingType('%v') = '%v' - UNHANDLED DEFAULT CASE", typ, result)
	}
	gmlPrintf("GML: moonbit/typeinfo.go: GetUnderlyingType('%v') = '%v'", typ, result)
	return result
}

func (lti *langTypeInfo) IsListType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	if !strings.HasSuffix(typ, "]") {
		return false
	}
	result := strings.HasPrefix(typ, "Array[") || strings.HasPrefix(typ, "FixedArray[")
	gmlPrintf("GML: moonbit/typeinfo.go: IsListType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsSliceType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	if !strings.HasSuffix(typ, "]") {
		return false
	}
	// MoonBit Arrays and FixedArrays are similar to Go slices.
	result := strings.HasPrefix(typ, "Array[") || strings.HasPrefix(typ, "FixedArray[")
	gmlPrintf("GML: moonbit/typeinfo.go: IsSliceType('%v') = %v", typ, result)
	return result
}

// MoonBit does not have an equivalent fixed-length array type where the
// length is declared in the type.  Instead, a MoonBit Array is a slice type.
func (lti *langTypeInfo) IsArrayType(typ string) bool {
	// if !strings.HasSuffix(typ, "]")  {
	// 	return false
	// }
	// // MoonBit Arrays do not have a fixed length, unlike Go, so Array[T] is _NOT_ an "array" type.
	// // Instead, a MoonBit Array is a slice type.
	// result := strings.HasPrefix(typ, "FixedArray[")
	// gmlPrintf("GML: moonbit/typeinfo.go: IsArrayType('%v') = %v", typ, result)
	// return result
	return false
}

func (lti *langTypeInfo) IsBooleanType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	result := strings.HasPrefix(typ, "Bool")
	gmlPrintf("GML: moonbit/typeinfo.go: IsBooleanType('%v') = %v", typ, result)
	return result
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
		gmlPrintf("GML: moonbit/typeinfo.go: IsByteSequenceType('%v') = true", typ)
		return true
	}

	if lti.IsArrayType(typ) {
		subtype := lti.GetListSubtype(typ)
		switch {
		case strings.HasPrefix(subtype, "Byte"):
			gmlPrintf("GML: B: moonbit/typeinfo.go: IsByteSequenceType('%v') = true", typ)
			return true
		}
	}

	gmlPrintf("GML: moonbit/typeinfo.go: IsByteSequenceType('%v') = false", typ)
	return false
}

func (lti *langTypeInfo) IsFloatType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	result := strings.HasPrefix(typ, "Float") || strings.HasPrefix(typ, "Double")
	gmlPrintf("GML: moonbit/typeinfo.go: IsFloatType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsIntegerType(typ string) bool {
	var hasError bool
	typ, hasError, _ = stripErrorAndOption(typ)

	if typ == "Unit" || hasError {
		gmlPrintf("GML: moonbit/typeinfo.go: IsIntegerType('Unit!Error') = true")
		return true // GML: Experiment
	}

	switch typ {
	case "Int", "Int16", "Int64",
		"UInt", "UInt16", "UInt64",
		"Byte", "Char":
		gmlPrintf("GML: moonbit/typeinfo.go: IsIntegerType('%v') = true", typ)
		return true
	default:
		gmlPrintf("GML: moonbit/typeinfo.go: IsIntegerType('%v') = false", typ)
		return false
	}
}

func (lti *langTypeInfo) IsMapType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	if !strings.HasSuffix(typ, "]") {
		return false
	}
	result := strings.HasPrefix(typ, "Map[")
	gmlPrintf("GML: moonbit/typeinfo.go: IsMapType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsNullableType(typ string) bool {
	typ, hasError, hasOption := stripErrorAndOption(typ)

	if typ == "Unit" && hasError {
		gmlPrintf("GML: moonbit/typeinfo.go: IsNullableType('Unit!Error') = true")
		return true // GML: Experiment
	}

	// result := lti.IsPointerType(typ) || lti.IsSliceType(typ) || lti.IsMapType(typ)
	result := hasError || hasOption
	gmlPrintf("GML: moonbit/typeinfo.go: IsNullableType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsOptionType(typ string) bool {
	t, hasError, hasOption := stripErrorAndOption(typ)

	if t == "Unit" && hasError {
		gmlPrintf("GML: moonbit/typeinfo.go: IsOptionType('Unit!Error') = true")
		return true // GML: Experiment
	}

	gmlPrintf("GML: moonbit/typeinfo.go: IsOptionType('%v') = %v", typ, hasOption)
	return hasOption
}

// NOTE! This is _NOT_ a pointer type in MoonBit! (but is the closest thing to a pointer type)
// To satisfy the languages.TypeInfo interface, we must implement this method!
// This is used within the handlers to determine, for example, if a Go-like `string`
// or a Go-like `*string` should be returned.
func (lti *langTypeInfo) IsPointerType(typ string) bool {
	return lti.IsOptionType(typ)
}

// func (lti *langTypeInfo) IsPointerType(typ string) bool {
// 	// TODO: Is Ref[T] the only pointer/reference type in MoonBit?
// 	// result := strings.HasPrefix(typ, "*")  // for Go
// 	// WRONG! Option[T] is _NOT_ a pointer type! // result := strings.HasSuffix(typ, "?")
// 	result := false
// 	// result := strings.HasSuffix(typ, "?") // This is currently needed for the test suite!!!  FIND OUT WHY!!!
// 	gmlPrintf("GML: moonbit/typeinfo.go: IsPointerType('%v') = %v", typ, result)
// 	return result
// }

func (lti *langTypeInfo) IsPrimitiveType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	result := lti.IsBooleanType(typ) || lti.IsIntegerType(typ) || lti.IsFloatType(typ) // TODO(gmlewis)
	gmlPrintf("GML: moonbit/typeinfo.go: IsPrimitiveType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsSignedIntegerType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	switch typ {
	case "Int", "Int16", "Int64":
		gmlPrintf("GML: moonbit/typeinfo.go: IsSignedIntegerType('%v') = true", typ)
		return true
	default:
		gmlPrintf("GML: moonbit/typeinfo.go: IsSignedIntegerType('%v') = false", typ)
		return false
	}
}

func (lti *langTypeInfo) IsStringType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	result := strings.HasPrefix(typ, "String")
	gmlPrintf("GML: moonbit/typeinfo.go: IsStringType('%v') = %v", typ, result)
	return result
}

func (lti *langTypeInfo) IsTimestampType(typ string) bool {
	typ, _, _ = stripErrorAndOption(typ)

	// Special case for MoonBit moonbitlang/x or wasi "timestamp"-like struct.
	result := strings.HasPrefix(typ, "@time.ZonedDateTime") || strings.HasPrefix(typ, "@wallClock.Datetime")
	gmlPrintf("GML: moonbit/typeinfo.go: IsTimestampType('%v') = %v", typ, result)
	return result
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

// TODO: DRY - copied from sdk/go/tools/modus-moonbit-build/packages/parse-params.go
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
	gmlPrintf("PROGRAMMING ERROR: GML: moonbit/typeinfo.go: ArrayLength('%v'): Bad array type!", typ)
	return 0, nil
	// gmlPrintf("GML: moonbit/typeinfo.go: ENTER ArrayLength('%v')", typ)
	// i := strings.Index(typ, "]")
	// if i == -1 {
	// 	return -1, fmt.Errorf("invalid array type: %s", typ)
	// }

	// size := typ[1:i]
	// if size == "" {
	// 	return -1, fmt.Errorf("invalid array type: %s", typ)
	// }

	// parsedSize, err := strconv.Atoi(size)
	// if err != nil {
	// 	return -1, err
	// }
	// if parsedSize < 0 || parsedSize > math.MaxUint32 {
	// 	return -1, fmt.Errorf("array size out of bounds: %s", size)
	// }

	// gmlPrintf("GML: moonbit/typeinfo.go: ArrayLength('%v') = %v", typ, parsedSize)
	// return parsedSize, nil
}

// func (lti *langTypeInfo) getSizeOfArray(ctx context.Context, typ string) (uint32, error) {
// 	gmlPrintf("PROGRAMMING ERROR: GML: moonbit/typeinfo.go: getSizeOfArray('%v'): Bad array type!", typ)
// 	return 0, nil
// }

func (lti *langTypeInfo) getSizeOfStruct(ctx context.Context, typ string) (uint32, error) {
	def, err := lti.GetTypeDefinition(ctx, typ)
	if err != nil {
		return 0, err
	}
	if len(def.Fields) == 0 {
		return 0, nil
	}

	offset := uint32(0)
	maxAlign := uint32(1)
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
		offset = langsupport.AlignOffset(offset, alignment)
		gmlPrintf("GML: typeinfo.go: getSizeOfStruct: name='%v.%v', size=%v, alignment=%v, offset=%v", def.Name, field.Name, size, alignment, offset)
		offset += size
	}

	size := langsupport.AlignOffset(offset, maxAlign)
	return size, nil
}

func (lti *langTypeInfo) GetAlignmentOfType(ctx context.Context, typ string) (uint32, error) {
	var hasOption bool
	typ, _, hasOption = stripErrorAndOption(typ)
	// reference: https://github.com/tinygo-org/tinygo/blob/release/compiler/sizes.go

	if hasOption {
		gmlPrintf("GML: moonbit/typeinfo.go: C: GetAlignmentOfType('%v?') = 4", typ)
		return 4, nil
	}

	// primitives align to their natural size
	if lti.IsPrimitiveType(typ) {
		result, err := lti.GetSizeOfType(ctx, typ)
		gmlPrintf("GML: moonbit/typeinfo.go: A: GetAlignmentOfType('%v') = %v, err=%v", typ, result, err)
		return result, err
	}

	// arrays align to the alignment of their element type
	if lti.IsArrayType(typ) {
		t := lti.GetListSubtype(typ)
		result, err := lti.GetAlignmentOfType(ctx, t)
		gmlPrintf("GML: moonbit/typeinfo.go: B: GetAlignmentOfType('%v') = %v, err=%v", typ, result, err)
		return result, err
	}

	// reference types align to the pointer size (4 bytes on 32-bit wasm)
	// if lti.IsPointerType(typ) || lti.IsSliceType(typ) || lti.IsStringType(typ) || lti.IsMapType(typ) {
	if lti.IsSliceType(typ) || lti.IsStringType(typ) || lti.IsMapType(typ) {
		gmlPrintf("GML: moonbit/typeinfo.go: C: GetAlignmentOfType('%v') = 4", typ)
		return 4, nil
	}

	// time.Time has 3 fields, the maximum alignment is 8 bytes
	if lti.IsTimestampType(typ) {
		gmlPrintf("GML: moonbit/typeinfo.go: D: GetAlignmentOfType('%v') = 8", typ)
		return 8, nil
	}

	// structs align to the maximum alignment of their fields
	result, err := lti.getAlignmentOfStruct(ctx, typ)
	gmlPrintf("GML: moonbit/typeinfo.go: E: GetAlignmentOfType('%v') = %v, err=%v", typ, result, err)
	return result, err
}

func (lti *langTypeInfo) ObjectsUseMaxFieldAlignment() bool {
	return true // TODO(gmlewis)
}

func (lti *langTypeInfo) getAlignmentOfStruct(ctx context.Context, typ string) (uint32, error) {
	def, err := lti.GetTypeDefinition(ctx, typ)
	if err != nil {
		return 0, err
	}

	max := uint32(1)
	for _, field := range def.Fields {
		align, err := lti.GetAlignmentOfType(ctx, field.Type)
		if err != nil {
			return 0, err
		}
		if align > max {
			max = align
		}
	}

	gmlPrintf("GML: moonbit/typeinfo.go: getAlignmentOfStruct('%v') = %v", typ, max)
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
		gmlPrintf("GML: moonbit/typeinfo.go: GetEncodingLengthOfType('%v?') = 4 + errorSize=%v", typ, errorSize)
		return 1 + errorSize, nil
	}

	// if lti.IsPrimitiveType(typ) || lti.IsPointerType(typ) || lti.IsMapType(typ) {
	if lti.IsPrimitiveType(typ) || lti.IsMapType(typ) {
		gmlPrintf("GML: moonbit/typeinfo.go: A: GetEncodingLengthOfType('%v') = 1 + errorSize=%v", typ, errorSize)
		return 1 + errorSize, nil
	} else if lti.IsStringType(typ) {
		gmlPrintf("GML: moonbit/typeinfo.go: B: GetEncodingLengthOfType('%v') = 1 + errorSize=%v", typ, errorSize)
		return 1 + errorSize, nil
	} else if lti.IsSliceType(typ) || lti.IsTimestampType(typ) {
		gmlPrintf("GML: moonbit/typeinfo.go: C: GetEncodingLengthOfType('%v') = 3 + errorSize=%v", typ, errorSize)
		return 3 + errorSize, nil
	} else if lti.IsArrayType(typ) {
		result, err := lti.getEncodingLengthOfArray(ctx, typ)
		gmlPrintf("GML: moonbit/typeinfo.go: D: GetEncodingLengthOfType('%v') = %v+%v, err=%v", typ, result, errorSize, err)
		return result + errorSize, err
	} else if lti.IsObjectType(typ) {
		result, err := lti.getEncodingLengthOfStruct(ctx, typ)
		gmlPrintf("GML: moonbit/typeinfo.go: E: GetEncodingLengthOfType('%v') = %v+%v, err=%v", typ, result, errorSize, err)
		return result + errorSize, err
	}

	return 0, fmt.Errorf("unable to determine encoding length for type: %s", typ)
}

func (lti *langTypeInfo) getEncodingLengthOfArray(ctx context.Context, typ string) (uint32, error) {
	gmlPrintf("PROGRAMMING ERROR: GML: moonbit/typeinfo.go: getEncodingLengthOfArray('%v'): Bad array type!", typ)
	return 0, nil
	// arrSize, err := lti.ArrayLength(typ)
	// if err != nil {
	// 	return 0, err
	// }
	// if arrSize == 0 {
	// 	gmlPrintf("GML: moonbit/typeinfo.go: A: getEncodingLengthOfArray('%v') = 0", typ)
	// 	return 0, nil
	// }

	// t := lti.GetListSubtype(typ)
	// elementLen, err := lti.GetEncodingLengthOfType(ctx, t)
	// if err != nil {
	// 	return 0, err
	// }

	// result := uint32(arrSize) * elementLen
	// gmlPrintf("GML: moonbit/typeinfo.go: B: getEncodingLengthOfArray('%v') = %v", typ, result)
	// return result, nil
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

	gmlPrintf("GML: moonbit/typeinfo.go: getEncodingLengthOfStruct('%v') = %v", typ, total)
	return total, nil
}

func (lti *langTypeInfo) GetSizeOfType(ctx context.Context, typ string) (uint32, error) {
	var hasError, hasOption bool
	typ, hasError, hasOption = stripErrorAndOption(typ) // TODO: Add size of error here?

	if hasOption {
		gmlPrintf("GML: moonbit/typeinfo.go: F: GetSizeOfType('%v?') = 4", typ)
		return 4, nil
	}

	if typ == "Unit" || hasError {
		gmlPrintf("GML: moonbit/typeinfo.go: GetSizeOfType('Unit!Error') = 1")
		return 1, nil // GML: Experiment
	}

	switch typ {
	case "Byte":
		gmlPrintf("GML: moonbit/typeinfo.go: A: GetSizeOfType('%v') = 1", typ)
		return 1, nil
	case "Char", "Int16", "UInt16":
		gmlPrintf("GML: moonbit/typeinfo.go: B: GetSizeOfType('%v') = 2", typ)
		return 2, nil
	case "Bool", "Int", "UInt", "Float": // we only support 32-bit wasm
		gmlPrintf("GML: moonbit/typeinfo.go: C: GetSizeOfType('%v') = 4", typ)
		return 4, nil
	case "Int64", "UInt64", "Double", "@time.Duration":
		gmlPrintf("GML: moonbit/typeinfo.go: D: GetSizeOfType('%v') = 8", typ)
		return 8, nil
	}

	if lti.IsStringType(typ) {
		gmlPrintf("GML: moonbit/typeinfo.go: E: GetSizeOfType('%v') = 4", typ)
		return 4, nil
	}

	if lti.IsMapType(typ) {
		// maps are passed by reference using a 4 byte pointer
		gmlPrintf("GML: moonbit/typeinfo.go: G: GetSizeOfType('%v') = 4", typ)
		return 4, nil
	}

	if lti.IsSliceType(typ) {
		gmlPrintf("GML: moonbit/typeinfo.go: H: GetSizeOfType('%v') = 4", typ)
		return 4, nil
	}

	if lti.IsTimestampType(typ) {
		// time.Time has 3 fields: 8 byte uint64, 8 byte int64, 4 byte pointer
		gmlPrintf("GML: moonbit/typeinfo.go: I: GetSizeOfType('%v') = 20", typ)
		return 20, nil
	}

	// MoonBit has _NO_ concept of a Go (fixed-length) "array" type.
	// if lti.IsArrayType(typ) {
	// 	result, err := lti.getSizeOfArray(ctx, typ)
	// 	gmlPrintf("GML: moonbit/typeinfo.go: J: GetSizeOfType('%v') = %v", typ, result)
	// 	return result, err
	// }

	result, err := lti.getSizeOfStruct(ctx, typ)
	gmlPrintf("GML: moonbit/typeinfo.go: K: GetSizeOfType('%v') = %v", typ, result)
	return result, err
}

func (lti *langTypeInfo) GetTypeDefinition(ctx context.Context, typ string) (*metadata.TypeDefinition, error) {
	// TODO: Does Unit!Error need to be handled here?

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
	gmlPrintf("GML: moonbit/typeinfo.go: GetTypeDefinition('%v') = %+v, err=%v", typ, result, err)
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
		result, err := lti.getReflectedType(typ, customTypes)
		gmlPrintf("GML: moonbit/typeinfo.go: A: GetReflectedType('%v') = %v, Kind()=%v", typ, result, result.Kind())
		return result, err
	}
	result, err := lti.getReflectedType(typ, nil)
	gmlPrintf("GML: moonbit/typeinfo.go: B: GetReflectedType('%v') = %v, Kind()=%v", typ, result, result.Kind())
	return result, err
}

func (lti *langTypeInfo) getReflectedType(typ string, customTypes map[string]reflect.Type) (reflect.Type, error) {
	gmlPrintf("GML: moonbit/typeinfo.go: ENTER getReflectedType('%v')", typ)
	if customTypes != nil {
		if rt, ok := customTypes[typ]; ok {
			gmlPrintf("GML: moonbit/typeinfo.go: A: getReflectedType('%v') = %v, Kind()=%v", typ, rt, rt.Kind())
			return rt, nil
		}
	}

	if rt, ok := reflectedTypeMap[typ]; ok {
		gmlPrintf("GML: moonbit/typeinfo.go: B: getReflectedType('%v') = %v, Kind()=%v", typ, rt, rt.Kind())
		return rt, nil
	}

	// TODO: Does Unit!Error need to be handled here?

	if lti.IsOptionType(typ) {
		tt := lti.GetUnderlyingType(typ)
		targetType, err := lti.getReflectedType(tt, customTypes)
		if err != nil {
			return nil, err
		}
		result := reflect.PointerTo(targetType)
		gmlPrintf("GML: moonbit/typeinfo.go: C: getReflectedType('%v') = %v, Kind()=%v", typ, result, result.Kind())
		return result, nil
	}

	// if lti.IsPointerType(typ) {
	// 	tt := lti.GetUnderlyingType(typ)
	// 	targetType, err := lti.getReflectedType(tt, customTypes)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	result := reflect.PointerTo(targetType)
	// 	gmlPrintf("GML: moonbit/typeinfo.go: C: getReflectedType('%v') = %v", typ, result)
	// 	return result, nil
	// }

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
		gmlPrintf("GML: moonbit/typeinfo.go: D: getReflectedType('%v') = %v, Kind()=%v", typ, result, result.Kind())
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
		gmlPrintf("GML: moonbit/typeinfo.go: E: getReflectedType('%v') = %v, Kind()=%v", typ, result, result.Kind())
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
		gmlPrintf("GML: moonbit/typeinfo.go: F: getReflectedType('%v') = %v, Kind()=%v", typ, result, result.Kind())
		return result, nil
	}

	// All other types are custom classes, which are represented as a map[string]any
	gmlPrintf("GML: moonbit/typeinfo.go: G: getReflectedType('%v') = %v, Kind()=%v", typ, rtMapStringAny, rtMapStringAny.Kind())
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
