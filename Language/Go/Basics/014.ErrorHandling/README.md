<div align=center>

# Chapter 14: Error Handling in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter explains **error handling in Go**, including:
    - Built-in `error` interface
    - Creating custom error types
    - Error wrapping and chaining
    - Using `errors.Is` and `errors.As` for error inspection
    - Implementing panic and recover for graceful recovery

It demonstrates:
    - How to create simple errors using `errors.New`
    - How to join multiple errors using `errors.Join`
    - Custom error structs that implement the `error` interface
    - Handling error chains using `errors.Is` and `errors.As`
    - `panic` for unrecoverable errors and `recover` for controlled recovery

## 2. Technology
- **Language**: Go 1.24
- **Packages Used**:
    - `errors`: for error creation and handling
    - `fmt`: for formatted output and error formatting
    - `math`: for square root function
    - `strings`: for string operations
- **Features Covered**:
    - Built-in `error` interface
    - Custom error types
    - Error wrapping with `%w` in `fmt.Errorf`
    - Panic and recover mechanism

## 3. Key Concepts
- **error Interface**:
    - Defined as:
        ```go
        type error interface {
            Error() string
        }
        ```
- **Creating Errors**:
    - Using `errors.New("message")`
    - Using `fmt.Errorf("formatted error")`
- **Error Wrapping**:
    - `fmt.Errorf("%w", err)` allows error unwrapping
- **errors.Is**:
    - Check if an error is or wraps a specific error
- **errors.As**:
    - Convert an error to a specific type
- **Custom Errors**:
    - Create struct implementing `Error() string` method
- **panic and recover**:
    - `panic`: stops execution, prints stack trace
    - `recover`: recovers from panic inside a deferred function

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
this is an error

First
Second

Error: math: square root of negative number

0

        Error - 520.
        Error Message: 520.
        Error Code: Not Contains "SeeChen"..


        Error - 520.
        Error Message: 520.
        Error Code: Not Contains "SeeChen"..



        Error - 520.
        Error Message: 520.
        Error Code: Not Contains "SeeChen"..



        Error - 520.
        Error Message: 520.
        Error Code: Not Contains "SeeChen"..


Yes, seechen seechen seechen
Error Code: 404
Error Message: Not Found

Yes!! SeeChen

Error: Lee, Panic Reason: The string does not contain 'SeeChen'!
Continue to run!
```

## 5. Practice
1. Write a custom error type that includes a timestamp and error severity level.
2. Modify `safeFunction` to log errors instead of printing them.
3. Implement a function that uses `errors.Join` to merge three different errors and output the combined message.
4. Write a panic scenario that simulates a file read error and recover from it gracefully.
5. Use `errors.As` to extract detailed error information from a wrapped custom error type.

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)
2. [Go by Example - Errors](https://gobyexample.com/errors)
3. [Go by Example - Panic and Recover](https://gobyexample.com/panic)
4. [Spec: Errors](https://go.dev/ref/spec#Errors)
5. [Go 1.13 Errors Package Enhancements](https://blog.golang.org/go1.13-errors)

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 01-SEPT-2025 16:46 UTC +08:00*
</div>