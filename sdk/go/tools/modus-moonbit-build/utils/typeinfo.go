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
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

// TODO: Remove debugging
var gmlDebugEnv bool

func gmlPrintf(fmtStr string, args ...any) {
	sync.OnceFunc(func() {
		log.SetFlags(0)
		if os.Getenv("GML_DEBUG") == "true" {
			gmlDebugEnv = true
		}
	})
	if gmlDebugEnv {
		log.Printf(fmtStr, args...)
	}
}

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
	gmlPrintf("GML: utils/typeinfo.go: GetNameForType('%v')", t)
	t, _ = StripError(t)

	sep := strings.LastIndex(t, ".")
	if sep == -1 {
		gmlPrintf("GML: utils/typeinfo.go: GetNameForType: A: = '%v'", t)
		return t
	}

	if IsOptionType(t) {
		result := GetNameForType(GetUnderlyingType(t), imports) + "?"
		gmlPrintf("GML: utils/typeinfo.go: GetNameForType: B: t: '%v' = '%v'", t, result)
		return result
	}

	// if IsPointerType(t) {
	// 	return "*" + GetNameForType(GetUnderlyingType(t), imports) // TODO
	// }

	if IsListType(t) {
		switch {
		case strings.HasPrefix(t, "Array["):
			result := "Array[" + GetNameForType(GetListSubtype(t), imports) + "]"
			gmlPrintf("GML: utils/typeinfo.go: A: GetNameForType: C: t: '%v' = '%v'", t, result)
			return result
		case strings.HasPrefix(t, "FixedArray["):
			result := "FixedArray[" + GetNameForType(GetListSubtype(t), imports) + "]"
			gmlPrintf("GML: utils/typeinfo.go: B: GetNameForType: D: t: '%v' = '%v'", t, result)
			return result
		default:
			gmlPrintf("PROGRAMMING ERROR: utils/typeinfo.go: GetNameForType('%v'): Bad list type!", t)
		}
	}

	if IsMapType(t) {
		kt, vt := GetMapSubtypes(t)
		keyTypeName := GetNameForType(kt, imports)
		valueTypeName := GetNameForType(vt, imports)
		result := "Map[" + keyTypeName + "," + valueTypeName + "]"
		gmlPrintf("GML: utils/typeinfo.go: C: GetNameForType: E: t: '%v' = '%v'", t, result)
		return result
	}

	pkgPath := t[:sep]
	pkgName := imports[pkgPath]
	typeName := t[sep+1:]
	if pkgName == "" {
		gmlPrintf("GML: utils/typeinfo.go: D: GetNameForType: F: t: '%v' = '%v'", t, typeName)
		return typeName
	}
	result := pkgName + "." + typeName
	gmlPrintf("GML: utils/typeinfo.go: E: GetNameForType: G: t: '%v' = '%v'", t, result)
	return result
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
		gmlPrintf("ERROR: utils/typeinfo.go: GetListSubtype('%v'): Bad list type!", typ)
		return ""
	}
	typ = strings.TrimSuffix(typ, "?")
	typ = strings.TrimSuffix(typ, "]")

	switch {
	case strings.HasPrefix(typ, "Array["):
		result := strings.TrimSpace(strings.TrimPrefix(typ, "Array["))
		gmlPrintf("GML: utils/typeinfo.go: GetListSubtype('%v') = '%v'", typ, result)
		return result
	case strings.HasPrefix(typ, "FixedArray["):
		result := strings.TrimSpace(strings.TrimPrefix(typ, "FixedArray["))
		gmlPrintf("GML: utils/typeinfo.go: GetListSubtype('%v') = '%v'", typ, result)
		return result
	default:
		gmlPrintf("ERROR: utils/typeinfo.go: GetListSubtype('%v'): Bad list type!", typ)
		return ""
	}
}

func GetMapSubtypes(typ string) (string, string) {
	typ, _ = StripError(typ)

	if !strings.HasSuffix(typ, "]") && !strings.HasSuffix(typ, "]?") {
		gmlPrintf("ERROR: utils/typeinfo.go: GetMapSubtypes('%v'): Bad map type!", typ)
		return "", ""
	}

	const prefix = "Map[" // e.g. Map[String, Int]
	if !strings.HasPrefix(typ, prefix) {
		gmlPrintf("GML: utils/typeinfo.go: A: GetMapSubtypes('%v') = ('', '')", typ)
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
				r1, r2 := strings.TrimSpace(typ[:i]), strings.TrimSpace(typ[i+1:])
				gmlPrintf("GML: utils/typeinfo.go: B: GetMapSubtypes('%v') = ('%v', '%v')", typ, r1, r2)
				return r1, r2
			}
		}
	}

	gmlPrintf("GML: utils/typeinfo.go: C: GetMapSubtypes('%v') = ('', '')", typ)
	return "", ""
}

func GetUnderlyingType(t string) string {
	t, _ = StripError(t)

	return strings.TrimSuffix(t, "?")
}

func IsListType(typ string) bool {
	typ, _ = StripError(typ)

	if !strings.HasSuffix(typ, "]") && !strings.HasSuffix(typ, "]?") {
		return false
	}
	result := strings.HasPrefix(typ, "Array[") || strings.HasPrefix(typ, "FixedArray[")
	gmlPrintf("GML: utils/typeinfo.go: IsListType('%v') = %v", typ, result)
	return result
}

func IsSliceType(typ string) bool {
	typ, _ = StripError(typ)

	if !strings.HasSuffix(typ, "]") && !strings.HasSuffix(typ, "]?") {
		return false
	}
	// MoonBit Arrays and FixedArrays are similar to Go slices.
	result := strings.HasPrefix(typ, "Array[") || strings.HasPrefix(typ, "FixedArray[")
	gmlPrintf("GML: utils/typeinfo.go: IsSliceType('%v') = %v", typ, result)
	return result
}

func IsArrayType(t string) bool {
	return false // no fixed-length array types in MoonBit
}

func IsMapType(typ string) bool {
	typ, _ = StripError(typ)

	if !strings.HasSuffix(typ, "]") && !strings.HasSuffix(typ, "]?") {
		return false
	}
	result := strings.HasPrefix(typ, "Map[")
	gmlPrintf("GML: utils/typeinfo.go: IsMapType('%v') = %v", typ, result)
	return result
}

func IsOptionType(typ string) bool {
	typ, _ = StripError(typ)

	result := strings.HasSuffix(typ, "?")
	gmlPrintf("GML: utils/typeinfo.go: IsOptionType('%v') = %v", typ, result)
	return result
}

// func IsPointerType(t string) bool { // TODO
// 	return strings.HasPrefix(t, "*")
// }

func IsStringType(typ string) bool {
	typ, _ = StripError(typ)

	result := strings.HasPrefix(typ, "String")
	gmlPrintf("GML: utils/typeinfo.go: IsStringType('%v') = %v", typ, result)
	return result
}

func IsStructType(t string) bool {
	// return !IsPointerType(t) && !IsPrimitiveType(t) && !IsListType(t) && !IsMapType(t) && !IsStringType(t)
	result := !IsVoidType(t) && !IsOptionType(t) && !IsPrimitiveType(t) && !IsListType(t) && !IsMapType(t) && !IsStringType(t)
	gmlPrintf("GML: utils/typeinfo.go: IsStructType('%v') = %v", t, result)
	return result
}

func IsPrimitiveType(typ string) bool {
	typ, _ = StripError(typ)

	switch typ {
	case "Bool",
		"Int", "Int16", "Int64",
		"UInt", "UInt16", "UInt64",
		"Byte", "Char",
		"Float", "Double",
		"@time.Duration":
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
