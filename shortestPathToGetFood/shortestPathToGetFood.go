package main

import (
	"demo/matrixCommon"
	"fmt"
)

const START = '*'
const FOOD = '#'
const FREE = 'O'
const OBSTACLE = 'X'

var directions = [][]int{
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // bottom
	{0, -1}, // left
}

type Cell struct {
	row    int
	column int
}

func getFood(m [][]byte) int {
	rows, columns := getRowsAndColumnsOfByteMatrix(m)

	startRow, startColumn := findStart(m)

	//fmt.Printf("Start: [%v, %v] \n", startRow, startColumn)

	// Actually, we don't need visited, we can just mark visited cells as obstacles
	//visited := createBoolMatrix(rows, columns)
	//visited[startRow][startColumn] = true

	// BFS from the start position, until we find the food cell or until the queue is exhausted.
	currentLevel := 0

	queue := make([]Cell, 0)
	queue = append(queue, Cell{startRow, startColumn})

	for len(queue) > 0 {
		currentLevelSize := len(queue)

		for range currentLevelSize {
			cell := queue[0]
			queue = queue[1:]

			//visited[cell.row][cell.column] = true
			m[cell.row][cell.column] = OBSTACLE

			for _, d := range directions {
				i := cell.row + d[0]
				j := cell.column + d[1]

				//if !cellExists(rows, columns, i, j) || visited[i][j] || m[i][j] == OBSTACLE {
				if !cellExists(rows, columns, i, j) || m[i][j] == OBSTACLE {
					// cannot move to the neighbor cell
					continue
				}

				if m[i][j] == FOOD { // food found -> return result immediately
					return currentLevel + 1
				}

				// add the cell to the next level
				queue = append(queue, Cell{i, j})

				//visited[i][j] = true
				m[i][j] = OBSTACLE
			}
		}

		currentLevel++
	}

	// food not found via DFS
	return -1
}

func findStart(m [][]byte) (row, column int) {
	rows, columns := getRowsAndColumnsOfByteMatrix(m)

	for i := range rows {
		for j := range columns {
			if m[i][j] == START {
				return i, j
			}
		}
	}

	panic(fmt.Sprintf("There is no start cell '%c' in the matrix.", START))
}

func getRowsAndColumnsOfByteMatrix(mat [][]byte) (rows, columns int) {
	if len(mat) <= 0 {
		return 0, 0
	}

	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func createBoolMatrix(rows, columns int) [][]bool {
	m := make([][]bool, rows)

	for i := range rows {
		m[i] = make([]bool, columns)
	}

	return m
}

func cellExists(rows, columns, row, column int) bool {
	return (0 <= row) &&
		(row < rows) &&
		(0 <= column) &&
		(column < columns)
}

func test(m [][]byte, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Food path matrix: \n")
	matrixCommon.PrintByteMatrix(m)

	result := getFood(m)

	fmt.Printf("Min distance to the food cell: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]byte{
		{'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'O', 'O', 'O', 'X'},
		{'X', 'O', 'O', '#', 'O', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X'},
	}

	test(m, 3)
}

func test2() {
	m := [][]byte{
		{'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'X', 'O', 'X'},
		{'X', 'O', 'X', '#', 'X'},
		{'X', 'X', 'X', 'X', 'X'},
	}

	test(m, -1)
}

func test3() {
	m := [][]byte{
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'O', 'X', 'O', '#', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O', 'O', 'X', 'X'},
		{'X', 'O', 'O', 'O', 'O', '#', 'O', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
	}

	test(m, 6)
}

func test4() {
	m := [][]byte{
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'O', 'X', 'O', '#', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O', 'O', 'X', 'X'},
		{'X', 'O', 'O', 'O', 'O', '#', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O'},
	}

	test(m, 6)
}

func test5() {
	m := [][]byte{
		{'X', 'X', 'X', 'X', 'X'},
		{'X', '*', 'X', 'O', 'X'},
		{'X', 'O', 'X', '#', 'X'},
		{'O', 'O', 'O', 'O', 'O'},
	}

	test(m, 5)
}

func main() {
	// 1730. Shortest Path to Get Food
	test1()
	test2()
	test3()
	test4()
	test5()
}
