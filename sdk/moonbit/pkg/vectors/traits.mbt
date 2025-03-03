// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(open) trait Addable {
  op_add(Self, Self) -> Self
}

///|
pub(open) trait Subtractable {
  op_sub(Self, Self) -> Self
}

///|
pub(open) trait Multipliable {
  op_mul(Self, Self) -> Self
}

///|
pub(open) trait Divideable {
  op_div(Self, Self) -> Self
}

///|
pub(open) trait SquareRootable {
  sqrt(Self) -> Self
}

///|
pub impl SquareRootable for Int with sqrt(self) {
  self.to_double().sqrt().round().to_int()
}

///|
pub impl SquareRootable for Int64 with sqrt(self) {
  self.to_double().sqrt().round().to_int64()
}

///|
pub(open) trait FromInt {
  from_int(Int) -> Self
}

///|
pub impl FromInt for Int with from_int(self) {
  self
}

///|
pub impl FromInt for Float with from_int(self) {
  self.to_float()
}

///|
pub(open) trait Absoluteable {
  abs(Self) -> Self
}
