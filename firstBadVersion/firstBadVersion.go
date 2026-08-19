package main

import "fmt"

func firstBadVersion(n int) int {
	// isBadVersion is provided by LeetCode environment
	return binarySearchOnRangeGeneric(1, n, isBadVersion) // versions start with 1
}

func binarySearchOnRangeGeneric(
	left int, // left of your range
	right int, // inclusive right of your range
	condition func(int) bool, // we will find the leftmost value satisfying this condition within [left; right] range
) int {
	// Using a template from:
	// https://leetcode.com/discuss/post/786126/python-powerful-ultimate-binary-search-t-rwv8/
	for left < right {
		mid := (left + right) / 2

		// NOT using an array
		if condition(mid) { // target condition
			right = mid // in this template it is always mid, NOT mid - 1
		} else {
			left = mid + 1
		}
	}

	// after exiting the while loop, left is the minimal k satisfying the condition function;
	return left
}

// this value we will override in the tests
var firstBadVersionNumber = -1

// this function will be provided by LeetCode
func isBadVersion(version int) bool {
	return version >= firstBadVersionNumber
}

func test(n int, badVersion int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number: %v \n", n)
	fmt.Printf("First bad version that will be provided by `isBadVersion` function: %v \n", badVersion)

	firstBadVersionNumber = badVersion
	result := firstBadVersion(n)

	fmt.Printf("First bad version: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(5, 4, 4)
}

func test2() {
	test(1, 1, 1)
}

func main() {
	// 278. First Bad Version
	test1()
	test2()
}
