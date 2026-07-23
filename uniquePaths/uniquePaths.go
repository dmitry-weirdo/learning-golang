package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	// m rows
	// n columns

	// todo: we can actually live with just one row, but it is an over-optimization that will hurt the readability
	// (m - 1) row -> last 0, all previous are 1 (since we can only go right in the last row)
	previousRow := make([]int, n)

	for i := range previousRow {
		previousRow[i] = 1 // the very last is actually 0, but we don't care
	}

	// (m - 2) row
	currentRow := make([]int, n) // will be 0 values by default

	// the cycle will work 0 times if there is just 1 row
	for i := m - 2; i >= 0; i-- { // actual current row number
		currentRow[n-1] = 1 // from the last column, we can only go down

		for j := n - 2; j >= 0; j-- { // from pre-last column to the 0-th column
			currentRow[j] = previousRow[j] + // go down
				currentRow[j+1] // go right
		}

		previousRow = currentRow
	}

	// cannot use currentRow, for just 1 row (m = 1) we won't have the currentRow filled
	return previousRow[0]
}

func test(m, n int, expectedResult int) {
	fmt.Println()
	fmt.Println("=================================")

	fmt.Printf("M (rows): %v \n", m)
	fmt.Printf("N (columns): %v \n", n)

	result := uniquePaths(m, n)

	fmt.Printf("Total distinct right-down paths from top-left to bottom-right in (%v x %v) matrix: %v \n", m, n, result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m, n := 3, 7
	expected := 28

	test(m, n, expected)
}

func test2() {
	m, n := 4, 4
	expected := 20

	test(m, n, expected)
}

func main() {
	// 62. Unique Paths
	test1()
	test2()
}
