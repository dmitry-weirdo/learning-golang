package main

import "fmt"

func stoneGameV(stoneValue []int) int {
	ps := getPrefixSums(stoneValue)
	//fmt.Printf("Prefix sums: %v \n", ps)

	n := len(stoneValue)

	// best sum when playing on [i; j] inclusive.
	// todo: maybe set default values as -1, since some result can be calculated as 0
	memo := createIntMatrix(n, n)

	var dfs func(left, right int) int

	dfs = func(left, right int) int {
		if left >= right { // 1 or 0 stones -> game stopped
			return 0
		}

		if memo[left][right] != 0 { // already computed -> return
			//fmt.Printf("Memo[%v][%v] = %v already calculated. Returning it. \n", left, right, memo[left][right])
			return memo[left][right]
		}

		maxScore := 0 // max Alice score
		leftSum := 0
		rightSum := 0

		aliceScore1 := 0
		aliceScore2 := 0

		// try all possible split points from left to right inclusive
		// both left and right piles must be non-empty.
		// left pile includes K, so we start with left
		// right is NON-inclusive, so we end K, so we end K at (right - 1)
		for k := left; k <= right-1; k++ {
			leftSum = ps[k+1] - ps[left]
			rightSum = ps[right+1] - ps[k+1]

			//fmt.Printf("Range [%v; %v], split (to left) = a[%v] = %v. Left sum = %v. Right sum = %v. \n", left, right, k, stoneValue[k], leftSum, rightSum)

			if leftSum < rightSum {
				// Bob takes right, Alice gets left

				// max possible win is rightSum (current take) + rightSum (all the remaining right array)
				// if maxScore is already more than this possible max, no need to run DFS further -> prune
				if maxScore > 2*leftSum {
					continue
				}

				aliceScore1 = leftSum + dfs(left, k)
				maxScore = max(maxScore, aliceScore1)
			} else if leftSum > rightSum {
				// Bob takes left, Alice gets right

				// max possible win is rightSum (current take) + rightSum (all the remaining right array)
				// if maxScore is already more than this possible max, no need to run DFS further -> prune
				if maxScore > 2*rightSum {
					continue
				}

				aliceScore1 = rightSum + dfs(k+1, right)
				maxScore = max(maxScore, aliceScore1)
			} else {
				// sums are equal -> Alice can try both option

				// Alice takes left
				if maxScore > 2*leftSum { // no need to DFS left
					aliceScore1 = -1
				} else {
					aliceScore1 = leftSum + dfs(left, k)
				}

				if maxScore > 2*rightSum { // no need to DFS right
					aliceScore2 = -1
				} else {
					aliceScore2 = rightSum + dfs(k+1, right)
				}

				maxScore = max(maxScore, aliceScore1)
				maxScore = max(maxScore, aliceScore2)
			}
		}

		//fmt.Printf("Max score for range [%v; %v]: %v \n", left, right, maxScore)
		memo[left][right] = maxScore
		return maxScore
	}

	return dfs(0, n-1)
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

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func test(arr []int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Stones: %v \n", arr)

	result := stoneGameV(arr)

	fmt.Printf("Maximum score of Alice (1st player): %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[]int{6, 2, 3, 4, 5, 5},
		18,
	)
}

func test2() {
	test(
		[]int{7, 7, 7, 7, 7, 7, 7},
		28,
	)
}

func test3() {
	test(
		[]int{4},
		0,
	)
}

func main() {
	// 1563. Stone Game V
	test1()
	test2()
	test3()
}
