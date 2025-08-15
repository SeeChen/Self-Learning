package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Type conversion allows us to convert a variable type to another variable type.

	// Number type conversion
	// In Golang, using type_name(expression) for number type conversion.
	// Define two intergers
	var a, b int = 57257, 5000
	// Perform division, the output will be an interger
	fmt.Printf("%d / %d = %v\n", a, b, a/b)

	// But after type conversion, the output will be float type.
	fmt.Printf("%d / %d = %v\n\n", a, b, float64(a)/float64(b))

	// String type conversion
	// To convert a string variable to a numeric type, or to convert from a number to a string
	// In Golang, must be use the package strconv
	// From a string to an int
	var strCorrect string = "5112151121"
	var strError string = "This is Golang"
	// Atoi will return two values, one is the converted number, the other is the error message
	// But we can use _ to ignore error message.
	num1, _ := strconv.Atoi(strCorrect)
	fmt.Printf("Before: [%s]%v; After: [%s]%v.\n", reflect.TypeOf(strCorrect).String(), strCorrect, reflect.TypeOf(num1).String(), num1)

	num2, err := strconv.Atoi(strError)
	if err != nil {
		// Use it to detect error
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(num2)
	}
}
