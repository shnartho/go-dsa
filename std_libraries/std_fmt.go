package main

import (
	"fmt"
	"os"
)

func main() {
	// Println: prints a new line after the message
	fmt.Println("Hello world!")

	// Printf: Formats and prints values using placeholder
	name := "Nayem"
	age := 1
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	// Print: prints without a new line
	fmt.Print("Hello world!")

	//Sprintf: Formats and returns a formatted string
	greetings := fmt.Sprintf("Welcome, %s!", name)
	fmt.Println(greetings)

	//Fprintf: Formats and writes to a specific io.Writer
	fmt.Fprintf(os.Stdout, "This is writter to nil writer. \n")

	// Errorf: creates a formatted error message
	err := fmt.Errorf("An error occured: %s", "something went wrong")
	fmt.Println(err)
}
