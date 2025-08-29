package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	// In Golang, arrays can be defined using var arrayName [size]dataType;
	var arrayInt [10]int

	// The default value of the "int" type in the array
	// Golang will set the value of all elements to 0.
	fmt.Println(arrayInt)

	// But can use the initial element value through var arrayName [size] dataType = {...}
	var arrayString [5]string = [5]string{
		"SeeChen",
		"Lee",
		"---",
		"李",
		"思净",
	}
	fmt.Println(arrayString)

	// Alternatively, can use a specific index to initialize the value at a different index position.
	specificIndex := [5]int{2: 10, 4: 100}
	fmt.Println(specificIndex)

	// In the example above, the array size must be specified.
	// But in real-world development, this is often not practical.
	// Therefore, Golang allows us to use [...] to initialize an array of indeterminate size.
	autoDetermineSize := [...]int{1, 2, 3, 4, 5}
	fmt.Println(autoDetermineSize)
	fmt.Println()

	// Below shows all operations on arrays in Golang.
	var arr [15]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}

	fmt.Printf("Full Arrays: %v.\n", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%2d]: %4d.\n", i, arr[i])
	}

	// Multidimensional arrays
	// All operations on multidimensional arrays in Golang are the same as those on one-dimensional arrays.
	// The only different is the definition.
	// To define a multidemensional arrays, use arraysName [SIZE1][SIZE2]...[SIZEN]arraysType
	// Define a 2D array.
	var arr2D [5][5]int
	fmt.Println(arr2D)

	// 3D array simple operation example
	arr3D := [][][]int{}
	for i := 0; i < 10; i++ {

		arr2D := [][]int{}
		for j := 0; j < 10; j++ {

			if (j % 4) == 0 {
				continue
			}

			arr1D := []int{}
			for k := 0; k < 10; k++ {
				if (k % 3) == 0 {
					continue
				}
				arr1D = append(arr1D, rand.Intn(1000))
			}
			arr2D = append(arr2D, arr1D)
		}
		arr3D = append(arr3D, arr2D)
	}
	for i := 0; i < len(arr3D); i++ {
		for j := 0; j < len(arr3D[i]); j++ {
			for k, v := range arr3D[i][j] {
				fmt.Printf("arr[%d][%d][%d]: %3d, ", i, j, k, v)
			}
			fmt.Println()
		}
		fmt.Println()
	}

	// In addition, the size of an array in Golang is part of its type.
	// Therefore, arrays of different sizes in Golang are not of the same type.
	// Let's print the type of the arrays
	fmt.Printf("Type of arr3D (slice-based 3D array): %s\n", reflect.TypeOf(arr3D).String())
	fmt.Printf("Type of arr   (fixed-size array    ): %s\n", reflect.TypeOf(arr).String())
}
