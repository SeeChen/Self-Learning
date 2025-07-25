
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

	fmt.Println();
	// #4
	// Multiple Assignments
	// Similar to Python, Golang also accepts multiple assignments to identifiers.
	// Type 1
	var var1, var2, var3 int;
	var var4, var5, var6 string = "Var 4", "Var 5", "Var 6";
	fmt.Printf("%d, %d, %d\n", var1, var2, var3);
	fmt.Printf("%s, %s, %s\n", var4, var5, var6);

	// Type 2
	var var7, var8, var9 = 1, "My String", true;
	fmt.Printf("%v, %v, %v\n", var7, var8, var9);

	// Type 3
	var10, var11, var12 := "This", false, 9;
	fmt.Printf("%v, %v, %v\n", var10, var11, var12);

	// Type 4
	// In generally, the factorization keyword is used to declare global variables.
	var (
		var13 int
		var14 string
	);
	var13 = 0;
	var14 = "String";
	fmt.Printf("%d, %s\n", var13, var14);
	fmt.Println();

	// Constants, keyword: const
	// Basically, constants in golang are similar to variables
	// Whether it is the declaration or use of identifiers.
	const const1 int = 0;
	const const2, const3 string = "Const 2", "Const 3";

	fmt.Printf("%d: %s, %s.\n", const1, const2, const3);

	// It is worth mentioning that the symbol := is not supported in statements defined as const.
	// In general, can use const () to define identifiers as enumerations.
	const (
		MALE   int = 0
		FEMALE int = 1
		OTHER  int = 2
	);

	fmt.Printf("Male   is %d\n", MALE  );
	fmt.Printf("Female is %d\n", FEMALE);
	fmt.Printf("Other  is %d\n", OTHER );
	fmt.Println();

	// A special identifier in Golang -- iota
	// It can be assumed that "iota" can be modified by the compiler
	const (
		iota1 int = iota
		iota2 int = iota
	)
	fmt.Println("iota 1: ", iota1);
	fmt.Println("iota 2: ", iota2);
	fmt.Println();

	// Or short-form
	const (
		iota3 = iota
		iota4
		iota5
	)
	fmt.Println("iota 1: ", iota3);
	fmt.Println("iota 2: ", iota4);
	fmt.Println("iota 3: ", iota5);
}