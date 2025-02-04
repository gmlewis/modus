// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn host_add(a : Int, b : Int) -> Int = "modus_test" "add"

///|
pub fn host_echo1(message : String?) -> String? = "modus_test" "echo1"

///|
pub fn host_echo2(message : String?) -> String? = "modus_test" "echo2"

///|
pub fn host_echo3(message : String?) -> String? = "modus_test" "echo3"

///|
pub fn host_echo4(message : String?) -> String? = "modus_test" "echo4"

///|
pub fn _hostEncodeStrings1(items : Any) -> String? = "modus_test" "encodeStrings1"

///|
//modus:import modus_test encodeStrings1
pub fn host_encode_strings1(items : Array[String]?) -> String {
  _hostEncodeStrings1(items)
}

///|
pub fn _hostEncodeStrings2(items : Any) -> String? = "modus_test" "encodeStrings2"

///|
//modus:import modus_test encodeStrings2
pub fn host_encode_strings2(items : Array[String?]?) -> String {
  _hostEncodeStrings2(items)
}
