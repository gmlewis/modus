// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| The openai package provides objects that conform to the OpenAI API specification.
/// It can be used with OpenAI, as well as other services that conform to the OpenAI API.

///| The usage statistics returned by the OpenAI API.
pub(all) struct Usage {
  // The number of completion tokens used.
  completion_tokens : Int

  // The number of prompt tokens used.
  prompt_tokens : Int

  // The total number of tokens used.
  total_tokens : Int
} derive(Show, Eq, FromJson, ToJson)
