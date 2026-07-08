package main

import "fmt"

func removeDuplicates(nums []int) int {
	writePos := 0
	current := nums[0]

	i := 0

	for i < len(nums) {
		// update current to the next different value
		current = nums[i]

		//fmt.Printf("\n")
		//fmt.Printf("i: %v, current: %v, a[i] = %v, writePos = %v \n", i, current, nums[i], writePos)

		// write unique value to the beginning of the array
		nums[writePos] = current
		writePos++

		for (i < len(nums)) && (nums[i] == current) {
			// skip all repeats of the current value
			i++
		}
	}

	return writePos // 0-based, and we need to return the count
}

func removeDuplicatesNaive(nums []int) int {
	writePos := 0
	current := nums[0]

	i := 0

	for i < len(nums) {
		for (i < len(nums)) && (nums[i] == current) {
			// skip all repeats of the current value
			i++
		}

		// write unique value to the beginning of the array
		nums[writePos] = current

		// todo: to avoid this internal check, we can more the current update to the start of the cycle (see the version of the funciton above)
		if i < len(nums) {
			current = nums[i]
			writePos++
		}
	}

	return writePos + 1 // 0-based, and we need to return the count
}

func test(nums []int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", nums)

	k := removeDuplicates(nums)

	fmt.Printf("Total distinct values in array: %v \n", k)
	fmt.Printf("Array after removing duplicates: %v \n", nums)
	fmt.Printf("First %v values without duplicates: %v \n", k, nums[0:k])

	if k != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, k)
	}
}

func test1() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := 5

	test(arr, expected)
}

func test2() {
	arr := []int{1, 1, 2}
	expected := 2

	test(arr, expected)
}

func test3() {
	arr := []int{1}
	expected := 1

	test(arr, expected)
}

func test4() {
	arr := []int{1, 2, 3}
	expected := 3

	test(arr, expected)
}

func test5() {
	arr := []int{1, 1, 2, 3, 4}
	expected := 4

	test(arr, expected)
}

func main() {
	// 26. Remove Duplicates from Sorted Array
	test1()
	test2()
	test3()
	test4()
	test5()
}
