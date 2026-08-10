package main

import "fmt"

func isUgly(n int) bool {
	if n == 0 { // corner-case - nothing to divide for 0.
		// 0 has all the possible numbers as divisors, so we return false
		return false
	}

	// same logic was used in "3348. Smallest Divisible Digit Product II"
	temp := n

	for i := 2; i <= 5; i++ {
		for temp%i == 0 {
			temp /= i
		}
	}

	return temp == 1
}

func test(x int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x)

	result := isUgly(x)

	fmt.Printf("Is ugly (only have 2, 3, 5 as prime divisors): %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(0, false) // 0 divides to any prime number
}

func test2() {
	test(6, true) // 6 = 2 * 3
}

func test3() {
	test(1, true) // 1 has no prime factors
}

func test4() {
	test(14, false) // 14 = 2 * 7
}

func main() {
	// 263. Ugly Number
	test1()
	test2()
	test3()
	test4()
}
