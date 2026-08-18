package main

import "fmt"

var directions = [][]int{
	{0, 1},  // right // we start from going right
	{1, 0},  // bottom
	{0, -1}, // left
	{-1, 0}, // top
}

// special value to mark visited values
const VISITED = -666

func spiralOrder(mat [][]int) []int {
	m, n := getRowsAndColumns(mat)

	i := 0
	j := -1 // start from -1-th column, so that the first move will set path[0]

	directionIndex := 0 // start with right
	direction := directions[directionIndex]

	pathLength := m * n
	path := make([]int, pathLength)
	pathIndex := 0

	// we mark visited cells as special value
	for pathIndex < pathLength { // until all cells passed
		for cellExists(m, n, i+direction[0], j+direction[1]) &&
			!isVisited(mat, i+direction[0], j+direction[1]) {
			i = i + direction[0]
			j = j + direction[1]

			path[pathIndex] = mat[i][j]
			pathIndex++

			// todo: if we don't want to change matrix data, we can set visited in a separate "visited" matrix. It will be O(m*n) space
			mat[i][j] = VISITED
		}

		directionIndex++
		directionIndex = directionIndex % len(directions)
		direction = directions[directionIndex]
	}

	return path
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

func isVisited(m [][]int, row, column int) bool {
	return m[row][column] == VISITED
}

func test(m [][]int, expectedResult []int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: %v \n", m)

	result := spiralOrder(m)

	fmt.Printf("Spiral part from top-right going right: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func test1() {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

	test(m, expected)
}

func test2() {
	m := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

	test(m, expected)
}

func test3() {
	m := [][]int{
		{1},
	}

	expected := []int{1}

	test(m, expected)
}

func test4() {
	// row-only
	m := [][]int{
		{1, 2, 3},
	}

	expected := []int{1, 2, 3}

	test(m, expected)
}

func test5() {
	// column-only
	m := [][]int{
		{1},
		{2},
		{3},
	}

	expected := []int{1, 2, 3}

	test(m, expected)
}

func test6() {
	m := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	expected := []int{1, 2, 4, 6, 5, 3}

	test(m, expected)
}

func main() {
	// 54. Spiral Matrix
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
