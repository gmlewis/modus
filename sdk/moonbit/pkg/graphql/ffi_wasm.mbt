///| Execute a GraphQL query with variables
pub fn host_execute_query(
  host_name : String,
  statement : String,
  variables : String
) -> String? {
  let return_area = @ffi.malloc(12)
  wasmImportExecuteQuery(
    @ffi.str2ptr(host_name),
    host_name.iter().count(),
    @ffi.str2ptr(statement),
    statement.iter().count(),
    @ffi.str2ptr(variables),
    variables.iter().count(),
    return_area,
  )
  let lifted : String? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      ignore(@ffi.load32(return_area + 8))
      let result = @ffi.ptr2str(@ffi.load32(return_area + 4))
      Option::Some(result)
    }
    _ => panic()
  }
  ignore(host_name)
  ignore(statement)
  ignore(variables)
  @ffi.free(return_area)
  return lifted
}
