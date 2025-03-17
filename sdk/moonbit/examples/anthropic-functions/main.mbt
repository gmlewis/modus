// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
pub struct StockPriceInput {
  symbol : String
} derive(Show, Eq, FromJson, ToJson)

///| This model name should match the one defined in the modus.json manifest file.
const MODEL_NAME = "text-generator"

///|
pub fn get_stock_price(company : String, use_tools : Bool) -> String!Error {
  let model : @anthropic.MessagesModel = @models.get_model!(MODEL_NAME)

  //
  let input = model.create_input([
    @anthropic.Message::new_user(
      @anthropic.String("what is the stock price of \{company}?"),
    ),
  ])

  // For Anthropic, system is passed as parameter to the invoke, not as a message
  input.system = "You are a helpful assistant. Do not answer if you do not have up-to-date information."
  // Optional parameters
  input.temperature = 1
  input.max_tokens = 100

  //
  if use_tools {
    input.tools = [
      {
        name: "stock_price",
        input_schema: {
          "type": "object",
          "properties": {
            "symbol": { "type": "string", "description": "The stock symbol" },
          },
          "required": ["symbol"],
        },
        description: Some("gets the stock price of a symbol"),
      },
    ]
    input.tool_choice = @anthropic.tool_choice_tool("stock_price")
  }

  // Here we invoke the model with the input we created.
  let output = model.invoke!(input)

  //
  if output.content.length() != 1 {
    raise @models.ModelsError(
      "unexpected output content length: \{output.content.length()}",
    )
  }
  // If tools are not used, the output will be a text block
  match output.content[0].type_ {
    "text" => return output.content[0].text.unwrap().trim_space()
    "tool_use" => ()
    t => raise @models.ModelsError("unexpected content type: \{t}")
  }

  //
  let tool_use = output.content[0]
  let inputs = tool_use.input.unwrap()

  //
  let parsed_input : StockPriceInput = @json.from_json!(inputs)
  let symbol = parsed_input.symbol
  let stock_price = call_stock_price_api!(symbol)
  "The stock price of \{symbol} is \{stock_price}"
}

///|
pub struct StockPriceAPIResponse {
  global_quote : GlobalQuote
}

///|
pub impl @json.FromJson for StockPriceAPIResponse with from_json(self, path) {
  let obj = self.as_object().unwrap()
  let global_quote : GlobalQuote = @json.from_json!(
    obj.get("Global Quote").unwrap(),
  )
  { global_quote, }
}

///|
pub struct GlobalQuote {
  symbol : String
  price : String
}

///|
pub impl @json.FromJson for GlobalQuote with from_json(self, _path) {
  let obj = self.as_object().unwrap()
  let symbol : String = @json.from_json!(obj.get("01. symbol").unwrap())
  let price : String = @json.from_json!(obj.get("05. price").unwrap())
  { symbol, price }
}

///|
priv type! QuoteError String derive(Show, Eq)

///|
fn call_stock_price_api(symbol : String) -> String!Error {
  let url = "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=\{symbol}"
  let req = @http.Request::new(url)
  let resp = @http.fetch!(req)
  guard resp.ok() else { raise QuoteError("Failed to fetch quote") }
  let data : StockPriceAPIResponse = resp.json!()
  data.global_quote.price
}
