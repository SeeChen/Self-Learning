package main

import (
	"fmt"
	"reflect"
)

// Define an interface
type firstInterface interface {
	method() string
}

// And define a struct to implement the interface method
type firstStruct struct {
	str string
}

// Define a function to implement the interface
func (firstStruct firstStruct) method() string {
	return "This is return from interface method, " + firstStruct.str
}

// Define a function that uses an empty interface to receive any interface type.
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func main() {
	// The interface in Golang is a collection of declarations of actions.

	// The interface features in Golang
	// 1. Implicit Implements
	// Golang has no keywords to declare which types implement which interfaces.
	// Once a type implements all the mehods required by an interface, it is considered to have implemented the interface.

	// 2. Interface type variable
	// Interface variables can hole any value that implements the interface.

	// 3. Nil
	// An uninitialized interface will have a zero value and will not contain any dynamic values or types.

	// 4.
	// Define an interface{} that can represent any type.

	// The interface in Golang has the following usage:
	// Polymorphism: Different types implement the same interface, achievinng polymorphic behavior.
	// Decoupling: Dependencies are defined through interfaces, reducing coupling between modules.
	// Generalization: Use the empty interface interface{} to represent any type.
	var firstInter firstInterface = firstStruct{str: "First Interface"}
	fmt.Println(firstInter.method())
	fmt.Println()

	// Nil interface
	// The empty interface is a special interface type in Golang
	// It's a superset that can represent any type.
	fmt.Println(typeof("Hello World"))
	fmt.Println(typeof(1234))
	fmt.Println(typeof(true))
	fmt.Println([]int{})
}
