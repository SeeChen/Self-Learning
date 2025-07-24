
package main

import (
	"fmt"
	"reflect"
)

func main() {
	
	// #1
	// Variables can be defined using the "var identifier type".
	var SeeChen string = "SeeChen";

	// It is also possible to define multiple variables of the same type.
	var See, Chen string = "See", "Chen";

	fmt.Printf("My name is %s.  \n", SeeChen  );
	fmt.Printf("My name is %s%s.\n", See, Chen);

	// #2
	// Not only can define variables with a specified type
	// Can also let golang automatically determine the type of the variable based on the value of the variable.
	var varBool = true     ;
	var varStr  = "SeeChen";
	var varInt  = 2000     ;

	fmt.Println();
	fmt.Println(reflect.TypeOf(varBool).String());
	fmt.Println(reflect.TypeOf(varStr ).String());
	fmt.Println(reflect.TypeOf(varInt ).String());
	// When using the above method, must assign a value to the variable.
	// If only use "var identifier", an error will be reported during compilation.

	// #3
	// := symbol.
	// When assigning a value to an identifier, be sure not to use "var identifier"
	// Otherwise an error will be reported.
	intValue := 1;
	strValue := "String";

	fmt.Println();
	fmt.Printf("%d, %s.\n", intValue, strValue);

	// #4
	// Multiple Assignments
	
}