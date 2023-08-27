package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	// Synchronization using sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(2) // We'll wait for 2 goroutines

	go func() {
		defer wg.Done() // Decrement the couter when done
		fmt.Println("Goroutine 1: Started!")
		// simulate some work
		for i := 1; i <= 5; i++ {
			fmt.Println("Goroutine 1:", i)
		}
		fmt.Println("Goroutine 1: Done")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2: started")
		for i := 6; i <= 10; i++ {
			fmt.Println("Goroutine 2:", i)
		}
		fmt.Println("Goroutine 2: Done")
	}()

	wg.Wait() // Wait for both goroutines to finish

	// buffered io using bufio
	fileName := "output.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() // Flush the buffer before closeing

	message := "This is a line written using bufio."
	_, writeErr := writer.WriteString(message)
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
		return
	}
	fmt.Println("Message written to", fileName, "using bufio.")
}
