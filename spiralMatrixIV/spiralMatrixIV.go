package main

import (
	. "demo/listsCommon" // not recommended, but ok for LeetCode -> to use TreeNode without prefix
)

var directions = [][]int{
	{0, 1},  // right // we start from going right
	{1, 0},  // bottom
	{0, -1}, // left
	{-1, 0}, // top
}

// special value to mark non-visited values
const NOT_VISITED = -666

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	// 0 value in the matrix means "non-visited"
	mat := createIntMatrixWithDefaultValues(m, n, NOT_VISITED)

	i := 0
	j := -1 // start from -1-th column, so that the first move will go to cell [0][0]

	dummyHead := &ListNode{-666, head}
	node := dummyHead

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

			if node != nil {
				node = node.Next
			}

			if node != nil {
				mat[i][j] = node.Val
			} else { // list exhausted -> fill the end of the matrix with -1
				mat[i][j] = -1
			}
		}

		directionIndex++
		directionIndex = directionIndex % len(directions)
		direction = directions[directionIndex]
	}

	return mat
}

func createIntMatrixWithDefaultValues(rows, columns int, defaultValue int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)

		for j := range columns { // !!! note that this is slow, will take O(m * n) additional operations :(
			m[i][j] = defaultValue
		}
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
	return m[row][column] != NOT_VISITED
}

func main() {
	// 2326. Spiral Matrix IV
	// Basically the same as "59. Spiral Matrix II",
	// we're just getting values from the given list instead of doing +1.
	// Non-square matrix makes no difference for my algorithm.

	// todo: add tests, will be a complex method of constructing a list from array, but generally it's still int -> matrix methoc
}
