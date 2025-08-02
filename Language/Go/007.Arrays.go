
package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	// In Golang, arrays can be defined using var arrayName [size]dataType;
	var arrayInt [10]int;

	// The default value of the "int" type in the array
	// Golang will set the value of all elements to 0.
	fmt.Println(arrayInt);

	// But can use the initial element value through var arrayName [size] dataType = {...}
	var arrayString [5]string = [5]string {
		"SeeChen",
		"Lee",
		"---",
		"李",
		"思净",
	};
	fmt.Println(arrayString);

	// Alternatively, can use a specific index to initialize the value at a different index position.
	specificIndex := [5]int {2: 10, 4: 100}
	fmt.Println(specificIndex);

	// The example above, array size needs to confirmed.
	// But in real-world development, this is ofter not possible.
	// Therefore, Golang allows us to use [...] to initialize an array of indeterminate size.
	autoDetermineSize := [...]int {1, 2, 3, 4, 5}
	fmt.Println(autoDetermineSize);
	fmt.Println();

	// Below shows all operations on arrays in Golang.
	var arr [15]int;
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000);
	}

	fmt.Printf("Full Arrays: %v", arr);
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%2d]: %4d.\n", i, arr[i]);
	}

	// 
}