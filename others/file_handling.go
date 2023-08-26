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

	// append in file
	appendContent := []byte("\nAppending more content in the same file")
	file, appendErr := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if appendErr != nil {
		fmt.Println("Error opening file:", appendErr)
		return
	}
	defer file.Close()

	_, writeErr := file.Write(appendContent)
	if writeErr != nil {
		fmt.Println("Error appending to file", writeErr)
		return
	}
	fmt.Printf("Successfully append %s to %s", string(appendContent), filename)

	// Read the contents after appending
	appendedContent, _ := os.ReadFile(filename)
	fmt.Println("Content: ", string(appendedContent))

	// openfile manually and close
	openFile, _ := os.Open(filename)
	defer file.Close()

	// read from open file
	buf := make([]byte, 200)
	n, _ := openFile.Read(buf)
	fmt.Println("Read from the opened file: ", string(buf[:n]))

}
