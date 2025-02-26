// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
// The name of the Neo4j connection, as specified in the modus.json manifest
const CONNECTION = "my-database"

///|
pub fn create_people_and_relationships() -> String!Error {
  let people : Array[Json] = [
    { "name": "Alice", "age": 42, "friends": ["Bob", "Peter", "Anna"] },
    { "name": "Bob", "age": 19 },
    { "name": "Peter", "age": 50 },
    { "name": "Anna", "age": 30 },
  ]
  for person in people {
    let _ = @neo4j.execute_query!(
      CONNECTION,
      "MERGE (p:Person {name: $person.name, age: $person.age})",
      { "person": person },
    )

  }
  for person in people.filter(fn(p) { p.value("friends") != None }) {
    let _ = @neo4j.execute_query!(
      CONNECTION,
      #|MATCH (p:Person {name: $person.name})
      #|UNWIND $person.friends AS friend_name
      #|MATCH (friend:Person {name: friend_name})
      #|MERGE (p)-[:KNOWS]->(friend)
      ,
      { "person": person },
    )

  }
  "People and relationships created successfully"
}

///|
pub struct Person {
  name : String
  age : Int64
} derive(Show, Eq, ToJson, FromJson)

///|
pub fn get_alice_friends_under_40() -> Array[Person]!Error {
  let response = @neo4j.execute_query!(
    CONNECTION,
    #|MATCH (p:Person {name: $name})-[:KNOWS]-(friend:Person)
    #|WHERE friend.age < $age
    #|RETURN friend
    ,
    { "name": "Alice", "age": 40 },
  )
  let records = response.records
  records.filter_map(fn(record) {
    let node = try {
      @neo4j.get_record_value!(record, "friend")
    } catch {
      _ => return None
    }
    match node {
      Map({ "name": String(name), "age": Number(age), .. }) =>
        Some({ name, age: age.to_int64() })
      _ => None
    }
  })
}

///|
pub fn get_alice_friends_under_40_ages() -> Array[Int64]!Error {
  let response = @neo4j.execute_query!(
    CONNECTION,
    #|MATCH (p:Person {name: $name})-[:KNOWS]-(friend:Person)
    #|WHERE friend.age < $age
    #|RETURN friend.age AS age
    ,
    { "name": "Alice", "age": 40 },
  )
  let records = response.records
  records.filter_map(fn(record) {
    let node = try {
      @neo4j.get_record_value!(record, "friend")
    } catch {
      _ => return None
    }
    match node {
      Map({ "age": Number(age), .. }) => Some(age.to_int64())
      _ => None
    }
  })
}

///|
pub fn delete_all_nodes() -> String!Error {
  let _ = @neo4j.execute_query!(
    CONNECTION,
    #|MATCH (n)
    #|DETACH DELETE n
    ,
    {},
  )
  "All nodes deleted"
}
