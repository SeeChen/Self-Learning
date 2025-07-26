
package main

import (
	"fmt"
)

func main() {
	
	// 1. Arithmetic Operators
	// Following are all the arithmetic operators in Golang.
	fmt.Printf("10 + 20 = %d.\n", 10   + 20  );	// Addition
	fmt.Printf("10 - 20 = %d.\n", 10   - 20  ); // Subtraction
	fmt.Printf("10 * 20 = %d.\n", 10   * 20  ); // Multiplication
	fmt.Printf("10 / 20 = %f.\n", 10.0 / 20.0); // Division

	fmt.Println()
	fmt.Printf("11 %% 10 = %d.\n", (11 % 10)); // Find the remainder

	fmt.Println();

	var a int;
	a = 20; a++;
	fmt.Printf("a++ = %d.\n", a);
	a = 15; a--;
	fmt.Printf("a-- = %d.\n", a);
}