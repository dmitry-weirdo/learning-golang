package main

import (
	"fmt"
	"slices"
)

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// From the right, iterate while we find a[i] < a[i + 1]
	// Therefore, all values to the right are descending

	// Then we swap a[i] with the next a[j] > a[i].
	// Of the same a[j] values, select the rightmost.
	// Therefore, the values after a[i] will be now in the descending order.
	// We reverse these values after a[i] to make it increasing.

	n := len(nums)

	i := n - 2

	for i = n - 2; i >= -1; i-- {
		if i < 0 { // this is the max permutation -> reverse the complete array to get the smallest permutation
			slices.Reverse(nums)
			return
		}

		if nums[i] < nums[i+1] {
			break
		}
	}

	// find index J to swap
	j := i + 1
	for j = i + 1; j < n; j++ {
		if nums[j] <= nums[i] { // found the first value that is <= a[i]
			break
		}
	}

	// switch to the last (rightmost) index that is > a[i]
	j--

	// swap I and J
	nums[i], nums[j] = nums[j], nums[i]

	// reverse array after I
	slices.Reverse(nums[i+1:])
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)

	nextPermutation(arr)
	result := arr // function does the permutation in place

	fmt.Printf("Next permutation (smallest if current is the biggest permutation): %v \n", result)
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
	test(
		[]int{1, 2, 3},
		[]int{1, 3, 2},
	)
}

func test2() {
	test(
		[]int{3, 2, 1},
		[]int{1, 2, 3}, // biggest permutation is changed to the smallest
	)
}

func test3() {
	test(
		[]int{1, 1, 5},
		[]int{1, 5, 1},
	)
}

func test4() {
	test(
		[]int{3, 5, 4, 4, 3, 1},
		[]int{4, 1, 3, 3, 4, 5},
	)
}

func test5() {
	test(
		[]int{1, 1},
		[]int{1, 1},
	)
}

func test6() {
	test(
		[]int{1},
		[]int{1},
	)
}

func main() {
	// 31. Next Permutation
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
