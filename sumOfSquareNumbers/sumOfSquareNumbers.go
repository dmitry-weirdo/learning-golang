package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	return judgeSquareSum_trivial(c) // O( sqrt(c) )
}

func judgeSquareSum_trivial(c int) bool {
	// a^2 + b^2 = c^2
	// b^2 = c - a^2

	// We iterate A in [0; sqrt(c)].
	// We check whether c - a^2 is a square

	sqrtC := sqrtCeilDown(c)

	if sqrtC*sqrtC == c { // a = sqrt(c), b = 0
		return true
	}

	for a := 0; a <= sqrtC; a++ {
		potentialBSquare := c - a*a

		if isSquare(potentialBSquare) {
			return true
		}
	}

	return false
}

func isSquare(x int) bool {
	sqrt := sqrtCeilDown(x)

	return sqrt*sqrt == x
}

func sqrtCeilDown(c int) int {
	return int(math.Sqrt(float64(c)))
}

func test(x int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", x) // todo: replace with your text if required

	result := judgeSquareSum(x)

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(5, true) // 5 = 1 + 4
}

func test2() {
	test(3, false)
}

func test3() {
	test(1, true) // 1 = 0 + 1
}

func test4() {
	test(0, true) // 0 = 0 + 0
}

func main() {
	// 633. Sum of Square Numbers
	test1()
	test2()
	test3()
	test4()
}
