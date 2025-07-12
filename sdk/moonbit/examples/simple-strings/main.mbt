// -*- compile-command: "./build.sh"; -*-

// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

// 10000: "\ff\ff\ff\ff\06\00\00PS\00t\00r\00i\00n\00g\00\00\00\00\00\ff\ff\ff\ff\00\00\00\00\00\00\00\00\00\00\00\00"
// 10040: "\ff\ff\ff\ff\0a\00\00P1\002\003\004\005\006\007\008\009\000\00\00\00\00\00"
// 10072: "\ff\ff\ff\ff\09\00\00P1\002\003\004\005\006\007\008\009\00\00\00\00\00\00\00"
// 10104: "\ff\ff\ff\ff\08\00\00P1\002\003\004\005\006\007\008\00\00\00\00\00\00\00\00\00"
// 10136: "\ff\ff\ff\ff\07\00\00P1\002\003\004\005\006\007\00\00\00"
// 10160: "\ff\ff\ff\ff\06\00\00P1\002\003\004\005\006\00\00\00\00\00"
// 10184: "\ff\ff\ff\ff\05\00\00P1\002\003\004\005\00\00\00\00\00\00\00"
// 10208: "\ff\ff\ff\ff\04\00\00P1\002\003\004\00\00\00\00\00\00\00\00\00"
// 10232: "\ff\ff\ff\ff\03\00\00P1\002\003\00\00\00"
// 10248: "\ff\ff\ff\ff\02\00\00P1\002\00\00\00\00\00"
// 10264: "\ff\ff\ff\ff\01\00\00P1\00\00\00\00\00\00\00"
// 10280: "\ff\ff\ff\ff\00\00\00P\00\00\00\00\00\00\00\00"

///|
pub fn test_string_output_len_0() -> String {
  ""
}

///|
pub fn test_string_output_len_1() -> String {
  "1"
}

///|
pub fn test_string_output_len_2() -> String {
  "12"
}

///|
pub fn test_string_output_len_3() -> String {
  "123"
}

///|
pub fn test_string_output_len_4() -> String {
  "1234"
}

///|
pub fn test_string_output_len_5() -> String {
  "12345"
}

///|
pub fn test_string_output_len_6() -> String {
  "123456"
}

///|
pub fn test_string_output_len_7() -> String {
  "1234567"
}

///|
pub fn test_string_output_len_8() -> String {
  "12345678"
}

///|
pub fn test_string_output_len_9() -> String {
  "123456789"
}

///|
pub fn test_string_output_len_10() -> String {
  "1234567890"
}
