// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
fn host_add(a : Int, b : Int) -> Int = "modus_test" "add"

///|
fn _host_echo1(message : Int) -> Int = "modus_test" "echo1"

///|
fn _host_echo2(message : Int) -> Int = "modus_test" "echo2"

///|
fn _host_echo3(message : Int) -> Int = "modus_test" "echo3"

///|
fn _host_echo4(message : Int) -> Int = "modus_test" "echo4"

///|
fn run_host_echo(message : String?, f : (Int) -> Int) -> String? {
  let result = match message {
    None => f(0)
    Some(m) => f(@ffi.str2ptr(m))
  }
  match result {
    0 => None
    _ => Some(@ffi.ptr2str(result))
  }
}

///|
fn host_echo1(message : String?) -> String? {
  run_host_echo(message, _host_echo1)
}

///|
fn host_echo2(message : String?) -> String? {
  run_host_echo(message, _host_echo2)
}

///|
fn host_echo3(message : String?) -> String? {
  run_host_echo(message, _host_echo3)
}

///|
fn host_echo4(message : String?) -> String? {
  run_host_echo(message, _host_echo4)
}

///|
fn _hostEncodeStrings1(items : Int) -> Int = "modus_test" "encodeStrings1"

///|
extern "wasm" fn _string_array2ptr(items : Int) -> Int =
  #|(func (param i32) (result i32) local.get 0 i32.const 8 i32.sub)

///|
fn host_encode_strings1(items : Array[String]?) -> String? {
  let result = match items {
    None => _hostEncodeStrings1(0)
    Some(m) => {
      let ptr = @ffi.malloc(4 * m.length())
      for i in 0..<m.length() {
        @ffi.store32(ptr + 4 * i, @ffi.str2ptr(m[i]))
      }
      let result = _hostEncodeStrings1(_string_array2ptr(ptr))
      @ffi.free(ptr)
      result
    }
  }
  match result {
    0 => None
    _ => Some(@ffi.ptr2str(result))
  }
}

///|
fn _hostEncodeStrings2(items : Int) -> Int = "modus_test" "encodeStrings2"

///|
fn host_encode_strings2(items : Array[String?]?) -> String? {
  let result = match items {
    None => _hostEncodeStrings2(0)
    Some(m) => {
      let ptr = @ffi.malloc(4 * m.length())
      for i in 0..<m.length() {
        match m[i] {
          None => @ffi.store32(ptr + 4 * i, 0)
          Some(s) => @ffi.store32(ptr + 4 * i, @ffi.str2ptr(s))
        }
      }
      let result = _hostEncodeStrings2(_string_array2ptr(ptr))
      @ffi.free(ptr)
      result
    }
  }
  match result {
    0 => None
    _ => Some(@ffi.ptr2str(result))
  }
}
