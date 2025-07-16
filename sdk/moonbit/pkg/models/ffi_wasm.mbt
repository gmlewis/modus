///|modus:import modus_models getModelInfo(model_name : String) -> ModelInfo?
fn wasmImportGetModelInfo(model_name : Int) -> Int = "modus_models" "getModelInfo"

///|modus:import modus_models invokeModel(model_name : String, input : String) -> String
fn wasmImportInvokeModel(model_name : Int, input : Int) -> Int = "modus_models" "invokeModel"

///|
pub fn host_get_model_info(model_name : String) -> ModelInfo!Error {
  let model_name : Int = @ffi.str2ptr(model_name) - 8
  let model_info = wasmImportGetModelInfo(model_name)
  if model_info == 0 {
    raise ModelsError(
      "getModelInfo failed. Check the logs for more information.",
    )
  }
  let model_info : ModelInfo? = cast(model_info)
  model_info.unwrap()
}

///|
pub fn host_invoke_model(model_name : String, input : String) -> String!Error {
  let model_name : Int = @ffi.str2ptr(model_name) - 8
  let input : Int = @ffi.str2ptr(input) - 8
  let response = wasmImportInvokeModel(model_name, input)
  if response == 0 {
    raise ModelsError(
      "invokeModel failed. Check the logs for more information.",
    )
  }
  @ffi.ptr2str(response + 8)
}

///|
fn cast[A, B](a : A) -> B = "%identity"
