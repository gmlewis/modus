// This example is part of the Modus project, licensed under the Apache License 2.0.
// You may modify and use this example in accordance with the license.
// See the LICENSE file that accompanied this code for further details.

///| Logs a message.
pub fn log_message(message : String) -> Unit {
  @console.log(message)
}

///| Adds two integers together and returns the result.
pub fn add(x : Int, y : Int) -> Int {
  x + y
}

///| Adds three integers together and returns the result.
/// The third integer is optional.
pub fn add3(a : Int, b : Int, c~ : Int = 0) -> Int {
  a + b + c
}

///| Adds any number of integers together and returns the result.
pub fn add_n(args : Array[Int]) -> Int {
  args.iter().fold(init=0, fn(acc, x) { acc + x })
}

///| Returns the current time.
pub fn get_current_time(
  /// now is optional so we can mock time in tests
  now~ : Int64 = @wasi.now(),
  zone~ : @time.Zone = @time.utc_zone
) -> @time.ZonedDateTime!Error {
  @time.unix!(now, zone~)
}

///| Returns the current time formatted as a string.
pub fn get_current_time_formatted(
  /// now is optional so we can mock time in tests
  now~ : Int64 = @wasi.now(),
  zone~ : @time.Zone = @time.utc_zone
) -> String!Error {
  get_current_time!(now~, zone~).to_string()
}

///| Combines the first and last name of a person, and returns the full name.
pub fn get_full_name(first_name : String, last_name : String) -> String {
  "\{first_name} \{last_name}"
}

///| Says hello to a person by name.
/// If the name is not provided, it will say hello without a name.
pub fn say_hello(name~ : String? = None) -> String {
  match name {
    Some(name) => "Hello, \{name}!"
    None => "Hello!"
  }
}

///| A simple object representing a person.
pub(all) struct Person {
  /// The person's first name.
  firstName : String
  /// The person's last name.
  lastName : String
  /// The person's age.
  age : Int
} derive(Show, Eq, FromJson, ToJson)

///| Gets a person object.
pub fn get_person() -> Person {
  { firstName: "John", lastName: "Doe", age: 42 }
}

///|
let rand : @random.Rand = @random.new()

///| Gets a random person object from a list of people.
pub fn get_random_person() -> Person {
  let people = get_people()
  let i = rand.int(limit=people.length())
  people[i]
}

///| Gets a list of people.
pub fn get_people() -> Array[Person] {
  [
    { firstName: "Bob", lastName: "Smith", age: 42 },
    { firstName: "Alice", lastName: "Jones", age: 35 },
    { firstName: "Charlie", lastName: "Brown", age: 8 },
  ]
}

///| Gets the name and age of a person.
pub fn get_name_and_age() -> (String, Int) {
  let p = get_person()
  (get_full_name(p.firstName, p.lastName), p.age)
}

///|
pub(all) type! SimpleError String derive(Show, Eq)

///| Tests returning an error.
pub fn test_normal_error(input : String) -> String!Error {
  // This is the preferred way to handle errors in functions.
  // Simply declare an !Error as part of the return value.

  if input == "" {
    raise SimpleError("input is empty")
  }
  "You said: \{input}"
}

///| Tests an alternative way to handle errors in functions.
pub fn test_alternative_error(input : String) -> String {
  // This is an alternative way to handle errors in functions.
  // It is identical in behavior to test_normal_error, but is not MoonBit idiomatic.

  if input == "" {
    @console.error("input is empty")
    return ""
  }
  "You said: \{input}"
}

///| Tests an abort.
pub fn test_abort() -> Unit {
  // This panics, will log the message as "fatal" and exits the function.
  // Generally, you should not panic.

  abort(
    #|This is a message from an abort.
    #|This is a second line from an abort.
    ,
  )
}

///| Tests an exit with a non-zero exit code.
pub fn test_exit() -> Unit {
  // If you need to exit prematurely without panicking, you can use os.Exit.
  // However, you cannot return any values from the function, so if you want
  // to log an error message, you should do so before calling os.Exit.
  // The exit code should be 0 for success, and non-zero for failure.

  @console.error("This is an error message.")
  @sys.exit(1)
  println("This line will not be executed.")
}

///| Tests logging at different levels.
pub fn test_logging() -> Unit {
  // This is a simple log message. It has no level.
  @console.log("This is a simple log message.")

  // These messages are logged at different levels.
  @console.debug("This is a debug message.")
  @console.info("This is an info message.")
  @console.warn("This is a warning message.")

  // This logs an error message, but allows the function to continue.
  @console.error("This is an error message.")
  @console.error(
    #|This is line 1 of a multi-line error message.
    #|This is line 2 of a multi-line error message.
    #|This is line 3 of a multi-line error message.
    ,
  )

  // You can also use MoonBit's built-in printing commands.
  println("This is a println message.")
  let interp = "with string interpolation"
  println("This is a println message \{interp}.")

  // NOTE: The main difference between using console functions and MoonBit's built-in functions, is in how newlines are handled.
  // - Using console functions, newlines are preserved and the entire message is logged as a single message.
  // - Using MoonBit's built-in function, each line is logged as a separate message.
  //
  // Thus, if you are logging data for debugging, we highly recommend using the console functions
  // to keep the data together in a single log message.
  //
  // The console functions also allow you to better control the reported logging level.
}
