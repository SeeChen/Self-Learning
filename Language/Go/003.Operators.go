
package main

import (
	"fmt"
)

func binaryBitOperator(p byte, q byte) {
	fmt.Printf("%b & %b = %b, %b | %b = %b, %b ^ %b = %b.\n", p, q, p & q, p, q, p | q, p, q, p ^ q);
}

func complicatedBitOperator(i int, j int, o string, x int) {
	fmt.Printf("%d %s %d = %d\n", i, o, j, x);
	fmt.Printf("%4d (%012b)\n", i, i);
	fmt.Printf("%4d (%012b)\n", j, j);
	fmt.Printf("%4d (%012b)\n", x, x);

	fmt.Println();
}

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

	// 3. Logical Operators
	// Logical operators are used to perform "AND" and "OR" logical calculations.
	fmt.Println("Symbol \"&&\"");
	fmt.Printf("true  && true  = %v.\n", true  && true );
	fmt.Printf("true  && false = %v.\n", true  && false);
	fmt.Printf("false && true  = %v.\n", false && true );
	fmt.Printf("false && false = %v.\n", false && false);

	fmt.Println();
	fmt.Println("Symbol \"||\"");
	fmt.Printf("true  || true  = %v.\n", true  || true );
	fmt.Printf("true  || false = %v.\n", true  || false);
	fmt.Printf("false || true  = %v.\n", false || true );
	fmt.Printf("false || false = %v.\n", false || false);

	fmt.Println();
	fmt.Println("Symbol \"!\"");
	fmt.Printf("!true  = %v.\n", !true );
	fmt.Printf("!false = %v.\n", !false);

	// 4. Bitwise Operators
	// Bitwise operators allow us to manipulate binary bits in memory.

	fmt.Println();
	fmt.Println("Simple example: ");
	var p, q byte = 0, 0; binaryBitOperator(p, q);
	p, q = 0, 1; binaryBitOperator(p, q);
	p, q = 1, 0; binaryBitOperator(p, q);
	p, q = 1, 1; binaryBitOperator(p, q);
	fmt.Println();

	// This is a complicated method to use
	var i, j int = 121, 10; 
	complicatedBitOperator(i, j, "&", i & j);
	complicatedBitOperator(i, j, "|", i | j);
	complicatedBitOperator(i, j, "^", i ^ j);
}