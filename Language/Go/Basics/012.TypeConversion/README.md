<div align=center>

# Chapter 12: Type Conversion in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **type conversion in Go**, which allows changing a variable from one type to another. It covers:
    - **Numeric type conversions** using `type(expression)`
    - **String conversions** using `strconv` package
    - **Interface conversions** including **type assertion** and **type conversion**

It demonstrates:
    - Converting integers to floats
    - Parsing strings to numbers and vice versa
    - Handling conversion errors
    - Converting between interfaces using type assertion and type conversion

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
    - `fmt`: for formatted output
    - `strconv`: for string-number conversions
    - `reflect`: for type inspection
- **Features Covered**:
    - Numeric type conversion
    - String ↔ Number conversion
    - Interface type conversion (`type assertion`, `type conversion`)

## 3. Key Concepts
- **Numeric Type Conversion**:
    - Syntax: `float64(a)`
    - Example:  
        ```go
        var x int = 10
        var y float64 = float64(x)
        ```
- **String ↔ Number Conversion**:
    - `strconv.Atoi(string)` → `int`
    - `strconv.ParseFloat(string, bitSize)` → `float64`
    - `strconv.Itoa(int)` → `string`
- **Error Handling in Conversion**:
    - Most `strconv` functions return `(value, error)`
    - Use `_` to ignore error or check `err != nil`
- **Interface Conversion**:
    - **Type Assertion**:
        ```go
        value, ok := interfaceVar.(TargetType)
        ```
    - **Type Conversion**:
        ```go
        TargetType(value)
        ```

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
57257 / 5000 = 11
57257 / 5000 = 11.4514

Before: [string] 5112151121; After: [int] 5112151121.
Error: strconv.Atoi: parsing "This is Golang": invalid syntax

Before: [string] 3.1415926, After: [float64] 3.1415926.

Java is the best programming language in the world.
Before: [main.Programming] Java is the best programming language in the world.
After : [main.Programming] Golang is the best programming language in the world.

Java is the best programming language in the world.
Golang is the best programming language in the world.
```

## 5. Practice
1. Convert a `float64` value to an `int` and print both values.
2. Write a function that converts a numeric string to an integer and returns an error message if conversion fails.
3. Implement an interface type assertion to safely check if an interface implements a specific method.
4. Parse a string `"123.456"` into a `float64` and round it to two decimal places.
5. Write a program that converts a `struct` implementing two interfaces and calls methods from both interfaces.

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Type Conversions](https://gobyexample.com/type-conversions)
3. [Spec: Conversions](https://go.dev/ref/spec#Conversions)
4. [strconv Package](https://pkg.go.dev/strconv)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 31-AUG-2025 00:10 UTC +08:00*
</div>