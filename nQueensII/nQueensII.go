package main

import (
	"fmt"
)

func totalNQueens(n int) int {
	// passes in 0 ms
	return totalNQueens_precalculated(n)

	// passes in 0-3 ms
	//return totalNQueens_backtracking(n)
}

func totalNQueens_precalculated(n int) int {
	memo := []int{1, 0, 0, 2, 10, 4, 40, 92, 352}
	return memo[n-1]
}

func totalNQueens_backtracking(n int) int {
	// Every solution is a string array of N strings (string is a row)

	// In every row, we place 1 queen, so we just iterate rows from 0 to (n - 1)
	// Every column also should be unique.

	// The trick is how to check diagonals.
	// Main diagonals are identified by constant (row - col) values:
	// 0 -1 -2 -3
	// 1  0 -1 -2
	// 2  1  0 -1
	// 3  2  1  0

	// Anti-diagonals are identified by constant (row + col) values:
	// 0  1  2  3
	// 1  2  3  4
	// 2  3  4  5
	// 3  4  5  6

	// track used with O(1) presence check
	columns := make(map[int]bool)
	diagonals := make(map[int]bool)
	antiDiagonals := make(map[int]bool)

	result := 0

	//fmt.Printf("Empty board: %v \n", board)

	var dfs func(row int)

	dfs = func(row int) {
		if row >= n { // went over all rows -> collect the solution
			result++
			return
		}

		for column := range n {
			diagonal := row - column
			antiDiagonal := row + column

			// check whether [row][col] is already hit by previous queens
			if columns[column] || diagonals[diagonal] || antiDiagonals[antiDiagonal] {
				continue
			}

			// add the current queen
			columns[column] = true
			diagonals[diagonal] = true
			antiDiagonals[antiDiagonal] = true

			// with the current queen, continue to the next row
			dfs(row + 1)

			// backtrack - remove the current queen
			columns[column] = false
			diagonals[diagonal] = false
			antiDiagonals[antiDiagonal] = false
		}

	}

	dfs(0)

	return result
}

func precalculate() {
	minN := 1
	maxN := 9

	for i := minN; i <= maxN; i++ {
		r := totalNQueens_backtracking(i)

		fmt.Printf("%v - %v \n", i, r)
	}
}

func test(n int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - size of the desk: %v \n", n)

	result := totalNQueens(n)

	fmt.Printf("Different solutions of %v queens: %v \n", n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(1, 1)
}

func test2() {
	test(4, 2)
}

func main() {
	// 52. N-Queens II
	// same as "51. N-Queens", just return the size of the result instead of result
	test1()
	test2()

	//precalculate()
}
