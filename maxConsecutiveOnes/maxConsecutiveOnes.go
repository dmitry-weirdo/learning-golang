package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	left := 0

	maxWindowSize := 0

	for right, v := range nums {
		if v != 1 { // met 0 -> dispose of the previous window
			left = right
			continue
		}

		// now we know that a[right] = 1, and we have a 1-window of size >= 1

		for nums[left] != 1 { // shrink the window from left until a[left] reaches the 1
			left++
		}

		windowSize := right - left + 1

		maxWindowSize = max(maxWindowSize, windowSize)
	}

	return maxWindowSize
}

func test(arr []int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)

	result := findMaxConsecutiveOnes(arr)

	fmt.Printf("Max consecutive ones: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	arr := []int{1, 1, 0, 1, 1, 1}
	expected := 3

	test(arr, expected)
}

func test2() {
	arr := []int{1, 0, 1, 1, 0, 1}
	expected := 2

	test(arr, expected)
}

func test3() {
	arr := []int{0, 0, 1, 0, 0}
	expected := 1

	test(arr, expected)
}

func test4() {
	arr := []int{0}
	expected := 0

	test(arr, expected)
}

func test5() {
	arr := []int{}
	expected := 0

	test(arr, expected)
}

func main() {
	// 485. Max Consecutive Ones
	test1()
	test2()
	test3()
	test4()
	test5()
}
