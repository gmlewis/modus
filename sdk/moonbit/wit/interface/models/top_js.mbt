// This files is provided so that `moon test --target js` can be run without errors.

///| Model information record
pub(all) struct ModelInfo {
  name : String
  full_name : String
} derive(Show, Eq, FromJson, ToJson)

///| Get information about a specific model
pub fn get_model_info(model_name : String) -> ModelInfo? {
  Some({ name: model_name, full_name: "\{model_name}-mock-model" })
}

///| Invoke a model with JSON-serialized input
pub fn invoke_model(model_name : String, input : String) -> String? {
  Some(input)
}
