package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// fileCreate creates a new empty file with the specified name.
// If the file already exists, its contents will be truncated.
// Always close the returned file to release system resources.
func fileCreate(fileName string) {
	file, err := os.Create(fileName) // Equivalent to OpenFile(fileName, O_RDWR|O_CREATE|O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Failed to create file: %q. Error: %v", fileName, err)
	}
	defer file.Close()
	log.Printf("File created: %q\n", fileName)
}

// fileOpen opens an existing file in read-only mode.
// Use os.OpenFile for more control over access flags and permissions.
func fileOpen(fileName string) {
	file, err := os.Open(fileName) // Opens in read-only mode (O_RDONLY)
	if err != nil {
		log.Printf("Failed to open file %q: %v\n", fileName, err)
		return
	}
	defer file.Close()

	log.Printf("File opened: %q\n", fileName)
}

// fileWrite demonstrates three different ways to write to a file:
// 1. file.WriteString
// 2. file.Write (byte slice)
// 3. fmt.Fprintf (formatted string)
func fileWrite(fileName string) {

	// In os.OpenFile, the second parameter specifies the file access mode using bitwise flags.
	// You can combine multiple flags using the bitwise OR operator (|).
	//
	// Available flags:
	//
	// os.O_RDONLY: Open the file in read-only mode (default).
	// os.O_WRONLY: Open the file in write-only mode.
	// os.O_RDWR  : Open the file in both read and write mode.
	// os.O_APPEND: Append data to the end of the file on each write.
	// os.O_CREATE: Create the file if it does not exist.
	// os.O_TRUNC : Truncate (clear) the file's content if it already exists.
	// os.O_EXCL  : Used with O_CREATE — fail if the file already exists.

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Printf("File: \"%s\" failed to open. Error: %v\n", fileName, err)
		return
	}

	defer file.Close()
	log.Printf("File: \"%s\" is opened!", fileName)

	// First: file.WriteString
	file.WriteString("Using function: file.WriteString to write data.\n")

	// Second: file.Write by byte
	data := []byte("Write data by Byte.\n")
	file.Write(data)

	// Third: Write by fmt.Fprintf
	fmt.Fprintf(file, "This is the file: %s\n", fileName)
	log.Printf("Success write data to file: %s.", fileName)
}

// fileWriteRow writes multiple lines to a file using bufio.Writer for buffered output.
func fileWriteRow(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Printf("File: \"%s\" failed to open! %v", fileName, err)
		return
	}

	defer file.Close()

	var values []string = []string{"Here", "is", "SeeChen", "Lee"}
	writer := bufio.NewWriter(file)

	for i, v := range values {
		log.Printf("Writing index %d: %s\n", i, v)
		fmt.Fprintf(writer, "Index %2d: %10s\n", i, v)
	}

	// Must flush buffered data to disk
	writer.Flush()
	log.Printf("Successfully wrote rows to file: %q\n", fileName)
}

// fileWriteOnce writes a complete string to a file in one operation using os.WriteFile.
func fileWriteOnce(fileName string) {
	var content string = `
Hi, SeeChen here.
This is my Golang self-learning document.
Each .go file contains my Golang learning notes and related code.
`

	// os.WriteFile is a convenience function (since Go 1.16)
	// It creates the file if not exists and truncates if exists.
	err := os.WriteFile(fileName, []byte(content), 0777)
	if err != nil {
		log.Printf("Failed to Write content to file: \"%s\". %v", fileName, err)
		return
	}

	log.Printf("Successfully wrote entire content to file: %q\n", fileName)
}

// fileWriteAppend appends data to the end of a file without truncating existing content.
func fileWriteAppend(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open file for append: %q\n", fileName)
		return
	}
	defer file.Close()

	// n represents the number of bytes written to the file.
	n, err := file.WriteString("\nThis line is appended by fileWriteAppend()")
	if err != nil {
		log.Printf("Failed to append content to file %q: %v\n", fileName, err)
		return
	}
	log.Printf("Successfully appended %d bytes to file: %q\n", n, fileName)
}

// fileRead reads a file line by line using bufio.Scanner (efficient for large files)
func fileRead(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Failed to open file \"%s\". Error: %v.", fileName, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // Read the complete contents of the file.
	log.Printf("File \"%s\" is Reading.", fileName)
	for scanner.Scan() {
		// This method will cause the content to be read line by line.
		log.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		log.Printf("Error to reading file: %v.", err)
	}
}

// fileReadAll reads the entire content of a file at once.
// Suitable for small files; avoid for large files due to memory usage.
func fileReadAll(fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("File \"%s\" failed to open! Error: %v.", fileName, err)
		return
	}

	var output string = fmt.Sprintf("%s.\n", content)
	log.Println(output)
}

// fileInfo retrieves and logs metadata about a file.
func fileInfo(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		// Check if the file exists.
		log.Printf("File \"%s\" does not exists.\n", fileName)
	}

	if err != nil {
		// If exitst but still failed get information.
		log.Printf("Failed to get file \"%s\" information.\n", fileName)
		return
	}

	log.Printf("%-15s: %s.\n", "File Name", fileInfo.Name())
	log.Printf("%-15s: %d Byte.\n", "File Size", fileInfo.Size())
	log.Printf("%-15s: %s.\n", "File Auth", fileInfo.Mode())
	log.Printf("%-15s: %s.\n", "Last Modified", fileInfo.ModTime())
	log.Printf("%-15s: %v.\n", "Is Direc", fileInfo.IsDir())
}

// directoryCreate creates a single directory and a nested directory using Mkdir and MkdirAll.
func directoryCreate(firstPath, secondPath string) {

	// Create single-level directory
	err := os.Mkdir(firstPath, 0777)
	if err != nil {
		log.Printf("Directory \"%s\" failed to create. Error: %v.", firstPath, err)
		return
	} else {
		log.Printf("Directory \"%s\" Created successfully.", firstPath)
	}

	// Create multi-level directory
	filePath := filepath.Join(firstPath, secondPath)
	err = os.MkdirAll(filePath, 0777)
	if err != nil {
		log.Printf("Directory \"%s\" failed to create. Error: %v.", filePath, err)
		return
	} else {
		log.Printf("Directory \"%s\" Created successfully.", filePath)
	}
}

// directoryRead lists entries in the current directory using os.ReadDir.
func directoryRead() {
	entries, err := os.ReadDir(".") // Read for this level dir
	if err != nil {
		log.Printf("Error to read directory. Error: %v.", err)
		return
	}

	// The entries is a list type
	for _, entry := range entries {
		info, _ := entry.Info() // ignoring error for simplicity
		log.Printf("%-20s: [Size] %08d, [Last Modifited] %v.\n", entry.Name(), info.Size(), info.ModTime())
	}
}

// fileAdvanced demonstrates file copy and file move (rename).
func fileAdvanced(fileName string, targetName string, firstDir string, secondDir string) {

	//  ====== File Copy ======
	srcFile, err := os.Open(fileName)
	if err != nil {
		log.Printf("Failed to Open File: %s. Error: %v.", fileName, err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(targetName)
	if err != nil {
		log.Printf("Failed to Create File: %s. Error: %v.", fileName, err)
		return
	}
	defer dstFile.Close()

	bytesCopied, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Printf("Failed to copy content from \"%s\" to \"%s\". Error: %v.\n", fileName, targetName, err)
		return
	}
	log.Printf("Successfully Copy content from \"%s\" to \"%s\". Total %d byte(s).\n", fileName, targetName, bytesCopied)

	//  ====== File Move and Rename ======
	firstPath := filepath.Join(".", firstDir)
	secondPath := filepath.Join(firstPath, secondDir)

	_ = os.Rename(targetName, filepath.Join(secondPath, targetName))
	_ = os.Rename(filepath.Join(secondPath, targetName), filepath.Join(secondPath, fileName))
}

// fileTemp demonstrates how to create and work with temporary files and directories in Go.
// It creates a temporary file, writes data to it, then creates a temporary directory and writes a file within it.
// All temporary resources are cleaned up automatically at the end of the function.
func fileTemp() {
	// 1. Create a temporary file in the system temp directory.
	// The first argument "" means use the default OS temp directory.
	// The pattern "example-*.txt" means the file will have a random suffix.
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		log.Fatalf("Failed to create temp file: %v", err)
	}
	// Ensure the file is deleted after use.
	defer os.Remove(tmpFile.Name())

	log.Printf("Temporary file created: %s", tmpFile.Name())

	// Write content to the temporary file.
	content := "Hello, this is a temp file example in Go.\n"
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		log.Fatalf("Failed to write to temp file: %v", err)
	}

	// Always close the file explicitly to flush data to disk.
	if err := tmpFile.Close(); err != nil {
		log.Fatalf("Failed to close temp file: %v", err)
	}

	// 2. Create a temporary directory in the system temp directory.
	tmpDir, err := os.MkdirTemp("", "example-dir-*")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	// Ensure the directory and its contents are removed after use.
	defer os.RemoveAll(tmpDir)

	log.Printf("Temporary directory created: %s", tmpDir)

	// Create a new file inside the temporary directory.
	newFilePath := filepath.Join(tmpDir, "myfile.txt")
	if err := os.WriteFile(newFilePath, []byte("Data inside temp directory."), 0644); err != nil {
		log.Fatalf("Failed to write file in temp directory: %v", err)
	}

	log.Printf("File created inside temp directory: %s", newFilePath)

	// 3. Display the default system temporary directory path.
	fmt.Println("System temporary directory:", os.TempDir())
}

func main() {
	// File handling is an intergral part of every language
	// Golang has 4 system packages to handle files

	// # 1. Create & Open
	var fileName string = "myFile.txt"
	fileCreate(fileName) // Creating file
	fileOpen(fileName)   // Opening file

	// # 2. Write & Read
	// Write
	fileWrite(fileName)       // Write file
	fileWriteRow(fileName)    // Write file by Row
	fileWriteOnce(fileName)   // Write file once
	fileWriteAppend(fileName) // Append content to an already existing file with content.

	// Read
	fileRead(fileName)    // Read a file
	fileReadAll(fileName) // Read file all content at Once

	// File Information
	fileInfo(fileName)

	// Directory Operation
	var firstDir, secondDir string = "File", "Handling"
	directoryCreate(firstDir, secondDir)
	directoryRead()

	log.Println()
	// File - Copy, Rename and Move
	fileAdvanced(fileName, "copied.txt", firstDir, secondDir)

	// File and directory Delete
	// Remove directory
	_ = os.RemoveAll(firstDir)

	// Remove File
	_ = os.Remove(fileName)

	// Temp file in Golang
	fileTemp()
}
