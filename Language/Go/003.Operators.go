
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
	fmt.Println();

	// 2. Relational Operator
	// Basically, relational operators are used in logical decisions and return a Boolean valu, true or false.
	fmt.Printf("Symbol \"==\" -- The \"Haha\" is     Equal to \"Haha\": %v.\n", "Haha" == "Haha");	// ==
	fmt.Printf("Symbol \"!=\" -- The \"Haha\" is Not Equal to \"haha\": %v.\n", "Haha" != "haha");	// !=
	fmt.Println();

	// > AND >=
	fmt.Printf("Symbol \"> \" -- 10 >  10 is %v\n", 10 >  10);
	fmt.Printf("Symbol \">=\" -- 10 >= 10 is %v\n", 10 >= 10);
	fmt.Println();

	// < AND <=
	fmt.Printf("Symbol \"< \" -- 10 <  20 is %v\n", 10 <  20);
	fmt.Printf("Symbol \"<=\" -- 10 <= 20 is %v\n", 10 <= 20);
	fmt.Println();

	
}