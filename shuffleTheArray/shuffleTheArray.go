package main

import (
	"fmt"
)

func shuffle(nums []int, n int) []int {
	// this is a genius algorithm that replaces the array in place
	// https://leetcode.com/problems/shuffle-the-array/solutions/684649/on-time-o1-space-no-bitwise-cheating-bea-ooon/
	// todo: MAYBE it is Cycle Leader Algorithm (Knuth, The Art of Computer Programming)

	// elements to the right of i are already in place
	// elements to the left of i are used to swap elements
	for i := 2*n - 1; i >= 0; i-- {
		k := i

		// do-while not present in go, so
		if k%2 == 1 {
			k = k/2 + n
		} else {
			k = k / 2
		}

		for k > i {
			if k%2 == 1 {
				k = k/2 + n
			} else {
				k = k / 2
			}
		}

		// swap a[i] and a[k]
		nums[i], nums[k] = nums[k], nums[i]
	}

	return nums
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	n := len(arr) / 2

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("N (half of array length): %v \n", n)

	result := shuffle(arr, n)

	fmt.Printf("Shuffled array: %v \n", result)
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
	arr := []int{2, 5, 1, 3, 4, 7}
	expected := []int{2, 3, 5, 4, 1, 7}

	test(arr, expected)
}

func test2() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 1}
	expected := []int{1, 4, 2, 3, 3, 2, 4, 1}

	test(arr, expected)
}

func test3() {
	arr := []int{1, 1, 2, 2}
	expected := []int{1, 2, 1, 2}

	test(arr, expected)
}

func main() {
	// 1470. Shuffle the Array
	test1()
	test2()
	test3()
}
