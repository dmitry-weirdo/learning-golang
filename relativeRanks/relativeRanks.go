package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	// This can also solved by heap, we put pairs of (score, index) to the heap and the pull from it.
	// It will be also O(n * log n) solution, same as sorting.

	// scores to original index
	// we can map from scores, since they're guaranteed to be unique
	n := len(score)

	m := make(map[int]int, n)

	for i, v := range score {
		m[v] = i
	}

	// sort by scores desc
	slices.SortFunc(score, func(a, b int) int {
		return -cmp.Compare(a, b)
	})

	//fmt.Printf("Sorted scores: %v \n", score)

	result := make([]string, n)
	index := -1

	if n > 0 {
		index = m[score[0]]
		result[index] = "Gold Medal"
	}

	if n > 1 {
		index = m[score[1]]
		result[index] = "Silver Medal"
	}

	if n > 2 {
		index = m[score[2]]
		result[index] = "Bronze Medal"
	}

	for i := 3; i < n; i++ { // for places after the first 3 values, assign 1-based places
		index = m[score[i]]
		result[index] = strconv.Itoa(i + 1) // places are 1-indexed, therefore (i + 1)
	}

	return result
}

func test(arr []int, expectedResult []string) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Scores (unique): %v \n", arr)

	result := findRelativeRanks(arr)

	fmt.Printf("Relative ranks: %v \n", result)
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
		[]int{5, 4, 3, 2, 1},
		[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
	)
}

func test2() {
	test(
		[]int{10, 3, 8, 9, 4},
		[]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
	)
}

func main() {
	// 506. Relative Ranks
	test1()
	test2()
}
