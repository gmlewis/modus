// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///|
pub fn hello_world_with_optional_arg(name~ : String = "World") -> String {
  "Hello, \{name}!"
}
