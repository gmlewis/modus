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

///| add adds two vectors together and returns the result.
pub fn add[T : Addable](a : Array[T], b : Array[T]) -> Array[T] {
  let len = a.length()
  guard len == b.length()
  let result = Array::new(capacity=len)
  for i in 0..<len {
    result.push(a[i] + b[i])
  }
  result
}

///| add_in_place adds two vectors together and stores the result in the first vector.
pub fn add_in_place[T : Addable](a : Array[T], b : Array[T]) -> Unit {
  let len = a.length()
  guard len == b.length()
  for i in 0..<len {
    a[i] += b[i]
  }
}

///|
pub(open) trait Subtractable {
  op_sub(Self, Self) -> Self
}

///| subtract subtracts one vector from another and returns the result.
pub fn subtract[T : Subtractable](a : Array[T], b : Array[T]) -> Array[T] {
  let len = a.length()
  guard len == b.length()
  let result = Array::new(capacity=len)
  for i in 0..<len {
    result.push(a[i] - b[i])
  }
  result
}

///| subtract_in_place subtracts one vector from another and stores the result in the first vector.
pub fn subtract_in_place[T : Subtractable](a : Array[T], b : Array[T]) -> Unit {
  let len = a.length()
  guard len == b.length()
  for i in 0..<len {
    a[i] -= b[i]
  }
}

///| add_number adds a number to each element of a vector and returns the result.
pub fn add_number[T : Addable](a : Array[T], b : T) -> Array[T] {
  let len = a.length()
  let result = Array::new(capacity=len)
  for i in 0..<len {
    result.push(a[i] + b)
  }
  result
}

///| add_number_in_place adds a number to each element of a vector and stores the result in the vector.
pub fn add_number_in_place[T : Addable](a : Array[T], b : T) -> Unit {
  let len = a.length()
  for i in 0..<len {
    a[i] += b
  }
}

///| subtract_number subtracts a number from each element of a vector and returns the result.
pub fn subtract_number[T : Subtractable](a : Array[T], b : T) -> Array[T] {
  let len = a.length()
  let result = Array::new(capacity=len)
  for i in 0..<len {
    result.push(a[i] - b)
  }
  result
}

///| subtract_number_in_place subtracts a number from each element of a vector and stores the result in the vector.
pub fn subtract_number_in_place[T : Subtractable](a : Array[T], b : T) -> Unit {
  let len = a.length()
  for i in 0..<len {
    a[i] -= b
  }
}

///|
pub(open) trait Multipliable {
  op_mul(Self, Self) -> Self
}

///| multiply_number Multiplys a number to each element of a vector and returns the result.
pub fn multiply_number[T : Multipliable](a : Array[T], b : T) -> Array[T] {
  let len = a.length()
  let result = Array::new(capacity=len)
  for i in 0..<len {
    result.push(a[i] * b)
  }
  result
}

///| multiply_number_in_place multiplies each element of a vector by a number and stores the result in the vector.
pub fn multiply_number_in_place[T : Multipliable](a : Array[T], b : T) -> Unit {
  let len = a.length()
  for i in 0..<len {
    a[i] *= b
  }
}

///|
pub(open) trait Divideable {
  op_div(Self, Self) -> Self
}

///| divide_number divides each element of a vector by a number and returns the result.
pub fn divide_number[T : Divideable](a : Array[T], b : T) -> Array[T] {
  let len = a.length()
  let result = Array::new(capacity=len)
  for i in 0..<len {
    result.push(a[i] / b)
  }
  result
}

///| divide_number_in_place divides each element of a vector by a number and stores the result in the vector.
pub fn divide_number_in_place[T : Divideable](a : Array[T], b : T) -> Unit {
  let len = a.length()
  for i in 0..<len {
    a[i] /= b
  }
}

///| dot computes the dot product of two vectors.
pub fn dot[T : Multipliable + Addable](a : Array[T], b : Array[T]) -> T {
  let len = a.length()
  guard len > 0
  guard len == b.length()
  let mut result = a[0] * b[0]
  for i in 1..<len {
    result += a[i] * b[i]
  }
  result
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

///| magnitude computes the magnitude of a vector.
pub fn magnitude[T : Multipliable + Addable + SquareRootable](
  a : Array[T]
) -> T {
  dot(a, a).sqrt()
}

///| normalize normalizes a vector to have a magnitude of 1.
pub fn normalize[T : Multipliable + Divideable + Addable + SquareRootable](
  a : Array[T]
) -> Array[T] {
  let mag = magnitude(a)
  divide_number(a, mag)
}

///| sum computes the sum of all elements in a vector.
pub fn sum[T : Addable](a : Array[T]) -> T {
  let len = a.length()
  guard len > 0
  let mut result = a[0] // 0
  for i in 1..<len {
    result += a[i]
  }
  result
}

///| product computes the product of all elements in a vector.
pub fn product[T : Multipliable](a : Array[T]) -> T {
  let len = a.length()
  guard len > 0
  let mut result = a[0] // 1
  for i in 1..<len {
    result *= a[i]
  }
  result
}

///| mean computes the mean of a vector.
pub fn mean[T : Addable + Divideable](a : Array[T]) -> T {
  let len = a.length()
  sum(a) / len.to_double()
}

///| min computes the minimum element in a vector.
pub fn min[T : Compare](a : Array[T]) -> T {
  let len = a.length()
  guard len > 0
  let mut result = a[0]
  for i in 1..<len {
    if a[i] < result {
      result = a[i]
    }
  }
  result
}

///| max computes the maximum element in a vector.
pub fn max[T : Compare](a : Array[T]) -> T {
  let len = a.length()
  guard len > 0
  let mut result = a[0]
  for i in 1..<len {
    if a[i] > result {
      result = a[i]
    }
  }
  result
}

///|
pub(open) trait Absoluteable {
  abs(Self) -> Self
}

///| abs computes the absolute value of each element in a vector.
pub fn abs[T : Absoluteable](a : Array[T]) -> Array[T] {
  let len = a.length()
  let result = Array::new(capacity=len)
  for i in 0..<len {
    result.push(a[i].abs())
  }
  result
}

///| abs_in_place computes the absolute value of each element in a vector and
/// stores the result in the vector.
pub fn abs_in_place[T : Absoluteable](a : Array[T]) -> Unit {
  let len = a.length()
  for i in 0..<len {
    a[i] = a[i].abs()
  }
}

///| euclidian_distance computes the Euclidian distance between two vectors.
pub fn euclidian_distance[T : Multipliable + SquareRootable + Addable + Subtractable](
  a : Array[T],
  b : Array[T]
) -> T {
  let len = a.length()
  guard len == b.length()
  let mut result = (a[0] - b[0]) * (a[0] - b[0])
  for i in 1..<len {
    result += (a[i] - b[i]) * (a[i] - b[i])
  }
  result.sqrt()
}
