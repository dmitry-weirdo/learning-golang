package matrixCommon

import "fmt"

var directions = [][]int{
	{-1, 0}, // top
	{0, 1},  // right
	{1, 0},  // bottom
	{0, -1}, // left
}

func PrintIntMatrix(mat [][]int) {
	rows, columns := getRowsAndColumns(mat)

	for i := range rows {
		for j := range columns {
			fmt.Printf("%v ", mat[i][j])
		}

		fmt.Println()
	}
}

func PrintByteMatrix(mat [][]byte) { // prints elements as characters, NOT as numeric byte values
	rows, columns := getRowsAndColumnsOfByteMatrix(mat)

	for i := range rows {
		for j := range columns {
			fmt.Printf("%v ", mat[i][j])
		}

		fmt.Println()
	}
}

func PrintBoolMatrix(mat [][]bool) {
	rows, columns := getRowsAndColumnsOfBoolMatrix(mat)

	for i := range rows {
		for j := range columns {
			var v string

			if mat[i][j] {
				v = "T"
			} else {
				v = "F"
			}

			fmt.Printf("%v ", v)
		}

		fmt.Println()
	}
}

func getRowsAndColumns(mat [][]int) (rows, columns int) {
	if len(mat) <= 0 {
		return 0, 0
	}

	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func getRowsAndColumnsOfByteMatrix(mat [][]byte) (rows, columns int) {
	if len(mat) <= 0 {
		return 0, 0
	}

	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func getRowsAndColumnsOfBoolMatrix(mat [][]bool) (rows, columns int) {
	if len(mat) <= 0 {
		return 0, 0
	}

	return len(mat), len(mat[0]) // !!! we assume that all rows have the same length
}

func createIntMatrix(rows, columns int) [][]int {
	m := make([][]int, rows)

	for i := range rows {
		m[i] = make([]int, columns)
	}

	return m
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

func isTopReached(row int) bool {
	return row == 0
}

func isBottomReached(row int, rows int) bool {
	return row == (rows - 1)
}

func isLeftReached(column int) bool {
	return column == 0
}

func isRightReached(column int, columns int) bool {
	return column == (columns - 1)
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
