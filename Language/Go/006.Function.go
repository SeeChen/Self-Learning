
package main

import (
	"fmt"
)

// Following is the function format in Golang
// func functionName ([parameter list]) [return type]
// for example:
func max(x, y int) int {
	if x > y {
		return x;
	} 
	return y;
}

// This is an example that returns two values
func sort2num(x, y int) (int, int) {
	if x > y {
		return x, y;
	}

	return y, x;
}

// In this function, changes in numbers will not affect the actual parameters.
func valuePass(x, y int) {
	x, y = 1000, 1000;

	fmt.Printf("[valuePass func] x = %d, y = %d.\n", x, y);
}
// In this function, changes in numbers will affect to actual parameters.
func addressPass(x, y *int) {
	*x, *y = 1000, 1000;
	fmt.Printf("[addressPass func] x = %d, y = %d.\n", *x, *y);
}

// closureFunc returns a closure that captures variable x from the outer scope.
func closureFunc(x int) (func(y int) (int, int, int)) {

	return func(y int) (int, int, int) {
		return x, y, x + y;
	}
}

// This is a Callback function example.
func callbackFunc(x, y int, callback func(int) (float64)) (float64) {
	var z int = x + y;
	return callback(z);
}

func main() {
	// Actually, the main() is a function
	fmt.Println("Hello World!");

	// We defined two numbers to call the function max()
	// This is the basic usage of functions in Golang.
	var x, y int = 2, 3;
	fmt.Printf("The max number between %d and %d is %d.\n", x, y, max(x, y));

	// Also, we can return multiple values.
	x, y = 21, 20;
	var z1, z2 int = sort2num(x, y);
	fmt.Printf("Sorted (%d, %d) = (%d, %d).\n\n", x, y, z1, z2);

	// In general, we need to use the following methoeds to operate some variables.
	// But any changes in the function will not effect actual parameters
	x, y = 100, 101;
	valuePass(x, y);
	fmt.Printf("[main func] x = %d, y = %d.\n\n", x, y);

	// Special case, because the parameter is passed by address.
	x, y = 100, 101;
	addressPass(&x, &y);
	fmt.Printf("[main func] x = %d, y = %d.\n\n", x, y);

	// And, function can be created during code execution.
	myFunc := func (x int) string {

		return fmt.Sprintf("%d", x);
	}
	fmt.Printf("%s.\n", myFunc(1));

	// Closure
	var closureX, closureY, closureZ int;
	closureFunc_10 := closureFunc(10);
	closureFunc_20 := closureFunc(20);

	// x = 10, y = 20
	closureX, closureY, closureZ = closureFunc_10(20);
	fmt.Printf("x: %d, y: %d, x + y: %d.\n", closureX, closureY, closureZ);
	// x = 10, y = 12
	closureX, closureY, closureZ = closureFunc_10(12);
	fmt.Printf("x: %d, y: %d, x + y: %d.\n", closureX, closureY, closureZ);
	// x = 20, y = 40
	closureX, closureY, closureZ = closureFunc_20(40);
	fmt.Printf("x: %d, y: %d, x + y: %d.\n", closureX, closureY, closureZ);

	// Callback
	x, y = 10, 20;
	var a float64 = callbackFunc(x, y, func(z int) float64 {
		return float64(z) * 1.4142135624;
	});
	fmt.Printf("callbackFunc: (%d + %d) * 1.4142135624 = %f\n", x, y, a);
}