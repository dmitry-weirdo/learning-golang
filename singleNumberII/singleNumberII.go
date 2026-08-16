package main

import (
	"fmt"
	"math"
)

func singleNumber(nums []int) int {
	// todo: implement a super-clever solution that is counting all the bits % 3 in just one array iteration

	// we iterate up to 32 + 1 times over the complete array O(n)
	// O(33 * n) is still O(n), but this is not perfect.
	return singleNumber_countingEveryBit(nums)
}

func singleNumber_countingEveryBit(nums []int) int {
	maxSignificantBit := -1

	for _, v := range nums {
		maxSignificantBit = max(maxSignificantBit, getMaxSignificantBit(v))

		fmt.Printf("Number %v = %032b \n", v, uint32(v))
	}

	fmt.Printf("Max significant bit: %v \n", maxSignificantBit)

	numberPresentOnlyOnce := 0

	// count every bit sum for all the numbers
	for bit := range maxSignificantBit + 1 {
		bitCount := 0

		for _, v := range nums {
			// shift the current bit to the right and check whether it's 1
			bitCount += (v >> bit) & 1
		}

		// take % 3 -> remove the bits that are repeated 3 times.
		// i.e. we get the bit of the number that is present only once
		bitCount %= 3

		// add this 0 or 1 bit to the bit's position in the result
		numberPresentOnlyOnce |= bitCount << bit
	}

	//fmt.Printf("Number present only once %v = %032b \n", numberPresentOnlyOnce, uint32(numberPresentOnlyOnce))

	// the trick to make 1st bit in 2^32 NOT overflow as the positive number in int64.
	// We must treat 2^32 = 1 as negative number, i.e. as int32
	return int(int32(numberPresentOnlyOnce))
}

func getMaxSignificantBit(n int) int {
	if n < 0 { // for the negative numbers, 1st bit (32-th bit) is significant and contains 1
		return 31
	}

	return int(math.Log2(float64(n)))
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := singleNumber(arr)

	fmt.Printf("Single number that is present 1 time instead of 3 times: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{2, 2, 3, 2},
		3,
	)
}

func test2() {
	test(
		[]int{0, 1, 0, 1, 0, 1, 99},
		99,
	)
}

func test3() {
	test(
		[]int{1, -1, 1, 1},
		-1,
	)
}

func main() {
	// 137. Single Number II
	test1()
	test2()
	test3()
}
