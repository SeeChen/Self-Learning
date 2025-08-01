
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

	// 3. for-each range
	// The for-each statement allows us to loop over strings, arrays, maps, etc.
	strings := []string{"Lee", "See", "Chen"};
	for i, s := range strings {
		fmt.Printf("%d. %s\n", i, s);
	}

	// Alternatively, we can skip value to loop over.
	// There are 3 ways to loop over a maps variables.
	var myName map[string]string = map[string]string {
		"FirstName" : "SeeChen",
		"FamilyName": "Lee",
	};

	// Loop over the map with key and value
	for key, value := range myName {
		fmt.Printf("My %s is %s. ", key, value);
	}
	fmt.Println();

	// Loop over without the value
	for key := range myName {
		fmt.Printf("My %s is %s. ", key, myName[key]);
	}
	fmt.Println();

	// Loop over without the key.
	// This method is similar to the key, value method.
	// However, if use "_" in Golang, will get an error -- "cannot use _ as a value or type".
	for _, value := range myName {
		fmt.Printf("%s ", value);
	}
	fmt.Println("\n");

	// 4. Nested Loop
	// Loop within loop
	// We can use loop statements inside loops and nest any of the above statements.
	// An example: 9 x 9 Multiplication Table
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9};
	for _, m := range numbers {
		for n := 1; n <= m; n++ {
			fmt.Printf("%d x %d = %d, ", m, n, m * n);
		}

		fmt.Println();
	}
	fmt.Println();

	// Loop Control Statements -- Break
	// In Golang, there are three ways to use break keyword.
	// 1. In for-loop statement
	var break_1 int = 0;
	for ; ; break_1++ {
		// Normally, this code will cause the codeblock to loop infinitely.
		fmt.Printf("%d ", break_1);

		// But add a condition to break out of this statement.
		if break_1 == 10 {
			fmt.Println("\n");
			break;
		}
	}

	// Or break out of the sub-loop
	break_1 = 0;
	for _, m := range numbers {
		// Using the 9x9 Multiples Table as an example.
		// In the above case, we know that all results will be displayed.

		for n := 1; n <= m; n++ {

			// But add a condition, if the result is greater than 50, then break.
			fmt.Printf("%d x %d = %d, ", m, n, m * n);
			if (m * n > 50) {
				break;
			}
		}
		fmt.Println();
	}

	fmt.Println("");
}