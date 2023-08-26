/*
Here is the list of golang standard libraries
*/

package main

import (
	"fmt"
)

func main() {
	standardLibraries := []struct {
		Name        string
		Description string
	}{
		{"fmt", "Formatting and printing."},
		{"io", "Input and output utilities."},
		{"os", "Operating system functionality."},
		{"strings", "String manipulation functions."},
		{"strconv", "Conversion functions for strings to other data types."},
		{"math", "Mathematical functions and constants."},
		{"time", "Time-related functionality."},
		{"sort", "Sorting slices and user-defined collections."},
		{"net/http", "HTTP client and server functionality."},
		{"encoding/json", "JSON encoding and decoding."},
		{"flag", "Command-line flag parsing."},
		{"log", "Simple logging package."},
		{"bufio", "Buffered I/O operations."},
		{"sync", "Basic synchronization primitives."},
		{"context", "Context management for goroutines."},
		{"path/filepath", "File path manipulation."},
		{"os/exec", "Running external commands."},
		{"regexp", "Regular expressions."},
		{"testing", "Package for writing tests and benchmarks."},
	}

	for _, lib := range standardLibraries {
		fmt.Printf("Name: %s\nDescription: %s\n\n", lib.Name, lib.Description)
	}
}
