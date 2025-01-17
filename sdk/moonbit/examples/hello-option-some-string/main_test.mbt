// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

// These tests are fairly straightforward.
// They are here to demonstrate how to write tests for your functions.

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
