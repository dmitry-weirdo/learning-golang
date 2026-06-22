package main

import "fmt"

var directions = [][]int{
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // bottom
	{0, -1}, // left
}

var nextIslandNumber byte

// var GRID [][]byte

func numIslands(grid [][]byte) int {
	numberOfIslands := 0
	currentIslandNumber := 2

	rows := len(grid)
	columns := len(grid[0])

	// to not pass a pointer in the stack
	//GRID = grid

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if isIsland(grid[i][j]) {
				//if grid[i][j] == '1' {
				numberOfIslands++

				//fmt.Printf("Number of islands increased to %v \n", numberOfIslands)

				// if we increase the bytes will overflow the byte on huge values!
				//nextIslandNumber = byte('0' + currentIslandNumber) // for debugging on small sets only
				nextIslandNumber = '2'

				currentIslandNumber++

				dfs(grid, rows, columns, i, j, nextIslandNumber)
				//dfs(grid, i, j, nextIslandNumber)
				//dfs(&grid, i, j)
				//dfs(i, j)

				//fmt.Println()
				//fmt.Printf("Matrix after DFS: \n")
				//printMatrix(grid)
			}
		}
	}

	return numberOfIslands
}

func dfs(grid [][]byte, rows int, columns int, i int, j int, nextIslandNumber byte) {
	//func dfs(i int, j int) {
	// trick to mark as visited - we'll mark it as '2', '3' etc.
	// Non-visited islands remain as '1'

	//fmt.Println()
	//fmt.Printf("[%v][%v]. Current island: %c Matrix: \n", i, j, nextIslandNumber)
	//printMatrix(grid)

	grid[i][j] = nextIslandNumber

	// iterate all 4 directions
	for _, d := range directions {
		row := i + d[0]
		column := j + d[1]

		exists := cellExists(rows, columns, row, column) // +1 stack

		/*		exists := (0 <= row) &&
				(row < len(GRID)) &&
				(0 <= column) &&
				(column < len(GRID[0]))
		*/

		if exists && isIsland(grid[row][column]) {
			dfs(grid, rows, columns, row, column, nextIslandNumber)
			//dfs(row, column) // shortest stack
			//dfs(grid, row, column, islandNumber)
			//dfs(grid, rows, columns, row, column, islandNumber)
		}
	}
}

func isIsland(b byte) bool {
	return b == '1'
}

func cellExists(rows, columns, row, column int) bool {
	return (0 <= row) &&
		(row < rows) &&
		(0 <= column) &&
		(column < columns)
}

func printMatrix(mat [][]byte) {
	rows := len(mat)
	columns := len(mat[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%c ", mat[i][j])
		}

		fmt.Println()
	}
}

func test(m [][]byte, expected int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Println("Initial islands:")
	printMatrix(m)

	numberOfIslands := numIslands(m)

	fmt.Println()
	fmt.Println("Islands after marking by numbers:")
	printMatrix(m)

	fmt.Println()
	fmt.Printf("Expected number of Islands: %v \n", expected)
	fmt.Printf("Number of Islands: %v \n", numberOfIslands)
}

func test1() {
	m := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	expected := 3

	test(m, expected)
}

func test2() {
	m := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	expected := 1

	test(m, expected)
}

func test3() {
	m := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}

	expected := 1

	test(m, expected)
}

func main() {
	test1()
	test2()
	test3()
}
