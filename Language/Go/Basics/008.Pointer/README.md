<div align=center>

# Chapter 8: Pointers and Array Pointers in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **pointers in Go**, including:
    - Declaring and using pointers
    - Understanding `nil` pointers and null checks
    - Working with **arrays of pointers**
    - Using **pointer to an array**
    - Passing pointers to functions for in-place modification

It demonstrates:
    - How to obtain variable addresses using `&`
    - Dereferencing pointers using `*`
    - Storing addresses of array elements in a pointer array
    - Modifying original data through pointers without returning values

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
    - `fmt`: for formatted printing
- **Features Covered**:
    - Pointers and `nil` checking
    - Arrays of pointers
    - Pointer to an array
    - Passing pointer to function

## 3. Key Concepts
- **Pointer Basics**:
    - A pointer stores the address of a variable
    - Declare a pointer: `var p *int`
    - Get the address: `&variable`
    - Access value: `*p` (dereference)
- **Nil Pointers**:
    - Default value for uninitialized pointer is `nil`
    - Check with: `if p == nil { ... }`
- **Array of Pointers**:
    - Declare as: `var arrPtr [size]*int`
    - Assign addresses of array elements to pointers
- **Pointer to an Array**:
    - Declare as: `var ptr *[size]int`
    - Access elements via `(*ptr)[index]` or `ptr[index]`
- **Passing Pointers to Functions**:
    - Allows modifying original variable/array without returning
    - Example:
    ```go
    func modifyArray(arrPtr *[10]int) {
        (*arrPtr)[5] = 100
    }
    ```

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
The address of a: c000118040.

0xc000118040
&b        : c000118048.
pointerInt: c000118048.
b: 20, *pointerInt: 20.
b: 30, *pointerInt: 30.
<nil>
P is a null Porinter.
[0xc00013a000 0xc00013a008 0xc00013a010 0xc00013a018 0xc00013a020 0xc00013a028 0xc00013a030 0xc00013a038 0xc00013a040 0xc00013a048]
1, 3, 5, 7, 9, 2, 4, 6, 8, 10,

[2 4 6 8 10 1 3 5 7 9]
&[2 4 6 8 10 1 3 5 7 9]
2
2
[2 4 6 8 10 1 3 5 7 9]
[2 4 6 8 10 100 3 5 7 9]
```

## 5. Practice
1. Declare a pointer to a `float64` and modify its value using dereferencing.
2. Create an array of `string` pointers and print each value.
3. Implement a function that swaps two integers using pointers.
4. Create a function that doubles every element in an array using a pointer to the array.
5. Check what happens if you dereference a `nil` pointer (handle safely).

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Pointers](https://gobyexample.com/pointers)
3. [Spec: Pointers](https://go.dev/ref/spec#Pointer_types)
4. [Spec: Address Operators](https://go.dev/ref/spec#Address_operators)
5. [Effective Go - Pointers vs Values](https://go.dev/doc/effective_go#pointers_vs_values)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 30-AUG-2025 23:13 UTC +08:00*
</div>