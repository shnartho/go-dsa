package main

import (
	"fmt"
	"sync"
)

func printEvenNumbers(evenChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		evenChan <- i
	}
	close(evenChan)
}

func printOddNumbers(oddChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		oddChan <- i
	}
	close(oddChan)
}

func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go printEvenNumbers(evenChan, &wg)
	go printOddNumbers(oddChan, &wg)

	go func() {
		wg.Wait()
		close(oddChan)
		close(evenChan)
	}()

	for even := range evenChan {
		fmt.Printf("Even: %d\n", even)
	}

	for odd := range oddChan {
		fmt.Printf("Odd: %d\n", odd)
	}
}
