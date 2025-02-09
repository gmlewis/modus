/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package languages

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/languages/assemblyscript"
	"github.com/gmlewis/modus/runtime/languages/golang"
	"github.com/gmlewis/modus/runtime/languages/moonbit"
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

var lang_AssemblyScript = langsupport.NewLanguage(
	"AssemblyScript",
	assemblyscript.LanguageTypeInfo(),
	assemblyscript.NewPlanner,
	assemblyscript.NewWasmAdapter,
)

var lang_Go = langsupport.NewLanguage(
	"Go",
	golang.LanguageTypeInfo(),
	golang.NewPlanner,
	golang.NewWasmAdapter,
)

var lang_MoonBit = langsupport.NewLanguage(
	"MoonBit",
	moonbit.LanguageTypeInfo(),
	moonbit.NewPlanner,
	moonbit.NewWasmAdapter,
)

func AssemblyScript() langsupport.Language {
	return lang_AssemblyScript
}

func GoLang() langsupport.Language {
	return lang_Go
}

func MoonBit() langsupport.Language {
	return lang_MoonBit
}

func GetLanguageForSDK(sdk string) (langsupport.Language, error) {
	// strip version if present
	sdkName := sdk
	if i := strings.Index(sdkName, "@"); i != -1 {
		sdkName = sdkName[:i]
	}

	gmlPrintf("GML: languages.go: GetLanguageForSDK(sdk=%q), sdkName=%q", sdk, sdkName)
	// each SDK has a corresponding language implementation
	switch sdkName {
	case "modus-sdk-as":
		return AssemblyScript(), nil
	case "modus-sdk-go":
		return GoLang(), nil
	case "modus-sdk-mbt":
		return MoonBit(), nil
	}

	return nil, fmt.Errorf("unsupported SDK: %s", sdk)
}
