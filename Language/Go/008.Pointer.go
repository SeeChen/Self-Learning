
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

	// The array pointer
	// In Golang, to define a pointer array, use var ptr [size]* type
	var arr = [10]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10};	// Declare an array
	var arrPtr1 [10]*int;

	// Use a loop to assign each index
	for i := 0; i < len(arr); i++ {
		arrPtr1[i] = &arr[i];
	}
	fmt.Println(arrPtr1);
	for i := 0; i< len(arr); i++ {
		fmt.Printf("%d, ", *arrPtr1[i]);
	}
	fmt.Println();
}
