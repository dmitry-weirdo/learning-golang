package main

import (
	"demo/matrixCommon"
	"fmt"
)

func hasValidPath(m [][]byte) bool {
	rows, columns := getRowsAndColumnsOfByteMatrix(m)

	// the valid path must start with '(' and end with ')'
	if (m[0][0] != '(') || (m[rows-1][columns-1] != ')') {
		return false
	}

	// path length must be even. Path length is (rows + columns - 1)
	pathLength := rows + columns - 1
	if pathLength%2 != 0 {
		return false
	}

	// we must have openBracketsCount - closingBracketsCount >= 0, else this is a bad path, and we have to stop
	// If current cell is `(`, we add to the balance
	// If current cell is `)`, we subtract from the balance

	// since we can only go right and down,
	// we can proceed every row, column by column
	// Next row will have the top cell from the previous row.
	// Right cell will have the left cell from earlier handling in this row.

	// For every cell, we will DP-memo the possible balances
	// todo: to save space, we can just save previous row and current row, not the complete matrix

	dp := createDpMatrix(rows, columns, pathLength)

	// starting opening bracket at top-left -> balance = 1
	dp[0][0][1] = true

	for i := range rows {
		for j := range columns {
			diff := 1           // opening bracket
			if m[i][j] == ')' { // closing bracket
				diff = -1
			}

			if i > 0 { // we can go from top
				// for all possible balances at the top, add balances with diff if they're valid, i.e. >= 0
				upperRow := dp[i-1][j]

				for upperRowBalance, v := range upperRow {
					if !v { // balance is not possible in the upper row
						continue
					}

					newBalance := upperRowBalance + diff
					if newBalance < 0 { // cannot proceed with a negative balance
						continue
					}

					dp[i][j][newBalance] = true
				}
			}

			if j > 0 { // we can go from left
				// for all possible balances at the top, add balances with diff if they're valid, i.e. >= 0
				leftRow := dp[i][j-1]

				for leftRowBalance, v := range leftRow {
					if !v { // balance is not possible in the upper row
						continue
					}

					newBalance := leftRowBalance + diff
					if newBalance < 0 { // cannot proceed with a negative balance
						continue
					}

					dp[i][j][newBalance] = true
				}
			}
		}
	}

	// we need to check whether we can reach the bottom left with 0 balance
	return dp[rows-1][columns-1][0]
}

func getRowsAndColumnsOfByteMatrix(mat [][]byte) (rows, columns int) {
	if len(mat) <= 0 {
		return 0, 0
	}

	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func createDpMatrix(rows, columns, pathLength int) [][][]bool {
	m := make([][][]bool, rows)

	for i := range rows {
		m[i] = make([][]bool, columns)

		for j := range columns {
			m[i][j] = make([]bool, pathLength)
		}
	}

	return m
}

func test(m [][]byte, expectedResult bool) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix of brackets: \n")
	matrixCommon.PrintByteMatrix(m)

	result := hasValidPath(m)

	fmt.Printf("Matrix has a valid brackets path from top-left to bottom-right: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test(
		[][]byte{
			{'(', '(', '('},
			{')', '(', ')'},
			{'(', '(', ')'},
			{'(', '(', ')'},
		},
		true,
	)
}

func test2() {
	test(
		[][]byte{
			{')', ')'},
			{'(', '('},
		},
		false,
	)
}

func main() {
	// 2267. Check if There Is a Valid Parentheses String Path
	test1()
	test2()
}
