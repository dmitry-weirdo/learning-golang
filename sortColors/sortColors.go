package main

import "fmt"

func sortColors(nums []int) {
	// todo: solve as "Dutch National Flag Problem" algorithm, that will be O(n)

	// O(2 * n)
	sortColors_counts(nums)
}

func sortColors_counts(nums []int) {
	counts := make([]int, 3) // counts of 0, 1, 2

	for _, v := range nums {
		counts[v]++
	}

	j := 0

	for i, v := range counts {
		for range v {
			nums[j] = i
			j++
		}
	}
}

func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array of 0, 1, 2 values: %v \n", arr)

	sortColors(arr)
	result := arr // function sorts in place

	fmt.Printf("Sorted array: %v \n", result)
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
		[]int{2, 0, 2, 1, 1, 0},
		[]int{0, 0, 1, 1, 2, 2},
	)
}

func test2() {
	test(
		[]int{2, 0, 1},
		[]int{0, 1, 2},
	)
}

func test3() {
	test(
		[]int{0},
		[]int{0},
	)
}

func main() {
	// 75. Sort Colors
	test1()
	test2()
	test3()
}
