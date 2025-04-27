///| Execute a SQL query with parameters
pub fn host_execute_query(
  _host_name : String,
  _db_type : String,
  _statement : String,
  _params_json : String
) -> HostQueryResponse? {
  @ffi.silence_warning()
  abort("not wasm")
}
