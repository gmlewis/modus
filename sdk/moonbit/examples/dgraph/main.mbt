// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
const CONNECTION = "dgraph"

///|
pub struct Person {
  uid : String // `json:"uid,omitempty"`
  firstName : String // `json:"firstName,omitempty"`
  lastName : String // `json:"lastName,omitempty"`
  dType : Array[String] // `json:"dgraph.type,omitempty"`
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct PeopleData {
  people : Array[Person] // `json:"people"`
} derive(Show, Eq, FromJson, ToJson)

///|
pub fn drop_all() -> String raise Error {
  @dgraph.drop_all(CONNECTION)
  "Success"
}

///|
pub fn drop_attr(attr : String) -> String raise Error {
  @dgraph.drop_attr(CONNECTION, attr)
  "Success"
}

///|
pub fn alter_schema() -> String raise Error {
  let schema =
    #|firstName: string @index(term) .
    #|lastName: string @index(term) .
    #|dgraph.type: [string] @index(exact) .
    #|
    #|type Person {
    #|	firstName
    #|	lastName
    #|}
  @dgraph.alter_schema(CONNECTION, schema)
  "Success"
}

///|
pub fn query_people() -> Array[Person] raise Error {
  let query =
    #|{
    #|  people(func: type(Person)) {
    #|    uid
    #|    firstName
    #|    lastName
    #|    dgraph.type
    #|  }
    #|}
  let response = @dgraph.execute_query(CONNECTION, @dgraph.new_query(query))
  let people_data : PeopleData = @json.from_json(@json.parse(response.json))
  people_data.people
}

///|
pub fn query_specific_person(
  first_name : String,
  last_name : String
) -> Person? raise Error {
  let query =
    #|query queryPerson($firstName: string, $lastName: string) {
    #|  people(func: eq(firstName, $firstName)) @filter(eq(lastName, $lastName)) {
    #|	  uid
    #|	  firstName
    #|	  lastName
    #|	  dgraph.type
    #|  }
    #|}
  let variables = { "$firstName": first_name, "$lastName": last_name }
  let response = @dgraph.execute_query(
    CONNECTION,
    @dgraph.new_query(query, variables~),
  )
  let people_data : PeopleData = @json.from_json(@json.parse(response.json))
  if people_data.people.is_empty() {
    None
  } else {
    Some(people_data.people[0])
  }
}

///|
pub fn add_person_as_rdf(
  first_name : String,
  last_name : String
) -> Map[String, String] raise Error {
  let mutation =
    $|_:user1 <firstName> "\{@dgraph.escape_rdf(first_name)}" .
    $|_:user1 <lastName> "\{@dgraph.escape_rdf(last_name)}" .
    #|_:user1 <dgraph.type> "Person" .
  let response = @dgraph.execute_mutations(CONNECTION, [
    @dgraph.set_nquads(mutation),
  ])
  response.uids
}

///|
pub fn add_person_as_json(
  firstName : String,
  lastName : String
) -> Map[String, String] raise Error {
  let person = { uid: "_:user1", firstName, lastName, dType: ["Person"] }
  let data = person.to_json().stringify()
  let response = @dgraph.execute_mutations(CONNECTION, [@dgraph.set_json(data)])
  response.uids
}

///|
pub fn upsert_person(
  name_to_change_from : String,
  name_to_change_to : String
) -> Map[String, String] raise Error {
  let query =
    #|query findPerson($name: string) {
    #|  person as var(func: eq(firstName, $name))
    #|}
  let variables = { "$name": name_to_change_from }
  let mutation =
    $|uid(person) <firstName> "\{@dgraph.escape_rdf(name_to_change_to)}" .
  let response = @dgraph.execute_query(CONNECTION, { query, variables }, mutations=[
    @dgraph.set_nquads(mutation),
  ])
  response.uids
}
