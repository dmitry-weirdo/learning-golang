package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		// we need to swap all [1...n] values that are in the i-th position and not yet in its expected position
		// it will be still O(n) since every value in range will be put into its position once
		for inRange(n, nums[i]) &&
			!inPosition(nums, i) { // if value is already in position -> stop swapping (this may happen in case of duplicates)
			// put current value into its position
			swap(nums, i, nums[i]-1)
		}
	}

	fmt.Printf("Array after hash-swapping: \n%v \n", nums)

	// check for the position
	for i, v := range nums {
		expectedValue := i + 1

		if v != expectedValue { // first position with incorrect value -> return this expected value
			fmt.Printf("arr[%v] = %v is not equal to its expected value = %v. Returning %v. \n", i, v, expectedValue, expectedValue)

			return expectedValue
		}
	}

	// all numbers in place -> return (n + 1)
	fmt.Printf("All array values [1..%v] are in their correct positions. Returning n + 1 = %v. \n", n, n+1)

	return n + 1
}

func inRange(n, value int) bool {
	return (1 <= value) && (value <= n)
}

func inPosition(arr []int, i int) bool {
	value := arr[i]
	expectedPosition := arr[i] - 1 // value 1 must be in position 0

	return arr[expectedPosition] == value
}

func swap(a []int, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func test(arr []int, expected int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr)

	result := firstMissingPositive(arr)

	fmt.Printf("First missing positive: %d \n", result)
	fmt.Printf("Expected value: %d \n", expected)

	if result != expected {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expected, result)
	}
}

func test1() {
	arr := []int{5, 4, 3, 2, 1}
	expected := 6

	test(arr, expected)
}

func test2() {
	arr := []int{1, 2, 0}
	expected := 3

	test(arr, expected)
}

func test3() {
	arr := []int{3, 4, -1, 1}
	expected := 2

	test(arr, expected)
}

func test4() {
	arr := []int{7, 8, 9, 11, 12}
	expected := 1

	test(arr, expected)
}

func test5() {
	arr := []int{1, 2, 1, 2, 4, 2}
	expected := 3

	test(arr, expected)
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
}
