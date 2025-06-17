// Copyright 2024 Hypermode Inc.
// Licensed under the terms of the Apache License, Version 2.0
// See the LICENSE file that accompanied this code for further details.
//
// SPDX-FileCopyrightText: 2024 Hypermode Inc. <hello@hypermode.com>
// SPDX-License-Identifier: Apache-2.0

///|
pub struct TestStruct1 {
  a : Bool
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestStruct1_map {
  a : Bool
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestStruct2 {
  a : Bool
  b : Int
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestStruct2_map {
  a : Bool
  b : Int
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestStruct3 {
  a : Bool
  b : Int
  c : String
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestStruct3_map {
  a : Bool
  b : Int
  c : String
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestStruct4 {
  a : Bool
  b : Int
  c : String?
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestStruct4_map {
  a : Bool
  b : Int
  c : String?
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestStruct5 {
  a : String
  b : String
  c : String
  d : Array[String]
  e : Double
  f : Double
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestStruct5_map {
  a : String
  b : String
  c : String
  d : Array[String]
  e : Double
  f : Double
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestRecursiveStruct {
  a : Bool
  mut b : TestRecursiveStruct?
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestRecursiveStruct_map {
  a : Bool
  mut b : TestRecursiveStruct_map?
} derive(Show, Eq, FromJson, ToJson)

///|
pub struct TestSmorgasbordStruct {
  bool : Bool
  byte : Byte
  c : Char
  f : Float
  d : Double
  i16 : Int16
  i32 : Int
  i64 : Int64
  s : String
  u16 : UInt16
  u32 : UInt
  u64 : UInt64
  someBool : Bool?
  noneBool : Bool?
  someByte : Byte?
  noneByte : Byte?
  someChar : Char?
  noneChar : Char?
  someFloat : Float?
  noneFloat : Float?
  someDouble : Double?
  noneDouble : Double?
  someI16 : Int16?
  noneI16 : Int16?
  someI32 : Int?
  noneI32 : Int?
  someI64 : Int64?
  noneI64 : Int64?
  someString : String?
  noneString : String?
  someU16 : UInt16?
  noneU16 : UInt16?
  someU32 : UInt?
  noneU32 : UInt?
  someU64 : UInt64?
  noneU64 : UInt64?
} derive(Show, Eq)

///|
pub struct TestSmorgasbordStruct_map {
  bool : Bool
  byte : Byte
  c : Char
  f : Float
  d : Double
  i16 : Int16
  i32 : Int
  i64 : Int64
  s : String
  u16 : UInt16
  u32 : UInt
  u64 : UInt64
  someBool : Bool?
  noneBool : Bool?
  someByte : Byte?
  noneByte : Byte?
  someChar : Char?
  noneChar : Char?
  someFloat : Float?
  noneFloat : Float?
  someDouble : Double?
  noneDouble : Double?
  someI16 : Int16?
  noneI16 : Int16?
  someI32 : Int?
  noneI32 : Int?
  someI64 : Int64?
  noneI64 : Int64?
  someString : String?
  noneString : String?
  someU16 : UInt16?
  noneU16 : UInt16?
  someU32 : UInt?
  noneU32 : UInt?
  someU64 : UInt64?
  noneU64 : UInt64?
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
  b: None, // initialized by the init function to point to test_recursive_struct_r2
}

///|
let test_recursive_struct_r2 : TestRecursiveStruct = TestRecursiveStruct::{
  a: false,
  b: None, // initialized by the init function to point to test_recursive_struct
}

///|
fn init {
  test_recursive_struct.b = Some(test_recursive_struct_r2)
  test_recursive_struct_r2.b = Some(test_recursive_struct)
}

///|
let test_recursive_struct_map : TestRecursiveStruct_map = TestRecursiveStruct_map::{
  a: true,
  b: None, // initialized by the init() function to point to test_recursive_struct_map_r2
}

///|
let test_recursive_struct_map_r2 : TestRecursiveStruct_map = TestRecursiveStruct_map::{
  a: false,
  b: None, // initialized by the init() function to point to test_recursive_struct_map
}

///|
fn init {
  test_recursive_struct_map.b = Some(test_recursive_struct_map_r2)
  test_recursive_struct_map_r2.b = Some(test_recursive_struct_map)
}

///|
let test_smorgasbord_struct : TestSmorgasbordStruct = {
  bool: true,
  byte: 0x12,
  c: 'c',
  f: 1.23,
  d: 4.56,
  i16: 123,
  i32: 456,
  i64: 789,
  s: "abc",
  u16: 123,
  u32: 456,
  u64: 789,
  someBool: Some(true),
  noneBool: None,
  someByte: Some(0x34),
  noneByte: None,
  someChar: Some('d'),
  noneChar: None,
  someFloat: Some(7.89),
  noneFloat: None,
  someDouble: Some(0.12),
  noneDouble: None,
  someI16: Some(234),
  noneI16: None,
  someI32: Some(567),
  noneI32: None,
  someI64: Some(890),
  noneI64: None,
  someString: Some("def"),
  noneString: None,
  someU16: Some(234),
  noneU16: None,
  someU32: Some(567),
  noneU32: None,
  someU64: Some(890),
  noneU64: None,
}

///|
let test_smorgasbord_struct_map : TestSmorgasbordStruct_map = {
  bool: true,
  byte: 0x12,
  c: 'c',
  f: 1.23,
  d: 4.56,
  i16: 123,
  i32: 456,
  i64: 789,
  s: "abc",
  u16: 123,
  u32: 456,
  u64: 789,
  someBool: Some(true),
  noneBool: None,
  someByte: Some(0x34),
  noneByte: None,
  someChar: Some('d'),
  noneChar: None,
  someFloat: Some(7.89),
  noneFloat: None,
  someDouble: Some(0.12),
  noneDouble: None,
  someI16: Some(234),
  noneI16: None,
  someI32: Some(567),
  noneI32: None,
  someI64: Some(890),
  noneI64: None,
  someString: Some("def"),
  noneString: None,
  someU16: Some(234),
  noneU16: None,
  someU32: Some(567),
  noneU32: None,
  someU64: Some(890),
  noneU64: None,
}

///|
pub fn test_struct_input1(o : TestStruct1) -> Unit raise Error {
  assert_eq(test_struct1, o)
}

///|
pub fn test_struct_input2(o : TestStruct2) -> Unit raise Error {
  assert_eq(test_struct2, o)
}

///|
pub fn test_struct_input3(o : TestStruct3) -> Unit raise Error {
  assert_eq(test_struct3, o)
}

///|
pub fn test_struct_input4(o : TestStruct4) -> Unit raise Error {
  assert_eq(test_struct4, o)
}

///|
pub fn test_struct_input5(o : TestStruct5) -> Unit raise Error {
  assert_eq(test_struct5, o)
}

///|
pub fn test_struct_input4_with_none(o : TestStruct4) -> Unit raise Error {
  assert_eq(test_struct4_with_none, o)
}

///|
pub fn test_smorgasbord_struct_input(
  o : TestSmorgasbordStruct
) -> Unit raise Error {
  assert_eq(test_smorgasbord_struct, o)
}

///|
fn[A] cast_struct_to_ptr(a : A) -> Int = "%identity"

///|
pub fn test_recursive_struct_input(
  r1 : TestRecursiveStruct
) -> Unit raise Error {
  let r1_ptr = cast_struct_to_ptr(r1)
  println("test_recursive_struct_input: r1_ptr=\{r1_ptr}, r1.a=\{r1.a}")
  assert_true(r1.a)
  guard r1.b is Some(r2) else { raise TestError("Expected Some, got None") }
  let r2_ptr = cast_struct_to_ptr(r2)
  println("test_recursive_struct_input: r2_ptr=\{r2_ptr}, r2.a=\{r2.a}")
  assert_false(r2.a)
  guard r2.b is Some(r3) else { raise TestError("Expected Some, got None") }
  let r3_ptr = cast_struct_to_ptr(r3)
  println("test_recursive_struct_input: r3_ptr=\{r3_ptr}, r3.a=\{r3.a}")
  assert_eq(r1_ptr, r3_ptr)
}

///|
pub fn test_struct_option_input1(o : TestStruct1?) -> Unit raise Error {
  assert_eq(Some(test_struct1), o)
}

///|
pub fn test_struct_option_input2(o : TestStruct2?) -> Unit raise Error {
  assert_eq(Some(test_struct2), o)
}

///|
pub fn test_struct_option_input3(o : TestStruct3?) -> Unit raise Error {
  assert_eq(Some(test_struct3), o)
}

///|
pub fn test_struct_option_input4(o : TestStruct4?) -> Unit raise Error {
  assert_eq(Some(test_struct4), o)
}

///|
pub fn test_struct_option_input5(o : TestStruct5?) -> Unit raise Error {
  assert_eq(Some(test_struct5), o)
}

///|
pub fn test_struct_option_input4_with_none(
  o : TestStruct4?
) -> Unit raise Error {
  assert_eq(Some(test_struct4_with_none), o)
}

///|
pub fn test_recursive_struct_option_input(
  o : TestRecursiveStruct?
) -> Unit raise Error {
  guard o is Some(r1) else { raise TestError("Expected Some, got None") }
  assert_true(r1.a)
  guard r1.b is Some(r2) else { raise TestError("Expected Some, got None") }
  assert_false(r2.a)
  guard r2.b is Some(r3) else { raise TestError("Expected Some, got None") }
  let r1_ptr = cast_struct_to_ptr(r1)
  let r3_ptr = cast_struct_to_ptr(r3)
  assert_eq(r1_ptr, r3_ptr)
}

///|
pub fn test_smorgasbord_struct_option_input(
  o : TestSmorgasbordStruct?
) -> Unit raise Error {
  assert_eq(Some(test_smorgasbord_struct), o)
}

///|
pub fn test_struct_option_input1_none(o : TestStruct1?) -> Unit raise Error {
  assert_eq(o, None)
}

///|
pub fn test_struct_option_input2_none(o : TestStruct2?) -> Unit raise Error {
  assert_eq(o, None)
}

///|
pub fn test_struct_option_input3_none(o : TestStruct3?) -> Unit raise Error {
  assert_eq(o, None)
}

///|
pub fn test_struct_option_input4_none(o : TestStruct4?) -> Unit raise Error {
  assert_eq(o, None)
}

///|
pub fn test_struct_option_input5_none(o : TestStruct5?) -> Unit raise Error {
  assert_eq(o, None)
}

///|
pub fn test_recursive_struct_option_input_none(
  o : TestRecursiveStruct?
) -> Unit raise Error {
  assert_eq(o, None)
}

///|
pub fn test_smorgasbord_struct_option_input_none(
  o : TestSmorgasbordStruct?
) -> Unit raise Error {
  assert_eq(o, None)
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
pub fn test_recursive_struct_output() -> TestRecursiveStruct raise Error {
  let r1 = test_recursive_struct
  println("test_recursive_struct_output: r1.a=\{r1.a}")
  match r1.b {
    Some(r2) => {
      println("test_recursive_struct_output: r2.a=\{r2.a}")
      match r2.b {
        Some(r3) => {
          println("test_recursive_struct_output: r3.a=\{r3.a}")
          let r1_ptr = cast_struct_to_ptr(r1)
          let r2_ptr = cast_struct_to_ptr(r2)
          let r3_ptr = cast_struct_to_ptr(r3)
          println("test_recursive_struct_output: r1_ptr=\{r1_ptr}")
          println("test_recursive_struct_output: r2_ptr=\{r2_ptr}")
          println("test_recursive_struct_output: r3_ptr=\{r3_ptr}")
        }
        _ => raise TestError("Expected Some, got None")
      }
    }
    _ => ()
  }
  test_recursive_struct
}

///|
pub fn test_recursive_struct_output_map() -> TestRecursiveStruct_map {
  test_recursive_struct_map
}

///|
pub fn test_smorgasbord_struct_output() -> TestSmorgasbordStruct {
  test_smorgasbord_struct
}

///|
pub fn test_smorgasbord_struct_output_map() -> TestSmorgasbordStruct_map {
  test_smorgasbord_struct_map
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
pub fn test_smorgasbord_struct_option_output() -> TestSmorgasbordStruct? {
  Some(test_smorgasbord_struct)
}

///|
pub fn test_smorgasbord_struct_option_output_map() -> TestSmorgasbordStruct_map? {
  Some(test_smorgasbord_struct_map)
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

///|
pub fn test_smorgasbord_struct_option_output_none() -> TestSmorgasbordStruct? {
  None
}

///|
pub fn test_smorgasbord_struct_option_output_map_none() -> TestSmorgasbordStruct_map? {
  None
}
