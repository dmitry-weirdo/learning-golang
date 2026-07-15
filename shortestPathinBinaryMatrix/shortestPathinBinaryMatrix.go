package main

import (
	"container/list"
	"fmt"
)

var directions = [][]int{ // octagonal directions, includes diagonals
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // bottom
	{0, -1}, // left

	{-1, -1}, // top-left
	{-1, 1},  // top-right
	{1, -1},  // bottom-left
	{1, 1},   // bottom-right
}

func cellExists(rows, columns, row, column int) bool {
	return (0 <= row) &&
		(row < rows) &&
		(0 <= column) &&
		(column < columns)
}

func shortestPathBinaryMatrix(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])

	if grid[0][0] == 1 { // cannot start
		fmt.Printf("Invalid grid: top-left is 1. Returning -1. \n")
		return -1
	}

	if grid[rows-1][columns-1] == 1 { // cannot end
		fmt.Printf("Invalid grid: bottom-right is 1. Returning -1. \n")
		return -1
	}

	// since it's forward-only BFS without back-tracking, we can set 1 to the original matrix instead of a separate matrix
	//visited := createVisitedMatrix(rows, columns)

	queue := list.New()

	// we're pushing row, col
	// start with top-left
	queue.PushBack([2]int{0, 0})
	//visited[0][0] = true
	grid[0][0] = 1

	level := 1

	for queue.Len() > 0 {
		currentQueueLength := queue.Len()

		for i := 0; i < currentQueueLength; i++ {
			// pop from the queue
			cell := queue.Remove(queue.Front()).([2]int)

			row, column := cell[0], cell[1]

			if (row == rows-1) && (column == columns-1) {
				// bottom-right reached
				return level
			}

			for _, d := range directions {
				newRow := row + d[0]
				newColumn := column + d[1]

				if !cellExists(rows, columns, newRow, newColumn) ||
					//visited[newRow][newColumn] || // this will be covered by setting grid[i][j] to 1 for the visited rows
					grid[newRow][newColumn] == 1 { // we can travel only via 0 values
					continue
				}

				grid[newRow][newColumn] = 1 // we're blocking the visited rows in the grid itself
				queue.PushBack([2]int{newRow, newColumn})
			}
		}

		level++
	}

	// no path found with full BFS
	return -1
}

func shortestPathBinaryMatrixWithVisited(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])

	if grid[0][0] == 1 { // cannot start
		fmt.Printf("Invalid grid: top-left is 1. Returning -1. \n")
		return -1
	}

	if grid[rows-1][columns-1] == 1 { // cannot end
		fmt.Printf("Invalid grid: bottom-right is 1. Returning -1. \n")
		return -1
	}

	visited := createVisitedMatrix(rows, columns)

	queue := list.New()

	// we're pushing row, col
	// start with top-left
	queue.PushBack([2]int{0, 0})
	visited[0][0] = true

	level := 1

	for queue.Len() > 0 {
		currentQueueLength := queue.Len()

		for i := 0; i < currentQueueLength; i++ {
			// pop from the queue
			cell := queue.Remove(queue.Front()).([2]int)

			row, column := cell[0], cell[1]

			if (row == rows-1) && (column == columns-1) {
				// bottom-right reached
				return level
			}

			for _, d := range directions {
				newRow := row + d[0]
				newColumn := column + d[1]

				if !cellExists(rows, columns, newRow, newColumn) ||
					visited[newRow][newColumn] ||
					grid[newRow][newColumn] == 1 { // we can travel only via 0 values
					continue
				}

				visited[newRow][newColumn] = true
				queue.PushBack([2]int{newRow, newColumn})
			}
		}

		level++
	}

	// no path found with full BFS
	return -1
}

func createVisitedMatrix(rows int, columns int) [][]bool {
	visited := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	return visited
}

func printMatrix(mat [][]int) {
	rows := len(mat)
	columns := len(mat[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%v ", mat[i][j])
		}

		fmt.Println()
	}
}

func test(m [][]int, expectedResult int) { // nodes can be null
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: \n")
	printMatrix(m)

	result := shortestPathBinaryMatrix(m)

	fmt.Printf("Shortest path from top-left to bottom-right: %v \n", result)
	fmt.Printf("Expected shortest path: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}

	expected := 4

	test(m, expected)
}

func test2() {
	// throws OOM when using list as a queue
	// fixed: I forgot to set the new added node to visited = true
	m := [][]int{
		{0, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 1, 0},
		{0, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}

	expected := 25

	test(m, expected)
}

func main() {
	// 1091. Shortest Path in Binary Matrix
	test1()
	test2()
}
