<div align=center>

# Chapter 11: Maps in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **Maps in Go**, which are:
    - An **unordered collection of key-value pairs**
    - Similar to JSON-like data structures
    - Useful for fast lookups and dynamic data storage

It demonstrates:
    - Creating maps using `make()`
    - Initializing maps with literal syntax
    - Accessing values by key
    - Checking for key existence
    - Adding and updating key-value pairs
    - Deleting keys with `delete()`
    - Checking if a map is empty using `len()`

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
    - `fmt`: for formatted printing
- **Features Covered**:
    - Map declaration and initialization
    - Key-value operations (`get`, `set`, `delete`)
    - Existence check (`key, ok := map[key]`)
    - Iterating over a map using `range`

## 3. Key Concepts
- **Map Declaration**:
    - Syntax: `map[KeyType]ValueType`
    - Example: `map[string]int`
- **Creating a Map**:
    - Using `make()`: `make(map[string]int)`
    - Using literal:  
    ```go
    map[string]string{
        "Key1": "Value1",
        "Key2": "Value2",
    }
    ```
- **Accessing Values**:
    - `value := myMap["Key"]`
    - Check existence:  
    ```go
    v, ok := myMap["Key"]
    if ok {
        // key exists
    }
    ```
- **Adding or Updating Values**:
    - `myMap["Key"] = "NewValue"`
- **Deleting a Key**:
    - `delete(myMap, "Key")`
- **Iterating a Map**:
    - `for k, v := range myMap { fmt.Println(k, v) }`
- **Check if Empty**:
    - `len(myMap) == 0`

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
map[]
My First Name is SeeChen.
Middle name does not exists.
FirstName : SeeChen.
FamilyName: Lee.
MiddleName: Handsome.

FirstName : SeeChen.
FamilyName: Lee.
MiddleName: Very Handsome.

map[FamilyName:Lee FirstName:SeeChen]

Map contains values.
```

## 5. Practice
1. Create a map that stores the names of countries as keys and their capitals as values.
2. Write a function that counts the frequency of each word in a given sentence using a map.
3. Modify the program to delete multiple keys from a map and print the result.
4. Implement a map that uses integers as keys and slices of strings as values.
5. Write code to merge two maps into a single map.

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Maps](https://gobyexample.com/maps)
3. [Spec: Map Types](https://go.dev/ref/spec#Map_types)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 31-AUG-2025 00:04 UTC +08:00*
</div>