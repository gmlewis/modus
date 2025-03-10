/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package datasource

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gmlewis/modus/runtime/logger"
	"github.com/gmlewis/modus/runtime/utils"
	"github.com/gmlewis/modus/runtime/wasmhost"

	"github.com/buger/jsonparser"
	"github.com/wundergraph/graphql-go-tools/v2/pkg/engine/datasource/httpclient"
	"github.com/wundergraph/graphql-go-tools/v2/pkg/engine/resolve"
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

const DataSourceName = "ModusDataSource"

type callInfo struct {
	FieldInfo    fieldInfo      `json:"field"`
	FunctionName string         `json:"function"`
	Parameters   map[string]any `json:"data"`
}

type ModusDataSource struct {
	WasmHost wasmhost.WasmHost
}

func (ds *ModusDataSource) Load(ctx context.Context, input []byte, out *bytes.Buffer) error {

	// Parse the input to get the function call info
	var ci callInfo
	err := utils.JsonDeserialize(input, &ci)
	if err != nil {
		return fmt.Errorf("error parsing input: %w", err)
	}

	// Load the data
	result, gqlErrors, err := ds.callFunction(ctx, &ci)

	// Write the response
	err = writeGraphQLResponse(ctx, out, result, gqlErrors, err, &ci)
	if err != nil {
		logger.Error(ctx).Err(err).Msg("Error creating GraphQL response.")
	}

	return err
}

func (*ModusDataSource) LoadWithFiles(ctx context.Context, input []byte, files []*httpclient.FileUpload, out *bytes.Buffer) (err error) {
	// See https://github.com/wundergraph/graphql-go-tools/pull/758
	panic("not implemented")
}

func (ds *ModusDataSource) callFunction(ctx context.Context, callInfo *callInfo) (any, []resolve.GraphQLError, error) {

	// Handle special case for __typename on root Query or Mutation
	if callInfo.FieldInfo.Name == "__typename" {
		return callInfo.FieldInfo.ParentType, nil, nil
	}

	// Get the function info
	fnInfo, err := ds.WasmHost.GetFunctionInfo(callInfo.FunctionName)
	if err != nil {
		return nil, nil, err
	}

	// Call the function
	execInfo, err := ds.WasmHost.CallFunction(ctx, fnInfo, callInfo.Parameters)
	if err != nil {
		// The full error message has already been logged.  Return a generic error to the caller, which will be included in the response.
		return nil, nil, errors.New("error calling function")
	}

	// Store the execution info into the function output map.
	outputMap := ctx.Value(utils.FunctionOutputContextKey).(map[string]wasmhost.ExecutionInfo)
	outputMap[callInfo.FieldInfo.AliasOrName()] = execInfo

	// Transform messages (and error lines in the output buffers) to GraphQL errors.
	messages := append(execInfo.Messages(), utils.TransformConsoleOutput(execInfo.Buffers())...)
	gqlErrors := transformErrors(messages, callInfo)

	// Get the result.
	result := execInfo.Result()

	// If we have multiple results, unpack them into a map that matches the schema generated type.
	if results, ok := result.([]any); ok && len(fnInfo.ExecutionPlan().ResultHandlers()) > 1 {
		fnMeta := fnInfo.Metadata()
		gmlPrintf("GML: graphql/datasource/source.go: callFunction: results=%+v, fnMeta=%+v", results, fnMeta)
		m := make(map[string]any, len(results))
		for i, r := range results {
			name := fnMeta.Results[i].Name
			if name == "" {
				name = fmt.Sprintf("item%d", i+1)
			}
			m[name] = r
		}
		result = m
	}

	return result, gqlErrors, err
}

func writeGraphQLResponse(ctx context.Context, out *bytes.Buffer, result any, gqlErrors []resolve.GraphQLError, fnErr error, ci *callInfo) error {

	fieldName := ci.FieldInfo.AliasOrName()

	// Include the function error
	if fnErr != nil {
		gqlErrors = append(gqlErrors, resolve.GraphQLError{
			Message: fnErr.Error(),
			Path:    []any{fieldName},
			Extensions: map[string]interface{}{
				"level": "error",
			},
		})
	}

	// If there are GraphQL errors, serialize them as json
	var jsonErrors []byte
	if len(gqlErrors) > 0 {
		var err error
		jsonErrors, err = utils.JsonSerialize(gqlErrors)
		if err != nil {
			return err
		}
	}

	// If there is any result data, or if the data is null without errors, serialize the data as json
	var jsonData []byte
	if result != nil || len(gqlErrors) == 0 {
		jsonResult, err := utils.JsonSerialize(result)
		if err != nil {
			if err, ok := err.(*json.UnsupportedValueError); ok {
				msg := fmt.Sprintf("Function completed successfully, but the result contains a %v value that cannot be serialized to JSON.", err.Value)
				logger.Warn(ctx).
					Bool("user_visible", true).
					Str("function", ci.FunctionName).
					Str("result", fmt.Sprintf("%+v", result)).
					Msg(msg)
				fmt.Fprintf(out, `{"errors":[{"message":"%s","path":["%s"],"extensions":{"level":"error"}}]}`, msg, fieldName)
				return nil
			}
			return err
		}

		// Transform the data
		if r, err := transformValue(jsonResult, &ci.FieldInfo); err != nil {
			return err
		} else {
			jsonData = r
		}
	}

	// Write the response.  This should be as efficient as possible, as it is called for every function invocation.
	out.Grow(len(jsonData) + len(jsonErrors) + len(fieldName) + 26)
	out.WriteByte('{')
	if len(jsonData) > 0 {
		out.WriteString(`"data":{"`)
		out.WriteString(fieldName)
		out.WriteString(`":`)
		out.Write(jsonData)
		out.WriteByte('}')
	}
	if len(jsonErrors) > 0 {
		if len(jsonData) > 0 {
			out.WriteByte(',')
		}
		out.WriteString(`"errors":`)
		out.Write(jsonErrors)
	}
	out.WriteByte('}')

	return nil
}

var nullWord = []byte("null")

func transformValue(data []byte, tf *fieldInfo) (result []byte, err error) {
	gmlPrintf("GML: graphql/datasource/source.go: transformValue: data='%s'", data)

	if len(tf.Fields) == 0 || len(data) == 0 || bytes.Equal(data, nullWord) {
		return data, nil
	}

	switch data[0] {
	case '{':
		if tf.IsMapType {
			return transformMap(data, tf)
		}
		return transformObject(data, tf)
	case '[':
		if tf.IsTupleType {
			return transformTuple(data, tf)
		}
		return transformArray(data, tf)
	default:
		return nil, fmt.Errorf("expected object or array, got: '%s'", data)
	}
}

func transformArray(data []byte, tf *fieldInfo) ([]byte, error) {
	gmlPrintf("GML: graphql/datasource/source.go: transformArray: data='%s'", data)
	buf := bytes.Buffer{}
	buf.WriteByte('[')

	var loopErr error
	_, err := jsonparser.ArrayEach(data, func(val []byte, _ jsonparser.ValueType, _ int, _ error) {
		if loopErr != nil {
			return
		}
		val, err := transformValue(val, tf)
		if err != nil {
			loopErr = err
			return
		}
		if buf.Len() > 1 {
			buf.WriteByte(',')
		}
		buf.Write(val)
	})
	if err != nil {
		return nil, err
	}
	if loopErr != nil {
		return nil, loopErr
	}

	buf.WriteByte(']')
	gmlPrintf("GML: graphql/datasource/source.go: transformArray: buf='%v'", buf.String())
	return buf.Bytes(), nil
}

func transformTuple(data []byte, tf *fieldInfo) ([]byte, error) {
	gmlPrintf("GML: graphql/datasource/source.go: transformTuple: data='%s'", data)
	buf := bytes.Buffer{}
	buf.WriteByte('{')

	var loopErr error
	var argCount int
	_, err := jsonparser.ArrayEach(data, func(value []byte, valueType jsonparser.ValueType, _ int, _ error) {
		if loopErr != nil {
			return
		}
		if argCount >= len(tf.Fields) {
			loopErr = fmt.Errorf("transformTuple: parsing field[%v] but len(Fields)=%v", argCount, len(tf.Fields))
			return
		}
		f := &tf.Fields[argCount]
		value, err := transformValue(value, f)
		if err != nil {
			loopErr = err
			return
		}

		if buf.Len() > 1 {
			buf.WriteByte(',')
		}

		// key := []byte(fmt.Sprintf("t%v", argCount))
		// b, err := writeKeyValuePair(key, value, "String", valueType)
		// if err != nil {
		// 	loopErr = err
		// 	return
		// }

		// f := &tf.Fields[argCount]
		// val, err := transformObject(b, f)
		// if err != nil {
		// 	loopErr = err
		// 	return
		// }
		// buf.Write(val)
		// buf.Write(b)

		buf.WriteString(fmt.Sprintf(`"t%v":`, argCount))
		if valueType == jsonparser.String {
			buf.WriteByte('"')
		}
		buf.Write(value)
		if valueType == jsonparser.String {
			buf.WriteByte('"')
		}
		argCount++
	})
	if err != nil {
		return nil, err
	}
	if loopErr != nil {
		return nil, loopErr
	}

	buf.WriteByte('}')
	gmlPrintf("GML: graphql/datasource/source.go: transformTuple: buf='%v'", buf.String())
	return buf.Bytes(), nil
}

func transformObject(data []byte, tf *fieldInfo) ([]byte, error) {
	gmlPrintf("GML: graphql/datasource/source.go: transformObject: data='%s'", data)
	buf := bytes.Buffer{}
	buf.WriteByte('{')
	for i, f := range tf.Fields {
		var val []byte
		if f.Name == "__typename" {
			val = []byte(`"` + tf.TypeName + `"`)
		} else {
			v, dataType, _, err := jsonparser.Get(data, f.Name)
			if err != nil {
				return nil, err
			}
			if dataType == jsonparser.String {
				// Note, string values here will be escaped for internal quotes, newlines, etc.,
				// but will be missing outer quotes.  So we need to add them back.
				v = []byte(`"` + string(v) + `"`)
			}
			val, err = transformValue(v, &f)
			if err != nil {
				return nil, err
			}
		}
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte('"')
		buf.WriteString(f.AliasOrName())
		buf.WriteString(`":`)
		buf.Write(val)
	}
	buf.WriteByte('}')
	gmlPrintf("GML: graphql/datasource/source.go: transformObject: buf='%v'", buf.String())
	return buf.Bytes(), nil
}

func transformMap(data []byte, tf *fieldInfo) ([]byte, error) {
	gmlPrintf("GML: graphql/datasource/source.go: transformMap: data='%s'", data)

	// check for pseudo map
	md, dt, _, err := jsonparser.Get(data, "$mapdata")
	if err == nil && dt == jsonparser.Array {
		return transformPseudoMap(md, tf)
	}

	var keyType string
	for _, f := range tf.Fields {
		if f.Name == "key" {
			keyType = f.TypeName
			break
		}
	}

	buf := bytes.Buffer{}
	buf.WriteByte('[')
	if err := jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if buf.Len() > 1 {
			buf.WriteByte(',')
		}

		b, err := writeKeyValuePair(key, value, keyType, dataType)
		if err != nil {
			return err
		}

		val, err := transformObject(b, tf)
		if err != nil {
			return err
		}
		buf.Write(val)

		return nil
	}); err != nil {
		return nil, err
	}

	buf.WriteByte(']')
	gmlPrintf("GML: graphql/datasource/source.go: transformMap: buf='%v'", buf.String())
	return buf.Bytes(), nil
}

func writeKeyValuePair(key, value []byte, keyType string, dataType jsonparser.ValueType) ([]byte, error) {
	b := bytes.Buffer{}
	b.WriteByte('{')
	b.WriteString(`"key":`)
	if keyType == "String" {
		k, err := utils.JsonSerialize(string(key))
		if err != nil {
			return nil, err
		}
		b.Write(k)
	} else {
		b.Write(key)
	}
	b.WriteString(`,"value":`)
	if dataType == jsonparser.String {
		b.WriteByte('"')
		b.Write(value)
		b.WriteByte('"')
	} else {
		b.Write(value)
	}
	b.WriteByte('}')
	gmlPrintf("GML: graphql/datasource/source.go: writeKeyValuePair: buf='%v'", b.String())
	return b.Bytes(), nil
}

func transformPseudoMap(data []byte, tf *fieldInfo) ([]byte, error) {
	gmlPrintf("GML: graphql/datasource/source.go: transformPseudoMap: data='%s'", data)
	buf := bytes.Buffer{}
	buf.WriteByte('[')

	var loopErr error
	_, err := jsonparser.ArrayEach(data, func(item []byte, _ jsonparser.ValueType, _ int, _ error) {
		if loopErr != nil {
			return
		}

		key, kdt, _, err := jsonparser.Get(item, "key")
		if err != nil {
			loopErr = err
			return
		}

		value, vdt, _, err := jsonparser.Get(item, "value")
		if err != nil {
			loopErr = err
			return
		}

		if buf.Len() > 1 {
			buf.WriteByte(',')
		}

		b := bytes.Buffer{}
		b.WriteByte('{')
		b.WriteString(`"key":`)
		if kdt == jsonparser.String {
			b.WriteString(`"`)
			b.Write(key)
			b.WriteString(`"`)
		} else {
			b.Write(key)
		}
		b.WriteString(`,"value":`)
		if vdt == jsonparser.String {
			b.WriteString(`"`)
			b.Write(value)
			b.WriteString(`"`)
		} else {
			b.Write(value)
		}
		b.WriteByte('}')

		val, err := transformObject(b.Bytes(), tf)
		if err != nil {
			loopErr = err
			return
		}
		buf.Write(val)
	})
	if err != nil {
		return nil, err
	}
	if loopErr != nil {
		return nil, loopErr
	}

	buf.WriteByte(']')
	return buf.Bytes(), nil
}

func transformErrors(messages []utils.LogMessage, ci *callInfo) []resolve.GraphQLError {
	gmlPrintf("GML: graphql/datasource/source.go: transformErrors: messages=%+v", messages)
	errors := make([]resolve.GraphQLError, 0, len(messages))
	for _, msg := range messages {
		// Only include errors.  Other messages will be captured later and
		// passed back as logs in the extensions section of the response.
		if msg.IsError() {
			errors = append(errors, resolve.GraphQLError{
				Message: msg.Message,
				Path:    []any{ci.FieldInfo.AliasOrName()},
				Extensions: map[string]interface{}{
					"level": msg.Level,
				},
			})
		}
	}
	return errors
}
