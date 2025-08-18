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

func (me *MyError) Error() error {
	var formatStr string = `
	Error - %d.
	Error Message: %d.
	Error Code: %s.
	`

	return errors.New(fmt.Sprintf(formatStr, me.ErrorCode, me.ErrorCode, me.ErrorMsgs))
}

func search(keyword string) (string, error) {

	if !strings.Contains(keyword, "SeeChen") {

		myError := MyError{ErrorCode: 520, ErrorMsgs: "Not Contains \"SeeChen\"."}
		return "", myError.Error()
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
}
