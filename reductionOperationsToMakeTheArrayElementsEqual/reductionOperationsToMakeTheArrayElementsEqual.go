package main

import (
	"fmt"
	"slices"
)

func reductionOperations(nums []int) int {
	slices.Sort(nums)

	if nums[len(nums)-1] == nums[0] { // all elements are equal -> nothing to do
		return 0
	}

	count := 0

	current := nums[len(nums)-1]
	currentCount := 0

	// we go from biggest element to smallest
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == current {
			currentCount++
			continue
		}

		// value changed
		count += currentCount
		current = nums[i]
		currentCount = currentCount + 1 // we added all previous bigger element to next smaller element

		if nums[i] == nums[0] { // we're switched to the smallest element -> no need to iterate anymore, return the current required operations
			return count
		}
	}

	return count
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := reductionOperations(arr)

	fmt.Printf("Required operations: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{5, 1, 3}
	expected := 3

	test(arr, expected)
}

func test2() {
	arr := []int{1, 1, 1} // nothing to decrease
	expected := 0

	test(arr, expected)
}

func test3() {
	arr := []int{1, 1, 2, 2, 3}
	expected := 4

	test(arr, expected)
}

func main() {
	// 1887. Reduction Operations to Make the Array Elements Equal
	test1()
	test2()
	test3()
}
