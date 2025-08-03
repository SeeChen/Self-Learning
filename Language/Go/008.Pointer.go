
package main

import (
	"fmt"
)

func main() {
	// The pointer points to the address of a variable.
	// In Golang, use the "&" symbol before a parameter to get th address of a variable.
	var a int = 10;
	fmt.Printf("The address of a: %x.\n", &a);
	fmt.Println();

	// To declare a variable as a pointer, must use the var varName *type format.
	var pointerInt *int = &a;
	fmt.Println(pointerInt);

	// Use of pointer.
	var b int = 10;
	pointerInt = &b;

	fmt.Printf("%-10s: %x.\n", "&b"        , &b        );
	fmt.Printf("%+10s: %x.\n", "pointerInt", pointerInt);

	// The following usage is the same operation, assigning values to the same variables.
	b = 20; 
	fmt.Printf("b: %d, *pointerInt: %d.\n", b, *pointerInt);
	*pointerInt = 30; 
	fmt.Printf("b: %d, *pointerInt: %d.\n", b, *pointerInt);

	// Null Pointer
	// When a pointer is declared but not assigned a value, it is a null pointer.
	// In Golang, the null pointer has only one value -- nil.
	// For example:
	var p *int;
	fmt.Println(p);

	// Check the null pointer.
	if p == nil {
		fmt.Println("P is a null Porinter.");
	}
}
