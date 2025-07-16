// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
pub fn get_embeddings_for_text_with_openai(text : String) -> Array[Float]!Error {
  let results = get_embeddings_for_texts_with_openai!([text])
  results[0]
}

///|
pub fn get_embeddings_for_texts_with_openai(
  texts : Array[String]
) -> Array[Array[Float]]!Error {
  let model : @experimental.EmbeddingsModel = @models.get_model!(
    "openai-embeddings",
  )
  let input = model.create_input!(texts)
  let output = model.invoke!(input)
  output.predictions
}
