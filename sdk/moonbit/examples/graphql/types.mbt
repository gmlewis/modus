// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
pub struct Country {
  code : String
  name : String
  capital : String?
} derive(Show, Eq, ToJson)

///|
pub impl @json.FromJson for Country with from_json(json, _path) {
  let json = json.as_object().unwrap()
  let code : String = @json.from_json!(json.get("code").unwrap())
  let name : String = @json.from_json!(json.get("name").unwrap())
  let capital : String? = match json.get("capital").unwrap() {
    String(capital) => Some(capital)
    _ => None
  }
  { code, name, capital }
}

///|
pub struct CountriesResponse {
  countries : Array[Country]
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct CountryResponse {
  country : Country?
} derive(Show, Eq, ToJson)

///|
pub impl @json.FromJson for CountryResponse with from_json(json, _path) {
  let json = json.as_object().unwrap()
  let country : Country? = match json.get("country").unwrap() {
    Null => None
    country => Some(@json.from_json!(country))
  }
  { country, }
}
