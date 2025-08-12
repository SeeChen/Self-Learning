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
		fmt.Printf("No %2d. %s; ", i, value)
	}
	fmt.Println()
	// Print the slice length
	fmt.Printf("The lenght of slice is %d\n", len(slice1))
	// And append an item to this slice
	slice1 = append(slice1, "Append")
	fmt.Println(slice1)
	fmt.Printf("The lenght of slice is %d\n", len(slice1))

	// Make function with capacity
	slice2 := make([]string, 10, 11)
	for i, _ := range slice2 {
		slice2[i] = fmt.Sprintf("Index_%d", i)
	}
	// Append one item to the slice, but that's to be expected.
	slice2 = append(slice2, "FirstAppend")
	// Print current slice size.
	fmt.Printf("Current slice2 size: %d\n", len(slice2))
	// Append again, and print again
	slice2 = append(slice2, "SecondAppend")
	fmt.Printf("Current slice2 size: %d\n", len(slice2))

	// len() function and cap() function
	// The `len` function can get the size of the current slice, which has been demonstrated above.
	// And `cap` function can get the capacity of the slice, which can calculate the maximum lenght of the slice.
	fmt.Printf("slice2 length: %d; capacity: %d.\n", len(slice2), cap(slice2))

	// Empty slice
	// The slice is an empty(nil) slice when the slice is only declare but not initialized.
	var emptySlice []int
	fmt.Printf("emptySlice is a nil slice -> %v.\n", nil == emptySlice)
	fmt.Println()

	// Slice Extraction
	// In Golang, can use [lower-bound: upper-bound] to extract a portion of a slice.
	fmt.Printf("Full slice: %v\n", slice2)
	fmt.Printf("Get the slice item between index 2 and index 5: %v\n", slice2[2:5])
	// From the output, we can see that the lower-bound is included and the upper-bound is out of the range.
	// Therefore, if need to extract from the items with index 0
	// Can use the [: upper-bound] format
	fmt.Printf("Start from 0: %v\n", slice2[:4])
	fmt.Printf("Start from 0: %v\n", slice2[0:4]) // Of course can use this

	// If want to extract the items of a slice up to the last item
	// Can use the [lower-bound:] format
	fmt.Printf("End in last item: %v.\n", slice2[3:])
	fmt.Printf("End in last item: %v.\n", slice2[3:len(slice2)]) // Actually, it's the same as this
}
