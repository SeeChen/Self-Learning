<div align=center>

# Chapter 10: Slices in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **Slices in Go**, which are more powerful and flexible than arrays:
    - **Dynamic size** (unlike arrays with fixed length)
    - **Reference type** that points to an underlying array
    - **Slice operations** such as `append`, `len()`, `cap()`
    - **Slice extraction** using `[lower:upper]` syntax

It demonstrates:
    - How to **declare and initialize slices**
    - Using `make()` to create slices with specific **length** and **capacity**
    - The difference between **nil slices**, empty slices, and initialized slices
    - **Appending elements** and observing capacity changes
    - **Slice slicing (sub-slices)** for extracting parts of a slice

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
    - `fmt`: for formatted output
- **Features Covered**:
    - Slice declaration: `var slice []int`
    - Slice initialization with literals: `[]int{1,2,3}`
    - `make([]T, length, capacity)`
    - Functions: `len()`, `cap()`
    - `append()` for dynamic resizing
    - Slice extraction: `slice[a:b]`

## 3. Key Concepts
- **Slice vs Array**:
    - Array: fixed size, value type
    - Slice: dynamic size, reference type
- **Slice Initialization**:
    - `slice := []int{1, 2, 3}`
    - `slice := make([]int, length)` (capacity defaults to length)
    - `slice := make([]int, length, capacity)` (explicit capacity)
- **Nil Slice**:
    - Declared but not initialized slice → `nil`
- **Slice Functions**:
    - `len(slice)`: current number of elements
    - `cap(slice)`: total capacity
- **Appending**:
    - `append(slice, element)`
    - If capacity exceeded → new underlying array is allocated
- **Slice Extraction**:
    - `[low:high]` (low inclusive, high exclusive)
    - `[:high]` (from start)
    - `[low:]` (to end)

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
[1 2 3]
No  0. Zero; No  1. One; No  2. ; No  3. ; No  4. ; No  5. String; No  6. ; No  7. ; No  8. ; No  9. ;
The lenght of slice is 10
[Zero One    String     Append]
The lenght of slice is 11
Current slice2 size: 11
Current slice2 size: 12
slice2 length: 12; capacity: 22.
emptySlice is a nil slice -> true.

Full slice: [Index_0 Index_1 Index_2 Index_3 Index_4 Index_5 Index_6 Index_7 Index_8 Index_9 FirstAppend SecondAppend]
Get the slice item between index 2 and index 5: [Index_2 Index_3 Index_4]
Start from 0: [Index_0 Index_1 Index_2 Index_3]
Start from 0: [Index_0 Index_1 Index_2 Index_3]
End in last item: [Index_3 Index_4 Index_5 Index_6 Index_7 Index_8 Index_9 FirstAppend SecondAppend].
End in last item: [Index_3 Index_4 Index_5 Index_6 Index_7 Index_8 Index_9 FirstAppend SecondAppend].
```

## 5. Practice
1. Create a slice of integers and append 5 new elements. Print its length and capacity after each append.
2. Write a program to copy one slice into another using the built-in `copy()` function.
3. Implement a function that accepts a slice and returns a new slice without modifying the original.
4. Experiment with sub-slices: create a sub-slice from index 3 to the end, modify it, and check if the original slice changes.
5. Create a multi-dimensional slice and print its elements.

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Slices](https://gobyexample.com/slices)
3. [Spec: Slice Types](https://go.dev/ref/spec#Slice_types)
4. [Effective Go: Slices](https://go.dev/doc/effective_go#slices)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 30-AUG-2025 23:25 UTC +08:00*
</div>