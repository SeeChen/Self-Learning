<div align=center>

# Chapter 4: Conditional Statements in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **conditional statements in Go**, including:
- `if` and `if...else` for decision making
- `switch` for multi-condition branching
- `fallthrough` in `switch` for forced execution of next case
- `select` for channel-based concurrency decisions

It demonstrates:
- How `if` works without parentheses
- Using inline variable declaration in `if`
- Type switch for interface type checking
- Channel-based `select` statement for goroutine communication

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
  - `fmt`: for formatted printing
  - `strings`: for substring checking
  - `time`: for concurrency simulation
- **Features Covered**:
  - `if`, `else if`, `else`
  - `switch` and type-switch
  - `fallthrough`
  - `select` with channels

## 3. Key Concepts
- **if Statement**:
  - Syntax: `if condition { ... }`
  - No parentheses required around condition
  - Optional inline variable initialization: `if x := 10; x > 5 { ... }`
- **if...else**:
  - Provides alternative execution path when condition is false
- **switch Statement**:
  - Simplifies multiple `if...else` chains
  - No need for `break` keyword (Go automatically breaks after case execution)
  - Supports type-switch for checking interface type
- **fallthrough**:
  - Forces execution of the next case in `switch`
- **select Statement**:
  - Similar to `switch` but for channel operations
  - Used for concurrent communication between goroutines
  - Includes `default` case when no channel is ready

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
True
Not false
"This is a string" is contains "str".
20

200 [i] > 100 [j].
100 [i] < 200 [j].
150 [i] = 150 [j].

Score:  65, Grade: C.
Score: -20, Grade: F.
Score:  77, Grade: B.
Score:  94, Grade: A.

nil

Without fallthrough
10

With fallthrough
10
20
No Number
No Communication ready.
```

## 5. Practice
1. Write an `if...else` statement that checks whether a number is positive, negative, or zero.
2. Modify the `switch` statement to handle weekdays and weekends.
3. Implement a type-switch that checks whether an `interface{}` value is `int`, `string`, or `bool`.
4. Create a `select` statement that reads from three channels and prints which one sends first.
5. Experiment with adding `fallthrough` in your `switch` statement and observe the difference.

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - If/Else](https://gobyexample.com/if-else)
3. [Go by Example - Switch](https://gobyexample.com/switch)
4. [Go by Example - Select](https://gobyexample.com/select)
5. [Spec: Statements](https://go.dev/ref/spec#Statements)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 26-AUG-2025 23:10 UTC +08:00*
</div>