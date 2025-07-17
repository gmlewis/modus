// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///| add adds two vectors together and returns the result.
/// This no longer works as of "moonc v0.1.20250310+a7a1e9804".
// pub impl[T : Addable] Addable for Array[T] with op_add(a, b) {
//   let len = a.length()
//   guard len == b.length()
//   Array::makei(len, fn(i) { a[i] + b[i] })
// }

///| add adds two vectors together and returns the result.
pub fn[T : Addable] add(a : Array[T], b : Array[T]) -> Array[T] {
  let len = a.length()
  guard len == b.length()
  Array::makei(len, fn(i) { T::add(a[i], b[i]) })
}

///| add_in_place adds two vectors together and stores the result in the first vector.
pub fn[T : Addable] add_in_place(a : Array[T], b : Array[T]) -> Unit {
  let len = a.length()
  guard len == b.length()
  for i in 0..<len {
    a[i] = T::add(a[i], b[i])
  }
}

///| subtract subtracts one vector from another and returns the result.
/// This no longer works as of "moonc v0.1.20250310+a7a1e9804".
// pub impl[T : Subtractable] Subtractable for Array[T] with sub(a, b) {
//   let len = a.length()
//   guard len == b.length()
//   Array::makei(len, fn(i) { a[i].sub(b[i]) })
// }

///| subtract subtracts one vector from another and returns the result.
pub fn[T : Subtractable] sub(a : Array[T], b : Array[T]) -> Array[T] {
  let len = a.length()
  guard len == b.length()
  Array::makei(len, fn(i) { a[i].sub(b[i]) })
}

///| subtract_in_place subtracts one vector from another and stores the result in the first vector.
pub fn[T : Subtractable] subtract_in_place(a : Array[T], b : Array[T]) -> Unit {
  let len = a.length()
  guard len == b.length()
  for i in 0..<len {
    a[i] = a[i].sub(b[i])
  }
}

///| add_number adds a number to each element of a vector and returns the result.
pub fn[T : Addable] add_number(a : Array[T], b : T) -> Array[T] {
  let len = a.length()
  Array::makei(len, fn(i) { T::add(a[i], b) })
}

///| add_number_in_place adds a number to each element of a vector and stores the result in the vector.
pub fn[T : Addable] add_number_in_place(a : Array[T], b : T) -> Unit {
  let len = a.length()
  for i in 0..<len {
    a[i] = T::add(a[i], b)
  }
}

///| subtract_number subtracts a number from each element of a vector and returns the result.
pub fn[T : Subtractable] subtract_number(a : Array[T], b : T) -> Array[T] {
  let len = a.length()
  Array::makei(len, fn(i) { a[i].sub(b) })
}

///| subtract_number_in_place subtracts a number from each element of a vector and stores the result in the vector.
pub fn[T : Subtractable] subtract_number_in_place(a : Array[T], b : T) -> Unit {
  let len = a.length()
  for i in 0..<len {
    a[i] = a[i].sub(b)
  }
}

///| multiply_number Multiplys a number to each element of a vector and returns the result.
pub fn[T : Multipliable] multiply_number(a : Array[T], b : T) -> Array[T] {
  let len = a.length()
  Array::makei(len, fn(i) { a[i] * b })
}

///| multiply_number_in_place multiplies each element of a vector by a number and stores the result in the vector.
pub fn[T : Multipliable] multiply_number_in_place(a : Array[T], b : T) -> Unit {
  let len = a.length()
  for i in 0..<len {
    a[i] *= b
  }
}

///| divide_number divides each element of a vector by a number and returns the result.
pub fn[T : Divideable] divide_number(a : Array[T], b : T) -> Array[T] {
  let len = a.length()
  Array::makei(len, fn(i) { a[i] / b })
}

///| divide_number_in_place divides each element of a vector by a number and stores the result in the vector.
pub fn[T : Divideable] divide_number_in_place(a : Array[T], b : T) -> Unit {
  let len = a.length()
  for i in 0..<len {
    a[i] /= b
  }
}

///| dot computes the dot product of two vectors.
pub fn[T : Multipliable + Addable + FromInt] dot(
  a : Array[T],
  b : Array[T]
) -> T {
  let len = a.length()
  guard len == b.length()
  let mut result = T::from_int(0)
  for i in 0..<len {
    result = T::add(result, a[i] * b[i])
  }
  result
}

///| magnitude computes the magnitude of a vector.
pub fn[T : Multipliable + Addable + SquareRootable + FromInt] magnitude(
  a : Array[T]
) -> T {
  dot(a, a).sqrt()
}

///| normalize normalizes a vector to have a magnitude of 1.
pub fn[T : Multipliable + Divideable + Addable + SquareRootable + FromInt] normalize(
  a : Array[T]
) -> Array[T] {
  let mag = magnitude(a)
  divide_number(a, mag)
}

///| sum computes the sum of all elements in a vector.
pub fn[T : Addable + FromInt] sum(a : Array[T]) -> T {
  let len = a.length()
  let mut result = T::from_int(0)
  for i in 0..<len {
    result = T::add(result, a[i])
  }
  result
}

///| product computes the product of all elements in a vector.
pub fn[T : Multipliable + FromInt] product(a : Array[T]) -> T {
  let len = a.length()
  let mut result = T::from_int(1)
  for i in 0..<len {
    result *= a[i]
  }
  result
}

///| mean computes the mean of a vector.
pub fn[T : Addable + Divideable + FromInt] mean(a : Array[T]) -> T {
  let len = T::from_int(a.length())
  sum(a) / len
}

///| min computes the minimum element in a vector.
pub fn[T : Compare] min(a : Array[T]) -> T {
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
pub fn[T : Compare] max(a : Array[T]) -> T {
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

///| abs computes the absolute value of each element in a vector.
pub fn[T : Absoluteable] abs(a : Array[T]) -> Array[T] {
  let len = a.length()
  Array::makei(len, fn(i) { a[i].abs() })
}

///| abs_in_place computes the absolute value of each element in a vector and
/// stores the result in the vector.
pub fn[T : Absoluteable] abs_in_place(a : Array[T]) -> Unit {
  let len = a.length()
  for i in 0..<len {
    a[i] = a[i].abs()
  }
}

///| euclidian_distance computes the Euclidian distance between two vectors.
pub fn[T : Multipliable + SquareRootable + Addable + Subtractable + FromInt] euclidian_distance(
  a : Array[T],
  b : Array[T]
) -> T {
  let len = a.length()
  guard len == b.length()
  let mut result = T::from_int(0)
  for i in 0..<len {
    result = T::add(result, a[i].sub(b[i]) * a[i].sub(b[i]))
  }
  result.sqrt()
}
