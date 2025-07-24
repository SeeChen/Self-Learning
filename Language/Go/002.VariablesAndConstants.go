
package main

import (
	"fmt"
)

func main() {
	
	// Variables can be defined using the "var identifier type".
	var SeeChen string = "SeeChen";

	// It is also possible to define multiple variables of the same type.
	var See, Chen string = "See", "Chen";

	fmt.Printf("My name is %s.  \n", SeeChen  );
	fmt.Printf("My name is %s%s.\n", See, Chen);
}