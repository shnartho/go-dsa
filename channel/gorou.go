package main

import (
	"fmt"
	"sync"
)

func generateNumbers(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i
	}
}

func calculateSquares(ch <-chan int, nch chan<- int, wg *sync.WaitGroup) {
	defer close(nch)
	defer wg.Done()

	for num := range ch {
		square := num * num
		nch <- square
	}
}

func main() {
	var wg sync.WaitGroup

	numChannel := make(chan int)
	squareChannel := make(chan int)

	wg.Add(2)
	go generateNumbers(numChannel, &wg)
	go calculateSquares(numChannel, squareChannel, &wg)

	// Wait for both goroutines to complete
	wg.Wait()

	// Close the squareChannel after all calculations are done
	close(squareChannel)

	for num := range numChannel {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	for square := range squareChannel {
		fmt.Printf("%d ", square)
	}
	fmt.Println()
}
