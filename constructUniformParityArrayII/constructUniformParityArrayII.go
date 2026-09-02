package main

import "fmt"

func uniformArray(nums1 []int) bool {
	// O(n), passes in 0-4 ms.
	// I don't think it's actually an optimization, probably just non-stable LeetCode environment speed.
	return uniformArray_optimized(nums1)

	// O(n), passes in 2-10 ms. Probably too many comparisons
	//return uniformArray_trivial(nums1)
}

func uniformArray_optimized(nums1 []int) bool {
	// We don't need minEven, we just need a presence of any even element
	// If both odd and even are present -> we need the min element to be odd

	const GREATER_THAN_MAX = 1_000_000_001 // values are up to 10_9

	bothEvenAndEndPresent := false
	fistElementMod2 := nums1[0] % 2

	minElement := GREATER_THAN_MAX

	for _, v := range nums1 {
		if v%2 != fistElementMod2 {
			bothEvenAndEndPresent = true
		}

		minElement = min(minElement, v)
	}

	if !bothEvenAndEndPresent { // all elements are odd or even -> nothing to change
		return true
	}

	// min element must be odd
	return minElement%2 == 1
}

func uniformArray_trivial(nums1 []int) bool {
	// If the array is already all odd or all even -> return true
	// Odd -> even: for the minOdd, we don't have a smaller odd element -> impossible
	// Even -> odd: for the minEven, we need a smaller odd element. I.e. the smallest element must be even.

	const GREATER_THAN_MAX = 1_000_000_001 // values are up to 10_9

	minEven := GREATER_THAN_MAX
	minOdd := GREATER_THAN_MAX

	for _, v := range nums1 {
		if v%2 == 0 {
			minEven = min(minEven, v)
		} else {
			minOdd = min(minOdd, v)
		}
	}

	if (minEven == GREATER_THAN_MAX) || (minOdd == GREATER_THAN_MAX) { // all elements are odd or even -> nothing to change
		return true
	}

	return minOdd < minEven
}

func test(arr []int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := uniformArray(arr)

	fmt.Printf("Possible to change value to an array with same parity elements: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{1, 4, 7}, true) // min element is odd
}

func test2() {
	test([]int{2, 3}, false) // min element is even
}

func test3() {
	test([]int{4, 6}, true) // only even parity already
}

func test4() {
	test([]int{1, 1}, true) // only odd parity already
}

func test5() {
	test([]int{2, 2}, true) // only even parity already
}

func main() {
	// 3876. Construct Uniform Parity Array II
	test1()
	test2()
	test3()
	test4()
	test5()
}
