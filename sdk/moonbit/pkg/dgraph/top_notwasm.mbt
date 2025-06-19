///| Executes a DQL query or mutation on the Dgraph database.
pub fn host_execute_query(
  _host_name : String,
  _request : Request
) -> Response? raise Error {
  @ffi.silence_warning()
  fail("not wasm")
}

///| Alters the schema of the dgraph database
pub fn host_alter_schema(
  _host_name : String,
  _schema : String
) -> Unit raise Error {
  fail("not wasm")
}

///| Drops an attribute from the schema.
pub fn host_drop_attribute(
  _host_name : String,
  _attr : String
) -> Unit raise Error {
  fail("not wasm")
}

///| Drops all data from the database.
pub fn host_drop_all_data(_host_name : String) -> Unit raise Error {
  fail("not wasm")
}
