package main

import "fmt"

func binarySearch(array []int, target int) string {
	if len(array) == 0 {
		return ""
	}

	first := 0
	last := len(array) - 1

	for first <= last {
		midpoint := (first + last) / 2
		if array[midpoint] == target {
			return fmt.Sprintf("%d", midpoint)
		}
		if array[midpoint] < target {
			first = midpoint + 1
		} else {
			last = midpoint - 1
		}
	}

	return ""

}

func main() {
	list := []int{2, 3, 4, 6, 7, 8}
	result := binarySearch(list, 8)
	if result != "" {
		fmt.Println("The index of the target value is " + result)
	} else {
		fmt.Println("The target value does not exist in the list")
	}
}
