///|
/// Debug function that shows us exactly what MoonBit is seeing
pub fn debug_array_bool_input(val : Array[Bool]) -> String {
  let mut result = ""
  result = result +
    "MoonBit received array with length: " +
    val.length().to_string() +
    "\n"
  result = result + "Array contents: ["
  for i = 0; i < val.length(); i = i + 1 {
    if i > 0 {
      result = result + ", "
    }
    result = result + val[i].to_string()
  }
  result = result + "]\n"
  result = result + "Array pointer: " + cast_to_ptr(val).to_string() + "\n"

  // Also create a fresh array for comparison
  let fresh_true = [true]
  result = result + "Fresh [true] array: [" + fresh_true[0].to_string() + "]\n"
  result = result +
    "Fresh [true] pointer: " +
    cast_to_ptr(fresh_true).to_string() +
    "\n"
  result = result + "Arrays equal: " + (val == fresh_true).to_string() + "\n"
  result
}

///|
/// Test function that uses the debug function
pub fn test_debug_array_bool_1(val : Array[Bool]) -> Unit raise Error {
  let debug_info = debug_array_bool_input(val)
  println(debug_info)

  // Now do the actual comparison
  let expected = test_array_output_bool_1()
  println("Expected: [" + expected[0].to_string() + "]")
  println("Got: [" + val[0].to_string() + "]")
  println("Expected pointer: " + cast_to_ptr(expected).to_string())
  println("Got pointer: " + cast_to_ptr(val).to_string())
  assert_eq(expected.length(), val.length())
  assert_eq(expected, val)
}
