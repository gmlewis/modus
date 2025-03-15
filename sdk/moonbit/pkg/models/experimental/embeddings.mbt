// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| A model that returns embeddings for a list of text strings.
///
/// NOTE: This model interface is experimental and may change in the future.
/// It is primarily intended for use with with embedding models hosted on Hypermode.
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

///| The input object for the embeddings model.
pub(all) struct EmbeddingsModelInput {
  // A list of one or more text strings to create vector embeddings for.
  instances : Array[String]
} derive(Show, Eq, FromJson, ToJson)

///| The output object for the embeddings model.
pub(all) struct EmbeddingsModelOutput {
  // A list of vector embeddings that correspond to each input text string.
  predictions : Array[Array[Double]]
} derive(Show, Eq, FromJson, ToJson)

///| Creates an input object for the embeddings model.
///
/// The content parameter is a list of one or more text strings to create vector embeddings for.
pub fn EmbeddingsModel::create_input(
  _self : EmbeddingsModel,
  content : Array[String]
) -> EmbeddingsModelInput!Error {
  if content.is_empty() {
    raise ExperimentalError("at least one text string must be provided")
  }
  { instances: content }
}
