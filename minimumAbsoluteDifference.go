package main

import (
	"fmt"
	"slices"
)

func minimumAbsDifference(arr []int) [][]int {
	// Since we're guaranteed that all the values are unique,
	// in the sorted array we should only compare to the next index
	//since the diff between [i] and [i + 2] will be bigger than diff between [i] and [i + 1]
	slices.Sort(arr) // O(N * log N)

	// we're guaranteed that there are at least 2 elements in the array
	minDiff := arr[1] - arr[0]

	for i := 1; i < len(arr)-1; i++ {
		minDiff = min(minDiff, arr[i+1]-arr[i])
	}

	// collect the result -> get all pairs that have minDiff
	result := make([][]int, 0)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == minDiff {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}

	return result
}

func test(arr []int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("Array: %v \n", arr) // todo: replace with your text if required

	result := minimumAbsDifference(arr)

	fmt.Printf("Pairs array with minimums diff between elements: %v \n", result) // todo: replace with your text if required
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i := 0; i < len(expectedResult); i++ {
		r := result[i]
		er := expectedResult[i]

		if r[0] != er[0] || r[1] != er[1] { // we only compare pairs, i.e. [0] and [1] elements
			fmt.Printf("FAILURE: expected result[%v] = [%v; %v], actual result[%v] = [%v; %v] \n", i, er[0], er[1], i, r[0], r[1])
		}
	}
}

func test1() {
	arr := []int{4, 2, 1, 3}
	expected := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}

	test(arr, expected)
}

func test2() {
	arr := []int{1, 3, 6, 10, 15}
	expected := [][]int{
		{1, 3},
	}

	test(arr, expected)
}

func test3() {
	arr := []int{3, 8, -10, 23, 19, -4, -14, 27}
	expected := [][]int{
		{-14, -10},
		{19, 23},
		{23, 27},
	}

	test(arr, expected)
}

func main() {
	// 1200. Minimum Absolute Difference
	test1()
	test2()
	test3()
}
