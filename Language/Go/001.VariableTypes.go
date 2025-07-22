package main

import (
	"fmt"
	"reflect"
)

// A typeof function
// Returns a string describing the variable type of the passed parameter.
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func main() {
	
	// The basic data type of Golang
	
	// 1. Bool
	// The bool type has only two types of data: true and false.
	var bool_A bool; 			// The default value of the bool type in golang is false.
	var bool_B = true;			// Golang can automatically detect and set the data type of the parameter. (No recommend for myself)
	var bool_C bool = false;	// Manually define parameter data types, show an error when incorrect value is given.

	var BoolFormatOutput = "%s: %5v, %s\n";
	fmt.Println("Bool");
	fmt.Println("=======================");
	fmt.Printf(BoolFormatOutput, "A", bool_A, typeof(bool_A));
	fmt.Printf(BoolFormatOutput, "B", bool_B, typeof(bool_B));
	fmt.Printf(BoolFormatOutput, "C", bool_C, typeof(bool_C));

	// 2. Numbers
	// 2.1 Integer
	var intSigned_00 int  ; // The actual number of bits for the int type depends on the size of the operating system.
	var intSigned_08 int8 ; //
	var intSigned_16 int16; //
	var intSigned_32 int32; //
	var intSigned_64 int64; //

	var intSignedOutput = "%-5s: %d\n";
	fmt.Println("\nSigned int");
	fmt.Println("=======================");
	fmt.Printf(intSignedOutput, typeof(intSigned_00), intSigned_00);
	fmt.Printf(intSignedOutput, typeof(intSigned_08), intSigned_08);
	fmt.Printf(intSignedOutput, typeof(intSigned_16), intSigned_16);
	fmt.Printf(intSignedOutput, typeof(intSigned_32), intSigned_32);
	fmt.Printf(intSignedOutput, typeof(intSigned_64), intSigned_64);
	// 2.2 Double

}