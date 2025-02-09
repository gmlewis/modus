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
	"go/ast"
	"regexp"
	"strings"
)

var (
	mbtCommentRE = regexp.MustCompile(`^/*\|?\s*`)
)

func getDocsForFunction(fnSignature, fullSrc string) *ast.CommentGroup {
	// gmlPrintf("GML: packages/parse-docs.go: getDocsForFunction: fnSignature=%v", fnSignature)
	fnIndex := strings.Index(fullSrc, fnSignature)
	if fnIndex < 0 {
		m := pubFnPrefixRE.FindStringSubmatch(fnSignature)
		if len(m) == 0 {
			gmlPrintf("PROGRAMMING ERROR: len(m) == 0 for fnSignature=%v", fnSignature)
			return nil
		}
		fnSignaturePrefix := m[0]
		fnIndex = strings.Index(fullSrc, fnSignaturePrefix)
		if fnIndex < 0 {
			gmlPrintf("PROGRAMMING ERROR: fnIndex < 0 for fnSignature=%v", fnSignature)
			return nil
		}
	}
	doubleBlankIndex := strings.LastIndex(fullSrc[:fnIndex], "\n\n")
	if doubleBlankIndex < 0 {
		gmlPrintf("PROGRAMMING ERROR: doubleBlankIndex < 0 for fnSignature=%v", fnSignature)
	}
	fnSrcLines := strings.Split(fullSrc[doubleBlankIndex+2:fnIndex], "\n")
	// gmlPrintf("GML: packages/parse-docs.go: getDocsForFunction: fnSrcLines=\n%v", fnSrcLines)
	docs := &ast.CommentGroup{}
	stripLeadingBlankLines := true
	for _, line := range fnSrcLines {
		if !strings.HasPrefix(line, "///") { // only process triple-slash comments for package docs
			continue
		}
		line = strings.TrimSpace(mbtCommentRE.ReplaceAllString(line, ""))
		if line == "" && stripLeadingBlankLines {
			continue
		}
		stripLeadingBlankLines = false
		docs.List = append(docs.List, &ast.Comment{Text: "// " + line}) // Conform to Go-style docs for later processing.
	}
	return docs
}
