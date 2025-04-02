/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func StripDefaultValue(typeSignature string) string {
	parts := strings.Split(typeSignature, "=")
	return strings.TrimSpace(parts[0])
}

func StripError(typeSignature string) (string, bool) {
	if i := strings.Index(typeSignature, "!"); i >= 0 {
		return typeSignature[:i], true
	}
	return typeSignature, false
}

func StripErrorAndOption(typeSignature string) (typ string, hasError, hasOption bool) {
	if i := strings.Index(typeSignature, "!"); i >= 0 {
		hasError = true
		typeSignature = typeSignature[:i]
	}
	hasOption = strings.HasSuffix(typeSignature, "?")
	return strings.TrimSuffix(typeSignature, "?"), hasError, hasOption
}

// TODO: This needs to be kept in sync with runtime/languages/moonbit/typeinfo.go GetNameForType()
func GetNameForType(t string, imports map[string]string) string {
	t, _ = StripError(t)

	sep := strings.LastIndex(t, ".")
	if sep == -1 {
		return t
	}

	if IsOptionType(t) {
		result := GetNameForType(GetUnderlyingType(t), imports) + "?"
		return result
	}

	if IsListType(t) {
		switch {
		case t == "Bytes":
			return t
		case strings.HasPrefix(t, "Array["):
			return "Array[" + GetNameForType(GetListSubtype(t), imports) + "]"
		case strings.HasPrefix(t, "ArrayView["):
			return "ArrayView[" + GetNameForType(GetListSubtype(t), imports) + "]"
		case strings.HasPrefix(t, "FixedArray["):
			return "FixedArray[" + GetNameForType(GetListSubtype(t), imports) + "]"
		default:
			log.Fatalf("PROGRAMMING ERROR: utils/typeinfo.go: GetNameForType('%v'): Bad list type!", t)
		}
	}

	if IsMapType(t) {
		kt, vt := GetMapSubtypes(t)
		keyTypeName := GetNameForType(kt, imports)
		valueTypeName := GetNameForType(vt, imports)
		return "Map[" + keyTypeName + "," + valueTypeName + "]"
	}

	pkgPath := t[:sep]
	pkgName := imports[pkgPath]
	typeName := t[sep+1:]
	if pkgName == "" {
		return typeName
	}

	return pkgName + "." + typeName
}

func GetPackageNamesForType(t string) []string {
	var hasOption bool
	t, _, hasOption = StripErrorAndOption(t)

	if hasOption {
		return GetPackageNamesForType(GetUnderlyingType(t))
	}

	// if IsPointerType(t) {
	// 	return GetPackageNamesForType(GetUnderlyingType(t))
	// }

	if IsListType(t) {
		return GetPackageNamesForType(GetListSubtype(t))
	}

	if IsMapType(t) {
		kt, vt := GetMapSubtypes(t)
		kp, vp := GetPackageNamesForType(kt), GetPackageNamesForType(vt)
		m := make(map[string]bool, len(kp)+len(vp))
		for _, p := range kp {
			m[p] = true
		}
		for _, p := range vp {
			m[p] = true
		}
		pkgs := MapKeys(m)
		sort.Strings(pkgs)
		return pkgs
	}

	// This is needed to find the package names of all fully-qualified struct types.
	if i := strings.LastIndex(t, "."); i != -1 { // "@..Type" => "@."
		return []string{t[:i]}
	}

	return nil
}

// func GetArraySubtype(t string) string {
// 	// TODO: FixedArray[] are fixed length, Array[] are dynamic
// 	// Go: return t[strings.Index(t, "]")+1:]
// 	return strings.TrimSuffix(strings.TrimPrefix(t, "Array["), "]")
// }

func GetListSubtype(typ string) string {
	typ, _ = StripError(typ)

	if !strings.HasSuffix(typ, "]") && !strings.HasSuffix(typ, "]?") {
		return ""
	}
	typ = strings.TrimSuffix(typ, "?")
	typ = strings.TrimSuffix(typ, "]")

	switch {
	case typ == "Bytes":
		return "Byte"
	case strings.HasPrefix(typ, "Array["):
		return strings.TrimSpace(strings.TrimPrefix(typ, "Array["))
	case strings.HasPrefix(typ, "ArrayView["):
		return strings.TrimSpace(strings.TrimPrefix(typ, "ArrayView["))
	case strings.HasPrefix(typ, "FixedArray["):
		return strings.TrimSpace(strings.TrimPrefix(typ, "FixedArray["))
	default:
		return ""
	}
}

func GetMapSubtypes(typ string) (string, string) {
	typ, _ = StripError(typ)

	if !strings.HasSuffix(typ, "]") && !strings.HasSuffix(typ, "]?") {
		return "", ""
	}

	const prefix = "Map[" // e.g. Map[String, Int]
	if !strings.HasPrefix(typ, prefix) {
		return "", ""
	}
	typ = strings.TrimSuffix(typ, "?")
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

func GetUnderlyingType(t string) string {
	t, _ = StripError(t)

	return strings.TrimSuffix(t, "?")
}

func IsListType(typ string) bool {
	typ, _ = StripError(typ)

	if typ == "Bytes" {
		return true
	}

	if !strings.HasSuffix(typ, "]") && !strings.HasSuffix(typ, "]?") {
		return false
	}
	return strings.HasPrefix(typ, "Array[") ||
		strings.HasPrefix(typ, "ArrayView[") ||
		strings.HasPrefix(typ, "FixedArray[")
}

func IsSliceType(typ string) bool {
	typ, _ = StripError(typ)

	if typ == "Bytes" {
		return true
	}

	if !strings.HasSuffix(typ, "]") && !strings.HasSuffix(typ, "]?") {
		return false
	}
	// MoonBit Arrays and FixedArrays are similar to Go slices.
	return strings.HasPrefix(typ, "Array[") ||
		strings.HasPrefix(typ, "ArrayView[") ||
		strings.HasPrefix(typ, "FixedArray[")
}

func IsArrayType(t string) bool {
	return false // no fixed-length array types in MoonBit
}

func IsMapType(typ string) bool {
	typ, _ = StripError(typ)

	if !strings.HasSuffix(typ, "]") && !strings.HasSuffix(typ, "]?") {
		return false
	}
	return strings.HasPrefix(typ, "Map[")
}

func IsOptionType(typ string) bool {
	typ, _ = StripError(typ)

	return strings.HasSuffix(typ, "?")
}

func IsStringType(typ string) bool {
	typ, _ = StripError(typ)

	return strings.HasPrefix(typ, "String")
}

func IsStructType(t string) bool {
	// special cases not covered elsewhere:
	switch {
	case strings.HasPrefix(t, "Iter["),
		strings.HasPrefix(t, "Iter2["),
		strings.HasPrefix(t, "Result["):
		return false
	}

	return !IsVoidType(t) && !IsOptionType(t) && !IsPrimitiveType(t) && !IsListType(t) && !IsMapType(t) && !IsStringType(t)
}

func IsPrimitiveType(typ string) bool {
	typ, _ = StripError(typ)

	switch typ {
	case "Bool",
		"Int", "Int16", "Int64",
		"UInt", "UInt16", "UInt64",
		"Byte", "Char",
		"Float", "Double",
		"@time.Duration",
		"Json":
		return true
	}

	return false
}

func IsVoidType(typ string) bool {
	return strings.HasPrefix(typ, "Unit")
}

func IsTupleType(typ string) bool {
	typ, _, _ = StripErrorAndOption(typ)
	return strings.HasPrefix(typ, "(") && strings.HasSuffix(typ, ")")
}

// IsWellKnownType is a hack to avoid listing types in external packages
// as custom types in that package.
func IsWellKnownType(typ string) bool {
	typ, _, _ = StripErrorAndOption(typ)
	if i := strings.Index(typ, "["); i >= 0 {
		typ = typ[:i+1] // include the '['
	}
	return wellKnownTypes[typ]
}

var wellKnownTypes = map[string]bool{
	"Array[":      true,
	"ArrayView[":  true,
	"Bool":        true,
	"Byte":        true,
	"Bytes":       true,
	"Char":        true,
	"Double":      true,
	"FixedArray[": true,
	"Float":       true,
	"Json":        true,
	"Int":         true,
	"Int16":       true,
	"Int64":       true,
	"Iter[":       true,
	"Iter2[":      true,
	"Map[":        true,
	"Result[":     true,
	"String":      true,
	"UInt":        true,
	"UInt16":      true,
	"UInt64":      true,
	"Unit":        true,
}

// FullyQualifyTypeName is a hack to attempt to convert a local type name
// to a fully-qualifyed type name (e.g. "Array[Record]" => "Array[@neo4j.Record]")
// by attempting to naively parse the MoonBit type and prefix any unknown
// type names with the package name.
// `accumulator` recursively builds up a map of all type names found
// (including `fqTypeName`).
func FullyQualifyTypeName(pkgName, typeName string) (fqTypeName string, accumulator map[string]struct{}) {
	accumulator = map[string]struct{}{}
	fqTypeName = fullyQualifyTypeNameWithAcc(pkgName, typeName, accumulator)
	return fqTypeName, accumulator
}

func fullyQualifyTypeNameWithAcc(pkgName, typeName string, accumulator map[string]struct{}) string {
	eqSignIndex := strings.Index(typeName, "=")
	if pkgName == "" {
		if typeName != "Unit" {
			typ := typeName
			// Don't include default values in the accumulator.
			if eqSignIndex >= 0 {
				typ = strings.TrimSpace(typeName[:eqSignIndex])
			}
			accumulator[typ] = struct{}{}
		}
		return typeName // no package to fully qualify
	}

	var defaultValue string
	if eqSignIndex >= 0 {
		defaultValue = strings.TrimSpace(typeName[eqSignIndex+1:])
		typeName = strings.TrimSpace(typeName[:eqSignIndex])
	}

	typ, hasError, hasOption := StripErrorAndOption(typeName)

	bracketIdx := strings.Index(typ, "[")
	if bracketIdx < 0 {
		if wellKnownTypes[typ] {
			if typeName != "Unit" {
				accumulator[typeName] = struct{}{}
			}
			return typeName // known type
		}

		if strings.HasPrefix(typ, "(") && strings.HasSuffix(typ, ")") {
			// Tuple type, recurse on each type.
			fields := SplitParamsWithBrackets(typ[1 : len(typ)-1])
			for i, f := range fields {
				fields[i] = fullyQualifyTypeNameWithAcc(pkgName, f, accumulator)
			}
			result := "(" + strings.Join(fields, ", ") + ")"
			accumulator[result] = struct{}{}
			return result
		}

		if strings.Contains(typ, ".") { // already fully qualified
			if typeName != "Unit" {
				accumulator[typeName] = struct{}{}
			}
			return typeName
		}

		result := pkgName + "." + typeName // unknown type.
		accumulator[result] = struct{}{}
		return result
	}

	restoreDefaultValue := func(s string) string {
		if defaultValue != "" {
			result := fmt.Sprintf("%v = %v", s, defaultValue)
			accumulator[s] = struct{}{}
			return result
		}
		if s != "Unit" {
			accumulator[s] = struct{}{}
		}
		return s
	}
	restoreSuffix := func(s string) string {
		if !hasOption && !hasError {
			return restoreDefaultValue(s)
		}
		if hasOption {
			s += "?"
		}
		if !hasError {
			return restoreDefaultValue(s)
		}
		i := strings.Index(typeName, "!")
		if i < 0 {
			log.Fatalf("PROGRAMMING ERROR: FullyQualifyTypeName not handling error for typeName='%v'", typeName)
		}
		return restoreDefaultValue(s + typeName[i:])
	}

	// type has bracket, recurse.
	if typ[len(typ)-1] != ']' {
		log.Fatalf("PROGRAMMING ERROR: FullyQualifyTypeName failed to parse typeName='%v'", typeName)
	}
	prefix := typ[:bracketIdx+1]
	inner := typ[bracketIdx+1 : len(typ)-1] // remove trailing ']'
	switch prefix {
	// single inner type
	case "Array[",
		"ArrayView[",
		"FixedArray[",
		"Iter[":
		fqInner := fullyQualifyTypeNameWithAcc(pkgName, inner, accumulator)
		return restoreSuffix(prefix + fqInner + "]")
		// double inner type
	case "Iter2[",
		"Map[",
		"Result[":
		innerFields := SplitParamsWithBrackets(inner)
		if len(innerFields) != 2 {
			log.Fatalf("PROGRAMMING ERROR: FullyQualifyTypeName unhandled split of typeName='%v'", typeName)
		}
		fqInner1 := fullyQualifyTypeNameWithAcc(pkgName, innerFields[0], accumulator)
		fqInner2 := fullyQualifyTypeNameWithAcc(pkgName, innerFields[1], accumulator)
		if prefix == "Map[" {
			accumulator[fmt.Sprintf("Array[%v]", fqInner1)] = struct{}{}
			accumulator[fmt.Sprintf("Array[%v]", fqInner2)] = struct{}{}
		}
		return restoreSuffix(fmt.Sprintf("%v%v, %v]", prefix, fqInner1, fqInner2))
	default:
		// This is possibly using generics, e.g. "CallStack[T]". Just return it prefixed with pkgName.
		result := pkgName + "." + typeName // unknown type.
		accumulator[result] = struct{}{}
		return result
	}
}
