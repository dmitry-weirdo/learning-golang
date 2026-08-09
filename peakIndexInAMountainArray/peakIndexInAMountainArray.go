package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	// Binary search.
	// We need to return the first (leftmost) element that is > next element

	// we need at least one element to the left of the peak, i.e. minimum value is 1
	left := 1

	// Index must be within the array.
	// We can even set ( len(arr) - 2  ) since there must be at least one element after the peak
	// By setting len(arr) - 2, we avoid the check that a[mid + 1] is out of the range.
	right := len(arr) - 2

	for left < right {
		mid := (left + right) / 2

		if arr[mid] > arr[mid+1] { // target condition
			right = mid // in this template it is always mid, NOT mid - 1
		} else {
			left = mid + 1
		}
	}

	// after exiting the while loop, left is the minimal k satisfying the condition function;
	return left
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Mountain array: %v \n", arr)

	result := peakIndexInMountainArray(arr)

	fmt.Printf("Peak index: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{0, 1, 0}
	expected := 1

	test(arr, expected)
}

func test2() {
	arr := []int{0, 2, 1, 0}
	expected := 1

	test(arr, expected)
}

func test3() {
	arr := []int{0, 10, 5, 2}
	expected := 1

	test(arr, expected)
}

func main() {
	// 852. Peak Index in a Mountain Array
	test1()
	test2()
	test3()
}
