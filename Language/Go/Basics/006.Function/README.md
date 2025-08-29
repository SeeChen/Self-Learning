<div align=center>

# Chapter 6: Functions in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **functions in Go**, including:
- Function declaration and usage
- Single return and multiple return values
- Pass-by-value vs pass-by-reference
- Anonymous functions and closures
- Callback functions using function types as parameters

It demonstrates:
- How to define functions with parameters and return types
- Using multiple return values to return sorted results
- Passing variables by value vs by address (pointer)
- Creating anonymous functions inside `main()`
- Implementing closures that capture external variables
- Using callbacks for custom operations

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
  - `fmt`: for formatted printing
- **Features Covered**:
  - Function declaration with and without return values
  - Multiple return values
  - Pass-by-value vs pass-by-reference
  - Anonymous functions
  - Closures
  - Callback functions

## 3. Key Concepts
- **Function Syntax**:
    ```go
    func functionName(parameters) returnType {
        // function body
    }
    ```
- **Multiple Return Values**:
    - Functions can return more than one value:
    ```go
    func sort2num(x, y int) (int, int) {
        if x > y {
            return x, y
        }
        return y, x
    }
    ```
- **Pass-by-Value vs Pass-by-Reference:**
    - **Value**: changes inside the function do not affect original variables
    - **Reference**: use pointers (`*int`) to modify original values
- **Anonymous Functions**:
    - Functions without names, created at runtime:
    ```go
    myFunc := func(x int) string {
        return fmt.Sprintf("%d", x)
    }
    ```
    **Closures**:
        - Functions that capture and use variables from outer scope
    **Callback Functions**:
        - Functions passed as arguments to other functions, useful for custom logic

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
Hello World!
The max number between 2 and 3 is 3.
Sorted (21, 20) = (21, 20).

[valuePass func] x = 1000, y = 1000.
[main func] x = 100, y = 101.

[addressPass func] x = 1000, y = 1000.
[main func] x = 1000, y = 1000.

1.
x: 10, y: 20, x + y: 30.
x: 10, y: 12, x + y: 22.
x: 20, y: 40, x + y: 60.
callbackFunc: (10 + 20) * 1.4142135624 = 42.426407
```

## 5. Practice
1. Write a function that calculates the factorial of a number using recursion.
2. Implement a function that swaps two integers using pointers.
3. Create a closure that keeps track of how many times it has been called.
4. Write a function that accepts another function as a parameter and applies it to a list of numbers.
5. Modify `callbackFunc` to accept two callbacks and choose one based on a condition.

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Functions](https://gobyexample.com/functions)
3. [Go by Example - Multiple Return Values](https://gobyexample.com/multiple-return-values)
4. [Go by Example - Closures](https://gobyexample.com/closures)
5. [Spec: Function declarations](https://go.dev/ref/spec#Function_declarations)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 29-AUG-2025 22:32 UTC +08:00*
</div>