///|
pub struct Quote {
  quote : String
  author : String
} derive(Show, Eq, ToJson, FromJson)

///|
priv suberror QuoteError String derive(Show, Eq)

///|
pub fn get_random_quote() -> Quote raise Error {
  let request = @http.Request::new("https://api.quotable.io/random")
  let response = request.fetch()
  guard response.ok() else { raise QuoteError("Failed to fetch quote") }
  let quotes : Array[Quote] = response.json()
  quotes[0]
}

///|
pub fn say_hello(name~ : String = "World") -> String {
  "Hello, \{name}!"
}
