// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

// These tests are fairly straightforward.
// They are here to demonstrate how to write tests for your functions.

test "hello_array_of_ints" {
  let got = hello_array_of_ints()
  inspect!(got, content="[1, 2, 3]")
  let got = hello_array_of_ints(n=0)
  inspect!(got, content="[]")
  let got = hello_array_of_ints(n=-1)
  inspect!(got, content="[]")
}

test "hello_array_of_ints_option" {
  let got = hello_array_of_ints_option()
  inspect!(got, content="Some([1, 2, 3])")
  let got = hello_array_of_ints_option(n=0)
  inspect!(got, content="Some([])")
  let got = hello_array_of_ints_option(n=-1)
  inspect!(got, content="None")
}
