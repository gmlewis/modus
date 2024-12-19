// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

test "add" {
  // Int
  let a = [1, 2, 3]
  let b = [4, 5, 6]
  let got = add(a, b)
  inspect!(a, content="[1, 2, 3]")
  inspect!(b, content="[4, 5, 6]")
  inspect!(got, content="[5, 7, 9]")

  // UInt
  let a = [1U, 2, 3]
  let b = [4U, 5, 6]
  let got = add(a, b)
  inspect!(a, content="[1, 2, 3]")
  inspect!(b, content="[4, 5, 6]")
  inspect!(got, content="[5, 7, 9]")

  // Int64
  let a = [1L, 2, 3]
  let b = [4L, 5, 6]
  let got = add(a, b)
  inspect!(got, content="[5, 7, 9]")

  // UInt64
  let a = [1UL, 2, 3]
  let b = [4UL, 5, 6]
  let got = add(a, b)
  inspect!(a, content="[1, 2, 3]")
  inspect!(b, content="[4, 5, 6]")
  inspect!(got, content="[5, 7, 9]")

  // Double
  let a = [1.0, 2, 3]
  let b = [4.0, 5, 6]
  let got = add(a, b)
  inspect!(got, content="[5, 7, 9]")

  // Byte
  let a = [b'\x01', 2, 3]
  let b = [b'\x04', 5, 6]
  let got = add(a, b)
  inspect!(got, content="[b'\\x05', b'\\x07', b'\\x09']")
}

test "add_in_place" {
  let a = [1, 2, 3]
  let b = [4, 5, 6]
  add_in_place(a, b)
  inspect!(a, content="[5, 7, 9]")
  inspect!(b, content="[4, 5, 6]")
}

test "subtract" {
  // Int
  let a = [4, 5, 6]
  let b = [1, 2, 3]
  let got = subtract(a, b)
  inspect!(a, content="[4, 5, 6]")
  inspect!(b, content="[1, 2, 3]")
  inspect!(got, content="[3, 3, 3]")
}

test "subtract_in_place" {
  let a = [4, 5, 6]
  let b = [1, 2, 3]
  subtract_in_place(a, b)
  inspect!(a, content="[3, 3, 3]")
  inspect!(b, content="[1, 2, 3]")
}

test "add_number" {
  let a = [1, 2, 3]
  let b = 4
  let got = add_number(a, b)
  inspect!(a, content="[1, 2, 3]")
  inspect!(got, content="[5, 6, 7]")
}

test "add_number_in_place" {
  let a = [1, 2, 3]
  let b = 4
  add_number_in_place(a, b)
  inspect!(a, content="[5, 6, 7]")
}

test "subtract_number" {
  let a = [4, 5, 6]
  let b = 1
  let got = subtract_number(a, b)
  inspect!(a, content="[4, 5, 6]")
  inspect!(got, content="[3, 4, 5]")
}

test "subtract_number_in_place" {
  let a = [4, 5, 6]
  let b = 1
  subtract_number_in_place(a, b)
  inspect!(a, content="[3, 4, 5]")
}

test "multiply_number" {
  let a = [1, 2, 3]
  let b = 4
  let got = multiply_number(a, b)
  inspect!(a, content="[1, 2, 3]")
  inspect!(got, content="[4, 8, 12]")
}

test "multiply_number_in_place" {
  let a = [1, 2, 3]
  let b = 4
  multiply_number_in_place(a, b)
  inspect!(a, content="[4, 8, 12]")
}

test "divide_number" {
  let a = [4, 8, 12]
  let b = 4
  let got = divide_number(a, b)
  inspect!(a, content="[4, 8, 12]")
  inspect!(got, content="[1, 2, 3]")
}

test "divide_number_in_place" {
  let a = [4, 8, 12]
  let b = 4
  divide_number_in_place(a, b)
  inspect!(a, content="[1, 2, 3]")
}

test "dot" {
  let a = [1.0, 2, 3]
  let b = [4.0, 5, 6]
  let got = dot(a, b)
  inspect!(a, content="[1, 2, 3]")
  inspect!(b, content="[4, 5, 6]")
  inspect!(got, content="32")
}

test "magnitude" {
  let a = [3.0, 4]
  let got = magnitude(a)
  inspect!(a, content="[3, 4]")
  inspect!(got, content="5")
}

test "normalize" {
  let a = [3.0, 4]
  let got = normalize(a)
  inspect!(a, content="[3, 4]")
  inspect!(got, content="[0.6, 0.8]")
}

test "sum" {
  let a = [1, 2, 3]
  let got = sum(a)
  inspect!(a, content="[1, 2, 3]")
  inspect!(got, content="6")
}

test "product" {
  let a = [1, 2, 3]
  let got = product(a)
  inspect!(a, content="[1, 2, 3]")
  inspect!(got, content="6")
}

test "mean" {
  let a = [1, 2, 3]
  let got = mean(a)
  inspect!(a, content="[1, 2, 3]")
  inspect!(got, content="2")
}

test "min" {
  let a = [1, 2, 3]
  let got = min(a)
  inspect!(got, content="1")
}

test "max" {
  let a = [1, 2, 3]
  let got = max(a)
  inspect!(got, content="3")
}

test "abs" {
  let a = [-1, 2, -3]
  let got = abs(a)
  inspect!(got, content="[1, 2, 3]")
}

test "abs_in_place" {
  let a = [-1, 2, -3]
  abs_in_place(a)
  inspect!(a, content="[1, 2, 3]")
}

test "euclidian_distance" {
  let a = [1.0, 2, 3]
  let b = [4.0, 5, 6]
  let got = euclidian_distance(a, b)
  inspect!(a, content="[1, 2, 3]")
  inspect!(b, content="[4, 5, 6]")
  inspect!(got, content="5.196152422706632")
}
