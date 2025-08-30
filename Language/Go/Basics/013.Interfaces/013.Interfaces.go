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

// ===========================
// Nested composite interfaces
// ===========================
type Reader interface {
	Read() string
}
type Writer interface {
	Write(data string)
}
type ReadWrite interface {
	Reader
	Writer
}

// This structure will implement the ReadWrite interface
type File struct {
	data string
}

func (f *File) Read() string {
	return f.data
}
func (f *File) Write(newData string) {
	f.data = newData
}

func main() {
	// The interface in Golang is a collection of declarations of actions.

	// The interface features in Golang
	// 1. Implicit Implements
	// Golang has no keywords to declare which types implement which interfaces.
	// Once a type implements all the mehods required by an interface, it is considered to have implemented the interface.

	// 2. Interface type variable
	// Interface variables can hold any value that implements the interface.

	// 3. Nil
	// An uninitialized interface will have a zero value and will not contain any dynamic values or types.

	// 4.
	// Define an interface{} that can represent any type.

	// The interface in Golang has the following usage:
	// Polymorphism: Different types implement the same interface, achieving polymorphic behavior.
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
	fmt.Println(typeof([]int{}))
	fmt.Println()

	// Interface Combination
	// This interface can be nested and combined to implement complex action descriptions
	var rw ReadWrite = &File{}
	rw.Write("This is a new data.") // Write data
	fmt.Println(rw.Read())          // Read the data

	// The Dynamic type and dynamic value
	// Define a value using an interface
	var dynamicInterface interface{} = 114514
	// And output the dynamic interface type and value
	fmt.Printf("Type: %-10T; Value: %5v\n", dynamicInterface, dynamicInterface)
	// Again
	dynamicInterface = "Hello World."
	fmt.Printf("Type: %-10T; Value: %5v\n", dynamicInterface, dynamicInterface)
	fmt.Println()

	// The empty interface in Golang
	// When an interface variable is declared but not initialized, the interface is `nil`
	var emptyInterface interface{}
	fmt.Println(emptyInterface)
	fmt.Println(nil == emptyInterface) // Can use double equal signs to check whether the interface is `nil`

	// Type switch
	// Type switch is a Golang syntax structure that can execute different code logic according to the type of interface.
	// A simple example
	var typeSwitchInterface interface{} = 5112151121
	switch v := typeSwitchInterface.(type) {
	case int:
		fmt.Println("Interger")
	case string:
		fmt.Println("String")
	default:
		fmt.Printf("Unknow Type: %v\n", v)
	}
}
