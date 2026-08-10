package main

import (
	"fmt"
	"math"
	"strings"
)

const MAX_N = 100_000

func winnerSquareGame(n int) bool {
	return winnerSquareGame_dp_array_no_recursion(n) // no recursion, we build the dp[] array for all values bottom-up. The quickest -> runs in 0 ms!
	//return winnerSquareGame_dp_array(n) // use array instead of map - yes, it's much faster, 55 ms instead of 340+ ms with a map!
	//return winnerSquareGame_dp(n) // straightforward DP solution
}

func winnerSquareGame_dp_array_no_recursion(n int) bool {
	dp := make([]bool, n+1)

	for i := 0; i <= n; i++ {
		if dp[i] { // winning position -> skip
			continue
		}

		// if we reached a losing position -> all (+ square) positions from it are winning
		for k := 1; i+k*k <= n; k++ {
			dp[i+k*k] = true
		}
	}

	return dp[n]
}

func winnerSquareGame_dp_array_collect_all_values(n int) bool {
	var t = true
	var f = false
	TRUE := &t
	FALSE := &f

	maxPossibleSquareBase := min(sqrtCeilDown(n), 316)

	// prefill for all squares
	dp := make([]*bool, MAX_N+1)                         // use an array instead of map since it's faster
	dp[0] = FALSE                                        // if there are no stones -> we have lost
	fillSquaresForArray(dp, maxPossibleSquareBase, TRUE) // max N = 10^5 = 100_000, so its root is ≈ 316.227766. If sqrt(N) < 316, no need to fill

	var dfs func(x int) bool

	dfs = func(x int) bool { // returns true if the current player wins, false if the current player loses
		// result already calculated -> return
		if dp[x] != nil {
			return *dp[x]
		}

		// squares are already in dp
		/*
			// base case - if current count is square -> we win
			if isSquare(x) {
				return true
			}
		*/

		// try to take all possible square stones within x, so that (x - taken) is false (i.e. opponent losing on the remaining amount)
		remaining := 0
		opponentResult := true

		for i := 1; i*i < x; i++ { // we have to take at least 1 square
			remaining = x - i*i

			opponentResult = dfs(remaining)
			if !opponentResult { // opponent loses -> we win
				dp[x] = TRUE
				return true
			}
		}

		// all options lead to opponent win -> we lose
		dp[x] = FALSE
		return false
	}

	for i := 0; i <= MAX_N; i++ {
		dfs(i)
	}

	boolResult := make([]bool, MAX_N+1)

	for i, v := range dp {
		boolResult[i] = *v
	}

	//fmt.Printf("All calculated values: %v \n", boolResult)

	sb := strings.Builder{}

	for i, v := range boolResult {
		sb.WriteString(fmt.Sprintf("%v, ", v))
		if i%100 == 0 {
			sb.WriteString("\n")
		}
	}

	fmt.Printf("All calculated values as string: \n%v \n", sb.String())

	return false
}

func winnerSquareGame_dp_array(n int) bool {
	var t = true
	var f = false
	TRUE := &t
	FALSE := &f

	maxPossibleSquareBase := min(sqrtCeilDown(n), 316)

	cacheLength := n + 1

	// prefill for all squares
	dp := make([]*bool, cacheLength)                     // use an array instead of map since it's faster
	dp[0] = FALSE                                        // if there are no stones -> we have lost
	fillSquaresForArray(dp, maxPossibleSquareBase, TRUE) // max N = 10^5 = 100_000, so its root is ≈ 316.227766. If sqrt(N) < 316, no need to fill

	var dfs func(x int) bool

	dfs = func(x int) bool { // returns true if the current player wins, false if the current player loses
		// result already calculated -> return
		if dp[x] != nil {
			return *dp[x]
		}

		// squares are already in dp
		/*
			// base case - if current count is square -> we win
			if isSquare(x) {
				return true
			}
		*/

		// try to take all possible square stones within x, so that (x - taken) is false (i.e. opponent losing on the remaining amount)
		remaining := 0
		opponentResult := true

		for i := 1; i*i < x; i++ { // we have to take at least 1 square
			remaining = x - i*i

			opponentResult = dfs(remaining)
			if !opponentResult { // opponent loses -> we win
				dp[x] = TRUE
				return true
			}
		}

		// all options lead to opponent win -> we lose
		dp[x] = FALSE
		return false
	}

	return dfs(n)
}

func winnerSquareGame_dp(n int) bool {
	// prefill for all squares
	dp := make(map[int]bool)
	dp[0] = false        // if there are no stones -> we have lost
	fillSquares(dp, 100) // max N = 10^5 =10000, so its root is 100

	var dfs func(x int) bool

	dfs = func(x int) bool { // returns true if the current player wins, false if the current player loses
		// result already calculated -> return
		if cached, ok := dp[x]; ok {
			return cached
		}

		// squares are already in dp
		/*
			// base case - if current count is square -> we win
			if isSquare(x) {
				return true
			}
		*/

		// try to take all possible square stones within x, so that (x - taken) is false (i.e. opponent losing on the remaining amount)
		remaining := 0
		opponentResult := true

		for i := 1; i*i < x; i++ { // we have to take at least 1 square
			remaining = x - i*i

			opponentResult = dfs(remaining)
			if !opponentResult { // opponent loses -> we win
				dp[x] = true
				return true
			}
		}

		// all options lead to opponent win -> we lose
		dp[x] = false
		return false
	}

	return dfs(n)
}

func fillSquares(dp map[int]bool, maxValue int) {
	for i := 1; i <= maxValue; i++ {
		dp[i*i] = true
	}
}

func fillSquaresForArray(dp []*bool, maxValue int, truePointer *bool) {
	for i := 1; i <= maxValue; i++ {
		dp[i*i] = truePointer
	}
}

func isSquare(n int) bool {
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}

func sqrtCeilDown(c int) int {
	return int(math.Sqrt(float64(c)))
}

func test(x int, expectedResult bool) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Number of stones: %v \n", x)

	result := winnerSquareGame(x)

	fmt.Printf("Result: %v \n", result) // todo: replace with your text
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(1, true)
}

func test2() {
	test(2, false)
}

func test3() {
	test(3, true) // take 1, then Bob has 2 and loses
}

func test4() {
	test(4, true)
}

func test5() {
	test(7, false)
}

func main() {
	// 1510. Stone Game IV

	//winnerSquareGame_dp_array_collect_all_values(MAX_N)

	test1()
	test2()
	test3()
	test4()
	test5()
}
