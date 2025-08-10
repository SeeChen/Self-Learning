package main

import "fmt"

func main() {
	// A Slice is actually similar to an array in its usage.
	// However, slices can be resized, and arrays of unknown size can be created.
	// Initialize a slice with no size.
	var slice0 []int

	// Use it
	slice0 = []int{1, 2, 3}
	fmt.Println(slice0)

	// Alternatively, in Golang, you can use the make() function to dynamically create arrays (slices).
	// The syntax for the make function is `make([]type, lenght, capacity)`.
	// The capacity parameter is optional.
	// The capacity is the total size of the underlying array referenced by the slice.
	// The capacity must be greater than or equal to the length. For example, make([]int, 10, 9) is not allowed.
	// If the capacity is omitted, the length is used by default. The capacity determines how many elements the slice can hold before the underlying array needs to be reallocated when the append operator is used.
	// Make function without capacity
	slice1 := make([]string, 10)
	slice1[0] = "Zero"
	slice1[1] = "One"
	slice1[5] = "String"
	for i, value := range slice1 {
		fmt.Printf("No %2d. %s\n", i, value)
	}
}
