package main

import "fmt"

func largestAltitude(gain []int) int {
	prefixSum := 0
	maxSum := 0

	for _, v := range gain {
		prefixSum += v

		maxSum = max(maxSum, prefixSum)
	}

	return maxSum
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array of altitude changes: %v \n", arr)

	result := largestAltitude(arr)

	fmt.Printf("Max height reached: %v \n", result)
	fmt.Printf("Expected last stone weight: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{-5, 1, 5, 0, -7}
	expected := 1

	test(arr, expected)
}

func test2() {
	arr := []int{-4, -3, -2, -1, 4, 3, 2}
	expected := 0

	test(arr, expected)
}

func main() {
	// 1732. Find the Highest Altitude
	test1()
	test2()
}
