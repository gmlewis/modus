// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn add(a : Int, b : Int) -> Int {
  host_add(a, b)
}

///|
pub fn echo1(message : String) -> String {
  host_echo1(message)
}

///|
pub fn echo2(message : String?) -> String {
  host_echo2(message)
}

///|
pub fn echo3(message : String) -> String? {
  host_echo3(message)
}

///|
pub fn echo4(message : String?) -> String? {
  host_echo4(message)
}

///|
pub fn encode_strings1(items : Array[String]?) -> String? {
  host_encode_strings1(items)
}

///|
pub fn encode_strings2(items : Array[String?]?) -> String? {
  host_encode_strings2(items)
}
