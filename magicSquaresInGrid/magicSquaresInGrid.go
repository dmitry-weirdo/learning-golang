package main

import "fmt"

const SUM = 15 // in 1-9 magic square, all rows, columns and sums must have a sum of 15

func numMagicSquaresInside(grid [][]int) int {
	rows, columns := getRowsAndColumns(grid)

	count := 0

	for i := 0; i <= rows-3; i++ {
		for j := 0; j <= columns-3; j++ {
			// i, j is top-right of 3x3 square
			if !check1To9Values(grid, i, j) {
				continue
			}

			if !checkSums(grid, i, j) {
				continue
			}

			count++
		}
	}

	return count
}

func getRowsAndColumns(mat [][]int) (rows, columns int) {
	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func check1To9Values(m [][]int, top, left int) bool {
	freq := make([]int, 10) // we need 1-9 values

	for i := top; i < top+3; i++ {
		for j := left; j < left+3; j++ {
			// we iterate 9 cells
			v := m[i][j]

			if v < 1 || v > 9 {
				return false
			}

			freq[v]++

			if freq[v] > 1 { // repeat of 1 to 9 found -> fail
				return false
			}
		}
	}

	// no repeats and all values within [1; 9]
	return true
}

func checkSums(m [][]int, i, j int) bool {
	// rows check
	if (m[i][j]+m[i][j+1]+m[i][j+2] != SUM) ||
		(m[i+1][j]+m[i+1][j+1]+m[i+1][j+2] != SUM) ||
		(m[i+2][j]+m[i+2][j+1]+m[i+2][j+2] != SUM) {
		return false
	}

	// columns check
	if (m[i][j]+m[i+1][j]+m[i+2][j] != SUM) ||
		(m[i][j+1]+m[i+1][j+1]+m[i+2][j+1] != SUM) ||
		(m[i][j+2]+m[i+1][j+2]+m[i+2][j+2] != SUM) {
		return false
	}

	// diagonals
	if (m[i][j]+m[i+1][j+1]+m[i+2][j+2] != SUM) ||
		(m[i][j+2]+m[i+1][j+1]+m[i+2][j] != SUM) {
		return false
	}

	return true
}

func test(m [][]int, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: %v \n", m)

	result := numMagicSquaresInside(m)

	fmt.Printf("Count of 3x3 magic squares with digits from 1 to 9: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	m := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}

	expected := 1

	test(m, expected)
}

func test2() {
	m := [][]int{
		{8},
	}

	expected := 0

	test(m, expected)
}

func main() {
	// 840. Magic Squares In Grid
	test1()
	test2()
}
