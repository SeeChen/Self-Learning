<div align=center>

# Chapter 0: Hello Go!

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This is the very first step in the **Golang Self-Learning Journey**.  
It demonstrates:
- How to print `Hello World!` in Go.
- How to use **string formatting** with `fmt.Sprintf` and `fmt.Printf`.
- How to assign a formatted string to a variable and print it.

## 2. Technology
- **Language**: Go 1.24
- **IDE**: VS Code
- **OS**: Linux / Windows
- **Packages Used**:
  - `fmt` (for formatted I/O)

## 3. Key Concepts
- **Basic Program Structure**:
  - `package main` -> Entry package for Go applications.
  - `func main()` -> Main function as the program entry point.
- **Output Methods**:
  - `fmt.Println()` -> Prints with a newline.
  - `fmt.Sprintf()` -> Returns a formatted string.
  - `fmt.Printf()` -> Formats and prints directly.
- **Format Specifiers**:
  - `%d` -> Integer
  - `%s` -> String

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
Hello World!
Number: 1, Date: 2025-07-22, Time: 08:00 GMT+08:00.
Number: 1, Date: 2025-07-22, Time: 08:00 GMT+08:00.
```

## 5. Practice
1. Change the value of Code to 2 and run the program.
2. Replace Date with the current date and print it.
3. Use fmt.Printf to output multiple formatted lines, for example:
```bash
Name: Alice, Age: 25
Name: Bob, Age: 30
```

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [fmt Package Docs](https://pkg.go.dev/fmt)
3. [Go by Example - Hello World](https://gobyexample.com/hello-world)
4. [Go by Example - String Formatting](https://gobyexample.com/string-formatting)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 25-AUG-2025 23:14 UTC +08:00*
</div>