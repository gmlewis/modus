/*
 * This example is part of the Modus project, licensed under the Apache License 2.0.
 * You may modify and use this example in accordance with the license.
 * See the LICENSE file that accompanied this code for further details.
 */

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gmlewis/modus/sdk/go/pkg/console"
)

// Logs a message.
func LogMessage(message string) {
	console.Log(message)
}

// Adds two integers together and returns the result.
func Add(x, y int) int {
	return x + y
}

// Adds three integers together and returns the result.
// The third integer is optional.
func Add3(a, b int, c *int) int {
	if c != nil {
		return a + b + *c
	}
	return a + b
}

// Adds any number of integers together and returns the result.
func AddN(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

// this indirection is so we can mock time.Now in tests
var nowFunc = time.Now

// Returns the current time.
func GetCurrentTime() time.Time {
	return nowFunc()
}

// Returns the current time formatted as a string.
func GetCurrentTimeFormatted() string {
	return nowFunc().Format(time.DateTime)
}

// Combines the first and last name of a person, and returns the full name.
func GetFullName(firstName, lastName string) string {
	return firstName + " " + lastName
}

// Says hello to a person by name.
// If the name is not provided, it will say hello without a name.
func SayHello(name *string) string {
	if name == nil {
		return "Hello!"
	} else {
		return "Hello, " + *name + "!"
	}
}

// A simple object representing a person.
type Person struct {

	// The person's first name.
	FirstName string `json:"firstName"`

	// The person's last name.
	LastName string `json:"lastName"`

	// The person's age.
	Age int `json:"age"`
}

// Gets a person object.
func GetPerson() Person {
	return Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       42,
	}
}

// Gets a random person object from a list of people.
func GetRandomPerson() Person {
	people := GetPeople()
	i := rand.Intn(len(people))
	return people[i]
}

// Gets a list of people.
func GetPeople() []Person {
	return []Person{
		{
			FirstName: "Bob",
			LastName:  "Smith",
			Age:       42,
		},
		{
			FirstName: "Alice",
			LastName:  "Jones",
			Age:       35,
		},
		{
			FirstName: "Charlie",
			LastName:  "Brown",
			Age:       8,
		},
	}
}

// Gets the name and age of a person.
func GetNameAndAge() (name string, age int) {
	p := GetPerson()
	return GetFullName(p.FirstName, p.LastName), p.Age
}

// Tests returning an error.
func TestNormalError(input string) (string, error) {

	// This is the preferred way to handle errors in functions.
	// Simply declare an error interface as the last return value.
	// You can use any object that implements the Go error interface.
	// For example, you can create a new error with errors.New("message"),
	// or with fmt.Errorf("message with %s", "parameters").

	if input == "" {
		return "", errors.New("input is empty")
	}
	output := "You said: " + input
	return output, nil
}

// Tests an alternative way to handle errors in functions.
func TestAlternativeError(input string) string {

	// This is an alternative way to handle errors in functions.
	// It is identical in behavior to TestNormalError, but is not Go idiomatic.

	if input == "" {
		console.Error("input is empty")
		return ""
	}
	output := "You said: " + input
	return output
}

// Tests a panic.
func TestPanic() {

	// This panics, will log the message as "fatal" and exits the function.
	// Generally, you should not panic.

	panic("This is a message from a panic.\nThis is a second line from a panic.\n")
}

// Tests an exit with a non-zero exit code.
func TestExit() {

	// If you need to exit prematurely without panicking, you can use os.Exit.
	// However, you cannot return any values from the function, so if you want
	// to log an error message, you should do so before calling os.Exit.
	// The exit code should be 0 for success, and non-zero for failure.

	console.Error("This is an error message.")
	os.Exit(1)
	println("This line will not be executed.")
}

// Tests logging at different levels.
func TestLogging() {
	// This is a simple log message. It has no level.
	console.Log("This is a simple log message.")

	// These messages are logged at different levels.
	console.Debug("This is a debug message.")
	console.Info("This is an info message.")
	console.Warn("This is a warning message.")

	// This logs an error message, but allows the function to continue.
	console.Error("This is an error message.")
	console.Error(
		`This is line 1 of a multi-line error message.
  This is line 2 of a multi-line error message.
  This is line 3 of a multi-line error message.`,
	)

	// You can also use Go's built-in printing commands.
	println("This is a println message.")
	fmt.Println("This is a fmt.Println message.")
	fmt.Printf("This is a fmt.Printf message (%s).\n", "with a parameter")

	// You can even just use stdout/stderr and the log level will be "info" or "error" respectively.
	fmt.Fprintln(os.Stdout, "This is an info message printed to stdout.")
	fmt.Fprintln(os.Stderr, "This is an error message printed to stderr.")

	// NOTE: The main difference between using console functions and Go's built-in functions, is in how newlines are handled.
	// - Using console functions, newlines are preserved and the entire message is logged as a single message.
	// - Using Go's built-in functions, each line is logged as a separate message.
	//
	// Thus, if you are logging data for debugging, we highly recommend using the console functions
	// to keep the data together in a single log message.
	//
	// The console functions also allow you to better control the reported logging level.
}
