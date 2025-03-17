// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) type! ExperimentalError String derive(Show, Eq, FromJson, ToJson)

///| A model that returns classification results for a list of text strings.
///
/// NOTE: This model interface is experimental and may change in the future.
/// It is primarily intended for use with with classification models hosted on Hypermode.
pub(all) type ClassificationModel @models.ModelBase derive (
  Show,
  Eq,
  FromJson,
  ToJson,
)

///|
pub fn invoke(
  self : ClassificationModel,
  input : ClassificationModelInput
) -> ClassificationModelOutput!Error {
  let input_json = input.to_json()
  let output_json = self._.invoke!(input_json)
  // println("got output_json=\{output_json}")
  @json.from_json!(output_json)
}

// type ClassificationModel struct {
// 	classificationModelBase
// }
//
// type classificationModelBase = models.ModelBase[ClassificationModelInput, ClassificationModelOutput]

///| The input object for the classification model.
pub(all) struct ClassificationModelInput {
  // A list of one or more text strings to classify.
  instances : Array[String]
} derive(Show, Eq, FromJson, ToJson)

///| The output object for the classification model.
pub(all) struct ClassificationModelOutput {
  // A list of prediction results that correspond to each input text string.
  predictions : Array[ClassifierResult]
} derive(Show, Eq, FromJson, ToJson)

///| A classification result for a single text string.
pub(all) struct ClassifierResult {
  // The classification label with the highest confidence.
  label : String

  // The confidence score for the classification label.
  confidence : Float

  // The list of all classification labels with their corresponding probabilities.
  probabilities : Array[ClassifierLabel]
} derive(Show, Eq, FromJson, ToJson)

///| A classification label with its corresponding probability.
pub(all) struct ClassifierLabel {
  // The classification label.
  label : String

  // The probability value.
  probability : Float
} derive(Show, Eq, FromJson, ToJson)

///| Creates an input object for the classification model.
///
/// The content parameter is a list of one or more text strings to classify.
pub fn ClassificationModel::create_input(
  _self : ClassificationModel,
  content : Array[String]
) -> ClassificationModelInput!Error {
  if content.length() == 0 {
    raise ExperimentalError("at least one text string must be provided")
  }
  { instances: content }
}
