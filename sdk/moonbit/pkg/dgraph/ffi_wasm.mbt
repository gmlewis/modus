// Generated by `wit-bindgen` 0.36.0. DO NOT EDIT!

///|modus:import modus_dgraph_client executeQuery(host_name : String, request : Request?) -> Response?!Error
fn wasmImportExecuteQuery(host_name : Int, request : Int) -> Int = "modus_dgraph_client" "executeQuery"

///|modus:import modus_dgraph_client alterSchema(host_name : String, schema : String) -> String!Error
fn wasmImportAlterSchema(host_name : Int, schema : Int) -> Int = "modus_dgraph_client" "alterSchema"

///|modus:import modus_dgraph_client dropAttr(host_name : String, attr : String) -> String!Error
fn wasmImportDropAttribute(host_name : Int, attr : Int) -> Int = "modus_dgraph_client" "dropAttribute"

///|modus:import modus_dgraph_client dropAll(host_name : String) -> String!Error
fn wasmImportDropAllData(host_name : Int) -> Int = "modus_dgraph_client" "dropAllData"

///|
fn cast[A, B](a : A) -> B = "%identity"

///| Executes a DQL query or mutation on the Dgraph database.
pub fn host_execute_query(
  host_name : String,
  request : Request
) -> Response?!Error {
  let host_name_ptr = @ffi.str2ptr(host_name) - 8
  let request_ptr : Int = cast(request)
  println(
    "GML: dgraph/ffi_wasm.mbt: CALLING wasmImportExecuteQuery(host_name_ptr=\{host_name_ptr}, request_ptr=\{request_ptr}",
  )
  let result = wasmImportExecuteQuery(host_name_ptr, request_ptr)
  println(
    "GML: dgraph/ffi_wasm.mbt: wasmImportExecuteQuery returned result=\{result}",
  )
  if result == 0 {
    raise DgraphError("failed to execute query")
  }
  let result : Response? = cast(result)
  println("GML: dgraph/ffi_wasm.mbt: host_execute_query: result=\{result}")
  result
}

///| Alters the schema of the dgraph database
pub fn host_alter_schema(host_name : String, schema : String) -> Unit!Error {
  let host_name_ptr = @ffi.str2ptr(host_name) - 8
  let schema_ptr = @ffi.str2ptr(schema) - 8
  println(
    "GML: dgraph/ffi_wasm.mbt: CALLING wasmImportAlterSchema(host_name_ptr=\{host_name_ptr}, schema_ptr=\{schema_ptr}",
  )
  let result = wasmImportAlterSchema(host_name_ptr, schema_ptr)
  println(
    "GML: dgraph/ffi_wasm.mbt: wasmImportAlterSchema returned result=\{result}",
  )
  if result == 0 {
    raise DgraphError("failed to alter the schema")
  }
}

///| Drops an attribute from the schema.
pub fn host_drop_attribute(host_name : String, attr : String) -> Unit!Error {
  let host_name_ptr = @ffi.str2ptr(host_name) - 8
  let attr_ptr = @ffi.str2ptr(attr) - 8
  println(
    "GML: dgraph/ffi_wasm.mbt: CALLING wasmImportDropAttribute(host_name_ptr=\{host_name_ptr}, attr_ptr=\{attr_ptr}",
  )
  let result = wasmImportDropAttribute(host_name_ptr, attr_ptr)
  println(
    "GML: dgraph/ffi_wasm.mbt: wasmImportDropAttribute returned result=\{result}",
  )
  if result == 0 {
    raise DgraphError("failed to drop the attribute")
  }
}

///| Drops all data from the database.
pub fn host_drop_all_data(host_name : String) -> Unit!Error {
  let host_name_ptr = @ffi.str2ptr(host_name) - 8
  println(
    "GML: dgraph/ffi_wasm.mbt: CALLING wasmImportDropAllData(host_name_ptr=\{host_name_ptr}",
  )
  let result = wasmImportDropAllData(host_name_ptr)
  println(
    "GML: dgraph/ffi_wasm.mbt: wasmImportDropAllData returned result=\{result}",
  )
  if result == 0 {
    raise DgraphError("failed to drop all data")
  }
}
