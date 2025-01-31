///| Execute a Cypher query with parameters
pub fn execute_query(
  host_name : String,
  db_name : String,
  query : String,
  parameters_json : String
) -> EagerResult? {
  let return_area = @ffi.malloc(20)
  wasmImportExecuteQuery(
    @ffi.str2ptr(host_name),
    host_name.iter().count(),
    @ffi.str2ptr(db_name),
    db_name.iter().count(),
    @ffi.str2ptr(query),
    query.iter().count(),
    @ffi.str2ptr(parameters_json),
    parameters_json.iter().count(),
    return_area,
  )
  let lifted : EagerResult? = match @ffi.load8_u(return_area + 0) {
    0 => Option::None
    1 => {
      let array : Array[String] = []
      for index = 0; index < @ffi.load32(return_area + 8); index = index + 1 {
        let base3 = @ffi.load32(return_area + 4) + index * 8
        ignore(@ffi.load32(base3 + 4))
        let result = @ffi.ptr2str(@ffi.load32(base3 + 0))
        array.push(result)
      }
      @ffi.free(@ffi.load32(return_area + 4))
      let array16 : Array[ResultRecord] = []
      for index17 = 0
          index17 < @ffi.load32(return_area + 16)
          index17 = index17 + 1 {
        let base5 = @ffi.load32(return_area + 12) + index17 * 16
        let array9 : Array[String] = []
        for index10 = 0; index10 < @ffi.load32(base5 + 4); index10 = index10 + 1 {
          let base7 = @ffi.load32(base5 + 0) + index10 * 8
          ignore(@ffi.load32(base7 + 4))
          let result8 = @ffi.ptr2str(@ffi.load32(base7 + 0))
          array9.push(result8)
        }
        @ffi.free(@ffi.load32(base5 + 0))
        let array14 : Array[String] = []
        for index15 = 0
            index15 < @ffi.load32(base5 + 12)
            index15 = index15 + 1 {
          let base12 = @ffi.load32(base5 + 8) + index15 * 8
          ignore(@ffi.load32(base12 + 4))
          let result13 = @ffi.ptr2str(@ffi.load32(base12 + 0))
          array14.push(result13)
        }
        @ffi.free(@ffi.load32(base5 + 8))
        array16.push(ResultRecord::{ values: array9, keys: array14 })
      }
      @ffi.free(@ffi.load32(return_area + 12))
      Option::Some(EagerResult::{ keys: array, records: array16 })
    }
    _ => panic()
  }
  ignore(host_name)
  ignore(db_name)
  ignore(query)
  ignore(parameters_json)
  @ffi.free(return_area)
  return lifted
}
