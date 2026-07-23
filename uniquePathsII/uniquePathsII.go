package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// todo: should we also handle if [0][0] is an obstacle?
	if obstacleGrid[m-1][n-1] == 1 {
		// cannot get to bottom-right
		return 0
	}

	// todo: we can actually live with just one row, but it is an over-optimization that will hurt the readability
	// (m - 1) row -> last 0, all previous are 1 (since we can only go right in the last row)
	previousRow := make([]int, n)

	// obstacle in last row sets all previous columns in this row to 0, we cannot go right after the obstacle anymore
	obstacleInLastRow := false

	for j := n - 1; j >= 0; j-- {
		if obstacleGrid[m-1][j] == 1 || obstacleInLastRow { // obstacle sets options to 0
			previousRow[j] = 0
			obstacleInLastRow = true
		} else {
			previousRow[j] = 1 // the very last is actually 0, but we don't care
		}
	}

	// (m - 2) row
	currentRow := make([]int, n) // will be 0 values by default

	// the cycle will work 0 times if there is just 1 row
	for i := m - 2; i >= 0; i-- { // actual current row number

		if obstacleGrid[i][n-1] == 1 || // obstacle in the current row, last column
			previousRow[n-1] == 0 { // obstacle blocked the previous row, last column
			currentRow[n-1] = 0
		} else {
			currentRow[n-1] = 1 // from the last column, we can only go down
		}

		for j := n - 2; j >= 0; j-- { // from pre-last column to the 0-th column
			if obstacleGrid[i][j] == 1 {
				currentRow[j] = 0 // obstacle in current row, current column
			} else {
				currentRow[j] = previousRow[j] + // go down
					currentRow[j+1] // go right
			}

		}

		previousRow = currentRow
	}

	// cannot use currentRow, for just 1 row (m = 1) we won't have the currentRow filled
	return previousRow[0]
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

func test(obstacleGrid [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("=================================")

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	fmt.Printf("M (rows): %v \n", m)
	fmt.Printf("N (columns): %v \n", n)

	fmt.Println("Obstacles matrix:")
	printMatrix(obstacleGrid)

	result := uniquePathsWithObstacles(obstacleGrid)

	fmt.Printf("Total distinct right-down paths from top-left to bottom-right in (%v x %v) matrix with obstacles: %v \n", m, n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]int{{0, 1, 0, 0}} // blocking obstacle in the last row

	expected := 0

	test(m, expected)
}

func test2() {
	m := [][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}

	expected := 0

	test(m, expected)
}

func test3() {
	m := [][]int{
		{0, 0},
		{0, 0},
		{0, 1}, // obstacle at the very end -> no ways
	}

	expected := 0

	test(m, expected)
}

func test4() {
	m := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	expected := 2

	test(m, expected)
}

func test5() {
	m := [][]int{
		{0, 1},
		{0, 0},
	}

	expected := 1

	test(m, expected)
}

func main() {
	// 63. Unique Paths II
	test1()
	test2()
	test3()
	test4()
	test5()
}
