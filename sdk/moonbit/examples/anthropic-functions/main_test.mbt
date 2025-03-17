///|
test "model input serialization" {
  let model : @anthropic.MessagesModel = @models.get_model!("test")
  let input = model.create_input([
    @anthropic.Message::new_user(
      @anthropic.String("what is the stock price of MSFT?"),
    ),
  ])
  inspect!(
    input.to_json(),
    content=
      #|Object({"model": String("test-mock-model"), "messages": Array([Object({"role": String("user"), "content": String("what is the stock price of MSFT?")})]), "max_tokens": Number(0), "system": String("")})
    ,
  )
}

///|
test "model response deserialization" {
  let response =
    #|{"id":"msg_01TePXAo2jkt1hVi1NkswTZU","type":"message","role":"assistant","model":"claude-3-opus-20240229","content":[{"type":"tool_use","id":"toolu_015VbuAApqwhQ34qxpgyXVv6","name":"stock_price","input":{"symbol":"MSFT"}}],"stop_reason":"tool_use","stop_sequence":null,"usage":{"input_tokens":403,"cache_creation_input_tokens":0,"cache_read_input_tokens":0,"output_tokens":35}}
  let json = @json.parse!(response)
  let output : @anthropic.MessagesModelOutput = @json.from_json!(json)
  inspect!(
    output,
    content=
      #|{id: "msg_01TePXAo2jkt1hVi1NkswTZU", type_: "message", role: "assistant", content: [{type_: "tool_use", text: None, id: Some("toolu_015VbuAApqwhQ34qxpgyXVv6"), name: Some("stock_price"), input: Some(Object({"symbol": String("MSFT")})), tool_use_id: None, is_error: None, content: None}], model: "claude-3-opus-20240229", stop_reason: "tool_use", stop_sequence: "", usage: {input_tokens: 403, output_tokens: 35}}
    ,
  )
}
