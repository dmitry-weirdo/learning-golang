package main

import (
	"demo/matrixCommon"
	"fmt"
)

func rotate(matrix [][]int) {
	// to rotate 90 degrees clockwise, we need to:
	// - transpose on main diagonal (i, j) <- -> (j, i)
	// - reverse every row

	// 1 2 3
	// 4 5 6
	// 7 8 9

	// transpose over main diagonal:
	// 1 4 7
	// 2 5 8
	// 3 6 9

	// transpose in place
	transposeSquareMatrixOverMainDiagonal(matrix)

	// reverse rows in place
	reverseMatrixRows(matrix)
}

func transposeSquareMatrixOverMainDiagonal(m [][]int) { // transposes in place
	n := len(m)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
}

func reverseMatrixRows(m [][]int) {
	rows, columns := getRowsAndColumns(m)

	for i := range rows {
		for j := range columns / 2 { // will ignore the center if columns is odd
			// let n = 3,
			// we need to swap 0 and 2
			// j = 0,
			// n - j - 1 = 2
			m[i][j], m[i][columns-j-1] = m[i][columns-j-1], m[i][j]
		}
	}
}

func getRowsAndColumns(mat [][]int) (rows, columns int) {
	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func test(m [][]int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: \n")
	matrixCommon.PrintIntMatrix(m)

	rotate(m)
	result := m // the function rotates in place

	fmt.Printf("Matrix rotated 90 degrees clockwise: \n")
	matrixCommon.PrintIntMatrix(result)

	fmt.Printf("Expected result: \n")
	matrixCommon.PrintIntMatrix(expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, resultRow := range result {
		expectedResultRow := expectedResult[i]

		// check that rows have the same length
		if len(resultRow) != len(expectedResultRow) {
			fmt.Printf("FAILURE: expectedResult[%v] length = %v, actualResult[%v] length = %v \n", i, len(expectedResultRow), i, len(resultRow))

			return
		}

		// same length -> check all row values
		for j, resultValue := range resultRow {
			expectedResultValue := expectedResultRow[j]

			if resultValue != expectedResultValue {
				fmt.Printf("FAILURE: expectedResult[%v][%v] = %v, actualResult[%v][%v]  = %v \n", i, j, expectedResultValue, i, j, resultValue)

				return
			}
		}
	}
}

func test1() {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	test(m, expected)
}

func test2() {
	m := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	expected := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}

	test(m, expected)
}

func main() {
	// 48. Rotate Image
	test1()
	test2()
}
