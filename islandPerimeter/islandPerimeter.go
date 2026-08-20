package main

import "fmt"

var directions = [][]int{
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // bottom
	{0, -1}, // left
}

func islandPerimeter(grid [][]int) int {
	// todo: we can execute dfs and find only island cells, but for 100 * 100 grid, it's enough to just iterate all the cells, it already can pass in 0 ms

	m, n := getRowsAndColumns(grid)

	perimeter := 0

	for i, row := range grid {
		for j, v := range row {
			if !isIsland(v) {
				continue
			}

			// check all directions whether they count as perimeter
			for _, d := range directions {
				if !cellExists(m, n, i+d[0], j+d[1]) || // border or grid -> nowhere to go
					!isIsland(grid[i+d[0]][j+d[1]]) { // water in this direction
					// border of the grid or no island in this direction -> this side counts as perimeter
					perimeter++
				}
			}
		}
	}

	return perimeter
}

func isIsland(x int) bool {
	return x == 1
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

func test(m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Grid (0 is water, 1 is island): %v \n", m) // todo: replace with your text if required

	result := islandPerimeter(m)

	fmt.Printf("Total island perimeter %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}

	expected := 16

	test(m, expected)
}

func test2() {
	m := [][]int{
		{1},
	}

	expected := 4

	test(m, expected)
}

func test3() {
	m := [][]int{
		{1, 0},
	}

	expected := 4

	test(m, expected)
}

func main() {
	// 463. Island Perimeter
	test1()
	test2()
	test3()
}
