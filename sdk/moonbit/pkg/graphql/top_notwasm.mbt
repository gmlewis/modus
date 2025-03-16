///| Execute a GraphQL query with variables
pub fn host_execute_query(
  _host_name : String,
  _statement : String,
  _variables : String
) -> String? {
  @ffi.silence_warning()
  abort("not wasm")
}
