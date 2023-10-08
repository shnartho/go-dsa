package main

import "fmt"

func factorial(val int) int {
	if val == 1 {
		return 1
	} else {
		return val * factorial(val-1)
	}
}

// 4 * 3 * 2 * 1

func main() {
	result := factorial(4)
	fmt.Println(result)
}
