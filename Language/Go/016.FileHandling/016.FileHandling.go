package main

import (
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

func main() {
	// File handling is an intergral part of every language
	// Golang has 4 system packages to handle files

	// # 1. Create & Open
	var fileName string = "myFile.txt"
	fileCreate(fileName) // Creating file
	fileOpen(fileName)   // Opening file

	// # 2. Write & Read
	fileWrite(fileName) // Write file
}
