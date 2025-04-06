///| Execute a Cypher query with parameters
pub fn host_execute_query(
  _host_name : String,
  _db_name : String,
  _query : String,
  _parameters_json : Map[String, Json]
) -> EagerResult?!Error {
  @ffi.silence_warning()
  fail!("not wasm")
}
