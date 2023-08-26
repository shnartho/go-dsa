/*
Problem: Collaborative Goroutines
Write a Go program that utilizes two goroutines. The first goroutine generates a sequence of numbers (e.g., 1, 2, 3, ...) and sends them to a channel. The second goroutine receives the numbers from the channel and calculates their squares. The main goroutine prints both the original number and its square.
*/

package main

import (
	"fmt"
)

func generateNumbers(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func calculateSquares(ch <-chan int, nch chan<- int) {
	for num := range ch {
		square := num * num
		nch <- square
	}
	close(nch)
}

func main() {
	numChannel := make(chan int)
	squareChannel := make(chan int)

	go generateNumbers(numChannel)
	go calculateSquares(numChannel, squareChannel)

	devider := 1
	for square := range squareChannel {
		fmt.Printf("Number: %d | Square: %d \n", square/devider, square)
		devider += 1
	}
	fmt.Println()
}
