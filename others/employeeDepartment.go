package main

import "fmt"

type Employee struct {
	Name       string
	Age        int
	Department string
}

func CalculateAverageAge(employees []Employee) map[string]float64 {
	deparmentAverageAge := make(map[string]float64)
	departmentAges := make(map[string][]int)

	for _, emp := range employees {
		departmentAges[emp.Department] = append(departmentAges[emp.Department], emp.Age)
	}

	for department, ages := range departmentAges {
		total := 0
		for _, age := range ages {
			total += age
		}
		deparmentAverageAge[department] = float64(total) / float64(len(ages))
	}

	return deparmentAverageAge
}

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
