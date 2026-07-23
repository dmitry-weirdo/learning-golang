package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	return numSquaresDp(n)
	//return numSquaresBfs(n)
}

func numSquaresDp(n int) int {
	dp := make([]int, n+1)

	// initialize all values to worst-case -> sum of all ones up to this value
	for i := 0; i <= n; i++ {
		dp[i] = i
	}

	//// for 0, we don't require any
	//dp[0] = 0

	for x := 1; x <= n; x++ { // calculate every next value based on previous
		for i := 1; i*i <= x; i++ { // iterate all squares up to X
			// subtract the square of i, this is one square (therefore + 1 in the "count of squares")
			// the dp[x - i*i] is the remaining part, we take it from the DP memo

			// example for x = 4
			// i goes from 1 to 2
			// x - 1*1 = 3, dp[3] = 3, + 1 = 4 -> this is the first result (1 + 1 + 1 + 1)
			// x - 2*2 = 0, dp[0] = 0, + 1 = 1 -> this is the better result (4 = 4)

			subtractISquareResult := dp[x-i*i] + 1
			dp[x] = min(dp[x], subtractISquareResult)
		}
	}

	return dp[n]
}

func numSquaresBfs(n int) int {
	squares := getSquaresNotBiggerThan(n)
	fmt.Printf("Squares up to %v: %v \n", n, squares)

	// !!! important trick -> on every level, do not handle repeating (duplicate) values N times,
	// use a unique structure to avoid duplicates
	queue := make(map[int]bool)
	queue[n] = true

	level := 1

	for len(queue) > 0 {
		// iterate the current level
		newQueue := make(map[int]bool)

		fmt.Println()
		fmt.Printf("Handling level %v... \n", level)

		for k := range queue {
			fmt.Printf("Trying to subtract squares from %v... \n", k)

			// generate next level
			// todo: it would be more optimal to iterate down from sqrt(n), to find the terminal case a bit faster
			for _, v := range squares {
				if k-v > 0 {
					fmt.Printf("Added %v - %v = %v to the next level. \n", k, v, k-v)
					newQueue[k-v] = true
				}

				if k-v == 0 { // reached 0 by subtracting the square -> solution found
					fmt.Printf("Solution found: %v - %v == 0. Returning level = %v. \n", k, v, level)
					return level
				}

				if k-v < 0 { // square is bigger then the remainder -> stop handling the further squares
					break
				}
			}
		}

		queue = newQueue
		level++
	}

	return level
}

func getSquaresNotBiggerThan(n int) []int {
	// form perfect squares arrays
	intSqrt := int(math.Ceil(math.Sqrt(float64(n))))
	if !isSquare(n) { // for non-squares 1 less size is required
		intSqrt--
	}

	squares := make([]int, intSqrt)
	for i := 1; i*i <= n; i++ {
		squares[i-1] = i * i
	}

	return squares
}

func isSquare(n int) bool {
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}

func test(n int, expectedResult int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("N: %v \n", n)

	result := numSquares(n)

	fmt.Printf("Min squares required to sum up to N = %v: %v \n", n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	n := 12
	expected := 3 // 4 + 4 + 4

	test(n, expected)
}

func test2() {
	n := 13
	expected := 2 // 9 + 4

	test(n, expected)
}

func test3() {
	n := 16
	expected := 1 // 16

	test(n, expected)
}

func main() {
	// 279. Perfect Squares
	test1()
	test2()
	test3()
}
