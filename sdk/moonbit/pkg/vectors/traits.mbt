// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub(open) trait Addable {
  // This no longer works as of "moonc v0.1.20250310+a7a1e9804".
  // op_add(Self, Self) -> Self
  add(Self, Self) -> Self
}

///|
pub impl Addable for Byte with add(a, b) {
  a + b
}

///|
pub impl Addable for Int with add(a, b) {
  a + b
}

///|
pub impl Addable for UInt with add(a, b) {
  a + b
}

///|
pub impl Addable for Int16 with add(a, b) {
  a + b
}

///|
pub impl Addable for UInt16 with add(a, b) {
  a + b
}

///|
pub impl Addable for Int64 with add(a, b) {
  a + b
}

///|
pub impl Addable for UInt64 with add(a, b) {
  a + b
}

///|
pub impl Addable for Double with add(a, b) {
  a + b
}

///|
pub impl Addable for Float with add(a, b) {
  a + b
}

///|
pub impl Addable for BigInt with add(a, b) {
  a + b
}

///|
pub(open) trait Subtractable {
  // This no longer works as of "moonc v0.1.20250310+a7a1e9804".
  // op_sub(Self, Self) -> Self
  sub(Self, Self) -> Self
}

///|
pub impl Subtractable for Byte with sub(a, b) {
  a - b
}

///|
pub impl Subtractable for Int with sub(a, b) {
  a - b
}

///|
pub impl Subtractable for UInt with sub(a, b) {
  a - b
}

///|
pub impl Subtractable for Int16 with sub(a, b) {
  a - b
}

///|
pub impl Subtractable for UInt16 with sub(a, b) {
  a - b
}

///|
pub impl Subtractable for Int64 with sub(a, b) {
  a - b
}

///|
pub impl Subtractable for UInt64 with sub(a, b) {
  a - b
}

///|
pub impl Subtractable for Double with sub(a, b) {
  a - b
}

///|
pub impl Subtractable for Float with sub(a, b) {
  a - b
}

///|
pub impl Subtractable for BigInt with sub(a, b) {
  a - b
}

///|
pub(open) trait Multipliable: Mul {
  op_mul(Self, Self) -> Self
}

///|
pub impl Multipliable for Byte with op_mul(a, b) {
  a * b
}

///|
pub impl Multipliable for Int with op_mul(a, b) {
  a * b
}

///|
pub impl Multipliable for UInt with op_mul(a, b) {
  a * b
}

///|
pub impl Multipliable for Int16 with op_mul(a, b) {
  a * b
}

///|
pub impl Multipliable for UInt16 with op_mul(a, b) {
  a * b
}

///|
pub impl Multipliable for Int64 with op_mul(a, b) {
  a * b
}

///|
pub impl Multipliable for UInt64 with op_mul(a, b) {
  a * b
}

///|
pub impl Multipliable for Double with op_mul(a, b) {
  a * b
}

///|
pub impl Multipliable for Float with op_mul(a, b) {
  a * b
}

///|
pub impl Multipliable for BigInt with op_mul(a, b) {
  a * b
}

///|
pub(open) trait Divideable: Div {
  op_div(Self, Self) -> Self
}

///|
pub impl Divideable for Byte with op_div(a, b) {
  a / b
}

///|
pub impl Divideable for Int with op_div(a, b) {
  a / b
}

///|
pub impl Divideable for UInt with op_div(a, b) {
  a / b
}

///|
pub impl Divideable for Int16 with op_div(a, b) {
  a / b
}

///|
pub impl Divideable for UInt16 with op_div(a, b) {
  a / b
}

///|
pub impl Divideable for Int64 with op_div(a, b) {
  a / b
}

///|
pub impl Divideable for UInt64 with op_div(a, b) {
  a / b
}

///|
pub impl Divideable for Double with op_div(a, b) {
  a / b
}

///|
pub impl Divideable for Float with op_div(a, b) {
  a / b
}

///|
pub impl Divideable for BigInt with op_div(a, b) {
  a / b
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
pub impl SquareRootable for Double with sqrt(self) {
  self.sqrt()
}

///|
pub impl SquareRootable for Float with sqrt(self) {
  self.sqrt()
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
pub impl FromInt for Double with from_int(self) {
  self.to_double()
}

///|
pub impl FromInt for Float with from_int(self) {
  self.to_float()
}

///|
pub(open) trait Absoluteable {
  abs(Self) -> Self
}

///|
pub impl Absoluteable for Int with abs(self) {
  self.to_double().abs().round().to_int()
}

///|
pub impl Absoluteable for Int64 with abs(self) {
  self.to_double().abs().round().to_int64()
}

///|
pub impl Absoluteable for Double with abs(self) {
  self.abs()
}

///|
pub impl Absoluteable for Float with abs(self) {
  self.abs()
}
