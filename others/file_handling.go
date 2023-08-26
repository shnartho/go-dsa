package main

import (
	"fmt"
	"os"
)

func main() {
	// create a new file
	filename := "example.txt"
	content := []byte("hello, File handling in Go!")
	err := os.WriteFile(filename, content, 0644) // 644 here owner rw and others r file permission
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File created", filename)

	// read content from file
	readContent, readErr := os.ReadFile(filename)
	if readErr != nil {
		fmt.Println("Error reading file:", readErr)
		return
	}
	fmt.Println("File content:", string(readContent))

}
