///|modus:import modus_graphql_client executeQuery(host_name: String, statement: String, variables: String) -> String?
fn wasmImportExecuteQuery(
  host_name : Int,
  statement : Int,
  variables : Int
) -> Int = "modus_graphql_client" "executeQuery"

///|
fn[A, B] cast(a : A) -> B = "%identity"

///| Execute a GraphQL query with variables
pub fn host_execute_query(
  host_name : String,
  statement : String,
  variables : String
) -> String raise Error {
  let host_name : Int = cast(host_name)
  let statement : Int = cast(statement)
  let variables : Int = cast(variables)
  let response : Int = wasmImportExecuteQuery(host_name, statement, variables)
  if response == 0 {
    raise GraphQLError(
      "Failed to execute the query. Check logs for more information.",
    )
  }
  let response_str : String? = cast(response)
  response_str.unwrap()
}
