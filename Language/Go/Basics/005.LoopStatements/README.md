<div align=center>

# Chapter 5: Loop Statements in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **loop statements in Go**, including:
- The single loop keyword: `for`
- Different forms of `for`:
    - Traditional `for init; condition; post {}` (similar to C-style loop)
    - `for condition {}` (similar to `while` loop)
    - `for range` (for iterating over collections)
- Nested loops for complex iterations
- Loop control statements:
    - `break`
    - `continue`
    - `goto`

It demonstrates:
- Looping through slices and maps
- Generating a 9×9 multiplication table
- Using `break` and `continue` to control flow
- Using `goto` for label-based jumps

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
    - `fmt`: for formatted printing
- **Features Covered**:
    - `for` loop variations
    - Range-based iteration for slices and maps
    - Nested loops
    - Loop control statements: `break`, `continue`, `goto`

## 3. Key Concepts
- **Basic `for` loop**:
  - Syntax:  
    ```go
    for init; condition; post { ... }
    ```
  - Example:
    ```go
    for i := 0; i < 10; i++ {
        sum += 1
    }
    ```
- **`for` as `while` loop**:
  - Omit initialization and increment:
    ```go
    for sum < 20 {
        sum += 1
    }
    ```
- **`for range` loop**:
  - Iterate over slices:
    ```go
    for i, s := range strings {
        fmt.Printf("%d. %s\n", i, s)
    }
    ```
  - Iterate over maps:
    ```go
    for key, value := range myName {
        fmt.Printf("My %s is %s. ", key, value)
    }
    ```
- **Nested loops**:
  - Used for multi-dimensional iterations (e.g., multiplication table)
- **Loop control statements**:
  - `break`: Exit from a loop or inner loop
  - `continue`: Skip the current iteration and move to the next
  - `goto`: Jump to a labeled statement

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
10
20
0. Lee
1. See
2. Chen
My FirstName is SeeChen. My FamilyName is Lee.
My FirstName is SeeChen. My FamilyName is Lee.
Lee SeeChen

1 x 1 = 1,
2 x 1 = 2, 2 x 2 = 4,
3 x 1 = 3, 3 x 2 = 6, 3 x 3 = 9,
4 x 1 = 4, 4 x 2 = 8, 4 x 3 = 12, 4 x 4 = 16,
5 x 1 = 5, 5 x 2 = 10, 5 x 3 = 15, 5 x 4 = 20, 5 x 5 = 25,
6 x 1 = 6, 6 x 2 = 12, 6 x 3 = 18, 6 x 4 = 24, 6 x 5 = 30, 6 x 6 = 36,
7 x 1 = 7, 7 x 2 = 14, 7 x 3 = 21, 7 x 4 = 28, 7 x 5 = 35, 7 x 6 = 42, 7 x 7 = 49,
8 x 1 = 8, 8 x 2 = 16, 8 x 3 = 24, 8 x 4 = 32, 8 x 5 = 40, 8 x 6 = 48, 8 x 7 = 56, 8 x 8 = 64,
9 x 1 = 9, 9 x 2 = 18, 9 x 3 = 27, 9 x 4 = 36, 9 x 5 = 45, 9 x 6 = 54, 9 x 7 = 63, 9 x 8 = 72, 9 x 9 = 81,

0 1 2 3 4 5 6 7 8 9 10

1 x 1 = 1,
2 x 1 = 2, 2 x 2 = 4,
3 x 1 = 3, 3 x 2 = 6, 3 x 3 = 9,
4 x 1 = 4, 4 x 2 = 8, 4 x 3 = 12, 4 x 4 = 16,
5 x 1 = 5, 5 x 2 = 10, 5 x 3 = 15, 5 x 4 = 20, 5 x 5 = 25,
6 x 1 = 6, 6 x 2 = 12, 6 x 3 = 18, 6 x 4 = 24, 6 x 5 = 30, 6 x 6 = 36,
7 x 1 = 7, 7 x 2 = 14, 7 x 3 = 21, 7 x 4 = 28, 7 x 5 = 35, 7 x 6 = 42, 7 x 7 = 49,
8 x 1 = 8, 8 x 2 = 16, 8 x 3 = 24, 8 x 4 = 32, 8 x 5 = 40, 8 x 6 = 48, 8 x 7 = 56,
9 x 1 = 9, 9 x 2 = 18, 9 x 3 = 27, 9 x 4 = 36, 9 x 5 = 45, 9 x 6 = 54,

is not a Prime Number.
1 is a Prime Number.
is not a Prime Number.
2 is a Prime Number.
is not a Prime Number.
3 is a Prime Number.
4 is not a Prime Number.
5 is a Prime Number.
6 is not a Prime Number.
7 is a Prime Number.
8 9 10 is not a Prime Number.
11 is a Prime Number.
12 is not a Prime Number.
13 is a Prime Number.
14 15 16 is not a Prime Number.
17 is a Prime Number.
18 is not a Prime Number.
19 is a Prime Number.
20 21 22 is not a Prime Number.
23 is a Prime Number.
24 25 26 27 28 is not a Prime Number.
29 is a Prime Number.
30 is not a Prime Number.
31 is a Prime Number.
32 33 34 35 36 is not a Prime Number.
37 is a Prime Number.
38 39 40 is not a Prime Number.
41 is a Prime Number.
42 is not a Prime Number.
43 is a Prime Number.
44 45 46 is not a Prime Number.
47 is a Prime Number.
48 49 50 51 52 is not a Prime Number.
53 is a Prime Number.
54 55 56 57 58 is not a Prime Number.
59 is a Prime Number.
60 is not a Prime Number.
61 is a Prime Number.
62 63 64 65 66 is not a Prime Number.
67 is a Prime Number.
68 69 70 is not a Prime Number.
71 is a Prime Number.
72 is not a Prime Number.
73 is a Prime Number.
74 75 76 77 78 is not a Prime Number.
79 is a Prime Number.
80 81 82 is not a Prime Number.
83 is a Prime Number.
84 85 86 87 88 is not a Prime Number.
89 is a Prime Number.
90 91 92 93 94 95 96 is not a Prime Number.
97 is a Prime Number.
98 99

1 2 3 4 5 6 7 8 9
```

## 5. Practice
1. Modify the program to calculate the factorial of numbers using a `for` loop.
2. Write a `for range` loop to sum all elements in an integer slice.
3. Use a nested loop to print a triangle pattern of `*`.
4. Implement a loop that uses `continue` to skip even numbers and print only odd numbers.
5. Experiment with `goto` by creating a loop that jumps to a label when a condition is met.

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - For](https://gobyexample.com/for)
3. [Go by Example - Range](https://gobyexample.com/range)
4. [Spec: For statements](https://go.dev/ref/spec#For_statements)
5. [Spec: Control statements](https://go.dev/ref/spec#For_statements)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 29-AUG-2025 22:22 UTC +08:00*
</div>