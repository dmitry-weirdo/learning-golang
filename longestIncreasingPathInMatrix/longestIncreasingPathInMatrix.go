package main

import (
	"demo/matrixCommon"
	"fmt"
)

var directions = [][]int{
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // bottom
	{0, -1}, // left
}

func cellExists(rows, columns, row, column int) bool {
	return (0 <= row) &&
		(row < rows) &&
		(0 <= column) &&
		(column < columns)
}

func longestIncreasingPath(matrix [][]int) int {
	rows := len(matrix)
	columns := len(matrix[0])

	// create a memo matrix of the same size
	memo := make([][]int, rows)
	for i := range rows {
		memo[i] = make([]int, columns)
	}

	maxPathLength := 0

	for i := range rows {
		for j := range columns {
			pathLength := dfs(matrix, memo, rows, columns, i, j)

			fmt.Println("========================")
			fmt.Printf("[%v][%v], max path length = %v \n", i, j, pathLength)

			fmt.Printf("Memo matrix: \n")
			matrixCommon.PrintIntMatrix(memo)

			maxPathLength = max(maxPathLength, pathLength)
		}
	}

	return maxPathLength
}

func dfs(m [][]int, memo [][]int, rows int, columns int, i int, j int) int {
	fmt.Printf("==================\n")
	fmt.Printf("dfs[%v][%v] \n", i, j)

	if memo[i][j] != 0 { // path for this cell already calculated -> return it
		fmt.Printf("Memo[%v][%v] already calculated and is %v. Returning it. \n", i, j, memo[i][j])
		return memo[i][j]
	}

	// iterate all 4 directions
	for _, d := range directions {
		row := i + d[0]
		column := j + d[1]

		if cellExists(rows, columns, row, column) &&
			(m[row][column] > m[i][j]) { // only move to cells that have bigger value
			// !!! the main trick is that not visiting the already visited cell is auto-covered by "go only to the bigger value" logic

			cellPath := dfs(m, memo, rows, columns, row, column)

			// this is the DP-logic: max of the current cell depends on pre-calculated values
			memo[i][j] = max(memo[i][j], cellPath)
		}
	}

	// cell itself counts -> add 1 to the distance for the cell itself
	memo[i][j] += 1
	return memo[i][j]
}

func test(m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Println("Matrix:")
	matrixCommon.PrintIntMatrix(m)

	result := longestIncreasingPath(m)

	fmt.Println()
	fmt.Printf("Expected longest path length: %v \n", expectedResult)
	fmt.Printf("Calculated longest path length: %v \n", result)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}

	expected := 4

	test(m, expected)
}

func test2() {
	m := [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}

	expected := 4

	test(m, expected)
}

func test3() {
	m := [][]int{
		{1, 2, 3},
		{6, 5, 4},
		{7, 8, 9},
	}

	expected := 9

	test(m, expected)
}

func main() {
	// 329. Longest Increasing Path in a Matrix
	test1()
	// test2()
	//test3()
}
