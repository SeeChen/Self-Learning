<div align=center>

# Chapter 2: Variables and Constants in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces variables and constants in Go, covering:
- Different ways to declare variables:
  - `var` keyword
  - Short declaration `:=`
  - Grouped declarations
- Constant declaration and usage
- Type inference and explicit typing
- Scope and initialization rules

It demonstrates:
- How variables can be declared with and without initial values
- The concept of default zero values
- Using constants for immutable values

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
  - `fmt`: for formatted printing
- **Features Covered**:
  - Variable declaration (`var`, `:=`)
  - Constants (`const`)
  - Type inference
  - Grouped declarations

## 3. Key Concepts
- **Variable Declaration**:
  - `var name type`: declares a variable with zero value
  - `var name = value`: type inferred
  - `name := value` : short declaration (only inside functions)
- **Constants**:
  - Declared using `const`
  - Must be assigned at compile time (cannot use runtime values)
- **Zero Values**:
  - Numeric: `0`
  - String: `""`
  - Boolean: `false`
- **Scope**:
  - Variables declared inside a function have local scope
  - Global variables declared outside functions

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
My name is SeeChen.
My name is SeeChen.

bool
string
int

1, String.

0, 0, 0
Var 4, Var 5, Var 6
1, My String, true
This, false, 9
0, String

0: Const 2, Const 3.
Male   is 0
Female is 1
Other  is 2

iota 1:  0
iota 2:  1

iota 1:  0
iota 2:  1
iota 3:  2
```

## 5. Practice
1. Declare three variables without initializing them. Print their zero values.
2. Use short declaration `:=` to declare and initialize multiple variables in one line.
3. Create a grouped declaration for 3 variables of different types.
4. Define a constant `MaxUsers` and try to reassign it. Observe the error.
5. Experiment with using constants in expressions (e.g., `const Radius = 10; const Area = PI * Radius * Radius`).

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Variables](https://gobyexample.com/variables)
3. [Go by Example - Constants](https://gobyexample.com/constants)
4. [Effective Go - Declarations and Scope](https://go.dev/doc/effective_go#declarations)
5. [Spec: Constants](https://go.dev/ref/spec#Constants)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 26-AUG-2025 22:59 UTC +08:00*
</div>