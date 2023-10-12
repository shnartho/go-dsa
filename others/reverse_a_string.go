package main

import "fmt"

func reverseAString(s string) string {
	var string2 string
	for i := len(s) - 1; i >= 0; i-- {
		string2 += s[i : i+1]
	}
	return string2
}

func main() {
	my_string := "HelloWorld"
	reversedString := reverseAString(my_string)
	fmt.Println(reversedString)
}
