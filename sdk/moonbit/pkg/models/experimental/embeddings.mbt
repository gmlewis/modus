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

///|
pub fn EmbeddingsModel::invoke(
  self : EmbeddingsModel,
  input : EmbeddingsModelInput
) -> EmbeddingsModelOutput!Error {
  let input_json = input.to_json()
  let output_json = self._.invoke!(input_json)
  // println("got output_json=\{output_json}")
  @json.from_json!(output_json)
}

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
  predictions : Array[Array[Float]]
} derive(Show, Eq, ToJson) // Float does not implement trait @json.FromJson

///|
pub impl @json.FromJson for EmbeddingsModelOutput with from_json(json, _path) {
  let json : Map[String, Json] = json.as_object().unwrap()
  let predictions : Array[Array[Double]] = @json.from_json!(
    json.get("predictions").unwrap(),
  )
  // TODO: There must be a better way to do this.
  let predictions = Array::makei(predictions.length(), fn(i) {
    let prediction = predictions.get(i).unwrap()
    Array::makei(prediction.length(), fn(j) {
      prediction.get(j).unwrap().to_float()
    })
  })
  { predictions, }
}

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
