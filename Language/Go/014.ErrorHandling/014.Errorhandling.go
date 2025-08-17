package main

import (
	"errors"
	"fmt"
	"math"
)

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
	example2, _ := Sqrt(4)
	fmt.Println(example2)

}
