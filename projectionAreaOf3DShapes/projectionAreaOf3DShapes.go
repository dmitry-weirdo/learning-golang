package main

import "fmt"

func projectionArea(grid [][]int) int {
	// v = grid[i][j]
	// i == x
	// j == y
	// v = z (top of the tower based on (x, y) and starting from z = 0)

	// from top
	// When we look from top, we see 1 square filled for every non-0 value in the grid
	fromTop := 0

	for _, x := range grid {
		for _, y := range x {
			if y > 0 {
				fromTop++
			}
		}
	}

	fmt.Printf("Square of fromTop projection (seeing X and Y axes): %d \n", fromTop)

	// for other top projection, we need the max Z value for either every row or every column

	// when we look from x (from right), we see the max from each grid row at every row
	fromX := 0

	for _, x := range grid {
		rowMax := 0

		for _, y := range x {
			rowMax = max(rowMax, y)
		}

		fromX += rowMax
	}

	fmt.Printf("Square of fromX projection (seeing X and Z axes): %d \n", fromX)

	// when we look from y (from left), we see the max from each grid column at every column
	fromY := 0

	for column := 0; column < len(grid[0]); column++ {
		columnMax := 0

		for row := range grid {
			columnMax = max(columnMax, grid[row][column])
		}

		fromY += columnMax
	}

	fmt.Printf("Square of fromY projection (seeing Y and Z axes): %d \n", fromY)

	return fromTop + fromX + fromY
}

func test(m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix of Z heights for all [x][y] coordinates: %v \n", m)

	result := projectionArea(m) // todo: replace with your function

	fmt.Printf("Total area of all 3 projections: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	heights := [][]int{
		{1, 2},
		{3, 4},
	}

	expected := 17 // 4 + 6 + 7

	test(heights, expected)
}

func test2() {
	heights := [][]int{
		{2},
	}

	expected := 5 // 1 + 2 + 2

	test(heights, expected)
}

func test3() {
	heights := [][]int{
		{1, 0},
		{0, 2},
	}

	expected := 8 // 2 + 3 + 3

	test(heights, expected)
}

func main() {
	// 883. Projection Area of 3D Shapes
	test1()
	test2()
	test3()
}
