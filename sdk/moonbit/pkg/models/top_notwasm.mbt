///|
pub fn host_get_model_info(_model_name : String) -> ModelInfo!Error {
  fail!("not wasm")
}

///|
pub fn host_invoke_model(_model_name : String, _input : String) -> String!Error {
  fail!("not wasm")
}
