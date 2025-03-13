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
  max_tokens : Int?

  // A `Metadata` object describing the request.
  metadata : Metadata?

  // Custom text sequences that will cause the model to stop generating.
  stop_sequences : Array[String]

  // System prompt.
  //
  // A system prompt is a way of providing context and instructions to Claude, such as
  // specifying a particular goal or role. See [guide to system prompts](https://docs.anthropic.com/en/docs/system-prompts).
  system : String

  // A number between `0.0` and `1.0` that controls the randomness injected into the response.
  //
  // It is recommended to use `temperature` closer to `0.0`
  // for analytical / multiple choice, and closer to `1.0` for creative tasks.
  //
  // Note that even with `temperature` of `0.0`, the results will not be fully deterministic.
  //
  // The default value is 1.0.
  temperature : Double

  // How the model should use the provided tools.
  //
  // Use either `tool_choice_auto`, `tool_choice_any`, or `ToolChoiceTool(name: string)`.
  tool_choice : ToolChoice?

  // Definitions of tools that the model may use.
  //
  // Tools can be used for workflows that include running client-side tools and functions,
  // or more generally whenever you want the model to produce a particular JSON structure of output.
  //
  // See Anthropic's [guide](https://docs.anthropic.com/en/docs/tool-use) for more details.
  tools : Array[Tool]

  // Only sample from the top K options for each subsequent token.
  //
  // Recommended for advanced use cases only. You usually only need to use `temperature`.
  top_k : Int?

  // Use nucleus sampling.
  //
  // You should either alter `temperature` or `top_p`, but not both.
  //
  // Recommended for advanced use cases only. You usually only need to use `temperature`.
  top_p : Double?
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) struct Metadata {
  // An external identifier for the user who is associated with the request.
  user_id : String // `json:"user_id,omitempty"`
} derive(Show, Eq, FromJson, ToJson)

///|
pub(all) enum ToolChoice {
  ///| Auto means that the model will automatically decide whether to use tools.
  Auto
  ///| Any means that the model will use any available tools.
  Any
  ///| Tool means that the model will use the specified tool.
  /// Tool(name, disable_parallel_tool_use)
  /// `name` is the name of the tool to use.
  /// `disable_parallel_tool_use` determines whether to disable parallel tool use
  /// (defaults to false).
  /// If set to true, the model will output at most one tool use.
  Tool(String, Bool)
} derive(Show, Eq)

///|
impl @json.FromJson for ToolChoice with from_json(json, path) {
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
impl ToJson for ToolChoice with to_json(self) {
  match self {
    Auto => Object({ "type": "auto" })
    Any => Object({ "type": "any" })
    Tool(name, true) =>
      Object({
        "type": "tool",
        "name": name.to_json(),
        "disable_parallel_tool_use": True,
      })
    Tool(name, false) => Object({ "type": "tool", "name": name.to_json() })
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
  stop_sequence : String?

  // The usage statistics for the request.
  usage : Usage
} derive(Show, Eq, FromJson, ToJson)

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

///|
pub(all) enum Content {
  // A string of text.
  String(StringContent)
  // An array of content blocks.
  Contents(Contents)
} derive(Show, Eq, FromJson, ToJson)

// Content may be either a single string or an array of content blocks,
// where each block has a specific type.
///|
pub(all) type StringContent String derive(Show, Eq, FromJson, ToJson)

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
} derive(Show, Eq, FromJson, ToJson)

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
  let info = self._.info()
  let model = info.full_name
  {
    model,
    messages,
    max_tokens: None,
    metadata: None,
    stop_sequences: [],
    system: "",
    temperature: 1.0,
    tool_choice: Some(ToolChoice::Auto),
    tools: [],
    top_k: None,
    top_p: None,
  }
}
