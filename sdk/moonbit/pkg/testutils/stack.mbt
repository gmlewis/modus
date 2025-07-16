// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(all) struct CallStack[T] {
  items : Array[Array[T]]
} derive(Show, Eq, FromJson, ToJson)

///|
pub fn CallStack::new[T]() -> CallStack[T] {
  { items: [] }
}

///|
pub fn CallStack::size[T](self : CallStack[T]) -> Int {
  self.items.length()
}

///|
pub fn push[T](self : CallStack[T], values : Array[T]) -> Unit {
  self.items.push(values)
}

///|
pub fn pop[T](self : CallStack[T]) -> Array[T] {
  match self.items.pop() {
    Some(values) => values
    None => []
  }
}
