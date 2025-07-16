// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
fn host_add(a : Int, b : Int) -> Int = "modus_test" "add"

// TODO: Handle import function type signature overrides:
///|modus:import modus_test echo1(String) -> String
fn _host_echo1(message : Int) -> Int = "modus_test" "echo1"

///|modus:import modus_test echo2(String?) -> String
fn _host_echo2(message : Int) -> Int = "modus_test" "echo2"

///|modus:import modus_test echo3(String) -> String?
fn _host_echo3(message : Int) -> Int = "modus_test" "echo3"

///|modus:import modus_test echo4(String?) -> String?
fn _host_echo4(message : Int) -> Int = "modus_test" "echo4"

///|
pub fn host_echo1(message : String) -> String {
  let ptr = @ffi.str2ptr(message) - 8
  let result = _host_echo1(ptr)
  @ffi.ptr2str(result + 8)
}

///|
pub fn host_echo2(message : String?) -> String {
  match message {
    None => {
      let result = _host_echo2(0)
      @ffi.ptr2str(result + 8)
    }
    Some(m) => {
      let ptr = @ffi.str2ptr(m) - 8
      let result = _host_echo2(ptr)
      @ffi.ptr2str(result + 8)
    }
  }
}

///|
pub fn host_echo3(message : String) -> String? {
  let ptr = @ffi.str2ptr(message) - 8
  let result = _host_echo3(ptr)
  match result {
    0 => None
    _ => Some(@ffi.ptr2str(result + 8))
  }
}

///|
pub fn host_echo4(message : String?) -> String? {
  let result = match message {
    None => _host_echo4(0)
    Some(m) => {
      let ptr = @ffi.str2ptr(m) - 8
      _host_echo4(ptr)
    }
  }
  match result {
    0 => None
    _ => Some(@ffi.ptr2str(result + 8))
  }
}

///|modus:import modus_test encodeStrings1(Array[String]?) -> String?
fn _host_encode_strings1(items : Int) -> Int = "modus_test" "encodeStrings1"

///|
fn host_encode_strings1(items : Array[String]?) -> String? {
  let result = match items {
    None => _host_encode_strings1(0)
    Some(m) => {
      let ptr : Int = cast(items)
      _host_encode_strings1(ptr)
    }
  }
  match result {
    0 => None
    _ => Some(@ffi.ptr2str(result + 8))
  }
}

///|modus:import modus_test encodeStrings2(Array[String?]?) -> String?
fn _host_encode_strings2(items : Int) -> Int = "modus_test" "encodeStrings2"

///|
fn host_encode_strings2(items : Array[String?]?) -> String? {
  let result = match items {
    None => _host_encode_strings2(0)
    Some(m) => {
      let ptr : Int = cast(items)
      _host_encode_strings2(ptr)
    }
  }
  match result {
    0 => None
    _ => Some(@ffi.ptr2str(result + 8))
  }
}
