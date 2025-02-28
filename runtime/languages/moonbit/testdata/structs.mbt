// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub struct TestStruct1 {
  a : Bool
} derive(Show, Eq)

///|
pub struct TestStruct1_map {
  a : Bool
} derive(Show, Eq)

///|
pub struct TestStruct2 {
  a : Bool
  b : Int
} derive(Show, Eq)

///|
pub struct TestStruct2_map {
  a : Bool
  b : Int
} derive(Show, Eq)

///|
pub struct TestStruct3 {
  a : Bool
  b : Int
  c : String
} derive(Show, Eq)

///|
pub struct TestStruct3_map {
  a : Bool
  b : Int
  c : String
} derive(Show, Eq)

///|
pub struct TestStruct4 {
  a : Bool
  b : Int
  c : String?
} derive(Show, Eq)

///|
pub struct TestStruct4_map {
  a : Bool
  b : Int
  c : String?
} derive(Show, Eq)

///|
pub struct TestStruct5 {
  a : String
  b : String
  c : String
  d : Array[String]
  e : Double
  f : Double
} derive(Show, Eq)

///|
pub struct TestStruct5_map {
  a : String
  b : String
  c : String
  d : Array[String]
  e : Double
  f : Double
} derive(Show, Eq)

///|
pub struct TestRecursiveStruct {
  a : Bool
  mut b : TestRecursiveStruct?
} derive(Show, Eq)

///|
pub struct TestRecursiveStruct_map {
  a : Bool
  mut b : TestRecursiveStruct_map?
} derive(Show, Eq)

///|
let test_struct1 : TestStruct1 = TestStruct1::{ a: true }

///|
let test_struct1_map : TestStruct1_map = TestStruct1_map::{ a: true }

///|
let test_struct2 : TestStruct2 = TestStruct2::{ a: true, b: 123 }

///|
let test_struct2_map : TestStruct2_map = TestStruct2_map::{ a: true, b: 123 }

///|
let test_struct3 : TestStruct3 = TestStruct3::{ a: true, b: 123, c: "abc" }

///|
let test_struct3_map : TestStruct3_map = TestStruct3_map::{
  a: true,
  b: 123,
  c: "abc",
}

///|
let test_struct4 : TestStruct4 = TestStruct4::{
  a: true,
  b: 123,
  c: Some("abc"),
}

///|
let test_struct4_map : TestStruct4_map = TestStruct4_map::{
  a: true,
  b: 123,
  c: Some("abc"),
}

///|
let test_struct4_with_none : TestStruct4 = TestStruct4::{
  a: true,
  b: 123,
  c: None,
}

///|
let test_struct4_map_with_none : TestStruct4_map = TestStruct4_map::{
  a: true,
  b: 123,
  c: None,
}

///|
let test_struct5 : TestStruct5 = TestStruct5::{
  a: "abc",
  b: "def",
  c: "ghi",
  d: ["jkl", "mno", "pqr"],
  e: 0.12345,
  f: 99.99999,
}

///|
let test_struct5_map : TestStruct5_map = TestStruct5_map::{
  a: "abc",
  b: "def",
  c: "ghi",
  d: ["jkl", "mno", "pqr"],
  e: 0.12345,
  f: 99.99999,
}

///|
let test_recursive_struct : TestRecursiveStruct = TestRecursiveStruct::{
  a: true,
  b: None, // initialized by the init() function to point to itself.
}

///|
fn init {
  test_recursive_struct.b = Some(test_recursive_struct)
}

///|
let test_recursive_struct_map : TestRecursiveStruct_map = TestRecursiveStruct_map::{
  a: true,
  b: None, // initialized by the init() function to point to itself.
}

///|
fn init {
  test_recursive_struct_map.b = Some(test_recursive_struct_map)
}

///|
pub fn test_struct_input1(o : TestStruct1) -> Unit!Error {
  assert_eq!(test_struct1, o)
}

///|
pub fn test_struct_input2(o : TestStruct2) -> Unit!Error {
  assert_eq!(test_struct2, o)
}

///|
pub fn test_struct_input3(o : TestStruct3) -> Unit!Error {
  assert_eq!(test_struct3, o)
}

///|
pub fn test_struct_input4(o : TestStruct4) -> Unit!Error {
  assert_eq!(test_struct4, o)
}

///|
pub fn test_struct_input5(o : TestStruct5) -> Unit!Error {
  assert_eq!(test_struct5, o)
}

///|
pub fn test_struct_input4_with_none(o : TestStruct4) -> Unit!Error {
  assert_eq!(test_struct4_with_none, o)
}

///|
pub fn test_recursive_struct_input(o : TestRecursiveStruct) -> Unit!Error {
  assert_eq!(test_recursive_struct, o)
}

///|
pub fn test_struct_option_input1(o : TestStruct1?) -> Unit!Error {
  assert_eq!(Some(test_struct1), o)
}

///|
pub fn test_struct_option_input2(o : TestStruct2?) -> Unit!Error {
  assert_eq!(Some(test_struct2), o)
}

///|
pub fn test_struct_option_input3(o : TestStruct3?) -> Unit!Error {
  assert_eq!(Some(test_struct3), o)
}

///|
pub fn test_struct_option_input4(o : TestStruct4?) -> Unit!Error {
  assert_eq!(Some(test_struct4), o)
}

///|
pub fn test_struct_option_input5(o : TestStruct5?) -> Unit!Error {
  assert_eq!(Some(test_struct5), o)
}

///|
pub fn test_struct_option_input4_with_none(o : TestStruct4?) -> Unit!Error {
  assert_eq!(Some(test_struct4_with_none), o)
}

///|
pub fn test_recursive_struct_option_input(
  o : TestRecursiveStruct?
) -> Unit!Error {
  assert_eq!(Some(test_recursive_struct), o)
}

///|
pub fn test_struct_option_input1_none(o : TestStruct1?) -> Unit!Error {
  assert_eq!(o, None)
}

///|
pub fn test_struct_option_input2_none(o : TestStruct2?) -> Unit!Error {
  assert_eq!(o, None)
}

///|
pub fn test_struct_option_input3_none(o : TestStruct3?) -> Unit!Error {
  assert_eq!(o, None)
}

///|
pub fn test_struct_option_input4_none(o : TestStruct4?) -> Unit!Error {
  assert_eq!(o, None)
}

///|
pub fn test_struct_option_input5_none(o : TestStruct5?) -> Unit!Error {
  assert_eq!(o, None)
}

///|
pub fn test_recursive_struct_option_input_none(
  o : TestRecursiveStruct?
) -> Unit!Error {
  assert_eq!(o, None)
}

///|
pub fn test_struct_output1() -> TestStruct1 {
  test_struct1
}

///|
pub fn test_struct_output1_map() -> TestStruct1_map {
  test_struct1_map
}

///|
pub fn test_struct_output2() -> TestStruct2 {
  test_struct2
}

///|
pub fn test_struct_output2_map() -> TestStruct2_map {
  test_struct2_map
}

///|
pub fn test_struct_output3() -> TestStruct3 {
  test_struct3
}

///|
pub fn test_struct_output3_map() -> TestStruct3_map {
  test_struct3_map
}

///|
pub fn test_struct_output4() -> TestStruct4 {
  test_struct4
}

///|
pub fn test_struct_output4_map() -> TestStruct4_map {
  test_struct4_map
}

///|
pub fn test_struct_output5() -> TestStruct5 {
  test_struct5
}

///|
pub fn test_struct_output5_map() -> TestStruct5_map {
  test_struct5_map
}

///|
pub fn test_struct_output4_with_none() -> TestStruct4 {
  test_struct4_with_none
}

///|
pub fn test_struct_output4_map_with_none() -> TestStruct4_map {
  test_struct4_map_with_none
}

///|
pub fn test_recursive_struct_output() -> TestRecursiveStruct {
  test_recursive_struct
}

///|
pub fn test_recursive_struct_output_map() -> TestRecursiveStruct_map {
  test_recursive_struct_map
}

///|
pub fn test_struct_option_output1() -> TestStruct1? {
  Some(test_struct1)
}

///|
pub fn test_struct_option_output1_map() -> TestStruct1_map? {
  Some(test_struct1_map)
}

///|
pub fn test_struct_option_output2() -> TestStruct2? {
  Some(test_struct2)
}

///|
pub fn test_struct_option_output2_map() -> TestStruct2_map? {
  Some(test_struct2_map)
}

///|
pub fn test_struct_option_output3() -> TestStruct3? {
  Some(test_struct3)
}

///|
pub fn test_struct_option_output3_map() -> TestStruct3_map? {
  Some(test_struct3_map)
}

///|
pub fn test_struct_option_output4() -> TestStruct4? {
  Some(test_struct4)
}

///|
pub fn test_struct_option_output4_map() -> TestStruct4_map? {
  Some(test_struct4_map)
}

///|
pub fn test_struct_option_output5() -> TestStruct5? {
  Some(test_struct5)
}

///|
pub fn test_struct_option_output5_map() -> TestStruct5_map? {
  Some(test_struct5_map)
}

///|
pub fn test_struct_option_output4_with_none() -> TestStruct4? {
  Some(test_struct4_with_none)
}

///|
pub fn test_struct_option_output4_map_with_none() -> TestStruct4_map? {
  Some(test_struct4_map_with_none)
}

///|
pub fn test_recursive_struct_option_output() -> TestRecursiveStruct? {
  Some(test_recursive_struct)
}

///|
pub fn test_recursive_struct_option_output_map() -> TestRecursiveStruct_map? {
  Some(test_recursive_struct_map)
}

///|
pub fn test_struct_option_output1_none() -> TestStruct1? {
  None
}

///|
pub fn test_struct_option_output2_none() -> TestStruct2? {
  None
}

///|
pub fn test_struct_option_output3_none() -> TestStruct3? {
  None
}

///|
pub fn test_struct_option_output4_none() -> TestStruct4? {
  None
}

///|
pub fn test_struct_option_output5_none() -> TestStruct5? {
  None
}

///|
pub fn test_recursive_struct_option_output_none() -> TestRecursiveStruct? {
  None
}
