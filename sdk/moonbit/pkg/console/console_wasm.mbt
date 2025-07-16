// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|modus:import modus_system logMessage(level : String, message : String) -> Unit
fn wasmImportLogMessage(level : Int, message : Int) = "modus_system" "logMessage"

///| Log a message with specified level
pub fn log_message(level : String, message : String) -> Unit {
  let level_ptr = @ffi.str2ptr(level) - 8
  let message_ptr = @ffi.str2ptr(message) - 8
  wasmImportLogMessage(level_ptr, message_ptr)
}

///|
pub fn assert_(condition : Bool, message : String) -> Unit {
  if not(condition) {
    error("Assertion failed: \{message}")
  }
}

///|
pub fn log(message : String) -> Unit {
  log_message("", message)
}

///|
pub fn debug(message : String) -> Unit {
  log_message("debug", message)
}

///|
pub fn info(message : String) -> Unit {
  log_message("info", message)
}

///|
pub fn warn(message : String) -> Unit {
  log_message("warning", message)
}

///|
pub fn error(message : String) -> Unit {
  log_message("error", message)
}
