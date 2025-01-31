// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn assert(condition : Bool, message : String) -> Unit {
  if not(condition) {
    error("Assertion failed: \{message}")
  }
}

///|
pub fn log(message : String) -> Unit {
  @host.log_message("", message)
}

///|
pub fn debug(message : String) -> Unit {
  @host.log_message("debug", message)
}

///|
pub fn info(message : String) -> Unit {
  @host.log_message("info", message)
}

///|
pub fn warn(message : String) -> Unit {
  @host.log_message("warning", message)
}

///|
pub fn error(message : String) -> Unit {
  @host.log_message("error", message)
}
