package main

import (
	"fmt"
)

func trap(height []int) int {
	n := len(height)

	if n < 3 { // no in-between space for water
		return 0
	}

	// calculate max left heights (including self)
	left := make([]int, n)

	leftMax := height[0]
	for i := 0; i < n; i++ {
		leftMax = max(leftMax, height[i])
		left[i] = leftMax
	}

	// calculate max right heights (including self)
	right := make([]int, n)

	rightMax := height[n-1]
	for i := n - 1; i >= 0; i-- {
		rightMax = max(rightMax, height[i])
		right[i] = rightMax
	}

	// calculate result
	sum := 0

	for i := 0; i < n; i++ {
		// we subtract the current height since we're only putting water at the top of the current height
		// todo: what if this is negative?
		sum += min(left[i], right[i]) - height[i]
	}

	return sum
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := trap(arr)

	fmt.Printf("Max water trapped: %v \n", result)
	fmt.Printf("Expected max water trapped: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expected := 6

	test(arr, expected)
}

func test2() {
	arr := []int{4, 2, 0, 3, 2, 5}
	expected := 9

	test(arr, expected)
}

func test3() {
	arr := []int{3, 0, 2, 0, 4}
	expected := 7

	test(arr, expected)
}

func main() {
	// 42. Trapping Rain Water
	test1()
	test2()
	test3()
}
