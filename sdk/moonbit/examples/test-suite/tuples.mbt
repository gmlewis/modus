// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub fn test_tuple_output() -> (Int, Bool, String) {
  (123, true, "hello")
}

///|
pub struct TupleSimulator {
  t0 : Int
  t1 : Bool
  t2 : String
}

///|
pub fn test_tuple_simulator() -> TupleSimulator {
  { t0: 123, t1: true, t2: "hello" }
}
