package main

import (
	"fmt"
	"strings"
)

func maxSumOfSquares(num int, sum int) string {
	if sum > num*9 { // sum is > than taking all 9 digits
		return ""
	}

	// since (x + y)^2 > x^2 + y^2,
	// we can just select the biggest digits until possible,
	// i.e. it's better than splitting one digit into multiple positions.

	// append 9 until possible
	var sb = strings.Builder{}

	totalNines := sum / 9

	for range totalNines {
		sb.WriteByte('9')
	}

	// append the remaining sum if there is any
	if sum%9 > 0 {
		sb.WriteByte(byte('0' + sum%9))
	}

	for range num - sb.Len() { // append remaining zeroes
		sb.WriteByte('0')
	}

	return sb.String()
}

func test(num, sum int, expectedResult string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number of digits: %v \n", num)
	fmt.Printf("Sum of digits: %v \n", sum)

	result := maxSumOfSquares(num, sum)

	fmt.Printf("Integer with max score: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(2, 3, "30")
}

func test2() {
	test(2, 17, "98")
}

func test3() {
	test(1, 10, "") // impossible to get sum 10 with just 1 digit
}

func main() {
	// 3723. Maximize Sum of Squares of Digits
	test1()
	test2()
	test3()
}
