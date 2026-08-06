package main

import "fmt"

func plusOne(digits []int) []int {
	addToNextDigit := 1 // we add 1 to the last digit
	currentDigit := 0

	for i := len(digits) - 1; i >= 0; i-- {
		currentDigit = digits[i] + addToNextDigit

		addToNextDigit = currentDigit / 10 // will be 1 if + 1 to the next digit
		digits[i] = currentDigit % 10
	}

	if addToNextDigit == 1 {
		// we need to append 1 to the beginning of the array, so the size will be N + 1
		return append([]int{1}, digits...)
	}

	// no +1 digit -> just return the updated digits array
	return digits
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array representing a number: %v \n", arr)

	result := plusOne(arr)

	fmt.Printf("Array for (number + 1): %v \n", result)
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
	arr := []int{1, 2, 3}
	expected := []int{1, 2, 4}

	test(arr, expected)
}

func test2() {
	arr := []int{9}
	expected := []int{1, 0}

	test(arr, expected)
}

func test3() {
	arr := []int{8, 9, 9}
	expected := []int{9, 0, 0}

	test(arr, expected)
}

func main() {
	// 66. Plus One
	test1()
	test2()
	test3()
}
