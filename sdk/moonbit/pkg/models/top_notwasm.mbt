///|
pub fn host_get_model_info(model_name : String) -> ModelInfo!Error {
  if model_name != "test" {
    raise ModelsError("bad model name: '\{model_name}', want 'test'")
  }
  { name: model_name, full_name: model_name + "-mock-model" }
}

///|
pub fn host_invoke_model(_model_name : String, _input : String) -> String!Error {
  @ffi.silence_warning()
  fail!("not wasm")
}
