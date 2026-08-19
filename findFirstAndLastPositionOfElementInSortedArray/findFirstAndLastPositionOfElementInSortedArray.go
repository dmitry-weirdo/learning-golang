package main

import "fmt"

func searchRange(nums []int, target int) []int {
	leftmostIndex := searchExactValueLeftmost(nums, target)

	if leftmostIndex < 0 { // element not found in the array -> return -1
		return []int{-1, -1}
	}

	// now we're sure there will be a rightmost index, since the leftmost search above returned a value already
	rightmostIndex := searchExactValueRightmost(nums, target)

	return []int{leftmostIndex, rightmostIndex}
}

func searchExactValueLeftmost(arr []int, target int) int { // returns -1 if element is not found
	if len(arr) < 1 { // empty array -> nothing to search
		return -1
	}

	condition := func(x int) bool {
		//return x == target // this will NOT work, e.g for {1, 1, 2, 3, 3}, target = 1 we will jump right
		return x >= target
	}

	index := binarySearchGeneric(arr, 0, len(arr)-1, condition)

	if arr[index] != target {
		return -1
	}

	return index
}

func searchExactValueRightmost(arr []int, target int) int { // returns -1 if element is not found
	if len(arr) < 1 { // empty array -> nothing to search
		return -1
	}

	condition := func(x int) bool {
		return x > target
	}

	// right is len(arr), so we'll jump over the array (to the insertion point)
	index := binarySearchGeneric(arr, 0, len(arr), condition)

	if index == 0 { // target smaller than the 0-th element of the array
		return -1
	}

	// rightmost is the previous from what we sought for
	index--

	if arr[index] != target {
		return -1
	}

	return index
}

func binarySearchGeneric(
	arr []int,
	left int, // usually it starts with 0, if we search in the complete array
	right int, // set len(arr) - 1 if you want to be within array. Set len(arr) if index after the array can be returned.
	condition func(int) bool, // we will find the leftmost index satisfying this condition within [left; right] range
) int {
	// Using a template from:
	// https://leetcode.com/discuss/post/786126/python-powerful-ultimate-binary-search-t-rwv8/
	for left < right {
		mid := (left + right) / 2

		if condition(arr[mid]) { // target condition
			right = mid // in this template it is always mid, NOT mid - 1
		} else {
			left = mid + 1
		}
	}

	// after exiting the while loop, left is the minimal k satisfying the condition function;
	return left
}

func test(arr []int, target int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Target value: %v \n", target)

	result := searchRange(arr, target)

	fmt.Printf("First and last positions of %v in the array: %v \n", target, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	arr := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := []int{3, 4}

	test(arr, target, expected)
}

func test2() {
	arr := []int{5, 7, 7, 8, 8, 10}
	target := 6
	expected := []int{-1, -1} // element not found

	test(arr, target, expected)
}

func test3() {
	arr := []int{}
	target := 0
	expected := []int{-1, -1} // element not found

	test(arr, target, expected)
}

func test4() {
	arr := []int{1, 2, 3, 3}
	target := 3
	expected := []int{2, 3} // end of the array

	test(arr, target, expected)
}

func test5() {
	arr := []int{1, 1, 2, 3, 3}
	target := 1
	expected := []int{0, 1} // start of the array

	test(arr, target, expected)
}

func test6() {
	arr := []int{1}
	target := 1
	expected := []int{0, 0} // start of the array = end of array

	test(arr, target, expected)
}

func test7() {
	arr := []int{1, 1, 1}
	target := 1
	expected := []int{0, 2} // start of the array - end of array (same element)

	test(arr, target, expected)
}

func test8() {
	arr := []int{1, 2, 3}
	target := 0 // target smaller than the first element
	expected := []int{-1, -1}

	test(arr, target, expected)
}

func test9() {
	arr := []int{1, 2, 3}
	target := 10 // target bigger than the last element
	expected := []int{-1, -1}

	test(arr, target, expected)
}

func main() {
	// 34. Find First and Last Position of Element in Sorted Array
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
}
