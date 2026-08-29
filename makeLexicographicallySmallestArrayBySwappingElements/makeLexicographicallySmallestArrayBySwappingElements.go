package main

import (
	"fmt"
	"slices"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	// we can swap any nearby elements if their diff is <= limit
	// in every group of joint elements, we put ascending to their original indices

	// a[i][0] - value
	// a[i][1] - index in nums
	n := len(nums)
	a := make([][]int, n)

	for i, v := range nums {
		a[i] = []int{v, i}
	}

	// sort by value, index
	slices.SortFunc(a, func(a, b []int) int {
		// same value -> sort by indices
		if a[0] == b[0] {
			return a[1] - b[1]
		}

		// sort by value
		return a[0] - b[0]
	})

	//fmt.Printf("Sorted by value, index: %v \n", a)

	result := make([]int, n)

	// within every limit-diff group, put their ordered values to their ordered indexes
	i := 0

	for i < n {
		groupEnd := i
		for (groupEnd < n-1) && (a[groupEnd+1][0]-a[groupEnd][0] <= limit) {
			groupEnd++
		}

		//fmt.Printf("Group of values with diff <= %v: %v \n", limit, a[i:groupEnd+1])

		// order the indexes of the array
		indexes := make([]int, groupEnd-i+1)

		for j := i; j <= groupEnd; j++ {
			indexes[j-i] = a[j][1]
		}

		slices.Sort(indexes)

		//fmt.Printf("Group indexes sored: %v \n", indexes)

		// put into the result array
		for j, index := range indexes {
			result[index] = a[i+j][0]
		}

		i = groupEnd + 1
	}

	return result
}

func test(arr []int, limit int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Array: %v \n", arr)
	fmt.Printf("Limit of swapped elements diff: %v \n", limit)

	result := lexicographicallySmallestArray(arr, limit)

	fmt.Printf("Maximum sorted array: %v \n", result)
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
		[]int{1, 5, 3, 9, 8},
		2,
		[]int{1, 3, 5, 8, 9},
	)
}

func test2() {
	test(
		[]int{1, 7, 6, 18, 2, 1},
		3,
		[]int{1, 6, 7, 18, 1, 2},
	)
}

func test3() {
	test(
		[]int{1, 7, 28, 19, 10},
		3,
		[]int{1, 7, 28, 19, 10},
	)
}

func main() {
	// 2948. Make Lexicographically Smallest Array by Swapping Elements
	test1()
	test2()
	test3()
}
