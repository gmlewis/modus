// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

// These tests are fairly straightforward.
// They are here to demonstrate how to write tests for your functions.

///|
test "say_hello" {
  let got = say_hello(name="MoonBit")
  inspect(got, content="Hello, MoonBit!")
  let got = say_hello()
  inspect(got, content="Hello, World!")
}
