package main

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
