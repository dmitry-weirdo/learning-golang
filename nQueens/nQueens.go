package main

import (
	"fmt"
	"strings"
)

const EMPTY = '.'
const QUEEN = 'Q'

func solveNQueens(n int) [][]string {
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

	board := createEmptyBoard(n)

	result := make([][]string, 0) // row is i-th solution, every solution is N strings representing N fields

	//fmt.Printf("Empty board: %v \n", board)

	var dfs func(row int)

	dfs = func(row int) {
		if row >= n { // went over all rows -> collect the solution
			result = append(result, copyBoard(board))
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
			board[row][column] = QUEEN

			// with the current queen, continue to the next row
			dfs(row + 1)

			// backtrack - remove the current queen
			columns[column] = false
			diagonals[diagonal] = false
			antiDiagonals[antiDiagonal] = false
			board[row][column] = EMPTY
		}

	}

	dfs(0)

	return result
}

func createEmptyBoard(n int) [][]byte {
	a := make([][]byte, n)

	for i := range n {
		s := strings.Repeat(".", n)
		a[i] = []byte(s)
	}

	return a
}

func copyBoard(board [][]byte) []string {
	a := make([]string, len(board))

	for i, v := range board {
		a[i] = string(v)
	}

	return a
}

func test(n int, expectedResult [][]string) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - size of the desk: %v \n", n)

	result := solveNQueens(n)

	fmt.Printf("%v queens solutions: %v \n", n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, resultRow := range result {
		expectedResultRow := expectedResult[i]

		// check that rows have the same length
		if len(resultRow) != len(expectedResultRow) {
			fmt.Printf("FAILURE: expectedResult[%v] length = %v, actualResult[%v] length = %v \n", i, len(expectedResultRow), i, len(resultRow))

			return
		}

		// same length -> check all row values
		for j, resultValue := range resultRow {
			expectedResultValue := expectedResultRow[j]

			if resultValue != expectedResultValue {
				fmt.Printf("FAILURE: expectedResult[%v][%v] = %v, actualResult[%v][%v]  = %v \n", i, j, expectedResultValue, i, j, resultValue)

				return
			}
		}
	}
}

func test1() {
	n := 4

	expected := [][]string{
		{
			".Q..",
			"...Q",
			"Q...",
			"..Q.",
		},
		{
			"..Q.",
			"Q...",
			"...Q",
			".Q..",
		},
	}

	test(n, expected)
}

func test2() {
	n := 1

	expected := [][]string{
		{
			"Q",
		},
	}

	test(n, expected)
}

func main() {
	// 51. N-Queens
	test1()
	test2()
}
