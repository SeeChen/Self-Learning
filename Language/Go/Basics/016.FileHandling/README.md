<div align=center>

# Chapter 6: File Handling in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **file handling in Go**, covering:
    - File creation and opening
    - Multiple ways of writing to files
    - Reading files line by line or entirely
    - Appending and buffered writes
    - Retrieving file metadata
    - Directory operations (create, read, delete)
    - File copy, move, and rename
    - Temporary file and directory handling

It demonstrates:
    - Usage of `os`, `bufio`, `io`, `log`, and `filepath` packages
    - Safe file operations with `defer file.Close()`
    - Buffered vs. unbuffered writes
    - Directory traversal and entry inspection
    - Handling temporary resources with cleanup

## 2. Technology
- **Language**: Go 1.24  
- **Packages Used**:
    - `os`: core file/directory operations
    - `bufio`: buffered I/O for efficiency
    - `io`: file content copying
    - `fmt`: formatted input/output
    - `log`: logging errors and info
    - `path/filepath`: path manipulation  
- **Features Covered**:
    - File create/open/write/read/append
    - File metadata (`os.Stat`)
    - Directory create/read
    - File copy, move, rename
    - Temporary file/directory handling

## 3. Key Concepts
- **File Creation & Opening**:
    - `os.Create`: creates/truncates file  
    - `os.Open`: opens file in read-only mode  
    - `os.OpenFile`: supports flags like `O_RDONLY`, `O_WRONLY`, `O_RDWR`, `O_CREATE`, `O_TRUNC`, `O_APPEND`

- **Writing to File**:
    - `file.WriteString`, `file.Write([]byte)`, `fmt.Fprintf`  
    - Buffered write: `bufio.Writer` + `writer.Flush()`  
    - One-shot write: `os.WriteFile`  

- **Reading File**:
    - Line by line: `bufio.Scanner`  
    - Entire content: `os.ReadFile`  

- **Appending**:
    - Open with `O_APPEND | O_WRONLY`  

- **File Info**:
    - `os.Stat`: name, size, mode, last modified, isDir  

- **Directory Operations**:
    - `os.Mkdir`, `os.MkdirAll`: create dirs  
    - `os.ReadDir`: list contents  
    - `os.Remove`, `os.RemoveAll`: delete  

- **File Copy/Move**:
    - `io.Copy` for copy  
    - `os.Rename` for move/rename  

- **Temporary Files & Dirs**:
    - `os.CreateTemp`, `os.MkdirTemp`  
    - `os.TempDir` to get system temp path

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
2025/09/01 16:53:44 File created: "myFile.txt"
2025/09/01 16:53:44 File opened: "myFile.txt"
2025/09/01 16:53:44 File: "myFile.txt" is opened!
2025/09/01 16:53:44 Success write data to file: myFile.txt.
2025/09/01 16:53:44 Writing index 0: Here
2025/09/01 16:53:44 Writing index 1: is
2025/09/01 16:53:44 Writing index 2: SeeChen
2025/09/01 16:53:44 Writing index 3: Lee
2025/09/01 16:53:44 Successfully wrote rows to file: "myFile.txt"
2025/09/01 16:53:44 Successfully wrote entire content to file: "myFile.txt"
2025/09/01 16:53:44 Successfully appended 43 bytes to file: "myFile.txt"
2025/09/01 16:53:44 File "myFile.txt" is Reading.
2025/09/01 16:53:44
2025/09/01 16:53:44 Hi, SeeChen here.
2025/09/01 16:53:44 This is my Golang self-learning document.
2025/09/01 16:53:44 Each .go file contains my Golang learning notes and related code.
2025/09/01 16:53:44
2025/09/01 16:53:44 This line is appended by fileWriteAppend()
2025/09/01 16:53:44
Hi, SeeChen here.
This is my Golang self-learning document.
Each .go file contains my Golang learning notes and related code.

This line is appended by fileWriteAppend().

2025/09/01 16:53:44 File Name      : myFile.txt.
2025/09/01 16:53:44 File Size      : 170 Byte.
2025/09/01 16:53:44 File Auth      : -rw-r--r--.
2025/09/01 16:53:44 Last Modified  : 2025-09-01 16:53:44.3974505 +0800 +08.
2025/09/01 16:53:44 Is Direc       : false.
2025/09/01 16:53:44 Directory "File" Created successfully.
2025/09/01 16:53:44 Directory "File/Handling" Created successfully.
2025/09/01 16:53:44 File                : [Size] 00004096, [Last Modifited] 2025-09-01 16:53:44.4496661 +0800 +08.
2025/09/01 16:53:44 README.md           : [Size] 00003319, [Last Modifited] 2025-08-30 23:35:04.556424 +0800 +08.
2025/09/01 16:53:44 main.go             : [Size] 00011588, [Last Modifited] 2025-08-30 23:46:57.9629877 +0800 +08.
2025/09/01 16:53:44 myFile.txt          : [Size] 00000170, [Last Modifited] 2025-09-01 16:53:44.3974505 +0800 +08.
2025/09/01 16:53:44
2025/09/01 16:53:44 Successfully Copy content from "myFile.txt" to "copied.txt". Total 170 byte(s).
2025/09/01 16:53:44 Temporary file created: /tmp/example-850878797.txt
2025/09/01 16:53:44 Temporary directory created: /tmp/example-dir-1480300175
2025/09/01 16:53:44 File created inside temp directory: /tmp/example-dir-1480300175/myfile.txt
```

## 5. Practice
1. Modify the program to create and write to **two different files** in parallel.  
2. Change `fileRead` to count the number of lines in a file.  
3. Experiment with opening a file in `O_APPEND` mode without truncating it.  
4. Implement a function to **delete files older than 7 days** in a directory.  
5. Use `fileTemp` to simulate a cache system with automatic cleanup.  

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)  
2. [Go by Example - File I/O](https://gobyexample.com/writing-files)  
3. [Go by Example - Reading Files](https://gobyexample.com/reading-files)  
4. [Go by Example - Directories](https://gobyexample.com/directories)  
5. [Spec: Package os](https://pkg.go.dev/os)  

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 01-SEPT-2025 17:04 UTC +08:00*
</div>