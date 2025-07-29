
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
	// 4. select
}