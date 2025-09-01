package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// This is a custom interface that extends the error interface
type MyError struct {
	ErrorCode int
	ErrorMsgs string
}

func (me *MyError) Error() string {
	var formatStr string = `
	Error - %d.
	Error Message: %d.
	Error Code: %s.
	`

	return fmt.Sprintf(formatStr, me.ErrorCode, me.ErrorCode, me.ErrorMsgs)
}

func search(keyword string) (string, error) {

	if !strings.Contains(keyword, "SeeChen") {

		myError := &MyError{ErrorCode: 520, ErrorMsgs: "Not Contains \"SeeChen\"."}
		return "", myError
	}

	return "SeeChen is Handsome", nil
}

// Use error in function
// This is an example
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}

	return math.Sqrt(f), nil
}

// Here is the errors.Is Example
var ErrorSeeChen error = errors.New("seechen seechen seechen")

func SeeChen(str string) error {
	return fmt.Errorf("%w", ErrorSeeChen)
}

func safeFunction(str string) (results string) {
	// Must use defer to declare the function before panic
	// Otherwise the panic error cannot be cought
	defer func() {
		if r := recover(); r != nil {
			results = fmt.Sprintf("Error: %s, Panic Reason: %v", str, r)
		}
	}()

	if !strings.Contains(str, "SeeChen") {
		panic("The string does not contain 'SeeChen'!")
	}

	results = fmt.Sprintf("Yes!! %s\n", str)
	return
}

func main() {
	// In the Golang standard library, an error interface is defined, which is an abstraction respresenting an error.
	// This is its definition
	// type error interface {
	// 		Error() string
	// }

	// We can use the error package in our code to create an error interface and display error messages
	err := errors.New("this is an error")
	fmt.Println(err)
	fmt.Println()

	// Starting from Go 1.20, we can use the Join function to merge multiple errors into one output.
	err1 := errors.New("First")
	err2 := errors.New("Second")

	errMerge := errors.Join(err1, err2)
	fmt.Println(errMerge)
	fmt.Println()

	// Try to use the sqrt function
	example1, err := Sqrt(-1)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(example1)
	}
	fmt.Println()

	// Or use underscore to ignore error messages
	example2, _ := Sqrt(-4)
	fmt.Println(example2)

	// Custom Error
	// We can extend the error interface with custom types
	_, err = search("Oh! Yeah~")
	if err != nil {
		fmt.Println(err)
	}

	// fmt package formatting error.
	// Golang fmt package supports error formatting output
	// %v : Default output
	// %+v: All verbose messages are displayed if supported.
	// %s : Display the message as a string.
	results, err := search("Lee")
	if err != nil {
		// Default output
		fmt.Printf("%v\n\n", err)

		// Displayed all verbose
		fmt.Printf("%+v\n\n", err)

		// Displayed as a string
		fmt.Printf("%s\n\n", err)
	} else {
		fmt.Println(results)
	}

	// Starting from Go 1.13, the package error imports errors.Is and errors.As to handle error chains.
	// errors.Is: Checks whether an error is a specific error or is wrapped by an errors.
	err = SeeChen("123")
	if errors.Is(err, ErrorSeeChen) {
		fmt.Printf("Yes, %v\n", err)
	} else {
		fmt.Printf("No, %v\n", err)
	}

	// errors.As: Convert the error to a special type to facilitate the next step of processing.
	err = &MyError{ErrorCode: 404, ErrorMsgs: "Not Found"}
	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Printf("Error Code: %d\nError Message: %s\n", myErr.ErrorCode, myErr.ErrorMsgs)
	}

	// panic and recover
	// `Panic` in Golang can handle unrecoverable errors, and the `recover` will recover from panic.

	// PANIC
	// Causes the program to crash and prints a stack trace
	// Commonly used when the program cannot continue

	// RECOVER
	// Catch `panic` to prevent program crash.
	fmt.Println()
	fmt.Println(safeFunction("SeeChen"))
	fmt.Println(safeFunction("Lee"))

	fmt.Println("Continue to run!")
	// Look like try catch finally
}
