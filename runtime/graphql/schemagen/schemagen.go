/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package schemagen

import (
	"bytes"
	"cmp"
	"context"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
	"sync"

	"github.com/gmlewis/modus/lib/metadata"
	"github.com/gmlewis/modus/runtime/langsupport"
	"github.com/gmlewis/modus/runtime/languages"
	"github.com/gmlewis/modus/runtime/utils"
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

type GraphQLSchema struct {
	Schema            string
	FieldsToFunctions map[string]string
	MapTypes          []string
	TupleTypes        []string
}

func GetGraphQLSchema(ctx context.Context, md *metadata.Metadata) (*GraphQLSchema, error) {
	span, _ := utils.NewSentrySpanForCurrentFunc(ctx)
	defer span.Finish()

	lang, err := languages.GetLanguageForSDK(md.SDK)
	if err != nil {
		return nil, err
	}

	lti := lang.TypeInfo()
	inputTypeDefs, errors := transformTypes(md.Types, lti, true)
	resultTypeDefs, errs := transformTypes(md.Types, lti, false)
	errors = append(errors, errs...)
	root, errs := transformFunctions(md.FnExports, inputTypeDefs, resultTypeDefs, lti)
	errors = append(errors, errs...)

	if len(errors) > 0 {
		return nil, fmt.Errorf("failed to generate schema: %+v", errors)
	}

	allFields := root.AllFields()
	scalarTypes := extractCustomScalarTypes(inputTypeDefs, resultTypeDefs)
	inputTypes := filterTypes(utils.MapValues(inputTypeDefs), allFields, true)
	resultTypes := filterTypes(utils.MapValues(resultTypeDefs), allFields, false)

	buf := bytes.Buffer{}
	writeSchema(&buf, root, scalarTypes, inputTypes, resultTypes)

	mapTypes := make([]string, 0, len(resultTypeDefs))
	tupleTypes := make([]string, 0, len(resultTypeDefs))
	for _, t := range resultTypeDefs {
		if t.IsMapType {
			mapTypes = append(mapTypes, t.Name)
		}
		if t.IsTupleType {
			tupleTypes = append(tupleTypes, t.Name)
		}
	}

	fieldsToFunctions := make(map[string]string, len(allFields))
	for _, f := range allFields {
		fieldsToFunctions[f.Name] = f.Function
	}

	return &GraphQLSchema{
		Schema:            buf.String(),
		FieldsToFunctions: fieldsToFunctions,
		MapTypes:          mapTypes,
		TupleTypes:        tupleTypes,
	}, nil
}

type TransformError struct {
	Source any
	Error  error
}

func (e *TransformError) String() string {
	return fmt.Sprintf("source: %+v, error: %v", e.Source, e.Error)
}

func transformTypes(types metadata.TypeMap, lti langsupport.LanguageTypeInfo, forInput bool) (map[string]*TypeDefinition, []*TransformError) {
	typeDefs := make(map[string]*TypeDefinition, len(types))
	errors := make([]*TransformError, 0)
	for _, t := range types {
		if lti.IsListType(t.Name) || lti.IsMapType(t.Name) || lti.IsTimestampType(t.Name) {
			continue
		}
		if lti.GetUnderlyingType(t.Name) != t.Name {
			continue
		}

		name := lti.GetNameForType(t.Name)
		if forInput {
			if len(t.Fields) > 0 && !strings.HasSuffix(name, "Input") {
				if _, found := types[t.Name+"Input"]; found {
					continue
				}
				name += "Input"
			}
		} else if _, found := types[strings.TrimSuffix(t.Name, "Input")]; !found {
			name = strings.TrimSuffix(name, "Input")
		}
		if _, ok := typeDefs[name]; ok {
			continue
		}

		fields, err := convertFields(t.Fields, lti, typeDefs, forInput)
		if err != nil {
			errors = append(errors, &TransformError{t, err})
			continue
		}

		typeDef := &TypeDefinition{
			Name:   name,
			Fields: fields,
		}

		if t.Docs != nil {
			typeDef.DocLines = t.Docs.Lines
		}

		typeDefs[name] = typeDef
	}
	return typeDefs, errors
}

type FieldDefinition struct {
	Name      string
	Type      string
	Arguments []*ArgumentDefinition
	Function  string
	DocLines  []string
}

type TypeDefinition struct {
	Name        string
	Fields      []*FieldDefinition
	IsMapType   bool
	IsTupleType bool
	DocLines    []string
}

type ArgumentDefinition struct {
	Name    string
	Type    string
	Default *any
}

type RootObjects struct {
	QueryFields    []*FieldDefinition
	MutationFields []*FieldDefinition
}

func (r *RootObjects) AllFields() []*FieldDefinition {
	return append(r.QueryFields, r.MutationFields...)
}

func transformFunctions(functions metadata.FunctionMap, inputTypeDefs, resultTypeDefs map[string]*TypeDefinition, lti langsupport.LanguageTypeInfo) (*RootObjects, []*TransformError) {
	queryFields := make([]*FieldDefinition, 0, len(functions))
	mutationFields := make([]*FieldDefinition, 0, len(functions))
	errors := make([]*TransformError, 0)
	filter := getFieldFilter()

	fnNames := utils.MapKeys(functions)
	sort.Strings(fnNames)
	for _, name := range fnNames {
		fn := functions[name]

		args, err := convertParameters(fn.Parameters, lti, inputTypeDefs)
		if err != nil {
			errors = append(errors, &TransformError{fn, err})
			continue
		}

		returnType, err := convertResults(fn.Results, lti, resultTypeDefs)
		if err != nil {
			errors = append(errors, &TransformError{fn, err})
			continue
		}

		fieldName := getFieldName(fn.Name)

		field := &FieldDefinition{
			Name:      fieldName,
			Arguments: args,
			Type:      returnType,
			Function:  fn.Name,
		}

		if fn.Docs != nil {
			field.DocLines = fn.Docs.Lines
		}

		if filter(field) {
			if isMutation(fn.Name) {
				mutationFields = append(mutationFields, field)
			} else {
				queryFields = append(queryFields, field)
			}
		}
	}

	results := &RootObjects{
		QueryFields:    queryFields,
		MutationFields: mutationFields,
	}

	return results, errors
}

func filterTypes(types []*TypeDefinition, fields []*FieldDefinition, forInput bool) []*TypeDefinition {
	// Filter out types that are not used by any field.
	// Also then recursively filter out types that are not used by any type.

	// Make a map of all types
	typeMap := make(map[string]*TypeDefinition, len(types))
	for _, t := range types {
		name := getBaseType(t.Name)
		typeMap[name] = t
	}

	// Get all types used by fields, including subtypes
	usedTypes := make(map[string]bool)
	for _, f := range fields {
		if forInput {
			for _, p := range f.Arguments {
				addUsedTypes(p.Type, typeMap, usedTypes)
			}
		} else {
			addUsedTypes(f.Type, typeMap, usedTypes)
		}
	}

	// Filter out types that are not used
	results := make([]*TypeDefinition, 0, len(types))
	for _, t := range types {
		name := getBaseType(t.Name)
		if usedTypes[name] {
			results = append(results, t)
		}
	}

	return results
}

func extractCustomScalarTypes(inputTypeDefs, resultTypeDefs map[string]*TypeDefinition) []string {
	scalarTypes := make(map[string]bool)
	for _, t := range inputTypeDefs {
		if len(t.Fields) == 0 {
			scalarTypes[t.Name] = true
			delete(inputTypeDefs, t.Name)
		}
	}
	for _, t := range resultTypeDefs {
		if len(t.Fields) == 0 {
			scalarTypes[t.Name] = true
			delete(resultTypeDefs, t.Name)
		}
	}

	return utils.MapKeys(scalarTypes)
}

func addUsedTypes(name string, types map[string]*TypeDefinition, usedTypes map[string]bool) {
	name = getBaseType(name)
	if usedTypes[name] {
		return
	}
	usedTypes[name] = true
	if t, ok := types[name]; ok {
		for _, f := range t.Fields {
			addUsedTypes(f.Type, types, usedTypes)
		}
	}
}

func getBaseType(name string) string {
	name = strings.TrimSuffix(name, "!")
	if strings.HasPrefix(name, "[") {
		return getBaseType(name[1 : len(name)-1])
	}

	return name
}

func writeSchema(buf *bytes.Buffer, root *RootObjects, scalarTypes []string, inputTypeDefs, resultTypeDefs []*TypeDefinition) {

	// write header
	buf.WriteString("# Modus GraphQL Schema (auto-generated)\n")

	// sort everything
	slices.SortFunc(root.QueryFields, func(a, b *FieldDefinition) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	slices.SortFunc(root.MutationFields, func(a, b *FieldDefinition) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	slices.SortFunc(scalarTypes, func(a, b string) int {
		return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	slices.SortFunc(inputTypeDefs, func(a, b *TypeDefinition) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	slices.SortFunc(resultTypeDefs, func(a, b *TypeDefinition) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})

	// write query object
	if len(root.QueryFields) > 0 {
		buf.WriteByte('\n')
		buf.WriteString("type Query {\n")
		for _, field := range root.QueryFields {
			writeField(buf, field)
		}
		buf.WriteString("}\n")
	}

	// write mutation object
	if len(root.MutationFields) > 0 {
		buf.WriteByte('\n')
		buf.WriteString("type Mutation {\n")
		for _, field := range root.MutationFields {
			writeField(buf, field)
		}
		buf.WriteString("}\n")
	}

	// write scalars
	if len(scalarTypes) > 0 {
		buf.WriteByte('\n')
		for _, scalar := range scalarTypes {
			buf.WriteString("scalar ")
			buf.WriteString(scalar)
			buf.WriteByte('\n')
		}
	}
	// write input types
	for _, t := range inputTypeDefs {
		buf.WriteByte('\n')

		if len(t.DocLines) > 0 {
			buf.WriteString("\"\"\"\n")
			for _, line := range t.DocLines {
				buf.WriteString(line)
				buf.WriteByte('\n')
			}
			buf.WriteString("\"\"\"\n")
		}

		buf.WriteString("input ")
		buf.WriteString(t.Name)
		buf.WriteString(" {\n")
		for _, f := range t.Fields {

			if len(f.DocLines) > 0 {
				buf.WriteString("  \"\"\"\n")
				for _, line := range f.DocLines {
					buf.WriteString("  ")
					buf.WriteString(line)
					buf.WriteByte('\n')
				}
				buf.WriteString("  \"\"\"\n")
			}

			buf.WriteString("  ")
			buf.WriteString(f.Name)
			buf.WriteString(": ")
			buf.WriteString(f.Type)
			buf.WriteByte('\n')
		}
		buf.WriteString("}\n")
	}

	// write result types
	for _, t := range resultTypeDefs {
		buf.WriteByte('\n')

		if len(t.DocLines) > 0 {
			buf.WriteString("\"\"\"\n")
			for _, line := range t.DocLines {
				buf.WriteString(line)
				buf.WriteByte('\n')
			}
			buf.WriteString("\"\"\"\n")
		}

		buf.WriteString("type ")
		buf.WriteString(t.Name)
		buf.WriteString(" {\n")
		for _, f := range t.Fields {

			if len(f.DocLines) > 0 {
				buf.WriteString("  \"\"\"\n")
				for _, line := range f.DocLines {
					buf.WriteString("  ")
					buf.WriteString(line)
					buf.WriteByte('\n')
				}
				buf.WriteString("  \"\"\"\n")
			}

			buf.WriteString("  ")
			buf.WriteString(f.Name)
			buf.WriteString(": ")
			buf.WriteString(f.Type)
			buf.WriteByte('\n')
		}
		buf.WriteString("}\n")
	}
}

func writeField(buf *bytes.Buffer, field *FieldDefinition) {

	if len(field.DocLines) > 0 {
		buf.WriteString("  \"\"\"\n")
		for _, line := range field.DocLines {
			buf.WriteString("  ")
			buf.WriteString(line)
			buf.WriteByte('\n')
		}
		buf.WriteString("  \"\"\"\n")
	}

	buf.WriteString("  ")
	buf.WriteString(field.Name)
	if len(field.Arguments) > 0 {
		buf.WriteByte('(')
		for i, p := range field.Arguments {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(p.Name)
			buf.WriteString(": ")
			buf.WriteString(p.Type)
			if p.Default != nil {
				val, err := utils.JsonSerialize(*p.Default)
				if err == nil {
					buf.WriteString(" = ")
					buf.Write(val)
				}
			}
		}
		buf.WriteByte(')')
	}
	buf.WriteString(": ")
	buf.WriteString(field.Type)
	buf.WriteByte('\n')
}

func convertParameters(parameters []*metadata.Parameter, lti langsupport.LanguageTypeInfo, typeDefs map[string]*TypeDefinition) ([]*ArgumentDefinition, error) {
	if len(parameters) == 0 {
		return nil, nil
	}

	args := make([]*ArgumentDefinition, len(parameters))
	for i, p := range parameters {

		t, err := convertType(p.Type, lti, typeDefs, false, true)
		if err != nil {
			return nil, err
		}

		args[i] = &ArgumentDefinition{
			Name:    p.Name,
			Type:    t,
			Default: p.Default,
		}
	}
	return args, nil
}

func convertResults(results []*metadata.Result, lti langsupport.LanguageTypeInfo, typeDefs map[string]*TypeDefinition) (string, error) {
	switch len(results) {
	case 0:
		return newScalar("Void", typeDefs), nil
	case 1:
		// Note: Single result doesn't use the name, even if it's present.
		return convertType(results[0].Type, lti, typeDefs, false, false)
	}

	fields := make([]*FieldDefinition, len(results))
	for i, r := range results {
		name := r.Name
		if name == "" {
			name = fmt.Sprintf("item%d", i+1)
		}

		typ, err := convertType(r.Type, lti, typeDefs, false, false)
		if err != nil {
			return "", err
		}

		fields[i] = &FieldDefinition{
			Name: name,
			Type: typ,
		}
	}

	t := getTypeForFields(fields, typeDefs)
	return t, nil
}

func getTypeForFields(fields []*FieldDefinition, typeDefs map[string]*TypeDefinition) string {
	// see if an existing type already matches
	for _, t := range typeDefs {
		if len(t.Fields) != len(fields) {
			continue
		}

		found := true
		for i, f := range fields {
			if t.Fields[i].Name != f.Name || t.Fields[i].Type != f.Type {
				found = false
				break
			}
		}

		if found {
			return t.Name
		}
	}

	// there's no existing type that matches, so create a new one
	var name string
	for i := 1; ; i++ {
		name = fmt.Sprintf("_type%d", i)
		if _, ok := typeDefs[name]; !ok {
			break
		}
	}

	return newType(name, fields, typeDefs)
}

func convertFields(fields []*metadata.Field, lti langsupport.LanguageTypeInfo, typeDefs map[string]*TypeDefinition, forInput bool) ([]*FieldDefinition, error) {
	if len(fields) == 0 {
		return nil, nil
	}

	results := make([]*FieldDefinition, len(fields))
	for i, f := range fields {
		t, err := convertType(f.Type, lti, typeDefs, true, forInput)
		if err != nil {
			return nil, err
		}

		fieldDef := &FieldDefinition{
			Name: f.Name,
			Type: t,
		}

		if f.Docs != nil {
			fieldDef.DocLines = f.Docs.Lines
		}

		results[i] = fieldDef
	}
	return results, nil
}

func convertType(typ string, lti langsupport.LanguageTypeInfo, typeDefs map[string]*TypeDefinition, firstPass, forInput bool) (string, error) {
	gmlPrintf("GML: schemagen.go: ENTER convertType: typ='%v'", typ)
	if underlyingType, ok := lti.IsErrorType(typ); ok {
		typ = underlyingType
	}

	if !lti.IsTupleType(typ) {
		// Unwrap parentheses if present
		if strings.HasPrefix(typ, "(") && strings.HasSuffix(typ, ")") {
			return convertType(typ[1:len(typ)-1], lti, typeDefs, firstPass, forInput)
		}
	}

	// Set the nullable flag.
	// In GraphQL, types are nullable by default,
	// and non-nullable types are indicated by a "!" suffix
	var n string
	if !lti.IsNullableType(typ) {
		n = "!"
	}

	// unwrap nullable types (and dereference pointers)
	for lti.IsNullableType(typ) {
		t := lti.GetUnderlyingType(typ)
		if t == typ {
			break
		}
		typ = t
	}

	// convert basic types
	// TODO: How do we want to provide GraphQL "ID" scalar types? Maybe they're annotated? or maybe by naming convention?

	if lti.IsStringType(typ) {
		gmlPrintf("GML: schemagen.go: convertType: A: typ='%v'", typ)
		return "String" + n, nil
	}

	if lti.IsByteSequenceType(typ) {
		// Note: If the bytes represent valid UTF-8 strings, Go will serialize them as actual strings.
		// Otherwise, the data will be base64 encoded.
		// TODO: We may want to ensure that the results are _always_ base64 encoded.
		gmlPrintf("GML: schemagen.go: convertType: B: typ='%v'", typ)
		return "String" + n, nil
	}

	if lti.IsBooleanType(typ) {
		gmlPrintf("GML: schemagen.go: convertType: C: typ='%v'", typ)
		return "Boolean" + n, nil
	}

	if lti.IsFloatType(typ) {
		gmlPrintf("GML: schemagen.go: convertType: D: typ='%v'", typ)
		return "Float" + n, nil
	}

	if lti.IsIntegerType(typ) {
		ctx := context.Background() // context is always unused for this purpose
		signed := lti.IsSignedIntegerType(typ)
		size, err := lti.GetSizeOfType(ctx, typ)
		if err != nil {
			return "", err
		}

		switch size {
		case 8:
			if signed {
				gmlPrintf("GML: schemagen.go: convertType: E: typ='%v'", typ)
				return newScalar("Int64", typeDefs) + n, nil
			} else {
				gmlPrintf("GML: schemagen.go: convertType: F: typ='%v'", typ)
				return newScalar("UInt64", typeDefs) + n, nil
			}
		case 4:
			if !signed {
				gmlPrintf("GML: schemagen.go: convertType: G: typ='%v'", typ)
				return newScalar("UInt", typeDefs) + n, nil
			}
		}

		gmlPrintf("GML: schemagen.go: convertType: H: typ='%v'", typ)
		return "Int" + n, nil
	}

	if lti.IsTimestampType(typ) {
		gmlPrintf("GML: schemagen.go: convertType: I: typ='%v'", typ)
		return newScalar("Timestamp", typeDefs) + n, nil
	}

	// check for array types
	if lti.IsListType(typ) {
		elem := lti.GetListSubtype(typ)
		t, err := convertType(elem, lti, typeDefs, firstPass, forInput)
		if err != nil {
			return "", err
		}
		gmlPrintf("GML: schemagen.go: convertType: J: typ='%v'", typ)
		return "[" + t + "]" + n, nil
	}

	if lti.IsTupleType(typ) {
		subTypes := lti.GetTupleSubtypes(typ)
		// tts := make([]string, 0, len(subTypes))
		ttns := make([]string, 0, len(subTypes))
		fields := make([]*FieldDefinition, 0, len(subTypes))
		for i, subType := range subTypes {
			tt, err := convertType(subType, lti, typeDefs, firstPass, forInput)
			if err != nil {
				return "", err
			}
			// tts = append(tts, tt)

			var ttn string
			if strings.HasSuffix(tt, "!") {
				ttn = tt[:len(tt)-1]
			} else if tt[0] == '[' {
				ttn = "[Nullable" + tt[1:]
			} else {
				ttn = "Nullable" + tt
			}
			if ttn[0] == '[' {
				t := ttn[1 : len(ttn)-2]
				if forInput {
					t = strings.TrimSuffix(t, "Input")
				}
				ttn = t + "List"
			} else if forInput {
				ttn = strings.TrimSuffix(ttn, "Input")
			}
			ttns = append(ttns, ttn)
			fields = append(fields, &FieldDefinition{Name: fmt.Sprintf("t%v", i), Type: tt})
		}

		typeName := strings.Join(ttns, "") + "Tuple"
		if forInput {
			typeName += "Input"
		}

		newTupleType(typeName, fields, typeDefs)

		// The tuple is represented as a map.
		// e.g. IntBoolStringTuple! or IntBoolStringTupleInput!
		gmlPrintf("GML: schemagen.go: convertType: K: typ='%v'", typ)
		return typeName + n, nil
	}

	// check for map types
	if lti.IsMapType(typ) {
		k, v := lti.GetMapSubtypes(typ)
		kt, err := convertType(k, lti, typeDefs, firstPass, forInput)
		if err != nil {
			return "", err
		}
		vt, err := convertType(v, lti, typeDefs, firstPass, forInput)
		if err != nil {
			return "", err
		}

		// The pair type name will be composed from the key and value types.
		// ex: StringStringPair, IntStringPair, StringNullableStringPair, etc.
		var ktn, vtn string
		if strings.HasSuffix(kt, "!") {
			ktn = kt[:len(kt)-1]
		} else if kt[0] == '[' {
			ktn = "[Nullable" + kt[1:]
		} else {
			ktn = "Nullable" + kt
		}

		if strings.HasSuffix(vt, "!") {
			vtn = vt[:len(vt)-1]
		} else if vt[0] == '[' {
			vtn = "[Nullable" + vt[1:]
		} else {
			vtn = "Nullable" + vt
		}

		if ktn[0] == '[' {
			t := ktn[1 : len(ktn)-2]
			if forInput {
				t = strings.TrimSuffix(t, "Input")
			}
			ktn = t + "List"
		} else if forInput {
			ktn = strings.TrimSuffix(ktn, "Input")
		}
		if vtn[0] == '[' {
			t := vtn[1 : len(vtn)-2]
			if forInput {
				t = strings.TrimSuffix(t, "Input")
			}
			vtn = t + "List"
		} else if forInput {
			vtn = strings.TrimSuffix(vtn, "Input")
		}

		typeName := ktn + vtn + "Pair"
		if forInput {
			typeName += "Input"
		}

		fields := []*FieldDefinition{{Name: "key", Type: kt}, {Name: "value", Type: vt}}
		newMapType(typeName, fields, typeDefs)

		// The map is represented as a list of the pair type.
		// The list might be nullable, but the pair type within the list is always non-nullable.
		// ex: [StringStringPair!] or [StringStringPair!]!
		gmlPrintf("GML: schemagen.go: convertType: K: typ='%v'", typ)
		return "[" + typeName + "!]" + n, nil
	}

	name := lti.GetNameForType(typ)
	if forInput {
		if !strings.HasSuffix(name, "Input") {
			name += "Input"
		}
	} else {
		name = strings.TrimSuffix(name, "Input")
	}

	// in the first pass, we convert input custom type definitions
	if firstPass {
		gmlPrintf("GML: schemagen.go: convertType: L: typ='%v'", typ)
		return name + n, nil
	}

	// going forward, convert custom types only if they have a type definition
	if _, ok := typeDefs[name]; ok {
		gmlPrintf("GML: schemagen.go: convertType: M: typ='%v'", typ)
		return name + n, nil
	}

	// edge case: a custom scalar used for input
	if forInput {
		name = strings.TrimSuffix(name, "Input")
		if _, ok := typeDefs[name]; ok {
			gmlPrintf("GML: schemagen.go: convertType: N: typ='%v'", typ)
			return name + n, nil
		}
	}

	return "", fmt.Errorf("unsupported type or missing type definition: %s", typ)
}

func newScalar(name string, typeDefs map[string]*TypeDefinition) string {
	return newType(name, nil, typeDefs)
}

func newType(name string, fields []*FieldDefinition, typeDefs map[string]*TypeDefinition) string {
	if _, ok := typeDefs[name]; !ok {
		typeDefs[name] = &TypeDefinition{
			Name:   name,
			Fields: fields,
		}
	}
	return name
}

func newMapType(name string, fields []*FieldDefinition, typeDefs map[string]*TypeDefinition) string {
	if _, ok := typeDefs[name]; !ok {
		typeDefs[name] = &TypeDefinition{
			Name:      name,
			Fields:    fields,
			IsMapType: true,
		}
	}
	return name
}

func newTupleType(name string, fields []*FieldDefinition, typeDefs map[string]*TypeDefinition) string {
	if _, ok := typeDefs[name]; !ok {
		typeDefs[name] = &TypeDefinition{
			Name:        name,
			Fields:      fields,
			IsTupleType: true,
		}
	}
	return name
}
