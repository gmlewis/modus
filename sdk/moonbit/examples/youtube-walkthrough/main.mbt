///|
pub fn generate_text(instruction : String, prompt : String) -> String!Error {
  let model : @openai.ChatModel = @models.get_model!("text-generator")
  let input = model.create_input([
    @openai.SystemMessage::new(instruction),
    @openai.UserMessage::new(prompt),
  ])
  input.temperature = 0.7
  let output = model.invoke!(input)
  output.choices[0].message.content
}

///|
pub struct Quote {
  // plaintext quote
  q : String
  // author of the quote
  a : String
  // html version of the quote
  h : String
} derive(Show, Eq, ToJson, FromJson)

///|
priv type! QuoteError String derive(Show, Eq)

///|
pub fn get_random_quote() -> Quote!Error {
  let request = @http.Request::new("https://zenquotes.io/api/random")
  let response = @http.fetch!(request)
  guard response.ok() else { raise QuoteError("Failed to fetch quote") }
  let quotes : Array[Quote] = response.json!()
  quotes[0]
}

///|
pub fn say_hello(name~ : String = "World") -> String {
  "Hello, \{name}!"
}
