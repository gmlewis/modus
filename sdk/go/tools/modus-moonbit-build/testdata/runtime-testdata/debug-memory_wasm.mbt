// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
fn cast_to_ptr[A](a : A) -> Int = "%identity"

///|
fn dump_mem_block(ptr : Int) -> String {
  let s = StringBuilder::new()
  s.write_string("ptr: \{ptr}: ")
  let refCount = load32(ptr)
  let part2 = load32(ptr + 4)
  let class_id = part2 & 0xff
  let words = part2 >> 8
  s.write_string(
    "refCount: \{refCount}, class_id: \{class_id}, words: \{words}: [",
  )
  for i in 0..<words {
    s.write_string("\{load32(ptr + 8 + i * 4)} ")
  }
  s.write_char(']')
  s.to_string()
}
