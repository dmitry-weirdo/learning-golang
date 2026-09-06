package main

import (
	"demo/matrixCommon"
	"fmt"
)

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := getRowsAndColumns(mat)

	if m*n != r*c { // cannot reshape -> return the original matrix
		return mat
	}

	result := createIntMatrix(r, c)

	for i := range m * n {
		row1 := i / n
		col1 := i % n

		row2 := i / c
		col2 := i % c

		result[row2][col2] = mat[row1][col1]
	}

	return result
}

func getRowsAndColumns(mat [][]int) (rows, columns int) {
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

func test(m [][]int, r, c int, expectedResult [][]int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("Matrix: \n")
	matrixCommon.PrintIntMatrix(m)

	result := matrixReshape(m, r, c)

	fmt.Printf("Matrix reshaped to (%v x %v) size: \n", r, c)
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
		{1, 2},
		{3, 4},
	}

	r := 1
	c := 4

	expected := [][]int{
		{1, 2, 3, 4},
	}

	test(m, r, c, expected)
}

func test2() {
	m := [][]int{
		{1, 2},
		{3, 4},
	}

	r := 2
	c := 4

	expected := [][]int{ // cannot reshape to 2*4
		{1, 2},
		{3, 4},
	}

	test(m, r, c, expected)
}

func main() {
	// 566. Reshape the Matrix
	test1()
	test2()
}
