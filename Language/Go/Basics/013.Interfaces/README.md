<div align=center>

# Chapter 13: Interfaces in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>


## 1. Summary
This chapter introduces **Interfaces in Go**, including:
    - **Defining and implementing interfaces**
    - **Empty interface (`interface{}`) for dynamic types**
    - **Composite interfaces (nested interfaces)**
    - **Polymorphism through interfaces**
    - **Type assertion and type switch**

It demonstrates:
    - Implicit interface implementation in Go (no `implements` keyword)
    - Using `interface{}` to handle any type
    - Combining multiple interfaces using embedding
    - Dynamic type and value storage within an interface
    - Type-switch for type-specific behavior

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
    - `fmt`: for formatted output
    - `reflect`: for type introspection
- **Features Covered**:
    - Interface definition and implementation
    - Empty interface and dynamic typing
    - Composite interfaces
    - Type switch for runtime type checking

## 3. Key Concepts
- **Interface in Go**:
    - A collection of method signatures (behavior).
    - Implemented implicitly: no explicit keyword required.
- **Empty Interface (`interface{}`)**:
    - Represents any type.
    - Useful for generic functions and dynamic data handling.
- **Composite Interface**:
    - Combine multiple interfaces into one.
    - Example:
        ```go
        type Reader interface { Read() string }
        type Writer interface { Write(data string) }
        type ReadWrite interface {
            Reader
            Writer
        }
        ```
- **Dynamic Type and Value**:
    - Interface variables store both **type** and **value** dynamically.
- **Type Switch**:
    - Syntax:
        ```go
        switch v := x.(type) {
        case int:
            fmt.Println("Integer")
        case string:
            fmt.Println("String")
        default:
            fmt.Println("Unknown")
        }
        ```

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
This is return from interface method, First Interface

string
int
bool
[]int

This is a new data.
Type: int       ; Value: 114514
Type: string    ; Value: Hello World.

<nil>
true
Interger
```

## 5. Practice
1. Create another interface `Printer` with method `Print()` and implement it in a struct.
2. Modify `typeof()` to return both **type** and **kind** of the value.
3. Write a function that accepts `interface{}` and:
   - Uses `type switch` to print custom messages for `int`, `float64`, `string`.
4. Implement a new struct `LogFile` that satisfies `ReadWrite` interface and adds `Log()` method.
5. Try declaring an interface variable without initializing it and check its value (`nil`).

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Interfaces](https://gobyexample.com/interfaces)
3. [Effective Go - Interfaces](https://go.dev/doc/effective_go#interfaces)
4. [Spec: Interfaces](https://go.dev/ref/spec#Interface_types)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 01-SEPT-2025 16:40 UTC +08:00*
</div>