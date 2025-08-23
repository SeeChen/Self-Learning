package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// To create a file, use the os.Create function, which returns a file object of type *os.File.
// After creating a file, must call the close function to close the file to release system resources.
func fileCreate(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close() // defer close file ensures that the file will be closed
	log.Printf("File: \"%s\" Created!\n", fileName)
}

// To open a file, use os.Open function.
// Same as the create function, os.Open will return an *os.File file object and must call close function to release system resources.
func fileOpen(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("File: \"%s\" failed to open. Error: %v\n", fileName, err)
		return
	}
	defer file.Close()

	log.Printf("File: \"%s\" Opened!\n", fileName)
}

// This demo shows 3 ways to write data to a file
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

// This function will write lines to the file.
func fileWriteRow(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Printf("File: \"%s\" failed to open! %v", fileName, err)
		return
	}

	defer file.Close()

	var arr []string = []string{"Here", "is", "SeeChen", "Lee"}
	writer := bufio.NewWriter(file)

	for i, value := range arr {
		log.Printf("Writing index %d: %s.\n", i, value)

		fmt.Fprintf(writer, "Index %2d: %10s.\n", i, value)
		writer.Flush()
	}

	log.Printf("Successs write content to file: %s.\n", fileName)
}

// This function will demonstrate how to insert the entire content into a file at once.
func fileWriteOnce(fileName string) {
	var content string = `
	Hi, SeeChen here.
	This is my Golang self-learning document.
	Each .go file contains my Golang learning notes and related code.
	`

	err := os.WriteFile(fileName, []byte(content), 0777)
	if err != nil {
		log.Printf("Failed to Write content to file: \"%s\". %v", fileName, err)
		return
	}

	log.Printf("Success write content to file: \"%s\".\n", fileName)
}

// This function shows how to append content or data to a file that already exists and has content.
func fileWriteAppend(fileName string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Printf("File: \"%s\" failed to open.\n", fileName)
		return
	}

	defer file.Close()

	// n represents the number of bytes written to the file.
	n, err := file.WriteString("\nThis row is append by function: fileWriteAppend()")
	if err != nil {
		log.Printf("Failed to append content to file: \"%s\".\n", fileName)
		return
	}

	log.Printf("Successfully append %d byte content to file: \"%s\".\n", n, fileName)
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
}
