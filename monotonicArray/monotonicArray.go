package main

import (
	"fmt"
)

func isMonotonic(nums []int) bool {
	n := len(nums)

	i := 0
	increasing := false
	decreasing := false

	for i < n-1 {
		if nums[i] < nums[i+1] {
			increasing = true
		} else if nums[i] > nums[i+1] {
			decreasing = true
		}

		if increasing && decreasing {
			return false
		}

		i++
	}

	return true
}

func isMonotonicNaive(nums []int) bool {
	n := len(nums)
	if n < 2 { // array of 1 element is always monotonic
		return true
	}

	// define whether array is increasing or decreasing
	i := 0

	// skip the equal starting values
	for (i < (n - 1)) && (nums[i] == nums[i+1]) {
		i++
	}

	if i >= (n - 1) {
		// all values in array are the same -> it's monotonic
		fmt.Printf("All values in the array are the same = %v. Returning true. \n", nums[0])
		return true
	}

	increasing := nums[i] < nums[i+1]

	if increasing { // rest must be increasing
		fmt.Printf("a[%v] = %v < a[%v] = %v. Array must be increasing. \n", i, nums[i], i+1, nums[i+1])

		for i < (n - 1) {
			if nums[i] > nums[i+1] {
				return false
			}

			i++
		}
	} else { // rest must be decreasing
		fmt.Printf("a[%v] = %v > a[%v] = %v. Array must be decreasing. \n", i, nums[i], i+1, nums[i+1])

		for i < (n - 1) {
			if nums[i] < nums[i+1] {
				return false
			}

			i++
		}
	}

	return true
}

func test(a []int, expectedResult bool) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("array: %v \n", a)

	result := isMonotonic(a)

	fmt.Printf("Is monotonic: %v  \n", result)
	fmt.Printf("Expected result: %v  \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	a := []int{1, 2, 2, 3}
	expected := true

	test(a, expected)
}

func test2() {
	a := []int{1, 1, 1}
	expected := true

	test(a, expected)
}

func test3() {
	a := []int{6, 5, 4, 4}
	expected := true

	test(a, expected)
}

func test4() {
	a := []int{1, 3, 2}
	expected := false

	test(a, expected)
}

func test5() {
	a := []int{1, 3}
	expected := true

	test(a, expected)
}

func test6() {
	a := []int{1}
	expected := true

	test(a, expected)
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
