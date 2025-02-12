// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

// These tests are fairly straightforward.
// They are here to demonstrate how to write tests for your functions.

///|
test "hello_option_empty_string" {
  let got = hello_option_empty_string()
  inspect!(
    got,
    content=
      #|Some("")
    ,
  )

  //
  let got = hello_option_empty_string(some=false)
  inspect!(
    got,
    content=
      #|None
    ,
  )
}

///|
test "hello_option_none" {
  let got = hello_option_none(some=true)
  inspect!(
    got,
    content=
      #|Some("")
    ,
  )

  //
  let got = hello_option_none()
  inspect!(
    got,
    content=
      #|None
    ,
  )
}

///|
test "hello_option_some_string" {
  let got = hello_option_some_string()
  inspect!(
    got,
    content=
      #|Some("Hello, World!")
    ,
  )

  //
  let got = hello_option_some_string(some=false)
  inspect!(
    got,
    content=
      #|None
    ,
  )
}

///|
test "hello_world_with_arg" {
  let got = hello_world_with_arg("MoonBit")
  inspect!(got, content="Hello, MoonBit!")
}

///|
test "hello_world_with_optional_arg" {
  let got = hello_world_with_optional_arg(name="MoonBit")
  inspect!(got, content="Hello, MoonBit!")
  let got = hello_world_with_optional_arg()
  inspect!(got, content="Hello, 世界 🌍 from MoonBit!")
}

///|
test "hello_world" {
  let got = hello_world()
  inspect!(got, content="Hello, World!")
}
