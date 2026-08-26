package main

import "fmt"

var directions = [][]int{
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // bottom
	{0, -1}, // left
}

const EMPTY = 0
const FRESH = 1
const ROTTEN = 2

func orangesRotting(grid [][]int) int {
	rows, columns := getRowsAndColumns(grid)

	// level 0 is all rotten oranges
	queue := make([][]int, 0)

	totalFreshOranges := 0

	for i, row := range grid {
		for j, cell := range row {
			if isFreshOrange(cell) {
				totalFreshOranges++
			} else if isRottenOrange(cell) {
				queue = append(queue, []int{i, j})
			}
		}
	}

	fmt.Printf("Total fresh oranges: %v \n", totalFreshOranges)
	fmt.Printf("Total rotten oranges: %v \n", len(queue))
	fmt.Printf("Rotten oranges: %v \n", queue)

	if totalFreshOranges == 0 { // no fresh oranges -> already done
		return 0
	}

	if len(queue) == 0 { // no rotten, there are fresh -> impossible to rot all fresh oranges
		return -1
	}

	freshOrangesThatBecomeRotten := 0
	level := 0

	for len(queue) > 0 {
		currentLevelRottenOranges := len(queue)
		level++

		for range currentLevelRottenOranges {
			// poll from queue
			currentRotten := queue[0]
			queue = queue[1:]

			i := currentRotten[0]
			j := currentRotten[1]

			for _, d := range directions {
				neighborI := i + d[0]
				neighborJ := j + d[1]

				if cellExists(rows, columns, neighborI, neighborJ) &&
					isFreshOrange(grid[neighborI][neighborJ]) {

					freshOrangesThatBecomeRotten++
					grid[neighborI][neighborJ] = ROTTEN // rotten works as visited marker

					if freshOrangesThatBecomeRotten >= totalFreshOranges {
						// rot all oranges -> stop
						return level
					}

					// add new rotten to the new level
					queue = append(queue, []int{neighborI, neighborJ})
				}
			}
		}
	}

	//if freshOrangesThatBecomeRotten < totalFreshOranges {
	//	return -1
	//}

	// if we haven't returned -> not all fresh oranges hit with rot
	return -1
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

func isFreshOrange(v int) bool {
	return v == FRESH
}

func isRottenOrange(v int) bool {
	return v == ROTTEN
}

func test(m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix of oranges: %v \n", m)

	result := orangesRotting(m)

	fmt.Printf("Minutes for all oranges to rot: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	expected := 4

	test(m, expected)
}

func test2() {
	m := [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}

	expected := -1 // 2, 0 will never rot

	test(m, expected)
}

func test3() {
	m := [][]int{
		{0, 2},
	}

	expected := 0 // already totally rotten

	test(m, expected)
}

func test4() {
	m := [][]int{
		{0, 1},
	}

	expected := -1 // no rotten oranges, there are fresh

	test(m, expected)
}

func test5() {
	m := [][]int{
		{0, 0},
	}

	expected := 0 // no rotten oranges, no fresh -> target completed

	test(m, expected)
}

func main() {
	// 994. Rotting Oranges
	test1()
	test2()
	test3()
	test4()
	test5()
}
