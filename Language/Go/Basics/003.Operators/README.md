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
This chapter introduces **operators in Go**, including:
- Arithmetic operators (`+`, `-`, `*`, `/`, `%`)
- Increment and decrement (`++`, `--`)
- Relational operators (`==`, `!=`, `>`, `<`, `>=`, `<=`)
- Logical operators (`&&`, `||`, `!`)
- Bitwise operators (`&`, `|`, `^`)
- Shift operators (`<<`, `>>`)

It demonstrates:
- How arithmetic operations work
- Boolean comparison using relational operators
- Logical conditions with `&&`, `||`, `!`
- Bit-level operations with `&`, `|`, `^`
- Binary shift operations

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
  - `fmt`: for formatted output
- **Features Covered**:
  - Arithmetic, relational, logical, and bitwise operators
  - Binary representation using `%b` in `fmt.Printf`
  - Shift operators (`<<`, `>>`)

## 3. Key Concepts
- **Arithmetic Operators**:
  - `+` addition, `-` subtraction, `*` multiplication, `/` division, `%` remainder
- **Increment/Decrement**:
  - `a++` increases by 1, `a--` decreases by 1
- **Relational Operators**:
  - `==`, `!=`, `>`, `<`, `>=`, `<=`
  - Return `true` or `false`
- **Logical Operators**:
  - `&&` (AND), `||` (OR), `!` (NOT)
- **Bitwise Operators**:
  - `&` (AND), `|` (OR), `^` (XOR)
  - Operate on bits
- **Shift Operators**:
  - `<<` shifts bits left
  - `>>` shifts bits right
- **Binary Representation**:
  - Use `fmt.Printf("%b", value)` to display binary format

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
10 + 20 = 30.
10 - 20 = -10.
10 * 20 = 200.
10 / 20 = 0.500000.

11 % 10 = 1.

a++ = 21.
a-- = 14.

Symbol "==" -- The "Haha" is     Equal to "Haha": true.
Symbol "!=" -- The "Haha" is Not Equal to "haha": true.

Symbol "> " -- 10 >  10 is false
Symbol ">=" -- 10 >= 10 is true

Symbol "< " -- 10 <  20 is true
Symbol "<=" -- 10 <= 20 is true

Symbol "&&"
true  && true  = true.
true  && false = false.
false && true  = false.
false && false = false.

Symbol "||"
true  || true  = true.
true  || false = true.
false || true  = true.
false || false = false.

Symbol "!"
!true  = false.
!false = true.

Simple example:
0 & 0 = 0, 0 | 0 = 0, 0 ^ 0 = 0.
0 & 1 = 0, 0 | 1 = 1, 0 ^ 1 = 1.
1 & 0 = 0, 1 | 0 = 1, 1 ^ 0 = 1.
1 & 1 = 1, 1 | 1 = 1, 1 ^ 1 = 0.

121 & 10 = 8
 121 (000001111001)
  10 (000000001010)
   8 (000000001000)

121 | 10 = 123
 121 (000001111001)
  10 (000000001010)
 123 (000001111011)

121 ^ 10 = 115
 121 (000001111001)
  10 (000000001010)
 115 (000001110011)

123 >> 2 = 30
 123 (000001111011)
  30 (000000011110)

123 << 2 = 492
 123 (000001111011)
 492 (000111101100)
```

## 5. Practice
1. Write a program to calculate `(a + b) * c` for given integers.
2. Test relational operators with strings and numbers.
3. Experiment with logical operators on different boolean values.
4. Try bitwise operators on numbers like `5` and `3` and show the binary result.
5. Perform shift operations on a number and explain how the value changes.

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Operators](https://gobyexample.com/arithmetic)
3. [Go by Example - Bitwise Operators](https://gobyexample.com/bitwise-operations)
4. [Spec: Operators](https://go.dev/ref/spec#Operators)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 26-AUG-2025 23:10 UTC +08:00*
</div>