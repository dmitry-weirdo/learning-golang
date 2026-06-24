package main

import "fmt"

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	l := len(s)

	result := 0

	for i, ch := range s {
		char := byte(ch)
		number := m[char]

		fmt.Printf("Char: %c, number %v \n", char, number)

		if i >= (l - 1) { // last digit - always add
			fmt.Printf("Char %c is last char. Adding %v to result \n", char, number)
			result += number

			fmt.Printf("Result: %v \n", result)
			continue
		}

		nextChar := s[i+1]
		nextNumber := m[nextChar]

		if number < nextNumber {
			fmt.Printf("Current number %c(%v) is smaller than the next number %c(%v). Subtracting the current number %d. \n", char, number, nextChar, nextNumber, number)

			result -= number
			fmt.Printf("Result: %v \n", result)
		} else {
			fmt.Printf("Current number %c(%v) is NOT smaller than the next number %c(%v). Adding the current number %d. \n", char, number, nextChar, nextNumber, number)

			result += number
			fmt.Printf("Result: %v \n", result)
		}
	}

	return result
}

func test(s string, expected int) {
	fmt.Println()
	fmt.Println("===============")

	fmt.Printf("Roman number string: %v \n", s)

	result := romanToInt(s)

	fmt.Printf("Expected number: %v \n", expected)
	fmt.Printf("Actual number: %v \n", result)

}

func test1() {
	s := "III"
	expected := 3

	test(s, expected)
}

func test2() {
	s := "LVIII"
	expected := 58

	test(s, expected)
}

func test3() {
	s := "MCMXCIV"
	expected := 1994

	test(s, expected)
}

func main() {
	test1()
	test2()
	test3()
}
