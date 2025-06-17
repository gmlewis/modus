// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
const HOST_NAME = "countries-api"

///| This application uses a public GraphQL API to retrieve information about countries.
/// We are using the API available at: https://github.com/trevorblades/countries
/// The `graphql` module from the Modus SDK allows us to interact with the API.
///
/// Define the host name of the API to be used in the requests.
/// Must match one of the http connections defined in the modus.json manifest file.

///| Function to list all countries returned by the API
pub fn countries() -> Array[Country] raise Error {
  let statement =
    #|query {
    #|  countries{
    #|    code
    #|    name
    #|    capital
    #|  }
    #|}
  // Execute the GraphQL query using the host name and query statement
  // The API returns a response of type `CountriesResponse` containing an array of `Country` objects
  let response : @graphql.Response[CountriesResponse] = @graphql.execute(
    HOST_NAME,
    statement,
  )
  if response.errors is Some(errors) && errors.length() > 0 {
    raise @graphql.GraphQLError("Failed to retrieve countries information")
  }
  response.data.unwrap().countries
}

///| Function to retrieve a specific country by its unique code
pub fn get_country_by_code(code : String) -> Country? raise Error {
  let statement =
    #|query ($code: ID!){
    #|  country(code: $code){
    #|    code
    #|    name
    #|    capital
    #|  }
    #|}
  // Create a vars map to pass in the `code` parameter
  let vars = { "code": code.to_json() }
  // Execute the GraphQL query using the host name, statement, and vars map
  // The API returns a response of type `CountryResponse` containing a single `Country` object
  let response : @graphql.Response[CountryResponse] = @graphql.execute(
    HOST_NAME,
    statement,
    vars~,
  )
  if response.errors is Some(errors) && errors.length() > 0 {
    raise @graphql.GraphQLError("Failed to retrieve country information")
  }
  response.data.unwrap().country
}
