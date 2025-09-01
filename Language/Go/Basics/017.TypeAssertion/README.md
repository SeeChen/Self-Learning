<div align=center>

# Chapter 17: Type Assertion and Reflection in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **type assertion and reflection in Go**, including:
    - Basic type assertion with and without the `ok` idiom
    - Using `type switch` for safer multi-type handling
    - Applying the `comma-ok` idiom with `if-else`
    - Using `reflect` to inspect type and value dynamically
    - Asserting interfaces to **custom error types**

It demonstrates:
    - Handling failed assertions gracefully with `defer` + `recover`
    - Checking multiple possible types at runtime
    - Reflecting type name, kind, and value
    - Differentiating between standard `error` and custom error implementations

## 2. Technology
- **Language**: Go 1.24  
- **Packages Used**:
    - `fmt`: for formatted output
    - `reflect`: for runtime type inspection
- **Features Covered**:
    - Type assertion (`value.(T)`)
    - `ok` idiom (`value.(T) → (result, ok)`)
    - `switch v := value.(type)`
    - Reflection with `reflect.TypeOf` and `reflect.ValueOf`
    - Custom error type assertion

## 3. Key Concepts
- **Type Assertion**:
    - Syntax: `value.(T)`
    - Panics if `value` is not of type `T`
    - Safer with `ok` idiom: `v, ok := value.(T)`
- **Assertion without ok**:
    - May panic, but can be recovered with `defer` + `recover`
- **Type Switch**:
    - Safer alternative to multiple assertions
    - Syntax:
        ```go
        switch v := value.(type) {
        case int:
            ...
        case string:
            ...
        default:
            ...
        }
        ```
- **if-else + ok Idiom**:
    - Sequential type checks with `comma-ok`
- **Reflection**:
    - `reflect.TypeOf`: retrieves type metadata
    - `reflect.ValueOf`: retrieves actual value
    - More generic than type assertion, but slower
- **Custom Error Assertion**:
    - Useful when working with multiple error implementations
    - `if ce, ok := err.(customError); ok { ... }`

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
Assertion to string ok=true, value="Hello World!"
Assertion to int ok=false, value=0

[Recover] Type assertion failed: interface conversion: interface {} is string, not int
Assertion success: original=12345, asserted(int)=12345

[string] "This is a string" + "1" = "This is a string1"
[int] 114514 + 1 = 114515
[bool] value = true
[Unknown type] value = 3.14 (type = float64)

[if-else] String: "Example string"
[if-else] Integer: 999
[if-else] Unknown type: [1 2 3] (type = []int)

[reflect] Type = string, Kind = string, Value = Hello
[reflect] Type = int, Kind = int, Value = 42
[reflect] Type = bool, Kind = bool, Value = false

[Custom Error] Something went wrong
[Custom Error] Not a customError, type=*errors.errorString
```

## 5. Practice
1. Modify `withoutOk` to handle both `int` and `string` assertions.  
2. Extend `switchMethod` to handle `float64` and print its square.  
3. Use `ifElseMethod` to detect slices (`[]int`, `[]string`).  
4. Implement a reflection-based function that prints **all fields of a struct**.  
5. Create a new custom error type (e.g., `networkError`) and extend `customErrorExample` to handle it.  

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)  
2. [Go by Example - Type Assertions](https://gobyexample.com/type-assertions)  
3. [Go by Example - Type Switches](https://gobyexample.com/type-switches)  
4. [Spec: Type assertions](https://go.dev/ref/spec#Type_assertions)  
5. [Spec: Switch statements](https://go.dev/ref/spec#Switch_statements)  
6. [Package reflect](https://pkg.go.dev/reflect)  

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 01-SEPT-2025 17:07 UTC +08:00*
</div>