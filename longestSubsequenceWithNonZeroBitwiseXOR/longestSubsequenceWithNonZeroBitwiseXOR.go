package main

import "fmt"

func longestSubsequence(nums []int) int {
	allZeroes := true

	xor := 0

	for _, v := range nums {
		if v != 0 {
			allZeroes = false
		}

		xor ^= v
	}

	if xor != 0 { // complete array already has non-0 xor
		return len(nums)
	}

	if allZeroes { // all elements are 0 -> no option to make it non-0
		return 0
	}

	// array xor is 0 but there is at least on non-0 element -> by excluding this element, we'll change the xor to non-0
	return len(nums) - 1
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := longestSubsequence(arr)

	fmt.Printf("Length of the longest xor != 0 subsequence: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{1, 2, 3}, 2) // 1 + 2 = 3, so array xor is 0, we need to exclude any element
}

func test2() {
	test([]int{2, 3, 4}, 3) // array xor != 0
}

func test3() {
	test([]int{0, 0, 0}, 0) // only 0-s in the array -> no way to make it non-0 xor
}

func main() {
	// 3702. Longest Subsequence With Non-Zero Bitwise XOR
	test1()
	test2()
	test3()
}
