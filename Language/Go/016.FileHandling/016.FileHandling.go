package main

import (
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

func main() {
	// File handling is an intergral part of every language
	// Golang has 4 system packages to handle files

	// # 1. Create & Open
	var fileName string = "myFile.txt"
	fileCreate(fileName) // Creating file
	fileOpen(fileName)   // Opening file
}
