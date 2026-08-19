package main

import "fmt"

func missingNumber(nums []int) int {
	result := 0

	for i := range nums { // from 0 to N - 1
		result ^= i
		result ^= nums[i]
	}

	result ^= len(nums) // N value -> index N not present in the array

	return result
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := missingNumber(arr)

	fmt.Printf("Single missing number in range [0..%v]: %v \n", len(arr), result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{3, 0, 1}, 2)
}

func test2() {
	test([]int{0, 1}, 2)
}

func test3() {
	test([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8)
}

func test4() {
	test([]int{1}, 0)
}

func test5() {
	test([]int{0}, 1)
}

func main() {
	// 268. Missing Number

	// Basically the same idea as "136. Single Number".
	// We just make a XOR of [0...N] ourselves, and XOR it with the array.
	test1()
	test2()
	test3()
	test4()
	test5()
}
