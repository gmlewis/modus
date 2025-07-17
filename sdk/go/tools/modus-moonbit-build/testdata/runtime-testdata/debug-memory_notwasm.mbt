// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
fn[A] cast_to_ptr(a : A) -> Int = "%identity"

///|
fn dump_mem_block(ptr : Int) -> String {
  let s = StringBuilder::new()
  s.write_string("ptr: \{ptr}: ")
  s.to_string()
}
