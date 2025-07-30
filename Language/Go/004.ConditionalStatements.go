
package main

import (
	"fmt"
	"strings"
)

func compareTwoNum(i int, j int) string {
	// This is an example to show how the if...else statement works.
	var strFormat string = "%03d [i] %s %03d [j].";
	var op string;
	if i > j {
		op = ">";
	} else if i < j {
		op = "<";
	} else {
		op = "=";
	}

	return fmt.Sprintf(strFormat, i, op, j);
}

func grade(score int) string {

	var grade string;
	switch {
		case score >= 90:
			grade = "A";
		case score >= 70:
			grade = "B";
		case score >= 60:
			grade = "C";
		case score >= 40:
			grade = "D";
		case score >= 20:
			grade = "E";
		default:
			grade = "F";
	}

	return grade;
}

func main() {
	// 1. if
	// The if conditional statement in Golang does not require brackets
	if true {
		fmt.Println("True");
	}

	if (!false) {
		fmt.Println("Not false");
	}

	var str string = "This is a string";
	if strings.Contains(str, "str") {
		fmt.Printf("\"%s\" is contains \"str\".\n", str);
	}

	// Or use the format below.
	if a := 20; a > 10 {
		fmt.Println(a);
	}
	
	fmt.Println();
	// 2. if...else
	// if...else allows us to handle conditional logic that evaluates to false.
	fmt.Println(compareTwoNum(200, 100));
	fmt.Println(compareTwoNum(100, 200));
	fmt.Println(compareTwoNum(150, 150));
	fmt.Println();

	// 3. switch
	// The switch statement allows us to handle multiple checks.
	// Golang will execute the logic code for the successful match.
	// Unlike other programming languages, Golang does not require adding a break keyword within sub-codes to prevent further execution of the logic for ohter matching conditions.
	// A simple code with switch logic demonstration.
	fmt.Printf("Score:  65, Grade: %s.\n", grade( 65));
	fmt.Printf("Score: -20, Grade: %s.\n", grade(-20));
	fmt.Printf("Score:  77, Grade: %s.\n", grade( 77));
	fmt.Printf("Score:  94, Grade: %s.\n", grade( 94));
	fmt.Println();

	// Another approach is to use a switch
	// Use type-switch to determine the true type of the interface in the memory.
	var value interface{};
	switch value.(type) {
		case nil:
			fmt.Println("nil");
		default:
			fmt.Println("Default");
	}
	fmt.Println();

	// "fallthrough" allows us to force execution of the next case statement.
	// Explain this with two examples.
	var fall int = 10;

	fmt.Println("Without fallthrough");
	switch fall {
		case 10:
			fmt.Println("10");
		case 20:
			fmt.Println("20");
		default:
			fmt.Println("No Number")
	}

	fmt.Println();
	fmt.Println("With fallthrough");
	switch fall {
		case 10:
			fmt.Println("10");
			fallthrough;
		case 20:
			fmt.Println("20");
			fallthrough;
		default:
			fmt.Println("No Number")
	}

	// 4. select
}