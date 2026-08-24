package main

import "fmt"

func stoneGameVIII(stones []int) int {
	n := len(stones)

	prefixSums := getPrefixSums(stones)

	// At every position i, we can:
	// - either take stones up to this position -> opponent plays from the next position (i + 1)
	// - skip this position and take stones up to the further position -> we are deciding from the next position (i + 1)

	memo := make([]any, n) // any is nil

	var dfs func(i int) int

	dfs = func(i int) int {
		if i >= (n - 1) { // last position -> we take all stones, no option to skip and go further
			return prefixSums[n]
		}

		if memo[i] != nil { // already precalculated
			return memo[i].(int)
		}

		// skip current - we continue from the next position
		skipCurrent := dfs(i + 1)

		// take current - we take the current prefix sum, opponent plays from dfs(i - 1)
		takeCurrent := prefixSums[i+1] - dfs(i+1) // prefixSums are 0-based

		currentOptimal := max(skipCurrent, takeCurrent)
		memo[i] = currentOptimal

		return currentOptimal
	}

	return dfs(1) // at pos 0, we should skip pos 0 in any case
}

func getPrefixSums(a []int) []int {
	// with prefixSums[0] = 0, sum up to a[i] will be in prefixSums[i+1]
	prefixSums := make([]int, len(a)+1)

	prefixSums[0] = 0

	for i, v := range a {
		prefixSumIndex := i + 1
		prefixSums[prefixSumIndex] = prefixSums[i] + v
	}

	return prefixSums
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Stones weights: %v \n", arr)

	result := stoneGameVIII(arr)

	fmt.Printf("Max difference between Alice and Bob scores: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test([]int{-1, 2, -3, 4, -5}, 5)
}

func test2() {
	test([]int{7, -6, 5, 10, 5, -2, -6}, 13)
}

func test3() {
	test([]int{-10, -12}, -22)
}

func test4() {
	test([]int{3, 7, 2, 3}, 15)
}

func main() {
	// 1872. Stone Game VIII
	test1()
	test2()
	test3()
	test4()
}
