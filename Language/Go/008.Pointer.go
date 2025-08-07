
package main

import (
	"fmt"
)

func modifyArray(arrPtr *[10]int) {
	// Modify an element via pointer to an array
	(*arrPtr)[5] = 100;
}

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

	// Array of pointers
	// In Golang, to define a pointer array, use var ptr [size]* type
	var arr = [10]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10};	// Declare an array
	var arrPtr1 [10]*int;

	// Assign each elemet's address to the pointer array
	for i := 0; i < len(arr); i++ {
		arrPtr1[i] = &arr[i];
	}
	fmt.Println(arrPtr1);
	for i := 0; i< len(arr); i++ {
		fmt.Printf("%d, ", *arrPtr1[i]);
	}
	fmt.Println("\n");

	// Pointer to an array
	// Use var ptr *[size] type declaration
	arr2 := [10]int {2, 4, 6, 8, 10, 1, 3, 5, 7, 9};
	fmt.Println(arr2);

	// Defined a pointer arrays
	var ptrArr2 *[10]int;
	ptrArr2 = &arr2;
	fmt.Println(ptrArr2);

	// Access array elements through the pointer
	fmt.Println((*ptrArr2)[0]);	// Using explicit dereference
	fmt.Println(ptrArr2[0]); 	// Shorthand (Go allow this)
	fmt.Println(*ptrArr2);		// Get the entire array

	// Use pointer to array in function
	// Can modify original array without return value.
	modifyArray(ptrArr2);
	fmt.Println(arr2);
}
