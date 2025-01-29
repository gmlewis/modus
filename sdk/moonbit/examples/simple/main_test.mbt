// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

// These tests are fairly straightforward.
// They are here to demonstrate how to write tests for your functions.

test "log_message" {
  let msg = "Hello, World!"
  log_message(msg)

  //
  let values = @console.log_call_stack.pop()
  inspect!(
    values,
    content=
      #|["", "Hello, World!"]
    ,
  )
}

test "add" {
  let result = add(1, 2)
  inspect!(result, content="3")
}

test "add3" {
  let a = 1
  let b = 2
  let c = 3
  let result = add3(a, b, c~)
  inspect!(result, content="6")
}

test "add_n" {
  // Test case 1: Empty arguments
  let result = add_n([])
  inspect!(result, content="0")

  // Test case 2: Single argument
  let result = add_n([5])
  inspect!(result, content="5")

  // Test case 3: Multiple arguments
  let result = add_n([1, 2, 3, 4, 5])
  inspect!(result, content="15")

  // Test case 4: Negative numbers
  let result = add_n([-1, -2, -3])
  inspect!(result, content="-6")
}

test "get_current_time" {
  let zone = @time.fixed_zone!("EST", -5 * 60 * 60)
  let result = get_current_time!(now=1704067200, zone~)
  inspect!(result, content="2023-12-31T19:00:00-05:00[EST]")
}

test "get_current_time_formatted" {
  let zone = @time.fixed_zone!("PST", -8 * 60 * 60)
  let result = get_current_time_formatted!(now=1704067200, zone~)
  inspect!(result, content="2023-12-31T16:00:00-08:00[PST]")
}

test "get_full_name" {
  let first_name = "John"
  let last_name = "Doe"
  let result = get_full_name(first_name, last_name)
  inspect!(result, content="John Doe")
}

test "say_hello" {
  let name = Some("Alice")
  let result = say_hello(name~)
  inspect!(result, content="Hello, Alice!")
  let result = say_hello()
  inspect!(result, content="Hello!")
}

test "get_person" {
  let result = get_person()
  inspect!(result, content="{firstName: \"John\", lastName: \"Doe\", age: 42}")
}

test "get_random_person" {
  // The same seed is always used, so the result is deterministic.
  let result = get_random_person()
  inspect!(
    result,
    content="{firstName: \"Charlie\", lastName: \"Brown\", age: 8}",
  )
}

test "get_people" {
  let result = get_people()
  inspect!(
    result,
    content="[{firstName: \"Bob\", lastName: \"Smith\", age: 42}, {firstName: \"Alice\", lastName: \"Jones\", age: 35}, {firstName: \"Charlie\", lastName: \"Brown\", age: 8}]",
  )
}

test "get_name_and_age" {
  let (name, age) = get_name_and_age()
  inspect!(name, content="John Doe")
  inspect!(age, content="42")
}
