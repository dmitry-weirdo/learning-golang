package main

import "fmt"

func maxProduct(nums []int) int {
	bestResult := nums[0]

	const FAKE_MAX_NEGATIVE = -10000 // -10 <= nums[i] <= 10

	minPositive := 1
	maxNegative := FAKE_MAX_NEGATIVE // min by abs, so it's logically maxNegative :)

	product := 1

	i := 0

	for i < len(nums) {
		if nums[i] == 0 { // 0 breaks the sums -> we have to restart
			bestResult = max(bestResult, 0) // if we only have negative products -> 0 is a better answer

			minPositive = 1
			maxNegative = FAKE_MAX_NEGATIVE
			product = 1
		} else {
			product *= nums[i]

			if product > 0 {
				bestResult = max(bestResult, product/minPositive)

				minPositive = min(minPositive, product) // todo: this doesn't make sense since minPositive is always 1
			} else {
				if maxNegative != FAKE_MAX_NEGATIVE { // there was a negative number -> consider a "negative / negative" product as well.
					bestResult = max(bestResult, product/maxNegative)
				}

				maxNegative = max(maxNegative, product)
			}
		}

		i++
	}

	return bestResult
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Array: %v \n", arr)

	result := maxProduct(arr)

	fmt.Printf("Max product of subarray: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{2, 3, -2, 4},
		6,
	)
}

func test2() {
	test(
		[]int{-2, 0, -1},
		0,
	)
}

func test3() {
	test(
		[]int{-1}, // only negative values
		-1,
	)
}

func main() {
	// 152. Maximum Product Subarray
	test1()
	test2()
	test3()
}
