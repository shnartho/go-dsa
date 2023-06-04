package main

import (
	"fmt"
)

func main() {
	employees := []Employee{
		{Name: "Nayem", Age: 26, Department: "IT"},
		{Name: "Sourab", Age: 18, Department: "Chemistry"},
		{Name: "Parvez", Age: 23, Department: "IT"},
		{Name: "Riaz", Age: 32, Department: "Chemistry"},
		{Name: "Lutfor", Age: 23, Department: "Chemistry"},
	}

	avrAges := CalculateAverageAge(employees)

	for department, avgAge := range avrAges {
		fmt.Println("Department: ", department, "| Average age: ", avgAge)
	}
}

/*
------ interface&struct -------

func main() {
	var newShape Shape

	newShape = Rectangle{width: 5, height: 4}
	fmt.Println("Area of rectangle: ", newShape.Area())
	fmt.Println("Perimeter of rectangle: ", newShape.Perimeter())

	newShape = Circle{radius: 4}
	fmt.Println("Area of circle: ", newShape.Area())
	fmt.Println("Perimeter of circle: ", newShape.Perimeter())
}

*/
/*
-------- maxSubArraySum --------

func main() {
	testArr := []int{3, 2, -5, -1, 5}
	fmt.Println(MaxSubArraySum(testArr))
}

func main() {
	mystring := "helehq"
	print(palindromeChecker(mystring))

}
*/

/*
--------- palindrome checker -------

func main() {
	mystring := "helehq"
	print(palindromeChecker(mystring))

}
*/
