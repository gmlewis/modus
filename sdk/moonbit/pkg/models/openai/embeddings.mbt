// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| Provides input and output types that conform to the OpenAI Embeddings API,
/// as described in the [API Reference] docs.
///
/// [API Reference]: https://platform.openai.com/docs/api-reference/embeddings
pub(all) type EmbeddingsModel @models.ModelBase derive (
  Show,
  Eq,
  FromJson,
  ToJson,
)

// type EmbeddingsModel struct {
// 	embeddingsModelBase
// }
//
// type embeddingsModelBase = models.ModelBase[EmbeddingsModelInput, EmbeddingsModelOutput]

///| The input object for the OpenAI Embeddings API.
pub(all) struct EmbeddingsModelInput {
  // The name of the model to use to generate the embeddings.
  //
  // Must be the exact string expected by the model provider.
  // For example, "text-embedding-3-small".
  model : String

  // The input content to vectorize.
  input : Content

  // The format for the output embeddings.
  // The default ("") is equivalent to [EncodingFormatFloat], which is currently the only supported format.
  encoding_format : EncodingFormat?

  // The maximum number of dimensions for the output embeddings.
  // The default (0) indicates that the model's default number of dimensions will be used.
  dimensions : Int?

  // The user ID to associate with the request, as described in the [documentation].
  // If not specified, the request will be anonymous.
  //
  // [documentation]: https://platform.openai.com/docs/guides/safety-best-practices/end-user-ids
  user : String?
} derive(Show, Eq, FromJson, ToJson)

///| The output object for the OpenAI Embeddings API.
pub(all) struct EmbeddingsModelOutput {
  // The name of the output object type returned by the API.
  // This will always be "list".
  object : String

  // The name of the model used to generate the embeddings.
  // In most cases, this will match the requested model field in the input.
  model : String

  // The usage statistics for the request.
  usage : Usage

  // The output vector embeddings data.
  data : Array[Embedding]
} derive(Show, Eq, FromJson, ToJson)

///| The encoding format for the output embeddings.
pub(all) type EncodingFormat String derive(Show, Eq, FromJson, ToJson)

///| The output embeddings are encoded as an array of floating-point numbers.
pub const ENCODING_FORMAT_FLOAT = "float"

///| The output embeddings are encoded as a base64-encoded string,
/// containing an binary representation of an array of floating-point numbers.
///
/// NOTE: This format is not currently supported.
pub const ENCODING_FORMAT_BASE64 = "base64"

///| The output vector embeddings data.
pub(all) struct Embedding {
  // The name of the output object type returned by the API.
  // This will always be "embedding".
  object : String

  // The index of the input text that corresponds to this embedding.
  // Used when requesting embeddings for multiple texts.
  index : Int

  // The vector embedding of the input text.
  embedding : Array[Double]
} derive(Show, Eq, FromJson, ToJson)

///| Creates an input object for the OpenAI Embeddings API.
///
/// The content parameter can be any of:
///   - A string representing the text to vectorize.
///   - A slice of strings representing multiple texts to vectorize.
///   - A slice of integers representing pre-tokenized text to vectorize.
///   - A slice of slices of integers representing multiple pre-tokenized texts to vectorize.
///
/// NOTE: The input content must not exceed the maximum token limit of the model.
pub fn create_input(
  self : EmbeddingsModel,
  content : Content
) -> EmbeddingsModelInput {
  {
    model: self.info.unwrap().name,
    input: content,
    encoding_format: None,
    dimensions: None,
    user: None,
  }
}

///|
pub(all) enum Content {
  String(String)
  ArrayString(Array[String])
  // https://github.com/moonbitlang/core/pull/1345
  // TODO: ArrayUInt(Array[UInt])
  // TODO: ArrayByte(Array[Byte])
  // TODO: ArrayUInt64(Array[UInt64])
  ArrayInt(Array[Int])
  ArrayInt64(Array[Int64])
} derive(Show, Eq, FromJson, ToJson)
