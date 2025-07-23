package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// A typeof function
// Returns a string describing the variable type of the passed parameter.
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

type Numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | 
		uint | uint8 | uint16 | uint32 | uint64 ;
}
func formatStringNumbers[T Numbers](parameter T) string {

	var str string = "%-7s: %d, Byte: %d.";
	return fmt.Sprintf(str, typeof(parameter), parameter, unsafe.Sizeof(parameter));
}

func main() {
	
	// The basic data type of Golang
	
	// 1. Bool
	// The bool type has only two types of data: true and false.
	var bool_A bool        ; // The default value of the bool type in golang is false.
	var bool_B      = true ; // Golang can automatically detect and set the data type of the parameter. (No recommend for myself)
	var bool_C bool = false; // Manually define parameter data types, show an error when incorrect value is given.

	var BoolFormatOutput = "%s: %5v, %s\n";
	fmt.Println("\n=======================");
	fmt.Println("Bool");
	fmt.Println("=======================");
	fmt.Printf(BoolFormatOutput, "A", bool_A, typeof(bool_A));
	fmt.Printf(BoolFormatOutput, "B", bool_B, typeof(bool_B));
	fmt.Printf(BoolFormatOutput, "C", bool_C, typeof(bool_C));

	// 2. Numbers
	// 2.1 Integer
	var intSigned_00 int  ; // The actual number of bits for the int type depends on the size of the operating system.
	var intSigned_08 int8 ;
	var intSigned_16 int16;
	var intSigned_32 int32;
	var intSigned_64 int64;

	fmt.Println("\n=======================");
	fmt.Println("Signed int");
	fmt.Println("=======================");
	fmt.Println(formatStringNumbers(intSigned_00));
	fmt.Println(formatStringNumbers(intSigned_08));
	fmt.Println(formatStringNumbers(intSigned_16));
	fmt.Println(formatStringNumbers(intSigned_32));
	fmt.Println(formatStringNumbers(intSigned_64));

	var intUnsigned_00 uint  ; // The unsigned int type in golang has the same purpose and definition as signed int, but has not sign.
	var intUnsigned_08 uint8 ;
	var intUnsigned_16 uint16;
	var intUnsigned_32 uint32;
	var intUnsigned_64 uint64;

	fmt.Println("\n=======================");
	fmt.Println("Unsigned int");
	fmt.Println("=======================");
	fmt.Println(formatStringNumbers(intUnsigned_00));
	fmt.Println(formatStringNumbers(intUnsigned_08));
	fmt.Println(formatStringNumbers(intUnsigned_16));
	fmt.Println(formatStringNumbers(intUnsigned_32));
	fmt.Println(formatStringNumbers(intUnsigned_64));

	var byte_uint8 byte; // The byte data type is an alias for uint8 in golang.
	var rune_int32 rune; // The rune data type is an alias for int32 in golang.

	fmt.Println("\n=======================");
	fmt.Println("byte and rune");
	fmt.Println("=======================");
	fmt.Printf("byte: %s.\n", typeof(byte_uint8));
	fmt.Printf("rune: %s.\n", typeof(rune_int32));

	// 2.2 Double
}