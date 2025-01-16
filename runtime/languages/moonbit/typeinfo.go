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
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/gmlewis/modus/lib/metadata"
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
	const prefix = "Map[" // e.g. Map[String,Int]
	if !strings.HasPrefix(typ, prefix) {
		log.Printf("GML: typeinfo.go: A: GetMapSubtypes('%v') = ('', '')", typ)
		return "", ""
	}

	n := 1
	for i := len(prefix); i < len(typ); i++ {
		switch typ[i] {
		case '[':
			n++
		case ']':
			n--
		case ',':
			if n == 1 {
				r1, r2 := typ[len(prefix):i], typ[i+1:len(typ)-1]
				log.Printf("GML: typeinfo.go: B: GetMapSubtypes('%v') = ('%v', '%v')", typ, r1, r2)
				return r1, r2
			}
		}
	}

	log.Printf("GML: typeinfo.go: C: GetMapSubtypes('%v') = ('', '')", typ)
	return "", ""
}

func (lti *langTypeInfo) GetNameForType(typ string) string {
	// "github.com/gmlewis/modus/sdk/go/examples/simple.Person" -> "Person"

	if lti.IsPointerType(typ) { // TODO
		result := "*" + lti.GetNameForType(lti.GetUnderlyingType(typ))
		log.Printf("GML: typeinfo.go: A: GetNameForType('%v') = '%v'", typ, result)
		return result
	}

	if lti.IsListType(typ) {
		result := "[" + lti.GetNameForType(lti.GetListSubtype(typ)) + "]"
		log.Printf("GML: typeinfo.go: B: GetNameForType('%v') = '%v'", typ, result)
		return result
	}

	if lti.IsMapType(typ) {
		kt, vt := lti.GetMapSubtypes(typ)
		result := "Map[" + lti.GetNameForType(kt) + "," + lti.GetNameForType(vt) + "]"
		log.Printf("GML: typeinfo.go: C: GetNameForType('%v') = '%v'", typ, result)
		return result
	}

	result := typ[strings.LastIndex(typ, ".")+1:]
	log.Printf("GML: typeinfo.go: D: GetNameForType('%v') = '%v'", typ, result)
	return result
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
	log.Printf("GML: typeinfo.go: ENTER ArrayLength('%v')", typ)
	i := strings.Index(typ, "]")
	if i == -1 {
		return -1, fmt.Errorf("invalid array type: %s", typ)
	}

	size := typ[1:i]
	if size == "" {
		return -1, fmt.Errorf("invalid array type: %s", typ)
	}

	parsedSize, err := strconv.Atoi(size)
	if err != nil {
		return -1, err
	}
	if parsedSize < 0 || parsedSize > math.MaxUint32 {
		return -1, fmt.Errorf("array size out of bounds: %s", size)
	}

	log.Printf("GML: typeinfo.go: ArrayLength('%v') = %v", typ, parsedSize)
	return parsedSize, nil
}

func (lti *langTypeInfo) getSizeOfArray(ctx context.Context, typ string) (uint32, error) {
	// array size is the element size times the number of elements, aligned to the element size
	arrSize, err := lti.ArrayLength(typ)
	if err != nil {
		return 0, err
	}
	if arrSize == 0 {
		return 0, nil
	}

	t := lti.GetListSubtype(typ)
	elementAlignment, err := lti.GetAlignmentOfType(ctx, t)
	if err != nil {
		return 0, err
	}
	elementSize, err := lti.GetSizeOfType(ctx, t)
	if err != nil {
		return 0, err
	}

	size := langsupport.AlignOffset(elementSize, elementAlignment)*uint32(arrSize-1) + elementSize
	return size, nil
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
		offset += size
	}

	size := langsupport.AlignOffset(offset, maxAlign)
	return size, nil
}

func (lti *langTypeInfo) GetAlignmentOfType(ctx context.Context, typ string) (uint32, error) {

	// reference: https://github.com/tinygo-org/tinygo/blob/release/compiler/sizes.go

	// primitives align to their natural size
	if lti.IsPrimitiveType(typ) {
		result, err := lti.GetSizeOfType(ctx, typ)
		log.Printf("GML: typeinfo.go: A: GetAlignmentOfType('%v') = %v, err=%v", typ, result, err)
		return result, err
	}

	// arrays align to the alignment of their element type
	if lti.IsArrayType(typ) {
		t := lti.GetListSubtype(typ)
		result, err := lti.GetAlignmentOfType(ctx, t)
		log.Printf("GML: typeinfo.go: B: GetAlignmentOfType('%v') = %v, err=%v", typ, result, err)
		return result, err
	}

	// reference types align to the pointer size (4 bytes on 32-bit wasm)
	if lti.IsPointerType(typ) || lti.IsSliceType(typ) || lti.IsStringType(typ) || lti.IsMapType(typ) {
		log.Printf("GML: typeinfo.go: C: GetAlignmentOfType('%v') = 4", typ)
		return 4, nil
	}

	// time.Time has 3 fields, the maximum alignment is 8 bytes
	if lti.IsTimestampType(typ) {
		log.Printf("GML: typeinfo.go: D: GetAlignmentOfType('%v') = 8", typ)
		return 8, nil
	}

	// structs align to the maximum alignment of their fields
	result, err := lti.getAlignmentOfStruct(ctx, typ)
	log.Printf("GML: typeinfo.go: E: GetAlignmentOfType('%v') = %v, err=%v", typ, result, err)
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

	log.Printf("GML: typeinfo.go: getAlignmentOfStruct('%v') = %v", typ, max)
	return max, nil
}

func (lti *langTypeInfo) GetDataSizeOfType(ctx context.Context, typ string) (uint32, error) {
	return lti.GetSizeOfType(ctx, typ)
}

func (lti *langTypeInfo) GetEncodingLengthOfType(ctx context.Context, typ string) (uint32, error) {
	if lti.IsPrimitiveType(typ) || lti.IsPointerType(typ) || lti.IsMapType(typ) {
		log.Printf("GML: typeinfo.go: A: GetEncodingLengthOfType('%v') = 1", typ)
		return 1, nil
	} else if lti.IsStringType(typ) {
		log.Printf("GML: typeinfo.go: B: GetEncodingLengthOfType('%v') = 2", typ)
		return 2, nil
	} else if lti.IsSliceType(typ) || lti.IsTimestampType(typ) {
		log.Printf("GML: typeinfo.go: C: GetEncodingLengthOfType('%v') = 3", typ)
		return 3, nil
	} else if lti.IsArrayType(typ) {
		result, err := lti.getEncodingLengthOfArray(ctx, typ)
		log.Printf("GML: typeinfo.go: D: GetEncodingLengthOfType('%v') = %v, err=%v", typ, result, err)
		return result, err
	} else if lti.IsObjectType(typ) {
		result, err := lti.getEncodingLengthOfStruct(ctx, typ)
		log.Printf("GML: typeinfo.go: E: GetEncodingLengthOfType('%v') = %v, err=%v", typ, result, err)
		return result, err
	}

	return 0, fmt.Errorf("unable to determine encoding length for type: %s", typ)
}

func (lti *langTypeInfo) getEncodingLengthOfArray(ctx context.Context, typ string) (uint32, error) {
	arrSize, err := lti.ArrayLength(typ)
	if err != nil {
		return 0, err
	}
	if arrSize == 0 {
		log.Printf("GML: typeinfo.go: A: getEncodingLengthOfArray('%v') = 0", typ)
		return 0, nil
	}

	t := lti.GetListSubtype(typ)
	elementLen, err := lti.GetEncodingLengthOfType(ctx, t)
	if err != nil {
		return 0, err
	}

	result := uint32(arrSize) * elementLen
	log.Printf("GML: typeinfo.go: B: getEncodingLengthOfArray('%v') = %v", typ, result)
	return result, nil
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

	log.Printf("GML: typeinfo.go: getEncodingLengthOfStruct('%v') = %v", typ, total)
	return total, nil
}

func (lti *langTypeInfo) GetSizeOfType(ctx context.Context, typ string) (uint32, error) {
	switch typ {
	case "int8", "uint8", "bool", "byte":
		log.Printf("GML: typeinfo.go: A: GetSizeOfType('%v') = 1", typ)
		return 1, nil
	case "int16", "uint16":
		log.Printf("GML: typeinfo.go: B: GetSizeOfType('%v') = 2", typ)
		return 2, nil
	case "int32", "uint32", "float32", "rune",
		"int", "uint", "uintptr", "unsafe.Pointer": // we only support 32-bit wasm
		log.Printf("GML: typeinfo.go: C: GetSizeOfType('%v') = 4", typ)
		return 4, nil
	case "int64", "uint64", "float64", "time.Duration":
		log.Printf("GML: typeinfo.go: D: GetSizeOfType('%v') = 8", typ)
		return 8, nil
	}

	if lti.IsStringType(typ) {
		// string header is a 4 byte pointer and 4 byte length
		log.Printf("GML: typeinfo.go: E: GetSizeOfType('%v') = 8", typ)
		return 8, nil
	}

	if lti.IsPointerType(typ) {
		log.Printf("GML: typeinfo.go: F: GetSizeOfType('%v') = 4", typ)
		return 4, nil
	}

	if lti.IsMapType(typ) {
		// maps are passed by reference using a 4 byte pointer
		log.Printf("GML: typeinfo.go: G: GetSizeOfType('%v') = 4", typ)
		return 4, nil
	}

	if lti.IsSliceType(typ) {
		// slice header is a 4 byte pointer, 4 byte length, 4 byte capacity
		log.Printf("GML: typeinfo.go: H: GetSizeOfType('%v') = 12", typ)
		return 12, nil
	}

	if lti.IsTimestampType(typ) {
		// time.Time has 3 fields: 8 byte uint64, 8 byte int64, 4 byte pointer
		log.Printf("GML: typeinfo.go: I: GetSizeOfType('%v') = 20", typ)
		return 20, nil
	}

	if lti.IsArrayType(typ) {
		result, err := lti.getSizeOfArray(ctx, typ)
		log.Printf("GML: typeinfo.go: J: GetSizeOfType('%v') = %v", typ, result)
		return result, err
	}

	result, err := lti.getSizeOfStruct(ctx, typ)
	log.Printf("GML: typeinfo.go: K: GetSizeOfType('%v') = %v", typ, result)
	return result, err
}

func (lti *langTypeInfo) GetTypeDefinition(ctx context.Context, typ string) (*metadata.TypeDefinition, error) {
	md := ctx.Value(utils.MetadataContextKey).(*metadata.Metadata)
	result, err := md.GetTypeDefinition(typ)
	log.Printf("GML: typeinfo.go: GetTypeDefinition('%v') = %v, err=%v", typ, result, err)
	return result, err
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
