package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)

	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			result = append(result, i)
		}
	}

	return result
}

func isSelfDividing(n int) bool {
	if n == 0 {
		return false
	}

	x := n

	digit := 0

	for x != 0 {
		digit = x % 10
		x = x / 10

		if digit == 0 { // cannot divide by 0
			return false
		}

		if n%digit != 0 { // does not divide on digit
			return false
		}
	}

	return true
}

func test(left, right int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Left (inclusive): %v \n", left)
	fmt.Printf("Right (inclusive): %v \n", right)

	result := selfDividingNumbers(left, right)

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	left := 1
	right := 22
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}

	test(left, right, expected)
}

func test2() {
	left := 47
	right := 85
	expected := []int{48, 55, 66, 77}

	test(left, right, expected)
}

func main() {
	// 728. Self Dividing Numbers
	test1()
	test2()
}
