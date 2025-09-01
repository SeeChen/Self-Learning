<div align=center>

# Chapter 18: Inheritance & Polymorphism in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **inheritance and polymorphism in Go**, focusing on:
    - Using **struct embedding** to simulate inheritance
    - Defining and implementing **interfaces**
    - Overriding embedded struct methods in child structs
    - Demonstrating **polymorphism** by handling different types with the same interface

It demonstrates:
    - How embedding works in Go
    - Overriding methods in child structs (`Dog`, `Cat`)
    - Polymorphic behavior through interface slices
    - Each type’s specific methods (`Speak` for `Dog` and `Cat`)

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
    - `fmt`: for formatted printing
- **Features Covered**:
    - Struct embedding
    - Interfaces
    - Method overriding
    - Polymorphism via interface

## 3. Key Concepts
- **Interface**:
    - Defines a set of methods that types must implement
    - Example: `MyName` interface with method `MyName() string`
- **Struct Embedding**:
    - Allows code reuse and simulates inheritance
    - Example: `Dog` and `Cat` both embed `Animal`
- **Method Overriding**:
    - Child struct can redefine the embedded struct’s behavior
    - Example: `Dog.MyName()` vs `Cat.MyName()`
- **Polymorphism**:
    - Different types (`Dog`, `Cat`) can be treated uniformly via `MyName` interface
- **Concrete Methods**:
    - Specific to struct and not part of the interface
    - Example: `Dog.Speak()` and `Cat.Speak()`

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
I am a Dog, my name is Buddy.
I am a Cat named Kitty with white fur.
Buddy barks: Woof Woof
Kitty meows softly.
```

## 5. Practice
1. Add another animal type (e.g., `Bird`) that implements the `MyName` interface.
2. Override `MyName()` differently for the new type and add a `Speak()` method.
3. Create a slice of `MyName` and insert `Dog`, `Cat`, and `Bird` to test polymorphism.
4. Add more attributes to `Animal` (e.g., `Age`) and use them in overridden methods.
5. Try removing `Dog.MyName()` and see how Go falls back to the embedded `Animal` (default behavior).

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Interfaces](https://gobyexample.com/interfaces)
3. [Effective Go - Embedding](https://go.dev/doc/effective_go#embedding)
4. [Spec: Method sets](https://go.dev/ref/spec#Method_sets)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 01-SEPT-2025 17:14 UTC +08:00*
</div>