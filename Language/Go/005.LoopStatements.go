
package main

import (
	"fmt"
)

func main() {
	// Loop statements in Golnag are different from other programming language
	// Golang loop statements have only one keyword -- "for".

	// 1. for init; condition; post {}
	// The logic is the same as other programming language, just the structure is different.
	var sum int = 0;
	for i := 0; i < 10; i++ {
		sum += 1;
	}
	fmt.Println(sum);

	// 2. for condition {}
	// This statement is similar to the "while" keyword in C++ language.
	for sum < 20 {
		sum += 1;
	}
	fmt.Println(sum);
}