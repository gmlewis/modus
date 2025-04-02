///|
test "CallStack::new/empty" {
  let stack : CallStack[Int] = CallStack::new()
  inspect!(stack.items, content="[]")
}

///|
test "CallStack::new/different_types" {
  let int_stack : CallStack[Int] = CallStack::new()
  let str_stack : CallStack[String] = CallStack::new()
  let bool_stack : CallStack[Bool] = CallStack::new()
  inspect!(int_stack.items, content="[]")
  inspect!(str_stack.items, content="[]")
  inspect!(bool_stack.items, content="[]")
}

///|
test "CallStack::new/unit_type" {
  let unit_stack : CallStack[Unit] = CallStack::new()
  inspect!(unit_stack.items, content="[]")
}

///|
test "CallStack::new/basic" {
  let call_stack : CallStack[Int] = CallStack::new()
  inspect!(call_stack, content="{items: []}")
}

///|
test "CallStack::new/string" {
  let call_stack : CallStack[String] = CallStack::new()
  inspect!(call_stack, content="{items: []}")
}

///|
test "CallStack::new/empty_call_stack" {
  let call_stack : CallStack[Unit] = CallStack::new()
  inspect!(call_stack.size(), content="0")
}

///|
test "CallStack::size/empty" {
  let stack : CallStack[Int] = CallStack::new()
  inspect!(stack.size(), content="0")
}

///|
test "CallStack::size/after_push" {
  let stack : CallStack[Int] = CallStack::new()
  stack.push([1, 2, 3])
  stack.push([4, 5])
  inspect!(stack.size(), content="2")
}

///|
test "CallStack::size/after_pop" {
  let stack : CallStack[Int] = CallStack::new()
  stack.push([1])
  stack.push([2])
  let _ = stack.pop()
  inspect!(stack.size(), content="1")
}

///|
test "CallStack::size/empty_stack" {
  let stack : CallStack[Int] = CallStack::new()
  inspect!(stack.size(), content="0")
}

///|
test "CallStack::size/one_element" {
  let stack : CallStack[Int] = CallStack::new()
  let values = [42]
  @testutils.push(stack, values)
  inspect!(stack.size(), content="1")
}

///|
test "CallStack::size/multiple_elements" {
  let stack : CallStack[Int] = CallStack::new()
  for i = 1; i <= 5; i = i + 1 {
    let values = [i]
    @testutils.push(stack, values)
  }
  inspect!(stack.size(), content="5")
}

///|
test "@testutils.push/basic" {
  let stack = CallStack::new()
  let values = [1, 2, 3]
  @testutils.push(stack, values)
  inspect!(stack.size(), content="1")
}

///|
test "@testutils.push/empty_array" {
  let stack = CallStack::new()
  let empty : Array[Int] = []
  @testutils.push(stack, empty)
  inspect!(stack.size(), content="1")
}

///|
test "@testutils.push/multiple" {
  let stack = CallStack::new()
  let values1 = [1, 2]
  let values2 = [3, 4]
  @testutils.push(stack, values1)
  @testutils.push(stack, values2)
  inspect!(stack.size(), content="2")
}

///|
test "@testutils.push/single_push" {
  let stack : CallStack[Int] = CallStack::new()
  @testutils.push(stack, [1])
  inspect!(stack.size(), content="1")
  let item : @moonbitlang/core/builtin.Array[Int]? = @moonbitlang/core/builtin.Array::get(
    stack.items,
    0,
  )
  inspect!(item, content="Some([1])")
}

///|
test "@testutils.push/multiple_pushes" {
  let stack : CallStack[String] = CallStack::new()
  @testutils.push(stack, ["first"])
  @testutils.push(stack, ["second", "third"])
  inspect!(stack.size(), content="2")
  let item1 : @moonbitlang/core/builtin.Array[String]? = @moonbitlang/core/builtin.Array::get(
    stack.items,
    0,
  )
  let item2 : @moonbitlang/core/builtin.Array[String]? = @moonbitlang/core/builtin.Array::get(
    stack.items,
    1,
  )
  inspect!(item1, content="Some([\"first\"])")
  inspect!(item2, content="Some([\"second\", \"third\"])")
}

///|
test "@testutils.push/empty_push" {
  let stack : CallStack[Char] = CallStack::new()
  @testutils.push(stack, [])
  inspect!(stack.size(), content="1")
  let item : @moonbitlang/core/builtin.Array[Char]? = @moonbitlang/core/builtin.Array::get(
    stack.items,
    0,
  )
  inspect!(item, content="Some([])")
}

///|
test "CallStack::pop/basic" {
  let stack : CallStack[Int] = CallStack::new()
  stack.push([1, 2, 3])
  inspect!(stack.pop(), content="[1, 2, 3]")
}

///|
test "CallStack::pop/empty" {
  let stack : CallStack[Int] = CallStack::new()
  inspect!(stack.pop(), content="[]")
}

///|
test "CallStack::pop/multiple" {
  let stack : CallStack[Int] = CallStack::new()
  stack.push([1, 2])
  stack.push([3, 4])
  inspect!(stack.pop(), content="[3, 4]")
  inspect!(stack.pop(), content="[1, 2]")
  inspect!(stack.pop(), content="[]")
}
