package main

import "fmt"

func main() {
	var newShape Shape

	newShape = Rectangle{width: 5, height: 4}
	fmt.Println("Area of rectangle: ", newShape.Area())
	fmt.Println("Perimeter of rectangle: ", newShape.Perimeter())

	newShape = Circle{radius: 4}
	fmt.Println("Area of circle: ", newShape.Area())
	fmt.Println("Perimeter of circle: ", newShape.Perimeter())
}

// func main() {
// 	testArr := []int{3, 2, -5, -1, 5}
// 	fmt.Println(MaxSubArraySum(testArr))
// }

// func main() {
// 	mystring := "helehq"
// 	print(palindromeChecker(mystring))

// }
