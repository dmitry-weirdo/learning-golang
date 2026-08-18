package main

import (
	"demo/matrixCommon"
	"fmt"
)

var directions = [][]int{
	{0, 1},  // right // we start from going right
	{1, 0},  // bottom
	{0, -1}, // left
	{-1, 0}, // top
}

func generateMatrix(n int) [][]int {
	m := n // square matrix

	// 0 value in the matrix means "non-visited"
	mat := createIntMatrix(m, n)

	i := 0
	j := -1 // start from -1-th column, so that the first move will go to cell [0][0]

	directionIndex := 0 // start with right
	direction := directions[directionIndex]

	pathLength := m * n
	pathIndex := 0

	// unvisited cells are cells with the initial value 0
	for pathIndex < pathLength { // until all cells passed
		for cellExists(m, n, i+direction[0], j+direction[1]) &&
			!isVisited(mat, i+direction[0], j+direction[1]) {
			i = i + direction[0]
			j = j + direction[1]

			pathIndex++

			mat[i][j] = pathIndex
		}

		directionIndex++
		directionIndex = directionIndex % len(directions)
		direction = directions[directionIndex]
	}

	return mat
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
}

func cellExists(rows, columns, row, column int) bool {
	return (0 <= row) &&
		(row < rows) &&
		(0 <= column) &&
		(column < columns)
}

func isVisited(m [][]int, row, column int) bool {
	return m[row][column] > 0
}

func test(n int, expectedResult [][]int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("N - number to generate the N * N matrix: %v \n", n)

	result := generateMatrix(n)
	fmt.Printf("Generated %v * %v matrix with values from 1 to %v: \n", n, n, n*n)
	matrixCommon.PrintIntMatrix(result)

	fmt.Printf("Generated matrix with values from 1 to %v: %v \n", n*n, result)
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
	n := 3

	expected := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}

	test(n, expected)
}

func test2() {
	n := 2

	expected := [][]int{
		{1, 2},
		{4, 3},
	}

	test(n, expected)
}

func test3() {
	n := 1

	expected := [][]int{
		{1},
	}

	test(n, expected)
}

func main() {
	// 59. Spiral Matrix II
	// Small adaptation of "54. Spiral Matrix", iteration is the same, but we create a matrix and fill it,
	// instead of collecting the values on the path.
	test1()
	test2()
	test3()
}
