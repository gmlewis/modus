// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///| These types are used by the example functions in the main.go file.
pub struct Quote {
  q : String
  a : String
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct Image {
  contentType : String
  data : Array[Int] // Array[Byte] is not supported by FromJson
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct Issue {
  title : String
  body : String
  html_url : String?
} derive(Show, Eq, FromJson, ToJson)
