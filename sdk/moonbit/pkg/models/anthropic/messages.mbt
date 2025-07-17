// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| Provides input and output types that conform to the Anthropic Messages API,
/// as described in the [API Reference] docs.
///
/// [API Reference]: https://docs.anthropic.com/en/api/messages
pub(all) type MessagesModel @models.ModelBase derive(Show, Eq, FromJson, ToJson)

///|
pub fn invoke(
  self : MessagesModel,
  input : MessagesModelInput
) -> MessagesModelOutput raise Error {
  let input_json = input.to_json()
  let output_json = self.inner().invoke(input_json)
  // println("got output_json=\{output_json}")
  @json.from_json(output_json)
}

// type MessagesModel struct {
// 	messagesModelBase
// }
//
// type messagesModelBase = models.ModelBase[MessagesModelInput, MessagesModelOutput]

///| The input object for the Anthropic Messages API.
pub(all) struct MessagesModelInput {
  // The model that will complete your prompt.
  //
  // Must be the exact string expected by the model provider.
  // For example, "claude-3-5-sonnet-20240620".
  //
  // See [models](https://docs.anthropic.com/en/docs/models-overview) for additional
  // details and options.
  model : String

  // Input messages.
  //
  // We do not currently support image content blocks, which are available starting with
  // Claude 3 models. This will be added in a future release.
  messages : Array[Message]

  // The maximum number of tokens to generate before stopping.
  //
  // Different models have different maximum values for this parameter. See
  // [models](https://docs.anthropic.com/en/docs/models-overview) for details.
  mut max_tokens : Int

  // A `Metadata` object describing the request.
  metadata : Metadata?

  // Custom text sequences that will cause the model to stop generating.
  stop_sequences : Array[String]

  // System prompt.
  //
  // A system prompt is a way of providing context and instructions to Claude, such as
  // specifying a particular goal or role. See [guide to system prompts](https://docs.anthropic.com/en/docs/system-prompts).
  mut system : String

  // A number between `0.0` and `1.0` that controls the randomness injected into the response.
  //
  // It is recommended to use `temperature` closer to `0.0`
  // for analytical / multiple choice, and closer to `1.0` for creative tasks.
  //
  // Note that even with `temperature` of `0.0`, the results will not be fully deterministic.
  //
  // The default value is 1.0.
  mut temperature : Double

  // How the model should use the provided tools.
  //
  // Use either `tool_choice_auto`, `tool_choice_any`, or `ToolChoiceTool(name: string)`.
  mut tool_choice : ToolChoice?

  // Definitions of tools that the model may use.
  //
  // Tools can be used for workflows that include running client-side tools and functions,
  // or more generally whenever you want the model to produce a particular JSON structure of output.
  //
  // See Anthropic's [guide](https://docs.anthropic.com/en/docs/tool-use) for more details.
  mut tools : Array[Tool]

  // Only sample from the top K options for each subsequent token.
  //
  // Recommended for advanced use cases only. You usually only need to use `temperature`.
  mut top_k : Int?

  // Use nucleus sampling.
  //
  // You should either alter `temperature` or `top_p`, but not both.
  //
  // Recommended for advanced use cases only. You usually only need to use `temperature`.
  mut top_p : Double?
} derive(Show, Eq, FromJson)

///|
pub impl ToJson for MessagesModelInput with to_json(self) {
  let obj = {
    "model": self.model.to_json(),
    "messages": self.messages.to_json(),
    "max_tokens": self.max_tokens.to_json(),
    "system": self.system.to_json(),
  }
  if self.metadata is Some(metadata) {
    obj["metadata"] = metadata.to_json()
  }
  if self.stop_sequences.length() > 0 {
    obj["stop_sequences"] = self.stop_sequences.to_json()
  }
  if self.temperature != 1 {
    obj["temperature"] = self.temperature.to_json()
  }
  if self.tool_choice is Some(tool_choice) && self.tools.length() > 0 {
    obj["tool_choice"] = tool_choice.to_json()
    obj["tools"] = self.tools.to_json()
  }
  if self.top_k is Some(top_k) {
    obj["top_k"] = top_k.to_json()
  }
  if self.top_p is Some(top_p) {
    obj["top_p"] = top_p.to_json()
  }
  obj.to_json()
}

///|
pub(all) struct Metadata {
  // An external identifier for the user who is associated with the request.
  user_id : String // `json:"user_id,omitempty"`
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) enum ToolChoice {
  /// Auto means that the model will automatically decide whether to use tools.
  Auto
  /// Any means that the model will use any available tools.
  Any
  /// Tool means that the model will use the specified tool.
  /// Tool(name, disable_parallel_tool_use)
  /// `name` is the name of the tool to use.
  /// `disable_parallel_tool_use` determines whether to disable parallel tool use
  /// (defaults to false).
  /// If set to true, the model will output at most one tool use.
  Tool(String, Bool)
} derive(Show, Eq)

///|
pub fn tool_choice_tool(
  name : String,
  disable_parallel_tool_use~ : Bool = false
) -> ToolChoice? {
  Some(Tool(name, disable_parallel_tool_use))
}

///|
pub impl @json.FromJson for ToolChoice with from_json(json, path) {
  match json {
    Object({ "type": String("auto"), .. }) => Auto
    Object({ "type": String("any"), .. }) => Any
    Object(
      {
        "type": String("tool"),
        "name": String(name),
        "disable_parallel_tool_use": True,
        ..
      }
    ) => Tool(name, true)
    Object({ "type": String("tool"), "name": String(name), .. }) =>
      Tool(name, false)
    _ => raise @json.JsonDecodeError((path, "ToolChoice: expected object"))
  }
}

///|
pub impl ToJson for ToolChoice with to_json(self) {
  match self {
    Auto => { "type": "auto" }
    Any => { "type": "any" }
    Tool(name, true) =>
      {
        "type": "tool",
        "name": name.to_json(),
        "disable_parallel_tool_use": true,
      }
    Tool(name, false) => { "type": "tool", "name": name.to_json() }
  }
}

///| The output object for the Anthropic Messages API.
pub(all) struct MessagesModelOutput {
  // Unique object identifier.
  id : String

  // Object type.
  //
  // For Messages, this is always "message".
  type_ : String

  // Conversational role of the generated message.
  //
  // This will always be "assistant".
  role : String

  // Content generated by the model.
  content : Array[ContentBlock]

  // The model that handled the request.
  model : String

  // The reason that the model stopped.
  stop_reason : String

  // Which custom stop sequence was generated, if any.
  stop_sequence : String

  // The usage statistics for the request.
  usage : Usage
} derive(Show, Eq)

///|
pub impl ToJson for MessagesModelOutput with to_json(self) {
  {
    "id": self.id.to_json(),
    "type": self.type_.to_json(), // "message"
    "role": self.role.to_json(),
    "content": self.content.to_json(),
    "model": self.model.to_json(),
    "stop_reason": self.stop_reason.to_json(),
    "stop_sequence": self.stop_sequence.to_json(),
    "usage": self.usage.to_json(),
  }.to_json()
}

///|
pub impl @json.FromJson for MessagesModelOutput with from_json(self, path) {
  let obj = match self {
    Object(obj) => obj
    _ =>
      raise @json.JsonDecodeError(
        (path, "MessagesModelOutput: expected object"),
      )
  }
  let id = match obj.get("id") {
    Some(String(id)) => id
    _ => raise @json.JsonDecodeError((path, "MessagesModelOutput: missing id"))
  }
  let type_ = match obj.get("type") {
    Some(String(type_)) => type_
    _ =>
      raise @json.JsonDecodeError((path, "MessagesModelOutput: missing type"))
  }
  let role = match obj.get("role") {
    Some(String(role)) => role
    _ =>
      raise @json.JsonDecodeError((path, "MessagesModelOutput: missing role"))
  }
  let model = match obj.get("model") {
    Some(String(model)) => model
    _ =>
      raise @json.JsonDecodeError((path, "MessagesModelOutput: missing model"))
  }
  let stop_reason = match obj.get("stop_reason") {
    Some(String(stop_reason)) => stop_reason
    _ =>
      raise @json.JsonDecodeError(
        (path, "MessagesModelOutput: missing stop_reason"),
      )
  }
  let stop_sequence = match obj.get("stop_sequence") {
    Some(String(stop_sequence)) => stop_sequence
    _ => ""
  }
  let content : Array[ContentBlock] = @json.from_json(
    obj.get("content").unwrap(),
  )
  let usage : Usage = @json.from_json(obj.get("usage").unwrap())
  { id, type_, role, content, model, stop_reason, stop_sequence, usage }
}

///|
pub(all) struct Usage {
  // The number of input tokens which were used.
  input_tokens : Int

  // The number of output tokens which were used.
  output_tokens : Int
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct Message {
  // The role of the author of this message.
  role : String

  // The content of the message.
  content : Content
} derive(Show, Eq, FromJson, ToJson)

///| Creates a new user message object.
pub fn Message::new_user(content : Content) -> Message {
  { role: "user", content }
}

///| Content may be either a single string or an array of content blocks,
// where each block has a specific type.
pub(all) enum Content {
  // A string of text.
  String(String)
  // An array of content blocks.
  Contents(Contents)
} derive(Show, Eq, FromJson)

///|
pub impl ToJson for Content with to_json(self) {
  match self {
    String(s) => s.to_json()
    Contents(contents) => contents.to_json()
  }
}

///|
pub(all) type Contents Array[ContentBlock] derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct ContentBlock {
  type_ : String
  text : String?
  id : String?
  name : String?
  input : Json?
  tool_use_id : String?
  is_error : Bool?
  content : String?
} derive(Show, Eq)

///|
pub impl @json.FromJson for ContentBlock with from_json(self, path) {
  let obj = match self {
    Object(obj) => obj
    _ => raise @json.JsonDecodeError((path, "ContentBlock: expected object"))
  }
  let type_ = match obj.get("type") {
    Some(String(type_)) => type_
    _ => raise @json.JsonDecodeError((path, "ContentBlock: missing type"))
  }
  let text = match obj.get("text") {
    Some(String(text)) => Some(text)
    _ => None
  }
  let id = match obj.get("id") {
    Some(String(id)) => Some(id)
    _ => None
  }
  let name = match obj.get("name") {
    Some(String(name)) => Some(name)
    _ => None
  }
  let input = match obj.get("input") {
    Some(input) => Some(input)
    _ => None
  }
  let tool_use_id = match obj.get("tool_use_id") {
    Some(String(tool_use_id)) => Some(tool_use_id)
    _ => None
  }
  let is_error = match obj.get("is_error") {
    Some(is_error) => Some(is_error.as_bool().unwrap())
    _ => None
  }
  let content = match obj.get("content") {
    Some(String(content)) => Some(content)
    _ => None
  }
  { type_, text, id, name, input, tool_use_id, is_error, content }
}

///|
pub impl ToJson for ContentBlock with to_json(self) {
  {
    "type": self.type_.to_json(),
    "text": match self.text {
      Some(text) => text.to_json()
      None => Json::null()
    },
    "id": match self.id {
      Some(id) => id.to_json()
      None => Json::null()
    },
    "name": match self.name {
      Some(name) => name.to_json()
      None => Json::null()
    },
    "input": match self.input {
      Some(input) => input
      None => Json::null()
    },
    "tool_use_id": match self.tool_use_id {
      Some(tool_use_id) => tool_use_id.to_json()
      None => Json::null()
    },
    "is_error": match self.is_error {
      Some(is_error) => is_error.to_json()
      None => Json::null()
    },
    "content": match self.content {
      Some(content) => content.to_json()
      None => Json::null()
    },
  }.to_json()
}

///|
pub fn ContentBlock::new(type_ : String) -> ContentBlock {
  {
    type_,
    text: None,
    id: None,
    name: None,
    input: None,
    tool_use_id: None,
    is_error: None,
    content: None,
  }
}

///|
pub fn text_content_block(text : String) -> ContentBlock {
  let content_block = ContentBlock::new("text")
  { ..content_block, text: Some(text) }
}

///|
pub fn tool_use_content_block(
  id : String,
  name : String,
  input : Json
) -> ContentBlock {
  let content_block = ContentBlock::new("tool_use")
  { ..content_block, id: Some(id), name: Some(name), input: Some(input) }
}

///|
pub fn tool_result_content_block(
  tool_use_id : String,
  content : String,
  is_error : Bool
) -> ContentBlock {
  let content_block = ContentBlock::new("tool_result")
  {
    ..content_block,
    tool_use_id: Some(tool_use_id),
    content: Some(content),
    is_error: Some(is_error),
  }
}

// A tool object that the model may call.

///|
pub(all) struct Tool {
  // Name of the tool.
  name : String

  // [JSON schema](https://json-schema.org/) for this tool's input.
  //
  // This defines the shape of the `input` that your tool accepts and that the model will produce.
  input_schema : Json

  // Optional, but strongly-recommended description of the tool.
  description : String?
} derive(Show, Eq, FromJson, ToJson)

// Creates an input object for the Anthropic Messages API.

///|
pub fn create_input(
  self : MessagesModel,
  messages : Array[Message]
) -> MessagesModelInput {
  let info = self.inner().info()
  let model = info.full_name
  {
    model,
    messages,
    max_tokens: 0,
    metadata: None,
    stop_sequences: [],
    system: "",
    temperature: 1,
    tool_choice: Some(ToolChoice::Auto),
    tools: [],
    top_k: None,
    top_p: None,
  }
}
