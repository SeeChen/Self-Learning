package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Java interface {
	Java() string
}
type Golang interface {
	Golang() string
}

type Programming struct {
}

func (p Programming) Java() string {
	return "Java is the best programming language in the world."
}
func (p Programming) Golang() string {
	return "Golang is the best programming language in the world."
}

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
	fmt.Printf("Before: [%s] %v; After: [%s] %v.\n", reflect.TypeOf(strCorrect).String(), strCorrect, reflect.TypeOf(num1).String(), num1)

	num2, err := strconv.Atoi(strError)
	if err != nil {
		// Use it to detect error
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(num2)
	}

	fmt.Println()
	// From numeric to a string
	// To convert a string to a float type, the specified bit size must be given in the second parameter.
	var piStr string = "3.1415926"
	pi, _ := strconv.ParseFloat(piStr, 64)
	fmt.Printf("Before: [%s] %v, After: [%s] %v.\n\n", reflect.TypeOf(piStr).String(), piStr, reflect.TypeOf(pi).String(), pi)

	// Interface type conversion
	// In Golang, there are two types of interface conversions
	// 1. Type Assertion
	// Type Assertion can convert interface type variables into special types
	var java Java = Programming{}
	fmt.Println(java.Java())

	// Do the convert
	golang, _ := java.(Golang)
	fmt.Printf("Before: [%s] %v\nAfter : [%s] %v\n", reflect.TypeOf(java).String(), java.Java(), reflect.TypeOf(golang).String(), golang.Golang())

	// 2. Type Conversion
	// Type conversion converts an interface type variable to another interface type variable/
	// The syntax is `T(value)`, where "T" represents the target interface type and "value" represents the value to be converted.
	fmt.Println()
	var ProgrammingLanguage Programming = Programming{}

	fmt.Println(Java(ProgrammingLanguage).Java())
	fmt.Println(Golang(ProgrammingLanguage).Golang())
}
