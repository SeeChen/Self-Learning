package main

import "fmt"

// My First Golang Program
func main() {
	fmt.Println("Hello World!");

	// Format String
	var Code = 1;
	var Date = "2025-07-22";
	var Time = "08:00 GMT+08:00";
	var StrFormat = "Number: %d, Date: %s, Time: %s.";

	// For this function, the formatted string is assigned to a variable.
	var StrOutput = fmt.Sprintf(StrFormat, Code, Date, Time);
	fmt.Println(StrOutput);

	// For this fucntion, the formatted string will be output in standard output format.
	fmt.Printf(StrFormat, Code, Date, Time);
	fmt.Println();
}