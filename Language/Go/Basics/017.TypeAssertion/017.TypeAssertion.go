package main

import (
	"fmt"
	"reflect"
)

// withoutOk demonstrates a type assertion without the `ok` idiom.
// If the assertion fails, it will panic. We use defer + recover to handle it gracefully.
func withoutOk(value interface{}) (result string) {
	defer func() {
		if r := recover(); r != nil {
			result = fmt.Sprintf("[Recover] Type assertion failed: %v", r)
		}
	}()

	v := value.(int) // Will panic if not int
	result = fmt.Sprintf("Assertion success: original=%v, asserted(int)=%d", value, v)
	return
}

// switchMethod demonstrates type switch, which is safer and more idiomatic than multiple assertions.
func switchMethod(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("[int] %d + 1 = %d\n", v, v+1)
	case string:
		fmt.Printf("[string] %q + \"1\" = %q\n", v, v+"1")
	case bool:
		fmt.Printf("[bool] value = %v\n", v)
	default:
		fmt.Printf("[Unknown type] value = %v (type = %T)\n", v, v)
	}
}

// ifElseMethod uses the `comma-ok` idiom for multiple types.
func ifElseMethod(value interface{}) {
	if str, ok := value.(string); ok {
		fmt.Printf("[if-else] String: %q\n", str)
	} else if i, ok := value.(int); ok {
		fmt.Printf("[if-else] Integer: %d\n", i)
	} else {
		fmt.Printf("[if-else] Unknown type: %v (type = %T)\n", value, value)
	}
}

// reflectExample shows how reflection can inspect type information dynamically.
// Unlike type assertion, reflect does not panic and supports generic inspection.
func reflectExample(value interface{}) {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	fmt.Printf("[reflect] Type = %s, Kind = %s, Value = %v\n", t.Name(), t.Kind(), v)
}

// customErrorExample demonstrates asserting interface to a specific custom type.
// This is common when working with error interfaces.
type customError struct {
	msg string
}

func (e customError) Error() string {
	return e.msg
}

func customErrorExample(err error) {
	if ce, ok := err.(customError); ok {
		fmt.Printf("[Custom Error] %s\n", ce.msg)
	} else {
		fmt.Printf("[Custom Error] Not a customError, type=%T\n", err)
	}
}

func main() {
	// === 1. Basic type assertion ===
	var i interface{} = "Hello World!"
	str, ok := i.(string)
	fmt.Printf("Assertion to string ok=%v, value=%q\n", ok, str)

	integer, ok := i.(int)
	fmt.Printf("Assertion to int ok=%v, value=%v\n\n", ok, integer)

	// === 2. Assertion without ok (may panic) ===
	fmt.Println(withoutOk("Hello")) // Will recover panic
	fmt.Println(withoutOk(12345))
	fmt.Println()

	// === 3. Type switch ===
	switchMethod("This is a string")
	switchMethod(114514)
	switchMethod(true)
	switchMethod(3.14) // Unknown type
	fmt.Println()

	// === 4. if-else method ===
	ifElseMethod("Example string")
	ifElseMethod(999)
	ifElseMethod([]int{1, 2, 3})
	fmt.Println()

	// === 5. Reflection ===
	reflectExample("Hello")
	reflectExample(42)
	reflectExample(false)
	fmt.Println()

	// === 6. Custom error assertion ===
	var err error = customError{"Something went wrong"}
	customErrorExample(err)
	customErrorExample(fmt.Errorf("standard error"))
}
