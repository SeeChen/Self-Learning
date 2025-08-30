<div align=center>

# Chapter 9: Structs in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **structs in Go**, including:
    - Defining and initializing structs
    - Struct pointers for modification
    - Methods with value and pointer receivers
    - Function fields inside structs (function as a member)
    - Nested structs for composition
    - Anonymous (embedded) fields for simplified field access

It demonstrates:
    - How to declare and initialize structs in different ways
    - Passing structs by value vs by pointer
    - Using methods to operate on struct data
    - Creating function-type fields for dynamic behavior
    - Composition as an alternative to inheritance

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
    - `fmt`: for formatted printing
- **Features Covered**:
    - Struct declaration and initialization
    - Methods with value and pointer receivers
    - Embedded structs (anonymous fields)
    - Higher-order functions via function fields

## 3. Key Concepts
- **Struct Declaration**:
  - Syntax:
    ```go
    type StructName struct {
        FieldName FieldType
        ...
    }
    ```
- **Initialization**:
    - Named fields: `User{Username: "Alice", Email: "...", Age: 25}`
    - Positional fields: `User{"Alice", "...", 25}`
    - Zero-value initialization and field assignment
- **Pointer vs Value Semantics**:
    - Passing by value does not change the original struct
    - Use `*Struct` and `&variable` to modify data directly
- **Methods**:
    - **Value receiver**: `(u User) MethodName()`
    - **Pointer receiver**: `(u *User) MethodName()` for mutation or efficiency
- **Function Fields**:
    - Struct can store functions:
    ```go
    type Calc struct {
        Operation func(int, int) int
    }
    ```
- **Nested Struct**:
    - Composition using structs inside structs
- **Anonymous Fields**:
    - Embedded struct allows direct access to inner fields
    - Example:
    ```go
    type Comics struct {
        Book
    }
    comics.Name // Access Book.Name directly
    ```

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
User1:
Username  : SeeChen
Email     : leeseechen@gmail.com
Age       : 25

User2:
Username  : SeeChen2
Email     : leeseechen@email.com
Age       : 100

User3:
Username  : SeeChen 3
Email     : seechenlee@gmail.com
Age       : 0

Original User data
Username  : SeeChen
Email     : leeseechen@gmail.com
Age       : 25

New User data
Username  : SeeChen SeeChen
Email     : leeseechen@gmail.com
Age       : 25

===
Username  : SeeChen
Email     : leeseechen@gmail.com
Age       : 25

Original User data
Username  : SeeChen
Email     : leeseechen@gmail.com
Age       : 25

New User data
Username  : SeeChen SeeChen
Email     : leeseechen@gmail.com
Age       : 25

===
Username  : SeeChen SeeChen
Email     : leeseechen@gmail.com
Age       : 25

Username  : SeeChen 22
Email     : leeseechen@email.com
Age       : 100

Hi, my name is SeeChen SeeChen, my email is leeseechen@gmail.com, my age is 25, my birthday is 06 Aug 2000.

Hi, my name is SeeChen 3, my email is seechenlee@gmail.com, my age is 0, my birthday is 06 Aug 2000.

10 + 5 = 15
10 - 5 = 5
{This is a novel {Novel 100 Novel Author}}
Novel Author
{{Comics 200 Comics Author}}
Comics
```

## 5. Practice
1. Create a `Car` struct with fields `Brand`, `Model`, and `Year`.  
   Write a method that prints: `"This is a <Brand> <Model> from <Year>."`
2. Modify the program to store multiple users in a `[]User` slice and print them all.
3. Implement a method on `Calc` to multiply and divide two integers.
4. Add another nested struct, such as `Publisher` inside `Book`, and access it from `Novel`.
5. Experiment with embedded structs and see how field name conflicts are resolved.

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Structs](https://gobyexample.com/structs)
3. [Effective Go - Structs](https://go.dev/doc/effective_go#structs)
4. [Go Spec: Struct Types](https://go.dev/ref/spec#Struct_types)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 30-AUG-2025 23:19 UTC +08:00*
</div>