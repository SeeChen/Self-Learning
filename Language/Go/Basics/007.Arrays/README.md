<div align=center>

# Chapter 7: Arrays in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **arrays in Go**, covering:
- Declaring and initializing arrays
- Accessing and modifying elements
- Using `[...]` for auto-sized arrays
- Multidimensional arrays (2D and 3D)
- Differences between arrays and slices
- Array type identity (size as part of type)

It demonstrates:
- Fixed-size arrays and zero-value initialization
- Index-based initialization
- Iterating over arrays with `for` loops
- Dynamically building 3D structures using slices
- Checking types with the `reflect` package

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
    - `fmt`: for printing arrays and values
    - `math/rand`: for generating random numbers
    - `reflect`: for inspecting array types
- **Features Covered**:
    - Array declaration and initialization
    - Loop-based array population
    - Multidimensional arrays
    - Slice-based dynamic structures
    - Type reflection in Go

## 3. Key Concepts
- **Array Declaration**:
    - Syntax: `var arr [size]Type`
    - Default values depend on element type (e.g., `0` for `int`, `""` for `string`)
- **Array Initialization**:
    - Full initialization: `var arr [5]int = [5]int{1, 2, 3, 4, 5}`
    - Partial initialization with index: `arr := [5]int{2: 10, 4: 100}`
    - Auto-size arrays using `...`: `arr := [...]int{1, 2, 3}`
- **Accessing Elements**:
    - Indexed access: `arr[i]`
    - Length with `len(arr)`
- **Multidimensional Arrays**:
    - Fixed-size 2D array: `var arr2D [3][3]int`
    - Dynamically constructed 3D arrays using slices
- **Array vs Slice**:
    - Arrays have fixed length; slices are dynamic
    - Array size is part of its type
- **Type Checking**:
    - Using `reflect.TypeOf()` to verify types

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
[0 0 0 0 0 0 0 0 0 0]
[SeeChen Lee --- 李 思净]
[0 0 10 0 100]
[1 2 3 4 5]

Full Arrays: [482 406 877 912 430 950 923 214 24 293 252 1 990 867 603].
arr[ 0]:  482.
arr[ 1]:  406.
arr[ 2]:  877.
arr[ 3]:  912.
arr[ 4]:  430.
arr[ 5]:  950.
arr[ 6]:  923.
arr[ 7]:  214.
arr[ 8]:   24.
arr[ 9]:  293.
arr[10]:  252.
arr[11]:    1.
arr[12]:  990.
arr[13]:  867.
arr[14]:  603.
[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
arr[0][0][0]: 992, arr[0][0][1]: 222, arr[0][0][2]: 525, arr[0][0][3]: 674, arr[0][0][4]: 358, arr[0][0][5]: 581,
arr[0][1][0]:  60, arr[0][1][1]: 296, arr[0][1][2]: 859, arr[0][1][3]: 388, arr[0][1][4]: 356, arr[0][1][5]: 636,
arr[0][2][0]: 774, arr[0][2][1]:  91, arr[0][2][2]: 715, arr[0][2][3]: 572, arr[0][2][4]: 669, arr[0][2][5]: 245,
arr[0][3][0]: 596, arr[0][3][1]: 793, arr[0][3][2]: 621, arr[0][3][3]: 861, arr[0][3][4]: 641, arr[0][3][5]: 281,
arr[0][4][0]:  92, arr[0][4][1]: 310, arr[0][4][2]: 207, arr[0][4][3]:  59, arr[0][4][4]: 567, arr[0][4][5]: 451,
arr[0][5][0]: 849, arr[0][5][1]: 671, arr[0][5][2]: 590, arr[0][5][3]: 286, arr[0][5][4]: 416, arr[0][5][5]: 449,
arr[0][6][0]: 770, arr[0][6][1]: 962, arr[0][6][2]: 563, arr[0][6][3]: 203, arr[0][6][4]: 231, arr[0][6][5]: 494,

arr[1][0][0]: 574, arr[1][0][1]: 134, arr[1][0][2]: 936, arr[1][0][3]: 647, arr[1][0][4]: 777, arr[1][0][5]: 247,
arr[1][1][0]: 346, arr[1][1][1]: 114, arr[1][1][2]: 496, arr[1][1][3]: 756, arr[1][1][4]: 353, arr[1][1][5]: 371,
arr[1][2][0]: 893, arr[1][2][1]: 129, arr[1][2][2]: 573, arr[1][2][3]: 367, arr[1][2][4]: 856, arr[1][2][5]: 553,
arr[1][3][0]: 795, arr[1][3][1]: 902, arr[1][3][2]: 877, arr[1][3][3]: 472, arr[1][3][4]: 322, arr[1][3][5]: 213,
arr[1][4][0]: 775, arr[1][4][1]: 126, arr[1][4][2]: 650, arr[1][4][3]: 229, arr[1][4][4]: 904, arr[1][4][5]:  34,
arr[1][5][0]: 726, arr[1][5][1]: 282, arr[1][5][2]: 925, arr[1][5][3]: 112, arr[1][5][4]: 620, arr[1][5][5]: 160,
arr[1][6][0]: 698, arr[1][6][1]: 291, arr[1][6][2]:  77, arr[1][6][3]: 580, arr[1][6][4]: 445, arr[1][6][5]: 267,

arr[2][0][0]: 111, arr[2][0][1]:  61, arr[2][0][2]:  14, arr[2][0][3]: 854, arr[2][0][4]: 107, arr[2][0][5]: 977,
arr[2][1][0]: 470, arr[2][1][1]: 862, arr[2][1][2]: 676, arr[2][1][3]: 618, arr[2][1][4]: 322, arr[2][1][5]: 882,
arr[2][2][0]: 255, arr[2][2][1]: 755, arr[2][2][2]: 393, arr[2][2][3]: 813, arr[2][2][4]:  69, arr[2][2][5]: 780,
arr[2][3][0]: 900, arr[2][3][1]: 955, arr[2][3][2]: 769, arr[2][3][3]: 451, arr[2][3][4]: 656, arr[2][3][5]: 403,
arr[2][4][0]: 871, arr[2][4][1]: 215, arr[2][4][2]: 523, arr[2][4][3]: 297, arr[2][4][4]: 756, arr[2][4][5]: 944,
arr[2][5][0]: 511, arr[2][5][1]: 526, arr[2][5][2]: 845, arr[2][5][3]: 744, arr[2][5][4]: 354, arr[2][5][5]: 892,
arr[2][6][0]: 645, arr[2][6][1]: 508, arr[2][6][2]: 511, arr[2][6][3]: 319, arr[2][6][4]: 841, arr[2][6][5]: 111,

arr[3][0][0]: 359, arr[3][0][1]: 326, arr[3][0][2]: 848, arr[3][0][3]: 257, arr[3][0][4]: 340, arr[3][0][5]: 155,
arr[3][1][0]: 304, arr[3][1][1]: 137, arr[3][1][2]: 163, arr[3][1][3]: 456, arr[3][1][4]: 839, arr[3][1][5]: 274,
arr[3][2][0]: 758, arr[3][2][1]: 809, arr[3][2][2]: 235, arr[3][2][3]: 300, arr[3][2][4]: 742, arr[3][2][5]: 263,
arr[3][3][0]: 982, arr[3][3][1]: 429, arr[3][3][2]: 230, arr[3][3][3]: 339, arr[3][3][4]: 278, arr[3][3][5]: 185,
arr[3][4][0]: 156, arr[3][4][1]: 944, arr[3][4][2]: 667, arr[3][4][3]: 645, arr[3][4][4]: 241, arr[3][4][5]: 908,
arr[3][5][0]: 456, arr[3][5][1]: 678, arr[3][5][2]: 442, arr[3][5][3]:  77, arr[3][5][4]:  76, arr[3][5][5]:  61,
arr[3][6][0]: 174, arr[3][6][1]: 241, arr[3][6][2]: 329, arr[3][6][3]: 140, arr[3][6][4]: 773, arr[3][6][5]: 509,

arr[4][0][0]: 594, arr[4][0][1]: 152, arr[4][0][2]: 501, arr[4][0][3]: 758, arr[4][0][4]: 888, arr[4][0][5]: 649,
arr[4][1][0]: 878, arr[4][1][1]: 526, arr[4][1][2]: 326, arr[4][1][3]: 427, arr[4][1][4]: 157, arr[4][1][5]: 890,
arr[4][2][0]: 950, arr[4][2][1]: 495, arr[4][2][2]: 619, arr[4][2][3]: 803, arr[4][2][4]: 590, arr[4][2][5]: 582,
arr[4][3][0]: 251, arr[4][3][1]: 690, arr[4][3][2]: 181, arr[4][3][3]: 289, arr[4][3][4]: 216, arr[4][3][5]: 834,
arr[4][4][0]: 487, arr[4][4][1]:   3, arr[4][4][2]: 480, arr[4][4][3]: 810, arr[4][4][4]: 973, arr[4][4][5]: 772,
arr[4][5][0]: 705, arr[4][5][1]: 623, arr[4][5][2]: 683, arr[4][5][3]: 703, arr[4][5][4]: 136, arr[4][5][5]: 817,
arr[4][6][0]: 450, arr[4][6][1]: 854, arr[4][6][2]: 103, arr[4][6][3]: 102, arr[4][6][4]: 677, arr[4][6][5]: 765,

arr[5][0][0]: 516, arr[5][0][1]: 605, arr[5][0][2]: 238, arr[5][0][3]: 927, arr[5][0][4]: 720, arr[5][0][5]: 656,
arr[5][1][0]: 432, arr[5][1][1]: 621, arr[5][1][2]: 413, arr[5][1][3]: 651, arr[5][1][4]:  73, arr[5][1][5]: 361,
arr[5][2][0]: 235, arr[5][2][1]: 631, arr[5][2][2]: 297, arr[5][2][3]:  33, arr[5][2][4]:  69, arr[5][2][5]: 541,
arr[5][3][0]: 487, arr[5][3][1]: 232, arr[5][3][2]: 456, arr[5][3][3]: 873, arr[5][3][4]:  22, arr[5][3][5]: 643,
arr[5][4][0]: 272, arr[5][4][1]: 243, arr[5][4][2]: 665, arr[5][4][3]: 190, arr[5][4][4]: 301, arr[5][4][5]: 633,
arr[5][5][0]:   5, arr[5][5][1]:  95, arr[5][5][2]: 749, arr[5][5][3]: 763, arr[5][5][4]: 732, arr[5][5][5]:  45,
arr[5][6][0]:  10, arr[5][6][1]: 539, arr[5][6][2]:  88, arr[5][6][3]: 543, arr[5][6][4]: 963, arr[5][6][5]: 153,

arr[6][0][0]: 262, arr[6][0][1]: 566, arr[6][0][2]:  59, arr[6][0][3]: 647, arr[6][0][4]: 630, arr[6][0][5]: 511,
arr[6][1][0]: 930, arr[6][1][1]: 335, arr[6][1][2]: 721, arr[6][1][3]: 278, arr[6][1][4]: 840, arr[6][1][5]: 685,
arr[6][2][0]: 571, arr[6][2][1]: 881, arr[6][2][2]: 926, arr[6][2][3]: 812, arr[6][2][4]: 889, arr[6][2][5]: 633,
arr[6][3][0]: 414, arr[6][3][1]: 110, arr[6][3][2]: 793, arr[6][3][3]:  25, arr[6][3][4]: 496, arr[6][3][5]: 900,
arr[6][4][0]: 803, arr[6][4][1]: 224, arr[6][4][2]: 505, arr[6][4][3]: 780, arr[6][4][4]: 109, arr[6][4][5]: 503,
arr[6][5][0]: 873, arr[6][5][1]: 270, arr[6][5][2]:  62, arr[6][5][3]: 130, arr[6][5][4]: 119, arr[6][5][5]: 350,
arr[6][6][0]: 767, arr[6][6][1]: 750, arr[6][6][2]: 880, arr[6][6][3]: 938, arr[6][6][4]: 819, arr[6][6][5]: 269,

arr[7][0][0]:  50, arr[7][0][1]: 870, arr[7][0][2]: 966, arr[7][0][3]: 701, arr[7][0][4]: 564, arr[7][0][5]:  19,
arr[7][1][0]: 897, arr[7][1][1]: 887, arr[7][1][2]: 615, arr[7][1][3]: 240, arr[7][1][4]: 959, arr[7][1][5]: 492,
arr[7][2][0]: 161, arr[7][2][1]: 243, arr[7][2][2]: 456, arr[7][2][3]:   8, arr[7][2][4]: 979, arr[7][2][5]: 497,
arr[7][3][0]: 422, arr[7][3][1]: 671, arr[7][3][2]:  75, arr[7][3][3]: 933, arr[7][3][4]: 147, arr[7][3][5]: 176,
arr[7][4][0]: 503, arr[7][4][1]: 586, arr[7][4][2]: 256, arr[7][4][3]: 246, arr[7][4][4]: 202, arr[7][4][5]: 456,
arr[7][5][0]:  96, arr[7][5][1]: 600, arr[7][5][2]: 868, arr[7][5][3]: 909, arr[7][5][4]: 956, arr[7][5][5]:  15,
arr[7][6][0]: 284, arr[7][6][1]: 207, arr[7][6][2]: 119, arr[7][6][3]: 369, arr[7][6][4]: 175, arr[7][6][5]: 783,

arr[8][0][0]: 876, arr[8][0][1]:  52, arr[8][0][2]: 417, arr[8][0][3]: 777, arr[8][0][4]: 638, arr[8][0][5]: 784,
arr[8][1][0]:  35, arr[8][1][1]: 873, arr[8][1][2]: 518, arr[8][1][3]: 918, arr[8][1][4]:  29, arr[8][1][5]: 203,
arr[8][2][0]: 604, arr[8][2][1]: 536, arr[8][2][2]: 581, arr[8][2][3]: 458, arr[8][2][4]: 323, arr[8][2][5]:  60,
arr[8][3][0]: 797, arr[8][3][1]: 918, arr[8][3][2]: 645, arr[8][3][3]: 700, arr[8][3][4]: 513, arr[8][3][5]: 780,
arr[8][4][0]: 916, arr[8][4][1]: 268, arr[8][4][2]:  31, arr[8][4][3]: 130, arr[8][4][4]: 153, arr[8][4][5]:  13,
arr[8][5][0]: 343, arr[8][5][1]: 451, arr[8][5][2]: 900, arr[8][5][3]: 621, arr[8][5][4]: 794, arr[8][5][5]: 698,
arr[8][6][0]: 356, arr[8][6][1]: 147, arr[8][6][2]: 815, arr[8][6][3]: 154, arr[8][6][4]: 359, arr[8][6][5]:   9,

arr[9][0][0]: 271, arr[9][0][1]: 266, arr[9][0][2]: 807, arr[9][0][3]:  71, arr[9][0][4]: 443, arr[9][0][5]: 174,
arr[9][1][0]: 745, arr[9][1][1]: 417, arr[9][1][2]: 854, arr[9][1][3]: 881, arr[9][1][4]: 852, arr[9][1][5]: 615,
arr[9][2][0]: 676, arr[9][2][1]: 528, arr[9][2][2]: 424, arr[9][2][3]:  57, arr[9][2][4]: 949, arr[9][2][5]: 651,
arr[9][3][0]: 542, arr[9][3][1]: 513, arr[9][3][2]: 950, arr[9][3][3]: 210, arr[9][3][4]: 603, arr[9][3][5]: 268,
arr[9][4][0]: 118, arr[9][4][1]: 746, arr[9][4][2]: 260, arr[9][4][3]: 422, arr[9][4][4]:  77, arr[9][4][5]: 705,
arr[9][5][0]: 441, arr[9][5][1]: 729, arr[9][5][2]: 565, arr[9][5][3]: 683, arr[9][5][4]: 332, arr[9][5][5]: 493,
arr[9][6][0]: 236, arr[9][6][1]: 746, arr[9][6][2]: 849, arr[9][6][3]: 239, arr[9][6][4]: 467, arr[9][6][5]: 241,

Type of arr3D (slice-based 3D array): [][][]int
Type of arr   (fixed-size array    ): [15]int
```

## 5. Practice
1. Declare an array of 10 strings and initialize it with your favorite words.
2. Create a 2D array (3x3) and fill it with multiplication table values.
3. Write a function that sums all elements of an integer array.
4. Convert the 3D slice-based structure into a function that returns the structure for given sizes.
5. Use `reflect.TypeOf()` to check if `[5]int` and `[10]int` are considered the same type in Go.

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Arrays](https://gobyexample.com/arrays)
3. [Go by Example - Slices](https://gobyexample.com/slices)
4. [Spec: Arrays](https://go.dev/ref/spec#Array_types)
5. [reflect Package](https://pkg.go.dev/reflect)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 29-AUG-2025 22:39 UTC +08:00*
</div>