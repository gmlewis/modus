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
	"sort"
	"strings"
)

func GetNameForType(t string, imports map[string]string) string {
	sep := strings.LastIndex(t, ".")
	if sep == -1 {
		return t
	}

	if IsOptionType(t) {
		return GetNameForType(GetUnderlyingType(t), imports) + "?"
	}

	// if IsPointerType(t) {
	// 	return "*" + GetNameForType(GetUnderlyingType(t), imports) // TODO
	// }

	if IsListType(t) {
		return "Array[" + GetNameForType(GetArraySubtype(t), imports) + "]"
	}

	if IsMapType(t) {
		kt, vt := GetMapSubtypes(t)
		return "Map[" + GetNameForType(kt, imports) + "," + GetNameForType(vt, imports) + "]"
	}

	pkgPath := t[:sep]
	pkgName := imports[pkgPath]
	typeName := t[sep+1:]
	if pkgName == "" {
		return typeName
	} else {
		return pkgName + "." + typeName
	}
}

func GetPackageNamesForType(t string) []string {
	if IsOptionType(t) {
		return GetPackageNamesForType(GetUnderlyingType(t))
	}

	// if IsPointerType(t) {
	// 	return GetPackageNamesForType(GetUnderlyingType(t))
	// }

	if IsListType(t) {
		return GetPackageNamesForType(GetArraySubtype(t))
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

	if i := strings.LastIndex(t, "."); i != -1 {
		return []string{t[:i]}
	}

	return nil
}

func GetArraySubtype(t string) string {
	// TODO: FixedArray[] are fixed length, Array[] are dynamic
	// Go: return t[strings.Index(t, "]")+1:]
	return strings.TrimSuffix(strings.TrimPrefix(t, "Array["), "]")
}

func GetMapSubtypes(t string) (string, string) {
	const prefix = "Map["
	if !strings.HasPrefix(t, prefix) {
		return "", ""
	}
	t = strings.TrimSuffix(t, "?")

	n := 1
	for i := len(prefix); i < len(t); i++ {
		switch t[i] {
		case '[':
			n++
		case ',':
			if n == 1 {
				return strings.TrimSpace(t[len(prefix):i]), strings.TrimSpace(t[i+1 : len(t)-1])
			}
		case ']':
			n--
		}
	}

	return "", ""
}

func GetUnderlyingType(t string) string {
	return strings.TrimSuffix(t, "?")
}

func IsListType(t string) bool {
	// Go: covers both slices and arrays
	// Go: return strings.HasPrefix(t, "[")
	return strings.HasPrefix(t, "Array[") && strings.HasSuffix(t, "]") // TODO
}

func IsSliceType(t string) bool {
	return strings.HasPrefix(t, "Array[") && strings.HasSuffix(t, "]")
}

func IsArrayType(t string) bool {
	// return IsListType(t) && !IsSliceType(t)
	return strings.HasPrefix(t, "FixedArray[") && strings.HasSuffix(t, "]")
}

func IsMapType(t string) bool {
	return strings.HasPrefix(t, "Map[") && strings.HasSuffix(t, "]")
}

func IsOptionType(t string) bool {
	return strings.HasSuffix(t, "?")
}

// func IsPointerType(t string) bool { // TODO
// 	return strings.HasPrefix(t, "*")
// }

func IsStringType(t string) bool {
	return t == "String"
}

func IsStructType(t string) bool {
	// return !IsPointerType(t) && !IsPrimitiveType(t) && !IsListType(t) && !IsMapType(t) && !IsStringType(t)
	return !IsOptionType(t) && !IsPrimitiveType(t) && !IsListType(t) && !IsMapType(t) && !IsStringType(t)
}

func IsPrimitiveType(t string) bool {
	switch t {
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
