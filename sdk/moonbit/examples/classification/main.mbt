// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
const MODEL_NAME = "my-classifier"

///| This function takes input text and a probability threshold, and returns the
/// classification label determined by the model, if the confidence is above the
/// threshold. Otherwise, it returns an empty string.
pub fn classify_text(text : String, threshold : Float) -> String!Error {
  let predictions = classify!([text])
  let prediction = predictions[0]
  if prediction.confidence < threshold {
    return ""
  }
  prediction.label
}

///| This function takes input text and returns the classification labels and their
/// corresponding probabilities, as determined by the model.
pub fn get_classification_labels(text : String) -> Map[String, Float]!Error {
  let predictions = classify!([text])
  get_labels(predictions[0])
}

///|
pub fn get_labels(
  prediction : @experimental.ClassifierResult
) -> Map[String, Float] {
  prediction.probabilities.fold(init={}, fn(acc, v) {
    acc[v.label] = v.probability
    acc
  })
}

///|
priv type! ClassificationError String derive(Show, Eq)

///| This function is similar to the previous, but allows multiple items to be classified at a time.
pub fn get_multiple_classification_labels(
  ids : Array[String],
  texts : Array[String]
) -> Map[String, Map[String, Float]]!Error {
  guard ids.length() == texts.length() else {
    raise ClassificationError(
      "number of IDs (\{ids.length()}) must match number of texts (\{texts.length()})",
    )
  }
  let predictions = classify!(texts)
  predictions.foldi(init={}, fn(i, acc, v) {
    acc[ids[i]] = get_labels(v)
    acc
  })
}

///|
pub fn classify(
  texts : Array[String]
) -> Array[@experimental.ClassifierResult]!Error {
  let model : @experimental.ClassificationModel = @models.get_model!(MODEL_NAME)
  let input = model.create_input!(texts)
  let output = model.invoke!(input)
  let predictions = output.predictions
  if predictions.length() != texts.length() {
    let word = if texts.length() > 1 { "predictions" } else { "prediction" }
    raise ClassificationError(
      "expected \{texts.length()} \{word}, got \{predictions.length()}",
    )
  }
  predictions
}
