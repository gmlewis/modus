///| Execute a SQL query with parameters
pub fn host_execute_query(
  host_name : String,
  db_type : String,
  statement : String,
  params_json : String
) -> HostQueryResponse? {
  let return_area = @ffi.malloc(32)
  wasmImportExecuteQuery(
    @ffi.str2ptr(host_name),
    host_name.iter().count(),
    @ffi.str2ptr(db_type),
    db_type.iter().count(),
    @ffi.str2ptr(statement),
    statement.iter().count(),
    @ffi.str2ptr(params_json),
    params_json.iter().count(),
    return_area,
  )
  let lifted12 : HostQueryResponse? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      let lifted : String? = match @ffi.load8_u(return_area + 4) {
        0 => Option::None
        1 => {
          ignore(@ffi.load32(return_area + 12))
          let result = @ffi.ptr2str(@ffi.load32(return_area + 8))
          Option::Some(result)
        }
        _ => panic()
      }
      let lifted11 : String? = match @ffi.load8_u(return_area + 16) {
        0 => Option::None
        1 => {
          ignore(@ffi.load32(return_area + 24))
          let result10 = @ffi.ptr2str(@ffi.load32(return_area + 20))
          Option::Some(result10)
        }
        _ => panic()
      }
      Option::Some(HostQueryResponse::{
        error: lifted,
        result_json: lifted11,
        rows_affected: @ffi.load32(return_area + 28).reinterpret_as_uint(),
      })
    }
    _ => panic()
  }
  ignore(host_name)
  ignore(db_type)
  ignore(statement)
  ignore(params_json)
  @ffi.free(return_area)
  return lifted12
}
