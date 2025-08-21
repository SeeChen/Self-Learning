package main

import (
	"log"
	"os"
)

func createFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // defer close file ensures that the file will be closed
	log.Printf("File \"%s\" Created!\n", fileName)
}

func main() {
	// File handling is an intergral part of every language
	// Golang has 4 system packages to handle files

	// 1. os
	// To create a file, use the os.Create function, which returns a file object of type *os.File
	// After creating a file, must call the close function to close the file to release system resources.
	createFile("myFile.txt")
}
