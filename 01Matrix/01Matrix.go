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

func updateMatrix(mat [][]int) [][]int {
	// we need to do BFS from the 0-cells.
	m, n := getRowsAndColumns(mat)

	queue := make([][]int, 0) // array of pairs

	// Collect into a separate matrix, collect all 0-coordinates
	// 0 - start
	// 1 - cells to calculate BFS-distance from 0-cells
	for i, row := range mat {
		for j, v := range row {
			if v == 0 {
				// save 0-cell coordinates
				queue = append(queue, []int{i, j})
			} else {
				// change non-0 cell to -1 - mark "non-visited"
				mat[i][j] = -1
			}
		}
	}

	fmt.Printf("Matrix after changing 1 to -1: \n")
	matrixCommon.PrintIntMatrix(mat)

	fmt.Printf("0-cells coordinates: %v \n", queue)

	// 0-th level is 0-cells
	level := 0

	for len(queue) > 0 {
		currentLevelLength := len(queue)

		level++

		for range currentLevelLength {
			// poll coords from queue
			current := queue[0]
			queue = queue[1:]

			i := current[0]
			j := current[1]

			// we set the neighbours of the next level to visited, so that on the next check they won't be added twice

			// set current level to the cell of the current level
			for _, d := range directions {
				neighborI := i + d[0]
				neighborJ := j + d[1]

				if cellExists(m, n, neighborI, neighborJ) &&
					!isVisited(mat[neighborI][neighborJ]) {
					mat[neighborI][neighborJ] = level                  // mark as visited
					queue = append(queue, []int{neighborI, neighborJ}) // add to the next level of queue
				}
			}
		}
	}

	return mat
}

func getRowsAndColumns(mat [][]int) (rows, columns int) {
	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func cellExists(rows, columns, row, column int) bool {
	return (0 <= row) &&
		(row < rows) &&
		(0 <= column) &&
		(column < columns)
}

func isVisited(x int) bool {
	return x >= 0
}

func test(m [][]int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: \n")
	matrixCommon.PrintIntMatrix(m)

	result := updateMatrix(m)

	fmt.Printf("Matrix of BFS-distances to 0: \n")
	matrixCommon.PrintIntMatrix(result)

	fmt.Printf("Expected result: \n")
	matrixCommon.PrintIntMatrix(expectedResult)

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
	m := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	expected := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	test(m, expected)
}

func test2() {
	m := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	expected := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 2, 1},
	}

	test(m, expected)
}

func main() {
	// 542. 01 Matrix
	// basically the same task as "1765. Map of Highest Peak", but with better input (0 means 0 (BFS start), 1 means 1)
	test1()
	test2()
}
