package main

import "fmt"

func searchInsert(nums []int, target int) int {
	// Using a template from:
	// https://leetcode.com/discuss/post/786126/python-powerful-ultimate-binary-search-t-rwv8/

	left := 0
	right := len(nums) // it can be after the end of the array // !!! this will fail the test if we set len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] >= target { // target condition
			right = mid // in this template it is always mid, NOT mid - 1
		} else {
			left = mid + 1
		}
	}

	// after exiting the while loop, left is the minimal k satisfying the condition function;
	return left
}

func searchInsert_working(nums []int, target int) int {
	// we need to find the first index where nums[i] >= target

	left := 0
	right := len(nums) - 1 // todo: there is a magic of setting this to len(nums)?

	// if no (nums[i] >= target) were found, return index after the end of the array
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] < target { // no match, move to the right
			// this will move to len(nums) if last value in the array is < target
			left = mid + 1
		} else { // match found -> save the mid index as a result, move to the left
			right = mid - 1
			//result = mid
		}
	}

	return left
}

func searchInsertNaive(nums []int, target int) int {
	// we need to find the first index where nums[i] >= target

	left := 0
	right := len(nums) - 1 // todo: there is a magic of setting this to len(nums)

	// if no (nums[i] >= target) were found, return index after the end of the array
	// this is a safe version, not relying on a heuristic that left will move correctly to the result
	result := len(nums)

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] < target { // no match, move to the right
			left = mid + 1
		} else { // match found -> save the mid index as a result, move to the left
			right = mid - 1
			result = mid
		}
	}

	return result
}

func test(nums []int, target int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", nums)
	fmt.Printf("Target value: %v \n", target)

	insertPos := searchInsert(nums, target)

	fmt.Printf("Insert index for value %v: %v \n", target, insertPos)
	fmt.Printf("Expected insert index for value %v: %v \n", target, expectedResult)

	if insertPos != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, insertPos)
	}
}

func test1() {
	arr := []int{1, 3, 5, 6}
	target := 5
	expected := 2

	test(arr, target, expected)
}

func test2() {
	arr := []int{1, 3, 5, 6}
	target := 2
	expected := 1

	test(arr, target, expected)
}

func test3() {
	arr := []int{1, 3, 5, 6}
	target := 7
	expected := 4 // after the end of the array

	test(arr, target, expected)
}

func test4() {
	arr := []int{1, 3, 5} // test the odd-length array
	target := 7
	expected := 3 // after the end of the array

	test(arr, target, expected)
}

func test5() {
	arr := []int{1, 3, 5} // test the odd-length array
	target := 0
	expected := 0 // 0-th position

	test(arr, target, expected)
}

func test6() {
	arr := []int{1, 3, 5} // test the odd-length array
	target := 3
	expected := 1 // exact target found

	test(arr, target, expected)
}

func main() {
	// 35. Search Insert Position
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
