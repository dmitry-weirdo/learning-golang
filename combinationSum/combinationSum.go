package main

import (
	"fmt"
	"slices"
)

var resultGlobal [][]int

func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	resultGlobal = make([][]int, 0)

	path := make([]int, 0)

	dfs(candidates, 0, target, path)

	return resultGlobal
}

func dfs(candidates []int, startIndex int, remainingTarget int, path []int) {
	// base-case -> we reached exactly the sum -> add it to the result
	if remainingTarget == 0 {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)

		fmt.Printf("Remaining sum 0 reached. Adding current path = %v to the result. \n", pathCopy)

		resultGlobal = append(resultGlobal, pathCopy)
		fmt.Printf("Result after adding: %v. \n", resultGlobal)
	}

	// If all candidates from the starting position are greater than the remaining sum -> stop this branch
	if candidates[startIndex] > remainingTarget {
		fmt.Printf("First candidate[%v] = %v is more than remainingTarget = %d. Stopping this branch. \n", startIndex, candidates[startIndex], remainingTarget)
		return
	}

	// adding the next number - try for every position starting from the current
	for i := startIndex; i < len(candidates); i++ {
		// add current candidate to path
		path = append(path, candidates[i])

		// select candidates[i]
		// proceed with the same starting index i, we can reuse the same last number
		// since candidates is sorted, we're always proceeding with the non-decreasing order
		newRemainingTarget := remainingTarget - candidates[i]
		dfs(candidates, i, newRemainingTarget, path)

		// remove the current added element from path
		path = path[:len(path)-1]
	}
}

func test(nums []int, targetSum int) {
	fmt.Println()
	fmt.Println("====================")
	fmt.Printf("Array: %v \n", nums)
	fmt.Printf("Target sum: %v \n", nums)

	result := combinationSum(nums, targetSum)

	fmt.Printf("All possible combinations with target sum = %v: %v \n", targetSum, result)
}

func test1() {
	arr := []int{2, 5, 6, 9}
	targetSum := 9

	test(arr, targetSum)
}

func test2() {
	arr := []int{3, 4, 5}
	targetSum := 16

	test(arr, targetSum)
}

func test3() {
	arr := []int{2, 3, 6, 7}
	targetSum := 7

	test(arr, targetSum)
}

func test4() {
	arr := []int{2, 3, 5}
	targetSum := 8

	test(arr, targetSum)
}

func test5() {
	arr := []int{2}
	targetSum := 1

	test(arr, targetSum)
}

func main() {
	// 39. Combination Sum
	test1()
	test2()
	test3()
	test4()
	test5()
}
