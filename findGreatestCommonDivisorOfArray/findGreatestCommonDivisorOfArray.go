package main

import "fmt"

func findGCD(nums []int) int {
	minValue, maxValue := minAndMaxInArray(nums)

	return gcd(minValue, maxValue)
}

func minAndMaxInArray(arr []int) (minValue, maxValue int) {
	// we assume the array is non-empty
	minValue = arr[0]
	maxValue = arr[0]

	for _, v := range arr {
		minValue = min(minValue, v)
		maxValue = max(maxValue, v)
	}

	return minValue, maxValue
}

func gcd(a, b int) int { // greatest common divisor
	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := findGCD(arr)

	fmt.Printf("GCD of min and max elements in the array: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{2, 5, 6, 9, 10}, 2)
}

func test2() {
	test([]int{7, 5, 6, 8, 3}, 1)
}

func test3() {
	test([]int{3, 3}, 3)
}

func main() {
	// 1979. Find Greatest Common Divisor of Array
	test1()
	test2()
	test3()
}
