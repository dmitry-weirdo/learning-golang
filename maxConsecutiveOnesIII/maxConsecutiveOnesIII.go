package main

import "fmt"

func longestOnes(nums []int, k int) int {
	maxWindowLen := 0 // if k = 0 and all elements are 0, we can have 0 result

	left := 0
	zeroesCount := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroesCount++
		}

		// todo: we can shrink left just by 1, so that we keep it at max size
		// but I find this logic more straightforward and comprehensive
		for zeroesCount > k { // move left until we have a zeroes-valid window
			if nums[left] == 0 {
				zeroesCount--
			}

			left++
		}

		maxWindowLen = max(maxWindowLen, right-left+1)
	}

	return maxWindowLen
}

func test(arr []int, k int, expectedResult int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("K - max zeroes to be replaced: %v \n", k)

	result := longestOnes(arr, k)

	fmt.Printf("Max consecutive ones with K = %v 0-replacements: %v \n", k, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	arr := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	expected := 6

	test(arr, k, expected)
}

func test2() {
	arr := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k := 3
	expected := 10

	test(arr, k, expected)
}

func test3() {
	arr := []int{0, 0, 0, 0}
	k := 0
	expected := 0

	test(arr, k, expected)
}

func main() {
	// 1004. Max Consecutive Ones III
	test1()
	test2()
	test3()
}
