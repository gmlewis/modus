// -*- compile-command: "go test -run ^TestFilterMetadata.*YoutubeWalkthrough$ ."; -*-

/*
 * Copyright 2024 Hypermode Inc.
 * Licensed under the terms of the Apache License, Version 2.0
 * See the LICENSE file that accompanied this code for further details.
 *
 * SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package wasm

import (
	"testing"

	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/config"
	"github.com/gmlewis/modus/sdk/go/tools/modus-moonbit-build/metadata"
)

const (
	youtubePath      = "../../../../moonbit/examples/youtube-walkthrough"
	youtubeBuildPath = youtubePath + "/build"
)

func TestFilterMetadata_YoutubeWalkthrough(t *testing.T) {
	config := &config.Config{
		SourceDir:    youtubePath,
		OutputDir:    youtubeBuildPath,
		WasmFileName: "youtube-walkthrough.wasm",
	}
	copyOfBefore := deepCopyMetadata(t, wantYoutubeWalkthroughBeforeFilter)
	testFilterMetadataHelper(t, config, copyOfBefore, wantYoutubeWalkthroughAfterFilter)
}

var wantYoutubeWalkthroughBeforeFilter = &metadata.Metadata{
	Plugin: "youtube-walkthrough",
	Module: "@youtube-walkthrough",
	FnExports: metadata.FunctionMap{
		"__modus_generate_text": {
			Name:       "__modus_generate_text",
			Parameters: []*metadata.Parameter{{Name: "instruction", Type: "String"}, {Name: "prompt", Type: "String"}},
			Results:    []*metadata.Result{{Type: "String raise Error"}},
		},
		"__modus_get_random_quote": {Name: "__modus_get_random_quote", Results: []*metadata.Result{{Type: "Quote raise Error"}}},
		"__modus_say_hello": {
			Name:       "__modus_say_hello",
			Parameters: []*metadata.Parameter{{Name: "name", Type: "String"}},
			Results:    []*metadata.Result{{Type: "String"}},
		},
		"__modus_say_hello_WithDefaults": {Name: "__modus_say_hello_WithDefaults", Results: []*metadata.Result{{Type: "String"}}},
		"cabi_realloc": {
			Name: "cabi_realloc",
			Parameters: []*metadata.Parameter{
				{Name: "src_offset", Type: "Int"}, {Name: "src_size", Type: "Int"},
				{Name: "_dst_alignment", Type: "Int"}, {Name: "dst_size", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
		"copy": {
			Name:       "copy",
			Parameters: []*metadata.Parameter{{Name: "dest", Type: "Int"}, {Name: "src", Type: "Int"}},
		},
		"generate_text": {
			Name:       "generate_text",
			Parameters: []*metadata.Parameter{{Name: "instruction", Type: "String"}, {Name: "prompt", Type: "String"}},
			Results:    []*metadata.Result{{Type: "String raise Error"}},
		},
		"get_random_quote": {Name: "get_random_quote", Results: []*metadata.Result{{Type: "Quote raise Error"}}},
		"malloc": {
			Name:       "malloc",
			Parameters: []*metadata.Parameter{{Name: "size", Type: "Int"}},
			Results:    []*metadata.Result{{Type: "Int"}},
		},
		"ptr_to_none": {Name: "ptr_to_none", Results: []*metadata.Result{{Type: "Int"}}},
		"read_map": {
			Name: "read_map",
			Parameters: []*metadata.Parameter{
				{Name: "key_type_name_ptr", Type: "Int"},
				{Name: "value_type_name_ptr", Type: "Int"}, {Name: "map_ptr", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int64"}},
		},
		"say_hello": {
			Name:       "say_hello",
			Parameters: []*metadata.Parameter{{Name: "name~", Type: "String", Default: AnyPtr(`"World"`)}},
			Results:    []*metadata.Result{{Type: "String"}},
		},
		"say_hello_WithDefaults": {Name: "say_hello_WithDefaults", Results: []*metadata.Result{{Type: "String"}}},
		"write_map": {
			Name: "write_map",
			Parameters: []*metadata.Parameter{
				{Name: "key_type_name_ptr", Type: "Int"},
				{Name: "value_type_name_ptr", Type: "Int"}, {Name: "keys_ptr", Type: "Int"},
				{Name: "values_ptr", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
	},
	FnImports: metadata.FunctionMap{
		"modus_http_client.fetch": {
			Name:       "modus_http_client.fetch",
			Parameters: []*metadata.Parameter{{Name: "request", Type: "@http.Request?"}},
			Results:    []*metadata.Result{{Type: "@http.Response?"}},
		},
		"modus_models.getModelInfo": {
			Name:       "modus_models.getModelInfo",
			Parameters: []*metadata.Parameter{{Name: "model_name", Type: "String"}},
			Results:    []*metadata.Result{{Type: "@models.ModelInfo?"}},
		},
		"modus_models.invokeModel": {
			Name:       "modus_models.invokeModel",
			Parameters: []*metadata.Parameter{{Name: "model_name", Type: "String"}, {Name: "input", Type: "String"}},
			Results:    []*metadata.Result{{Type: "String"}},
		},
		"modus_system.logMessage": {
			Name:       "modus_system.logMessage",
			Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
		},
	},
	Types: metadata.TypeMap{
		"(String)":         {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
		"(String, String)": {Name: "(String, String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}, {Name: "1", Type: "String"}}},
		"@http.Content":    {Name: "@http.Content", Fields: []*metadata.Field{{Name: "data", Type: "String"}}},
		"@http.Header": {Name: "@http.Header", Fields: []*metadata.Field{
			{Name: "name", Type: "String"},
			{Name: "values", Type: "Array[String]"},
		}},
		"@http.Header?":  {Name: "@http.Header?"},
		"@http.Headers":  {Name: "@http.Headers", Fields: []*metadata.Field{{Name: "data", Type: "Map[String, @http.Header?]"}}},
		"@http.Headers?": {Name: "@http.Headers?", Fields: []*metadata.Field{{Name: "data", Type: "Map[String, @http.Header?]"}}},
		"@http.Request": {Name: "@http.Request", Fields: []*metadata.Field{
			{Name: "url", Type: "String"},
			{Name: "method_", Type: "String"},
			{Name: "headers", Type: "@http.Headers?"},
			{Name: "body", Type: "Array[Byte]"},
		}},
		"@http.Request?": {
			Name: "@http.Request?",
			Fields: []*metadata.Field{
				{Name: "url", Type: "String"},
				{Name: "method_", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		"@http.RequestOptions": {Name: "@http.RequestOptions"},
		"@http.Response": {
			Name: "@http.Response",
			Fields: []*metadata.Field{
				{Name: "status", Type: "UInt16"},
				{Name: "status_text", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		"@http.Response raise Error": {
			Name: "@http.Response raise Error",
			Fields: []*metadata.Field{
				{Name: "status", Type: "UInt16"},
				{Name: "status_text", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		"@http.Response?": {
			Name: "@http.Response?",
			Fields: []*metadata.Field{
				{Name: "status", Type: "UInt16"},
				{Name: "status_text", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		// "@http.T":             {Name: "@http.T"},
		// "@http.T raise Error": {Name: "@http.T raise Error"},
		"@models.ModelBase": {
			Name:   "@models.ModelBase",
			Fields: []*metadata.Field{{Name: "mut info", Type: "@models.ModelInfo?"}, {Name: "debug", Type: "Bool"}},
		},
		"@models.ModelBase raise Error": {
			Name:   "@models.ModelBase raise Error",
			Fields: []*metadata.Field{{Name: "mut info", Type: "@models.ModelInfo?"}, {Name: "debug", Type: "Bool"}},
		},
		"@models.ModelInfo": {
			Name:   "@models.ModelInfo",
			Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "full_name", Type: "String"}},
		},
		"@models.ModelInfo raise Error": {
			Name:   "@models.ModelInfo raise Error",
			Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "full_name", Type: "String"}},
		},
		"@models.ModelInfo?": {
			Name:   "@models.ModelInfo?",
			Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "full_name", Type: "String"}},
		},
		"@openai.&RequestMessage": {Name: "@openai.&RequestMessage"},
		"@openai.AssistantMessage": {
			Name: "@openai.AssistantMessage",
			Fields: []*metadata.Field{
				{Name: "role", Type: "String"}, {Name: "content", Type: "String"},
				{Name: "name", Type: "String?"},
				{Name: "tool_calls", Type: "Array[@openai.ToolCall]"},
			},
		},
		"@openai.ChatModel": {Name: "@openai.ChatModel"},
		"@openai.ChatModelInput": {
			Name: "@openai.ChatModelInput",
			Fields: []*metadata.Field{
				{Name: "model", Type: "String"},
				{Name: "messages", Type: "Array[@openai.&RequestMessage]"},
				{Name: "frequency_penalty", Type: "Double?"},
				{Name: "logit_bias", Type: "Map[String, Double]?"},
				{Name: "logprobs", Type: "Bool?"}, {Name: "top_logprobs", Type: "Int?"},
				{Name: "max_tokens", Type: "Int?"}, {Name: "n", Type: "Int?"},
				{Name: "presence_penalty", Type: "Double?"},
				{Name: "response_format", Type: "@openai.ResponseFormat"},
				{Name: "seed", Type: "Int?"},
				{Name: "service_tier", Type: "@openai.ServiceTier?"},
				{Name: "stop", Type: "Array[String]"},
				{Name: "mut temperature", Type: "Double"},
				{Name: "top_p", Type: "Double?"},
				{Name: "tools", Type: "Array[@openai.Tool]"},
				{Name: "tool_choice", Type: "@openai.ToolChoice"},
				{Name: "parallel_tool_calls", Type: "Bool?"},
				{Name: "user", Type: "String?"},
			},
		},
		"@openai.ChatModelOutput": {
			Name: "@openai.ChatModelOutput",
			Fields: []*metadata.Field{
				{Name: "id", Type: "String"}, {Name: "object", Type: "String"},
				{Name: "choices", Type: "Array[@openai.Choice]"},
				{Name: "created", Type: "Int"}, {Name: "model", Type: "String"},
				{Name: "service_tier", Type: "@openai.ServiceTier?"},
				{Name: "usage", Type: "@openai.Usage"},
			},
		},
		"@openai.ChatModelOutput raise Error": {
			Name: "@openai.ChatModelOutput raise Error",
			Fields: []*metadata.Field{
				{Name: "id", Type: "String"}, {Name: "object", Type: "String"},
				{Name: "choices", Type: "Array[@openai.Choice]"},
				{Name: "created", Type: "Int"}, {Name: "model", Type: "String"},
				{Name: "service_tier", Type: "@openai.ServiceTier?"},
				{Name: "usage", Type: "@openai.Usage"},
			},
		},
		"@openai.Choice": {
			Name: "@openai.Choice",
			Fields: []*metadata.Field{
				{Name: "finish_reason", Type: "String"}, {Name: "index", Type: "Int"},
				{Name: "message", Type: "@openai.CompletionMessage"},
			},
		},
		"@openai.CompletionMessage": {
			Name: "@openai.CompletionMessage",
			Fields: []*metadata.Field{
				{Name: "role", Type: "String"}, {Name: "content", Type: "String"},
				{Name: "refusal", Type: "String?"},
			},
		},
		"@openai.Content": {Name: "@openai.Content"},
		"@openai.Embedding": {
			Name: "@openai.Embedding",
			Fields: []*metadata.Field{
				{Name: "object", Type: "String"}, {Name: "index", Type: "Int"},
				{Name: "embedding", Type: "Array[Double]"},
			},
		},
		"@openai.EmbeddingsModel": {Name: "@openai.EmbeddingsModel"},
		"@openai.EmbeddingsModelInput": {
			Name: "@openai.EmbeddingsModelInput",
			Fields: []*metadata.Field{
				{Name: "model", Type: "String"}, {Name: "input", Type: "@openai.Content"},
				{Name: "encoding_format", Type: "@openai.EncodingFormat?"},
				{Name: "dimensions", Type: "Int?"}, {Name: "user", Type: "String?"},
			},
		},
		"@openai.EncodingFormat":  {Name: "@openai.EncodingFormat"},
		"@openai.EncodingFormat?": {Name: "@openai.EncodingFormat?"},
		"@openai.Function": {
			Name:   "@openai.Function",
			Fields: []*metadata.Field{{Name: "name", Type: "String"}},
		},
		"@openai.Function?": {
			Name:   "@openai.Function?",
			Fields: []*metadata.Field{{Name: "name", Type: "String"}},
		},
		"@openai.FunctionCall": {
			Name:   "@openai.FunctionCall",
			Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "arguments", Type: "String"}},
		},
		"@openai.FunctionDefinition": {
			Name: "@openai.FunctionDefinition",
			Fields: []*metadata.Field{
				{Name: "name", Type: "String"}, {Name: "description", Type: "String?"},
				{Name: "strict", Type: "Bool?"}, {Name: "parameters", Type: "Json?"},
			},
		},
		"@openai.LogprobsContent": {
			Name: "@openai.LogprobsContent",
			Fields: []*metadata.Field{
				{Name: "token", Type: "String"}, {Name: "logprob", Type: "Double"},
				{Name: "bytes", Type: "Json"},
				{Name: "top_logprobs", Type: "Array[@openai.LogprobsContentObject]"},
			},
		},
		"@openai.LogprobsContentObject": {
			Name: "@openai.LogprobsContentObject",
			Fields: []*metadata.Field{
				{Name: "token", Type: "String"}, {Name: "logprob", Type: "Double"},
				{Name: "bytes", Type: "Json"},
			},
		},
		"@openai.ResponseFormat": {
			Name:   "@openai.ResponseFormat",
			Fields: []*metadata.Field{{Name: "type_", Type: "String"}, {Name: "json_schema", Type: "Json?"}},
		},
		"@openai.ServiceTier":  {Name: "@openai.ServiceTier"},
		"@openai.ServiceTier?": {Name: "@openai.ServiceTier?"},
		"@openai.SystemMessage": {
			Name: "@openai.SystemMessage",
			Fields: []*metadata.Field{
				{Name: "role", Type: "String"}, {Name: "content", Type: "String"},
				{Name: "name", Type: "String?"},
			},
		},
		"@openai.Tool": {
			Name: "@openai.Tool",
			Fields: []*metadata.Field{
				{Name: "type_", Type: "String"},
				{Name: "function", Type: "@openai.FunctionDefinition"},
			},
		},
		"@openai.ToolCall": {
			Name: "@openai.ToolCall",
			Fields: []*metadata.Field{
				{Name: "id", Type: "String"}, {Name: "type_", Type: "String"},
				{Name: "function", Type: "@openai.FunctionCall"},
			},
		},
		"@openai.ToolChoice": {
			Name:   "@openai.ToolChoice",
			Fields: []*metadata.Field{{Name: "type_", Type: "String"}, {Name: "function", Type: "@openai.Function?"}},
		},
		"@openai.ToolMessage": {
			Name: "@openai.ToolMessage",
			Fields: []*metadata.Field{
				{Name: "role", Type: "String"}, {Name: "content", Type: "String"},
				{Name: "tool_call_id", Type: "String"},
			},
		},
		"@openai.Usage": {
			Name: "@openai.Usage",
			Fields: []*metadata.Field{
				{Name: "completion_tokens", Type: "Int"},
				{Name: "prompt_tokens", Type: "Int"}, {Name: "total_tokens", Type: "Int"},
			},
		},
		"@openai.UserMessage": {
			Name: "@openai.UserMessage",
			Fields: []*metadata.Field{
				{Name: "role", Type: "String"}, {Name: "content", Type: "String"},
				{Name: "name", Type: "String?"},
			},
		},
		// "@testutils.CallStack[T]":              {Name: "@testutils.CallStack[T]", Fields: []*metadata.Field{{Name: "items", Type: "Array[Array[@testutils.T]]"}}},
		"@testutils.T":                         {Name: "@testutils.T"},
		"ArrayView[Byte]":                      {Name: "ArrayView[Byte]"},
		"Array[(String, String)]":              {Name: "Array[(String, String)]"},
		"Array[@http.Header?]":                 {Name: "Array[@http.Header?]"},
		"Array[@http.RequestOptions]":          {Name: "Array[@http.RequestOptions]"},
		"Array[@openai.&RequestMessage]":       {Name: "Array[@openai.&RequestMessage]"},
		"Array[@openai.Choice]":                {Name: "Array[@openai.Choice]"},
		"Array[@openai.Embedding]":             {Name: "Array[@openai.Embedding]"},
		"Array[@openai.LogprobsContentObject]": {Name: "Array[@openai.LogprobsContentObject]"},
		"Array[@openai.LogprobsContent]":       {Name: "Array[@openai.LogprobsContent]"},
		"Array[@openai.ToolCall]":              {Name: "Array[@openai.ToolCall]"},
		"Array[@openai.Tool]":                  {Name: "Array[@openai.Tool]"},
		"Array[@testutils.T]":                  {Name: "Array[@testutils.T]"},
		"Array[Array[@testutils.T]]":           {Name: "Array[Array[@testutils.T]]"},
		"Array[Array[String]]":                 {Name: "Array[Array[String]]"},
		"Array[Byte]":                          {Name: "Array[Byte]"},
		"Array[Double]":                        {Name: "Array[Double]"},
		"Array[Int]":                           {Name: "Array[Int]"},
		"Array[String]":                        {Name: "Array[String]"},
		"Bool":                                 {Name: "Bool"},
		"Bool?":                                {Name: "Bool?"},
		"Byte":                                 {Name: "Byte"},
		"Bytes":                                {Name: "Bytes"},
		// "Bytes raise CorruptInputError":        {Name: "Bytes raise CorruptInputError"},
		"Bytes raise Error":          {Name: "Bytes raise Error"},
		"Char":                       {Name: "Char"},
		"Double":                     {Name: "Double"},
		"Double?":                    {Name: "Double?"},
		"FixedArray[Byte]":           {Name: "FixedArray[Byte]"},
		"FixedArray[Double]":         {Name: "FixedArray[Double]"},
		"FixedArray[Float]":          {Name: "FixedArray[Float]"},
		"FixedArray[Int64]":          {Name: "FixedArray[Int64]"},
		"FixedArray[Int]":            {Name: "FixedArray[Int]"},
		"FixedArray[UInt64]":         {Name: "FixedArray[UInt64]"},
		"FixedArray[UInt]":           {Name: "FixedArray[UInt]"},
		"Float":                      {Name: "Float"},
		"Int":                        {Name: "Int"},
		"Int64":                      {Name: "Int64"},
		"Int?":                       {Name: "Int?"},
		"Iter[Byte]":                 {Name: "Iter[Byte]"},
		"Iter[Char]":                 {Name: "Iter[Char]"},
		"Json":                       {Name: "Json"},
		"Json raise Error":           {Name: "Json raise Error"},
		"Json?":                      {Name: "Json?"},
		"Map[String, @http.Header?]": {Name: "Map[String, @http.Header?]"},
		"Map[String, Array[String]]": {Name: "Map[String, Array[String]]"},
		"Map[String, Double]":        {Name: "Map[String, Double]"},
		"Map[String, Double]?":       {Name: "Map[String, Double]?"},
		"Map[String, String]":        {Name: "Map[String, String]"},
		"Quote": {
			Name: "Quote",
			Fields: []*metadata.Field{
				{Name: "q", Type: "String"}, {Name: "a", Type: "String"},
				{Name: "h", Type: "String"},
			},
		},
		"Quote raise Error": {
			Name: "Quote raise Error",
			Fields: []*metadata.Field{
				{Name: "q", Type: "String"}, {Name: "a", Type: "String"},
				{Name: "h", Type: "String"},
			},
		},
		"String": {Name: "String"},
		// "String raise CorruptInputError": {Name: "String raise CorruptInputError"},
		"String raise Error": {Name: "String raise Error"},
		"String?":            {Name: "String?"},
		"UInt":               {Name: "UInt"},
		"UInt16":             {Name: "UInt16"},
		"UInt64":             {Name: "UInt64"},
	},
}

var wantYoutubeWalkthroughAfterFilter = &metadata.Metadata{
	Plugin: "youtube-walkthrough",
	Module: "@youtube-walkthrough",
	FnExports: metadata.FunctionMap{
		"cabi_realloc": {
			Name: "cabi_realloc",
			Parameters: []*metadata.Parameter{
				{Name: "src_offset", Type: "Int"}, {Name: "src_size", Type: "Int"},
				{Name: "_dst_alignment", Type: "Int"}, {Name: "dst_size", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
		"copy": {
			Name:       "copy",
			Parameters: []*metadata.Parameter{{Name: "dest", Type: "Int"}, {Name: "src", Type: "Int"}},
		},
		"generate_text": {
			Name:       "generate_text",
			Parameters: []*metadata.Parameter{{Name: "instruction", Type: "String"}, {Name: "prompt", Type: "String"}},
			Results:    []*metadata.Result{{Type: "String raise Error"}},
		},
		"get_random_quote": {Name: "get_random_quote", Results: []*metadata.Result{{Type: "Quote raise Error"}}},
		"malloc": {
			Name:       "malloc",
			Parameters: []*metadata.Parameter{{Name: "size", Type: "Int"}},
			Results:    []*metadata.Result{{Type: "Int"}},
		},
		"ptr_to_none": {Name: "ptr_to_none", Results: []*metadata.Result{{Type: "Int"}}},
		"read_map": {
			Name: "read_map",
			Parameters: []*metadata.Parameter{
				{Name: "key_type_name_ptr", Type: "Int"},
				{Name: "value_type_name_ptr", Type: "Int"}, {Name: "map_ptr", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int64"}},
		},
		"say_hello": {
			Name:       "say_hello",
			Parameters: []*metadata.Parameter{{Name: "name~", Type: "String", Default: AnyPtr(`"World"`)}},
			Results:    []*metadata.Result{{Type: "String"}},
		},
		"say_hello_WithDefaults": {Name: "say_hello_WithDefaults", Results: []*metadata.Result{{Type: "String"}}},
		"write_map": {
			Name: "write_map",
			Parameters: []*metadata.Parameter{
				{Name: "key_type_name_ptr", Type: "Int"},
				{Name: "value_type_name_ptr", Type: "Int"}, {Name: "keys_ptr", Type: "Int"},
				{Name: "values_ptr", Type: "Int"},
			},
			Results: []*metadata.Result{{Type: "Int"}},
		},
	},
	FnImports: metadata.FunctionMap{
		"modus_http_client.fetch": {
			Name:       "modus_http_client.fetch",
			Parameters: []*metadata.Parameter{{Name: "request", Type: "@http.Request?"}},
			Results:    []*metadata.Result{{Type: "@http.Response?"}},
		},
		"modus_models.getModelInfo": {
			Name:       "modus_models.getModelInfo",
			Parameters: []*metadata.Parameter{{Name: "model_name", Type: "String"}},
			Results:    []*metadata.Result{{Type: "@models.ModelInfo?"}},
		},
		"modus_models.invokeModel": {
			Name:       "modus_models.invokeModel",
			Parameters: []*metadata.Parameter{{Name: "model_name", Type: "String"}, {Name: "input", Type: "String"}},
			Results:    []*metadata.Result{{Type: "String"}},
		},
		"modus_system.logMessage": {
			Name:       "modus_system.logMessage",
			Parameters: []*metadata.Parameter{{Name: "level", Type: "String"}, {Name: "message", Type: "String"}},
		},
	},
	Types: metadata.TypeMap{
		"(String)": {Name: "(String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}}},
		// "(String, String)": {Name: "(String, String)", Fields: []*metadata.Field{{Name: "0", Type: "String"}, {Name: "1", Type: "String"}}},
		// "@http.Content":    {Name: "@http.Content", Fields: []*metadata.Field{{Name: "data", Type: "String"}}},
		"@http.Header": {Name: "@http.Header", Fields: []*metadata.Field{
			{Name: "name", Type: "String"},
			{Name: "values", Type: "Array[String]"},
		}},
		"@http.Header?":  {Name: "@http.Header?"},
		"@http.Headers":  {Name: "@http.Headers", Fields: []*metadata.Field{{Name: "data", Type: "Map[String, @http.Header?]"}}},
		"@http.Headers?": {Name: "@http.Headers?", Fields: []*metadata.Field{{Name: "data", Type: "Map[String, @http.Header?]"}}},
		"@http.Request": {Name: "@http.Request", Fields: []*metadata.Field{
			{Name: "url", Type: "String"},
			{Name: "method_", Type: "String"},
			{Name: "headers", Type: "@http.Headers?"},
			{Name: "body", Type: "Array[Byte]"},
		}},
		"@http.Request?": {
			Name: "@http.Request?",
			Fields: []*metadata.Field{
				{Name: "url", Type: "String"},
				{Name: "method_", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		// "@http.RequestOptions": {Name: "@http.RequestOptions"},
		"@http.Response": {
			Name: "@http.Response",
			Fields: []*metadata.Field{
				{Name: "status", Type: "UInt16"},
				{Name: "status_text", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		// "@http.Response raise Error": {
		// 	Name: "@http.Response raise Error",
		// 	Fields: []*metadata.Field{
		// 		{Name: "status", Type: "UInt16"},
		// 		{Name: "status_text", Type: "String"},
		// 		{Name: "headers", Type: "@http.Headers?"},
		// 		{Name: "body", Type: "Array[Byte]"},
		// 	},
		// },
		"@http.Response?": {
			Name: "@http.Response?",
			Fields: []*metadata.Field{
				{Name: "status", Type: "UInt16"},
				{Name: "status_text", Type: "String"},
				{Name: "headers", Type: "@http.Headers?"},
				{Name: "body", Type: "Array[Byte]"},
			},
		},
		// "@http.T":                     {Name: "@http.T"},
		// "@http.T raise Error":               {Name: "@http.T raise Error"},
		"@models.ModelInfo": {
			Name:   "@models.ModelInfo",
			Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "full_name", Type: "String"}},
		},
		"@models.ModelInfo?": {
			Name:   "@models.ModelInfo?",
			Fields: []*metadata.Field{{Name: "name", Type: "String"}, {Name: "full_name", Type: "String"}},
		},
		// "@testutils.CallStack[T]":     {Name: "@testutils.CallStack[T]", Fields: []*metadata.Field{{Name: "items", Type: "Array[Array[@testutils.T]]"}}},
		// "@testutils.T":                {Name: "@testutils.T"},
		// "Array[(String, String)]":     {Name: "Array[(String, String)]"},
		"Array[@http.Header?]": {Name: "Array[@http.Header?]"},
		// "Array[@http.RequestOptions]": {Name: "Array[@http.RequestOptions]"},
		// "Array[@testutils.T]":         {Name: "Array[@testutils.T]"},
		// "Array[Array[@testutils.T]]":  {Name: "Array[Array[@testutils.T]]"},
		// "Array[Array[String]]":        {Name: "Array[Array[String]]"},
		"Array[Byte]":   {Name: "Array[Byte]"},
		"Array[String]": {Name: "Array[String]"},
		// "Bool":                        {Name: "Bool"},
		"Byte": {Name: "Byte"},
		// "FixedArray[Double]":          {Name: "FixedArray[Double]"},
		// "FixedArray[Float]":           {Name: "FixedArray[Float]"},
		// "FixedArray[Int64]":           {Name: "FixedArray[Int64]"},
		// "FixedArray[Int]":             {Name: "FixedArray[Int]"},
		// "FixedArray[UInt64]":          {Name: "FixedArray[UInt64]"},
		// "FixedArray[UInt]":            {Name: "FixedArray[UInt]"},
		// "Float":                       {Name: "Float"},
		// "Double":                      {Name: "Double"},
		"Int":   {Name: "Int"},
		"Int64": {Name: "Int64"},
		// "Json":                       {Name: "Json"},
		"Map[String, @http.Header?]": {Name: "Map[String, @http.Header?]"},
		// "Map[String, Array[String]]": {Name: "Map[String, Array[String]]"},
		// "Map[String, String]":        {Name: "Map[String, String]"},
		"Quote": {
			Name:   "Quote",
			Fields: []*metadata.Field{{Name: "q", Type: "String"}, {Name: "a", Type: "String"}, {Name: "h", Type: "String"}},
		},
		"Quote raise Error": {
			Name:   "Quote raise Error",
			Fields: []*metadata.Field{{Name: "q", Type: "String"}, {Name: "a", Type: "String"}, {Name: "h", Type: "String"}},
		},
		"String":             {Name: "String"},
		"String raise Error": {Name: "String raise Error"},
		// "UInt":   {Name: "UInt"},
		"UInt16": {Name: "UInt16"},
		// "UInt64": {Name: "UInt64"},
	},
}
