// This file is based on: https://cs.opensource.google/go/x/tools/+/refs/tags/v0.28.0:go/packages/packages.go
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packages

import (
	"go/ast"
	"log"
	"regexp"
	"strings"
)

var (
	mbtCommentRE = regexp.MustCompile(`^/*\|?\s*`)
)

func getDocsForFunction(fnSignature, fullSrc string) *ast.CommentGroup {
	// log.Printf("GML: packages/parse-docs.go: getDocsForFunction: fnSignature=%v", fnSignature)
	fnIndex := strings.Index(fullSrc, fnSignature)
	if fnIndex < 0 {
		m := pubFnPrefixRE.FindStringSubmatch(fnSignature)
		if len(m) == 0 {
			log.Printf("PROGRAMMING ERROR: len(m) == 0 for fnSignature=%v", fnSignature)
			return nil
		}
		fnSignaturePrefix := m[0]
		fnIndex = strings.Index(fullSrc, fnSignaturePrefix)
		if fnIndex < 0 {
			log.Printf("PROGRAMMING ERROR: fnIndex < 0 for fnSignature=%v", fnSignature)
			return nil
		}
	}
	doubleBlankIndex := strings.LastIndex(fullSrc[:fnIndex], "\n\n")
	if doubleBlankIndex < 0 {
		log.Printf("PROGRAMMING ERROR: doubleBlankIndex < 0 for fnSignature=%v", fnSignature)
	}
	fnSrcLines := strings.Split(fullSrc[doubleBlankIndex+2:fnIndex], "\n")
	// log.Printf("GML: packages/parse-docs.go: getDocsForFunction: fnSrcLines=\n%v", fnSrcLines)
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
