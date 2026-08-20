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

func highestPeak(isWater [][]int) [][]int {
	// we need to do BFS from the water cells.
	// water cells are initially 0
	m, n := getRowsAndColumns(isWater)

	queue := make([][]int, 0) // array of pairs

	// Collect into a separate matrix, collect all water coordinates
	// 0 - water
	// -1 - land
	for i, row := range isWater {
		for j, v := range row {
			if cellIsWater(v) {
				// save water coordinates
				queue = append(queue, []int{i, j})

				// change water cell to 0
				isWater[i][j] = 0
			} else {
				// change land cell to -1 - mark "non-visited"
				isWater[i][j] = -1
			}
		}
	}

	fmt.Printf("Matrix after changing water to 0, land to -1: \n")
	matrixCommon.PrintIntMatrix(isWater)

	fmt.Printf("Water cells coordinates: %v \n", queue)

	// 0-th level is waterCoordinates
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

			// we set the neighbors of the next level to visited, so that on the next check they won't be added twice

			// set current level to the cell of the current level
			for _, d := range directions {
				neighborI := i + d[0]
				neighborJ := j + d[1]

				if cellExists(m, n, neighborI, neighborJ) &&
					!isVisited(isWater[neighborI][neighborJ]) {
					isWater[neighborI][neighborJ] = level              // mark as visited
					queue = append(queue, []int{neighborI, neighborJ}) // add to the next level of queue
				}
			}
		}
	}

	return isWater
}

func cellExists(rows, columns, row, column int) bool {
	return (0 <= row) &&
		(row < rows) &&
		(0 <= column) &&
		(column < columns)
}

func cellIsWater(x int) bool {
	return x == 1
}

func isVisited(x int) bool {
	return x >= 0
}

func getRowsAndColumns(mat [][]int) (rows, columns int) {
	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func test(m [][]int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix of water = 1, land = 0: \n")
	matrixCommon.PrintIntMatrix(m)

	result := highestPeak(m)

	fmt.Printf("Matrix of BFS from water: \n")
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
		{0, 1},
		{0, 0},
	}

	expected := [][]int{
		{1, 0},
		{2, 1},
	}

	test(m, expected)
}

func test2() {
	// Input: isWater = [[0,0,1],[1,0,0],[0,0,0]]
	//Output: [[1,1,0],[0,1,1],[1,2,2]]

	m := [][]int{
		{0, 0, 1},
		{1, 0, 0},
		{0, 0, 0},
	}

	expected := [][]int{
		{1, 1, 0},
		{0, 1, 1},
		{1, 2, 2},
	}

	test(m, expected)
}

func main() {
	// 1765. Map of Highest Peak
	test1()
	test2()
}
