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
      #|        UNWIND $person.friends AS friend_name
      #|        MATCH (friend:Person {name: friend_name})
      #|        MERGE (p)-[:KNOWS]->(friend)
      ,
      { "person": person },
    )

  }
  "People and relationships created successfully"
}

///|
pub struct Person {
  name : String
  age : Int
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
  let persons = []
  for record in response.records {
    let node : @neo4j.Node = @neo4j.get_record_value!(record, "friend")
    let name : String = @neo4j.get_property!(node, "name")
    let age : Int = @neo4j.get_property!(node, "age")
    persons.push({ name, age })
  }
  persons
}

///|
pub fn get_alice_friends_under_40_ages() -> Array[Int]!Error {
  let response = @neo4j.execute_query!(
    CONNECTION,
    #|MATCH (p:Person {name: $name})-[:KNOWS]-(friend:Person)
    #|WHERE friend.age < $age
    #|RETURN friend.age AS age
    ,
    { "name": "Alice", "age": 40 },
  )
  let ages = []
  for record in response.records {
    let age : Int = @neo4j.get_record_value!(record, "age")
    println("pushing age=\{age}")
    ages.push(age)
  }
  ages
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
