package main

import "fmt"

func minimumDeletions(nums []int) int {
	minValue := nums[0]
	minIndex := 0

	maxValue := nums[0]
	maxIndex := 0

	// we're guaranteed that values are distinct -> no check of duplicate values
	for i, v := range nums {
		if v < minValue {
			minValue = v
			minIndex = i
		}

		if v > maxValue {
			maxValue = v
			maxIndex = i
		}
	}

	left, right := getMinAndMax(minIndex, maxIndex)

	//fmt.Printf("min value: %v, min index: %v \n", minValue, minIndex)
	//fmt.Printf("max value: %v, max index: %v \n", maxValue, maxIndex)

	n := len(nums)
	removeFromLeft := right + 1                     // remove both values from left
	removeFromRight := n - left                     // remove both values from right
	removeFromBothSides := (left + 1) + (n - right) // remove from left side and right side

	return min(removeFromBothSides, min(removeFromLeft, removeFromRight))
}

func getMinAndMax(a, b int) (smaller, greater int) {
	if a <= b {
		return a, b
	}

	return b, a
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := minimumDeletions(arr)

	fmt.Printf("Minimum number of deletions to delete min and max elements: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{2, 10, 7, 5, 4, 1, 8, 6}, 5)
}

func test2() {
	test([]int{0, -4, 19, 1, 8, -2, -3, 5}, 3)
}

func test3() {
	test([]int{101}, 1)
}

func main() {
	// 2091. Removing Minimum and Maximum From Array
	test1()
	test2()
	test3()
}
