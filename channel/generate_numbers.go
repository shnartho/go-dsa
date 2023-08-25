/*
Problem: Goroutine Communication with Channels
Write a Go program that creates a channel, launches a goroutine to generate a sequence of numbers (e.g., 1, 2, 3, ...), and sends these numbers to the main goroutine through the channel. The main goroutine should print the received numbers.
*/

package main

import "fmt"

func generateNumbers(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	numChannel := make(chan int)
	go generateNumbers(numChannel)

	for num := range numChannel {
		fmt.Printf("%d", num)
	}
}
