
package main

import (
	"fmt"
	"strings"
)

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
	
	// 2. if...else
	// 3. switch
	// 4. select
}