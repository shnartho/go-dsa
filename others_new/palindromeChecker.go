package main

func palindromeChecker(givenString string) bool {
	givenStringLen := len(givenString)
	stack := make([]rune, 0, givenStringLen)
	var i int

	for i = 0; i < givenStringLen/2; i++ {
		stack = append(stack, rune(givenString[i]))
	}

	if givenStringLen%2 == 1 {
		i++
	}

	for ; i < givenStringLen; i++ {
		if rune(givenString[i]) != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return true
}
